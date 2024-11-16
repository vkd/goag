package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseSchema(t *testing.T) {
	ctx := context.Background()
	api := API{
		GetPetHandler: func(_ context.Context, _ GetPetRequest) GetPetResponse {
			return NewGetPetResponse200JSON(RawResponse{
				AdditionalProperties: map[string]any{
					"field1": "hello",
				},
			})
		},
	}

	resp, err := api.Client().GetPet(ctx, GetPetParams{})
	require.NoError(t, err)
	assert.Equal(t, "hello", resp.(GetPetResponse200JSON).Body.AdditionalProperties["field1"])
}

func (a API) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	a.ServeHTTP(w, req)
	return w.Result(), nil
}

func (a API) Client() *Client {
	return NewClient("", a)
}
