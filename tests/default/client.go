package test

import (
	"bytes"
	"context"
	"encoding/json"
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

// Get
// GET /
func (c *Client) Get(ctx context.Context, request GetParams) (GetResponse, error) {
	var requestURL = c.BaseURL + "/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShops
// GET /shops
func (c *Client) GetShops(ctx context.Context, request GetShopsParams) (GetShopsResponse, error) {
	var requestURL = c.BaseURL + "/shops"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsRT
// GET /shops/
func (c *Client) GetShopsRT(ctx context.Context, request GetShopsRTParams) (GetShopsRTResponse, error) {
	var requestURL = c.BaseURL + "/shops/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsRTResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsActivate
// GET /shops/activate
func (c *Client) GetShopsActivate(ctx context.Context, request GetShopsActivateParams) (GetShopsActivateResponse, error) {
	var requestURL = c.BaseURL + "/shops/activate"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsActivateResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsActivateRT
// GET /shops/activate/
func (c *Client) GetShopsActivateRT(ctx context.Context, request GetShopsActivateRTParams) (GetShopsActivateRTResponse, error) {
	var requestURL = c.BaseURL + "/shops/activate/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsActivateRTResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsActivateTag
// GET /shops/activate/tag
func (c *Client) GetShopsActivateTag(ctx context.Context, request GetShopsActivateTagParams) (GetShopsActivateTagResponse, error) {
	var requestURL = c.BaseURL + "/shops/activate/tag"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsActivateTagResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShop
// GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(strconv.FormatInt(int64(request.Path.Shop), 10))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsShopResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShopRT
// GET /shops/{shop}/
func (c *Client) GetShopsShopRT(ctx context.Context, request GetShopsShopRTParams) (GetShopsShopRTResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(strconv.FormatInt(int64(request.Path.Shop), 10)) + "/"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsShopRTResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShopPets
// GET /shops/{shop}/pets
func (c *Client) GetShopsShopPets(ctx context.Context, request GetShopsShopPetsParams) (GetShopsShopPetsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(strconv.FormatInt(int64(request.Path.Shop), 10)) + "/pets"

	query := make(url.Values, 2)
	if request.Query.Page.IsSet {
		query["page"] = []string{strconv.FormatInt(int64(request.Query.Page.Value), 10)}
	}
	query["page_size"] = []string{strconv.FormatInt(int64(request.Query.PageSize), 10)}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
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

		var response GetShopsShopPetsResponse200JSON
		var hs []string
		hs = resp.Header.Values("x-next")
		if len(hs) > 0 {
			vOpt := hs[0]
			response.Headers.XNext.Set(vOpt)
		}

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetShopsShopPetsResponse200JSON' response body: %w", err)
		}
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetShopsShopPetsResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetShopsShopPetsResponseDefaultJSON' response body: %w", err)
		}
		return response, nil
	}
}

// ReviewShop - Review shop.
// Returns a current pet.
// POST /shops/{shop}/review
func (c *Client) ReviewShop(ctx context.Context, request ReviewShopParams) (ReviewShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(strconv.FormatInt(int64(request.Path.Shop), 10)) + "/review"

	query := make(url.Values, 4)
	if request.Query.Page.IsSet {
		query["page"] = []string{strconv.FormatInt(int64(request.Query.Page.Value), 10)}
	}
	query["page_size"] = []string{strconv.FormatInt(int64(request.Query.PageSize), 10)}
	if request.Query.Tag.IsSet {
		query["tag"] = request.Query.Tag.Value
	}
	if request.Query.Filter.IsSet {
		qv := make([]string, 0, len(request.Query.Filter.Value))
		for _, v := range request.Query.Filter.Value {
			qv = append(qv, strconv.FormatInt(int64(v), 10))
		}
		query["filter"] = qv
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
	if request.Headers.RequestID.IsSet {
		req.Header.Set("request-id", request.Headers.RequestID.Value)
	}
	req.Header.Set("user-id", request.Headers.UserID)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	case 200:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response ReviewShopResponse200JSON
		var hs []string
		hs = resp.Header.Values("x-next")
		if len(hs) > 0 {
			vOpt := hs[0]
			response.Headers.XNext.Set(vOpt)
		}

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ReviewShopResponse200JSON' response body: %w", err)
		}
		return response, nil
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response ReviewShopResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ReviewShopResponseDefaultJSON' response body: %w", err)
		}
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
