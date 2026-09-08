package tokens

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

func TestResolveRouterRef(t *testing.T) {
	sel := chain_selectors.TEST_90000001.Selector
	routerAddr := "0x1111111111111111111111111111111111111111"
	offRampAddr := "0x2222222222222222222222222222222222222222"

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: sel, Type: "Router", Qualifier: "shared", Address: routerAddr, Version: semver.MustParse("1.0.0"),
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: sel, Type: "OffRamp", Qualifier: "shared", Address: offRampAddr, Version: semver.MustParse("1.0.0"),
	}))
	sealed := ds.Seal()

	t.Run("explicit address bypasses datastore", func(t *testing.T) {
		got, err := resolveRouterRef(sealed, sel, datastore.AddressRef{Address: "0x3333333333333333333333333333333333333333"})
		require.NoError(t, err)
		require.Equal(t, "0x3333333333333333333333333333333333333333", got)
	})

	t.Run("type-less ref defaults to Router", func(t *testing.T) {
		got, err := resolveRouterRef(sealed, sel, datastore.AddressRef{Qualifier: "shared"})
		require.NoError(t, err)
		require.Equal(t, routerAddr, got, "qualifier shared by Router and OffRamp must resolve to the Router")
	})

	t.Run("explicit type is honoured", func(t *testing.T) {
		got, err := resolveRouterRef(sealed, sel, datastore.AddressRef{Type: "OffRamp", Qualifier: "shared"})
		require.NoError(t, err)
		require.Equal(t, offRampAddr, got)
	})

	t.Run("missing ref errors", func(t *testing.T) {
		_, err := resolveRouterRef(sealed, sel, datastore.AddressRef{Qualifier: "nope"})
		require.Error(t, err)
	})
}
