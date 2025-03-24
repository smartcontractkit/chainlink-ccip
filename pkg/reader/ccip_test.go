package reader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-ccip/internal"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	writer_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common"
	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	chainA = cciptypes.ChainSelector(1)
	chainB = cciptypes.ChainSelector(2)
	chainC = cciptypes.ChainSelector(3)
	chainD = cciptypes.ChainSelector(4)
)

func TestCCIPChainReader_getSourceChainsConfig(t *testing.T) {
	sourceCRs := make(map[cciptypes.ChainSelector]*reader_mocks.MockContractReaderFacade)
	for _, chain := range []cciptypes.ChainSelector{chainA, chainB} {
		sourceCRs[chain] = reader_mocks.NewMockContractReaderFacade(t)
		sourceCRs[chain].EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	}

	destCR := reader_mocks.NewMockContractReaderFacade(t)
	destCR.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	destCR.EXPECT().HealthReport().Return(nil)
	destCR.EXPECT().BatchGetLatestValues(
		mock.Anything,
		mock.Anything,
	).RunAndReturn(func(
		ctx context.Context,
		request types.BatchGetLatestValuesRequest,
	) (types.BatchGetLatestValuesResult, error) {
		results := make(types.BatchGetLatestValuesResult, 0)
		for contractName, batch := range request {
			for _, readReq := range batch {
				res := types.BatchReadResult{
					ReadName: readReq.ReadName,
				}
				params := readReq.Params.(map[string]any)
				sourceChain := params["sourceChainSelector"].(cciptypes.ChainSelector)
				v := readReq.ReturnVal.(*SourceChainConfig)

				fromString, err := cciptypes.NewBytesFromString(fmt.Sprintf(
					"0x%d000000000000000000000000000000000000000", sourceChain),
				)
				require.NoError(t, err)
				v.OnRamp = cciptypes.UnknownAddress(fromString)
				v.IsEnabled = true
				v.Router = fromString
				res.SetResult(v, nil)
				results[contractName] = append(results[contractName], res)
			}
		}
		return results, nil
	})

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	offrampAddress := []byte{0x3}
	ccipReader := newCCIPChainReaderInternal(
		tests.Context(t),
		logger.Test(t),
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			chainA: sourceCRs[chainA],
			chainB: sourceCRs[chainB],
			chainC: destCR,
		}, nil, chainC, offrampAddress, mockAddrCodec,
	)

	// Add cleanup to ensure resources are released
	t.Cleanup(func() {
		err := ccipReader.Close()
		if err != nil {
			t.Logf("Error closing ccipReader: %v", err)
		}
	})

	addrStr, err := mockAddrCodec.AddressBytesToString(offrampAddress, 111_111)
	require.NoError(t, err)

	require.NoError(t, ccipReader.contractReaders[chainA].Bind(
		context.Background(), []types.BoundContract{{Name: "OnRamp", Address: "0x1"}}))
	require.NoError(t, ccipReader.contractReaders[chainB].Bind(
		context.Background(), []types.BoundContract{{Name: "OnRamp", Address: "0x2"}}))
	require.NoError(t, ccipReader.contractReaders[chainC].Bind(
		context.Background(), []types.BoundContract{{Name: "OffRamp",
			Address: addrStr}}))

	ctx := context.Background()
	cfgs, err := ccipReader.getOffRampSourceChainsConfig(
		ctx, logger.Test(t), []cciptypes.ChainSelector{chainA, chainB}, false)
	assert.NoError(t, err)
	assert.Len(t, cfgs, 2)
	assert.Equal(t, "0x1000000000000000000000000000000000000000", cfgs[chainA].OnRamp.String())
	assert.Equal(t, "0x2000000000000000000000000000000000000000", cfgs[chainB].OnRamp.String())
}

func TestCCIPChainReader_GetContractAddress(t *testing.T) {
	ecr := reader_mocks.NewMockExtended(t)

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	ccipReader := ccipChainReader{
		lggr: logger.Test(t),
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			chainA: ecr,
		},
		addrCodec: mockAddrCodec,
	}

	someAddr := "0x1234567890123456789012345678901234567890"
	someAddrBytes, err := mockAddrCodec.AddressStringToBytes(someAddr, chainA)
	require.NoError(t, err)

	t.Run("happy path", func(t *testing.T) {
		ecr.EXPECT().GetBindings(consts.ContractNameOnRamp).Return([]contractreader.ExtendedBoundContract{
			{
				BoundAt: time.Now().UTC(),
				Binding: types.BoundContract{Address: someAddr, Name: consts.ContractNameOnRamp},
			},
		}).Once()
		addr, err := ccipReader.GetContractAddress(consts.ContractNameOnRamp, chainA)
		assert.NoError(t, err)
		assert.Equal(t, someAddrBytes, cciptypes.UnknownAddress(addr))
	})

	t.Run("multiple bindings leads to error", func(t *testing.T) {
		ecr.EXPECT().GetBindings(consts.ContractNameOnRamp).Return([]contractreader.ExtendedBoundContract{
			{
				BoundAt: time.Now().UTC(),
				Binding: types.BoundContract{Address: someAddr, Name: consts.ContractNameOnRamp},
			},
			{
				BoundAt: time.Now().UTC(),
				Binding: types.BoundContract{Address: someAddr, Name: consts.ContractNameOnRamp},
			},
		}).Once()
		_, err := ccipReader.GetContractAddress(consts.ContractNameOnRamp, chainA)
		assert.Error(t, err)
	})

	t.Run("no binding leads to error", func(t *testing.T) {
		ecr.EXPECT().GetBindings(consts.ContractNameOnRamp).Return([]contractreader.ExtendedBoundContract{}).Once()
		_, err := ccipReader.GetContractAddress(consts.ContractNameOnRamp, chainA)
		assert.Error(t, err)
	})

	t.Run("invalid address leads to error", func(t *testing.T) {
		ecr.EXPECT().GetBindings(consts.ContractNameOnRamp).Return([]contractreader.ExtendedBoundContract{
			{
				BoundAt: time.Now().UTC(),
				Binding: types.BoundContract{Address: "some wrong address fmt", Name: consts.ContractNameOnRamp},
			},
		}).Once()
		_, err := ccipReader.GetContractAddress(consts.ContractNameOnRamp, chainA)
		assert.Error(t, err)
	})
}

func TestCCIPChainReader_Sync_HappyPath_BindsContractsSuccessfully(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destNonceMgr := []byte{0x3}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	destNonceMgrAddrStr, err := mockAddrCodec.AddressBytesToString(destNonceMgr, destChain)
	require.NoError(t, err)
	destExtended := reader_mocks.NewMockExtended(t)
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: destNonceMgrAddrStr,
		},
	}).Return(nil)

	s1OnrampAddrStr, err := mockAddrCodec.AddressBytesToString(s1Onramp, sourceChain1)
	require.NoError(t, err)
	source1Extended := reader_mocks.NewMockExtended(t)
	source1Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: s1OnrampAddrStr,
		},
	}).Return(nil)

	sourceChain2AddrStr, err := mockAddrCodec.AddressBytesToString(s2Onramp, sourceChain2)
	require.NoError(t, err)
	source2Extended := reader_mocks.NewMockExtended(t)
	source2Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: sourceChain2AddrStr,
		},
	}).Return(nil)

	defer destExtended.AssertExpectations(t)
	defer source1Extended.AssertExpectations(t)
	defer source2Extended.AssertExpectations(t)

	ccipReader := &ccipChainReader{
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain:    destExtended,
			sourceChain1: source1Extended,
			sourceChain2: source2Extended,
		},
		destChain: destChain,
		lggr:      logger.Test(t),
		addrCodec: mockAddrCodec,
	}

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
}

func TestCCIPChainReader_Sync_HappyPath_SkipsEmptyAddress(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}

	// empty address, should get skipped
	s2Onramp := []byte{}

	destNonceMgr := []byte{0x3}
	destExtended := reader_mocks.NewMockExtended(t)
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	destNonceMgrAddrStr, err := mockAddrCodec.AddressBytesToString(destNonceMgr, destChain)
	require.NoError(t, err)
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: destNonceMgrAddrStr,
		},
	}).Return(nil)

	sourceChain1AddrStr, err := mockAddrCodec.AddressBytesToString(s1Onramp, sourceChain1)
	require.NoError(t, err)
	source1Extended := reader_mocks.NewMockExtended(t)
	source1Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: sourceChain1AddrStr,
		},
	}).Return(nil)

	// bind should not be called on this one.
	source2Extended := reader_mocks.NewMockExtended(t)

	defer destExtended.AssertExpectations(t)
	defer source1Extended.AssertExpectations(t)
	defer source2Extended.AssertExpectations(t)

	ccipReader := &ccipChainReader{
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain:    destExtended,
			sourceChain1: source1Extended,
			sourceChain2: source2Extended,
		},
		destChain: destChain,
		lggr:      logger.Test(t),
		addrCodec: mockAddrCodec,
	}

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
}

func TestCCIPChainReader_Sync_HappyPath_DontSupportAllChains(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destNonceMgr := []byte{0x3}
	destExtended := reader_mocks.NewMockExtended(t)
	mockAddrCodec := internal.NewMockAddressCodecHex(t)

	destNonceMgrAddrStr, err := mockAddrCodec.AddressBytesToString(destNonceMgr, destChain)
	require.NoError(t, err)
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: destNonceMgrAddrStr,
		},
	}).Return(nil)

	// only support source2, source1 unsupported.
	sourceChain2AddrStr, err := mockAddrCodec.AddressBytesToString(s2Onramp, sourceChain2)
	require.NoError(t, err)
	source2Extended := reader_mocks.NewMockExtended(t)
	source2Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: sourceChain2AddrStr,
		},
	}).Return(nil)

	defer destExtended.AssertExpectations(t)
	defer source2Extended.AssertExpectations(t)

	ccipReader := &ccipChainReader{
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain:    destExtended,
			sourceChain2: source2Extended,
		},
		destChain: destChain,
		lggr:      logger.Test(t),
		addrCodec: mockAddrCodec,
	}

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
}

func TestCCIPChainReader_Sync_BindError(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destNonceMgr := []byte{0x3}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	destNonceMgrAddrStr, err := mockAddrCodec.AddressBytesToString(destNonceMgr, destChain)
	require.NoError(t, err)
	destExtended := reader_mocks.NewMockExtended(t)
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: destNonceMgrAddrStr,
		},
	}).Return(nil)

	s1OnrampAddrStr, err := mockAddrCodec.AddressBytesToString(s1Onramp, sourceChain1)
	require.NoError(t, err)
	expectedErr := errors.New("some error")
	source1Extended := reader_mocks.NewMockExtended(t)
	source1Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: s1OnrampAddrStr,
		},
	}).Return(expectedErr)

	s2OnrampAddrStr, err := mockAddrCodec.AddressBytesToString(s2Onramp, sourceChain2)
	require.NoError(t, err)
	source2Extended := reader_mocks.NewMockExtended(t)
	source2Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: s2OnrampAddrStr,
		},
	}).Return(nil)

	defer destExtended.AssertExpectations(t)
	defer source1Extended.AssertExpectations(t)
	defer source2Extended.AssertExpectations(t)

	ccipReader := &ccipChainReader{
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain:    destExtended,
			sourceChain1: source1Extended,
			sourceChain2: source2Extended,
		},
		destChain: destChain,
		lggr:      logger.Test(t),
		addrCodec: mockAddrCodec,
	}

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
}

// The round1 version returns NoBindingFound errors for onramp contracts to simulate
// the two-phase approach to discovering those contracts.
func TestCCIPChainReader_DiscoverContracts_HappyPath_Round1(t *testing.T) {
	ctx := tests.Context(t)
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

	mockReaders := make(map[cciptypes.ChainSelector]*reader_mocks.MockExtended)
	mockReaders[destChain] = reader_mocks.NewMockExtended(t)

	// Setup cache mock and configurations
	mockCache := new(mockConfigCache)
	// Destination chain config
	destChainConfig := ChainConfigSnapshot{
		Offramp: OfframpConfig{
			StaticConfig: offRampStaticChainConfig{
				NonceManager: destNonceMgr,
				RmnRemote:    destRMNRemote,
			},
			DynamicConfig: offRampDynamicChainConfig{
				FeeQuoter: destFeeQuoter,
			},
		},
	}
	// Set up cache expectations for destination chain and source chains
	mockCache.On("GetChainConfig", mock.Anything, destChain).Return(destChainConfig, nil).Once()
	mockCache.On("GetChainConfig", mock.Anything, sourceChain[0]).Return(
		ChainConfigSnapshot{}, contractreader.ErrNoBindings).Maybe()
	mockCache.On("GetChainConfig", mock.Anything, sourceChain[1]).Return(
		ChainConfigSnapshot{}, contractreader.ErrNoBindings).Maybe()
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
		lggr:            lggr,
		configPoller:    mockCache,
	}

	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx, sourceChain[:])
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
	ctx := tests.Context(t)
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

	mockReaders := make(map[cciptypes.ChainSelector]*reader_mocks.MockExtended)
	mockReaders[destChain] = reader_mocks.NewMockExtended(t)

	// Setup cache mock and configurations
	mockCache := new(mockConfigCache)
	// Destination chain config
	destChainConfig := ChainConfigSnapshot{
		Offramp: OfframpConfig{
			StaticConfig: offRampStaticChainConfig{
				NonceManager: destNonceMgr,
				RmnRemote:    destRMNRemote,
			},
			DynamicConfig: offRampDynamicChainConfig{
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
		mockReaders[chain] = reader_mocks.NewMockExtended(t)

		srcChainConfig := ChainConfigSnapshot{
			OnRamp: OnRampConfig{
				DynamicConfig: getOnRampDynamicConfigResponse{
					DynamicConfig: onRampDynamicConfig{
						FeeQuoter: srcFeeQuoters[i],
					},
				},
				DestChainConfig: onRampDestChainConfig{
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

	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx, sourceChain[:])
	require.NoError(t, err)
	require.Equal(t, expectedContractAddresses, contractAddresses)
	mockCache.AssertExpectations(t)
}

func TestCCIPChainReader_DiscoverContracts_GetAllSourceChainConfig_Errors(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)

	// Setup cache mock and configuration
	mockCache := new(mockConfigCache)
	chainConfig := ChainConfigSnapshot{
		Offramp: OfframpConfig{
			// We can leave the configs empty since we just need GetChainConfig to succeed
			StaticConfig:  offRampStaticChainConfig{},
			DynamicConfig: offRampDynamicChainConfig{},
		},
	}
	mockCache.On("GetChainConfig", mock.Anything, destChain).Return(chainConfig, nil)

	// Setup mock cache to return an error
	destExtended := reader_mocks.NewMockExtended(t)
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
			sourceChain1: reader_mocks.NewMockExtended(t),
			sourceChain2: reader_mocks.NewMockExtended(t),
		},
		lggr:         logger.Test(t),
		configPoller: mockCache,
	}

	_, err := ccipChainReader.DiscoverContracts(ctx, []cciptypes.ChainSelector{sourceChain1, sourceChain2})
	require.Error(t, err)
	require.ErrorIs(t, err, getLatestValueErr)
	mockCache.AssertExpectations(t)
}

func TestCCIPChainReader_DiscoverContracts_GetOfframpStaticConfig_Errors(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)

	// Setup mock cache to return error
	// mock the call to get the static config - failure
	getLatestValueErr := errors.New("some error")
	mockCache := new(mockConfigCache)
	mockCache.On("GetChainConfig", mock.Anything, destChain).Return(ChainConfigSnapshot{}, getLatestValueErr)

	// create the reader with cache
	ccipChainReader := &ccipChainReader{
		destChain: destChain,
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain: reader_mocks.NewMockExtended(t),
			// these won't be used in this test, but are needed because
			// we determine the source chain selectors to query from the chains
			// that we have readers for.
			sourceChain1: reader_mocks.NewMockExtended(t),
			sourceChain2: reader_mocks.NewMockExtended(t),
		},
		lggr:         logger.Test(t),
		configPoller: mockCache,
	}

	_, err := ccipChainReader.DiscoverContracts(ctx, []cciptypes.ChainSelector{sourceChain1, sourceChain2})
	require.Error(t, err)
	require.ErrorIs(t, err, getLatestValueErr)
	mockCache.AssertExpectations(t)
}

// withReturnValueOverridden is a helper function to override the return value of a mocked out
// ExtendedGetLatestValue call.
func withReturnValueOverridden(mapper func(returnVal interface{})) func(ctx context.Context,
	contractName,
	methodName string,
	confidenceLevel primitives.ConfidenceLevel,
	params,
	returnVal interface{}) {
	return func(ctx context.Context,
		contractName,
		methodName string,
		confidenceLevel primitives.ConfidenceLevel,
		params,
		returnVal interface{}) {
		mapper(returnVal)
	}
}

func TestCCIPChainReader_getDestFeeQuoterStaticConfig(t *testing.T) {
	ctx := context.Background()

	// Setup expected values
	offrampAddress := []byte{0x3}
	expectedConfig := feeQuoterStaticConfig{
		MaxFeeJuelsPerMsg:  cciptypes.NewBigIntFromInt64(10),
		LinkToken:          []byte{0x3, 0x4},
		StalenessThreshold: 12,
	}

	// Setup cache with the expected config
	mockCache := new(mockConfigCache)
	chainConfig := ChainConfigSnapshot{
		FeeQuoter: FeeQuoterConfig{
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
	destCR := reader_mocks.NewMockContractReaderFacade(t)
	destCR.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	destCR.EXPECT().HealthReport().Return(nil)
	destCR.EXPECT().GetLatestValue(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(func(
		ctx context.Context,
		readIdentifier string,
		confidenceLevel primitives.ConfidenceLevel,
		params interface{},
		returnVal interface{},
	) {
		givenTokenAddr := params.(map[string]any)["token"].([]byte)
		if bytes.Equal(tokenAddr, givenTokenAddr) {
			price := returnVal.(*plugintypes.TimestampedUnixBig)
			price.Value = big.NewInt(145)
		}
	}).Return(nil)

	offrampAddress := []byte{0x3}
	feeQuoterAddress := []byte{0x4}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)

	ccipReader := newCCIPChainReaderInternal(
		tests.Context(t),
		logger.Test(t),
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			chainC: destCR,
		}, nil, chainC, offrampAddress, mockAddrCodec,
	)

	// Add cleanup to properly shut down the background polling
	t.Cleanup(func() {
		err := ccipReader.Close()
		if err != nil {
			t.Logf("Error closing ccipReader: %v", err)
		}
	})

	feeQuoterAddressStr, err := mockAddrCodec.AddressBytesToString(feeQuoterAddress, 111_111)
	require.NoError(t, err)
	require.NoError(t, ccipReader.contractReaders[chainC].Bind(
		context.Background(), []types.BoundContract{{Name: "FeeQuoter",
			Address: feeQuoterAddressStr}}))

	ctx := context.Background()
	price, err := ccipReader.getFeeQuoterTokenPriceUSD(ctx, []byte{0x3, 0x4})
	assert.NoError(t, err)
	assert.Equal(t, cciptypes.NewBigIntFromInt64(145), price)
}

func TestCCIPFeeComponents_HappyPath(t *testing.T) {
	cw := writer_mocks.NewMockContractWriter(t)
	cw.EXPECT().GetFeeComponents(mock.Anything).Return(
		&types.ChainFeeComponents{
			ExecutionFee:        big.NewInt(1),
			DataAvailabilityFee: big.NewInt(2),
		}, nil,
	)

	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	// Missing writer for chainB
	contractWriters[chainA] = cw
	contractWriters[chainC] = cw

	ccipReader := newCCIPChainReaderInternal(
		tests.Context(t),
		logger.Test(t),
		nil,
		contractWriters,
		chainC,
		[]byte{0x3},
		internal.NewMockAddressCodecHex(t),
	)

	// Add cleanup to ensure resources are released
	t.Cleanup(func() {
		err := ccipReader.Close()
		if err != nil {
			t.Logf("Error closing ccipReader: %v", err)
		}
	})

	ctx := context.Background()
	feeComponents := ccipReader.GetChainsFeeComponents(ctx, []cciptypes.ChainSelector{chainA, chainB, chainC})
	assert.Len(t, feeComponents, 2)
	assert.Equal(t, big.NewInt(1), feeComponents[chainA].ExecutionFee)
	assert.Equal(t, big.NewInt(2), feeComponents[chainA].DataAvailabilityFee)
	assert.Equal(t, big.NewInt(1), feeComponents[chainC].ExecutionFee)
	assert.Equal(t, big.NewInt(2), feeComponents[chainC].DataAvailabilityFee)

	destChainFeeComponent, err := ccipReader.GetDestChainFeeComponents(ctx)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1), destChainFeeComponent.ExecutionFee)
	assert.Equal(t, big.NewInt(2), destChainFeeComponent.DataAvailabilityFee)
}

func TestCCIPFeeComponents_NotFoundErrors(t *testing.T) {
	cw := writer_mocks.NewMockContractWriter(t)
	contractWriters := make(map[cciptypes.ChainSelector]types.ContractWriter)
	// Missing writer for dest chain chainC
	contractWriters[chainA] = cw
	ccipReader := newCCIPChainReaderInternal(
		tests.Context(t),
		logger.Test(t),
		nil,
		contractWriters,
		chainC,
		[]byte{0x3},
		internal.NewMockAddressCodecHex(t),
	)

	// Add cleanup to ensure resources are released
	t.Cleanup(func() {
		err := ccipReader.Close()
		if err != nil {
			t.Logf("Error closing ccipReader: %v", err)
		}
	})

	ctx := context.Background()
	_, err := ccipReader.GetDestChainFeeComponents(ctx)
	require.Error(t, err)
}

func TestCCIPChainReader_LinkPriceUSD(t *testing.T) {
	ctx := context.Background()
	tokenAddr := []byte{0x3, 0x4}
	offrampAddress := []byte{0x3}

	// Setup mock cache with the fee quoter static config
	mockCache := new(mockConfigCache)
	chainConfig := ChainConfigSnapshot{
		FeeQuoter: FeeQuoterConfig{
			StaticConfig: feeQuoterStaticConfig{
				MaxFeeJuelsPerMsg:  cciptypes.NewBigIntFromInt64(10),
				LinkToken:          tokenAddr,
				StalenessThreshold: 12,
			},
		},
	}
	mockCache.On("GetChainConfig", mock.Anything, chainC).Return(chainConfig, nil)

	// Setup contract reader for getting token price
	destCR := reader_mocks.NewMockExtended(t)

	// mock the call to get the token price
	destCR.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameFeeQuoter,
		consts.MethodNameFeeQuoterGetTokenPrice,
		primitives.Unconfirmed,
		map[string]interface{}{"token": tokenAddr},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		price := returnVal.(*plugintypes.TimestampedUnixBig)
		price.Value = big.NewInt(145)
	}))

	// Setup ccipReader with both cache and contract readers
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	offrampAddressStr, err := mockAddrCodec.AddressBytesToString(offrampAddress, chainC)
	require.NoError(t, err)
	ccipReader := &ccipChainReader{
		lggr:           logger.Test(t),
		destChain:      chainC,
		configPoller:   mockCache,
		offrampAddress: offrampAddressStr,
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			chainC: destCR,
		},
	}

	price, err := ccipReader.LinkPriceUSD(ctx)
	require.NoError(t, err)
	assert.Equal(t, cciptypes.NewBigIntFromInt64(145), price)

	mockCache.AssertExpectations(t)
}

func TestCCIPChainReader_GetMedianDataAvailabilityGasConfig(t *testing.T) {
	type mockValue struct {
		overhead   uint32
		perByte    uint16
		multiplier uint16
		enabled    bool
	}

	setupConfigMocks := func(
		readers map[cciptypes.ChainSelector]*reader_mocks.MockExtended,
		chains []cciptypes.ChainSelector,
		values []mockValue) {
		for i, chain := range chains {
			readers[chain].EXPECT().
				ExtendedGetLatestValue(
					mock.Anything,
					consts.ContractNameFeeQuoter,
					consts.MethodNameGetDestChainConfig,
					primitives.Unconfirmed,
					mock.Anything,
					mock.Anything,
				).
				Return(nil).
				Run(withReturnValueOverridden(func(returnVal interface{}) {
					cfg := returnVal.(*cciptypes.FeeQuoterDestChainConfig)
					cfg.DestDataAvailabilityOverheadGas = values[i].overhead
					cfg.DestGasPerDataAvailabilityByte = values[i].perByte
					cfg.DestDataAvailabilityMultiplierBps = values[i].multiplier
					cfg.IsEnabled = values[i].enabled
				})).Once()
		}
	}

	tests := []struct {
		name           string
		expectedConfig cciptypes.DataAvailabilityGasConfig
		expectError    bool
		chains         []cciptypes.ChainSelector
		setupMocks     func(readers map[cciptypes.ChainSelector]*reader_mocks.MockExtended)
	}{
		{
			name: "success - returns median values from multiple configs",
			expectedConfig: cciptypes.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   200,
				DestGasPerDataAvailabilityByte:    20,
				DestDataAvailabilityMultiplierBps: 2000,
			},
			chains: []cciptypes.ChainSelector{chainA, chainB, chainC},
			setupMocks: func(readers map[cciptypes.ChainSelector]*reader_mocks.MockExtended) {
				values := []mockValue{
					{100, 10, 1000, true},
					{200, 20, 2000, true},
					{300, 30, 3000, true},
				}
				setupConfigMocks(readers, []cciptypes.ChainSelector{chainA, chainB, chainC}, values)
			},
		},
		{
			name: "success - skips disabled configs",
			expectedConfig: cciptypes.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   300,
				DestGasPerDataAvailabilityByte:    30,
				DestDataAvailabilityMultiplierBps: 3000,
			},
			chains: []cciptypes.ChainSelector{chainA, chainB, chainC},
			setupMocks: func(readers map[cciptypes.ChainSelector]*reader_mocks.MockExtended) {
				values := []mockValue{
					{100, 10, 1000, true},
					{200, 20, 2000, false},
					{300, 30, 3000, true},
				}
				setupConfigMocks(readers, []cciptypes.ChainSelector{chainA, chainB, chainC}, values)
			},
		},
		{
			name: "no valid configs found due to empty DA params",
			expectedConfig: cciptypes.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   0,
				DestGasPerDataAvailabilityByte:    0,
				DestDataAvailabilityMultiplierBps: 0,
			},
			expectError: false,
			chains:      []cciptypes.ChainSelector{chainA, chainB, chainC},
			setupMocks: func(readers map[cciptypes.ChainSelector]*reader_mocks.MockExtended) {
				values := []mockValue{
					{0, 0, 0, true}, // Empty DA params
					{0, 0, 0, true}, // Empty DA params
					{0, 0, 0, true}, // Empty DA params
				}
				setupConfigMocks(readers, []cciptypes.ChainSelector{chainA, chainB, chainC}, values)
			},
		},
		{
			name:        "all configs disabled",
			expectError: false,
			expectedConfig: cciptypes.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   0,
				DestGasPerDataAvailabilityByte:    0,
				DestDataAvailabilityMultiplierBps: 0,
			},
			chains: []cciptypes.ChainSelector{chainA, chainB},
			setupMocks: func(readers map[cciptypes.ChainSelector]*reader_mocks.MockExtended) {
				values := []mockValue{
					{100, 10, 1000, false},
					{200, 20, 2000, false},
				}
				setupConfigMocks(readers, []cciptypes.ChainSelector{chainA, chainB}, values)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReaders := make(map[cciptypes.ChainSelector]*reader_mocks.MockExtended)
			contractReaders := make(map[cciptypes.ChainSelector]contractreader.Extended)

			// Initialize mocks
			for _, chain := range tt.chains {
				mockReaders[chain] = reader_mocks.NewMockExtended(t)
				contractReaders[chain] = mockReaders[chain]
			}

			tt.setupMocks(mockReaders)

			reader := &ccipChainReader{
				lggr:            logger.Test(t),
				contractReaders: contractReaders,
				destChain:       chainC,
			}
			config, err := reader.GetMedianDataAvailabilityGasConfig(context.Background())

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedConfig, config)
			}
		})
	}
}

func Test_getCurseInfoFromCursedSubjects(t *testing.T) {
	testCases := []struct {
		name              string
		cursedSubjectsSet mapset.Set[[16]byte]
		destChainSelector cciptypes.ChainSelector
		expCurseInfo      CurseInfo
	}{
		{
			name:              "no cursed subjects",
			cursedSubjectsSet: mapset.NewSet[[16]byte](),
			destChainSelector: chainA,
			expCurseInfo: CurseInfo{
				CursedSourceChains: map[cciptypes.ChainSelector]bool{},
				CursedDestination:  false,
				GlobalCurse:        false,
			},
		},
		{
			name: "everything cursed",
			cursedSubjectsSet: mapset.NewSet[[16]byte](
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
				chainSelectorToBytes16(chainA), // dest
				GlobalCurseSubject,
			),
			destChainSelector: chainA,
			expCurseInfo: CurseInfo{
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
			cursedSubjectsSet: mapset.NewSet[[16]byte](
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
				chainSelectorToBytes16(chainA), // dest
			),
			destChainSelector: chainA,
			expCurseInfo: CurseInfo{
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
			cursedSubjectsSet: mapset.NewSet[[16]byte](
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
				GlobalCurseSubject,
			),
			destChainSelector: chainA,
			expCurseInfo: CurseInfo{
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
			cursedSubjectsSet: mapset.NewSet[[16]byte](
				chainSelectorToBytes16(chainB),
				chainSelectorToBytes16(chainC),
			),
			destChainSelector: chainA,
			expCurseInfo: CurseInfo{
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
			cursedSubjectsSet: mapset.NewSet[[16]byte](
				chainSelectorToBytes16(chainC),
				chainSelectorToBytes16(chainA), // dest
				GlobalCurseSubject,
			),
			destChainSelector: chainA,
			expCurseInfo: CurseInfo{
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
	addr1 := "0x1234567890123456789012345678901234567890"
	addr2 := "0x2234567890123456789012345678901234567890"
	nonceManagerAddr := "0x3234567890123456789012345678901234567890"

	destReader := reader_mocks.NewMockExtended(t)

	// Mock Bind call
	destReader.EXPECT().Bind(
		mock.Anything,
		[]types.BoundContract{{
			Name:    consts.ContractNameNonceManager,
			Address: nonceManagerAddr,
		}},
	).Return(nil)

	// Setup expected responses
	nonce1, nonce2 := uint64(5), uint64(10)

	result1 := &types.BatchReadResult{ReadName: consts.MethodNameGetInboundNonce}
	result1.SetResult(&nonce1, nil)

	result2 := &types.BatchReadResult{ReadName: consts.MethodNameGetInboundNonce}
	result2.SetResult(&nonce2, nil)

	responses := types.BatchGetLatestValuesResult{
		types.BoundContract{
			Name:    consts.ContractNameNonceManager,
			Address: nonceManagerAddr,
		}: []types.BatchReadResult{*result1, *result2},
	}
	destReader.EXPECT().ExtendedBatchGetLatestValues(
		mock.Anything,
		mock.MatchedBy(func(req contractreader.ExtendedBatchGetLatestValuesRequest) bool {
			batch := req[consts.ContractNameNonceManager]
			return len(batch) == 2 &&
				batch[0].ReadName == consts.MethodNameGetInboundNonce &&
				batch[1].ReadName == consts.MethodNameGetInboundNonce
		}),
		false,
	).Return(responses, []string{}, nil)
	ccipReader := &ccipChainReader{
		lggr: logger.Test(t),
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			chainB: destReader,
		},
		destChain: chainB,
		addrCodec: internal.NewMockAddressCodecHex(t),
	}

	// Bind the contract first
	err := ccipReader.contractReaders[chainB].Bind(
		context.Background(),
		[]types.BoundContract{{
			Name:    consts.ContractNameNonceManager,
			Address: nonceManagerAddr,
		}},
	)
	require.NoError(t, err)

	// Call Nonces
	nonces, err := ccipReader.Nonces(
		context.Background(),
		chainA,
		[]string{addr1, addr2},
	)

	require.NoError(t, err)
	assert.Equal(t, map[string]uint64{
		addr1: 5,
		addr2: 10,
	}, nonces)
}

func TestCCIPChainReader_DiscoverContracts_Parallel(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChains := []cciptypes.ChainSelector{2, 3, 4} // Adding one more chain for better parallelism testing

	// Setup mock cache and configurations
	mockCache := new(mockConfigCache)

	// Destination chain config with a delay
	destChainConfig := ChainConfigSnapshot{
		Offramp: OfframpConfig{
			StaticConfig: offRampStaticChainConfig{
				NonceManager: []byte{0x3},
				RmnRemote:    []byte{0x4},
			},
			DynamicConfig: offRampDynamicChainConfig{
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
		srcChainConfig := ChainConfigSnapshot{
			OnRamp: OnRampConfig{
				DynamicConfig: getOnRampDynamicConfigResponse{
					DynamicConfig: onRampDynamicConfig{
						FeeQuoter: []byte{0x7},
					},
				},
				DestChainConfig: onRampDestChainConfig{
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
	mockReaders := make(map[cciptypes.ChainSelector]*reader_mocks.MockExtended)
	contractReaders := make(map[cciptypes.ChainSelector]contractreader.Extended)

	// Setup dest chain reader
	mockReaders[destChain] = reader_mocks.NewMockExtended(t)
	contractReaders[destChain] = mockReaders[destChain]

	// Setup source chain readers
	for _, chain := range sourceChains {
		mockReaders[chain] = reader_mocks.NewMockExtended(t)
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
	contractAddresses, err := ccipReader.DiscoverContracts(ctx, sourceChains)
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
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)

	wrappedNative1 := cciptypes.Bytes{0x1}
	wrappedNative2 := cciptypes.Bytes{0x2}

	t.Run("happy path - gets prices for all chains", func(t *testing.T) {
		// Setup mock cache with configs containing wrapped native addresses
		mockCache := new(mockConfigCache)
		sourceChain1Config := ChainConfigSnapshot{
			Router: RouterConfig{
				WrappedNativeAddress: wrappedNative1,
			},
		}
		sourceChain2Config := ChainConfigSnapshot{
			Router: RouterConfig{
				WrappedNativeAddress: wrappedNative2,
			},
		}

		mockCache.On("GetChainConfig", mock.Anything, sourceChain1).Return(sourceChain1Config, nil)
		mockCache.On("GetChainConfig", mock.Anything, sourceChain2).Return(sourceChain2Config, nil)

		// Setup readers with price responses
		sourceReader1 := reader_mocks.NewMockExtended(t)
		price1 := plugintypes.TimestampedUnixBig{
			Value:     big.NewInt(100),
			Timestamp: uint32(time.Now().Unix()),
		}
		sourceReader1.EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameFeeQuoter,
			consts.MethodNameFeeQuoterGetTokenPrice,
			primitives.Unconfirmed,
			map[string]interface{}{"token": wrappedNative1},
			mock.Anything,
		).Run(
			func(
				ctx context.Context,
				contractName, methodName string,
				confidence primitives.ConfidenceLevel,
				params any,
				returnVal any) {
				pricePtr := returnVal.(*plugintypes.TimestampedUnixBig)
				*pricePtr = price1
			}).Return(nil)

		sourceReader2 := reader_mocks.NewMockExtended(t)
		price2 := plugintypes.TimestampedUnixBig{
			Value:     big.NewInt(200),
			Timestamp: uint32(time.Now().Unix()),
		}
		sourceReader2.EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameFeeQuoter,
			consts.MethodNameFeeQuoterGetTokenPrice,
			primitives.Unconfirmed,
			map[string]interface{}{"token": wrappedNative2},
			mock.Anything,
		).Run(
			func(
				ctx context.Context,
				contractName, methodName string,
				confidence primitives.ConfidenceLevel,
				params any,
				returnVal any) {
				pricePtr := returnVal.(*plugintypes.TimestampedUnixBig)
				*pricePtr = price2
			}).Return(nil)

		ccipReader := &ccipChainReader{
			destChain: destChain,
			contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
				sourceChain1: sourceReader1,
				sourceChain2: sourceReader2,
			},
			configPoller: mockCache,
			lggr:         logger.Test(t),
		}

		prices := ccipReader.GetWrappedNativeTokenPriceUSD(ctx, []cciptypes.ChainSelector{sourceChain1, sourceChain2})
		require.Len(t, prices, 2)
		assert.Equal(t, cciptypes.NewBigInt(big.NewInt(100)), prices[sourceChain1])
		assert.Equal(t, cciptypes.NewBigInt(big.NewInt(200)), prices[sourceChain2])

		mockCache.AssertExpectations(t)
	})

	t.Run("handles missing chain configs", func(t *testing.T) {
		mockCache := new(mockConfigCache)
		mockCache.On("GetChainConfig", mock.Anything, sourceChain1).Return(ChainConfigSnapshot{}, fmt.Errorf("not found"))
		mockCache.On("GetChainConfig", mock.Anything, sourceChain2).Return(ChainConfigSnapshot{
			Router: RouterConfig{
				WrappedNativeAddress: wrappedNative2,
			},
		}, nil)

		sourceReader2 := reader_mocks.NewMockExtended(t)
		price2 := plugintypes.TimestampedUnixBig{
			Value:     big.NewInt(200),
			Timestamp: uint32(time.Now().Unix()),
		}
		sourceReader2.EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameFeeQuoter,
			consts.MethodNameFeeQuoterGetTokenPrice,
			primitives.Unconfirmed,
			map[string]interface{}{"token": wrappedNative2},
			mock.Anything,
		).Run(func(
			ctx context.Context,
			contractName, methodName string,
			confidence primitives.ConfidenceLevel,
			params any,
			returnVal any) {
			pricePtr := returnVal.(*plugintypes.TimestampedUnixBig)
			*pricePtr = price2
		}).Return(nil)

		ccipReader := &ccipChainReader{
			destChain: destChain,
			contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
				sourceChain1: reader_mocks.NewMockExtended(t),
				sourceChain2: sourceReader2,
			},
			configPoller: mockCache,
			lggr:         logger.Test(t),
		}

		prices := ccipReader.GetWrappedNativeTokenPriceUSD(ctx, []cciptypes.ChainSelector{sourceChain1, sourceChain2})
		require.Len(t, prices, 1)
		assert.Equal(t, cciptypes.NewBigInt(big.NewInt(200)), prices[sourceChain2])

		mockCache.AssertExpectations(t)
	})

	t.Run("handles price fetch error", func(t *testing.T) {
		mockCache := new(mockConfigCache)
		sourceConfig := ChainConfigSnapshot{
			Router: RouterConfig{
				WrappedNativeAddress: wrappedNative1,
			},
		}
		mockCache.On("GetChainConfig", mock.Anything, sourceChain1).Return(sourceConfig, nil)

		sourceReader := reader_mocks.NewMockExtended(t)
		sourceReader.EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameFeeQuoter,
			consts.MethodNameFeeQuoterGetTokenPrice,
			primitives.Unconfirmed,
			map[string]interface{}{"token": wrappedNative1},
			mock.Anything,
		).Return(fmt.Errorf("price fetch failed"))

		ccipReader := &ccipChainReader{
			destChain: destChain,
			contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
				sourceChain1: sourceReader,
			},
			configPoller: mockCache,
			lggr:         logger.Test(t),
		}

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
		require.IsType(t, &getOnRampDynamicConfigResponse{}, onRampRequests[0].ReturnVal)

		// Verify OnRamp dest chain config request
		require.Equal(t, consts.MethodNameOnRampGetDestChainConfig, onRampRequests[1].ReadName)
		require.Equal(t, map[string]any{"destChainSelector": destChain}, onRampRequests[1].Params)
		require.IsType(t, &onRampDestChainConfig{}, onRampRequests[1].ReturnVal)

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
		require.IsType(t, &OCRConfigResponse{}, offRampRequests[0].ReturnVal)

		// Check OffRamp execute config request
		require.Equal(t, consts.MethodNameOffRampLatestConfigDetails, offRampRequests[1].ReadName)
		require.Equal(t, map[string]any{"ocrPluginType": consts.PluginTypeExecute}, offRampRequests[1].Params)
		require.IsType(t, &OCRConfigResponse{}, offRampRequests[1].ReturnVal)

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
		require.IsType(t, &feeQuoterStaticConfig{}, feeQuoterRequests[0].ReturnVal)
	})
}

type mockConfigCache struct {
	mock.Mock
}

func (m *mockConfigCache) GetChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error) {
	args := m.Called(ctx, chainSel)
	return args.Get(0).(ChainConfigSnapshot), args.Error(1)
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
