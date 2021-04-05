package test

import (
	"net/http"
	"strings"
)

type API struct {
	PostPetsHandler PostPetsHandlerer

	// not found
	NotFoundHandler http.Handler
}

func (a API) Router() http.Handler {
	r := router{
		PostPetsHandler: PostPetsHandler(a.PostPetsHandler),

		NotFoundHandler: a.NotFoundHandler,
	}
	if r.NotFoundHandler == nil {
		r.NotFoundHandler = http.NotFoundHandler()
	}
	return &r
}

type router struct {
	PostPetsHandler http.Handler

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

	if path == "" {
		switch prefix {
		case "/pets":
			switch method {
			case http.MethodPost:
				return rt.PostPetsHandler
			}
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
