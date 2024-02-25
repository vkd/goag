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

type HTTPClientFunc func(*http.Request) (*http.Response, error)

var _ HTTPClient = HTTPClientFunc(nil)

func (f HTTPClientFunc) Do(r *http.Request) (*http.Response, error) { return f(r) }

var _ HTTPClient = (*http.Client)(nil)

func NewClient(baseURL string, httpClient HTTPClient) *Client {
	return &Client{BaseURL: baseURL, HTTPClient: httpClient}
}

// ListPets - GET /pets
func (c *Client) ListPets(ctx context.Context, request ListPetsParams) (ListPetsResponse, error) {
	var requestURL = c.BaseURL + "/pets"

	query := make(url.Values, 1)
	if request.Query.Limit != nil {
		query["limit"] = []string{strconv.FormatInt(int64(*request.Query.Limit), 10)}
	}
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
		var response ListPetsResponse200JSON
		response.Headers.XNext = resp.Header.Get("x-next")

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ListPetsResponse200JSON' response body: %w", err)
		}
		return response, nil
	default:
		var response ListPetsResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ListPetsResponseDefaultJSON' response body: %w", err)
		}
		return response, nil
	}
}

// CreatePets - POST /pets
func (c *Client) CreatePets(ctx context.Context, request CreatePetsParams) (CreatePetsResponse, error) {
	var requestURL = c.BaseURL + "/pets"

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
		var response CreatePetsResponse201
		return response, nil
	default:
		var response CreatePetsResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'CreatePetsResponseDefaultJSON' response body: %w", err)
		}
		return response, nil
	}
}

// ShowPetByID - GET /pets/{petId}
func (c *Client) ShowPetByID(ctx context.Context, request ShowPetByIDParams) (ShowPetByIDResponse, error) {
	var requestURL = c.BaseURL + "/pets/" + request.Path.PetID

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
		var response ShowPetByIDResponse200JSON
		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ShowPetByIDResponse200JSON' response body: %w", err)
		}
		return response, nil
	default:
		var response ShowPetByIDResponseDefaultJSON
		response.Code = resp.StatusCode

		err := json.NewDecoder(resp.Body).Decode(&response.Body)
		if err != nil {
			return nil, fmt.Errorf("decode 'ShowPetByIDResponseDefaultJSON' response body: %w", err)
		}
		return response, nil
	}
}
