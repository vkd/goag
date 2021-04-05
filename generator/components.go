package generator

import (
	"fmt"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

type Components struct {
	Schemas []GoTypeDef
}

func NewComponents(spec openapi3.Components) (zero Components, _ error) {
	tds, err := NewGoTypeDefs(NewSchemas(spec.Schemas))
	if err != nil {
		return zero, fmt.Errorf("new schemas: %w", err)
	}
	return Components{
		Schemas: tds,
	}, nil
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
