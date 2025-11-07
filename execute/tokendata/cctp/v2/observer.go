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
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// CCTPv2TokenDataObserver observes CCTPv2 USDC messages and fetches attestations from Circle's v2 API.
type CCTPv2TokenDataObserver struct {
	lggr                     logger.Logger
	destChainSelector        cciptypes.ChainSelector
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
	attestationEncoder       AttestationEncoder
	httpClient               CCTPv2HTTPClient
}

// CCTPv2RequestArgs contains arguments for fetching CCTP messages from Circle API.
type CCTPv2RequestArgs struct {
	SourceChain     cciptypes.ChainSelector
	SourceDomainID  uint32
	TransactionHash string
	MsgTokenIDs     []reader.MessageTokenID // Tokens belonging to this transaction
}

// FetchResult contains the result of fetching CCTP messages for a transaction.
type FetchResult struct {
	Args              CCTPv2RequestArgs
	ProcessedMessages []ProcessedCCTPMessage
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

	// Step 1: Get CCTPv2 token payloads
	v2TokenPayloads := o.getCCTPv2TokenPayloads(lggr, messages)
	if len(v2TokenPayloads) == 0 {
		lggr.Debug("no CCTPv2 messages found, skipping observation")
		return o.convertToTokenDataObservations(ctx, lggr, messages, nil), nil
	}

	// Step 2: Fetch CCTPv2 messages from Circle API (grouped by transaction)
	fetchResults, err := o.fetchCCTPv2Messages(ctx, messages, v2TokenPayloads)
	if err != nil {
		return nil, fmt.Errorf("fetch CCTP messages: %w", err)
	}

	// Step 3: Assign and validate attestations (transaction-scoped pooling)
	attestations, err := o.assignAttestations(lggr, v2TokenPayloads, fetchResults)
	if err != nil {
		return nil, fmt.Errorf("assign attestations: %w", err)
	}

	// Step 4: Convert to final output format
	return o.convertToTokenDataObservations(ctx, lggr, messages, attestations), nil
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

// processSingleToken attempts to process a single token and returns its payload if valid.
// Returns nil, nil for unsupported tokens (not an error case).
// Returns nil, error for tokens that fail to decode.
func (o *CCTPv2TokenDataObserver) processSingleToken(
	chainSelector cciptypes.ChainSelector,
	seqNum cciptypes.SeqNum,
	tokenIndex int,
	tokenAmount cciptypes.RampTokenAmount,
	lggr logger.Logger,
) (*SourceTokenDataPayloadV2, error) {
	// Check if token is from a supported pool
	if !o.IsTokenSupported(chainSelector, tokenAmount) {
		return nil, nil // Not an error, just not supported
	}

	// Try to decode CCTPv2 payload from ExtraData
	payload, err := DecodeSourceTokenDataPayloadV2(tokenAmount.ExtraData)
	if err != nil {
		lggr.Warnw(
			"Failed to decode CCTPv2 source token data",
			"chainSelector", chainSelector,
			"seqNum", seqNum,
			"tokenIndex", tokenIndex,
			"error", err,
		)
		return nil, err
	}

	return payload, nil
}

// getCCTPv2TokenPayloads identifies CCTPv2 USDC transfers within messages.
// It filters for supported USDC pool addresses and decodes their CCTPv2 payloads.
func (o *CCTPv2TokenDataObserver) getCCTPv2TokenPayloads(
	lggr logger.Logger,
	messages exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2 {
	result := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2)

	// Iterate through each chain
	for chainSelector, chainMessages := range messages {
		// Iterate through each message
		for seqNum, message := range chainMessages {
			// Process each token in the message
			for i, tokenAmount := range message.TokenAmounts {
				// Attempt to process this token
				payload, err := o.processSingleToken(chainSelector, seqNum, i, tokenAmount, lggr)
				if err != nil || payload == nil {
					// Skip unsupported tokens or tokens that failed to decode
					continue
				}

				// Lazy initialize chain map if needed
				if result[chainSelector] == nil {
					result[chainSelector] = make(map[reader.MessageTokenID]*SourceTokenDataPayloadV2)
				}

				// Store the successfully decoded payload
				msgTokenID := reader.NewMessageTokenID(seqNum, i)
				result[chainSelector][msgTokenID] = payload
			}
		}
	}

	return result
}

// fetchCCTPv2Messages fetches CCTPv2 messages from Circle's HTTP API.
// It groups transfers by transaction for efficient batch API calls, then fetches processed messages for each transaction.
// Returns FetchResults containing transaction context and processed CCTP messages.
func (o *CCTPv2TokenDataObserver) fetchCCTPv2Messages(
	ctx context.Context,
	messages exectypes.MessageObservations,
	v2TokenPayloads map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) ([]FetchResult, error) {
	// Group tokens by transaction
	reqArgs := o.getCCTPv2RequestArgs(messages, v2TokenPayloads)

	var results []FetchResult
	for _, arg := range reqArgs {
		processedMessages, err := o.httpClient.GetProcessedMessages(
			ctx,
			arg.SourceChain,
			arg.SourceDomainID,
			arg.TransactionHash,
		)
		if err != nil {
			o.lggr.Warnw(
				"Failed to fetch CCTP messages from Circle API",
				"chainSelector", arg.SourceChain,
				"sourceDomain", arg.SourceDomainID,
				"txHash", arg.TransactionHash,
				"error", err,
			)
			continue // Continue processing other transactions
		}

		o.lggr.Debugw(
			"Fetched CCTP messages from Circle API",
			"chainSelector", arg.SourceChain,
			"sourceDomain", arg.SourceDomainID,
			"txHash", arg.TransactionHash,
			"numMessages", len(processedMessages),
		)

		results = append(results, FetchResult{
			Args:             arg,
			ProcessedMessages: processedMessages,
		})
	}

	return results, nil
}

// getCCTPv2RequestArgs groups tokens by transaction and creates request arguments.
func (o *CCTPv2TokenDataObserver) getCCTPv2RequestArgs(
	messages exectypes.MessageObservations,
	v2TokenPayloads map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) []CCTPv2RequestArgs {
	// Group by (chainSelector, sourceDomain, txHash)
	type txKey struct {
		chainSelector cciptypes.ChainSelector
		sourceDomain  uint32
		txHash        string
	}

	groupedTokens := make(map[txKey][]reader.MessageTokenID)

	for chainSelector, chainPayloads := range v2TokenPayloads {
		for msgTokenID, payload := range chainPayloads {
			// Get transaction hash from the message
			message, ok := messages[chainSelector][msgTokenID.SeqNr]
			if !ok {
				o.lggr.Warnw(
					"Message not found for token payload",
					"chainSelector", chainSelector,
					"msgTokenID", msgTokenID,
				)
				continue
			}

			key := txKey{
				chainSelector: chainSelector,
				sourceDomain:  payload.SourceDomain,
				txHash:        message.Header.TxHash,
			}

			groupedTokens[key] = append(groupedTokens[key], msgTokenID)
		}
	}

	// Convert map to slice of request args
	var reqArgs []CCTPv2RequestArgs
	for key, msgTokenIDs := range groupedTokens {
		reqArgs = append(reqArgs, CCTPv2RequestArgs{
			SourceChain:     key.chainSelector,
			SourceDomainID:  key.sourceDomain,
			TransactionHash: key.txHash,
			MsgTokenIDs:     msgTokenIDs,
		})
	}

	return reqArgs
}

// assignAttestations assigns attestations using transaction-scoped pooling.
// It performs depositHash matching within each transaction using the already-processed messages.
// Returns AttestationStatus objects ready for encoding.
func (o *CCTPv2TokenDataObserver) assignAttestations(
	lggr logger.Logger,
	v2TokenPayloads map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
	fetchResults []FetchResult,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus, error) {
	result := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus)

	// Process each transaction's results
	for _, fr := range fetchResults {
		// Build depositHash map FOR THIS TRANSACTION ONLY (transaction-scoped pooling)
		depositHashToMessages := make(map[[32]byte][]ProcessedCCTPMessage)
		for _, msg := range fr.ProcessedMessages {
			depositHashToMessages[msg.DepositHash] = append(depositHashToMessages[msg.DepositHash], msg)
		}

		// Assign attestations to tokens IN THIS TRANSACTION
		for _, msgTokenID := range fr.Args.MsgTokenIDs {
			payload := v2TokenPayloads[fr.Args.SourceChain][msgTokenID]
			if payload == nil {
				lggr.Warnw(
					"Payload not found for MessageTokenID",
					"chainSelector", fr.Args.SourceChain,
					"msgTokenID", msgTokenID,
				)
				continue
			}

			// Pop attestation from THIS transaction's pool
			messages := depositHashToMessages[payload.DepositHash]
			if len(messages) == 0 {
				lggr.Warnw(
					"No matching Circle message found for depositHash",
					"chainSelector", fr.Args.SourceChain,
					"txHash", fr.Args.TransactionHash,
					"depositHash", fmt.Sprintf("%x", payload.DepositHash),
				)
				if result[fr.Args.SourceChain] == nil {
					result[fr.Args.SourceChain] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)
				}
				result[fr.Args.SourceChain][msgTokenID] = tokendata.ErrorAttestationStatus(tokendata.ErrDataMissing)
				continue
			}

			// Pop the first available message (destructive - modifies the map)
			processedMsg := messages[0]
			depositHashToMessages[payload.DepositHash] = messages[1:]

			lggr.Debugw(
				"Matched Circle message to depositHash",
				"chainSelector", fr.Args.SourceChain,
				"txHash", fr.Args.TransactionHash,
				"depositHash", fmt.Sprintf("%x", payload.DepositHash),
			)

			// Initialize chain map if needed
			if result[fr.Args.SourceChain] == nil {
				result[fr.Args.SourceChain] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)
			}

			// Create successful AttestationStatus with the processed message data
			result[fr.Args.SourceChain][msgTokenID] = tokendata.SuccessAttestationStatus(
				processedMsg.DepositHash[:],
				processedMsg.MessageBytes,
				processedMsg.AttestationBytes,
			)
		}
	}

	// Add error status for any payloads that weren't in any fetch result (API fetch failed)
	for chainSelector, chainPayloads := range v2TokenPayloads {
		if result[chainSelector] == nil {
			result[chainSelector] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)
		}
		for msgTokenID := range chainPayloads {
			if _, exists := result[chainSelector][msgTokenID]; !exists {
				result[chainSelector][msgTokenID] = tokendata.ErrorAttestationStatus(tokendata.ErrDataMissing)
			}
		}
	}

	return result, nil
}

// convertToTokenDataObservations converts attestations to final TokenDataObservations format.
// It iterates through all messages and their tokens, encoding attestations for CCTPv2 USDC tokens
// and returning NotSupportedTokenData for unsupported tokens.
func (o *CCTPv2TokenDataObserver) convertToTokenDataObservations(
	ctx context.Context,
	lggr logger.Logger,
	messages exectypes.MessageObservations,
	attestations map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus,
) exectypes.TokenDataObservations {
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

	return tokenObservations
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
