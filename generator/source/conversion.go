package source

import (
	"text/template"
)

func TypeConversion(goType, variable string) Render {
	return typeConversion{goType, variable}
}

var tmTypeConversion = InitTemplate("TypeConversion", `{{.GoType}}({{.Var}})`)

type typeConversion struct {
	GoType string
	Var    string
}

func (c typeConversion) String() (string, error) { return tmTypeConversion.String(c) }

// --- deprecated ---

// int, int32, int64

func ParseInt(to string, from string, newError ErrorWrapper) Render {
	return Renders{
		ParseIntXX{0, from, "vInt", newError},
		AssignNew(to, TypeConversion("int", "vInt")),
	}
}

func ParseInt32(to string, from string, newError ErrorWrapper) Render {
	return Renders{
		ParseIntXX{32, from, "vInt", newError},
		AssignNew(to, TypeConversion("int32", "vInt")),
	}
}

func ParseInt64(to string, from string, newError ErrorWrapper) Render {
	return ParseIntXX{64, from, to, newError}
}

type ParseIntXX struct {
	BitSize int
	From    string
	ToNew   string
	Error   ErrorWrapper
}

var tmParseIntXX = template.Must(template.New("ParseIntXX").Parse(`
{{- $bitSize := ""}}
{{- if .BitSize}}{{$bitSize = (printf "%d" .BitSize)}}{{else}}{{$bitSize = ""}}{{end -}}
{{.ToNew}}, err := strconv.ParseInt({{.From}}, 10, {{.BitSize}})
if err != nil {
	return zero, {{.Error.Wrap (print "parse int" $bitSize)}}
}`))

func (c ParseIntXX) String() (string, error) { return String(tmParseIntXX, c) }

// float32, float64

func ConvertToFloat32(from, to string, newError ErrorWrapper) Render {
	return Renders{
		ConvertToFloatXX{32, from, "vf", newError},
		AssignNew(to, TypeConversion("float32", "vf")),
	}
}

func ConvertToFloat64(from, to string, newError ErrorWrapper) Render {
	return ConvertToFloatXX{64, from, to, newError}
}

type ConvertToFloatXX struct {
	BitSize     int
	From, ToNew string
	Error       ErrorWrapper
}

var tmConvertToFloatXX = template.Must(template.New("ConvertToFloatXX").Parse(`
{{.ToNew}}, err := strconv.ParseFloat({{.From}}, {{.BitSize}})
if err != nil {
	return zero, {{.Error.Wrap (print "parse float" (printf "%d" .BitSize))}}
}`))

func (c ConvertToFloatXX) String() (string, error) { return String(tmConvertToFloatXX, c) }

// bool

func ConvertToBool(from, to string, newError ErrorWrapper) Render {
	return convertToBool{from, to, newError}
}

type convertToBool struct {
	From, ToNew string
	Error       ErrorWrapper
}

var tmConvertToBool = template.Must(template.New("convertToBool").Parse(`
{{.ToNew}}, err := strconv.ParseBool({{.From}})
if err != nil {
	return zero, {{.Error.Wrap "parse bool"}}
}`))

func (c convertToBool) String() (string, error) { return String(tmConvertToBool, c) }
