package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostBody(t *testing.T) {
	ctx := context.Background()

	api := API{
		PostPetsHandler: func(ctx context.Context, r PostPetsRequest) PostPetsResponse {
			req, err := r.Parse()
			if err != nil {
				assert.NoError(t, err)
				assert.Fail(t, "Should not call invalid function")
				return NewPostPetsResponseDefault(400)
			}
			assert.Equal(t, "mike", req.Body.Name)
			assert.Equal(t, "cat", req.Body.Tag)
			return NewPostPetsResponse201()
		},
		PostPets2Handler: func(ctx context.Context, r PostPets2Request) PostPets2Response {
			req, err := r.Parse()
			if err != nil {
				assert.NoError(t, err)
				assert.Fail(t, "Should not call invalid function")
				return NewPostPets2ResponseDefault(400)
			}
			assert.Equal(t, "mike2", req.Body.Name)
			assert.Equal(t, "cat2", req.Body.Tag)
			return NewPostPets2Response201()
		},
	}

	{
		resp, err := api.Client().PostPets(ctx, PostPetsParams{
			Body: NewPet{
				Name: "mike",
				Tag:  "cat",
			},
		})
		require.NoError(t, err)
		assert.IsType(t, PostPetsResponse201{}, resp)
	}
	{
		resp, err := api.Client().PostPets2(ctx, PostPets2Params{
			Body: Pets2JSON(NewPetJSON(NewPet{
				Name: "mike2",
				Tag:  "cat2",
			})),
		})
		require.NoError(t, err)
		assert.IsType(t, PostPets2Response201{}, resp)
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
