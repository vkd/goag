package generator

import (
	"bytes"
	"embed"
	"fmt"
	"reflect"
	"text/template"
)

//go:embed *.gotmpl
var templatesFS embed.FS
var templates *template.Template

func init() {
	templates = template.Must(
		template.
			New("Generator").
			Funcs(template.FuncMap{
				"private":    privateTemplateFunc,
				"parseError": parseErrorFunc,
			}).
			ParseFS(templatesFS, "*.gotmpl"),
	)

}

func ExecuteTemplate(name string, data any) (string, error) {
	var bs bytes.Buffer
	err := templates.ExecuteTemplate(&bs, name, data)
	if err != nil {
		return "", fmt.Errorf("execute template (%s): %w", name, err)
	}
	return bs.String(), nil
}

type TData map[string]any

// --- Functions ---

func privateTemplateFunc(t reflect.Value) (string, error) {
	switch t := t.Interface().(type) {
	case string:
		return PrivateFieldName(t), nil
	}

	return "", fmt.Errorf("%T is not string", t.Interface())
}

func parseErrorFunc(k reflect.Value, p reflect.Value) (ErrorRender, error) {
	var kindParam string
	switch kind := k.Interface().(type) {
	case string:
		kindParam = kind
	default:
		return nil, fmt.Errorf("kind of params: %T is not string", k.Interface())
	}

	var paramName string
	switch par := p.Interface().(type) {
	case string:
		paramName = par
	default:
		return nil, fmt.Errorf("param name: %T is not string", p.Interface())
	}

	return parseError{kindParam, paramName}, nil
}
