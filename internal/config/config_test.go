package config_test

import (
	"os"
	"testing"

	"github.com/Blancduman/banners-rotation/internal/config"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	env := map[string]string{
		"APP_NAME":  "twirler",
		"LOG_LEVEL": "error",
	}

	for k, v := range env {
		err := os.Setenv(k, v)
		require.NoError(t, err)
	}

	conf, err := config.Load()
	require.NoError(t, err)

	lvl, err := conf.LogLevel()
	require.NoError(t, err)
	require.Equal(t, lvl, zerolog.ErrorLevel)
}
