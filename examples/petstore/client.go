package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL    string
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

func NewClient(baseURL string, httpClient HTTPClient) *Client {
	return &Client{baseURL: baseURL, HTTPClient: httpClient}
}

func (c *Client) GetPets(ctx context.Context, request GetPetsRequest) (GetPetsResponse, error) {
	var requestURL = c.baseURL + "/pets"

	{
		var q = url.Values{}

		//q["limit"] = request.Query.Limit

		requestURL += "?" + q.Encode()
	}

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
		var response GetPetsResponse200JSON

		response.Headers.XNext = resp.Header.Get("x-next")

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetPetsResponse200JSON' response body: %w", err)
		}
		return response, nil

	default:
		var response GetPetsResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetPetsResponseDefaultJSON' response body: %w", err)
		}
		return response, nil

	}
}

func (c *Client) PostPets(ctx context.Context, request PostPetsRequest) (PostPetsResponse, error) {
	var requestURL = c.baseURL + "/pets"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, nil)
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
	case 201:
		var response PostPetsResponse201

		return response, nil

	default:
		var response PostPetsResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'PostPetsResponseDefaultJSON' response body: %w", err)
		}
		return response, nil

	}
}

func (c *Client) GetPetsPetID(ctx context.Context, request GetPetsPetIDRequest) (GetPetsPetIDResponse, error) {
	var requestURL = c.baseURL + "/pets/{petId}"

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
		var response GetPetsPetIDResponse200JSON

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetPetsPetIDResponse200JSON' response body: %w", err)
		}
		return response, nil

	default:
		var response GetPetsPetIDResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'GetPetsPetIDResponseDefaultJSON' response body: %w", err)
		}
		return response, nil

	}
}

func decodeJSON(r io.Reader, v interface{}, name string) {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		LogError(fmt.Errorf("decode json response %q: %w", name, err))
	}
}
