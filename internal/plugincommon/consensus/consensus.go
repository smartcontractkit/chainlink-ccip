package consensus

import (
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
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
				lggr.Warnf("failed to reach consensus on a %s's for key %+v "+
					"because no single item was observed more than the expected min (%d) times, "+
					"all observed items: %v",
					objectName, key, minThresh, items)
			} else {
				consensus[key] = items[0]
			}
		} else {
			// TODO: metrics
			lggr.Warnf("getConsensus(%s): min not found for chain %d", objectName, key)
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
		if thresh, ok := f.Get(key); ok && len(values) < int(thresh) {
			lggr.Warnf("could not reach consensus on %s for key %v", objectName, key)
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

func TimestampComparator(a, b time.Time) bool {
	return a.Before(b)
}

func BigIntComparator(a, b cciptypes.BigInt) bool {
	return a.Cmp(b.Int) == -1
}

func TokenPriceComparator(a, b cciptypes.TokenPrice) bool {
	return a.Price.Int.Cmp(b.Price.Int) == -1
}
