package changesets_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/latest/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestDeployRMN_VerifyPreconditions(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	require.True(t, ok)

	mcmsRegistry := cs_core.GetRegistry()
	err = changesets.DeployRMN(mcmsRegistry).VerifyPreconditions(*e, cs_core.WithMCMS[changesets.DeployRMNCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployRMNCfg{
			ChainSel:    chainSelector,
			CurseAdmins: []common.Address{chain.DeployerKey.From},
		},
	})
	require.NoError(t, err)
}

func TestDeployRMN_Apply(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	require.True(t, ok)

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.DeployRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployRMNCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployRMNCfg{
			ChainSel:    chainSelector,
			CurseAdmins: []common.Address{chain.DeployerKey.From},
		},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, addrs, 1)
	require.Equal(t, datastore.ContractType(rmn.ContractType), addrs[0].Type)
	require.Equal(t, rmn.Version.String(), addrs[0].Version.String())
}
