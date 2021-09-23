package test

import (
	"context"
	"net/http"
	"strings"
)

type API struct {
	GetPetsIDsHandler GetPetsIDsHandlerFunc

	// not found
	NotFoundHandler http.Handler

	Middlewares []func(h http.Handler) http.Handler
}

func (rt *API) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	h, path := rt.route(path, r.Method)
	if h == nil {
		h = rt.NotFoundHandler
		if h == nil {
			h = http.NotFoundHandler()
		}
		h.ServeHTTP(rw, r)
		return
	}

	for _, m := range rt.Middlewares {
		h = m(h)
	}
	r = r.WithContext(context.WithValue(r.Context(), pathKey{}, path))
	h.ServeHTTP(rw, r)
}

func (rt *API) route(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/pets":
			return rt.routePets(path, method)
		}
	}

	return nil, ""
}

func (rt *API) routePets(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/ids":
			switch method {
			case http.MethodGet:
				return rt.GetPetsIDsHandler, "/pets/ids"
			}
		}
	}

	return nil, ""
}

type pathKey struct{}

func SchemaPath(r *http.Request) (string, bool) {
	if s, ok := r.Context().Value(pathKey{}).(string); ok {
		return s, true
	}
	return r.URL.Path, false
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
