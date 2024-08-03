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

// PostShopsShopStringSepShopSchemaPets
// POST /shops/{shop_string}/sep/{shop_schema}/pets
func (c *Client) PostShopsShopStringSepShopSchemaPets(ctx context.Context, request PostShopsShopStringSepShopSchemaPetsParams) (PostShopsShopStringSepShopSchemaPetsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.ShopString + "/sep/" + request.Path.ShopSchema.String() + "/pets"

	query := make(url.Values, 5)
	if request.Query.PageInt.IsSet {
		query["page_int"] = []string{strconv.FormatInt(int64(request.Query.PageInt.Value), 10)}
	}
	if request.Query.PageSchema.IsSet {
		query["page_schema"] = []string{request.Query.PageSchema.Value.String()}
	}
	if request.Query.PagesSchema.IsSet {
		query["pages_schema"] = request.Query.PagesSchema.Value.Strings()
	}
	query["page_int_req"] = []string{strconv.FormatInt(int64(request.Query.PageIntReq), 10)}
	query["page_schema_req"] = []string{request.Query.PageSchemaReq.String()}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if request.Headers.XOrganizationInt.IsSet {
		req.Header.Set("X-Organization-Int", strconv.FormatInt(int64(request.Headers.XOrganizationInt.Value), 10))
	}
	if request.Headers.XOrganizationSchema.IsSet {
		req.Header.Set("X-Organization-Schema", request.Headers.XOrganizationSchema.Value.String())
	}
	req.Header.Set("X-Organization-Int-Required", strconv.FormatInt(int64(request.Headers.XOrganizationIntRequired), 10))
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
