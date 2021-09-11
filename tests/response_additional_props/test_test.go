package test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseSchema(t *testing.T) {
	handler := GetPetHandlerFunc(func(_ GetPetParamsParser) GetPetResponser {
		return GetPetResponse200JSON(GetPetResponse200JSONBody{
			Groups: map[string]Pets{
				"cats": {Pet{Name: "mike"}, Pet{Name: "alex"}},
			},
		})
	})

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"groups\":{\"cats\":[{\"name\":\"mike\"},{\"name\":\"alex\"}]}}\n", w.Body.String())
}
