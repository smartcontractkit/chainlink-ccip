package verification

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	cldverification "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/verification"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/verification/evm"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

var _ evm.ContractInputsProvider = (*EVMContractInputsProvider)(nil)
var _ hooks.ContractVerification = (*EVMContractInputsProvider)(nil)

func init() {
	hooks.GetContractVerificationRegistry().Register(chain_selectors.FamilyEVM, &EVMContractInputsProvider{})
}

// EVMContractInputsProvider implements hooks.ContractVerification for EVM chains:
// it supplies Solidity metadata and builds per-address verifiers using the framework's EVM strategy.
type EVMContractInputsProvider struct{}

func (p *EVMContractInputsProvider) NeedsVerification(ref datastore.AddressRef) bool {
	contractType := cldf.ContractType(ref.Type)
	version := ref.Version
	if versionedMetadata, versionedMetadataExists := contractMetadata[contractType]; versionedMetadataExists {
		if version == nil {
			return false
		}
		if _, versionExists := versionedMetadata[version.String()]; versionExists {
			return true
		}
	}

	return false
}

func (p *EVMContractInputsProvider) FilterNetworks(envName string, dom domain.Domain, lggr logger.Logger) (*cfgnet.Config, error) {
	networkCfg, err := config.LoadNetworks(envName, dom, lggr)
	if err != nil {
		return nil, err
	}

	return networkCfg.FilterWith(cfgnet.ChainFamilyFilter(chain_selectors.FamilyEVM)), nil
}

func (p *EVMContractInputsProvider) ForEachNetwork(
	ctx context.Context,
	network cfgnet.Network,
	selector uint64,
	lggr logger.Logger,
	logPrefix string) (
	hooks.VerifierBuilderForNetwork,
	bool,
) {
	strategy := evm.GetVerificationStrategyForNetwork(network)
	if strategy == evm.StrategyUnknown {
		lggr.Warnf("%s: no verification strategy for %d, skipping network", logPrefix, selector)
		return nil, true
	}
	lggr.Infof("%s: using verification strategy %d for %d", logPrefix, strategy, selector)
	return func(_ context.Context, ref datastore.AddressRef) (cldverification.Verifiable, error) {
		metadata, err := p.GetInputs(ref.Type, ref.Version)
		if err != nil {
			return nil, err
		}
		ch, ok := chain_selectors.ChainBySelector(selector)
		if !ok {
			return nil, fmt.Errorf("chain selector %d not found in chain selectors", selector)
		}
		v, err := evm.NewVerifier(strategy, evm.VerifierConfig{
			Chain:        ch,
			Network:      network,
			Address:      ref.Address,
			Metadata:     metadata,
			ContractType: string(ref.Type),
			Version:      ref.Version.String(),
			PollInterval: hooks.DefaultVerifyPollInterval,
			Logger:       lggr,
		})
		if err != nil {
			return nil, fmt.Errorf("%s %s (%s on %s): %w",
				ref.Type, ref.Version, ref.Address, ch.Name, err)
		}

		return v, nil
	}, false
}

// GetInputs returns metadata for the given contract type and version.
func (e *EVMContractInputsProvider) GetInputs(contractType datastore.ContractType, version *semver.Version) (evm.SolidityContractMetadata, error) {
	if version == nil {
		return evm.SolidityContractMetadata{}, nil
	}
	meta, err := LoadSolidityContractMetadata(cldf.ContractType(contractType), version)
	if err != nil {
		return evm.SolidityContractMetadata{}, err
	}

	return evm.SolidityContractMetadata{
		Version:  meta.Version,
		Language: meta.Language,
		Settings: meta.Settings,
		Sources:  meta.Sources,
		Bytecode: meta.Bytecode,
		Name:     meta.Name,
	}, nil
}

// RawContractInfo holds explorer verification inputs for a single contract version.
type rawContractInfo struct {
	solidityStandardJSONInput string
	bytecode                  string
	name                      string
}

// LoadSolidityContractMetadata loads the metadata for a contract type and version, including the standard JSON input, bytecode, and name.
func LoadSolidityContractMetadata(
	contractType cldf.ContractType,
	version *semver.Version,
) (evm.SolidityContractMetadata, error) {
	contract, ok := contractMetadata[contractType]
	if !ok {
		return evm.SolidityContractMetadata{}, fmt.Errorf("no contract found for type %s", contractType)
	}
	contractWithVersion, ok := contract[version.String()]
	if !ok {
		return evm.SolidityContractMetadata{}, fmt.Errorf("no contract found for type %s with version %s", contractType, version)
	}

	var input evm.SolidityContractMetadata
	err := json.Unmarshal([]byte(contractWithVersion.solidityStandardJSONInput), &input)
	if err != nil {
		return evm.SolidityContractMetadata{}, fmt.Errorf("failed to unmarshal solidity standard JSON input for contract type %s: %w", contractType, err)
	}
	// Add remaining fields that don't exist in the standard JSON input
	input.Bytecode = contractWithVersion.bytecode
	input.Name = contractWithVersion.name

	return input, nil
}
