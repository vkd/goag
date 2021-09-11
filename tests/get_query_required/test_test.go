package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryRequired(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(r GetPetsRequester) GetPetsResponser {
			_, err := r.Parse()
			if err != nil {
				return GetPetsResponseDefault(400)
			}
			return GetPetsResponse200()
		},
	)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 400, w.Code)
}
