package reader

import (
	"context"
	"fmt"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	chainA       = cciptypes.ChainSelector(1)
	chainB       = cciptypes.ChainSelector(2)
	chainC       = cciptypes.ChainSelector(3)
	oracleAId    = commontypes.OracleID(1)
	p2pOracleAId = libocrtypes.PeerID{byte(oracleAId)}
	oracleBId    = commontypes.OracleID(2)
	p2pOracleBId = libocrtypes.PeerID{byte(oracleBId)}
	oracleCId    = commontypes.OracleID(3)
	p2pOracleCId = libocrtypes.PeerID{byte(oracleCId)}
)

func TestHomeChainConfigPoller_HealthReport(t *testing.T) {
	homeChainReader := mocks.NewContractReaderMock()
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		consts.ContractNameCCIPConfig,
		consts.MethodNameGetAllChainConfigs,
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
				Config: []byte{0},
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
				Config: []byte{0},
			},
		},
		{
			ChainSelector: chainC,
			ChainConfig: HomeChainConfigMapper{
				FChain: 3,
				Readers: []libocrtypes.PeerID{
					p2pOracleCId,
				},
				Config: []byte{0},
			},
		},
	}
	homeChainConfig := map[cciptypes.ChainSelector]ChainConfig{
		chainA: {
			FChain:         1,
			SupportedNodes: mapset.NewSet(p2pOracleAId, p2pOracleBId, p2pOracleCId),
		},
		chainB: {
			FChain:         2,
			SupportedNodes: mapset.NewSet(p2pOracleAId, p2pOracleBId),
		},
		chainC: {
			FChain:         3,
			SupportedNodes: mapset.NewSet(p2pOracleCId),
		},
	}

	homeChainReader := mocks.NewContractReaderMock()
	var firstCall = true
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		consts.ContractNameCCIPConfig,
		consts.MethodNameGetAllChainConfigs,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(
		func(args mock.Arguments) {
			arg := args.Get(5).(*[]ChainConfigInfo)
			if firstCall {
				*arg = onChainConfigs
				firstCall = false
			} else {
				*arg = []ChainConfigInfo{} // return empty for other pages
			}
		}).Return(nil)

	defer homeChainReader.AssertExpectations(t)

	var (
		tickTime       = 2 * time.Millisecond
		totalSleepTime = tickTime * 0
	)

	configPoller := NewHomeChainConfigPoller(
		homeChainReader,
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
	// called at least 2 times, one for start and one for the first tick
	require.GreaterOrEqual(t, callCount, 2)

	configs, err := configPoller.GetAllChainConfigs()
	require.NoError(t, err)
	require.Equal(t, homeChainConfig, configs)
}

func Test_HomeChainPoller_GetOCRConfig(t *testing.T) {
	donID := uint32(1)
	pluginType := uint8(1) // execution
	homeChainReader := mocks.NewContractReaderMock()
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		consts.ContractNameCCIPConfig,
		consts.MethodNameGetOCRConfig,
		mock.Anything,
		map[string]any{
			"donId":      donID,
			"pluginType": pluginType,
		},
		mock.AnythingOfType("*[]reader.OCR3ConfigWithMeta"),
	).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(5).(*[]OCR3ConfigWithMeta)
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
	)

	configs, err := configPoller.GetOCRConfigs(context.Background(), donID, pluginType)
	require.NoError(t, err)
	require.Len(t, configs, 1)
	require.Equal(t, uint8(1), configs[0].Config.PluginType)
	require.Equal(t, cciptypes.ChainSelector(1), configs[0].Config.ChainSelector)
	require.Equal(t, uint8(1), configs[0].Config.F)
	require.Equal(t, []byte("offramp"), configs[0].Config.OfframpAddress)
}
