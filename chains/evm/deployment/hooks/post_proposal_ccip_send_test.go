package hooks

import (
	"context"
	"os"
	"testing"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"

	cciphooks "github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

func blockChainsWithSelectors(selectors ...uint64) cldf_chain.BlockChains {
	chains := make(map[uint64]cldf_chain.BlockChain)
	for _, sel := range selectors {
		chains[sel] = cldf_evm.Chain{Selector: sel}
	}
	return cldf_chain.NewBlockChains(chains)
}

func writeMinimalEnvDatastore(t *testing.T, dom domain.Domain, env string) {
	t.Helper()

	envDir := dom.EnvDir(env)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))
}

func useOnlyEVMPostProposalProvider(t *testing.T) {
	t.Helper()

	cciphooks.ResetPostProposalCCIPSendRegistryForTest()
	cciphooks.GetPostProposalCCIPSendRegistry().Register(chain_selectors.FamilyEVM, &EVMPostProposalCCIPSend{})

	// Keep registry in the package's default state for tests that run later.
	t.Cleanup(func() {
		cciphooks.ResetPostProposalCCIPSendRegistryForTest()
		cciphooks.GetPostProposalCCIPSendRegistry().Register(chain_selectors.FamilyEVM, &EVMPostProposalCCIPSend{})
	})
}

func TestEVMPostProposalCCIPSend_SkinSend(t *testing.T) {
	p := &EVMPostProposalCCIPSend{}

	tests := []struct {
		name string
		env  cldf.Environment
		want bool
	}{
		{
			name: "non forked env skips send",
			env:  cldf.Environment{Name: "staging"},
			want: true,
		},
		{
			name: "forked env without node ids runs send",
			env:  cldf.Environment{Name: "proposal-fork"},
			want: false,
		},
		{
			name: "forked env with node ids skips send",
			env: cldf.Environment{
				Name:    "proposal-fork",
				NodeIDs: []string{"node-1"},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, p.SkipSend(tt.env))
		})
	}
}

func TestEVMPostProposalCCIPSend_PreSendValidation(t *testing.T) {
	p := &EVMPostProposalCCIPSend{}
	ethSel := chain_selectors.ETHEREUM_MAINNET.Selector

	t.Run("invalid selector", func(t *testing.T) {
		err := p.PreSendValidation(cldf.Environment{}, 0)
		require.Error(t, err)
		require.ErrorContains(t, err, "get selector family")
	})

	t.Run("selector not in env chains", func(t *testing.T) {
		env := cldf.Environment{
			BlockChains: blockChainsWithSelectors(chain_selectors.POLYGON_MAINNET.Selector),
		}
		err := p.PreSendValidation(env, ethSel)
		require.Error(t, err)
		require.ErrorContains(t, err, "not found in chain selectors")
	})

	t.Run("valid selector in env chains", func(t *testing.T) {
		env := cldf.Environment{
			BlockChains: blockChainsWithSelectors(ethSel),
		}
		require.NoError(t, p.PreSendValidation(env, ethSel))
	})
}

func TestEVMPostProposalCCIPSend_SupportedDestinations_Error(t *testing.T) {
	p := &EVMPostProposalCCIPSend{}
	env := cldf.Environment{DataStore: datastore.NewMemoryDataStore().Seal()}
	_, err := p.SupportedDestinations(env, chain_selectors.ETHEREUM_MAINNET.Selector)
	require.Error(t, err)
	require.ErrorContains(t, err, "derive all lane versions from source")
}

func TestEVMPostProposalCCIPSend_AdapterVersionForLane_Error(t *testing.T) {
	p := &EVMPostProposalCCIPSend{}
	env := cldf.Environment{DataStore: datastore.NewMemoryDataStore().Seal()}
	_, err := p.AdapterVersionForLane(
		env,
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.POLYGON_MAINNET.Selector,
	)
	require.Error(t, err)
	require.ErrorContains(t, err, "derive all lane versions from source")
}

func TestEVMPostProposalCCIPSend_SupportedFeeTokens_ErrorWhenOnRampMissing(t *testing.T) {
	p := &EVMPostProposalCCIPSend{}
	ethSel := chain_selectors.ETHEREUM_MAINNET.Selector
	env := cldf.Environment{
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		BlockChains: blockChainsWithSelectors(ethSel),
	}

	_, err := p.SupportedFeeTokens(env, ethSel)
	require.Error(t, err)
	require.ErrorContains(t, err, "failed to find onramp address")
}

func TestGlobalPostProposalCCIPSendHook_EndToEnd_WithEVMProviderInRegistry(t *testing.T) {
	useOnlyEVMPostProposalProvider(t)

	dom := domain.NewDomain(t.TempDir(), "test")
	const envName = "staging"
	writeMinimalEnvDatastore(t, dom, envName)

	registered, ok := cciphooks.GetPostProposalCCIPSendRegistry().Get(chain_selectors.FamilyEVM)
	require.True(t, ok)
	require.IsType(t, &EVMPostProposalCCIPSend{}, registered)

	hook := cciphooks.GlobalPostProposalCCIPSendHook(dom)
	err := hook.Func(context.Background(), cldf_changeset.PostProposalHookParams{
		Env: cldf_changeset.ProposalHookEnv{
			Name:   envName,
			Logger: logger.Test(t),
		},
		Reports: []cldf_changeset.MCMSTimelockExecuteReport{
			{
				Type:   cldf_changeset.MCMSTimelockExecuteReportType,
				Status: "SUCCESS",
				Input: cldf_changeset.MCMSTimelockExecuteReportInput{
					ChainSelector: chain_selectors.ETHEREUM_MAINNET.Selector,
				},
			},
		},
	})
	require.NoError(t, err)
}
