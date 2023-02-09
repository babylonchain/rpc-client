package client

import (
	"fmt"

	"github.com/babylonchain/rpc-client/config"
	"github.com/babylonchain/rpc-client/query"
	lensclient "github.com/strangelove-ventures/lens/client"
)

var _ BabylonClient = &Client{}

type Client struct {
	*lensclient.ChainClient
	*query.QueryClient

	cfg *config.BabylonConfig
}

func New(cfg *config.BabylonConfig) (*Client, error) {
	// ensure cfg is valid
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// create a Tendermint/Cosmos client for Babylon
	cc, err := newLensClient(cfg.Unwrap())
	if err != nil {
		return nil, err
	}

	// create a queryClient so that the Client inherits all query functions
	queryClient := &query.QueryClient{
		RPCClient: cc.RPCClient,
		Cfg:       cfg,
	}

	// wrap to our type
	client := &Client{cc, queryClient, cfg}

	return client, nil
}

func (c *Client) GetConfig() *config.BabylonConfig {
	return c.cfg
}

func (c *Client) GetTagIdx() uint8 {
	tagIdxStr := c.cfg.TagIdx
	if len(tagIdxStr) != 1 {
		panic(fmt.Errorf("tag index should be one byte"))
	}
	// convert tagIdx from string to its ascii value
	return uint8(rune(tagIdxStr[0]))
}

func (c *Client) Stop() {
	if c.ChainClient.RPCClient != nil && c.ChainClient.RPCClient.IsRunning() {
		<-c.ChainClient.RPCClient.Quit()
	}
}
