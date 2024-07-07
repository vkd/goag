package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type API struct {
	GetHandler                 GetHandlerFunc
	GetShopsHandler            GetShopsHandlerFunc
	GetShopsRTHandler          GetShopsRTHandlerFunc
	GetShopsActivateHandler    GetShopsActivateHandlerFunc
	GetShopsActivateRTHandler  GetShopsActivateRTHandlerFunc
	GetShopsActivateTagHandler GetShopsActivateTagHandlerFunc
	GetShopsShopHandler        GetShopsShopHandlerFunc
	GetShopsShopRTHandler      GetShopsShopRTHandlerFunc
	GetShopsShopPetsHandler    GetShopsShopPetsHandlerFunc
	ReviewShopHandler          ReviewShopHandlerFunc

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
		case "/":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetHandler)
				return h, "/", true
			}
		case "/shops":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsHandler)
				return h, "/shops", true
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
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsRTHandler)
				return h, "/shops/", true
			}
		case "/activate":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsActivateHandler)
				return h, "/shops/activate", true
			}
		}
		switch method {
		case http.MethodGet:
			h := http.Handler(rt.GetShopsShopHandler)
			return h, "/shops/{shop}", true
		}
		return nil, "", false
	}

	switch prefix {
	case "/activate":
		h, out, hasPath := rt.routeShopsActivate(path, method)
		if h != nil {
			return h, out, hasPath
		}
	}

	return rt.routeShopsShop(path, method)
}

func (rt *API) routeShopsActivate(path, method string) (http.Handler, string, bool) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsActivateRTHandler)
				return h, "/shops/activate/", true
			}
		case "/tag":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsActivateTagHandler)
				return h, "/shops/activate/tag", true
			}
		}
		return nil, "", false
	}

	return nil, "", false
}

func (rt *API) routeShopsShop(path, method string) (http.Handler, string, bool) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsShopRTHandler)
				return h, "/shops/{shop}/", true
			}
		case "/pets":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsShopPetsHandler)
				return h, "/shops/{shop}/pets", true
			}
		case "/review":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.ReviewShopHandler)
				return h, "/shops/{shop}/review", true
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
