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

// GetShopsShop
// GET /shops/{shop}/pages/{page}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop.String() + "/pages/" + request.Path.Page.String()

	query := make(url.Values, 5)
	if request.Query.Page.IsSet {
		cv := request.Query.Page.Value.Int32()
		query["page"] = []string{strconv.FormatInt(int64(cv), 10)}
	}
	cv := request.Query.PageReq.Int32()
	query["page_req"] = []string{strconv.FormatInt(int64(cv), 10)}
	if request.Query.Pages.IsSet {
		qv := make([]string, 0, len(request.Query.Pages.Value))
		for _, v := range request.Query.Pages.Value {
			qv = append(qv, strconv.FormatInt(int64(v.Int32()), 10))
		}
		query["pages"] = qv
	}
	if request.Query.PagesArray.IsSet {
		cv := request.Query.PagesArray.Value.Int32s()
		qv := make([]string, 0, len(cv))
		for _, v := range cv {
			qv = append(qv, strconv.FormatInt(int64(v), 10))
		}
		query["pages_array"] = qv
	}
	if request.Query.PageCustom.IsSet {
		cv := request.Query.PageCustom.Value.String()
		query["page_custom"] = []string{cv}
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

	switch resp.StatusCode {
	case 200:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsShopResponse200
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsShopResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

func (a API) TestClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
