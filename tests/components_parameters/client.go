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

// PostShopsShopStringSepShopSchemaPets - POST /shops/{shop_string}/sep/{shop_schema}/pets
func (c *Client) PostShopsShopStringSepShopSchemaPets(ctx context.Context, request PostShopsShopStringSepShopSchemaPetsParams) (PostShopsShopStringSepShopSchemaPetsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.ShopString.String() + "/sep/" + request.Path.ShopSchema.String() + "/pets"

	query := make(url.Values, 4)
	if request.Query.PageInt != nil {
		query["page_int"] = []string{request.Query.PageInt.String()}
	}
	if request.Query.PageSchema != nil {
		query["page_schema"] = []string{request.Query.PageSchema.String()}
	}
	query["page_int_req"] = []string{request.Query.PageIntReq.String()}
	query["page_schema_req"] = []string{request.Query.PageSchemaReq.String()}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if request.Headers.XOrganizationInt != nil {
		req.Header.Set("X-Organization-Int", request.Headers.XOrganizationInt.String())
	}
	if request.Headers.XOrganizationSchema != nil {
		req.Header.Set("X-Organization-Schema", request.Headers.XOrganizationSchema.String())
	}
	req.Header.Set("X-Organization-Int-Required", request.Headers.XOrganizationIntRequired.String())
	req.Header.Set("X-Organization-Schema-Required", request.Headers.XOrganizationSchemaRequired.String())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 200:
		var response PostShopsShopStringSepShopSchemaPetsResponse200
		return response, nil

	default:
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}
