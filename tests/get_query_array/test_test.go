package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryArray_Strings(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(r GetPetsRequester) GetPetsResponder {
			req, err := r.Parse()
			if err != nil {
				return GetPetsResponseDefault(400)
			}
			assert.Equal(t, []string{"cat", "dog"}, req.Query.Tag)
			return GetPetsResponse200()
		},
	)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets?tag=cat&tag=dog", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetQueryArray_Ints(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(r GetPetsRequester) GetPetsResponder {
			req, err := r.Parse()
			if err != nil {
				return GetPetsResponseDefault(400)
			}
			assert.Equal(t, []int64{2, 4}, req.Query.Page)
			return GetPetsResponse200()
		},
	)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets?page=2&page=4", nil))

	assert.Equal(t, 200, w.Code)
}
