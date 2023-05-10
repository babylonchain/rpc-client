package config_test

import (
	bbn "github.com/babylonchain/babylon/app"
	"github.com/cosmos/cosmos-sdk/types/module"
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
	lensConfigModulesMap := make(map[string]module.AppModuleBasic, len(lensConfig.Modules))
	for i := range lensConfig.Modules {
		lensConfigModulesMap[lensConfig.Modules[i].Name()] = lensConfig.Modules[i]
	}
	require.Equal(t, len(bbn.ModuleBasics), len(lensConfig.Modules))
	for modName := range lensConfigModulesMap {
		require.Equal(t, bbn.ModuleBasics[modName].Name(), lensConfigModulesMap[modName].Name())
	}
}
