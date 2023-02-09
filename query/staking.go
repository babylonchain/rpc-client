package query

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// QueryStaking queries the Staking module of the Babylon node
// according to the given function
func (c *QueryClient) QueryStaking(f func(ctx context.Context, queryClient stakingtypes.QueryClient)) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.RPCClient}
	queryClient := stakingtypes.NewQueryClient(clientCtx)

	f(ctx, queryClient)
}

// QueryStakingParams queries btccheckpoint module's parameters via ChainClient
func (c *QueryClient) StakingParams() (*stakingtypes.Params, error) {
	var (
		resp *stakingtypes.QueryParamsResponse
		err  error
	)
	c.QueryStaking(func(ctx context.Context, queryClient stakingtypes.QueryClient) {
		req := &stakingtypes.QueryParamsRequest{}
		resp, err = queryClient.Params(ctx, req)
	})

	if err != nil {
		return nil, err
	}
	return &resp.Params, nil
}
