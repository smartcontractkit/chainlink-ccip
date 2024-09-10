package shared

import (
	"sort"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/shared/filter"
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
			minObservations := filter.NewMinObservation[T](min, nil)
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

func GetConsensusMapMedian[K comparable, T any](
	lggr logger.Logger,
	objectName string,
	items map[K][]T,
	f int,
	comp func(a, b T) bool,
) map[K]T {
	consensus := make(map[K]T)

	for key, items := range items {
		if len(items) < f {
			lggr.Warnf("could not reach consensus on %s for key %v", objectName, key)
			continue
		}
		consensus[key] = Median(items, comp)
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
