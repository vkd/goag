package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

// PostShopsShopPets
// POST /shops/{shop}/pets
func (c *Client) PostShopsShopPets(ctx context.Context, request PostShopsShopPetsParams) (PostShopsShopPetsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop.String() + "/pets"

	query := make(url.Values, 1)
	if request.Query.Filter.IsSet {
		query["filter"] = request.Query.Filter.Value.Strings()
	}
	requestURL += "?" + query.Encode()

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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 201:
		var response PostShopsShopPetsResponse201
		return response, nil
	default:
		var response PostShopsShopPetsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}
