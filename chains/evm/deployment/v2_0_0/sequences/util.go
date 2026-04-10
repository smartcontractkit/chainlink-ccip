package sequences

import (
	"github.com/ethereum/go-ethereum/common"
)

// UnorderedSliceEqual returns true if a and b contain the same elements (order-independent).
// eq reports whether two elements are equal.
func UnorderedSliceEqual[T any](a, b []T, eq func(T, T) bool) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, va := range a {
		found := false
		for j, vb := range b {
			if !used[j] && eq(va, vb) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// SliceNotIn returns elements of superset that are not in subset (set difference).
// eq reports whether two elements are equal.
func SliceNotIn[T any](superset, subset []T, eq func(T, T) bool) []T {
	var out []T
	for _, a := range superset {
		found := false
		for _, b := range subset {
			if eq(a, b) {
				found = true
				break
			}
		}
		if !found {
			out = append(out, a)
		}
	}
	return out
}

// AddressesNotIn returns elements of superset that are not in subset (set difference).
func AddressesNotIn(superset, subset []common.Address) []common.Address {
	return SliceNotIn(superset, subset, func(a, b common.Address) bool { return a == b })
}
