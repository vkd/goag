package generator

import (
	"github.com/vkd/goag/specification"
)

type Generator struct {
	skipDoNotEdit bool
	packageName   string
	spec          *specification.Spec
}

type GenOption func(*Generator)

func SkipDoNotEdit(ifArgs ...bool) GenOption {
	for _, a := range ifArgs {
		if !a {
			return func(g *Generator) {}
		}
	}
	return func(g *Generator) {
		g.skipDoNotEdit = true
	}
}

func NewGenerator(spec *specification.Spec, packageName string, opts ...GenOption) (*Generator, error) {
	g := &Generator{
		packageName: packageName,
		spec:        spec,
	}
	for _, o := range opts {
		o(g)
	}
	return g, nil
}

func (g *Generator) HandlersFile(hs []Handler, isJSON bool) (Templater, error) {
	file := HandlersFile{
		Handlers:        hs,
		IsWriteJSONFunc: isJSON,
	}

	return g.goFile([]string{
		"encoding/json",
		"fmt",
		"io",
		"log",
		"net/http",
		"strconv",
		"strings",
	}, file), nil
}

func (g *Generator) RouterFile(hs []Handler, oldRouter any) (Templater, error) {
	file := NewRouterFile(g.spec, hs, oldRouter)

	return g.goFile([]string{
		"net/http",
		"strings",
	}, file), nil
}

func (g *Generator) goFile(ims []string, body executor) Templater {
	return OldTemplater(GoFile{
		SkipDoNotEdit: g.skipDoNotEdit,
		PackageName:   g.packageName,
		Imports:       ims,
		Body:          OldTemplater(body),
	})
}
