package test

import (
	"net/http"
	"strings"
)

type API struct {
	GetShopsShopPetsPetIDHandler GetShopsShopPetsPetIDHandlerer

	// not found
	NotFoundHandler http.Handler
}

func (a API) Router() http.Handler {
	r := router{
		GetShopsShopPetsPetIDHandler: GetShopsShopPetsPetIDHandler(a.GetShopsShopPetsPetIDHandler),

		NotFoundHandler: a.NotFoundHandler,
	}
	if r.NotFoundHandler == nil {
		r.NotFoundHandler = http.NotFoundHandler()
	}
	return &r
}

type router struct {
	GetShopsShopPetsPetIDHandler http.Handler

	// not found
	NotFoundHandler http.Handler
}

func (rt *router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	h := rt.route(path, r.Method)
	if h == nil {
		h = rt.NotFoundHandler
	}
	h.ServeHTTP(rw, r)
}

func (rt *router) route(path, method string) http.Handler {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/shops":
			return rt.routeShops(path, method)
		}
	}

	return nil
}

func (rt *router) routeShops(path, method string) http.Handler {
	_, path = splitPath(path)

	if path != "" {

		return rt.routeShopsShop(path, method)
	}

	return nil
}

func (rt *router) routeShopsShop(path, method string) http.Handler {
	prefix, path := splitPath(path)

	if path != "" {
		switch prefix {
		case "/pets":
			return rt.routeShopsShopPets(path, method)
		}
	}

	return nil
}

func (rt *router) routeShopsShopPets(path, method string) http.Handler {
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
