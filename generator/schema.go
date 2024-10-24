package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Schema struct {
	Description string

	Ref  *SchemaComponent
	Type InternalSchemaType

	Nullable   string
	CustomType string
}

type Componenter interface {
	GetSchema(string) (*SchemaComponent, bool)
	AddSchema(string, Schema, Config) *SchemaComponent
}

func NewSchemaRef(sc *SchemaComponent) Schema {
	return Schema{
		Type: nil,
		Ref:  sc,
	}
}

func NewSchema(s specification.Ref[specification.Schema], components Componenter, cfg Config) (zero Schema, _ Imports, _ error) {
	var schemaRef *SchemaComponent
	if ref := s.Ref(); ref != nil {
		refOut, ok := components.GetSchema(ref.Name)
		if !ok {
			return zero, nil, fmt.Errorf("ref schema %q not found in schemas", ref.Name)
		}
		return NewSchemaRef(refOut), nil, nil
	}

	schema := s.Value()

	st, ims, err := newSchemaType(schema, components, cfg)
	if err != nil {
		return zero, nil, err
	}

	var customType string
	if specCustom, ok := schema.Value().Custom.Get(); ok {
		ct, is := NewCustomType(specCustom, st)
		customType = ct.Value
		ims = append(ims, is...)
	}

	var nullable string
	if schema.Nullable {
		nullable = "Nullable"
		if cfg.Nullable.Type != "" {
			nullable = cfg.Nullable.Type
			ims = append(ims, Import(cfg.Nullable.Import))
		}
	}

	return Schema{
		Description: s.Value().Description,

		Type: st,
		Ref:  schemaRef,

		Nullable:   nullable,
		CustomType: customType,
	}, ims, nil
}

func NewSchemaWithType(s InternalSchemaType) Schema {
	return Schema{
		Type: s,
	}
}

func (s Schema) CopyBase() Schema {
	return Schema{
		Description: s.Description,
		Ref:         s.Ref,
		Type:        s.Type,
		Nullable:    "",
		CustomType:  "",
	}
}

func (s Schema) BaseSchemaType() InternalSchemaType {
	return s.Base().Type
}

func (s Schema) Base() Schema {
	if s.Ref != nil {
		return s.Ref.Schema.Base()
	}
	return s
}

func (s Schema) RenderBaseFrom(prefix, from, suffix string) (string, error) {
	if s.CustomType != "" {
		from = from + "." + s.FuncTypeName() + "()"
	}
	if s.Nullable != "" {
		return NullableType{V: s.Type, TypeName: s.Nullable}.RenderBaseFrom(prefix, from, suffix)
	}
	return prefix + from + suffix, nil
}

func (s Schema) RenderToBaseType(to, from string) (string, error) {
	if s.Ref != nil {
		isStruct := s.Ref.Schema.Kind() == SchemaKindObject
		isArray := s.Ref.Schema.Kind() == SchemaKindArray
		if !s.Ref.Schema.IsCustom() && !isStruct && !isArray {
			from = from + "." + s.Ref.Schema.FuncTypeName() + "()"
		}
		return s.Ref.Schema.RenderToBaseType(to, from)
	}
	tp := s.Type
	if s.CustomType != "" {
		tp = CustomType{Value: s.CustomType, Type: s.Type}
	}
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderToBaseType(to, from)
}

func (s Schema) FuncTypeName() string {
	if s.Ref != nil {
		return s.Ref.Name
	}
	return s.Type.FuncTypeName()
}

func (s Schema) Kind() SchemaKind {
	if s.Ref != nil {
		return SchemaKindRef
	}
	return s.Type.Kind()
}

func (s Schema) RenderGoType() (string, error) {
	if s.Ref != nil {
		tp := s.Ref.Name
		// if s.Base().Nullable != "" {
		// 	tp = NullableType{TypeName: s.Base().Nullable}.GoType(tp)
		// }
		return tp, nil
	}
	tp := InternalSchemaType(s.Type)
	if s.CustomType != "" {
		tp = CustomType{Value: s.CustomType, Type: s.Type}
	}
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderGoType()
}

func (s Schema) RenderBaseGoType() (string, error) {
	if s.Ref != nil {
		tp := s.Ref.Name
		// if s.Base().Nullable != "" {
		// 	tp = NullableType{TypeName: s.Base().Nullable}.GoType(tp)
		// }
		return tp, nil
	}
	tp := InternalSchemaType(s.Type)
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderGoType()
}

// TODO: refactor to remove the method
func (s Schema) IsCustom() bool {
	return s.CustomType != ""
}

func (s Schema) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if s.Ref != nil {
		if s.Ref.Schema.IsCustom() {
			return s.Ref.Schema.ParseString(to, from, isNew, mkErr)
		}
		return ExecuteTemplate("Schema_Ref_ParseString", TData{
			"FuncName": s.Ref.Schema.FuncTypeName(),
			"Type":     s.Ref.Schema,
			"Name":     s.Ref.Name,

			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	}
	tp := InternalSchemaType(s.Type)
	if s.CustomType != "" {
		tp = CustomType{Value: s.CustomType, Type: s.Type}
	}
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.ParseString(to, from, isNew, mkErr)
}

func (s Schema) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if s.Ref != nil {
		if s.Ref.Schema.IsCustom() {
			return s.Ref.Schema.ParseStrings(to, from, isNew, mkErr)
		}
		return ExecuteTemplate("Schema_Ref_ParseStrings", TData{
			"FuncName": s.Ref.Schema.FuncTypeName(),
			"Type":     s.Ref.Schema,
			"Name":     s.Ref.Name,

			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	}
	switch st := s.Type.(type) {
	case SliceType:
		if st.Items.Ref == nil && !st.Items.IsCustom() {
			switch items := st.Items.Base().Type.(type) {
			case Primitive:
				switch items.PrimitiveIface.(type) {
				case StringType:
					return ExecuteTemplate("Schema_Assign", TData{
						"To":    to,
						"From":  from,
						"IsNew": isNew,
					})
				}
			}
		}
	}
	tp := s.Type
	if s.CustomType != "" {
		tp = CustomType{Value: s.CustomType, Type: s.Type}
	}
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.ParseStrings(to, from, isNew, mkErr)
}

func (s Schema) RenderFormat(from string) (string, error) {
	if s.Ref != nil {
		isStruct := s.Ref.Schema.Kind() == SchemaKindObject
		isArray := s.Ref.Schema.Kind() == SchemaKindArray
		if !s.Ref.Schema.IsCustom() && !isStruct && !isArray {
			from = from + "." + s.Ref.Schema.FuncTypeName() + "()"
		}
		return s.Ref.Schema.RenderFormat(from)
	}
	tp := s.Type
	if s.CustomType != "" {
		tp = CustomType{Value: s.CustomType, Type: s.Type}
	}
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderFormat(from)
}

func (s Schema) RenderConvertToBaseSchema(from string) (string, error) {
	if s.Ref != nil {
		if !s.Ref.Schema.IsCustom() {
			from = from + "." + s.Ref.Schema.FuncTypeName() + "()"
		}
		return s.Ref.Schema.RenderConvertToBaseSchema(from)
	}

	if s.CustomType != "" {
		from = from + "." + s.Type.FuncTypeName() + "()"
	}
	return from, nil
}

func (s Schema) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	if s.Ref != nil {
		if !s.Ref.Schema.IsCustom() {
			from = from + "." + s.Ref.Schema.FuncTypeName() + "()"
		}
		tp := SchemaType(s.Ref.Schema)
		return tp.RenderFormatStrings(to, from, isNew)
	}
	switch st := s.Type.(type) {
	case SliceType:
		if st.Items.Ref == nil && !st.Items.IsCustom() {
			switch items := st.Items.Base().Type.(type) {
			case Primitive:
				switch items.PrimitiveIface.(type) {
				case StringType:
					return ExecuteTemplate("Schema_Assign", TData{
						"To":    to,
						"From":  from,
						"IsNew": isNew,
					})
				}
			}
		}
	}
	tp := InternalSchemaType(s.Type)
	if s.CustomType != "" {
		tp = CustomType{Value: s.CustomType, Type: s.Type}
	}
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderFormatStrings(to, from, isNew)
}

func (s Schema) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if s.Ref != nil {
		// if !s.Ref.Schema.IsCustom() {
		// 	from = from + "." + s.Ref.Schema.FuncTypeName() + "()"
		// }
		return s.Ref.Schema.RenderUnmarshalJSON(to, from, isNew, mkErr)
	}

	var tp InternalSchemaType = s.Type
	if s.CustomType != "" {
		// from = from + "." + s.Type.FuncTypeName() + "()"
		tp = CustomType{Value: s.CustomType, Type: tp}
	}
	if s.Nullable != "" {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderUnmarshalJSON(to, from, isNew, mkErr)
}

type InternalSchemaType interface {
	GoTypeRender
	Parser
	Formatter

	RenderToBaseType(to, from string) (string, error)

	FuncTypeName() string
	Kind() SchemaKind

	RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
}

type SchemaType interface {
	GoTypeRender
	Parser
	Formatter

	RenderBaseFrom(prefix, from, suffix string) (string, error)
	RenderToBaseType(to, from string) (string, error)

	FuncTypeName() string
	Kind() SchemaKind

	RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
}

type SchemaKind string

const (
	SchemaKindPrimitive SchemaKind = "primitive"
	SchemaKindArray     SchemaKind = "array"
	SchemaKindObject    SchemaKind = "object"
	SchemaKindRef       SchemaKind = "ref"
)

func newSchemaType(spec *specification.Schema, components Componenter, cfg Config) (InternalSchemaType, Imports, error) {
	if len(spec.AllOf) > 0 {
		var s StructureType
		var imports Imports
		for i, a := range spec.AllOf {
			if ref := a.Ref(); ref != nil {
				s.Fields = append(s.Fields, StructureField{GoTypeFn: StringRender(ref.Name).Render, Name: ref.Name, Embedded: true})
			} else {
				st, ims, err := NewStructureType(a.Value(), components, cfg)
				if err != nil {
					return nil, nil, fmt.Errorf("allOf: %d-th element: new structure type: %w", i, err)
				}
				imports = append(imports, ims...)
				s.Fields = append(s.Fields, st.Fields...)
			}
		}
		return s, imports, nil
	}

	// https://datatracker.ietf.org/doc/html/draft-wright-json-schema-00#section-4
	switch spec.Type {
	case "boolean":
		return NewPrimitive(BoolType{}), nil, nil
	case "object":
		r, ims, err := NewStructureType(spec, components, cfg)
		if err != nil {
			return nil, nil, fmt.Errorf("'object' type: %w", err)
		}
		return r, ims, nil
	case "array":
		itemType, is, err := NewSchema(spec.Value().Items, NamedComponenter{Componenter: components, Name: "Items"}, cfg)
		if err != nil {
			return nil, nil, fmt.Errorf("items schema: %w", err)
		}
		return SliceType{Items: itemType}, is, nil
	case "number":
		switch spec.Format {
		case "float":
			return NewPrimitive(FloatType{BitSize: 32}), nil, nil
		case "double", "":
			return NewPrimitive(FloatType{BitSize: 64}), nil, nil
		default:
			return nil, nil, fmt.Errorf("unsupported 'number' format %q", spec.Format)
		}
	case "string":
		switch spec.Format {
		case "":
			return NewPrimitive(StringType{}), nil, nil
		case "byte": // base64 encoded characters
		case "binary": // any sequence of octets
		case "date": // full-date = 4DIGIT "-" 01-12 "-" 01-31
		case "date-time": // full-date "T" 00-23 ":" 00-59 ":" 00-60 "Z" / ("+" / "-") 00-23 ":" 00-60
			format := "time.RFC3339Nano"
			if v, ok := spec.Extentions[specification.ExtTagGoTimeFormat]; ok {
				format = v
			}
			dt, ims := NewDateTime(format)
			return NewPrimitive(dt), ims, nil
		case "password":
		default:
			return nil, nil, fmt.Errorf("unsupported 'string' format %q", spec.Format)
		}
		return NewPrimitive(StringType{}), nil, nil
	case "integer":
		switch spec.Format {
		case "int32":
			return NewPrimitive(IntType{BitSize: 32}), nil, nil
		case "int64":
			return NewPrimitive(IntType{BitSize: 64}), nil, nil
		default:
			return NewPrimitive(IntType{}), nil, nil
		}
	}

	return nil, nil, fmt.Errorf("unknown schema type: %q", spec.Type)
}
