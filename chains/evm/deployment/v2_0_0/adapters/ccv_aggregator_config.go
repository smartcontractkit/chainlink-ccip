package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	dsutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

type EVMCCVAggregatorConfigAdapter struct{}

// ResolveDestinationVerifierAddress implements [adapters.AggregatorConfigAdapter].
func (a *EVMCCVAggregatorConfigAdapter) ResolveDestinationVerifierAddress(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error) {
	return a.resolveVerifierAddress(ds, chainSelector, qualifier)
}

// ResolveSourceVerifierAddress implements [adapters.AggregatorConfigAdapter].
func (a *EVMCCVAggregatorConfigAdapter) ResolveSourceVerifierAddress(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error) {
	return a.resolveVerifierAddress(ds, chainSelector, qualifier)
}

var _ ccvdeploymentadapters.AggregatorConfigAdapter = (*EVMCCVAggregatorConfigAdapter)(nil)

func (a *EVMCCVAggregatorConfigAdapter) resolveVerifierAddress(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error) {
	return dsutils.FindAndFormatFirstRef(ds, chainSelector,
		func(r datastore.AddressRef) (string, error) { return r.Address, nil },
		datastore.AddressRef{
			Type:      datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
			Qualifier: qualifier,
		},
		datastore.AddressRef{
			Type:      datastore.ContractType(committee_verifier.ContractType),
			Qualifier: qualifier,
		},
	)
}

func (a *EVMCCVAggregatorConfigAdapter) GetDeployedChains(ds datastore.DataStore, qualifier string) []uint64 {
	if ds == nil {
		return nil
	}
	resolverRefs := ds.Addresses().Filter(
		datastore.AddressRefByQualifier(qualifier),
		datastore.AddressRefByType(datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType)),
		datastore.AddressRefByVersion(versioned_verifier_resolver.Version),
	)
	verifierRefs := ds.Addresses().Filter(
		datastore.AddressRefByQualifier(qualifier),
		datastore.AddressRefByType(datastore.ContractType(committee_verifier.ContractType)),
		datastore.AddressRefByVersion(committee_verifier.Version),
	)
	seen := make(map[uint64]struct{}, len(resolverRefs)+len(verifierRefs))
	chains := make([]uint64, 0, len(resolverRefs)+len(verifierRefs))
	for _, refs := range [][]datastore.AddressRef{resolverRefs, verifierRefs} {
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
	}
	return chains
}
