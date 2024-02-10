package generator

import (
	"encoding/json"
	"fmt"

	"github.com/vkd/goag/specification"
)

func NewSchema(spec specification.Schema) (Render, error) {
	if spec.Ref != "" {
		return NewRef(spec.Ref), nil
	}

	if v, ok := spec.Schema.ExtensionProps.Extensions[ExtTagGoType]; ok {
		if raw, ok := v.(json.RawMessage); ok {
			s := string(raw)
			if len(s) > 2 {
				s = s[1 : len(s)-1]
			}
			return NewCustomType(s), nil
		} else {
			return nil, fmt.Errorf("unexpected type for 'Extension Properties' %T - expected 'json.RawMessage", v)
		}
	}

	if len(spec.AllOf) > 0 {
		var s StructureType
		for i, a := range spec.AllOf {
			if a.Ref != "" {
				s.Fields = append(s.Fields, StructureField{Type: NewRef(a.Ref)})
			} else {
				for _, p := range a.Properties {
					sf, err := NewStructureField(p)
					if err != nil {
						return nil, fmt.Errorf("allOf: %d-th element: new structure field: %w", i, err)
					}
					s.Fields = append(s.Fields, sf)
				}
			}
		}
		return s, nil
	}

	switch spec.Type {
	case "object":
		r, err := NewStructureType(spec)
		if err != nil {
			return nil, fmt.Errorf("'object' type: %w", err)
		}
		return r, nil
	case "array":
		itemType, err := NewSchema(*spec.Items)
		if err != nil {
			return nil, fmt.Errorf("items schema: %w", err)
		}
		return SliceType{Items: itemType}, nil
	case "string":
		return StringType{}, nil
	case "integer":
		switch spec.Schema.Format {
		case "int32":
			return IntType{32}, nil
		case "int64":
			return IntType{64}, nil
		default:
			return IntType{}, nil
		}
	case "number":
		switch spec.Schema.Format {
		case "float":
			return FloatType{BitSize: 32}, nil
		case "double", "":
		default:
			return nil, fmt.Errorf("unsupported 'number' format %q", spec.Schema.Format)
		}
		return FloatType{BitSize: 64}, nil
	case "boolean":
		return BoolType{}, nil
	}

	return nil, fmt.Errorf("unknown schema type: %q", spec.Schema.Type)
}

func NewParameterSchema(spec specification.Schema) (interface {
	Render
	Formatter
}, error) {
	if v, ok := spec.Schema.ExtensionProps.Extensions[ExtTagGoType]; ok {
		if raw, ok := v.(json.RawMessage); ok {
			s := string(raw)
			if len(s) > 2 {
				s = s[1 : len(s)-1]
			}
			return NewCustomType(s), nil
		} else {
			return nil, fmt.Errorf("unexpected type for 'Extension Properties' %T - expected 'json.RawMessage", v)
		}
	}

	switch spec.Schema.Type {
	case "array":
		itemType, err := NewParameterSchema(*spec.Items)
		if err != nil {
			return nil, fmt.Errorf("items schema: %w", err)
		}
		return SliceType{Items: itemType}, nil
	case "string":
		return StringType{}, nil
	case "integer":
		switch spec.Schema.Format {
		case "int32":
			return IntType{32}, nil
		case "int64":
			return IntType{64}, nil
		default:
			return IntType{}, nil
		}
	case "number":
		switch spec.Schema.Format {
		case "float":
			return FloatType{BitSize: 32}, nil
		case "double", "":
		default:
			return nil, fmt.Errorf("unsupported 'number' format %q", spec.Schema.Format)
		}
		return FloatType{BitSize: 64}, nil
	case "boolean":
		return BoolType{}, nil
	}

	return nil, fmt.Errorf("unknown schema type: %q", spec.Schema.Type)
}
