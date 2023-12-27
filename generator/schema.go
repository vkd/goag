package generator

import (
	"encoding/json"
	"fmt"
	"text/template"

	"github.com/vkd/goag/generator-v0/source"
	"github.com/vkd/goag/specification"
)

const ExtTagGoType = "x-goag-go-type"

type Schema interface {
	// FormatQuery()
	TemplateToString(Templater) Templater
	// FormatAssignTemplater(from, to Templater, isNew bool) Templater
	// ExecuteParse() (string, error)
	// Format() (Templater, error)
}

func NewSchema(spec specification.Schema) Schema {
	// if spec.AllOf != nil {
	// 	var fields []GoStructField
	// 	for _, s := range spec.AllOf {
	// 		if s.Ref != "" {
	// 			fields = append(fields, GoStructField{Type: NewRef(s.Ref)})
	// 			continue
	// 		}
	// 		sfs := NewGoStructFields(NewSchemas(s.Value.Properties))
	// 		fields = append(fields, sfs...)
	// 	}
	// 	return GoStruct{Fields: fields}
	// }

	if v, ok := spec.Schema.ExtensionProps.Extensions[ExtTagGoType]; ok {
		if raw, ok := v.(json.RawMessage); ok {
			s := string(raw)
			if len(s) > 2 {
				s = s[1 : len(s)-1]
			}
			return StringerType{}
		}
	}

	switch spec.Schema.Type {
	// case "object":
	// sfs := NewGoStructFields(NewSchemas(spec.Properties))
	// goStruct := GoStruct{Fields: sfs}

	// if spec.AdditionalProperties != nil {
	// 	addSchema := NewSchemaRef(spec.AdditionalProperties)
	// 	if len(sfs) == 0 {
	// 		return GoMap{Key: StringType, Value: addSchema}
	// 	}
	// 	goStruct.Fields = append(goStruct.Fields, GoStructField{
	// 		Name: "AdditionalProperties",
	// 		Type: GoMap{Key: StringType, Value: addSchema},
	// 		Tags: []GoFieldTag{{Key: "json", Value: "-"}},
	// 	})
	// }
	// return goStruct
	case "array":
		// sr := NewSchemaRef(spec.Items)
		return SliceType{Items: NewSchema(*spec.Items)}
	case "string":
		return StringType{}
	case "integer":
		switch spec.Schema.Format {
		case "int32":
			return Int32Type{}
			// return source.GoIntXXVar(32)
		case "int64":
			return Int64Type{}
			// return source.GoIntXXVar(64)
		default:
			return IntType{}
		}
		// case "number":
		// 	switch spec.Format {
		// 	case "float":
		// 		return Float32
		// 	case "double":
		// 		return Float64
		// 	}
		// 	return Float64
		// case "boolean":
		// 	return BooleanType
	}

	panic(fmt.Errorf("unknown schema type: %q", spec.Schema.Type))
}

type StringConst string

func (s StringConst) Execute() (string, error) { return "\"" + string(s) + "\"", nil }
func (s StringConst) String() (string, error)  { return s.Execute() }

// ---

// type Formatter struct{}

// type StringVar struct{}

// type StringType struct{}

// func (_ StringType) FormatTemplate(t Templater) Templater {
// 	return t
// }

// func (_ StringType) ToString(varName string) (string, error) { return varName, nil }

// func Int64ToString(t Templater) Templater {
// 	return Int64Type{t}
// }

// type Int64Type struct {
// 	t Templater
// 	a Int64ToString
// }

// func (t Int64Type) String() (string, error) {
// 	return templates.ExecuteTemplate("Int64ToString", t.t)
// }

// func (_ Int64Type) FormatTemplate(t Templater) Templater {
// 	return templates.ExecuteTemplate("Int64ToString", t)
// 	// return RawTemplate("strconv.FormatInt(" + varName + ", 10)")
// }

// type IntType struct{}

// //	func (_ IntType) ExecuteArgs(args ...any) (string, error) {
// //		log.Printf("args: %+v", args)
// //		return "IntType{}", nil
// //	}
// func (_ IntType) ToInt64(varName string) Templater {
// 	return InitTemplate("int->int64", "int64({{ exec . }})").Execute(Int64Type{})
// 	// Int64Type{}.ToString(varName)
// 	// return RawTemplate("str->int(" + varName + ")")
// }

type CustomType string

func (c CustomType) String() (string, error) {
	return string(c), nil
}

func (c CustomType) Parser(from, to string, mkErr source.ErrorWrapper) source.Render {
	return CustomTypeParser{string(c), from, to, mkErr}
}

func (c CustomType) Format(s string) source.Templater {
	panic("not implemented")
}

type CustomTypeParser struct {
	Type  string
	From  string
	To    string
	Error source.ErrorWrapper
}

var tmCustomTypeParser = template.Must(template.New("CustomTypeParser").Parse(`
var v {{ .Type }}
err := v.UnmarshalText({{.From}})
if err != nil {
	return zero, {{.Error.Wrap (print "unmarshal text")}}
}`))

func (c CustomTypeParser) String() (string, error) { return source.String(tmCustomTypeParser, c) }
