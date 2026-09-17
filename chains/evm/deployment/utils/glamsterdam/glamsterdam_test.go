package glamsterdam_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/glamsterdam"
	"github.com/stretchr/testify/require"
)

func TestApplyRatio(t *testing.T) {
	tests := []struct {
		desc             string
		prague           uint32
		glamsterdam      uint32
		current          uint32
		expectedFallback uint32
	}{
		{
			desc:             "2x ratio, exact",
			prague:           200_000,
			glamsterdam:      400_000,
			current:          100_000,
			expectedFallback: 200_000,
		},
		{
			desc:             "3x ratio, exact",
			prague:           90_000,
			glamsterdam:      270_000,
			current:          123_000,
			expectedFallback: 369_000,
		},
		{
			desc:             "3.2x ratio, rounds to nearest",
			prague:           20,
			glamsterdam:      64,
			current:          25,
			expectedFallback: 80, // 25 * 3.2 = 80 exactly
		},
		{
			desc:             "ratio requiring rounding",
			prague:           75_000,
			glamsterdam:      85_000,
			current:          100_000,
			expectedFallback: 113_333, // 100000 * 85000/75000 = 113333.33 -> rounds to 113333
		},
		{
			desc:             "zero current stays zero",
			prague:           200_000,
			glamsterdam:      400_000,
			current:          0,
			expectedFallback: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fallback := glamsterdam.ApplyRatio(test.prague, test.glamsterdam)
			require.Equal(t, test.expectedFallback, fallback(test.current))
		})
	}
}

func TestResolve(t *testing.T) {
	spec := glamsterdam.FieldSpec[uint32]{
		Name:             "TestField",
		ExpectedPrague:   200_000,
		GlamsterdamValue: 400_000,
		Fallback:         glamsterdam.ApplyRatio[uint32](200_000, 400_000),
	}

	t.Run("matched baseline applies literal glamsterdam value", func(t *testing.T) {
		result := glamsterdam.Resolve(spec, uint32(200_000))
		require.True(t, result.Matched)
		require.Equal(t, uint32(200_000), result.Current)
		require.Equal(t, uint32(400_000), result.AppliedValue)
	})

	t.Run("mismatched baseline applies fallback", func(t *testing.T) {
		result := glamsterdam.Resolve(spec, uint32(150_000))
		require.False(t, result.Matched)
		require.Equal(t, uint32(150_000), result.Current)
		require.Equal(t, uint32(300_000), result.AppliedValue) // 150000 * 2x
	})
}

func TestFieldResultString(t *testing.T) {
	spec := glamsterdam.FieldSpec[uint32]{
		Name:             "TestField",
		ExpectedPrague:   200_000,
		GlamsterdamValue: 400_000,
		Fallback:         glamsterdam.ApplyRatio[uint32](200_000, 400_000),
	}

	t.Run("matched", func(t *testing.T) {
		result := glamsterdam.Resolve(spec, uint32(200_000))
		line := glamsterdam.FieldResultString(uint64(1), result)
		require.Contains(t, line, "chain 1")
		require.Contains(t, line, "TestField")
		require.Contains(t, line, "matched expected Prague value 200000")
		require.Contains(t, line, "applying Glamsterdam value 400000")
	})

	t.Run("mismatched", func(t *testing.T) {
		result := glamsterdam.Resolve(spec, uint32(150_000))
		line := glamsterdam.FieldResultString(uint64(2), result)
		require.Contains(t, line, "chain 2")
		require.Contains(t, line, "MISMATCH")
		require.Contains(t, line, "current value 150000")
		require.Contains(t, line, "expected Prague value 200000")
		require.Contains(t, line, "fallback value 300000")
		require.Contains(t, line, "instead of literal Glamsterdam value 400000")
	})
}

func TestReport(t *testing.T) {
	r := glamsterdam.NewReport()
	r.AddSkipped(1)
	r.AddNoLane(2)
	r.AddUnresolvedContract(3, "FeeQuoter")

	spec := glamsterdam.FieldSpec[uint32]{
		Name:             "TestField",
		ExpectedPrague:   200_000,
		GlamsterdamValue: 400_000,
		Fallback:         glamsterdam.ApplyRatio[uint32](200_000, 400_000),
	}
	glamsterdam.AddField(r, uint64(4), glamsterdam.Resolve(spec, uint32(200_000)))
	glamsterdam.AddField(r, uint64(5), glamsterdam.Resolve(spec, uint32(150_000)))

	out := r.String()
	require.Contains(t, out, "chain 1: skipped (explicit SkipChainSelectors entry)")
	require.Contains(t, out, "chain 2: no lane to target chain, skipped")
	require.Contains(t, out, "chain 3: ERROR - could not resolve FeeQuoter address, skipping this chain")
	require.Contains(t, out, "chain 4: TestField matched expected Prague value 200000")
	require.Contains(t, out, "chain 5: TestField MISMATCH")

	lines := []string{
		"chain 1: skipped (explicit SkipChainSelectors entry)",
		"chain 2: no lane to target chain, skipped",
		"chain 3: ERROR - could not resolve FeeQuoter address, skipping this chain",
	}
	for i, want := range lines {
		require.Contains(t, out, want, "line %d", i)
	}
}
