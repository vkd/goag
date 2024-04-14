package test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseSchema(t *testing.T) {
	api := API{
		GetPetHandler: func(_ context.Context, _ GetPetRequest) GetPetResponse {
			return NewPetResponse(Pet{ID: 1, Name: "mike"})
		},
		GetV2PetHandler: func(ctx context.Context, r GetV2PetRequest) GetV2PetResponse {
			return NewPetResponse(Pet{ID: 2, Name: "luke"})
		},
		GetV3PetHandler: func(ctx context.Context, r GetV3PetRequest) GetV3PetResponse {
			return NewPetResponse(Pet{ID: 3, Name: "luke"})
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pet", nil))
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"id\":1,\"name\":\"mike\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	w = httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/v2/pet", nil))
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{\"id\":2,\"name\":\"luke\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	w = httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/v3/pet", nil))
	assert.Equal(t, 202, w.Code)
	assert.Equal(t, "{\"id\":3,\"name\":\"luke\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
