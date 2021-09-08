package test

import (
	"net/http"
	"strings"
)

type API struct {
	GetShopsShopPetsPetIDHandler GetShopsShopPetsPetIDHandlerFunc

	// not found
	NotFoundHandler http.Handler
}

func (rt *API) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	h := rt.route(path, r.Method)
	if h == nil {
		h = rt.NotFoundHandler
		if h == nil {
			h = http.NotFoundHandler()
		}
	}
	h.ServeHTTP(rw, r)
}

func (rt *API) route(path, method string) http.Handler {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/shops":
			return rt.routeShops(path, method)
		}
	}

	return nil
}

func (rt *API) routeShops(path, method string) http.Handler {
	_, path = splitPath(path)

	if path != "" {

		return rt.routeShopsShop(path, method)
	}

	return nil
}

func (rt *API) routeShopsShop(path, method string) http.Handler {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/pets":
			return rt.routeShopsShopPets(path, method)
		}
	}

	return nil
}

func (rt *API) routeShopsShopPets(path, method string) http.Handler {
	_, path = splitPath(path)

	if path == "" {

		switch method {
		case http.MethodGet:
			return rt.GetShopsShopPetsPetIDHandler
		}
	}

	return nil
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
