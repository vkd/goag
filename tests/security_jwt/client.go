package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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

// PostLogin - POST /login
func (c *Client) PostLogin(ctx context.Context, request PostLoginParams) (PostLoginResponse, error) {
	var requestURL = c.BaseURL + "/login"

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

		var response PostLoginResponse200
		return response, nil
	case 401:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostLoginResponse401
		return response, nil

	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}

// PostShops - POST /shops
func (c *Client) PostShops(ctx context.Context, request PostShopsParams) (PostShopsResponse, error) {
	var requestURL = c.BaseURL + "/shops"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Authorization", request.Headers.Authorization)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	case 200:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostShopsResponse200
		return response, nil
	case 401:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostShopsResponse401
		return response, nil

	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}

func (a API) LocalClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
