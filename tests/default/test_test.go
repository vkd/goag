package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefault(t *testing.T) {
	ctx := context.Background()

	api := API{
		GetHandler:        func(r GetRequestParser) GetResponse { return GetResponseDefault{Code: 201} },
		GetShopsHandler:   func(r GetShopsRequestParser) GetShopsResponse { return GetShopsResponseDefault{Code: 202} },
		GetShopsRTHandler: func(r GetShopsRTRequestParser) GetShopsRTResponse { return GetShopsRTResponseDefault{Code: 203} },
		GetShopsActivateHandler: func(r GetShopsActivateRequestParser) GetShopsActivateResponse {
			return GetShopsActivateResponseDefault{Code: 204}
		},
		GetShopsActivateRTHandler: func(r GetShopsActivateRTRequestParser) GetShopsActivateRTResponse {
			return GetShopsActivateRTResponseDefault{Code: 205}
		},
		GetShopsShopPetsHandler: func(r GetShopsShopPetsRequestParser) GetShopsShopPetsResponse {
			return GetShopsShopPetsResponse200JSON{Headers: struct {
				Body  GetShopsShopPetsResponse200JSONBody
				XNext string
			}{XNext: "test-next-value"}}
		},
	}

	{
		resp, err := api.Client().Get(ctx, GetRequest{})
		require.NoError(t, err)
		assert.Equal(t, 201, resp.(GetResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShops(ctx, GetShopsRequest{})
		require.NoError(t, err)
		assert.Equal(t, 202, resp.(GetShopsResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsRT(ctx, GetShopsRTRequest{})
		require.NoError(t, err)
		assert.Equal(t, 203, resp.(GetShopsRTResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsActivate(ctx, GetShopsActivateRequest{})
		require.NoError(t, err)
		assert.Equal(t, 204, resp.(GetShopsActivateResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsActivateRT(ctx, GetShopsActivateRTRequest{})
		require.NoError(t, err)
		assert.Equal(t, 205, resp.(GetShopsActivateRTResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsShopPets(ctx, GetShopsShopPetsRequest{})
		require.NoError(t, err)
		require.IsType(t, GetShopsShopPetsResponse200JSON{}, resp, resp)
		assert.Equal(t, "test-next-value", resp.(GetShopsShopPetsResponse200JSON).Headers.XNext)
	}
}

func (a API) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	a.ServeHTTP(w, req)
	return w.Result(), nil
}

func (a API) Client() *Client {
	return NewClient("", a)
}

var testFunc = GetHandlerFunc(func(r GetRequestParser) GetResponse {
	return GetResponseDefault{Code: 101}
})

func TestCustom(t *testing.T) {
	resp := testFunc(GetRequest{}).(GetResponseDefault)
	assert.Equal(t, 101, resp.Code)
}

func TestCustom2(t *testing.T) {
	resp := testFunc(GetHTTPRequest(httptest.NewRequest("GET", "/", nil))).(GetResponseDefault)
	assert.Equal(t, 101, resp.Code)
}
