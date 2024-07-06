package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorsDefault(t *testing.T) {
	testShop := "paw"
	testPage := int32(2)
	testRequestID := "abcdef"

	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			req, err := r.Parse()
			if err != nil {
				return NewGetShopsShopResponseDefault(400)
			}
			assert.Equal(t, testShop, req.Path.Shop)
			assert.Equal(t, Just(testPage), req.Query.Page)
			assert.Equal(t, Just(testRequestID), req.Headers.RequestID)
			return NewGetShopsShopResponse200()
		},

		CORSHandler: func(methods, headers []string) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNoContent)
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ", "))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ", "))
			})
		},
	}

	// /shops/{shop}
	target := fmt.Sprintf("/shops/%s", testShop)
	req := httptest.NewRequest(http.MethodOptions, target, nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "GET, POST", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Request-Id, Query-Id", w.Header().Get("Access-Control-Allow-Headers"))

	// /shops
	target = "/shops"
	req = httptest.NewRequest(http.MethodOptions, target, nil)
	w = httptest.NewRecorder()
	api.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, "GET", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Access-Key", w.Header().Get("Access-Control-Allow-Headers"))
}
