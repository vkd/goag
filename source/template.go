package source

import (
	"bytes"
	"fmt"
	"text/template"
)

type Templater interface {
	Execute() (string, error)
}

func MustTemplate(name, text string) *Template {
	return &Template{
		tm: template.Must(template.New(name).Parse(text)),
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
