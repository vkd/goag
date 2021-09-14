package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMultiParams(t *testing.T) {
	testShop := "paw"
	testPetID := int64(1)
	testColor := "white"
	testPage := int32(2)

	api := API{
		GetShopsShopPetsPetIDHandler: func(r GetShopsShopPetsPetIDRequester) GetShopsShopPetsPetIDResponder {
			req, err := r.Parse()
			if err != nil {
				return GetShopsShopPetsPetIDResponseDefault(400)
			}
			assert.Equal(t, testShop, req.Shop)
			assert.Equal(t, testPetID, req.PetID)
			assert.Equal(t, testColor, req.Color)
			assert.Equal(t, testPage, *req.Page)
			return GetShopsShopPetsPetIDResponse200()
		},
	}

	target := fmt.Sprintf("/shops/%s/pets/%d?color=%s&page=%d", testShop, testPetID, testColor, testPage)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", target, nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetMultiParams_Optional(t *testing.T) {
	api := API{
		GetShopsShopPetsPetIDHandler: func(r GetShopsShopPetsPetIDRequester) GetShopsShopPetsPetIDResponder {
			req, err := r.Parse()
			if err != nil {
				return GetShopsShopPetsPetIDResponseDefault(400)
			}
			assert.Nil(t, req.Page)
			return GetShopsShopPetsPetIDResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/shops/paw/pets/1?color=white", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetMultiParams_BadRequest(t *testing.T) {
	api := API{
		GetShopsShopPetsPetIDHandler: func(r GetShopsShopPetsPetIDRequester) GetShopsShopPetsPetIDResponder {
			_, err := r.Parse()
			if err != nil {
				return GetShopsShopPetsPetIDResponseDefault(400)
			}
			return GetShopsShopPetsPetIDResponse200()
		},
	}

	for _, target := range []string{
		// "/shops/paw/pets/1?color=white&page=2",
		"/shops/paw/pets/a?color=white&page=2",
		"/shops/paw/pets/1?color=white&page=b",
		"/shops//pets/1?color=white&page=2",
		"/shops/paw/pets/?color=white&page=2",
		"/shops/paw/pets/1?page=2",
	} {
		w := httptest.NewRecorder()
		api.ServeHTTP(w, httptest.NewRequest("GET", target, nil))

		assert.Equal(t, 400, w.Code, "suppose to be a bad request: %s", target)
	}
}
