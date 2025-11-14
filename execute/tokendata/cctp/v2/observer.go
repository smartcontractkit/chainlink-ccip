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
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)
type DepositHash = [32]byte
type TxHash = string

type CCTPv2RequestParams struct {
	chainSelector cciptypes.ChainSelector
	sourceDomain  uint32
	txHash        TxHash
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

// Observe TODO: doc
func (o *CCTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	cctpV2RequestParams := o.getCCTPv2RequestParams(messages)
	cctpV2Messages := o.makeCCTPv2Requests(ctx, cctpV2RequestParams)
	tokenData := o.convertCCTPv2MessagesToTokenData(ctx, cctpV2Messages)
	return o.assignTokenData(messages, tokenData), nil
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

// getCCTPv2RequestParams extracts CCTP v2 request parameters from observed messages.
// It iterates through all messages and their token amounts, identifies supported CCTP v2 tokens,
// decodes their payload data, and collects the parameters needed to query Circle's API.
// TODO: mention that no duplicates requests are made, and no unnecessary requests are made
func (o *CCTPv2TokenDataObserver) getCCTPv2RequestParams(messages exectypes.MessageObservations) mapset.Set[CCTPv2RequestParams] {
	result := mapset.NewSet[CCTPv2RequestParams]()

	// Iterate over each chain selector and its messages
	for chainSelector, chainMessages := range messages {
		// Iterate over each message in the chain
		for _, msg := range chainMessages {
			// Skip messages without a transaction hash
			if msg.Header.TxHash == "" {
				continue
			}

			// Iterate over each token amount in the message
			for _, tokenAmount := range msg.TokenAmounts {
				// Check if this token is a supported CCTP v2 token
				if !o.IsTokenSupported(chainSelector, tokenAmount) {
					continue
				}

				// Decode the CCTP v2 payload from the token's ExtraData
				payload, err := DecodeSourceTokenDataPayloadV2(tokenAmount.ExtraData)
				if err != nil {
					// Skip tokens that fail to decode
					continue
				}

				// Add the request parameters to the result set
				result.Add(CCTPv2RequestParams{
					chainSelector: chainSelector,
					sourceDomain:  payload.SourceDomain,
					txHash:        msg.Header.TxHash,
				})
			}
		}
	}

	return result
}

// makeCCTPv2Requests fetches CCTP v2 messages from Circle's API for each request parameter.
// It makes HTTP calls to Circle's attestation API and maps the responses back to their request parameters.
// Errors are logged but don't fail the entire operation - successful requests are still returned.
func (o *CCTPv2TokenDataObserver) makeCCTPv2Requests(
	ctx context.Context,
	cctpV2RequestParams mapset.Set[CCTPv2RequestParams],
) map[CCTPv2RequestParams]CCTPv2Messages {
	result := make(map[CCTPv2RequestParams]CCTPv2Messages)

	// Iterate over each request parameter
	for params := range cctpV2RequestParams.Iter() {
		// Call Circle's API to get messages for this transaction
		messages, err := o.httpClient.GetMessages(ctx, params.chainSelector, params.sourceDomain, params.txHash)
		if err != nil {
			// Log error but continue processing other requests
			o.lggr.Warnw("failed to get CCTP v2 messages from Circle API",
				"chainSelector", params.chainSelector,
				"sourceDomain", params.sourceDomain,
				"txHash", params.txHash,
				"err", err,
			)
			continue
		}

		// Map the successful response to its request parameters
		result[params] = messages
	}

	return result
}

// convertCCTPv2MessagesToTokenData transforms Circle API responses into token data.
// It processes each CCTP v2 message, converts it to token data format, and organizes the results
// by request parameters and deposit hash for efficient lookup in the assignment phase.
func (o *CCTPv2TokenDataObserver) convertCCTPv2MessagesToTokenData(
	ctx context.Context,
	cctpV2Messages map[CCTPv2RequestParams]CCTPv2Messages,
) map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData {
	result := make(map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData)

	// Iterate over each batch of messages keyed by request parameters
	for requestParams, messages := range cctpV2Messages {
		// Initialize the inner map for this request params
		result[requestParams] = make(map[DepositHash][]exectypes.TokenData)

		// Process each individual CCTP v2 message in the batch
		for _, msg := range messages.Messages {
			// Calculate the deposit hash for this message
			depositHash, err := getCCTPv2MessageDepositHash(msg)
			if err != nil {
				o.lggr.Warnw("failed to calculate deposit hash for CCTP v2 message",
					"err", err,
				)
				continue
			}

			// Convert the CCTP v2 message to token data
			tokenData := CCTPv2MessageToTokenData(ctx, msg, o.attestationEncoder)

			// Append the token data to the slice for this deposit hash
			result[requestParams][depositHash] = append(result[requestParams][depositHash], tokenData)
		}
	}

	return result
}

// assignTokenData matches fetched token data from Circle's API back to the original messages and tokens.
// For each token in each message:
// - If NOT supported: returns NotSupportedTokenData()
// - If supported but data not found: returns NewErrorTokenData(ErrDataMissing)
// - If supported and data found: assigns the token data (each data item used only once)
// The structure is preserved: len(TokenAmounts) == len(result.TokenData) for each message.
func (o *CCTPv2TokenDataObserver) assignTokenData(
	messages exectypes.MessageObservations,
	tokenData map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData,
) exectypes.TokenDataObservations {
	result := make(exectypes.TokenDataObservations)

	// Track which token data items have been consumed (to enforce "use once" requirement)
	consumed := make(map[CCTPv2RequestParams]map[DepositHash]int)

	// Iterate over each chain selector and its messages
	for chainSelector, chainMessages := range messages {
		result[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		// Iterate over each message in the chain
		for seqNum, msg := range chainMessages {
			tokenDataSlice := make([]exectypes.TokenData, 0, len(msg.TokenAmounts))

			// Process each token in the message
			for _, tokenAmount := range msg.TokenAmounts {
				// Check if this token is supported
				if !o.IsTokenSupported(chainSelector, tokenAmount) {
					tokenDataSlice = append(tokenDataSlice, exectypes.NotSupportedTokenData())
					continue
				}

				// Token is supported - try to find its data
				// Decode the payload to get the deposit hash
				payload, err := DecodeSourceTokenDataPayloadV2(tokenAmount.ExtraData)
				if err != nil {
					// Failed to decode - treat as missing data
					tokenDataSlice = append(tokenDataSlice, exectypes.NewErrorTokenData(tokendata.ErrDataMissing))
					continue
				}

				// Build request params for lookup
				requestParams := CCTPv2RequestParams{
					chainSelector: chainSelector,
					sourceDomain:  payload.SourceDomain,
					txHash:        msg.Header.TxHash,
				}

				// Look up the token data
				depositHashMap, found := tokenData[requestParams]
				if !found {
					tokenDataSlice = append(tokenDataSlice, exectypes.NewErrorTokenData(tokendata.ErrDataMissing))
					continue
				}

				tokenDataList, found := depositHashMap[payload.DepositHash]
				if !found {
					tokenDataSlice = append(tokenDataSlice, exectypes.NewErrorTokenData(tokendata.ErrDataMissing))
					continue
				}

				// Initialize consumed tracking for this request params if needed
				if consumed[requestParams] == nil {
					consumed[requestParams] = make(map[DepositHash]int)
				}

				// Get the next available index for this deposit hash
				consumedIndex := consumed[requestParams][payload.DepositHash]

				// Check if we have any unused items left
				if consumedIndex >= len(tokenDataList) {
					// All items have been consumed
					tokenDataSlice = append(tokenDataSlice, exectypes.NewErrorTokenData(tokendata.ErrDataMissing))
					continue
				}

				// Use the next available token data item
				tokenDataSlice = append(tokenDataSlice, tokenDataList[consumedIndex])

				// Mark this item as consumed
				consumed[requestParams][payload.DepositHash]++
			}

			// Assign the token data slice to the result
			result[chainSelector][seqNum] = exectypes.NewMessageTokenData(tokenDataSlice...)
		}
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

// Do not impl yet
func getCCTPv2MessageDepositHash(msg CCTPv2Message) (DepositHash, error) {
	return DepositHash{}, nil
}
