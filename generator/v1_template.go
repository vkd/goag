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
var templates *Template

func init() {
	templates = &Template{
		template.Must(
			template.
				New("Generator").
				Funcs(template.FuncMap{
					"exec":    execTemplateFunc,
					"private": privateTemplateFunc,
					"raw":     rawTemplateFunc,
					"render":  renderTemplateFunc,
				}).
				ParseFS(templatesFS, "*.gotmpl"),
		),
	}
}

type Templater interface {
	// Execute() (string, error)
	String() (string, error)
}

type TemplaterFunc func() (string, error)

func (t TemplaterFunc) String() (string, error) { return t() }

type ExecTemplater interface {
	Execute() (string, error)
	// String() (string, error)
}

type executor interface {
	Execute() (string, error)
}

func OldTemplater(t interface{ Execute() (string, error) }) Templater {
	return executorWrapper{t: t}
}

type executorWrapper struct {
	t interface{ Execute() (string, error) }
}

func (e executorWrapper) String() (string, error) { return e.t.Execute() }

func InitTemplate(name, text string) *Template {
	return &Template{
		tm: template.Must(template.New(name).Funcs(template.FuncMap{
			"exec":    execTemplateFunc,
			"private": privateTemplateFunc,
		}).Parse(text)),
	}
}

func ExecuteTemplate(name string, data any) (string, error) {
	var bs bytes.Buffer
	err := templates.tm.ExecuteTemplate(&bs, name, data)
	if err != nil {
		return "", fmt.Errorf("execute template (%s): %w", name, err)
	}
	return bs.String(), nil
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

func (t *Template) ExecuteTemplate(name string, data interface{}) (string, error) {
	return ExecuteTemplate(name, data)
}

type Templaters []Templater

var tmTemplaters = InitTemplate("Templaters", `{{ range $_, $t := . }}{{ exec $t }}{{ end }}`)

func (t Templaters) Execute() (string, error) { return tmTemplaters.Execute(t) }

func (t Templaters) ExecuteArgs(args ...any) (string, error) {
	return tmTemplaters.Execute(t)
}

func TemplateData(name string, data interface{}) Templater {
	return TemplaterFunc(func() (string, error) { return ExecuteTemplate(name, data) })
}

type RawTemplate string

func (t RawTemplate) Execute() (string, error) {
	return string(t), nil
}
func (t RawTemplate) String() (string, error) { return t.Execute() }

type TData map[string]any

// --- Functions ---

func execTemplateFunc(t reflect.Value, args ...any) (string, error) {
	v := t.Interface()
	switch t := v.(type) {
	case Templater:
		return t.String()
	case string:
		return t, nil
	}

	if r, ok := v.(interface{ Render() (string, error) }); ok {
		return r.Render()
	}
	if t, ok := v.(interface {
		Execute() (string, error)
	}); ok {
		return t.Execute()
	}

	return "", fmt.Errorf("%T does not implement Templater: missing method Execute() (string, error)", t.Interface())
}

func privateTemplateFunc(t reflect.Value) (string, error) {
	switch t := t.Interface().(type) {
	case string:
		return PrivateFieldName(t), nil
	}

	return "", fmt.Errorf("%T is not string", t.Interface())
}

func rawTemplateFunc(t reflect.Value) (Render, error) {
	switch t := t.Interface().(type) {
	case string:
		return StringRender(t), nil
	}

	return nil, fmt.Errorf("%T is not string", t.Interface())
}

func renderTemplateFunc(value reflect.Value) (string, error) {
	return IfaceRender(value)
}

func IfaceRender(value reflect.Value) (string, error) {
	for value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	v := value.Interface()

	switch t := v.(type) {
	case string:
		return t, nil
	case interface {
		Render() (string, error)
	}:
		return t.Render()
	}

	return ExecuteTemplate(value.Type().Name(), v)
}
