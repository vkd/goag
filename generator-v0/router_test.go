package generator

import (
	"fmt"
	"log"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"
	"github.com/vkd/goag/specification"
)

func MustHandler(method, path string) Handler {
	h, err := NewHandler(&openapi3.Operation{}, specification.Path(path), method, nil)
	if err != nil {
		panic(fmt.Errorf("new handler %q: %w", path, err))
	}
	return h
}

func TestNewRoutes(t *testing.T) {
	rs, err := NewRoutes("", []Handler{
		MustHandler("GET", "/"),
		MustHandler("GET", "/pets"),
		MustHandler("POST", "/pets"),
		MustHandler("POST", "/activate"),
		MustHandler("GET", "/pets/{petId}"),
		MustHandler("DELETE", "/pets/{petId}"),
		MustHandler("POST", "/pets/{petId}/buy"),
		MustHandler("POST", "/pets/{petId}/rent"),
		// MustHandler("GET", "/shops"),
		// MustHandler("POST", "/shops"),
		// MustHandler("GET", "/shops/{shopId}"),
		// MustHandler("POST", "/shops/{shopId}/coupons/activate"),
	})
	require.NoError(t, err)

	for _, r := range rs {
		print(r)
	}
}

func print(r Route) {
	log.Printf("Route: %q", r.Name)
	for _, h := range r.Handlers {
		log.Printf("Handler: %q %v", h.Prefix, h.Methods)
	}
	if r.WildcardHandler != nil {
		log.Printf("Wildcard handler: %v", r.WildcardHandler.Methods)
	}
	for _, route := range r.Routes {
		log.Printf("ReRoute: %q -> %q", route.Prefix, route.RouteName)
	}
	if r.WildcardRouteName != "" {
		log.Printf("Wildcard router: %q", r.WildcardRouteName)
	}
}
