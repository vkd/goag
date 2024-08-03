package test

import (
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

// GetShopsShop
// GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop.String()

	query := make(url.Values, 2)
	if request.Query.PageSchemaRefQuery.IsSet {
		query["page_schema_ref_query"] = []string{request.Query.PageSchemaRefQuery.Value.String()}
	}
	if request.Query.PageCustomTypeQuery.IsSet {
		query["page_custom_type_query"] = request.Query.PageCustomTypeQuery.Value.Strings()
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 200:
		var response GetShopsShopResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetShopsShopResponse200JSON' response body: %w", err)
		}
		return response, nil
	default:
		var response GetShopsShopResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}
