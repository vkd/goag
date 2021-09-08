package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	api := API{
		GetRTHandler: GetRTHandlerFunc(func(gr GetRTParamsParser) GetRTResponser { return GetRTResponseDefault(201) }),

		GetShopsHandler: GetShopsHandlerFunc(func(gr GetShopsParamsParser) GetShopsResponser { return GetShopsResponseDefault(202) }),

		GetShopsRTHandler: GetShopsRTHandlerFunc(func(gr GetShopsRTParamsParser) GetShopsRTResponser { return GetShopsRTResponseDefault(203) }),

		GetShopsShopHandler: GetShopsShopHandlerFunc(func(gssp GetShopsShopParamsParser) GetShopsShopResponser { return GetShopsShopResponseDefault(204) }),

		GetShopsShopRTHandler: GetShopsShopRTHandlerFunc(func(gssr GetShopsShopRTParamsParser) GetShopsShopRTResponser {
			return GetShopsShopRTResponseDefault(205)
		}),

		GetShopsShopPetsHandler: GetShopsShopPetsHandlerFunc(func(gsspp GetShopsShopPetsParamsParser) GetShopsShopPetsResponser {
			return GetShopsShopPetsResponseDefault(206)
		}),

		GetShopsActivateHandler: GetShopsActivateHandlerFunc(func(gsap GetShopsActivateParamsParser) GetShopsActivateResponser {
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
