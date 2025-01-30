package configcache

import (
	"errors"
	"math/big"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// createMockBatchResult creates a mock batch read result with the given return value
func createMockBatchResult(retVal interface{}) types.BatchReadResult {
	result := &types.BatchReadResult{}
	result.SetResult(retVal, nil)
	return *result
}

func setupConfigCacheTest(t *testing.T) (*configCache, *reader_mocks.MockExtended) {
	mockReader := reader_mocks.NewMockExtended(t)

	cache := NewConfigCache(mockReader).(*configCache)

	return cache, mockReader
}

func TestConfigCache_Refresh(t *testing.T) {
	t.Run("handles partial results in batch response", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// Only return some of the expected results
		mockBatchResults := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.Bytes{1, 2, 3}),
			},
			// Intentionally missing other contracts
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		err := cache.refresh(tests.Context(t))
		require.NoError(t, err)

		// Verify only router config was updated
		addr, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.Bytes{1, 2, 3}, addr)
	})

	t.Run("handles type mismatch in results", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// Return wrong type in result
		mockBatchResults := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
				createMockBatchResult("wrong type"), // String instead of Bytes
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		err := cache.refresh(tests.Context(t))
		require.NoError(t, err) // Should not error but skip invalid result
	})
}

func TestConfigCache_GetOffRampConfigDigest(t *testing.T) {
	t.Run("returns commit config digest", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		expectedDigest := [32]byte{1, 2, 3}
		mockBatchResults := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameOffRamp}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.OCRConfigResponse{
					OCRConfig: cciptypes.OCRConfig{
						ConfigInfo: cciptypes.ConfigInfo{
							ConfigDigest: expectedDigest,
						},
					},
				}),
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		digest, err := cache.GetOffRampConfigDigest(tests.Context(t), consts.PluginTypeCommit)
		require.NoError(t, err)
		assert.Equal(t, expectedDigest, digest)
	})

	t.Run("returns execute config digest", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		expectedDigest := [32]byte{4, 5, 6}
		mockBatchResults := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameOffRamp}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.OCRConfigResponse{}), // commit config
				createMockBatchResult(&cciptypes.OCRConfigResponse{ // execute config
					OCRConfig: cciptypes.OCRConfig{
						ConfigInfo: cciptypes.ConfigInfo{
							ConfigDigest: expectedDigest,
						},
					},
				}),
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		digest, err := cache.GetOffRampConfigDigest(tests.Context(t), consts.PluginTypeExecute)
		require.NoError(t, err)
		assert.Equal(t, expectedDigest, digest)
	})
}

func TestConfigCache_GetRMNVersionedConfig(t *testing.T) {
	cache, mockReader := setupConfigCacheTest(t)

	expectedConfig := cciptypes.VersionedConfigRemote{
		Version: 1,
		Config: cciptypes.Config{
			RMNHomeContractConfigDigest: [32]byte{1, 2, 3},
			Signers: []cciptypes.Signer{
				{
					OnchainPublicKey: []byte{4, 5, 6},
					NodeIndex:        1,
				},
			},
			F: 2,
		},
	}

	mockBatchResults := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameRMNRemote}: []types.BatchReadResult{
			createMockBatchResult(&cciptypes.RMNDigestHeader{}), // digest header
			createMockBatchResult(&expectedConfig),              // versioned config
		},
	}

	mockReader.EXPECT().ExtendedBatchGetLatestValues(
		mock.Anything,
		mock.Anything,
	).Return(mockBatchResults, nil)

	config, err := cache.GetRMNVersionedConfig(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestConfigCache_GetOnRampDynamicConfig(t *testing.T) {
	cache, mockReader := setupConfigCacheTest(t)

	expectedConfig := cciptypes.GetOnRampDynamicConfigResponse{
		DynamicConfig: cciptypes.OnRampDynamicConfig{
			FeeQuoter:          []byte{1, 2, 3},
			AllowListAdmin:     []byte{4, 5, 6},
			MessageInterceptor: []byte{7, 8, 9},
		},
	}

	mockBatchResults := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOnRamp}: []types.BatchReadResult{
			createMockBatchResult(&expectedConfig),
		},
	}

	mockReader.EXPECT().ExtendedBatchGetLatestValues(
		mock.Anything,
		mock.Anything,
	).Return(mockBatchResults, nil)

	config, err := cache.GetOnRampDynamicConfig(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestConfigCache_GetOffRampAllChains(t *testing.T) {
	cache, mockReader := setupConfigCacheTest(t)

	expectedConfig := cciptypes.SelectorsAndConfigs{
		Selectors: []uint64{1, 2, 3},
		SourceChainConfigs: []cciptypes.SourceChainConfig{
			{
				Router:    []byte{1, 2, 3},
				IsEnabled: true,
				MinSeqNr:  100,
			},
			{
				Router:    []byte{4, 5, 6},
				IsEnabled: true,
				MinSeqNr:  200,
			},
		},
	}

	mockBatchResults := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameOffRamp}: []types.BatchReadResult{
			createMockBatchResult(&cciptypes.OCRConfigResponse{}),         // commit config
			createMockBatchResult(&cciptypes.OCRConfigResponse{}),         // execute config
			createMockBatchResult(&cciptypes.OffRampStaticChainConfig{}),  // static config
			createMockBatchResult(&cciptypes.OffRampDynamicChainConfig{}), // dynamic config
			createMockBatchResult(&expectedConfig),                        // all chains config
		},
	}

	mockReader.EXPECT().ExtendedBatchGetLatestValues(
		mock.Anything,
		mock.Anything,
	).Return(mockBatchResults, nil)

	config, err := cache.GetOffRampAllChains(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestConfigCache_ConcurrentAccess(t *testing.T) {
	cache, mockReader := setupConfigCacheTest(t)

	// Setup mock to always return some data
	mockBatchResults := types.BatchGetLatestValuesResult{
		types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
			createMockBatchResult(&cciptypes.Bytes{1, 2, 3}),
		},
	}

	mockReader.EXPECT().ExtendedBatchGetLatestValues(
		mock.Anything,
		mock.Anything,
	).Return(mockBatchResults, nil).Maybe()

	// Run multiple goroutines accessing cache simultaneously
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := cache.GetNativeTokenAddress(tests.Context(t))
			require.NoError(t, err)
		}()
	}
	wg.Wait()
}

func TestConfigCache_UpdateFromResults(t *testing.T) {
	t.Run("handles error getting result", func(t *testing.T) {
		cache, _ := setupConfigCacheTest(t)

		// Create a mock BatchReadResult that returns an error
		errResult := &types.BatchReadResult{}
		errResult.SetResult(nil, errors.New("failed to get result"))

		results := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
				*errResult,
			},
		}

		err := cache.updateFromResults(results)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "get router result")
	})

	t.Run("handles nil results", func(t *testing.T) {
		cache, _ := setupConfigCacheTest(t)

		results := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
				createMockBatchResult(nil),
			},
		}

		err := cache.updateFromResults(results)
		require.NoError(t, err)
	})
}

func TestConfigCache_GetFeeQuoterConfig(t *testing.T) {
	t.Run("returns cached fee quoter config", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		expectedConfig := cciptypes.FeeQuoterStaticConfig{
			MaxFeeJuelsPerMsg:  cciptypes.BigInt{Int: big.NewInt(1000)},
			LinkToken:          []byte{1, 2, 3, 4},
			StalenessThreshold: uint32(300),
		}

		mockBatchResults := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameFeeQuoter}: []types.BatchReadResult{
				createMockBatchResult(&expectedConfig),
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		config, err := cache.GetFeeQuoterConfig(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, expectedConfig.MaxFeeJuelsPerMsg, config.MaxFeeJuelsPerMsg)
		assert.Equal(t, expectedConfig.LinkToken, config.LinkToken)
		assert.Equal(t, expectedConfig.StalenessThreshold, config.StalenessThreshold)
	})

	t.Run("handles invalid fee quoter config", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		invalidConfig := cciptypes.FeeQuoterStaticConfig{
			MaxFeeJuelsPerMsg:  cciptypes.BigInt{Int: big.NewInt(0)}, // Invalid zero max fee
			LinkToken:          []byte{},                             // Invalid empty address
			StalenessThreshold: uint32(0),                            // Invalid zero threshold
		}

		mockBatchResults := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameFeeQuoter}: []types.BatchReadResult{
				createMockBatchResult(&invalidConfig),
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		config, err := cache.GetFeeQuoterConfig(tests.Context(t))
		require.NoError(t, err) // Should not error as validation is not part of the cache
		assert.Equal(t, invalidConfig, config)
	})

	t.Run("handles missing fee quoter config", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// Return empty results
		mockBatchResults := types.BatchGetLatestValuesResult{}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		config, err := cache.GetFeeQuoterConfig(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.FeeQuoterStaticConfig{}, config) // Should return empty config
	})
}

func TestConfigCache_GetRMNRemoteAddress(t *testing.T) {
	t.Run("returns cached RMN remote address", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		expectedAddress := cciptypes.Bytes{1, 2, 3, 4}
		mockBatchResults := types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameRMNProxy}: []types.BatchReadResult{
				createMockBatchResult(&expectedAddress),
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValues(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil)

		address, err := cache.GetRMNRemoteAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, expectedAddress, address)
	})
}
