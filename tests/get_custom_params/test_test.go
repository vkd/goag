package test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMultiParams(t *testing.T) {
	testShop := Shop("paw")
	testPage := Page(2)
	testPageReq := Page(3)
	testPages := []Page{4}
	testRequestID := RequestID("abcdef")

	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			req, err := r.Parse()
			require.NoError(t, err)

			assert.Equal(t, testShop, req.Path.Shop)
			assert.Equal(t, testPage, *req.Query.Page)
			assert.Equal(t, testPageReq, req.Query.PageReq)
			assert.Equal(t, testPages, req.Query.Pages)
			assert.Equal(t, testRequestID, *req.Headers.RequestID)
			return NewGetShopsShopResponse200()
		},
	}

	target := fmt.Sprintf("/shops/%s?page=%d&page_req=%d&pages=%d", testShop, testPage, testPageReq, testPages[0])
	req := httptest.NewRequest("GET", target, nil)
	req.Header.Set("request-id", testRequestID.String())
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
