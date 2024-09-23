package plugincommon

import (
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// GetConsensusMap takes a mapping from chains to a list of items,
// return a mapping from chains to a single consensus item.
// The consensus item for a given chain is the item with the
// most observations that was observed at least fChain times.
func GetConsensusMap[T any](
	lggr logger.Logger,
	objectName string,
	itemsByChain map[cciptypes.ChainSelector][]T,
	minObs map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]T {
	consensus := make(map[cciptypes.ChainSelector]T)

	for chain, items := range itemsByChain {
		if min, exists := minObs[chain]; exists {
			minObservations := NewMinObservation[T](min, nil)
			for _, item := range items {
				minObservations.Add(item)
			}
			items = minObservations.GetValid()
			if len(items) != 1 {
				// TODO: metrics
				lggr.Warnf("failed to reach consensus on a %s's for chain %d "+
					"because no single item was observed more than the expected min (%d) times, "+
					"all observed items: %v",
					objectName, chain, min, items)
			} else {
				consensus[chain] = items[0]
			}
		} else {
			// TODO: metrics
			lggr.Warnf("getConsensus(%s): min not found for chain %d", objectName, chain)
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
	minObs map[K]int,
	agg Aggregator[T],
) map[K]T {
	consensus := make(map[K]T)

	for key, values := range items {
		if _, exists := minObs[key]; !exists {
			lggr.Warnf("no F value found for key %d", key)
			continue
		}
		if len(values) < minObs[key] {
			lggr.Warnf("not enough observations to reach consensus for %s on key %v", objectName, key)
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

var TimestampComparator = func(a, b time.Time) bool {
	return a.Before(b)
}

var BigIntComparator = func(a, b cciptypes.BigInt) bool {
	return a.Cmp(b.Int) == -1
}

var TokenPriceComparator = func(a, b cciptypes.TokenPrice) bool {
	return a.Price.Int.Cmp(b.Price.Int) == -1
}

// TimestampedBigAggregator aggregates the fee quoter updates by taking the median of the prices and timestamps
var TimestampedBigAggregator = func(updates []plugintypes.TimestampedBig) plugintypes.TimestampedBig {
	timestamps := make([]time.Time, len(updates))
	prices := make([]cciptypes.BigInt, len(updates))
	for i := range updates {
		timestamps[i] = updates[i].Timestamp
		prices[i] = updates[i].Value
	}
	medianPrice := Median(prices, BigIntComparator)
	medianTimestamp := Median(timestamps, TimestampComparator)
	return plugintypes.TimestampedBig{
		Value:     medianPrice,
		Timestamp: medianTimestamp,
	}
}

func EarliestTimestamp(updates []plugintypes.TimestampedBig, current time.Time) time.Time {
	earliest := current
	for _, update := range updates {
		if update.Timestamp.Before(earliest) {
			earliest = update.Timestamp
		}
	}
	return earliest
}
