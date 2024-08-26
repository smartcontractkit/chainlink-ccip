package utils

import (
	"math/big"
	"sort"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// MedianTimestamp returns the middle timestamp after sorting the provided timestamps.
// empty/default is returned if the provided slice is empty.
func MedianTimestamp(timestamps []time.Time) time.Time {
	if len(timestamps) == 0 {
		return time.Time{}
	}
	valsCopy := make([]time.Time, len(timestamps))
	copy(valsCopy[:], timestamps[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Before(valsCopy[j])
	})

	return valsCopy[len(valsCopy)/2]
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

// MedianBigInt returns the middle number after sorting the provided numbers.
// nil is returned if the provided slice is empty.
// If length of the provided slice is even, the right-hand-side value of the middle 2 numbers is returned.
// The objective of this function is to always pick within the range of values reported
// by honest nodes when we have 2f+1 values.
func MedianBigInt(vals []cciptypes.BigInt) cciptypes.BigInt {
	if len(vals) == 0 {
		return cciptypes.BigInt{}
	}

	valsCopy := make([]cciptypes.BigInt, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j].Int) == -1
	})
	return valsCopy[len(valsCopy)/2]
}

// Given a list of elems, return the elem that occurs most frequently and how often it occurs
func MostFrequentElem[T comparable](elems []T) (T, int) {
	var mostFrequentElem T

	counts := Counts(elems)
	maxCount := 0

	for elem, count := range counts {
		if count > maxCount {
			mostFrequentElem = elem
			maxCount = count
		}
	}

	return mostFrequentElem, maxCount
}

// Given a list of elems, return a map from elems to how often they occur in the given list
func Counts[T comparable](elems []T) map[T]int {
	m := make(map[T]int)
	for _, elem := range elems {
		m[elem]++
	}

	return m
}
