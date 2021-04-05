package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	api := API{
		GetPetsHandler: NewGetPetsHandlerer(func(_ GetPetsParams) GetPetsResponser {
			return GetPetsResponse200()
		}),
	}

	w := httptest.NewRecorder()
	api.Router().ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetRequest_NotFound(t *testing.T) {
	fn := NewGetPetsHandlerer(func(gpp GetPetsParams) GetPetsResponser { return GetPetsResponse200() })

	w := httptest.NewRecorder()
	API{GetPetsHandler: fn}.Router().ServeHTTP(w, httptest.NewRequest("GET", "/not_found", nil))

	assert.Equal(t, 404, w.Code)
}
