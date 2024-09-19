package reader

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader/contractreader"
)

var (
	rmnHomeBoundContract = types.BoundContract{
		Address: "0xRMNHomeFakeAddress",
		Name:    consts.ContractNameRMNHome,
	}
)

func TestRMNHomeChainConfigPoller_HealthReport(t *testing.T) {
	homeChainReader := readermock.NewMockContractReaderFacade(t)
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(fmt.Errorf("error"))

	var (
		tickTime       = 1 * time.Millisecond
		totalSleepTime = 50 * time.Millisecond // give enough time for 10 ticks
	)
	configPoller := NewRMNHomePoller(
		homeChainReader,
		rmnHomeBoundContract,
		logger.Test(t),
		tickTime,
	)
	require.NoError(t, configPoller.Start(context.Background()))
	// Initially it's healthy
	healthy := configPoller.HealthReport()
	require.Equal(t, map[string]error{configPoller.Name(): error(nil)}, healthy)
	// After one second it will try polling 10 times and fail
	time.Sleep(totalSleepTime)
	errors := configPoller.HealthReport()
	require.Equal(t, 1, len(errors))
	require.Errorf(t, errors[configPoller.Name()], "polling failed %d times in a row", MaxFailedPolls)
	require.NoError(t, configPoller.Close())
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
			rmnHomeOnChainConfigs := []rmntypes.VersionedConfigWithDigest{primaryConfig, secondaryConfig}

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
					arg := args.Get(4).(*[]rmntypes.VersionedConfigWithDigest)
					*arg = rmnHomeOnChainConfigs
				}).Return(nil)

			defer homeChainReader.AssertExpectations(t)

			var (
				tickTime       = 2 * time.Millisecond
				totalSleepTime = tickTime * 20
			)

			configPoller := NewRMNHomePoller(
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

			for i, config := range rmnHomeOnChainConfigs {
				isEmpty := (i == 0 && tt.primaryEmpty) || (i == 1 && tt.secondaryEmpty)

				rmnNodes, err := configPoller.GetRMNNodesInfo(config.ConfigDigest)
				require.NoError(t, err)
				if isEmpty {
					require.Empty(t, rmnNodes)
				} else {
					require.NotEmpty(t, rmnNodes)
				}

				isValid, err := configPoller.IsRMNHomeConfigDigestSet(config.ConfigDigest)
				require.NoError(t, err)
				if isEmpty {
					require.False(t, isValid)
				} else {
					require.True(t, isValid)
				}

				offchainConfig, err := configPoller.GetOffChainConfig(config.ConfigDigest)
				require.NoError(t, err)
				if isEmpty {
					require.Empty(t, offchainConfig)
				} else {
					require.NotEmpty(t, offchainConfig)
				}

				minObsMap, err := configPoller.GetMinObservers(config.ConfigDigest)
				require.NoError(t, err)
				if isEmpty {
					require.Empty(t, minObsMap)
				} else {
					require.Len(t, minObsMap, 1)
					expectedChainSelector := cciptypes.ChainSelector(uint64(i + 1))
					minObs, exists := minObsMap[expectedChainSelector]
					require.True(t, exists)
					require.Equal(t, uint64(i+1), minObs)
				}
			}
		})
	}
}

func TestIsNodeObserver(t *testing.T) {
	tests := []struct {
		name           string
		sourceChain    rmntypes.SourceChain
		nodeIndex      int
		totalNodes     int
		expectedResult bool
		expectedError  string
	}{
		{
			name: "Node is observer",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        3,
				ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(7)), // 111 in binary
			},
			nodeIndex:      1,
			totalNodes:     3,
			expectedResult: true,
			expectedError:  "",
		},
		{
			name: "Node is not observer",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        3,
				ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(5)), // 101 in binary
			},
			nodeIndex:      1,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "",
		},
		{
			name: "Node index out of range (high)",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        3,
				ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(7)), // 111 in binary
			},
			nodeIndex:      3,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "invalid node index: 3",
		},
		{
			name: "Negative node index",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        3,
				ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(7)), // 111 in binary
			},
			nodeIndex:      -1,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "invalid node index: -1",
		},
		{
			name: "Invalid bitmap (out of bounds)",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        3,
				ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(8)), // 1000 in binary
			},
			nodeIndex:      0,
			totalNodes:     3,
			expectedResult: false,
			expectedError:  "invalid observer nodes bitmap",
		},
		{
			name: "Zero total nodes",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        3,
				ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(1)),
			},
			nodeIndex:      0,
			totalNodes:     0,
			expectedResult: false,
			expectedError:  "invalid total nodes: 0",
		},
		{
			name: "Total nodes exceeds 256",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        3,
				ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(1)),
			},
			nodeIndex:      0,
			totalNodes:     257,
			expectedResult: false,
			expectedError:  "invalid total nodes: 257",
		},
		{
			name: "Last valid node is observer",
			sourceChain: rmntypes.SourceChain{
				ChainSelector:       cciptypes.ChainSelector(1),
				MinObservers:        1,
				ObserverNodesBitmap: cciptypes.NewBigInt(new(big.Int).SetBit(big.NewInt(0), 255, 1)), // Only the 256th bit is set
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
	secondaryEmpty bool) (primary, secondary rmntypes.VersionedConfigWithDigest) {
	createConfig := func(id byte, isEmpty bool) rmntypes.VersionedConfigWithDigest {
		if isEmpty {
			return rmntypes.VersionedConfigWithDigest{}
		}
		return rmntypes.VersionedConfigWithDigest{
			ConfigDigest: cciptypes.Bytes32{id},
			VersionedConfig: rmntypes.VersionedConfig{
				Version: uint32(id),
				Config: rmntypes.Config{
					Nodes: []rmntypes.Node{
						{
							PeerID:            cciptypes.Bytes32{10 * id},
							OffchainPublicKey: cciptypes.Bytes32{20 * id},
						},
					},
					SourceChains: []rmntypes.SourceChain{
						{
							ChainSelector:       cciptypes.ChainSelector(uint64(id)),
							MinObservers:        uint64(id),
							ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(int64(id))),
						},
					},
					OffchainConfig: cciptypes.Bytes{30 * id},
				},
			},
		}
	}

	primary = createConfig(1, primaryEmpty)
	secondary = createConfig(2, secondaryEmpty)
	return primary, secondary
}
