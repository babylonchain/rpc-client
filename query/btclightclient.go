package query

import (
	"context"

	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/cosmos/cosmos-sdk/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

func (c *QueryClient) QueryBTCLightclient(f func(ctx context.Context, queryClient btclctypes.QueryClient)) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.RPCClient}
	queryClient := btclctypes.NewQueryClient(clientCtx)

	f(ctx, queryClient)
}

// BTCLightClientParams queries btclightclient module's parameters via ChainClient
func (c *QueryClient) BTCLightClientParams() (*btclctypes.Params, error) {
	var (
		resp *btclctypes.QueryParamsResponse
		err  error
	)
	c.QueryBTCLightclient(func(ctx context.Context, queryClient btclctypes.QueryClient) {
		req := &btclctypes.QueryParamsRequest{}
		resp, err = queryClient.Params(ctx, req)
	})

	if err != nil {
		return nil, err
	}
	return &resp.Params, err
}

// BTCHeaderChainTip queries hash/height of the latest BTC block in the btclightclient module
func (c *QueryClient) BTCHeaderChainTip() (*btclctypes.QueryTipResponse, error) {
	var (
		resp *btclctypes.QueryTipResponse
		err  error
	)
	c.QueryBTCLightclient(func(ctx context.Context, queryClient btclctypes.QueryClient) {
		req := &btclctypes.QueryTipRequest{}
		resp, err = queryClient.Tip(ctx, req)
	})

	return resp, err
}

// BTCHeader queries the base BTC header of the btclightclient module
func (c *QueryClient) BTCHeader() (*btclctypes.QueryBaseHeaderResponse, error) {
	var (
		resp *btclctypes.QueryBaseHeaderResponse
		err  error
	)
	c.QueryBTCLightclient(func(ctx context.Context, queryClient btclctypes.QueryClient) {
		req := &btclctypes.QueryBaseHeaderRequest{}
		resp, err = queryClient.BaseHeader(ctx, req)
	})

	return resp, err
}

// ContainsBTCBlock queries the btclightclient module for the existence of a block hash
func (c *QueryClient) ContainsBTCBlock(blockHash *chainhash.Hash) (*btclctypes.QueryContainsBytesResponse, error) {
	var (
		resp *btclctypes.QueryContainsBytesResponse
		err  error
	)
	c.QueryBTCLightclient(func(ctx context.Context, queryClient btclctypes.QueryClient) {
		req := &btclctypes.QueryContainsBytesRequest{
			Hash: blockHash.CloneBytes(),
		}
		resp, err = queryClient.ContainsBytes(ctx, req)
	})

	return resp, err
}

// BTCMainChain queries the btclightclient module for the BTC canonical chain
func (c *QueryClient) BTCMainChain(pagination *sdkquerytypes.PageRequest) (*btclctypes.QueryMainChainResponse, error) {
	var (
		resp *btclctypes.QueryMainChainResponse
		err  error
	)
	c.QueryBTCLightclient(func(ctx context.Context, queryClient btclctypes.QueryClient) {
		req := &btclctypes.QueryMainChainRequest{
			Pagination: pagination,
		}
		resp, err = queryClient.MainChain(ctx, req)
	})

	return resp, err
}
