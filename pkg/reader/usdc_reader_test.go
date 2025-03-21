package reader

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/internal"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	sel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	reader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_USDCMessageReader_New(t *testing.T) {
	address1 := "0x0000000000000000000000000000000000000001"
	address2 := "0x0000000000000000000000000000000000000002"

	emptyReaders := func() map[cciptypes.ChainSelector]*reader.MockContractReaderFacade {
		return map[cciptypes.ChainSelector]*reader.MockContractReaderFacade{}
	}

	tt := []struct {
		name         string
		tokensConfig map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig
		readers      func() map[cciptypes.ChainSelector]*reader.MockContractReaderFacade
		errorMessage string
	}{
		{
			name:         "empty tokens and readers works",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{},
			readers:      emptyReaders,
		},
		{
			name: "missing readers",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				cciptypes.ChainSelector(1): {
					SourcePoolAddress:            address1,
					SourceMessageTransmitterAddr: address2,
				},
			},
			readers:      emptyReaders,
			errorMessage: "validate reader existence: chain 1: contract reader not found",
		},
		{
			name: "binding errors",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				cciptypes.ChainSelector(1): {
					SourcePoolAddress:            address1,
					SourceMessageTransmitterAddr: address2,
				},
			},
			readers: func() map[cciptypes.ChainSelector]*reader.MockContractReaderFacade {
				readers := make(map[cciptypes.ChainSelector]*reader.MockContractReaderFacade)
				m := reader.NewMockContractReaderFacade(t)
				m.EXPECT().Bind(mock.Anything, mock.Anything).Return(errors.New("error"))
				readers[cciptypes.ChainSelector(1)] = m
				return readers
			},
			errorMessage: "unable to bind MessageTransmitter 0x0000000000000000000000000000000000000002 for chain 1",
		},
		{
			name: "happy path",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				cciptypes.ChainSelector(2): {
					SourcePoolAddress:            address1,
					SourceMessageTransmitterAddr: address2,
				},
			},
			readers: func() map[cciptypes.ChainSelector]*reader.MockContractReaderFacade {
				readers := make(map[cciptypes.ChainSelector]*reader.MockContractReaderFacade)
				m := reader.NewMockContractReaderFacade(t)
				m.EXPECT().Bind(mock.Anything, []types.BoundContract{
					{
						Address: address2,
						Name:    consts.ContractNameCCTPMessageTransmitter,
					},
				}).Return(nil)

				readers[cciptypes.ChainSelector(2)] = m
				return readers
			},
		},
	}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tests.Context(t)
			readers := make(map[cciptypes.ChainSelector]contractreader.ContractReaderFacade)
			for k, v := range tc.readers() {
				readers[k] = v
			}

			r, err := NewUSDCMessageReader(ctx, logger.Test(t), tc.tokensConfig, readers, mockAddrCodec)
			if tc.errorMessage != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorMessage)
			} else {
				require.NoError(t, err)
				require.NotNil(t, r)
			}
		})
	}
}

func Test_USDCMessageReader_MessagesByTokenID(t *testing.T) {
	ctx := tests.Context(t)
	emptyChain := cciptypes.ChainSelector(sel.ETHEREUM_MAINNET.Selector)
	emptyReader := reader.NewMockContractReaderFacade(t)
	emptyReader.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	emptyReader.EXPECT().QueryKey(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return([]types.Sequence{}, nil).Maybe()

	faultyChain := cciptypes.ChainSelector(sel.AVALANCHE_MAINNET.Selector)
	faultyReader := reader.NewMockContractReaderFacade(t)
	faultyReader.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	faultyReader.EXPECT().QueryKey(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil, errors.New("error")).Maybe()

	validSequence := []types.Sequence{
		{
			Data: &MessageSentEvent{Arg0: make([]byte, 128)},
		},
	}

	validChain := cciptypes.ChainSelector(sel.ETHEREUM_MAINNET_ARBITRUM_1.Selector)
	validChainCCTP := CCTPDestDomains[uint64(validChain)]
	validReader := reader.NewMockContractReaderFacade(t)
	validReader.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	validReader.EXPECT().QueryKey(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(validSequence, nil).Maybe()

	tokensConfigs := map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
		faultyChain: {
			SourcePoolAddress:            "0x000000000000000000000000000000000000000A",
			SourceMessageTransmitterAddr: "0x000000000000000000000000000000000000000B",
		},
		emptyChain: {
			SourcePoolAddress:            "0x000000000000000000000000000000000000000C",
			SourceMessageTransmitterAddr: "0x000000000000000000000000000000000000000D",
		},
		validChain: {
			SourcePoolAddress:            "0x000000000000000000000000000000000000000E",
			SourceMessageTransmitterAddr: "0x000000000000000000000000000000000000000F",
		},
	}

	contactReaders := map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
		faultyChain: faultyReader,
		emptyChain:  emptyReader,
		validChain:  validReader,
	}

	tokens := map[MessageTokenID]cciptypes.RampTokenAmount{
		NewMessageTokenID(1, 1): {
			ExtraData: NewSourceTokenDataPayload(11, validChainCCTP).ToBytes(),
		},
	}

	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	usdcReader, err := NewUSDCMessageReader(ctx, logger.Test(t), tokensConfigs, contactReaders, mockAddrCodec)
	require.NoError(t, err)

	tt := []struct {
		name           string
		sourceSelector cciptypes.ChainSelector
		destSelector   cciptypes.ChainSelector
		expectedMsgIDs []MessageTokenID
		errorMessage   string
	}{
		{
			name:           "should return empty dataset when chain doesn't have events",
			sourceSelector: emptyChain,
			destSelector:   faultyChain,
			expectedMsgIDs: []MessageTokenID{},
		},
		{
			name:           "should return error when chain reader errors",
			sourceSelector: faultyChain,
			destSelector:   emptyChain,
			errorMessage:   fmt.Sprintf("error querying contract reader for chain %d", faultyChain),
		},
		{
			name:           "should return error when CCTP domain is not supported",
			sourceSelector: emptyChain,
			destSelector:   cciptypes.ChainSelector(2),
			errorMessage:   "destination domain not found for chain 2",
		},
		{
			name:           "should return error when CCTP domain is not supported",
			sourceSelector: cciptypes.ChainSelector(sel.POLYGON_MAINNET.Selector),
			destSelector:   emptyChain,
			errorMessage:   fmt.Sprintf("no contract bound for chain %d", sel.POLYGON_MAINNET.Selector),
		},
		{
			name:           "valid chain return events but nothing is matched",
			sourceSelector: validChain,
			destSelector:   emptyChain,
			expectedMsgIDs: []MessageTokenID{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			messages, err1 := usdcReader.MessagesByTokenID(
				tests.Context(t),
				tc.sourceSelector,
				tc.destSelector,
				tokens,
			)

			if tc.errorMessage != "" {
				require.Error(t, err1)
				require.ErrorContains(t, err1, tc.errorMessage)
			} else {
				require.NoError(t, err)
				require.NotNil(t, messages)
				require.Equal(t, tc.expectedMsgIDs, maps.Keys(messages))
			}
		})
	}
}

func Test_MessageSentEvent_unpackID(t *testing.T) {
	nonEmptyEvent := eventID{}
	for i := 0; i < 32; i++ {
		nonEmptyEvent[i] = byte(i)
	}

	fullPayload := make([]byte, 0, 64)
	fullPayload = append(fullPayload, nonEmptyEvent[:]...)
	fullPayload = append(fullPayload, nonEmptyEvent[:]...)

	tt := []struct {
		name    string
		data    []byte
		want    eventID
		wantErr bool
	}{
		{
			name:    "event too short",
			data:    make([]byte, 31),
			wantErr: true,
		},
		{
			name: "event with proper length but empty",
			data: make([]byte, 32),
			want: eventID{},
		},
		{
			name: "event with proper length and data",
			data: fullPayload,
			want: nonEmptyEvent,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			event := MessageSentEvent{Arg0: tc.data}
			got, err := event.unpackID()

			if !tc.wantErr {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			} else {
				require.Error(t, err)
			}
		})

	}
}

func Test_SourceTokenDataPayload_ToBytes(t *testing.T) {
	tt := []struct {
		nonce        uint64
		sourceDomain uint32
	}{
		{
			nonce:        0,
			sourceDomain: 0,
		},
		{
			nonce:        2137,
			sourceDomain: 4,
		},
		{
			nonce:        2,
			sourceDomain: 2137,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("nonce=%d,sourceDomain=%d", tc.nonce, tc.sourceDomain), func(t *testing.T) {
			payload1 := NewSourceTokenDataPayload(tc.nonce, tc.sourceDomain)
			bytes := payload1.ToBytes()

			payload2, err := NewSourceTokenDataPayloadFromBytes(bytes)
			require.NoError(t, err)
			require.Equal(t, *payload1, *payload2)
		})
	}
}

func Test_SourceTokenDataPayload_FromBytes(t *testing.T) {
	tt := []struct {
		name    string
		data    []byte
		want    *SourceTokenDataPayload
		wantErr bool
	}{
		{
			name:    "too short data",
			data:    []byte{0x01, 0x02, 0x03},
			wantErr: true,
		},
		{
			name: "empty data but with proper length",
			data: make([]byte, 64),
			want: NewSourceTokenDataPayload(0, 0),
		},
		{
			name: "data with nonce and source domain",
			data: []byte{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
			},
			want: NewSourceTokenDataPayload(1, 2),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewSourceTokenDataPayloadFromBytes(tc.data)

			if !tc.wantErr {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			} else {
				require.Error(t, err)
			}
		})
	}
}
