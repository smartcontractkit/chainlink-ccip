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

				feeQuoterResult := &types.BatchReadResult{ReadName: consts.MethodNameFeeQuoterGetStaticConfig}
				feeQuoterResult.SetResult(&feeQuoterStaticConfig{}, nil)

				return types.BatchGetLatestValuesResult{
					types.BoundContract{Name: consts.ContractNameOffRamp}:   {},
					types.BoundContract{Name: consts.ContractNameRMNProxy}:  {*rmnProxyResult},
					types.BoundContract{Name: consts.ContractNameRMNRemote}: {*rmnDigestResult, *rmnConfigResult},
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

				feeQuoterResult := &types.BatchReadResult{ReadName: consts.MethodNameFeeQuoterGetStaticConfig}
				feeQuoterResult.SetResult(&feeQuoterStaticConfig{}, nil)

				return types.BatchGetLatestValuesResult{
					types.BoundContract{Name: consts.ContractNameOffRamp}: {
						*result1, *result2, *result3, *result4,
					},
					types.BoundContract{Name: consts.ContractNameRMNProxy}:  {*rmnProxyResult},
					types.BoundContract{Name: consts.ContractNameRMNRemote}: {*rmnDigestResult, *rmnConfigResult},
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
