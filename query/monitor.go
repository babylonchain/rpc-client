package query

import (
	monitortypes "github.com/babylonchain/babylon/x/monitor/types"
	"github.com/babylonchain/rpc-client/client"
)

// EndedEpochBTCHeight queries the tip height of BTC light client at epoch ends
func EndedEpochBTCHeight(c *client.Client, epochNum uint64) (*monitortypes.QueryEndedEpochBtcHeightResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := monitortypes.NewQueryClient(c.ChainClient)
	req := &monitortypes.QueryEndedEpochBtcHeightRequest{EpochNum: epochNum}
	resp, err := queryClient.EndedEpochBtcHeight(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ReportedCheckpointBTCHeight queries the tip height of BTC light client when a given checkpoint is reported
func ReportedCheckpointBTCHeight(c *client.Client, hashStr string) (*monitortypes.QueryReportedCheckpointBtcHeightResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := monitortypes.NewQueryClient(c.ChainClient)
	req := &monitortypes.QueryReportedCheckpointBtcHeightRequest{CkptHash: hashStr}
	resp, err := queryClient.ReportedCheckpointBtcHeight(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
