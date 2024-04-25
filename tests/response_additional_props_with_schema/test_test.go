package test

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseSchema(t *testing.T) {
	handler := GetPetHandlerFunc(func(_ context.Context, _ GetPetRequest) GetPetResponse {
		return NewGetPetResponse200JSON(GetPetResponse200JSONBody{
			Length: 1,
			AdditionalProperties: map[string]Pets{
				"cats": {
					Pet{
						Name: "mike",
						AdditionalProperties: map[string]json.RawMessage{
							"status": json.RawMessage(`"ok"`),
						},
						Custom: PetCustom(`{"field1": "pet_custom_field"}`),
					},
				},
			},
		})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	var v struct {
		Length int `json:"length"`
		Cats   []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
			Custom struct {
				Field1 string `json:"field1"`
			} `json:"custom"`
		} `json:"cats"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &v)
	require.NoError(t, err, w.Body.String())
	assert.Equal(t, 1, v.Length)
	assert.Equal(t, "mike", v.Cats[0].Name)
	assert.Equal(t, "ok", v.Cats[0].Status, w.Body.String())
	assert.Equal(t, "pet_custom_field", v.Cats[0].Custom.Field1, w.Body.String())
}
