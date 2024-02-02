package query

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/query"
	"strconv"
	"time"

	"github.com/babylonchain/rpc-client/config"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	cwclient "github.com/cosmos/relayer/v2/client"
	"github.com/strangelove-ventures/cometbft-client/client"
	"google.golang.org/grpc/metadata"
)

// QueryClient is a client that can only perform queries to a Babylon node
// It only requires `Cfg` to have `Timeout` and `RPCAddr`, but not other fields
// such as keyring, chain ID, etc..
type QueryClient struct {
	RPCClient cwclient.RPCClient
	timeout   time.Duration
}

// New creates a new QueryClient according to the given config
func New(cfg *config.BabylonQueryConfig) (*QueryClient, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	tmClient, err := client.NewClient(cfg.RPCAddr, cfg.Timeout)
	if err != nil {
		return nil, err
	}

	return &QueryClient{
		RPCClient: cwclient.NewRPCClient(tmClient),
		timeout:   cfg.Timeout,
	}, nil
}

// NewWithClient creates a new QueryClient with a given existing rpcClient and timeout
// used by `client/` where `ChainClient` already creates an rpc client
func NewWithClient(rpcClient cwclient.RPCClient, timeout time.Duration) (*QueryClient, error) {
	if timeout <= 0 {
		return nil, fmt.Errorf("timeout must be positive")
	}

	client := &QueryClient{
		RPCClient: rpcClient,
		timeout:   timeout,
	}

	return client, nil
}

// getQueryContext returns a context that includes the height and uses the timeout from the config
// (adapted from https://github.com/strangelove-ventures/lens/blob/v0.5.4/client/query/query_options.go#L29-L36)
func (c *QueryClient) getQueryContext() (context.Context, context.CancelFunc) {
	defaultOptions := DefaultQueryOptions()
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	strHeight := strconv.Itoa(int(defaultOptions.Height))
	ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, strHeight)
	return ctx, cancel
}

type QueryOptions struct {
	Pagination *query.PageRequest
	Height     int64
}

func DefaultQueryOptions() *QueryOptions {
	return &QueryOptions{
		Pagination: &query.PageRequest{
			Key:        []byte(""),
			Offset:     0,
			Limit:      1000,
			CountTotal: true,
		},
		Height: 0,
	}
}
