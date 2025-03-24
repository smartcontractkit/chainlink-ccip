package reader

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func setupBasicCache(t *testing.T) (*configPoller, *reader_mocks.MockExtended) {
	mockReader := reader_mocks.NewMockExtended(t)

	reader := &ccipChainReader{
		lggr: logger.Test(t),
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			chainA: mockReader,
		},
		destChain: chainA,
	}

	cache := newConfigPoller(logger.Test(t), reader, 1*time.Second)
	return cache, mockReader
}

// Helper to setup a standard mock response for chain configuration
func setupMockResponse(reader *reader_mocks.MockExtended) types.BatchGetLatestValuesResult {
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil)

	return responses
}

// Helper to setup a batch response containing both chain config and source chain configs
func setupBatchMockResponse(reader *reader_mocks.MockExtended) {
	// Chain config part
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	// Source chain config part
	resultB := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	resultB.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	resultC := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	resultC.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	// Combined response
	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			// Standard chain config results
			*result1, *result2, *result3, *result4,
			// Source chain config results
			*resultB, *resultC,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil)
}

// Helper to setup initial data in the cache
func setupInitialData(ctx context.Context, cache *configPoller, reader *reader_mocks.MockExtended) {
	setupMockResponse(reader)

	// Call GetChainConfig to populate the cache
	_, err := cache.GetChainConfig(ctx, chainA)
	if err != nil {
		panic(fmt.Sprintf("Failed to setup initial data: %v", err))
	}

	// Setup source chain configs
	sourceChains := []cciptypes.ChainSelector{chainB, chainC}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	response := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response, []string{}, nil).Once()

	// Call GetOfframpSourceChainConfigs to populate the source chain cache
	_, err = cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	if err != nil {
		panic(fmt.Sprintf("Failed to setup initial source chain data: %v", err))
	}
}

// Helper to setup a second chain in the cache
func setupSecondChain(ctx context.Context, t *testing.T, cache *configPoller) *reader_mocks.MockExtended {
	// Create a new mock reader for the second chain
	mockReader := reader_mocks.NewMockExtended(t)

	// Type assertion to access the underlying ccipChainReader
	// This is acceptable in test code since we know the concrete type
	reader, ok := cache.reader.(*ccipChainReader)
	if !ok {
		t.Fatalf("Expected cache.reader to be *ccipChainReader, got %T", cache.reader)
	}

	// Now we can access contractReaders directly
	reader.contractReaders[chainB] = mockReader

	// Setup mock response for second chain
	setupMockResponse(mockReader)

	// Call GetChainConfig to populate the cache for the second chain
	_, err := cache.GetChainConfig(ctx, chainB)
	if err != nil {
		panic(fmt.Sprintf("Failed to setup second chain data: %v", err))
	}

	return mockReader
}

// Helper to setup refresh expectations
func setupRefreshExpectations(reader *reader_mocks.MockExtended) {
	// Setup with updated values to verify refresh happened
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 2, N: 6}, // Different from initial
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil)
}

func TestConfigPoller_StartStop(t *testing.T) {
	cache, reader := setupBasicCache(t)

	// Setup basic response for initial fetch and background refresh
	setupMockResponse(reader)

	// Start the background poller
	err := cache.Start(t.Context())
	require.NoError(t, err, "Starting config poller should not error")

	// Verify it's running by letting it execute at least once
	time.Sleep(2 * cache.refreshPeriod)

	// Stop the poller
	err = cache.Close()
	require.NoError(t, err, "Stopping config poller should not error")

	// Reset the mock counts/expectations
	reader.ExpectedCalls = nil

	// Setup expectation that should not be called
	reader.On("ExtendedBatchGetLatestValues", mock.Anything, mock.Anything, mock.Anything).
		Maybe().Return(nil, nil, nil)

	// Sleep for refresh period again - no calls should occur
	time.Sleep(2 * cache.refreshPeriod)

	// Verify no calls occurred after stopping
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 0)
}

func TestConfigPoller_BatchRefresh(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup source chains
	sourceChains := []cciptypes.ChainSelector{chainB, chainC}

	// Setup expected batch results
	setupBatchMockResponse(reader)

	// Call the batch refresh method
	err := cache.batchRefreshChainAndSourceConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)

	// Verify both chain config and source configs were updated in a single call
	chainConfig, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.NotEqual(t, ChainConfigSnapshot{}, chainConfig)

	sourceConfigs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	assert.Len(t, sourceConfigs, 2)

	// Verify only one call was made for both sets of data
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 1)
}

func TestConfigPoller_RefreshAllKnownChains(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// First populate the cache with initial data
	setupInitialData(ctx, cache, reader)

	// Add a second chain to the cache
	secondReader := setupSecondChain(ctx, t, cache)

	// Reset mock counts for both readers
	reader.ExpectedCalls = nil
	secondReader.ExpectedCalls = nil

	// Create a standard mock response for any request
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	standardResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}

	// Setup expectations for both readers
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(standardResponse, []string{}, nil)

	secondReader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(standardResponse, []string{}, nil)

	// Before refreshing, remember how many chains we have
	numChains := len(cache.chainCaches)
	require.GreaterOrEqual(t, numChains, 2, "Test should have at least 2 chains in cache")

	// Track initial refresh times
	refreshTimes := make(map[cciptypes.ChainSelector]time.Time)
	for chainSel, chainCache := range cache.chainCaches {
		chainCache.chainConfigMu.RLock()
		refreshTimes[chainSel] = chainCache.chainConfigRefresh
		chainCache.chainConfigMu.RUnlock()
	}

	// Call the refresh method
	cache.refreshAllKnownChains()

	// Verify all chains were refreshed by checking their refresh timestamps updated
	for chainSel, chainCache := range cache.chainCaches {
		chainCache.chainConfigMu.RLock()
		newRefreshTime := chainCache.chainConfigRefresh
		chainCache.chainConfigMu.RUnlock()

		assert.True(t, newRefreshTime.After(refreshTimes[chainSel]),
			"Chain %v should have been refreshed", chainSel)
	}

	// Verify both readers were called
	reader.AssertCalled(t, "ExtendedBatchGetLatestValues", mock.Anything, mock.Anything, mock.Anything)
	secondReader.AssertCalled(t, "ExtendedBatchGetLatestValues", mock.Anything, mock.Anything, mock.Anything)
}

func TestConfigPoller_TrackSourceChain(t *testing.T) {
	cache, reader := setupBasicCache(t)

	// Track a source chain
	success := cache.trackSourceChain(chainA, chainB)
	assert.True(t, success)

	// Track another source chain
	success = cache.trackSourceChain(chainA, chainC)
	assert.True(t, success)

	// Attempt to track destination as its own source (should fail)
	success = cache.trackSourceChain(chainA, chainA)
	assert.False(t, success)

	// First populate the cache with an initial request
	ctx := tests.Context(t)
	setupMockResponse(reader)
	_, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)

	// Reset mock expectations
	reader.ExpectedCalls = nil

	// Setup refresh expectations
	setupRefreshExpectations(reader)

	// Start the background poller
	err = cache.Start(t.Context())
	require.NoError(t, err)

	// Let it run for a refresh cycle (increased duration)
	time.Sleep(3 * cache.refreshPeriod)

	// Stop the poller
	err = cache.Close()
	require.NoError(t, err)

	// Verify the mock was called
	reader.AssertCalled(t, "ExtendedBatchGetLatestValues", mock.Anything, mock.Anything, true)
}

func TestConfigPoller_BackgroundErrorHandling(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup initial successful fetch
	setupInitialData(ctx, cache, reader)

	// Reset mock counts
	reader.ExpectedCalls = nil

	// Setup error response for background refresh
	reader.On("ExtendedBatchGetLatestValues", mock.Anything, mock.Anything, mock.Anything).
		Return(nil, nil, errors.New("simulated error"))

	// Start the poller
	err := cache.Start(t.Context())
	require.NoError(t, err)

	// Let it run and encounter the error
	time.Sleep(2 * cache.refreshPeriod)

	// Stop the poller
	err = cache.Close()
	require.NoError(t, err)

	// Verify the service continues running despite errors
	// This can be done by checking that initial data is still available
	config, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.NotEqual(t, ChainConfigSnapshot{}, config)
}

func TestConfigPoller_ConcurrentWithBackground(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup with initial data
	setupInitialData(ctx, cache, reader)

	// Setup slow refresh response that will hold the lock for a while
	reader.On("ExtendedBatchGetLatestValues", mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			// Simulate slow RPC call
			time.Sleep(500 * time.Millisecond)
		}).Return(setupMockResponse(reader), []string{}, nil)

	// Start the background poller
	err := cache.Start(t.Context())
	require.NoError(t, err)

	// Sleep briefly to ensure background poller has started a refresh
	time.Sleep(100 * time.Millisecond)

	// Now try to read concurrently while the refresh is in progress
	start := time.Now()
	config, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	elapsed := time.Since(start)

	// Stop the poller
	err = cache.Close()
	require.NoError(t, err)

	// Verify the read was fast (non-blocking) despite ongoing refresh
	require.NoError(t, err)
	assert.Less(t, elapsed, 100*time.Millisecond, "Read should not be blocked by background refresh")
	assert.NotEqual(t, ChainConfigSnapshot{}, config)
}

func TestConfigCache_GetChainConfig_CacheHit(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup mock for initial fetch
	mockCommitOCRConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}
	mockExecOCRConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 2, N: 6},
		},
	}

	// Setup batch results
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockCommitOCRConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockExecOCRConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil).Once()

	// First call should fetch
	config1, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), config1.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)
	assert.Equal(t, uint8(4), config1.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.N)

	// Second call within refresh period should hit cache
	config2, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, config1, config2)

	// Verify the mock was called exactly once
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 1)
}

func TestConfigCache_GetChainConfig_CacheUpdate(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup mock responses for two different fetches
	setupMockBatchResponse := func(f uint8, n uint8) types.BatchGetLatestValuesResult {
		mockConfig := OCRConfigResponse{
			OCRConfig: OCRConfig{
				ConfigInfo: ConfigInfo{F: f, N: n},
			},
		}

		result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
		result1.SetResult(&mockConfig, nil)
		result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
		result2.SetResult(&mockConfig, nil)
		result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
		result3.SetResult(&offRampStaticChainConfig{}, nil)
		result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
		result4.SetResult(&offRampDynamicChainConfig{}, nil)

		return types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameOffRamp}: {
				*result1, *result2, *result3, *result4,
			},
		}
	}

	// Setup first response with F=1, N=4
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(setupMockBatchResponse(1, 4), []string{}, nil).Once()

	// First call should fetch initial config
	config1, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), config1.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)
	assert.Equal(t, uint8(4), config1.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.N)

	// Setup second response with F=2, N=6
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(setupMockBatchResponse(2, 6), []string{}, nil).Once()

	// Simulate what the background poller would do: manually refresh the cache
	cache.refreshAllKnownChains()

	// The next call to GetChainConfig should return the refreshed data
	config2, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(2), config2.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)
	assert.Equal(t, uint8(6), config2.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.N)
	assert.NotEqual(t, config1, config2)

	// Verify the contract reader was called exactly twice
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 2)
}

func TestConfigCache_GetChainConfig_Error(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	expectedErr := errors.New("fetch error")
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(nil, nil, expectedErr)

	_, err := cache.GetChainConfig(ctx, chainA)
	require.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)
}

func TestConfigCache_NoReader(t *testing.T) {
	cache, _ := setupBasicCache(t)
	ctx := tests.Context(t)

	// Test with a chain that has no reader
	_, err := cache.GetChainConfig(ctx, chainB)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no contract reader for chain")
}

func TestConfigCache_ErrorWithCachedData(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup initial successful fetch
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil).Once()

	// First call should succeed and cache data
	config1, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)

	// Setup error for background refresh attempt
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(nil, nil, errors.New("fetch error")).Once()

	// Simulate background polling with error
	cache.refreshAllKnownChains()

	// Call should still return cached data despite failed refresh
	config2, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, config1, config2)
}

func TestConfigCache_RefreshChainConfig(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup mock response
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil).Once()

	// Force refresh should fetch regardless of cache state
	config, err := cache.refreshChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), config.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)

	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 1)
}

func TestConfigCache_ConcurrentAccess(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup mock response
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}

	// First, load the cache with an initial fetch
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil).Once()

	// Prefill the cache
	_, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)

	// Run concurrent fetches on the filled cache
	const numGoroutines = 10
	errCh := make(chan error, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			_, err := cache.GetChainConfig(ctx, chainA)
			errCh <- err
		}()
	}

	// Collect results
	for i := 0; i < numGoroutines; i++ {
		err := <-errCh
		require.NoError(t, err)
	}

	// Should only fetch once (for the prefill) despite concurrent access
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 1)
}

func TestConfigCache_Initialization(t *testing.T) {
	testCases := []struct {
		name          string
		setupReader   func() *ccipChainReader
		refreshPeriod time.Duration
		chainToTest   cciptypes.ChainSelector
		expectedErr   string
	}{
		{
			name: "nil readers map",
			setupReader: func() *ccipChainReader {
				return &ccipChainReader{
					lggr:            logger.Test(t),
					contractReaders: nil,
					destChain:       chainA,
				}
			},
			refreshPeriod: time.Second,
			chainToTest:   chainA,
			expectedErr:   "no contract reader for chain",
		},
		{
			name: "empty readers map",
			setupReader: func() *ccipChainReader {
				return &ccipChainReader{
					lggr:            logger.Test(t),
					contractReaders: make(map[cciptypes.ChainSelector]contractreader.Extended),
					destChain:       chainA,
				}
			},
			refreshPeriod: time.Second,
			chainToTest:   chainA,
			expectedErr:   "no contract reader for chain",
		},
		{
			name: "missing specific chain",
			setupReader: func() *ccipChainReader {
				return &ccipChainReader{
					lggr: logger.Test(t),
					contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
						chainB: nil, // Different chain than we'll test
					},
					destChain: chainA,
				}
			},
			refreshPeriod: time.Second,
			chainToTest:   chainA,
			expectedErr:   "no contract reader for chain",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)
			ctx := tests.Context(t)

			reader := tc.setupReader()
			cache := newConfigPoller(lggr, reader, tc.refreshPeriod)
			require.NotNil(t, cache, "cache should never be nil after initialization")

			require.NotNil(t, cache.chainCaches, "chainCaches map should never be nil")
			assert.Equal(t, tc.refreshPeriod, cache.refreshPeriod)
			assert.Equal(t, reader, cache.reader)

			_, err := cache.GetChainConfig(ctx, tc.chainToTest)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}

func TestConfigCache_GetChainConfig_SkippedContracts(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup mock response with skipped contracts
	mockConfig := OCRConfigResponse{
		OCRConfig: OCRConfig{
			ConfigInfo: ConfigInfo{F: 1, N: 4},
		},
	}

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result1.SetResult(&mockConfig, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
	result2.SetResult(&mockConfig, nil)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
	result3.SetResult(&offRampStaticChainConfig{}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
	result4.SetResult(&offRampDynamicChainConfig{}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, *result3, *result4,
		},
	}
	skippedContracts := []string{consts.ContractNameRMNProxy}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, skippedContracts, nil).Once()

	// Should succeed even with skipped contracts
	config, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), config.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)
}

func TestConfigCache_InvalidResults(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	testCases := []struct {
		name        string
		setupMock   func() types.BatchGetLatestValuesResult
		expectedErr string
	}{
		{
			name: "missing offramp results",
			setupMock: func() types.BatchGetLatestValuesResult {
				// Setup minimum valid responses for other contracts
				rmnProxyResult := &types.BatchReadResult{ReadName: consts.MethodNameGetARM}
				rmnProxyResult.SetResult(&[]byte{1, 2, 3}, nil)

				rmnDigestResult := &types.BatchReadResult{ReadName: consts.MethodNameGetReportDigestHeader}
				rmnDigestResult.SetResult(&rmnDigestHeader{}, nil)
				rmnConfigResult := &types.BatchReadResult{ReadName: consts.MethodNameGetVersionedConfig}
				rmnConfigResult.SetResult(&versionedConfig{}, nil)
				rmnCurseResult := &types.BatchReadResult{ReadName: consts.MethodNameGetCursedSubjects}
				rmnCurseResult.SetResult(&RMNCurseResponse{}, nil)

				feeQuoterResult := &types.BatchReadResult{ReadName: consts.MethodNameFeeQuoterGetStaticConfig}
				feeQuoterResult.SetResult(&feeQuoterStaticConfig{}, nil)

				return types.BatchGetLatestValuesResult{
					types.BoundContract{Name: consts.ContractNameOffRamp}:   {},
					types.BoundContract{Name: consts.ContractNameRMNProxy}:  {*rmnProxyResult},
					types.BoundContract{Name: consts.ContractNameRMNRemote}: {*rmnDigestResult, *rmnConfigResult, *rmnCurseResult},
					types.BoundContract{Name: consts.ContractNameFeeQuoter}: {*feeQuoterResult},
				}
			},
			expectedErr: "expected 4 offramp results",
		},
		{
			name: "invalid commit config type",
			setupMock: func() types.BatchGetLatestValuesResult {
				// Setup all 4 required offramp results, but make the first one invalid
				result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
				result1.SetResult("invalid type", nil)
				result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
				result2.SetResult(&OCRConfigResponse{}, nil)
				result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
				result3.SetResult(&offRampStaticChainConfig{}, nil)
				result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
				result4.SetResult(&offRampDynamicChainConfig{}, nil)

				// Setup valid responses for other contracts
				rmnProxyResult := &types.BatchReadResult{ReadName: consts.MethodNameGetARM}
				rmnProxyResult.SetResult(&[]byte{1, 2, 3}, nil)

				rmnDigestResult := &types.BatchReadResult{ReadName: consts.MethodNameGetReportDigestHeader}
				rmnDigestResult.SetResult(&rmnDigestHeader{}, nil)
				rmnConfigResult := &types.BatchReadResult{ReadName: consts.MethodNameGetVersionedConfig}
				rmnConfigResult.SetResult(&versionedConfig{}, nil)
				rmnCurseResult := &types.BatchReadResult{ReadName: consts.MethodNameGetCursedSubjects}
				rmnCurseResult.SetResult(&RMNCurseResponse{}, nil)

				feeQuoterResult := &types.BatchReadResult{ReadName: consts.MethodNameFeeQuoterGetStaticConfig}
				feeQuoterResult.SetResult(&feeQuoterStaticConfig{}, nil)

				return types.BatchGetLatestValuesResult{
					types.BoundContract{Name: consts.ContractNameOffRamp}: {
						*result1, *result2, *result3, *result4,
					},
					types.BoundContract{Name: consts.ContractNameRMNProxy}:  {*rmnProxyResult},
					types.BoundContract{Name: consts.ContractNameRMNRemote}: {*rmnDigestResult, *rmnConfigResult, *rmnCurseResult},
					types.BoundContract{Name: consts.ContractNameFeeQuoter}: {*feeQuoterResult},
				}
			},
			expectedErr: "invalid type for CommitLatestOCRConfig",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader.On("ExtendedBatchGetLatestValues",
				mock.Anything,
				mock.Anything,
				true,
			).Return(tc.setupMock(), []string{}, nil).Once()

			_, err := cache.GetChainConfig(ctx, chainA)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.expectedErr)

			reader.AssertExpectations(t)
		})
	}
}

func TestConfigCache_MultipleChains(t *testing.T) {
	readerA := reader_mocks.NewMockExtended(t)
	readerB := reader_mocks.NewMockExtended(t)

	// Create ccipChainReader with multiple chain readers
	reader := &ccipChainReader{
		lggr: logger.Test(t),
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			chainA: readerA,
			chainB: readerB,
		},
		destChain: chainA,
	}

	cache := newConfigPoller(logger.Test(t), reader, 1*time.Second)
	ctx := tests.Context(t)

	// Setup mock response for both chains
	setupMockResponse := func(f uint8) types.BatchGetLatestValuesResult {
		mockConfig := OCRConfigResponse{
			OCRConfig: OCRConfig{
				ConfigInfo: ConfigInfo{F: f, N: 4},
			},
		}

		result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
		result1.SetResult(&mockConfig, nil)
		result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
		result2.SetResult(&mockConfig, nil)
		result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
		result3.SetResult(&offRampStaticChainConfig{}, nil)
		result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
		result4.SetResult(&offRampDynamicChainConfig{}, nil)

		return types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameOffRamp}: {
				*result1, *result2, *result3, *result4,
			},
		}
	}

	readerA.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(setupMockResponse(1), []string{}, nil).Once()

	readerB.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(setupMockResponse(2), []string{}, nil).Once()

	// Get configs for both chains
	configA, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), configA.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)

	configB, err := cache.GetChainConfig(ctx, chainB)
	require.NoError(t, err)
	assert.Equal(t, uint8(2), configB.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)

	// Each reader should be called exactly once
	readerA.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 1)
	readerB.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 1)
}

func TestConfigCache_BackgroundRefreshPeriod(t *testing.T) {
	// Test with different refresh periods
	testCases := []struct {
		name          string
		refreshPeriod time.Duration
		waitTime      time.Duration
		expectedCalls int
	}{
		{
			name:          "multiple refresh cycles",
			refreshPeriod: 100 * time.Millisecond,
			waitTime:      250 * time.Millisecond, // Should trigger ~2 refreshes
			expectedCalls: 3,                      // Initial + 2 background refreshes
		},
		{
			name:          "single refresh cycle",
			refreshPeriod: 200 * time.Millisecond,
			waitTime:      250 * time.Millisecond, // Should trigger 1 refresh
			expectedCalls: 2,                      // Initial + 1 background refresh
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockReader := reader_mocks.NewMockExtended(t)

			// Create ccipChainReader with the mock reader
			reader := &ccipChainReader{
				lggr: logger.Test(t),
				contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
					chainA: mockReader,
				},
				destChain: chainA,
			}

			cache := newConfigPoller(logger.Test(t), reader, tc.refreshPeriod)
			ctx := tests.Context(t)

			mockConfig := OCRConfigResponse{
				OCRConfig: OCRConfig{
					ConfigInfo: ConfigInfo{F: 1, N: 4},
				},
			}

			result1 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
			result1.SetResult(&mockConfig, nil)
			result2 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampLatestConfigDetails}
			result2.SetResult(&mockConfig, nil)
			result3 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetStaticConfig}
			result3.SetResult(&offRampStaticChainConfig{}, nil)
			result4 := &types.BatchReadResult{ReadName: consts.MethodNameOffRampGetDynamicConfig}
			result4.SetResult(&offRampDynamicChainConfig{}, nil)

			responses := types.BatchGetLatestValuesResult{
				types.BoundContract{Name: consts.ContractNameOffRamp}: {
					*result1, *result2, *result3, *result4,
				},
			}

			// Setup mock to return the same response for all calls
			mockReader.On("ExtendedBatchGetLatestValues",
				mock.Anything,
				mock.Anything,
				true,
			).Return(responses, []string{}, nil).Times(tc.expectedCalls)

			// First call to initialize cache
			_, err := cache.GetChainConfig(ctx, chainA)
			require.NoError(t, err)

			// Start the background poller
			err = cache.Start(t.Context())
			require.NoError(t, err)

			// Wait for the specified time to allow background polling
			time.Sleep(tc.waitTime)

			// Stop the background poller
			err = cache.Close()
			require.NoError(t, err)

			// Verify the correct number of calls were made
			mockReader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", tc.expectedCalls)
		})
	}
}

func TestConfigCache_GetOfframpSourceChainConfigs_CacheHit(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Setup mock response for source chain configs
	sourceChains := []cciptypes.ChainSelector{chainB, chainC}

	// Create batch read results for source chains
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(responses, []string{}, nil).Once()

	// First call should fetch
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs, 2)
	assert.True(t, configs[chainB].IsEnabled)
	assert.Equal(t, cciptypes.UnknownAddress{1, 2, 3}, configs[chainB].OnRamp)
	assert.True(t, configs[chainC].IsEnabled)
	assert.Equal(t, cciptypes.UnknownAddress{4, 5, 6}, configs[chainC].OnRamp)

	// Second call within refresh period should hit cache
	configs2, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs2, 2)
	assert.Equal(t, configs, configs2)

	// Verify the mock was called exactly once
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 1)
}

func TestConfigCache_GetOfframpSourceChainConfigs_Update(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	sourceChains := []cciptypes.ChainSelector{chainB, chainC}

	// Setup mock response for first fetch
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	firstResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(firstResponse, []string{}, nil).Once()

	// First call should fetch initial data
	configs1, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs1, 2)
	assert.Equal(t, cciptypes.UnknownAddress{1, 2, 3}, configs1[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{4, 5, 6}, configs1[chainC].OnRamp)
	assert.True(t, configs1[chainB].IsEnabled)
	assert.True(t, configs1[chainC].IsEnabled)

	// Setup mock response for second fetch with different data
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result3.SetResult(&SourceChainConfig{IsEnabled: false, OnRamp: cciptypes.UnknownAddress{7, 8, 9}}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result4.SetResult(&SourceChainConfig{IsEnabled: false, OnRamp: cciptypes.UnknownAddress{10, 11, 12}}, nil)

	secondResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result3, *result4,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(secondResponse, []string{}, nil).Once()

	// Manually refresh the source chain configs (simulating background refresh)
	newConfigs, err := cache.refreshSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, newConfigs, 2)
	assert.Equal(t, cciptypes.UnknownAddress{7, 8, 9}, newConfigs[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{10, 11, 12}, newConfigs[chainC].OnRamp)
	assert.False(t, newConfigs[chainB].IsEnabled)
	assert.False(t, newConfigs[chainC].IsEnabled)

	// Next call should return the updated values
	configs2, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs2, 2)

	// Verify the updated values match what we expect
	assert.Equal(t, cciptypes.UnknownAddress{7, 8, 9}, configs2[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{10, 11, 12}, configs2[chainC].OnRamp)
	assert.False(t, configs2[chainB].IsEnabled)
	assert.False(t, configs2[chainC].IsEnabled)

	// Verify they don't match the original values
	assert.NotEqual(t, configs1, configs2)

	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 2)
}

func TestConfigCache_GetOfframpSourceChainConfigs_MixedSet(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// First request for chains B and C
	sourceChains1 := []cciptypes.ChainSelector{chainB, chainC}

	// Setup mock response for first fetch
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	firstResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(firstResponse, []string{}, nil).Once()

	// First call should fetch both B and C
	configs1, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains1)
	require.NoError(t, err)
	require.Len(t, configs1, 2)

	// Second request with chains B and D (mix of cached and new)
	sourceChains2 := []cciptypes.ChainSelector{chainB, chainD}

	// Setup mock response for second fetch (only D should be fetched)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result3.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{7, 8, 9}}, nil)

	secondResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result3,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(secondResponse, []string{}, nil).Once()

	// Second call should only fetch chain D and use cached value for B
	configs2, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains2)
	require.NoError(t, err)
	require.Len(t, configs2, 2)

	// Chain B should be the same as in first request
	assert.Equal(t, configs1[chainB], configs2[chainB])

	// Chain D should be newly fetched
	assert.Equal(t, cciptypes.UnknownAddress{7, 8, 9}, configs2[chainD].OnRamp)

	// Verify the mock was called twice (once for each fetch)
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 2)
}

func TestConfigCache_RefreshSourceChainConfigs(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	sourceChains := []cciptypes.ChainSelector{chainB, chainC}

	// Setup mock response
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	response := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response, []string{}, nil).Once()

	// Force refresh should fetch regardless of cache state
	configs, err := cache.refreshSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs, 2)
	assert.Equal(t, cciptypes.UnknownAddress{1, 2, 3}, configs[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{4, 5, 6}, configs[chainC].OnRamp)

	// Setup mock for a second call with different data
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result3.SetResult(&SourceChainConfig{IsEnabled: false, OnRamp: cciptypes.UnknownAddress{7, 8, 9}}, nil)
	result4 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result4.SetResult(&SourceChainConfig{IsEnabled: false, OnRamp: cciptypes.UnknownAddress{10, 11, 12}}, nil)

	response2 := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result3, *result4,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response2, []string{}, nil).Once()

	// Force refresh again, should fetch new data
	configs2, err := cache.refreshSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs2, 2)
	assert.NotEqual(t, configs, configs2)
	assert.Equal(t, cciptypes.UnknownAddress{7, 8, 9}, configs2[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{10, 11, 12}, configs2[chainC].OnRamp)

	// Getting from cache now should give the refreshed values
	configs3, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	assert.Equal(t, configs2, configs3)

	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 2)
}

func TestConfigCache_GetOfframpSourceChainConfigs_Error(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	sourceChains := []cciptypes.ChainSelector{chainB, chainC}

	// First fetch - one chain succeeds, one fails with error
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)

	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(nil, errors.New("read error"))

	response := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response, []string{}, nil).Once()

	// Should fail with error from the failing chain
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)

	// Expect an error related to the failing chain
	require.Error(t, err)
	assert.Contains(t, err.Error(), "read error")

	// Result should be empty due to the error
	assert.Empty(t, configs)
}

func TestConfigCache_GlobalSourceChainRefreshTime(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// First set of chains to request
	sourceChains1 := []cciptypes.ChainSelector{chainB, chainC}

	// Setup mock response for first fetch
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	firstResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.MatchedBy(func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
			batch, ok := req[consts.ContractNameOffRamp]
			return ok && len(batch) == 2
		}),
		false,
	).Return(firstResponse, []string{}, nil).Once()

	// First call should fetch B and C
	configs1, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains1)
	require.NoError(t, err)
	require.Len(t, configs1, 2)

	// Request for a new chain D - should only fetch D
	sourceChains2 := []cciptypes.ChainSelector{chainD}

	// Setup mock for D fetch only
	resultD := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	resultD.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{7, 8, 9}}, nil)

	secondResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*resultD,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.MatchedBy(func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
			batch, ok := req[consts.ContractNameOffRamp]
			return ok && len(batch) == 1
		}),
		false,
	).Return(secondResponse, []string{}, nil).Once()

	// Should only fetch D
	configs2, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains2)
	require.NoError(t, err)
	require.Len(t, configs2, 1)

	// Chain D should be newly fetched
	assert.Equal(t, cciptypes.UnknownAddress{7, 8, 9}, configs2[chainD].OnRamp)

	// Request all chains (B, C, D) - should use cache for all of them
	sourceChains3 := []cciptypes.ChainSelector{chainB, chainC, chainD}

	// Setup mock for manual refresh of all chains
	resultB2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	resultB2.SetResult(&SourceChainConfig{IsEnabled: false, OnRamp: cciptypes.UnknownAddress{10, 11, 12}}, nil)
	resultC2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	resultC2.SetResult(&SourceChainConfig{IsEnabled: false, OnRamp: cciptypes.UnknownAddress{13, 14, 15}}, nil)
	resultD2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	resultD2.SetResult(&SourceChainConfig{IsEnabled: false, OnRamp: cciptypes.UnknownAddress{16, 17, 18}}, nil)

	thirdResponse := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*resultB2, *resultC2, *resultD2,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.MatchedBy(func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
			batch, ok := req[consts.ContractNameOffRamp]
			return ok && len(batch) == 3
		}),
		false,
	).Return(thirdResponse, []string{}, nil).Once()

	// First get all chains from cache (should be all present)
	cached, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains3)
	require.NoError(t, err)
	require.Len(t, cached, 3)

	// Now manually refresh all chains (simulating background refresh)
	_, err = cache.refreshSourceChainConfigs(ctx, chainA, sourceChains3)
	require.NoError(t, err)

	// Get all chains again - should now have the updated values
	configs3, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains3)
	require.NoError(t, err)
	require.Len(t, configs3, 3)

	// Verify all configs have been refreshed with the new data
	assert.Equal(t, cciptypes.UnknownAddress{10, 11, 12}, configs3[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{13, 14, 15}, configs3[chainC].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{16, 17, 18}, configs3[chainD].OnRamp)
	assert.False(t, configs3[chainB].IsEnabled)
	assert.False(t, configs3[chainC].IsEnabled)
	assert.False(t, configs3[chainD].IsEnabled)

	// Verify the calls were made as expected
	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 3)
}

func TestConfigCache_GetOrCreateChainCache_InitializesSourceChainConfig(t *testing.T) {
	cache, _ := setupBasicCache(t)

	// Get a cache entry
	chainCache := cache.getOrCreateChainCache(chainA)
	require.NotNil(t, chainCache)

	// Verify the sourceChainConfigs map is initialized
	require.NotNil(t, chainCache.staticSourceChainConfigs)
	assert.Len(t, chainCache.staticSourceChainConfigs, 0)

	// Verify the initial refresh time is zero
	assert.True(t, chainCache.sourceChainRefresh.IsZero())
}

func TestConfigCache_RefreshSourceChainConfigs_SetsGlobalTimestamp(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	sourceChains := []cciptypes.ChainSelector{chainB}

	// Setup mock response
	result := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)

	response := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response, []string{}, nil).Once()

	// Before refresh, get the chainCache
	chainCache := cache.getOrCreateChainCache(chainA)
	initialRefreshTime := chainCache.sourceChainRefresh
	assert.True(t, initialRefreshTime.IsZero())

	// Refresh the source chain configs
	configs, err := cache.refreshSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs, 1)

	// Verify the global timestamp was updated
	updatedRefreshTime := chainCache.sourceChainRefresh
	assert.False(t, updatedRefreshTime.IsZero())
	assert.True(t, updatedRefreshTime.After(initialRefreshTime))
}

func TestConfigPoller_GetChainsToRefresh(t *testing.T) {
	// Setup test environment with destination chain A
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// We need to first populate the cache for chain A by making a call to GetChainConfig
	setupMockResponse(reader)
	_, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err, "Failed to populate cache for chain A")

	// Add a second chain (chain B) to the cache
	_ = setupSecondChain(ctx, t, cache)

	// Track source chains B and C for destination chain A
	sourceChains := []cciptypes.ChainSelector{chainB, chainC}
	for _, chain := range sourceChains {
		success := cache.trackSourceChain(chainA, chain)
		require.True(t, success, "Failed to track source chain %d for dest chain %d", chain, chainA)
	}

	// Try to track a source chain D for chain B (which is not the destination chain)
	// This should be ignored by getChainsToRefresh since B is not the destination
	success := cache.trackSourceChain(chainB, chainD)
	require.True(t, success, "Failed to track chain D as source for chain B")

	// Call the method we want to test
	chains, sourceChainsMap := cache.getChainsToRefresh()

	// Verify all chains in cache are returned
	require.Len(t, chains, 2, "Should return both chains in the cache")
	assert.Contains(t, chains, chainA, "Chain A should be included")
	assert.Contains(t, chains, chainB, "Chain B should be included")

	// Verify only source chains for the destination chain are included
	require.Len(t, sourceChainsMap, 1, "Should only have source chains for the destination chain")
	assert.Contains(t, sourceChainsMap, chainA, "Destination chain should be in the map")

	// Verify the source chains for destination chain are correct
	destSourceChains := sourceChainsMap[chainA]
	require.Len(t, destSourceChains, 2, "Should have 2 source chains for destination")
	assert.Contains(t, destSourceChains, chainB, "Chain B should be in source chains")
	assert.Contains(t, destSourceChains, chainC, "Chain C should be in source chains")

	// Verify chain B's source chains are not included since it's not the destination chain
	assert.NotContains(t, sourceChainsMap, chainB, "Non-destination chains should not have source chains in the map")

	// Test edge case: Clear the source chains map and verify empty result
	cache.Lock()
	cache.knownSourceChains = make(map[cciptypes.ChainSelector]map[cciptypes.ChainSelector]bool)
	cache.Unlock()

	chains2, sourceChainsMap2 := cache.getChainsToRefresh()
	require.Len(t, chains2, 2, "Should still return all chains in cache")
	assert.Contains(t, chains2, chainA, "Chain A should still be included")
	assert.Contains(t, chains2, chainB, "Chain B should still be included")
	assert.Empty(t, sourceChainsMap2, "Source chains map should be empty when none tracked")
}
