package client

import (
	btclctypes "github.com/babylonchain/babylon/x/btclightclient/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/strangelove-ventures/lens/client/query"
)

// QueryBTCLightclientParams queries btclightclient module's parameters via ChainClient
func (c *Client) QueryBTCLightclientParams() (*btclctypes.Params, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := &btclctypes.QueryParamsRequest{}
	resp, err := queryClient.Params(ctx, req)
	if err != nil {
		return &btclctypes.Params{}, err
	}
	return &resp.Params, nil
}

// QueryHeaderChainTip queries hash/height of the latest BTC block in the btclightclient module
func (c *Client) QueryHeaderChainTip() (*chainhash.Hash, uint64, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := &btclctypes.QueryTipRequest{}
	resp, err := queryClient.Tip(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	return resp.Header.Hash.ToChainhash(), resp.Header.Height, nil
}

func (c *Client) QueryBaseHeader() (*wire.BlockHeader, uint64, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)

	req := &btclctypes.QueryBaseHeaderRequest{}
	resp, err := queryClient.BaseHeader(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	header := resp.Header.Header.ToBlockHeader()
	height := resp.Header.Height

	return header, height, nil
}

func (c *Client) QueryContainsBlock(blockHash *chainhash.Hash) (bool, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := btclctypes.NewQueryClient(c.ChainClient)
	req := btclctypes.QueryContainsBytesRequest{Hash: blockHash.CloneBytes()}
	resp, err := queryClient.ContainsBytes(ctx, &req)
	if err != nil {
		return false, err
	}

	return resp.Contains, nil
}
