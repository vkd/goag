package test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseSchema(t *testing.T) {
	handler := GetPetHandlerFunc(func(_ context.Context, _ GetPetRequest) GetPetResponse {
		return NewGetPetResponse200JSON(Pet{ID: 1, Name: "mike"})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"id\":1,\"name\":\"mike\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
