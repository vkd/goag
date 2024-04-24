package test

import (
	"context"
	"fmt"
	"net/http"
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

// GetShopsMinePetsMikeTails - GET /shops/mine/pets/mike/tails
func (c *Client) GetShopsMinePetsMikeTails(ctx context.Context, request GetShopsMinePetsMikeTailsParams) (GetShopsMinePetsMikeTailsResponse, error) {
	var requestURL = c.BaseURL + "/shops/mine/pets/mike/tails"

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
		var response GetShopsMinePetsMikeTailsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShop - GET /shops/{shop}
func (c *Client) GetShopsShop(ctx context.Context, request GetShopsShopParams) (GetShopsShopResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop

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
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop + "/"

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
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop + "/pets"

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
		var response GetShopsShopPetsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}

// GetShopsShopPetsMikePaws - GET /shops/{shop}/pets/mike/paws
func (c *Client) GetShopsShopPetsMikePaws(ctx context.Context, request GetShopsShopPetsMikePawsParams) (GetShopsShopPetsMikePawsResponse, error) {
	var requestURL = c.BaseURL + "/shops/" + request.Path.Shop + "/pets/mike/paws"

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
		var response GetShopsShopPetsMikePawsResponseDefault
		response.Code = resp.StatusCode

		return response, nil
	}
}
