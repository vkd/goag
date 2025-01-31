package test

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOctetStream(t *testing.T) {
	testBody := `hello`

	ctx := context.Background()
	api := API{
		GetPetsHandler: func(ctx context.Context, r GetPetsRequest) GetPetsResponse {
			_ = r.Parse()
			return NewGetPetsResponse200(io.NopCloser(strings.NewReader(testBody)))
		},
		PostPetsHandler: func(ctx context.Context, r PostPetsRequest) PostPetsResponse {
			param := r.Parse()
			bs, err := io.ReadAll(param.Body)
			require.NoError(t, err)

			assert.Equal(t, testBody, string(bs))
			return NewPostPetsResponse200()
		},
	}

	client := api.TestClient()

	// GetPets
	resp, err := client.GetPets(ctx, GetPetsParams{})
	require.NoError(t, err)

	body := resp.(GetPetsResponse200).Body

	bs, err := io.ReadAll(body)
	require.NoError(t, err)

	err = body.Close()
	require.NoError(t, err)

	assert.Equal(t, testBody, string(bs))

	// PostPets
	postResp, err := client.PostPets(ctx, PostPetsParams{
		Body: strings.NewReader(testBody),
	})
	require.NoError(t, err)

	assert.Equal(t, PostPetsResponse200{}, postResp.(PostPetsResponse200))
}
