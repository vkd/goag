package generator

import (
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
	From Render
	To   string
}

var tmAssign = template.Must(template.New("Assign").Parse(`{{.To}} = {{.From.String}}`))

func (c Assign) String() (string, error) { return String(tmAssign, c) }

type AssignNew struct {
	From Render
	To   string
}

var tmAssignNew = template.Must(template.New("AssignNew").Parse(`{{.To}} := {{.From.String}}`))

func (c AssignNew) String() (string, error) { return String(tmAssignNew, c) }

type FuncNewError func(s string) string

type StructParser struct {
	From, To string
	NewError FuncNewError
}

var tmStructParser = template.Must(template.New("StructParser").Parse(`err := {{.To}}.UnmarshalText({{.From}})
if err != nil {
	return zero, {{call .NewError "fmt.Errorf(\"parse struct: %w\", err)"}}
}`,
))

func (s StructParser) String() (string, error) { return String(tmStructParser, s) }

// ----

type TypeConversion struct {
	Type string
	From string
}

var tmTypeConversion = template.Must(template.New("TypeConversion").Parse(
	`{{.Type}}({{.From}})`,
))

func (c TypeConversion) String() (string, error) { return String(tmTypeConversion, c) }

// int, int32, int64

func ConvertToInt(from, to string, newError FuncNewError) Render {
	return Combine{
		ConvertToIntXX{0, from, "vInt", newError},
		AssignNew{TypeConversion{"int", "vInt"}, to},
	}
}

func ConvertToInt32(from, to string, newError FuncNewError) Render {
	return Combine{
		ConvertToIntXX{32, from, "vInt", newError},
		AssignNew{TypeConversion{"int32", "vInt"}, to},
	}
}

func ConvertToInt64(from, to string, newError FuncNewError) Render {
	return ConvertToIntXX{64, from, to, newError}
}

type ConvertToIntXX struct {
	BitSize     int
	From, ToNew string
	NewError    FuncNewError
}

var tmConvertToIntXX = template.Must(template.New("ConvertToIntXX").Parse(`
{{- $bitSize := ""}}
{{- if .BitSize}}{{$bitSize = (printf "%d" .BitSize)}}{{else}}{{$bitSize = ""}}{{end -}}
{{.ToNew}}, err := strconv.ParseInt({{.From}}, 10, {{.BitSize}})
if err != nil {
	return zero, {{call .NewError (print "fmt.Errorf(\"parse int" $bitSize ": %w\", err)") }}
}`))

func (c ConvertToIntXX) String() (string, error) { return String(tmConvertToIntXX, c) }

// float32, float64

func ConvertToFloat32(from, to string, newError FuncNewError) Render {
	return Combine{
		ConvertToFloatXX{32, from, "vf", newError},
		AssignNew{TypeConversion{"float32", "vf"}, to},
	}
}

func ConvertToFloat64(from, to string, newError FuncNewError) Render {
	return ConvertToFloatXX{64, from, to, newError}
}

type ConvertToFloatXX struct {
	BitSize     int
	From, ToNew string
	NewError    FuncNewError
}

var tmConvertToFloatXX = template.Must(template.New("ConvertToFloatXX").Parse(`{{.ToNew}}, err := strconv.ParseFloat({{.From}}, {{.BitSize}})
if err != nil {
	return zero, {{call .NewError (print "fmt.Errorf(\"parse float" (printf "%d" .BitSize) ": %w\", err)")}}
}`))

func (c ConvertToFloatXX) String() (string, error) { return String(tmConvertToFloatXX, c) }
