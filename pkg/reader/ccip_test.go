package reader

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
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
	cfgs, err := ccipReader.getSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB})
	assert.NoError(t, err)
	assert.Len(t, cfgs, 2)
	assert.Equal(t, []byte("onramp-1"), cfgs[chainA].OnRamp)
	assert.Equal(t, []byte("onramp-2"), cfgs[chainB].OnRamp)
}

func TestCCIPChainReader_GetContractAddress(t *testing.T) {
	ecr := reader_mocks.NewMockExtended(t)

	ccipReader := ccipChainReader{
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

func TestCCIPChainReader_DiscoverContracts_HappyPath(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destNonceMgr := []byte{0x3}
	expectedContractAddresses := ContractAddresses{
		consts.ContractNameOnRamp: {
			sourceChain1: s1Onramp,
			sourceChain2: s2Onramp,
		},
		consts.ContractNameNonceManager: {
			destChain: destNonceMgr,
		},
	}
	destExtended := reader_mocks.NewMockExtended(t)

	// mock the call for sourceChain1
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		primitives.Unconfirmed,
		map[string]any{"sourceChainSelector": sourceChain1},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*sourceChainConfig)
		v.OnRamp = s1Onramp
		v.IsEnabled = true
	}))
	// mock the call for sourceChain2
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		primitives.Unconfirmed,
		map[string]any{"sourceChainSelector": sourceChain2},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*sourceChainConfig)
		v.OnRamp = s2Onramp
		v.IsEnabled = true
	}))
	// mock the call to get the nonce manager
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOfframpGetStaticConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*offrampStaticChainConfig)
		v.NonceManager = destNonceMgr
	}))

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

	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx, destChain)
	require.NoError(t, err)

	require.Equal(t, expectedContractAddresses, contractAddresses)
}

func TestCCIPChainReader_DiscoverContracts_HappyPath_OnlySupportDest(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	destNonceMgr := []byte{0x3}
	expectedContractAddresses := ContractAddresses{
		// since the source chains are not supported, we should not have any onramp addresses
		// after discovery.
		consts.ContractNameOnRamp: {},
		// the nonce manager doesn't require source chain support though,
		// so we should discover that always if we support the dest.
		consts.ContractNameNonceManager: {
			destChain: destNonceMgr,
		},
	}
	destExtended := reader_mocks.NewMockExtended(t)

	// mock the call to get the nonce manager
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOfframpGetStaticConfig,
		primitives.Unconfirmed,
		map[string]any{},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*offrampStaticChainConfig)
		v.NonceManager = destNonceMgr
	}))

	// create the reader
	ccipChainReader := &ccipChainReader{
		destChain: destChain,
		contractReaders: map[cciptypes.ChainSelector]contractreader.Extended{
			destChain: destExtended,
		},
		lggr: logger.Test(t),
	}

	contractAddresses, err := ccipChainReader.DiscoverContracts(ctx, destChain)
	require.NoError(t, err)
	require.Equal(t, expectedContractAddresses, contractAddresses)
}

func TestCCIPChainReader_DiscoverContracts_GetSourceChainConfig_Errors(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	destExtended := reader_mocks.NewMockExtended(t)

	// mock the call for sourceChain1 - successful
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		primitives.Unconfirmed,
		map[string]any{"sourceChainSelector": sourceChain1},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*sourceChainConfig)
		v.OnRamp = s1Onramp
		v.IsEnabled = true
	}))

	// mock the call for sourceChain2 - failure
	getLatestValueErr := errors.New("some error")
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		primitives.Unconfirmed,
		map[string]any{"sourceChainSelector": sourceChain2},
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

	_, err := ccipChainReader.DiscoverContracts(ctx, destChain)
	require.Error(t, err)
	require.ErrorIs(t, err, getLatestValueErr)
}

func TestCCIPChainReader_DiscoverContracts_GetOfframpStaticConfig_Errors(t *testing.T) {
	ctx := tests.Context(t)
	destChain := cciptypes.ChainSelector(1)
	sourceChain1 := cciptypes.ChainSelector(2)
	sourceChain2 := cciptypes.ChainSelector(3)
	s1Onramp := []byte{0x1}
	s2Onramp := []byte{0x2}
	destExtended := reader_mocks.NewMockExtended(t)

	// mock the call for sourceChain1 - successful
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		primitives.Unconfirmed,
		map[string]any{"sourceChainSelector": sourceChain1},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*sourceChainConfig)
		v.OnRamp = s1Onramp
		v.IsEnabled = true
	}))
	// mock the call for sourceChain2 - successful
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameGetSourceChainConfig,
		primitives.Unconfirmed,
		map[string]any{"sourceChainSelector": sourceChain2},
		mock.Anything,
	).Return(nil).Run(withReturnValueOverridden(func(returnVal interface{}) {
		v := returnVal.(*sourceChainConfig)
		v.OnRamp = s2Onramp
		v.IsEnabled = true
	}))
	// mock the call to get the nonce manager - failure
	getLatestValueErr := errors.New("some error")
	destExtended.EXPECT().ExtendedGetLatestValue(
		mock.Anything,
		consts.ContractNameOffRamp,
		consts.MethodNameOfframpGetStaticConfig,
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

	_, err := ccipChainReader.DiscoverContracts(ctx, destChain)
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
