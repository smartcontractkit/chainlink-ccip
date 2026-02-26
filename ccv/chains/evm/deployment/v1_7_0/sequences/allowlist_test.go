package sequences

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestMakeAllowlistUpdates(t *testing.T) {
	addr := func(s string) common.Address { return common.HexToAddress(s) }

	t.Run("no current no added no removed", func(t *testing.T) {
		toAdd, toRemove := makeAllowlistUpdates(nil, nil, nil)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})

	t.Run("current only unchanged", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		toAdd, toRemove := makeAllowlistUpdates(current, nil, nil)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})

	t.Run("add one address", func(t *testing.T) {
		current := []common.Address{addr("0x01")}
		added := []string{"0x02"}
		toAdd, toRemove := makeAllowlistUpdates(current, added, nil)
		require.Len(t, toAdd, 1)
		require.Equal(t, addr("0x02"), toAdd[0])
		require.Empty(t, toRemove)
	})

	t.Run("remove one address", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		removed := []string{"0x02"}
		toAdd, toRemove := makeAllowlistUpdates(current, nil, removed)
		require.Empty(t, toAdd)
		require.Len(t, toRemove, 1)
		require.Equal(t, addr("0x02"), toRemove[0])
	})

	t.Run("desired equals current union added minus removed", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		added := []string{"0x03"}
		removed := []string{"0x02"}
		toAdd, toRemove := makeAllowlistUpdates(current, added, removed)
		// desired = {0x01, 0x03}; toAdd = desired \ current = {0x03}, toRemove = current \ desired = {0x02}
		require.Len(t, toAdd, 1)
		require.Equal(t, addr("0x03"), toAdd[0])
		require.Len(t, toRemove, 1)
		require.Equal(t, addr("0x02"), toRemove[0])
	})

	t.Run("idempotent when already desired", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		// No adds or removes: desired = current
		toAdd, toRemove := makeAllowlistUpdates(current, nil, nil)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})

	t.Run("add already in current no op", func(t *testing.T) {
		current := []common.Address{addr("0x01")}
		added := []string{"0x01"}
		toAdd, toRemove := makeAllowlistUpdates(current, added, nil)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})
}
