package generator

type GoFile struct {
	PackageName string
	Imports     []string
	Body        Templater
}

var tmGoFile = InitTemplate("GoFile", `
package {{ .PackageName }}

{{ if .Imports -}}
import (
	{{ range $_, $i := .Imports }}
	"{{ $i }}"
	{{- end }}
)
{{- end }}

{{ exec .Body }}
`)

func (g GoFile) Execute() (string, error) { return tmGoFile.Execute(g) }
