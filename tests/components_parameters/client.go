package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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
	var requestURL = c.BaseURL + "/shops/" + request.Path.ShopString + "/sep/" + request.Path.ShopSchema.Shopa().Shopb().Shopc().String() + "/pets"

	query := make(url.Values, 5)
	if request.Query.PageInt.IsSet {
		query["page_int"] = []string{strconv.FormatInt(int64(request.Query.PageInt.Value), 10)}
	}
	if request.Query.PageSchema.IsSet {
		query["page_schema"] = []string{strconv.FormatInt(int64(request.Query.PageSchema.Value.Int32()), 10)}
	}
	if request.Query.PagesSchema.IsSet {
		qv := make([]string, 0, len(request.Query.PagesSchema.Value.Int32s()))
		for _, v := range request.Query.PagesSchema.Value.Int32s() {
			qv = append(qv, strconv.FormatInt(int64(v), 10))
		}
		query["pages_schema"] = qv
	}
	query["page_int_req"] = []string{strconv.FormatInt(int64(request.Query.PageIntReq), 10)}
	query["page_schema_req"] = []string{strconv.FormatInt(int64(request.Query.PageSchemaReq.Int32()), 10)}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if request.Headers.XOrganizationInt.IsSet {
		req.Header.Set("X-Organization-Int", strconv.FormatInt(int64(request.Headers.XOrganizationInt.Value), 10))
	}
	if request.Headers.XOrganizationSchema.IsSet {
		req.Header.Set("X-Organization-Schema", strconv.FormatInt(int64(request.Headers.XOrganizationSchema.Value.Int()), 10))
	}
	req.Header.Set("X-Organization-Int-Required", strconv.FormatInt(int64(request.Headers.XOrganizationIntRequired), 10))
	req.Header.Set("X-Organization-Schema-Required", strconv.FormatInt(int64(request.Headers.XOrganizationSchemaRequired.Int()), 10))

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

func (a API) TestClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
