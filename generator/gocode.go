package generator

import (
	"fmt"
	"strings"
	"text/template"
)

type GoFile struct {
	DoNotEdit   bool
	PackageName string

	Imports []GoFileImport

	Renders []Render
}

//nolint:gci
var tmGoFile = template.Must(template.New("GoFile").Parse(`
{{- if .DoNotEdit}}// Code is generated by goag. DO NOT EDIT!
{{end -}}
package {{.PackageName}}

{{- if .Imports}}

import (
	{{- range $i, $imp := .Imports}}
	{{if $imp.Alias}}{{$imp.Alias}} {{end}}"{{$imp.Package}}"
	{{- end}}
)
{{- end}}

{{- range $_, $body := .Renders}}

{{$body.String}}
{{- end}}
`))

func (g GoFile) String() (string, error) { return String(tmGoFile, g) }

type GoFileImport struct {
	Alias   string
	Package string
}

type GoVarDef struct {
	Name  string
	Type  Render
	Value Render
}

var tmGoVarDef = template.Must(template.New("GoVarDef").Parse(`var {{.Name}} {{.Type.String}} = {{.Value.String}}`))

func (g GoVarDef) String() (string, error) { return String(tmGoVarDef, g) }

type GoConstDef struct {
	Name  string
	Type  Render
	Value Render
}

var tmGoConstDef = template.Must(template.New("GoConstDef").Parse(`const {{.Name}} {{.Type.String}} = {{.Value.String}}`))

func (g GoConstDef) String() (string, error) { return String(tmGoConstDef, g) }

type GoTypeDef struct {
	Comment string
	Name    string
	Type    Render

	Methods []Render
}

func NewGoTypeDef(i SchemasItem) GoTypeDef {
	sr := NewSchemaRef(i.Schema)
	return GoTypeDef{
		Name:    i.Name,
		Comment: i.Schema.Value.Description,
		Type:    sr,
	}
}

func NewGoTypeDefs(si SchemasItems) []GoTypeDef {
	out := make([]GoTypeDef, 0, len(si))
	for _, i := range si {
		td := NewGoTypeDef(i)
		out = append(out, td)
	}
	return out
}

var tmGoTypeDef = template.Must(template.New("GoTypeDef").Parse(`
{{- if .Comment}}// {{.Name}} - {{.Comment}}
{{end -}}
type {{.Name}} {{.Type.String -}}

{{- range $_, $m := .Methods}}

{{$m.String}}
{{- end}}
`))

func (g GoTypeDef) String() (string, error) { return String(tmGoTypeDef, g) }

func (g GoTypeDef) GoTypeRef() GoType { return GoType(g.Name) }

type GoStruct struct {
	Fields []GoStructField
}

var tmGoStruct = template.Must(template.New("GoStruct").Parse(`
{{- if .Fields -}}
struct {
	{{- range $_, $f := .Fields}}
	{{$f.String -}}
	{{end}}
}
{{- else -}}
struct{}
{{- end}}`))

func (g GoStruct) String() (string, error) { return String(tmGoStruct, g) }

func (g GoStruct) Parser(from, to string, mkErr FuncNewError) Render {
	return StructParser{from, to, mkErr}
}

type GoStructField struct {
	Name    string
	Comment string
	Type    Render
	Tags    []GoFieldTag
}

func NewGoStructField(i SchemasItem) (zero GoStructField) {
	sr := NewSchemaRef(i.Schema)
	sf := GoStructField{
		Name: PublicFieldName(i.Name),
		Type: sr,
	}
	if sf.Name != i.Name {
		sf.Tags = append(sf.Tags, GoFieldTag{"json", i.Name})
	}
	return sf
}

func NewGoStructFields(si SchemasItems) []GoStructField {
	out := make([]GoStructField, 0, len(si))
	for _, i := range si {
		sf := NewGoStructField(i)
		out = append(out, sf)
	}
	return out
}

var tmGoStructField = template.Must(template.New("GoStructField").Parse(`
{{- if .Comment}}// {{.Name}} - {{.Comment}}
{{end -}}
{{.Name}} {{.Type.String}}
{{- if .Tags}} ` + "`" + `
	{{- range $i, $t := .Tags}}
		{{- if $i}} {{end}}
		{{- $t.String}}
	{{- end}}` + "`" + `
{{- end -}}
`))

func (g GoStructField) String() (string, error) { return String(tmGoStructField, g) }

func (g GoStructField) GetTag(key string) (GoFieldTag, bool) {
	for _, t := range g.Tags {
		if t.Key == key {
			return t, true
		}
	}
	return GoFieldTag{}, false
}

type GoFieldTag struct {
	Key   string
	Value string
}

func (g GoFieldTag) String() (string, error) { return fmt.Sprintf("%s:%q", g.Key, g.Value), nil }

type GoType string

func (g GoType) String() (string, error) { return string(g), nil }

const (
	StringType GoType = "string"

	Int   GoType = "int"
	Int32 GoType = "int32"
	Int64 GoType = "int64"
	// Struct GoType = "struct"
	// Slice GoType = "slice"

	Float32 GoType = "float32"
	Float64 GoType = "float64"
)

func (g GoType) Parser(from, to string, mkErr FuncNewError) Render {
	switch g {
	case StringType:
		return AssignNew{from, to}
	case Int:
		return ConvertToInt{from, to, mkErr}
	case Int32:
		return ConvertToInt32{from, to, mkErr}
	case Int64:
		return ConvertToInt64{from, to, mkErr}
	case Float32:
		return ConvertToFloat32{from, to, mkErr}
	case Float64:
		return ConvertToFloat64{from, to, mkErr}
	}
	panic(fmt.Errorf("unsupported GoType: %q", g))
}

type GoValue string

func (g GoValue) String() (string, error) { return string(g), nil }

type StringValue GoValue

func (s StringValue) String() (string, error) {
	if strings.Contains(string(s), "\n") {
		s = StringValue(strings.ReplaceAll(string(s), "`", "`+\"`\"+`"))
		s = "`" + s + "`"
	} else {
		s = StringValue(strings.ReplaceAll(string(s), `"`, `\"`))
		s = `"` + s + `"`
	}
	return GoValue(s).String()
}

type GoSlice struct {
	Items SchemaRender
}

var tmGoSlice = template.Must(template.New("GoSlice").Parse(`[]{{ .Items.String }}`))

func (s GoSlice) String() (string, error) { return String(tmGoSlice, s) }

func (s GoSlice) Parser(from, to string, mkErr FuncNewError) Render {
	switch t := s.Items.(type) {
	case GoType:
		switch t {
		case StringType:
			return Assign{from + "[0]", to}
		}
	}
	panic("not implemented")
}

func (GoSlice) Optionable() {}

func (s GoSlice) StringsParser(from, to string, mkErr FuncNewError) Render {
	switch t := s.Items.(type) {
	case GoType:
		switch t {
		case StringType:
			return Assign{from, to}
		}
	}
	return ConvertStrings{s.Items, from, to, mkErr}
}

type GoMap struct {
	Key, Value SchemaRender
}

var tmGoMap = template.Must(template.New("GoMap").Parse(`map[{{ .Key.String }}]{{ .Value.String }}`))

func (s GoMap) String() (string, error) { return String(tmGoMap, s) }

func (s GoMap) Parser(from, to string, mkErr FuncNewError) Render {
	panic("not implemented")
}

func (GoMap) Optionable() {}

type ConvertStrings struct {
	ItemType SchemaRender
	From, To string
	MkErr    FuncNewError
}

var tmConvertStrings = template.Must(template.New("ConvertStrings").Parse(`
{{- .To}} = make([]{{.ItemType.String}}, len({{.From}}))
for i := range {{.From}} {
	{{.ItemRender (print .From "[i]") (print .To "[i]")}}
}`))

func (c ConvertStrings) ItemRender(from, toOrig string) (string, error) {
	to := "v1"
	r := c.ItemType.Parser(from, to, c.MkErr)
	r = Combine{r, Assign{to, toOrig}}
	return r.String()
}

func (c ConvertStrings) String() (string, error) { return String(tmConvertStrings, c) }

type GoFunction struct {
	Receiver GoFunctionArg
	Name     string
	Args     []GoFunctionArg
	Returns  []GoFunctionArg
	Body     Render
}

var tmGoFunction = template.Must(template.New("GoFunction").Parse(`func{{if .Receiver}} ({{.Receiver.String}}){{end}}{{if .Name}} {{.Name}}{{end}}({{range $i, $a := .Args}}{{if $i}}, {{end}}{{$a.Var}} {{$a.Type.String}}{{end}}) ({{range $i, $r := .Returns}}{{if $i}}, {{end}}{{$r.Var}} {{$r.Type.String}}{{end}}) {
	{{.Body.String}}
}`))

func (g GoFunction) String() (string, error) {
	return String(tmGoFunction, g)
}

type GoFunctionArg struct {
	Var  string
	Type Render
}

func (g GoFunctionArg) String() (string, error) {
	tp, err := g.Type.String()
	if err != nil {
		return "", fmt.Errorf("render type: %w", err)
	}

	if g.Var == "" && len(tp) > 0 {
		g.Var = strings.ToLower(tp[:1])
	}

	out := g.Var
	if len(out) > 0 {
		out += " "
	}
	out += tp

	return out, nil
}

func MarshalJSONFunc(orig GoType, body GoStruct) GoFunction {
	bodyFunc := `m := make(map[string]interface{})
for k, v := range b.AdditionalProperties {
	m[k] = v
}
`
	for _, f := range body.Fields {
		if t, ok := f.GetTag("json"); ok && t.Value != "-" {
			bodyFunc += `m["` + t.Value + `"] = b.` + f.Name + "\n"
		}
	}
	bodyFunc += `return json.Marshal(m)` + "\n"

	return GoFunction{
		Receiver: GoFunctionArg{Var: "b", Type: orig},
		Name:     "MarshalJSON",
		Returns:  []GoFunctionArg{{Type: GoType("[]byte")}, {Type: GoType("error")}},
		Body:     GoValue(bodyFunc),
	}
}
