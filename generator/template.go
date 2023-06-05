package source

import (
	"bytes"
	"fmt"
	"reflect"
	"text/template"
)

type Templater interface {
	Execute() (string, error)
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

// --- Functions ---

func execTemplateFunc(t reflect.Value) (string, error) {
	if t.IsNil() {
		return "", fmt.Errorf("cannot execute 'nil'")
	}
	tmp, ok := t.Interface().(Templater)
	if !ok {
		return "", fmt.Errorf("%T does not implement Templater: missing method Execute() (string, error)", t.Interface())
	}
	return tmp.Execute()
}
