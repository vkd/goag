package test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaArray(t *testing.T) {
	h := GetPetsHandlerFunc(func(_ context.Context, _ GetPetsRequest) GetPetsResponse {
		return NewGetPetsResponse200JSON([]Pet{{ID: 1, Name: "mike"}})
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	require.Equal(t, 200, w.Code)
	var pets []Pet
	err := json.Unmarshal(w.Body.Bytes(), &pets)
	require.NoError(t, err)
	require.Len(t, pets, 1)
	assert.Equal(t, int64(1), pets[0].ID)
	assert.Equal(t, "mike", pets[0].Name)
}

func TestSchemaArray_Names(t *testing.T) {
	h := GetPetsNamesHandlerFunc(func(_ context.Context, _ GetPetsNamesRequest) GetPetsNamesResponse {
		return NewGetPetsNamesResponse200JSON([]string{"mike"})
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/pets/names", nil))

	require.Equal(t, 200, w.Code)
	var pets []string
	err := json.Unmarshal(w.Body.Bytes(), &pets)
	require.NoError(t, err)
	require.Len(t, pets, 1)
	assert.Equal(t, "mike", pets[0])
}
