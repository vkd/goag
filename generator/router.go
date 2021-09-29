package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Router struct {
	PackageName string
	BasePath    string
	SpecFile    string
	SpecFileExt string

	Handlers []Handler

	Routes []Route
}

func NewRouter(packageName string, handlers []Handler, spec *openapi3.Swagger, specRaw []byte, baseFilename string) (Router, error) {
	var out Router
	out.PackageName = packageName
	if len(spec.Servers) > 0 {
		out.BasePath = spec.Servers[0].URL
	}
	out.SpecFile = string(specRaw)
	out.SpecFileExt = filepath.Ext(out.SpecFile)

	out.Handlers = handlers
	rts, err := NewRoutes("", handlers)
	if err != nil {
		return Router{}, fmt.Errorf("new routes: %w", err)
	}
	if len(rts) == 0 {
		rts = append(rts, Route{})
	}
	rts[0].Handlers = append(rts[0].Handlers, RouteHandler{
		Prefix: "/" + baseFilename,
		Methods: []RouteMethod{{
			Method:      "Get",
			HandlerName: "SpecFile",
			Path:        "/" + baseFilename,
		}},
	})
	out.Routes = rts
	return out, nil
}

func NewRoutes(routePrefix string, handlers []Handler) ([]Route, error) {
	var route Route
	route.Name = PathName(routePrefix)

	m, ps := groupByPrefix(routePrefix, handlers)

	var out []Route
	for _, p := range ps {
		hs := m[p]
		if p.isRoute {
			if p.IsWildcard() {
				pn := PathName(routePrefix + p.prefix)
				if route.WildcardRouteName != "" && route.WildcardRouteName != pn {
					return nil, fmt.Errorf("unsupported multiple wildcards: exist %q, new %q", route.WildcardRouteName, pn)
				}
				route.WildcardRouteName = pn
			} else {
				route.Routes = append(route.Routes, ReRoute{p.prefix, PathName(routePrefix + p.prefix)})
			}

			rts, err := NewRoutes(routePrefix+p.prefix, hs)
			if err != nil {
				return nil, fmt.Errorf("new routes: %w", err)
			}
			out = append(out, rts...)
		} else {
			if p.IsWildcard() {
				if route.WildcardHandler != nil && route.WildcardHandler.Prefix != p.prefix {
					return nil, fmt.Errorf("unsupported multiple wildcards handlers: exist %q, new %q", route.WildcardHandler.Prefix, p.prefix)
				}
				if route.WildcardHandler == nil {
					route.WildcardHandler = &RouteHandler{Prefix: p.prefix}
				}
				for _, h := range hs {
					route.WildcardHandler.Methods = append(route.WildcardHandler.Methods, RouteMethod{
						Method:      strings.Title(strings.ToLower(h.Method)),
						HandlerName: h.Name,
						Path:        h.Path,
					})
				}
			} else {
				rh := RouteHandler{Prefix: p.prefix}
				for _, h := range hs {
					rh.Methods = append(rh.Methods, RouteMethod{
						Method:      strings.Title(strings.ToLower(h.Method)),
						HandlerName: h.Name,
						Path:        h.Path,
					})
				}
				route.Handlers = append(route.Handlers, rh)
			}
		}
	}

	return append([]Route{route}, out...), nil
}

type pathPrefix struct {
	prefix  string
	isRoute bool
}

func (p pathPrefix) IsWildcard() bool {
	return strings.HasPrefix(p.prefix, "/{") && strings.HasSuffix(p.prefix, "}")
}

func groupByPrefix(routePrefix string, hs []Handler) (map[pathPrefix][]Handler, []pathPrefix) {
	var prefixes []pathPrefix
	m := make(map[pathPrefix][]Handler)

	for i, h := range hs {
		path := h.Path
		path = strings.TrimPrefix(path, routePrefix)
		prefix, path := splitPath(path)

		k := pathPrefix{prefix, path != ""}
		if _, ok := m[k]; !ok {
			prefixes = append(prefixes, k)
		}
		m[k] = append(m[k], hs[i])
	}

	return m, prefixes
}

type Route struct {
	Name string

	Handlers        []RouteHandler
	WildcardHandler *RouteHandler

	Routes            []ReRoute
	WildcardRouteName string
}

type RouteHandler struct {
	Prefix  string
	Methods []RouteMethod
}

type ReRoute struct {
	Prefix    string
	RouteName string
}

type RouteMethod struct {
	Method      string
	HandlerName string
	Path        string
}

func splitPath(s string) (string, string) {
	if !strings.HasPrefix(s, "/") {
		return s, ""
	}
	idx := strings.Index(s[1:], "/")
	if idx == -1 {
		return s, ""
	}
	return s[:idx+1], s[idx+1:]
}
