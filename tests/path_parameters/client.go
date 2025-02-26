package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

type Client struct {
	BaseURL    string
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HTTPClientFunc func(*http.Request) (*http.Response, error)

var _ HTTPClient = HTTPClientFunc(nil)

func (f HTTPClientFunc) Do(r *http.Request) (*http.Response, error) { return f(r) }

var _ HTTPClient = (*http.Client)(nil)

func NewClient(baseURL string, httpClient HTTPClient) *Client {
	return &Client{BaseURL: baseURL, HTTPClient: httpClient}
}

// PostShopsShopStringShopSchemaPets - POST /shops/{shop_string}/{shop_schema}/pets
func (c *Client) PostShopsShopStringShopSchemaPets(ctx context.Context, request PostShopsShopStringShopSchemaPetsParams) (PostShopsShopStringShopSchemaPetsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(request.Path.ShopString) + "/" + url.PathEscape(request.Path.ShopSchema.String()) + "/pets"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	case 200:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostShopsShopStringShopSchemaPetsResponse200
		return response, nil

	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}

func (a API) TestClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
