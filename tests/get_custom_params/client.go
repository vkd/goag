package test

import (
	"context"
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

// GetShopsShop - GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop.String()

	query := make(url.Values, 4)
	if request.Query.Page.IsSet {
		query["page"] = request.Query.Page.Value.Strings()
	}
	query["page_req"] = request.Query.PageReq.Strings()
	if request.Query.Pages.IsSet {
		qv := make([]string, 0, len(request.Query.Pages.Value))
		for _, v := range request.Query.Pages.Value {
			qv = append(qv, v.String())
		}
		query["pages"] = qv
	}
	if request.Query.PageCustom.IsSet {
		query["page_custom"] = []string{request.Query.PageCustom.Value.String()}
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if request.Headers.RequestID.IsSet {
		req.Header.Set("request-id", request.Headers.RequestID.Value.String())
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
