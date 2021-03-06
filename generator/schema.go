package generator

import (
	"fmt"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type SchemaRender interface {
	Render
	Parser(from, to string, _ FuncNewError) Render
}

func NewSchemaRef(spec *openapi3.SchemaRef) SchemaRender {
	if spec.Ref != "" {
		return NewRef(spec.Ref)
	}
	return NewSchema(spec.Value)
}

func NewSchema(spec *openapi3.Schema) SchemaRender {
	if spec.AllOf != nil {
		var fields []GoStructField
		for _, s := range spec.AllOf {
			if s.Ref != "" {
				fields = append(fields, GoStructField{Type: NewRef(s.Ref)})
				continue
			}
			sfs := NewGoStructFields(NewSchemas(s.Value.Properties))
			fields = append(fields, sfs...)
		}
		return GoStruct{Fields: fields}
	}

	switch spec.Type {
	case "object":
		sfs := NewGoStructFields(NewSchemas(spec.Properties))
		goStruct := GoStruct{Fields: sfs}

		if spec.AdditionalProperties != nil {
			addSchema := NewSchemaRef(spec.AdditionalProperties)
			if len(sfs) == 0 {
				return GoMap{Key: StringType, Value: addSchema}
			}
			goStruct.Fields = append(goStruct.Fields, GoStructField{
				Name: "AdditionalProperties",
				Type: GoMap{Key: StringType, Value: addSchema},
				Tags: []GoFieldTag{{Key: "json", Value: "-"}},
			})
		}
		return goStruct
	case "array":
		sr := NewSchemaRef(spec.Items)
		return GoSlice{Items: sr}
	case "string":
		return StringType
	case "integer":
		switch spec.Format {
		case "int32":
			return Int32
		case "int64":
			return Int64
		default:
			return Int
		}
	case "number":
		switch spec.Format {
		case "float":
			return Float32
		case "double":
			return Float64
		}
		return Float64
	}

	panic(fmt.Errorf("unknown schema type: %q", spec.Type))
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

func (r Ref) Parser(from, to string, mkErr FuncNewError) Render {
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
