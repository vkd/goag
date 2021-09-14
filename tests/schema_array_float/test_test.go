package test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaArrayFloat(t *testing.T) {
	h := GetPetsIDsHandlerFunc(func(_ GetPetsIDsRequester) GetPetsIDsResponder {
		return GetPetsIDsResponse200JSON([]float64{0.8})
	})

	w := httptest.NewRecorder()
	h.ServeHTTP(w, httptest.NewRequest("GET", "/pets", nil))

	require.Equal(t, 200, w.Code)
	var out []float64
	err := json.Unmarshal(w.Body.Bytes(), &out)
	require.NoError(t, err)
	require.Len(t, out, 1)
	assert.Equal(t, float64(0.8), out[0])
}
