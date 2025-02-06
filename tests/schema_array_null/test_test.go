package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaArrayNull(t *testing.T) {
	ctx := context.Background()
	api := API{
		GetPetsHandler: func(_ context.Context, _ GetPetsRequest) GetPetsResponse {
			return NewGetPetsResponse200JSON(GetPetsResponse200JSONBody{
				Array:    nil,
				Nullable: Null[[]string](),
			})
		},
	}

	client := api.TestClient()

	// GetPets
	resp, err := client.GetPets(ctx, GetPetsParams{})
	require.NoError(t, err)

	body := resp.(GetPetsResponse200JSON).Body

	assert.NotNil(t, body.Array)
	assert.Equal(t, false, body.Nullable.IsSet)
}
