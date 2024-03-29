package test

import (
	"context"
	"fmt"
	"net/http"
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 200:
		var response GetPetsResponse200
		var hs []string
		hs = resp.Header.Values("x-next")
		if len(hs) > 0 {
			response.Headers.XNext = hs[0]
		}

		hs = resp.Header.Values("x-next-two")
		if len(hs) == 0 {
			return nil, fmt.Errorf("header parameter 'x-next-two': is required")
		}
		if len(hs) > 0 {
			response.Headers.XNextTwo = hs[0]
		}

		return response, nil

	default:
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}
