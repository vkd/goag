package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type API struct {
	PostShopsShopStringSepShopSchemaPetsHandler PostShopsShopStringSepShopSchemaPetsHandlerFunc

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

	h, path := rt.route(path, r.Method)
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
	r = r.WithContext(context.WithValue(r.Context(), pathKey{}, path))
	h.ServeHTTP(rw, r)
}

func (rt *API) route(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/shops":
			return rt.routeShops(path, method)
		}
	}

	return nil, ""
}

func (rt *API) routeShops(path, method string) (http.Handler, string) {
	_, path = splitPath(path)

	if path != "" {

		return rt.routeShopsShopString(path, method)
	}

	return nil, ""
}

func (rt *API) routeShopsShopString(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/sep":
			return rt.routeShopsShopStringSep(path, method)
		}
	}

	return nil, ""
}

func (rt *API) routeShopsShopStringSep(path, method string) (http.Handler, string) {
	_, path = splitPath(path)

	if path != "" {

		return rt.routeShopsShopStringSepShopSchema(path, method)
	}

	return nil, ""
}

func (rt *API) routeShopsShopStringSepShopSchema(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/pets":
			switch method {
			case http.MethodPost:
				return rt.PostShopsShopStringSepShopSchemaPetsHandler, "/shops/{shop_string}/sep/{shop_schema}/pets"
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