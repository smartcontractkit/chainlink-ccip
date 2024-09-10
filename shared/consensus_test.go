package shared

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Test_fChainConsensus is a legacy test that was replaced by GetConsensusMap, now it's a regression test.
func Test_fChainConsensus(t *testing.T) {
	lggr := logger.Test(t)
	f := 3
	fChainValues := map[cciptypes.ChainSelector][]int{
		cciptypes.ChainSelector(1): {5, 5, 5, 5, 5},
		cciptypes.ChainSelector(2): {5, 5, 5, 3, 5},
		cciptypes.ChainSelector(3): {5, 5},             // not enough observations, must be observed at least f times.
		cciptypes.ChainSelector(4): {5, 3, 5, 3, 5, 3}, // both values appear at least f times, no consensus
	}

	minObs := map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1): f,
		cciptypes.ChainSelector(2): f,
		cciptypes.ChainSelector(3): f,
		cciptypes.ChainSelector(4): f,
	}
	fChainFinal := GetConsensusMap(lggr, "fChain", fChainValues, minObs)

	assert.Equal(t, map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1): 5,
		cciptypes.ChainSelector(2): 5,
	}, fChainFinal)
}

func Test_GetConsensusMapMedianTimestamp(t *testing.T) {
	lggr := logger.Test(t)
	f := 3

	ts1 := time.Now()
	ts2 := ts1.Add(time.Minute)
	ts3 := ts1.Add(2 * time.Minute)

	timeValues := map[int][]time.Time{
		1: {ts1, ts1, ts1, ts1, ts1},
		2: {ts2, ts2, ts1, ts1, ts1},
		3: {ts1, ts1}, // not enough observations, must be observed at least f times.
		4: {ts1, ts2, ts3},
	}

	timeFinal := GetConsensusMapMedian(lggr, "time", timeValues, f, func(a, b time.Time) bool {
		return a.Before(b)
	})

	assert.Equal(t, map[int]time.Time{
		1: ts1,
		2: ts1,
		4: ts2,
	}, timeFinal)
}

// Test_GetConsensusMapMedian tests the GetConsensusMapMedian function.
func Test_GetConsensusMapMedianInt(t *testing.T) {
	lggr := logger.Test(t)
	f := 3

	// Test with integers
	intValues := map[int][]int{
		1: {5, 5, 5, 5, 5},
		2: {5, 5, 5, 3, 5},
		3: {5, 5}, // not enough observations, must be observed at least f times.
		4: {1, 2, 3, 4, 5, 6},
		5: {5, 4, 3, 2, 1},
	}

	intFinal := GetConsensusMapMedian(lggr, "int", intValues, f, func(a, b int) bool {
		return a < b
	})

	assert.Equal(t, map[int]int{
		1: 5,
		2: 5,
		4: 4,
		5: 3,
	}, intFinal)
}
