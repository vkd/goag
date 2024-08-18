package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Schema struct {
	Type   SchemaType
	Ref    *SchemaComponent
	Custom Maybe[string]
}

func NewSchema(s specification.Ref[specification.Schema], components Components) (zero Schema, _ Imports, _ error) {
	var schemaRef *SchemaComponent
	if ref := s.Ref(); ref != nil {
		refOut, ok := components.Schemas.Get(ref.V)
		if !ok {
			return zero, nil, fmt.Errorf("ref schema %q not found in schemas", ref.Name)
		}
		schemaRef = refOut
	}

	st, ims, err := NewSchemaType(s, components)
	if err != nil {
		return zero, nil, fmt.Errorf("new schema type: %w", err)
	}

	return Schema{
		Type: st,
		Ref:  schemaRef,
	}, ims, nil
}

func (s Schema) Base() SchemaType { return s.Type }

func (s Schema) RenderType() (string, error) {
	return s.Type.Render()
}

func (s Schema) TypeRender() Render { return RenderFunc(s.RenderType) }

func (s Schema) ParseString(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return s.Type.ParseString(to, from, isNew, mkErr)
}

func (s Schema) ParseStrings(to, from string, isNew bool, mkErr ErrorRender) (string, error) {
	return s.Type.ParseStrings(to, from, isNew, mkErr)
}

func (s Schema) IsMultivalue() bool { return s.Type.IsMultivalue() }

func (s Schema) RenderFormat(from string) (string, error) {
	return s.Type.RenderFormat(from)
}

func (s Schema) RenderFormatStrings(to, from string, isNew bool) (string, error) {
	return s.Type.RenderFormatStrings(to, from, isNew)
}

type SchemaType interface {
	Base() SchemaType

	Render
	Parser
	Formatter

	Kind() SchemaKind
}

type SchemaKind string

const (
	SchemaKindPrimitive SchemaKind = "primitive"
	SchemaKindArray     SchemaKind = "array"
	SchemaKindObject    SchemaKind = "object"
	SchemaKindCustom    SchemaKind = "custom"
	SchemaKindRef       SchemaKind = "ref"
	SchemaKindMap       SchemaKind = "map"
)

func NewSchemaType(s specification.Ref[specification.Schema], components Components) (SchemaType, Imports, error) {
	if ref := s.Ref(); ref != nil {
		refOut, ok := components.Schemas.Get(ref.V)
		if !ok {
			return nil, nil, fmt.Errorf("ref schema %q not found in schemas", ref.Name)
		}
		return NewRefSchemaType(ref.Name, refOut), nil, nil
	}

	spec := s.Value()

	out, ims, err := newSchemaType(spec, components)
	if err != nil {
		return nil, nil, err
	}

	if specCustom, ok := spec.Custom.Get(); ok {
		ct, is := NewCustomType(specCustom, out)

		out = ct
		ims = append(ims, is...)
		// return ct, append(is, ims...), nil
	}

	return out, ims, nil
}

func newSchemaType(spec *specification.Schema, components Components) (SchemaType, Imports, error) {
	if len(spec.AllOf) > 0 {
		var s StructureType
		var imports Imports
		for i, a := range spec.AllOf {
			if ref := a.Ref(); ref != nil {
				s.Fields = append(s.Fields, StructureField{Type: StringRender(ref.Name), Embedded: true})
			} else {
				st, ims, err := NewStructureType(a.Value(), components)
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
		if specAdditionalProperties, ok := spec.AdditionalProperties.Get(); ok && len(spec.Properties) == 0 {
			additional, ims, err := NewSchemaType(specAdditionalProperties, components)
			if err != nil {
				return nil, nil, fmt.Errorf("additional properties: %w", err)
			}
			return NewMapType(NewPrimitive(StringType{}), additional), ims, nil
		}
		r, ims, err := NewStructureType(spec, components)
		if err != nil {
			return nil, nil, fmt.Errorf("'object' type: %w", err)
		}
		return r, ims, nil
	case "array":
		itemType, is, err := NewSchemaType(spec.Value().Items, components)
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
