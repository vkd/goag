package test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostBody(t *testing.T) {
	handler := PostPetsHandlerFunc(func(r PostPetsRequest) PostPetsResponse {
		req, err := r.Parse()
		if err != nil {
			assert.NoError(t, err)
			assert.Fail(t, "Should not call invalid function")
			return NewPostPetsResponseDefault(400)
		}
		assert.Equal(t, "mike", req.Body.Name)
		assert.Equal(t, "cat", req.Body.Tag)
		return NewPostPetsResponse201()
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("POST", "/pets", strings.NewReader(`{"name": "mike", "tag": "cat"}`)))

	assert.Equal(t, 201, w.Code)
}
