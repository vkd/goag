package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMultiParams(t *testing.T) {
	testShop := "paw"
	testPage := int32(2)
	testRequestID := "abcdef"

	api := API{
		GetShopsShopHandler: func(r GetShopsShopRequester) GetShopsShopResponder {
			req, err := r.Parse()
			if err != nil {
				return GetShopsShopResponseDefault(400)
			}
			assert.Equal(t, testShop, req.Shop)
			assert.Equal(t, testPage, *req.Page)
			assert.Equal(t, testRequestID, *req.RequestID)
			return GetShopsShopResponse200()
		},
	}

	target := fmt.Sprintf("/shops/%s?page=%d", testShop, testPage)
	req := httptest.NewRequest("GET", target, nil)
	req.Header.Set("request-id", testRequestID)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetMultiParams_Optional(t *testing.T) {
	api := API{
		GetShopsShopHandler: func(r GetShopsShopRequester) GetShopsShopResponder {
			req, err := r.Parse()
			if err != nil {
				return GetShopsShopResponseDefault(400)
			}
			assert.Nil(t, req.Page)
			return GetShopsShopResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/shops/paw", nil))

	assert.Equal(t, 200, w.Code)
}

func TestGetMultiParams_BadRequest(t *testing.T) {
	api := API{
		GetShopsShopHandler: func(r GetShopsShopRequester) GetShopsShopResponder {
			_, err := r.Parse()
			if err != nil {
				return GetShopsShopResponseDefault(400)
			}
			return GetShopsShopResponse200()
		},
	}

	for _, target := range []string{
		// "/shops/paw/pets/1?color=white&page=2",
		"/shops/paw?page=b",
		"/shops/?page=2",
	} {
		w := httptest.NewRecorder()
		api.ServeHTTP(w, httptest.NewRequest("GET", target, nil))

		assert.Equal(t, 400, w.Code, "suppose to be a bad request: %s", target)
	}
}
