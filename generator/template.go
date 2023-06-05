package generator

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"
)

type Templater interface {
	// Execute() (string, error)
}

func InitTemplate(name, text string) *Template {
	return &Template{
		tm: template.Must(template.New(name).Funcs(template.FuncMap{
			"exec": execTemplateFunc,
		}).Parse(text)),
	}
}

type Template struct {
	tm *template.Template
}

func (t *Template) Execute(data interface{}) (string, error) {
	var bs bytes.Buffer
	err := t.tm.Execute(&bs, data)
	if err != nil {
		return "", fmt.Errorf("execute template (%s): %w", t.tm.Name(), err)
	}
	return bs.String(), nil
}

type Templaters []Template

var tmTemplaters = InitTemplate("Templaters", `
{{ range $_, $t := . }}
{{ exec $t }}
{{ end }}
`)

func (t Templaters) Execute() (string, error) { return tmTemplaters.Execute(t) }

// --- Functions ---

func execTemplateFunc(t reflect.Value) (string, error) {
	str, ok := t.Interface().(interface {
		String() (string, error)
	})
	if ok {
		return str.String()
	}
	tmp, ok := t.Interface().(interface {
		Execute() (string, error)
	})
	if !ok {
		return "", fmt.Errorf("%T does not implement Templater: missing method Execute() (string, error)", t.Interface())
	}
	return tmp.Execute()
}
