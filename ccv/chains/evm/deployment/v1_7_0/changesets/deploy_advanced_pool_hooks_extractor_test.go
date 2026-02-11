package changesets_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks_extractor"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestDeployAdvancedPoolHooksExtractor_VerifyPreconditions(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	mcmsRegistry := cs_core.GetRegistry()
	err = changesets.DeployAdvancedPoolHooksExtractor(mcmsRegistry).VerifyPreconditions(*e, cs_core.WithMCMS[changesets.DeployAdvancedPoolHooksExtractorCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployAdvancedPoolHooksExtractorCfg{
			ChainSel: chainSelector,
		},
	})
	require.NoError(t, err)
}

func TestDeployAdvancedPoolHooksExtractor_Apply(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.DeployAdvancedPoolHooksExtractor(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployAdvancedPoolHooksExtractorCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployAdvancedPoolHooksExtractorCfg{
			ChainSel: chainSelector,
		},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, addrs, 1)
	require.Equal(t, datastore.ContractType(advanced_pool_hooks_extractor.ContractType), addrs[0].Type)
	require.Equal(t, advanced_pool_hooks_extractor.Version.String(), addrs[0].Version.String())
}
