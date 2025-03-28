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
	CustomType Maybe[CustomType]
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

	var customType Maybe[CustomType]
	if specCustom, ok := schema.Value().Custom.Get(); ok {
		ct, is := NewCustomType(specCustom, st)
		customType = Just(ct)
		ims = append(ims, is...)
	}

	var nullable string
	if schema.Nullable {
		nullable = "Nullable"
		if cfg.Nullable.Type != "" {
			nullable = cfg.Nullable.Type
			if cfg.Nullable.Import != "" {
				ims = append(ims, NewImport(cfg.Nullable.Import, ""))
			}
		}
	}

	out := Schema{
		Description: s.Value().Description,

		Type: st,
		Ref:  schemaRef,

		Nullable:   nullable,
		CustomType: customType,
	}
	return out, ims, nil
}

func NewSchemaWithType(s InternalSchemaType) Schema {
	return Schema{
		Type: s,
	}
}

func (s Schema) IsNullable() bool {
	return s.Base().Nullable != ""
}

func (s Schema) Custom() *CustomType {
	if ct, ok := s.CustomType.Get(); ok {
		return &ct
	}
	return nil
}

func (s Schema) NullableType() NullableType {
	return NullableType{V: s.Type, TypeName: s.Nullable}
}

func (s Schema) GetParser() (Parser, bool) {
	switch schema := s.Type.(type) {
	case RawBytesType:
		return nil, false
	case StructureType:
		return nil, false
	case SliceType:
		switch schema.Items.Base().Type.(type) {
		case StructureType:
		default:
			return s, true
		}
		return nil, false
	case OneOfStructure:
		return nil, false
	}
	return s, true
}

func (s Schema) GetFormatter() (Formatter, bool) {
	switch schema := s.Type.(type) {
	case RawBytesType:
		return nil, false
	case StructureType:
		return nil, false
	case SliceType:
		switch schema.Items.Base().Type.(type) {
		case StructureType:
		default:
			return s, true
		}
		return nil, false
	case OneOfStructure:
		return nil, false
	}
	return s, true
}

func (s Schema) Base() Schema {
	if s.Ref != nil {
		return s.Ref.Schema.Base()
	}
	return s
}

func (s Schema) RenderToBaseType(to, from string) (string, error) {
	if s.Ref != nil {
		return ExecuteTemplate("Schema_RenderToBaseType_Ref", TData{
			"To":   to,
			"From": from,

			"Schema": s,
		})
	}
	return ExecuteTemplate("Schema_RenderToBaseType", TData{
		"To":   to,
		"From": from,

		"Schema": s,
	})
}

func (s Schema) FuncTypeName() string {
	if s.Ref != nil {
		return s.Ref.Name
	}
	return s.Type.FuncTypeName()
}

func (s Schema) Kind() SchemaKind {
	return s.Base().Type.Kind()
}

func (s Schema) RenderTypeDefinition() (string, error) {
	if s.Ref != nil {
		return s.Ref.Name, nil
	}
	return s.Type.RenderGoType()
}

func (s Schema) RenderFieldType() (string, error) {
	if s.Ref != nil {
		tp := s.Ref.Name
		if ct, ok := s.Ref.Schema.CustomType.Get(); ok {
			tp = ct.Value
		}
		if s.IsNullable() {
			tp = NullableType{TypeName: s.Base().Nullable}.GoType(tp)
		}
		return tp, nil
	}
	if ct, ok := s.CustomType.Get(); ok {
		if s.IsNullable() {
			return NullableType{TypeName: s.Nullable}.GoType(ct.Value), nil
		}
		return ct.Value, nil
	}
	tp := s.Type
	if s.IsNullable() {
		tp = s.NullableType()
	}
	return tp.RenderGoType()
}

func (s Schema) RenderGoType() (string, error) {
	if s.Ref != nil {
		tp := s.Ref.Name
		if s.IsNullable() {
			tp = NullableType{TypeName: s.Base().Nullable}.GoType(tp)
		}
		return tp, nil
	}

	tp := s.Type
	if ct, ok := s.CustomType.Get(); ok {
		tp = ct
	}
	if s.IsNullable() {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderGoType()
}

func (s Schema) RenderBaseGoType() (string, error) {
	if s.Ref != nil {
		tp := s.Ref.Name
		// if s.Base().IsNullable() {
		// 	tp = NullableType{TypeName: s.Base().Nullable}.GoType(tp)
		// }
		return tp, nil
	}
	tp := s.Type
	if s.IsNullable() {
		tp = s.NullableType()
	}
	return tp.RenderGoType()
}

// TODO: refactor to remove the method
func (s Schema) IsCustom() bool {
	return s.CustomType.IsSet || (s.Ref != nil && s.Ref.Schema.IsCustom())
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
	tp := s.Type
	if ct, ok := s.CustomType.Get(); ok {
		tp = ct
	}
	if s.IsNullable() {
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
	if ct, ok := s.CustomType.Get(); ok {
		tp = ct
	}
	if s.IsNullable() {
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
	if ct, ok := s.CustomType.Get(); ok {
		tp = ct
	}
	if s.IsNullable() {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderFormat(from)
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
	if ct, ok := s.CustomType.Get(); ok {
		tp = ct
	}
	if s.IsNullable() {
		tp = NullableType{V: tp, TypeName: s.Nullable}
	}
	return tp.RenderFormatStrings(to, from, isNew)
}

func (s Schema) RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if s.Ref != nil {
		return ExecuteTemplate("Schema_Ref_RenderUnmarshalJSON", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"Ref":                   s.Ref,
			"Name":                  s.Ref.Name,
			"Schema":                s.Ref.Schema,
			"IsCustomUnmarshalJSON": s.Ref.Schema.Kind() == SchemaKindObject || s.Ref.Schema.Kind() == SchemaKindArray,
			"RenderFieldTypeFn":     s.RenderFieldType,
			"RenderGoTypeFn":        s.RenderBaseGoType,
		})
	}

	if custom, ok := s.CustomType.Get(); ok {
		return ExecuteTemplate("Schema_Custom_RenderUnmarshalJSON", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"Schema":            s,
			"CustomType":        custom,
			"RenderFieldTypeFn": s.RenderFieldType,
			"RenderGoTypeFn":    s.RenderBaseGoType,
		})
	}

	return ExecuteTemplate("Schema_RenderUnmarshalJSON", TData{
		"To":    to,
		"From":  from,
		"IsNew": isNew,
		"MkErr": mkErr,

		"Schema":            s,
		"RenderFieldTypeFn": s.RenderFieldType,
		"RenderGoTypeFn":    s.RenderBaseGoType,
	})
}

func (s Schema) RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if s.Ref != nil {
		if ct, ok := s.Ref.Schema.CustomType.Get(); ok {
			return ExecuteTemplate("Schema_Ref_CustomType_RenderMarshalJSON", TData{
				"To":    to,
				"From":  from,
				"IsNew": isNew,
				"MkErr": mkErr,

				"Ref":    s.Ref,
				"Type":   ct,
				"Schema": s,
			})
		}
		return ExecuteTemplate("Schema_Ref_RenderMarshalJSON", TData{
			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,

			"Ref": s.Ref,
		})
	}

	tp := s.Type
	if s.IsNullable() {
		tp = s.NullableType()
	}
	return tp.RenderMarshalJSON(to, from, isNew, mkErr)
}

type InternalSchemaType interface {
	GoTypeRender
	Parser
	Formatter

	RenderToBaseType(to, from string) (string, error)

	FuncTypeName() string
	Kind() SchemaKind

	RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
	RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
}

type SchemaType interface {
	GoTypeRender
	Parser
	Formatter

	RenderToBaseType(to, from string) (string, error)
	RenderFieldType() (string, error)

	FuncTypeName() string
	Kind() SchemaKind

	RenderUnmarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
	RenderMarshalJSON(to, from string, isNew bool, mkErr ErrorRender) (string, error)
}

type SchemaKind string

const (
	SchemaKindPrimitive SchemaKind = "primitive"
	SchemaKindArray     SchemaKind = "array"
	SchemaKindObject    SchemaKind = "object"
	SchemaKindRawBytes  SchemaKind = "raw_bytes"
)

func newSchemaType(spec *specification.Schema, components Componenter, cfg Config) (InternalSchemaType, Imports, error) {
	if len(spec.AllOf) > 0 {
		var s StructureType
		var imports Imports
		for i, a := range spec.AllOf {
			schema, ims, err := NewSchema(a, components, cfg)
			if err != nil {
				return nil, nil, fmt.Errorf("allOf: %d-th element: new schema: %w", i, err)
			}
			imports = append(imports, ims...)

			if schema.Ref != nil {
				s.Fields = append(s.Fields, StructureField{
					Name:               schema.Ref.Name,
					Type:               schema,
					Schema:             schema,
					GoTypeFn:           schema.RenderGoType,
					FieldTypeFn:        schema.RenderFieldType,
					Embedded:           true,
					RenderToBaseTypeFn: schema.RenderToBaseType,
				})
			} else if schema.Kind() == SchemaKindObject {
				st, ims, err := NewStructureType(a.Value(), components, cfg)
				if err != nil {
					return nil, nil, fmt.Errorf("allOf: %d-th element: new structure type: %w", i, err)
				}
				imports = append(imports, ims...)
				s.Fields = append(s.Fields, st.Fields...)
			} else {
				return nil, nil, fmt.Errorf("allOf: %d-th element: wrong schema type: only type 'object' is supported: object type: %q", i, schema.Kind())
			}
		}
		return s, imports, nil
	}
	if len(spec.OneOf) > 0 {
		out, imports, err := NewOneOfStructure(spec.OneOf, spec.Discriminator, components, cfg)
		return out, imports, err
	}

	// https://datatracker.ietf.org/doc/html/draft-wright-json-schema-00#section-4
	switch spec.Type {
	case "boolean":
		return NewBoolType(), nil, nil
	case "object":
		r, ims, err := NewStructureType(spec, components, cfg)
		if err != nil {
			return nil, nil, fmt.Errorf("'object' type: %w", err)
		}
		return r, ims, nil
	case "": // any
		return RawBytesType{}, nil, nil
	case "array":
		itemType, is, err := NewSchema(spec.Value().Items, NamedComponenter{Componenter: components, Name: "Items"}, cfg)
		if err != nil {
			return nil, nil, fmt.Errorf("items schema: %w", err)
		}
		if itemType.Ref == nil && itemType.Kind() == SchemaKindObject {
			sc := components.AddSchema("Item", itemType, cfg)
			itemType.Ref = sc
		}
		return SliceType{Items: itemType}, is, nil
	case "number":
		switch spec.Format {
		case "float":
			return NewFloatType(32), nil, nil
		case "double", "":
			return NewFloatType(64), nil, nil
		default:
			return nil, nil, fmt.Errorf("unsupported 'number' format %q", spec.Format)
		}
	case "string":
		switch spec.Format {
		case "":
			return NewStringType(), nil, nil
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
		return NewStringType(), nil, nil
	case "integer":
		switch spec.Format {
		case "int32":
			return NewIntTypeXX(32), nil, nil
		case "int64":
			return NewIntTypeXX(64), nil, nil
		default:
			return NewIntType(), nil, nil
		}
	}

	return nil, nil, fmt.Errorf("unknown schema type: %q", spec.Type)
}
