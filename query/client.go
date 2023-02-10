package query

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	lensquery "github.com/strangelove-ventures/lens/client/query"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"google.golang.org/grpc/metadata"
)

var _ BabylonQueryClient = &QueryClient{}

// QueryClient is a client that can only perform queries to a Babylon node
// It only requires `Cfg` to have `Timeout` and `RPCAddr`, but not other fields
// such as keyring, chain ID, etc..
type QueryClient struct {
	RPCClient rpcclient.Client
	timeout   time.Duration
}

// NewWithClient creates a new QueryClient to a given rpcAddr and a given timeout
func New(rpcAddr string, timeout time.Duration) (*QueryClient, error) {
	if timeout <= 0 {
		return nil, fmt.Errorf("timeout must be positive")
	}

	tmClient, err := client.NewClientFromNode(rpcAddr)
	if err != nil {
		return nil, err
	}

	client := &QueryClient{
		RPCClient: tmClient,
		timeout:   timeout,
	}

	return client, nil
}

// NewWithClient creates a new QueryClient with a given existing rpcClient and timeout
func NewWithClient(rpcClient rpcclient.Client, timeout time.Duration) (*QueryClient, error) {
	if timeout <= 0 {
		return nil, fmt.Errorf("timeout must be positive")
	}

	client := &QueryClient{
		RPCClient: rpcClient,
		timeout:   timeout,
	}

	return client, nil
}

func (c *QueryClient) Stop() {
	if c.RPCClient != nil && c.RPCClient.IsRunning() {
		<-c.RPCClient.Quit()
	}
}

// getQueryContext returns a context that includes the height and uses the timeout from the config
// (adapted from https://github.com/strangelove-ventures/lens/blob/v0.5.4/client/query/query_options.go#L29-L36)
func (c *QueryClient) getQueryContext() (context.Context, context.CancelFunc) {
	defaultOptions := lensquery.DefaultOptions()
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	strHeight := strconv.Itoa(int(defaultOptions.Height))
	ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, strHeight)
	return ctx, cancel
}
