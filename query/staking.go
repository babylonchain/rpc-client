package query

import (
	"github.com/babylonchain/rpc-client/client"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// QueryStakingParams queries staking module's parameters via ChainClient
// code is adapted from https://github.com/strangelove-ventures/lens/blob/v0.5.1/cmd/staking.go#L128-L149
func QueryStakingParams(c *client.Client) (*stakingtypes.QueryParamsResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := stakingtypes.NewQueryClient(c.ChainClient)
	req := &stakingtypes.QueryParamsRequest{}
	resp, err := queryClient.Params(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
