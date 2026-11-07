package glamsterdam_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/glamsterdam"
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
		require.False(t, result.AlreadyApplied)
		require.Equal(t, uint32(200_000), result.Current)
		require.Equal(t, uint32(400_000), result.AppliedValue)
	})

	t.Run("mismatched baseline applies fallback", func(t *testing.T) {
		result := glamsterdam.Resolve(spec, uint32(150_000))
		require.False(t, result.Matched)
		require.False(t, result.AlreadyApplied)
		require.Equal(t, uint32(150_000), result.Current)
		require.Equal(t, uint32(300_000), result.AppliedValue) // 150000 * 2x
	})

	t.Run("already at glamsterdam value is a no-op, not a re-scaled fallback", func(t *testing.T) {
		// This is the rerun-compounding regression: a prior run (or an earlier step of the same
		// batch) already wrote the literal GlamsterdamValue on-chain. Resolving it again must
		// leave it unchanged rather than treating 400_000 as a "mismatch" against ExpectedPrague
		// and applying the fallback ratio on top of it (which would have produced 800_000).
		result := glamsterdam.Resolve(spec, uint32(400_000))
		require.True(t, result.AlreadyApplied)
		require.False(t, result.Matched)
		require.Equal(t, uint32(400_000), result.Current)
		require.Equal(t, uint32(400_000), result.AppliedValue)
	})

	t.Run("prague and glamsterdam value can coincide without ambiguity", func(t *testing.T) {
		// Some fields (e.g. FeeQuoter.MaxPerMsgGasLimit) have ExpectedPrague == GlamsterdamValue
		// by design (no-op fields). Resolve must still report a sensible outcome rather than
		// depending on branch order to avoid a spurious AlreadyApplied/Matched conflict.
		noOpSpec := glamsterdam.FieldSpec[uint32]{
			Name:             "NoOpField",
			ExpectedPrague:   15_000_000,
			GlamsterdamValue: 15_000_000,
			Fallback:         func(current uint32) uint32 { return current },
		}
		result := glamsterdam.Resolve(noOpSpec, uint32(15_000_000))
		require.Equal(t, uint32(15_000_000), result.AppliedValue)
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

	t.Run("already applied", func(t *testing.T) {
		result := glamsterdam.Resolve(spec, uint32(400_000))
		line := glamsterdam.FieldResultString(uint64(3), result)
		require.Contains(t, line, "chain 3")
		require.Contains(t, line, "TestField")
		require.Contains(t, line, "already matches Glamsterdam value 400000")
		require.NotContains(t, line, "MISMATCH")
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
	glamsterdam.AddField(r, uint64(6), glamsterdam.Resolve(spec, uint32(400_000)))

	out := r.String()
	require.Contains(t, out, "chain 1: skipped (explicit SkipChainSelectors entry)")
	require.Contains(t, out, "chain 2: no lane to target chain, skipped")
	require.Contains(t, out, "chain 3: ERROR - could not resolve FeeQuoter address, skipping this chain")
	require.Contains(t, out, "chain 4: TestField matched expected Prague value 200000")
	require.Contains(t, out, "chain 5: TestField MISMATCH")
	require.Contains(t, out, "chain 6: TestField already matches Glamsterdam value 400000")
}

// TestResolveRerunIsIdempotent is a regression test for the exact rerun-compounding bug reported
// in review: a chain's DestGasOverhead migrates from a matched Prague baseline to the literal
// Glamsterdam value, and a *second* run against the now-migrated on-chain value must reapply the
// same Glamsterdam value rather than treating it as a fresh mismatch and re-scaling it.
func TestResolveRerunIsIdempotent(t *testing.T) {
	spec := glamsterdam.FieldSpec[uint32]{
		Name:             "FeeQuoter.DestChainConfig.DestGasOverhead",
		ExpectedPrague:   300_000,
		GlamsterdamValue: 500_000,
		Fallback:         glamsterdam.ApplyRatio[uint32](300_000, 500_000),
	}

	firstRun := glamsterdam.Resolve(spec, uint32(300_000))
	require.Equal(t, uint32(500_000), firstRun.AppliedValue)

	// Simulate the proposal having executed: on-chain state is now firstRun.AppliedValue.
	secondRun := glamsterdam.Resolve(spec, firstRun.AppliedValue)
	require.True(t, secondRun.AlreadyApplied)
	require.Equal(t, uint32(500_000), secondRun.AppliedValue, "rerun must not compound the fallback ratio on top of an already-migrated value")
}
