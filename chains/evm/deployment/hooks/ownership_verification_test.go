package hooks

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	seq2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func TestTimelocksInOwnershipCheck_LoadsAndCaches(t *testing.T) {
	selector := chainsel.ETHEREUM_MAINNET.Selector
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType(common_utils.RBACTimelock),
		Qualifier:     common_utils.CLLQualifier,
		Address:       "0x00000000000000000000000000000000000000A1",
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType(common_utils.RBACTimelock),
		Qualifier:     common_utils.RMNTimelockQualifier,
		Address:       "0x00000000000000000000000000000000000000B1",
	}))

	e := &EVMContractOwnership{}
	require.NoError(t, e.timelocksInOwnershipCheck(ds.Seal(), selector))

	cll, ok := e.cllccipTimelockAddr.Load(selector)
	require.True(t, ok)
	require.Equal(t, common.HexToAddress("0x00000000000000000000000000000000000000A1"), cll.(common.Address))
	rmn, ok := e.rmntimelockAddr.Load(selector)
	require.True(t, ok)
	require.Equal(t, common.HexToAddress("0x00000000000000000000000000000000000000B1"), rmn.(common.Address))

	// cache hit should not need datastore lookups again.
	require.NoError(t, e.timelocksInOwnershipCheck(datastore.NewMemoryDataStore().Seal(), selector))
}

func TestTimelocksInOwnershipCheck_MissingCLLTimelock(t *testing.T) {
	selector := chainsel.ETHEREUM_MAINNET.Selector
	ds := datastore.NewMemoryDataStore()
	e := &EVMContractOwnership{}
	err := e.timelocksInOwnershipCheck(ds.Seal(), selector)
	require.Error(t, err)
	require.ErrorContains(t, err, "ownership transfer requires CLLCCIP RBACTimelock")
}

func TestExpectedOwnerForRef_UsesRMNTimelockForRMNRemote(t *testing.T) {
	selector := chainsel.ETHEREUM_MAINNET.Selector
	e := &EVMContractOwnership{}
	e.cllccipTimelockAddr.Store(selector, common.HexToAddress("0x00000000000000000000000000000000000000A1"))
	e.rmntimelockAddr.Store(selector, common.HexToAddress("0x00000000000000000000000000000000000000B1"))

	normal, err := e.expectedOwnerForRef(datastore.AddressRef{
		ChainSelector: selector,
		Type:          "AnyType",
	})
	require.NoError(t, err)
	rmn, err := e.expectedOwnerForRef(datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType(rmn_remote.ContractType),
	})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress("0x00000000000000000000000000000000000000A1"), normal)
	require.Equal(t, common.HexToAddress("0x00000000000000000000000000000000000000B1"), rmn)
}

func TestExpectedOwnerForRef_MissingTimelockReturnsError(t *testing.T) {
	selector := chainsel.ETHEREUM_MAINNET.Selector
	e := &EVMContractOwnership{}

	_, err := e.expectedOwnerForRef(datastore.AddressRef{
		ChainSelector: selector,
		Type:          "AnyType",
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "CLLCCIP RBACTimelock address not found")

	_, err = e.expectedOwnerForRef(datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType(rmn_remote.ContractType),
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "RMNMCMS RBACTimelock address not found")
}

func TestNeedsOwnershipCheck_UsesLaneMigratorContractTypes(t *testing.T) {
	e := &EVMContractOwnership{}
	allowedTypes := []datastore.ContractType{
		datastore.ContractType(committee_verifier.ContractType),
		datastore.ContractType(executor.ContractType),
		datastore.ContractType(seq2_0.ExecutorProxyType),
		datastore.ContractType(onrampops.ContractType),
		datastore.ContractType(offrampops.ContractType),
		datastore.ContractType(fqops.ContractType),
		datastore.ContractType(routerops.ContractType),
		datastore.ContractType(rmn_remote.ContractType),
		datastore.ContractType(rmn_proxy.ContractType),
		datastore.ContractType(token_admin_registry.ContractType),
	}

	for _, ct := range allowedTypes {
		require.True(t, e.NeedsOwnershipCheck(datastore.AddressRef{Type: ct}), "expected allowed type %s to require ownership check", ct)
	}
	require.False(t, e.NeedsOwnershipCheck(datastore.AddressRef{Type: "UnknownType"}))
}

func TestVerifyContractOwnership_NoRPCConfigured(t *testing.T) {
	e := &EVMContractOwnership{}
	err := e.VerifyContractOwnership(t.Context(), logger.Test(t), datastore.NewMemoryDataStore().Seal(), cfgnet.Network{
		ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
		RPCs:          nil,
	}, nil)
	require.Error(t, err)
	require.ErrorContains(t, err, "has no HTTP RPC configured")
}

func TestVerifyContractOwnership_InvalidPreferredURLScheme(t *testing.T) {
	e := &EVMContractOwnership{}
	err := e.VerifyContractOwnership(t.Context(), logger.Test(t), datastore.NewMemoryDataStore().Seal(), cfgnet.Network{
		ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
		RPCs: []cfgnet.RPC{{
			RPCName:            "bad-rpc",
			HTTPURL:            "http://localhost:8545",
			PreferredURLScheme: "definitely-invalid",
		}},
	}, nil)
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid preferred URL scheme")
}

func TestVerifyContractOwnership_DialRPCFailure(t *testing.T) {
	e := &EVMContractOwnership{}
	err := e.VerifyContractOwnership(t.Context(), logger.Test(t), datastore.NewMemoryDataStore().Seal(), cfgnet.Network{
		ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
		RPCs: []cfgnet.RPC{{
			RPCName:            "local",
			HTTPURL:            "http://localhost:8545",
			PreferredURLScheme: "http",
		}},
	}, nil)
	require.Error(t, err)
	require.ErrorContains(t, err, "dial RPC for chain")
}
