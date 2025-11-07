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

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

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
	lggr.Debugw("CCTPv2TokenDataObserver.Observe")
	result := make(exectypes.TokenDataObservations)

	for chainSelector, seqNumToMessage := range messages {
		v2TokenPayloads := getV2TokenPayloads(seqNumToMessage)
		sourceDomainID := getSourceDomainID(v2TokenPayloads)
		txHashes := getTxHashes(seqNumToMessage)
		attestations := assignAttestationsToV2TokenPayloads(
			seqNumToMessage, txHashes, sourceDomainID, v2TokenPayloads,
		)
		tokenData := o.createTokenData(ctx, chainSelector, seqNumToMessage, attestations)
		result[chainSelector] = tokenData
	}

	return result, nil
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

func getV2TokenPayloads(
	messages map[cciptypes.SeqNum]cciptypes.Message,
) map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2 {
	return nil
}

// TODO: impl
// will need to return an error
func getSourceDomainID(v2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2) uint32 {
	return 0
}

// TODO: impl
func getTxHashes(messages map[cciptypes.SeqNum]cciptypes.Message) map[string][]cciptypes.SeqNum {
	return nil
}

// TODO: impl
// Will call CCTPv2HTTPClient.GetMessages
func getCCTPv2Messages(sourceDomainID uint32, txHash string) CCTPv2Messages {
	return CCTPv2Messages{}
}

// TODO: doc
func assignAttestationsToV2TokenPayloads(
	messages map[cciptypes.SeqNum]cciptypes.Message,
	txHashToSeqNums map[string][]cciptypes.SeqNum,
	sourceDomainID uint32,
	seqNumToV2TokenPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2,
) map[cciptypes.SeqNum]map[int]tokendata.AttestationStatus {
	result := make(map[cciptypes.SeqNum]map[int]tokendata.AttestationStatus)
	for txHash, seqNums := range txHashToSeqNums {
		cctpv2Messages := getCCTPv2Messages(sourceDomainID, txHash)
		attestations := extractAttestations(cctpv2Messages)
		assignedAttestations := make(map[int]tokendata.AttestationStatus)
		for _, seqNum := range seqNums {
			v2TokenPayloads, ok := seqNumToV2TokenPayloads[seqNum]
			if !ok {
				// TODO: is this correct?
				continue
			}

			for idx, v2TokenPayload := range v2TokenPayloads {
				assignedAttestation := assignAttestationForV2TokenPayload(attestations, v2TokenPayload)
				msgID := messages[seqNum].Header.MessageID
				assignedAttestation.ID = cciptypes.Bytes(msgID[:])
				assignedAttestations[idx] = assignedAttestation
			}
			result[seqNum] = assignedAttestations
		}
	}
	return result
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

// [32]byte = depositHash
// TODO: impl
func extractAttestations(cctpV2Messages CCTPv2Messages) map[[32]byte][]tokendata.AttestationStatus {
	// filter out not complete, not v2
	return nil
}

// Need to mutate attestations
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
