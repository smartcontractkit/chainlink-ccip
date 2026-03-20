package adapters

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	dsutil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type EVMTokenVerifierConfigAdapter struct{}

var _ ccvadapters.TokenVerifierConfigAdapter = (*EVMTokenVerifierConfigAdapter)(nil)

func (a *EVMTokenVerifierConfigAdapter) ResolveTokenVerifierAddresses(
	ds datastore.DataStore,
	chainSelector uint64,
	cctpQualifier string,
	lombardQualifier string,
) (*ccvadapters.TokenVerifierChainAddresses, error) {
	toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

	onRampAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type: datastore.ContractType(onramp.ContractType),
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get on ramp address for chain %d: %w", chainSelector, err)
	}

	rmnRemoteAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type: datastore.ContractType(rmn_remote.ContractType),
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get rmn remote address for chain %d: %w", chainSelector, err)
	}

	result := &ccvadapters.TokenVerifierChainAddresses{
		OnRampAddress:    onRampAddr,
		RMNRemoteAddress: rmnRemoteAddr,
	}

	cctpVerifierAddr, cctpVerifierErr := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(cctp_verifier.ContractType),
		Qualifier: cctpQualifier,
	}, chainSelector, toAddress)

	cctpResolverAddr, cctpResolverErr := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
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

	lombardResolverAddr, lombardResolverErr := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(versioned_verifier_resolver.LombardVerifierResolverType),
		Qualifier: lombardQualifier,
	}, chainSelector, toAddress)

	if lombardResolverErr == nil {
		result.LombardVerifierResolverAddress = lombardResolverAddr
	}

	return result, nil
}
