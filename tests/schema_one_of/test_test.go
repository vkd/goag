package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vkd/goag/tests/schema_one_of/pkg"
)

func TestSchemaAllOf_Get(t *testing.T) {
	ctx := context.Background()

	api := API{
		GetPetHandler: func(_ context.Context, req GetPetRequest) GetPetResponse {
			params, err := req.Parse()
			require.NoError(t, err)

			return NewGetPetResponse200JSON(Resp{
				Pet: params.Body,
			})
		},
	}
	client := api.LocalClient()

	// oneOf0 - Cat
	testCat := "test_cat"
	resp, err := client.GetPet(ctx, GetPetParams{
		Body: NewPetCat(Cat{
			Label: testCat,
		}),
	})
	require.NoError(t, err)

	assert.Equal(t, testCat, resp.(GetPetResponse200JSON).Body.Pet.Cat.Value.Label)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Dog.Value.Name)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf2.Value.ID)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf3.Value)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.OneOf4.Value)

	// oneOf1 - Dog
	testDog := "test_dog"
	resp, err = client.GetPet(ctx, GetPetParams{
		Body: NewPetDog(pkg.Dog{
			Name: testDog,
		}),
	})
	require.NoError(t, err)

	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Cat.Value.Label)
	assert.Equal(t, testDog, resp.(GetPetResponse200JSON).Body.Pet.Dog.Value.Name)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf2.Value.ID)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf3.Value)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.OneOf4.Value)

	// oneOf2 - OneOf2
	testOneOf2ID := int64(44)
	resp, err = client.GetPet(ctx, GetPetParams{
		Body: NewPetOneOf2(PetOneOf2{
			ID: testOneOf2ID,
		}),
	})
	require.NoError(t, err)

	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Cat.Value.Label)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Dog.Value.Name)
	assert.Equal(t, testOneOf2ID, resp.(GetPetResponse200JSON).Body.Pet.OneOf2.Value.ID)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf3.Value)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.OneOf4.Value)

	// oneOf3 - int64
	testOneOf3Int64 := int64(66)
	resp, err = client.GetPet(ctx, GetPetParams{
		Body: NewPetOneOf3(testOneOf3Int64),
	})
	require.NoError(t, err)

	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Cat.Value.Label)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Dog.Value.Name)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf2.Value.ID)
	assert.Equal(t, testOneOf3Int64, resp.(GetPetResponse200JSON).Body.Pet.OneOf3.Value)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.OneOf4.Value)

	// oneOf4 - string
	testOneOf4String := "test_one_of_4_string"
	resp, err = client.GetPet(ctx, GetPetParams{
		Body: NewPetOneOf4(testOneOf4String),
	})
	require.NoError(t, err)

	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Cat.Value.Label)
	assert.Equal(t, "", resp.(GetPetResponse200JSON).Body.Pet.Dog.Value.Name)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf2.Value.ID)
	assert.Equal(t, int64(0), resp.(GetPetResponse200JSON).Body.Pet.OneOf3.Value)
	assert.Equal(t, testOneOf4String, resp.(GetPetResponse200JSON).Body.Pet.OneOf4.Value)
}

func TestSchemaAllOf_Get2(t *testing.T) {
	ctx := context.Background()

	api := API{
		GetPet2Handler: func(_ context.Context, req GetPet2Request) GetPet2Response {
			params, err := req.Parse()
			require.NoError(t, err)

			return NewGetPet2Response200JSON(Resp2{
				Pet: params.Body,
			})
		},
	}
	client := api.LocalClient()

	// oneOf0 - Cat
	testCat := "test_cat"
	resp, err := client.GetPet2(ctx, GetPet2Params{
		Body: NewPet2Cat2(Cat2{
			Name:   pkg.Just(testCat),
			PetType: "cati",
		}),
	})
	require.NoError(t, err)

	assert.Equal(t, testCat, resp.(GetPet2Response200JSON).Body.Pet.Cat2.Value.Name.Value)
	assert.Equal(t, "", resp.(GetPet2Response200JSON).Body.Pet.Dog2.Value.Name.Value)

	// oneOf1 - Dog
	testDog := "test_dog"
	resp, err = client.GetPet2(ctx, GetPet2Params{
		Body: NewPet2Dog2(pkg.Dog2{
			Name:    pkg.Just(testDog),
			PetType: "dogi",
		}),
	})
	require.NoError(t, err)

	assert.Equal(t, "", resp.(GetPet2Response200JSON).Body.Pet.Cat2.Value.Name.Value)
	assert.Equal(t, testDog, resp.(GetPet2Response200JSON).Body.Pet.Dog2.Value.Name.Value)
}
