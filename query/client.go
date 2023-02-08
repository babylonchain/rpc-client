package query

import (
	"context"
	"strconv"

	btcctypes "github.com/babylonchain/babylon/x/btccheckpoint/types"
	"github.com/babylonchain/rpc-client/config"
	"github.com/cosmos/cosmos-sdk/client"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	lensquery "github.com/strangelove-ventures/lens/client/query"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	cfg *config.BabylonConfig
}

func New(cfg *config.BabylonConfig) (*Client, error) {
	// ensure cfg is valid
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{cfg}, nil
}

// getQueryContext returns a context that includes the height and uses the timeout from the config
// (adapted from https://github.com/strangelove-ventures/lens/blob/v0.5.4/client/query/query_options.go#L29-L36)
func (c *Client) getQueryContext() (context.Context, context.CancelFunc) {
	defaultOptions := lensquery.DefaultOptions()
	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.Timeout)
	strHeight := strconv.Itoa(int(defaultOptions.Height))
	ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, strHeight)
	return ctx, cancel
}

func (c *Client) query(f func(ctx context.Context, clientCtx client.Context)) {
	ctx, cancel := c.getQueryContext()
	defer cancel()

	// pass the clientCtx, which implements grpc.ClientConn interface, to f()
	// so that f() uses NewQueryClient to generate a querier to a certain module
	clientCtx := client.Context{NodeURI: c.cfg.GRPCAddr}
	f(ctx, clientCtx)
}

// BTCCheckpointParams queries btccheckpoint module's parameters via ChainClient
func (c *Client) BTCCheckpointParams() (*btcctypes.Params, error) {
	var (
		resp *btcctypes.QueryParamsResponse
		err  error
	)
	c.query(func(ctx context.Context, clientCtx client.Context) {
		queryClient := btcctypes.NewQueryClient(clientCtx)
		req := &btcctypes.QueryParamsRequest{}
		resp, err = queryClient.Params(ctx, req)
	})

	if err != nil {
		return nil, err
	}
	return &resp.Params, nil
}
