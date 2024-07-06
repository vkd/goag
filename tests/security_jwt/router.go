package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type API struct {
	PostLoginHandler PostLoginHandlerFunc
	PostShopsHandler PostShopsHandlerFunc

	// not found
	NotFoundHandler http.Handler
	// spec file
	SpecFileHandler http.Handler

	Middlewares []func(h http.Handler) http.Handler

	SecurityBearerAuth SecurityBearerAuthMiddleware
}

func (rt *API) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if rt.SpecFileHandler != nil && path == "/openapi.yaml" {
		rt.SpecFileHandler.ServeHTTP(rw, r)
		return
	}

	h, r := rt.Route(r)
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
	h.ServeHTTP(rw, r)
}

func (rt *API) Route(r *http.Request) (http.Handler, *http.Request) {
	h, path := rt.route(r.URL.Path, r.Method)
	if h == nil {
		return nil, r
	}

	r = r.WithContext(context.WithValue(r.Context(), pathKey{}, path))
	return h, r
}

func (rt *API) route(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/login":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.PostLoginHandler)
				return h, "/login"
			}
		case "/shops":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.PostShopsHandler)
				h = middlewares(h, rt.SecurityBearerAuth)
				return h, "/shops"
			}
		}
		return nil, ""
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

func middlewares(h http.Handler, ms ...interface {
	Middleware(http.Handler) http.Handler
}) http.Handler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i].Middleware(h)
	}
	return h
}

type SecurityBearerAuthMiddleware func(w http.ResponseWriter, r *http.Request, token string, next http.Handler)

func (m SecurityBearerAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string
		hs := r.Header.Values("Authorization")
		if len(hs) > 0 {
			token = strings.TrimPrefix(hs[0], "Bearer ")
		}

		m(w, r, token, next)
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
