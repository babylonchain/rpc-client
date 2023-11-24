package client

import (
	"context"
	bbn "github.com/babylonchain/babylon/app"
	"github.com/cosmos/relayer/v2/relayer/chains/cosmos"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/babylonchain/rpc-client/config"
	"github.com/babylonchain/rpc-client/query"
)

type Client struct {
	*query.QueryClient

	provider *cosmos.CosmosProvider
	timeout  time.Duration
	logger   *logrus.Logger
	cfg      *config.BabylonConfig
}

func New(cfg *config.BabylonConfig, logger *logrus.Logger) (*Client, error) {
	// ensure cfg is valid
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	zapLogger, err := newRootLogger("console", true)
	if err != nil {
		return nil, err
	}

	// Create tmp Babylon app to retrieve and register codecs
	tmpBabylon := bbn.NewTmpBabylonApp()

	cosmosConfig := cfg.ToCosmosProviderConfig()
	provider, err := cosmosConfig.NewProvider(
		zapLogger,
		"", // TODO: set home path
		true,
		"babylon",
	)
	if err != nil {
		return nil, err
	}

	cp := provider.(*cosmos.CosmosProvider)
	cp.PCfg.KeyDirectory = cfg.KeyDirectory
	// Need to override this manually as otherwise option from config is ignored
	cp.Cdc = cosmos.Codec{
		InterfaceRegistry: tmpBabylon.InterfaceRegistry(),
		Marshaler:         tmpBabylon.AppCodec(),
		TxConfig:          tmpBabylon.TxConfig(),
		Amino:             tmpBabylon.LegacyAmino(),
	}

	err = cp.Init(context.Background())
	if err != nil {
		return nil, err
	}

	// create a queryClient so that the Client inherits all query functions
	queryClient, err := query.NewWithClient(cp.RPCClient, cfg.Timeout)
	if err != nil {
		return nil, err
	}

	return &Client{
		queryClient,
		cp,
		cfg.Timeout,
		logger,
		cfg,
	}, nil
}

func (c *Client) GetConfig() *config.BabylonConfig {
	return c.cfg
}

func (c *Client) Stop() error {
	if !c.provider.RPCClient.IsRunning() {
		return nil
	}

	return c.provider.RPCClient.Stop()
}
