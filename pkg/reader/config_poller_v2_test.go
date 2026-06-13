package reader

import (
	"context"
	"errors"
	"maps"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"

	mock_ccipocr3 "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/ccipocr3"
)

// Test chain selectors for config_poller_v2
const (
	destChain    cciptypes.ChainSelector = 1
	sourceChain1 cciptypes.ChainSelector = 2
	sourceChain2 cciptypes.ChainSelector = 3
	sourceChain3 cciptypes.ChainSelector = 4
	sourceChain4 cciptypes.ChainSelector = 5
	sourceChain5 cciptypes.ChainSelector = 6
)

// setupConfigPollerV2 creates a test instance of configPollerV2
func setupConfigPollerV2(t *testing.T) (*configPollerV2, map[cciptypes.ChainSelector]*mock_ccipocr3.MockChainAccessor) {
	accessors := make(map[cciptypes.ChainSelector]*mock_ccipocr3.MockChainAccessor)
	chainAccessors := make(map[cciptypes.ChainSelector]cciptypes.ChainAccessor)

	// Create accessors for test chains
	for _, chain := range []cciptypes.ChainSelector{
		destChain, sourceChain1, sourceChain2, sourceChain3, sourceChain4, sourceChain5,
	} {
		accessor := mock_ccipocr3.NewMockChainAccessor(t)
		accessors[chain] = accessor
		chainAccessors[chain] = accessor
	}

	cPollerV2 := newConfigPollerV2(
		logger.Test(t),
		chainAccessors,
		destChain,            // destination chain
		100*time.Millisecond, // short refresh period for testing
	)

	return cPollerV2, accessors
}

func TestConfigPollerV2_GetChainConfig_CacheHit(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// Setup mock for initial fetch
	expectedConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(expectedConfig, expectedSourceConfigs, nil).Once()

	// First call should fetch from chain accessor
	config1, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, config1)

	// Second call should hit cache
	config2, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, config2)

	// Verify chain accessor was called only once
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 1)
}

func TestConfigPollerV2_GetChainConfig_CacheMiss(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// Setup mock for fetch
	expectedConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(expectedConfig, expectedSourceConfigs, nil).Once()

	// Call should trigger fetch
	config, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, config)

	// Verify cache was populated
	cache := cPollerV2.getOrCreateChainCache(destChain)
	require.NotNil(t, cache)

	cache.chainConfigMu.RLock()
	assert.False(t, cache.chainConfigRefresh.IsZero())
	assert.Equal(t, expectedConfig, cache.chainConfigData)
	cache.chainConfigMu.RUnlock()
}

func TestConfigPollerV2_GetChainConfig_NoAccessor(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)
	ctx := context.Background()

	// Remove accessor for destChain to simulate missing accessor
	delete(cPollerV2.chainAccessors, destChain)

	// Should return error
	_, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no chain accessor for")
}

func TestConfigPollerV2_GetOfframpSourceChainConfigs_CacheHit(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := createMockSourceChainConfigs(sourceChains)

	// Setup mock for initial fetch
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, expectedSourceConfigs, nil).Once()

	// First call should fetch
	configs1, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs1, 2)

	// Second call should hit cache
	configs2, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Equal(t, configs1, configs2)

	// Verify accessor was called only once
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 1)
}

func TestConfigPollerV2_GetOfframpSourceChainConfigs_PartialCacheHit(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// First request: sourceChain1 and sourceChain2
	sourceChains1 := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs1 := createMockSourceChainConfigs(sourceChains1)

	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains1))).
		Return(expectedChainConfig, expectedSourceConfigs1, nil).Once()

	// Populate cache with sourceChain1 and sourceChain2
	_, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains1)
	require.NoError(t, err)

	// Second request: sourceChain1 (cached) and sourceChain3 (new)
	sourceChains2 := []cciptypes.ChainSelector{sourceChain1, sourceChain3}
	allSourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2, sourceChain3}
	expectedSourceConfigs2 := createMockSourceChainConfigs(allSourceChains)

	// GetAllConfigsLegacy should return all source chain configs (sourceChain1, sourceChain2, and sourceChain3)
	// Use MatchedBy to handle non-deterministic ordering of source chains
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(allSourceChains))).
		Return(expectedChainConfig, expectedSourceConfigs2, nil).Once()

	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains2)
	require.NoError(t, err)
	require.Len(t, configs, 2)

	// Should have sourceChain1 from cache and sourceChain3 from fetch
	assert.Contains(t, configs, sourceChain1)
	assert.Contains(t, configs, sourceChain3)

	// Verify accessor was called twice (once for each batch)
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 2)
}

func TestConfigPollerV2_GetOfframpSourceChainConfigs_FilterDestChain(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)
	ctx := context.Background()

	// Should return empty map for destination chain only
	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, []cciptypes.ChainSelector{destChain})
	require.NoError(t, err)
	assert.Empty(t, configs)

	// Test with mixed chains - should filter out destination chain
	sourceChainsNoDest := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := createMockSourceChainConfigs(sourceChainsNoDest)
	accessor := cPollerV2.chainAccessors[destChain].(*mock_ccipocr3.MockChainAccessor)

	// Mock the call to expect no dest chain in the slice of source chains
	accessor.On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChainsNoDest))).
		Return(expectedChainConfig, expectedSourceConfigs, nil).Once()

	// Pass in a slice that includes the destination chain to confirm destChain is filtered out
	sourceChainsMixedWithDest := []cciptypes.ChainSelector{destChain, sourceChain1, sourceChain2}
	configs2, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChainsMixedWithDest)
	require.NoError(t, err)
	require.Len(t, configs2, 2) // Should exclude destChain
	assert.Contains(t, configs2, sourceChain1)
	assert.Contains(t, configs2, sourceChain2)
	assert.NotContains(t, configs2, destChain)
}

func TestConfigPollerV2_BatchRefreshChainAndSourceConfigs(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	expectedChainConfig := createMockChainConfigSnapshot()
	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedSourceConfigs := createMockSourceChainConfigs(sourceChains)

	// Track source chains first
	cPollerV2.trackSourceChainForDest(sourceChain1)
	cPollerV2.trackSourceChainForDest(sourceChain2)

	// Setup mock for batch refresh, we should expect to see the two source chains sourceChain1, sourceChain2
	// that are already being tracked for destChain
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, expectedSourceConfigs, nil).Once()

	// Verify chain cache is empty before refresh
	cache := cPollerV2.getOrCreateChainCache(destChain)
	require.NotNil(t, cache)

	cache.chainConfigMu.RLock()
	assert.True(t, cache.chainConfigRefresh.IsZero())
	assert.Empty(t, cache.chainConfigData)
	cache.chainConfigMu.RUnlock()

	// Call batch refresh
	err := cPollerV2.batchRefreshChainAndSourceConfigs(ctx, destChain)
	require.NoError(t, err)

	// Verify chain cache is non-nil after refresh
	cache = cPollerV2.getOrCreateChainCache(destChain)
	require.NotNil(t, cache)

	cache.chainConfigMu.RLock()
	assert.False(t, cache.chainConfigRefresh.IsZero())
	assert.Equal(t, expectedChainConfig, cache.chainConfigData)
	cache.chainConfigMu.RUnlock()

	// Verify source configs were cached
	cache.sourceChainMu.RLock()
	assert.False(t, cache.sourceChainRefresh.IsZero())
	assert.Len(t, cache.staticSourceChainConfigs, 2)
	cache.sourceChainMu.RUnlock()
}

func TestConfigPollerV2_BatchRefreshChainAndSourceConfigs_Error(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// First, populate the cache with initial data
	initialChainConfig := createMockChainConfigSnapshot()
	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	initialSourceConfigs := createMockSourceChainConfigs(sourceChains)

	// Track source chains first so they'll be included in the batch refresh
	cPollerV2.trackSourceChainForDest(sourceChain1)
	cPollerV2.trackSourceChainForDest(sourceChain2)

	// Setup mock for initial successful population
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(initialChainConfig, initialSourceConfigs, nil).Once()

	// Populate the cache initially
	err := cPollerV2.batchRefreshChainAndSourceConfigs(ctx, destChain)
	require.NoError(t, err)

	// Capture the initial cache state after successful population
	cache := cPollerV2.getOrCreateChainCache(destChain)
	require.NotNil(t, cache)

	cache.chainConfigMu.RLock()
	initialChainRefreshTime := cache.chainConfigRefresh
	initialChainConfigData := cache.chainConfigData
	cache.chainConfigMu.RUnlock()

	cache.sourceChainMu.RLock()
	initialSourceRefreshTime := cache.sourceChainRefresh
	initialSourceConfigData := make(map[cciptypes.ChainSelector]StaticSourceChainConfig)
	maps.Copy(initialSourceConfigData, cache.staticSourceChainConfigs)
	cache.sourceChainMu.RUnlock()

	// Verify cache was populated
	assert.False(t, initialChainRefreshTime.IsZero())
	assert.Equal(t, initialChainConfig, initialChainConfigData)
	assert.False(t, initialSourceRefreshTime.IsZero())
	assert.Len(t, initialSourceConfigData, 2)

	// Now setup mock to return error on subsequent call
	expectedError := errors.New("chain accessor error")
	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(cciptypes.ChainConfigSnapshot{},
			map[cciptypes.ChainSelector]cciptypes.SourceChainConfig{},
			expectedError).Once()

	// Call should return error
	err = cPollerV2.batchRefreshChainAndSourceConfigs(ctx, destChain)
	require.Error(t, err)
	assert.Contains(t, err.Error(), expectedError.Error())

	// Verify cache was NOT updated (remains the same as before the failed call)
	cache.chainConfigMu.RLock()
	finalChainRefreshTime := cache.chainConfigRefresh
	finalChainConfigData := cache.chainConfigData
	cache.chainConfigMu.RUnlock()

	cache.sourceChainMu.RLock()
	finalSourceRefreshTime := cache.sourceChainRefresh
	finalSourceConfigData := cache.staticSourceChainConfigs
	cache.sourceChainMu.RUnlock()

	// Cache should remain unchanged after the error
	assert.Equal(t, initialChainRefreshTime, finalChainRefreshTime)
	assert.Equal(t, initialChainConfigData, finalChainConfigData)
	assert.Equal(t, initialSourceRefreshTime, finalSourceRefreshTime)
	assert.Equal(t, len(initialSourceConfigData), len(finalSourceConfigData))
	for chain, initialConfig := range initialSourceConfigData {
		assert.Equal(t, initialConfig, finalSourceConfigData[chain])
	}
}

func TestConfigPollerV2_ConcurrentAccess_GetChainConfig(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	expectedConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	// Setup mock with slow response to test concurrent access
	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Run(func(args mock.Arguments) {
			// Simulate slow network call
			time.Sleep(100 * time.Millisecond)
		}).
		Return(expectedConfig, expectedSourceConfigs, nil).Times(10)

	// Run multiple concurrent requests
	const numGoroutines = 10
	var wg sync.WaitGroup
	var successCount int32
	var errorCount int32

	for range numGoroutines {
		wg.Go(func() {
			_, err := cPollerV2.GetChainConfig(ctx, destChain)
			if err != nil {
				atomic.AddInt32(&errorCount, 1)
			} else {
				atomic.AddInt32(&successCount, 1)
			}
		})
	}

	wg.Wait()

	// All requests should succeed
	assert.Equal(t, int32(numGoroutines), successCount)
	assert.Equal(t, int32(0), errorCount)

	// Chain accessor should only be called once despite concurrent requests
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 10)
}

func TestConfigPollerV2_ConcurrentCacheAccess_PrepopulatedCache(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	expectedConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	// Setup mock for initial population
	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(expectedConfig, expectedSourceConfigs, nil).Once()

	// Populate cache first
	_, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)

	// Now run concurrent reads from cache
	const numGoroutines = 50
	var wg sync.WaitGroup
	var successCount int32

	for range numGoroutines {
		wg.Go(func() {
			_, err := cPollerV2.GetChainConfig(ctx, destChain)
			if err == nil {
				// All reads should succeed from cache
				atomic.AddInt32(&successCount, 1)
			}
		})
	}

	wg.Wait()

	// All reads should succeed
	assert.Equal(t, int32(numGoroutines), successCount)

	// Chain accessor should only be called once (for initial population)
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 1)
}

func TestConfigPollerV2_TrackSourceChain(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)

	// Track source chain
	success := cPollerV2.trackSourceChainForDest(sourceChain1)
	assert.True(t, success)

	// Verify it was tracked
	cPollerV2.RLock()
	_, exists := cPollerV2.knownSourceChains[sourceChain1]
	assert.Len(t, cPollerV2.knownSourceChains, 1)
	assert.True(t, exists)
	cPollerV2.RUnlock()

	// Try to track destination as its own source (should fail)
	success = cPollerV2.trackSourceChainForDest(destChain)
	assert.False(t, success)

	// Add another source chain
	success = cPollerV2.trackSourceChainForDest(sourceChain2)
	assert.True(t, success)

	// Verify both source chains are tracked
	cPollerV2.RLock()
	assert.Len(t, cPollerV2.knownSourceChains, 2)
	assert.Contains(t, cPollerV2.knownSourceChains, sourceChain1)
	assert.Contains(t, cPollerV2.knownSourceChains, sourceChain2)
	cPollerV2.RUnlock()
}

func TestConfigPollerV2_RefreshAllKnownChains(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// Setup initial cache with destChain ChainConfigSnapshot
	expectedConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(expectedConfig, expectedSourceConfigs, nil).Once()

	// Populate cache with destChain ChainConfigSnapshot
	_, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)

	// Track some source chains
	cPollerV2.trackSourceChainForDest(sourceChain1)
	cPollerV2.trackSourceChainForDest(sourceChain2)

	// Setup mock for refresh with updated config
	updatedConfig := createMockChainConfigSnapshot()
	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	updatedSourceConfigs := createMockSourceChainConfigs(sourceChains)

	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(updatedConfig, updatedSourceConfigs, nil).Once()

	// Mock config calls for known source chains
	emptySourceChains := make([]cciptypes.ChainSelector, 0)
	accessors[sourceChain1].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(emptySourceChains))).
		Return(updatedConfig, updatedSourceConfigs, nil).Once()
	accessors[sourceChain2].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(emptySourceChains))).
		Return(updatedConfig, updatedSourceConfigs, nil).Once()

	// Get initial refresh time
	cache := cPollerV2.getOrCreateChainCache(destChain)
	cache.chainConfigMu.RLock()
	initialChainConfigRefreshTime := cache.chainConfigRefresh
	cache.chainConfigMu.RUnlock()

	// Call refresh
	cPollerV2.refreshAllKnownChains()

	// Verify cache was updated
	cache.chainConfigMu.RLock()
	newChainConfigRefreshTime := cache.chainConfigRefresh
	newChainConfig := cache.chainConfigData
	cache.chainConfigMu.RUnlock()

	assert.True(t, newChainConfigRefreshTime.After(initialChainConfigRefreshTime))
	assert.Equal(t, updatedConfig, newChainConfig)

	// Verify source configs were updated
	cache.sourceChainMu.RLock()
	assert.Len(t, cache.staticSourceChainConfigs, 2)
	cache.sourceChainMu.RUnlock()
}

func TestConfigPollerV2_RefreshAllKnownChains_ErrorHandling(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// Setup initial cache
	expectedConfig := createMockChainConfigSnapshot()
	expectedSourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(expectedConfig, expectedSourceConfigs, nil).Once()

	// Populate cache
	_, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)

	// Setup mock to return error on refresh
	expectedError := errors.New("refresh error")
	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(cciptypes.ChainConfigSnapshot{},
			map[cciptypes.ChainSelector]cciptypes.SourceChainConfig{},
			expectedError).Once()

	// Track initial failed polls count
	initialFailedPolls := cPollerV2.consecutiveFailedPolls.Load()

	// Call refresh (should handle error gracefully)
	cPollerV2.refreshAllKnownChains()

	// Verify failed polls counter increased
	finalFailedPolls := cPollerV2.consecutiveFailedPolls.Load()
	assert.Greater(t, finalFailedPolls, initialFailedPolls)

	// Verify old data is still accessible
	config, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestConfigPollerV2_GetChainsToRefresh(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)

	// Create caches for multiple chains
	cPollerV2.getOrCreateChainCache(destChain)
	cPollerV2.getOrCreateChainCache(sourceChain1)

	// Track source chains for destination
	cPollerV2.trackSourceChainForDest(sourceChain1)
	cPollerV2.trackSourceChainForDest(sourceChain2)

	// Get chains to refresh
	chains := cPollerV2.getChainsToRefresh()

	// Should include all chains in cache
	assert.Contains(t, chains, destChain)
	assert.Contains(t, chains, sourceChain1)
	assert.Contains(t, chains, sourceChain2)
	assert.Len(t, chains, 3)
}

func TestConfigPollerV2_BackgroundPolling(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// Setup initial cache
	expectedConfigDest := createMockChainConfigSnapshot()
	expectedConfigSourceB := createMockChainConfigSnapshot()
	expectedConfigSourceC := createMockChainConfigSnapshot()
	emptyExpectedSourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	// Mock for initial population
	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Return(expectedConfigDest, emptyExpectedSourceConfigs, nil).Once()

	// Populate cache
	_, err := cPollerV2.GetChainConfig(ctx, destChain)
	require.NoError(t, err)

	// Insert some source chains to track
	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedSourceConfigs := createMockSourceChainConfigs(sourceChains)

	// Test that cache structure is properly maintained
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedConfigDest, expectedSourceConfigs, nil).Once()
	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Len(t, configs, len(sourceChains))

	// Mock for background refreshes - allow multiple calls
	updatedDestConfig := createMockChainConfigSnapshot()
	emptySourceChains := make([]cciptypes.ChainSelector, 0)
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(updatedDestConfig, expectedSourceConfigs, nil).Maybe()
	accessors[sourceChain1].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(emptySourceChains))).
		Return(expectedConfigSourceB, emptyExpectedSourceConfigs, nil).Maybe()
	accessors[sourceChain2].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(emptySourceChains))).
		Return(expectedConfigSourceC, emptyExpectedSourceConfigs, nil).Maybe()

	// Start background polling
	err = cPollerV2.Start(ctx)
	require.NoError(t, err)

	// Wait for a few refresh cycles (test config poller is configured to run every 100ms)
	time.Sleep(400 * time.Millisecond)

	// Stop background polling
	err = cPollerV2.Close()
	require.NoError(t, err)

	// Verify background polling was working (should have made additional calls)
	// We can't assert exact call count due to timing, but should be more than 1
	accessors[destChain].AssertCalled(t, "GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything)
	accessors[sourceChain1].AssertCalled(t, "GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything)
	accessors[sourceChain2].AssertCalled(t, "GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything)
}

func TestConfigPollerV2_HealthReport(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)

	// Initially should be healthy
	health := cPollerV2.HealthReport()
	assert.Len(t, health, 1)

	// Set failed polls to exceed maximum
	cPollerV2.consecutiveFailedPolls.Store(MaxFailedPolls + 1)

	// Should now report unhealthy
	health = cPollerV2.HealthReport()
	assert.Len(t, health, 1)
}

func TestConfigPollerV2_StartStop(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)
	ctx := context.Background()

	// Start should succeed
	err := cPollerV2.Start(ctx)
	require.NoError(t, err)

	// Double start should fail
	err = cPollerV2.Start(ctx)
	require.Error(t, err)

	// Insert some failed polls and verify
	cPollerV2.consecutiveFailedPolls.Store(1)
	assert.Equal(t, uint32(1), cPollerV2.consecutiveFailedPolls.Load())

	// Stop should succeed
	err = cPollerV2.Close()
	require.NoError(t, err)

	// Double stop should fail
	err = cPollerV2.Close()
	require.Error(t, err)

	// Verify failed polls counter is reset after close
	assert.Equal(t, uint32(0), cPollerV2.consecutiveFailedPolls.Load())
}

func TestConfigPollerV2_ConsecutivePollsAtomicity(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)

	// Test atomic operations on consecutiveFailedPolls
	initialValue := cPollerV2.consecutiveFailedPolls.Load()
	assert.Equal(t, uint32(0), initialValue)

	// Simulate failed polls
	for range 5 {
		cPollerV2.consecutiveFailedPolls.Add(1)
	}

	assert.Equal(t, uint32(5), cPollerV2.consecutiveFailedPolls.Load())

	// Reset counter
	cPollerV2.consecutiveFailedPolls.Store(0)
	assert.Equal(t, uint32(0), cPollerV2.consecutiveFailedPolls.Load())
}

func TestConfigPollerV2_GetOrCreateChainCache(t *testing.T) {
	cPollerV2, _ := setupConfigPollerV2(t)

	// Should create new cache
	cache1 := cPollerV2.getOrCreateChainCache(destChain)
	require.NotNil(t, cache1)
	assert.NotNil(t, cache1.staticSourceChainConfigs)

	// Should return same cache on second call
	cache2 := cPollerV2.getOrCreateChainCache(destChain)
	assert.Same(t, cache1, cache2)

	// Should return nil for chain without accessor
	delete(cPollerV2.chainAccessors, sourceChain1)
	cache3 := cPollerV2.getOrCreateChainCache(sourceChain1)
	assert.Nil(t, cache3)
}

func TestConfigPollerV2_BasicCacheFunctionality(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// Setup mock
	expectedConfig := createMockChainConfigSnapshot()
	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedSourceConfigs := createMockSourceChainConfigs(sourceChains)

	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedConfig, expectedSourceConfigs, nil).Once()

	// Populate and verify everything is in cache that we expect
	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	cache := cPollerV2.getOrCreateChainCache(destChain)
	require.NotNil(t, cache)
	cache.chainConfigMu.RLock()
	assert.False(t, cache.chainConfigRefresh.IsZero())
	cache.chainConfigMu.RUnlock()

	// Check source config cache
	cache.sourceChainMu.RLock()
	assert.False(t, cache.sourceChainRefresh.IsZero())
	assert.Len(t, cache.staticSourceChainConfigs, len(sourceChains))
	for _, chain := range sourceChains {
		assert.Contains(t, cache.staticSourceChainConfigs, chain)
	}
	cache.sourceChainMu.RUnlock()

	// Verify returned configs match cached configs
	assert.Len(t, configs, len(sourceChains))
	for _, chain := range sourceChains {
		assert.Contains(t, configs, chain)
	}
}

func TestConfigPollerV2_NoDeadlocks(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	// Dest
	expectedConfigDest := createMockChainConfigSnapshot()
	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2, sourceChain3}
	expectedSourceChainConfigs := createMockSourceChainConfigs(sourceChains)

	// Source chain D
	expectedConfigSourceD := createMockChainConfigSnapshot()
	expectedEmptySourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	// Setup mock with slow response
	accessors[destChain].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Run(func(args mock.Arguments) {
			time.Sleep(50 * time.Millisecond)
		}).
		Return(expectedConfigDest, expectedSourceChainConfigs, nil)
	accessors[sourceChain2].On("GetAllConfigsLegacy", mock.Anything, destChain, mock.Anything).
		Run(func(args mock.Arguments) {
			time.Sleep(50 * time.Millisecond)
		}).
		Return(expectedConfigSourceD, expectedEmptySourceConfigs, nil)

	// Run some concurrent operations that could maybe cause deadlock
	const numOperations = 40
	var wg sync.WaitGroup

	// Mix of read and write operations
	for i := range numOperations {
		wg.Add(1)
		go func(iteration int) {
			defer wg.Done()

			switch iteration % 4 {
			case 0:
				_, _ = cPollerV2.GetChainConfig(ctx, destChain)
				_, _ = cPollerV2.GetChainConfig(ctx, sourceChain2)
			case 1:
				_, _ = cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, []cciptypes.ChainSelector{sourceChain1})
			case 2:
				cPollerV2.trackSourceChainForDest(cciptypes.ChainSelector(iteration))
			case 3:
				cPollerV2.getOrCreateChainCache(cciptypes.ChainSelector(iteration))
			}
		}(i)
	}

	// Use a timeout to detect potential deadlocks
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Operations timed out, potential deadlock detected")
	}
}

// Helper to create a mock chain config snapshot with some randomized addresses to help assert
// correctness in tests
func createMockChainConfigSnapshot() cciptypes.ChainConfigSnapshot {
	return cciptypes.ChainConfigSnapshot{
		Offramp: cciptypes.OfframpConfig{},
		RMNProxy: cciptypes.RMNProxyConfig{
			RemoteAddress: rand.RandomAddressBytes(),
		},
		RMNRemote: cciptypes.RMNRemoteConfig{},
		FeeQuoter: cciptypes.FeeQuoterConfig{},
		OnRamp:    cciptypes.OnRampConfig{},
		Router: cciptypes.RouterConfig{
			WrappedNativeAddress: rand.RandomAddressBytes(),
		},
		CurseInfo: cciptypes.CurseInfo{},
	}
}

// Helper to create mock source chain configs
func createMockSourceChainConfigs(
	chains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.SourceChainConfig {
	configs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)
	for _, chain := range chains {
		configs[chain] = cciptypes.SourceChainConfig{
			IsEnabled: true,
			OnRamp:    cciptypes.UnknownAddress(rand.RandomAddressBytes()),
		}
	}
	return configs
}

// Helper function to create a matcher for chain selector slices (order-independent)
func chainSelectorSliceMatcher(expected []cciptypes.ChainSelector) func([]cciptypes.ChainSelector) bool {
	return func(actual []cciptypes.ChainSelector) bool {
		if len(expected) != len(actual) {
			return false
		}

		expectedSet := make(map[cciptypes.ChainSelector]bool)
		for _, chain := range expected {
			expectedSet[chain] = true
		}

		for _, chain := range actual {
			if !expectedSet[chain] {
				return false
			}
		}

		return true
	}
}

// TestConfigPollerV2_FailedSourceChainDoesNotRetriggerInlineFetch verifies that when a source chain
// config fetch fails (e.g. RPC error), subsequent calls to GetOfframpSourceChainConfigs do NOT
// trigger another inline batchRefreshChainAndSourceConfigs call within the suppression window.
// This prevents the "doom-loop" where persistently failing chains cause redundant RPC calls
// and error log spam on every tick.
func TestConfigPollerV2_FailedSourceChainDoesNotRetriggerInlineFetch(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	// Use a long refresh period so the inline-retry backoff comfortably covers the test
	// (suppression is `time.Since(lastAttempted) < refreshPeriod`).
	cPollerV2.refreshPeriod = 30 * time.Second
	ctx := context.Background()

	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()

	// First fetch: sourceChain1 succeeds, sourceChain2 is missing from the response (simulates RPC failure
	// where processSourceChainConfigResults skips the chain due to an error).
	partialSourceConfigs := createMockSourceChainConfigs([]cciptypes.ChainSelector{sourceChain1})

	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, partialSourceConfigs, nil).Once()

	// First call: triggers inline fetch since nothing is cached
	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	// Only sourceChain1 should be returned (sourceChain2 failed)
	assert.Len(t, configs, 1)
	assert.Contains(t, configs, sourceChain1)
	assert.NotContains(t, configs, sourceChain2)

	// Second call with the SAME chains: should NOT trigger another GetAllConfigsLegacy call
	// because sourceChain2 was already attempted (it's in attemptedSourceChains).
	configs2, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	// Still only sourceChain1 (sourceChain2 remains failed but no re-fetch happened)
	assert.Len(t, configs2, 1)
	assert.Contains(t, configs2, sourceChain1)

	// Verify GetAllConfigsLegacy was only called ONCE total, the second call did not trigger a re-fetch
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 1)
}

// TestConfigPollerV2_BackgroundPollerRetriesFailedSourceChains verifies that the background
// poller successfully retries and recovers source chains that previously failed.
func TestConfigPollerV2_BackgroundPollerRetriesFailedSourceChains(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()
	emptySourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	// First fetch: only sourceChain1 succeeds (sourceChain2 is missing, simulates RPC failure)
	partialSourceConfigs := createMockSourceChainConfigs([]cciptypes.ChainSelector{sourceChain1})

	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, partialSourceConfigs, nil).Once()

	// Initial inline fetch
	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Len(t, configs, 1)
	assert.Contains(t, configs, sourceChain1)

	// Now simulate the background poller succeeding for both chains on the dest chain refresh
	fullSourceConfigs := createMockSourceChainConfigs(sourceChains)
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, fullSourceConfigs, nil).Maybe()

	// refreshAllKnownChains also refreshes individual source chains (not just dest),
	// so we need to mock those too.
	emptyChains := make([]cciptypes.ChainSelector, 0)
	for _, chain := range sourceChains {
		accessors[chain].On(
			"GetAllConfigsLegacy",
			mock.Anything,
			destChain,
			mock.MatchedBy(chainSelectorSliceMatcher(emptyChains))).
			Return(createMockChainConfigSnapshot(), emptySourceConfigs, nil).Maybe()
	}

	// Simulate background poller running
	cPollerV2.refreshAllKnownChains()

	// Now GetOfframpSourceChainConfigs should return BOTH chains from cache
	configs2, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Len(t, configs2, 2)
	assert.Contains(t, configs2, sourceChain1)
	assert.Contains(t, configs2, sourceChain2)
}

// TestConfigPollerV2_NewSourceChainStillTriggersInlineFetch verifies that a truly new source
// chain (never attempted before) still triggers an inline fetch, even if other chains have
// already been attempted.
func TestConfigPollerV2_NewSourceChainStillTriggersInlineFetch(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	ctx := context.Background()

	initialChains := []cciptypes.ChainSelector{sourceChain1}
	expectedChainConfig := createMockChainConfigSnapshot()
	initialSourceConfigs := createMockSourceChainConfigs(initialChains)

	// First fetch: sourceChain1 only
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(initialChains))).
		Return(expectedChainConfig, initialSourceConfigs, nil).Once()

	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, initialChains)
	require.NoError(t, err)
	assert.Len(t, configs, 1)

	// Now request with a NEW chain (sourceChain3) that was never attempted
	newChains := []cciptypes.ChainSelector{sourceChain1, sourceChain3}
	allKnownChains := []cciptypes.ChainSelector{sourceChain1, sourceChain3}
	allSourceConfigs := createMockSourceChainConfigs(allKnownChains)

	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(allKnownChains))).
		Return(expectedChainConfig, allSourceConfigs, nil).Once()

	// This should trigger an inline fetch because sourceChain3 was never attempted
	configs2, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, newChains)
	require.NoError(t, err)
	assert.Len(t, configs2, 2)
	assert.Contains(t, configs2, sourceChain1)
	assert.Contains(t, configs2, sourceChain3)

	// Verify GetAllConfigsLegacy was called twice (once for initial, once for new chain)
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 2)
}

// TestConfigPollerV2_AllSourceChainsFail_NoInlineRetry verifies that when ALL source chains
// fail on the initial fetch, subsequent calls within the suppression window still don't
// trigger inline re-fetches.
func TestConfigPollerV2_AllSourceChainsFail_NoInlineRetry(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	cPollerV2.refreshPeriod = 30 * time.Second
	ctx := context.Background()

	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()

	// All source chains fail, empty map returned (simulates all RPC calls failing)
	emptySourceConfigs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, emptySourceConfigs, nil).Once()

	// First call: triggers inline fetch
	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Empty(t, configs)

	// Second call: should NOT trigger another fetch
	configs2, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Empty(t, configs2)

	// Verify only one call was made
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 1)
}

// TestConfigPollerV2_BatchErrorDoesNotPoisonAttemptedSet verifies that when the top-level
// GetAllConfigsLegacy call returns an error (as opposed to a successful response with partial
// per-chain failures), attemptedSourceChains is NOT populated. Inline retries must continue
// across transient batch failures (e.g., provider RPC unavailable) so we don't get wedged on
// the first failed call.
func TestConfigPollerV2_BatchErrorDoesNotPoisonAttemptedSet(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	cPollerV2.refreshPeriod = 30 * time.Second
	ctx := context.Background()

	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()
	fullSourceConfigs := createMockSourceChainConfigs(sourceChains)
	var nilSourceConfigs map[cciptypes.ChainSelector]cciptypes.SourceChainConfig

	// First call: top-level batch errors (simulates RPC provider unavailable).
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(cciptypes.ChainConfigSnapshot{}, nilSourceConfigs, errors.New("rpc unavailable")).Once()

	_, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.Error(t, err)

	// attemptedSourceChains must remain empty after a batch-level error so the next call
	// is treated as a true cache miss and triggers another inline fetch.
	destCache := cPollerV2.getOrCreateChainCache(destChain)
	require.NotNil(t, destCache)
	destCache.sourceChainMu.RLock()
	assert.Empty(t, destCache.attemptedSourceChains,
		"attemptedSourceChains must not be populated when GetAllConfigsLegacy returns an error")
	destCache.sourceChainMu.RUnlock()

	// Second call: top-level batch now succeeds for both chains.
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, fullSourceConfigs, nil).Once()

	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Len(t, configs, 2)
	assert.Contains(t, configs, sourceChain1)
	assert.Contains(t, configs, sourceChain2)

	// Verify GetAllConfigsLegacy was called exactly twice (one error + one success).
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 2)
}

// TestConfigPollerV2_InlineRetryResumesAfterBackoff verifies that once the suppression window
// elapses, a still-missing source chain triggers another inline fetch. This guards the recovery
// path used when the background poller is unhealthy or has been killed.
func TestConfigPollerV2_InlineRetryResumesAfterBackoff(t *testing.T) {
	cPollerV2, accessors := setupConfigPollerV2(t)
	// Very short backoff so the test can wait it out without flakiness.
	cPollerV2.refreshPeriod = 50 * time.Millisecond
	ctx := context.Background()

	sourceChains := []cciptypes.ChainSelector{sourceChain1, sourceChain2}
	expectedChainConfig := createMockChainConfigSnapshot()

	// First fetch: sourceChain1 succeeds, sourceChain2 is missing.
	partialSourceConfigs := createMockSourceChainConfigs([]cciptypes.ChainSelector{sourceChain1})
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, partialSourceConfigs, nil).Once()

	configs, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Len(t, configs, 1)
	assert.Contains(t, configs, sourceChain1)

	// While the suppression window is open, no additional fetch should fire.
	configs2, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Len(t, configs2, 1)
	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 1)

	// Wait for the backoff to elapse, then expect an inline retry that recovers sourceChain2.
	time.Sleep(75 * time.Millisecond)

	fullSourceConfigs := createMockSourceChainConfigs(sourceChains)
	accessors[destChain].On(
		"GetAllConfigsLegacy",
		mock.Anything,
		destChain,
		mock.MatchedBy(chainSelectorSliceMatcher(sourceChains))).
		Return(expectedChainConfig, fullSourceConfigs, nil).Once()

	configs3, err := cPollerV2.GetOfframpSourceChainConfigs(ctx, destChain, sourceChains)
	require.NoError(t, err)
	assert.Len(t, configs3, 2)
	assert.Contains(t, configs3, sourceChain1)
	assert.Contains(t, configs3, sourceChain2)

	accessors[destChain].AssertNumberOfCalls(t, "GetAllConfigsLegacy", 2)
}
