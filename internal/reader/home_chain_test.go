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
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip/consts"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
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

// makeValidOnChainConfig returns a ChainConfigInfo that passes validateChainConfigInfos.
func makeValidOnChainConfig(
	t *testing.T,
	sel cciptypes.ChainSelector,
	fChain uint8,
	peers ...libocrtypes.PeerID,
) ChainConfigInfo {
	t.Helper()

	chainConfig := chainconfig.ChainConfig{}
	encoded, err := chainconfig.EncodeChainConfig(chainConfig)
	require.NoError(t, err)
	return ChainConfigInfo{
		ChainSelector: sel,
		ChainConfig: HomeChainConfigMapper{
			FChain:  fChain,
			Readers: peers,
			Config:  encoded,
		},
	}
}

// makeInvalidOnChainConfig returns a ChainConfigInfo that fails validateChainConfigInfos
// (FChain == 0 is the simplest trigger).
func makeInvalidOnChainConfig(sel cciptypes.ChainSelector) ChainConfigInfo {
	return ChainConfigInfo{
		ChainSelector: sel,
		ChainConfig:   HomeChainConfigMapper{FChain: 0},
	}
}

// paginatedMock sets up GetLatestValue to return successive pages from pages.
// Each call pops the next page from the slice.
func paginatedMock(t *testing.T, pages [][]ChainConfigInfo) *readermock.MockContractReaderFacade {
	t.Helper()
	m := readermock.NewMockContractReaderFacade(t)
	call := 0
	m.On("GetLatestValue", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			out := args.Get(4).(*[]ChainConfigInfo)
			if call < len(pages) {
				*out = pages[call]
			} else {
				*out = nil
			}
			call++
		}).
		Return(nil)
	return m
}

// Test_FetchAndSetConfigs_ZeroConfigs verifies that fetchAndSetConfigs returns
// promptly when no configs exist on-chain (empty first page) instead of
// hanging forever.
func Test_FetchAndSetConfigs_ZeroConfigs(t *testing.T) {
	m := paginatedMock(t, [][]ChainConfigInfo{
		{}, // page 0: empty
	})

	poller := NewHomeChainConfigPoller(m, logger.Test(t), time.Hour, ccipConfigBoundContract)
	hcp := poller.(*homeChainPoller)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	require.NoError(t, hcp.fetchAndSetConfigs(ctx))

	configs, err := poller.GetAllChainConfigs()
	require.NoError(t, err)
	require.Empty(t, configs)
}

// Test_FetchAndSetConfigs_ExactPageMultiple verifies that the loop terminates
// when the total config count is an exact multiple of defaultConfigPageSize.
// The mock returns one full page of valid configs followed by an empty page.
func Test_FetchAndSetConfigs_ExactPageMultiple(t *testing.T) {
	pageSize := int(defaultConfigPageSize)
	fullPage := make([]ChainConfigInfo, pageSize)
	for i := 0; i < pageSize; i++ {
		fullPage[i] = makeValidOnChainConfig(t, cciptypes.ChainSelector(i+1), 1, p2pOracleAId)
	}

	m := paginatedMock(t, [][]ChainConfigInfo{
		fullPage, // page 0: exactly 100 valid entries
		{},       // page 1: empty — terminates pagination
	})

	poller := NewHomeChainConfigPoller(m, logger.Test(t), time.Hour, ccipConfigBoundContract)
	hcp := poller.(*homeChainPoller)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	require.NoError(t, hcp.fetchAndSetConfigs(ctx))

	configs, err := poller.GetAllChainConfigs()
	require.NoError(t, err)
	require.Len(t, configs, pageSize)
}

// Test_FetchAndSetConfigs_MixedValidInvalid verifies that only valid entries
// reach setState when a page contains a mix of valid and invalid configs.
func Test_FetchAndSetConfigs_MixedValidInvalid(t *testing.T) {
	m := paginatedMock(t, [][]ChainConfigInfo{
		{
			makeValidOnChainConfig(t, chainA, 1, p2pOracleAId),
			makeInvalidOnChainConfig(chainB), // FChain == 0: invalid
			makeValidOnChainConfig(t, chainC, 2, p2pOracleCId),
		},
	})

	poller := NewHomeChainConfigPoller(m, logger.Test(t), time.Hour, ccipConfigBoundContract)
	hcp := poller.(*homeChainPoller)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	require.NoError(t, hcp.fetchAndSetConfigs(ctx))

	configs, err := poller.GetAllChainConfigs()
	require.NoError(t, err)
	require.Len(t, configs, 2)
	require.Contains(t, configs, chainA)
	require.Contains(t, configs, chainC)
	require.NotContains(t, configs, chainB)
}

// Test_FetchAndSetConfigs_ContextCancelled verifies that fetchAndSetConfigs
// returns ctx.Err() when the context is cancelled mid-pagination.
// The mock returns a full page on the first call and cancels the context so
// that the ctx.Done() guard at the top of the second iteration fires.
func Test_FetchAndSetConfigs_ContextCancelled(t *testing.T) {
	pageSize := int(defaultConfigPageSize)
	fullPage := make([]ChainConfigInfo, pageSize)
	for i := 0; i < pageSize; i++ {
		fullPage[i] = makeValidOnChainConfig(t, cciptypes.ChainSelector(i+1), 1, p2pOracleAId)
	}

	ctx, cancel := context.WithCancel(context.Background())

	m := readermock.NewMockContractReaderFacade(t)
	m.On("GetLatestValue", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			out := args.Get(4).(*[]ChainConfigInfo)
			*out = fullPage
			// Cancel after the first page so the next iteration's ctx.Done() fires.
			cancel()
		}).
		Return(nil).
		Once()

	poller := NewHomeChainConfigPoller(m, logger.Test(t), time.Hour, ccipConfigBoundContract)
	hcp := poller.(*homeChainPoller)

	err := hcp.fetchAndSetConfigs(ctx)
	require.ErrorIs(t, err, context.Canceled)
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
