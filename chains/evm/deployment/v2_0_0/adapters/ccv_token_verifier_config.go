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
		Type:    datastore.ContractType(onrampop.ContractType),
		Version: onrampop.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get on ramp address for chain %d: %w", chainSelector, err)
	}

	rmnRemoteAddr, err := dsutils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(rmnremote.ContractType),
		Version: rmnremote.Version,
	}, chainSelector, toAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get rmn remote address for chain %d: %w", chainSelector, err)
	}

	result := &ccvdeploymentadapters.TokenVerifierChainAddresses{
		OnRampAddress:    onRampAddr,
		RMNRemoteAddress: rmnRemoteAddr,
	}

	cctpVerifierRefs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByType(datastore.ContractType(cctpverifier.ContractType)),
		datastore.AddressRefByQualifier(cctpQualifier),
		datastore.AddressRefByVersion(cctpverifier.Version),
	)
	if len(cctpVerifierRefs) > 1 {
		return nil, fmt.Errorf("chain %d: expected at most 1 CCTPVerifier with qualifier %q, found %d", chainSelector, cctpQualifier, len(cctpVerifierRefs))
	}

	cctpResolverRefs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByType(datastore.ContractType(versioned_verifier_resolver.CCTPVerifierResolverType)),
		datastore.AddressRefByQualifier(cctpQualifier),
		datastore.AddressRefByVersion(versioned_verifier_resolver.Version),
	)
	if len(cctpResolverRefs) > 1 {
		return nil, fmt.Errorf("chain %d: expected at most 1 CCTPVerifierResolver with qualifier %q, found %d", chainSelector, cctpQualifier, len(cctpResolverRefs))
	}

	if (len(cctpVerifierRefs) == 1) != (len(cctpResolverRefs) == 1) {
		return nil, fmt.Errorf(
			"chain %d: CCTP verifier and resolver must both exist or both be absent (verifier found: %v, resolver found: %v)",
			chainSelector, len(cctpVerifierRefs) == 1, len(cctpResolverRefs) == 1,
		)
	}
	if len(cctpVerifierRefs) == 1 {
		result.CCTPVerifierAddress = cctpVerifierRefs[0].Address
		result.CCTPVerifierResolverAddress = cctpResolverRefs[0].Address
	}

	lombardResolverRefs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByType(datastore.ContractType(versioned_verifier_resolver.LombardVerifierResolverType)),
		datastore.AddressRefByQualifier(lombardQualifier),
		datastore.AddressRefByVersion(versioned_verifier_resolver.Version),
	)
	if len(lombardResolverRefs) > 1 {
		return nil, fmt.Errorf("chain %d: expected at most 1 LombardVerifierResolver with qualifier %q, found %d", chainSelector, lombardQualifier, len(lombardResolverRefs))
	}
	if len(lombardResolverRefs) == 1 {
		result.LombardVerifierResolverAddress = lombardResolverRefs[0].Address
	}

	return result, nil
}
