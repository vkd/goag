package test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponentsParams(t *testing.T) {
	testShop := "paw"
	testPage := int32(2)

	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			req, err := r.Parse()
			require.NoError(t, err)
			assert.Equal(t, testShop, string(req.Path.Shop))
			assert.Equal(t, testPage, int32(*req.Query.Page))
			return NewGetShopsShopResponse200()
		},
	}

	target := fmt.Sprintf("/shops/%s?page=%d", testShop, testPage)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", target, nil))

	assert.Equal(t, 200, w.Code)
}

func TestComponentsParams_Optional(t *testing.T) {
	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			req, err := r.Parse()
			if err != nil {
				return NewGetShopsShopResponseDefault(400)
			}
			assert.Nil(t, req.Query.Page)
			return NewGetShopsShopResponse200()
		},
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/shops/paw", nil))

	assert.Equal(t, 200, w.Code)
}

func TestComponentsParams_BadRequest(t *testing.T) {
	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			_, err := r.Parse()
			if err != nil {
				return NewGetShopsShopResponseDefault(400)
			}
			return NewGetShopsShopResponse200()
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
