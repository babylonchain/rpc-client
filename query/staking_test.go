package query_test

import (
	"testing"

	"github.com/babylonchain/babylon/app"
	"github.com/babylonchain/rpc-client/config"
	"github.com/babylonchain/rpc-client/query"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/stretchr/testify/require"
)

// TestStakingParams sets up a vanilla Cosmos SDK chain and tests
// the RPC call to StakingParams
// This ensures that the clientCtx is functional and can be used for
// creating QueryClient for all modules
func TestStakingParams(t *testing.T) {
	// setup virtual network
	cfg := network.DefaultConfig()
	encodingCfg := app.MakeTestEncodingConfig()
	cfg.InterfaceRegistry = encodingCfg.InterfaceRegistry
	cfg.TxConfig = encodingCfg.TxConfig
	cfg.NumValidators = 1
	cfg.RPCAddress = "tcp://0.0.0.0:26657"
	testNetwork, err := network.New(t, t.TempDir(), cfg)
	require.NoError(t, err)
	defer testNetwork.Cleanup()

	err = testNetwork.WaitForNextBlock()
	require.NoError(t, err)

	clientCfg := config.DefaultBabylonConfig()
	client, err := query.New(&clientCfg)
	require.NoError(t, err)
	params, err := client.StakingParams()
	require.NoError(t, err)
	t.Log(params)
}
