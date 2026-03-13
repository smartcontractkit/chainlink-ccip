package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type EVMIndexerConfigAdapter struct{}

var _ ccvadapters.IndexerConfigAdapter = (*EVMIndexerConfigAdapter)(nil)

func (a *EVMIndexerConfigAdapter) ResolveVerifierAddresses(
	ds datastore.DataStore, chainSelector uint64, qualifier string, kind ccvadapters.VerifierKind,
) ([]string, error) {
	resolverType, version, err := resolveContractMeta(kind)
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
		return nil, fmt.Errorf("no %s verifier addresses found for chain %d with qualifier %q", kind, chainSelector, qualifier)
	}

	addresses := make([]string, 0, len(refs))
	for _, r := range refs {
		addresses = append(addresses, r.Address)
	}

	return addresses, nil
}

func resolveContractMeta(kind ccvadapters.VerifierKind) (deployment.ContractType, *semver.Version, error) {
	switch kind {
	case ccvadapters.CommitteeVerifierKind:
		return versioned_verifier_resolver.CommitteeVerifierResolverType, versioned_verifier_resolver.Version, nil
	case ccvadapters.CCTPVerifierKind:
		return cctp_verifier.ResolverType, cctp_verifier.Version, nil
	case ccvadapters.LombardVerifierKind:
		return lombard_verifier.ResolverType, lombard_verifier.Version, nil
	default:
		return "", nil, fmt.Errorf("unknown verifier kind %q", kind)
	}
}
