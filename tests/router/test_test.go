package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	api := API{
		GetRTHandler: GetRTHandlerFunc(func(_ GetRTRequester) GetRTResponder { return GetRTResponseDefault(201) }),

		GetShopsHandler: GetShopsHandlerFunc(func(_ GetShopsRequester) GetShopsResponder { return GetShopsResponseDefault(202) }),

		GetShopsRTHandler: GetShopsRTHandlerFunc(func(_ GetShopsRTRequester) GetShopsRTResponder { return GetShopsRTResponseDefault(203) }),

		GetShopsShopHandler: GetShopsShopHandlerFunc(func(_ GetShopsShopRequester) GetShopsShopResponder { return GetShopsShopResponseDefault(204) }),

		GetShopsShopRTHandler: GetShopsShopRTHandlerFunc(func(_ GetShopsShopRTRequester) GetShopsShopRTResponder {
			return GetShopsShopRTResponseDefault(205)
		}),

		GetShopsShopPetsHandler: GetShopsShopPetsHandlerFunc(func(_ GetShopsShopPetsRequester) GetShopsShopPetsResponder {
			return GetShopsShopPetsResponseDefault(206)
		}),

		GetShopsActivateHandler: GetShopsActivateHandlerFunc(func(_ GetShopsActivateRequester) GetShopsActivateResponder {
			return GetShopsActivateResponseDefault(207)
		}),
	}

	for _, tt := range []struct {
		path       string
		code       int
		schemaPath string
	}{
		{"/", 201, "/"},
		{"/shops", 202, "/shops"},
		{"/shops/", 203, "/shops/"},
		{"/shops/my_shop", 204, "/shops/{shop}"},

		{"/shops/my_shop/", 205, "/shops/{shop}/"},
		{"/shops//", 205, "/shops/{shop}/"},

		{"/shops/my_shop/pets", 206, "/shops/{shop}/pets"},

		{"/shops/activate", 207, "/shops/activate"},

		{"/not_found", 404, "/this_is_not_gonna_be_checked"},
	} {
		tt := tt
		t.Run(tt.path, func(t *testing.T) {
			api := api
			api.Middlewares = append(api.Middlewares, func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					assert.Equal(t, tt.schemaPath, SchemaPath(r))
					h.ServeHTTP(rw, r)
				})
			})
			w := httptest.NewRecorder()
			path := "/api/v1" + tt.path
			api.ServeHTTP(w, httptest.NewRequest("GET", path, nil))
			assert.Equal(t, tt.code, w.Code, "path: %s", path)
		})
	}
}
