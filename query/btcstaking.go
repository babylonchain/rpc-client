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

// BTCStakingParams queries the BTC staking module parameters
func (c *QueryClient) BTCStakingParams() (*btcstakingtypes.QueryParamsResponse, error) {
	var resp *btcstakingtypes.QueryParamsResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryParamsRequest{}
		resp, err = queryClient.Params(ctx, req)
		return err
	})

	return resp, err
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
func (c *QueryClient) BTCValidatorDelegations(valBtcPkHex string, pagination *sdkquerytypes.PageRequest) (*btcstakingtypes.QueryBTCValidatorDelegationsResponse, error) {
	var resp *btcstakingtypes.QueryBTCValidatorDelegationsResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryBTCValidatorDelegationsRequest{
			ValBtcPkHex: valBtcPkHex,
			Pagination:  pagination,
		}
		resp, err = queryClient.BTCValidatorDelegations(ctx, req)
		return err
	})

	return resp, err
}

// BTCDelegations queries the BTCStaking module for all delegations under a given status
func (c *QueryClient) BTCDelegations(status btcstakingtypes.BTCDelegationStatus, pagination *sdkquerytypes.PageRequest) (*btcstakingtypes.QueryBTCDelegationsResponse, error) {
	var resp *btcstakingtypes.QueryBTCDelegationsResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryBTCDelegationsRequest{
			Status:     status,
			Pagination: pagination,
		}
		resp, err = queryClient.BTCDelegations(ctx, req)
		return err
	})

	return resp, err
}

// BTCDelegation queries the BTCStaking module to retrieve delegation by corresponding staking tx hash
func (c *QueryClient) BTCDelegation(stakingTxHashHex string) (*btcstakingtypes.QueryBTCDelegationResponse, error) {
	var resp *btcstakingtypes.QueryBTCDelegationResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryBTCDelegationRequest{
			StakingTxHashHex: stakingTxHashHex,
		}
		resp, err = queryClient.BTCDelegation(ctx, req)
		return err
	})

	return resp, err
}

// ActiveBTCValidatorsAtHeight queries the BTCStaking module for all btc validators
// with non-zero voting power at a given height
func (c *QueryClient) ActiveBTCValidatorsAtHeight(height uint64, pagination *sdkquerytypes.PageRequest) (*btcstakingtypes.QueryActiveBTCValidatorsAtHeightResponse, error) {
	var resp *btcstakingtypes.QueryActiveBTCValidatorsAtHeightResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryActiveBTCValidatorsAtHeightRequest{
			Height:     height,
			Pagination: pagination,
		}
		resp, err = queryClient.ActiveBTCValidatorsAtHeight(ctx, req)
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

func (c *QueryClient) ActivatedHeight() (*btcstakingtypes.QueryActivatedHeightResponse, error) {
	var resp *btcstakingtypes.QueryActivatedHeightResponse
	err := c.QueryBTCStaking(func(ctx context.Context, queryClient btcstakingtypes.QueryClient) error {
		var err error
		req := &btcstakingtypes.QueryActivatedHeightRequest{}
		resp, err = queryClient.ActivatedHeight(ctx, req)
		return err
	})

	return resp, err
}
