package client

import (
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/strangelove-ventures/lens/client/query"
)

// QueryStakingParams queries staking module's parameters via ChainClient
// code is adapted from https://github.com/strangelove-ventures/lens/blob/v0.5.1/cmd/staking.go#L128-L149
func (c *Client) QueryStakingParams() (*stakingtypes.Params, error) {
	query := query.Query{Client: c.ChainClient, Options: query.DefaultOptions()} // TODO: what's the impact of DefaultOptions()?
	resp, err := query.Staking_Params()
	if err != nil {
		return &stakingtypes.Params{}, err
	}

	return &resp.Params, nil
}
