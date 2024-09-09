package test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestResponseSchema(t *testing.T) {
	handler := GetPetHandlerFunc(func(_ context.Context, _ GetPetRequest) GetPetResponse {
		return NewGetPetResponse200JSON(Pet{ID: 1, Name: "mike", CreatedAt: time.Date(2004, 5, 14, 16, 22, 7, 0, time.UTC)})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"created_at\":\"2004-05-14T16:22:07Z\",\"id\":1,\"name\":\"mike\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
