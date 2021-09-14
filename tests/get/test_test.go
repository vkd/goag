package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	api := API{
		GetPetsHandler: func(_ GetPetsRequester) GetPetsResponder { return GetPetsResponse200() },
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetRequest_NotFound(t *testing.T) {
	api := API{
		GetPetsHandler: func(_ GetPetsRequester) GetPetsResponder { return GetPetsResponse200() },
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/not_found", nil))

	assert.Equal(t, 404, w.Code)
}
