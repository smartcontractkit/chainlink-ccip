package changesets_test

import (
	"context"
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
)

// newGlamsterdamTestEnv builds a minimal Environment with one real EVM chain (so
// chain_selectors.GetSelectorFamily resolves it), matching
// ConfigureChainsForLanesFromTopology's test convention.
func newGlamsterdamTestEnv(t *testing.T, chainSelectors []uint64) deployment.Environment {
	t.Helper()
	chains := make(map[uint64]cldf_chain.BlockChain, len(chainSelectors))
	for _, sel := range chainSelectors {
		chains[sel] = cldfevm.Chain{Selector: sel}
	}
	lggr := logger.Test(t)
	return deployment.Environment{
		BlockChains: cldf_chain.NewBlockChains(chains),
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		Logger:      lggr,
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

// TestUpdateGasConfigForGlamsterdamV200_NoAdapterRegistered is a regression test: a fresh,
// un-registered GasUpdateAdapterRegistry (injected rather than the global singleton, so no other
// test's registration can leak in) must fail the changeset explicitly for a chain family with
// candidate chains but no registered adapter — automation must not be able to mistake an
// unexecuted migration for a completed one.
func TestUpdateGasConfigForGlamsterdamV200_NoAdapterRegistered(t *testing.T) {
	const targetSel = uint64(999)
	srcSel := chainsel.TEST_90000001.Selector

	registry := cs_core.GetRegistry()
	adapterRegistry := adapters.NewGasUpdateAdapterRegistry()
	e := newGlamsterdamTestEnv(t, []uint64{srcSel, targetSel})

	_, err := changesets.UpdateGasConfigForGlamsterdamV200(registry, adapterRegistry).Apply(e, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV200Cfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.GlamsterdamGasUpdateV200Cfg{
			TargetChainSelector: targetSel,
		},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "no gas update adapter registered for chain family")
	require.Contains(t, err.Error(), chainsel.FamilyEVM)
}

func TestUpdateGasConfigForGlamsterdamV200_ValidateRequiresTarget(t *testing.T) {
	registry := cs_core.GetRegistry()
	adapterRegistry := adapters.NewGasUpdateAdapterRegistry()
	_, err := changesets.UpdateGasConfigForGlamsterdamV200(registry, adapterRegistry).Apply(deployment.Environment{}, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV200Cfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.GlamsterdamGasUpdateV200Cfg{},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "TargetChainSelector must be set")
}

func TestUpdateGasConfigForGlamsterdamV200_ValidateRequiresAdapterRegistry(t *testing.T) {
	registry := cs_core.GetRegistry()
	_, err := changesets.UpdateGasConfigForGlamsterdamV200(registry, nil).Apply(deployment.Environment{}, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV200Cfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.GlamsterdamGasUpdateV200Cfg{
			TargetChainSelector: 123,
		},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "gas update adapter registry is required")
}
