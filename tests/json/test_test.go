package test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	ctx := context.Background()

	api := API{
		PostPetsHandler: func(ctx context.Context, r PostPetsRequest) PostPetsResponse {
			params, err := r.Parse()
			require.NoError(t, err)
			return NewPetResponse(Pet{
				ID:     1,
				NewPet: params.Body,
			})
		},
	}
	client := api.TestClient()

	for _, tt := range []struct {
		Name   string
		Params PostPetsParams
	}{
		{"all",
			PostPetsParams{
				Body: NewPet{
					Name:     "test_name",
					Tag:      Pointer("test_tag"),
					Tago:     Just(Pointer("test_tago")),
					Birthday: time.Date(2005, 12, 13, 14, 31, 11, 0, time.UTC),
					Metadata: Just(Metadata{
						Owner: "test_metadata_owner",
						Tags:  Just(Tags{Tag{Name: "test_metadata_tags_0_name", Value: "test_metadata_tags_0_value"}}),
					}),
				},
			},
		},
	} {
		t.Run(tt.Name, func(t *testing.T) {
			resp, err := client.PostPets(ctx, tt.Params)
			require.NoError(t, err)

			body := resp.(PetResponse).Body
			assert.Equal(t, Pet{ID: 1, NewPet: tt.Params.Body}, body)
		})
	}
}
