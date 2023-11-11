package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostRequest(t *testing.T) {
	handler := PostPetsHandlerFunc(func(_ PostPetsRequestParser) PostPetsResponse {
		return NewPostPetsResponse200()
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("POST", "/pets", nil))

	assert.Equal(t, 200, w.Code)
}
