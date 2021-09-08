package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathParam(t *testing.T) {
	handler := GetPetsPetIDHandlerFunc(func(p GetPetsPetIDParamsParser) GetPetsPetIDResponser {
		ps, err := p.Parse()
		if err != nil {
			return GetPetsPetIDResponseDefault(400)
		}
		assert.Equal(t, 1, ps.PetID)
		return GetPetsPetIDResponse200()
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets/1", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetPathParam_Invalid(t *testing.T) {
	handler := GetPetsPetIDHandlerFunc(func(p GetPetsPetIDParamsParser) GetPetsPetIDResponser {
		_, err := p.Parse()
		if err != nil {
			return GetPetsPetIDResponseDefault(400)
		}
		assert.Fail(t, "petId is not a number")
		return GetPetsPetIDResponse200()
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets/a", nil))

	assert.Equal(t, 400, w.Code)
}
