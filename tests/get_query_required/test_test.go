package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryRequired(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(ps GetPetsParams) GetPetsResponser {
			return GetPetsResponse200()
		},
		func(err error) GetPetsResponser {
			return GetPetsResponseDefault(400)
		})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 400, w.Code)
}
