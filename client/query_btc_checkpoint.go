package client

import (
	"github.com/babylonchain/babylon/types/retry"
	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	"github.com/strangelove-ventures/lens/client/query"
)

// QueryBTCCheckpointParams queries btccheckpoint module's parameters via ChainClient
func (c *Client) QueryBTCCheckpointParams() (*btcctypes.Params, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
	ctx, cancel := query.GetQueryContext()
	defer cancel()

	queryClient := btcctypes.NewQueryClient(c.ChainClient)
	req := &btcctypes.QueryParamsRequest{}
	resp, err := queryClient.Params(ctx, req)
	if err != nil {
		return &btcctypes.Params{}, err
	}
	return &resp.Params, nil
}

func (c *Client) MustQueryBTCCheckpointParams() *btcctypes.Params {
	var (
		params *btcctypes.Params
		err    error
	)

	err = retry.Do(c.retrySleepTime, c.maxRetrySleepTime, func() error {
		params, err = c.QueryBTCCheckpointParams()
		return err
	})
	if err != nil {
		panic(err)
	}
	return params
}
