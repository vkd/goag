package generator

import (
	"github.com/vkd/goag/specification"
)

type Generator struct {
	skipDoNotEdit bool
	packageName   string
	spec          *specification.Spec

	Paths      []*PathItem
	Operations []*Operation
}

func SkipDoNotEdit() GenOption {
	return genOptionFunc(func(g *Generator) {
		g.skipDoNotEdit = true
	})
}

func NewGenerator(spec *specification.Spec, packageName string, opts ...GenOption) (*Generator, error) {
	g := &Generator{
		packageName: packageName,
		spec:        spec,
	}
	for _, opt := range opts {
		opt.apply(g)
	}

	for _, pi := range spec.Paths {
		pathItem := &PathItem{
			PathItem: pi,
		}
		for _, o := range pi.Operations {
			operation := NewOperation(o)
			g.Operations = append(g.Operations, operation)
			pathItem.Operations = append(pathItem.Operations, operation)
		}
		g.Paths = append(g.Paths, pathItem)
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

func (g *Generator) goFile(ims []string, body executor) Templater {
	return OldTemplater(GoFile{
		SkipDoNotEdit: g.skipDoNotEdit,
		PackageName:   g.packageName,
		Imports:       ims,
		Body:          OldTemplater(body),
	})
}

type GenOption interface {
	apply(*Generator)
}

type genOptionFunc func(*Generator)

func (f genOptionFunc) apply(g *Generator) { f(g) }

func IfOption(opt GenOption, ifCond ...bool) GenOption {
	for _, cnd := range ifCond {
		if !cnd {
			return genOptionFunc(func(g *Generator) {})
		}
	}
	return opt
}
