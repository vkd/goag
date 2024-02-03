package generator

import (
	"encoding/json"
	"fmt"

	"github.com/vkd/goag/specification"
)

func NewParameterSchema(spec specification.Schema) (Formatter, error) {
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
