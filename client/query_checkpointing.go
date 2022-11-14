package client

import (
	checkpointingtypes "github.com/babylonchain/babylon/x/checkpointing/types"
	"github.com/strangelove-ventures/lens/client/query"
)

func (c *Client) QueryRawCheckpoint(epochNumber uint64) (*checkpointingtypes.RawCheckpointWithMeta, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := checkpointingtypes.NewQueryClient(c.ChainClient)
	req := &checkpointingtypes.QueryRawCheckpointRequest{
		EpochNum: epochNumber,
	}
	resp, err := queryClient.RawCheckpoint(ctx, req)
	if err != nil {
		return &checkpointingtypes.RawCheckpointWithMeta{}, err
	}
	return resp.RawCheckpoint, nil
}

func (c *Client) QueryRawCheckpointList(status checkpointingtypes.CheckpointStatus) ([]*checkpointingtypes.RawCheckpointWithMeta, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := checkpointingtypes.NewQueryClient(c.ChainClient)
	req := &checkpointingtypes.QueryRawCheckpointListRequest{
		Status:     status,
		Pagination: query.Options.Pagination, // TODO: parameterise/customise pagination options
	}
	resp, err := queryClient.RawCheckpointList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.RawCheckpoints, nil
}
