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
		case "/login":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.PostLoginHandler)
				return h, "/login", true
			}
		case "/shops":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.PostShopsHandler)
				h = middlewares(h, authMiddlewareOr(
					rt.SecurityBearerAuth,
				))
				return h, "/shops", true
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

type Middleware interface {
	Middleware(http.Handler) http.Handler
}

type MiddlewareFunc func(http.Handler) http.Handler

func (m MiddlewareFunc) Middleware(next http.Handler) http.Handler {
	return m(next)
}

func middlewares(h http.Handler, ms ...Middleware) http.Handler {
	for i := len(ms) - 1; i >= 0; i-- {
		h = ms[i].Middleware(h)
	}
	return h
}

func authMiddlewareOr(fns ...AuthMiddleware) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, fn := range fns {
				if fn == nil {
					continue
				}
				authReq, ok := fn.Auth(r)
				if ok {
					next.ServeHTTP(w, authReq)
					return
				}
			}
			w.WriteHeader(401)
		})
	}
}

type AuthMiddleware interface {
	Auth(r *http.Request) (*http.Request, bool)
}

type AuthMiddlewareFunc func(*http.Request) (*http.Request, bool)

type SecurityBearerAuthMiddleware func(r *http.Request, token string) (*http.Request, bool)

func (s SecurityBearerAuthMiddleware) Auth(r *http.Request) (*http.Request, bool) {
	var token string
	hs := r.Header.Values("Authorization")
	if len(hs) == 0 {
		return nil, false
	}
	token = hs[0]

	token = strings.TrimPrefix(token, "Bearer ")

	return s(r, token)
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
