package test

import (
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
		path string
		code int
	}{
		{"/", 201},
		{"/shops", 202},
		{"/shops/", 203},
		{"/shops/my_shop", 204},

		{"/shops/my_shop/", 205},
		{"/shops//", 205},

		{"/shops/my_shop/pets", 206},

		{"/shops/activate", 207},

		{"/not_found", 404},
	} {
		tt := tt
		t.Run(tt.path, func(t *testing.T) {
			w := httptest.NewRecorder()
			api.ServeHTTP(w, httptest.NewRequest("GET", tt.path, nil))
			assert.Equal(t, tt.code, w.Code)
		})
	}
}
