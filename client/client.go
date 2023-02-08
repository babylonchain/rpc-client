package client

import (
	"context"
	"fmt"
	"github.com/babylonchain/rpc-client/config"
	lensclient "github.com/strangelove-ventures/lens/client"
	"github.com/strangelove-ventures/lens/client/query"
)

var _ BabylonClient = &Client{}

type Client struct {
	*lensclient.ChainClient
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

	// wrap to our type
	client := &Client{cc, cfg}

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

func (c *Client) GetDefaultQuery() *query.Query {
	return &query.Query{Client: c.ChainClient, Options: query.DefaultOptions()}
}

func (c *Client) GetDefaultQueryContext() (context.Context, context.CancelFunc) {
	query := c.GetDefaultQuery()
	ctx, cancel := query.GetQueryContext()
	return ctx, cancel
}

func (c *Client) Stop() {
	if c.ChainClient.RPCClient != nil && c.ChainClient.RPCClient.IsRunning() {
		<-c.ChainClient.RPCClient.Quit()
	}
}
