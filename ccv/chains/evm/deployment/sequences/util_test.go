package sequences

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnorderedSliceEqual(t *testing.T) {
	eqBytes := func(a, b byte) bool { return a == b }
	eqAddr := func(a, b common.Address) bool { return a == b }

	tests := []struct {
		name string
		a    []byte
		b    []byte
		eq   func(byte, byte) bool
		want bool
	}{
		{"both nil", nil, nil, eqBytes, true},
		{"both empty", []byte{}, []byte{}, eqBytes, true},
		{"same order", []byte{1, 2, 3}, []byte{1, 2, 3}, eqBytes, true},
		{"different order", []byte{1, 2, 3}, []byte{3, 1, 2}, eqBytes, true},
		{"different lengths", []byte{1, 2}, []byte{1, 2, 3}, eqBytes, false},
		{"same elements duplicate", []byte{1, 1, 2}, []byte{1, 2, 1}, eqBytes, true},
		{"mismatch", []byte{1, 2, 3}, []byte{1, 2, 4}, eqBytes, false},
		{"subset same length", []byte{1, 2, 2}, []byte{1, 1, 2}, eqBytes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnorderedSliceEqual(tt.a, tt.b, tt.eq)
			assert.Equal(t, tt.want, got, "UnorderedSliceEqual(%v, %v)", tt.a, tt.b)
		})
	}

	// Address slice (order-independent)
	addr1 := common.HexToAddress("0x01")
	addr2 := common.HexToAddress("0x02")
	addr3 := common.HexToAddress("0x03")
	require.True(t, UnorderedSliceEqual(
		[]common.Address{addr1, addr2, addr3},
		[]common.Address{addr3, addr1, addr2},
		eqAddr,
	))
	require.False(t, UnorderedSliceEqual(
		[]common.Address{addr1, addr2},
		[]common.Address{addr1, addr3},
		eqAddr,
	))
}

func TestSliceNotIn(t *testing.T) {
	eqInt := func(a, b int) bool { return a == b }

	tests := []struct {
		name     string
		superset []int
		subset   []int
		eq       func(int, int) bool
		want     []int
	}{
		{"empty superset", []int{}, []int{1, 2}, eqInt, nil},
		{"empty subset", []int{1, 2, 3}, []int{}, eqInt, []int{1, 2, 3}},
		{"subset equals superset", []int{1, 2, 3}, []int{1, 2, 3}, eqInt, nil},
		{"proper subset", []int{1, 2, 3, 4}, []int{1, 3}, eqInt, []int{2, 4}},
		{"no overlap", []int{1, 2}, []int{3, 4}, eqInt, []int{1, 2}},
		{"duplicates in superset", []int{1, 1, 2}, []int{1}, eqInt, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceNotIn(tt.superset, tt.subset, tt.eq)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAddressesNotIn(t *testing.T) {
	a := common.HexToAddress("0x01")
	b := common.HexToAddress("0x02")
	c := common.HexToAddress("0x03")

	got := AddressesNotIn(
		[]common.Address{a, b, c},
		[]common.Address{a, c},
	)
	require.Len(t, got, 1)
	require.Equal(t, b, got[0])

	got = AddressesNotIn([]common.Address{a, b}, []common.Address{a, b})
	require.Empty(t, got)

	got = AddressesNotIn([]common.Address{}, []common.Address{a})
	require.Empty(t, got)

	got = AddressesNotIn([]common.Address{a, b}, []common.Address{})
	require.Len(t, got, 2)
	require.Contains(t, got, a)
	require.Contains(t, got, b)
}

func TestUnorderedSliceEqual_Bytes(t *testing.T) {
	// Used for OnRamps comparison (bytes.Equal semantics)
	onRampsA := [][]byte{common.LeftPadBytes([]byte{1}, 32), common.LeftPadBytes([]byte{2}, 32)}
	onRampsB := [][]byte{common.LeftPadBytes([]byte{2}, 32), common.LeftPadBytes([]byte{1}, 32)}
	require.True(t, UnorderedSliceEqual(onRampsA, onRampsB, bytes.Equal))
	require.False(t, UnorderedSliceEqual(onRampsA, [][]byte{common.LeftPadBytes([]byte{1}, 32)}, bytes.Equal))
}
