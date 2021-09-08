package test

import (
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

func (rt *API) routeShops(path, method string) http.Handler {
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

func (rt *API) routeShopsShop(path, method string) http.Handler {
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
