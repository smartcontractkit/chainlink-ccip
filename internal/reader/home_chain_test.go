package reader

import (
	"context"
	"fmt"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	ccipConfigBoundContract = types.BoundContract{
		Address: "0xCCIPConfigFakeAddress",
		Name:    consts.ContractNameCCIPConfig,
	}
)

func TestHomeChainConfigPoller_HealthReport(t *testing.T) {
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
	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		tickTime,
		ccipConfigBoundContract,
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
	onChainConfigs := []ChainConfigInfo{
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

	homeChainReader := readermock.NewMockContractReaderFacade(t)
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(
		func(args mock.Arguments) {
			arg := args.Get(4).(*[]ChainConfigInfo)
			*arg = onChainConfigs
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
	// called at least 2 times, one for start and one for the first tick
	require.GreaterOrEqual(t, callCount, 2)

	configs, err := configPoller.GetAllChainConfigs()
	require.NoError(t, err)
	require.Equal(t, homeChainConfig, configs)
}

func Test_HomeChainPoller_GetOCRConfig(t *testing.T) {
	donID := uint32(1)
	pluginType := uint8(1) // execution
	homeChainReader := readermock.NewMockContractReaderFacade(t)
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		map[string]any{
			"donId":      donID,
			"pluginType": pluginType,
		},
		mock.AnythingOfType("*reader.ActiveAndCandidate"),
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(4).(*ActiveAndCandidate)
		*arg = ActiveAndCandidate{
			ActiveConfig: OCR3ConfigWithMeta{
				Version: 1,
				Config: OCR3Config{
					PluginType:     pluginType,
					ChainSelector:  1,
					FRoleDON:       1,
					OfframpAddress: []byte("offramp"),
				},
				ConfigDigest: [32]byte{1},
			},
		}
	})
	defer homeChainReader.AssertExpectations(t)

	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
		logger.Test(t),
		10*time.Millisecond,
		ccipConfigBoundContract,
	)

	configs, err := configPoller.GetOCRConfigs(context.Background(), donID, pluginType)
	require.NoError(t, err)
	require.NotEqual(t, [32]byte{}, configs.ActiveConfig.ConfigDigest)
	require.Equal(t, [32]byte{}, configs.CandidateConfig.ConfigDigest)
	require.Equal(t, uint8(1), configs.ActiveConfig.Config.PluginType)
	require.Equal(t, cciptypes.ChainSelector(1), configs.ActiveConfig.Config.ChainSelector)
	require.Equal(t, uint8(1), configs.ActiveConfig.Config.FRoleDON)
	require.Equal(t, []byte("offramp"), configs.ActiveConfig.Config.OfframpAddress)
}
