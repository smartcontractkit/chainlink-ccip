package v2

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Common test constants
const (
	defaultTxHash = "0xabc123"

	// Standard test addresses
	supportedPoolAddr   = "0x1111111111111111111111111111111111111111"
	destTokenAddr       = "0x2222222222222222222222222222222222222222"
	unsupportedPoolAddr = "0x9999999999999999999999999999999999999999"

	// 32-byte addresses (used in CCTP v2 payloads)
	burnTokenAddr32         = "0x2222222222222222222222222222222222222222222222222222222222222222"
	destinationCallerAddr32 = "0x3333333333333333333333333333333333333333333333333333333333333333"
	mintRecipientAddr32     = "0x0000000000000000000000001234567890abcdef1234567890abcdef12345678"
)

// Common test data and setup
var (
	testCtx         = context.Background()
	testLogger      = mocks.NullLogger
	testSourceChain = cciptypes.ChainSelector(1)
)

// Helper functions for creating test data
func createExpectedMessageTokenData(tokenData ...exectypes.TokenData) exectypes.MessageTokenData {
	return exectypes.NewMessageTokenData(tokenData...)
}

func notSupportedToken() exectypes.TokenData {
	return exectypes.NotSupportedTokenData()
}

func errorToken(err error) exectypes.TokenData {
	return exectypes.NewErrorTokenData(err)
}

func successToken(data []byte) exectypes.TokenData {
	return exectypes.NewSuccessTokenData(data)
}

// TokenConfig represents configuration for creating a token in test messages
type TokenConfig struct {
	SourcePoolAddress string
	Amount            int64
	// Note: Nonce is always 0 for CCTP v2, so no longer configurable
}

// createTokenDataPayloadWithAmount creates a source token data payload hex string with the specified amount
func createTokenDataPayloadWithAmount(amount int64) string {
	// Structure: 0x + nonce(64) + sourceDomain(64) + cctpVersion(64) + amount(64) + rest
	return fmt.Sprintf("0x%064x%s%064x%s",
		0,                                   // nonce is always 0 (64 hex chars)
		sourceTokenDataPayloadHexV2[66:194], // sourceDomain + cctpVersion (128 hex chars)
		amount,                              // amount (64 hex chars)
		sourceTokenDataPayloadHexV2[258:])   // rest of the payload (from position 258)
}

// Helper functions for creating test data
func createCCIPMessageWithTokenConfigs(seqNum cciptypes.SeqNum, txHash string, tokens []TokenConfig) cciptypes.Message {
	tokenAmounts := make([]cciptypes.RampTokenAmount, len(tokens))
	destTokenAddress, _ := cciptypes.NewUnknownAddressFromHex(destTokenAddr)

	for i, token := range tokens {
		sourcePoolAddress, _ := cciptypes.NewUnknownAddressFromHex(token.SourcePoolAddress)

		// Create payload with the specific amount for this token
		payloadHex := createTokenDataPayloadWithAmount(token.Amount)
		extraData := mustBytes(payloadHex)

		tokenAmounts[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: sourcePoolAddress,
			DestTokenAddress:  destTokenAddress,
			Amount:            cciptypes.NewBigIntFromInt64(token.Amount),
			ExtraData:         extraData,
		}
	}

	return cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:      [32]byte{byte(seqNum)},
			SequenceNumber: seqNum,
			TxHash:         txHash,
		},
		TokenAmounts: tokenAmounts,
	}
}

// Convenience functions for common scenarios
func createCCIPMessageWithValidCCTPv2Tokens(
	seqNum cciptypes.SeqNum,
	txHash string,
) cciptypes.Message {
	tokens := []TokenConfig{{
		SourcePoolAddress: supportedPoolAddr,
		Amount:            1,
	}}
	return createCCIPMessageWithTokenConfigs(seqNum, txHash, tokens)
}

func createCCIPMessageWithMixedTokens(seqNum cciptypes.SeqNum, txHash string) cciptypes.Message {
	tokens := []TokenConfig{
		{
			SourcePoolAddress: supportedPoolAddr, // Supported pool
			Amount:            1000,
		},
		{
			SourcePoolAddress: unsupportedPoolAddr, // Unsupported pool
			Amount:            2000,
		},
	}
	return createCCIPMessageWithTokenConfigs(seqNum, txHash, tokens)
}

func createCCIPMessageWithMultipleTokens(
	seqNum cciptypes.SeqNum,
	txHash string,
	tokenCount int,
) cciptypes.Message {
	tokens := make([]TokenConfig, tokenCount)
	for i := 0; i < tokenCount; i++ {
		tokens[i] = TokenConfig{
			SourcePoolAddress: supportedPoolAddr,
			Amount:            int64(i + 1),
		}
	}
	return createCCIPMessageWithTokenConfigs(seqNum, txHash, tokens)
}

// createMatchingCCTPv2Message creates a CCTP v2 message that matches the sourceTokenDataPayloadHexV2 constant
func createMatchingCCTPv2Message(nonce string) Message {
	return Message{
		EventNonce:  nonce,
		CCTPVersion: 2,
		Status:      "complete",
		Message:     "0x1234567890abcdef",
		Attestation: "0xfedcba0987654321",
		DecodedMessage: DecodedMessage{
			SourceDomain:         "111",
			DestinationDomain:    "305419896",
			Nonce:                nonce,
			DestinationCaller:    destinationCallerAddr32,
			MinFinalityThreshold: "5",
			DecodedMessageBody: DecodedMessageBody{
				Amount:        "1000",
				BurnToken:     burnTokenAddr32,
				MintRecipient: mintRecipientAddr32,
				MaxFee:        "50",
			},
		},
	}
}

// Common test helpers
func createMessages(messages ...Message) Messages {
	return Messages{Messages: messages}
}

func defaultSupportedPools() map[cciptypes.ChainSelector]string {
	return map[cciptypes.ChainSelector]string{
		testSourceChain: supportedPoolAddr,
	}
}

func noopMockClient() func(*MockCCTPv2AttestationClient) {
	return func(m *MockCCTPv2AttestationClient) {}
}

func noopMockEncoder() func(*mockAttestationEncoder) {
	return func(m *mockAttestationEncoder) {}
}

// mustBytes32 converts a 0x-prefixed hex string to a [32]byte.
func mustBytes32(h string) (out [32]byte) {
	b, err := hex.DecodeString(strings.TrimPrefix(h, "0x"))
	if err != nil {
		panic(err)
	}
	copy(out[32-len(b):], b) // right-align like EVM abi-encoding
	return
}

func mustBytes(h string) []byte {
	b, err := hex.DecodeString(strings.TrimPrefix(h, "0x"))
	if err != nil {
		panic(err)
	}
	return b
}
