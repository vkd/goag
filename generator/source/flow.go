package source

import "text/template"

type Combine []Render

var tmCombine = template.Must(template.New("Combine").Parse(`
{{- range $i, $c := . }}{{ if $i }}
{{ end }}
{{- $c.String }}
{{- end }}`))

func (c Combine) String() (string, error) { return String(tmCombine, c) }
