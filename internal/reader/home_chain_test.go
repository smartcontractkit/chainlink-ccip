package reader

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	chainreadermocks "github.com/smartcontractkit/chainlink-ccip/mocks/cl-common/chainreader"
)

var (
	chainA                  = cciptypes.ChainSelector(1)
	chainB                  = cciptypes.ChainSelector(2)
	chainC                  = cciptypes.ChainSelector(3)
	oracleAId               = commontypes.OracleID(1)
	p2pOracleAId            = libocrtypes.PeerID{byte(oracleAId)}
	oracleBId               = commontypes.OracleID(2)
	p2pOracleBId            = libocrtypes.PeerID{byte(oracleBId)}
	oracleCId               = commontypes.OracleID(3)
	p2pOracleCId            = libocrtypes.PeerID{byte(oracleCId)}
	ccipConfigBoundContract = types.BoundContract{
		Address: "0xCCIPConfigFakeAddress",
		Name:    consts.ContractNameCCIPConfig,
	}
	rmnHomeBoundContract = types.BoundContract{
		Address: "0xRMNHomeFakeAddress",
		Name:    consts.ContractNameRMNHome,
	}
)

func TestHomeChainConfigPoller_HealthReport(t *testing.T) {
	homeChainReader := chainreadermocks.NewMockContractReader(t)
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
	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		tickTime,
		ccipConfigBoundContract,
		rmnHomeBoundContract,
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

func Test_PollingWorking(t *testing.T) {
	chainConfig := chainconfig.ChainConfig{}
	encodedChainConfig, err := chainconfig.EncodeChainConfig(chainConfig)
	require.NoError(t, err)
	ccipOnChainConfigs := []ChainConfigInfo{
		{
			ChainSelector: chainA,
			ChainConfig: HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					p2pOracleAId,
					p2pOracleBId,
					p2pOracleCId,
				},
				Config: encodedChainConfig,
			},
		},
		{
			ChainSelector: chainB,
			ChainConfig: HomeChainConfigMapper{
				FChain: 2,
				Readers: []libocrtypes.PeerID{
					p2pOracleAId,
					p2pOracleBId,
				},
				Config: encodedChainConfig,
			},
		},
		{
			ChainSelector: chainC,
			ChainConfig: HomeChainConfigMapper{
				FChain: 3,
				Readers: []libocrtypes.PeerID{
					p2pOracleCId,
				},
				Config: encodedChainConfig,
			},
		},
	}
	homeChainConfig := map[cciptypes.ChainSelector]ChainConfig{
		chainA: {
			FChain:         1,
			SupportedNodes: mapset.NewSet(p2pOracleAId, p2pOracleBId, p2pOracleCId),
			Config:         chainConfig,
		},
		chainB: {
			FChain:         2,
			SupportedNodes: mapset.NewSet(p2pOracleAId, p2pOracleBId),
			Config:         chainConfig,
		},
		chainC: {
			FChain:         3,
			SupportedNodes: mapset.NewSet(p2pOracleCId),
			Config:         chainConfig,
		},
	}

	rmnHomeOnChainConfigs := createTestRMNHomeConfigs()

	homeChainReader := chainreadermocks.NewMockContractReader(t)
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		ccipConfigBoundContract.ReadIdentifier(consts.MethodNameGetAllChainConfigs),
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(
		func(args mock.Arguments) {
			arg := args.Get(4).(*[]ChainConfigInfo)
			*arg = ccipOnChainConfigs
		}).Return(nil)

	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		rmnHomeBoundContract.ReadIdentifier(consts.MethodNameGetVersionedConfigsWithDigests),
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

	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		tickTime,
		ccipConfigBoundContract,
		rmnHomeBoundContract,
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
	// called at least 4 times, one for start and one for the first tick for each contract (ccip *2 and rmn *2)
	require.GreaterOrEqual(t, callCount, 4)
	configs, err := configPoller.GetAllChainConfigs()
	require.NoError(t, err)
	require.Equal(t, homeChainConfig, configs)

	for _, configs := range rmnHomeOnChainConfigs {
		rmnNodes, err := configPoller.GetRMNNodesInfo(configs.ConfigDigest)
		require.NoError(t, err)
		require.NotNil(t, rmnNodes)

		isValid, err := configPoller.IsRMNHomeConfigDigestSet(configs.ConfigDigest)
		require.NoError(t, err)
		require.True(t, isValid)

		offchainConfig, err := configPoller.GetOffChainConfig(configs.ConfigDigest)
		require.NoError(t, err)
		require.NotNil(t, offchainConfig)

		minObs, err := configPoller.GetMinObservers(configs.ConfigDigest)
		require.NoError(t, err)
		require.NotNil(t, minObs)
	}
}

func Test_HomeChainPoller_GetOCRConfig(t *testing.T) {
	donID := uint32(1)
	pluginType := uint8(1) // execution
	homeChainReader := chainreadermocks.NewMockContractReader(t)
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		map[string]any{
			"donId":      donID,
			"pluginType": pluginType,
		},
		mock.AnythingOfType("*[]reader.OCR3ConfigWithMeta"),
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(4).(*[]OCR3ConfigWithMeta)
		*arg = append(*arg, OCR3ConfigWithMeta{
			ConfigCount: 1,
			Config: OCR3Config{
				PluginType:     pluginType,
				ChainSelector:  1,
				F:              1,
				OfframpAddress: []byte("offramp"),
			},
		})
	})
	defer homeChainReader.AssertExpectations(t)

	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		10*time.Millisecond,
		ccipConfigBoundContract,
		rmnHomeBoundContract,
	)

	configs, err := configPoller.GetOCRConfigs(context.Background(), donID, pluginType)
	require.NoError(t, err)
	require.Len(t, configs, 1)
	require.Equal(t, uint8(1), configs[0].Config.PluginType)
	require.Equal(t, cciptypes.ChainSelector(1), configs[0].Config.ChainSelector)
	require.Equal(t, uint8(1), configs[0].Config.F)
	require.Equal(t, []byte("offramp"), configs[0].Config.OfframpAddress)
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

func createTestRMNHomeConfigs() []rmntypes.VersionedConfigWithDigest {
	return []rmntypes.VersionedConfigWithDigest{
		{
			ConfigDigest: cciptypes.Bytes32{1, 2, 3, 4, 5},
			VersionedConfig: rmntypes.VersionedConfig{
				Version: 1,
				Config: rmntypes.Config{
					Nodes: []rmntypes.Node{
						{
							PeerID:            cciptypes.Bytes32{10, 11, 12, 13, 14},
							OffchainPublicKey: cciptypes.Bytes32{20, 21, 22, 23, 24},
						},
						{
							PeerID:            cciptypes.Bytes32{15, 16, 17, 18, 19},
							OffchainPublicKey: cciptypes.Bytes32{25, 26, 27, 28, 29},
						},
					},
					SourceChains: []rmntypes.SourceChain{
						{
							ChainSelector:       cciptypes.ChainSelector(1),
							MinObservers:        1,
							ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(2)), // 10 in binary = 1 is observer and 0 is not
						},
						{
							ChainSelector:       cciptypes.ChainSelector(2),
							MinObservers:        2,
							ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(3)), // 11 in binary = 0 and 1 are observers
						},
					},
					OffchainConfig: cciptypes.Bytes{30, 31, 32, 33, 34},
				},
			},
		},
		{
			ConfigDigest: cciptypes.Bytes32{6, 7, 8, 9, 10},
			VersionedConfig: rmntypes.VersionedConfig{
				Version: 2,
				Config: rmntypes.Config{
					Nodes: []rmntypes.Node{
						{
							PeerID:            cciptypes.Bytes32{40, 41, 42, 43, 44},
							OffchainPublicKey: cciptypes.Bytes32{50, 51, 52, 53, 54},
						},
					},
					SourceChains: []rmntypes.SourceChain{
						{
							ChainSelector:       cciptypes.ChainSelector(1),
							MinObservers:        1,
							ObserverNodesBitmap: cciptypes.NewBigInt(big.NewInt(1)), // 1 in binary = 0 is an observer
						},
					},
					OffchainConfig: cciptypes.Bytes{60, 61, 62, 63, 64},
				},
			},
		},
	}
}
