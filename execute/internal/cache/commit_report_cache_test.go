package cache

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	defaultQueryLimit  = 10
	defaultReturnLimit = 10
)

func TestCommitReportsCache_GetCachedAndNewReports(t *testing.T) {
	lggr := logger.Test(t)

	now := time.Now().UTC()
	fetchFrom := now.Add(-10 * time.Minute)
	r1 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-8 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0xa1}},
			},
		},
	}
	r2 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-5 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0xa2}},
			},
		},
	}

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 2*time.Hour, mockReader, defaultQueryLimit, defaultReturnLimit)

	// We capture the timestamp parameters to ensure they advance between calls.
	var tsFinalizedCalls []time.Time

	// First round expectations
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, defaultQueryLimit,
		).
		RunAndReturn(func(
			_ context.Context,
			ts time.Time,
			_ primitives.ConfidenceLevel,
			_ int) ([]ccipocr3.CommitPluginReportWithMeta, error) {
			tsFinalizedCalls = append(tsFinalizedCalls, ts)
			return []ccipocr3.CommitPluginReportWithMeta{r1}, nil
		}).Once()

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, mock.AnythingOfType("time.Time"), primitives.Unconfirmed, defaultQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{r1, r2}, nil).Once()

	// First invocation
	got, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)
	assert.Len(t, got, 2)

	// Second round expectations
	// finalized query returns nothing new but we record the ts to ensure advancement
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, defaultQueryLimit,
		).
		RunAndReturn(func(_ context.Context, ts time.Time, _ primitives.ConfidenceLevel, _ int) (
			[]ccipocr3.CommitPluginReportWithMeta,
			error) {
			tsFinalizedCalls = append(tsFinalizedCalls, ts)
			return nil, nil
		}).Once()

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, mock.AnythingOfType("time.Time"), primitives.Unconfirmed, defaultQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{r1, r2}, nil).Once()

	// Second invocation (no new finalized data)
	got2, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)
	assert.Len(t, got2, 2)

	// Validate timestamp advancement
	assert.Len(t, tsFinalizedCalls, 2)
	firstTs := tsFinalizedCalls[0]
	secondTs := tsFinalizedCalls[1]
	assert.Equal(t, fetchFrom, firstTs) // First call uses fetchFrom
	// Second call's queryFrom should be the timestamp of the latest finalized report from the first call (r1)
	// or fetchFrom if no finalized reports were returned in the first call. Here it's r1.
	assert.Equal(t, r1.Timestamp, secondTs)
	assert.True(t, secondTs.After(firstTs)) // Ensure time advanced
	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_LimitAndDeduplication(t *testing.T) {
	lggr := logger.Test(t)
	testQueryLimit := 3
	testReturnLimit := 3

	now := time.Now().UTC()
	initialFetchFrom := now.Add(-10 * time.Minute)

	// Create 5 reports with unique timestamps (descending)
	reports := make([]ccipocr3.CommitPluginReportWithMeta, 5)
	for i := 0; i < 5; i++ {
		// Add unique MerkleRoot to ensure generateKey creates unique keys
		reports[i] = ccipocr3.CommitPluginReportWithMeta{
			Timestamp: now.Add(-time.Duration(i) * time.Minute),
			Report: ccipocr3.CommitPluginReport{
				BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
					{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{byte(i + 1)}},
				},
			},
		}
	}

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, testQueryLimit, testReturnLimit)

	// First call: finalized returns first 3 (reports[0], reports[1], reports[2]).
	// queryLimit is 3 for reader calls.
	// Early exit will occur because len(3) >= returnLimit(3), so unconfirmed is not called.
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, initialFetchFrom, primitives.Finalized, testQueryLimit,
		).
		Return(reports[:3], nil).Once()

	// returnLimit = 3 should truncate to newest 3
	got, err := cache.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)
	assert.Len(t, got, testReturnLimit)
	// Ensure dedup preserved order newest first (reports[0] is newest)
	assert.Equal(t, reports[0].Timestamp, got[0].Timestamp)
	assert.Equal(t, reports[1].Timestamp, got[1].Timestamp)
	assert.Equal(t, reports[2].Timestamp, got[2].Timestamp)

	// Second call:
	// lastQueryTimestamp should be reports[0].Timestamp (most recent from previous finalized set).
	// Finalized returns nothing.
	// Early exit will occur again because cache still has 3 items and returnLimit is 3.
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, reports[0].Timestamp, primitives.Finalized, testQueryLimit,
		).
		Return(nil, nil).Once()

	got2, err := cache.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)
	assert.Len(t, got2, testReturnLimit)
	assert.Equal(t, reports[0].Timestamp, got2[0].Timestamp)
	assert.Equal(t, reports[1].Timestamp, got2[1].Timestamp)
	assert.Equal(t, reports[2].Timestamp, got2[2].Timestamp)
	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_DedupAndFetchWindow(t *testing.T) {
	lggr := logger.Test(t)
	testQueryLimit := 5
	testReturnLimit := 5

	now := time.Now().UTC()
	// This is the initial fetchFrom for the first call to GetCachedAndNewReports
	initialFetchFrom := now.Add(-10 * time.Minute)
	// This is the report that will be fetched and cached
	r := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-2 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0xb1}},
			},
		},
	}

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 30*time.Minute, mockReader, testQueryLimit, testReturnLimit)

	// --- First invocation ---
	// Purpose: Populate the cache with report 'r'.
	// Expected queryFrom for reader: initialFetchFrom
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, initialFetchFrom, primitives.Finalized, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once()
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, initialFetchFrom, primitives.Unconfirmed, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once()

	got1, err1 := cache.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err1)
	assert.Len(t, got1, 1, "First call should return the single fetched report")
	// After this call, cache.lastQueryTimestamp should be r.Timestamp (now - 2m)

	// --- Second invocation ---
	// Purpose: Test deduplication. Fetch 'r' again; should still get only one 'r'.
	// Expected queryFrom for reader: r.Timestamp (because lastQueryTimestamp was updated)
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, r.Timestamp, primitives.Finalized, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once() // Reader returns 'r' again
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, r.Timestamp, primitives.Unconfirmed, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once() // Reader returns 'r' again

	// The 'fetchFrom' for GetCachedAndNewReports itself is still initialFetchFrom for this logical step
	got2, err2 := cache.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err2)
	assert.Len(t, got2, 1, "Second call should return the single report due to deduplication")
	// After this, cache.lastQueryTimestamp should remain r.Timestamp

	// --- Third invocation ---
	// Purpose: Test that fetching with a window newer than the cached item yields nothing.
	// The 'fetchFrom' for GetCachedAndNewReports is now newer than 'r.Timestamp'.
	newerFetchWindowStart := now.Add(-1 * time.Minute)
	// Expected queryFrom for reader: r.Timestamp (lastQueryTimestamp hasn't changed meaningfully beyond r.Timestamp)
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, r.Timestamp, primitives.Finalized, testQueryLimit,
		).
		Return(nil, nil).Once() // Reader finds nothing new from r.Timestamp
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, r.Timestamp, primitives.Unconfirmed, testQueryLimit,
		).
		Return(nil, nil).Once() // Reader finds nothing new from r.Timestamp

	got3, err3 := cache.GetCachedAndNewReports(context.Background(), newerFetchWindowStart)
	assert.NoError(t, err3)
	assert.Len(t, got3, 0, "Third call with newer fetch window should return no reports")

	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_ReaderError_Finalized(t *testing.T) {
	lggr := logger.Test(t)
	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	expectedErr := errors.New("finalized reader error")

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, mock.Anything, primitives.Finalized, defaultQueryLimit,
		).
		Return(nil, expectedErr).Once()

	c := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, defaultQueryLimit, defaultReturnLimit)
	_, err := c.GetCachedAndNewReports(context.Background(), time.Now())

	assert.Error(t, err)
	assert.True(t, errors.Is(err, expectedErr), "Error should wrap the expected reader error")
	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_ReaderError_Unconfirmed(t *testing.T) {
	lggr := logger.Test(t)
	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	expectedErr := errors.New("unconfirmed reader error")

	// Finalized succeeds
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, mock.Anything, primitives.Finalized, defaultQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{}, nil).Once()

	// Unconfirmed fails
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, mock.Anything, primitives.Unconfirmed, defaultQueryLimit,
		).
		Return(nil, expectedErr).Once()

	c := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, defaultQueryLimit, defaultReturnLimit)
	_, err := c.GetCachedAndNewReports(context.Background(), time.Now())

	assert.Error(t, err)
	assert.True(t, errors.Is(err, expectedErr), "Error should wrap the expected reader error")
	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_MergeSortLimit(t *testing.T) {
	lggr := logger.Test(t)
	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	now := time.Now().UTC()

	testQueryLimit := 3
	testReturnLimit := 3

	// Reports to be initially in cache (simulated by first fetch)
	rOld1 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-20 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x01}},
			},
		},
	}
	rOld2 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-10 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x02}},
			},
		},
	}

	// Reports to be returned by the unconfirmed reader call
	rNew1 := ccipocr3.CommitPluginReportWithMeta{ // Newest
		Timestamp: now.Add(-5 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x03}},
			},
		},
	}
	rNew2DupOld2 := ccipocr3.CommitPluginReportWithMeta{ // Duplicate of rOld2, same timestamp
		Timestamp: now.Add(-10 * time.Minute), // Key will be same as rOld2
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x02}},
			},
		},
	}
	rNew3 := ccipocr3.CommitPluginReportWithMeta{ // Older than rOld2/rNew2_DupOld2 but newer than rOld1
		Timestamp: now.Add(-15 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x04}},
			},
		},
	}
	rNew4 := ccipocr3.CommitPluginReportWithMeta{ // Oldest
		Timestamp: now.Add(-25 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x05}},
			},
		},
	}

	c := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, testQueryLimit, testReturnLimit)
	initialFetchFrom := now.Add(-30 * time.Minute)

	// First call: Populate cache with rOld1, rOld2 from finalized & unconfirmed
	// queryLimit is testQueryLimit (3). Reader returns 2 reports, which is fine.
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, initialFetchFrom, primitives.Finalized, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rOld1, rOld2}, nil).Once()
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, initialFetchFrom, primitives.Unconfirmed, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rOld1, rOld2}, nil).Once()

	// The result of this call will be limited by testReturnLimit (3). Here it returns 2.
	_, err := c.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)
	// lastQueryTimestamp is now rOld2.Timestamp

	// Second call: Test merge, sort, limit
	// Cached: rOld1, rOld2 (rOld2 is newer). LastQueryTimestamp should be rOld2.Timestamp.
	// Unconfirmed fetch will return: rNew1, rNew2_DupOld2, rNew3, rNew4

	// Expect finalized to be called from rOld2.Timestamp, return nothing new for simplicity
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rOld2.Timestamp, primitives.Finalized, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{}, nil).Once()
	// Expect unconfirmed to be called from rOld2.Timestamp
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rOld2.Timestamp, primitives.Unconfirmed, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rNew1, rNew2DupOld2, rNew3, rNew4}, nil).Once()

	// For this call, we are interested in reports >= initialFetchFrom. All our reports satisfy this.
	finalReports, err := c.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)

	// Expected unique reports after merge and sort, before limit:
	// rNew1 (ts: now-5m, key 0x03)
	// rNew2_DupOld2 (ts: now-10m, key 0x02) - this overwrites rOld2 due to same key
	// rNew3 (ts: now-15m, key 0x04)
	// rOld1 (ts: now-20m, key 0x01)
	// rNew4 (ts: now-25m, key 0x05) - this would be included if limit was higher and queryLimit was higher

	// After sorting (newest first) and limiting to testReturnLimit (3):
	// 1. rNew1 (now-5m)
	// 2. rNew2_DupOld2 (now-10m)
	// 3. rNew3 (now-15m)

	assert.Len(t, finalReports, testReturnLimit)
	assert.Equal(t,
		rNew1.Report.BlessedMerkleRoots[0].MerkleRoot,
		finalReports[0].Report.BlessedMerkleRoots[0].MerkleRoot,
	)
	assert.Equal(t,
		rNew2DupOld2.Report.BlessedMerkleRoots[0].MerkleRoot,
		finalReports[1].Report.BlessedMerkleRoots[0].MerkleRoot,
	)
	assert.Equal(t,
		rNew3.Report.BlessedMerkleRoots[0].MerkleRoot,
		finalReports[2].Report.BlessedMerkleRoots[0].MerkleRoot,
	)

	// Ensure timestamps are in correct descending order
	assert.True(t, finalReports[0].Timestamp.After(finalReports[1].Timestamp))
	assert.True(t, finalReports[1].Timestamp.After(finalReports[2].Timestamp))

	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_EarlyExitOptimization(t *testing.T) {
	lggr := logger.Test(t)
	now := time.Now().UTC()
	fetchFrom := now.Add(-30 * time.Minute)

	queryLimit := 5
	returnLimit := 2 // We want to exit early if cache has 2 or more

	rCached1 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-20 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0xc1}},
			},
		},
	}
	rFinalized1 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-5 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0xf1}},
			},
		},
	}
	rFinalized2 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-10 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0xf2}},
			},
		},
	}
	// This report would be fetched by unconfirmed if early exit didn't happen
	// rUnconfirmedSkipped := ccipocr3.CommitPluginReportWithMeta{
	// 	Timestamp: now.Add(-1 * time.Minute),
	// 	Report: ccipocr3.CommitPluginReport{
	// 		BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
	// 			{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x5A}},
	// 		},
	// 	},
	// }

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, queryLimit, returnLimit)

	// 1. Pre-populate cache with rCached1 (simulating it's already there from a previous run)
	// For this, we make a setup call.
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, fetchFrom, primitives.Finalized, queryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rCached1}, nil).Once()
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, fetchFrom, primitives.Unconfirmed, queryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rCached1}, nil).Once()
	_, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)
	// lastQueryTimestamp is now rCached1.Timestamp

	// 2. Test call that should trigger early exit
	// Finalized reports fetch will bring in rFinalized1, rFinalized2
	// Query from rCached1.Timestamp
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rCached1.Timestamp, primitives.Finalized, queryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rFinalized1, rFinalized2}, nil).Once()

	// IMPORTANT: Unconfirmed reader should NOT be called due to early exit.
	// We don't set an EXPECT for it for this specific interaction path.
	// If it were called, the mock would complain about an unexpected call.

	reports, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)

	// After new finalized reports are added, cache has: rCached1, rFinalized1, rFinalized2
	// Sorted by time (newest first): rFinalized1, rFinalized2, rCached1
	// Since len (3) >= returnLimit (2), we should get the newest 2.
	assert.Len(t, reports, returnLimit)
	assert.Equal(t, rFinalized1, reports[0]) // Newest
	assert.Equal(t, rFinalized2, reports[1]) // Second newest

	// Verify that unconfirmed was not called for the second GetCachedAndNewReports interaction.
	// AssertExpectations will check if all *defined* expectations were met.
	// The absence of an expectation for unconfirmed for the second call is key.
	mockReader.AssertExpectations(t)

	// As an extra check, if we somehow defined an unconfirmed call that should not happen:
	// mockReader.EXPECT().
	// 	CommitReportsGTETimestamp(mock.Anything, mock.Anything, primitives.Unconfirmed, mock.Anything).
	// 	Return([]ccipocr3.CommitPluginReportWithMeta{rUnconfirmedSkipped}, nil).Times(0)
	// This .Times(0) would also work but not setting it is cleaner if that's the only
	// interaction we want to avoid for a path.
}
