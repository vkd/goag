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
	testShopSchemaPath := Shop("testShopSchemaPath")
	testPageIntQuery := int(3)
	testPageSchemaQuery := Page(4)
	testPageIntQueryRequired := int(5)
	testPageSchemaQueryRequired := Page(6)
	testOrgIntHeader := int(7)
	testOrgSchemaHeader := Organization(8)
	testOrgIntHeaderRequired := int(9)
	testOrgSchemaHeaderRequired := Organization(10)

	ctx := context.Background()

	api := API{
		PostShopsShopStringSepShopSchemaPetsHandler: func(ctx context.Context, r PostShopsShopStringSepShopSchemaPetsRequest) PostShopsShopStringSepShopSchemaPetsResponse {
			params, err := r.Parse()
			require.NoError(t, err)

			assert.Equal(t, params.Path.ShopString, testShopStringPath)
			assert.Equal(t, params.Path.ShopSchema, testShopSchemaPath)
			assert.Equal(t, *params.Query.PageInt, testPageIntQuery)
			assert.Equal(t, *params.Query.PageSchema, testPageSchemaQuery)
			assert.Equal(t, params.Query.PageIntReq, testPageIntQueryRequired)
			assert.Equal(t, params.Query.PageSchemaReq, testPageSchemaQueryRequired)
			assert.Equal(t, *params.Headers.XOrganizationInt, testOrgIntHeader)
			assert.Equal(t, *params.Headers.XOrganizationSchema, testOrgSchemaHeader)
			assert.Equal(t, params.Headers.XOrganizationIntRequired, testOrgIntHeaderRequired)
			assert.Equal(t, params.Headers.XOrganizationSchemaRequired, testOrgSchemaHeaderRequired)

			return NewPostShopsShopStringSepShopSchemaPetsResponse200()
		},
	}

	client := NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		api.ServeHTTP(w, r)
		return w.Result(), nil
	}))

	var params PostShopsShopStringSepShopSchemaPetsParams
	params.Path.ShopString = testShopStringPath
	params.Path.ShopSchema = testShopSchemaPath
	params.Query.PageInt = &testPageIntQuery
	params.Query.PageSchema = &testPageSchemaQuery
	params.Query.PageIntReq = testPageIntQueryRequired
	params.Query.PageSchemaReq = testPageSchemaQueryRequired
	params.Headers.XOrganizationInt = &testOrgIntHeader
	params.Headers.XOrganizationSchema = &testOrgSchemaHeader
	params.Headers.XOrganizationIntRequired = testOrgIntHeaderRequired
	params.Headers.XOrganizationSchemaRequired = testOrgSchemaHeaderRequired

	resp, err := client.PostShopsShopStringSepShopSchemaPets(ctx, params)
	require.NoError(t, err)
	assert.IsType(t, PostShopsShopStringSepShopSchemaPetsResponse200{}, resp)
}
