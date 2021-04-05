package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMultiParams(t *testing.T) {
	handler := NewGetShopsShopPetsPetIDHandlerer(
		func(ps GetShopsShopPetsPetIDParams) GetShopsShopPetsPetIDResponser {
			assert.Equal(t, "paw", ps.Shop)
			assert.Equal(t, int64(1), ps.PetID)
			assert.Equal(t, "white", ps.Color)
			assert.Equal(t, int32(1), ps.Page)
			return GetShopsShopPetsPetIDResponse200()
		},
		func(_ error) GetShopsShopPetsPetIDResponser {
			return GetShopsShopPetsPetIDResponseDefault(400)
		})

	w := httptest.NewRecorder()
	API{GetShopsShopPetsPetIDHandler: handler}.Router().ServeHTTP(w, httptest.NewRequest("GET", "/shops/paw/pets/1?color=white&page=1", nil))

	assert.Equal(t, 200, w.Code)
}
