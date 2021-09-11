package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQueryRequest(t *testing.T) {
	testPetID := int32(1)

	api := API{
		GetPetsHandler: func(r GetPetsRequester) GetPetsResponser {
			params, err := r.Parse()
			if err != nil {
				return GetPetsResponseDefault(400)
			}
			assert.Equal(t, testPetID, *params.Limit)
			return GetPetsResponse200()
		},
	}

	target := fmt.Sprintf("/pets?limit=%d", testPetID)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", target, nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetQueryRequest_BadRequest(t *testing.T) {
	api := API{
		GetPetsHandler: func(r GetPetsRequester) GetPetsResponser {
			_, err := r.Parse()
			if err != nil {
				return GetPetsResponseDefault(http.StatusBadRequest)
			}
			assert.Fail(t, "petId is not a number")
			return GetPetsResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pets?limit=a", nil))

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
