package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponents(t *testing.T) {
	ctx := context.Background()

	api := API{
		PostShopsHandler: func(ctx context.Context, r PostShopsRequest) PostShopsResponse {
			_, err := r.Parse()
			if err != nil {
				return NewErrorResponseResponse(Error{
					Detail: err.Error(),
				}, Just(10031))
			}
			return NewPostShopsResponse200JSON(Pets{})
		},
	}

	client := NewClient("", HTTPClientFunc(func(r *http.Request) (*http.Response, error) {
		w := httptest.NewRecorder()
		api.ServeHTTP(w, r)
		return w.Result(), nil
	}))

	resp, err := client.PostShops(ctx, PostShopsParams{})
	require.NoError(t, err)

	assert.Len(t, resp.(PostShopsResponse200JSON).Body, 0)
}
