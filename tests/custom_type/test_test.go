package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vkd/goag/tests/custom_type/pkg"
)

func TestGetMultiParams(t *testing.T) {
	ctx := context.Background()
	testShop := pkg.Shop("paw")
	testPageSchemaRefQuery := PageCustom("testPage")
	testPageCustomTypeQuery := pkg.PageCustomTypeQuery("testPage2")
	testShopName := Shop(ShopName(pkg.Page("testPage3")))

	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			req, err := r.Parse()
			if err != nil {
				return NewGetShopsShopResponseDefault(400)
			}

			assert.Equal(t, testShop, req.Path.Shop)
			assert.Equal(t, Just(testPageSchemaRefQuery), req.Query.PageSchemaRefQuery)
			assert.Equal(t, Just(testPageCustomTypeQuery), req.Query.PageCustomTypeQuery)

			return NewGetShopsShopResponse200JSON(Shop(testShopName))
		},
	}

	client := NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		api.ServeHTTP(w, r)
		return w.Result(), nil
	}))

	var params GetShopsShopParams
	params.Path.Shop = testShop
	params.Query.PageSchemaRefQuery = Just(testPageSchemaRefQuery)
	params.Query.PageCustomTypeQuery = Just(testPageCustomTypeQuery)
	resp, err := client.GetShopsShop(ctx, params)
	require.NoError(t, err)

	body := resp.(GetShopsShopResponse200JSON)
	assert.Equal(t, testShopName, body.Body)
}
