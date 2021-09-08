package test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostBody(t *testing.T) {
	handler := PostPetsHandlerFunc(func(ps PostPetsParamsParser) PostPetsResponser {
		p, err := ps.Parse()
		if err != nil {
			assert.NoError(t, err)
			assert.Fail(t, "Should not call invalid function")
			return PostPetsResponseDefault(400)
		}
		assert.Equal(t, "mike", p.Body.Name)
		assert.Equal(t, "cat", p.Body.Tag)
		return PostPetsResponse201()
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("POST", "/pets", strings.NewReader(`{"name": "mike", "tag": "cat"}`)))

	assert.Equal(t, 201, w.Code)
}
