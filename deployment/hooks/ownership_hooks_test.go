package hooks_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"

	"github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

type stubContractOwnership struct{}

type configurableContractOwnership struct {
	filterCfg          *cfgnet.Config
	filterErr          error
	needsOwnershipFunc func(datastore.AddressRef) bool
	verifyErr          error

	mu         sync.Mutex
	verifyRefs map[uint64][]datastore.AddressRef
}

func (s *stubContractOwnership) FilterNetworks(_ string, _ domain.Domain, _ logger.Logger) (*cfgnet.Config, error) {
	return cfgnet.NewConfig(nil), nil
}

func (s *stubContractOwnership) NeedsOwnershipCheck(_ datastore.AddressRef) bool {
	return true
}

func (s *stubContractOwnership) VerifyContractOwnership(
	_ context.Context,
	_ logger.Logger,
	_ datastore.DataStore,
	_ cfgnet.Network,
	_ []datastore.AddressRef,
) error {
	return nil
}

func (c *configurableContractOwnership) FilterNetworks(_ string, _ domain.Domain, _ logger.Logger) (*cfgnet.Config, error) {
	if c.filterErr != nil {
		return nil, c.filterErr
	}
	if c.filterCfg != nil {
		return c.filterCfg, nil
	}
	return cfgnet.NewConfig(nil), nil
}

func (c *configurableContractOwnership) NeedsOwnershipCheck(ref datastore.AddressRef) bool {
	if c.needsOwnershipFunc == nil {
		return true
	}
	return c.needsOwnershipFunc(ref)
}

func (c *configurableContractOwnership) VerifyContractOwnership(
	_ context.Context,
	_ logger.Logger,
	_ datastore.DataStore,
	network cfgnet.Network,
	refsToCheck []datastore.AddressRef,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.verifyRefs == nil {
		c.verifyRefs = map[uint64][]datastore.AddressRef{}
	}
	c.verifyRefs[network.ChainSelector] = append(c.verifyRefs[network.ChainSelector], refsToCheck...)
	return c.verifyErr
}

func writeEnvDatastoreForOwnershipTest(t *testing.T, envName string, dom domain.Domain, refs []datastore.AddressRef) {
	t.Helper()
	envDir := dom.EnvDir(envName)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))
	rawRefs, err := json.Marshal(refs)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), rawRefs, 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))
}

func TestRequireOwnedEnvContractsPreHookForMultipleChainFamilies_DedupesSameFamily(t *testing.T) {
	hooks.ResetContractOwnershipRegistryForTest()
	hooks.GetContractOwnershipRegistry().Register(chainsel.FamilyEVM, &stubContractOwnership{})

	dom := domain.NewDomain(t.TempDir(), "test")
	preHooks := hooks.RequireOwnedEnvContractsPreHookForMultipleChainFamilies(dom, []string{
		chainsel.FamilyEVM,
		chainsel.FamilyEVM,
	})
	require.Len(t, preHooks, 1)
	require.Equal(t, hooks.RequireOwnedEnvContractsHookName, preHooks[0].Name)
	require.Equal(t, changeset.Abort, preHooks[0].FailurePolicy)
}

func TestRequireOwnedEnvContractsPreHookForMultipleChainFamilies_OneHookPerDistinctFamily(t *testing.T) {
	hooks.ResetContractOwnershipRegistryForTest()
	hooks.GetContractOwnershipRegistry().Register(chainsel.FamilyEVM, &stubContractOwnership{})
	hooks.GetContractOwnershipRegistry().Register(chainsel.FamilySolana, &stubContractOwnership{})

	dom := domain.NewDomain(t.TempDir(), "test")
	preHooks := hooks.RequireOwnedEnvContractsPreHookForMultipleChainFamilies(dom, []string{
		chainsel.FamilyEVM,
		chainsel.FamilySolana,
	})
	require.Len(t, preHooks, 2)
	for _, h := range preHooks {
		require.Equal(t, hooks.RequireOwnedEnvContractsHookName, h.Name)
		require.NotNil(t, h.Func)
	}
}

func TestRequireOwnedEnvContractsPreHookForMultipleChainFamilies_PanicsWhenProviderNotRegistered(t *testing.T) {
	hooks.ResetContractOwnershipRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	require.PanicsWithValue(t,
		"no contract ownership provider registered for chain family evm",
		func() {
			_ = hooks.RequireOwnedEnvContractsPreHookForMultipleChainFamilies(dom, []string{chainsel.FamilyEVM})
		},
	)
}

func TestContractOwnershipRegistry_FirstRegistrationWinsAndReset(t *testing.T) {
	hooks.ResetContractOwnershipRegistryForTest()
	r := hooks.GetContractOwnershipRegistry()
	first := &stubContractOwnership{}
	second := &stubContractOwnership{}
	r.Register(chainsel.FamilyEVM, first)
	r.Register(chainsel.FamilyEVM, second)
	got, ok := r.Get(chainsel.FamilyEVM)
	require.True(t, ok)
	require.Same(t, first, got)

	hooks.ResetContractOwnershipRegistryForTest()
	_, ok = hooks.GetContractOwnershipRegistry().Get(chainsel.FamilyEVM)
	require.False(t, ok)
}

func TestNewRequireOwnedEnvContractsPreHook_Definition(t *testing.T) {
	h := hooks.NewRequireOwnedEnvContractsPreHook(domain.Domain{}, &stubContractOwnership{})
	require.Equal(t, hooks.RequireOwnedEnvContractsHookName, h.Name)
	require.Equal(t, changeset.Abort, h.FailurePolicy)
	require.NotNil(t, h.Func)
}

func TestNewRequireOwnedEnvContractsPreHook_FuncRunsWithMinimalEnv(t *testing.T) {
	hooks.ResetContractOwnershipRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"
	writeEnvDatastoreForOwnershipTest(t, envName, dom, nil)

	h := hooks.NewRequireOwnedEnvContractsPreHook(dom, &stubContractOwnership{})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.NoError(t, err)
}

func TestNewRequireOwnedEnvContractsPreHook_DataStoreError(t *testing.T) {
	dom := domain.NewDomain(t.TempDir(), "test")
	h := hooks.NewRequireOwnedEnvContractsPreHook(dom, &stubContractOwnership{})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: "missing-env", Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "require ownership pre-hook: load datastore")
}

func TestNewRequireOwnedEnvContractsPreHook_FilterNetworksError(t *testing.T) {
	dom := domain.NewDomain(t.TempDir(), "test")
	envName := "hooktest"
	writeEnvDatastoreForOwnershipTest(t, envName, dom, nil)
	h := hooks.NewRequireOwnedEnvContractsPreHook(dom, &configurableContractOwnership{
		filterErr: errors.New("boom"),
	})
	err := h.Func(t.Context(), changeset.PreHookParams{
		Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "require ownership pre-hook: load networks")
}

func TestIterateOwnershipCheckers_SkipsNetworksWithoutRefs(t *testing.T) {
	ds := datastore.NewMemoryDataStore().Seal()
	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
		RPCs:          []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})
	verifier := &configurableContractOwnership{}
	err := hooks.IterateOwnershipCheckers(t.Context(), logger.Test(t), ds, cfg, verifier)
	require.NoError(t, err)
	verifier.mu.Lock()
	defer verifier.mu.Unlock()
	require.Empty(t, verifier.verifyRefs)
}

func TestIterateOwnershipCheckers_InvalidRefAndVerifyErrorAreAggregated(t *testing.T) {
	v := semver.MustParse("1.0.0")
	selector := chainsel.ETHEREUM_MAINNET.Selector
	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Type:          "T",
		Version:       nil,
		Address:       "0x1",
		Qualifier:     "bad",
	}))
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Type:          "T",
		Version:       v,
		Address:       "0x0000000000000000000000000000000000000001",
		Qualifier:     "ok",
	}))
	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: selector,
		RPCs:          []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})
	verifier := &configurableContractOwnership{verifyErr: errors.New("verify failed")}

	err := hooks.IterateOwnershipCheckers(t.Context(), logger.Test(t), mds.Seal(), cfg, verifier)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid address ref")
	require.ErrorContains(t, err, "verify failed")
}

func TestIterateOwnershipCheckers_AppliesNeedsOwnershipCheckFilter(t *testing.T) {
	v := semver.MustParse("1.0.0")
	selector := chainsel.ETHEREUM_MAINNET.Selector
	mds := datastore.NewMemoryDataStore()
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Type:          "Keep",
		Version:       v,
		Address:       "0x0000000000000000000000000000000000000001",
		Qualifier:     "a",
	}))
	require.NoError(t, mds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Type:          "Drop",
		Version:       v,
		Address:       "0x0000000000000000000000000000000000000002",
		Qualifier:     "b",
	}))
	cfg := cfgnet.NewConfig([]cfgnet.Network{{
		Type:          cfgnet.NetworkTypeMainnet,
		ChainSelector: selector,
		RPCs:          []cfgnet.RPC{{HTTPURL: "http://localhost"}},
	}})
	verifier := &configurableContractOwnership{
		needsOwnershipFunc: func(ref datastore.AddressRef) bool { return ref.Type == "Keep" },
	}

	err := hooks.IterateOwnershipCheckers(t.Context(), logger.Test(t), mds.Seal(), cfg, verifier)
	require.NoError(t, err)
	verifier.mu.Lock()
	refs := verifier.verifyRefs[selector]
	verifier.mu.Unlock()
	require.Len(t, refs, 1)
	require.Equal(t, datastore.ContractType("Keep"), refs[0].Type)
}
