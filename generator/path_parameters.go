package generator

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type PathParameter struct {
	Name      string
	FieldName string
	// GoType    GoType
	Type SchemaRender
}

func NewPathParameter(p *openapi3.Parameter) PathParameter {
	var out PathParameter
	out.Name = p.Name
	out.FieldName = PublicFieldName(p.Name)
	// out.GoType = NewGoType(p.Schema)
	sr := NewSchema(p.Schema.Value)
	out.Type = sr
	return out
}

func NewPathParamsParsers(path string, params []PathParameter) ([]Render, error) {
	m := make(map[string]PathParameter)
	for _, p := range params {
		m[p.Name] = p
	}

	var out []Render

	p := path
	for len(p) > 0 {
		l := strings.Index(p, "{")
		if l == -1 {
			out = append(out, PathConstantParser{p, path})
			break
		}
		r := strings.Index(p, "}")
		if r == -1 {
			return nil, fmt.Errorf("wrong path: '}' not found")
		}
		if r < l {
			return nil, fmt.Errorf("wrong path: '}' found before '{'")
		}

		out = append(out, PathConstantParser{p[:l], path})

		paramName := p[l+1 : r]

		param, ok := m[paramName]
		if !ok {
			return nil, fmt.Errorf("wrong spec: param %q not found", paramName)
		}
		delete(m, paramName)

		// conv, err := NewConvertFromString(
		// 	param.GoType,
		// 	"v",
		// 	"params."+param.FieldName,
		// 	NewPathErrorFunc(param.Name),
		// )

		to := "params." + param.FieldName

		conv := param.Type.Parser("vPath", "v", ParseError{"path", param.Name})
		out = append(out, PathParameterParser{
			"vPath",
			Combine{conv, Assign{GoValue("v"), to}},
			ParseError{"path", param.Name},
		})

		p = p[r+1:]
	}

	if len(m) > 0 {
		for _, v := range m {
			return nil, fmt.Errorf("wrong spec: %q path param is not used", v.Name)
		}
	}

	return out, nil
}

type PathConstantParser struct {
	Prefix   string
	FullPath string
}

var tmPathConstantParser = template.Must(template.New("PathConstantParser").Parse(`if !strings.HasPrefix(p, "{{.Prefix}}") {
	return zero, fmt.Errorf("wrong path: expected '{{.FullPath}}'")
}
p = p[{{len .Prefix}}:] // "{{.Prefix}}"`))

func (p PathConstantParser) String() (string, error) {
	return String(tmPathConstantParser, p)
}

type PathParameterParser struct {
	Variable string

	Convert Render
	Error   ErrorBuilder
}

var tmPathParameterParser = template.Must(template.New("PathParameterParser").Parse(`{
	idx := strings.Index(p, "/")
	if idx == -1 {
		idx = len(p)
	}
	{{.Variable}} := p[:idx]
	p = p[idx:]

	if len({{.Variable}}) == 0 {
		return zero, {{.Error.New "required"}}
	}

	{{.Convert.String}}
}`))

func (p PathParameterParser) String() (string, error) {
	return String(tmPathParameterParser, p)
}
