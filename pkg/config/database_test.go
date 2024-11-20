package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestConfig_ConnectToPgxPool test connecting to the database pool
func TestConfig_ConnectToPgxPool(t *testing.T) {
	loadedConfig, err := LoadConfig("../../")
	require.NoError(t, err)
	conn, err := loadedConfig.ConnectToPgxPool()
	require.NoError(t, err)
	require.NotEmpty(t, conn)
}
