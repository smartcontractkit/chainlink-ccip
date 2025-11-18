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
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
)

const (
	// CCTPMessageStatusComplete indicates a CCTP message has been attested and is ready for processing
	CCTPMessageStatusComplete = "complete"
)

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)
type DepositHash = [32]byte
type TxHash = string

// CCTPv2RequestParams contains the args needed to call CCTPv2HTTPClient.GetMessages()
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
	calculateDepositHashFn   func(CCTPv2DecodedMessage) ([32]byte, error)
	messageToTokenDataFn     func(context.Context, logger.Logger, CCTPv2Message, AttestationEncoder) exectypes.TokenData
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
		calculateDepositHashFn:   CalculateDepositHash,
		messageToTokenDataFn:     CCTPv2MessageToTokenData,
	}
}

// Observe fetches CCTPv2 attestations for USDC tokens in the provided messages.
//
// This method implements the TokenDataObserver interface and orchestrates the complete
// CCTPv2 attestation flow:
//  1. Extracts unique CCTPv2 API request parameters from supported CCTPv2 tokens in the messages
//  2. Fetches attestations from Circle's CCTP v2 API for each unique transaction
//  3. Converts the API responses to token data format organized by deposit hash
//  4. Assigns the fetched token data back to the original messages
//
// For each token in the input messages:
//   - If NOT supported (wrong pool or invalid CCTPv2 payload): returns NotSupportedTokenData
//   - If supported but attestation not found: returns NewErrorTokenData(ErrDataMissing)
//   - If supported and attestation found: returns NewSuccessTokenData with the attestation
//
// The returned TokenDataObservations preserves the exact structure of the input messages:
// for each message, len(TokenAmounts) == len(result.TokenData).
//
// Errors from Circle's API are logged but don't fail the entire operation - other tokens
// will still be processed successfully.
func (o *CCTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	// Extract the CCTPv2 API requests that need to be made
	cctpV2RequestParams := o.getCCTPv2RequestParams(messages)

	// Execute these API requests to fetch CCTPv2Messages
	cctpV2Messages := o.makeCCTPv2Requests(ctx, cctpV2RequestParams)

	// Encode the fetched CCTPv2Messages into TokenData
	// CCTPv2Message contains Attestation and MessageBody fields, and TokenData is simply the serialization of
	// Attestation + MessageBody
	tokenData := o.convertCCTPv2MessagesToTokenData(ctx, cctpV2Messages)

	// For each token transfer in each message, assign TokenData to this token transfer
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

// getCCTPv2RequestParams extracts CCTP v2 HTTP API request parameters from observed messages.
// It iterates through all messages and their token amounts, identifies supported CCTP v2 tokens,
// decodes their payload data, and collects the parameters needed to query Circle's API.
// Ensures that the returned request params are:
// - Unique: ensures we don't make duplicate requests to the CCTPv2 API
// - Are for CCTPv2 transfers: ensures we don't make CCTPv2 API calls for tokens that are not CCTPv2 transfers
func (o *CCTPv2TokenDataObserver) getCCTPv2RequestParams(
	messages exectypes.MessageObservations,
) mapset.Set[CCTPv2RequestParams] {
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
// Note that each CCTPv2Message contains an Attestation and MessageBody, and TokenData is simply the serialization of
// (Attestation, MessageBody)
//
// It's necessary to group by DepositHash so that we can associate TokenData with a TokenAmount. Given a CCIP Message
// and one of its TokenAmount, we can look up the TokenData by:
// - parsing TokenAmount.ExtraData: tokenPayload := SourceTokenDataPayloadV2
// - constructing reqParams := CCTPv2RequestParams { Message.Header.TxHash, tokenPayload.sourceDomain }
// - look up TokenData via result[reqParams][tokenPayload.DepositHash]
// and then popping the first TokenData in the list
//
// DepositHash maps to []exectypes.TokenData instead of a single TokenData because there can be multiple
// USDC CCTPv2 transfers with identical params (e.g. same amount, receiver, etc), which means there will be multiple
// CCTPv2Messages with the same DepositHash. In this case, the attestations are fungible. For example, if we have:
//
// tokenPayloadA with DepositHashX
// tokenPayloadB with DepositHashX
// cctpV2Message1 with attestation1 and DepositHashX
// cctpV2Message2 with attestation2 and DepositHashX
//
// then attestation1 can be assigned to **either** tokenPayloadA or tokenPayloadB, and attestation2 assigned to the
// other tokenPayload.
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
			depositHash, err := o.calculateDepositHashFn(msg.DecodedMessage)
			if err != nil {
				// TODO: metrics
				o.lggr.Warnw("failed to calculate deposit hash for CCTP v2 message",
					"eventNonce", msg.EventNonce,
					"sourceDomain", msg.DecodedMessage.SourceDomain,
					"err", err,
				)
				continue
			}

			// Convert the CCTP v2 message to token data
			tokenData := o.messageToTokenDataFn(ctx, o.lggr, msg, o.attestationEncoder)

			// Append the token data to the slice for this deposit hash
			result[requestParams][depositHash] = append(result[requestParams][depositHash], tokenData)
		}
	}

	return result
}

// CCTPv2MessageToTokenData converts a CCTPv2Message from Circle's API into TokenData.
// It validates the message status, decodes the hex-encoded message and attestation fields,
// and uses the attestationEncoder to create the final token data payload.
func CCTPv2MessageToTokenData(
	ctx context.Context,
	lggr logger.Logger,
	msg CCTPv2Message,
	attestationEncoder AttestationEncoder,
) exectypes.TokenData {
	if msg.Status != CCTPMessageStatusComplete {
		lggr.Debugw("CCTPv2 message not ready",
			"status", msg.Status,
			"eventNonce", msg.EventNonce,
			"sourceDomain", msg.DecodedMessage.SourceDomain,
		)
		return exectypes.NewErrorTokenData(tokendata.ErrNotReady)
	}

	messageBytes, err := hex.DecodeString(strings.TrimPrefix(msg.Message, "0x"))
	if err != nil {
		lggr.Warnw("failed to decode CCTPv2 message hex",
			"err", err,
			"eventNonce", msg.EventNonce,
			"sourceDomain", msg.DecodedMessage.SourceDomain,
		)
		return exectypes.NewErrorTokenData(err)
	}
	if len(messageBytes) == 0 {
		lggr.Warnw("CCTPv2 message is empty",
			"eventNonce", msg.EventNonce,
			"sourceDomain", msg.DecodedMessage.SourceDomain,
		)
		return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	}

	attestationBytes, err := hex.DecodeString(strings.TrimPrefix(msg.Attestation, "0x"))
	if err != nil {
		lggr.Warnw("failed to decode CCTPv2 attestation hex",
			"err", err,
			"eventNonce", msg.EventNonce,
			"sourceDomain", msg.DecodedMessage.SourceDomain,
		)
		return exectypes.NewErrorTokenData(err)
	}
	if len(attestationBytes) == 0 {
		lggr.Warnw("CCTPv2 attestation is empty",
			"eventNonce", msg.EventNonce,
			"sourceDomain", msg.DecodedMessage.SourceDomain,
		)
		return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	}

	encodedData, err := attestationEncoder(ctx, messageBytes, attestationBytes)
	if err != nil {
		lggr.Warnw("failed to encode CCTPv2 attestation",
			"err", err,
			"eventNonce", msg.EventNonce,
			"sourceDomain", msg.DecodedMessage.SourceDomain,
		)
		return exectypes.NewErrorTokenData(fmt.Errorf("unable to encode attestation: %w", err))
	}

	return exectypes.NewSuccessTokenData(encodedData)
}

// assignSingleTokenData finds and assigns token data for a single token amount.
// It handles the complete lookup and consumption logic, returning the appropriate TokenData.
// The function mutates tokenData by removing consumed items from the slices.
// Returns:
//   - NotSupportedTokenData() if token is not supported (wrong pool or invalid ExtraData)
//   - NewErrorTokenData(ErrDataMissing) if token is supported but data not found or all consumed
//   - The first available TokenData item if found (and removes it from the slice)
func (o *CCTPv2TokenDataObserver) assignSingleTokenData(
	chainSelector cciptypes.ChainSelector,
	txHash TxHash,
	tokenAmount cciptypes.RampTokenAmount,
	tokenData map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData,
) exectypes.TokenData {
	// Check if this token is supported
	if !o.IsTokenSupported(chainSelector, tokenAmount) {
		return exectypes.NotSupportedTokenData()
	}

	// Decode the payload to get the deposit hash
	payload, err := DecodeSourceTokenDataPayloadV2(tokenAmount.ExtraData)
	if err != nil {
		return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	}

	// Build request params for lookup
	requestParams := CCTPv2RequestParams{
		chainSelector: chainSelector,
		sourceDomain:  payload.SourceDomain,
		txHash:        txHash,
	}

	// Look up the token data
	depositHashMap, found := tokenData[requestParams]
	if !found {
		return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	}

	tokenDataList, found := depositHashMap[payload.DepositHash]
	if !found || len(tokenDataList) == 0 {
		return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	}

	// Get the first available token data item
	result := tokenDataList[0]

	// Remove the consumed item from the slice
	depositHashMap[payload.DepositHash] = tokenDataList[1:]

	return result
}

// assignTokenData matches fetched token data from Circle's API back to the original messages and tokens.
// For each token in each message:
// - If NOT supported: returns NotSupportedTokenData()
// - If supported but data not found: returns NewErrorTokenData(ErrDataMissing)
// - If supported and data found: assigns the token data (each data item used only once via mutation)
// The structure is preserved: len(TokenAmounts) == len(result.TokenData) for each message.
// Note: This method mutates tokenData by removing consumed items.
func (o *CCTPv2TokenDataObserver) assignTokenData(
	messages exectypes.MessageObservations,
	tokenData map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData,
) exectypes.TokenDataObservations {
	result := make(exectypes.TokenDataObservations)

	// For each chain selector
	for chainSelector, chainMessages := range messages {
		result[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		// For each message
		for seqNum, msg := range chainMessages {
			tokenDataSlice := make([]exectypes.TokenData, 0, len(msg.TokenAmounts))

			// For each TokenAmount in this message
			for _, tokenAmount := range msg.TokenAmounts {
				// assign TokenData to this TokenAmount
				td := o.assignSingleTokenData(chainSelector, msg.Header.TxHash, tokenAmount, tokenData)
				tokenDataSlice = append(tokenDataSlice, td)
			}

			result[chainSelector][seqNum] = exectypes.NewMessageTokenData(tokenDataSlice...)
		}
	}

	return result
}
