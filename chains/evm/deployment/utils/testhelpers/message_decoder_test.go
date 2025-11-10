package testhelpers

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestDecodeString(t *testing.T) {
	// Load a hex-encoded CCIP message (v1) from a string
	hexMessage := "0x010000000000000001000000000000000200000000000000010005091000030d400000cbe110ff8af6bb46e2a4e352d0c9026e3902ad3f1369ad49c6a57fc740dee6a114912cf1ee1d4a3a302dfc64d70bd023e6839a5dd814baa5fa2b829e9c90d06cd4824ccb54f550386ebc1400007e64e1fb0c487f25dd6d3601ff6af8d32e4e1400007e64e1fb0c487f25dd6d3601ff6af8d32e4e00000097010000000000000000000000000000000000000000000000000de0b6b3a764000014ce4ec7b524851e51d5c55eefbbb8e58e8ce2515f14e55c1374bc6a38cea4e86f3a7e5b1cf17db5418514e2c2bb2f43b91f65b5519708e340310394c72d8f1400007e64e1fb0c487f25dd6d3601ff6af8d32e4e00200000000000000000000000000000000000000000000000000000000000000012000d65326520746573742064617461"
	// Remove "0x" prefix if present
	if len(hexMessage) >= 2 && hexMessage[0:2] == "0x" {
		hexMessage = hexMessage[2:]
	}
	encodedMessage, err := hex.DecodeString(hexMessage)
	require.NoError(t, err, "Failed to decode hex message")

	msg, err := DecodeMessageV1(encodedMessage)
	require.NoError(t, err, "Failed to decode message")
	require.NotNil(t, msg, "Decoded message is nil")

	// Pretty print the decoded message
	t.Log("\n" + PrettyPrintMessage(msg))

	// Also print the hex-encoded message for reference
	t.Logf("\nOriginal encoded message (%d bytes):\n%s\n", len(encodedMessage), hex.EncodeToString(encodedMessage))
}

// MessageV1 represents a decoded CCIP MessageV1
type MessageV1 struct {
	Version             uint8
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	ExecutionGasLimit   uint32
	CallbackGasLimit    uint32
	Finality            uint16
	CCVAndExecutorHash  [32]byte
	OnRampAddress       []byte
	OffRampAddress      []byte
	Sender              []byte
	Receiver            []byte
	DestBlob            []byte
	TokenTransfers      []TokenTransferV1
	Data                []byte
}

// TokenTransferV1 represents a decoded token transfer
type TokenTransferV1 struct {
	Version            uint8
	Amount             *big.Int
	SourcePoolAddress  []byte
	SourceTokenAddress []byte
	DestTokenAddress   []byte
	TokenReceiver      []byte
	ExtraData          []byte
}

// DecodeMessageV1 decodes a byte slice into a MessageV1 struct following the v1 protocol format.
func DecodeMessageV1(encoded []byte) (*MessageV1, error) {
	if len(encoded) < 77 {
		return nil, fmt.Errorf("message too short: %d bytes, minimum is 77", len(encoded))
	}

	msg := &MessageV1{}
	offset := 0

	// Version (1 byte)
	msg.Version = encoded[offset]
	offset++
	if msg.Version != 1 {
		return nil, fmt.Errorf("invalid version: %d", msg.Version)
	}

	// sourceChainSelector (8 bytes, big endian)
	msg.SourceChainSelector = binary.BigEndian.Uint64(encoded[offset : offset+8])
	offset += 8

	// destChainSelector (8 bytes, big endian)
	msg.DestChainSelector = binary.BigEndian.Uint64(encoded[offset : offset+8])
	offset += 8

	// sequenceNumber (8 bytes, big endian)
	msg.SequenceNumber = binary.BigEndian.Uint64(encoded[offset : offset+8])
	offset += 8

	// executionGasLimit (4 bytes, big endian)
	if offset+4 > len(encoded) {
		return nil, fmt.Errorf("executionGasLimit out of bounds")
	}
	msg.ExecutionGasLimit = binary.BigEndian.Uint32(encoded[offset : offset+4])
	offset += 4

	// callbackGasLimit (4 bytes, big endian)
	if offset+4 > len(encoded) {
		return nil, fmt.Errorf("callbackGasLimit out of bounds")
	}
	msg.CallbackGasLimit = binary.BigEndian.Uint32(encoded[offset : offset+4])
	offset += 4

	// finality (2 bytes, big endian)
	if offset+2 > len(encoded) {
		return nil, fmt.Errorf("finality out of bounds")
	}
	msg.Finality = binary.BigEndian.Uint16(encoded[offset : offset+2])
	offset += 2

	// ccvAndExecutorHash (32 bytes)
	if offset+32 > len(encoded) {
		return nil, fmt.Errorf("ccvAndExecutorHash out of bounds")
	}
	copy(msg.CCVAndExecutorHash[:], encoded[offset:offset+32])
	offset += 32

	// onRampAddressLength and onRampAddress
	if offset >= len(encoded) {
		return nil, fmt.Errorf("onRamp length out of bounds")
	}
	onRampLen := int(encoded[offset])
	offset++
	if offset+onRampLen > len(encoded) {
		return nil, fmt.Errorf("onRamp address out of bounds")
	}
	msg.OnRampAddress = make([]byte, onRampLen)
	copy(msg.OnRampAddress, encoded[offset:offset+onRampLen])
	offset += onRampLen

	// offRampAddressLength and offRampAddress
	if offset >= len(encoded) {
		return nil, fmt.Errorf("offRamp length out of bounds")
	}
	offRampLen := int(encoded[offset])
	offset++
	if offset+offRampLen > len(encoded) {
		return nil, fmt.Errorf("offRamp address out of bounds")
	}
	msg.OffRampAddress = make([]byte, offRampLen)
	copy(msg.OffRampAddress, encoded[offset:offset+offRampLen])
	offset += offRampLen

	// senderLength and sender
	if offset >= len(encoded) {
		return nil, fmt.Errorf("sender length out of bounds")
	}
	senderLen := int(encoded[offset])
	offset++
	if offset+senderLen > len(encoded) {
		return nil, fmt.Errorf("sender address out of bounds")
	}
	msg.Sender = make([]byte, senderLen)
	copy(msg.Sender, encoded[offset:offset+senderLen])
	offset += senderLen

	// receiverLength and receiver
	if offset >= len(encoded) {
		return nil, fmt.Errorf("receiver length out of bounds")
	}
	receiverLen := int(encoded[offset])
	offset++
	if offset+receiverLen > len(encoded) {
		return nil, fmt.Errorf("receiver address out of bounds")
	}
	msg.Receiver = make([]byte, receiverLen)
	copy(msg.Receiver, encoded[offset:offset+receiverLen])
	offset += receiverLen

	// destBlobLength and destBlob
	if offset+2 > len(encoded) {
		return nil, fmt.Errorf("destBlob length out of bounds")
	}
	destBlobLen := int(binary.BigEndian.Uint16(encoded[offset : offset+2]))
	offset += 2
	if offset+destBlobLen > len(encoded) {
		return nil, fmt.Errorf("destBlob content out of bounds")
	}
	msg.DestBlob = make([]byte, destBlobLen)
	copy(msg.DestBlob, encoded[offset:offset+destBlobLen])
	offset += destBlobLen

	// tokenTransferLength and tokenTransfer
	if offset+2 > len(encoded) {
		return nil, fmt.Errorf("tokenTransfer length out of bounds")
	}
	tokenTransferLen := int(binary.BigEndian.Uint16(encoded[offset : offset+2]))
	offset += 2

	// Decode token transfers (0 or 1)
	if tokenTransferLen > 0 {
		expectedEnd := offset + tokenTransferLen
		tokenTransfer, newOffset, err := DecodeTokenTransferV1(encoded, offset)
		if err != nil {
			return nil, fmt.Errorf("failed to decode token transfer: %w", err)
		}
		if newOffset != expectedEnd {
			return nil, fmt.Errorf("token transfer length mismatch: expected %d, got %d", expectedEnd, newOffset)
		}
		msg.TokenTransfers = []TokenTransferV1{*tokenTransfer}
		offset = newOffset
	}

	// dataLength and data
	if offset+2 > len(encoded) {
		return nil, fmt.Errorf("data length out of bounds")
	}
	dataLen := int(binary.BigEndian.Uint16(encoded[offset : offset+2]))
	offset += 2
	if offset+dataLen > len(encoded) {
		return nil, fmt.Errorf("data content out of bounds")
	}
	msg.Data = make([]byte, dataLen)
	copy(msg.Data, encoded[offset:offset+dataLen])
	offset += dataLen

	// Ensure all bytes consumed
	if offset != len(encoded) {
		return nil, fmt.Errorf("message has extra bytes: consumed %d of %d", offset, len(encoded))
	}

	return msg, nil
}

// DecodeTokenTransferV1 decodes a token transfer from the encoded bytes
func DecodeTokenTransferV1(encoded []byte, offset int) (*TokenTransferV1, int, error) {
	tt := &TokenTransferV1{}

	// version (1 byte)
	if offset >= len(encoded) {
		return nil, 0, fmt.Errorf("version out of bounds")
	}
	tt.Version = encoded[offset]
	offset++
	if tt.Version != 1 {
		return nil, 0, fmt.Errorf("invalid token transfer version: %d", tt.Version)
	}

	// amount (32 bytes)
	if offset+32 > len(encoded) {
		return nil, 0, fmt.Errorf("amount out of bounds")
	}
	tt.Amount = new(big.Int).SetBytes(encoded[offset : offset+32])
	offset += 32

	// sourcePoolAddressLength and sourcePoolAddress
	if offset >= len(encoded) {
		return nil, 0, fmt.Errorf("sourcePool length out of bounds")
	}
	sourcePoolLen := int(encoded[offset])
	offset++
	if offset+sourcePoolLen > len(encoded) {
		return nil, 0, fmt.Errorf("sourcePool address out of bounds")
	}
	tt.SourcePoolAddress = make([]byte, sourcePoolLen)
	copy(tt.SourcePoolAddress, encoded[offset:offset+sourcePoolLen])
	offset += sourcePoolLen

	// sourceTokenAddressLength and sourceTokenAddress
	if offset >= len(encoded) {
		return nil, 0, fmt.Errorf("sourceToken length out of bounds")
	}
	sourceTokenLen := int(encoded[offset])
	offset++
	if offset+sourceTokenLen > len(encoded) {
		return nil, 0, fmt.Errorf("sourceToken address out of bounds")
	}
	tt.SourceTokenAddress = make([]byte, sourceTokenLen)
	copy(tt.SourceTokenAddress, encoded[offset:offset+sourceTokenLen])
	offset += sourceTokenLen

	// destTokenAddressLength and destTokenAddress
	if offset >= len(encoded) {
		return nil, 0, fmt.Errorf("destToken length out of bounds")
	}
	destTokenLen := int(encoded[offset])
	offset++
	if offset+destTokenLen > len(encoded) {
		return nil, 0, fmt.Errorf("destToken address out of bounds")
	}
	tt.DestTokenAddress = make([]byte, destTokenLen)
	copy(tt.DestTokenAddress, encoded[offset:offset+destTokenLen])
	offset += destTokenLen

	// tokenReceiverLength and tokenReceiver
	if offset >= len(encoded) {
		return nil, 0, fmt.Errorf("tokenReceiver length out of bounds")
	}
	tokenReceiverLen := int(encoded[offset])
	offset++
	if offset+tokenReceiverLen > len(encoded) {
		return nil, 0, fmt.Errorf("tokenReceiver address out of bounds")
	}
	tt.TokenReceiver = make([]byte, tokenReceiverLen)
	copy(tt.TokenReceiver, encoded[offset:offset+tokenReceiverLen])
	offset += tokenReceiverLen

	// extraDataLength and extraData
	if offset+2 > len(encoded) {
		return nil, 0, fmt.Errorf("extraData length out of bounds")
	}
	extraDataLen := int(binary.BigEndian.Uint16(encoded[offset : offset+2]))
	offset += 2
	if offset+extraDataLen > len(encoded) {
		return nil, 0, fmt.Errorf("extraData content out of bounds")
	}
	tt.ExtraData = make([]byte, extraDataLen)
	copy(tt.ExtraData, encoded[offset:offset+extraDataLen])
	offset += extraDataLen

	return tt, offset, nil
}

// PrettyPrintMessage formats a decoded message for readable output
func PrettyPrintMessage(msg *MessageV1) string {
	s := "=== CCIP Message V1 ===\n\n"
	s += "Protocol Header:\n"
	s += fmt.Sprintf("  Version:               %d\n", msg.Version)
	s += fmt.Sprintf("  Source Chain Selector: %d (0x%x)\n", msg.SourceChainSelector, msg.SourceChainSelector)
	s += fmt.Sprintf("  Dest Chain Selector:   %d (0x%x)\n", msg.DestChainSelector, msg.DestChainSelector)
	s += fmt.Sprintf("  Sequence Number:       %d\n", msg.SequenceNumber)
	s += fmt.Sprintf("  Execution Gas Limit:   %d\n", msg.ExecutionGasLimit)
	s += fmt.Sprintf("  Callback Gas Limit:    %d\n", msg.CallbackGasLimit)
	s += fmt.Sprintf("  Finality:              %d\n", msg.Finality)
	s += fmt.Sprintf("  CCV & Executor Hash:   0x%x\n", msg.CCVAndExecutorHash)
	s += fmt.Sprintf("  OnRamp Address:        %s\n", formatAddress(msg.OnRampAddress))
	s += fmt.Sprintf("  OffRamp Address:       %s\n", formatAddress(msg.OffRampAddress))
	s += "\n"

	s += "User Controlled Data:\n"
	s += fmt.Sprintf("  Sender:                %s\n", formatAddress(msg.Sender))
	s += fmt.Sprintf("  Receiver:              %s\n", formatAddress(msg.Receiver))
	s += fmt.Sprintf("  Dest Blob:             %s (%d bytes)\n", formatHex(msg.DestBlob), len(msg.DestBlob))
	s += "\n"

	s += fmt.Sprintf("Token Transfers: %d\n", len(msg.TokenTransfers))
	for i, tt := range msg.TokenTransfers {
		s += fmt.Sprintf("\n  Token Transfer #%d:\n", i+1)
		s += fmt.Sprintf("    Version:             %d\n", tt.Version)
		s += fmt.Sprintf("    Amount:              %s (wei)\n", tt.Amount.String())
		s += fmt.Sprintf("    Source Pool:         %s\n", formatAddress(tt.SourcePoolAddress))
		s += fmt.Sprintf("    Source Token:        %s\n", formatAddress(tt.SourceTokenAddress))
		s += fmt.Sprintf("    Dest Token:          %s\n", formatAddress(tt.DestTokenAddress))
		s += fmt.Sprintf("    Token Receiver:      %s\n", formatAddress(tt.TokenReceiver))
		s += fmt.Sprintf("    Extra Data:          %s (%d bytes)\n", formatHex(tt.ExtraData), len(tt.ExtraData))
	}
	s += "\n"

	s += fmt.Sprintf("Data Payload:          %s (%d bytes)\n", formatHex(msg.Data), len(msg.Data))
	s += "\n"

	return s
}

// formatAddress formats a byte slice as an Ethereum address or hex string
func formatAddress(addr []byte) string {
	if len(addr) == 0 {
		return "<empty>"
	}
	if len(addr) == 20 {
		return common.BytesToAddress(addr).Hex()
	}
	return "0x" + hex.EncodeToString(addr)
}

// formatHex formats a byte slice as a hex string with proper handling of empty slices
func formatHex(data []byte) string {
	if len(data) == 0 {
		return "<empty>"
	}
	if len(data) <= 32 {
		return "0x" + hex.EncodeToString(data)
	}
	return fmt.Sprintf("0x%s... (truncated)", hex.EncodeToString(data[:32]))
}

func TestMessageDecoding(t *testing.T) {
	// Create a hardcoded encoded MessageV1 following the protocol format
	// This represents a CCIP message from Ethereum Sepolia (16015286601757825753)
	// to Arbitrum Sepolia (3478487238524512106) with a token transfer

	var encoded []byte

	// Version (1 byte): 1
	encoded = append(encoded, 1)

	// Source Chain Selector (8 bytes): Ethereum Sepolia = 16015286601757825753
	srcChain := make([]byte, 8)
	binary.BigEndian.PutUint64(srcChain, 16015286601757825753)
	encoded = append(encoded, srcChain...)

	// Dest Chain Selector (8 bytes): Arbitrum Sepolia = 3478487238524512106
	dstChain := make([]byte, 8)
	binary.BigEndian.PutUint64(dstChain, 3478487238524512106)
	encoded = append(encoded, dstChain...)

	// Sequence Number (8 bytes): 42
	seqNum := make([]byte, 8)
	binary.BigEndian.PutUint64(seqNum, 42)
	encoded = append(encoded, seqNum...)

	// Execution Gas Limit (4 bytes): 330000
	execGasLimit := make([]byte, 4)
	binary.BigEndian.PutUint32(execGasLimit, 330000)
	encoded = append(encoded, execGasLimit...)

	// Callback Gas Limit (4 bytes): 200000
	callbackGasLimit := make([]byte, 4)
	binary.BigEndian.PutUint32(callbackGasLimit, 200000)
	encoded = append(encoded, callbackGasLimit...)

	// Finality (2 bytes): 100 blocks
	finality := make([]byte, 2)
	binary.BigEndian.PutUint16(finality, 100)
	encoded = append(encoded, finality...)

	// CCV and Executor Hash (32 bytes) - example hash
	ccvAndExecutorHash := [32]byte{
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
		0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0,
	}
	encoded = append(encoded, ccvAndExecutorHash[:]...)

	// OnRamp Address (1 byte length + 20 bytes address)
	onRampAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")
	encoded = append(encoded, 20) // length
	encoded = append(encoded, onRampAddr.Bytes()...)

	// OffRamp Address (1 byte length + 20 bytes address)
	offRampAddr := common.HexToAddress("0x0987654321098765432109876543210987654321")
	encoded = append(encoded, 20) // length
	encoded = append(encoded, offRampAddr.Bytes()...)

	// Sender Address (1 byte length + 20 bytes address)
	senderAddr := common.HexToAddress("0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa")
	encoded = append(encoded, 20) // length
	encoded = append(encoded, senderAddr.Bytes()...)

	// Receiver Address (1 byte length + 20 bytes address)
	receiverAddr := common.HexToAddress("0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB")
	encoded = append(encoded, 20) // length
	encoded = append(encoded, receiverAddr.Bytes()...)

	// Dest Blob (2 bytes length + 0 bytes content) - empty for this example
	destBlobLen := make([]byte, 2)
	binary.BigEndian.PutUint16(destBlobLen, 0)
	encoded = append(encoded, destBlobLen...)

	// Token Transfer
	// First, encode the token transfer
	var tokenTransfer []byte

	// Token Transfer Version (1 byte): 1
	tokenTransfer = append(tokenTransfer, 1)

	// Amount (32 bytes): 1000000000000000000 (1 token with 18 decimals)
	amount := big.NewInt(1000000000000000000)
	amountBytes := make([]byte, 32)
	amount.FillBytes(amountBytes)
	tokenTransfer = append(tokenTransfer, amountBytes...)

	// Source Pool Address (1 byte length + 20 bytes address)
	sourcePoolAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	tokenTransfer = append(tokenTransfer, 20) // length
	tokenTransfer = append(tokenTransfer, sourcePoolAddr.Bytes()...)

	// Source Token Address (1 byte length + 20 bytes address)
	sourceTokenAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")
	tokenTransfer = append(tokenTransfer, 20) // length
	tokenTransfer = append(tokenTransfer, sourceTokenAddr.Bytes()...)

	// Dest Token Address (1 byte length + 20 bytes address)
	destTokenAddr := common.HexToAddress("0x3333333333333333333333333333333333333333")
	tokenTransfer = append(tokenTransfer, 20) // length
	tokenTransfer = append(tokenTransfer, destTokenAddr.Bytes()...)

	// Token Receiver Address (1 byte length + 20 bytes address)
	tokenReceiverAddr := common.HexToAddress("0x4444444444444444444444444444444444444444")
	tokenTransfer = append(tokenTransfer, 20) // length
	tokenTransfer = append(tokenTransfer, tokenReceiverAddr.Bytes()...)

	// Extra Data (2 bytes length + 4 bytes content)
	extraData := []byte{0xde, 0xad, 0xbe, 0xef}
	extraDataLen := make([]byte, 2)
	binary.BigEndian.PutUint16(extraDataLen, uint16(len(extraData)))
	tokenTransfer = append(tokenTransfer, extraDataLen...)
	tokenTransfer = append(tokenTransfer, extraData...)

	// Add token transfer to message with its length prefix
	tokenTransferLenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(tokenTransferLenBytes, uint16(len(tokenTransfer)))
	encoded = append(encoded, tokenTransferLenBytes...)
	encoded = append(encoded, tokenTransfer...)

	// Data payload (2 bytes length + content)
	dataPayload := []byte("Hello, CCIP!")
	dataLen := make([]byte, 2)
	binary.BigEndian.PutUint16(dataLen, uint16(len(dataPayload)))
	encoded = append(encoded, dataLen...)
	encoded = append(encoded, dataPayload...)

	// Test decoding
	t.Run("decode_and_pretty_print", func(t *testing.T) {
		msg, err := DecodeMessageV1(encoded)
		require.NoError(t, err, "Failed to decode message")
		require.NotNil(t, msg, "Decoded message is nil")

		// Verify decoded values
		require.Equal(t, uint8(1), msg.Version)
		require.Equal(t, uint64(16015286601757825753), msg.SourceChainSelector)
		require.Equal(t, uint64(3478487238524512106), msg.DestChainSelector)
		require.Equal(t, uint64(42), msg.SequenceNumber)
		require.Equal(t, uint32(330000), msg.ExecutionGasLimit)
		require.Equal(t, uint32(200000), msg.CallbackGasLimit)
		require.Equal(t, uint16(100), msg.Finality)
		require.Equal(t, ccvAndExecutorHash, msg.CCVAndExecutorHash)
		require.Equal(t, onRampAddr.Bytes(), msg.OnRampAddress)
		require.Equal(t, offRampAddr.Bytes(), msg.OffRampAddress)
		require.Equal(t, senderAddr.Bytes(), msg.Sender)
		require.Equal(t, receiverAddr.Bytes(), msg.Receiver)
		require.Empty(t, msg.DestBlob)
		require.Len(t, msg.TokenTransfers, 1)

		// Verify token transfer
		tt := msg.TokenTransfers[0]
		require.Equal(t, uint8(1), tt.Version)
		require.Equal(t, amount, tt.Amount)
		require.Equal(t, sourcePoolAddr.Bytes(), tt.SourcePoolAddress)
		require.Equal(t, sourceTokenAddr.Bytes(), tt.SourceTokenAddress)
		require.Equal(t, destTokenAddr.Bytes(), tt.DestTokenAddress)
		require.Equal(t, tokenReceiverAddr.Bytes(), tt.TokenReceiver)
		require.Equal(t, extraData, tt.ExtraData)

		// Verify data payload
		require.Equal(t, dataPayload, msg.Data)

		// Pretty print the decoded message
		t.Log("\n" + PrettyPrintMessage(msg))

		// Also print the hex-encoded message for reference
		t.Logf("\nOriginal encoded message (%d bytes):\n%s\n", len(encoded), hex.EncodeToString(encoded))
	})

	t.Run("decode_empty_message", func(t *testing.T) {
		// Create a minimal valid message with no token transfers and no data
		var minEncoded []byte

		// Version
		minEncoded = append(minEncoded, 1)

		// Chain selectors and sequence number
		minEncoded = append(minEncoded, srcChain...)
		minEncoded = append(minEncoded, dstChain...)
		minEncoded = append(minEncoded, seqNum...)

		// Execution Gas Limit
		minEncoded = append(minEncoded, execGasLimit...)

		// Callback Gas Limit
		minEncoded = append(minEncoded, callbackGasLimit...)

		// Finality
		minEncoded = append(minEncoded, finality...)

		// CCV and Executor Hash
		minEncoded = append(minEncoded, ccvAndExecutorHash[:]...)

		// OnRamp Address
		minEncoded = append(minEncoded, 20)
		minEncoded = append(minEncoded, onRampAddr.Bytes()...)

		// OffRamp Address
		minEncoded = append(minEncoded, 20)
		minEncoded = append(minEncoded, offRampAddr.Bytes()...)

		// Sender and receiver
		minEncoded = append(minEncoded, 20)
		minEncoded = append(minEncoded, senderAddr.Bytes()...)
		minEncoded = append(minEncoded, 20)
		minEncoded = append(minEncoded, receiverAddr.Bytes()...)

		// Empty destBlob
		minEncoded = append(minEncoded, 0, 0)

		// Empty token transfers
		minEncoded = append(minEncoded, 0, 0)

		// Empty data
		minEncoded = append(minEncoded, 0, 0)

		msg, err := DecodeMessageV1(minEncoded)
		require.NoError(t, err, "Failed to decode minimal message")
		require.NotNil(t, msg)
		require.Empty(t, msg.TokenTransfers)
		require.Empty(t, msg.Data)
		require.Empty(t, msg.DestBlob)

		t.Log("\n" + PrettyPrintMessage(msg))
	})

	t.Run("decode_invalid_version", func(t *testing.T) {
		invalidEncoded := make([]byte, len(encoded))
		copy(invalidEncoded, encoded)
		invalidEncoded[0] = 2 // Invalid version

		_, err := DecodeMessageV1(invalidEncoded)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid version")
	})

	t.Run("decode_too_short", func(t *testing.T) {
		_, err := DecodeMessageV1([]byte{1, 2, 3})
		require.Error(t, err)
		require.Contains(t, err.Error(), "too short")
	})
}
