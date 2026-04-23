package sequences

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestMakeAllowlistUpdates(t *testing.T) {
	addr := func(s string) common.Address { return common.HexToAddress(s) }

	t.Run("no current no added no removed", func(t *testing.T) {
		toAdd, toRemove, err := makeAllowlistUpdates(nil, nil, nil)
		require.NoError(t, err)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})

	t.Run("current only unchanged", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		toAdd, toRemove, err := makeAllowlistUpdates(current, nil, nil)
		require.NoError(t, err)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})

	t.Run("add one address", func(t *testing.T) {
		current := []common.Address{addr("0x01")}
		added := []string{"0x0000000000000000000000000000000000000002"}
		toAdd, toRemove, err := makeAllowlistUpdates(current, added, nil)
		require.NoError(t, err)
		require.Len(t, toAdd, 1)
		require.Equal(t, addr("0x02"), toAdd[0])
		require.Empty(t, toRemove)
	})

	t.Run("remove one address", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		removed := []string{"0x0000000000000000000000000000000000000002"}
		toAdd, toRemove, err := makeAllowlistUpdates(current, nil, removed)
		require.NoError(t, err)
		require.Empty(t, toAdd)
		require.Len(t, toRemove, 1)
		require.Equal(t, addr("0x02"), toRemove[0])
	})

	t.Run("desired equals current union added minus removed", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		added := []string{"0x0000000000000000000000000000000000000003"}
		removed := []string{"0x0000000000000000000000000000000000000002"}
		toAdd, toRemove, err := makeAllowlistUpdates(current, added, removed)
		require.NoError(t, err)
		require.Len(t, toAdd, 1)
		require.Equal(t, addr("0x03"), toAdd[0])
		require.Len(t, toRemove, 1)
		require.Equal(t, addr("0x02"), toRemove[0])
	})

	t.Run("idempotent when already desired", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		toAdd, toRemove, err := makeAllowlistUpdates(current, nil, nil)
		require.NoError(t, err)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})

	t.Run("add already in current no op", func(t *testing.T) {
		current := []common.Address{addr("0x01")}
		added := []string{"0x0000000000000000000000000000000000000001"}
		toAdd, toRemove, err := makeAllowlistUpdates(current, added, nil)
		require.NoError(t, err)
		require.Empty(t, toAdd)
		require.Empty(t, toRemove)
	})

	t.Run("empty added and empty removed slices fallback to current", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		toAdd, toRemove, err := makeAllowlistUpdates(current, []string{}, []string{})
		require.NoError(t, err)
		require.Empty(t, toAdd, "empty added should not produce toAdd")
		require.Empty(t, toRemove, "empty removed should not produce toRemove")
	})

	t.Run("empty added with non-empty removed", func(t *testing.T) {
		current := []common.Address{addr("0x01"), addr("0x02")}
		toAdd, toRemove, err := makeAllowlistUpdates(current, []string{}, []string{"0x0000000000000000000000000000000000000002"})
		require.NoError(t, err)
		require.Empty(t, toAdd)
		require.Len(t, toRemove, 1)
		require.Equal(t, addr("0x02"), toRemove[0])
	})

	t.Run("empty removed with non-empty added", func(t *testing.T) {
		current := []common.Address{addr("0x01")}
		toAdd, toRemove, err := makeAllowlistUpdates(current, []string{"0x0000000000000000000000000000000000000002"}, []string{})
		require.NoError(t, err)
		require.Len(t, toAdd, 1)
		require.Equal(t, addr("0x02"), toAdd[0])
		require.Empty(t, toRemove)
	})

	t.Run("invalid hex in added returns error", func(t *testing.T) {
		_, _, err := makeAllowlistUpdates(nil, []string{"not-hex"}, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid hex address")
	})

	t.Run("invalid hex in removed returns error", func(t *testing.T) {
		_, _, err := makeAllowlistUpdates(nil, nil, []string{"xyz"})
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid hex address")
	})
}
