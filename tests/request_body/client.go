package test

import (
	"bytes"
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

// PostPets - POST /pets
func (c *Client) PostPets(ctx context.Context, request PostPetsParams) (PostPetsResponse, error) {
	var requestURL = c.BaseURL + "/pets"

	bs, err := json.Marshal(request.Body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewReader(bs))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	case 201:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostPetsResponse201
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostPetsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// PostPets2 - POST /pets2
func (c *Client) PostPets2(ctx context.Context, request PostPets2Params) (PostPets2Response, error) {
	var requestURL = c.BaseURL + "/pets2"

	bs, err := json.Marshal(request.Body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewReader(bs))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	case 201:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostPets2Response201
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response PostPets2ResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

func (a API) LocalClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
