package query

import (
	"context"
	"strconv"

	"github.com/babylonchain/rpc-client/config"
	"github.com/cosmos/cosmos-sdk/client"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	lensquery "github.com/strangelove-ventures/lens/client/query"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"google.golang.org/grpc/metadata"
)

type Client struct {
	rpcclient.Client
	cfg *config.BabylonConfig
}

func New(cfg *config.BabylonConfig) (*Client, error) {
	// ensure cfg is valid
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	tmClient, err := client.NewClientFromNode(cfg.RPCAddr)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Client: tmClient,
		cfg:    cfg,
	}

	return client, nil
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
