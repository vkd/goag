package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseDefault(t *testing.T) {
	handler := GetPetsHandlerFunc(func(_ GetPetsRequester) GetPetsResponder {
		return GetPetsResponseDefaultJSON(400, Error{Message: "test default response"})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"message\":\"test default response\"}\n", w.Body.String())
}
