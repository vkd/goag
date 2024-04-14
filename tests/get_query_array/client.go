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

	query := make(url.Values, 2)
	if request.Query.Tag.IsSet {
		query["tag"] = request.Query.Tag.Value
	}
	if request.Query.Page.IsSet {
		{
			query_values := make([]string, 0, len(request.Query.Page.Value))
			for _, v := range request.Query.Page.Value {
				query_values = append(query_values, strconv.FormatInt(v, 10))
			}
			query["page"] = query_values
		}
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
		var response GetPetsResponse200
		return response, nil
	default:
		var response GetPetsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}
