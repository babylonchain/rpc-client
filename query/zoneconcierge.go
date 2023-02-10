package query

import (
	cztypes "github.com/babylonchain/babylon/x/zoneconcierge/types"
	"github.com/babylonchain/rpc-client/client"
	sdkquerytypes "github.com/cosmos/cosmos-sdk/types/query"
)

// FinalizedConnectedChainInfo queries the zoneconcierge module to get the finalization information for a connected chain
func FinalizedConnectedChainInfo(c *client.Client, chainID string) (*cztypes.QueryFinalizedChainInfoResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := cztypes.NewQueryClient(c.ChainClient)
	req := &cztypes.QueryFinalizedChainInfoRequest{
		ChainId: chainID,
	}

	resp, err := queryClient.FinalizedChainInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ConnectedChainInfo queries the zoneconcierge module to get information for a connected chain
func ConnectedChainInfo(c *client.Client, chainID string) (*cztypes.ChainInfo, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := cztypes.NewQueryClient(c.ChainClient)
	req := &cztypes.QueryChainInfoRequest{
		ChainId: chainID,
	}

	resp, err := queryClient.ChainInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.ChainInfo, nil
}

// ConnectedChainList queries the zoneconierge module for the chain IDs of the connected chains
func ConnectedChainList(c *client.Client) (*cztypes.QueryChainListResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := cztypes.NewQueryClient(c.ChainClient)
	req := &cztypes.QueryChainListRequest{}

	resp, err := queryClient.ChainList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ConnectedChainHeaders queries the zoneconcierge module for the headers of a connected chain
func ConnectedChainHeaders(c *client.Client, chainID string, pagination *sdkquerytypes.PageRequest) (*cztypes.QueryListHeadersResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := cztypes.NewQueryClient(c.ChainClient)
	req := &cztypes.QueryListHeadersRequest{
		ChainId:    chainID,
		Pagination: pagination,
	}

	resp, err := queryClient.ListHeaders(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ConnectedChainEpochInfo queries the zoneconcierge module for the chain information of a connected chain at a particular epoch
func ConnectedChainEpochInfo(c *client.Client, chainID string, epochNum uint64) (*cztypes.QueryEpochChainInfoResponse, error) {
	ctx, cancel := c.GetDefaultQueryContext()
	defer cancel()

	queryClient := cztypes.NewQueryClient(c.ChainClient)
	req := &cztypes.QueryEpochChainInfoRequest{
		ChainId:  chainID,
		EpochNum: epochNum,
	}

	resp, err := queryClient.EpochChainInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
