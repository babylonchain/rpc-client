package query

import (
	"context"

	btcstakingtypes "github.com/babylonchain/babylon/x/btcstaking/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// QueryBTCStaking queries the BTCStaking module of the Babylon node according to the given function
func (c *QueryClient) QueryBTCStaking(f func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error) error {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	clientCtx := client.Context{Client: c.RPCClient}
	queryClient := btcstakingtypes.NewQueryClient(clientCtx)

	return f(ctx, queryClient)
}

// BTCValidators queries the BTCStaking module for all btc validators
func (c *QueryClient) BTCValidators(pagination *sdkquerytypes.PageRequest) (*btcstakingtypes.QueryBTCValidatorsResponse, error) {
	var resp *btcstakingtypes.QueryBTCValidatorsResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryBTCValidatorsRequest{
			Pagination: pagination,
		}
		resp, err = queryClient.BTCValidators(ctx, req)
		return err
	})

	return resp, err
}

// BTCValidatorDelegations queries the BTCStaking module for all delegations of a btc validator
func (c *QueryClient) BTCValidatorDelegations(valBtcPkHex string, noJurySigOnly bool, pagination *sdkquerytypes.PageRequest) (*btcstakingtypes.QueryBTCValidatorDelegationsResponse, error) {
	var resp *btcstakingtypes.QueryBTCValidatorDelegationsResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryBTCValidatorDelegationsRequest{
			ValBtcPkHex:   valBtcPkHex,
			NoJurySigOnly: noJurySigOnly,
			Pagination:    pagination,
		}
		resp, err = queryClient.BTCValidatorDelegations(ctx, req)
		return err
	})

	return resp, err
}

// BTCValidatorsAtHeight queries the BTCStaking module for all btc validators at a given height
func (c *QueryClient) BTCValidatorsAtHeight(height uint64, pagination *sdkquerytypes.PageRequest) (*btcstakingtypes.QueryBTCValidatorsAtHeightResponse, error) {
	var resp *btcstakingtypes.QueryBTCValidatorsAtHeightResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryBTCValidatorsAtHeightRequest{
			Height:     height,
			Pagination: pagination,
		}
		resp, err = queryClient.BTCValidatorsAtHeight(ctx, req)
		return err
	})

	return resp, err
}

// BTCValidatorPowerAtHeight queries the BTCStaking module for the power of a btc validator at a given height
func (c *QueryClient) BTCValidatorPowerAtHeight(valBtcPkHex string, height uint64) (*btcstakingtypes.QueryBTCValidatorPowerAtHeightResponse, error) {
	var resp *btcstakingtypes.QueryBTCValidatorPowerAtHeightResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryBTCValidatorPowerAtHeightRequest{
			ValBtcPkHex: valBtcPkHex,
			Height:      height,
		}
		resp, err = queryClient.BTCValidatorPowerAtHeight(ctx, req)
		return err
	})

	return resp, err
}
