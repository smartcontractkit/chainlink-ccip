package evmimpls

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/evmimpls/gethwrappers"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

type EVMMessageCodec struct {
	abi abi.ABI
}

var _ modsectypes.MessageCodec = (*EVMMessageCodec)(nil)

func NewEVMMessageCodec() (*EVMMessageCodec, error) {
	contractAbi, err := abi.JSON(strings.NewReader(gethwrappers.CCIPMessageSentEmitterABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract abi: %w", err)
	}

	return &EVMMessageCodec{
		abi: contractAbi,
	}, nil
}

// TODO: might not be right
func (c *EVMMessageCodec) extractGasLimit(receipts []modsectypes.Receipt) uint32 {
	gasLimit := uint32(0)
	for _, r := range receipts {
		gasLimit += uint32(r.DestGasLimit)
	}
	return gasLimit
}

func (c *EVMMessageCodec) Encode(ctx context.Context, message modsectypes.Message) ([]byte, error) {
	gasLimit := c.extractGasLimit(message.Receipts)

	requiredVerifiers := make([][]byte, len(message.Receipts))
	for i, r := range message.Receipts {
		requiredVerifiers[i] = r.Issuer
	}

	tokenAmts := []gethwrappers.CCIPMessageSentEmitterAny2EVMMultiProofTokenTransfer{}
	// Assuming only one token transfer for now, based on modsectypes.Message struct
	if message.TokenTransfer.Amount != nil && message.TokenTransfer.Amount.Cmp(big.NewInt(0)) > 0 {
		tokenAmts = append(tokenAmts, gethwrappers.CCIPMessageSentEmitterAny2EVMMultiProofTokenTransfer{
			SourcePoolAddress: message.TokenTransfer.SourcePoolAddress,
			DestTokenAddress:  common.BytesToAddress(message.TokenTransfer.DestTokenAddress),
			ExtraData:         message.TokenTransfer.ExtraData,
			Amount:            message.TokenTransfer.Amount,
		})
	}

	any2EVM := gethwrappers.CCIPMessageSentEmitterAny2EVMMultiProofMessage{
		Header: gethwrappers.CCIPMessageSentEmitterHeader{
			MessageId:           message.Header.MessageID,
			SourceChainSelector: message.Header.SourceChainSelector,
			DestChainSelector:   message.Header.DestChainSelector,
			SequenceNumber:      message.Header.SequenceNumber,
		},
		Sender:            message.Sender,
		Data:              message.Data,
		Receiver:          common.BytesToAddress(message.Receiver),
		GasLimit:          gasLimit,
		TokenAmounts:      tokenAmts,
		RequiredVerifiers: requiredVerifiers,
	}

	executeMethod, ok := c.abi.Methods["execute"]
	if !ok {
		return nil, fmt.Errorf("execute method not found in abi")
	}

	// We are encoding the first argument to the execute method, which is the message struct.
	messageArg := executeMethod.Inputs[0]
	args := abi.Arguments{messageArg}

	packed, err := args.Pack(any2EVM)
	if err != nil {
		return nil, fmt.Errorf("failed to pack Any2EVMMultiProofMessage: %w", err)
	}
	return packed, nil
}

// Decode implements modsectypes.MessageCodec.
func (c *EVMMessageCodec) Decode(ctx context.Context, data []byte) (modsectypes.Message, error) {
	emitMethod, ok := c.abi.Methods["emitCCIPMessageSent"]
	if !ok {
		return modsectypes.Message{}, fmt.Errorf("emitCCIPMessageSent method not found in abi")
	}

	unpacked, err := emitMethod.Inputs.Unpack(data)
	if err != nil {
		return modsectypes.Message{}, fmt.Errorf("failed to unpack message: %w", err)
	}
	if len(unpacked) == 0 {
		return modsectypes.Message{}, fmt.Errorf("unpacked to empty slice")
	}

	message := *abi.ConvertType(unpacked[0], new(gethwrappers.CCIPMessageSentEmitterEVM2AnyVerifierMessage)).(*gethwrappers.CCIPMessageSentEmitterEVM2AnyVerifierMessage)

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
