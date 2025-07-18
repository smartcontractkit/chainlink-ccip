package cache

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	readerMocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
)

// mockTimeProvider allows controlling time in tests
type mockTimeProvider struct {
	currentTime time.Time
	mu          sync.RWMutex
}

func newMockTimeProvider(t time.Time) *mockTimeProvider {
	return &mockTimeProvider{currentTime: t}
}

func (m *mockTimeProvider) Now() time.Time {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.currentTime
}

func (m *mockTimeProvider) SetTime(t time.Time) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentTime = t
}

func (m *mockTimeProvider) Advance(d time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentTime = m.currentTime.Add(d)
}

func TestCommitReportCache_NewCommitReportCache_DefaultLookback(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	tp := newMockTimeProvider(time.Now())
	mockRdr := readerMocks.NewMockCCIPReader(t)

	visInterval := 8 * time.Hour
	cleanupInterval := 30 * time.Minute

	// Test case 1: LookbackGracePeriod = 0, defaults to fixed 30 minutes
	cfg1 := CommitReportCacheConfig{
		MessageVisibilityInterval: visInterval,
		EvictionGracePeriod:       1 * time.Hour,
		CleanupInterval:           cleanupInterval,
		LookbackGracePeriod:       0, // Test this
	}
	cache1 := NewCommitReportCache(lggr, cfg1, tp, mockRdr).(*commitReportCache)
	expectedLookback := 30 * time.Minute
	assert.Equal(t, expectedLookback, cache1.cfg.LookbackGracePeriod, "Test Case 1 Failed")

	// Test case 2: LookbackGracePeriod = 0, same fixed default used regardless of visInterval
	shortVisInterval := 30 * time.Minute
	cfg2 := CommitReportCacheConfig{
		MessageVisibilityInterval: shortVisInterval,
		EvictionGracePeriod:       1 * time.Hour,
		CleanupInterval:           8 * time.Minute,
		LookbackGracePeriod:       0,
	}
	cache2 := NewCommitReportCache(lggr, cfg2, tp, mockRdr).(*commitReportCache)
	assert.Equal(t, expectedLookback, cache2.cfg.LookbackGracePeriod, "Test Case 2 Failed")

	// Test case 3: LookbackGracePeriod = 0, same fixed default even for very short intervals
	veryShortVisInterval := 1 * time.Minute
	shortCleanupInterval := 2 * time.Minute
	cfg3 := CommitReportCacheConfig{
		MessageVisibilityInterval: veryShortVisInterval,
		EvictionGracePeriod:       1 * time.Hour,
		CleanupInterval:           shortCleanupInterval,
		LookbackGracePeriod:       0,
	}
	cache3 := NewCommitReportCache(lggr, cfg3, tp, mockRdr).(*commitReportCache)
	assert.Equal(t, expectedLookback, cache3.cfg.LookbackGracePeriod, "Test Case 3 Failed")

	// Test case 4: LookbackGracePeriod is already set
	presetLookback := 15 * time.Minute
	cfg4 := CommitReportCacheConfig{
		MessageVisibilityInterval: visInterval,
		EvictionGracePeriod:       1 * time.Hour,
		CleanupInterval:           cleanupInterval,
		LookbackGracePeriod:       presetLookback,
	}
	cache4 := NewCommitReportCache(lggr, cfg4, tp, mockRdr).(*commitReportCache)
	assert.Equal(t, presetLookback, cache4.cfg.LookbackGracePeriod, "Test Case 4 Failed")
}

func TestCommitReportCache_RefreshCache_StoresOnlyReportsWithMerkleRoots(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	tp := newMockTimeProvider(time.Now().UTC())
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cfg := CommitReportCacheConfig{
		MessageVisibilityInterval: 1 * time.Hour,
		EvictionGracePeriod:       10 * time.Minute,
		CleanupInterval:           5 * time.Minute,
		LookbackGracePeriod:       5 * time.Minute,
	}
	cache := NewCommitReportCache(lggr, cfg, tp, mockRdr).(*commitReportCache)

	reportWithBlessedRoots := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}},
			},
		},
		Timestamp: tp.Now().Add(-10 * time.Minute), BlockNum: 1,
	}
	reportWithUnblessedRoots := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{2}},
			},
		},
		Timestamp: tp.Now().Add(-9 * time.Minute), BlockNum: 2,
	}
	reportWithoutRoots := ccipocr3.CommitPluginReportWithMeta{
		Report:    ccipocr3.CommitPluginReport{},
		Timestamp: tp.Now().Add(-8 * time.Minute), BlockNum: 3,
	}

	reportsToReturn := []ccipocr3.CommitPluginReportWithMeta{
		reportWithBlessedRoots, reportWithoutRoots, reportWithUnblessedRoots,
	}
	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return(reportsToReturn, nil).Once()

	err := cache.RefreshCache(context.Background())
	require.NoError(t, err)

	assert.Equal(t, 2, cache.reportsCache.ItemCount(), "Should only store 2 reports with roots")

	cachedReports := cache.GetCachedReports(tp.Now().Add(-1 * time.Hour))
	assert.Len(t, cachedReports, 2)
	expectedReports := []ccipocr3.CommitPluginReportWithMeta{reportWithBlessedRoots, reportWithUnblessedRoots}
	sort.Slice(expectedReports, func(i, j int) bool {
		if expectedReports[i].Timestamp.Equal(expectedReports[j].Timestamp) {
			return expectedReports[i].BlockNum < expectedReports[j].BlockNum
		}
		return expectedReports[i].Timestamp.Before(expectedReports[j].Timestamp)
	})
	assert.Equal(t, expectedReports, cachedReports)
	mockRdr.AssertExpectations(t)
}

func TestCommitReportCache_RefreshCache_ConvertsLatestReportTimestampToUTC(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	fixedTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	tp := newMockTimeProvider(fixedTime)
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cacheCfg := CommitReportCacheConfig{
		MessageVisibilityInterval: 1 * time.Hour,
		LookbackGracePeriod:       5 * time.Minute,
	}
	cache := NewCommitReportCache(lggr, cacheCfg, tp, mockRdr).(*commitReportCache)

	cest, err := time.LoadLocation("Europe/Berlin")
	require.NoError(t, err)
	reportTimeNonUTC := time.Date(2023, 1, 1, 13, 0, 0, 0, cest)

	reportsToReturn := []ccipocr3.CommitPluginReportWithMeta{
		{
			Report: ccipocr3.CommitPluginReport{
				BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
					{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}},
				},
			},
			Timestamp: reportTimeNonUTC, BlockNum: 1,
		},
	}
	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return(reportsToReturn, nil).Once()

	err = cache.RefreshCache(context.Background())
	require.NoError(t, err)

	assert.Equal(t, time.UTC, cache.latestFinalizedReportTimestamp.Location())
	assert.True(t, cache.latestFinalizedReportTimestamp.Equal(reportTimeNonUTC.UTC()))
	mockRdr.AssertExpectations(t)
}

func TestCommitReportCache_RefreshCache_UpdatesLatestReportTimestampToMostRecent(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	baseTime := time.Now().UTC()
	tp := newMockTimeProvider(baseTime)
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cacheCfg := CommitReportCacheConfig{
		MessageVisibilityInterval: 1 * time.Hour,
		LookbackGracePeriod:       5 * time.Minute,
	}
	cache := NewCommitReportCache(lggr, cacheCfg, tp, mockRdr).(*commitReportCache)

	r1 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}},
			},
		},
		Timestamp: baseTime.Add(-20 * time.Minute), BlockNum: 1,
	}
	r2 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{2}},
			},
		},
		Timestamp: baseTime.Add(-10 * time.Minute), BlockNum: 2, // Most recent
	}
	r3 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{3}},
			},
		},
		Timestamp: baseTime.Add(-15 * time.Minute), BlockNum: 3,
	}

	reportsToReturn := []ccipocr3.CommitPluginReportWithMeta{r1, r2, r3}
	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return(reportsToReturn, nil).Once()

	err := cache.RefreshCache(context.Background())
	require.NoError(t, err)
	assert.True(t, r2.Timestamp.UTC().Equal(cache.latestFinalizedReportTimestamp))
	mockRdr.AssertExpectations(t)
}

func TestGenerateKey(t *testing.T) {
	t.Parallel()
	ts := time.Now()
	reportBlessed := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{10}}},
		},
		Timestamp: ts, BlockNum: 1,
	}
	expectedKeyBlessed := fmt.Sprintf("%d_%s", uint64(1), reportBlessed.Report.BlessedMerkleRoots[0].MerkleRoot)
	key, err := generateKey(reportBlessed)
	require.NoError(t, err)
	assert.Equal(t, expectedKeyBlessed, key)

	reportUnblessed := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{20}}},
		},
		Timestamp: ts, BlockNum: 2,
	}
	expectedKeyUnblessed := fmt.Sprintf("%d_%s", uint64(2), reportUnblessed.Report.UnblessedMerkleRoots[0].MerkleRoot)
	key, err = generateKey(reportUnblessed)
	require.NoError(t, err)
	assert.Equal(t, expectedKeyUnblessed, key)

	reportBoth := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots:   []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{10}}},
			UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{20}}},
		},
		Timestamp: ts, BlockNum: 3,
	}
	key, err = generateKey(reportBoth)
	require.NoError(t, err)
	assert.Equal(t, expectedKeyBlessed, key, "Should prioritize blessed root")

	reportNone := ccipocr3.CommitPluginReportWithMeta{Report: ccipocr3.CommitPluginReport{}, Timestamp: ts, BlockNum: 4}
	key, err = generateKey(reportNone)
	require.Error(t, err, "Should return error for report with no roots")
	assert.Empty(t, key, "Key should be empty when there are no roots")
}

func TestCommitReportCache_GetReportsToQueryFromTimestamp_EmptyCache(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	visInterval := 2 * time.Hour
	now := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	tp := newMockTimeProvider(now)
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cacheCfg := CommitReportCacheConfig{
		MessageVisibilityInterval: visInterval,
		LookbackGracePeriod:       15 * time.Minute,
	}
	cache := NewCommitReportCache(lggr, cacheCfg, tp, mockRdr)

	expectedQueryTs := now.Add(-visInterval)
	assert.Equal(t, expectedQueryTs, cache.GetReportsToQueryFromTimestamp())
}

func TestCommitReportCache_GetReportsToQueryFromTimestamp_LookbackApplied(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	visInterval := 2 * time.Hour
	lookback := 30 * time.Minute
	now := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	tp := newMockTimeProvider(now)
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cacheCfg := CommitReportCacheConfig{
		MessageVisibilityInterval: visInterval,
		LookbackGracePeriod:       lookback,
	}
	cache := NewCommitReportCache(lggr, cacheCfg, tp, mockRdr).(*commitReportCache)

	cache.latestFinalizedReportTimestamp = now.Add(-1 * time.Hour)

	expectedQueryTs := now.Add(-1 * time.Hour).Add(-lookback)
	assert.True(t, expectedQueryTs.Equal(cache.GetReportsToQueryFromTimestamp()))
}

func TestCommitReportCache_GetReportsToQueryFromTimestamp_VisibilityWindowDominates(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	visInterval := 2 * time.Hour
	lookback := 30 * time.Minute
	now := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	tp := newMockTimeProvider(now)
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cacheCfg := CommitReportCacheConfig{
		MessageVisibilityInterval: visInterval,
		LookbackGracePeriod:       lookback,
	}
	cache := NewCommitReportCache(lggr, cacheCfg, tp, mockRdr).(*commitReportCache)

	cache.latestFinalizedReportTimestamp = now.Add(-3 * time.Hour)

	expectedQueryTs := now.Add(-visInterval)
	assert.True(t, expectedQueryTs.Equal(cache.GetReportsToQueryFromTimestamp()))
}

func TestCommitReportCache_GetCachedReports_FiltersAndSorts(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	now := time.Now().UTC()
	tp := newMockTimeProvider(now)
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cacheCfg := CommitReportCacheConfig{
		MessageVisibilityInterval: 1 * time.Hour,
		LookbackGracePeriod:       5 * time.Minute,
	}
	cache := NewCommitReportCache(lggr, cacheCfg, tp, mockRdr)

	r1 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}},
			},
		},
		Timestamp: now.Add(-30 * time.Minute), BlockNum: 100,
	}
	r2SameTsBlockGreater := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{2}},
			},
		},
		Timestamp: now.Add(-20 * time.Minute), BlockNum: 201,
	}
	r3 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{3}},
			},
		},
		Timestamp: now.Add(-10 * time.Minute), BlockNum: 300,
	}
	r4TooOld := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{4}},
			},
		},
		Timestamp: now.Add(-40 * time.Minute), BlockNum: 50,
	}
	r5SameTsBlockSmaller := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{5}},
			},
		},
		Timestamp: now.Add(-20 * time.Minute), BlockNum: 200,
	}

	reportsToReturn := []ccipocr3.CommitPluginReportWithMeta{r3, r1, r2SameTsBlockGreater, r4TooOld, r5SameTsBlockSmaller}
	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return(
		func(ctx context.Context, ts time.Time, conf primitives.ConfidenceLevel, limit int) (
			[]ccipocr3.CommitPluginReportWithMeta, error,
		) {
			var filtered []ccipocr3.CommitPluginReportWithMeta
			for _, r := range reportsToReturn {
				if r.Timestamp.After(ts) || r.Timestamp.Equal(ts) {
					filtered = append(filtered, r)
				}
			}
			return filtered, nil
		},
	).Once()

	err := cache.RefreshCache(context.Background())
	require.NoError(t, err)

	fromTs := now.Add(-35 * time.Minute)
	cached := cache.GetCachedReports(fromTs)

	expected := []ccipocr3.CommitPluginReportWithMeta{r1, r5SameTsBlockSmaller, r2SameTsBlockGreater, r3}
	assert.Equal(t, expected, cached)
	assert.Len(t, cached, 4)
	mockRdr.AssertExpectations(t)
}

func TestCommitReportCache_GetCachedReports_ExpiredItemsIgnored(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	now := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	tp := newMockTimeProvider(now)
	mockRdr := readerMocks.NewMockCCIPReader(t)

	visInterval := 30 * time.Minute
	evictionGrace := 5 * time.Minute

	cfg := CommitReportCacheConfig{
		MessageVisibilityInterval: visInterval,
		EvictionGracePeriod:       evictionGrace,
		CleanupInterval:           1 * time.Hour,
		LookbackGracePeriod:       5 * time.Minute,
	}
	cacheInstance := NewCommitReportCache(lggr, cfg, tp, mockRdr).(*commitReportCache)

	reportToStayKey := "stay_key"
	reportToStay := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{2}}},
		},
		Timestamp: now, BlockNum: 2,
	}
	cacheInstance.reportsCache.SetDefault(reportToStayKey, reportToStay)

	reportToExpireKey := "expire_key"
	reportToExpire := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}}},
		},
		Timestamp: now.Add(-1 * time.Minute), BlockNum: 1,
	}
	cacheInstance.reportsCache.Set(reportToExpireKey, reportToExpire, 1*time.Millisecond)

	time.Sleep(5 * time.Millisecond)

	cached := cacheInstance.GetCachedReports(now.Add(-2 * time.Hour))

	assert.Len(t, cached, 1, "Only one report should remain after specific item expiry")
	if len(cached) == 1 {
		found := false
		for _, r := range cached {
			if r.BlockNum == reportToStay.BlockNum {
				found = true
				assert.Equal(t, reportToStay, r)
				break
			}
		}
		assert.True(t, found, "reportToStay was not found in cached reports")
	}
}

func TestCommitReportCache_RefreshCache_ConcurrentCalls(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	now := time.Now().UTC()
	tp := newMockTimeProvider(now)
	mockRdr := readerMocks.NewMockCCIPReader(t)

	cfg := CommitReportCacheConfig{
		MessageVisibilityInterval: 1 * time.Hour,
		LookbackGracePeriod:       5 * time.Minute,
	}
	cache := NewCommitReportCache(lggr, cfg, tp, mockRdr)

	numGoroutines := 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	var initialReports []ccipocr3.CommitPluginReportWithMeta
	for i := 0; i < 5; i++ {
		initialReports = append(initialReports, ccipocr3.CommitPluginReportWithMeta{
			Report: ccipocr3.CommitPluginReport{
				BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
					{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{byte(i + 100)}},
				},
			},
			Timestamp: now.Add(time.Duration(i+100) * time.Minute), BlockNum: uint64(i + 100),
		})
	}

	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return(initialReports, nil)

	for i := 0; i < numGoroutines; i++ {
		go func(_ int) {
			defer wg.Done()
			err := cache.RefreshCache(context.Background())
			assert.NoError(t, err)
		}(i)
	}
	wg.Wait()
	internalCache := cache.(*commitReportCache)
	assert.LessOrEqual(t, internalCache.reportsCache.ItemCount(), len(initialReports))
	mockRdr.AssertExpectations(t)
}

func TestCommitReportCache_RefreshCache_RespectsReaderLimit(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	tp := newMockTimeProvider(time.Now().UTC())
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cacheCfg := CommitReportCacheConfig{MessageVisibilityInterval: 1 * time.Hour}
	cache := NewCommitReportCache(lggr, cacheCfg, tp, mockRdr)

	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return([]ccipocr3.CommitPluginReportWithMeta{}, nil).Once()

	err := cache.RefreshCache(context.Background())
	require.NoError(t, err)

	mockRdr.AssertExpectations(t)
}

func TestCommitReportCache_DeduplicateReports(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	now := time.Now().UTC()
	// tp and mockRdr are not used by the package-level DeduplicateReports function,
	// but are kept here for consistency with other tests in case they are needed later.
	newMockTimeProvider(now)
	readerMocks.NewMockCCIPReader(t)

	// Create test reports with different configurations

	// Report with blessed root 1
	r1 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}},
			},
		},
		Timestamp: now.Add(-30 * time.Minute),
		BlockNum:  100,
	}

	// Report with blessed root 2
	r2 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{2}},
			},
		},
		Timestamp: now.Add(-25 * time.Minute),
		BlockNum:  200,
	}

	// Duplicate of r1 (same root) but with different timestamp and block
	r3Dupe := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}},
			},
		},
		Timestamp: now.Add(-20 * time.Minute),
		BlockNum:  300,
	}

	// Report with unblessed root
	r4 := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{3}},
			},
		},
		Timestamp: now.Add(-15 * time.Minute),
		BlockNum:  400,
	}

	// No merkle roots - should be skipped
	r5NoRoots := ccipocr3.CommitPluginReportWithMeta{
		Report:    ccipocr3.CommitPluginReport{},
		Timestamp: now.Add(-10 * time.Minute),
		BlockNum:  500,
	}

	// Duplicate of r4 (same chain selector and root)
	r6Dupe := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{3}},
			},
		},
		Timestamp: now.Add(-5 * time.Minute),
		BlockNum:  600,
	}

	// Different chain selector but same root value as r1
	r7DiffChain := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
				{ChainSel: 3, MerkleRoot: ccipocr3.Bytes32{1}}, // Same root value but different chain
			},
		},
		Timestamp: now,
		BlockNum:  700,
	}

	// Test cases
	testCases := []struct {
		name     string
		reports  []ccipocr3.CommitPluginReportWithMeta
		expected []ccipocr3.CommitPluginReportWithMeta
	}{
		{
			name:     "Empty input",
			reports:  []ccipocr3.CommitPluginReportWithMeta{},
			expected: []ccipocr3.CommitPluginReportWithMeta{},
		},
		{
			name:     "Single report",
			reports:  []ccipocr3.CommitPluginReportWithMeta{r1},
			expected: []ccipocr3.CommitPluginReportWithMeta{r1},
		},
		{
			name:     "No duplicates",
			reports:  []ccipocr3.CommitPluginReportWithMeta{r1, r2},
			expected: []ccipocr3.CommitPluginReportWithMeta{r1, r2},
		},
		{
			name:     "With duplicates",
			reports:  []ccipocr3.CommitPluginReportWithMeta{r1, r2, r3Dupe},
			expected: []ccipocr3.CommitPluginReportWithMeta{r1, r2}, // r3Dupe should be filtered out
		},
		{
			name:     "Mix of blessed and unblessed roots",
			reports:  []ccipocr3.CommitPluginReportWithMeta{r1, r4},
			expected: []ccipocr3.CommitPluginReportWithMeta{r1, r4},
		},
		{
			name:     "Report without roots skipped",
			reports:  []ccipocr3.CommitPluginReportWithMeta{r1, r5NoRoots},
			expected: []ccipocr3.CommitPluginReportWithMeta{r1}, // r5 should be filtered out
		},
		{
			name:    "Multiple dupes with different chain selectors",
			reports: []ccipocr3.CommitPluginReportWithMeta{r1, r2, r3Dupe, r4, r5NoRoots, r6Dupe, r7DiffChain},
			expected: []ccipocr3.CommitPluginReportWithMeta{
				r1, r2, r4, r7DiffChain,
			}, // r3Dupe, r5NoRoots, r6Dupe should be filtered out
		},
		{
			name:     "Multiple reports without roots",
			reports:  []ccipocr3.CommitPluginReportWithMeta{r5NoRoots, r5NoRoots, r5NoRoots},
			expected: []ccipocr3.CommitPluginReportWithMeta{}, // All reports should be skipped
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DeduplicateReports(lggr, tc.reports)

			// Verify length
			assert.Equal(t, len(tc.expected), len(result), "Result length mismatch")

			// Convert to map by key for easier comparison
			resultKeys := make(map[string]bool)
			for _, r := range result {
				key, err := generateKey(r)
				require.NoError(t, err)
				resultKeys[key] = true
			}

			// Verify all expected reports are present by checking their keys
			for _, r := range tc.expected {
				key, err := generateKey(r)
				require.NoError(t, err)
				assert.True(t, resultKeys[key], "Expected report with key %s not found in result", key)
			}

			// Verify first occurrence is kept when duplicates exist
			if tc.name == "With duplicates" {
				for _, r := range result {
					key, _ := generateKey(r)
					if key == "1_0100000000000000000000000000000000000000000000000000000000000000" {
						assert.Equal(t, r1.BlockNum, r.BlockNum, "First occurrence should be kept")
					}
				}
			}
		})
	}
}

func TestCommitReportCache_RefreshCache_LatestTimestampAdvancesWithRootlessReports(t *testing.T) {
	t.Parallel()
	lggr := logger.Test(t)
	baseTime := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	tp := newMockTimeProvider(baseTime)
	mockRdr := readerMocks.NewMockCCIPReader(t)
	cfg := CommitReportCacheConfig{
		MessageVisibilityInterval: 1 * time.Hour,
		EvictionGracePeriod:       10 * time.Minute,
		CleanupInterval:           5 * time.Minute,
		LookbackGracePeriod:       5 * time.Minute,
	}
	cache := NewCommitReportCache(lggr, cfg, tp, mockRdr).(*commitReportCache)

	// Initial state: no reports, latest timestamp is zero
	require.True(t, cache.latestFinalizedReportTimestamp.IsZero())

	// --- First Refresh: Report with no roots but a newer timestamp ---
	rootlessReportTime := baseTime.Add(10 * time.Minute)
	reportWithoutRoots := ccipocr3.CommitPluginReportWithMeta{
		Report:    ccipocr3.CommitPluginReport{ /* No roots */ },
		Timestamp: rootlessReportTime,
		BlockNum:  100,
	}
	reportWithOldRoot := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{1}}},
		},
		Timestamp: baseTime.Add(-5 * time.Minute), // Older than rootlessReportTime
		BlockNum:  99,
	}

	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return([]ccipocr3.CommitPluginReportWithMeta{reportWithOldRoot, reportWithoutRoots}, nil).Once()

	err := cache.RefreshCache(context.Background())
	require.NoError(t, err)

	// Assert: latestFinalizedReportTimestamp should be updated to the timestamp of the rootless report
	// because it was the latest processed by the reader.
	assert.True(t, cache.latestFinalizedReportTimestamp.Equal(rootlessReportTime.UTC()))
	// Assert: cache should still be empty or contain only reportWithOldRoot, as reportWithoutRoots has no key
	assert.Equal(t, 1, cache.reportsCache.ItemCount())
	keyOldRoot, _ := generateKey(reportWithOldRoot)
	_, foundOld := cache.reportsCache.Get(keyOldRoot)
	assert.True(t, foundOld, "reportWithOldRoot should be in cache")

	mockRdr.AssertExpectations(t) // Ensures the first call was made

	// --- Second Refresh: Report with roots and an even newer timestamp ---
	// Reset the mock for a new set of expectations for the second call
	mockRdr = readerMocks.NewMockCCIPReader(t)         // New mock to avoid issues with previous .Once()
	cache.reader = mockRdr                             // Update cache to use the new mock
	cache.latestFinalizedReportTimestamp = time.Time{} // Reset for clarity, though After check would handle it
	cache.reportsCache.Flush()                         // Clear cache for this part of test
	tp.SetTime(baseTime.Add(20 * time.Minute))         // Advance time

	newReportTime := baseTime.Add(15 * time.Minute) // Newer than rootlessReportTime
	reportWithNewRoot := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{2}}},
		},
		Timestamp: newReportTime,
		BlockNum:  101,
	}

	mockRdr.On("CommitReportsGTETimestamp",
		mock.Anything, mock.AnythingOfType("time.Time"), primitives.Finalized, DefaultMaxCommitReportsToFetch,
	).Return([]ccipocr3.CommitPluginReportWithMeta{reportWithNewRoot}, nil).Once()

	err = cache.RefreshCache(context.Background())
	require.NoError(t, err)

	// Assert: latestFinalizedReportTimestamp should be updated to newReportTime
	assert.True(t, cache.latestFinalizedReportTimestamp.Equal(newReportTime.UTC()))
	// Assert: reportWithNewRoot should be in the cache
	assert.Equal(t, 1, cache.reportsCache.ItemCount())
	keyNewRoot, _ := generateKey(reportWithNewRoot)
	_, foundNew := cache.reportsCache.Get(keyNewRoot)
	assert.True(t, foundNew, "reportWithNewRoot should be in cache")
	mockRdr.AssertExpectations(t)
}
