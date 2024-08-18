package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type API struct {
	PostPetsHandler  PostPetsHandlerFunc
	PostPets2Handler PostPets2HandlerFunc

	// not found
	NotFoundHandler http.Handler
	// spec file
	SpecFileHandler http.Handler

	Middlewares []func(h http.Handler) http.Handler
}

func (rt *API) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if rt.SpecFileHandler != nil && path == "/openapi.yaml" {
		rt.SpecFileHandler.ServeHTTP(rw, r)
		return
	}

	h, path, hasPath := rt.route(path, r.Method)
	if h == nil {
		h = rt.NotFoundHandler
		if h == nil {
			h = http.NotFoundHandler()
		}

		hasPath = false
	}

	if hasPath {
		r = r.WithContext(context.WithValue(r.Context(), pathKey{}, path))

		for i := len(rt.Middlewares) - 1; i >= 0; i-- {
			h = rt.Middlewares[i](h)
		}
	}

	h.ServeHTTP(rw, r)
}

func (rt *API) route(path, method string) (http.Handler, string, bool) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/pets":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.PostPetsHandler)
				return h, "/pets", true
			}
		case "/pets2":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.PostPets2Handler)
				return h, "/pets2", true
			}
		}
		return nil, "", false
	}

	return nil, "", false
}

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
		rw.Header().Set("Content-Type", "application/yaml")
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
