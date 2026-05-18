package adapters

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func TestResolveRouterAddress(t *testing.T) {
	const chainSelector uint64 = 5009297550715157269
	prodRouter := common.HexToAddress("0x1111111111111111111111111111111111111111")
	testRouter := common.HexToAddress("0x2222222222222222222222222222222222222222")
	override := common.HexToAddress("0x3333333333333333333333333333333333333333")

	newStore := func(t *testing.T) datastore.DataStore {
		t.Helper()
		ds := datastore.NewMemoryDataStore()
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(router.ContractType),
			Version:       router.Version,
			Address:       prodRouter.Hex(),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(router.TestRouterContractType),
			Version:       router.Version,
			Address:       testRouter.Hex(),
		}))
		return ds.Seal()
	}

	t.Run("nil ref defaults to production router", func(t *testing.T) {
		got, err := resolveRouterAddress(newStore(t), chainSelector, nil)
		require.NoError(t, err)
		require.Equal(t, prodRouter, got)
	})

	t.Run("ref with TestRouter type resolves to test router", func(t *testing.T) {
		got, err := resolveRouterAddress(newStore(t), chainSelector, &datastore.AddressRef{
			Type: datastore.ContractType(router.TestRouterContractType),
		})
		require.NoError(t, err)
		require.Equal(t, testRouter, got)
	})

	t.Run("ref with explicit Address bypasses datastore", func(t *testing.T) {
		// Empty datastore would normally cause a lookup failure; an explicit
		// Address must skip the lookup entirely.
		got, err := resolveRouterAddress(datastore.NewMemoryDataStore().Seal(), chainSelector, &datastore.AddressRef{
			Address: override.Hex(),
		})
		require.NoError(t, err)
		require.Equal(t, override, got)
	})

	t.Run("ref with non-hex Address errors", func(t *testing.T) {
		// "0x123" is not 20 bytes; common.HexToAddress would silently pad it,
		// so resolveRouterAddress must reject it up-front.
		_, err := resolveRouterAddress(newStore(t), chainSelector, &datastore.AddressRef{
			Address: "0x123",
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "not a hex address")
	})

	t.Run("ref with zero Address errors", func(t *testing.T) {
		_, err := resolveRouterAddress(newStore(t), chainSelector, &datastore.AddressRef{
			Address: "0x0000000000000000000000000000000000000000",
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "zero address")
	})

	t.Run("ref forces chain selector to target chain", func(t *testing.T) {
		// User passes a ref with the wrong chain selector — the helper must
		// rewrite it to the target chain so the lookup succeeds.
		got, err := resolveRouterAddress(newStore(t), chainSelector, &datastore.AddressRef{
			ChainSelector: chainSelector + 1,
			Type:          datastore.ContractType(router.TestRouterContractType),
		})
		require.NoError(t, err)
		require.Equal(t, testRouter, got)
	})

	t.Run("missing router type in datastore returns error", func(t *testing.T) {
		emptyDS := datastore.NewMemoryDataStore().Seal()
		_, err := resolveRouterAddress(emptyDS, chainSelector, nil)
		require.Error(t, err)
	})

	t.Run("ref with only Qualifier resolves via lookup", func(t *testing.T) {
		ds := datastore.NewMemoryDataStore()
		// Two routers with the same Type but different qualifiers — a qualifier
		// on the ref disambiguates.
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(router.ContractType),
			Version:       router.Version,
			Qualifier:     "canary",
			Address:       override.Hex(),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(router.ContractType),
			Version:       router.Version,
			Address:       prodRouter.Hex(),
		}))
		got, err := resolveRouterAddress(ds.Seal(), chainSelector, &datastore.AddressRef{
			Qualifier: "canary",
		})
		require.NoError(t, err)
		require.Equal(t, override, got)
	})
}
