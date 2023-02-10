package config

import (
	"github.com/babylonchain/babylon/x/btccheckpoint"
	"github.com/babylonchain/babylon/x/btclightclient"
	"github.com/babylonchain/babylon/x/checkpointing"
	"github.com/babylonchain/babylon/x/epoching"
	"github.com/babylonchain/babylon/x/monitor"
	"github.com/babylonchain/babylon/x/zoneconcierge"
	"github.com/strangelove-ventures/lens/client"
	lensclient "github.com/strangelove-ventures/lens/client"
)

// ModuleBasics is the list of modules used in Babylon
// necessary for serialising/deserialising Babylon messages/queries
var ModuleBasics = append(
	client.ModuleBasics,
	epoching.AppModuleBasic{},
	checkpointing.AppModuleBasic{},
	btclightclient.AppModuleBasic{},
	btccheckpoint.AppModuleBasic{},
	zoneconcierge.AppModuleBasic{},
	monitor.AppModuleBasic{},
)

// GetBabylonCdc returns Codec, i.e., struct for encoding/decoding messages, for Babylon
func GetCodec() lensclient.Codec {
	return lensclient.MakeCodec(ModuleBasics)
}
