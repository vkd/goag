package test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testKey struct{}

func TestGetRequest(t *testing.T) {
	testValue := "test_value"

	api := API{
		GetPetsHandler: GetPetsHandlerFunc(func(r GetPetsRequester) GetPetsResponser {
			req := r.Parse()
			assert.Equal(t, testValue, req.HTTPRequest.Context().Value(testKey{}).(string))
			return GetPetsResponse200()
		}),
	}

	r := httptest.NewRequest("GET", "/pets", nil)
	r = r.WithContext(context.WithValue(context.Background(), testKey{}, testValue))
	w := httptest.NewRecorder()
	api.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)
}
