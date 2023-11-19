package test

import (
	"context"
	"fmt"
	"net/http"
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

var _ HTTPClient = (*http.Client)(nil)

func NewClient(baseURL string, httpClient HTTPClient) *Client {
	return &Client{BaseURL: baseURL, HTTPClient: httpClient}
}

// GetShopsShop - GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopRequest) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop

	query := make(url.Values, 1)
	if request.Query.Page != nil {
		query["page"] = []string{strconv.FormatInt(int64(*request.Query.Page), 10)}
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if request.Headers.RequestID != nil {
		req.Header.Set("request-id", *request.Headers.RequestID)
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
		var response GetShopsShopResponse200
		return response, nil
	default:
		var response GetShopsShopResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}
