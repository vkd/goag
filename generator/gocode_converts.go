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

type ErrorWrapper interface {
	Wrap(reason string) string
	Error(reason string) string
}

type ParseErrorWrapper struct {
	In, Parameter string
}

func (p ParseErrorWrapper) Wrap(reason string) string {
	return `ErrParseParam{In: "` + p.In + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `", Err: err}`
}

func (p ParseErrorWrapper) Error(reason string) string {
	return `ErrParseParam{In: "` + p.In + `", Parameter: "` + p.Parameter + `", Reason: "` + reason + `"}`
}

type StructParser struct {
	From, To string
	NewError ErrorWrapper
}

var tmStructParser = template.Must(template.New("StructParser").Parse(`err := {{.To}}.UnmarshalText({{.From}})
if err != nil {
	return zero, {{.NewError.Wrap "parse struct"}}
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

func ConvertToInt(from, to string, newError ErrorWrapper) Render {
	return Combine{
		ConvertToIntXX{0, from, "vInt", newError},
		AssignNew{TypeConversion{"int", "vInt"}, to},
	}
}

func ConvertToInt32(from, to string, newError ErrorWrapper) Render {
	return Combine{
		ConvertToIntXX{32, from, "vInt", newError},
		AssignNew{TypeConversion{"int32", "vInt"}, to},
	}
}

func ConvertToInt64(from, to string, newError ErrorWrapper) Render {
	return ConvertToIntXX{64, from, to, newError}
}

type ConvertToIntXX struct {
	BitSize     int
	From, ToNew string
	NewError    ErrorWrapper
}

var tmConvertToIntXX = template.Must(template.New("ConvertToIntXX").Parse(`
{{- $bitSize := ""}}
{{- if .BitSize}}{{$bitSize = (printf "%d" .BitSize)}}{{else}}{{$bitSize = ""}}{{end -}}
{{.ToNew}}, err := strconv.ParseInt({{.From}}, 10, {{.BitSize}})
if err != nil {
	return zero, {{.NewError.Wrap (print "parse int" $bitSize)}}
}`))

func (c ConvertToIntXX) String() (string, error) { return String(tmConvertToIntXX, c) }

// float32, float64

func ConvertToFloat32(from, to string, newError ErrorWrapper) Render {
	return Combine{
		ConvertToFloatXX{32, from, "vf", newError},
		AssignNew{TypeConversion{"float32", "vf"}, to},
	}
}

func ConvertToFloat64(from, to string, newError ErrorWrapper) Render {
	return ConvertToFloatXX{64, from, to, newError}
}

type ConvertToFloatXX struct {
	BitSize     int
	From, ToNew string
	NewError    ErrorWrapper
}

var tmConvertToFloatXX = template.Must(template.New("ConvertToFloatXX").Parse(`{{.ToNew}}, err := strconv.ParseFloat({{.From}}, {{.BitSize}})
if err != nil {
	return zero, {{.NewError.Wrap (print "parse float" (printf "%d" .BitSize))}}
}`))

func (c ConvertToFloatXX) String() (string, error) { return String(tmConvertToFloatXX, c) }
