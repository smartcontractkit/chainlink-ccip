package reader

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	rmnHomeBoundContract = types.BoundContract{
		Address: "0xRMNHomeFakeAddress",
		Name:    consts.ContractNameRMNHome,
	}
)

func TestRMNHomeChainConfigPoller_Ready(t *testing.T) {
	homeChainReader := readermock.NewMockContractReaderFacade(t)
	configPoller := newRMNHomePoller(
		homeChainReader,
		rmnHomeBoundContract,
		logger.Test(t),
		1*time.Millisecond,
	)
	// Return any result as we are testing the ready method
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(fmt.Errorf("error"))

	// Initially it's not ready
	require.Error(t, configPoller.Ready())

	require.NoError(t, configPoller.Start(context.Background()))
	// After starting it's ready
	require.NoError(t, configPoller.Ready())

	require.NoError(t, configPoller.Close())
}

func TestRMNHomePoller_HealthReport(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		failedPolls uint
		wantErr     bool
	}{
		{
			name:        "Healthy state",
			failedPolls: 0,
			wantErr:     false,
		},
		{
			name:        "Unhealthy state",
			failedPolls: MaxFailedPolls,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			homeChainReader := readermock.NewMockContractReaderFacade(t)

			homeChainReader.On("GetLatestValue",
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Run(func(args mock.Arguments) {
				result := args.Get(4).(*GetAllConfigsResponse)
				*result = GetAllConfigsResponse{
					ActiveConfig: VersionedConfig{
						ConfigDigest:  [32]byte{1},
						Version:       1,
						StaticConfig:  StaticConfig{Nodes: []Node{}},
						DynamicConfig: DynamicConfig{SourceChains: []SourceChain{}},
					},
				}
			}).Return(nil)

			poller := newRMNHomePoller(
				homeChainReader,
				rmnHomeBoundContract,
				logger.Test(t),
				10*time.Millisecond,
			)

			require.NoError(t, poller.Start(context.Background()))

			// Wait for the initial fetch to complete
			require.Eventually(t, func() bool {
				return homeChainReader.AssertCalled(
					t,
					"GetLatestValue",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything)
			}, 5*time.Second, 10*time.Millisecond, "GetLatestValue was not called within the expected timeframe")

			poller.mutex.Lock()
			poller.failedPolls = tt.failedPolls
			poller.mutex.Unlock()

			report := poller.HealthReport()

			require.Len(t, report, 1)
			err := report[poller.Name()]

			if tt.wantErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), "polling failed")
			} else {
				require.NoError(t, err)
			}

			require.NoError(t, poller.Close())
		})
	}
}

func Test_RMNHomePollingWorking(t *testing.T) {
	tests := []struct {
		name              string
		primaryEmpty      bool
		secondaryEmpty    bool
		expectedCallCount int
	}{
		{
			name:              "Both configs non-empty",
			primaryEmpty:      false,
			secondaryEmpty:    false,
			expectedCallCount: 4,
		},
		{
			name:              "Primary config empty",
			primaryEmpty:      true,
			secondaryEmpty:    false,
			expectedCallCount: 4,
		},
		{
			name:              "Secondary config empty",
			primaryEmpty:      false,
			secondaryEmpty:    true,
			expectedCallCount: 4,
		},
		{
			name:              "Both configs empty",
			primaryEmpty:      true,
			secondaryEmpty:    true,
			expectedCallCount: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			primaryConfig, secondaryConfig := createTestRMNHomeConfigs(tt.primaryEmpty, tt.secondaryEmpty)
			rmnHomeOnChainConfigs := GetAllConfigsResponse{
				primaryConfig,
				secondaryConfig,
			}

			homeChainReader := readermock.NewMockContractReaderFacade(t)
			homeChainReader.On(
				"GetLatestValue",
				mock.Anything,
				rmnHomeBoundContract.ReadIdentifier(consts.MethodNameGetAllConfigs),
				mock.Anything,
				mock.Anything,
				mock.Anything,
			).Run(
				func(args mock.Arguments) {
					arg := args.Get(4).(*GetAllConfigsResponse)
					*arg = rmnHomeOnChainConfigs
				}).Return(nil)

			defer homeChainReader.AssertExpectations(t)

			var (
				tickTime       = 2 * time.Millisecond
				totalSleepTime = tickTime * 20
			)

			configPoller := newRMNHomePoller(
				homeChainReader,
				rmnHomeBoundContract,
				logger.Test(t),
				tickTime,
			)

			require.NoError(t, configPoller.Start(context.Background()))
			// sleep to allow polling to happen
			time.Sleep(totalSleepTime)
			require.NoError(t, configPoller.Close())

			calls := homeChainReader.Calls
			callCount := 0
			for _, call := range calls {
				if call.Method == "GetLatestValue" {
					callCount++
				}
			}
			require.GreaterOrEqual(t, callCount, tt.expectedCallCount)

			for i, config := range []VersionedConfig{primaryConfig, secondaryConfig} {
				isEmpty := (i == 0 && tt.primaryEmpty) || (i == 1 && tt.secondaryEmpty)

				rmnNodes, err := configPoller.GetRMNNodesInfo(config.ConfigDigest)
				if isEmpty {
					require.Error(t, err)
					require.Empty(t, rmnNodes)
				} else {
					require.NoError(t, err)
					require.NotEmpty(t, rmnNodes)
				}

				isValid := configPoller.IsRMNHomeConfigDigestSet(config.ConfigDigest)
				if isEmpty {
					require.False(t, isValid)
				} else {
					require.True(t, isValid)
				}

				offchainConfig, err := configPoller.GetOffChainConfig(config.ConfigDigest)
				if isEmpty {
					require.Error(t, err)
					require.Empty(t, offchainConfig)
				} else {
					require.NoError(t, err)
					require.NotEmpty(t, offchainConfig)
				}

				minObsMap, err := configPoller.GetFObserve(config.ConfigDigest)
				if isEmpty {
					require.Error(t, err)
					require.Empty(t, minObsMap)
				} else {
					require.NoError(t, err)
					require.Len(t, minObsMap, 1)
					expectedChainSelector := cciptypes.ChainSelector(uint64(i + 1))
					minObs, exists := minObsMap[expectedChainSelector]
					require.True(t, exists)
					require.Equal(t, i+1, minObs)
				}

				activeConfigDigest, candidateConfigDigest := configPoller.GetAllConfigDigests()
				require.Equal(t, primaryConfig.ConfigDigest, activeConfigDigest)
				require.Equal(t, secondaryConfig.ConfigDigest, candidateConfigDigest)

			}
		})
	}
}

func Test_RMNHomePoller_Close(t *testing.T) {
	homeChainReader := readermock.NewMockContractReaderFacade(t)
	poller := newRMNHomePoller(
		homeChainReader,
		rmnHomeBoundContract,
		logger.Test(t),
		10*time.Millisecond,
	)

	homeChainReader.On("GetLatestValue",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(func(args mock.Arguments) {
		result := args.Get(4).(*GetAllConfigsResponse)
		*result = GetAllConfigsResponse{
			ActiveConfig: VersionedConfig{
				ConfigDigest:  [32]byte{1},
				Version:       1,
				StaticConfig:  StaticConfig{Nodes: []Node{}},
				DynamicConfig: DynamicConfig{SourceChains: []SourceChain{}},
			},
		}
	}).Return(nil)

	ctx := tests.Context(t)
	require.NoError(t, poller.Start(ctx))

	err1 := poller.Close()
	require.NoError(t, err1)
	err2 := poller.Close()
	require.NoError(t, err2)
	err3 := poller.Close()
	require.NoError(t, err3)
}

func TestIsNodeObserver(t *testing.T) {
	tests := []struct {
		name           string
		sourceChain    SourceChain
		nodeIndex      int
		totalNodes     int
		expectedResult bool
		expectedError  string
	}{
		{
			name: "Node is observer",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            3,
				ObserverNodesBitmap: big.NewInt(7), // 111 in binary
			},
			nodeIndex:      1,
			totalNodes:     3,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name: "Node is not observer",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            3,
				ObserverNodesBitmap: big.NewInt(5), // 101 in binary
			},
			nodeIndex:      1,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "",
		},
		{
			name: "Node index out of range (high)",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            3,
				ObserverNodesBitmap: big.NewInt(7), // 111 in binary
			},
			nodeIndex:      3,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "invalid node index: 3",
		},
		{
			name: "Negative node index",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            3,
				ObserverNodesBitmap: big.NewInt(7), // 111 in binary
			},
			nodeIndex:      -1,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "invalid node index: -1",
		},
		{
			name: "Invalid bitmap (out of bounds)",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            3,
				ObserverNodesBitmap: big.NewInt(8), // 1000 in binary
			},
			nodeIndex:      0,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "invalid observer nodes bitmap",
		},
		{
			name: "Zero total nodes",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            3,
				ObserverNodesBitmap: big.NewInt(1),
			},
			nodeIndex:      0,
			totalNodes:     0,
			expectedResult: false,
			expectedError:  "invalid total nodes: 0",
		},
		{
			name: "Total nodes exceeds 256",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            3,
				ObserverNodesBitmap: big.NewInt(1),
			},
			nodeIndex:      0,
			totalNodes:     257,
			expectedResult: false,
			expectedError:  "invalid total nodes: 257",
		},
		{
			name: "Last valid node is observer",
			sourceChain: SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				FObserve:            1,
				ObserverNodesBitmap: new(big.Int).SetBit(big.NewInt(0), 255, 1), // Only the 256th bit is set
			},
			nodeIndex:      255,
			totalNodes:     256,
			expectedResult: true,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IsNodeObserver(tt.sourceChain, tt.nodeIndex, tt.totalNodes)

			if tt.expectedError != "" {
				require.Error(t, err)
				require.Equal(t, tt.expectedError, err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func createTestRMNHomeConfigs(
	primaryEmpty bool,
	secondaryEmpty bool) (primary, secondary VersionedConfig) {
	createConfig := func(id byte, isEmpty bool) VersionedConfig {
		if isEmpty {
			return VersionedConfig{}
		}
		return VersionedConfig{
			ConfigDigest: cciptypes.Bytes32{id},
			Version:      uint32(id),
			DynamicConfig: DynamicConfig{
				SourceChains: []SourceChain{
					{
						ChainSelector:       cciptypes.ChainSelector(id),
						FObserve:            uint64(id),
						ObserverNodesBitmap: big.NewInt(int64(id)),
					},
				},
				OffchainConfig: cciptypes.Bytes{30 * id},
			},
			StaticConfig: StaticConfig{
				Nodes: []Node{
					{
						PeerID:            cciptypes.Bytes32{10 * id},
						OffchainPublicKey: cciptypes.Bytes32{20 * id},
					},
				},
				OffchainConfig: cciptypes.Bytes{30 * id},
			},
		}
	}

	primary = createConfig(1, primaryEmpty)
	secondary = createConfig(2, secondaryEmpty)
	return primary, secondary
}

func Test_CachingInstances(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)

	chain1 := readermock.NewMockContractReaderFacade(t)
	chain2 := readermock.NewMockContractReaderFacade(t)

	for _, chain := range []*readermock.MockContractReaderFacade{chain1, chain2} {
		chain.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil).Maybe()
		chain.EXPECT().GetLatestValue(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
	}

	t.Run("reusing instance for the same chain and address", func(t *testing.T) {
		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()
		poller1, err := GetRMNHomePoller(ctx, lggr, chainSelector, address, chain1, HomeChainPollingInterval)
		require.NoError(t, err)

		poller2, err := GetRMNHomePoller(ctx, lggr, chainSelector, address, chain1, HomeChainPollingInterval)
		require.NoError(t, err)

		poller3, err := GetRMNHomePoller(ctx, lggr, chainSelector, address, chain1, HomeChainPollingInterval)
		require.NoError(t, err)

		require.True(t, poller1 == poller2)
		require.True(t, poller2 == poller3)

		require.NoError(t, poller1.Close())
		require.NoError(t, poller2.Close())
		require.NoError(t, poller3.Close())
	})

	t.Run("creating new instance for different addresses on a single chain", func(t *testing.T) {
		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address1 := rand.RandomAddressBytes()
		address2 := rand.RandomAddressBytes()

		poller1, err := GetRMNHomePoller(ctx, lggr, chainSelector, address1, chain1, HomeChainPollingInterval)
		require.NoError(t, err)

		poller2, err := GetRMNHomePoller(ctx, lggr, chainSelector, address2, chain1, HomeChainPollingInterval)
		require.NoError(t, err)

		require.False(t, poller1 == poller2)
		require.NoError(t, poller1.Close())
		require.NoError(t, poller2.Close())
	})

	t.Run("creating new instance for different chains but same addresses", func(t *testing.T) {
		chainSelector1 := cciptypes.ChainSelector(rand.RandomInt64())
		chainSelector2 := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()

		poller1, err := GetRMNHomePoller(ctx, lggr, chainSelector1, address, chain1, HomeChainPollingInterval)
		require.NoError(t, err)

		poller2, err := GetRMNHomePoller(ctx, lggr, chainSelector2, address, chain2, HomeChainPollingInterval)
		require.NoError(t, err)

		require.False(t, poller1 == poller2)
		require.NoError(t, poller1.Close())
		require.NoError(t, poller2.Close())
	})

	t.Run("parallel creation of instances doesn't cause any failures", func(t *testing.T) {
		instancesMu.Lock()
		instances = make(map[string]*rmnHomePoller)
		instancesMu.Unlock()

		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()

		eg := new(errgroup.Group)
		for i := 0; i < 1000; i++ {
			eg.Go(func() error {
				poller, err := GetRMNHomePoller(ctx, lggr, chainSelector, address, chain1, HomeChainPollingInterval)
				require.NotNil(t, poller)
				require.NoError(t, err)
				return nil
			})
		}
		require.NoError(t, eg.Wait())
		require.Len(t, instances, 1)

		poller, err := GetRMNHomePoller(ctx, lggr, chainSelector, address, chain1, HomeChainPollingInterval)
		require.NoError(t, err)
		require.NoError(t, poller.Close())
	})

	t.Run("create new instance when old one is already stopped", func(t *testing.T) {
		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()

		poller1, err := GetRMNHomePoller(ctx, lggr, chainSelector, address, chain1, HomeChainPollingInterval)
		require.NoError(t, err)
		require.NoError(t, poller1.Close())
		require.Error(t, poller1.Start(ctx))

		poller2, err := GetRMNHomePoller(ctx, lggr, chainSelector, address, chain1, HomeChainPollingInterval)
		require.NoError(t, err)

		castedPoller, ok := poller2.(*rmnHomePoller)
		require.True(t, ok)
		require.NoError(t, castedPoller.Ready())

		require.NoError(t, poller2.Close())
	})
}
