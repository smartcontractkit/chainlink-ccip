package reader

import (
	"errors"
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

func TestConfigCache_GetChainConfig_CacheMiss(t *testing.T) {
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

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(setupMockBatchResponse(1, 4), []string{}, nil).Once()

	// First call should fetch initial config
	config1, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(1), config1.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)

	// Wait for cache to expire
	time.Sleep(1100 * time.Millisecond)

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(setupMockBatchResponse(2, 6), []string{}, nil).Once()

	// Second call after refresh period should fetch new config
	config2, err := cache.GetChainConfig(ctx, chainA)
	require.NoError(t, err)
	assert.Equal(t, uint8(2), config2.Offramp.CommitLatestOCRConfig.OCRConfig.ConfigInfo.F)
	assert.NotEqual(t, config1, config2)

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

	// Wait for cache to expire
	time.Sleep(1100 * time.Millisecond)

	// Setup error for second fetch attempt
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(nil, nil, errors.New("fetch error"))

	// Second call should return cached data despite error
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
	config, err := cache.RefreshChainConfig(ctx, chainA)
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

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		true,
	).Return(responses, []string{}, nil)

	// Run concurrent fetches
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

	// Should only fetch once despite concurrent access
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

func TestConfigCache_RefreshPeriod(t *testing.T) {
	// Test with different refresh periods
	testCases := []struct {
		name          string
		refreshPeriod time.Duration
		sleepTime     time.Duration
		expectRefresh bool
	}{
		{
			name:          "short refresh period",
			refreshPeriod: 100 * time.Millisecond,
			sleepTime:     150 * time.Millisecond,
			expectRefresh: true,
		},
		{
			name:          "long refresh period",
			refreshPeriod: 1 * time.Second,
			sleepTime:     500 * time.Millisecond,
			expectRefresh: false,
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

			// Setup expected number of calls
			expectedCalls := 1
			if tc.expectRefresh {
				expectedCalls = 2
			}

			mockReader.On("ExtendedBatchGetLatestValues",
				mock.Anything,
				mock.Anything,
				true,
			).Return(responses, []string{}, nil).Times(expectedCalls)

			// First call
			_, err := cache.GetChainConfig(ctx, chainA)
			require.NoError(t, err)

			// Wait
			time.Sleep(tc.sleepTime)

			// Second call
			_, err = cache.GetChainConfig(ctx, chainA)
			require.NoError(t, err)

			// Verify number of calls
			mockReader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", expectedCalls)
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

func TestConfigCache_GetOfframpSourceChainConfigs_CacheMiss(t *testing.T) {
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

	// First call should fetch
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs, 2)

	// Wait for cache to expire
	time.Sleep(1100 * time.Millisecond)

	// Setup mock response for second fetch (with different data)
	result3 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result3.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{7, 8, 9}}, nil)
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

	// Second call after refresh period should fetch new configs
	configs2, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs2, 2)
	assert.NotEqual(t, configs, configs2)
	assert.Equal(t, cciptypes.UnknownAddress{7, 8, 9}, configs2[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{10, 11, 12}, configs2[chainC].OnRamp)
	assert.False(t, configs2[chainC].IsEnabled)

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
	configs, err := cache.RefreshSourceChainConfigs(ctx, chainA, sourceChains)
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
	configs2, err := cache.RefreshSourceChainConfigs(ctx, chainA, sourceChains)
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

	// Setup mock to return an error
	expectedErr := errors.New("fetch error")
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(nil, nil, expectedErr).Once()

	// Should return error on first fetch
	_, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)

	// Setup successful mock response for second attempt
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: []byte{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: []byte{4, 5, 6}}, nil)

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

	// Second call should succeed
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs, 2)

	// Wait for cache to expire
	time.Sleep(1100 * time.Millisecond)

	// Setup error for third fetch after cache is populated
	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(nil, nil, expectedErr).Once()

	// Should return error after cache expired
	_, err = cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)

	reader.AssertNumberOfCalls(t, "ExtendedBatchGetLatestValues", 3)
}

func TestConfigCache_GetOfframpSourceChainConfigs_NoReader(t *testing.T) {
	cache, _ := setupBasicCache(t)
	ctx := tests.Context(t)

	// Test with a chain that has no reader
	_, err := cache.GetOfframpSourceChainConfigs(ctx, chainB, []cciptypes.ChainSelector{chainC})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no contract reader for destination chain")
}

func TestConfigCache_GetOfframpSourceChainConfigs_EmptyChains(t *testing.T) {
	cache, _ := setupBasicCache(t)
	ctx := tests.Context(t)

	// Test with empty source chains slice
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, []cciptypes.ChainSelector{})
	require.NoError(t, err)
	assert.Empty(t, configs)
}

func TestConfigCache_GetOfframpSourceChainConfigs_SkippedContracts(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	sourceChains := []cciptypes.ChainSelector{chainB, chainC}

	// Setup mock response with skipped contracts
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	response := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2,
		},
	}
	skippedContracts := []string{consts.ContractNameRouter}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response, skippedContracts, nil).Once()

	// Should succeed even with skipped contracts
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs, 2)
	assert.Equal(t, cciptypes.UnknownAddress{1, 2, 3}, configs[chainB].OnRamp)
	assert.Equal(t, cciptypes.UnknownAddress{4, 5, 6}, configs[chainC].OnRamp)
}

func TestConfigCache_GetOfframpSourceChainConfigs_InvalidResults(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	sourceChains := []cciptypes.ChainSelector{chainB}

	// Setup mock with invalid result type
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult("invalid type", nil)

	response := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1,
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response, []string{}, nil).Once()

	// Should return an error since the result type was invalid
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)

	// Expect an error about invalid result type
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid result type")

	// Result should be empty due to the error
	assert.Empty(t, configs)
}

func TestConfigCache_FetchPartialResults(t *testing.T) {
	cache, reader := setupBasicCache(t)
	ctx := tests.Context(t)

	// Request for 3 chains
	sourceChains := []cciptypes.ChainSelector{chainB, chainC, chainD}

	// Setup mock to return only 2 results (partial success)
	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result1.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{1, 2, 3}}, nil)
	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetSourceChainConfig}
	result2.SetResult(&SourceChainConfig{IsEnabled: true, OnRamp: cciptypes.UnknownAddress{4, 5, 6}}, nil)

	response := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: {
			*result1, *result2, // Only 2 results for 3 requested chains
		},
	}

	reader.On("ExtendedBatchGetLatestValues",
		mock.Anything,
		mock.Anything,
		false,
	).Return(response, []string{}, nil).Once()

	// Should fail because the number of results doesn't match requested chains
	configs, err := cache.GetOfframpSourceChainConfigs(ctx, chainA, sourceChains)

	// Expect an error about mismatch
	require.Error(t, err)
	assert.Contains(t, err.Error(), "length mismatch")

	// Result should be empty due to the error
	assert.Empty(t, configs)
}

func TestConfigCache_GetOfframpSourceChainConfigs_ErrorHandling(t *testing.T) {
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

	// Wait briefly
	time.Sleep(100 * time.Millisecond)

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
	assert.Equal(t, cciptypes.UnknownAddress{7, 8, 9}, configs2[chainD].OnRamp)

	// Instead of waiting for refresh period, directly force the refresh timestamp to be expired
	chainCache := cache.getOrCreateChainCache(chainA)
	chainCache.sourceChainMu.Lock()
	// Set timestamp to be well in the past (expired)
	chainCache.sourceChainRefresh = time.Now().Add(-2 * time.Second)
	chainCache.sourceChainMu.Unlock()

	// Request all chains (B, C, D) - all should be refreshed due to explicitly expired global timestamp
	sourceChains3 := []cciptypes.ChainSelector{chainB, chainC, chainD}

	// Setup mock for full refresh
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

	// Should refresh all chains due to expired global timestamp
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
	require.NotNil(t, chainCache.sourceChainConfigs)
	assert.Len(t, chainCache.sourceChainConfigs, 0)

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
	configs, err := cache.RefreshSourceChainConfigs(ctx, chainA, sourceChains)
	require.NoError(t, err)
	require.Len(t, configs, 1)

	// Verify the global timestamp was updated
	updatedRefreshTime := chainCache.sourceChainRefresh
	assert.False(t, updatedRefreshTime.IsZero())
	assert.True(t, updatedRefreshTime.After(initialRefreshTime))
}
