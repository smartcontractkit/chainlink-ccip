package consensus

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

	minObs := MakeConstantThreshold[cciptypes.ChainSelector](Threshold(f))
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

	threshold := MakeConstantThreshold[int](Threshold(f))
	timeFinal := GetConsensusMapAggregator(lggr, "time", timeValues, threshold, func(vals []time.Time) time.Time {
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

	threshold := MakeConstantThreshold[int](Threshold(f))
	intFinal := GetConsensusMapAggregator(lggr, "int", intValues, threshold, func(vals []int) int {
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

func TestMakeConstantThreshold(t *testing.T) {
	{
		f := 3
		threshold := MakeConstantThreshold[cciptypes.ChainSelector](Threshold(f))

		for i := 0; i < 10; i++ {
			thresh, ok := threshold.Get(cciptypes.ChainSelector(i))
			assert.True(t, ok)
			assert.Equal(t, Threshold(f), thresh)
		}
	}
	{
		f := 3.5
		threshold := MakeConstantThreshold[float64](Threshold(f))

		for i := 0; i < 10; i++ {
			thresh, ok := threshold.Get(float64(i))
			assert.True(t, ok)
			assert.Equal(t, Threshold(f), thresh)
		}
	}
}

func TestMakeMultiThreshold(t *testing.T) {
	fChain := map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1): 1,
		cciptypes.ChainSelector(2): 2,
		cciptypes.ChainSelector(3): 3,
	}

	{
		// Using TwoFPlus1
		threshold := MakeMultiThreshold(fChain, TwoFPlus1)

		thresh, ok := threshold.Get(cciptypes.ChainSelector(1))
		assert.True(t, ok)
		assert.Equal(t, Threshold(3), thresh)

		thresh, ok = threshold.Get(cciptypes.ChainSelector(2))
		assert.True(t, ok)
		assert.Equal(t, Threshold(5), thresh)

		thresh, ok = threshold.Get(cciptypes.ChainSelector(3))
		assert.True(t, ok)
		assert.Equal(t, Threshold(7), thresh)

		_, ok = threshold.Get(cciptypes.ChainSelector(4))
		assert.False(t, ok)
	}
	{
		// Using FPlus1
		threshold := MakeMultiThreshold(fChain, FPlus1)

		thresh, ok := threshold.Get(cciptypes.ChainSelector(1))
		assert.True(t, ok)
		assert.Equal(t, Threshold(2), thresh)

		thresh, ok = threshold.Get(cciptypes.ChainSelector(2))
		assert.True(t, ok)
		assert.Equal(t, Threshold(3), thresh)

		thresh, ok = threshold.Get(cciptypes.ChainSelector(3))
		assert.True(t, ok)
		assert.Equal(t, Threshold(4), thresh)

		_, ok = threshold.Get(cciptypes.ChainSelector(4))
		assert.False(t, ok)
	}
	{
		// Custom mapping which sets it to a constant.
		threshold := MakeMultiThreshold(fChain, func(i int) Threshold {
			return 1337
		})

		thresh, ok := threshold.Get(cciptypes.ChainSelector(1))
		assert.True(t, ok)
		assert.Equal(t, Threshold(1337), thresh)

		thresh, ok = threshold.Get(cciptypes.ChainSelector(2))
		assert.True(t, ok)
		assert.Equal(t, Threshold(1337), thresh)

		thresh, ok = threshold.Get(cciptypes.ChainSelector(3))
		assert.True(t, ok)
		assert.Equal(t, Threshold(1337), thresh)

		_, ok = threshold.Get(cciptypes.ChainSelector(4))
		assert.False(t, ok)
	}
}

func TestMakeMultiThreshold_GenericKey(t *testing.T) {
	fChain := map[float64]int{
		1.23: 1,
		2.34: 2,
		8.91: 3,
	}

	// Not a chain selector for the key.
	threshold := MakeMultiThreshold(fChain, FPlus1)

	thresh, ok := threshold.Get(1.23)
	assert.True(t, ok)
	assert.Equal(t, Threshold(2), thresh)

	thresh, ok = threshold.Get(2.34)
	assert.True(t, ok)
	assert.Equal(t, Threshold(3), thresh)

	thresh, ok = threshold.Get(8.91)
	assert.True(t, ok)
	assert.Equal(t, Threshold(4), thresh)

	_, ok = threshold.Get(4)
	assert.False(t, ok)
}