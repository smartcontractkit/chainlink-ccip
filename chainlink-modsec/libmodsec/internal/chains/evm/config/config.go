package config

import (
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/config"
	evmcfg "github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
)

func CreateNewEVMChainFromTOML(lggr logger.Logger, configTOML string) (*config.ChainScoped, error) {
	var cfg struct {
		EVM evmcfg.EVMConfig
	}

	d := toml.NewDecoder(strings.NewReader(configTOML))
	d.DisallowUnknownFields()

	if err := d.Decode(&cfg); err != nil {
		lggr.Panicf("failed to decode config toml: %w:\n\t%s", err, configTOML)
	}

	evmConfig := config.NewTOMLChainScopedConfig(&cfg.EVM)
	return evmConfig, nil
}
