package goag

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*.gotmpl
var templates embed.FS

var tmps *template.Template = template.New("generator").Funcs(template.FuncMap{
	"stringsTitle": strings.Title,
	"privateField": func(s string) string {
		return strings.ToLower(s[:1]) + s[1:]
	},
	"stringsJoin": strings.Join,
})

func ParseTemplates() error {
	var t *template.Template = tmps
	fs, err := templates.ReadDir("templates")
	if err != nil {
		return fmt.Errorf("read dir 'templates/': %w", err)
	}
	for _, f := range fs {
		t, err = parseFiles(t, templates, "templates/"+f.Name())
		if err != nil {
			return fmt.Errorf("parse file %q: %w", f.Name(), err)
		}
	}
	return nil
}

type fileStorager interface {
	ReadFile(filename string) ([]byte, error)
}

// parseFiles is the helper for the method and function. If the argument
// template is nil, it is created from the first file.
func parseFiles(t *template.Template, f fileStorager, filename string) (*template.Template, error) {
	s, err := f.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("get file content (%s): %q", filename, err)
	}
	name := filepath.Base(filename)
	// First template becomes return value if not already defined,
	// and we use that one for subsequent New calls to associate
	// all the templates together. Also, if this file has the same name
	// as t, this file becomes the contents of t, so
	//  t, err := New(name).Funcs(xxx).ParseFiles(name)
	// works. Otherwise we create a new template associated with t.
	var tmpl *template.Template
	if t == nil {
		t = template.New(name)
	}
	if name == t.Name() {
		tmpl = t
	} else {
		tmpl = t.New(name)
	}
	_, err = tmpl.Parse(string(s))
	if err != nil {
		return nil, fmt.Errorf("parse template for filename %s: %w", name, err)
	}
	return t, nil
}
