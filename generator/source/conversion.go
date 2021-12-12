package source

import "text/template"

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
	Error       ErrorWrapper
}

var tmConvertToIntXX = template.Must(template.New("ConvertToIntXX").Parse(`
{{- $bitSize := ""}}
{{- if .BitSize}}{{$bitSize = (printf "%d" .BitSize)}}{{else}}{{$bitSize = ""}}{{end -}}
{{.ToNew}}, err := strconv.ParseInt({{.From}}, 10, {{.BitSize}})
if err != nil {
	return zero, {{.Error.Wrap (print "parse int" $bitSize)}}
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
	return zero, {{.Error.Wrap (print "parse float" (printf "%d" .BitSize))}}
}`))

func (c ConvertToFloatXX) String() (string, error) { return String(tmConvertToFloatXX, c) }
