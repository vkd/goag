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
