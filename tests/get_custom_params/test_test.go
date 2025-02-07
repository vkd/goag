package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vkd/goag/tests/get_custom_params/pkg"
)

func TestGetMultiParams(t *testing.T) {
	ctx := context.Background()
	testShop := Shop("paw")
	testPage := Page(2)
	testPageReq := Page(3)
	testPages := []Page{4}
	testPagesArray := Pages{5, 6}
	testRequestID := RequestID("abcdef")
	testPageCustom := pkg.Page("7")
	testPageCustomPath := pkg.Page("8")

	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			req, err := r.Parse()
			if err != nil {
				return NewGetShopsShopResponseDefault(400)
			}
			assert.Equal(t, testShop, req.Path.Shop)
			assert.Equal(t, Just(testPage), req.Query.Page)
			assert.Equal(t, testPageReq, req.Query.PageReq)
			assert.Equal(t, Just(testPages), req.Query.Pages)
			assert.Equal(t, Just(testPagesArray), req.Query.PagesArray)
			assert.Equal(t, Just(testRequestID), req.Headers.RequestID)
			assert.Equal(t, Just(testPageCustom), req.Query.PageCustom)
			assert.Equal(t, testPageCustomPath, req.Path.Page)
			return NewGetShopsShopResponse200()
		},
	}

	client := NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		api.ServeHTTP(w, r)
		return w.Result(), nil
	}))

	resp, err := client.GetShopsShop(ctx, GetShopsShopParams{
		Path: GetShopsShopParamsPath{
			Shop: testShop,
			Page: testPageCustomPath,
		},
		Query: GetShopsShopParamsQuery{
			Page:       Just(testPage),
			PageReq:    testPageReq,
			Pages:      Just(testPages),
			PagesArray: Just(testPagesArray),
			PageCustom: Just(testPageCustom),
		},
		Headers: GetShopsShopParamsHeaders{
			RequestID: Just(testRequestID),
		},
	})
	require.NoError(t, err)

	body := resp.(GetShopsShopResponse200)
	assert.NotNil(t, body)
}
