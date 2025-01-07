package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
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
	e.SaveJSONOutputFile(crib.AddressBookFileName, addresses)
}

func (e *CCIPEnvState) SaveChainConfigs(chainConfigs []crib.ChainConfig) {
	e.SaveJSONOutputFile(crib.ChainsConfigsFileName, chainConfigs)
}

func (e *CCIPEnvState) AddressBookExists() bool {
	if _, err := os.Stat(e.getOutputFilePath(crib.AddressBookFileName)); err == nil {
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
		e.logger.Fatal("unable to create file", "err", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.WriteString(nodeTomlOverride)
	if err != nil {
		e.logger.Fatal("unable to write to file", "err", err)
		os.Exit(1)
	}
}

func (e *CCIPEnvState) SaveNodeDetails(details crib.NodesDetails) {
	e.SaveJSONOutputFile(crib.NodesDetailsFileName, details)
}

func (e *CCIPEnvState) SaveJSONOutputFile(filename string, data interface{}) string {
	err := os.MkdirAll(e.devspaceEnv.TmpDir, os.ModeDir)
	if err != nil {
		e.logger.Error("unable to create temporary directory: %s", e.devspaceEnv.TmpDir)
		os.Exit(1)
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
		e.logger.Fatal("unable to encode data", "err", err)
		os.Exit(1)
	}

	return file.Name()
}
