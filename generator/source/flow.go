package source

type Renders []Render

var tmRenders = InitTemplate("Renders", `
{{- range $i, $c := . }}{{ if $i }}
{{ end }}
{{- $c.String }}
{{- end }}`)

func (c Renders) String() (string, error) {
	switch len(c) {
	case 0:
		return "", nil
	case 1:
		return c[0].String()
	}
	return tmRenders.String(c)
}
