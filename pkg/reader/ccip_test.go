package reader

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	commonccipocr3 "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/ccipocr3"
	writermocks "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/types"
	readermocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/addressbook"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

var (
	chainA = cciptypes.ChainSelector(1)
	chainB = cciptypes.ChainSelector(2)
	chainC = cciptypes.ChainSelector(3)
	chainD = cciptypes.ChainSelector(4)
)

func TestCCIPChainReader_GetContractAddress(t *testing.T) {
	ab := addressbook.NewBook()
	require.NoError(t, ab.InsertOrUpdate(addressbook.ContractAddresses{
		"abc": map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			123: []byte("0x123"),
		},
	}))

	r := &ccipChainReader{}
	r.donAddressBook = ab

	addr, err := r.GetContractAddress("abc", 123)
	require.NoError(t, err)
	require.Equal(t, []byte("0x123"), addr)
}

func TestCCIPChainReader_Sync_HappyPath_BindsContractsSuccessfully(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destNonceMgr := []byte{0x3}
	offRamp := []byte{0x4}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	destExtended := readermocks.NewMockExtended(t)
	source1Extended := readermocks.NewMockExtended(t)
	source2Extended := readermocks.NewMockExtended(t)

	cw := writermocks.NewMockContractWriter(t)
	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	contractWriters[destChain] = cw
	contractWriters[sourceChain1] = cw
	contractWriters[sourceChain2] = cw

	chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1, sourceChain2)
	// OnRamp sourceChain1
	mockExpectChainAccessorSyncCall(chainAccessors[sourceChain1], consts.ContractNameOnRamp, s1Onramp, nil)
	// OnRamp sourceChain2
	mockExpectChainAccessorSyncCall(chainAccessors[sourceChain2], consts.ContractNameOnRamp, s2Onramp, nil)
	// OffRamp dest chain
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRamp, nil)
	// NonceManager dest chain
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameNonceManager, destNonceMgr, nil)
	ccipReader, err := newCCIPChainReaderInternal(
		ctx,
		logger.Test(t),
		chainAccessors,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			destChain:    destExtended,
			sourceChain1: source1Extended,
			sourceChain2: source2Extended,
		},
		contractWriters,
		destChain,
		offRamp,
		mockAddrCodec,
	)
	require.NoError(t, err)

	contracts := ContractAddresses{
		consts.ContractNameOnRamp: {
			sourceChain1: s1Onramp,
			sourceChain2: s2Onramp,
		},
		consts.ContractNameNonceManager: {
			destChain: destNonceMgr,
		},
	}

	err = ccipReader.Sync(ctx, contracts)
	require.NoError(t, err)
	err = ccipReader.Close()
	require.NoError(t, err)
}

func TestCCIPChainReader_Sync_HappyPath_SkipsEmptyAddress(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	offRamp := []byte{0x4}

	s2Onramp := []byte{}

	destNonceMgr := []byte{0x3}
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	destExtended := readermocks.NewMockExtended(t)
	source1Extended := readermocks.NewMockExtended(t)
	source2Extended := readermocks.NewMockExtended(t)

	cw := writermocks.NewMockContractWriter(t)
	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	contractWriters[destChain] = cw
	contractWriters[sourceChain1] = cw
	contractWriters[sourceChain2] = cw

	chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1)
	// Sync() on OnRamp sourceChain1
	mockExpectChainAccessorSyncCall(chainAccessors[sourceChain1], consts.ContractNameOnRamp, s1Onramp, nil)
	// Sync() OffRamp (from constructor) and NonceManager (from this test)
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRamp, nil)
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameNonceManager, destNonceMgr, nil)
	ccipReader, err := newCCIPChainReaderInternal(
		ctx,
		logger.Test(t),
		chainAccessors,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			destChain:    destExtended,
			sourceChain1: source1Extended,
			sourceChain2: source2Extended,
		},
		contractWriters,
		destChain,
		offRamp,
		mockAddrCodec,
	)
	require.NoError(t, err)

	contracts := ContractAddresses{
		consts.ContractNameOnRamp: {
			sourceChain1: s1Onramp,
			sourceChain2: s2Onramp,
		},
		consts.ContractNameNonceManager: {
			destChain: destNonceMgr,
		},
	}

	err = ccipReader.Sync(ctx, contracts)
	require.NoError(t, err)
	err = ccipReader.Close()
	require.NoError(t, err)
}

func TestCCIPChainReader_Sync_HappyPath_DontSupportAllChains(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destNonceMgr := []byte{0x3}
	offRamp := []byte{0x4}
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	destExtended := readermocks.NewMockExtended(t)
	// only support source2, source1 unsupported.
	source2Extended := readermocks.NewMockExtended(t)

	cw := writermocks.NewMockContractWriter(t)
	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	contractWriters[destChain] = cw
	contractWriters[sourceChain2] = cw

	chainAccessors := createMockedChainAccessors(t, destChain, sourceChain2)
	// OnRamp sourceChain2
	mockExpectChainAccessorSyncCall(chainAccessors[sourceChain2], consts.ContractNameOnRamp, s2Onramp, nil)
	// OffRamp (during init) and NonceManager (from test)
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRamp, nil)
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameNonceManager, destNonceMgr, nil)
	ccipReader, err := newCCIPChainReaderInternal(
		ctx,
		logger.Test(t),
		chainAccessors,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			destChain:    destExtended,
			sourceChain2: source2Extended,
		},
		contractWriters,
		destChain,
		offRamp,
		mockAddrCodec,
	)
	require.NoError(t, err)

	contracts := ContractAddresses{
		consts.ContractNameOnRamp: {
			sourceChain1: s1Onramp,
			sourceChain2: s2Onramp,
		},
		consts.ContractNameNonceManager: {
			destChain: destNonceMgr,
		},
	}

	err = ccipReader.Sync(ctx, contracts)
	require.NoError(t, err)
	err = ccipReader.Close()
	require.NoError(t, err)
}

func TestCCIPChainReader_Sync_BindError(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destNonceMgr := []byte{0x3}
	offRamp := []byte{0x4}

	expectedErr := errors.New("some error")
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	destExtended := readermocks.NewMockExtended(t)
	source1Extended := readermocks.NewMockExtended(t)
	source2Extended := readermocks.NewMockExtended(t)

	cw := writermocks.NewMockContractWriter(t)
	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	contractWriters[destChain] = cw
	contractWriters[sourceChain1] = cw
	contractWriters[sourceChain2] = cw

	chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1, sourceChain2)
	// sourceChain1 accessor will fail with an error
	mockExpectChainAccessorSyncCall(chainAccessors[sourceChain1], consts.ContractNameOnRamp, s1Onramp, expectedErr)
	mockExpectChainAccessorSyncCall(chainAccessors[sourceChain2], consts.ContractNameOnRamp, s2Onramp, nil)
	// OffRamp (during init) + NonceManager (from test)
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRamp, nil)
	mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameNonceManager, destNonceMgr, nil)
	ccipReader, err := newCCIPChainReaderInternal(
		ctx,
		logger.Test(t),
		chainAccessors,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			destChain:    destExtended,
			sourceChain1: source1Extended,
			sourceChain2: source2Extended,
		},
		contractWriters,
		destChain,
		offRamp,
		mockAddrCodec,
	)
	require.NoError(t, err)

	contracts := ContractAddresses{
		consts.ContractNameOnRamp: {
			sourceChain1: s1Onramp,
			sourceChain2: s2Onramp,
		},
		consts.ContractNameNonceManager: {
			destChain: destNonceMgr,
		},
	}

	err = ccipReader.Sync(ctx, contracts)
	require.Error(t, err)
	require.ErrorIs(t, err, expectedErr)
	err = ccipReader.Close()
	require.NoError(t, err)
}

// The round1 version returns NoBindingFound errors for onramp contracts to simulate
// the two-phase approach to discovering those contracts.
func TestCCIPChainReader_DiscoverContracts_HappyPath_Round1(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain := [2]cciptypes.ChainSelector{2, 3}
	onramps := [2][]byte{{0x1}, {0x2}}
	destNonceMgr := []byte{0x3}
	destRMNRemote := []byte{0x4}
	destFeeQuoter := []byte{0x5}
	destRouter := []byte{0x6}

	sourceChainConfigs := make(map[cciptypes.ChainSelector]StaticSourceChainConfig, len(sourceChain))
	for i, chain := range sourceChain {
		sourceChainConfigs[chain] = StaticSourceChainConfig{
			Router:    destRouter,
			IsEnabled: true,
			OnRamp:    onramps[i],
		}
	}

	// Build expected addresses.
	var expectedContractAddresses ContractAddresses
	for i := range onramps {
		expectedContractAddresses = expectedContractAddresses.Append(
			consts.ContractNameOnRamp, sourceChain[i], onramps[i])
	}
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameRouter, destChain, destRouter)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameFeeQuoter, destChain, destFeeQuoter)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameRMNRemote, destChain, destRMNRemote)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameNonceManager, destChain, destNonceMgr)

	mockReaders := make(map[cciptypes.ChainSelector]*readermocks.MockExtended)
	mockReaders[destChain] = readermocks.NewMockExtended(t)

	// Setup cache mock and configurations
	mockCache := new(mockConfigCache)
	// Destination chain config
	destChainConfig := cciptypes.ChainConfigSnapshot{
		Offramp: cciptypes.OfframpConfig{
			StaticConfig: cciptypes.OffRampStaticChainConfig{
				NonceManager: destNonceMgr,
				RmnRemote:    destRMNRemote,
			},
			DynamicConfig: cciptypes.OffRampDynamicChainConfig{
				FeeQuoter: destFeeQuoter,
			},
		},
	}
	// Set up cache expectations for destination chain and source chains
	mockCache.On("GetChainConfig", mock.Anything, destChain).Return(destChainConfig, nil).Once()
	mockCache.On("GetChainConfig", mock.Anything, sourceChain[0]).Return(
		cciptypes.ChainConfigSnapshot{}, contractreader.ErrNoBindings).Maybe()
	mockCache.On("GetChainConfig", mock.Anything, sourceChain[1]).Return(
		cciptypes.ChainConfigSnapshot{}, contractreader.ErrNoBindings).Maybe()
	mockCache.On(
		"GetOfframpSourceChainConfigs",
		mock.Anything,
		destChain,
		sourceChain[:]).Return(sourceChainConfigs, nil).Once()

	castToExtended := make(map[cciptypes.ChainSelector]contractreader.Extended)
	for sel, v := range mockReaders {
		castToExtended[sel] = v
	}

	lggr, hook := logger.TestObserved(t, zapcore.InfoLevel)

	// create the reader with cache
	ccipChainReader := &ccipChainReader{
		destChain:       destChain,
		contractReaders: castToExtended,
		donAddressBook:  addressbook.NewBook(),
		lggr:            lggr,
		configPoller:    mockCache,
	}

	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx, []cciptypes.ChainSelector{destChain}, sourceChain[:])
	require.NoError(t, err)

	assert.Equal(t, expectedContractAddresses, contractAddresses)
	require.Equal(t, 3, hook.Len())

	assert.Contains(
		t,
		hook.All()[0].Message,
		"appending router contract address",
	)
	assert.Contains(
		t,
		hook.All()[1].Message,
		"appending RMN remote contract address",
	)
	assert.Contains(
		t,
		hook.All()[2].Message,
		"appending fee quoter contract address",
	)

	mockCache.AssertExpectations(t)
}

// The round2 version includes calls to the onRamp contracts.
func TestCCIPChainReader_DiscoverContracts_HappyPath_Round2(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain := [2]cciptypes.ChainSelector{2, 3}
	onramps := [2][]byte{{0x1}, {0x2}}
	destNonceMgr := []byte{0x3}
	destRMNRemote := []byte{0x4}
	destFeeQuoter := []byte{0x5}
	destRouter := [2][]byte{{0x6}, {0xFF}} // We should never see 0xFF in the result.
	srcFeeQuoters := [2][]byte{{0x7}, {0x8}}
	srcRouters := [2][]byte{{0x9}, {0x10}}

	sourceChainConfigs := make(map[cciptypes.ChainSelector]StaticSourceChainConfig, len(sourceChain))
	for i, chain := range sourceChain {
		sourceChainConfigs[chain] = StaticSourceChainConfig{
			Router:    destRouter[i], // Using the corresponding router from destRouter array
			IsEnabled: true,
			OnRamp:    onramps[i],
		}
	}

	// Build expected addresses.
	var expectedContractAddresses ContractAddresses
	for i := range onramps {
		expectedContractAddresses = expectedContractAddresses.Append(
			consts.ContractNameOnRamp, sourceChain[i], onramps[i])
	}
	for i := range srcFeeQuoters {
		expectedContractAddresses = expectedContractAddresses.Append(
			consts.ContractNameFeeQuoter, sourceChain[i], srcFeeQuoters[i])
	}
	for i := range srcRouters {
		expectedContractAddresses = expectedContractAddresses.Append(
			consts.ContractNameRouter, sourceChain[i], srcRouters[i])
	}
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameFeeQuoter, destChain, destFeeQuoter)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameRMNRemote, destChain, destRMNRemote)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameNonceManager, destChain, destNonceMgr)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameRouter, destChain, destRouter[0])

	mockReaders := make(map[cciptypes.ChainSelector]*readermocks.MockExtended)
	mockReaders[destChain] = readermocks.NewMockExtended(t)

	// Setup cache mock and configurations
	mockCache := new(mockConfigCache)
	// Destination chain config
	destChainConfig := cciptypes.ChainConfigSnapshot{
		Offramp: cciptypes.OfframpConfig{
			StaticConfig: cciptypes.OffRampStaticChainConfig{
				NonceManager: destNonceMgr,
				RmnRemote:    destRMNRemote,
			},
			DynamicConfig: cciptypes.OffRampDynamicChainConfig{
				FeeQuoter: destFeeQuoter,
			},
		},
	}
	// Set up destination chain expectation
	mockCache.On("GetChainConfig", mock.Anything, destChain).Return(destChainConfig, nil).Once()
	mockCache.On(
		"GetOfframpSourceChainConfigs",
		mock.Anything,
		destChain,
		sourceChain[:]).Return(sourceChainConfigs, nil).Once()

	// Set up source chain expectations with proper OnRamp configs
	for i, chain := range sourceChain {
		// Create mock reader for source chain
		mockReaders[chain] = readermocks.NewMockExtended(t)

		srcChainConfig := cciptypes.ChainConfigSnapshot{
			OnRamp: cciptypes.OnRampConfig{
				DynamicConfig: cciptypes.GetOnRampDynamicConfigResponse{
					DynamicConfig: cciptypes.OnRampDynamicConfig{
						FeeQuoter: srcFeeQuoters[i],
					},
				},
				DestChainConfig: cciptypes.OnRampDestChainConfig{
					Router: srcRouters[i],
				},
			},
		}
		mockCache.On("GetChainConfig", mock.Anything, chain).Return(srcChainConfig, nil).Once()
	}

	castToExtended := make(map[cciptypes.ChainSelector]contractreader.Extended)
	for sel, v := range mockReaders {
		castToExtended[sel] = v
	}

	// create the reader with cache
	ccipChainReader := &ccipChainReader{
		destChain:       destChain,
		contractReaders: castToExtended,
		lggr:            logger.Test(t),
		configPoller:    mockCache,
	}

	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx, []cciptypes.ChainSelector{
		sourceChain[0], sourceChain[1], destChain}, sourceChain[:])
	require.NoError(t, err)
	require.Equal(t, expectedContractAddresses, contractAddresses)
	mockCache.AssertExpectations(t)
}

func TestCCIPChainReader_DiscoverContracts_GetAllSourceChainConfig_Errors(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)

	// Setup cache mock and configuration
	mockCache := new(mockConfigCache)
	chainConfig := cciptypes.ChainConfigSnapshot{
		Offramp: cciptypes.OfframpConfig{
			// We can leave the configs empty since we just need GetChainConfig to succeed
			StaticConfig:  cciptypes.OffRampStaticChainConfig{},
			DynamicConfig: cciptypes.OffRampDynamicChainConfig{},
		},
	}
	mockCache.On("GetChainConfig", mock.Anything, destChain).Return(chainConfig, nil)

	// Setup mock cache to return an error
	destExtended := readermocks.NewMockExtended(t)
	getLatestValueErr := errors.New("some error")
	mockCache.On("GetOfframpSourceChainConfigs", mock.Anything, destChain,
		[]cciptypes.ChainSelector{sourceChain1, sourceChain2}).
		Return(map[cciptypes.ChainSelector]StaticSourceChainConfig{}, getLatestValueErr).Once()

	// create the reader with cache
	ccipChainReader := &ccipChainReader{
		destChain: destChain,
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain: destExtended,
			// these won't be used in this test, but are needed because
			// we determine the source chain selectors to query from the chains
			// that we have readers for.
			sourceChain1: readermocks.NewMockExtended(t),
			sourceChain2: readermocks.NewMockExtended(t),
		},
		lggr:         logger.Test(t),
		configPoller: mockCache,
	}

	_, err := ccipChainReader.DiscoverContracts(
		ctx,
		[]cciptypes.ChainSelector{sourceChain1, sourceChain2, destChain},
		[]cciptypes.ChainSelector{sourceChain1, sourceChain2},
	)
	require.Error(t, err)
	require.ErrorIs(t, err, getLatestValueErr)
	mockCache.AssertExpectations(t)
}

func TestCCIPChainReader_DiscoverContracts_GetOfframpStaticConfig_Errors(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)

	// Setup mock cache to return error
	// mock the call to get the static config - failure
	getLatestValueErr := errors.New("some error")
	mockCache := new(mockConfigCache)
	mockCache.On("GetChainConfig", mock.Anything, destChain).Return(cciptypes.ChainConfigSnapshot{}, getLatestValueErr)

	// create the reader with cache
	ccipChainReader := &ccipChainReader{
		destChain: destChain,
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain: readermocks.NewMockExtended(t),
			// these won't be used in this test, but are needed because
			// we determine the source chain selectors to query from the chains
			// that we have readers for.
			sourceChain1: readermocks.NewMockExtended(t),
			sourceChain2: readermocks.NewMockExtended(t),
		},
		lggr:         logger.Test(t),
		configPoller: mockCache,
	}

	_, err := ccipChainReader.DiscoverContracts(ctx,
		[]cciptypes.ChainSelector{sourceChain1, sourceChain2, destChain},
		[]cciptypes.ChainSelector{sourceChain1, sourceChain2})
	require.Error(t, err)
	require.ErrorIs(t, err, getLatestValueErr)
	mockCache.AssertExpectations(t)
}

func TestCCIPChainReader_getDestFeeQuoterStaticConfig(t *testing.T) {
	ctx := context.Background()

	// Setup expected values
	offrampAddress := []byte{0x3}
	expectedConfig := cciptypes.FeeQuoterStaticConfig{
		MaxFeeJuelsPerMsg:  cciptypes.NewBigIntFromInt64(10),
		LinkToken:          []byte{0x3, 0x4},
		StalenessThreshold: 12,
	}

	// Setup cache with the expected config
	mockCache := new(mockConfigCache)
	chainConfig := cciptypes.ChainConfigSnapshot{
		FeeQuoter: cciptypes.FeeQuoterConfig{
			StaticConfig: expectedConfig,
		},
	}
	mockCache.On("GetChainConfig", mock.Anything, chainC).Return(chainConfig, nil)

	mockAddrCodec := internal.NewMockAddressCodecHex(t)

	offrampAddressStr, err := mockAddrCodec.AddressBytesToString(offrampAddress, chainC)
	require.NoError(t, err)
	ccipReader := &ccipChainReader{
		lggr:           logger.Test(t),
		destChain:      chainC,
		configPoller:   mockCache,
		offrampAddress: offrampAddressStr,
	}

	cfg, err := ccipReader.getDestFeeQuoterStaticConfig(ctx)
	require.NoError(t, err)

	assert.Equal(t, expectedConfig.MaxFeeJuelsPerMsg, cfg.MaxFeeJuelsPerMsg)
	assert.Equal(t, expectedConfig.LinkToken, cfg.LinkToken)
	assert.Equal(t, expectedConfig.StalenessThreshold, cfg.StalenessThreshold)

	mockCache.AssertExpectations(t)
}

func TestCCIPChainReader_getFeeQuoterTokenPriceUSD(t *testing.T) {
	tokenAddr := []byte{0x3, 0x4}
	offrampAddress := []byte{0x3}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)

	cw := writermocks.NewMockContractWriter(t)
	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	contractWriters[chainA] = cw
	contractWriters[chainB] = cw
	contractWriters[chainC] = cw

	chainAccessors := createMockedChainAccessors(t, chainA, chainB, chainC)
	expectedPrice := cciptypes.TimestampedUnixBig{
		Value:     big.NewInt(145),
		Timestamp: uint32(time.Now().Unix()),
	}
	mockExpectChainAccessorSyncCall(chainAccessors[chainC], consts.ContractNameOffRamp, offrampAddress, nil)
	mockExpectChainAccessorGetTokenPriceUSD(chainAccessors[chainC], tokenAddr, expectedPrice)

	ccipReader, err := newCCIPChainReaderInternal(
		t.Context(),
		logger.Test(t),
		chainAccessors,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			chainC: readermocks.NewMockContractReaderFacade(t),
		}, contractWriters, chainC, offrampAddress, mockAddrCodec,
	)
	require.NoError(t, err)

	// Add cleanup to properly shut down the background polling
	t.Cleanup(func() {
		err := ccipReader.Close()
		if err != nil {
			t.Logf("Error closing ccipReader: %v", err)
		}
	})

	ctx := context.Background()
	price, err := ccipReader.getFeeQuoterTokenPriceUSD(ctx, []byte{0x3, 0x4})
	assert.NoError(t, err)
	assert.Equal(t, cciptypes.NewBigIntFromInt64(145), price)
}

func TestCCIPFeeComponents_HappyPath(t *testing.T) {
	cw := writermocks.NewMockContractWriter(t)
	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	contractWriters[chainA] = cw
	contractWriters[chainB] = cw
	contractWriters[chainC] = cw

	sourceCRs := make(map[cciptypes.ChainSelector]*readermocks.MockContractReaderFacade)
	for _, chain := range []cciptypes.ChainSelector{chainA, chainB, chainC} {
		sourceCRs[chain] = readermocks.NewMockContractReaderFacade(t)
	}

	expectedFeeComponents := cciptypes.ChainFeeComponents{
		ExecutionFee:        big.NewInt(1),
		DataAvailabilityFee: big.NewInt(2),
	}

	offRampAddress := []byte{0x3}
	chainAccessors := createMockedChainAccessors(t, chainA, chainB, chainC)
	mockExpectChainAccessorSyncCall(chainAccessors[chainC], consts.ContractNameOffRamp, offRampAddress, nil)
	mockExpectChainAccessorGetChainFeeComponentsCall(chainAccessors[chainA], expectedFeeComponents, nil)
	mockExpectChainAccessorGetChainFeeComponentsCall(chainAccessors[chainB], expectedFeeComponents, nil)
	mockExpectChainAccessorGetChainFeeComponentsCall(chainAccessors[chainC], expectedFeeComponents, nil)
	ccipReader, err := newCCIPChainReaderInternal(
		t.Context(),
		logger.Test(t),
		chainAccessors,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			chainA: sourceCRs[chainA],
			chainB: sourceCRs[chainB],
			chainC: sourceCRs[chainC],
		},
		contractWriters,
		chainC,
		offRampAddress,
		internal.NewMockAddressCodecHex(t),
	)
	require.NoError(t, err)

	// Add cleanup to ensure resources are released
	t.Cleanup(func() {
		err := ccipReader.Close()
		if err != nil {
			t.Logf("Error closing ccipReader: %v", err)
		}
	})

	ctx := context.Background()
	feeComponents := ccipReader.GetChainsFeeComponents(ctx, []cciptypes.ChainSelector{chainA, chainB, chainC})
	assert.Len(t, feeComponents, 3)
	assert.Equal(t, big.NewInt(1), feeComponents[chainA].ExecutionFee)
	assert.Equal(t, big.NewInt(1), feeComponents[chainB].ExecutionFee)
	assert.Equal(t, big.NewInt(1), feeComponents[chainC].ExecutionFee)
	assert.Equal(t, big.NewInt(2), feeComponents[chainA].DataAvailabilityFee)
	assert.Equal(t, big.NewInt(2), feeComponents[chainB].DataAvailabilityFee)
	assert.Equal(t, big.NewInt(2), feeComponents[chainC].DataAvailabilityFee)

	destChainFeeComponent, err := ccipReader.GetDestChainFeeComponents(ctx)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1), destChainFeeComponent.ExecutionFee)
	assert.Equal(t, big.NewInt(2), destChainFeeComponent.DataAvailabilityFee)
}

func TestCCIPChainReader_LinkPriceUSD(t *testing.T) {
	ctx := context.Background()
	tokenAddr := []byte{0x3, 0x4}
	offrampAddress := []byte{0x3}

	// Setup mock cache with the fee quoter static config
	mockCache := new(mockConfigCache)
	chainConfig := cciptypes.ChainConfigSnapshot{
		FeeQuoter: cciptypes.FeeQuoterConfig{
			StaticConfig: cciptypes.FeeQuoterStaticConfig{
				MaxFeeJuelsPerMsg:  cciptypes.NewBigIntFromInt64(10),
				LinkToken:          tokenAddr,
				StalenessThreshold: 12,
			},
		},
	}
	mockCache.On("GetChainConfig", mock.Anything, chainC).Return(chainConfig, nil)

	// Setup contract reader for getting token price
	destCR := readermocks.NewMockExtended(t)

	// Mock accessors
	chainAccessors := createMockedChainAccessors(t, chainC)
	expectedPrice := cciptypes.TimestampedUnixBig{
		Value:     big.NewInt(145),
		Timestamp: uint32(time.Now().Unix()),
	}
	mockExpectChainAccessorGetTokenPriceUSD(chainAccessors[chainC], tokenAddr, expectedPrice)

	// Setup ccipReader with both cache and contract readers
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	offrampAddressStr, err := mockAddrCodec.AddressBytesToString(offrampAddress, chainC)
	require.NoError(t, err)
	ccipReader := &ccipChainReader{
		lggr:           logger.Test(t),
		destChain:      chainC,
		configPoller:   mockCache,
		offrampAddress: offrampAddressStr,
		accessors:      chainAccessors,
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			chainC: destCR,
		},
	}

	price, err := ccipReader.LinkPriceUSD(ctx)
	require.NoError(t, err)
	assert.Equal(t, cciptypes.NewBigIntFromInt64(145), price)

	mockCache.AssertExpectations(t)
}

func Test_getCurseInfoFromCursedSubjects(t *testing.T) {
	testCases := []struct {
		name              string
		cursedSubjectsSet mapset.Set[[16]byte]
		destChainSelector cciptypes.ChainSelector
		expCurseInfo      cciptypes.CurseInfo
	}{
		{
			name:              "no cursed subjects",
			cursedSubjectsSet: mapset.NewSet[[16]byte](),
			destChainSelector: chainA,
			expCurseInfo: cciptypes.CurseInfo{
				CursedSourceChains: map[cciptypes.ChainSelector]bool{},
				CursedDestination:  false,
				GlobalCurse:        false,
			},
		},
		{
			name: "everything cursed",
			cursedSubjectsSet: mapset.NewSet(
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
				chainSelectorToBytes16(chainA), // dest
				cciptypes.GlobalCurseSubject,
			),
			destChainSelector: chainA,
			expCurseInfo: cciptypes.CurseInfo{
				CursedSourceChains: map[cciptypes.ChainSelector]bool{
					chainB: true,
					chainC: true,
				},
				CursedDestination: true,
				GlobalCurse:       true,
			},
		},
		{
			name: "no global curse",
			cursedSubjectsSet: mapset.NewSet(
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
				chainSelectorToBytes16(chainA), // dest
			),
			destChainSelector: chainA,
			expCurseInfo: cciptypes.CurseInfo{
				CursedSourceChains: map[cciptypes.ChainSelector]bool{
					chainB: true,
					chainC: true,
				},
				CursedDestination: true,
				GlobalCurse:       false,
			},
		},
		{
			name: "dest cursed due to global curse",
			cursedSubjectsSet: mapset.NewSet(
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
				cciptypes.GlobalCurseSubject,
			),
			destChainSelector: chainA,
			expCurseInfo: cciptypes.CurseInfo{
				CursedSourceChains: map[cciptypes.ChainSelector]bool{
					chainB: true,
					chainC: true,
				},
				CursedDestination: true,
				GlobalCurse:       true,
			},
		},
		{
			name: "dest not cursed",
			cursedSubjectsSet: mapset.NewSet(
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
			),
			destChainSelector: chainA,
			expCurseInfo: cciptypes.CurseInfo{
				CursedSourceChains: map[cciptypes.ChainSelector]bool{
					chainB: true,
					chainC: true,
				},
				CursedDestination: false,
				GlobalCurse:       false,
			},
		},
		{
			name: "source chain B not cursed",
			cursedSubjectsSet: mapset.NewSet(
				chainSelectorToBytes16(chainC),
				chainSelectorToBytes16(chainA), // dest
				cciptypes.GlobalCurseSubject,
			),
			destChainSelector: chainA,
			expCurseInfo: cciptypes.CurseInfo{
				CursedSourceChains: map[cciptypes.ChainSelector]bool{
					chainC: true,
				},
				CursedDestination: true,
				GlobalCurse:       true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			curseInfo := getCurseInfoFromCursedSubjects(tc.cursedSubjectsSet, tc.destChainSelector)
			assert.Equal(t, tc.expCurseInfo, *curseInfo)
		})
	}
}

func TestCCIPChainReader_Nonces(t *testing.T) {
	type testCase struct {
		name           string
		funcInput      map[cciptypes.ChainSelector][]string
		expectedNonces map[cciptypes.ChainSelector]map[string]uint64
		matchedBy      func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool
		mockResults    []*uint64
	}

	// copies the behavior of the address codec
	// https://github.com/smartcontractkit/chainlink/blob/develop/integration-tests/smoke/ccip/ccip_reader_test.go#L2013
	transformAddress := func(addr string) cciptypes.UnknownAddress {
		addrBytes, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(addr, "0x")))
		assert.NoError(t, err)
		return slicelib.LeftPadBytes(addrBytes, 32)
	}

	var (
		addr1  = "0x1234567890123456789012345678901234567890"
		addr2  = "0x2234567890123456789012345678901234567890"
		addr3  = "0x3234567890123456789012345678901234567890"
		addr4  = "0x4234567890123456789012345678901234567890"
		nonce1 = uint64(5)
		nonce2 = uint64(10)
		nonce3 = uint64(15)
		nonce4 = uint64(20)
	)
	testCases := []testCase{
		{
			name:      "single chain, two addresses",
			funcInput: map[cciptypes.ChainSelector][]string{chainB: {addr1, addr2}},
			expectedNonces: map[cciptypes.ChainSelector]map[string]uint64{
				chainB: {
					addr1: 5,
					addr2: 10,
				},
			},
			mockResults: []*uint64{&nonce1, &nonce2},
			matchedBy: func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
				batch := req[consts.ContractNameNonceManager]

				assert.Len(t, batch, 2)
				assert.Equal(t, transformAddress(addr1), batch[0].Params.(map[string]any)["sender"])
				assert.Equal(t, transformAddress(addr2), batch[1].Params.(map[string]any)["sender"])
				assert.Equal(t, consts.MethodNameGetInboundNonce, batch[0].ReadName)
				assert.Equal(t, consts.MethodNameGetInboundNonce, batch[1].ReadName)

				return true
			},
		},
		{
			name:      "two chains",
			funcInput: map[cciptypes.ChainSelector][]string{chainA: {addr1, addr2}, chainB: {addr3, addr4}},
			expectedNonces: map[cciptypes.ChainSelector]map[string]uint64{
				chainA: {
					addr1: nonce1,
					addr2: nonce2,
				},
				chainB: {
					addr3: nonce3,
					addr4: nonce4,
				},
			},
			mockResults: []*uint64{&nonce1, &nonce2, &nonce3, &nonce4},
			matchedBy: func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
				batch := req[consts.ContractNameNonceManager]
				for _, batchRead := range batch {
					if batchRead.ReadName != consts.MethodNameGetInboundNonce {
						return false
					}
				}

				assert.Len(t, batch, 4)
				assert.Equal(t, transformAddress(addr1), batch[0].Params.(map[string]any)["sender"])
				assert.Equal(t, transformAddress(addr2), batch[1].Params.(map[string]any)["sender"])
				assert.Equal(t, transformAddress(addr3), batch[2].Params.(map[string]any)["sender"])
				assert.Equal(t, transformAddress(addr4), batch[3].Params.(map[string]any)["sender"])
				assert.Equal(t, chainA, batch[0].Params.(map[string]any)["sourceChainSelector"])
				assert.Equal(t, chainA, batch[1].Params.(map[string]any)["sourceChainSelector"])
				assert.Equal(t, chainB, batch[2].Params.(map[string]any)["sourceChainSelector"])
				assert.Equal(t, chainB, batch[3].Params.(map[string]any)["sourceChainSelector"])
				return true
			},
		},
		{
			name:      "same address multiple chains",
			funcInput: map[cciptypes.ChainSelector][]string{chainA: {addr1}, chainB: {addr1, addr2}},
			expectedNonces: map[cciptypes.ChainSelector]map[string]uint64{
				chainA: {
					addr1: nonce1,
				},
				chainB: {
					addr1: nonce2,
					addr2: nonce3,
				},
			},
			mockResults: []*uint64{&nonce1, &nonce2, &nonce3},
			matchedBy: func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
				batch := req[consts.ContractNameNonceManager]

				assert.Len(t, batch, 3)
				assert.Equal(t, transformAddress(addr1), batch[0].Params.(map[string]any)["sender"])
				assert.Equal(t, transformAddress(addr1), batch[1].Params.(map[string]any)["sender"])
				assert.Equal(t, transformAddress(addr2), batch[2].Params.(map[string]any)["sender"])
				assert.Equal(t, chainA, batch[0].Params.(map[string]any)["sourceChainSelector"])
				assert.Equal(t, chainB, batch[1].Params.(map[string]any)["sourceChainSelector"])
				assert.Equal(t, chainB, batch[2].Params.(map[string]any)["sourceChainSelector"])
				return true
			},
		},
		{
			name:      "empty chain",
			funcInput: map[cciptypes.ChainSelector][]string{chainA: {addr1}, chainB: {}},
			expectedNonces: map[cciptypes.ChainSelector]map[string]uint64{
				chainA: {
					addr1: nonce1,
				},
			},
			mockResults: []*uint64{&nonce1},
			matchedBy: func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
				batch := req[consts.ContractNameNonceManager]
				assert.Len(t, batch, 1)
				assert.Equal(t, transformAddress(addr1), batch[0].Params.(map[string]any)["sender"])
				assert.Equal(t, chainA, batch[0].Params.(map[string]any)["sourceChainSelector"])
				return true
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			destReader := readermocks.NewMockContractReaderFacade(t)
			cw := writermocks.NewMockContractWriter(t)
			contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
			contractWriters[chainB] = cw
			chainAccessors := createMockedChainAccessors(t, chainB)
			offRampAddress := []byte{0x3}
			mockExpectChainAccessorSyncCall(chainAccessors[chainB], consts.ContractNameOffRamp, offRampAddress, nil)
			mockExpectChainAccessorNoncesCall(chainAccessors[chainB], tc.expectedNonces)
			ccipReader, err := newCCIPChainReaderInternal(
				t.Context(),
				logger.Test(t),
				chainAccessors,
				map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
					chainB: destReader,
				},
				contractWriters,
				chainB,
				offRampAddress,
				internal.NewMockAddressCodecHex(t),
			)
			require.NoError(t, err)

			// Call Nonces
			nonces, err := ccipReader.Nonces(
				context.Background(),
				tc.funcInput,
			)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedNonces, nonces)

			err = ccipReader.Close()
			require.NoError(t, err)
		})
	}
}

func TestCCIPChainReader_DiscoverContracts_Parallel(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChains := []cciptypes.ChainSelector{2, 3, 4} // Adding one more chain for better parallelism testing

	// Setup mock cache and configurations
	mockCache := new(mockConfigCache)

	// Destination chain config with a delay
	destChainConfig := cciptypes.ChainConfigSnapshot{
		Offramp: cciptypes.OfframpConfig{
			StaticConfig: cciptypes.OffRampStaticChainConfig{
				NonceManager: []byte{0x3},
				RmnRemote:    []byte{0x4},
			},
			DynamicConfig: cciptypes.OffRampDynamicChainConfig{
				FeeQuoter: []byte{0x5},
			},
		},
	}

	sourceChainConfigs := make(map[cciptypes.ChainSelector]StaticSourceChainConfig, len(sourceChains))
	for i, chain := range sourceChains {
		sourceChainConfigs[chain] = StaticSourceChainConfig{
			Router:    []byte{0x6}, // Same router for all source chains
			IsEnabled: true,
			OnRamp:    []byte{byte(i + 1)}, // 0x1, 0x2, 0x3 as in the batch response
		}
	}

	// Set up destination chain expectation with delay
	mockCache.On("GetChainConfig", mock.Anything, destChain).
		Run(func(args mock.Arguments) {
			time.Sleep(75 * time.Millisecond) // Simulate network delay
		}).
		Return(destChainConfig, nil).Once()

	// Set up source chain expectations with delays
	for _, chain := range sourceChains {
		srcChainConfig := cciptypes.ChainConfigSnapshot{
			OnRamp: cciptypes.OnRampConfig{
				DynamicConfig: cciptypes.GetOnRampDynamicConfigResponse{
					DynamicConfig: cciptypes.OnRampDynamicConfig{
						FeeQuoter: []byte{0x7},
					},
				},
				DestChainConfig: cciptypes.OnRampDestChainConfig{
					Router: []byte{0x9},
				},
			},
		}
		mockCache.On("GetChainConfig", mock.Anything, chain).
			Run(func(args mock.Arguments) {
				time.Sleep(75 * time.Millisecond) // Simulate network delay
			}).
			Return(srcChainConfig, nil).Once()
	}

	// Setup contract readers
	mockReaders := make(map[cciptypes.ChainSelector]*readermocks.MockExtended)
	contractReaders := make(map[cciptypes.ChainSelector]contractreader.Extended)

	// Setup dest chain reader
	mockReaders[destChain] = readermocks.NewMockExtended(t)
	contractReaders[destChain] = mockReaders[destChain]

	// Setup source chain readers
	for _, chain := range sourceChains {
		mockReaders[chain] = readermocks.NewMockExtended(t)
		contractReaders[chain] = mockReaders[chain]
	}

	// Setup dest chain batch get values expectation
	mockCache.On("GetOfframpSourceChainConfigs", mock.Anything, destChain, sourceChains).
		Run(func(args mock.Arguments) {
			time.Sleep(100 * time.Millisecond) // Simulate network delay
		}).
		Return(sourceChainConfigs, nil).Once()

	ccipReader := &ccipChainReader{
		destChain:       destChain,
		contractReaders: contractReaders,
		lggr:            logger.Test(t),
		configPoller:    mockCache,
	}

	// Measure execution time
	start := time.Now()
	contractAddresses, err := ccipReader.DiscoverContracts(ctx, append(sourceChains, destChain), sourceChains)
	duration := time.Since(start)

	// Verify execution
	require.NoError(t, err)
	require.NotNil(t, contractAddresses)

	// If operations were sequential, it would take ~400ms (100ms * 4 chains)
	// With parallel execution, it should take ~100ms plus some overhead
	// Allow for reasonable overhead in CI environment
	assert.Less(t, duration, 300*time.Millisecond,
		"execution took too long, suggesting sequential rather than parallel execution")

	// Verify all mock expectations were met
	mockCache.AssertExpectations(t)

	// Verify the content of contractAddresses
	expectedCounts := map[string]struct {
		count    int
		chainsIn []cciptypes.ChainSelector
	}{
		consts.ContractNameOnRamp: {
			count:    len(sourceChains),
			chainsIn: sourceChains,
		},
		consts.ContractNameRouter: {
			count:    len(sourceChains) + 1,
			chainsIn: append(sourceChains, destChain),
		}, // source + dest
		consts.ContractNameFeeQuoter: {
			count:    len(sourceChains) + 1,
			chainsIn: append(sourceChains, destChain),
		}, // source + dest
		consts.ContractNameRMNRemote: {
			count:    1,
			chainsIn: []cciptypes.ChainSelector{destChain},
		}, // dest only
		consts.ContractNameNonceManager: {
			count:    1,
			chainsIn: []cciptypes.ChainSelector{destChain},
		}, // dest only
	}

	for contractName, expected := range expectedCounts {
		addresses := contractAddresses[contractName]
		assert.Lenf(t, addresses, expected.count,
			"expected %d addresses for %s, got %d", expected.count, contractName, len(addresses))

		// Verify contracts are on the correct chains
		for _, chain := range expected.chainsIn {
			assert.Contains(t, addresses, chain,
				"expected %s contract on chain %d", contractName, chain)
		}
	}
}

func TestCCIPChainReader_GetWrappedNativeTokenPriceUSD(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)

	wrappedNative1 := cciptypes.Bytes{0x1}
	wrappedNative2 := cciptypes.Bytes{0x2}

	t.Run("happy path - gets prices for all chains", func(t *testing.T) {
		// Setup mock cache with configs containing wrapped native addresses
		mockCache := new(mockConfigCache)
		sourceChain1Config := cciptypes.ChainConfigSnapshot{
			Router: cciptypes.RouterConfig{
				WrappedNativeAddress: wrappedNative1,
			},
		}
		sourceChain2Config := cciptypes.ChainConfigSnapshot{
			Router: cciptypes.RouterConfig{
				WrappedNativeAddress: wrappedNative2,
			},
		}

		mockCache.On("GetChainConfig", mock.Anything, sourceChain1).Return(sourceChain1Config, nil)
		mockCache.On("GetChainConfig", mock.Anything, sourceChain2).Return(sourceChain2Config, nil)
		mockCache.On("Start", mock.Anything).Return(nil)

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[sourceChain1] = cw
		contractWriters[sourceChain2] = cw

		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1, sourceChain2)
		mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRampAddress, nil)
		mockExpectChainAccessorGetTokenPriceUSD(
			chainAccessors[sourceChain1],
			cciptypes.UnknownAddress(wrappedNative1),
			cciptypes.TimestampedUnixBig{Value: big.NewInt(100), Timestamp: uint32(time.Now().Unix())},
		)
		mockExpectChainAccessorGetTokenPriceUSD(
			chainAccessors[sourceChain2],
			cciptypes.UnknownAddress(wrappedNative2),
			cciptypes.TimestampedUnixBig{Value: big.NewInt(200), Timestamp: uint32(time.Now().Unix())},
		)
		ccipReader, err := newCCIPChainReaderWithConfigPollerInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				sourceChain1: readermocks.NewMockContractReaderFacade(t),
				sourceChain2: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			destChain,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
			mockCache,
		)
		require.NoError(t, err)

		prices := ccipReader.GetWrappedNativeTokenPriceUSD(ctx, []cciptypes.ChainSelector{sourceChain1, sourceChain2})
		require.Len(t, prices, 2)
		assert.Equal(t, cciptypes.NewBigInt(big.NewInt(100)), prices[sourceChain1])
		assert.Equal(t, cciptypes.NewBigInt(big.NewInt(200)), prices[sourceChain2])

		mockCache.AssertExpectations(t)
	})

	t.Run("handles missing chain configs", func(t *testing.T) {
		mockCache := new(mockConfigCache)
		mockCache.On("GetChainConfig", mock.Anything, sourceChain1).Return(cciptypes.ChainConfigSnapshot{},
			fmt.Errorf("not found"))
		mockCache.On("GetChainConfig", mock.Anything, sourceChain2).Return(cciptypes.ChainConfigSnapshot{
			Router: cciptypes.RouterConfig{
				WrappedNativeAddress: wrappedNative2,
			},
		}, nil)
		mockCache.On("Start", mock.Anything).Return(nil)

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[sourceChain1] = cw
		contractWriters[sourceChain2] = cw

		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1, sourceChain2)
		mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRampAddress, nil)
		mockExpectChainAccessorGetTokenPriceUSD(
			chainAccessors[sourceChain2],
			cciptypes.UnknownAddress(wrappedNative2),
			cciptypes.TimestampedUnixBig{Value: big.NewInt(200), Timestamp: uint32(time.Now().Unix())},
		)
		ccipReader, err := newCCIPChainReaderWithConfigPollerInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				sourceChain1: readermocks.NewMockContractReaderFacade(t),
				sourceChain2: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			destChain,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
			mockCache,
		)
		require.NoError(t, err)

		prices := ccipReader.GetWrappedNativeTokenPriceUSD(ctx, []cciptypes.ChainSelector{sourceChain1, sourceChain2})
		require.Len(t, prices, 1)
		assert.Equal(t, cciptypes.NewBigInt(big.NewInt(200)), prices[sourceChain2])

		mockCache.AssertExpectations(t)
	})

	t.Run("handles price fetch error", func(t *testing.T) {
		mockCache := new(mockConfigCache)
		sourceConfig := cciptypes.ChainConfigSnapshot{
			Router: cciptypes.RouterConfig{
				WrappedNativeAddress: wrappedNative1,
			},
		}
		mockCache.On("GetChainConfig", mock.Anything, sourceChain1).Return(sourceConfig, nil)
		mockCache.On("Start", mock.Anything).Return(nil)

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[sourceChain1] = cw

		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1, sourceChain2)
		mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRampAddress, nil)
		chainAccessors[sourceChain1].(*commonccipocr3.MockChainAccessor).EXPECT().
			GetTokenPriceUSD(mock.Anything, cciptypes.UnknownAddress(wrappedNative1)).
			Return(cciptypes.TimestampedUnixBig{}, fmt.Errorf("price fetch failed"))
		ccipReader, err := newCCIPChainReaderWithConfigPollerInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				sourceChain1: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			destChain,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
			mockCache,
		)
		require.NoError(t, err)

		prices := ccipReader.GetWrappedNativeTokenPriceUSD(ctx, []cciptypes.ChainSelector{sourceChain1})
		require.Empty(t, prices)

		mockCache.AssertExpectations(t)
	})
}

func TestCCIPChainReader_prepareBatchConfigRequests(t *testing.T) {
	destChain := cciptypes.ChainSelector(1)
	sourceChain := cciptypes.ChainSelector(2)

	ccipReader := &ccipChainReader{
		destChain: destChain,
	}

	t.Run("source chain requests", func(t *testing.T) {
		requests := ccipReader.prepareBatchConfigRequests(sourceChain)

		// Should contain OnRamp and Router requests
		require.Len(t, requests, 2)
		require.Contains(t, requests, consts.ContractNameOnRamp)
		require.Contains(t, requests, consts.ContractNameRouter)

		onRampRequests := requests[consts.ContractNameOnRamp]
		require.Len(t, onRampRequests, 2)

		// Verify OnRamp dynamic config request
		require.Equal(t, consts.MethodNameOnRampGetDynamicConfig, onRampRequests[0].ReadName)
		require.Empty(t, onRampRequests[0].Params)
		require.IsType(t, &cciptypes.GetOnRampDynamicConfigResponse{}, onRampRequests[0].ReturnVal)

		// Verify OnRamp dest chain config request
		require.Equal(t, consts.MethodNameOnRampGetDestChainConfig, onRampRequests[1].ReadName)
		require.Equal(t, map[string]any{"destChainSelector": destChain}, onRampRequests[1].Params)
		require.IsType(t, &cciptypes.OnRampDestChainConfig{}, onRampRequests[1].ReturnVal)

		// Verify Router requests
		routerRequests := requests[consts.ContractNameRouter]
		require.Len(t, routerRequests, 1)
		require.Equal(t, consts.MethodNameRouterGetWrappedNative, routerRequests[0].ReadName)
		require.Empty(t, routerRequests[0].Params)
		require.IsType(t, &[]byte{}, routerRequests[0].ReturnVal)
	})

	t.Run("destination chain requests", func(t *testing.T) {
		requests := ccipReader.prepareBatchConfigRequests(destChain)

		// Should contain all contract requests except OnRamp and Router
		require.Len(t, requests, 4)
		require.Contains(t, requests, consts.ContractNameOffRamp)
		require.Contains(t, requests, consts.ContractNameRMNProxy)
		require.Contains(t, requests, consts.ContractNameRMNRemote)
		require.Contains(t, requests, consts.ContractNameFeeQuoter)
		require.NotContains(t, requests, consts.ContractNameOnRamp)
		require.NotContains(t, requests, consts.ContractNameRouter)

		// Verify OffRamp requests
		offRampRequests := requests[consts.ContractNameOffRamp]
		require.Len(t, offRampRequests, 4)

		// Check OffRamp commit config request
		require.Equal(t, consts.MethodNameOffRampLatestConfigDetails, offRampRequests[0].ReadName)
		require.Equal(t, map[string]any{"ocrPluginType": consts.PluginTypeCommit}, offRampRequests[0].Params)
		require.IsType(t, &cciptypes.OCRConfigResponse{}, offRampRequests[0].ReturnVal)

		// Check OffRamp execute config request
		require.Equal(t, consts.MethodNameOffRampLatestConfigDetails, offRampRequests[1].ReadName)
		require.Equal(t, map[string]any{"ocrPluginType": consts.PluginTypeExecute}, offRampRequests[1].Params)
		require.IsType(t, &cciptypes.OCRConfigResponse{}, offRampRequests[1].ReturnVal)

		// Verify RMN requests
		rmnProxyRequests := requests[consts.ContractNameRMNProxy]
		require.Len(t, rmnProxyRequests, 1)
		require.Equal(t, consts.MethodNameGetARM, rmnProxyRequests[0].ReadName)
		require.Empty(t, rmnProxyRequests[0].Params)
		require.IsType(t, &[]byte{}, rmnProxyRequests[0].ReturnVal)

		// Verify FeeQuoter request
		feeQuoterRequests := requests[consts.ContractNameFeeQuoter]
		require.Len(t, feeQuoterRequests, 1)
		require.Equal(t, consts.MethodNameFeeQuoterGetStaticConfig, feeQuoterRequests[0].ReadName)
		require.Empty(t, feeQuoterRequests[0].Params)
		require.IsType(t, &cciptypes.FeeQuoterStaticConfig{}, feeQuoterRequests[0].ReturnVal)
	})
}

func TestCCIPChainReader_GetChainFeePriceUpdate(t *testing.T) {
	ctx := t.Context()
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	sourceChain3 := cciptypes.ChainSelector(4) // Chain with zero gas price result

	lggr := logger.Test(t)

	t.Run("happy path", func(t *testing.T) {
		selectors := []cciptypes.ChainSelector{sourceChain1, sourceChain2}

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[chainB] = cw
		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, chainB, sourceChain1, sourceChain2)
		mockExpectChainAccessorSyncCall(chainAccessors[chainB], consts.ContractNameOffRamp, offRampAddress, nil)
		expectedResult := map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig{
			sourceChain1: {Value: big.NewInt(100), Timestamp: uint32(time.Now().Unix())},
			sourceChain2: {Value: big.NewInt(200), Timestamp: uint32(time.Now().Unix())},
		}
		mockExpectChainAccessorGetChainFeePriceUpdate(chainAccessors[chainB], selectors, expectedResult)
		ccipReader, err := newCCIPChainReaderInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				chainB: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			chainB,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
		)
		require.NoError(t, err)

		feeUpdates := ccipReader.GetChainFeePriceUpdate(ctx, selectors)

		require.Len(t, feeUpdates, 2)
		assert.NotNil(t, feeUpdates[sourceChain1].Value)
		assert.Equal(t, 0, feeUpdates[sourceChain1].Value.Cmp(big.NewInt(100)))
		assert.NotZero(t, feeUpdates[sourceChain1].Timestamp)
		assert.NotNil(t, feeUpdates[sourceChain2].Value)
		assert.Equal(t, 0, feeUpdates[sourceChain2].Value.Cmp(big.NewInt(200)))
		assert.NotZero(t, feeUpdates[sourceChain2].Timestamp)

		err = ccipReader.Close()
		require.NoError(t, err)
	})

	t.Run("empty selectors", func(t *testing.T) {
		mockReader := readermocks.NewMockExtended(t)
		ccipReader := &ccipChainReader{
			lggr:      lggr,
			destChain: destChain,
			contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
				destChain: mockReader,
			},
		}

		feeUpdates := ccipReader.GetChainFeePriceUpdate(ctx, []cciptypes.ChainSelector{})
		require.Empty(t, feeUpdates)
		// No calls expected to the reader
		mockReader.AssertExpectations(t)
	})

	t.Run("batch call error", func(t *testing.T) {
		selectors := []cciptypes.ChainSelector{sourceChain1}

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[destChain] = cw
		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1)
		mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRampAddress, nil)
		chainAccessors[destChain].(*commonccipocr3.MockChainAccessor).EXPECT().
			GetChainFeePriceUpdate(mock.Anything, selectors).
			Return(map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig{}, nil)
		ccipReader, err := newCCIPChainReaderInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				destChain: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			destChain,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
		)
		require.NoError(t, err)

		feeUpdates := ccipReader.GetChainFeePriceUpdate(ctx, selectors)
		require.Empty(t, feeUpdates)

		err = ccipReader.Close()
		require.NoError(t, err)
	})

	t.Run("partial success - one result empty", func(t *testing.T) {
		selectors := []cciptypes.ChainSelector{sourceChain1, sourceChain3}

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[destChain] = cw
		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1, sourceChain3)
		mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRampAddress, nil)
		expectedResult := map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig{
			sourceChain1: {Value: big.NewInt(100), Timestamp: uint32(time.Now().Unix())},
			sourceChain3: {Value: big.NewInt(0), Timestamp: uint32(time.Now().Unix())},
		}
		mockExpectChainAccessorGetChainFeePriceUpdate(chainAccessors[destChain], selectors, expectedResult)
		ccipReader, err := newCCIPChainReaderInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				destChain: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			destChain,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
		)
		require.NoError(t, err)

		feeUpdates := ccipReader.GetChainFeePriceUpdate(ctx, selectors)

		require.Len(t, feeUpdates, 2)
		require.Contains(t, feeUpdates, sourceChain1)
		assert.NotNil(t, feeUpdates[sourceChain1].Value)
		assert.Equal(t, 0, feeUpdates[sourceChain1].Value.Cmp(big.NewInt(100)))
		assert.Contains(t, feeUpdates, sourceChain3)

		err = ccipReader.Close()
		require.NoError(t, err)
	})

	t.Run("result count mismatch", func(t *testing.T) {
		// Request two selectors, but mock response only has one result
		selectors := []cciptypes.ChainSelector{sourceChain1, sourceChain2}

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[destChain] = cw
		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1, sourceChain2)
		mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRampAddress, nil)
		expectedResult := map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig{
			sourceChain1: {Value: big.NewInt(100), Timestamp: uint32(time.Now().Unix())},
		}
		mockExpectChainAccessorGetChainFeePriceUpdate(chainAccessors[destChain], selectors, expectedResult)
		ccipReader, err := newCCIPChainReaderInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				destChain: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			destChain,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
		)
		require.NoError(t, err)

		feeUpdates := ccipReader.GetChainFeePriceUpdate(ctx, selectors)

		// Should still process the results it received
		require.Len(t, feeUpdates, 1)
		require.Contains(t, feeUpdates, sourceChain1)
		assert.NotNil(t, feeUpdates[sourceChain1].Value)
		assert.Equal(t, 0, feeUpdates[sourceChain1].Value.Cmp(big.NewInt(100)))
		assert.NotContains(t, feeUpdates, sourceChain2)

		err = ccipReader.Close()
		require.NoError(t, err)
	})

	t.Run("missing fee quoter result in batch response", func(t *testing.T) {
		selectors := []cciptypes.ChainSelector{sourceChain1}

		cw := writermocks.NewMockContractWriter(t)
		contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
		contractWriters[destChain] = cw
		offRampAddress := []byte{0x3}
		chainAccessors := createMockedChainAccessors(t, destChain, sourceChain1)
		mockExpectChainAccessorSyncCall(chainAccessors[destChain], consts.ContractNameOffRamp, offRampAddress, nil)
		mockExpectChainAccessorGetChainFeePriceUpdate(
			chainAccessors[destChain],
			selectors,
			map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig{},
		)
		ccipReader, err := newCCIPChainReaderInternal(
			t.Context(),
			logger.Test(t),
			chainAccessors,
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				destChain: readermocks.NewMockContractReaderFacade(t),
			},
			contractWriters,
			destChain,
			offRampAddress,
			internal.NewMockAddressCodecHex(t),
		)
		require.NoError(t, err)

		feeUpdates := ccipReader.GetChainFeePriceUpdate(ctx, selectors)
		require.Empty(t, feeUpdates)

		err = ccipReader.Close()
		require.NoError(t, err)
	})

	t.Run("accessor does not exist for dest chain", func(t *testing.T) {
		ccipReader, err := newCCIPChainReaderInternal(
			t.Context(),
			logger.Test(t),
			nil,
			nil,
			nil,
			destChain,
			[]byte("0x3"),
			internal.NewMockAddressCodecHex(t),
		)
		require.NoError(t, err)

		feeUpdates := ccipReader.GetChainFeePriceUpdate(ctx, []cciptypes.ChainSelector{sourceChain1})
		// Original logic returned nil in this case
		assert.Nil(t, feeUpdates)

		err = ccipReader.Close()
		require.NoError(t, err)
	})
}

func createMockedChainAccessors(
	t *testing.T,
	chains ...cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.ChainAccessor {
	chainAccessors := make(map[cciptypes.ChainSelector]cciptypes.ChainAccessor)
	for _, chain := range chains {
		mockAccessor := commonccipocr3.NewMockChainAccessor(t)
		chainAccessors[chain] = mockAccessor
	}
	return chainAccessors
}

func mockExpectChainAccessorSyncCall(
	chainAccessor cciptypes.ChainAccessor,
	expectedContractName string,
	expectedContractAddress []byte,
	err error,
) {
	chainAccessor.(*commonccipocr3.MockChainAccessor).EXPECT().
		Sync(mock.Anything, expectedContractName, cciptypes.UnknownAddress(expectedContractAddress)).
		Once().Return(err)
}

func mockExpectChainAccessorNoncesCall(
	chainAccessor cciptypes.ChainAccessor,
	expectedResult map[cciptypes.ChainSelector]map[string]uint64,
) {
	chainAccessor.(*commonccipocr3.MockChainAccessor).EXPECT().
		Nonces(mock.Anything, mock.Anything).Return(expectedResult, nil)
}

func mockExpectChainAccessorGetTokenPriceUSD(
	chainAccessor cciptypes.ChainAccessor,
	token cciptypes.UnknownAddress,
	price cciptypes.TimestampedUnixBig,
) {
	chainAccessor.(*commonccipocr3.MockChainAccessor).EXPECT().
		GetTokenPriceUSD(mock.Anything, token).Return(price, nil)
}

func mockExpectChainAccessorGetChainFeeComponentsCall(
	chainAccessor cciptypes.ChainAccessor,
	feeComponents cciptypes.ChainFeeComponents,
	err error,
) {
	chainAccessor.(*commonccipocr3.MockChainAccessor).EXPECT().
		GetChainFeeComponents(mock.Anything).Return(feeComponents, err)
}

func mockExpectChainAccessorGetChainFeePriceUpdate(
	chainAccessor cciptypes.ChainAccessor,
	selectors []cciptypes.ChainSelector,
	expectedResult map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig,
) {
	chainAccessor.(*commonccipocr3.MockChainAccessor).EXPECT().
		GetChainFeePriceUpdate(mock.Anything, selectors).Return(expectedResult, nil)
}

type mockConfigCache struct {
	mock.Mock
}

func (m *mockConfigCache) GetChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector) (cciptypes.ChainConfigSnapshot, error) {
	args := m.Called(ctx, chainSel)
	return args.Get(0).(cciptypes.ChainConfigSnapshot), args.Error(1)
}

func (m *mockConfigCache) GetOfframpSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	args := m.Called(ctx, destChain, sourceChains)
	return args.Get(0).(map[cciptypes.ChainSelector]StaticSourceChainConfig), args.Error(1)
}

// Update Start method to accept context parameter
func (m *mockConfigCache) Start(ctx context.Context) error {
	return m.Called(ctx).Error(0)
}

func (m *mockConfigCache) Close() error {
	return m.Called().Error(0)
}

// Implement HealthReport method for services.Service interface
func (m *mockConfigCache) HealthReport() map[string]error {
	args := m.Called()
	return args.Get(0).(map[string]error)
}

// Implement Name method for the Service interface
func (m *mockConfigCache) Name() string {
	return m.Called().String(0)
}

func (m *mockConfigCache) Ready() error {
	return m.Called().Error(0)
}
