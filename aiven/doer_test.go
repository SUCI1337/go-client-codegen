package aiven

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	token := os.Getenv("AIVEN_TOKEN")
	if token == "" {
		t.Skip("token is required for the test")
	}

	client, err := NewClient(DebugOpt())
	require.NoError(t, err)

	ctx := context.Background()
	tokens, err := client.User.AccessTokenList(ctx)
	require.NoError(t, err)

	found := 0
	for _, to := range tokens {
		if strings.HasPrefix(token, to.TokenPrefix) {
			found++
		}
	}
	assert.Equal(t, 1, found)
}
