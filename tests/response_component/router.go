package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type API struct {
	GetPetHandler   GetPetHandlerFunc
	GetV2PetHandler GetV2PetHandlerFunc
	GetV3PetHandler GetV3PetHandlerFunc

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
		case "/pet":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetPetHandler)
				return h, "/pet", true
			}
		}
		return nil, "", false
	}

	switch prefix {
	case "/v2":
		return rt.routeV2(path, method)
	case "/v3":
		return rt.routeV3(path, method)
	}
	return nil, "", false
}

func (rt *API) routeV2(path, method string) (http.Handler, string, bool) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/pet":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetV2PetHandler)
				return h, "/v2/pet", true
			}
		}
		return nil, "", false
	}

	return nil, "", false
}

func (rt *API) routeV3(path, method string) (http.Handler, string, bool) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/pet":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetV3PetHandler)
				return h, "/v3/pet", true
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
