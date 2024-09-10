package shared

import (
	"math/big"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

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

var TimestampComparator = func(a, b time.Time) bool {
	return a.Before(b)
}

var BigIntComparator = func(a, b cciptypes.BigInt) bool {
	return a.Cmp(b.Int) == -1
}

// MedianTimestampedBig returns median of the provided TimestampedBig values.
// It calculates the median of the timestamps and the median of the values.
func MedianTimestampedBig(vals []TimestampedBig) TimestampedBig {
	timestamps := make([]time.Time, len(vals))
	prices := make([]cciptypes.BigInt, len(vals))
	for i := range vals {
		timestamps[i] = vals[i].Timestamp
		prices[i] = vals[i].Value
	}
	return TimestampedBig{
		Timestamp: MedianTimestamp(timestamps),
		Value:     MedianBigInt(prices),
	}
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
