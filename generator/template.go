package generator

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"reflect"
	"strings"
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
				"newError":   newErrorFunc,
				"returns":    newReturnsFunc,
				"comment":    commentFunc,
				"title":      titleFunc,
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
	if os.Getenv("TEMPLATE_DEBUG") != "" {
		return `/** ` + name + ` >>> */` + bs.String() + `/** <<< ` + name + ` */`, nil
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

func parseErrorFunc(k string, p string) (ErrorRender, error) {
	return parseParamError{k, p}, nil
}

func newErrorFunc() (ErrorRender, error) {
	return newError{}, nil
}

func newReturnsFunc(returns string, e ErrorRender) (ErrorRender, error) {
	return returnsArgs{returns, e}, nil
}

func commentFunc(s string) (string, error) {
	return strings.ReplaceAll(s, "\n", "\n// "), nil
}

func titleFunc(s string) (string, error) {
	return stringsTitle(s), nil
}
