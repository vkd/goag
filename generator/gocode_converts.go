package generator

import (
	"text/template"
)

type Assign struct {
	From, To string
}

var tmAssign = template.Must(template.New("Assign").Parse(`
{{- .To}} = {{.From}}`))

func (c Assign) String() (string, error) { return String(tmAssign, c) }

type FuncNewError func(s string) string

type ConvertToInt struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToInt = template.Must(template.New("ConvertToInt").Parse(`vInt, err := strconv.Atoi({{.Variable}})
if err != nil {
	return zero, {{call .NewError "parse int"}}
}
{{.Field}} = vInt`))

func (c ConvertToInt) String() (string, error) { return String(tmConvertToInt, c) }

type ConvertToInt32 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToInt32 = template.Must(template.New("ConvertToInt32").Parse(`vInt, err := strconv.ParseInt({{.Variable}}, 10, 32)
if err != nil {
	return zero, {{call .NewError "parse int32"}}
}
{{.Field}} = int32(vInt)`))

func (c ConvertToInt32) String() (string, error) { return String(tmConvertToInt32, c) }

type ConvertToInt64 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToInt64 = template.Must(template.New("ConvertToInt64").Parse(`vInt, err := strconv.ParseInt({{.Variable}}, 10, 64)
if err != nil {
	return zero, {{call .NewError "parse int64"}}
}
{{.Field}} = vInt`))

func (c ConvertToInt64) String() (string, error) { return String(tmConvertToInt64, c) }

type ConvertToFloat32 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToFloat32 = template.Must(template.New("ConvertToFloat32").Parse(`vf, err := strconv.ParseFloat({{.Variable}}, 32)
if err != nil {
	return zero, {{call .NewError "parse float32"}}
}
{{.Field}} = float32(vf)`))

func (c ConvertToFloat32) String() (string, error) { return String(tmConvertToFloat32, c) }

type ConvertToFloat64 struct {
	Variable, Field string
	NewError        FuncNewError
}

var tmConvertToFloat64 = template.Must(template.New("ConvertToFloat64").Parse(`vf, err := strconv.ParseFloat({{.Variable}}, 64)
if err != nil {
	return zero, {{call .NewError "parse float64"}}
}
{{.Field}} = vf`))

func (c ConvertToFloat64) String() (string, error) { return String(tmConvertToFloat64, c) }

type StructParser struct {
	From, To string
	MkErr    FuncNewError
}

var tmStructParser = template.Must(template.New("StructParser").Parse(`err := {{.To}}.UnmarshalText({{.From}})
if err != nil {
	return zero, {{call .MkErr "parse struct"}}
}`))

func (s StructParser) String() (string, error) { return String(tmStructParser, s) }
