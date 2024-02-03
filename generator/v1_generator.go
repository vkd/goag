package generator

import (
	"fmt"

	"github.com/vkd/goag/specification"
)

type Generator struct {
	Options Options

	Spec *specification.Spec

	Handlers []Handler
	Client   Client

	// deprecated

	Paths      []*PathItem
	Operations []*Operation
}

type PathItem struct {
	PathItem   *specification.PathItem
	Operations []*Operation
}

var defaultOptions = Options{
	DoNotEdit:   true,
	PackageName: "goag",
}

type Options struct {
	DoNotEdit   bool
	PackageName string
}

func NewGenerator(spec *specification.Spec, opts ...GenOption) (*Generator, error) {
	g := &Generator{
		Options: defaultOptions,
		Spec:    spec,
	}
	for _, opt := range opts {
		opt.apply(&g.Options)
	}

	for _, pi := range spec.Paths {
		pathItem := &PathItem{
			PathItem: pi,
		}
		for _, o := range pi.Operations {
			operation := NewOperation(o)
			g.Operations = append(g.Operations, operation)
			pathItem.Operations = append(pathItem.Operations, operation)

			h, err := NewHandler(o)
			if err != nil {
				return nil, fmt.Errorf("new handler %q: %w", operation.Name, err)
			}
			g.Handlers = append(g.Handlers, h)
		}
		g.Paths = append(g.Paths, pathItem)
	}

	g.Client = NewClient(g.Handlers)

	return g, nil
}

func SkipDoNotEdit() GenOption {
	return genOptionFunc(func(o *Options) {
		o.DoNotEdit = false
	})
}

func PackageName(packageName string) GenOption {
	return genOptionFunc(func(o *Options) {
		o.PackageName = packageName
	})
}

type GenOption interface {
	apply(*Options)
}

type genOptionFunc func(*Options)

func (f genOptionFunc) apply(o *Options) { f(o) }

func IfOption(opt GenOption, ifCond ...bool) GenOption {
	for _, cnd := range ifCond {
		if !cnd {
			return genOptionFunc(func(g *Options) {})
		}
	}
	return opt
}
