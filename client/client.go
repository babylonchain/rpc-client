package client

import (
	"fmt"
	"time"

	"github.com/babylonchain/rpc-client/config"
	"github.com/sirupsen/logrus"
	lensclient "github.com/strangelove-ventures/lens/client"
)

var _ BabylonClient = &Client{}

type Client struct {
	*lensclient.ChainClient
	cfg *config.BabylonConfig

	log *logrus.Logger

	// retry attributes
	retrySleepTime    time.Duration
	maxRetrySleepTime time.Duration
}

func New(cfg *config.BabylonConfig, log *logrus.Logger, retrySleepTime, maxRetrySleepTime time.Duration) (*Client, error) {
	// create a Tendermint/Cosmos client for Babylon
	cc, err := newLensClient(cfg.Unwrap())
	if err != nil {
		return nil, err
	}

	// wrap to our type
	client := &Client{cc, cfg, log, retrySleepTime, maxRetrySleepTime}

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
	if c.RPCClient != nil && c.RPCClient.IsRunning() {
		<-c.RPCClient.Quit()
	}
}
