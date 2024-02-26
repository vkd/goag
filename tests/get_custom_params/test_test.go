package test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Page int

func (p Page) String() string { return strconv.Itoa(int(p)) }

func (p *Page) UnmarshalText(data []byte) error {
	i, err := strconv.Atoi(string(data))
	if err != nil {
		return fmt.Errorf("parse int: %w", err)
	}
	*p = Page(i)
	return nil
}

func TestGetMultiParams(t *testing.T) {
	testShop := "paw"
	testPage := Page(2)
	testPageReq := Page(3)
	testPages := []Page{4}
	testRequestID := "abcdef"

	api := API{
		GetShopsShopHandler: func(ctx context.Context, r GetShopsShopRequest) GetShopsShopResponse {
			req, err := r.Parse()
			if err != nil {
				return NewGetShopsShopResponseDefault(400)
			}
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
	req.Header.Set("request-id", testRequestID)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
