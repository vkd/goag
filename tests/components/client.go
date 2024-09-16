package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

// PostPets
// POST /pets
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 200:
		var response PostPetsResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'PostPetsResponse200JSON' response body: %w", err)
		}
		return response, nil
	case 500:
		var response ErrorResponseResponse
		var hs []string
		hs = resp.Header.Values("X-Error-Code")
		if len(hs) > 0 {
			vInt64, err := strconv.ParseInt(hs[0], 10, 0)
			if err != nil {
				return nil, ErrParseParam{In: "header", Parameter: "X-Error-Code", Reason: "parse int", Err: err}
			}
			vOpt := int(vInt64)
			response.Headers.XErrorCode.Set(vOpt)
		}

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ErrorResponseResponse' response body: %w", err)
		}
		return response, nil

	default:
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}

// PostShops
// POST /shops
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

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case 200:
		var response PostShopsResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'PostShopsResponse200JSON' response body: %w", err)
		}
		return response, nil
	case 500:
		var response ErrorResponseResponse
		var hs []string
		hs = resp.Header.Values("X-Error-Code")
		if len(hs) > 0 {
			vInt64, err := strconv.ParseInt(hs[0], 10, 0)
			if err != nil {
				return nil, ErrParseParam{In: "header", Parameter: "X-Error-Code", Reason: "parse int", Err: err}
			}
			vOpt := int(vInt64)
			response.Headers.XErrorCode.Set(vOpt)
		}

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ErrorResponseResponse' response body: %w", err)
		}
		return response, nil

	default:
		return nil, fmt.Errorf("status code %d: not implemented", resp.StatusCode)
	}
}
