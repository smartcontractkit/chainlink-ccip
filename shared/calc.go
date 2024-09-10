package shared

import (
	"math/big"
	"sort"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Median returns the middle element after sorting the provided slice.
// For an empty slice, it returns the zero value of the type.
// If the length of the slice is even, it returns the right-hand-side value of the middle 2 elements.
func Median[T any](vals []T, less func(T, T) bool) T {
	if len(vals) == 0 {
		var zero T
		return zero
	}
	valsCopy := make([]T, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return less(valsCopy[i], valsCopy[j])
	})
	return valsCopy[len(valsCopy)/2]
}

// MedianTimestamp returns the middle timestamp after sorting the provided timestamps.
func MedianTimestamp(timestamps []time.Time) time.Time {
	return Median(timestamps, func(a, b time.Time) bool {
		return a.Before(b)
	})
}

// MedianBigInt returns the middle number after sorting the provided numbers.
func MedianBigInt(vals []cciptypes.BigInt) cciptypes.BigInt {
	return Median(vals, func(a, b cciptypes.BigInt) bool {
		return a.Cmp(b.Int) == -1
	})
}

// Deviates checks if x1 and x2 deviates based on the provided ppb (parts per billion)
// ppb is calculated based on the smaller value of the two
// e.g, if x1 > x2, deviation_parts_per_billion = ((x1 - x2) / x2) * 1e9
func Deviates(x1, x2 *big.Int, ppb int64) bool {
	// if x1 == 0 or x2 == 0, deviates if x2 != x1, to avoid the relative division by 0 error
	if x1.BitLen() == 0 || x2.BitLen() == 0 {
		return x1.Cmp(x2) != 0
	}
	diff := big.NewInt(0).Sub(x1, x2) // diff = x1-x2
	diff.Mul(diff, big.NewInt(1e9))   // diff = diff * 1e9
	// dividing by the smaller value gives consistent ppb regardless of input order, and supports >100% deviation.
	if x1.Cmp(x2) > 0 {
		diff.Div(diff, x2)
	} else {
		diff.Div(diff, x1)
	}
	return diff.CmpAbs(big.NewInt(ppb)) > 0 // abs(diff) > ppb
}
