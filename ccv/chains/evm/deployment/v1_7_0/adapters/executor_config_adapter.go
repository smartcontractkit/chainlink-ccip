package adapters

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	execcontract "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	offrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	dsutil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type EVMExecutorConfigAdapter struct{}

var _ adapters.ExecutorConfigAdapter = (*EVMExecutorConfigAdapter)(nil)

func (a *EVMExecutorConfigAdapter) GetDeployedChains(ds datastore.DataStore, qualifier string) []uint64 {
	if ds == nil {
		return nil
	}
	refs := ds.Addresses().Filter(
		datastore.AddressRefByQualifier(qualifier),
		datastore.AddressRefByType(datastore.ContractType(sequences.ExecutorProxyType)),
	)
	seen := make(map[uint64]struct{}, len(refs))
	chains := make([]uint64, 0, len(refs))
	for _, ref := range refs {
		if _, exists := seen[ref.ChainSelector]; !exists {
			seen[ref.ChainSelector] = struct{}{}
			chains = append(chains, ref.ChainSelector)
		}
	}
	return chains
}

func (a *EVMExecutorConfigAdapter) BuildChainConfig(ds datastore.DataStore, chainSelector uint64, qualifier string) (adapters.ExecutorChainConfig, error) {
	toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

	offRampAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(offrampoperations.ContractType),
		Version: offrampoperations.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return adapters.ExecutorChainConfig{}, fmt.Errorf("failed to get off ramp address for chain %d: %w", chainSelector, err)
	}

	rmnRemoteAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(rmn_remote.ContractType),
		Version: rmn_remote.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return adapters.ExecutorChainConfig{}, fmt.Errorf("failed to get rmn remote address for chain %d: %w", chainSelector, err)
	}

	executorAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(sequences.ExecutorProxyType),
		Qualifier: qualifier,
		Version:   execcontract.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return adapters.ExecutorChainConfig{}, fmt.Errorf("failed to get executor proxy address for chain %d: %w", chainSelector, err)
	}

	return adapters.ExecutorChainConfig{
		OffRampAddress:       offRampAddr,
		RmnAddress:           rmnRemoteAddr,
		ExecutorProxyAddress: executorAddr,
	}, nil
}
