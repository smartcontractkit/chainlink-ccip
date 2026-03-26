package changesets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
)

// lombard tests use two well-known EVM test selectors so GetSelectorFamily succeeds.
var (
	lombSelA = chainsel.TEST_90000001.Selector
	lombSelB = chainsel.TEST_90000002.Selector
)

func lombardRegistry() *cs_core.MCMSReaderRegistry { return cs_core.GetRegistry() }

func TestDeployLombardChains_Validate_Valid(t *testing.T) {
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelB: {}}},
			lombSelB: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelA: {}}},
		},
	}
	require.NoError(t, cs.VerifyPreconditions(deployment.Environment{}, cfg))
}

func TestDeployLombardChains_Validate_InvalidChainSelector(t *testing.T) {
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{99999: {}},
	}
	require.Error(t, cs.VerifyPreconditions(deployment.Environment{}, cfg))
}

func TestDeployLombardChains_Validate_InvalidLocalAdapterAddress(t *testing.T) {
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {LocalAdapter: "not-a-hex-address"},
		},
	}
	err := cs.VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid local adapter address")
}

func TestDeployLombardChains_Validate_InvalidRemoteChainSelector(t *testing.T) {
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {
				RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{99999: {}},
			},
		},
	}
	require.Error(t, cs.VerifyPreconditions(deployment.Environment{}, cfg))
}

func TestDeployLombardChains_Validate_AsymmetricRemoteConfig(t *testing.T) {
	// A → B remote defined, but B does not list A back.
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelB: {}}},
			lombSelB: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{}},
		},
	}
	err := cs.VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "does not define a remote chain config")
}

func TestDeployLombardChains_Validate_RemoteChainAbsentFromTopLevel(t *testing.T) {
	// A lists B as remote but B is not present in Chains at all.
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelB: {}}},
		},
	}
	err := cs.VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not found in chains")
}

func TestDeployLombardChains_Validate_InvalidRemoteAdapterEncoding(t *testing.T) {
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{
				lombSelB: {RemoteAdapter: "0xZZZZ"}, // invalid hex
			}},
			lombSelB: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelA: {}}},
		},
	}
	err := cs.VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid remote adapter")
}

func TestDeployLombardChains_Validate_RemoteAdapterExceeds32Bytes(t *testing.T) {
	// 33 bytes encoded as hex (0x + 66 hex chars)
	oversized := "0x"
	for i := 0; i < 33; i++ {
		oversized += "aa"
	}
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	cfg := changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{
				lombSelB: {RemoteAdapter: oversized},
			}},
			lombSelB: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelA: {}}},
		},
	}
	err := cs.VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "exceeds 32 bytes")
}

func TestDeployLombardChains_Validate_InvalidMCMSConfig(t *testing.T) {
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	badMCMS := &mcms.Input{} // zero value: TimelockAction is empty → invalid
	cfg := changesets.DeployLombardChainsConfig{
		MCMS:   badMCMS,
		Chains: map[uint64]changesets.LombardChainConfig{},
	}
	err := cs.VerifyPreconditions(deployment.Environment{}, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to validate MCMS input")
}

// ---- Apply tests ----

// mockLombardAdapter is a controllable LombardChain adapter for Apply tests.
type mockLombardAdapter struct {
	deployErr    error
	configureErr error
	deployAddr   string
}

var _ adapters.LombardChain = (*mockLombardAdapter)(nil)

func (m *mockLombardAdapter) DeployLombardChain() *cldf_ops.Sequence[adapters.DeployLombardInput, sequences.OnChainOutput, adapters.DeployLombardChainDeps] {
	return cldf_ops.NewSequence(
		"mock-deploy-lombard",
		semver.MustParse("1.0.0"),
		"Mock lombard deploy",
		func(_ cldf_ops.Bundle, _ adapters.DeployLombardChainDeps, in adapters.DeployLombardInput) (sequences.OnChainOutput, error) {
			if m.deployErr != nil {
				return sequences.OnChainOutput{}, m.deployErr
			}
			addrs := []datastore.AddressRef{}
			if m.deployAddr != "" {
				addrs = append(addrs, datastore.AddressRef{
					ChainSelector: in.ChainSelector,
					Address:       m.deployAddr,
					Type:          datastore.ContractType("LombardTokenPool"),
					Version:       semver.MustParse("1.0.0"),
				})
			}
			return sequences.OnChainOutput{Addresses: addrs}, nil
		},
	)
}

func (m *mockLombardAdapter) ConfigureLombardChainForLanes() *cldf_ops.Sequence[adapters.ConfigureLombardChainForLanesInput, sequences.OnChainOutput, adapters.ConfigureLombardChainForLanesDeps] {
	return cldf_ops.NewSequence(
		"mock-configure-lombard",
		semver.MustParse("1.0.0"),
		"Mock lombard configure",
		func(_ cldf_ops.Bundle, _ adapters.ConfigureLombardChainForLanesDeps, _ adapters.ConfigureLombardChainForLanesInput) (sequences.OnChainOutput, error) {
			if m.configureErr != nil {
				return sequences.OnChainOutput{}, m.configureErr
			}
			return sequences.OnChainOutput{}, nil
		},
	)
}

func (m *mockLombardAdapter) AllowedCallerOnDest(_ datastore.DataStore, _ cldf_chain.BlockChains, _ uint64) ([]byte, error) {
	return []byte("caller"), nil
}
func (m *mockLombardAdapter) RemoteTokenAddress(_ cldf_ops.Bundle, _ datastore.DataStore, _ cldf_chain.BlockChains, _ uint64, _ string) ([]byte, error) {
	return []byte("token"), nil
}
func (m *mockLombardAdapter) RemoteTokenPoolAddress(_ datastore.DataStore, _ cldf_chain.BlockChains, _ uint64, _ string) ([]byte, error) {
	return []byte("pool"), nil
}
func (m *mockLombardAdapter) AddressRefToBytes(_ datastore.AddressRef) ([]byte, error) {
	return []byte("bytes"), nil
}

func newLombardTestEnv(t *testing.T) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	return deployment.Environment{
		Name:        "test",
		BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{}),
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		Logger:      lggr,
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

func TestDeployLombardChains_Apply_NoAdapterRegistered(t *testing.T) {
	env := newLombardTestEnv(t)
	cs := changesets.DeployLombardChains(adapters.NewLombardChainRegistry(), lombardRegistry())

	_, err := cs.Apply(env, changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no CCTP adapter registered")
}

func TestDeployLombardChains_Apply_Success(t *testing.T) {
	env := newLombardTestEnv(t)

	mock := &mockLombardAdapter{deployAddr: "0xLombardPool"}
	registry := adapters.NewLombardChainRegistry()
	registry.RegisterLombardChain(chainsel.FamilyEVM, mock)

	cs := changesets.DeployLombardChains(registry, lombardRegistry())
	out, err := cs.Apply(env, changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, out.DataStore)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, addrs, 1)
	assert.Equal(t, "0xLombardPool", addrs[0].Address)
}

func TestDeployLombardChains_Apply_DeploySequenceError(t *testing.T) {
	env := newLombardTestEnv(t)

	mock := &mockLombardAdapter{deployErr: errors.New("deploy boom")}
	registry := adapters.NewLombardChainRegistry()
	registry.RegisterLombardChain(chainsel.FamilyEVM, mock)

	cs := changesets.DeployLombardChains(registry, lombardRegistry())
	_, err := cs.Apply(env, changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to deploy Lombard")
}

func TestDeployLombardChains_Apply_TwoChainSuccess_WithRemoteChains(t *testing.T) {
	// Exercises combinedDS merge and ConfigureLombardChainForLanes with non-empty RemoteChains.
	env := newLombardTestEnv(t)

	mock := &mockLombardAdapter{deployAddr: "0xLombardPoolA"}
	registry := adapters.NewLombardChainRegistry()
	registry.RegisterLombardChain(chainsel.FamilyEVM, mock)

	cs := changesets.DeployLombardChains(registry, lombardRegistry())
	out, err := cs.Apply(env, changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelB: {}}},
			lombSelB: {RemoteChains: map[uint64]adapters.RemoteLombardChainConfig{lombSelA: {}}},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, out.DataStore)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	// One address per chain (both use the same deployAddr from the mock).
	assert.Len(t, addrs, 2)
}

func TestDeployLombardChains_Apply_ConfigureSequenceError(t *testing.T) {
	env := newLombardTestEnv(t)

	mock := &mockLombardAdapter{configureErr: errors.New("configure boom")}
	registry := adapters.NewLombardChainRegistry()
	registry.RegisterLombardChain(chainsel.FamilyEVM, mock)

	cs := changesets.DeployLombardChains(registry, lombardRegistry())
	_, err := cs.Apply(env, changesets.DeployLombardChainsConfig{
		Chains: map[uint64]changesets.LombardChainConfig{
			lombSelA: {},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to configure Lombard")
}
