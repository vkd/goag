package test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResponseSchema(t *testing.T) {
	handler := GetPetHandlerFunc(func(_ GetPetRequester) GetPetResponser {
		return GetPetResponse200JSON(GetPetResponse200JSONBody{
			Length: 1,
			AdditionalProperties: map[string]Pets{
				"cats": {Pet{Name: "mike"}},
			},
		})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	var v struct {
		Length int  `json:"length"`
		Cats   Pets `json:"cats"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &v)
	require.NoError(t, err)
	assert.Equal(t, 1, v.Length)
	assert.Equal(t, Pets{Pet{Name: "mike"}}, v.Cats)
}
