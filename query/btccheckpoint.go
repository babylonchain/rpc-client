package query

import (
	"context"

	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

func (c *Client) QueryBTCCheckpoint(f func(ctx context.Context, queryClient btcctypes.QueryClient)) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.Client}
	queryClient := btcctypes.NewQueryClient(clientCtx)

	f(ctx, queryClient)
}

// BTCCheckpointParams queries btccheckpoint module's parameters via ChainClient
func (c *Client) BTCCheckpointParams() (*btcctypes.Params, error) {
	var (
		resp *btcctypes.QueryParamsResponse
		err  error
	)
	c.QueryBTCCheckpoint(func(ctx context.Context, queryClient btcctypes.QueryClient) {
		req := &btcctypes.QueryParamsRequest{}
		resp, err = queryClient.Params(ctx, req)
	})

	if err != nil {
		return nil, err
	}
	return &resp.Params, err
}

// BTCPositionAtEpoch queries btccheckpoint module for the Bitcoin position of an epoch
func (c *Client) BTCPositionAtEpoch(epochNumber uint64) (*btcctypes.QueryBtcCheckpointInfoResponse, error) {
	var (
		resp *btcctypes.QueryBtcCheckpointInfoResponse
		err  error
	)
	c.QueryBTCCheckpoint(func(ctx context.Context, queryClient btcctypes.QueryClient) {
		req := &btcctypes.QueryBtcCheckpointInfoRequest{
			EpochNum: epochNumber,
		}
		resp, err = queryClient.BtcCheckpointInfo(ctx, req)
	})

	return resp, err
}

// BTCPositionForEpochRange queries btccheckpoint module for the Bitcoin position of an epoch range
func (c *Client) BTCPositionForEpochRange(startEpoch uint64, endEpoch uint64, pagination *sdkquerytypes.PageRequest) (*btcctypes.QueryBtcCheckpointsInfoResponse, error) {
	var (
		resp *btcctypes.QueryBtcCheckpointsInfoResponse
		err  error
	)
	c.QueryBTCCheckpoint(func(ctx context.Context, queryClient btcctypes.QueryClient) {
		req := &btcctypes.QueryBtcCheckpointsInfoRequest{
			StartEpoch: startEpoch,
			EndEpoch:   endEpoch,
			Pagination: pagination,
		}
		resp, err = queryClient.BtcCheckpointsInfo(ctx, req)
	})

	return resp, err
}
