package generator

import (
	"fmt"
	"strings"

	"github.com/vkd/goag/specification"
)

type Schema struct {
	Description string

	Type   SchemaType
	Ref    *SchemaComponent
	Custom Maybe[string]
}

type Componenter interface {
	GetSchema(string) (*SchemaComponent, bool)
	AddSchema(string, Schema, Config) *SchemaComponent
}

func NewSchemaRef(sc *SchemaComponent) Schema {
	return Schema{
		Type:   nil,
		Ref:    sc,
		Custom: Nothing[string](),
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

	st, ims, err := NewSchemaType(s.Value(), components, cfg)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema type: %w", err)
	}

	var custom Maybe[string]
	if specCustom, ok := s.Value().Custom.Get(); ok {
		var customImport, customType string = "", specCustom
		slIdx := strings.LastIndex(specCustom, "/")
		if slIdx >= 0 {
			customImport = specCustom[:slIdx]
			customType = specCustom[slIdx+1:]

			dotIdx := strings.LastIndex(specCustom, ".")
			if dotIdx >= 0 {
				customImport = specCustom[:dotIdx]
			}
		}

		custom = Just(customType)
		ims = append(ims, NewImportsS(customImport)...)
	}

	return Schema{
		Description: s.Value().Description,

		Type:   st,
		Ref:    schemaRef,
		Custom: custom,
	}, ims, nil
}

func NewSchemaWithType(s SchemaType) Schema {
	return Schema{
		Type: s,
	}
}

func (s Schema) BaseSchemaType() SchemaType {
	if s.Ref != nil {
		return s.Ref.Schema.BaseSchemaType()
	}
	return s.Type
}

func (s Schema) Base() Schema {
	if s.Ref != nil {
		return s.Ref.Schema
	}
	return s
}

func (s Schema) FuncTypeName() string {
	if s.Ref != nil {
		return s.Ref.Name
	}
	if custom, ok := s.Custom.Get(); ok {
		return strings.ReplaceAll(custom, ".", "")
	}
	return s.Type.FuncTypeName()
}

func (s Schema) Kind() SchemaKind {
	if s.Ref != nil {
		return SchemaKindRef
	}
	return s.Type.Kind()
}

func (s Schema) Render() (string, error) {
	if s.Ref != nil {
		return s.Ref.Name, nil
	}
	if custom, ok := s.Custom.Get(); ok {
		return custom, nil
	}
	return s.Type.Render()
}

func (s Schema) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if s.Ref != nil {
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
	if custom, ok := s.Custom.Get(); ok {
		return ExecuteTemplate("Schema_Custom_ParseString", TData{
			"Base":         s.Type,
			"Type":         StringRender(custom),
			"FuncTypeName": s.Type.FuncTypeName(),

			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	}
	return s.Type.ParseString(to, from, isNew, mkErr)
}

func (s Schema) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	if s.Ref != nil {
		return ExecuteTemplate("Schema_Ref_ParseStrings", TData{
			"IsCustom": s.Ref.Schema.Custom.IsSet,
			"FuncName": s.Ref.Schema.FuncTypeName(),
			"Type":     s.Ref.Schema,
			"Name":     s.Ref.Name,

			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	}
	if custom, ok := s.Custom.Get(); ok {
		return ExecuteTemplate("Schema_Custom_ParseStrings", TData{
			"Base":         s.Type,
			"Type":         StringRender(custom),
			"FuncTypeName": s.Type.FuncTypeName(),

			"To":    to,
			"From":  from,
			"IsNew": isNew,
			"MkErr": mkErr,
		})
	}
	switch st := s.Type.(type) {
	case SliceType:
		if st.Items.Ref == nil && !st.Items.Custom.IsSet {
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
	return s.Type.ParseStrings(to, from, isNew, mkErr)
}

func (s Schema) IsMultivalue() bool { return s.Type.IsMultivalue() }

func (s Schema) RenderFormat(from string) (string, error) {
	if s.Ref != nil {
		if !s.Ref.Schema.Custom.IsSet {
			from = from + "." + s.Ref.Schema.FuncTypeName() + "()"
		}
		return s.Ref.Schema.RenderFormat(from)
	}
	if _, ok := s.Custom.Get(); ok {
		return s.Type.RenderFormat(from + "." + s.Type.FuncTypeName() + "()")
	}
	return s.Type.RenderFormat(from)
}

func (s Schema) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	if s.Ref != nil {
		if !s.Ref.Schema.Custom.IsSet {
			from = from + "." + s.Ref.Schema.FuncTypeName() + "()"
		}
		return s.Ref.Schema.RenderFormatStrings(to, from, isNew)
	}
	if _, ok := s.Custom.Get(); ok {
		return ExecuteTemplate("Schema_Custom_RenderFormatStrings", TData{
			"Base":         s.Type,
			"IsMultivalue": s.Type.IsMultivalue(),

			"To":    to,
			"From":  from + "." + s.Type.FuncTypeName() + "()",
			"IsNew": isNew,
		})
	}
	switch st := s.Type.(type) {
	case SliceType:
		if st.Items.Ref == nil && !st.Items.Custom.IsSet {
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
	return s.Type.RenderFormatStrings(to, from, isNew)
}

type SchemaType interface {
	Render
	Parser
	Formatter

	FuncTypeName() string
	Kind() SchemaKind
}

type SchemaKind string

const (
	SchemaKindPrimitive SchemaKind = "primitive"
	SchemaKindArray     SchemaKind = "array"
	SchemaKindObject    SchemaKind = "object"
	// SchemaKindCustom    SchemaKind = "custom"
	SchemaKindRef SchemaKind = "ref"
)

func NewSchemaType(s *specification.Schema, components Componenter, cfg Config) (SchemaType, Imports, error) {
	out, ims, err := newSchemaType(s, components, cfg)
	if err != nil {
		return nil, nil, err
	}

	return out, ims, nil
}

func newSchemaType(spec *specification.Schema, components Componenter, cfg Config) (SchemaType, Imports, error) {
	if len(spec.AllOf) > 0 {
		var s StructureType
		var imports Imports
		for i, a := range spec.AllOf {
			if ref := a.Ref(); ref != nil {
				s.Fields = append(s.Fields, StructureField{Type: StringRender(ref.Name), Embedded: true})
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
		itemType, is, err := NewSchema(spec.Value().Items, components, cfg)
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
