package hooks_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	cldverification "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/verification"

	"github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

// stubContractVerification implements hooks.ContractVerification with no-op behavior so
// post/pre hooks exit quickly (empty network list, skip every network).
type stubContractVerification struct {
	filterErr error
}

func (s *stubContractVerification) FilterNetworks(_ string, _ domain.Domain, _ logger.Logger) (*cfgnet.Config, error) {
	if s.filterErr != nil {
		return nil, s.filterErr
	}
	return cfgnet.NewConfig(nil), nil
}

func (s *stubContractVerification) NeedsVerification(_ datastore.AddressRef) bool {
	return false
}

func (s *stubContractVerification) ForEachNetwork(
	_ context.Context,
	_ cfgnet.Network,
	_ uint64,
	_ logger.Logger,
	_ string,
) (hooks.VerifierBuilderForNetwork, bool) {
	return func(_ context.Context, _ datastore.AddressRef) (cldverification.Verifiable, error) {
		return nil, nil
	}, true
}

type noopVerifiable struct{}

func (noopVerifiable) String() string { return "noop" }

func (noopVerifiable) IsVerified(context.Context) (bool, error) { return true, nil }

func (noopVerifiable) Verify(context.Context) error { return nil }

// iterateParallelTestVerifier drives IterateVerifiers with configurable networks and always
// returns a cheap noop Verifiable so the step callback is the locus of concurrency.
type iterateParallelTestVerifier struct {
	networks []cfgnet.Network
}

func (v *iterateParallelTestVerifier) FilterNetworks(_ string, _ domain.Domain, _ logger.Logger) (*cfgnet.Config, error) {
	return cfgnet.NewConfig(v.networks), nil
}

func (v *iterateParallelTestVerifier) NeedsVerification(_ datastore.AddressRef) bool {
	return true
}

func (v *iterateParallelTestVerifier) ForEachNetwork(
	_ context.Context,
	_ cfgnet.Network,
	_ uint64,
	_ logger.Logger,
	_ string,
) (hooks.VerifierBuilderForNetwork, bool) {
	return func(_ context.Context, _ datastore.AddressRef) (cldverification.Verifiable, error) {
		return noopVerifiable{}, nil
	}, false
}

// selectorRecordingVerifier exposes three EVM networks and records which chain selector each
// built address ref had, so tests can assert refsToVerify filtering kept the intended chain.
type selectorRecordingVerifier struct {
	mu               sync.Mutex
	seenChainSel     []uint64
	ethereumSelector uint64
	polygonSelector  uint64
	arbitrumSelector uint64
}

func newSelectorRecordingVerifier(ethSel, polySel, arbSel uint64) *selectorRecordingVerifier {
	return &selectorRecordingVerifier{
		ethereumSelector: ethSel,
		polygonSelector:  polySel,
		arbitrumSelector: arbSel,
	}
}

func (v *selectorRecordingVerifier) FilterNetworks(_ string, _ domain.Domain, _ logger.Logger) (*cfgnet.Config, error) {
	return cfgnet.NewConfig([]cfgnet.Network{
		{Type: cfgnet.NetworkTypeMainnet, ChainSelector: v.ethereumSelector, RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}}},
		{Type: cfgnet.NetworkTypeMainnet, ChainSelector: v.polygonSelector, RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}}},
		{Type: cfgnet.NetworkTypeMainnet, ChainSelector: v.arbitrumSelector, RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}}},
	}), nil
}

func (v *selectorRecordingVerifier) NeedsVerification(_ datastore.AddressRef) bool {
	return true
}

func (v *selectorRecordingVerifier) ForEachNetwork(
	_ context.Context,
	_ cfgnet.Network,
	_ uint64,
	_ logger.Logger,
	_ string,
) (hooks.VerifierBuilderForNetwork, bool) {
	return func(_ context.Context, ref datastore.AddressRef) (cldverification.Verifiable, error) {
		v.mu.Lock()
		v.seenChainSel = append(v.seenChainSel, ref.ChainSelector)
		v.mu.Unlock()
		return noopVerifiable{}, nil
	}, false
}

func writeEnvDatastoreWithRefs(t *testing.T, env string, dom domain.Domain, refs []datastore.AddressRef) {
	t.Helper()

	envDir := dom.EnvDir(env)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))

	refsJSON, err := json.Marshal(refs)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), refsJSON, 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_DedupesSameFamily(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{
		chainsel.FamilyEVM,
	})
	require.Len(t, postHooks, 1, "EVM chains share one family; expect one post hook")
	require.Equal(t, hooks.VerifyDeployedContractsHookName, postHooks[0].Name)
	require.Equal(t, changeset.Warn, postHooks[0].FailurePolicy)
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_OneHookPerDistinctFamily(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})
	r.Register(chainsel.FamilySolana, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{
		chainsel.FamilyEVM,
		chainsel.FamilySolana,
	})
	require.Len(t, postHooks, 2)
	for _, h := range postHooks {
		require.Equal(t, hooks.VerifyDeployedContractsHookName, h.Name)
		require.NotNil(t, h.Func)
	}
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_PanicsWhenProviderNotRegistered(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()

	dom := domain.NewDomain(t.TempDir(), "test")
	require.PanicsWithValue(t,
		"no contract verification provider registered for chain family evm",
		func() {
			_ = hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{chainsel.FamilyEVM})
		},
	)
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_PanicsWhenSomeProvidersNotRegistered(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	require.PanicsWithValue(t,
		fmt.Sprintf("no contract verification provider registered for chain family solana"),
		func() {
			_ = hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{
				chainsel.FamilyEVM,
				chainsel.FamilySolana, // solana provider not registered
			})
		},
	)
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_PanicsOnUnknownSelector(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	require.Panics(t, func() {
		_ = hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{"unsupported"})
	})
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_FuncRuns(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{chainsel.FamilyEVM})
	require.Len(t, postHooks, 1)

	err := postHooks[0].Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: datastore.NewMemoryDataStore(),
		},
	})
	require.NoError(t, err)
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_FuncRunsWithChangesetParams(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{chainsel.FamilyEVM})
	require.Len(t, postHooks, 1)

	err := postHooks[0].Func(t.Context(), changeset.PostHookParams{
		Env:          changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		ChangesetKey: "test-changeset",
		Config:       map[string]any{"foo": "bar"},
		Output: cldf.ChangesetOutput{
			DataStore: datastore.NewMemoryDataStore(),
		},
	})
	require.NoError(t, err)
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_DedupesSameFamily(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	preHooks := hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{
		chainsel.ETHEREUM_MAINNET.Selector,
		chainsel.POLYGON_MAINNET.Selector,
	}, nil)
	require.Len(t, preHooks, 1)
	require.Equal(t, hooks.RequireVerifiedEnvContractsHookName, preHooks[0].Name)
	require.Equal(t, changeset.Abort, preHooks[0].FailurePolicy)
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_OneHookPerDistinctFamily(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})
	r.Register(chainsel.FamilySolana, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	preHooks := hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{
		chainsel.ETHEREUM_MAINNET.Selector,
		chainsel.SOLANA_MAINNET.Selector,
	}, nil)
	require.Len(t, preHooks, 2)
	for _, h := range preHooks {
		require.Equal(t, hooks.RequireVerifiedEnvContractsHookName, h.Name)
		require.NotNil(t, h.Func)
	}
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_PanicsWhenProviderNotRegistered(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()

	dom := domain.NewDomain(t.TempDir(), "test")
	require.PanicsWithValue(t,
		"no contract verification provider registered for chain family evm",
		func() {
			_ = hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{chainsel.ETHEREUM_MAINNET.Selector}, nil)
		},
	)
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_PanicsWhenSomeProvidersNotRegistered(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	require.PanicsWithValue(t,
		fmt.Sprintf("no contract verification provider registered for chain family solana"),
		func() {
			_ = hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{
				chainsel.ETHEREUM_MAINNET.Selector,
				chainsel.SOLANA_MAINNET.Selector, // solana provider not registered
			}, nil)
		},
	)
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_FuncRunsWithMinimalEnv(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"
	envDir := dom.EnvDir(envName)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))

	preHooks := hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{chainsel.ETHEREUM_MAINNET.Selector}, nil)
	require.Len(t, preHooks, 1)

	err := preHooks[0].Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.NoError(t, err)
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_PassesRefsToEachFamilyHook(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})
	r.Register(chainsel.FamilySolana, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"

	refsInput := []datastore.AddressRef{{
		Type:    "C",
		Version: nil,
	}}
	refs := []datastore.AddressRef{
		{
			ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
			Type:          "C",
			Version:       nil,
			Address:       "0x0000000000000000000000000000000000000001",
		},
		{
			ChainSelector: chainsel.SOLANA_MAINNET.Selector,
			Type:          "C",
			Version:       nil,
			Address:       "0x0000000000000000000000000000000000000002",
		},
	}
	writeEnvDatastoreWithRefs(t, envName, dom, refs)
	preHooks := hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{
		chainsel.ETHEREUM_MAINNET.Selector,
		chainsel.SOLANA_MAINNET.Selector,
	}, refsInput)
	require.Len(t, preHooks, 2)

	// Both family hooks should accept the same refs slice without error.
	for _, h := range preHooks {
		err := h.Func(t.Context(), changeset.PreHookParams{
			Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
		})
		require.NoError(t, err)
	}
}

// TestIterateVerifiers_VerificationStepsSerializedForExplorerRateLimit spawns more address refs
// than the per-network errgroup limit, but IterateVerifiers holds a global mutex around each
// verification step plus a post-step cooldown for explorer APIs. Peak concurrent step executions
// should therefore stay at one.
func TestIterateVerifiers_VerificationStepsSerializedForExplorerRateLimit(t *testing.T) {
	t.Parallel()

	chain, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)

	var (
		mu            sync.Mutex
		inFlight      int
		maxConcurrent int
	)

	verifier := &iterateParallelTestVerifier{
		networks: []cfgnet.Network{{
			Type:          cfgnet.NetworkTypeMainnet,
			ChainSelector: chain.Selector,
			RPCs:          []cfgnet.RPC{{HTTPURL: "http://localhost"}},
		}},
	}

	// AddressRef store keys are (chain, type, version, qualifier) — not address. Use a unique
	// qualifier per row so six contracts on the same chain can coexist.
	addrs := []string{
		"0x1000000000000000000000000000000000000001",
		"0x2000000000000000000000000000000000000001",
		"0x3000000000000000000000000000000000000001",
		"0x4000000000000000000000000000000000000001",
		"0x5000000000000000000000000000000000000001",
		"0x6000000000000000000000000000000000000001",
	}

	ds := datastore.NewMemoryDataStore()
	for i, addr := range addrs {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain.Selector,
			Type:          "C",
			Version:       semver.MustParse("1.0.0"),
			Qualifier:     fmt.Sprintf("p%d", i),
			Address:       addr,
		}))
	}

	cfg, err := verifier.FilterNetworks("", domain.Domain{}, logger.Test(t))
	require.NoError(t, err)

	err = hooks.IterateVerifiers(t.Context(), ds.Seal(), cfg, logger.Test(t), "test", verifier,
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			mu.Lock()
			inFlight++
			if inFlight > maxConcurrent {
				maxConcurrent = inFlight
			}
			cur := inFlight
			mu.Unlock()

			if cur > 1 {
				t.Errorf("in-flight steps %d exceeds serialized limit 1", cur)
			}

			time.Sleep(5 * time.Millisecond)

			mu.Lock()
			inFlight--
			mu.Unlock()
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, 1, maxConcurrent,
		"explorer rate-limit mutex should serialize steps so peak in-flight is 1")
}

// TestIterateVerifiers_MultipleNetworksAllRefsProcessed runs three networks with one ref each and
// asserts every ref receives a step. Network errgroups still run concurrently, but verification
// steps are serialized globally with a cooldown, so wall time is not used as a signal here.
func TestIterateVerifiers_MultipleNetworksAllRefsProcessed(t *testing.T) {
	t.Parallel()

	eth, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)
	poly, ok := chainsel.ChainBySelector(chainsel.POLYGON_MAINNET.Selector)
	require.True(t, ok)
	arb, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET_ARBITRUM_1.Selector)
	require.True(t, ok)

	verifier := &iterateParallelTestVerifier{
		networks: []cfgnet.Network{
			{Type: cfgnet.NetworkTypeMainnet, ChainSelector: eth.Selector, RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}}},
			{Type: cfgnet.NetworkTypeMainnet, ChainSelector: poly.Selector, RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}}},
			{Type: cfgnet.NetworkTypeMainnet, ChainSelector: arb.Selector, RPCs: []cfgnet.RPC{{HTTPURL: "http://localhost"}}},
		},
	}

	ds := datastore.NewMemoryDataStore()
	for _, ch := range []chainsel.Chain{eth, poly, arb} {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: ch.Selector,
			Type:          "C",
			Version:       semver.MustParse("1.0.0"),
			Address:       "0x0000000000000000000000000000000000000001",
		}))
	}

	cfg, err := verifier.FilterNetworks("", domain.Domain{}, logger.Test(t))
	require.NoError(t, err)

	var stepCount atomic.Int32
	err = hooks.IterateVerifiers(t.Context(), ds.Seal(), cfg, logger.Test(t), "test", verifier,
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			stepCount.Add(1)
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, int32(3), stepCount.Load(), "each network's ref should run the step once")
}

func TestNewVerifyDeployedContractsPostHook_SkipsWhenApplyFailed(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	h := hooks.NewVerifyDeployedContractsPostHook(dom, &stubContractVerification{})
	err := h.Func(t.Context(), changeset.PostHookParams{
		Err: errors.New("apply failed"),
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: datastore.NewMemoryDataStore(),
		},
	})
	require.NoError(t, err)
}

func TestNewVerifyDeployedContractsPostHook_SkipsWhenDataStoreNil(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	h := hooks.NewVerifyDeployedContractsPostHook(dom, &stubContractVerification{})
	err := h.Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: nil,
		},
	})
	require.NoError(t, err)
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_SkipsWhenApplyFailed(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []string{chainsel.FamilyEVM})
	require.Len(t, postHooks, 1)

	err := postHooks[0].Func(t.Context(), changeset.PostHookParams{
		Err: errors.New("apply failed"),
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
		Output: cldf.ChangesetOutput{
			DataStore: datastore.NewMemoryDataStore(),
		},
	})
	require.NoError(t, err)
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_PanicsOnInvalidChainSelector(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	require.PanicsWithValue(t,
		"invalid chain selector 0: unknown chain selector 0",
		func() {
			_ = hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{0}, nil)
		},
	)
}

func TestNewRequireVerifiedEnvContractsPreHook_FilterRefs_InvalidEmptyCriteria(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"
	envDir := dom.EnvDir(envName)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))

	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, &stubContractVerification{},
		[]datastore.AddressRef{{}}, []uint64{chainsel.ETHEREUM_MAINNET.Selector})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid filter criteria")
}

func TestNewRequireVerifiedEnvContractsPreHook_FilterRefs_NoMatchingAddressRef(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"

	refs := []datastore.AddressRef{{
		ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
		Type:          "OnChain",
		Address:       "0x0000000000000000000000000000000000000001",
	}}
	writeEnvDatastoreWithRefs(t, envName, dom, refs)

	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, &stubContractVerification{},
		[]datastore.AddressRef{{Type: "OtherType"}}, []uint64{chainsel.ETHEREUM_MAINNET.Selector})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "no address ref found for filter criteria")
}

func TestNewRequireVerifiedEnvContractsPreHook_FilterRefs_SelectsMatchingRefs(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"

	v := semver.MustParse("1.0.0")
	refs := []datastore.AddressRef{
		{
			ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
			Type:          "Keep",
			Version:       v,
			Address:       "0x0000000000000000000000000000000000000001",
		},
		{
			ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
			Type:          "Drop",
			Version:       v,
			Address:       "0x0000000000000000000000000000000000000002",
		},
	}
	writeEnvDatastoreWithRefs(t, envName, dom, refs)

	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, &stubContractVerification{},
		[]datastore.AddressRef{{Type: "Keep"}}, []uint64{chainsel.ETHEREUM_MAINNET.Selector})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.NoError(t, err)
}

func TestNewRequireVerifiedEnvContractsPreHook_FilterRefs_SelectsMatchingSelectors(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"

	eth := chainsel.ETHEREUM_MAINNET.Selector
	poly := chainsel.POLYGON_MAINNET.Selector
	arb := chainsel.ETHEREUM_MAINNET_ARBITRUM_1.Selector

	v := semver.MustParse("1.0.0")
	refs := []datastore.AddressRef{
		{
			ChainSelector: eth,
			Type:          "SameType",
			Version:       v,
			Address:       "0x0000000000000000000000000000000000000001",
		},
		{
			ChainSelector: poly,
			Type:          "SameType",
			Version:       v,
			Address:       "0x0000000000000000000000000000000000000002",
		},
		{
			ChainSelector: arb,
			Type:          "SameType",
			Version:       v,
			Address:       "0x0000000000000000000000000000000000000003",
		},
	}
	writeEnvDatastoreWithRefs(t, envName, dom, refs)

	rec := newSelectorRecordingVerifier(eth, poly, arb)
	h := hooks.NewRequireVerifiedEnvContractsPreHook(dom, rec,
		[]datastore.AddressRef{{
			ChainSelector: poly,
			Type:          "SameType",
			Version:       v,
		}},
		[]uint64{poly},
	)
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.NoError(t, err)

	rec.mu.Lock()
	seen := append([]uint64(nil), rec.seenChainSel...)
	rec.mu.Unlock()
	require.Equal(t, []uint64{poly}, seen,
		"refsToVerify should narrow to the Polygon ref; other selectors must not be verified")
	require.NotContains(t, seen, eth, "Ethereum ref should not be verified")
	require.NotContains(t, seen, arb, "Arbitrum ref should not be verified")
}

func TestRequireVerifiedEnvContractsPreHookForMultipleChainFamilies_FilterRefs_InvalidEmptyCriteria(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"
	envDir := dom.EnvDir(envName)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))

	preHooks := hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom,
		[]uint64{chainsel.ETHEREUM_MAINNET.Selector}, []datastore.AddressRef{{}})
	require.Len(t, preHooks, 1)

	err := preHooks[0].Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid filter criteria")
}
