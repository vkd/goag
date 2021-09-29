package test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const SpecFile string = `paths:
  /shops/{shop}/pets/{petId}:
    get:
      parameters:
        - name: shop
          in: path
          required: true
          schema:
            type: string
        - name: petId
          in: path
          required: true
          schema:
            type: integer
            format: int64
        - name: color
          in: query
          required: true
          schema:
            type: string
        - name: page
          in: query
          schema:
            type: integer
            format: int32
      responses:
        '200': {}
        default: {}
`

type API struct {
	GetShopsShopPetsPetIDHandler GetShopsShopPetsPetIDHandlerFunc

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

	for i := len(rt.Middlewares) - 1; i >= 0; i-- {
		h = rt.Middlewares[i](h)
	}
	r = r.WithContext(context.WithValue(r.Context(), pathKey{}, path))
	h.ServeHTTP(rw, r)
}

func (rt *API) route(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
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
	_, path = splitPath(path)

	if path != "" {

		return rt.routeShopsShop(path, method)
	}

	return nil, ""
}

func (rt *API) routeShopsShop(path, method string) (http.Handler, string) {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/pets":
			return rt.routeShopsShopPets(path, method)
		}
	}

	return nil, ""
}

func (rt *API) routeShopsShopPets(path, method string) (http.Handler, string) {
	_, path = splitPath(path)

	if path == "" {

		switch method {
		case http.MethodGet:
			return rt.GetShopsShopPetsPetIDHandler, "/shops/{shop}/pets/{petId}"
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
