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
			return NewPetResponse(Pet{ID: 1, Name: "mike"})
		},
		GetV2PetHandler: func(ctx context.Context, r GetV2PetRequest) GetV2PetResponse {
			return NewPetResponse(Pet{ID: 2, Name: "luke"})
		},
		GetV3PetHandler: func(ctx context.Context, r GetV3PetRequest) GetV3PetResponse {
			return NewPetResponse(Pet{ID: 3, Name: "luke"})
		},
	}

	{
		resp, err := api.Client().GetPet(ctx, GetPetParams{})
		require.NoError(t, err)
		assert.Equal(t, int64(1), resp.(PetResponse).Body.ID)
	}
	{
		resp, err := api.Client().GetV2Pet(ctx, GetV2PetParams{})
		require.NoError(t, err)
		assert.Equal(t, int64(2), resp.(PetResponse).Body.ID)
	}
	{
		resp, err := api.Client().GetV3Pet(ctx, GetV3PetParams{})
		require.NoError(t, err)
		assert.Equal(t, int64(3), resp.(PetResponse).Body.ID)
	}

	w := httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/pet", nil))
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"id\":1,\"name\":\"mike\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	w = httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/v2/pet", nil))
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "{\"id\":2,\"name\":\"luke\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	w = httptest.NewRecorder()
	api.ServeHTTP(w, httptest.NewRequest("GET", "/v3/pet", nil))
	assert.Equal(t, 202, w.Code)
	assert.Equal(t, "{\"id\":3,\"name\":\"luke\"}\n", w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

func (a API) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	a.ServeHTTP(w, req)
	return w.Result(), nil
}

func (a API) Client() *Client {
	return NewClient("", a)
}
