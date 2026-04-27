package adapters

import (
	"fmt"

	rmnremote "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	cctpverifier "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_verifier"
	onrampop "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	dsutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

type EVMCCVTokenVerifierConfigAdapter struct{}

var _ ccvdeploymentadapters.TokenVerifierConfigAdapter = (*EVMCCVTokenVerifierConfigAdapter)(nil)

func (a *EVMCCVTokenVerifierConfigAdapter) ResolveTokenVerifierAddresses(
	ds datastore.DataStore,
	chainSelector uint64,
	cctpQualifier string,
	lombardQualifier string,
) (*ccvdeploymentadapters.TokenVerifierChainAddresses, error) {
	toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

	onRampAddr, err := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type: datastore.ContractType(onrampop.ContractType),
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get on ramp address for chain %d: %w", chainSelector, err)
	}

	rmnRemoteAddr, err := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type: datastore.ContractType(rmnremote.ContractType),
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get rmn remote address for chain %d: %w", chainSelector, err)
	}

	result := &ccvdeploymentadapters.TokenVerifierChainAddresses{
		OnRampAddress:    onRampAddr,
		RMNRemoteAddress: rmnRemoteAddr,
	}

	cctpVerifierAddr, cctpVerifierErr := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(cctpverifier.ContractType),
		Qualifier: cctpQualifier,
	}, chainSelector, toAddress)

	cctpResolverAddr, cctpResolverErr := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(versioned_verifier_resolver.CCTPVerifierResolverType),
		Qualifier: cctpQualifier,
	}, chainSelector, toAddress)

	if (cctpVerifierErr == nil) != (cctpResolverErr == nil) {
		return nil, fmt.Errorf(
			"chain %d: cctp verifier and resolver must both exist or both be absent (verifier error: %v, resolver error: %v)",
			chainSelector, cctpVerifierErr, cctpResolverErr,
		)
	}

	if cctpVerifierErr == nil {
		result.CCTPVerifierAddress = cctpVerifierAddr
		result.CCTPVerifierResolverAddress = cctpResolverAddr
	}

	lombardResolverAddr, lombardResolverErr := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(versioned_verifier_resolver.LombardVerifierResolverType),
		Qualifier: lombardQualifier,
	}, chainSelector, toAddress)

	if lombardResolverErr == nil {
		result.LombardVerifierResolverAddress = lombardResolverAddr
	}

	return result, nil
}
