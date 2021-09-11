package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaRender interface {
	Render
	Parser(from, to string, _ FuncNewError) (Render, error)
}

func NewSchemaRef(spec *openapi3.SchemaRef) (SchemaRender, error) {
	if spec.Ref != "" {
		return NewRef(spec.Ref), nil
	}
	return NewSchema(spec.Value)
}

func NewSchema(spec *openapi3.Schema) (SchemaRender, error) {
	if spec.AllOf != nil {
		var fields []GoStructField
		for _, s := range spec.AllOf {
			if s.Ref != "" {
				fields = append(fields, GoStructField{Type: NewRef(s.Ref)})
				continue
			}
			if s.Value.Type != "object" {
				return nil, fmt.Errorf("allOf works only for objects: unknown type %q", s.Value.Type)
			}
			sfs, err := NewGoStructFields(NewSchemas(s.Value.Properties))
			if err != nil {
				return nil, fmt.Errorf("new schemas of allOf: %w", err)
			}
			fields = append(fields, sfs...)
		}
		return GoStruct{Fields: fields}, nil
	}

	switch spec.Type {
	case "object":
		sfs, err := NewGoStructFields(NewSchemas(spec.Properties))
		if err != nil {
			return nil, fmt.Errorf("new schemas of 'object' type: %w", err)
		}
		if spec.AdditionalProperties != nil {
			if len(sfs) > 0 {
				panic("not implemented")
			}
			addSchema, err := NewSchemaRef(spec.AdditionalProperties)
			if err != nil {
				return nil, fmt.Errorf("new schema ref for value type of additional properties: %w", err)
			}
			return GoMap{Key: StringType, Value: addSchema}, nil
		}
		return GoStruct{Fields: sfs}, nil
	case "array":
		sr, err := NewSchemaRef(spec.Items)
		if err != nil {
			return nil, fmt.Errorf("new schema ref: %w", err)
		}
		return GoSlice{Items: sr}, nil
	case "string":
		return StringType, nil
	case "integer":
		switch spec.Format {
		case "int32":
			return Int32, nil
		case "int64":
			return Int64, nil
		default:
			return Int, nil
		}
	case "number":
		switch spec.Format {
		case "float":
			return Float32, nil
		case "double":
			return Float64, nil
		}
		return Float64, nil
	}
	return nil, fmt.Errorf("unknown schema type: %q", spec.Type)
}

type Ref string

var _ Render = Ref("")

func NewRef(ref string) Ref {
	ref = ref[strings.LastIndex(ref, "/"):]
	ref = strings.TrimPrefix(ref, "/")
	return Ref(ref)
}

func (r Ref) String() (string, error) {
	return string(r), nil
}

func (r Ref) Parser(from, to string, mkErr FuncNewError) (Render, error) {
	panic("not implemented")
}

type SchemasItem struct {
	Name   string
	Schema *openapi3.SchemaRef
}

type SchemasItems []SchemasItem

func NewSchemas(ss openapi3.Schemas) SchemasItems {
	fields := make(SchemasItems, 0, len(ss))
	for fieldname, s := range ss {
		fields = append(fields, SchemasItem{fieldname, s})
	}
	sort.Slice(fields, func(i, j int) bool { return fields[i].Name < fields[j].Name })

	return fields
}
