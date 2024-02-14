package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefault(t *testing.T) {
	type mykey struct{}

	api := API{
		PostLoginHandler: func(ctx context.Context, r PostLoginRequest) PostLoginResponse { return NewPostLoginResponse200() },
		PostShopsHandler: func(ctx context.Context, r PostShopsRequest) PostShopsResponse { return NewPostShopsResponse200() },

		SecurityBearerAuth: func(w http.ResponseWriter, r *http.Request, token string, next http.Handler) {
			if token != "valid" {
				w.WriteHeader(401)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), mykey{}, "my-value")))
		},
	}

	for _, tt := range []struct {
		path         string
		responseCode int
	}{
		{"/login", 200},
		{"/shops", 401},
	} {
		t.Run("no_auth|"+tt.path, func(t *testing.T) {
			resp, err := api.Do(httptest.NewRequest("POST", tt.path, nil))
			require.NoError(t, err)
			assert.Equal(t, tt.responseCode, resp.StatusCode)
		})
	}

	for _, tt := range []struct {
		path         string
		responseCode int
	}{
		{"/login", 200},
		{"/shops", 200},
	} {
		t.Run("auth_header|"+tt.path, func(t *testing.T) {
			req := httptest.NewRequest("POST", tt.path, nil)
			req.Header.Set("Authorization", "Bearer valid")
			resp, err := api.Do(req)
			require.NoError(t, err)
			assert.Equal(t, tt.responseCode, resp.StatusCode)
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
