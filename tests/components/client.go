package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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

// PostPets - POST /pets
func (c *Client) PostPets(ctx context.Context, request PostPetsParams) (PostPetsResponse, error) {
	var requestURL = c.BaseURL + "/pets"

	bs, err := json.Marshal(request.Body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewReader(bs))
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

		var response PostPetsResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'PostPetsResponse200JSON' response body: %w", err)
		}
		return response, nil
	case 500:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response ErrorResponseResponse
		var hs []string
		hs = resp.Header.Values("X-Error-Code")
		if len(hs) > 0 {
			var vOpt int
			if len(hs) == 1 {
				vInt64, err := strconv.ParseInt(hs[0], 10, 0)
				if err != nil {
					return nil, ErrParseParam{In: "header", Parameter: "X-Error-Code", Reason: "parse int", Err: err}
				}
				vOpt = int(vInt64)
			} else {
				return nil, ErrParseParam{In: "header", Parameter: "X-Error-Code", Reason: "multiple values found: single value expected"}
			}
			response.Headers.XErrorCode.Set(vOpt)
		}

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ErrorResponseResponse' response body: %w", err)
		}
		return response, nil

	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}

// PostShops - POST /shops
func (c *Client) PostShops(ctx context.Context, request PostShopsParams) (PostShopsResponse, error) {
	var requestURL = c.BaseURL + "/shops"

	bs, err := json.Marshal(request.Body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewReader(bs))
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

		var response PostShopsResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'PostShopsResponse200JSON' response body: %w", err)
		}
		return response, nil
	case 500:
		if resp.Body != nil {
			defer resp.Body.Close()
		}

		var response ErrorResponseResponse
		var hs []string
		hs = resp.Header.Values("X-Error-Code")
		if len(hs) > 0 {
			var vOpt int
			if len(hs) == 1 {
				vInt64, err := strconv.ParseInt(hs[0], 10, 0)
				if err != nil {
					return nil, ErrParseParam{In: "header", Parameter: "X-Error-Code", Reason: "parse int", Err: err}
				}
				vOpt = int(vInt64)
			} else {
				return nil, ErrParseParam{In: "header", Parameter: "X-Error-Code", Reason: "multiple values found: single value expected"}
			}
			response.Headers.XErrorCode.Set(vOpt)
		}

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ErrorResponseResponse' response body: %w", err)
		}
		return response, nil

	default:
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}

func (a API) LocalClient() *Client {
	return NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
		return w.Result(), nil
	}))
}
