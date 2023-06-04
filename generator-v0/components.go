package generator

import (
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type Components struct {
	Schemas []GoTypeDef
}

func NewComponents(spec openapi3.Components) Components {
	tds := NewGoTypeDefs(NewSchemas(spec.Schemas))
	return Components{
		Schemas: tds,
	}
}

var tmComponents = template.Must(template.New("Components").Parse(`
{{- if .Schemas -}}
// ------------------------
//         Schemas
// ------------------------

{{range $_, $s := .Schemas}}
{{$s.String}}
{{end}}
{{- end}}`))

func (c Components) String() (string, error) {
	return String(tmComponents, c)
}
