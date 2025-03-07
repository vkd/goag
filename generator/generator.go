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

	Client     ClientTemplate
	Components *Components

	Router Router

	HandlersFile    HandlersFileTemplate
	Handlers        []*Handler
	IsWriteJSONFunc bool
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
	IsCors       bool
}

func NewGenerator(spec *specification.Spec, cfg Config, opts ...GenOption) (*Generator, error) {
	g := &Generator{
		Options: defaultOptions,
		Spec:    spec,
	}
	if cfg.Cors.Enable {
		g.Options.IsCors = true
	}
	for _, opt := range opts {
		opt.apply(&g.Options)
	}

	components, ims, err := NewComponents(spec.Components, cfg)
	if err != nil {
		return nil, fmt.Errorf("file components: %w", err)
	}
	g.Components = &components
	g.Imports = append(g.Imports, ims...)

	// ---

	for _, pi := range spec.PathItems {
		pathItem := &PathItem{
			PathItem: pi,
		}
		for _, o := range pi.Operations {
			operation, ims, err := NewOperation(o, g.Components, cfg)
			if err != nil {
				return nil, fmt.Errorf(": %w", err)
			}
			g.Imports = append(g.Imports, ims...)

			g.Operations = append(g.Operations, operation)
			pathItem.Operations = append(pathItem.Operations, operation)
		}
		g.Paths = append(g.Paths, pathItem)
	}

	var isWriteJSONFunc bool
	for _, o := range g.Operations {
		h, ims, err := NewHandler(o, g.Options.BasePath, g.Components, cfg)
		if err != nil {
			return nil, fmt.Errorf("handler %q: %w", o.Name, err)
		}
		g.Imports = append(g.Imports, ims...)
		g.Handlers = append(g.Handlers, h)

		for _, r := range h.Responses {
			if _, ok := r.ContentJSON.Get(); ok {
				isWriteJSONFunc = true
			}
		}
		if h.DefaultResponse != nil && h.DefaultResponse.ContentJSON.IsSet {
			isWriteJSONFunc = true
		}
	}

	if g.Components.HasContentJSON {
		isWriteJSONFunc = true
	}

	g.HandlersFile = NewHandlersFileTemplate(g.Handlers, isWriteJSONFunc, cfg)

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
