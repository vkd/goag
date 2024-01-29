package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

// GetPetsPetID - GET /pets/{pet_id}
func (c *Client) GetPetsPetID(ctx context.Context, request GetPetsPetIDParams) (GetPetsPetIDResponse, error) {
	var requestURL = c.BaseURL + "/pets/" + strconv.FormatInt(int64(request.Path.PetID), 10)

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
		var response GetPetsPetIDResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetPetsPetIDResponse200JSON' response body: %w", err)
		}
		return response, nil

	default:
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}
