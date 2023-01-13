package config_test

import (
	"testing"

	"github.com/babylonchain/rpc-client/config"
	"github.com/stretchr/testify/require"
)

// TestBabylonConfig ensures that the default Babylon config is valid
func TestBabylonConfig(t *testing.T) {
	defaultConfig := config.DefaultBabylonConfig()
	err := defaultConfig.Validate()
	require.NoError(t, err)

	// ensure the unwrapped config has Babylon codec formats
	lensConfig := defaultConfig.Unwrap()
	require.Equal(t, len(config.ModuleBasics), len(lensConfig.Modules))
	for i := range config.ModuleBasics {
		require.Equal(t, config.ModuleBasics[i].Name(), lensConfig.Modules[i].Name())
	}
}
