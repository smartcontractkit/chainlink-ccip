package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/environment/devenv"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/sdk/ccip"
	"go.uber.org/zap"
)

const nodeOverridesTomlFilePath = "ccip-v2-scripts-node-overrides.toml"

func NewEnvState(logger *zap.SugaredLogger, env config.DevspaceEnv) CCIPEnvState {
	return CCIPEnvState{
		logger:      logger,
		devspaceEnv: env,
	}
}

type CCIPEnvState struct {
	devspaceEnv config.DevspaceEnv
	logger      *zap.SugaredLogger
}

func (e *CCIPEnvState) getOutputFilePath(fileName string) string {
	return path.Join(e.devspaceEnv.TmpDir, fileName)
}

func (e *CCIPEnvState) SaveAddressBook(addresses map[uint64]map[string]deployment.TypeAndVersion) {
	e.SaveJSONOutputFile(ccip.AddressBookFileName, addresses)
}

func (e *CCIPEnvState) SaveChainConfigs(chainConfigs []devenv.ChainConfig) {
	configs := make([]ccip.ChainConfig, 0)
	for _, chainConfig := range chainConfigs {
		configs = append(configs, TransmittedConfig(chainConfig))
	}
	e.SaveJSONOutputFile(ccip.ChainsConfigsFileName, configs)
}

func TransmittedConfig(c devenv.ChainConfig) ccip.ChainConfig {
	return ccip.ChainConfig{
		ChainID:   c.ChainID,
		ChainName: c.ChainName,
		ChainType: c.ChainType,
		WSRPCs:    c.WSRPCs,
		HTTPRPCs:  c.HTTPRPCs,
	}
}

func (e *CCIPEnvState) AddressBookExists() bool {
	if _, err := os.Stat(e.getOutputFilePath(ccip.AddressBookFileName)); err == nil {
		return true
	}
	return false
}

func (e *CCIPEnvState) SaveNodesTomlOverride(capRegConfig deployment.CapabilityRegistryConfig, homeChainID uint64) {
	nodeTomlOverride := fmt.Sprintf(`
[Capabilities]
[Capabilities.ExternalRegistry]
Address = '%s'
NetworkID = '%s'
ChainID = '%d'`, capRegConfig.Contract.String(), "evm", homeChainID)

	// Create or open a file for writing
	file, err := os.Create(e.getOutputFilePath(nodeOverridesTomlFilePath))
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.WriteString(nodeTomlOverride)
	if err != nil {
		panic(err)
	}
}

func (e *CCIPEnvState) SaveNodeDetails(details ccip.NodesDetails) {
	e.SaveJSONOutputFile(ccip.NodesDetailsFileName, details)
}

func (e *CCIPEnvState) SaveJSONOutputFile(filename string, data interface{}) string {
	err := os.MkdirAll(e.devspaceEnv.TmpDir, os.ModeDir)
	if err != nil {
		e.logger.Error("unable to create temporary directory: %s", e.devspaceEnv.TmpDir)
		panic(err)
	}
	file, err := os.Create(e.getOutputFilePath(filename))
	if err != nil {
		e.logger.Errorf("unable to create JSON file: %s", e.getOutputFilePath(filename))
		return ""
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	return file.Name()
}
