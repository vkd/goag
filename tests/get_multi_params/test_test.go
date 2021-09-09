package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMultiParams(t *testing.T) {
	api := API{
		GetShopsShopPetsPetIDHandler: func(p GetShopsShopPetsPetIDParamsParser) GetShopsShopPetsPetIDResponser {
			ps, err := p.Parse()
			if err != nil {
				return GetShopsShopPetsPetIDResponseDefault(400)
			}
			assert.Equal(t, "paw", ps.Shop)
			assert.Equal(t, int64(1), ps.PetID)
			assert.Equal(t, stringP("white"), ps.Color)
			assert.Equal(t, int32P(1), ps.Page)
			return GetShopsShopPetsPetIDResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/shops/paw/pets/1?color=white&page=1", nil))

	assert.Equal(t, 200, w.Code)
}

func int32P(i int32) *int32    { return &i }
func stringP(s string) *string { return &s }
