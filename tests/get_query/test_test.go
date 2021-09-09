package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryRequest(t *testing.T) {
	api := API{
		GetPetsHandler: func(p GetPetsParamsParser) GetPetsResponser {
			params, err := p.Parse()
			if err != nil {
				return GetPetsResponseDefault(400)
			}
			assert.Equal(t, pInt32(1), params.Limit)
			return GetPetsResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pets?limit=1", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetQueryRequest_BadRequest(t *testing.T) {
	api := API{
		GetPetsHandler: func(ps GetPetsParamsParser) GetPetsResponser {
			params, err := ps.Parse()
			if err != nil {
				return GetPetsResponseDefault(http.StatusBadRequest)
			}
			assert.Equal(t, pInt32(1), params.Limit)
			return GetPetsResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pets?limit=a", nil))

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func pInt32(i int32) *int32 { return &i }
