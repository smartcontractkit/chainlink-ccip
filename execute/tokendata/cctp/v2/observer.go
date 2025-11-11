// Package v2 implements CCTPv2 token data observation for USDC cross-chain transfers.
//
// This package observes CCIP messages containing USDC tokens and fetches attestations
// from Circle's CCTP v2 API. These attestations prove that USDC was burned on the source
// chain and allow minting on the destination chain.
package v2

import (
	"context"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
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

func (o *CCTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	return nil, nil
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

			payload, err := DecodeSourceTokenDataPayloadV2(tokenAmount.ExtraData)
			if err != nil {
				continue
			}

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
