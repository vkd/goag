package test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseHeader(t *testing.T) {
	testHeader := "test_header"

	handler := GetPetsHandlerFunc(func(_ context.Context, _ GetPetsRequest) GetPetsResponse {
		return NewGetPetsResponse200(testHeader)
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, testHeader, w.Header().Get("x-next"))
}
