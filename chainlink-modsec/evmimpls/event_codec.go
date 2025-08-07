package evmimpls

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/libmodsec/modsectypes"
)

type EVMEventCodec struct {
	abi abi.ABI
}

func NewEVMEventCodec(theAbi abi.ABI) EVMEventCodec {
	return EVMEventCodec{
		abi: theAbi,
	}
}

func (c *EVMEventCodec) Decode(ctx context.Context, data []byte) (modsectypes.Message, error) {
	unpacked, err := c.abi.Events["CCIPMessageSent"].Inputs.Unpack(data)
	if err != nil {
		return modsectypes.Message{}, fmt.Errorf("failed to unpack log data: %w", err)
	}
	if len(unpacked) != 1 {
		return modsectypes.Message{}, fmt.Errorf("expected 1 argument, got %d", len(unpacked))
	}
	message, ok := unpacked[0].(struct {
		Header struct {
			MessageId           [32]byte "json:\"messageId\""
			SourceChainSelector uint64   "json:\"sourceChainSelector\""
			DestChainSelector   uint64   "json:\"destChainSelector\""
			SequenceNumber      uint64   "json:\"sequenceNumber\""
		} "json:\"header\""
		Sender         common.Address "json:\"sender\""
		Data           []byte         "json:\"data\""
		Receiver       []byte         "json:\"receiver\""
		FeeToken       common.Address "json:\"feeToken\""
		FeeTokenAmount *big.Int       "json:\"feeTokenAmount\""
		FeeValueJuels  *big.Int       "json:\"feeValueJuels\""
		TokenTransfer  struct {
			SourceTokenAddress common.Address "json:\"sourceTokenAddress\""
			SourcePoolAddress  common.Address "json:\"sourcePoolAddress\""
			DestTokenAddress   []byte         "json:\"destTokenAddress\""
			ExtraData          []byte         "json:\"extraData\""
			Amount             *big.Int       "json:\"amount\""
			DestExecData       []byte         "json:\"destExecData\""
			RequiredVerifierId [32]byte       "json:\"requiredVerifierId\""
		} "json:\"tokenTransfer\""
		Receipts []struct {
			ReceiptType       uint8          "json:\"receiptType\""
			Issuer            common.Address "json:\"issuer\""
			FeeTokenAmount    *big.Int       "json:\"feeTokenAmount\""
			DestGasLimit      uint64         "json:\"destGasLimit\""
			DestBytesOverhead uint32         "json:\"destBytesOverhead\""
			ExtraArgs         []byte         "json:\"extraArgs\""
		} "json:\"receipts\""
	})
	if !ok {
		return modsectypes.Message{}, fmt.Errorf("failed to cast unpacked message to expected struct")
	}

	receipts := make([]modsectypes.Receipt, len(message.Receipts))
	for i, r := range message.Receipts {
		receipts[i] = modsectypes.Receipt{
			ReceiptType:       r.ReceiptType,
			Issuer:            r.Issuer.Bytes(),
			FeeTokenAmount:    r.FeeTokenAmount,
			DestGasLimit:      r.DestGasLimit,
			DestBytesOverhead: r.DestBytesOverhead,
			ExtraArgs:         r.ExtraArgs,
		}
	}

	return modsectypes.Message{
		Header: modsectypes.Header{
			MessageID:           message.Header.MessageId,
			SourceChainSelector: message.Header.SourceChainSelector,
			DestChainSelector:   message.Header.DestChainSelector,
			SequenceNumber:      message.Header.SequenceNumber,
		},
		Sender:         message.Sender.Bytes(),
		Data:           message.Data,
		Receiver:       message.Receiver,
		FeeToken:       message.FeeToken.Bytes(),
		FeeTokenAmount: message.FeeTokenAmount,
		FeeValueJuels:  message.FeeValueJuels,
		TokenTransfer: modsectypes.TokenTransfer{
			SourceTokenAddress: message.TokenTransfer.SourceTokenAddress.Bytes(),
			SourcePoolAddress:  message.TokenTransfer.SourcePoolAddress.Bytes(),
			DestTokenAddress:   message.TokenTransfer.DestTokenAddress,
			ExtraData:          message.TokenTransfer.ExtraData,
			Amount:             message.TokenTransfer.Amount,
			DestExecData:       message.TokenTransfer.DestExecData,
			RequiredVerifierID: message.TokenTransfer.RequiredVerifierId,
		},
		Receipts: receipts,
	}, nil
}

// InternalABI is the ABI for the CCIPMessageSent event.
// TODO: should use the ABI from the gethwrappers generated next to the real contracts instead.
const InternalABI = `
[
  {
    "type": "event",
    "name": "CCIPMessageSent",
    "inputs": [
      {
        "name": "destChainSelector",
        "type": "uint64",
        "indexed": true,
        "internalType": "uint64"
      },
      {
        "name": "sequenceNumber",
        "type": "uint64",
        "indexed": true,
        "internalType": "uint64"
      },
      {
        "name": "message",
        "type": "tuple",
        "indexed": false,
        "internalType": "struct Internal.EVM2AnyVerifierMessage",
        "components": [
          {
            "name": "header",
            "type": "tuple",
            "internalType": "struct Internal.Header",
            "components": [
              { "name": "messageId", "type": "bytes32", "internalType": "bytes32" },
              { "name": "sourceChainSelector", "type": "uint64", "internalType": "uint64" },
              { "name": "destChainSelector", "type": "uint64", "internalType": "uint64" },
              { "name": "sequenceNumber", "type": "uint64", "internalType": "uint64" }
            ]
          },
          { "name": "sender", "type": "address", "internalType": "address" },
          { "name": "data", "type": "bytes", "internalType": "bytes" },
          { "name": "receiver", "type": "bytes", "internalType": "bytes" },
          { "name": "feeToken", "type": "address", "internalType": "address" },
          { "name": "feeTokenAmount", "type": "uint256", "internalType": "uint256" },
          { "name": "feeValueJuels", "type": "uint256", "internalType": "uint256" },
          {
            "name": "tokenTransfer",
            "type": "tuple",
            "internalType": "struct Internal.EVMTokenTransfer",
            "components": [
              { "name": "sourceTokenAddress", "type": "address", "internalType": "address" },
              { "name": "sourcePoolAddress", "type": "address", "internalType": "address" },
              { "name": "destTokenAddress", "type": "bytes", "internalType": "bytes" },
              { "name": "extraData", "type": "bytes", "internalType": "bytes" },
              { "name": "amount", "type": "uint256", "internalType": "uint256" },
              { "name": "destExecData", "type": "bytes", "internalType": "bytes" },
              { "name": "requiredVerifierId", "type": "bytes32", "internalType": "bytes32" }
            ]
          },
          {
            "name": "receipts",
            "type": "tuple[]",
            "internalType": "struct Internal.Receipt[]",
            "components": [
              { "name": "receiptType", "type": "uint8", "internalType": "enum Internal.ReceiptType" },
              { "name": "issuer", "type": "address", "internalType": "address" },
              { "name": "feeTokenAmount", "type": "uint256", "internalType": "uint256" },
              { "name": "destGasLimit", "type": "uint64", "internalType": "uint64" },
              { "name": "destBytesOverhead", "type": "uint32", "internalType": "uint32" },
              { "name": "extraArgs", "type": "bytes", "internalType": "bytes" }
            ]
          }
        ]
      }
    ]
  }
]
`
