package test

import (
	"net/http"
	"strings"
)

type API struct {
	GetPetsPetIDHandler GetPetsPetIDHandlerer

	// not found
	NotFoundHandler http.Handler
}

func (a API) Router() http.Handler {
	r := router{
		GetPetsPetIDHandler: GetPetsPetIDHandler(a.GetPetsPetIDHandler),

		NotFoundHandler: a.NotFoundHandler,
	}
	if r.NotFoundHandler == nil {
		r.NotFoundHandler = http.NotFoundHandler()
	}
	return &r
}

type router struct {
	GetPetsPetIDHandler http.Handler

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
		case "/pets":
			return rt.routePets(path, method)
		}
	}

	return nil
}

func (rt *router) routePets(path, method string) http.Handler {
	_, path = splitPath(path)

	if path == "" {

		switch method {
		case http.MethodGet:
			return rt.GetPetsPetIDHandler
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
