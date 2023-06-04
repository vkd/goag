package source

type GoFile struct {
	PackageName string
	Imports     []string
	Body        []Templater
}

var tmGoFile = InitTemplate("GoFile", `
package {{ .PackageName }}

{{ if .Imports -}}
import (
	{{ range $_, $i := .Imports }}
	"{{$i}}"
	{{- end }}
)
{{- end }}

{{ range $_, $b := .Body }}
{{ exec $b }}
{{ end }}
`)

func (g GoFile) String() (string, error) { return tmGoFile.String(g) }
