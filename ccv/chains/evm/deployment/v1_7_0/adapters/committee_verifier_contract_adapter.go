package adapters

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type EVMCommitteeVerifierContractAdapter struct{}

var _ ccvadapters.CommitteeVerifierContractAdapter = (*EVMCommitteeVerifierContractAdapter)(nil)

func (a *EVMCommitteeVerifierContractAdapter) ResolveCommitteeVerifierContracts(
	ds datastore.DataStore,
	chainSelector uint64,
	qualifier string,
) ([]datastore.AddressRef, error) {
	verifier, err := ds.Addresses().Get(datastore.NewAddressRefKey(
		chainSelector,
		datastore.ContractType(committee_verifier.ContractType),
		committee_verifier.Version,
		qualifier,
	))
	if err != nil {
		return nil, fmt.Errorf("committee verifier not found for chain %d qualifier %q: %w", chainSelector, qualifier, err)
	}

	resolver, err := ds.Addresses().Get(datastore.NewAddressRefKey(
		chainSelector,
		datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
		versioned_verifier_resolver.Version,
		qualifier,
	))
	if err != nil {
		return nil, fmt.Errorf("committee verifier resolver not found for chain %d qualifier %q: %w", chainSelector, qualifier, err)
	}

	return []datastore.AddressRef{verifier, resolver}, nil
}
