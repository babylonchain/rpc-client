package query

import (
	checkpointingtypes "github.com/babylonchain/babylon/x/checkpointing/types"
	"github.com/babylonchain/rpc-client/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// RawCheckpoint queries the checkpointing module for the raw checkpoint for an epoch number
func RawCheckpoint(c *client.Client, epochNumber uint64) (*checkpointingtypes.QueryRawCheckpointResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := checkpointingtypes.NewQueryClient(c.ChainClient)
	req := &checkpointingtypes.QueryRawCheckpointRequest{
		EpochNum: epochNumber,
	}
	resp, err := queryClient.RawCheckpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RawCheckpointList queries the checkpointing module for a list of raw checkpoints
func RawCheckpointList(c *client.Client, status checkpointingtypes.CheckpointStatus, pagination *sdkquerytypes.PageRequest) (*checkpointingtypes.QueryRawCheckpointListResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := checkpointingtypes.NewQueryClient(c.ChainClient)
	req := &checkpointingtypes.QueryRawCheckpointListRequest{
		Status:     status,
		Pagination: pagination,
	}
	resp, err := queryClient.RawCheckpointList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BlsPublicKeyList queries the checkpointing module for the list of BLS keys for an epoch
func BlsPublicKeyList(c *client.Client, epochNumber uint64) (*checkpointingtypes.QueryBlsPublicKeyListResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := checkpointingtypes.NewQueryClient(c.ChainClient)
	req := &checkpointingtypes.QueryBlsPublicKeyListRequest{
		EpochNum: epochNumber,
	}
	resp, err := queryClient.BlsPublicKeyList(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// EpochStatusCount queries the checkpointing module for the status of the latest `epochCount` epochs`
func EpochStatusCount(c *client.Client, epochCount uint64) (*checkpointingtypes.QueryRecentEpochStatusCountResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := checkpointingtypes.NewQueryClient(c.ChainClient)
	req := &checkpointingtypes.QueryRecentEpochStatusCountRequest{
		EpochCount: epochCount,
	}

	resp, err := queryClient.RecentEpochStatusCount(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LatestEpochFromStatus queries the checkpointing module for the last checkpoint with a particular status
func LatestEpochFromStatus(c *client.Client, status checkpointingtypes.CheckpointStatus) (*checkpointingtypes.QueryLastCheckpointWithStatusResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := checkpointingtypes.NewQueryClient(c.ChainClient)
	req := &checkpointingtypes.QueryLastCheckpointWithStatusRequest{
		Status: status,
	}

	resp, err := queryClient.LastCheckpointWithStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
