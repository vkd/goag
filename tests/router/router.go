package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const SpecFile string = `servers:
  - url: /api/v1
paths:
  /: {get: {responses: {default: {}}}}
  /shops: {get: {responses: {default: {}}}}
  /shops/: {get: {responses: {default: {}}}}
  /shops/{shop}: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/{shop}/: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/{shop}/pets: {get: {parameters: [{in: path, name: shop, required: true, schema: {type: string}}], responses: {default: {}}}}
  /shops/activate: {get: {responses: {default: {}}}}
`

type API struct {
	GetRTHandler            GetRTHandlerFunc
	GetShopsHandler         GetShopsHandlerFunc
	GetShopsRTHandler       GetShopsRTHandlerFunc
	GetShopsActivateHandler GetShopsActivateHandlerFunc
	GetShopsShopHandler     GetShopsShopHandlerFunc
	GetShopsShopRTHandler   GetShopsShopRTHandlerFunc
	GetShopsShopPetsHandler GetShopsShopPetsHandlerFunc

	// not found
	NotFoundHandler http.Handler
	// spec file
	SpecFileHandler http.Handler

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
	if !strings.HasPrefix(path, "/api/v1") {
		return nil, ""
	}
	path = path[7:] // "/api/v1"

	if !strings.HasPrefix(path, "/") {
		return nil, ""
	}

	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				return rt.GetRTHandler, "/"
			}
		case "/shops":
			switch method {
			case http.MethodGet:
				return rt.GetShopsHandler, "/shops"
			}
		case "/openapi.yaml":
			switch method {
			case http.MethodGet:
				return rt.SpecFileHandler, "/openapi.yaml"
			}
		}
	}

	if path != "" {
		switch prefix {
		case "/shops":
			return rt.routeShops(path, method)
		}
	}

	return nil, ""
}

func (rt *API) routeShops(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				return rt.GetShopsRTHandler, "/shops/"
			}
		case "/activate":
			switch method {
			case http.MethodGet:
				return rt.GetShopsActivateHandler, "/shops/activate"
			}
		}
		switch method {
		case http.MethodGet:
			return rt.GetShopsShopHandler, "/shops/{shop}"
		}
	}

	if path != "" {

		return rt.routeShopsShop(path, method)
	}

	return nil, ""
}

func (rt *API) routeShopsShop(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				return rt.GetShopsShopRTHandler, "/shops/{shop}/"
			}
		case "/pets":
			switch method {
			case http.MethodGet:
				return rt.GetShopsShopPetsHandler, "/shops/{shop}/pets"
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

var specFileBs = []byte(SpecFile)

func SpecFileHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/")
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
