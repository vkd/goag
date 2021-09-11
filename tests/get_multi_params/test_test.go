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
		GetShopsShopPetsPetIDHandler: func(r GetShopsShopPetsPetIDRequester) GetShopsShopPetsPetIDResponser {
			req, err := r.Parse()
			if err != nil {
				return GetShopsShopPetsPetIDResponseDefault(400)
			}
			assert.Equal(t, testShop, req.Shop)
			assert.Equal(t, testPetID, req.PetID)
			assert.Equal(t, testColor, *req.Color)
			assert.Equal(t, testPage, *req.Page)
			return GetShopsShopPetsPetIDResponse200()
		},
	}

	target := fmt.Sprintf("/shops/%s/pets/%d?color=%s&page=%d", testShop, testPetID, testColor, testPage)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", target, nil))

	assert.Equal(t, 200, w.Code)
}
