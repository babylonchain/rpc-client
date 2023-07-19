package query

import (
	"context"

	finalitytypes "github.com/babylonchain/babylon/x/finality/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// QueryFinality queries the Finality module of the Babylon node according to the given function
func (c *QueryClient) QueryFinality(f func(ctx context.Context, queryClient finalitytypes.QueryClient) error) error {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.RPCClient}
	queryClient := finalitytypes.NewQueryClient(clientCtx)

	return f(ctx, queryClient)
}

// VotesAtHeight queries the Finality module to get signature set at a given babylon block height
func (c *QueryClient) VotesAtHeight(height uint64) (*finalitytypes.QueryVotesAtHeightResponse, error) {
	var resp *finalitytypes.QueryVotesAtHeightResponse
	err := c.QueryFinality(func(ctx context.Context, queryClient finalitytypes.QueryClient) error {
		var err error
		req := &finalitytypes.QueryVotesAtHeightRequest{
			Height: height,
		}
		resp, err = queryClient.VotesAtHeight(ctx, req)
		return err
	})

	return resp, err
}

// ListBlocks queries the Finality module to get list of finalized or non-finalized blocks
func (c *QueryClient) ListBlocks(finalized bool, pagination *sdkquerytypes.PageRequest) (*finalitytypes.QueryListBlocksResponse, error) {
	var resp *finalitytypes.QueryListBlocksResponse
	err := c.QueryFinality(func(ctx context.Context, queryClient finalitytypes.QueryClient) error {
		var err error
		req := &finalitytypes.QueryListBlocksRequest{
			Finalized:  finalized,
			Pagination: pagination,
		}
		resp, err = queryClient.ListBlocks(ctx, req)
		return err
	})

	return resp, err
}
