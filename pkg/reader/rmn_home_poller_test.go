package reader

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

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
	configPoller := newRMNHomeForTests(t, homeChainReader, rmnHomeBoundContract, 1*time.Millisecond)

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
	homeChainReader := readermock.NewMockContractReaderFacade(t)

	mocked := homeChainReader.On("GetLatestValue",
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

	poller := newRMNHomeForTests(
		t,
		homeChainReader,
		rmnHomeBoundContract,
		2*time.Millisecond,
	)
	require.NoError(t, poller.Start(tests.Context(t)))

	require.Eventually(t, func() bool {
		active, _ := poller.GetAllConfigDigests()
		return active == [32]byte{1}
	}, tests.WaitTimeout(t), 10*time.Millisecond)

	mocked.Return(fmt.Errorf("polling failed"))
	require.Eventually(t, func() bool {
		report := poller.HealthReport()
		err := report[poller.bgPoller.Name()]
		return err != nil && strings.Contains(err.Error(), "polling failed")
	}, tests.WaitTimeout(t), 10*time.Millisecond)

	require.NoError(t, poller.Close())
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

			configPoller := newRMNHomeForTests(
				t,
				homeChainReader,
				rmnHomeBoundContract,
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
	poller := newRMNHomeForTests(
		t,
		homeChainReader,
		rmnHomeBoundContract,
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
			result, err := isNodeObserver(tt.sourceChain, tt.nodeIndex, tt.totalNodes)

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

func newRMNHomeForTests(
	t *testing.T,
	reader *readermock.MockContractReaderFacade,
	contract types.BoundContract,
	duration time.Duration,
) *rmnHome {
	poller := newRMNHomePoller(logger.Test(t), reader, contract, duration)
	return &rmnHome{bgPoller: poller}
}
