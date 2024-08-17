package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Schema struct {
	Spec specification.Ref[specification.Schema]
	Type SchemaType
}

type SchemaType interface {
	Base() SchemaType

	Render
	Parser
	Formatter
}

func NewSchema(s specification.Ref[specification.Schema], components Components) (SchemaType, Imports, error) {
	if ref := s.Ref(); ref != nil {
		refOut, ok := components.Schemas.Get(ref.V)
		if !ok {
			return nil, nil, fmt.Errorf("ref schema %q not found in schemas", ref.Name)
		}
		return NewRefSchemaType(ref.Name, refOut), nil, nil
	}

	spec := s.Value()

	out, ims, err := newSchema(spec, components)
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

func newSchema(spec *specification.Schema, components Components) (SchemaType, Imports, error) {
	if len(spec.AllOf) > 0 {
		var s StructureType
		var imports Imports
		for i, a := range spec.AllOf {
			if ref := a.Ref(); ref != nil {
				s.Fields = append(s.Fields, StructureField{Type: StringRender(ref.Name), Embedded: true})
			} else {
				for _, p := range a.Value().Properties {
					sf, ims, err := NewStructureField(p, components)
					if err != nil {
						return nil, nil, fmt.Errorf("allOf: %d-th element: new structure field: %w", i, err)
					}
					s.Fields = append(s.Fields, sf)
					imports = append(imports, ims...)
				}
			}
		}
		return s, imports, nil
	}

	switch spec.Type {
	case "object":
		if specAdditionalProperties, ok := spec.AdditionalProperties.Get(); ok && len(spec.Properties) == 0 {
			additional, ims, err := NewSchema(specAdditionalProperties, components)
			if err != nil {
				return nil, nil, fmt.Errorf("additional properties: %w", err)
			}
			return NewMapType(StringType{}, additional), ims, nil
		}
		r, ims, err := NewStructureType(spec, components)
		if err != nil {
			return nil, nil, fmt.Errorf("'object' type: %w", err)
		}
		return r, ims, nil
	case "array":
		itemType, is, err := NewSchema(spec.Value().Items, components)
		if err != nil {
			return nil, nil, fmt.Errorf("items schema: %w", err)
		}
		return SliceType{Items: itemType}, is, nil
	case "string":
		return StringType{}, nil, nil
	case "integer":
		switch spec.Format {
		case "int32":
			return IntType{BitSize: 32}, nil, nil
		case "int64":
			return IntType{BitSize: 64}, nil, nil
		default:
			return IntType{}, nil, nil
		}
	case "number":
		switch spec.Format {
		case "float":
			return FloatType{BitSize: 32}, nil, nil
		case "double", "":
		default:
			return nil, nil, fmt.Errorf("unsupported 'number' format %q", spec.Format)
		}
		return FloatType{BitSize: 64}, nil, nil
	case "boolean":
		return BoolType{}, nil, nil
	}

	return nil, nil, fmt.Errorf("unknown schema type: %q", spec.Type)
}
