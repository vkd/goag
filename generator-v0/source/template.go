package source

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"
)

type Templater interface {
	String() (string, error)
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

func (t *Template) String(data interface{}) (string, error) {
	var bs bytes.Buffer
	err := t.tm.Execute(&bs, data)
	if err != nil {
		return "", fmt.Errorf("execute template: %w", err)
	}
	return bs.String(), nil
}

func execTemplateFunc(t reflect.Value) (string, error) {
	if t.IsNil() {
		return "", fmt.Errorf("cannot exec 'nil'")
	}
	tmp, ok := t.Interface().(Templater)
	if !ok {
		return "", fmt.Errorf("unexpected type %T: 'Templater' is expected", t.Interface())
	}
	return tmp.String()
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
