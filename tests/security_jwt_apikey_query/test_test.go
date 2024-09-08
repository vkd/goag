package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefault(t *testing.T) {
	type mykey struct{}

	api := API{
		PostLoginHandler: func(ctx context.Context, r PostLoginRequest) PostLoginResponse {
			var output string
			if s, ok := r.HTTP().Context().Value(mykey{}).(string); ok {
				output = s
			}
			return NewPostLoginResponse200JSON(PostLoginResponse200JSONBody{
				Output: output,
			})
		},
		PostShopsHandler: func(ctx context.Context, r PostShopsRequest) PostShopsResponse {
			return NewPostShopsResponse200JSON(PostShopsResponse200JSONBody{
				Output: r.HTTP().Context().Value(mykey{}).(string),
			})
		},

		SecurityBearerAuth: func(r *http.Request, token string) (*http.Request, bool) {
			if token != "valid" {
				return nil, false
			}
			return r.WithContext(context.WithValue(r.Context(), mykey{}, "my-value")), true
		},
		SecurityAPIKeyAuthAccessToken: func(r *http.Request, token string) (*http.Request, bool) {
			if token != "personal access token" {
				return nil, false
			}
			return r.WithContext(context.WithValue(r.Context(), mykey{}, "my-value-2")), true
		},
		SecurityAPIKeyAuthPersonalAccessToken: func(r *http.Request, token string) (*http.Request, bool) {
			if token != "pat2" {
				return nil, false
			}
			return r.WithContext(context.WithValue(r.Context(), mykey{}, "my-value-3")), true
		},
	}

	for _, tt := range []struct {
		path         string
		responseCode int
		response     string
	}{
		{"/login", 200, ""},
		{"/shops", 401, ""},
	} {
		t.Run("no_auth|"+tt.path, func(t *testing.T) {
			resp, err := api.Do(httptest.NewRequest("POST", tt.path, nil))
			require.NoError(t, err)
			assert.Equal(t, tt.responseCode, resp.StatusCode)
			if resp.StatusCode != 200 {
				return
			}
			var out struct{ Output string }
			err = json.NewDecoder(resp.Body).Decode(&out)
			assert.NoError(t, err)
			assert.Equal(t, tt.response, out.Output)
		})
	}

	for _, tt := range []struct {
		path         string
		responseCode int
		response     string
	}{
		{"/login", 200, ""},
		{"/shops", 200, "my-value"},
	} {
		t.Run("auth_bearer_header|"+tt.path, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt.path, nil)
			req.Header.Set("Authorization", "Bearer valid")
			resp, err := api.Do(req)
			require.NoError(t, err)
			assert.Equal(t, tt.responseCode, resp.StatusCode)
			var out struct{ Output string }
			err = json.NewDecoder(resp.Body).Decode(&out)
			assert.NoError(t, err)
			assert.Equal(t, tt.response, out.Output)
		})
	}

	for _, tt := range []struct {
		path         string
		responseCode int
		response     string
	}{
		{"/login", 200, ""},
		{"/shops", 200, "my-value-2"},
	} {
		t.Run("auth_apikey_header|"+tt.path, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt.path, nil)
			req.Header.Set("Access-Token", "personal access token")
			resp, err := api.Do(req)
			require.NoError(t, err)
			assert.Equal(t, tt.responseCode, resp.StatusCode)
			var out struct{ Output string }
			err = json.NewDecoder(resp.Body).Decode(&out)
			assert.NoError(t, err)
			assert.Equal(t, tt.response, out.Output)
		})
	}

	for _, tt := range []struct {
		path         string
		responseCode int
		response     string
	}{
		{"/login", 200, ""},
		{"/shops", 200, "my-value-3"},
	} {
		t.Run("auth_apikey_query|"+tt.path, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt.path+"?Personal-Access-Token=pat2", nil)
			resp, err := api.Do(req)
			require.NoError(t, err)
			assert.Equal(t, tt.responseCode, resp.StatusCode)
			var out struct{ Output string }
			err = json.NewDecoder(resp.Body).Decode(&out)
			assert.NoError(t, err)
			assert.Equal(t, tt.response, out.Output)
		})
	}
}

func (a API) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	a.ServeHTTP(w, req)
	return w.Result(), nil
}

func (a API) Client() *Client {
	return NewClient("", a)
}
