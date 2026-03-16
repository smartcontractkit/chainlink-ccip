package adapters

import (
	"testing"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"
)

func TestConfigImportAdapter_ConnectedChains_cacheHitReturnsCopy(t *testing.T) {
	t.Parallel()
	ci := &ConfigImportAdapter{}
	chainSel := uint64(123)
	cachedChains := []uint64{456, 789}
	ci.connectedChainsCache = map[uint64][]uint64{
		chainSel: cachedChains,
	}

	var e cldf.Environment

	first, err := ci.ConnectedChains(e, chainSel)
	require.NoError(t, err)
	require.Equal(t, []uint64{456, 789}, first)

	// Mutate the returned slice; cache should be unaffected.
	first[0] = 999
	first = append(first, 111)

	second, err := ci.ConnectedChains(e, chainSel)
	require.NoError(t, err)
	require.Equal(t, []uint64{456, 789}, second, "cache should return a copy; mutating first result must not affect second call")
}

func TestConfigImportAdapter_ConnectedChains_cachePerChainSelector(t *testing.T) {
	t.Parallel()
	ci := &ConfigImportAdapter{}
	ci.connectedChainsCache = map[uint64][]uint64{
		100: {200, 201},
		300: {301},
	}

	var e cldf.Environment

	got100, err := ci.ConnectedChains(e, 100)
	require.NoError(t, err)
	require.Equal(t, []uint64{200, 201}, got100)

	got300, err := ci.ConnectedChains(e, 300)
	require.NoError(t, err)
	require.Equal(t, []uint64{301}, got300)
}
