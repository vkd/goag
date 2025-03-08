package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostRequest(t *testing.T) {
	ctx := context.Background()
	testID := int64(1)

	api := API{
		PostPetsHandler: func(ctx context.Context, r PostPetsRequest) PostPetsResponse {
			params, err := r.Parse()
			require.NoError(t, err)

			return NewPetResponse(Pet{
				ID:   testID,
				Name: params.Body.Name,
				Tag:  params.Body.Tag,
				Tago: params.Body.Tago,
			})
		},
	}
	cli := api.LocalClient()

	for _, tt := range []struct {
		Name string
		Body NewPet
	}{
		{"all set", NewPet{
			Name: "all_set",
			Tag:  Pointer("tag_all_set"),
			Tago: Just(Pointer("tago_all_set")),
		}},

		{"all null", NewPet{
			Name: "all_set",
			Tag:  Null[string](),
			Tago: Just(Null[string]()),
		}},

		{"all null and optional", NewPet{
			Name: "all_set",
			Tag:  Null[string](),
			Tago: Nothing[Nullable[string]](),
		}},
	} {
		t.Run(tt.Name, func(t *testing.T) {
			resp, err := cli.PostPets(ctx, PostPetsParams{
				Body: tt.Body,
			})
			require.NoError(t, err)
			body := resp.(PetResponse).Body
			assert.Equal(t, testID, body.ID)
			assert.Equal(t, tt.Body.Name, body.Name)
			assert.Equal(t, tt.Body.Tag, body.Tag)
			assert.Equal(t, tt.Body.Tago, body.Tago)
		})
	}
}
