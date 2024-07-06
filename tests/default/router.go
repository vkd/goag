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
		case "/":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetHandler)
				return h, "/"
			}
		case "/shops":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsHandler)
				return h, "/shops"
			}
		}
		return nil, ""
	}

	switch prefix {
	case "/shops":
		return rt.routeShops(path, method)
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
				h := http.Handler(rt.GetShopsRTHandler)
				return h, "/shops/"
			}
		case "/activate":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsActivateHandler)
				return h, "/shops/activate"
			}
		}
		switch method {
		case http.MethodGet:
			h := http.Handler(rt.GetShopsShopHandler)
			return h, "/shops/{shop}"
		}
		return nil, ""
	}

	switch prefix {
	case "/activate":
		h, out := rt.routeShopsActivate(path, method)
		if h != nil {
			return h, out
		}
	}

	return rt.routeShopsShop(path, method)
}

func (rt *API) routeShopsActivate(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsActivateRTHandler)
				return h, "/shops/activate/"
			}
		case "/tag":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsActivateTagHandler)
				return h, "/shops/activate/tag"
			}
		}
		return nil, ""
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
				h := http.Handler(rt.GetShopsShopRTHandler)
				return h, "/shops/{shop}/"
			}
		case "/pets":
			switch method {
			case http.MethodGet:
				h := http.Handler(rt.GetShopsShopPetsHandler)
				return h, "/shops/{shop}/pets"
			}
		case "/review":
			switch method {
			case http.MethodPost:
				h := http.Handler(rt.ReviewShopHandler)
				return h, "/shops/{shop}/review"
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
