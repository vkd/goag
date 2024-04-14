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
	ctx := context.Background()

	var params GetReviewsParams
	// Path
	params.Path.Shop = 1000

	// Query
	params.Query.IntReq = 1
	params.Query.Int = Just(int(2))
	params.Query.Int32Req = 3
	params.Query.Int32 = Just(int32(4))
	params.Query.Int64Req = 5
	params.Query.Int64 = Just(int64(6))
	params.Query.Float32Req = 7
	params.Query.Float32 = Just(float32(8))
	params.Query.Float64Req = 9
	params.Query.Float64 = Just(float64(10))

	params.Query.StringReq = "11"
	params.Query.String = Just("12")

	params.Query.Tag = Just([]string{"13", "14"})
	params.Query.Filter = Just([]int32{15, 16})

	// Headers
	params.Headers.RequestID = Just("2000")
	params.Headers.UserID = "2001"

	api := API{GetReviewsHandler: func(ctx context.Context, r GetReviewsRequest) GetReviewsResponse {
		p, err := r.Parse()
		require.NoError(t, err)
		// Path
		assert.Equal(t, params.Path.Shop, p.Path.Shop)

		// Query
		assert.Equal(t, params.Query.IntReq, p.Query.IntReq)
		assert.Equal(t, params.Query.Int, p.Query.Int)
		assert.Equal(t, params.Query.Int32Req, p.Query.Int32Req)
		assert.Equal(t, params.Query.Int32, p.Query.Int32)
		assert.Equal(t, params.Query.Int64Req, p.Query.Int64Req)
		assert.Equal(t, params.Query.Int64, p.Query.Int64)
		assert.Equal(t, params.Query.Float32Req, p.Query.Float32Req)
		assert.Equal(t, params.Query.Float32, p.Query.Float32)
		assert.Equal(t, params.Query.Float64Req, p.Query.Float64Req)
		assert.Equal(t, params.Query.Float64, p.Query.Float64)

		assert.Equal(t, params.Query.StringReq, p.Query.StringReq)
		assert.Equal(t, params.Query.String, p.Query.String)

		assert.Equal(t, params.Query.Tag, p.Query.Tag)
		assert.Equal(t, params.Query.Filter, p.Query.Filter)

		// Headers
		assert.Equal(t, params.Headers.RequestID, p.Headers.RequestID)
		assert.Equal(t, params.Headers.UserID, p.Headers.UserID)

		return NewGetReviewsResponseDefault(200)
	}}
	client := api.Client()
	resp, err := client.GetReviews(ctx, params)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.(GetReviewsResponseDefault).Code)
}

func (a API) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	a.ServeHTTP(w, req)
	return w.Result(), nil
}

func (a API) Client() *Client {
	return NewClient("", a)
}

func ptr[T any](v T) *T { return &v }
