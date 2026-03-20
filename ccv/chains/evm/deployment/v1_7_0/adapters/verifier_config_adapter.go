package adapters

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	dsutil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type EVMVerifierJobConfigAdapter struct{}

var _ ccvadapters.VerifierConfigAdapter = (*EVMVerifierJobConfigAdapter)(nil)

func (a *EVMVerifierJobConfigAdapter) ResolveVerifierContractAddresses(
	ds datastore.DataStore,
	chainSelector uint64,
	committeeQualifier string,
	executorQualifier string,
) (*ccvadapters.VerifierContractAddresses, error) {
	toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

	committeeVerifierAddr, err := dsutil.FindAndFormatFirstRef(ds, chainSelector, toAddress,
		datastore.AddressRef{
			Type:      datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
			Qualifier: committeeQualifier,
		},
		datastore.AddressRef{
			Type:      datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierContractType),
			Qualifier: committeeQualifier,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get committee verifier address for chain %d: %w", chainSelector, err)
	}

	onRampAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(onramp.ContractType),
		Version: onramp.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get on ramp address for chain %d: %w", chainSelector, err)
	}

	executorAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(sequences.ExecutorProxyType),
		Qualifier: executorQualifier,
		Version:   executor.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get executor proxy address for chain %d: %w", chainSelector, err)
	}

	rmnRemoteAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(rmn_remote.ContractType),
		Version: rmn_remote.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get rmn remote address for chain %d: %w", chainSelector, err)
	}

	return &ccvadapters.VerifierContractAddresses{
		CommitteeVerifierAddress: committeeVerifierAddr,
		OnRampAddress:            onRampAddr,
		ExecutorProxyAddress:     executorAddr,
		RMNRemoteAddress:         rmnRemoteAddr,
	}, nil
}
