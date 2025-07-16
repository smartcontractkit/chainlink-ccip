// Package v2 implements CCTP v2 token data observation for Chainlink CCIP.
//
// # Overview
//
// Cross-Chain Transfer Protocol (CCTP) is Circle's protocol for transferring USDC natively across
// blockchains. CCTP v2 is the second version of this protocol that enables secure, trust-minimized
// transfer of USDC between supported chains.
//
// Attestations are cryptographic proofs that validate cross-chain token transfers. When USDC is
// burned on a source chain, Circle's attestation service provides signed attestations that can be
// used to mint the equivalent amount on the destination chain. These attestations contain the
// original message data and Circle's signature, proving the burn occurred.
//
// # File Structure
//
// This file contains the CCTPv2TokenDataObserver, which processes CCIP messages to extract CCTP v2
// token data and fetch corresponding attestations. The main components are:
//
//   - CCTPv2TokenDataObserver: Main struct that implements token data observation
//   - Helper functions for processing CCIP messages and matching them to CCTP messages
//   - Functions for fetching attestations from Circle's API
//   - Token data transformation and validation logic
//
// # Main Processing Flow
//
// The primary transformation happens in the Observe method, which processes CCIP messages through
// several stages:
//
//  1. **Token Identification**: Iterate through all CCIP messages and identify tokens that represent
//     CCTP v2 transfers. A token is considered a CCTP v2 transfer if:
//     - Its source pool address matches a configured CCTP v2-enabled pool
//     - Its ExtraData field contains a valid SourceTokenDataPayload with CCTP version 2
//     - The payload can be successfully ABI-decoded
//
//  2. **Transaction Hash Collection**: Extract unique transaction hashes from messages containing
//     CCTP v2 tokens. This batching is crucial for efficiency - multiple tokens in the same
//     transaction share the same attestation data, so we fetch attestations by transaction hash
//     to minimize API calls to Circle's attestation service.
//
//  3. **Attestation Fetching**: Query Circle's attestation API for all collected transaction hashes
//     to retrieve CCTP v2 messages and their attestations.
//
//  4. **Message Matching**: Match each SourceTokenDataPayload with its corresponding CCTP v2 Message.
//     This matching is complex because:
//     - SourceTokenDataPayload contains the original burn parameters but lacks the nonce
//     - CCTP v2 contracts don't return the nonce in events, so it's not available in the payload
//     - We must match on all available fields (amount, domains, addresses, etc.) to ensure
//     each CCIP token transfer matches the correct attestation
//     - CCTPv2 Messages are consumed after matching to prevent double-assignment
//
// 5. **Token Data Generation**: Transform matched pairs into MessageTokenData:
//   - SourceTokenDataPayload + CCTP v2 Message → MessageTokenData
//   - The message bytes and attestation are encoded together for on-chain use
//   - Unmatched tokens become error tokens, unsupported tokens remain as not-supported
//
// # On-Chain Usage
//
// The resulting MessageTokenData contains encoded message bytes and attestation data that will be
// passed to CCTP v2 contracts on the destination chain. The contracts will:
//   - Verify the attestation signature against Circle's known public keys
//   - Validate the message format and parameters
//   - Mint the specified amount of USDC to the intended recipient
//
// # Data Structure Requirements
//
// The TokenDataObservations returned by Observe must maintain strict structural consistency:
//   - For each input Message, there must be exactly one MessageTokenData in the result
//   - MessageTokenData.TokenData length must equal the corresponding Message.TokenAmounts length
//   - Array indexes must correspond (TokenData[i] relates to TokenAmounts[i])
//   - This ensures downstream processing can correlate token data with the original message tokens
//
// # Error Handling
//
// The observer implements graceful error handling:
//   - Network failures when fetching attestations are logged but don't stop processing
//   - Invalid or unmatched tokens are marked with appropriate error states
//   - Unsupported tokens (non-CCTP v2) are marked as not supported
//   - Domain ID conflicts result in error tokens for all affected messages
package v2

import (
	"context"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

// AttestationEncoder encodes attestation data for on-chain use
type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

// CCTPv2TokenDataObserver observes CCTP v2 token data and fetches attestations for cross-chain transfers
type CCTPv2TokenDataObserver struct {
	lggr                     logger.Logger
	destChainSelector        cciptypes.ChainSelector
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
	attestationEncoder       AttestationEncoder
	attestationClient        CCTPv2AttestationClient
	metricsReporter          MetricsReporter
}

// NewCCTPv2TokenDataObserver creates a new CCTP v2 token data observer
func NewCCTPv2TokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	usdcConfig pluginconfig.USDCCCTPObserverConfig,
	attestationEncoder AttestationEncoder,
) (*CCTPv2TokenDataObserver, error) {
	metricsReporter, err := NewMetricsReporter(lggr, destChainSelector)
	if err != nil {
		lggr.Errorw("Failed to create CCTP v2 metrics reporter",
			"error", err,
			"destChainSelector", destChainSelector,
		)
		// Use no-op reporter instead of failing
		metricsReporter = NewNoOpMetricsReporter()
	}

	attestationClient, err := NewCCTPv2AttestationClientHTTP(lggr, usdcConfig, metricsReporter)
	if err != nil {
		lggr.Errorw("Failed to create CCTP v2 attestation client",
			"error", err,
			"destChainSelector", destChainSelector,
		)
		return nil, fmt.Errorf("create attestation client: %w", err)
	}

	supportedPoolsBySelector := make(map[cciptypes.ChainSelector]string)
	for chainSelector, tokenConfig := range usdcConfig.Tokens {
		supportedPoolsBySelector[chainSelector] = tokenConfig.SourcePoolAddress
	}
	lggr.Infow("Created CCTPv2 Token Data Observer",
		"supportedTokenPools", supportedPoolsBySelector,
	)
	return &CCTPv2TokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		attestationClient:        attestationClient,
		metricsReporter:          metricsReporter,
	}, nil
}

// InitCCTPv2TokenDataObserver initializes a CCTP v2 token data observer with pre-configured dependencies
func InitCCTPv2TokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	attestationEncoder AttestationEncoder,
	attestationClient *CCTPv2AttestationClientHTTP,
	metricsReporter MetricsReporter,
) *CCTPv2TokenDataObserver {

	return &CCTPv2TokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		attestationClient:        attestationClient,
		metricsReporter:          metricsReporter,
	}
}

// Observe processes a set of CCIP messages and returns token data observations.
// For each source chain, it extracts CCTP v2 token data from message payloads,
// fetches corresponding attestations, and converts them to MessageTokenData.
func (u *CCTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	startTime := time.Now()

	tokenDataObservations := make(exectypes.TokenDataObservations)
	for chainSelector, chainMessages := range messages {
		chainStartTime := time.Now()

		tokenDataObservations[chainSelector] = getMessageTokenDataForSourceChain(
			ctx, u.lggr, chainSelector, chainMessages, u.supportedPoolsBySelector, u.attestationEncoder,
			u.attestationClient, u.metricsReporter,
		)

		u.metricsReporter.TrackObservationLatency(
			chainSelector, "getMessageTokenDataForSourceChain", time.Since(chainStartTime),
		)
	}

	u.metricsReporter.TrackObservationLatency(0, "observe_total", time.Since(startTime))

	return tokenDataObservations, nil
}

// IsTokenSupported checks if a token from a given source chain is supported for CCTP v2 processing.
// It verifies that the token's source pool address matches a configured pool and that
// the token data payload can be successfully decoded as CCTP v2 format.
func (u *CCTPv2TokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	_, err := getCCTPv2SourceTokenDataPayload(u.supportedPoolsBySelector[sourceChain], msgToken)
	return err == nil
}

func (u *CCTPv2TokenDataObserver) Close() error {
	return nil
}

// getMessageTokenDataForSourceChain takes a collection of CCIP Messages for a source chain, collects the messages'
// TokenAmounts that are CCTP v2 transfers, fetches the attestations for these CCTP v2 transfers, and converts them to
// MessageTokenData.
//
// Transaction hashes are collected and used to batch-fetch CCTP messages from the attestation service, as multiple
// tokens within the same transaction will share the same attestation data. The function returns a map with the same
// keys as the input ccipMessages map, ensuring each input message gets corresponding token data.
func getMessageTokenDataForSourceChain(
	ctx context.Context,
	lggr logger.Logger,
	sourceChain cciptypes.ChainSelector,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	attestationEncoder AttestationEncoder,
	attestationClient CCTPv2AttestationClient,
	metricsReporter MetricsReporter,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	// Step 1: Check if source chain has CCTP v2 support configured
	// If not, all tokens are marked as not supported
	cctpV2EnabledTokenPoolAddress, ok := supportedPoolsBySelector[sourceChain]
	if !ok {
		return notSupportedMessageTokenData(ccipMessages)
	}

	// Step 2: Extract and decode CCTP v2 token data payloads from CCIP messages
	sourceTokenDataPayloads := make(map[cciptypes.SeqNum]map[int]SourceTokenDataPayload)
	for seqNum, ccipMessage := range ccipMessages {
		sourceTokenDataPayloads[seqNum] = getSourceTokenDataPayloads(ccipMessage, cctpV2EnabledTokenPoolAddress)
	}

	// Step 3: Validate that all CCTP v2 tokens have the same source domain ID
	// This is required because we need to query Circle's API with a single domain ID
	// If tokens have different domain IDs, it indicates a configuration error
	sourceDomainID, err := getSourceDomainID(lggr, sourceChain, sourceTokenDataPayloads)
	if err != nil {
		// Return error tokens for all CCTP v2 tokens, others remain not supported
		return errorMessageTokenData(err, ccipMessages, sourceTokenDataPayloads)
	}

	// Step 4: Extract unique transaction hashes from messages with CCTP v2 tokens
	txHashes := getTxHashes(lggr, sourceTokenDataPayloads, ccipMessages)

	// Step 5: Fetch CCTP v2 messages and attestations from Circle's API
	// Query the attestation service for all collected transaction hashes
	// Each response contains the message data and attestation needed for on-chain minting
	cctpV2Messages := getCCTPv2Messages(ctx, lggr, attestationClient, sourceDomainID, txHashes)

	// Step 6: Match each SourceTokenDataPayload to its corresponding CCTP v2 Message
	tokenIndexToCCTPv2Message := matchCCTPv2MessagesToSourceTokenDataPayloads(
		lggr, cctpV2Messages, sourceTokenDataPayloads, matchesCctpMessage, sourceChain, metricsReporter)

	// Step 7: Convert matched CCTP messages to final MessageTokenData format
	return convertCCTPv2MessagesToMessageTokenData(
		ctx, ccipMessages, tokenIndexToCCTPv2Message, attestationEncoder, sourceChain, metricsReporter,
	)
}

// notSupportedMessageTokenData creates a MessageTokenData map where all tokens are marked as not supported.
// This is used when a source chain doesn't have CCTP v2 support configured.
func notSupportedMessageTokenData(
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
	for seqNum, message := range ccipMessages {
		tokenData := make([]exectypes.TokenData, 0, len(message.TokenAmounts))
		for range message.TokenAmounts {
			tokenData = append(tokenData, exectypes.NotSupportedTokenData())
		}
		result[seqNum] = exectypes.NewMessageTokenData(tokenData...)
	}

	return result
}

// errorMessageTokenData creates a MessageTokenData map where tokens with valid CCTP v2 payloads
// are marked with error status, while other tokens remain as not supported.
// This is used when an error occurs during CCTP v2 processing (e.g., domain ID conflicts).
func errorMessageTokenData(
	err error,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
	sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := notSupportedMessageTokenData(ccipMessages)
	for seqNum, tokenPayloads := range sourceTokenDataPayloads {
		msgTokenData, exists := result[seqNum]
		if !exists {
			continue
		}

		for tokenIndex := range tokenPayloads {
			if tokenIndex >= len(msgTokenData.TokenData) {
				continue
			}
			msgTokenData.TokenData[tokenIndex] = exectypes.NewErrorTokenData(err)
		}
	}

	return result
}

// CCTPv2MessageOrError represents either a successful CCTP v2 message or an error
type CCTPv2MessageOrError struct {
	message Message
	err     error
}

// convertCCTPv2MessagesToMessageTokenData converts matched CCTP v2 messages to MessageTokenData.
// For each token, it either creates successful token data (if matching message exists),
// error token data (if message has errors), or not supported token data (if no match).
func convertCCTPv2MessagesToMessageTokenData(
	ctx context.Context,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
	tokenIndexToCCTPv2Message map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError,
	attestationEncoder AttestationEncoder,
	sourceChain cciptypes.ChainSelector,
	metricsReporter MetricsReporter,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

	successCount := 0
	errorCount := 0
	notSupportedCount := 0

	for seqNum, ccipMessage := range ccipMessages {
		tokenDataList := make([]exectypes.TokenData, 0, len(ccipMessage.TokenAmounts))
		for tokenIndex := range ccipMessage.TokenAmounts {
			var tokenData exectypes.TokenData
			if tokenIndexToCCTPv2Message[seqNum] == nil {
				tokenData = exectypes.NotSupportedTokenData()
				notSupportedCount++
			} else if cctpMessageOrError, ok := tokenIndexToCCTPv2Message[seqNum][tokenIndex]; !ok {
				tokenData = exectypes.NotSupportedTokenData()
				notSupportedCount++
			} else if cctpMessageOrError.err != nil {
				tokenData = exectypes.NewErrorTokenData(cctpMessageOrError.err)
				errorCount++
			} else {
				tokenData = cctpMessageOrError.message.TokenData(ctx, attestationEncoder)
				successCount++
			}

			tokenDataList = append(tokenDataList, tokenData)
		}
		result[seqNum] = exectypes.NewMessageTokenData(tokenDataList...)
	}

	// Track token processing metrics
	metricsReporter.TrackTokenProcessed(sourceChain, "success", successCount)
	metricsReporter.TrackTokenProcessed(sourceChain, "error", errorCount)
	metricsReporter.TrackTokenProcessed(sourceChain, "not_supported", notSupportedCount)

	return result
}

// getCCTPv2Messages fetches CCTP v2 messages from the attestation service for given transaction hashes.
// It queries the attestation client for each transaction hash and collects all returned messages,
// indexed by their event nonce. Errors are logged but don't stop processing of other transactions.
func getCCTPv2Messages(
	ctx context.Context,
	lggr logger.Logger,
	attestationClient CCTPv2AttestationClient,
	sourceDomainID uint32,
	txHashes mapset.Set[string],
) map[string]Message {
	cctpV2Messages := make(map[string]Message)
	for txHash := range txHashes.Iter() {
		cctpResponse, err := attestationClient.GetMessages(ctx, sourceDomainID, txHash)

		if err != nil {
			lggr.Warnw("Failed to get CCTPv2 messages from attestation service",
				"sourceDomainID", sourceDomainID,
				"txHash", txHash,
				"error", err,
			)
		} else {
			for _, msg := range cctpResponse.Messages {
				cctpV2Messages[msg.EventNonce] = msg
			}
		}
	}

	return cctpV2Messages
}

// getTxHashes extracts unique transaction hashes from CCIP messages that contain valid CCTP v2 token data.
// It only includes transaction hashes for messages that have at least one valid CCTP v2 token payload.
func getTxHashes(
	lggr logger.Logger,
	sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
) mapset.Set[string] {
	txHashes := mapset.NewSet[string]()
	for seqNum, payloads := range sourceTokenDataPayloads {
		if len(payloads) > 0 {
			if message, exists := ccipMessages[seqNum]; exists {
				txHashes.Add(message.Header.TxHash)
			} else {
				lggr.Warnw("CCIP message not found for sequence number with token payloads",
					"seqNum", seqNum,
					"numPayloads", len(payloads),
				)
			}
		}
	}

	return txHashes
}

// matchCCTPv2MessagesToSourceTokenDataPayloads matches CCTP v2 messages to source token data payloads.
// For each token payload, it finds the corresponding CCTP message using the provided matching function.
// Messages are consumed (removed from the map) after being matched to prevent double-matching.
func matchCCTPv2MessagesToSourceTokenDataPayloads(
	lggr logger.Logger,
	cctpV2Messages map[string]Message,
	sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
	isMatch func(SourceTokenDataPayload, Message) bool,
	sourceChain cciptypes.ChainSelector,
	metricsReporter MetricsReporter,
) map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError {
	matchedCCTPv2Messages := make(map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError)

	matchedCount := 0
	unmatchedCount := 0

	for seqNum, tokenPayloads := range sourceTokenDataPayloads {
		if len(tokenPayloads) > 0 {
			matchedCCTPv2Messages[seqNum] = make(map[int]CCTPv2MessageOrError)
		}
		for tokenIndex, sourceTokenData := range tokenPayloads {
			foundNonce := ""
			for nonce, cctpMessage := range cctpV2Messages {
				if isMatch(sourceTokenData, cctpMessage) {
					matchedCCTPv2Messages[seqNum][tokenIndex] = CCTPv2MessageOrError{
						message: cctpMessage,
					}
					foundNonce = nonce
					matchedCount++
					break
				}
			}
			if foundNonce == "" {
				lggr.Warnw("No matching CCTP v2 message found for token payload",
					"seqNum", seqNum,
					"tokenIndex", tokenIndex,
					"sourceDomain", sourceTokenData.SourceDomain,
					"destinationDomain", sourceTokenData.DestinationDomain,
					"amount", sourceTokenData.Amount.String(),
					"availableMessages", len(cctpV2Messages),
				)
				matchedCCTPv2Messages[seqNum][tokenIndex] = CCTPv2MessageOrError{
					err: fmt.Errorf(
						"no CCTPv2 message found for source token data payload, seqNum: %d, tokenIndex: %d",
						seqNum, tokenIndex,
					),
				}
				unmatchedCount++
			} else {
				delete(cctpV2Messages, foundNonce)
			}
		}
	}

	// Track matching metrics
	metricsReporter.TrackMessageMatching(sourceChain, "matched", matchedCount)
	metricsReporter.TrackMessageMatching(sourceChain, "unmatched", unmatchedCount)

	return matchedCCTPv2Messages
}

// getSourceTokenDataPayloads iterates over ccipMessage.TokenAmounts and returns a map of token indexes to
// SourceTokenDataPayload if the TokenAmount's ExtraData can decode to a valid (CCTP v2) SourceTokenDataPayload.
func getSourceTokenDataPayloads(
	ccipMessage cciptypes.Message,
	cctpV2EnabledTokenPoolAddress string,
) map[int]SourceTokenDataPayload {
	sourceTokenDataPayloads := make(map[int]SourceTokenDataPayload)
	for i, tokenAmount := range ccipMessage.TokenAmounts {
		payload, err := getCCTPv2SourceTokenDataPayload(cctpV2EnabledTokenPoolAddress, tokenAmount)
		if err == nil {
			sourceTokenDataPayloads[i] = *payload
		}
	}
	return sourceTokenDataPayloads
}

// getSourceDomainId returns the source domain ID for the provided source chain. All SourceTokenDataPayloads for the
// given source chain must have the same source domain ID. If no SourceTokenDataPayloads are found for the
func getSourceDomainID(
	lggr logger.Logger,
	sourceChain cciptypes.ChainSelector,
	seqNumToSourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
) (uint32, error) {
	sourceDomainIDSet := false
	var sourceDomainID uint32
	for seqNum, sourceTokenDataPayloads := range seqNumToSourceTokenDataPayloads {
		for _, sourceTokenDataPayload := range sourceTokenDataPayloads {
			if !sourceDomainIDSet {
				sourceDomainID = sourceTokenDataPayload.SourceDomain
				sourceDomainIDSet = true
			} else if sourceDomainID != sourceTokenDataPayload.SourceDomain {
				lggr.Errorw("Multiple source domain IDs detected for single source chain",
					"sourceChain", sourceChain,
					"firstDomainID", sourceDomainID,
					"conflictingDomainID", sourceTokenDataPayload.SourceDomain,
					"seqNum", seqNum,
				)
				return 0, fmt.Errorf("multiple source domain IDs found for the same source chain: sourceChain %d, "+
					"sourceDomainIDs %d and %d, seqNum %d", sourceChain, sourceDomainID,
					sourceTokenDataPayload.SourceDomain, seqNum)
			}
		}
	}

	if !sourceDomainIDSet {
		return 0, fmt.Errorf("no source domain ID found for source chain %s", sourceChain)
	}

	return sourceDomainID, nil
}
