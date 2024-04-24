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
		} `json:"cats"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &v)
	require.NoError(t, err)
	assert.Equal(t, 1, v.Length)
	assert.Equal(t, "mike", v.Cats[0].Name)
	assert.Equal(t, "ok", v.Cats[0].Status, w.Body.String())
}
