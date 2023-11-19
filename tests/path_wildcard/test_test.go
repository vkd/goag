package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRequest_Names(t *testing.T) {
	api := API{
		GetPetsPetIDHandler: func(r GetPetsPetIDRequestParser) GetPetsPetIDResponse {
			params, err := r.Parse()
			require.NoError(t, err)

			var resp GetPetsPetIDResponse200JSON
			resp.Body.ID = params.Path.PetID
			return resp
		},
	}

	client := NewClient("", api)
	ctx := context.Background()

	{
		var req GetPetsPetIDRequest
		req.Path.PetID = 103
		resp, err := client.GetPetsPetID(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, int32(103), resp.(GetPetsPetIDResponse200JSON).Body.ID)
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
