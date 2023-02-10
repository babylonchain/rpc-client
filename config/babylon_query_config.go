package config

import (
	"errors"
	"time"
)

// BabylonConfig defines configuration for the Babylon query client
type BabylonQueryConfig struct {
	RPCAddr string        `mapstructure:"rpc-addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func (cfg *BabylonQueryConfig) Validate() error {
	if cfg.Timeout <= 0 {
		return errors.New("cfg.Timeout must be positive")
	}
	return nil
}

func DefaultBabylonQueryConfig() BabylonQueryConfig {
	return BabylonQueryConfig{
		RPCAddr: "http://localhost:26657",
		Timeout: 20 * time.Second,
	}
}
