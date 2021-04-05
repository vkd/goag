package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	api := API{
		GetRTHandler: NewGetRTHandlerer(func(gr GetRTParams) GetRTResponser { return GetRTResponseDefault(201) }),

		GetShopsHandler: NewGetShopsHandlerer(func(gr GetShopsParams) GetShopsResponser { return GetShopsResponseDefault(202) }),

		GetShopsRTHandler: NewGetShopsRTHandlerer(func(gr GetShopsRTParams) GetShopsRTResponser { return GetShopsRTResponseDefault(203) }),

		GetShopsShopHandler: NewGetShopsShopHandlerer(func(gssp GetShopsShopParams) GetShopsShopResponser { return GetShopsShopResponseDefault(204) }),

		GetShopsShopRTHandler: NewGetShopsShopRTHandlerer(func(gssr GetShopsShopRTParams) GetShopsShopRTResponser { return GetShopsShopRTResponseDefault(205) }),

		GetShopsShopPetsHandler: NewGetShopsShopPetsHandlerer(func(gsspp GetShopsShopPetsParams) GetShopsShopPetsResponser {
			return GetShopsShopPetsResponseDefault(206)
		}),

		GetShopsActivateHandler: NewGetShopsActivateHandlerer(func(gsap GetShopsActivateParams) GetShopsActivateResponser {
			return GetShopsActivateResponseDefault(207)
		}),
	}

	router := api.Router()

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
			router.ServeHTTP(w, httptest.NewRequest("GET", tt.path, nil))
			assert.Equal(t, tt.code, w.Code)
		})
	}
}
