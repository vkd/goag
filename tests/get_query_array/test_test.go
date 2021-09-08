package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryArray(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(p GetPetsParamsParser) GetPetsResponser {
			ps, err := p.Parse()
			if err != nil {
				return GetPetsResponseDefault(400)
			}
			assert.Len(t, ps.Tag, 2)
			assert.Equal(t, []string{"cat", "dog"}, ps.Tag)
			return GetPetsResponse200()
		},
	)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets?tag=cat&tag=dog", nil))

	assert.Equal(t, 200, w.Code)
}
