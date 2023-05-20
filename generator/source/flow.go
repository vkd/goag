package source

import "text/template"

type Renders []Render

var tmRenders = template.Must(template.New("Renders").Parse(`
{{- range $i, $c := . }}{{ if $i }}
{{ end }}
{{- $c.String }}
{{- end }}`))

func (c Renders) String() (string, error) { return String(tmRenders, c) }
