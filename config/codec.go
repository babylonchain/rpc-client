package config

import (
	appparams "github.com/babylonchain/babylon/app/params"
	"github.com/babylonchain/babylon/x/btccheckpoint"
	"github.com/babylonchain/babylon/x/btclightclient"
	"github.com/babylonchain/babylon/x/checkpointing"
	"github.com/babylonchain/babylon/x/epoching"
	"github.com/babylonchain/babylon/x/monitor"
	"github.com/babylonchain/babylon/x/zoneconcierge"
	"github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
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
	ibc.AppModuleBasic{},
	ibctm.AppModuleBasic{},
	transfer.AppModuleBasic{},
)

// GetEncodingConfig returns EncodingConfig for Babylon
func GetEncodingConfig() appparams.EncodingConfig {
	lensCdc := lensclient.MakeCodec(ModuleBasics, []string{})
	return appparams.EncodingConfig{
		InterfaceRegistry: lensCdc.InterfaceRegistry,
		Marshaler:         lensCdc.Marshaler,
		TxConfig:          lensCdc.TxConfig,
		Amino:             lensCdc.Amino,
	}
}
