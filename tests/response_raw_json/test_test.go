package test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseSchema(t *testing.T) {
	rawBody := `{"id":"test"}`

	handler := GetPetHandlerFunc(func(_ context.Context, _ GetPetRequest) GetPetResponse {
		return NewGetPetResponse200JSON(json.RawMessage(rawBody))
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, rawBody+"\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
