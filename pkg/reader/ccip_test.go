package reader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	chainA = cciptypes.ChainSelector(1)
	chainB = cciptypes.ChainSelector(2)
	chainC = cciptypes.ChainSelector(3)
)

func TestCCIPChainReader_getSourceChainsConfig(t *testing.T) {
	sourceCRs := make(map[cciptypes.ChainSelector]*reader_mocks.MockContractReaderFacade)
	for _, chain := range []cciptypes.ChainSelector{chainA, chainB} {
		sourceCRs[chain] = reader_mocks.NewMockContractReaderFacade(t)
		sourceCRs[chain].EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	}

	destCR := reader_mocks.NewMockContractReaderFacade(t)
	destCR.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
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
		sourceChain := params.(map[string]any)["sourceChainSelector"].(cciptypes.ChainSelector)
		v := returnVal.(*sourceChainConfig)
		v.OnRamp = []byte(fmt.Sprintf("onramp-%d", sourceChain))
		v.IsEnabled = true
	}).Return(nil)

	offrampAddress := []byte{0x3}
	ccipReader := newCCIPChainReaderInternal(
		tests.Context(t),
		logger.Test(t),
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			chainA: sourceCRs[chainA],
			chainB: sourceCRs[chainB],
			chainC: destCR,
		}, nil, chainC, offrampAddress,
	)

	require.NoError(t, ccipReader.contractReaders[chainA].Bind(
		context.Background(), []types.BoundContract{{Name: "OnRamp", Address: "0x1"}}))
	require.NoError(t, ccipReader.contractReaders[chainB].Bind(
		context.Background(), []types.BoundContract{{Name: "OnRamp", Address: "0x2"}}))
	require.NoError(t, ccipReader.contractReaders[chainC].Bind(
		context.Background(), []types.BoundContract{{Name: "OffRamp",
			Address: typeconv.AddressBytesToString(offrampAddress, 111_111)}}))

	ctx := context.Background()
	cfgs, err := ccipReader.getOffRampSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB})
	assert.NoError(t, err)
	assert.Len(t, cfgs, 2)
	assert.Equal(t, []byte("onramp-1"), cfgs[chainA].OnRamp)
	assert.Equal(t, []byte("onramp-2"), cfgs[chainB].OnRamp)
}

func TestCCIPChainReader_GetContractAddress(t *testing.T) {
	ecr := reader_mocks.NewMockExtended(t)

	ccipReader := ccipChainReader{
		lggr: logger.Test(t),
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			chainA: ecr,
		},
	}

	someAddr := "0x1234567890123456789012345678901234567890"
	someAddrBytes, err := typeconv.AddressStringToBytes(someAddr, uint64(chainA))
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
		assert.Equal(t, someAddrBytes, addr)
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
	destExtended := reader_mocks.NewMockExtended(t)
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: typeconv.AddressBytesToString(destNonceMgr, uint64(destChain)),
		},
	}).Return(nil)

	source1Extended := reader_mocks.NewMockExtended(t)
	source1Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: typeconv.AddressBytesToString(s1Onramp, uint64(sourceChain1)),
		},
	}).Return(nil)

	source2Extended := reader_mocks.NewMockExtended(t)
	source2Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: typeconv.AddressBytesToString(s2Onramp, uint64(sourceChain2)),
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

	err := ccipReader.Sync(ctx, contracts)
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
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: typeconv.AddressBytesToString(destNonceMgr, uint64(destChain)),
		},
	}).Return(nil)

	source1Extended := reader_mocks.NewMockExtended(t)
	source1Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: typeconv.AddressBytesToString(s1Onramp, uint64(sourceChain1)),
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

	err := ccipReader.Sync(ctx, contracts)
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
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: typeconv.AddressBytesToString(destNonceMgr, uint64(destChain)),
		},
	}).Return(nil)

	// only support source2, source1 unsupported.
	source2Extended := reader_mocks.NewMockExtended(t)
	source2Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: typeconv.AddressBytesToString(s2Onramp, uint64(sourceChain2)),
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

	err := ccipReader.Sync(ctx, contracts)
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
	destExtended := reader_mocks.NewMockExtended(t)
	destExtended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameNonceManager,
			Address: typeconv.AddressBytesToString(destNonceMgr, uint64(destChain)),
		},
	}).Return(nil)

	expectedErr := errors.New("some error")
	source1Extended := reader_mocks.NewMockExtended(t)
	source1Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: typeconv.AddressBytesToString(s1Onramp, uint64(sourceChain1)),
		},
	}).Return(expectedErr)

	source2Extended := reader_mocks.NewMockExtended(t)
	source2Extended.EXPECT().Bind(mock.Anything, []types.BoundContract{
		{
			Name:    consts.ContractNameOnRamp,
			Address: typeconv.AddressBytesToString(s2Onramp, uint64(sourceChain2)),
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

	err := ccipReader.Sync(ctx, contracts)
	require.Error(t, err)
	require.ErrorIs(t, err, expectedErr)
}

func addDestinationContractAssertions(
	extended *reader_mocks.MockExtended,
	destNonceMgr, destRMNRemote, destFeeQuoter []byte,
) {
	// mock the call to get the nonce manager
	extended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetStaticConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*offRampStaticChainConfig)
		v.NonceManager = destNonceMgr
		v.RmnRemote = destRMNRemote
	}))
	// mock the call to get the fee quoter
	extended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetDynamicConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*offRampDynamicChainConfig)
		v.FeeQuoter = destFeeQuoter
	}))
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
	//srcRouters := []byte{0x7, 0x8}
	//srcFeeQuoters := [2][]byte{{0x7}, {0x8}}

	// Build expected addresses.
	var expectedContractAddresses ContractAddresses
	// Source FeeQuoter's and destRouter are missing.
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
	addDestinationContractAssertions(mockReaders[destChain], destNonceMgr, destRMNRemote, destFeeQuoter)

	mockReaders[destChain].EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetAllSourceChainConfigs,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*selectorsAndConfigs)
		v.Selectors = []uint64{uint64(sourceChain[0]), uint64(sourceChain[1])}
		v.SourceChainConfigs = []sourceChainConfig{
			{
				OnRamp:    onramps[0],
				Router:    destRouter,
				IsEnabled: true,
			},
			{
				OnRamp:    onramps[1],
				Router:    destRouter,
				IsEnabled: true,
			},
		}
	}))

	// mock calls to get fee quoter from onramps and source chain config from offramp.
	for _, selector := range sourceChain {
		mockReaders[selector] = reader_mocks.NewMockExtended(t)

		// ErrNoBindings is ignored.
		mockReaders[selector].EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameOnRamp,
			consts.MethodNameOnRampGetDynamicConfig,
			primitives.Unconfirmed,
			map[string]any{},
			mock.Anything,
		).Return(contractreader.ErrNoBindings)

		mockReaders[selector].EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameOnRamp,
			consts.MethodNameOnRampGetDestChainConfig,
			primitives.Unconfirmed,
			map[string]any{
				"destChainSelector": selector,
			},
			mock.Anything,
		).Return(contractreader.ErrNoBindings)
	}

	castToExtended := make(map[cciptypes.ChainSelector]contractreader.Extended)
	for sel, v := range mockReaders {
		castToExtended[sel] = v
	}

	lggr, hook := logger.TestObserved(t, zapcore.InfoLevel)
	// create the reader
	ccipChainReader := &ccipChainReader{
		destChain:       destChain,
		contractReaders: castToExtended,
		lggr:            lggr,
	}

	// TODO: allChains should be initialized to "append(onramps, destChain)" when that feature is implemented.
	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx)
	require.NoError(t, err)

	assert.Equal(t, expectedContractAddresses, contractAddresses)
	require.Equal(t, 3, hook.Len())

	assert.Contains(
		t,
		"appending RMN remote contract address",
		hook.All()[0].Message,
	)
	assert.Contains(
		t,
		"unable to lookup source fee quoters, this is expected during initialization",
		hook.All()[1].Message,
	)
	assert.Contains(
		t,
		"unable to lookup source routers, this is expected during initialization",
		hook.All()[2].Message,
	)
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
	addDestinationContractAssertions(mockReaders[destChain], destNonceMgr, destRMNRemote, destFeeQuoter)

	mockReaders[destChain].EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetAllSourceChainConfigs,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*selectorsAndConfigs)
		v.Selectors = []uint64{uint64(sourceChain[0]), uint64(sourceChain[1])}
		v.SourceChainConfigs = []sourceChainConfig{
			{
				OnRamp:    onramps[0],
				Router:    destRouter[0],
				IsEnabled: true,
			},
			{
				OnRamp:    onramps[1],
				Router:    destRouter[1],
				IsEnabled: true,
			},
		}
	}))

	// mock calls to get fee quoter from onramps and source chain config from offramp.
	for i, selector := range sourceChain {
		mockReaders[selector] = reader_mocks.NewMockExtended(t)

		mockReaders[selector].EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameOnRamp,
			consts.MethodNameOnRampGetDynamicConfig,
			primitives.Unconfirmed,
			map[string]any{},
			mock.Anything,
		).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
			v := returnVal.(*onRampDynamicChainConfig)
			v.FeeQuoter = srcFeeQuoters[i]
		}))

		mockReaders[selector].EXPECT().ExtendedGetLatestValue(
			mock.Anything,
			consts.ContractNameOnRamp,
			consts.MethodNameOnRampGetDestChainConfig,
			primitives.Unconfirmed,
			map[string]any{
				"destChainSelector": selector,
			},
			mock.Anything,
		).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
			v := returnVal.(*onRampDestChainConfig)
			v.Router = srcRouters[i]
		}))
	}

	castToExtended := make(map[cciptypes.ChainSelector]contractreader.Extended)
	for sel, v := range mockReaders {
		castToExtended[sel] = v
	}

	// create the reader
	ccipChainReader := &ccipChainReader{
		destChain:       destChain,
		contractReaders: castToExtended,
		lggr:            logger.Test(t),
	}

	// TODO: allChains should be initialized to "append(onramps, destChain)" when that feature is implemented.
	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx)
	require.NoError(t, err)

	require.Equal(t, expectedContractAddresses, contractAddresses)
}

// TODO: Remove this test when allChains is implemented.
// This test checks that onramps are not discovered if there are only dest readers available.
func TestCCIPChainReader_DiscoverContracts_HappyPath_OnlySupportDest(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain := [2]cciptypes.ChainSelector{2, 3}
	onramps := [2][]byte{{0x1}, {0x2}}
	destNonceMgr := []byte{0x3}
	destRMNRemote := []byte{0x4}
	destFeeQuoter := []byte{0x5}
	destRouter := []byte{0x6}

	var expectedContractAddresses ContractAddresses
	for i := range onramps {
		expectedContractAddresses = expectedContractAddresses.Append(
			consts.ContractNameOnRamp, sourceChain[i], onramps[i])
	}
	// All dest chain contracts should be discovered, they do not require source chain support.
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameRouter, destChain, destRouter)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameFeeQuoter, destChain, destFeeQuoter)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameRMNRemote, destChain, destRMNRemote)
	expectedContractAddresses = expectedContractAddresses.Append(consts.ContractNameNonceManager, destChain, destNonceMgr)

	destExtended := reader_mocks.NewMockExtended(t)

	// mock the call to get the nonce manager
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetStaticConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*offRampStaticChainConfig)
		v.NonceManager = destNonceMgr
		v.RmnRemote = destRMNRemote
	}))
	// mock the call to get the fee quoter
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetDynamicConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*offRampDynamicChainConfig)
		v.FeeQuoter = destFeeQuoter
	}))
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetAllSourceChainConfigs,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*selectorsAndConfigs)
		v.Selectors = []uint64{uint64(sourceChain[0]), uint64(sourceChain[1])}
		v.SourceChainConfigs = []sourceChainConfig{
			{
				OnRamp:    onramps[0],
				Router:    destRouter,
				IsEnabled: true,
			},
			{
				OnRamp:    onramps[1],
				Router:    destRouter,
				IsEnabled: true,
			},
		}
	}))

	// create the reader
	ccipChainReader := &ccipChainReader{
		destChain: destChain,
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain: destExtended,
		},
		lggr: logger.Test(t),
	}

	// TODO: allChains should be initialized to "append(onramps, destChain)" when that feature is implemented.
	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx)
	require.NoError(t, err)
	require.Equal(t, expectedContractAddresses, contractAddresses)
}

func TestCCIPChainReader_DiscoverContracts_GetAllSourceChainConfig_Errors(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	destExtended := reader_mocks.NewMockExtended(t)

	// mock the call for sourceChain2 - failure
	getLatestValueErr := errors.New("some error")
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetAllSourceChainConfigs,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(getLatestValueErr)

	// get static config call won't occur because the source chain config call failed.

	// create the reader
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
		lggr: logger.Test(t),
	}

	// TODO: allChains should be initialized to "append(onramps, destChain)" when that feature is implemented.
	_, err := ccipChainReader.DiscoverContracts(ctx)
	require.Error(t, err)
	require.ErrorIs(t, err, getLatestValueErr)
}

func TestCCIPChainReader_DiscoverContracts_GetOfframpStaticConfig_Errors(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	destExtended := reader_mocks.NewMockExtended(t)

	// mock the call for source chain configs
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetAllSourceChainConfigs,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil) // doesn't matter for this test
	// mock the call to get the nonce manager - failure
	getLatestValueErr := errors.New("some error")
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetStaticConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(getLatestValueErr)

	// create the reader
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
		lggr: logger.Test(t),
	}

	// TODO: allChains should be initialized to "append(onramps, destChain)" when that feature is implemented.
	_, err := ccipChainReader.DiscoverContracts(ctx)
	require.Error(t, err)
	require.ErrorIs(t, err, getLatestValueErr)
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
	destCR := reader_mocks.NewMockContractReaderFacade(t)
	destCR.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
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
		cfg := returnVal.(*feeQuoterStaticConfig)
		cfg.MaxFeeJuelsPerMsg = cciptypes.NewBigIntFromInt64(10)
		cfg.LinkToken = []byte{0x3, 0x4}
		cfg.StalenessThreshold = 12
	}).Return(nil)

	offrampAddress := []byte{0x3}
	feeQuoterAddress := []byte{0x4}
	ccipReader := newCCIPChainReaderInternal(
		tests.Context(t),
		logger.Test(t),
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			chainC: destCR,
		}, nil, chainC, offrampAddress,
	)

	require.NoError(t, ccipReader.contractReaders[chainC].Bind(
		context.Background(), []types.BoundContract{{Name: "FeeQuoter",
			Address: typeconv.AddressBytesToString(feeQuoterAddress, 111_111)}}))

	ctx := context.Background()
	cfg, err := ccipReader.getDestFeeQuoterStaticConfig(ctx)
	assert.NoError(t, err)
	assert.Equal(t, cciptypes.NewBigIntFromInt64(10), cfg.MaxFeeJuelsPerMsg)
	assert.Equal(t, []byte{0x3, 0x4}, cfg.LinkToken)
	assert.Equal(t, uint32(12), cfg.StalenessThreshold)
}

func TestCCIPChainReader_getFeeQuoterTokenPriceUSD(t *testing.T) {
	tokenAddr := []byte{0x3, 0x4}
	destCR := reader_mocks.NewMockContractReaderFacade(t)
	destCR.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
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
	ccipReader := newCCIPChainReaderInternal(
		tests.Context(t),
		logger.Test(t),
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			chainC: destCR,
		}, nil, chainC, offrampAddress,
	)

	require.NoError(t, ccipReader.contractReaders[chainC].Bind(
		context.Background(), []types.BoundContract{{Name: "FeeQuoter",
			Address: typeconv.AddressBytesToString(feeQuoterAddress, 111_111)}}))

	ctx := context.Background()
	price, err := ccipReader.getFeeQuoterTokenPriceUSD(ctx, []byte{0x3, 0x4})
	assert.NoError(t, err)
	assert.Equal(t, cciptypes.NewBigIntFromInt64(145), price)
}

func TestCCIPChainReader_LinkPriceUSD(t *testing.T) {
	tokenAddr := []byte{0x3, 0x4}
	destCR := reader_mocks.NewMockExtended(t)
	destCR.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)

	destCR.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameFeeQuoter,
		consts.MethodNameFeeQuoterGetStaticConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		cfg := returnVal.(*feeQuoterStaticConfig)
		cfg.MaxFeeJuelsPerMsg = cciptypes.NewBigIntFromInt64(10)
		cfg.LinkToken = []byte{0x3, 0x4}
		cfg.StalenessThreshold = 12
	}))

	// mock the call to get the fee quoter
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

	offrampAddress := []byte{0x3}
	feeQuoterAddress := []byte{0x4}
	contractReaders := make(map[cciptypes.ChainSelector]contractreader.Extended)
	contractReaders[chainC] = destCR
	ccipReader := ccipChainReader{
		logger.Test(t),
		contractReaders,
		nil,
		chainC,
		string(offrampAddress),
	}

	require.NoError(t, ccipReader.contractReaders[chainC].Bind(
		context.Background(), []types.BoundContract{{Name: "FeeQuoter",
			Address: typeconv.AddressBytesToString(feeQuoterAddress, 111_111)}}))

	ctx := context.Background()
	price, err := ccipReader.LinkPriceUSD(ctx)
	assert.NoError(t, err)
	assert.Equal(t, cciptypes.NewBigIntFromInt64(145), price)
}
