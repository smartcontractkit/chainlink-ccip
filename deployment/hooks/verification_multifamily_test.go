package hooks_test

import (
	"context"
	"fmt"
	"os"
	"testing"

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

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_DedupesSameFamily(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []uint64{
		chainsel.ETHEREUM_MAINNET.Selector,
		chainsel.POLYGON_MAINNET.Selector,
	})
	require.Len(t, postHooks, 1, "EVM chains share one family; expect one post hook")
	require.Equal(t, hooks.VerifyDeployedContractsHookName, postHooks[0].Name)
	require.Equal(t, changeset.Abort, postHooks[0].FailurePolicy)
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_OneHookPerDistinctFamily(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})
	r.Register(chainsel.FamilySolana, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []uint64{
		chainsel.ETHEREUM_MAINNET.Selector,
		chainsel.SOLANA_MAINNET.Selector,
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
		"no contract verification provider registered for chain selector 5009297550715157269",
		func() {
			_ = hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []uint64{chainsel.ETHEREUM_MAINNET.Selector})
		},
	)
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_PanicsWhenSomeProvidersNotRegistered(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	require.PanicsWithValue(t,
		fmt.Sprintf("no contract verification provider registered for chain selector %d", chainsel.SOLANA_MAINNET.Selector),
		func() {
			_ = hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []uint64{
				chainsel.ETHEREUM_MAINNET.Selector,
				chainsel.SOLANA_MAINNET.Selector, // solana provider not registered
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
		_ = hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []uint64{9999999999999999999})
	})
}

func TestVerifyDeployedContractsPostHookForMultipleChainFamilies_FuncRuns(t *testing.T) {
	hooks.ResetContractVerificationRegistryForTest()
	r := hooks.GetContractVerificationRegistry()
	r.Register(chainsel.FamilyEVM, &stubContractVerification{})

	dom := domain.NewDomain(t.TempDir(), "test")
	postHooks := hooks.VerifyDeployedContractsPostHookForMultipleChainFamilies(dom, []uint64{chainsel.ETHEREUM_MAINNET.Selector})
	require.Len(t, postHooks, 1)

	err := postHooks[0].Func(t.Context(), changeset.PostHookParams{
		Env: changeset.HookEnv{Name: "staging", Logger: logger.Test(t)},
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
		"no contract verification provider registered for chain selector 5009297550715157269",
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
		fmt.Sprintf("no contract verification provider registered for chain selector %d", chainsel.SOLANA_MAINNET.Selector),
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
	envDir := dom.EnvDir(envName)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))

	refs := []datastore.AddressRef{{
		Type:    "C",
		Version: nil,
	}}
	preHooks := hooks.RequireVerifiedEnvContractsPreHookForMultipleChainFamilies(dom, []uint64{
		chainsel.ETHEREUM_MAINNET.Selector,
		chainsel.SOLANA_MAINNET.Selector,
	}, refs)
	require.Len(t, preHooks, 2)

	// Both family hooks should accept the same refs slice without error (empty datastore matches no refs after filter).
	for _, h := range preHooks {
		err := h.Func(t.Context(), changeset.PreHookParams{
			Env: changeset.HookEnv{Name: envName, Logger: logger.Test(t)},
		})
		require.NoError(t, err)
	}
}
