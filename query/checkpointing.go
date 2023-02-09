package query

import (
	"context"

	checkpointingtypes "github.com/babylonchain/babylon/x/checkpointing/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// QueryCheckpointing queries the Checkpointing module of the Babylon node
// according to the given function
func (c *QueryClient) QueryCheckpointing(f func(ctx context.Context, queryClient checkpointingtypes.QueryClient)) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.RPCClient}
	queryClient := checkpointingtypes.NewQueryClient(clientCtx)

	f(ctx, queryClient)
}

// RawCheckpoint queries the checkpointing module for the raw checkpoint for an epoch number
func (c *QueryClient) RawCheckpoint(epochNumber uint64) (*checkpointingtypes.QueryRawCheckpointResponse, error) {
	var (
		resp *checkpointingtypes.QueryRawCheckpointResponse
		err  error
	)
	c.QueryCheckpointing(func(ctx context.Context, queryClient checkpointingtypes.QueryClient) {
		req := &checkpointingtypes.QueryRawCheckpointRequest{
			EpochNum: epochNumber,
		}
		resp, err = queryClient.RawCheckpoint(ctx, req)
	})

	return resp, err
}

// RawCheckpointList queries the checkpointing module for a list of raw checkpoints
func (c *QueryClient) RawCheckpointList(status checkpointingtypes.CheckpointStatus, pagination *sdkquerytypes.PageRequest) (*checkpointingtypes.QueryRawCheckpointListResponse, error) {
	var (
		resp *checkpointingtypes.QueryRawCheckpointListResponse
		err  error
	)
	c.QueryCheckpointing(func(ctx context.Context, queryClient checkpointingtypes.QueryClient) {
		req := &checkpointingtypes.QueryRawCheckpointListRequest{
			Status:     status,
			Pagination: pagination,
		}
		resp, err = queryClient.RawCheckpointList(ctx, req)
	})

	return resp, err
}

// BlsPublicKeyList queries the checkpointing module for the list of BLS keys for an epoch
func (c *QueryClient) BlsPublicKeyList(epochNumber uint64, pagination *sdkquerytypes.PageRequest) (*checkpointingtypes.QueryBlsPublicKeyListResponse, error) {
	var (
		resp *checkpointingtypes.QueryBlsPublicKeyListResponse
		err  error
	)
	c.QueryCheckpointing(func(ctx context.Context, queryClient checkpointingtypes.QueryClient) {
		req := &checkpointingtypes.QueryBlsPublicKeyListRequest{
			EpochNum:   epochNumber,
			Pagination: pagination,
		}
		resp, err = queryClient.BlsPublicKeyList(ctx, req)
	})

	return resp, err
}

// EpochStatusCount queries the checkpointing module for the status of the latest `epochCount` epochs`
func (c *QueryClient) EpochStatusCount(epochCount uint64) (*checkpointingtypes.QueryRecentEpochStatusCountResponse, error) {
	var (
		resp *checkpointingtypes.QueryRecentEpochStatusCountResponse
		err  error
	)
	c.QueryCheckpointing(func(ctx context.Context, queryClient checkpointingtypes.QueryClient) {
		req := &checkpointingtypes.QueryRecentEpochStatusCountRequest{
			EpochCount: epochCount,
		}
		resp, err = queryClient.RecentEpochStatusCount(ctx, req)
	})

	return resp, err
}

// LatestEpochFromStatus queries the checkpointing module for the last checkpoint with a particular status
func (c *QueryClient) LatestEpochFromStatus(status checkpointingtypes.CheckpointStatus) (*checkpointingtypes.QueryLastCheckpointWithStatusResponse, error) {
	var (
		resp *checkpointingtypes.QueryLastCheckpointWithStatusResponse
		err  error
	)
	c.QueryCheckpointing(func(ctx context.Context, queryClient checkpointingtypes.QueryClient) {
		req := &checkpointingtypes.QueryLastCheckpointWithStatusRequest{
			Status: status,
		}
		resp, err = queryClient.LastCheckpointWithStatus(ctx, req)
	})

	return resp, err
}
