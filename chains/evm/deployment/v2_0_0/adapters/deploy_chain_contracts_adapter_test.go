package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
)

func newResolveTestEnv(sel uint64, ds datastore.DataStore) deployment.Environment {
	return deployment.Environment{
		BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{
			sel: cldfevm.Chain{Selector: sel},
		}),
		DataStore: ds,
	}
}

func TestEVMDeployChainContractsAdapter_ResolveDeployAddresses_FailsWithoutCREATE2Factory(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	env := newResolveTestEnv(sel, datastore.NewMemoryDataStore().Seal())

	adapter := &adapters.EVMDeployChainContractsAdapter{}
	_, err := adapter.ResolveDeployAddresses(env, sel)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no CREATE2Factory")
}

func TestEVMDeployChainContractsAdapter_ResolveDeployAddresses_ReturnsCREATE2FactoryFromDatastore(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	factory := "0x490fD85cD094E0A71EC440AA6C1e8eC2882aeA22"

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Address:       factory,
		ChainSelector: sel,
		Type:          datastore.ContractType(create2_factory.ContractType),
		Version:       create2_factory.Version,
	}))
	env := newResolveTestEnv(sel, ds.Seal())

	adapter := &adapters.EVMDeployChainContractsAdapter{}
	resolved, err := adapter.ResolveDeployAddresses(env, sel)
	require.NoError(t, err)
	assert.Equal(t, factory, resolved.DeployerContract)
	assert.Empty(t, resolved.NewAddressRefs)
}
