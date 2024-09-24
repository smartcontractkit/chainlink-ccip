package plugincommon

import (
	"fmt"
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

	timeFinal := GetConsensusMapAggregator(lggr, "time", timeValues, f, func(vals []time.Time) time.Time {
		return Median(vals, TimestampComparator)
	})

	assert.Equal(t, map[int]time.Time{
		1: ts1,
		2: ts1,
		4: ts2,
	}, timeFinal)
}

// Test_GetConsensusMapMedian tests the GetConsensusMapAggregator function.
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

	intFinal := GetConsensusMapAggregator(lggr, "int", intValues, f, func(vals []int) int {
		return Median(vals, func(a, b int) bool {
			return a < b
		})
	})

	assert.Equal(t, map[int]int{
		1: 5,
		2: 5,
		4: 4,
		5: 3,
	}, intFinal)
}

var BI = func(x int64) cciptypes.BigInt {
	return cciptypes.NewBigIntFromInt64(x)
}

func TestMedianBigInt(t *testing.T) {
	tests := []struct {
		name string
		vals []cciptypes.BigInt
		want cciptypes.BigInt
	}{
		{
			name: "base case",
			vals: []cciptypes.BigInt{BI(1), BI(2), BI(4), BI(5)},
			want: BI(4),
		},
		{
			name: "not sorted",
			vals: []cciptypes.BigInt{BI(100), BI(50), BI(30), BI(110)},
			want: cciptypes.NewBigIntFromInt64(100),
		},
		{
			name: "empty slice",
			vals: []cciptypes.BigInt{},
			want: cciptypes.BigInt{},
		},
		{
			name: "one item",
			vals: []cciptypes.BigInt{BI(123)},
			want: BI(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Median(tt.vals, BigIntComparator), "Median(%v)", tt.vals)
		})
	}
}

func Test_HonestMajorityThreshold(t *testing.T) {
	tests := []struct {
		f    int
		want int
	}{
		{f: 0, want: 1},
		{f: 1, want: 3},
		{f: 2, want: 5},
		{f: 3, want: 7},
		{f: 10, want: 21},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("f=%d", tt.f), func(t *testing.T) {
			got := HonestMajorityThreshold(tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_HonestMajorityThresholdMap(t *testing.T) {
	tests := []struct {
		name string
		fMap map[cciptypes.ChainSelector]int
		want map[cciptypes.ChainSelector]int
	}{
		{
			name: "base case",
			fMap: map[cciptypes.ChainSelector]int{
				cciptypes.ChainSelector(1): 0,
				cciptypes.ChainSelector(2): 1,
				cciptypes.ChainSelector(3): 2,
			},
			want: map[cciptypes.ChainSelector]int{
				cciptypes.ChainSelector(1): 1,
				cciptypes.ChainSelector(2): 3,
				cciptypes.ChainSelector(3): 5,
			},
		},
		{
			name: "empty map",
			fMap: map[cciptypes.ChainSelector]int{},
			want: map[cciptypes.ChainSelector]int{},
		},
		{
			name: "single entry",
			fMap: map[cciptypes.ChainSelector]int{
				cciptypes.ChainSelector(1): 10,
			},
			want: map[cciptypes.ChainSelector]int{
				cciptypes.ChainSelector(1): 21,
			},
		},
		{
			name: "multiple entries",
			fMap: map[cciptypes.ChainSelector]int{
				cciptypes.ChainSelector(1): 3,
				cciptypes.ChainSelector(2): 5,
				cciptypes.ChainSelector(3): 7,
			},
			want: map[cciptypes.ChainSelector]int{
				cciptypes.ChainSelector(1): 7,
				cciptypes.ChainSelector(2): 11,
				cciptypes.ChainSelector(3): 15,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HonestMajorityThresholdMap(tt.fMap)
			assert.Equal(t, tt.want, got)
		})
	}
}
