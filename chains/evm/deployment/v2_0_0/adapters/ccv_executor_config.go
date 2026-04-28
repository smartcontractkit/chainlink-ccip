package adapters

import (
	"fmt"

	chainsel "github.com/smartcontractkit/chain-selectors"
	rmnremote "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	execop "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	offrampop "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	dsutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
	"github.com/smartcontractkit/chainlink-ccv/executor"
	"github.com/smartcontractkit/chainlink-ccv/pkg/chainaccess"
)

type EVMCCVExecutorConfigAdapter struct{}

var _ ccvdeploymentadapters.ExecutorConfigAdapter = (*EVMCCVExecutorConfigAdapter)(nil)

func (a *EVMCCVExecutorConfigAdapter) GetDeployedChains(ds datastore.DataStore, qualifier string) []uint64 {
	if ds == nil {
		return nil
	}
	refs := ds.Addresses().Filter(
		datastore.AddressRefByQualifier(qualifier),
		datastore.AddressRefByType(datastore.ContractType(sequences.ExecutorProxyType)),
		datastore.AddressRefByVersion(execop.Version),
	)
	seen := make(map[uint64]struct{}, len(refs))
	chains := make([]uint64, 0, len(refs))
	for _, ref := range refs {
		if _, exists := seen[ref.ChainSelector]; exists {
			continue
		}
		family, err := chainsel.GetSelectorFamily(ref.ChainSelector)
		if err != nil || family != chainsel.FamilyEVM {
			continue
		}
		seen[ref.ChainSelector] = struct{}{}
		chains = append(chains, ref.ChainSelector)
	}
	return chains
}

func (a *EVMCCVExecutorConfigAdapter) BuildChainConfig(ds datastore.DataStore, chainSelector uint64, qualifier string) (executor.ChainConfiguration, error) {
	toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

	offRampAddr, err := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(offrampop.ContractType),
		Version: offrampop.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return executor.ChainConfiguration{}, fmt.Errorf("failed to get off ramp address for chain %d: %w", chainSelector, err)
	}

	rmnRemoteAddr, err := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(rmnremote.ContractType),
		Version: rmnremote.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return executor.ChainConfiguration{}, fmt.Errorf("failed to get rmn remote address for chain %d: %w", chainSelector, err)
	}

	executorAddr, err := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(sequences.ExecutorProxyType),
		Qualifier: qualifier,
		Version:   execop.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return executor.ChainConfiguration{}, fmt.Errorf("failed to get executor proxy address for chain %d: %w", chainSelector, err)
	}

	return executor.ChainConfiguration{
		DestinationChainConfig: chainaccess.DestinationChainConfig{
			OffRampAddress: offRampAddr,
			RmnAddress:     rmnRemoteAddr,
		},
		DefaultExecutorAddress: executorAddr,
	}, nil
}
