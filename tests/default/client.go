package test

import (
	"context"
	"encoding/json"
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

// Get - GET /
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	default:
		var response GetResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShops - GET /shops
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	default:
		var response GetShopsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsRT - GET /shops/
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	default:
		var response GetShopsRTResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsActivate - GET /shops/activate
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	default:
		var response GetShopsActivateResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsActivateRT - GET /shops/activate/
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	default:
		var response GetShopsActivateRTResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsActivateTag - GET /shops/activate/tag
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	default:
		var response GetShopsActivateTagResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShop - GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + strconv.FormatInt(int64(request.Path.Shop), 10)

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
	default:
		var response GetShopsShopResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShopRT - GET /shops/{shop}/
func (c *Client) GetShopsShopRT(ctx context.Context, request GetShopsShopRTParams) (GetShopsShopRTResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + strconv.FormatInt(int64(request.Path.Shop), 10) + "/"

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
	default:
		var response GetShopsShopRTResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShopPets - GET /shops/{shop}/pets
func (c *Client) GetShopsShopPets(ctx context.Context, request GetShopsShopPetsParams) (GetShopsShopPetsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + strconv.FormatInt(int64(request.Path.Shop), 10) + "/pets"

	query := make(url.Values, 2)
	if request.Query.Page != nil {
		query["page"] = []string{strconv.FormatInt(int64(*request.Query.Page), 10)}
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 200:
		var response GetShopsShopPetsResponse200JSON
		response.Headers.XNext = resp.Header.Get("x-next")

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetShopsShopPetsResponse200JSON' response body: %w", err)
		}
		return response, nil
	default:
		var response GetShopsShopPetsResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetShopsShopPetsResponseDefaultJSON' response body: %w", err)
		}
		return response, nil
	}
}

// PostShopsShopReview - POST /shops/{shop}/review
func (c *Client) PostShopsShopReview(ctx context.Context, request PostShopsShopReviewParams) (PostShopsShopReviewResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + strconv.FormatInt(int64(request.Path.Shop), 10) + "/review"

	query := make(url.Values, 4)
	if request.Query.Page != nil {
		query["page"] = []string{strconv.FormatInt(int64(*request.Query.Page), 10)}
	}
	query["page_size"] = []string{strconv.FormatInt(int64(request.Query.PageSize), 10)}
	query["tag"] = request.Query.Tag
	{
		query_values := make([]string, 0, len(request.Query.Filter))
		for _, v := range request.Query.Filter {
			query_values = append(query_values, strconv.FormatInt(int64(v), 10))
		}
		query["filter"] = query_values
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if request.Headers.RequestID != nil {
		req.Header.Set("request-id", *request.Headers.RequestID)
	}
	req.Header.Set("user-id", request.Headers.UserID)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 200:
		var response PostShopsShopReviewResponse200JSON
		response.Headers.XNext = resp.Header.Get("x-next")

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'PostShopsShopReviewResponse200JSON' response body: %w", err)
		}
		return response, nil
	default:
		var response PostShopsShopReviewResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'PostShopsShopReviewResponseDefaultJSON' response body: %w", err)
		}
		return response, nil
	}
}
