// Package v2 implements CCTPv2 token data observation for USDC cross-chain transfers.
//
// This package observes CCIP messages containing USDC tokens and fetches attestations
// from Circle's CCTP v2 API. These attestations prove that USDC was burned on the source
// chain and allow minting on the destination chain.
package v2

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)
type DepositHash = [32]byte
type TxHash = string

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
			result[chainSelector] = assignTokenData(seqNumToMessage, v2TokenPayloads, nil)
			continue
		}

		txHashes := getTxHashes(seqNumToMessage, v2TokenPayloads)
		cctpV2Messages := fetchCCTPv2Messages(ctx, lggr, o.httpClient, chainSelector, txHashes, sourceDomainID)
		tokenData := getTokenData(ctx, cctpV2Messages, o.attestationEncoder)
		result[chainSelector] = assignTokenData(seqNumToMessage, v2TokenPayloads, tokenData)
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

// getTxHashes returns the set of unique transaction hashes of each message in "messages" that
// contains a CCTPv2 USDC token transfer. These tx hashes will be used to fetch attestations from
// the CCTPv2 API, and we don't want to call the API with tx hashes that don't contain any CCTPv2
// token transfers.
func getTxHashes(
	messages map[cciptypes.SeqNum]cciptypes.Message,
	v2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2,
) mapset.Set[TxHash] {
	result := mapset.NewSet[TxHash]()

	for seqNum, _ := range v2TokenPayloads {
		if msg, ok := messages[seqNum]; ok {
			result.Add(msg.Header.TxHash)
		} else {
			// TODO: log error
		}
	}

	return result
}

// fetchCCTPv2Messages TODO: doc
func fetchCCTPv2Messages(
	ctx context.Context,
	lggr logger.Logger,
	httpClient CCTPv2HTTPClient,
	chainSelector cciptypes.ChainSelector,
	txHashes mapset.Set[TxHash],
	sourceDomainID uint32,
) map[TxHash]CCTPv2Messages {
	result := make(map[TxHash]CCTPv2Messages)
	for txHash := range txHashes.Iter() {
		cctpv2Messages, err := httpClient.GetMessages(ctx, chainSelector, sourceDomainID, txHash)
		if err != nil {
			lggr.Warnw(
				"Failed to fetch CCTPv2 messages from Circle API",
				"txHash", txHash,
				"sourceDomainID", sourceDomainID,
				"chainSelector", chainSelector,
				"error", err,
			)
			continue
		}
		result[txHash] = cctpv2Messages
	}

	return result
}

// getTokenData TODO: doc
func getTokenData(
	ctx context.Context,
	txHashToCCTPV2Messages map[TxHash]CCTPv2Messages,
	attestationEncoder AttestationEncoder,
) map[TxHash]map[DepositHash][]exectypes.TokenData {
	result := make(map[TxHash]map[DepositHash][]exectypes.TokenData)
	for txHash, cctpV2Messages := range txHashToCCTPV2Messages {
		result[txHash] = convertCCTPV2MessagesToTokenData(ctx, cctpV2Messages, attestationEncoder)
	}

	return result
}

// TODO: doc
func assignTokenData(
	seqNumToMessage map[cciptypes.SeqNum]cciptypes.Message,
	seqNumToV2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2,
	tokenDataMap map[TxHash]map[DepositHash][]exectypes.TokenData,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

	for seqNum, msg := range seqNumToMessage {
		v2TokenPayloads, ok := seqNumToV2TokenPayloads[seqNum]
		if !ok {
			result[seqNum] = makeNotSupportedTokenData(msg)
			continue
		}

		messageTokenData := make([]exectypes.TokenData, 0)
		for tokenIndex := range msg.TokenAmounts {
			v2TokenPayload, ok := v2TokenPayloads[tokenIndex]
			if !ok {
				messageTokenData = append(messageTokenData, exectypes.NotSupportedTokenData())
				continue
			}

			tokenData := lookUpTokenData(msg.Header.TxHash, v2TokenPayload.DepositHash, tokenDataMap)
			messageTokenData = append(messageTokenData, tokenData)
		}
	}

	return result
}

// TODO: doc, destructive pop
func lookUpTokenData(
	txHash TxHash,
	depositHash DepositHash,
	txHashToDepositHashToTokenDataList map[TxHash]map[DepositHash][]exectypes.TokenData,
) exectypes.TokenData {
	if depositHashToTokenDataList, ok := txHashToDepositHashToTokenDataList[txHash]; ok {
		if tokenDataList, ok := depositHashToTokenDataList[depositHash]; ok {
			if len(tokenDataList) > 0 {
				tokenData := tokenDataList[0]
				depositHashToTokenDataList[depositHash] = tokenDataList[1:]
				return tokenData
			}
		}
	}

	return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
}

// TODO: doc
func makeNotSupportedTokenData(msg cciptypes.Message) exectypes.MessageTokenData {
	tokenData := make([]exectypes.TokenData, 0)
	for _ = range msg.TokenAmounts {
		tokenData = append(tokenData, exectypes.NotSupportedTokenData())
	}

	return exectypes.MessageTokenData{TokenData: tokenData}
}

// extractAttestations processes CCTPv2 messages from Circle's API and groups them by depositHash.
// - Calculates the depositHash for each message
// - Decodes the message and attestation from hex
// - Groups attestations by depositHash
//
// Returns a map where each depositHash maps to a list of attestations that can be used
// interchangeably for tokens with the same depositHash within the transaction.
func convertCCTPV2MessagesToTokenData(
	ctx context.Context,
	cctpV2Messages CCTPv2Messages,
	attestationEncoder AttestationEncoder,
) map[DepositHash][]exectypes.TokenData {
	result := make(map[DepositHash][]exectypes.TokenData)

	for _, msg := range cctpV2Messages.Messages {
		depositHash, err := CalculateDepositHash(msg.DecodedMessage)
		if err != nil {
			// Skip messages where we can't calculate the depositHash
			// This shouldn't happen with valid API responses
			// TODO: log, metrics
			continue
		}

		tokenData := CCTPv2MessageToTokenData(ctx, msg, attestationEncoder)
		result[depositHash] = append(result[depositHash], tokenData)
	}

	return result
}

// CCTPv2MessageToTokenData TODO: doc
func CCTPv2MessageToTokenData(
	ctx context.Context,
	msg CCTPv2Message,
	attestationEncoder AttestationEncoder,
) exectypes.TokenData {
	// TODO: const
	if msg.Status != "complete" {
		// TODO: log, metrics
		return exectypes.NewErrorTokenData(tokendata.ErrNotReady)
	}

	messageBytes, err := hex.DecodeString(strings.TrimPrefix(msg.Message, "0x"))
	if err != nil {
		// TODO: log, metrics
		return exectypes.NewErrorTokenData(err)
	}

	attestationBytes, err := hex.DecodeString(strings.TrimPrefix(msg.Attestation, "0x"))
	if err != nil {
		// TODO: log, metrics
		return exectypes.NewErrorTokenData(err)
	}

	encodedData, err := attestationEncoder(ctx, messageBytes, attestationBytes)
	if err != nil {
		// TODO: log, metrics
		return exectypes.NewErrorTokenData(fmt.Errorf("unable to encode attestation: %w", err))
	}

	return exectypes.NewSuccessTokenData(encodedData)
}
