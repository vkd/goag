package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryRequest(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(ps GetPetsParams) GetPetsResponser {
			assert.Equal(t, int32(1), ps.Limit)
			return GetPetsResponse200()
		},
		func(_ error) GetPetsResponser { return GetPetsResponseDefault(400) })

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets?limit=1", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetQueryRequest_BadRequest(t *testing.T) {
	handler := GetPetsHandlerFunc(
		func(ps GetPetsParams) GetPetsResponser {
			assert.Equal(t, int32(1), ps.Limit)
			return GetPetsResponse200()
		},
		func(_ error) GetPetsResponser { return GetPetsResponseDefault(http.StatusBadRequest) })

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets?limit=a", nil))

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
