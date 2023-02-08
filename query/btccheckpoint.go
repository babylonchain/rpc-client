package query

import (
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	"github.com/babylonchain/rpc-client/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// BTCCheckpointParams queries btccheckpoint module's parameters via ChainClient
func BTCCheckpointParams(c *client.Client) (*btcctypes.Params, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btcctypes.NewQueryClient(c.ChainClient)
	req := &btcctypes.QueryParamsRequest{}
	resp, err := queryClient.Params(ctx, req)
	if err != nil {
		return &btcctypes.Params{}, err
	}
	return &resp.Params, nil
}

// BTCPositionAtEpoch queries btccheckpoint module for the Bitcoin position of an epoch
func BTCPositionAtEpoch(c *client.Client, epochNumber uint64) (*btcctypes.QueryBtcCheckpointInfoResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btcctypes.NewQueryClient(c.ChainClient)
	req := &btcctypes.QueryBtcCheckpointInfoRequest{
		EpochNum: epochNumber,
	}

	resp, err := queryClient.BtcCheckpointInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BTCPositionForEpochRange queries btccheckpoint module for the Bitcoin position of an epoch range
func BTCPositionForEpochRange(c *client.Client, startEpoch uint64, endEpoch uint64, pagination *sdkquerytypes.PageRequest) (*btcctypes.QueryBtcCheckpointsInfoResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := btcctypes.NewQueryClient(c.ChainClient)
	req := &btcctypes.QueryBtcCheckpointsInfoRequest{
		StartEpoch: startEpoch,
		EndEpoch:   endEpoch,
		Pagination: pagination,
	}

	resp, err := queryClient.BtcCheckpointsInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
