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
	"fmt"
	"strings"

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
	cctpMessages, err := o.fetchCCTPv2Attestations(ctx, txGroups)
	if err != nil {
		return nil, fmt.Errorf("fetch CCTPv2 attestations: %w", err)
	}

	// Step 5: Match Circle API responses to CCIP messages using depositHash
	// Validates that attestation matches the expected message parameters
	attestations := o.matchAttestationsToMessages(cctpMessages, v2Messages)

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
	// TODO: Implement filtering and decoding logic
	// - Iterate through all messages and their TokenAmounts
	// - Check if token is supported via o.IsTokenSupported()
	// - Verify ExtraData starts with CCTP_VERSION_2_TAG
	// - Decode SourceTokenDataPayloadV2 from ExtraData
	// - Return: chainSelector -> MessageTokenID -> SourceTokenDataPayloadV2
	return make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2)
}

// extractTransactionHashes gets transaction hash for each V2 message token.
func (o *CCTPv2TokenDataObserver) extractTransactionHashes(
	messages exectypes.MessageObservations,
	v2Messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]string {
	// TODO: Implement transaction hash extraction
	// - Extract message.Header.TxHash for each message
	// - Return: chainSelector -> MessageTokenID -> txHash (string)
	return make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]string)
}

// groupMessagesByTransaction groups messages by (sourceDomain, txHash) for batch API calls.
func (o *CCTPv2TokenDataObserver) groupMessagesByTransaction(
	txHashes map[cciptypes.ChainSelector]map[reader.MessageTokenID]string,
	v2Messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID {
	// TODO: Implement grouping logic
	// - Group messages by (sourceDomain, txHash) tuple
	// - Extract sourceDomain from SourceTokenDataPayloadV2 (NO mapping needed!)
	// - Return: chainSelector -> TxKey -> []MessageTokenID
	return make(map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID)
}

// fetchCCTPv2Attestations queries Circle API for attestations.
func (o *CCTPv2TokenDataObserver) fetchCCTPv2Attestations(
	ctx context.Context,
	txGroups map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message, error) {
	// TODO: Implement Circle API calls
	// - For each chain and txHash, call o.httpClient.GetMessages(ctx, chainSel, sourceDomain, txHash)
	// - Un-group results back to individual MessageTokenID
	// - Return: chainSelector -> MessageTokenID -> CCTPv2Message
	return make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message), nil
}

// matchAttestationsToMessages validates attestations match expected messages using depositHash.
func (o *CCTPv2TokenDataObserver) matchAttestationsToMessages(
	cctpMessages map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message,
	v2Messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus {
	// TODO: Implement depositHash validation
	// - For each CCIP message, compare:
	//   Expected: v2Messages[chain][msgTokenID].depositHash
	//   Actual: Calculate depositHash from cctpMessages[chain][msgTokenID].DecodedMessage fields
	// - If match: Create successful AttestationStatus
	// - If mismatch or missing: Create error AttestationStatus
	// - Return: chainSelector -> MessageTokenID -> AttestationStatus
	return make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus)
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
