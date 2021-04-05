package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryArrayInt(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(ps GetPetsParams) GetPetsResponser {
			assert.Len(t, ps.Tag, 2)
			assert.Equal(t, []int64{2, 4}, ps.Tag)
			return GetPetsResponse200()
		},
		func(_ error) GetPetsResponser { return GetPetsResponseDefault(400) })

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets?tag=2&tag=4", nil))

	assert.Equal(t, 200, w.Code)
}
