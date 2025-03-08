package test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponents(t *testing.T) {
	testName := "test_name"
	testPayload := `"payload-001"`

	ctx := context.Background()

	api := API{
		PostPetsHandler: func(ctx context.Context, r PostPetsRequest) PostPetsResponse {
			params, err := r.Parse()
			if err != nil {
				t.Fatalf("Parse params: %v", err)
			}
			return NewPostPetsResponse200JSON(Pets{Pet{
				Name:    params.Body.Name,
				Payload: params.Body.Payload,
			}})
		},
	}

	client := api.LocalClient()

	resp, err := client.PostPets(ctx, PostPetsParams{
		Body: CreatePetJSON{
			Name:    testName,
			Payload: json.RawMessage(testPayload),
		},
	})
	require.NoError(t, err)

	assert.Equal(t, testName, resp.(PostPetsResponse200JSON).Body[0].Name)
	assert.Equal(t, testPayload, string(resp.(PostPetsResponse200JSON).Body[0].Payload))
}
