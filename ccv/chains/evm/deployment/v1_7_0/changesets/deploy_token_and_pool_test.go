package changesets_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
)

const dtpChainSelector = uint64(5009297550715157269)

// newDTPEnvironment creates a minimal EVM sim environment for VerifyPreconditions tests.
func newDTPEnvironment(t *testing.T, ds *datastore.MemoryDataStore) deployment.Environment {
	t.Helper()
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{dtpChainSelector}),
	)
	require.NoError(t, err)
	if ds != nil {
		e.DataStore = ds.Seal()
	}
	return *e
}

func rmnProxyRef() datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: dtpChainSelector,
		Type:          datastore.ContractType(rmn_proxy.ContractType),
		Version:       rmn_proxy.Version,
		Address:       common.HexToAddress("0x01").Hex(),
	}
}

func routerRef() datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: dtpChainSelector,
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       common.HexToAddress("0x02").Hex(),
	}
}

func tokenAdminRegistryRef() datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: dtpChainSelector,
		Type:          datastore.ContractType(token_admin_registry.ContractType),
		Version:       token_admin_registry.Version,
		Address:       common.HexToAddress("0x03").Hex(),
	}
}

func dtpValidCfg() cs_core.WithMCMS[changesets.DeployTokenAndPoolCfg] {
	return cs_core.WithMCMS[changesets.DeployTokenAndPoolCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployTokenAndPoolCfg{
			ChainSel:    dtpChainSelector,
			TokenSymbol: "TEST",
			Router:      datastore.AddressRef{Type: datastore.ContractType(router.ContractType), Version: router.Version},
			TokenAdminRegistryRef: datastore.AddressRef{
				Type:    datastore.ContractType(token_admin_registry.ContractType),
				Version: token_admin_registry.Version,
			},
		},
	}
}

func TestDeployTokenAndPool_VerifyPreconditions_MissingRMNProxy(t *testing.T) {
	// Completely empty datastore → RMN proxy lookup fails first.
	env := newDTPEnvironment(t, datastore.NewMemoryDataStore())
	cs := changesets.DeployTokenAndPool(cs_core.GetRegistry())

	err := cs.VerifyPreconditions(env, dtpValidCfg())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "rmn proxy")
}

func TestDeployTokenAndPool_VerifyPreconditions_MissingRouter(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnProxyRef()))
	// Router is absent.

	env := newDTPEnvironment(t, ds)
	cs := changesets.DeployTokenAndPool(cs_core.GetRegistry())

	err := cs.VerifyPreconditions(env, dtpValidCfg())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "router")
}

func TestDeployTokenAndPool_VerifyPreconditions_EmptyTokenAdminRegistryType(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnProxyRef()))
	require.NoError(t, ds.Addresses().Add(routerRef()))

	env := newDTPEnvironment(t, ds)
	cs := changesets.DeployTokenAndPool(cs_core.GetRegistry())

	cfg := dtpValidCfg()
	cfg.Cfg.TokenAdminRegistryRef = datastore.AddressRef{} // empty Type

	err := cs.VerifyPreconditions(env, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "TokenAdminRegistryRef is required")
}

func TestDeployTokenAndPool_VerifyPreconditions_MissingTokenAdminRegistry(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnProxyRef()))
	require.NoError(t, ds.Addresses().Add(routerRef()))
	// Token admin registry absent despite Type being set.

	env := newDTPEnvironment(t, ds)
	cs := changesets.DeployTokenAndPool(cs_core.GetRegistry())

	err := cs.VerifyPreconditions(env, dtpValidCfg())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "token admin registry")
}

func TestDeployTokenAndPool_VerifyPreconditions_AllRefsPresent(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnProxyRef()))
	require.NoError(t, ds.Addresses().Add(routerRef()))
	require.NoError(t, ds.Addresses().Add(tokenAdminRegistryRef()))

	env := newDTPEnvironment(t, ds)
	cs := changesets.DeployTokenAndPool(cs_core.GetRegistry())

	err := cs.VerifyPreconditions(env, dtpValidCfg())
	require.NoError(t, err)
}
