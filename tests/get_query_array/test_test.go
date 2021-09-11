package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryArray(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(r GetPetsRequester) GetPetsResponser {
			req, err := r.Parse()
			if err != nil {
				return GetPetsResponseDefault(400)
			}
			assert.Equal(t, []string{"cat", "dog"}, req.Tag)
			return GetPetsResponse200()
		},
	)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets?tag=cat&tag=dog", nil))

	assert.Equal(t, 200, w.Code)
}
