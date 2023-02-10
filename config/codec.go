package config

import (
	appparams "github.com/babylonchain/babylon/app/params"
	"github.com/babylonchain/babylon/x/btccheckpoint"
	"github.com/babylonchain/babylon/x/btclightclient"
	"github.com/babylonchain/babylon/x/checkpointing"
	"github.com/babylonchain/babylon/x/epoching"
	"github.com/babylonchain/babylon/x/monitor"
	"github.com/babylonchain/babylon/x/zoneconcierge"
	lensclient "github.com/strangelove-ventures/lens/client"
)

// ModuleBasics is the list of modules used in Babylon
// necessary for serialising/deserialising Babylon messages/queries
var ModuleBasics = append(
	lensclient.ModuleBasics,
	epoching.AppModuleBasic{},
	checkpointing.AppModuleBasic{},
	btclightclient.AppModuleBasic{},
	btccheckpoint.AppModuleBasic{},
	zoneconcierge.AppModuleBasic{},
	monitor.AppModuleBasic{},
)

// GetEncodingConfig returns EncodingConfig for Babylon
func GetEncodingConfig() appparams.EncodingConfig {
	lensCdc := lensclient.MakeCodec(ModuleBasics)
	return appparams.EncodingConfig{
		InterfaceRegistry: lensCdc.InterfaceRegistry,
		Marshaler:         lensCdc.Marshaler,
		TxConfig:          lensCdc.TxConfig,
		Amino:             lensCdc.Amino,
	}
}
