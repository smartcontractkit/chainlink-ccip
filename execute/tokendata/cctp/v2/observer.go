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
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

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
	result := make(exectypes.TokenDataObservations)

	for chainSelector, seqNumToMessage := range messages {
		v2TokenPayloads := o.getV2TokenPayloads(chainSelector, seqNumToMessage)

		// Extract and validate source domain ID
		sourceDomainID, err := getSourceDomainID(v2TokenPayloads)
		if err != nil {
			// Log the error and skip this chain - no CCTPv2 tokens or inconsistent data
			lggr.Warnw(
				"Failed to get source domain ID for chain",
				"chainSelector", chainSelector,
				"error", err,
			)
			// Create empty token data for all messages to preserve structure
			tokenData := o.createTokenData(ctx, chainSelector, seqNumToMessage, nil)
			result[chainSelector] = tokenData
			continue
		}

		txHashes := getTxHashes(seqNumToMessage)
		attestations, err := o.assignAttestationsToV2TokenPayloads(
			ctx, chainSelector, seqNumToMessage, txHashes, sourceDomainID, v2TokenPayloads,
		)
		if err != nil {
			// Log the error but continue with partial results
			// The method already creates error attestations for failed fetches
			lggr.Warnw(
				"Error assigning attestations to token payloads",
				"chainSelector", chainSelector,
				"error", err,
			)
		}

		tokenData := o.createTokenData(ctx, chainSelector, seqNumToMessage, attestations)
		result[chainSelector] = tokenData
	}

	return result, nil
}

// IsTokenSupported checks if the given token is a supported CCTPv2 USDC token.
// A token is considered supported if:
//  1. Its source pool address matches the configured CCTPv2 pool for the chain
//  2. Its ExtraData field contains a valid CCTPv2 payload that can be decoded
//
// This ensures that only tokens with both the correct pool AND valid payload
// structure are processed as CCTPv2 tokens.
func (o *CCTPv2TokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	// First check if the pool address matches
	if !strings.EqualFold(o.supportedPoolsBySelector[sourceChain], msgToken.SourcePoolAddress.String()) {
		return false
	}

	// Then verify the ExtraData can be decoded as a valid CCTPv2 payload
	_, err := DecodeSourceTokenDataPayloadV2(msgToken.ExtraData)
	return err == nil
}

// Close cleans up resources used by the observer.
func (o *CCTPv2TokenDataObserver) Close() error {
	return nil
}

// getV2TokenPayloads extracts and decodes CCTPv2 token payloads from CCIP messages.
// It processes messages from a single chain, identifying CCTPv2 USDC tokens using
// IsTokenSupported (which checks both pool address and payload validity).
//
// Returns a nested map: SeqNum -> tokenIndex -> SourceTokenDataPayloadV2
// Only valid CCTPv2 tokens are included. Tokens from unsupported pools or with
// invalid payloads are silently skipped.
func (o *CCTPv2TokenDataObserver) getV2TokenPayloads(
	chainSelector cciptypes.ChainSelector,
	messages map[cciptypes.SeqNum]cciptypes.Message,
) map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2 {
	result := make(map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2)

	// Iterate through each message
	for seqNum, message := range messages {
		messagePayloads := make(map[int]SourceTokenDataPayloadV2)

		// Process each token in the message
		for tokenIndex, tokenAmount := range message.TokenAmounts {
			if !o.IsTokenSupported(chainSelector, tokenAmount) {
				continue
			}

			// IsTokenSupported already validated that this decodes successfully
			payload, _ := DecodeSourceTokenDataPayloadV2(tokenAmount.ExtraData)

			// Add to results
			messagePayloads[tokenIndex] = *payload
		}

		// Only add to result if this message has CCTPv2 tokens
		if len(messagePayloads) > 0 {
			result[seqNum] = messagePayloads
		}
	}

	return result
}

// getSourceDomainID extracts the source domain ID from CCTPv2 token payloads.
// It validates that all payloads have the same source domain ID, as required
// for transfers from the same chain.
//
// Returns an error if:
//   - The payload map is empty (no CCTPv2 tokens found)
//   - Multiple different source domain IDs are found (data inconsistency)
func getSourceDomainID(v2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2) (uint32, error) {
	// Check if we have any payloads
	if len(v2TokenPayloads) == 0 {
		return 0, fmt.Errorf("no CCTPv2 token payloads found")
	}

	var sourceDomainID uint32
	var initialized bool

	// Iterate through all payloads to extract and validate source domain IDs
	for seqNum, tokenPayloads := range v2TokenPayloads {
		for tokenIndex, payload := range tokenPayloads {
			if !initialized {
				// First payload - use its domain ID as reference
				sourceDomainID = payload.SourceDomain
				initialized = true
			} else if payload.SourceDomain != sourceDomainID {
				// Found a different domain ID - this is an error
				return 0, fmt.Errorf(
					"inconsistent source domain IDs found: expected %d but got %d (seqNum=%d, tokenIndex=%d)",
					sourceDomainID, payload.SourceDomain, seqNum, tokenIndex,
				)
			}
		}
	}

	if !initialized {
		// This should not happen if len check passed, but be defensive
		return 0, fmt.Errorf("no token payloads found in non-empty map")
	}

	return sourceDomainID, nil
}

// getTxHashes groups messages by their transaction hash for efficient API calls.
// It extracts the TxHash from each message's header and creates a mapping from
// each unique transaction hash to the list of sequence numbers that share it.
func getTxHashes(messages map[cciptypes.SeqNum]cciptypes.Message) map[string][]cciptypes.SeqNum {
	result := make(map[string][]cciptypes.SeqNum)

	// Group messages by their transaction hash
	for seqNum, message := range messages {
		txHash := message.Header.TxHash

		// Append this sequence number to the list for this transaction hash
		result[txHash] = append(result[txHash], seqNum)
	}

	return result
}

// assignAttestationsToV2TokenPayloads fetches CCTPv2 attestations from Circle's API
// and assigns them to the corresponding token payloads. It uses transaction-scoped
// pooling to handle fungible attestations within the same transaction.
func (o *CCTPv2TokenDataObserver) assignAttestationsToV2TokenPayloads(
	ctx context.Context,
	chainSelector cciptypes.ChainSelector,
	messages map[cciptypes.SeqNum]cciptypes.Message,
	txHashToSeqNums map[string][]cciptypes.SeqNum,
	sourceDomainID uint32,
	seqNumToV2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2,
) (map[cciptypes.SeqNum]map[int]tokendata.AttestationStatus, error) {
	lggr := logutil.WithContextValues(ctx, o.lggr)
	result := make(map[cciptypes.SeqNum]map[int]tokendata.AttestationStatus)

	for txHash, seqNums := range txHashToSeqNums {
		cctpv2Messages, err := o.httpClient.GetMessages(ctx, chainSelector, sourceDomainID, txHash)
		if err != nil {
			lggr.Warnw(
				"Failed to fetch CCTPv2 messages from Circle API",
				"txHash", txHash,
				"sourceDomainID", sourceDomainID,
				"error", err,
			)
			o.assignErrorAttestations(result, messages, seqNumToV2TokenPayloads, seqNums, err)
			continue
		}

		attestations := extractAttestations(cctpv2Messages)
		o.assignSuccessAttestations(result, messages, seqNumToV2TokenPayloads, seqNums, attestations)
	}
	return result, nil
}

// assignErrorAttestations creates error attestations for all tokens in the given sequence numbers
func (o *CCTPv2TokenDataObserver) assignErrorAttestations(
	result map[cciptypes.SeqNum]map[int]tokendata.AttestationStatus,
	messages map[cciptypes.SeqNum]cciptypes.Message,
	seqNumToV2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2,
	seqNums []cciptypes.SeqNum,
	err error,
) {
	for _, seqNum := range seqNums {
		v2TokenPayloads, ok := seqNumToV2TokenPayloads[seqNum]
		if !ok {
			continue
		}
		result[seqNum] = createErrorAttestationsForTokens(messages, seqNum, v2TokenPayloads, err)
	}
}

// assignSuccessAttestations assigns fetched attestations to tokens in the given sequence numbers
func (o *CCTPv2TokenDataObserver) assignSuccessAttestations(
	result map[cciptypes.SeqNum]map[int]tokendata.AttestationStatus,
	messages map[cciptypes.SeqNum]cciptypes.Message,
	seqNumToV2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2,
	seqNums []cciptypes.SeqNum,
	attestations map[[32]byte][]tokendata.AttestationStatus,
) {
	for _, seqNum := range seqNums {
		v2TokenPayloads, ok := seqNumToV2TokenPayloads[seqNum]
		if !ok {
			continue
		}
		result[seqNum] = assignAttestationsForTokens(messages, seqNum, v2TokenPayloads, attestations)
	}
}

// createTokenData converts attestations to final TokenDataObservations format.
// It iterates through all messages and their tokens, encoding attestations for CCTPv2 USDC tokens
// and returning NotSupportedTokenData for unsupported tokens.
func (o *CCTPv2TokenDataObserver) createTokenData(
	ctx context.Context,
	chainSelector cciptypes.ChainSelector,
	messages map[cciptypes.SeqNum]cciptypes.Message,
	attestations map[cciptypes.SeqNum]map[int]tokendata.AttestationStatus,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

	for seqNum, message := range messages {
		tokenDataSlice := make([]exectypes.TokenData, len(message.TokenAmounts))

		for tokenIndex, tokenAmount := range message.TokenAmounts {
			// Check if token is supported
			if !o.IsTokenSupported(chainSelector, tokenAmount) {
				// Token is not supported
				tokenDataSlice[tokenIndex] = exectypes.NotSupportedTokenData()
				continue
			}

			// Look up attestation for this token
			attestationsBySeq, hasSeq := attestations[seqNum]
			if !hasSeq {
				// No attestations for this message
				tokenDataSlice[tokenIndex] = exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
				continue
			}

			attestation, hasToken := attestationsBySeq[tokenIndex]
			if !hasToken {
				// No attestation for this specific token
				tokenDataSlice[tokenIndex] = exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
				continue
			}

			// Check if attestation has an error
			if attestation.Error != nil {
				tokenDataSlice[tokenIndex] = exectypes.NewErrorTokenData(attestation.Error)
				continue
			}

			// Encode the attestation
			encodedData, err := o.attestationEncoder(ctx, attestation.MessageBody, attestation.Attestation)
			if err != nil {
				tokenDataSlice[tokenIndex] = exectypes.NewErrorTokenData(fmt.Errorf("unable to encode attestation: %w", err))
			} else {
				tokenDataSlice[tokenIndex] = exectypes.NewSuccessTokenData(encodedData)
			}
		}

		// Create MessageTokenData from the slice
		result[seqNum] = exectypes.NewMessageTokenData(tokenDataSlice...)
	}

	return result
}

// extractAttestations processes CCTPv2 messages from Circle's API and groups them by depositHash.
// Since GetMessages now filters for complete V2 messages, this function only needs to:
// 1. Calculate the depositHash for each message
// 2. Decode the message and attestation from hex
// 3. Group attestations by depositHash for fungible assignment
//
// Returns a map where each depositHash maps to a list of attestations that can be used
// interchangeably for tokens with the same depositHash within the transaction.
func extractAttestations(cctpV2Messages CCTPv2Messages) map[[32]byte][]tokendata.AttestationStatus {
	result := make(map[[32]byte][]tokendata.AttestationStatus)

	for _, msg := range cctpV2Messages.Messages {
		// Calculate the depositHash for this message
		depositHash, err := CalculateDepositHash(msg.DecodedMessage)
		if err != nil {
			// Skip messages where we can't calculate the depositHash
			// This shouldn't happen with valid API responses
			continue
		}

		// Decode the message bytes from hex
		messageBytes, err := hex.DecodeString(strings.TrimPrefix(msg.Message, "0x"))
		if err != nil {
			// Skip if we can't decode the message
			continue
		}

		// Decode the attestation bytes from hex
		attestationBytes, err := hex.DecodeString(strings.TrimPrefix(msg.Attestation, "0x"))
		if err != nil {
			// Skip if we can't decode the attestation
			continue
		}

		// Create the attestation status
		attestation := tokendata.AttestationStatus{
			// ID will be filled in by assignAttestationsToV2TokenPayloads
			MessageBody: messageBytes,
			Attestation: attestationBytes,
			Error:       nil,
		}

		// Add to the list for this depositHash
		result[depositHash] = append(result[depositHash], attestation)
	}

	return result
}

// TODO: doc, explain that it mutates attestations
func assignAttestationForV2TokenPayload(
	attestations map[[32]byte][]tokendata.AttestationStatus,
	v2TokenPayload SourceTokenDataPayloadV2,
) tokendata.AttestationStatus {
	attestationStatuses, ok := attestations[v2TokenPayload.DepositHash]
	if !ok || len(attestationStatuses) == 0 {
		return tokendata.ErrorAttestationStatus(tokendata.ErrDataMissing)
	}

	attestation := attestationStatuses[0]
	attestations[v2TokenPayload.DepositHash] = attestationStatuses[1:]
	return attestation
}

// getMessageIDBytes extracts message ID as bytes from a message
func getMessageIDBytes(messages map[cciptypes.SeqNum]cciptypes.Message, seqNum cciptypes.SeqNum) cciptypes.Bytes {
	msgID := messages[seqNum].Header.MessageID
	return cciptypes.Bytes(msgID[:])
}

// createErrorAttestationsForTokens creates error attestations for all tokens in a message
func createErrorAttestationsForTokens(
	messages map[cciptypes.SeqNum]cciptypes.Message,
	seqNum cciptypes.SeqNum,
	v2TokenPayloads map[int]SourceTokenDataPayloadV2,
	err error,
) map[int]tokendata.AttestationStatus {
	assignedAttestations := make(map[int]tokendata.AttestationStatus)
	msgIDBytes := getMessageIDBytes(messages, seqNum)

	for idx := range v2TokenPayloads {
		assignedAttestations[idx] = tokendata.AttestationStatus{
			ID:    msgIDBytes,
			Error: fmt.Errorf("failed to fetch attestation: %w", err),
		}
	}
	return assignedAttestations
}

// assignAttestationsForTokens assigns attestations to all tokens in a message
func assignAttestationsForTokens(
	messages map[cciptypes.SeqNum]cciptypes.Message,
	seqNum cciptypes.SeqNum,
	v2TokenPayloads map[int]SourceTokenDataPayloadV2,
	attestations map[[32]byte][]tokendata.AttestationStatus,
) map[int]tokendata.AttestationStatus {
	assignedAttestations := make(map[int]tokendata.AttestationStatus)
	msgIDBytes := getMessageIDBytes(messages, seqNum)

	for idx, v2TokenPayload := range v2TokenPayloads {
		assignedAttestation := assignAttestationForV2TokenPayload(attestations, v2TokenPayload)
		assignedAttestation.ID = msgIDBytes
		assignedAttestations[idx] = assignedAttestation
	}
	return assignedAttestations
}
