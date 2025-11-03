// Package v2 implements CCTPv2 token data observation for USDC cross-chain transfers.
//
// # Overview
//
// This package observes CCIP messages containing USDC tokens and fetches attestations
// from Circle's CCTP v2 API. These attestations prove that USDC was burned on the source
// chain and allow minting on the destination chain.
//
// # CCTPv2 vs CCTPv1
//
// CCTPv2 differs from v1 in several key ways:
//   - V1 uses nonce-based identification (sequential counter per source chain)
//   - V2 uses depositHash-based identification (content-addressable hash of transfer parameters)
//   - V1 API: GET /v1/attestations/{messageHash}
//   - V2 API: GET /v2/messages/{sourceDomain}?transactionHash={txHash}
//   - V2 allows multiple identical transfers to share attestations (fungible)
package v2

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// AttestationEncoder encodes CCTP message and attestation into format expected by USDC token pool.
type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

// CCTP version tags for identifying V2 messages.
const (
	// CCTP_VERSION_2_TAG identifies standard CCTP V2 transfers (slow transfers).
	// Preimage: keccak256("CCTP_V2")
	CCTP_VERSION_2_TAG = 0xb148ea5f

	// CCTP_VERSION_2_CCV_TAG identifies CCTP V2 transfers with CCIP v1.7 fast transfer support.
	// CCV = Cross-Chain Verification. Enables fast transfers with verification infrastructure.
	// Preimage: keccak256("CCTP_V2_CCV")
	CCTP_VERSION_2_CCV_TAG = 0x3047587c
)

// SourceTokenDataPayloadV2 represents the CCTPv2 source pool data embedded in message token data.
type SourceTokenDataPayloadV2 struct {
	SourceDomain uint32
	DepositHash  [32]byte
}

// CCTPv2TokenDataObserver observes CCTPv2 USDC messages and fetches attestations from Circle's v2 API.
type CCTPv2TokenDataObserver struct {
	lggr                     logger.Logger
	destChainSelector        cciptypes.ChainSelector
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
	attestationEncoder       AttestationEncoder
	httpClient               CCTPv2HTTPClient
}

// NewCCTPv2TokenDataObserver creates a new CCTPv2 token data observer.
func NewCCTPv2TokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	attestationEncoder AttestationEncoder,
	httpClient CCTPv2HTTPClient,
) *CCTPv2TokenDataObserver {
	return &CCTPv2TokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		httpClient:               httpClient,
	}
}

// Observe fetches CCTPv2 attestations for USDC messages from Circle's API.
func (o *CCTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	lggr := logutil.WithContextValues(ctx, o.lggr)

	// Step 1: Filter and decode CCTPv2 USDC messages
	// Identifies which messages contain CCTPv2 tokens and extracts SourceTokenDataPayloadV2
	v2Messages := o.pickOnlyCCTPv2Messages(lggr, messages)
	if len(v2Messages) == 0 {
		// No V2 messages to process, return observations for all messages
		return o.extractTokenData(ctx, lggr, messages, nil)
	}

	// Step 2: Extract transaction hashes from message headers
	// We need txHash to query Circle's CCTPv2 API
	txHashes := o.extractTransactionHashes(messages, v2Messages)

	// Step 3: Group messages by (sourceDomain, txHash) for efficient API calls
	// One transaction may have multiple USDC transfers
	txGroups := o.groupMessagesByTransaction(txHashes, v2Messages)

	// Step 4: Fetch CCTPv2 messages and attestations from Circle API
	// Calls GET /v2/messages/{sourceDomain}?transactionHash={hash}
	// Matches messages to MessageTokenIDs using depositHash
	cctpMessages, err := o.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)
	if err != nil {
		return nil, fmt.Errorf("fetch CCTPv2 attestations: %w", err)
	}

	// Step 5: Build AttestationStatus objects from matched CCTP messages
	// Decodes hex-encoded message and attestation bytes into AttestationStatus structures
	attestations := o.buildAttestationStatuses(cctpMessages, v2Messages)

	// Step 6: Build final TokenDataObservations
	// Encodes attestations and creates TokenData for each message token
	return o.extractTokenData(ctx, lggr, messages, attestations)
}

// IsTokenSupported checks if the given token is a supported CCTPv2 USDC token.
func (o *CCTPv2TokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	return strings.EqualFold(o.supportedPoolsBySelector[sourceChain], msgToken.SourcePoolAddress.String())
}

// Close cleans up resources used by the observer.
func (o *CCTPv2TokenDataObserver) Close() error {
	return nil
}

// pickOnlyCCTPv2Messages filters messages containing CCTPv2 USDC tokens and decodes their payloads.
func (o *CCTPv2TokenDataObserver) pickOnlyCCTPv2Messages(
	lggr logger.Logger,
	messages exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2 {
	result := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2)

	// Iterate through each chain
	for chainSelector, chainMessages := range messages {
		// Iterate through each message
		for seqNum, message := range chainMessages {
			// Iterate through each token
			for i, tokenAmount := range message.TokenAmounts {
				// Check if token is from a supported pool
				if !o.IsTokenSupported(chainSelector, tokenAmount) {
					continue
				}

				// Try to decode CCTPv2 payload from ExtraData
				payload, err := DecodeSourceTokenDataPayloadV2(tokenAmount.ExtraData)
				if err != nil {
					lggr.Warnw(
						"Failed to decode CCTPv2 source token data",
						"chainSelector", chainSelector,
						"seqNum", seqNum,
						"tokenIndex", i,
						"error", err,
					)
					continue
				}

				// Lazy initialize chain map
				if result[chainSelector] == nil {
					result[chainSelector] = make(map[reader.MessageTokenID]*SourceTokenDataPayloadV2)
				}

				// Store the decoded payload
				msgTokenID := reader.NewMessageTokenID(seqNum, i)
				result[chainSelector][msgTokenID] = payload
			}
		}
	}

	return result
}

// extractTransactionHashes gets transaction hash for each V2 message token.
func (o *CCTPv2TokenDataObserver) extractTransactionHashes(
	messages exectypes.MessageObservations,
	v2Messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]string {
	result := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]string)

	// Iterate through each chain's V2 messages
	for chainSelector, chainV2Messages := range v2Messages {
		// Initialize map for this chain
		result[chainSelector] = make(map[reader.MessageTokenID]string)

		// For each V2 message token, extract its transaction hash
		for msgTokenID := range chainV2Messages {
			// Extract sequence number from MessageTokenID
			seqNum := msgTokenID.SeqNr

			// Look up the full message using chainSelector and seqNum
			message, ok := messages[chainSelector][seqNum]
			if !ok {
				// Message not found - this shouldn't happen but be defensive
				o.lggr.Warnw(
					"Message not found for V2 token",
					"chainSelector", chainSelector,
					"seqNum", seqNum,
					"messageTokenID", msgTokenID,
				)
				continue
			}

			// Extract TxHash from message header
			txHash := message.Header.TxHash

			// Store the mapping
			result[chainSelector][msgTokenID] = txHash
		}
	}

	return result
}

// groupMessagesByTransaction groups messages by (sourceDomain, txHash) for batch API calls.
func (o *CCTPv2TokenDataObserver) groupMessagesByTransaction(
	txHashes map[cciptypes.ChainSelector]map[reader.MessageTokenID]string,
	v2Messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID {
	result := make(map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID)

	// Iterate through each chain
	for chainSelector, chainV2Messages := range v2Messages {
		// Initialize map for this chain
		if result[chainSelector] == nil {
			result[chainSelector] = make(map[TxKey][]reader.MessageTokenID)
		}

		// Group messages by (sourceDomain, txHash)
		for msgTokenID, v2Msg := range chainV2Messages {
			// Get transaction hash for this message token
			txHash, ok := txHashes[chainSelector][msgTokenID]
			if !ok {
				// This shouldn't happen if extractTransactionHashes is implemented correctly,
				// but we skip messages without transaction hashes to be defensive
				o.lggr.Warnw(
					"Transaction hash not found for MessageTokenID",
					"chainSelector", chainSelector,
					"messageTokenID", msgTokenID,
				)
				continue
			}

			// Create composite key from sourceDomain (from v2Msg) and txHash
			txKey := TxKey{
				SourceDomain: v2Msg.SourceDomain,
				TxHash:       txHash,
			}

			// Append MessageTokenID to the group for this transaction
			result[chainSelector][txKey] = append(result[chainSelector][txKey], msgTokenID)
		}
	}

	return result
}

// depositHashResult holds the result of a depositHash calculation.
type depositHashResult struct {
	hash [32]byte
	err  error
}

// calculateDepositHashes calculates deposit hashes for all Circle messages
func (o *CCTPv2TokenDataObserver) calculateDepositHashes(
	cctpMessages []CCTPv2Message,
	chainSelector cciptypes.ChainSelector,
	txHash string,
) map[int]depositHashResult {
	calculatedHashes := make(map[int]depositHashResult)

	for i, cctpMsg := range cctpMessages {
		hash, err := calculateDepositHash(cctpMsg.DecodedMessage)
		if err != nil {
			o.lggr.Warnw(
				"Failed to calculate depositHash for Circle message",
				"chainSelector", chainSelector,
				"txHash", txHash,
				"messageIndex", i,
				"error", err,
			)
		}
		calculatedHashes[i] = depositHashResult{hash: hash, err: err}
	}

	return calculatedHashes
}

// fetchMessagesForTransaction calls Circle API for a single transaction.
func (o *CCTPv2TokenDataObserver) fetchMessagesForTransaction(
	ctx context.Context,
	chainSelector cciptypes.ChainSelector,
	txKey TxKey,
) (CCTPv2Messages, error) {
	cctpMessages, err := o.httpClient.GetMessages(
		ctx,
		chainSelector,
		txKey.SourceDomain,
		txKey.TxHash,
	)

	if err != nil {
		o.lggr.Warnw(
			"Failed to fetch CCTP messages from Circle API",
			"chainSelector", chainSelector,
			"sourceDomain", txKey.SourceDomain,
			"txHash", txKey.TxHash,
			"error", err,
		)
		return CCTPv2Messages{}, err
	}

	o.lggr.Debugw(
		"Fetched CCTP messages from Circle API",
		"chainSelector", chainSelector,
		"sourceDomain", txKey.SourceDomain,
		"txHash", txKey.TxHash,
		"numMessages", len(cctpMessages.Messages),
	)

	return cctpMessages, nil
}

// matchMessageToCircleMessages finds a Circle message matching the expected depositHash.
// Returns true and the message index if found, false otherwise.
func (o *CCTPv2TokenDataObserver) matchMessageToCircleMessages(
	expectedHash [32]byte,
	cctpMessages []CCTPv2Message,
	calculatedHashes map[int]depositHashResult,
	usedMessageIndices map[int]bool,
) (matched bool, messageIndex int) {
	for i := range cctpMessages {
		// Skip already-assigned messages
		if usedMessageIndices[i] {
			continue
		}

		// Get pre-calculated depositHash
		hashResult := calculatedHashes[i]
		if hashResult.err != nil {
			continue
		}

		// Check if hashes match
		if hashResult.hash == expectedHash {
			return true, i
		}
	}

	// No match found
	return false, -1
}

// processTransactionMessages fetches and matches all messages for a single transaction.
func (o *CCTPv2TokenDataObserver) processTransactionMessages(
	ctx context.Context,
	chainSelector cciptypes.ChainSelector,
	txKey TxKey,
	msgTokenIDs []reader.MessageTokenID,
	v2MessagesForChain map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) map[reader.MessageTokenID]CCTPv2Message {
	result := make(map[reader.MessageTokenID]CCTPv2Message)

	// Fetch Circle messages for this transaction
	cctpMessages, err := o.fetchMessagesForTransaction(ctx, chainSelector, txKey)
	if err != nil {
		return result
	}

	// Pre-calculate deposit hashes for all Circle messages
	calculatedHashes := o.calculateDepositHashes(cctpMessages.Messages, chainSelector, txKey.TxHash)

	// Track which Circle messages have been assigned
	usedMessageIndices := make(map[int]bool)

	// Match each MessageTokenID to a Circle message
	for _, msgTokenID := range msgTokenIDs {
		// Get expected depositHash
		v2Msg, ok := v2MessagesForChain[msgTokenID]
		if !ok {
			o.lggr.Warnw(
				"v2Message not found for MessageTokenID",
				"chainSelector", chainSelector,
				"messageTokenID", msgTokenID,
			)
			continue
		}

		// Find matching Circle message
		matched, msgIndex := o.matchMessageToCircleMessages(
			v2Msg.DepositHash,
			cctpMessages.Messages,
			calculatedHashes,
			usedMessageIndices,
		)

		if matched {
			o.lggr.Debugw(
				"Matched Circle message to MessageTokenID by depositHash",
				"chainSelector", chainSelector,
				"messageTokenID", msgTokenID,
				"messageIndex", msgIndex,
				"depositHash", fmt.Sprintf("%x", v2Msg.DepositHash),
			)
			result[msgTokenID] = cctpMessages.Messages[msgIndex]
			usedMessageIndices[msgIndex] = true
		} else {
			o.lggr.Warnw(
				"No matching Circle message found for MessageTokenID",
				"chainSelector", chainSelector,
				"txHash", txKey.TxHash,
				"messageTokenID", msgTokenID,
				"expectedDepositHash", fmt.Sprintf("%x", v2Msg.DepositHash),
				"numCircleMessages", len(cctpMessages.Messages),
			)
		}
	}

	return result
}

// fetchCCTPv2Attestations queries Circle API for attestations and matches them to MessageTokenIDs using depositHash.
func (o *CCTPv2TokenDataObserver) fetchCCTPv2Attestations(
	ctx context.Context,
	txGroups map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID,
	v2Messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message, error) {
	result := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message)

	// Iterate through each chain
	for chainSelector, txKeyToMsgIDs := range txGroups {
		result[chainSelector] = make(map[reader.MessageTokenID]CCTPv2Message)

		// Process each transaction
		for txKey, msgTokenIDs := range txKeyToMsgIDs {
			txMessages := o.processTransactionMessages(
				ctx,
				chainSelector,
				txKey,
				msgTokenIDs,
				v2Messages[chainSelector],
			)

			// Merge transaction results into overall result
			for msgTokenID, cctpMsg := range txMessages {
				result[chainSelector][msgTokenID] = cctpMsg
			}
		}
	}

	return result, nil
}

// buildAttestationStatuses transforms matched CCTP messages into AttestationStatus objects.
// It assumes that cctpMessages have already been validated and matched by depositHash in fetchCCTPv2Attestations.
func (o *CCTPv2TokenDataObserver) buildAttestationStatuses(
	cctpMessages map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message,
	v2Messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus {
	result := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus)

	// Iterate through all v2Messages to transform each one
	for chainSelector, chainMessages := range v2Messages {
		if result[chainSelector] == nil {
			result[chainSelector] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)
		}

		for msgTokenID, v2Msg := range chainMessages {
			// Look up the corresponding CCTP message from Circle API
			cctpMsg, found := cctpMessages[chainSelector][msgTokenID]
			if !found {
				// CCTP message not found in API response (no match was found in fetchCCTPv2Attestations)
				result[chainSelector][msgTokenID] = tokendata.ErrorAttestationStatus(tokendata.ErrDataMissing)
				continue
			}

			// Decode message and attestation from hex strings
			messageBytes, err := hexDecode(cctpMsg.Message)
			if err != nil {
				result[chainSelector][msgTokenID] = tokendata.ErrorAttestationStatus(
					fmt.Errorf("decode message hex: %w", err))
				continue
			}

			attestationBytes, err := hexDecode(cctpMsg.Attestation)
			if err != nil {
				result[chainSelector][msgTokenID] = tokendata.ErrorAttestationStatus(
					fmt.Errorf("decode attestation hex: %w", err))
				continue
			}

			// Success - create successful AttestationStatus with the expected depositHash
			result[chainSelector][msgTokenID] = tokendata.SuccessAttestationStatus(
				v2Msg.DepositHash[:],
				messageBytes,
				attestationBytes,
			)
		}
	}

	return result
}

// hexDecode decodes a hex string (with or without 0x prefix) to bytes.
func hexDecode(hexStr string) ([]byte, error) {
	hexStr = strings.TrimPrefix(hexStr, "0x")
	return hex.DecodeString(hexStr)
}

// extractTokenData builds final TokenDataObservations by iterating through all messages and their tokens.
// For CCTPv2 USDC tokens, it fetches attestation data in attestationToTokenData.
// For unsupported tokens, it returns NotSupportedTokenData. Handles nil attestations when no V2 messages exist.
func (o *CCTPv2TokenDataObserver) extractTokenData(
	ctx context.Context,
	lggr logger.Logger,
	messages exectypes.MessageObservations,
	attestations map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus,
) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)

	for chainSelector, chainMessages := range messages {
		tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seqNum, message := range chainMessages {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))

			for i, tokenAmount := range message.TokenAmounts {
				if !o.IsTokenSupported(chainSelector, tokenAmount) {
					lggr.Debugw(
						"Ignoring unsupported token",
						"seqNum", seqNum,
						"sourceChainSelector", chainSelector,
						"sourcePoolAddress", tokenAmount.SourcePoolAddress.String(),
						"destTokenAddress", tokenAmount.DestTokenAddress.String(),
					)
					tokenData[i] = exectypes.NotSupportedTokenData()
				} else {
					var chainAttestations map[reader.MessageTokenID]tokendata.AttestationStatus
					if attestations != nil {
						chainAttestations = attestations[chainSelector]
					}
					tokenData[i] = o.attestationToTokenData(ctx, seqNum, i, chainAttestations)
				}
			}

			tokenObservations[chainSelector][seqNum] = exectypes.NewMessageTokenData(tokenData...)
		}
	}

	return tokenObservations, nil
}

// attestationToTokenData looks up and encodes attestation for a specific token in a message.
// Returns ErrorTokenData if attestation is missing, has errors, or encoding fails.
// On success, uses attestationEncoder to format data for the USDC token pool.
func (o *CCTPv2TokenDataObserver) attestationToTokenData(
	ctx context.Context,
	seqNum cciptypes.SeqNum,
	tokenIndex int,
	attestations map[reader.MessageTokenID]tokendata.AttestationStatus,
) exectypes.TokenData {
	status, ok := attestations[reader.NewMessageTokenID(seqNum, tokenIndex)]
	if !ok {
		return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	}
	if status.Error != nil {
		return exectypes.NewErrorTokenData(status.Error)
	}
	tokenData, err := o.attestationEncoder(ctx, status.MessageBody, status.Attestation)
	if err != nil {
		return exectypes.NewErrorTokenData(fmt.Errorf("unable to encode attestation: %w", err))
	}
	return exectypes.NewSuccessTokenData(tokenData)
}

// TxKey represents a unique identifier for grouping messages by transaction.
type TxKey struct {
	SourceDomain uint32
	TxHash       string
}

// calculateDepositHash calculates the depositHash from CCTPv2 message fields.
// This must match Solidity's _calculateDepositHash in USDCSourcePoolDataCodec.sol.
// The hash is: keccak256(abi.encode(sourceDomain, amount, destinationDomain, mintRecipient,
// burnToken, destinationCaller, maxFee, minFinalityThreshold))
func calculateDepositHash(msg CCTPv2DecodedMessage) ([32]byte, error) {
	var zero [32]byte

	// Parse sourceDomain
	sourceDomain, err := strconv.ParseUint(msg.SourceDomain, 10, 32)
	if err != nil {
		return zero, fmt.Errorf("parse sourceDomain: %w", err)
	}

	// Parse destinationDomain
	destinationDomain, err := strconv.ParseUint(msg.DestinationDomain, 10, 32)
	if err != nil {
		return zero, fmt.Errorf("parse destinationDomain: %w", err)
	}

	// Parse amount
	amount := new(big.Int)
	if _, ok := amount.SetString(msg.DecodedMessageBody.Amount, 10); !ok {
		return zero, fmt.Errorf("parse amount: invalid number")
	}

	// Parse maxFee (defaults to 0 if empty)
	maxFee := new(big.Int)
	if msg.DecodedMessageBody.MaxFee != "" {
		if _, ok := maxFee.SetString(msg.DecodedMessageBody.MaxFee, 10); !ok {
			return zero, fmt.Errorf("parse maxFee: invalid number")
		}
	}

	// Parse minFinalityThreshold
	minFinalityThreshold, err := strconv.ParseUint(msg.MinFinalityThreshold, 10, 32)
	if err != nil {
		return zero, fmt.Errorf("parse minFinalityThreshold: %w", err)
	}

	// Parse hex addresses to bytes32
	mintRecipient, err := hexToBytes32(msg.DecodedMessageBody.MintRecipient)
	if err != nil {
		return zero, fmt.Errorf("parse mintRecipient: %w", err)
	}

	burnToken, err := hexToBytes32(msg.DecodedMessageBody.BurnToken)
	if err != nil {
		return zero, fmt.Errorf("parse burnToken: %w", err)
	}

	destinationCaller, err := hexToBytes32(msg.DestinationCaller)
	if err != nil {
		return zero, fmt.Errorf("parse destinationCaller: %w", err)
	}

	// Create ABI types for encoding
	uint32Type, _ := abi.NewType("uint32", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)

	// Define ABI encoding arguments (matches Solidity's abi.encode order)
	arguments := abi.Arguments{
		{Type: uint32Type},  // sourceDomain
		{Type: uint256Type}, // amount
		{Type: uint32Type},  // destinationDomain
		{Type: bytes32Type}, // mintRecipient
		{Type: bytes32Type}, // burnToken
		{Type: bytes32Type}, // destinationCaller
		{Type: uint256Type}, // maxFee
		{Type: uint32Type},  // minFinalityThreshold
	}

	// Pack the arguments using ABI encoding
	encoded, err := arguments.Pack(
		uint32(sourceDomain),
		amount,
		uint32(destinationDomain),
		mintRecipient,
		burnToken,
		destinationCaller,
		maxFee,
		uint32(minFinalityThreshold),
	)
	if err != nil {
		return zero, fmt.Errorf("ABI encode: %w", err)
	}

	// Calculate Keccak256 hash
	hash := crypto.Keccak256Hash(encoded)
	return hash, nil
}

// hexToBytes32 converts a hex string (with or without 0x prefix) to a [32]byte array.
func hexToBytes32(hexStr string) ([32]byte, error) {
	var result [32]byte

	// Remove 0x prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")

	// Decode hex string
	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		return result, fmt.Errorf("decode hex: %w", err)
	}

	// Validate length (should be 20 bytes for address, but we pad to 32)
	if len(decoded) > 32 {
		return result, fmt.Errorf("hex string too long: %d bytes", len(decoded))
	}

	// Copy to bytes32 (left-padded with zeros for addresses, which are 20 bytes)
	copy(result[32-len(decoded):], decoded)

	return result, nil
}

// DecodeSourceTokenDataPayloadV2 decodes SourceTokenDataPayloadV2 from ExtraData bytes.
// The payload is encoded as: bytes4(versionTag) + uint32(sourceDomain) + bytes32(depositHash)
// Total length: 40 bytes (4 + 4 + 32)
func DecodeSourceTokenDataPayloadV2(extraData cciptypes.Bytes) (*SourceTokenDataPayloadV2, error) {
	// Validate length
	if len(extraData) != 40 {
		return nil, fmt.Errorf("invalid V2 source pool data length: expected 40 bytes, got %d", len(extraData))
	}

	// Extract and validate version tag (bytes 0-3)
	versionTag := binary.BigEndian.Uint32(extraData[0:4])
	if versionTag != CCTP_VERSION_2_TAG && versionTag != CCTP_VERSION_2_CCV_TAG {
		return nil, fmt.Errorf("invalid CCTPv2 version tag: expected 0x%x or 0x%x, got 0x%x",
			CCTP_VERSION_2_TAG, CCTP_VERSION_2_CCV_TAG, versionTag)
	}

	// Extract sourceDomain (bytes 4-7, big-endian uint32)
	sourceDomain := binary.BigEndian.Uint32(extraData[4:8])

	// Extract depositHash (bytes 8-39)
	var depositHash [32]byte
	copy(depositHash[:], extraData[8:40])

	return &SourceTokenDataPayloadV2{
		SourceDomain: sourceDomain,
		DepositHash:  depositHash,
	}, nil
}
