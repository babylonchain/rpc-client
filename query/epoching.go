package query

import (
	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	"github.com/babylonchain/rpc-client/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// EpochingParams queries epoching module's parameters via ChainClient
func EpochingParams(c *client.Client) (*epochingtypes.QueryParamsResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryParamsRequest{}
	resp, err := queryClient.Params(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CurrentEpoch queries the current epoch number via ChainClient
func CurrentEpoch(c *client.Client) (*epochingtypes.QueryCurrentEpochResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryCurrentEpochRequest{}
	resp, err := queryClient.CurrentEpoch(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// EpochsInfoForEpochRange queries the epoching module for epochs in the given range
func EpochsInfoForEpochRange(c *client.Client, startEpoch uint64, endEpoch uint64) (*epochingtypes.QueryEpochsInfoResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryEpochsInfoRequest{
		StartEpoch: startEpoch,
		EndEpoch:   endEpoch,
	}

	resp, err := queryClient.EpochsInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// EpochsInfo queries the epoching module for the maintained epochs
func EpochsInfo(c *client.Client, pagination *sdkquerytypes.PageRequest) (*epochingtypes.QueryEpochsInfoResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryEpochsInfoRequest{
		Pagination: pagination,
	}

	resp, err := queryClient.EpochsInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LatestEpochMsgs queries the epoching module for the latest messages maintained in its delayed
// staking queue until a specified endEpoch.
func LatestEpochMsgs(c *client.Client, endEpoch uint64, epochCount uint64, pagination *sdkquerytypes.PageRequest) (*epochingtypes.QueryLatestEpochMsgsResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryLatestEpochMsgsRequest{
		EndEpoch:   endEpoch,
		EpochCount: epochCount,
		Pagination: pagination,
	}

	resp, err := queryClient.LatestEpochMsgs(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DelegationLifecycle queries the epoching module for the lifecycle of a delegator.
func DelegationLifecycle(c *client.Client, delegator string) (*epochingtypes.QueryDelegationLifecycleResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := epochingtypes.NewQueryClient(c.ChainClient)
	req := &epochingtypes.QueryDelegationLifecycleRequest{
		DelAddr: delegator,
	}

	resp, err := queryClient.DelegationLifecycle(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
