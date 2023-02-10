package query

import (
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	"github.com/babylonchain/rpc-client/client"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// BTCLightClientParams queries btclightclient module's parameters via ChainClient
func BTCLightClientParams(c *client.Client) (*btclctypes.Params, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := &btclctypes.QueryParamsRequest{}
	resp, err := queryClient.Params(ctx, req)
	if err != nil {
		return &btclctypes.Params{}, err
	}
	return &resp.Params, nil
}

// BTCHeaderChainTip queries hash/height of the latest BTC block in the btclightclient module
func BTCHeaderChainTip(c *client.Client) (*btclctypes.QueryTipResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := &btclctypes.QueryTipRequest{}
	resp, err := queryClient.Tip(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// BTCHeader queries the base BTC header of the btclightclient module
func BTCHeader(c *client.Client) (*btclctypes.QueryBaseHeaderResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := &btclctypes.QueryBaseHeaderRequest{}
	resp, err := queryClient.BaseHeader(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ContainsBTCBlock queries the btclightclient module for the existence of a block hash
func ContainsBTCBlock(c *client.Client, blockHash *chainhash.Hash) (*btclctypes.QueryContainsBytesResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := btclctypes.QueryContainsBytesRequest{Hash: blockHash.CloneBytes()}
	resp, err := queryClient.ContainsBytes(ctx, &req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// BTCMainChain queries the btclightclient module for the BTC canonical chain
func BTCMainChain(c *client.Client, pagination *sdkquerytypes.PageRequest) (*btclctypes.QueryMainChainResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := &btclctypes.QueryMainChainRequest{
		Pagination: pagination,
	}

	resp, err := queryClient.MainChain(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
