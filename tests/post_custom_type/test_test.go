package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vkd/goag/tests/post_custom_type/pkg"
)

func TestPostBody(t *testing.T) {
	ctx := context.Background()
	api := API{
		PostShopsShopPetsHandler: func(r PostShopsShopPetsRequest) PostShopsShopPetsResponse {
			req, err := r.Parse()
			if err != nil {
				assert.NoError(t, err)
				assert.Fail(t, "Should not call invalid function")
				return NewPostShopsShopPetsResponseDefault(400)
			}
			assert.Equal(t, "testshop", req.Path.Shop.V)
			assert.Equal(t, "tiger", req.Body.Tag.V)
			return NewPostShopsShopPetsResponse201()

		},
	}

	cli := NewClient("", api)

	resp, err := cli.PostShopsShopPets(ctx, PostShopsShopPetsParams{Path: struct{ Shop pkg.ShopType }{Shop: pkg.ShopType{V: "testshop"}}, Body: NewPet{Tag: pkg.PetTag{V: "tiger"}}})
	require.NoError(t, err)
	_, ok := resp.(PostShopsShopPetsResponse201)
	assert.True(t, ok)
}

func (a API) Do(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	a.ServeHTTP(w, req)
	return w.Result(), nil
}
