package test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostBody(t *testing.T) {
	handler := PostPetsHandlerFunc(func(r PostPetsRequester) PostPetsResponser {
		req, err := r.Parse()
		if err != nil {
			assert.NoError(t, err)
			assert.Fail(t, "Should not call invalid function")
			return PostPetsResponseDefault(400)
		}
		assert.Equal(t, "mike", req.Body.Name)
		assert.Equal(t, "cat", req.Body.Tag)
		return PostPetsResponse201()
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("POST", "/pets", strings.NewReader(`{"name": "mike", "tag": "cat"}`)))

	assert.Equal(t, 201, w.Code)
}
