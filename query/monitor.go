package query

import (
	"context"

	monitortypes "github.com/babylonchain/babylon/x/monitor/types"
	"github.com/cosmos/cosmos-sdk/client"
)

func (c *Client) QueryMonitor(f func(ctx context.Context, queryClient monitortypes.QueryClient)) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.Client}
	queryClient := monitortypes.NewQueryClient(clientCtx)

	f(ctx, queryClient)
}

// EndedEpochBTCHeight queries the tip height of BTC light client at epoch ends
func (c *Client) EndedEpochBTCHeight(epochNum uint64) (*monitortypes.QueryEndedEpochBtcHeightResponse, error) {
	var (
		resp *monitortypes.QueryEndedEpochBtcHeightResponse
		err  error
	)
	c.QueryMonitor(func(ctx context.Context, queryClient monitortypes.QueryClient) {
		req := &monitortypes.QueryEndedEpochBtcHeightRequest{
			EpochNum: epochNum,
		}
		resp, err = queryClient.EndedEpochBtcHeight(ctx, req)
	})

	return resp, err
}

// ReportedCheckpointBTCHeight queries the tip height of BTC light client when a given checkpoint is reported
func (c *Client) ReportedCheckpointBTCHeight(hashStr string) (*monitortypes.QueryReportedCheckpointBtcHeightResponse, error) {
	var (
		resp *monitortypes.QueryReportedCheckpointBtcHeightResponse
		err  error
	)
	c.QueryMonitor(func(ctx context.Context, queryClient monitortypes.QueryClient) {
		req := &monitortypes.QueryReportedCheckpointBtcHeightRequest{
			CkptHash: hashStr,
		}
		resp, err = queryClient.ReportedCheckpointBtcHeight(ctx, req)
	})

	return resp, err
}
