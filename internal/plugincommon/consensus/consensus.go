package consensus

import (
	"cmp"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// GetConsensusMap takes a mapping from chains to a list of items,
// return a mapping from chains to a single consensus item.
// The consensus item for a given chain is the item with the
// most observations that was observed at least fChain times.
func GetConsensusMap[K comparable, T any](
	lggr logger.Logger,
	objectName string,
	itemsByKey map[K][]T,
	minObs MultiThreshold[K],
) map[K]T {
	consensus := make(map[K]T)

	for key, items := range itemsByKey {
		if minThresh, exists := minObs.Get(key); exists {
			minObservations := NewMinObservation[T](minThresh, nil)
			for _, item := range items {
				minObservations.Add(item)
			}
			items = minObservations.GetValid()
			if len(items) != 1 {
				// TODO: metrics
				lggr.Debugw("could not reach consensus due to not enough observations meeting the minimum threshold",
					"objectName", objectName,
					"key", key,
					"minThreshold", minThresh,
					"items", items)
			} else {
				consensus[key] = items[0]
			}
		} else {
			// TODO: metrics
			lggr.Warnw("min threshold not found defined", "objectName", objectName, "key", key)
		}
	}
	return consensus
}

// Aggregator is a function type that aggregates a slice of values into a single value.
type Aggregator[T any] func(vals []T) T

func GetConsensusMapAggregator[K comparable, T any](
	lggr logger.Logger,
	objectName string,
	items map[K][]T,
	f MultiThreshold[K],
	agg Aggregator[T],
) map[K]T {
	consensus := make(map[K]T)

	for key, values := range items {
		if thresh, ok := f.Get(key); !ok || len(values) < int(thresh) {
			lggr.Debugw("could not reach consensus in consensusMapAggregator",
				"objectName", objectName,
				"key", key)
			continue
		}
		consensus[key] = agg(values)
	}
	return consensus
}

// Median returns the middle element after sorting the provided slice.
// For an empty slice, it returns the zero value of the type.
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

// GetOrderedConsensus returns the max sequence number for each chain that meets the min obs threshold
// taking into account thresholds which is the maximum number of faults across the whole DON.
// Here we'll either have thresholds+1 parsed honest values here,
// 2f+1 parsed values with thresholds adversarial values or somewhere in between.
// We choose the more "conservative" seqNum[thresholds] so:
// - We are ensured that at least one honest oracle has seen the max, so adversary cannot set it lower and
// cause the maxSeqNum < minSeqNum errors
// - If an honest oracle reports sorted_max[thresholds] which happens to be stale i.e. that oracle
// has a delayed view of the source chain, then we simply lose a little bit of throughput.
// - If we were to pick sorted_max[-thresholds]
// i.e. the maximum honest node view (a more "aggressive" setting in terms of throughput),
// then an adversary can continually send high values e.g. imagine we have observations from all 4 nodes
// [honest 1, honest 1, honest 2, malicious 2], in this case we pick 2, but it's not enough to be able
// to build a report since the first 2 honest nodes are unaware of message 2.
// In an observation of [1, 2, 3, 4], we would pick 1, which is the most conservative choice.
func GetOrderedConsensus[K comparable, T cmp.Ordered](
	lggr logger.Logger,
	objectName string,
	itemsByKey map[K][]T,
	minObs MultiThreshold[K]) map[K]T {
	result := make(map[K]T)

	for key, items := range itemsByKey {
		if _, exists := minObs.Get(key); !exists {
			lggr.Warnw("could not find threshold value", "objectName", objectName, "key", key)
			continue
		}

		minThresh, _ := minObs.Get(key)
		if minThresh <= 0 {
			lggr.Errorw("found a negative or 0 threshold",
				"objectName", objectName,
				"key", key,
				"minThresh", minThresh)
			continue
		}

		if len(items) < 2*int(minThresh)+1 {
			lggr.Errorw("insufficient items to reach consensus",
				"objectName", objectName,
				"key", key,
				"minThresh", minThresh,
				"items", items)
			continue
		}

		sort.Slice(items, func(i, j int) bool { return items[i] < items[j] })
		result[key] = items[minThresh]
	}
	return result
}

func TimestampComparator(a, b time.Time) bool {
	return a.Before(b)
}

func TimestampsMedian(timestamps []time.Time) time.Time {
	return Median(timestamps, TimestampComparator)
}

func BigIntComparator(a, b cciptypes.BigInt) bool {
	return a.Cmp(b.Int) == -1
}

func TokenPriceComparator(a, b cciptypes.TokenPrice) bool {
	return a.Price.Int.Cmp(b.Price.Int) == -1
}

// TimestampedBigAggregator aggregates the fee quoter updates by taking the median of the prices and timestamps
func TimestampedBigAggregator(updates []cciptypes.TimestampedBig) cciptypes.TimestampedBig {
	timestamps := make([]time.Time, len(updates))
	prices := make([]cciptypes.BigInt, len(updates))
	for i := range updates {
		timestamps[i] = updates[i].Timestamp
		prices[i] = updates[i].Value
	}
	medianPrice := Median(prices, BigIntComparator)
	medianTimestamp := Median(timestamps, TimestampComparator)
	return cciptypes.TimestampedBig{
		Value:     medianPrice,
		Timestamp: medianTimestamp,
	}
}
