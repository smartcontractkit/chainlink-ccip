// Package configuration provides configuration management for the aggregator service.
package configuration

import (
	"github.com/BurntSushi/toml"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
)

// LoadConfig loads the aggregator configuration from a file.
func LoadConfig(filePath string) (*model.AggregatorConfig, error) {
	var config model.AggregatorConfig
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
