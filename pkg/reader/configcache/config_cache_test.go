package configcache

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
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

	testLogger := logger.Test(t)

	cache := NewConfigCache(mockReader, testLogger).(*configCache)

	return cache, mockReader
}

func TestConfigCache_Refresh_BasicScenario(t *testing.T) {
	cache, mockReader := setupConfigCacheTest(t)

	// Create sample test data for all contract types
	expectedResults := contractreader.BatchGetLatestValuesGracefulResult{
		Results: types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.Bytes{1, 2, 3}), // native token address
			},
			types.BoundContract{Name: consts.ContractNameOnRamp}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.GetOnRampDynamicConfigResponse{
					DynamicConfig: cciptypes.OnRampDynamicConfig{
						FeeQuoter:              []byte{4, 5, 6},
						ReentrancyGuardEntered: false,
						MessageInterceptor:     []byte{7, 8, 9},
						FeeAggregator:          []byte{10, 11, 12},
						AllowListAdmin:         []byte{13, 14, 15},
					},
				}),
			},
			types.BoundContract{Name: consts.ContractNameOffRamp}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.OCRConfigResponse{ // commit config
					OCRConfig: cciptypes.OCRConfig{
						ConfigInfo: cciptypes.ConfigInfo{
							ConfigDigest:                   [32]byte{1},
							F:                              uint8(1),
							N:                              uint8(3),
							IsSignatureVerificationEnabled: true,
						},
						Signers:      [][]byte{{1}, {2}, {3}},
						Transmitters: [][]byte{{4}, {5}, {6}},
					},
				}),
				createMockBatchResult(&cciptypes.OCRConfigResponse{ // execute config
					OCRConfig: cciptypes.OCRConfig{
						ConfigInfo: cciptypes.ConfigInfo{
							ConfigDigest:                   [32]byte{2},
							F:                              uint8(2),
							N:                              uint8(4),
							IsSignatureVerificationEnabled: true,
						},
						Signers:      [][]byte{{7}, {8}, {9}, {10}},
						Transmitters: [][]byte{{11}, {12}, {13}, {14}},
					},
				}),
				createMockBatchResult(&cciptypes.OffRampStaticChainConfig{
					ChainSelector:        200,
					GasForCallExactCheck: 10000,
					RmnRemote:            []byte{15, 16, 17},
					TokenAdminRegistry:   []byte{18, 19, 20},
					NonceManager:         []byte{21, 22, 23},
				}),
				createMockBatchResult(&cciptypes.OffRampDynamicChainConfig{
					FeeQuoter:                               []byte{24, 25, 26},
					PermissionLessExecutionThresholdSeconds: 3600,
					IsRMNVerificationDisabled:               false,
					MessageInterceptor:                      []byte{27, 28, 29},
				}),
				createMockBatchResult(&cciptypes.SelectorsAndConfigs{
					Selectors: []uint64{1, 2, 3},
					SourceChainConfigs: []cciptypes.SourceChainConfig{
						{
							Router:    []byte{30, 31, 32},
							IsEnabled: true,
							MinSeqNr:  100,
							OnRamp:    []byte{33, 34, 35},
						},
					},
				}),
			},
			types.BoundContract{Name: consts.ContractNameRMNRemote}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.RMNDigestHeader{
					DigestHeader: [32]byte{3},
				}),
				createMockBatchResult(&cciptypes.VersionedConfigRemote{
					Version: 1,
					Config: cciptypes.Config{
						RMNHomeContractConfigDigest: [32]byte{4},
						Signers: []cciptypes.Signer{
							{
								OnchainPublicKey: []byte{36, 37, 38},
								NodeIndex:        1,
							},
							{
								OnchainPublicKey: []byte{39, 40, 41},
								NodeIndex:        2,
							},
						},
						F: 1,
					},
				}),
			},
			types.BoundContract{Name: consts.ContractNameRMNProxy}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.Bytes{42, 43, 44}), // remote address
			},
			types.BoundContract{Name: consts.ContractNameFeeQuoter}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.FeeQuoterStaticConfig{
					MaxFeeJuelsPerMsg:  cciptypes.BigInt{Int: big.NewInt(1000000)},
					LinkToken:          []byte{45, 46, 47},
					StalenessThreshold: 300,
				}),
			},
		},
		SkippedNoBinds: []string{}, // No skipped contracts
	}

	// Setup mock to return our test data
	mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
		mock.Anything,
		mock.Anything,
	).Return(expectedResults, nil)

	// Force refresh to populate cache
	err := cache.refresh(tests.Context(t))
	require.NoError(t, err)

	// Verify Router values
	addr, err := cache.GetNativeTokenAddress(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, cciptypes.Bytes{1, 2, 3}, addr)

	// Verify OnRamp values
	onRampConfig, err := cache.GetOnRampDynamicConfig(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, []byte{4, 5, 6}, onRampConfig.DynamicConfig.FeeQuoter)
	assert.Equal(t, []byte{7, 8, 9}, onRampConfig.DynamicConfig.MessageInterceptor)
	assert.Equal(t, []byte{10, 11, 12}, onRampConfig.DynamicConfig.FeeAggregator)
	assert.Equal(t, []byte{13, 14, 15}, onRampConfig.DynamicConfig.AllowListAdmin)
	assert.False(t, onRampConfig.DynamicConfig.ReentrancyGuardEntered)

	// Verify OffRamp values
	commitDigest, err := cache.GetOffRampConfigDigest(tests.Context(t), consts.PluginTypeCommit)
	require.NoError(t, err)
	assert.Equal(t, [32]byte{1}, commitDigest)

	execDigest, err := cache.GetOffRampConfigDigest(tests.Context(t), consts.PluginTypeExecute)
	require.NoError(t, err)
	assert.Equal(t, [32]byte{2}, execDigest)

	staticConfig, err := cache.GetOffRampStaticConfig(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, uint16(10000), staticConfig.GasForCallExactCheck)
	assert.Equal(t, []byte{15, 16, 17}, staticConfig.RmnRemote)
	assert.Equal(t, []byte{18, 19, 20}, staticConfig.TokenAdminRegistry)
	assert.Equal(t, []byte{21, 22, 23}, staticConfig.NonceManager)

	dynamicConfig, err := cache.GetOffRampDynamicConfig(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, []byte{24, 25, 26}, dynamicConfig.FeeQuoter)
	assert.Equal(t, uint32(3600), dynamicConfig.PermissionLessExecutionThresholdSeconds)
	assert.False(t, dynamicConfig.IsRMNVerificationDisabled)
	assert.Equal(t, []byte{27, 28, 29}, dynamicConfig.MessageInterceptor)

	allChains, err := cache.GetOffRampAllChains(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, []uint64{1, 2, 3}, allChains.Selectors)
	require.Len(t, allChains.SourceChainConfigs, 1)
	assert.Equal(t, []byte{30, 31, 32}, allChains.SourceChainConfigs[0].Router)
	assert.True(t, allChains.SourceChainConfigs[0].IsEnabled)
	assert.Equal(t, uint64(100), allChains.SourceChainConfigs[0].MinSeqNr)
	assert.Equal(t, cciptypes.UnknownAddress([]byte{33, 34, 35}), allChains.SourceChainConfigs[0].OnRamp)

	// Verify RMN values
	digestHeader, err := cache.GetRMNDigestHeader(tests.Context(t))
	require.NoError(t, err)
	expectedDigestHeader := cciptypes.Bytes32{3} // Using Bytes32 type directly
	assert.Equal(t, expectedDigestHeader, digestHeader.DigestHeader)

	versionedConfig, err := cache.GetRMNVersionedConfig(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, uint32(1), versionedConfig.Version)
	expectedHomeDigest := cciptypes.Bytes32{4} // Using Bytes32 type directly
	assert.Equal(t, expectedHomeDigest, versionedConfig.Config.RMNHomeContractConfigDigest)
	require.Len(t, versionedConfig.Config.Signers, 2)
	assert.Equal(t, []byte{36, 37, 38}, versionedConfig.Config.Signers[0].OnchainPublicKey)
	assert.Equal(t, uint64(1), versionedConfig.Config.Signers[0].NodeIndex)
	assert.Equal(t, []byte{39, 40, 41}, versionedConfig.Config.Signers[1].OnchainPublicKey)
	assert.Equal(t, uint64(2), versionedConfig.Config.Signers[1].NodeIndex)
	assert.Equal(t, uint64(1), versionedConfig.Config.F)

	rmnRemote, err := cache.GetRMNRemoteAddress(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, cciptypes.Bytes{42, 43, 44}, rmnRemote)

	// Verify FeeQuoter values
	feeQuoterConfig, err := cache.GetFeeQuoterConfig(tests.Context(t))
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1000000), feeQuoterConfig.MaxFeeJuelsPerMsg.Int)
	assert.Equal(t, []byte{45, 46, 47}, feeQuoterConfig.LinkToken)
	assert.Equal(t, uint32(300), feeQuoterConfig.StalenessThreshold)
}

func TestConfigCache_ConcurrentAccess(t *testing.T) {
	cache, mockReader := setupConfigCacheTest(t)

	// Setup expected response
	mockBatchResults := contractreader.BatchGetLatestValuesGracefulResult{
		Results: types.BatchGetLatestValuesResult{
			types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.Bytes{1, 2, 3}),
			},
			types.BoundContract{Name: consts.ContractNameOnRamp}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.GetOnRampDynamicConfigResponse{
					DynamicConfig: cciptypes.OnRampDynamicConfig{
						FeeQuoter: []byte{4, 5, 6},
					},
				}),
			},
			types.BoundContract{Name: consts.ContractNameOffRamp}: []types.BatchReadResult{
				createMockBatchResult(&cciptypes.OCRConfigResponse{
					OCRConfig: cciptypes.OCRConfig{
						ConfigInfo: cciptypes.ConfigInfo{
							ConfigDigest: [32]byte{1},
						},
					},
				}),
			},
		},
	}

	// We expect only one call because of caching
	mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
		mock.Anything,
		mock.Anything,
	).Return(mockBatchResults, nil).Once()

	// Run multiple goroutines accessing different methods
	var wg sync.WaitGroup
	numGoroutines := 10
	wg.Add(numGoroutines * 3)

	// Use this channel to coordinate start time
	start := make(chan struct{})

	// Launch goroutines
	for i := 0; i < numGoroutines; i++ {
		// Test different getter methods concurrently
		go func() {
			defer wg.Done()
			<-start // Wait for start signal
			addr, err := cache.GetNativeTokenAddress(tests.Context(t))
			require.NoError(t, err)
			assert.Equal(t, cciptypes.Bytes{1, 2, 3}, addr)
		}()

		go func() {
			defer wg.Done()
			<-start // Wait for start signal
			config, err := cache.GetOnRampDynamicConfig(tests.Context(t))
			require.NoError(t, err)
			assert.Equal(t, []byte{4, 5, 6}, config.DynamicConfig.FeeQuoter)
		}()

		go func() {
			defer wg.Done()
			<-start // Wait for start signal
			digest, err := cache.GetOffRampConfigDigest(tests.Context(t), consts.PluginTypeCommit)
			require.NoError(t, err)
			assert.Equal(t, [32]byte{1}, digest)
		}()
	}

	// Start all goroutines simultaneously
	close(start)

	// Wait for all goroutines to complete
	wg.Wait()

	// Verify the cache worked by checking we only got one mock call
	mockReader.AssertExpectations(t)
}

func TestConfigCache_RefreshInterval(t *testing.T) {
	t.Run("respects refresh interval", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// Setup initial call
		mockBatchResults := contractreader.BatchGetLatestValuesGracefulResult{
			Results: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
					createMockBatchResult(&cciptypes.Bytes{1, 2, 3}),
				},
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil).Once()

		// First call should trigger refresh
		addr, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.Bytes{1, 2, 3}, addr)

		// Second immediate call should use cached value without calling mock
		addr2, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, addr, addr2)
	})
}

func TestConfigCache_ErrorHandling(t *testing.T) {
	t.Run("handles reader error", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(contractreader.BatchGetLatestValuesGracefulResult{}, fmt.Errorf("mock error")).Once()

		_, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.Error(t, err)
		assert.Contains(t, err.Error(), "refresh cache")
	})

	t.Run("keeps last valid value when type mismatch occurs", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// First set a valid value
		initialResults := contractreader.BatchGetLatestValuesGracefulResult{
			Results: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
					createMockBatchResult(&cciptypes.Bytes{1, 2, 3}),
				},
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(initialResults, nil).Once()

		// Initial call to set value
		addr, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.Bytes{1, 2, 3}, addr)

		// Force cache expiry
		cache.lastUpdateAt = time.Time{}

		// Now try with invalid type
		invalidResults := contractreader.BatchGetLatestValuesGracefulResult{
			Results: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
					createMockBatchResult("wrong type"), // String instead of Bytes
				},
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(invalidResults, nil).Once()

		// Get value after invalid result - should keep old value
		addr, err = cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.Bytes{1, 2, 3}, addr, "should retain last valid value")
	})

	t.Run("propagates type conversion errors", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// Create a BatchReadResult that will fail GetResult()
		badResult := &types.BatchReadResult{}
		badResult.SetResult(nil, fmt.Errorf("conversion error"))

		mockBatchResults := contractreader.BatchGetLatestValuesGracefulResult{
			Results: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
					*badResult,
				},
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil).Once()

		// Should propagate the error
		_, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.Error(t, err)
		assert.Contains(t, err.Error(), "conversion error")
	})
}

func TestConfigCache_SkippedContractsHandling(t *testing.T) {
	t.Run("clears skipped contract values", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// First response with values
		initialResults := contractreader.BatchGetLatestValuesGracefulResult{
			Results: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: consts.ContractNameRouter}: []types.BatchReadResult{
					createMockBatchResult(&cciptypes.Bytes{1, 2, 3}),
				},
				types.BoundContract{Name: consts.ContractNameOnRamp}: []types.BatchReadResult{
					createMockBatchResult(&cciptypes.GetOnRampDynamicConfigResponse{
						DynamicConfig: cciptypes.OnRampDynamicConfig{
							FeeQuoter: []byte{4, 5, 6},
						},
					}),
				},
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(initialResults, nil).Once()

		// Initial calls to populate cache
		addr, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.Bytes{1, 2, 3}, addr)

		onRampConfig, err := cache.GetOnRampDynamicConfig(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, []byte{4, 5, 6}, onRampConfig.DynamicConfig.FeeQuoter)

		// Second response with router skipped
		updatedResults := contractreader.BatchGetLatestValuesGracefulResult{
			Results: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: consts.ContractNameOnRamp}: []types.BatchReadResult{
					createMockBatchResult(&cciptypes.GetOnRampDynamicConfigResponse{
						DynamicConfig: cciptypes.OnRampDynamicConfig{
							FeeQuoter: []byte{7, 8, 9},
						},
					}),
				},
			},
			SkippedNoBinds: []string{consts.ContractNameRouter},
		}

		// Force cache expiry
		cache.lastUpdateAt = time.Time{}

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(updatedResults, nil).Once()

		// Router value should be cleared
		addr, err = cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.Bytes{}, addr, "router value should be cleared")

		// OnRamp value should be updated
		onRampConfig, err = cache.GetOnRampDynamicConfig(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, []byte{7, 8, 9}, onRampConfig.DynamicConfig.FeeQuoter)
	})
}

func TestConfigCache_EmptyValues(t *testing.T) {
	t.Run("returns zero values for empty results", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		mockBatchResults := contractreader.BatchGetLatestValuesGracefulResult{
			Results: types.BatchGetLatestValuesResult{},
			SkippedNoBinds: []string{
				consts.ContractNameRouter,
				consts.ContractNameOnRamp,
				consts.ContractNameOffRamp,
				consts.ContractNameRMNRemote,
				consts.ContractNameRMNProxy,
				consts.ContractNameFeeQuoter,
			},
		}

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(mockBatchResults, nil).Once()

		// Check all getters return zero values
		addr, err := cache.GetNativeTokenAddress(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.Bytes{}, addr)

		onRampConfig, err := cache.GetOnRampDynamicConfig(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.GetOnRampDynamicConfigResponse{}, onRampConfig)

		offRampConfig, err := cache.GetOffRampStaticConfig(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.OffRampStaticChainConfig{}, offRampConfig)

		feeQuoterConfig, err := cache.GetFeeQuoterConfig(tests.Context(t))
		require.NoError(t, err)
		assert.Equal(t, cciptypes.FeeQuoterStaticConfig{}, feeQuoterConfig)
	})
}

func TestConfigCache_ContextHandling(t *testing.T) {
	t.Run("handles cancelled context", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		// Create cancelled context
		ctx, cancel := context.WithCancel(tests.Context(t))
		cancel()

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(contractreader.BatchGetLatestValuesGracefulResult{}, context.Canceled).Once()

		_, err := cache.GetNativeTokenAddress(ctx)
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("handles deadline exceeded", func(t *testing.T) {
		cache, mockReader := setupConfigCacheTest(t)

		ctx, cancel := context.WithTimeout(tests.Context(t), time.Nanosecond)
		defer cancel()
		time.Sleep(time.Microsecond)

		mockReader.EXPECT().ExtendedBatchGetLatestValuesGraceful(
			mock.Anything,
			mock.Anything,
		).Return(contractreader.BatchGetLatestValuesGracefulResult{}, context.DeadlineExceeded).Once()

		_, err := cache.GetNativeTokenAddress(ctx)
		require.Error(t, err)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}
