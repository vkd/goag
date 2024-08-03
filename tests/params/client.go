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

// GetReviews
// GET /shops/{shop}/reviews
func (c *Client) GetReviews(ctx context.Context, request GetReviewsParams) (GetReviewsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + strconv.FormatInt(int64(request.Path.Shop), 10) + "/reviews"

	query := make(url.Values, 14)
	query["int_req"] = []string{strconv.FormatInt(int64(request.Query.IntReq), 10)}
	if request.Query.Int.IsSet {
		query["int"] = []string{strconv.FormatInt(int64(request.Query.Int.Value), 10)}
	}
	query["int32_req"] = []string{strconv.FormatInt(int64(request.Query.Int32Req), 10)}
	if request.Query.Int32.IsSet {
		query["int32"] = []string{strconv.FormatInt(int64(request.Query.Int32.Value), 10)}
	}
	query["int64_req"] = []string{strconv.FormatInt(request.Query.Int64Req, 10)}
	if request.Query.Int64.IsSet {
		query["int64"] = []string{strconv.FormatInt(request.Query.Int64.Value, 10)}
	}
	query["float32_req"] = []string{strconv.FormatFloat(float64(request.Query.Float32Req), 'e', -1, 32)}
	if request.Query.Float32.IsSet {
		query["float32"] = []string{strconv.FormatFloat(float64(request.Query.Float32.Value), 'e', -1, 32)}
	}
	query["float64_req"] = []string{strconv.FormatFloat(request.Query.Float64Req, 'e', -1, 64)}
	if request.Query.Float64.IsSet {
		query["float64"] = []string{strconv.FormatFloat(request.Query.Float64.Value, 'e', -1, 64)}
	}
	query["string_req"] = []string{request.Query.StringReq}
	if request.Query.String.IsSet {
		query["string"] = []string{request.Query.String.Value}
	}
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
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
