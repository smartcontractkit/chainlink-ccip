package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	cctpverifier "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_verifier"
	lombardverifier "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldfdeployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

type EVMCCVIndexerConfigAdapter struct{}

var _ ccvdeploymentadapters.IndexerConfigAdapter = (*EVMCCVIndexerConfigAdapter)(nil)

func (a *EVMCCVIndexerConfigAdapter) ResolveVerifierAddresses(
	ds datastore.DataStore,
	chainSelector uint64,
	qualifier string,
	kind ccvdeploymentadapters.VerifierKind,
) ([]string, error) {
	resolverType, version, err := resolveEVMCCVIndexerContractMeta(kind)
	if err != nil {
		return nil, err
	}

	refs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByQualifier(qualifier),
		datastore.AddressRefByType(datastore.ContractType(resolverType)),
		datastore.AddressRefByVersion(version),
	)

	if len(refs) == 0 {
		return nil, &ccvdeploymentadapters.MissingIndexerVerifierAddressesError{
			Kind:          kind,
			ChainSelector: chainSelector,
			Qualifier:     qualifier,
		}
	}

	addresses := make([]string, 0, len(refs))
	for _, r := range refs {
		addresses = append(addresses, r.Address)
	}

	return addresses, nil
}

func resolveEVMCCVIndexerContractMeta(kind ccvdeploymentadapters.VerifierKind) (cldfdeployment.ContractType, *semver.Version, error) {
	switch kind {
	case ccvdeploymentadapters.CommitteeVerifierKind:
		return versioned_verifier_resolver.CommitteeVerifierResolverType, versioned_verifier_resolver.Version, nil
	case ccvdeploymentadapters.CCTPVerifierKind:
		return versioned_verifier_resolver.CCTPVerifierResolverType, cctpverifier.Version, nil
	case ccvdeploymentadapters.LombardVerifierKind:
		return versioned_verifier_resolver.LombardVerifierResolverType, lombardverifier.Version, nil
	default:
		return "", nil, fmt.Errorf("unknown verifier kind %q", kind)
	}
}
