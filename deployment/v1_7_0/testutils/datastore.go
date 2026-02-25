package testutils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldfchain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/onchain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// NewTestBundle creates a new operations bundle for testing.
func NewTestBundle() operations.Bundle {
	lggr, _ := logger.New()
	return operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)
}

// NewStubBlockChains creates stub BlockChains with the given chain selectors.
// These are minimal stubs that only implement ListChainSelectors().
func NewStubBlockChains(selectors []uint64) cldfchain.BlockChains {
	chains := make([]cldfchain.BlockChain, 0, len(selectors))
	for _, sel := range selectors {
		chains = append(chains, cldfevm.Chain{Selector: sel})
	}
	return cldfchain.NewBlockChainsFromSlice(chains)
}

// NewSimulatedEVMEnvironment creates a test environment with simulated EVM chains.
// This uses the same chain infrastructure as the deployments framework, enabling
// real contract deployment and on-chain state reading in tests.
func NewSimulatedEVMEnvironment(t *testing.T, selectors []uint64) (deployment.Environment, []cldfevm.Chain) {
	t.Helper()

	evmChains, blockchains := loadSimulatedChains(t, selectors)
	ds := datastore.NewMemoryDataStore()

	env := newBaseEnvironment(blockchains)
	env.DataStore = ds.Seal()

	return env, evmChains
}

// NewSimulatedEVMEnvironmentWithDataStore creates a test environment with simulated EVM chains
// and returns a mutable datastore that can be used to add contracts before sealing.
func NewSimulatedEVMEnvironmentWithDataStore(t *testing.T, selectors []uint64) (deployment.Environment, datastore.MutableDataStore) {
	t.Helper()

	_, blockchains := loadSimulatedChains(t, selectors)
	ds := datastore.NewMemoryDataStore()

	env := newBaseEnvironment(blockchains)

	return env, ds
}

func loadSimulatedChains(t *testing.T, selectors []uint64) ([]cldfevm.Chain, cldfchain.BlockChains) {
	chains, err := onchain.NewEVMSimLoader().Load(t, selectors)
	require.NoError(t, err)

	evmChains := make([]cldfevm.Chain, 0, len(chains))
	for _, c := range chains {
		evmChain, ok := c.(cldfevm.Chain)
		require.True(t, ok, "expected cldfevm.Chain")
		evmChains = append(evmChains, evmChain)
	}

	return evmChains, cldfchain.NewBlockChainsFromSlice(chains)
}

func newBaseEnvironment(blockchains cldfchain.BlockChains) deployment.Environment {
	lggr, _ := logger.New()
	return deployment.Environment{
		GetContext: func() context.Context { return context.Background() },
		Logger:     lggr,
		OperationsBundle: operations.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			operations.NewMemoryReporter(),
		),
		BlockChains: blockchains,
	}
}
