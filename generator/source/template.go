package source

import (
	"bytes"
	"fmt"
	"text/template"
)

type Templater interface {
	String() (string, error)
}

func InitTemplate(name, text string) *Template {
	return &Template{
		tm: template.Must(template.New(name).Parse(text)),
	}
}

type Template struct {
	tm *template.Template
}

func (t *Template) String(data interface{}) (string, error) {
	var bs bytes.Buffer
	err := t.tm.Execute(&bs, data)
	if err != nil {
		return "", fmt.Errorf("execute template: %w", err)
	}
	return bs.String(), nil
}

// --- deprecated ---

type Render = Templater

func String(tm *template.Template, data interface{}) (string, error) {
	var bs bytes.Buffer
	err := tm.Execute(&bs, data)
	if err != nil {
		return "", fmt.Errorf("to string: %w", err)
	}
	return bs.String(), nil
}
