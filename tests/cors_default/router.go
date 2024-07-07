package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type API struct {
	GetShopsHandler      GetShopsHandlerFunc
	GetShopsShopHandler  GetShopsShopHandlerFunc
	PostShopsShopHandler PostShopsShopHandlerFunc

	// not found
	NotFoundHandler http.Handler
	// spec file
	SpecFileHandler http.Handler
	CORSHandler     CorsHandlerFunc

	Middlewares []func(h http.Handler) http.Handler
}

type CorsHandlerFunc func(methods, headers []string) http.Handler

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
		case "/shops":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsHandler)
				return h, "/shops", true
			case http.MethodOptions:
				if rt.CORSHandler == nil {
					return nil, "", false
				}
				h := rt.CORSHandler([]string{"GET"}, []string{"Access-Key"})
				return h, "", false
			}
		}
		return nil, "", false
	}

	switch prefix {
	case "/shops":
		return rt.routeShops(path, method)
	}
	return nil, "", false
}

func (rt *API) routeShops(path, method string) (http.Handler, string, bool) {
	_, path = splitPath(path)

	if path == "" {

		switch method {
		case http.MethodGet:
			h := http.Handler(rt.GetShopsShopHandler)
			return h, "/shops/{shop}", true
		case http.MethodOptions:
			if rt.CORSHandler == nil {
				return nil, "", false
			}
			h := rt.CORSHandler([]string{"GET", "POST"}, []string{"Request-Id", "Query-Id"})
			return h, "", false
		case http.MethodPost:
			h := http.Handler(rt.PostShopsShopHandler)
			return h, "/shops/{shop}", true
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
