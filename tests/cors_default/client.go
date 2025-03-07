package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
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

// GetShops - GET /shops
func (c *Client) GetShops(ctx context.Context, request GetShopsParams) (GetShopsResponse, error) {
	var requestURL = c.BaseURL + "/shops"

	query := make(url.Values, 1)
	if qvOpt, ok := request.Query.Page.Get(); ok {
		query["page"] = []string{strconv.FormatInt(int64(qvOpt), 10)}
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if hvOpt, ok := request.Headers.AccessKey.Get(); ok {
		req.Header.Set("access-key", hvOpt)
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

		var response GetShopsResponse200
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShop - GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(request.Path.Shop)

	query := make(url.Values, 1)
	if qvOpt, ok := request.Query.Page.Get(); ok {
		query["page"] = []string{strconv.FormatInt(int64(qvOpt), 10)}
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if hvOpt, ok := request.Headers.RequestID.Get(); ok {
		req.Header.Set("request-id", hvOpt)
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

		var response GetShopsShopResponse200
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsShopResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// PostShopsShop - POST /shops/{shop}
func (c *Client) PostShopsShop(ctx context.Context, request PostShopsShopParams) (PostShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(request.Path.Shop)

	query := make(url.Values, 1)
	if qvOpt, ok := request.Query.Page.Get(); ok {
		query["page"] = []string{strconv.FormatInt(int64(qvOpt), 10)}
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if hvOpt, ok := request.Headers.QueryID.Get(); ok {
		req.Header.Set("query-id", hvOpt)
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

		var response PostShopsShopResponse200
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostShopsShopResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

func (a API) TestClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
