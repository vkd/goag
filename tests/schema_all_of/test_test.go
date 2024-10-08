package test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaAllOf(t *testing.T) {
	handler := GetPetHandlerFunc(func(_ context.Context, _ GetPetRequest) GetPetResponse {
		return NewGetPetResponse200JSON(Pet{ID: 1, NewPet: NewPet{Name: "mike", Tag: Just("cat")}})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	var pet Pet
	err := json.Unmarshal(w.Body.Bytes(), &pet)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), pet.ID)
	assert.Equal(t, "mike", pet.Name)
	assert.Equal(t, Just("cat"), pet.Tag)
}
