package generator

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/vkd/goag/specification"
)

type Router struct {
	BasePath     string
	SpecFilename string
	SpecFileExt  string

	Imports    Imports
	PathItems  []*RouterPathItem
	Operations []*Operation
	Routes     []*Route

	JWT  bool
	Cors bool
}

func NewRouter(s *specification.Spec, ps []*PathItem, os []*Operation, opt GeneratorOptions) Router {
	r := Router{
		BasePath:     opt.BasePath,
		SpecFilename: opt.SpecFilename,
		SpecFileExt:  strings.TrimPrefix(filepath.Ext(opt.SpecFilename), "."),

		Operations: os,

		Cors: opt.IsCors,
	}
	for _, pi := range ps {
		p := &RouterPathItem{
			RawPath:    pi.RawPath,
			HasOptions: pi.PathItem.HasOperation(http.MethodOptions),
		}
		for _, o := range pi.Operations {
			p.Operations = append(p.Operations, RouterPathItemOperation{
				Name:     o.Name,
				Method:   o.Method,
				PathSpec: o.PathItem.Path.Spec,
				Handler:  string(o.Name) + "Handler",
			})
			if !p.HasOptions && opt.IsCors {
				p.Operations = append(p.Operations, RouterPathItemOperation{
					Name:     o.Name,
					Method:   "Options",
					PathSpec: o.PathItem.Path.Spec,
					Handler:  "CORSHandler",
				})
				p.HasOptions = true
			}
		}

		for _, o := range pi.Operations {
			for _, s := range o.Security {
				if s.Scheme.Type == specification.SecuritySchemeTypeHTTP && s.Scheme.Scheme == "bearer" {
					r.JWT = true
					p.JWT = true
				}
			}
		}

		r.PathItems = append(r.PathItems, p)
	}

	root := &Route{
		BasePath: r.BasePath,
		mRoutes:  make(map[string]*Route),
	}

	for _, pi := range r.PathItems {
		root.Add(pi)
	}
	r.Routes = append(r.Routes, root.GetRoutes()...)

	return r
}

func (r Router) Render() (string, error) {
	return ExecuteTemplate("Router", r)
}

type RouterPathItem struct {
	RawPath string

	Operations []RouterPathItemOperation

	JWT        bool
	HasOptions bool
}

type RouterPathItemOperation struct {
	Name     OperationName
	Method   specification.HTTPMethodTitle
	PathSpec string
	Handler  string
}

type Route struct {
	Name     string
	BasePath string
	Prefix   string

	PrefixPathItems []*RoutePathItem
	Variable        *RoutePathItem

	Routes        []*Route
	mRoutes       map[string]*Route
	VariableRoute *Route
}

func (r *Route) GetRoutes() []*Route {
	if r == nil {
		return nil
	}
	var out []*Route
	out = append(out, r)
	for _, route := range r.Routes {
		out = append(out, route.GetRoutes()...)
	}
	out = append(out, r.VariableRoute.GetRoutes()...)
	return out
}

func (r *Route) Add(pi *RouterPathItem) {
	path := pi.RawPath
	r.add(pi, strings.Split(strings.TrimPrefix(path, "/"), "/"))
}

func (r *Route) add(pi *RouterPathItem, dirs []string) {
	if len(dirs) == 0 {
		return
	}

	d := dirs[0]
	if len(dirs) == 1 {
		if strings.HasPrefix(d, "{") && strings.HasSuffix(d, "}") {
			r.Variable = &RoutePathItem{
				RouterPathItem: pi,
				Prefix:         "/" + d[1:len(d)-1],
			}
		} else {
			r.PrefixPathItems = append(r.PrefixPathItems, &RoutePathItem{
				RouterPathItem: pi,
				Prefix:         "/" + d,
			})
		}
		return
	}

	dirs = dirs[1:]

	if strings.HasPrefix(d, "{") && strings.HasSuffix(d, "}") {
		variableRoute := r.VariableRoute
		if variableRoute == nil {
			variableRoute = &Route{
				Name:    r.Name + PublicFieldName(d[1:len(d)-1]),
				Prefix:  "/" + d,
				mRoutes: make(map[string]*Route),
			}
			r.VariableRoute = variableRoute
		}
		variableRoute.add(pi, dirs)
	} else {
		if next, ok := r.mRoutes[d]; ok {
			next.add(pi, dirs)
		} else {
			route := &Route{
				Name:    r.Name + PublicFieldName(d),
				Prefix:  "/" + d,
				mRoutes: make(map[string]*Route),
			}
			r.Routes = append(r.Routes, route)
			r.mRoutes[d] = route
			route.add(pi, dirs)
		}
	}
}

type RoutePathItem struct {
	*RouterPathItem
	Prefix string
}
