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

// GetReviews - GET /shops/{shop}/reviews
func (c *Client) GetReviews(ctx context.Context, request GetReviewsParams) (GetReviewsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop.String() + "/reviews"

	query := make(url.Values, 14)
	query["int_req"] = []string{request.Query.IntReq.String()}
	if request.Query.Int != nil {
		query["int"] = []string{request.Query.Int.String()}
	}
	query["int32_req"] = []string{request.Query.Int32Req.String()}
	if request.Query.Int32 != nil {
		query["int32"] = []string{request.Query.Int32.String()}
	}
	query["int64_req"] = []string{request.Query.Int64Req.String()}
	if request.Query.Int64 != nil {
		query["int64"] = []string{request.Query.Int64.String()}
	}
	query["float32_req"] = []string{request.Query.Float32Req.String()}
	if request.Query.Float32 != nil {
		query["float32"] = []string{request.Query.Float32.String()}
	}
	query["float64_req"] = []string{request.Query.Float64Req.String()}
	if request.Query.Float64 != nil {
		query["float64"] = []string{request.Query.Float64.String()}
	}
	query["string_req"] = []string{request.Query.StringReq.String()}
	if request.Query.String != nil {
		query["string"] = []string{request.Query.String.String()}
	}
	query["tag"] = request.Query.Tag
	{
		query_values := make([]string, 0, len(request.Query.Filter))
		for _, v := range request.Query.Filter {
			query_values = append(query_values, strconv.FormatInt(int64(v), 10))
		}
		query["filter"] = query_values
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
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
	default:
		var response GetReviewsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}
