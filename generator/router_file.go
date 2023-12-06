package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/vkd/goag/specification"
)

type RouterFile struct {
	BasePath         string
	BaseSpecFilename string
	SpecFileExt      string

	Handlers []routerHandler
	Routes   []routeMethod
}

type routerHandler struct {
	Name string
	Type string
}

func (g *Generator) RouterFile(basePath, baseFilename string, hs []HandlerOld, oldRouter any) (Templater, error) {
	file := RouterFile{
		BasePath:         basePath,
		BaseSpecFilename: baseFilename,
		SpecFileExt:      strings.TrimPrefix(filepath.Ext(baseFilename), "."),
		// Data:             oldRouter,
	}
	if len(hs) != len(g.Operations) {
		panic("wrong Operations")
	}
	for i, o := range g.Operations {
		o.Handler = &hs[i]
		file.Handlers = append(file.Handlers, routerHandler{
			Name: o.Name + "Handler",
			Type: o.HandlerTypeName,
		})
	}

	routes, err := g.newRoutes()
	if err != nil {
		return nil, fmt.Errorf("new routes: %w", err)
	}
	file.Routes = routes

	// for _, pi := range g.Paths {
	// 	err := file.Router.add(pi.pathItem.Path, pi)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("add %q route: %w", pi.pathItem.Path, err)
	// 	}
	// }

	return g.goFile([]string{
		"net/http",
		"strings",
	}, file), nil
}

var tmRouterFile = InitTemplate("RouterFile", `
type API struct {
	{{range $_, $h := .Handlers}}
	{{$h.Name}} {{$h.Type}}
	{{- end}}

	// not found
	NotFoundHandler http.Handler
	// spec file
	SpecFileHandler http.Handler

	Middlewares []func(h http.Handler) http.Handler
}

func (rt *API) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if rt.SpecFileHandler != nil && path == "{{.BasePath}}/{{.BaseSpecFilename}}" {
		rt.SpecFileHandler.ServeHTTP(rw, r)
		return
	}

	h, path := rt.route(path, r.Method)
	if h == nil {
		h = rt.NotFoundHandler
		if h == nil {
			h = http.NotFoundHandler()
		}
		h.ServeHTTP(rw, r)
		return
	}

	for i := len(rt.Middlewares) - 1; i >= 0; i-- {
		h = rt.Middlewares[i](h)
	}
	r = r.WithContext(context.WithValue(r.Context(), pathKey{}, path))
	h.ServeHTTP(rw, r)
}

{{- $basePath := .BasePath }}
{{ range $_, $r := .Routes }}
func (rt *API) route{{$r.Name}}(path, method string) (http.Handler, string) {
	{{- if and $basePath (not $r.Name) }}
	if !strings.HasPrefix(path, "{{$basePath}}") {
		return nil, ""
	}
	path = path[{{len $basePath }}:] // "{{$basePath}}"

	if !strings.HasPrefix(path, "/") {
		return nil, ""
	}

	{{ end -}}

	{{if or .PrefixPathItems .Routes}}prefix, path :{{else}}_, path {{end -}}= splitPath(path)

	{{if or .PrefixPathItems .PathItem -}}
	if path == "" {
		{{if .PrefixPathItems -}}
		switch prefix {
			{{range $_, $h := .PrefixPathItems -}}
		case "{{.Prefix}}":
			switch method {
				{{- range $_, $m := .PathItem.Operations}}
			case http.Method{{.Operation.Method}}:
				return rt.{{.Handler.Name}}Handler, "{{.Operation.PathItem.Path.Spec}}"
				{{- end}}
			}
			{{end -}}
		}
		{{- end}}
		{{- if .PathItem}}
		switch method {
			{{- range $_, $m := .PathItem.Operations}}
		case http.Method{{.Operation.Method}}:
			return rt.{{.Handler.Name}}Handler, "{{.Operation.PathItem.Path.Spec}}"
			{{- end}}
		}
		{{- end}}
	}
	{{- end}}

	{{if or .Routes .VariableRoute -}}
	if path != "" {
		{{if .Routes -}}
		switch prefix {
		{{range $_, $r := .Routes -}}
		case "{{.Prefix}}":
			return rt.route{{.Name}}(path, method)
		{{end -}}
		}
		{{- end}}

		{{- if .VariableRoute}}

		return rt.route{{.VariableRoute.Name}}(path, method)
		{{- end}}
	}
	{{- end}}

	return nil, ""
}
{{ end }}

type pathKey struct{}

func SchemaPath(r *http.Request) (string, bool) {
	if s, ok := r.Context().Value(pathKey{}).(string); ok {
		return s, true
	}
	return r.URL.Path, false
}

var specFileBs = []byte(SpecFile)

func SpecFileHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/{{.SpecFileExt}}")
		rw.WriteHeader(http.StatusOK)
		_, err := rw.Write(specFileBs)
		if err != nil {
			LogError(fmt.Errorf("serve spec file: %w", err))
		}
	})
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
`)

func (r RouterFile) Execute() (string, error) { return tmRouterFile.Execute(r) }

func (g *Generator) newRoutes() ([]routeMethod, error) {
	var route routeMethod
	for _, pi := range g.Paths {
		err := route.add(pi.PathItem.PathOld, pi)
		if err != nil {
			return nil, fmt.Errorf("add %q path: %w", pi.PathItem.Path.Spec, err)
		}
	}

	rs := route.routes()
	return rs, nil

	// mRoutes := make(map[string]*routeMethod)

	// for _, pi := range g.Paths {
	// 	path := pi.pathItem.Path
	// 	prefix, path, ok := path.Cut()
	// 	if !ok {
	// 		rm, ok := mRoutes[prefix.Name()]
	// 		if !ok {
	// 			rm = &routeMethod{
	// 				Prefix: prefix,
	// 			}
	// 			routes = append(routes, rm)
	// 			mRoutes[prefix.Name()] = rm
	// 		}
	// 		rm.
	// 	}
	// }
}

type routeMethod struct {
	Name     string
	Prefix   specification.Prefix
	BasePath string

	PrefixPathItems []*routePrefixPathItem
	mPathItems      map[string]*routePrefixPathItem

	PathItem *PathItem

	Routes  []*routeMethod
	mRoutes map[string]*routeMethod

	VariableRoute *routeMethod

	// PrefixHandlers []PrefixHandlers

	// Routes []*router
	// ptrs   map[specification.Prefix]*router
}

// func (r routeMethod) String() string {
// 	var sb strings.Builder
// 	fmt.Fprintf(&sb, "func %s | %s\n", r.Name, r.Prefix)
// 	for _, pi := range r.PrefixPathItems {
// 		fmt.Fprintf(&sb, "\t%q -> %v\n", pi.Prefix, pi.PathItem)
// 	}
// 	fmt.Fprintf(&sb, "~~~%v\n", r.PathItem)
// 	for _, rt := range r.Routes {
// 		fmt.Fprintf(&sb, "\t%v\n", rt)
// 	}
// 	if r.VariableRoute != nil {
// 		fmt.Fprintf(&sb, ">>> %v\n", r.VariableRoute)
// 	}
// 	return sb.String()
// }

type routePrefixPathItem struct {
	Prefix   specification.Prefix
	PathItem *PathItem
}

// func (r routePrefixPathItem) String() string {
// 	return fmt.Sprintf("routePrefixPathItem{%q\n\t%v}", r.Prefix, r.PathItem)
// }

// type routePrefixRouteMethod struct {
// 	Prefix specification.Prefix
// 	Route  routeMethod
// }

// func newrouter(prefix specification.Prefix) *router {
// 	return &router{
// 		ptrs: make(map[specification.Prefix]*router),
// 	}
// }

func (r *routeMethod) routes() []routeMethod {
	var out = []routeMethod{*r}
	for _, r := range r.Routes {
		out = append(out, r.routes()...)
		// log.Printf("routes state: %v", out)
	}
	if r.VariableRoute != nil {
		out = append(out, r.VariableRoute.routes()...)
		// log.Printf("routes vtate: %v", out)
	}
	return out
}

func (r *routeMethod) add(path specification.PathOld, pathItem *PathItem) error {
	prefix, path, ok := path.Cut()
	if !ok {
		if prefix.IsVariable() {
			r.PathItem = pathItem
		} else {
			if r.mPathItems == nil {
				r.mPathItems = make(map[string]*routePrefixPathItem)
			}
			if _, ok := r.mPathItems[prefix.Name()]; ok {
				return fmt.Errorf("%q prefix already exists", prefix)
			}
			pi := &routePrefixPathItem{
				Prefix:   prefix,
				PathItem: pathItem,
			}
			r.mPathItems[prefix.Name()] = pi
			r.PrefixPathItems = append(r.PrefixPathItems, pi)
		}
	} else {
		if prefix.IsVariable() {
			if r.VariableRoute == nil {
				rm := &routeMethod{
					Name:   r.Name + PublicFieldName(prefix.Name()),
					Prefix: prefix,
				}
				r.VariableRoute = rm
			}
			return r.VariableRoute.add(path, pathItem)
		} else {
			if r.mRoutes == nil {
				r.mRoutes = make(map[string]*routeMethod)
			}
			if route, ok := r.mRoutes[prefix.Name()]; ok {
				return route.add(path, pathItem)
			} else {
				route := &routeMethod{
					Name:   r.Name + PublicFieldName(prefix.Name()),
					Prefix: prefix,
				}
				r.mRoutes[prefix.Name()] = route
				r.Routes = append(r.Routes, route)
				return route.add(path, pathItem)
			}
		}
	}

	return nil
}
