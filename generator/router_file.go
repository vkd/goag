package generator

import "github.com/vkd/goag/specification"

type RouterFile struct {
	Handlers []Handler
	Data     any
}

func NewRouterFile(spec *specification.Spec, handlers []Handler, oldRouter any) RouterFile {
	return RouterFile{
		Handlers: handlers,
		Data:     oldRouter,
	}
}

var tmRouterFile = InitTemplate("RouterFile", `
{{- $router := .Data}}

type API struct {
	{{range $_, $h := .Handlers}}
	{{$h.Name}}Handler {{$h.HandlerFuncName}}
	{{- end}}

	// not found
	NotFoundHandler http.Handler
	// spec file
	SpecFileHandler http.Handler

	Middlewares []func(h http.Handler) http.Handler
}

func (rt *API) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if rt.SpecFileHandler != nil && path == "{{$router.BasePath}}/{{.Data.BaseSpecFilename}}" {
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

{{ range $_, $r := .Data.Routes }}
func (rt *API) route{{$r.Name}}(path, method string) (http.Handler, string) {
	{{- $basePath := $router.BasePath }}
	{{- if and $basePath (not $r.Name) }}
	if !strings.HasPrefix(path, "{{$basePath}}") {
		return nil, ""
	}
	path = path[{{len $basePath }}:] // "{{$basePath}}"

	if !strings.HasPrefix(path, "/") {
		return nil, ""
	}

	{{ end -}}

	{{if or .Handlers .Routes}}prefix, path :{{else}}_, path {{end -}}= splitPath(path)

	{{if or .Handlers .WildcardHandler -}}
	if path == "" {
		{{if .Handlers -}}
		switch prefix {
			{{range $_, $h := .Handlers -}}
		case "{{.Prefix}}":
			switch method {
				{{- range $_, $m := .Methods}}
			case http.Method{{.Method}}:
				return rt.{{.HandlerName}}Handler, "{{.Path}}"
				{{- end}}
			}
			{{end -}}
		}
		{{- end}}
		{{- if .WildcardHandler}}
		switch method {
			{{- range $_, $m := .WildcardHandler.Methods}}
		case http.Method{{.Method}}:
			return rt.{{.HandlerName}}Handler, "{{.Path}}"
			{{- end}}
		}
		{{- end}}
	}
	{{- end}}

	{{if or .Routes .WildcardRouteName -}}
	if path != "" {
		{{if .Routes -}}
		switch prefix {
		{{range $_, $r := .Routes -}}
		case "{{.Prefix}}":
			return rt.route{{.RouteName}}(path, method)
		{{end -}}
		}
		{{- end}}

		{{- if .WildcardRouteName}}

		return rt.route{{.WildcardRouteName}}(path, method)
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
		rw.Header().Set("Content-Type", "application/{{.Data.SpecFileExt}}")
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
