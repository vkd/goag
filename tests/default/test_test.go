package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefault(t *testing.T) {
	ctx := context.Background()

	api := API{
		GetHandler: func(ctx context.Context, r GetRequest) GetResponse { return GetResponseDefault{Code: 201} },
		GetShopsHandler: func(ctx context.Context, r GetShopsRequest) GetShopsResponse {
			return GetShopsResponseDefault{Code: 202}
		},
		GetShopsRTHandler: func(ctx context.Context, r GetShopsRTRequest) GetShopsRTResponse {
			return GetShopsRTResponseDefault{Code: 203}
		},
		GetShopsActivateHandler: func(ctx context.Context, r GetShopsActivateRequest) GetShopsActivateResponse {
			return GetShopsActivateResponseDefault{Code: 204}
		},
		GetShopsActivateRTHandler: func(ctx context.Context, r GetShopsActivateRTRequest) GetShopsActivateRTResponse {
			return GetShopsActivateRTResponseDefault{Code: 205}
		},
		GetShopsShopPetsHandler: func(ctx context.Context, r GetShopsShopPetsRequest) GetShopsShopPetsResponse {
			return GetShopsShopPetsResponse200JSON{Headers: struct {
				XNext Maybe[string]
			}{Just("test-next-value")}}
		},
		ReviewShopHandler: func(ctx context.Context, r ReviewShopRequest) ReviewShopResponse {
			params, err := r.Parse()
			if err != nil {
				return NewReviewShopResponseDefaultJSON(500, Error{Message: fmt.Errorf("parse request: %w", err).Error()})
			}
			return NewReviewShopResponseDefaultJSON(206, Error{Message: params.Body.Name})
		},
	}

	{
		resp, err := api.Client().Get(ctx, GetParams{})
		require.NoError(t, err)
		assert.Equal(t, 201, resp.(GetResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShops(ctx, GetShopsParams{})
		require.NoError(t, err)
		assert.Equal(t, 202, resp.(GetShopsResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsRT(ctx, GetShopsRTParams{})
		require.NoError(t, err)
		assert.Equal(t, 203, resp.(GetShopsRTResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsActivate(ctx, GetShopsActivateParams{})
		require.NoError(t, err)
		assert.Equal(t, 204, resp.(GetShopsActivateResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsActivateRT(ctx, GetShopsActivateRTParams{})
		require.NoError(t, err)
		assert.Equal(t, 205, resp.(GetShopsActivateRTResponseDefault).Code)
	}
	{
		resp, err := api.Client().GetShopsShopPets(ctx, GetShopsShopPetsParams{})
		require.NoError(t, err)
		require.IsType(t, GetShopsShopPetsResponse200JSON{}, resp, resp)
		assert.Equal(t, Just("test-next-value"), resp.(GetShopsShopPetsResponse200JSON).Headers.XNext)
	}
	{
		resp, err := api.Client().ReviewShop(ctx, ReviewShopParams{Body: NewPet{Name: "207"}})
		require.NoError(t, err)
		require.IsType(t, ReviewShopResponseDefaultJSON{}, resp, resp)
		assert.Equal(t, 206, resp.(ReviewShopResponseDefaultJSON).Code)
		assert.Equal(t, "207", resp.(ReviewShopResponseDefaultJSON).Body.Message)
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

var testFunc = GetHandlerFunc(func(ctx context.Context, r GetRequest) GetResponse {
	return GetResponseDefault{Code: 101}
})

func TestCustom(t *testing.T) {
	resp := testFunc(context.Background(), GetParams{}).(GetResponseDefault)
	assert.Equal(t, 101, resp.Code)
}

func TestCustom2(t *testing.T) {
	resp := testFunc(context.Background(), GetHTTPRequest(httptest.NewRequest("GET", "/", nil))).(GetResponseDefault)
	assert.Equal(t, 101, resp.Code)
}
