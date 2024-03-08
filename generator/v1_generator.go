package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Generator struct {
	Options GeneratorOptions

	Spec *specification.Spec

	Imports    Imports
	Operations []*Operation
	Paths      []*PathItem

	FileHandler FileHandler
	Router      Router
	Client      Client
}

type PathItem struct {
	*specification.PathItem
	Operations []*Operation
}

var defaultOptions = GeneratorOptions{
	DoNotEdit:   true,
	PackageName: "goag",
}

type GeneratorOptions struct {
	DoNotEdit    bool
	PackageName  string
	BasePath     string
	SpecFilename string
}

func NewGenerator(spec *specification.Spec, opts ...GenOption) (*Generator, error) {
	g := &Generator{
		Options: defaultOptions,
		Spec:    spec,
	}
	for _, opt := range opts {
		opt.apply(&g.Options)
	}

	for _, pi := range spec.PathItems {
		pathItem := &PathItem{
			PathItem: pi,
		}
		for _, o := range pi.Operations {
			operation, ims, err := NewOperation(o, spec.Components)
			if err != nil {
				return nil, fmt.Errorf(": %w", err)
			}
			g.Imports = append(g.Imports, ims...)

			g.Operations = append(g.Operations, operation)
			pathItem.Operations = append(pathItem.Operations, operation)
		}
		g.Paths = append(g.Paths, pathItem)
	}

	var err error

	g.FileHandler, err = NewFileHandler(g.Operations, g.Options.BasePath)
	if err != nil {
		return nil, fmt.Errorf("file handler: %w", err)
	}
	g.Client = NewClient(spec, g.Operations)
	g.Router = NewRouter(spec, g.Paths, g.Operations, g.Options)

	return g, nil
}

func SkipDoNotEdit() GenOption {
	return genOptionFunc(func(o *GeneratorOptions) {
		o.DoNotEdit = false
	})
}

func PackageName(packageName string) GenOption {
	return genOptionFunc(func(o *GeneratorOptions) {
		o.PackageName = packageName
	})
}

func BasePath(basePath string) GenOption {
	return genOptionFunc(func(o *GeneratorOptions) {
		o.BasePath = basePath
	})
}

func SpecFilename(specFilename string) GenOption {
	return genOptionFunc(func(o *GeneratorOptions) {
		o.SpecFilename = specFilename
	})
}

type GenOption interface {
	apply(*GeneratorOptions)
}

type genOptionFunc func(*GeneratorOptions)

func (f genOptionFunc) apply(o *GeneratorOptions) { f(o) }

func IfOption(opt GenOption, ifCond ...bool) GenOption {
	for _, cnd := range ifCond {
		if !cnd {
			return genOptionFunc(func(g *GeneratorOptions) {})
		}
	}
	return opt
}

func (g *Generator) goFile(ims []Import, body any) Templater {
	return TemplaterFunc(GoFile{
		SkipDoNotEdit: !g.Options.DoNotEdit,
		PackageName:   g.Options.PackageName,
		Imports:       ims,
		Body:          body,
	}.Render)
}

func (g *Generator) FileHandlerTemplater(hs []Render, isJSON bool) (Templater, error) {
	return g.goFile(nil, g.FileHandler), nil
}
