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

	readerMocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

	// Test case 1: LookbackGracePeriod = 0, defaults to visInterval / 8
	cfg1 := CommitReportCacheConfig{
		MessageVisibilityInterval: visInterval,
		EvictionGracePeriod:       1 * time.Hour,
		CleanupInterval:           cleanupInterval,
		LookbackGracePeriod:       0, // Test this
	}
	cache1 := NewCommitReportCache(lggr, cfg1, tp, mockRdr).(*commitReportCache)
	expectedLookback1 := visInterval / 8
	assert.Equal(t, expectedLookback1, cache1.cfg.LookbackGracePeriod, "Test Case 1 Failed")

	// Test case 2: visInterval / 8 is less than 1 minute, cleanupInterval / 4 is used
	shortVisInterval := 30 * time.Minute // -> visInterval/8 = 3.75 min
	cfg2 := CommitReportCacheConfig{
		MessageVisibilityInterval: shortVisInterval,
		EvictionGracePeriod:       1 * time.Hour,
		CleanupInterval:           8 * time.Minute, // -> cleanupInterval/4 = 2 min
		LookbackGracePeriod:       0,
	}
	cache2 := NewCommitReportCache(lggr, cfg2, tp, mockRdr).(*commitReportCache)
	//expectedLookback2 := cfg2.CleanupInterval / 4  // Original logic had this if defaultLookback < 1min
	// Current logic: defaultLookback = shortVisInterval / 8 = 3.75min. This is > 1min.
	expectedLookback2 := shortVisInterval / 8
	assert.Equal(t, expectedLookback2, cache2.cfg.LookbackGracePeriod, "Test Case 2 Failed")

	// Test case 3: visInterval / 8 is very small, cleanupInterval / 4 is also small, defaults to 1 minute
	veryShortVisInterval := 1 * time.Minute // -> visInterval/8 = 7.5s
	shortCleanupInterval := 2 * time.Minute // -> cleanupInterval/4 = 30s
	cfg3 := CommitReportCacheConfig{
		MessageVisibilityInterval: veryShortVisInterval,
		EvictionGracePeriod:       1 * time.Hour,
		CleanupInterval:           shortCleanupInterval,
		LookbackGracePeriod:       0,
	}
	cache3 := NewCommitReportCache(lggr, cfg3, tp, mockRdr).(*commitReportCache)
	// defaultLookback = 7.5s. 7.5s < 1min. cleanupInterval > 0. defaultLookback becomes cleanupInterval / 4 = 30s.
	// 30s < 1min. defaultLookback becomes 1min.
	assert.Equal(t, 1*time.Minute, cache3.cfg.LookbackGracePeriod, "Test Case 3 Failed")

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

	assert.Equal(t, time.UTC, cache.latestReportTimestamp.Location())
	assert.True(t, cache.latestReportTimestamp.Equal(reportTimeNonUTC.UTC()))
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
	assert.True(t, r2.Timestamp.UTC().Equal(cache.latestReportTimestamp))
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
	expectedKeyBlessed := fmt.Sprintf("%s_%s", ccipocr3.ChainSelector(1).String(), ccipocr3.Bytes32{10}.String())
	assert.Equal(t, expectedKeyBlessed, generateKey(reportBlessed))

	reportUnblessed := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{20}}},
		},
		Timestamp: ts, BlockNum: 2,
	}
	expectedKeyUnblessed := fmt.Sprintf("%s_%s", ccipocr3.ChainSelector(2).String(), ccipocr3.Bytes32{20}.String())
	assert.Equal(t, expectedKeyUnblessed, generateKey(reportUnblessed))

	reportBoth := ccipocr3.CommitPluginReportWithMeta{
		Report: ccipocr3.CommitPluginReport{
			BlessedMerkleRoots:   []ccipocr3.MerkleRootChain{{ChainSel: 1, MerkleRoot: ccipocr3.Bytes32{10}}},
			UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{{ChainSel: 2, MerkleRoot: ccipocr3.Bytes32{20}}},
		},
		Timestamp: ts, BlockNum: 3,
	}
	assert.Equal(t, expectedKeyBlessed, generateKey(reportBoth), "Should prioritize blessed root")

	reportNone := ccipocr3.CommitPluginReportWithMeta{Report: ccipocr3.CommitPluginReport{}, Timestamp: ts, BlockNum: 4}
	expectedKeyNone := fmt.Sprintf("%d_%d", ts.Unix(), uint64(4))
	assert.Equal(t, expectedKeyNone, generateKey(reportNone))
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

	cache.latestReportTimestamp = now.Add(-1 * time.Hour)

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

	cache.latestReportTimestamp = now.Add(-3 * time.Hour)

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
