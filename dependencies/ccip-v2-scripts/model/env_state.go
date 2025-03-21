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

const (
	WorkerNodeInitialConfigOverridesFileNamePattern = "config-override-%d.toml"
	BootNodeInitialConfigOverridesFileNamePattern   = "config-override-bt-%d.toml"
	NodeInitialSecretsOverridesFileNamePattern      = "secrets-override-%d.toml"
	BootNodeInitialSecretsOverridesFileNamePattern  = "secrets-override-bt-%d.toml"
	nodeCapabilityRegistryOverrideFileName          = "node-capability-registry-overrides.toml"
	rmnSharedConfigTomlFileName                     = "rmn-shared-config.toml"
	rmnLocalConfigTomlFileName                      = "rmn-local-config.toml"
)

const (
	DataSubdir         = "data"
	DONOverridesSubdir = "ccip"
)

type EnvStateFileType int

const (
	DataFile EnvStateFileType = iota
	CLNodeConfigInputs
	RMNConfigInput
)

func SubDirForFileType(fileType EnvStateFileType) string {
	switch fileType {
	case DataFile:
		return "data"
	case CLNodeConfigInputs:
		return "ccip"
	case RMNConfigInput:
		return "rmn"
	default:
		return ""
	}
}

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

func (e *CCIPEnvState) getOutputFilePath(fileName string, fileType EnvStateFileType) string {
	return path.Join(e.devspaceEnv.TmpDir, SubDirForFileType(fileType), fileName)
}

func (e *CCIPEnvState) SaveAddressBook(addresses map[uint64]map[string]deployment.TypeAndVersion) {
	e.SaveJSONOutputFile(crib.AddressBookFileName, addresses, DataFile)
}

func (e *CCIPEnvState) SaveChainConfigs(chainConfigs []crib.ChainConfig) {
	e.SaveJSONOutputFile(crib.ChainsConfigsFileName, chainConfigs, DataFile)
}

func (e *CCIPEnvState) SaveRMNNodeConfigs(configs []crib.RMNNodeConfig) {
	e.SaveJSONOutputFile(crib.RMNNodeIdentitiesFileName, configs, DataFile)
}

func (e *CCIPEnvState) AddressBookExists() bool {
	if _, err := os.Stat(e.getOutputFilePath(crib.AddressBookFileName, DataFile)); err == nil {
		return true
	}
	return false
}

func (e *CCIPEnvState) NodeOverridesExist() bool {
	if _, err := os.Stat(e.getOutputFilePath(fmt.Sprintf(NodeInitialSecretsOverridesFileNamePattern, 0), CLNodeConfigInputs)); err == nil {
		return true
	}
	return false
}

func (e *CCIPEnvState) RMNIdentitiesExists() bool {
	if _, err := os.Stat(e.getOutputFilePath(crib.RMNNodeIdentitiesFileName, DataFile)); err == nil {
		return true
	}
	return false
}

func (e *CCIPEnvState) RMNTomlConfigsExists() bool {
	if _, err := os.Stat(e.getOutputFilePath(rmnSharedConfigTomlFileName, DataFile)); err != nil {
		return false
	}
	if _, err := os.Stat(e.getOutputFilePath(rmnLocalConfigTomlFileName, DataFile)); err != nil {
		return false
	}
	return true
}

func (e *CCIPEnvState) SaveCapRegistryNodeOverride(capRegConfig deployment.CapabilityRegistryConfig, homeChainID uint64) {
	nodeTomlOverride := fmt.Sprintf(`
[Capabilities]
[Capabilities.ExternalRegistry]
Address = '%s'
NetworkID = '%s'
ChainID = '%d'`, capRegConfig.Contract.String(), "evm", homeChainID)

	e.SaveStringToFile(nodeTomlOverride, nodeCapabilityRegistryOverrideFileName, CLNodeConfigInputs)
}

func (e *CCIPEnvState) SaveStringToFile(content string, fileName string, fileType EnvStateFileType) {
	e.ensureSubDirExists(fileType)

	file, err := os.Create(e.getOutputFilePath(fileName, fileType))
	if err != nil {
		e.logger.Fatal("unable to create file", "err", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		e.logger.Fatal("unable to write to file", "err", err)
		os.Exit(1)
	}
}

func (e *CCIPEnvState) SaveNodeDetails(details crib.NodesDetails) {
	e.SaveJSONOutputFile(crib.NodesDetailsFileName, details, DataFile)
}

func (e *CCIPEnvState) SaveRMNSharedToml(tomlData []byte) {
	e.SaveStringToFile(string(tomlData), rmnSharedConfigTomlFileName, RMNConfigInput)
}

func (e *CCIPEnvState) SaveRMNLocalToml(tomlData []byte) {
	e.SaveStringToFile(string(tomlData), rmnLocalConfigTomlFileName, RMNConfigInput)
}

func (e *CCIPEnvState) SaveJSONOutputFile(filename string, data interface{}, fileType EnvStateFileType) string {
	e.ensureSubDirExists(fileType)
	file, err := os.Create(e.getOutputFilePath(filename, DataFile))
	if err != nil {
		e.logger.Errorf("unable to create JSON file: %s", e.getOutputFilePath(filename, DataFile))
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

func (e *CCIPEnvState) ensureSubDirExists(fileType EnvStateFileType) {
	dirPath := path.Join(e.devspaceEnv.TmpDir, SubDirForFileType(fileType))
	err := os.MkdirAll(dirPath, 0o755)
	if err != nil {
		e.logger.Error("unable to create temporary directory: %s", e.devspaceEnv.TmpDir)
		os.Exit(1)
	}
}
