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

// Helper to create a report with a single blessed root for brevity in tests
func newReportWithRoot(ts time.Time, rootByte byte) ccipocr3.CommitPluginReportWithMeta {
	return ccipocr3.CommitPluginReportWithMeta{
		Timestamp: ts,
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{rootByte}},
			},
		},
	}
}

func newRootlessReport(ts time.Time) ccipocr3.CommitPluginReportWithMeta {
	return ccipocr3.CommitPluginReportWithMeta{
		Timestamp: ts,
		Report:    ccipocr3.CommitPluginReport{ /* No roots */ },
	}
}

func TestCommitReportsCache_GetCachedAndNewReports(t *testing.T) {
	lggr := logger.Test(t)

	now := time.Now().UTC()
	fetchFrom := now.Add(-10 * time.Minute)

	rWithRoots1 := newReportWithRoot(now.Add(-8*time.Minute), 0xa1)      // Older, with roots
	rWithRoots2 := newReportWithRoot(now.Add(-5*time.Minute), 0xa2)      // Newer, with roots
	rFinalizedRootless := newRootlessReport(now.Add(-7 * time.Minute))   // Newer than rWithRoots1, but rootless
	rUnconfirmedRootless := newRootlessReport(now.Add(-4 * time.Minute)) // Newest overall, but rootless

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 2*time.Hour, mockReader, defaultQueryLimit, defaultReturnLimit)

	var tsFinalizedCalls []time.Time

	// --- First invocation ---
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, fetchFrom, primitives.Finalized, defaultQueryLimit,
		).
		RunAndReturn(func(
			_ context.Context,
			ts time.Time,
			_ primitives.ConfidenceLevel,
			_ int) ([]ccipocr3.CommitPluginReportWithMeta, error) {
			tsFinalizedCalls = append(tsFinalizedCalls, ts)
			// Return one with roots, one without. rWithRoots1 is older.
			return []ccipocr3.CommitPluginReportWithMeta{rWithRoots1, rFinalizedRootless}, nil
		}).Once()

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, fetchFrom, primitives.Unconfirmed, defaultQueryLimit,
		).
		// rUnconfirmedRootless is newest overall but should be filtered.
		Return([]ccipocr3.CommitPluginReportWithMeta{rWithRoots1, rWithRoots2, rUnconfirmedRootless}, nil).Once()

	got, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)
	// Expected: rWithRoots2, rWithRoots1 (sorted, rootless ones filtered)
	assert.Len(t, got, 2)
	assert.Equal(t, rWithRoots2.Report.BlessedMerkleRoots[0].MerkleRoot, got[0].Report.BlessedMerkleRoots[0].MerkleRoot)
	assert.Equal(t, rWithRoots1.Report.BlessedMerkleRoots[0].MerkleRoot, got[1].Report.BlessedMerkleRoots[0].MerkleRoot)

	// lastQueryTimestamp should be rWithRoots1.Timestamp as rFinalizedRootless was skipped for caching
	assert.Equal(t, rWithRoots1.Timestamp, cache.lastQueryTimestamp)

	// --- Second invocation ---
	// queryFrom should be rWithRoots1.Timestamp (previous lastQueryTimestamp)
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rWithRoots1.Timestamp, primitives.Finalized, defaultQueryLimit,
		).
		RunAndReturn(func(_ context.Context, ts time.Time, _ primitives.ConfidenceLevel, _ int) (
			[]ccipocr3.CommitPluginReportWithMeta,
			error) {
			tsFinalizedCalls = append(tsFinalizedCalls, ts)
			return nil, nil // No new finalized reports with roots
		}).Once()

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rWithRoots1.Timestamp, primitives.Unconfirmed, defaultQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rWithRoots1, rWithRoots2, rUnconfirmedRootless}, nil).Once()

	got2, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)
	// Expected: rWithRoots2, rWithRoots1 (rootless still filtered)
	assert.Len(t, got2, 2)
	assert.Equal(t, rWithRoots2.Report.BlessedMerkleRoots[0].MerkleRoot, got2[0].Report.BlessedMerkleRoots[0].MerkleRoot)
	assert.Equal(t, rWithRoots1.Report.BlessedMerkleRoots[0].MerkleRoot, got2[1].Report.BlessedMerkleRoots[0].MerkleRoot)

	// Validate timestamp advancement for CommitReportsGTETimestamp(Finalized) calls
	assert.Len(t, tsFinalizedCalls, 2)
	assert.Equal(t, fetchFrom, tsFinalizedCalls[0])             // First call uses fetchFrom
	assert.Equal(t, rWithRoots1.Timestamp, tsFinalizedCalls[1]) // Second call uses previous lastQueryTimestamp
	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_LimitAndDeduplication(t *testing.T) {
	lggr := logger.Test(t)
	testQueryLimit := 3
	testReturnLimit := 3

	now := time.Now().UTC()
	initialFetchFrom := now.Add(-10 * time.Minute)

	// Create reports with unique timestamps (descending)
	reportsWithRoots := make([]ccipocr3.CommitPluginReportWithMeta, 5)
	for i := 0; i < 5; i++ {
		// reportsWithRoots[0] is newest
		reportsWithRoots[i] = newReportWithRoot(now.Add(-time.Duration(i)*time.Minute), byte(i+1))
	}

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, testQueryLimit, testReturnLimit)

	// --- First call ---
	// Finalized returns first 3 reports with roots.
	// Early exit will occur because len(3) >= returnLimit(3), so unconfirmed is not called.
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, initialFetchFrom, primitives.Finalized, testQueryLimit,
		).
		Return(reportsWithRoots[:3], nil).Once()

	got, err := cache.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)
	assert.Len(t, got, testReturnLimit)
	assert.Equal(t, reportsWithRoots[0].Timestamp, got[0].Timestamp)
	assert.Equal(t, reportsWithRoots[1].Timestamp, got[1].Timestamp)
	assert.Equal(t, reportsWithRoots[2].Timestamp, got[2].Timestamp)

	// --- Second call ---
	// lastQueryTimestamp should be reportsWithRoots[0].Timestamp (most recent from previous finalized set).
	// Finalized returns nothing.
	// Early exit will occur again because cache still has 3 items and returnLimit is 3.
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, reportsWithRoots[0].Timestamp, primitives.Finalized, testQueryLimit,
		).
		Return(nil, nil).Once()

	got2, err := cache.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)
	assert.Len(t, got2, testReturnLimit)
	assert.Equal(t, reportsWithRoots[0].Timestamp, got2[0].Timestamp)
	assert.Equal(t, reportsWithRoots[1].Timestamp, got2[1].Timestamp)
	assert.Equal(t, reportsWithRoots[2].Timestamp, got2[2].Timestamp)
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

	testQueryLimit := 5  // Allow fetching more to test merge logic properly
	testReturnLimit := 3 // Final result should be 3

	// Reports to be initially in cache (simulated by first fetch)
	rOld1 := newReportWithRoot(now.Add(-20*time.Minute), 0x01)
	rOld2 := newReportWithRoot(now.Add(-10*time.Minute), 0x02)

	// Reports to be returned by the unconfirmed reader call
	rNewWithRoots1 := newReportWithRoot(now.Add(-5*time.Minute), 0x03)         // Newest with roots
	rNewRootless := newRootlessReport(now.Add(-2 * time.Minute))               // Absolute newest, but rootless
	rNewWithRoots2DupOld2 := newReportWithRoot(now.Add(-10*time.Minute), 0x02) // Same as rOld2
	rNewWithRoots3 := newReportWithRoot(now.Add(-15*time.Minute), 0x04)

	c := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, testQueryLimit, testReturnLimit)
	initialFetchFrom := now.Add(-30 * time.Minute)

	// --- First call: Populate cache with rOld1, rOld2 ---
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

	_, err := c.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)
	// lastQueryTimestamp is now rOld2.Timestamp (newest from rOld1, rOld2)

	// --- Second call: Test merge, sort, limit with rootless filtering ---
	// Cached: rOld1, rOld2. LastQueryTimestamp should be rOld2.Timestamp.
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rOld2.Timestamp, primitives.Finalized, testQueryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{}, nil).Once() // No new finalized

	newUnconfirmedReports := []ccipocr3.CommitPluginReportWithMeta{
		rNewWithRoots1, rNewRootless, rNewWithRoots2DupOld2, rNewWithRoots3,
	}
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rOld2.Timestamp, primitives.Unconfirmed, testQueryLimit,
		).
		Return(newUnconfirmedReports, nil).Once()

	finalReports, err := c.GetCachedAndNewReports(context.Background(), initialFetchFrom)
	assert.NoError(t, err)

	// Expected unique reports *with roots* after merge and sort, before limit:
	// rNewWithRoots1 (ts: now-5m)
	// rNewWithRoots2DupOld2 (ts: now-10m) (effectively rOld2)
	// rNewWithRoots3 (ts: now-15m)
	// rOld1 (ts: now-20m)
	// rNewRootless (ts: now-2m) is filtered out.

	// After sorting (newest first) and limiting to testReturnLimit (3):
	// 1. rNewWithRoots1
	// 2. rNewWithRoots2DupOld2 (or rOld2)
	// 3. rNewWithRoots3

	assert.Len(t, finalReports, testReturnLimit)
	assert.Equal(t,
		rNewWithRoots1.Report.BlessedMerkleRoots[0].MerkleRoot,
		finalReports[0].Report.BlessedMerkleRoots[0].MerkleRoot,
	)
	assert.Equal(t,
		rNewWithRoots2DupOld2.Report.BlessedMerkleRoots[0].MerkleRoot,
		finalReports[1].Report.BlessedMerkleRoots[0].MerkleRoot,
	)
	assert.Equal(t,
		rNewWithRoots3.Report.BlessedMerkleRoots[0].MerkleRoot,
		finalReports[2].Report.BlessedMerkleRoots[0].MerkleRoot,
	)

	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_EarlyExitOptimization(t *testing.T) {
	lggr := logger.Test(t)
	now := time.Now().UTC()
	fetchFrom := now.Add(-30 * time.Minute)

	queryLimit := 5
	returnLimit := 2 // We want to exit early if cache has 2 or more reports with roots

	rCachedWithRoots1 := newReportWithRoot(now.Add(-20*time.Minute), 0xc1)
	// rFinalizedRootless1 is newer but has no roots, should be processed by addReportsToCache but not cached
	rFinalizedRootless1 := newRootlessReport(now.Add(-5 * time.Minute))
	rFinalizedWithRoots2 := newReportWithRoot(now.Add(-10*time.Minute), 0xf2)

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 1*time.Hour, mockReader, queryLimit, returnLimit)

	// --- 1. Pre-populate cache with rCachedWithRoots1 ---
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, fetchFrom, primitives.Finalized, queryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rCachedWithRoots1}, nil).Once()
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, fetchFrom, primitives.Unconfirmed, queryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rCachedWithRoots1}, nil).Once()
	_, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)
	// lastQueryTimestamp is now rCachedWithRoots1.Timestamp
	assert.Equal(t, rCachedWithRoots1.Timestamp, cache.lastQueryTimestamp)

	// --- 2. Test call that should trigger early exit ---
	// Finalized reports fetch will bring in rFinalizedRootless1 and rFinalizedWithRoots2.
	// Query from rCachedWithRoots1.Timestamp.
	mockReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything, rCachedWithRoots1.Timestamp, primitives.Finalized, queryLimit,
		).
		Return([]ccipocr3.CommitPluginReportWithMeta{rFinalizedRootless1, rFinalizedWithRoots2}, nil).Once()

	// IMPORTANT: Unconfirmed reader should NOT be called due to early exit.

	reports, err := cache.GetCachedAndNewReports(context.Background(), fetchFrom)
	assert.NoError(t, err)

	// After new finalized reports are processed:
	// - rFinalizedRootless1 is skipped by addReportsToCache.
	// - rFinalizedWithRoots2 is added to cache.
	// Cache now effectively contains: {rCachedWithRoots1, rFinalizedWithRoots2}
	// Sorted by time (newest first for early exit check): {rFinalizedWithRoots2, rCachedWithRoots1}
	// Since len (2) >= returnLimit (2), we should get these newest 2 with roots.

	assert.Len(t, reports, returnLimit)
	assert.Equal(t, rFinalizedWithRoots2, reports[0]) // Newest with roots
	assert.Equal(t, rCachedWithRoots1, reports[1])    // Second newest with roots

	// lastQueryTimestamp should update to rFinalizedWithRoots2.Timestamp, as rFinalizedRootless1 was skipped
	assert.Equal(t, rFinalizedWithRoots2.Timestamp, cache.lastQueryTimestamp)

	mockReader.AssertExpectations(t)
}
