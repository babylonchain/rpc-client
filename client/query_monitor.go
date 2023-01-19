package client

import (
	monitortypes "github.com/babylonchain/babylon/x/monitor/types"
	"github.com/strangelove-ventures/lens/client/query"
)

// QueryFinishedEpochBtcHeight queries the tip height of BTC light client at epoch ends
func (c *Client) QueryFinishedEpochBtcHeight(epochNum uint64) (uint64, error) {
	var (
		btcHeight uint64
		err       error
	)

	q := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	queryClient := monitortypes.NewQueryClient(c.ChainClient)
	req := &monitortypes.QueryFinishedEpochBtcHeightRequest{EpochNum: epochNum}
	resp, err := queryClient.FinishedEpochBtcHeight(ctx, req)
	if err != nil {
		return btcHeight, err
	}

	btcHeight = resp.BtcLightClientHeight

	return btcHeight, nil
}

// QueryReportedCheckpointBtcHeight queries the tip height of BTC light client when a given checkpoint is reported
func (c *Client) QueryReportedCheckpointBtcHeight(hash []byte) (uint64, error) {
	var (
		btcHeight uint64
		err       error
	)

	q := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := q.GetQueryContext()
	defer cancel()

	queryClient := monitortypes.NewQueryClient(c.ChainClient)
	req := &monitortypes.QueryReportedCheckpointBtcHeightRequest{CkptHash: hash}
	resp, err := queryClient.ReportedCheckpointBtcHeight(ctx, req)
	if err != nil {
		return btcHeight, err
	}

	btcHeight = resp.BtcLightClientHeight

	return btcHeight, nil
}
