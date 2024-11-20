package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestConfig_LoadConfig test loading configuration from file
func TestConfig_LoadConfig(t *testing.T) {
	loadedConfig, err := LoadConfig("../../")
	require.NoError(t, err)
	require.NotEmpty(t, loadedConfig)
}

// TestConfig_GetDatabaseSource test getting database source
func TestConfig_GetDatabaseSource(t *testing.T) {
	loadedConfig, err := LoadConfig("../../")
	require.NoError(t, err)
	require.NotEmpty(t, loadedConfig.GetDatabaseSource())
}

// TestConfig_GetServerAddress test getting server address
func TestConfig_GetServerAddress(t *testing.T) {
	loadedConfig, err := LoadConfig("../../")
	require.NoError(t, err)
	serverAddress := loadedConfig.GetServerAddress()
	require.NotNil(t, serverAddress)
}
