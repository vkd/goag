package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathParam(t *testing.T) {
	testPetID := int(1)

	handler := GetPetsPetIDHandlerFunc(func(r GetPetsPetIDRequester) GetPetsPetIDResponser {
		ps, err := r.Parse()
		if err != nil {
			return GetPetsPetIDResponseDefault(400)
		}
		assert.Equal(t, testPetID, ps.PetID)
		return GetPetsPetIDResponse200()
	})

	target := fmt.Sprintf("/pets/%d", testPetID)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", target, nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetPathParam_Invalid(t *testing.T) {
	handler := GetPetsPetIDHandlerFunc(func(r GetPetsPetIDRequester) GetPetsPetIDResponser {
		_, err := r.Parse()
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
