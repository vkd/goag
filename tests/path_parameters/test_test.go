package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponents(t *testing.T) {
	testShopStringPath := "testShopStringPath"
	testShopSchemaPath := NewShop("testShopSchemaPath")

	ctx := context.Background()

	api := API{
		PostShopsShopStringShopSchemaPetsHandler: func(ctx context.Context, r PostShopsShopStringShopSchemaPetsRequest) PostShopsShopStringShopSchemaPetsResponse {
			params, err := r.Parse()
			require.NoError(t, err)

			assert.Equal(t, params.Path.ShopString, testShopStringPath)
			assert.Equal(t, params.Path.ShopSchema, testShopSchemaPath)

			return NewPostShopsShopStringShopSchemaPetsResponse200()
		},
	}

	client := NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		api.ServeHTTP(w, r)
		return w.Result(), nil
	}))

	var params PostShopsShopStringShopSchemaPetsParams
	params.Path.ShopString = testShopStringPath
	params.Path.ShopSchema = testShopSchemaPath

	resp, err := client.PostShopsShopStringShopSchemaPets(ctx, params)
	require.NoError(t, err)
	assert.IsType(t, PostShopsShopStringShopSchemaPetsResponse200{}, resp)
}
