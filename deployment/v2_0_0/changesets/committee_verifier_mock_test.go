package changesets_test

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type mockCommitteeVerifierContractAdapter struct {
	contractsByChainAndQualifier map[string][]datastore.AddressRef
	resolveErr                   error
}

func (m *mockCommitteeVerifierContractAdapter) ResolveCommitteeVerifierContracts(
	_ datastore.DataStore,
	chainSelector uint64,
	qualifier string,
) ([]datastore.AddressRef, error) {
	if m.resolveErr != nil {
		return nil, m.resolveErr
	}
	key := fmt.Sprintf("%d:%s", chainSelector, qualifier)
	contracts, ok := m.contractsByChainAndQualifier[key]
	if !ok {
		return nil, fmt.Errorf("no contracts for chain %d qualifier %q", chainSelector, qualifier)
	}
	return contracts, nil
}

func (m *mockCommitteeVerifierContractAdapter) GetCommitteeVerifierResolver(
	ds datastore.DataStore,
	chainSelector uint64,
	qualifier string,
) ([]datastore.AddressRef, error) {
	refs, err := m.ResolveCommitteeVerifierContracts(ds, chainSelector, qualifier)
	if err != nil {
		return nil, err
	}
	resolverType := datastore.ContractType("CommitteeVerifierResolver")
	var out []datastore.AddressRef
	for _, ref := range refs {
		if ref.Type == resolverType {
			out = append(out, ref)
		}
	}
	if len(out) > 0 {
		return out, nil
	}
	if len(refs) == 1 {
		return refs, nil
	}
	return nil, fmt.Errorf("no CommitteeVerifierResolver ref for default lane CCVs (chain %d qualifier %q)", chainSelector, qualifier)
}
