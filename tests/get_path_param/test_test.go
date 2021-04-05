package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathParam(t *testing.T) {
	handler := GetPetsPetIDHandlerFunc(func(ps GetPetsPetIDParams) GetPetsPetIDResponser {
		assert.Equal(t, 1, ps.PetID)
		return GetPetsPetIDResponse200()
	}, func(_ error) GetPetsPetIDResponser {
		return GetPetsPetIDResponseDefault(400)
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets/1", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetPathParam_Invalid(t *testing.T) {
	handler := GetPetsPetIDHandlerFunc(func(ps GetPetsPetIDParams) GetPetsPetIDResponser {
		assert.Fail(t, "petId is not a number")
		return GetPetsPetIDResponse200()
	}, func(err error) GetPetsPetIDResponser {
		// log.Printf("process 400: %q %q %v", in, param, err)
		// assert.Equal(t, "path", in)
		// assert.Equal(t, "petId", param)
		// assert.Error(t, err)
		return GetPetsPetIDResponseDefault(400)
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets/a", nil))

	assert.Equal(t, 400, w.Code)
}
