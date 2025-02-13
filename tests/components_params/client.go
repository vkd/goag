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

// PostShopsNew
// POST /shops/new
func (c *Client) PostShopsNew(ctx context.Context, request PostShopsNewParams) (PostShopsNewResponse, error) {
	var requestURL = c.BaseURL + "/shops/new"

	query := make(url.Values, 1)
	if request.Query.Page.IsSet {
		query["page"] = []string{strconv.FormatInt(int64(request.Query.Page.Value), 10)}
	}
	requestURL += "?" + query.Encode()

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

		var response PostShopsNewResponse200
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostShopsNewResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShop
// GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(request.Path.Shop)

	query := make(url.Values, 1)
	if request.Query.Page.IsSet {
		query["page"] = []string{strconv.FormatInt(int64(request.Query.Page.Value), 10)}
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
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

// GetShopsShopReviews
// GET /shops/{shop}/reviews
func (c *Client) GetShopsShopReviews(ctx context.Context, request GetShopsShopReviewsParams) (GetShopsShopReviewsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(request.Path.Shop) + "/reviews"

	query := make(url.Values, 1)
	if request.Query.Page.IsSet {
		query["page"] = []string{strconv.FormatInt(int64(request.Query.Page.Value), 10)}
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
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

		var response GetShopsShopReviewsResponse200
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsShopReviewsResponseDefault
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
