package generator

import (
	"fmt"
	"text/template"
)

type Combine []Render

var tmCombine = template.Must(template.New("Combine").Parse(`
{{- range $i, $c := . }}{{ if $i }}
{{ end }}
{{- $c.String }}
{{- end }}`))

func (c Combine) String() (string, error) { return String(tmCombine, c) }

type Assign struct {
	From, To string
}

func (c Assign) String() (string, error) { return fmt.Sprintf(`%s = %s`, c.To, c.From), nil }

type AssignNew struct {
	From, To string
}

func (c AssignNew) String() (string, error) { return fmt.Sprintf(`%s := %s`, c.To, c.From), nil }

type FuncNewError func(s string) string

type ConvertToInt struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToInt = template.Must(template.New("ConvertToInt").Parse(`vInt, err := strconv.Atoi({{.Variable}})
if err != nil {
	return zero, {{call .NewError "fmt.Errorf(\"parse int: %w\", err)"}}
}
{{.Field}} := vInt`))

func (c ConvertToInt) String() (string, error) { return String(tmConvertToInt, c) }

type ConvertToInt32 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToInt32 = template.Must(template.New("ConvertToInt32").Parse(`vInt, err := strconv.ParseInt({{.Variable}}, 10, 32)
if err != nil {
	return zero, {{call .NewError "fmt.Errorf(\"parse int32: %w\", err)"}}
}
{{.Field}} := int32(vInt)`))

func (c ConvertToInt32) String() (string, error) { return String(tmConvertToInt32, c) }

type ConvertToInt64 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToInt64 = template.Must(template.New("ConvertToInt64").Parse(`vInt, err := strconv.ParseInt({{.Variable}}, 10, 64)
if err != nil {
	return zero, {{call .NewError "fmt.Errorf(\"parse int64: %w\", err)"}}
}
{{.Field}} := vInt`))

func (c ConvertToInt64) String() (string, error) { return String(tmConvertToInt64, c) }

type ConvertToFloat32 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToFloat32 = template.Must(template.New("ConvertToFloat32").Parse(`vf, err := strconv.ParseFloat({{.Variable}}, 32)
if err != nil {
	return zero, {{call .NewError "fmt.Errorf(\"parse float32: %w\", err)"}}
}
{{.Field}} := float32(vf)`))

func (c ConvertToFloat32) String() (string, error) { return String(tmConvertToFloat32, c) }

type ConvertToFloat64 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToFloat64 = template.Must(template.New("ConvertToFloat64").Parse(`vf, err := strconv.ParseFloat({{.Variable}}, 64)
if err != nil {
	return zero, {{call .NewError "fmt.Errorf(\"parse float64: %w\", err)"}}
}
{{.Field}} := vf`))

func (c ConvertToFloat64) String() (string, error) { return String(tmConvertToFloat64, c) }

type StructParser struct {
	From, To string
	NewError FuncNewError
}

var tmStructParser = template.Must(template.New("StructParser").Parse(`err := {{.To}}.UnmarshalText({{.From}})
if err != nil {
	return zero, {{call .NewError "fmt.Errorf(\"parse struct: %w\", err)"}}
}`))

func (s StructParser) String() (string, error) { return String(tmStructParser, s) }
