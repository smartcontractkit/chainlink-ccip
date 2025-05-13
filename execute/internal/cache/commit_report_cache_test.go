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

// Test using the generated testify-mock CCIPReader rather than a hand-rolled fake.

func TestCommitReportsCache_GetCachedAndNewReports(t *testing.T) {
	lggr := logger.Test(t)

	now := time.Now().UTC()
	fetchFrom := now.Add(-10 * time.Minute)
	r1 := ccipocr3.CommitPluginReportWithMeta{Timestamp: now.Add(-8 * time.Minute)}
	r2 := ccipocr3.CommitPluginReportWithMeta{Timestamp: now.Add(-5 * time.Minute)}

	mockReader := readerpkg_mock.NewMockCCIPReader(t)

	// We capture the timestamp parameters to ensure they advance between calls.
	var tsFinalizedCalls []time.Time

	// First round expectations --------------------------------------
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, 10).
		RunAndReturn(func(
			_ context.Context,
			ts time.Time,
			_ primitives.ConfidenceLevel,
			_ int) ([]ccipocr3.CommitPluginReportWithMeta, error) {
			tsFinalizedCalls = append(tsFinalizedCalls, ts)
			return []ccipocr3.CommitPluginReportWithMeta{r1}, nil
		})

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.AnythingOfType("time.Time"), primitives.Unconfirmed, 10).
		Return([]ccipocr3.CommitPluginReportWithMeta{r1, r2}, nil)

	// Second round expectations -------------------------------------
	// finalized query returns nothing new but we record the ts to ensure advancement
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, 10).
		RunAndReturn(func(_ context.Context, ts time.Time, _ primitives.ConfidenceLevel, _ int) (
			[]ccipocr3.CommitPluginReportWithMeta,
			error) {
			tsFinalizedCalls = append(tsFinalizedCalls, ts)
			return nil, nil
		})

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.AnythingOfType("time.Time"), primitives.Unconfirmed, 10).
		Return([]ccipocr3.CommitPluginReportWithMeta{r1, r2}, nil)

	cache := NewCommitReportsCache(lggr, 2*time.Hour)

	// First invocation
	got, err := cache.GetCachedAndNewReports(context.Background(), mockReader, fetchFrom, 10)
	assert.NoError(t, err)
	assert.Len(t, got, 2)

	// Second invocation (no new finalized data)
	got2, err := cache.GetCachedAndNewReports(context.Background(), mockReader, fetchFrom, 10)
	assert.NoError(t, err)
	assert.Len(t, got2, 2)

	// Validate timestamp advancement
	assert.Len(t, tsFinalizedCalls, 2)
	firstTs := tsFinalizedCalls[0]
	secondTs := tsFinalizedCalls[1]
	assert.Equal(t, fetchFrom, firstTs)
	assert.True(t, secondTs.After(firstTs))
}

func TestCommitReportsCache_LimitAndDeduplication(t *testing.T) {
	lggr := logger.Test(t)

	now := time.Now().UTC()
	// Create 5 reports with unique timestamps (descending)
	reports := make([]ccipocr3.CommitPluginReportWithMeta, 5)
	for i := 0; i < 5; i++ {
		reports[i] = ccipocr3.CommitPluginReportWithMeta{Timestamp: now.Add(-time.Duration(i) * time.Minute)}
	}

	mockReader := readerpkg_mock.NewMockCCIPReader(t)

	// First call: finalized returns first 3, unconfirmed returns all 5.
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, 3).
		Return(reports[:3], nil)

	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.AnythingOfType("time.Time"), primitives.Unconfirmed, 3).
		Return(reports, nil)

	cache := NewCommitReportsCache(lggr, 1*time.Hour)
	// limit = 3 should truncate to newest 3 (those with smallest i)
	got, err := cache.GetCachedAndNewReports(context.Background(), mockReader, now.Add(-10*time.Minute), 3)
	assert.NoError(t, err)
	assert.Len(t, got, 3)
	// Ensure dedup preserved order newest first
	assert.True(t, got[0].Timestamp.After(got[1].Timestamp))
	assert.True(t, got[1].Timestamp.After(got[2].Timestamp))

	// Call again with empty reader responses to verify cached dedup still applies and limit holds
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.Anything, primitives.Finalized, 3).
		Return(nil, nil)
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.Anything, primitives.Unconfirmed, 3).
		Return(reports, nil)

	got2, err := cache.GetCachedAndNewReports(context.Background(), mockReader, now.Add(-10*time.Minute), 3)
	assert.NoError(t, err)
	assert.Len(t, got2, 3)
}

func TestCommitReportsCache_DedupAndFetchWindow(t *testing.T) {
	lggr := logger.Test(t)

	now := time.Now().UTC()
	// This is the initial fetchFrom for the first call to GetCachedAndNewReports
	initialFetchFrom := now.Add(-10 * time.Minute)
	// This is the report that will be fetched and cached
	r := ccipocr3.CommitPluginReportWithMeta{Timestamp: now.Add(-2 * time.Minute)}

	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	cache := NewCommitReportsCache(lggr, 30*time.Minute)

	// --- First invocation ---
	// Purpose: Populate the cache with report 'r'.
	// Expected queryFrom for reader: initialFetchFrom
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, initialFetchFrom, primitives.Finalized, 5).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once()
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, initialFetchFrom, primitives.Unconfirmed, 5).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once()

	got1, err1 := cache.GetCachedAndNewReports(context.Background(), mockReader, initialFetchFrom, 5)
	assert.NoError(t, err1)
	assert.Len(t, got1, 1, "First call should return the single fetched report")
	// After this call, cache.lastQueryTimestamp should be r.Timestamp (now - 2m)

	// --- Second invocation ---
	// Purpose: Test deduplication. Fetch 'r' again; should still get only one 'r'.
	// Expected queryFrom for reader: r.Timestamp (because lastQueryTimestamp was updated)
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, r.Timestamp, primitives.Finalized, 5).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once() // Reader returns 'r' again
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, r.Timestamp, primitives.Unconfirmed, 5).
		Return([]ccipocr3.CommitPluginReportWithMeta{r}, nil).Once() // Reader returns 'r' again

	// The 'fetchFrom' for GetCachedAndNewReports itself is still initialFetchFrom for this logical step
	got2, err2 := cache.GetCachedAndNewReports(context.Background(), mockReader, initialFetchFrom, 5)
	assert.NoError(t, err2)
	assert.Len(t, got2, 1, "Second call should return the single report due to deduplication")
	// After this, cache.lastQueryTimestamp should remain r.Timestamp

	// --- Third invocation ---
	// Purpose: Test that fetching with a window newer than the cached item yields nothing.
	// The 'fetchFrom' for GetCachedAndNewReports is now newer than 'r.Timestamp'.
	newerFetchWindowStart := now.Add(-1 * time.Minute)
	// Expected queryFrom for reader: r.Timestamp (lastQueryTimestamp hasn't changed)
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, r.Timestamp, primitives.Finalized, 5).
		Return(nil, nil).Once() // Reader finds nothing new from r.Timestamp
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, r.Timestamp, primitives.Unconfirmed, 5).
		Return(nil, nil).Once() // Reader finds nothing new from r.Timestamp

	got3, err3 := cache.GetCachedAndNewReports(context.Background(), mockReader, newerFetchWindowStart, 5)
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
		CommitReportsGTETimestamp(mock.Anything, mock.Anything, primitives.Finalized, mock.Anything).
		Return(nil, expectedErr).Once()

	c := NewCommitReportsCache(lggr, 1*time.Hour) // Use exported NewCommitReportsCache
	_, err := c.GetCachedAndNewReports(context.Background(), mockReader, time.Now(), 10)

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
		CommitReportsGTETimestamp(mock.Anything, mock.Anything, primitives.Finalized, mock.Anything).
		Return([]ccipocr3.CommitPluginReportWithMeta{}, nil).Once()

	// Unconfirmed fails
	mockReader.
		EXPECT().
		CommitReportsGTETimestamp(mock.Anything, mock.Anything, primitives.Unconfirmed, mock.Anything).
		Return(nil, expectedErr).Once()

	c := NewCommitReportsCache(lggr, 1*time.Hour) // Use exported NewCommitReportsCache
	_, err := c.GetCachedAndNewReports(context.Background(), mockReader, time.Now(), 10)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, expectedErr), "Error should wrap the expected reader error")
	mockReader.AssertExpectations(t)
}

func TestCommitReportsCache_MergeSortLimit(t *testing.T) {
	lggr := logger.Test(t)
	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	now := time.Now().UTC()

	// Reports to be initially in cache (simulated by first fetch)
	rOld1 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-20 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x01}}},
		},
	}
	rOld2 := ccipocr3.CommitPluginReportWithMeta{
		Timestamp: now.Add(-10 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x02}}},
		},
	}

	// Reports to be returned by the unconfirmed reader call
	rNew1 := ccipocr3.CommitPluginReportWithMeta{ // Newest
		Timestamp: now.Add(-5 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x03}}},
		},
	}
	rNew2DupOld2 := ccipocr3.CommitPluginReportWithMeta{ // Duplicate of rOld2, same timestamp
		Timestamp: now.Add(-10 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x02}}},
		},
	}
	rNew3 := ccipocr3.CommitPluginReportWithMeta{ // Older than rOld2/rNew2_DupOld2 but newer than rOld1
		Timestamp: now.Add(-15 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x04}}},
		},
	}
	rNew4 := ccipocr3.CommitPluginReportWithMeta{ // Oldest
		Timestamp: now.Add(-25 * time.Minute),
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{0x05}}},
		},
	}

	c := NewCommitReportsCache(lggr, 1*time.Hour)
	initialFetchFrom := now.Add(-30 * time.Minute)
	limit := 3

	// First call: Populate cache with rOld1, rOld2 from finalized & unconfirmed
	mockReader.EXPECT().CommitReportsGTETimestamp(mock.Anything, initialFetchFrom, primitives.Finalized, mock.Anything).
		Return([]ccipocr3.CommitPluginReportWithMeta{rOld1, rOld2}, nil).Once()
	mockReader.EXPECT().CommitReportsGTETimestamp(mock.Anything, initialFetchFrom, primitives.Unconfirmed, mock.Anything).
		Return([]ccipocr3.CommitPluginReportWithMeta{rOld1, rOld2}, nil).Once()

	// Use a larger limit for setup
	_, err := c.GetCachedAndNewReports(context.Background(), mockReader, initialFetchFrom, 10)
	assert.NoError(t, err)

	// Second call: Test merge, sort, limit
	// Cached: rOld1, rOld2. LastQueryTimestamp should be rOld2.Timestamp.
	// Unconfirmed fetch will return: rNew1, rNew2_DupOld2, rNew3, rNew4

	// Expect finalized to be called from rOld2.Timestamp, return nothing new for simplicity
	mockReader.EXPECT().CommitReportsGTETimestamp(mock.Anything, rOld2.Timestamp, primitives.Finalized, limit).
		Return([]ccipocr3.CommitPluginReportWithMeta{}, nil).Once()
	// Expect unconfirmed to be called from rOld2.Timestamp
	mockReader.EXPECT().CommitReportsGTETimestamp(mock.Anything, rOld2.Timestamp, primitives.Unconfirmed, limit).
		Return([]ccipocr3.CommitPluginReportWithMeta{rNew1, rNew2DupOld2, rNew3, rNew4}, nil).Once()

	// For this call, we are interested in reports >= initialFetchFrom. All our reports satisfy this.
	finalReports, err := c.GetCachedAndNewReports(context.Background(), mockReader, initialFetchFrom, limit)
	assert.NoError(t, err)

	// Expected unique reports before limit, sorted: rNew1, rOld2/rNew2_DupOld2, rNew3, rOld1, rNew4
	// Merged (unique by key, rNew2_DupOld2 overwrites rOld2):
	// rNew1 (ts: now-5m)
	// rNew2_DupOld2 (ts: now-10m) -> this is effectively rOld2 updated if content changed, or just rOld2
	// rNew3 (ts: now-15m)
	// rOld1 (ts: now-20m)
	// rNew4 (ts: now-25m)

	// After sorting (newest first) and limiting to 3:
	// 1. rNew1 (now-5m)
	// 2. rNew2_DupOld2 (now-10m)
	// 3. rNew3 (now-15m)

	assert.Len(t, finalReports, limit)
	assert.Equal(t, rNew1.Report.BlessedMerkleRoots[0].MerkleRoot, finalReports[0].Report.BlessedMerkleRoots[0].MerkleRoot)
	assert.Equal(t,
		rNew2DupOld2.Report.BlessedMerkleRoots[0].MerkleRoot,
		finalReports[1].Report.BlessedMerkleRoots[0].MerkleRoot,
	)
	assert.Equal(t, rNew3.Report.BlessedMerkleRoots[0].MerkleRoot, finalReports[2].Report.BlessedMerkleRoots[0].MerkleRoot)

	// Ensure timestamps are in correct descending order
	assert.True(t, finalReports[0].Timestamp.After(finalReports[1].Timestamp))
	assert.True(t, finalReports[1].Timestamp.After(finalReports[2].Timestamp))

	mockReader.AssertExpectations(t)
}
