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

// GetReviews - GET /shops/{shop}/reviews
func (c *Client) GetReviews(ctx context.Context, request GetReviewsParams) (GetReviewsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + url.PathEscape(strconv.FormatInt(int64(request.Path.Shop), 10)) + "/reviews"

	query := make(url.Values, 14)
	query["int_req"] = []string{strconv.FormatInt(int64(request.Query.IntReq), 10)}
	if qvOpt, ok := request.Query.Int.Get(); ok {
		query["int"] = []string{strconv.FormatInt(int64(qvOpt), 10)}
	}
	query["int32_req"] = []string{strconv.FormatInt(int64(request.Query.Int32Req), 10)}
	if qvOpt, ok := request.Query.Int32.Get(); ok {
		query["int32"] = []string{strconv.FormatInt(int64(qvOpt), 10)}
	}
	query["int64_req"] = []string{strconv.FormatInt(request.Query.Int64Req, 10)}
	if qvOpt, ok := request.Query.Int64.Get(); ok {
		query["int64"] = []string{strconv.FormatInt(qvOpt, 10)}
	}
	query["float32_req"] = []string{strconv.FormatFloat(float64(request.Query.Float32Req), 'e', -1, 32)}
	if qvOpt, ok := request.Query.Float32.Get(); ok {
		query["float32"] = []string{strconv.FormatFloat(float64(qvOpt), 'e', -1, 32)}
	}
	query["float64_req"] = []string{strconv.FormatFloat(request.Query.Float64Req, 'e', -1, 64)}
	if qvOpt, ok := request.Query.Float64.Get(); ok {
		query["float64"] = []string{strconv.FormatFloat(qvOpt, 'e', -1, 64)}
	}
	query["string_req"] = []string{request.Query.StringReq}
	if qvOpt, ok := request.Query.String.Get(); ok {
		query["string"] = []string{qvOpt}
	}
	if qvOpt, ok := request.Query.Tag.Get(); ok {
		query["tag"] = qvOpt
	}
	if qvOpt, ok := request.Query.Filter.Get(); ok {
		qv := make([]string, 0, len(qvOpt))
		for _, v := range qvOpt {
			qv = append(qv, strconv.FormatInt(int64(v), 10))
		}
		query["filter"] = qv
	}
	requestURL += "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if hvOpt, ok := request.Headers.RequestID.Get(); ok {
		req.Header.Set("request-id", hvOpt)
	}
	req.Header.Set("user-id", request.Headers.UserID)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http client Do(): %w", err)
	}

	switch resp.StatusCode {
	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response GetReviewsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

func (a API) LocalClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
