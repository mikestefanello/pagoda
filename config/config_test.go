package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	_, err := GetConfig()
	require.NoError(t, err)

	var env environment
	env = "abc"
	SwitchEnvironment(env)
	cfg, err := GetConfig()
	require.NoError(t, err)
	assert.Equal(t, env, cfg.App.Environment)

	key := cfg.App.EncryptionKey
	assert.Len(t, key, 64)

	cfg, err = GetConfig()
	require.NoError(t, err)
	assert.NotEqual(t, key, cfg.App.EncryptionKey)
}
