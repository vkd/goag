package test

import (
	"context"
	"encoding/json"
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

// GetPets - GET /pets
func (c *Client) GetPets(ctx context.Context, request GetPetsParams) (GetPetsResponse, error) {
	var requestURL = c.BaseURL + "/pets"

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

		var response GetPetsResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetPetsResponse200JSON' response body: %w", err)
		}
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
