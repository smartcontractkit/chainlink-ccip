package evmimpls_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls/gethwrappers"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/modsectypes"
)

func TestMessageCodec_Decode(t *testing.T) {
	codec, err := evmimpls.NewEVMMessageCodec()
	require.NoError(t, err)

	expectedMsg := gethwrappers.CCIPMessageSentEmitterEVM2AnyVerifierMessage{
		Header: gethwrappers.CCIPMessageSentEmitterHeader{
			MessageId:           [32]byte{1},
			SourceChainSelector: 1,
			DestChainSelector:   2,
			SequenceNumber:      3,
		},
		Sender:         common.HexToAddress("0x123"),
		Data:           []byte("test data"),
		Receiver:       []byte("test receiver"),
		FeeToken:       common.HexToAddress("0x456"),
		FeeTokenAmount: big.NewInt(100),
		FeeValueJuels:  big.NewInt(200),
		TokenTransfer: gethwrappers.CCIPMessageSentEmitterEVMTokenTransfer{
			SourceTokenAddress: common.HexToAddress("0x789"),
			SourcePoolAddress:  common.HexToAddress("0xabc"),
			DestTokenAddress:   []byte("dest token address"),
			ExtraData:          []byte("extra data"),
			Amount:             big.NewInt(300),
			DestExecData:       []byte("dest exec data"),
			RequiredVerifierId: [32]byte{2},
		},
		Receipts: []gethwrappers.CCIPMessageSentEmitterReceipt{
			{
				ReceiptType:       0,
				Issuer:            common.HexToAddress("0xdef"),
				FeeTokenAmount:    big.NewInt(50),
				DestGasLimit:      100000,
				DestBytesOverhead: 200,
				ExtraArgs:         []byte("receipt extra args"),
			},
		},
	}

	abi, err := gethwrappers.CCIPMessageSentEmitterMetaData.GetAbi()
	require.NoError(t, err)

	emitMethod, ok := abi.Methods["emitCCIPMessageSent"]
	require.True(t, ok)

	encoded, err := emitMethod.Inputs.Pack(expectedMsg)
	require.NoError(t, err)

	decodedMsg, err := codec.Decode(t.Context(), encoded)
	require.NoError(t, err)

	// Assertions
	require.Equal(t, expectedMsg.Header.MessageId, decodedMsg.Header.MessageID)
	// ... (add all other assertions)
}

func TestMessageCodec_Encode(t *testing.T) {
	codec, err := evmimpls.NewEVMMessageCodec()
	require.NoError(t, err)

	msgToEncode := modsectypes.Message{
		Header: modsectypes.Header{
			MessageID:           [32]byte{1},
			SourceChainSelector: 1,
			DestChainSelector:   2,
			SequenceNumber:      3,
		},
		Sender:   common.HexToAddress("0x123").Bytes(),
		Data:     []byte("test data"),
		Receiver: common.HexToAddress("0x456").Bytes(),
		TokenTransfer: modsectypes.TokenTransfer{
			SourcePoolAddress: common.HexToAddress("0xabc").Bytes(),
			DestTokenAddress:  common.HexToAddress("0xdef").Bytes(),
			Amount:            big.NewInt(300),
		},
		Receipts: []modsectypes.Receipt{
			{
				Issuer:            common.HexToAddress("0x123").Bytes(),
				ReceiptType:       modsectypes.ReceiptTypeExecutor,
				FeeTokenAmount:    big.NewInt(100),
				DestGasLimit:      100000,
				DestBytesOverhead: 200,
				ExtraArgs:         []byte("receipt extra args"),
			},
		},
	}

	_, err = codec.Encode(t.Context(), msgToEncode)
	require.NoError(t, err)
}
