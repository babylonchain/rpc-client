package query

import (
	"context"

	epochingtypes "github.com/babylonchain/babylon/x/epoching/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

func (c *Client) QueryEpoching(f func(ctx context.Context, queryClient epochingtypes.QueryClient)) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.Client}
	queryClient := epochingtypes.NewQueryClient(clientCtx)

	f(ctx, queryClient)
}

// EpochingParams queries epoching module's parameters via ChainClient
func (c *Client) EpochingParams() (*epochingtypes.QueryParamsResponse, error) {
	var (
		resp *epochingtypes.QueryParamsResponse
		err  error
	)
	c.QueryEpoching(func(ctx context.Context, queryClient epochingtypes.QueryClient) {
		req := &epochingtypes.QueryParamsRequest{}
		resp, err = queryClient.Params(ctx, req)
	})

	return resp, err
}

// CurrentEpoch queries the current epoch number via ChainClient
func (c *Client) CurrentEpoch() (*epochingtypes.QueryCurrentEpochResponse, error) {
	var (
		resp *epochingtypes.QueryCurrentEpochResponse
		err  error
	)
	c.QueryEpoching(func(ctx context.Context, queryClient epochingtypes.QueryClient) {
		req := &epochingtypes.QueryCurrentEpochRequest{}
		resp, err = queryClient.CurrentEpoch(ctx, req)
	})

	return resp, err
}

// EpochsInfoForEpochRange queries the epoching module for epochs in the given range
func (c *Client) EpochsInfoForEpochRange(startEpoch uint64, endEpoch uint64) (*epochingtypes.QueryEpochsInfoResponse, error) {
	var (
		resp *epochingtypes.QueryEpochsInfoResponse
		err  error
	)
	c.QueryEpoching(func(ctx context.Context, queryClient epochingtypes.QueryClient) {
		req := &epochingtypes.QueryEpochsInfoRequest{
			StartEpoch: startEpoch,
			EndEpoch:   endEpoch,
		}
		resp, err = queryClient.EpochsInfo(ctx, req)
	})

	return resp, err
}

// EpochsInfo queries the epoching module for the maintained epochs
func (c *Client) EpochsInfo(pagination *sdkquerytypes.PageRequest) (*epochingtypes.QueryEpochsInfoResponse, error) {
	var (
		resp *epochingtypes.QueryEpochsInfoResponse
		err  error
	)
	c.QueryEpoching(func(ctx context.Context, queryClient epochingtypes.QueryClient) {
		req := &epochingtypes.QueryEpochsInfoRequest{
			Pagination: pagination,
		}
		resp, err = queryClient.EpochsInfo(ctx, req)
	})

	return resp, err
}

// LatestEpochMsgs queries the epoching module for the latest messages maintained in its delayed
// staking queue until a specified endEpoch.
func (c *Client) LatestEpochMsgs(endEpoch uint64, epochCount uint64, pagination *sdkquerytypes.PageRequest) (*epochingtypes.QueryLatestEpochMsgsResponse, error) {
	var (
		resp *epochingtypes.QueryLatestEpochMsgsResponse
		err  error
	)
	c.QueryEpoching(func(ctx context.Context, queryClient epochingtypes.QueryClient) {
		req := &epochingtypes.QueryLatestEpochMsgsRequest{
			EndEpoch:   endEpoch,
			EpochCount: epochCount,
			Pagination: pagination,
		}
		resp, err = queryClient.LatestEpochMsgs(ctx, req)
	})

	return resp, err
}

// DelegationLifecycle queries the epoching module for the lifecycle of a delegator.
func (c *Client) DelegationLifecycle(delegator string) (*epochingtypes.QueryDelegationLifecycleResponse, error) {
	var (
		resp *epochingtypes.QueryDelegationLifecycleResponse
		err  error
	)
	c.QueryEpoching(func(ctx context.Context, queryClient epochingtypes.QueryClient) {
		req := &epochingtypes.QueryDelegationLifecycleRequest{
			DelAddr: delegator,
		}
		resp, err = queryClient.DelegationLifecycle(ctx, req)
	})

	return resp, err
}
