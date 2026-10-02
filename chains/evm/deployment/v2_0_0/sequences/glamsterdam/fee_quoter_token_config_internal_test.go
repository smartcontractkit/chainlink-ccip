package glamsterdam

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCollectPaged(t *testing.T) {
	addrs := make([]common.Address, 250)
	for i := range addrs {
		addrs[i] = common.BigToAddress(new(big.Int).SetUint64(uint64(i + 1)))
	}
	fetchFrom := func(all []common.Address, calls *int) func(start, maxCount uint64) ([]common.Address, error) {
		return func(start, maxCount uint64) ([]common.Address, error) {
			*calls++
			if start >= uint64(len(all)) {
				return nil, nil
			}
			end := min(start+maxCount, uint64(len(all)))
			return all[start:end], nil
		}
	}

	t.Run("drains multiple pages and stops on the short one", func(t *testing.T) {
		calls := 0
		got, err := collectPaged(100, fetchFrom(addrs, &calls))
		require.NoError(t, err)
		require.Equal(t, addrs, got)
		require.Equal(t, 3, calls) // 100 + 100 + 50
	})

	t.Run("an exact multiple of the page size needs one trailing empty page", func(t *testing.T) {
		calls := 0
		got, err := collectPaged(100, fetchFrom(addrs[:200], &calls))
		require.NoError(t, err)
		require.Len(t, got, 200)
		require.Equal(t, 3, calls)
	})

	t.Run("empty registry", func(t *testing.T) {
		calls := 0
		got, err := collectPaged(100, fetchFrom(nil, &calls))
		require.NoError(t, err)
		require.Empty(t, got)
		require.Equal(t, 1, calls)
	})

	t.Run("propagates errors", func(t *testing.T) {
		_, err := collectPaged(100, func(uint64, uint64) ([]common.Address, error) { return nil, errors.New("boom") })
		require.ErrorContains(t, err, "boom")
	})
}
