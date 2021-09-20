package test

import (
	"context"
	"net/http"
	"strings"
)

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

func SchemaPath(r *http.Request) string {
	if s, ok := r.Context().Value(pathKey{}).(string); ok {
		return s
	}
	return r.URL.Path
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
