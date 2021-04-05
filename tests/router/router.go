package test

import (
	"net/http"
	"strings"
)

type API struct {
	GetRTHandler            GetRTHandlerer
	GetShopsHandler         GetShopsHandlerer
	GetShopsRTHandler       GetShopsRTHandlerer
	GetShopsActivateHandler GetShopsActivateHandlerer
	GetShopsShopHandler     GetShopsShopHandlerer
	GetShopsShopRTHandler   GetShopsShopRTHandlerer
	GetShopsShopPetsHandler GetShopsShopPetsHandlerer

	// not found
	NotFoundHandler http.Handler
}

func (a API) Router() http.Handler {
	r := router{
		GetRTHandler:            GetRTHandler(a.GetRTHandler),
		GetShopsHandler:         GetShopsHandler(a.GetShopsHandler),
		GetShopsRTHandler:       GetShopsRTHandler(a.GetShopsRTHandler),
		GetShopsActivateHandler: GetShopsActivateHandler(a.GetShopsActivateHandler),
		GetShopsShopHandler:     GetShopsShopHandler(a.GetShopsShopHandler),
		GetShopsShopRTHandler:   GetShopsShopRTHandler(a.GetShopsShopRTHandler),
		GetShopsShopPetsHandler: GetShopsShopPetsHandler(a.GetShopsShopPetsHandler),

		NotFoundHandler: a.NotFoundHandler,
	}
	if r.NotFoundHandler == nil {
		r.NotFoundHandler = http.NotFoundHandler()
	}
	return &r
}

type router struct {
	GetRTHandler            http.Handler
	GetShopsHandler         http.Handler
	GetShopsRTHandler       http.Handler
	GetShopsActivateHandler http.Handler
	GetShopsShopHandler     http.Handler
	GetShopsShopRTHandler   http.Handler
	GetShopsShopPetsHandler http.Handler

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
		case "/":
			switch method {
			case http.MethodGet:
				return rt.GetRTHandler
			}
		case "/shops":
			switch method {
			case http.MethodGet:
				return rt.GetShopsHandler
			}
		}
	}

	if path != "" {
		switch prefix {
		case "/shops":
			return rt.routeShops(path, method)
		}
	}

	return nil
}

func (rt *router) routeShops(path, method string) http.Handler {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				return rt.GetShopsRTHandler
			}
		case "/activate":
			switch method {
			case http.MethodGet:
				return rt.GetShopsActivateHandler
			}
		}
		switch method {
		case http.MethodGet:
			return rt.GetShopsShopHandler
		}
	}

	if path != "" {

		return rt.routeShopsShop(path, method)
	}

	return nil
}

func (rt *router) routeShopsShop(path, method string) http.Handler {
	prefix, path := splitPath(path)

	if path == "" {
		switch prefix {
		case "/":
			switch method {
			case http.MethodGet:
				return rt.GetShopsShopRTHandler
			}
		case "/pets":
			switch method {
			case http.MethodGet:
				return rt.GetShopsShopPetsHandler
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
