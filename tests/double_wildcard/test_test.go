package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest_Names(t *testing.T) {
	api := API{
		GetPetsPetIDNamesHandler: func(r GetPetsPetIDNamesRequest) GetPetsPetIDNamesResponse {
			return NewGetPetsPetIDNamesResponse200()
		},
		GetPetsPetIDShopsHandler: func(r GetPetsPetIDShopsRequest) GetPetsPetIDShopsResponse { panic("wrong handler") },
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pets/mike/names", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetRequest_Shops(t *testing.T) {
	api := API{
		GetPetsPetIDShopsHandler: func(r GetPetsPetIDShopsRequest) GetPetsPetIDShopsResponse {
			return NewGetPetsPetIDShopsResponse200()
		},
		GetPetsPetIDNamesHandler: func(r GetPetsPetIDNamesRequest) GetPetsPetIDNamesResponse { panic("wrong handler") },
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pets/mike/shops", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetRequest_NotFound(t *testing.T) {
	api := API{
		GetPetsPetIDNamesHandler: func(r GetPetsPetIDNamesRequest) GetPetsPetIDNamesResponse {
			return NewGetPetsPetIDNamesResponse200()
		},
		GetPetsPetIDShopsHandler: func(r GetPetsPetIDShopsRequest) GetPetsPetIDShopsResponse {
			return NewGetPetsPetIDShopsResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pets/mike/not_found", nil))

	assert.Equal(t, 404, w.Code)
}
