package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vkd/goag"
)

func TestGenerateTests(t *testing.T) {
	err := goag.GenerateDir("./", "test", "openapi.yaml", "")
	require.NoError(t, err)
}
