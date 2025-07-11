package v2

import (
	"context"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

type CCTPv2TokenDataObserver struct {
	lggr                     logger.Logger
	destChainSelector        cciptypes.ChainSelector
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
	attestationEncoder       AttestationEncoder
	attestationClient        *CCTPv2AttestationClient
}

func NewCCTPv2TokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	usdcConfig pluginconfig.USDCCCTPObserverConfig,
	attestationEncoder AttestationEncoder,
) (*CCTPv2TokenDataObserver, error) {
	attestationClient, err := NewCCTPv2AttestationClient(lggr, usdcConfig)
	if err != nil {
		return nil, fmt.Errorf("create attestation client: %w", err)
	}
	supportedPoolsBySelector := make(map[cciptypes.ChainSelector]string)
	for chainSelector, tokenConfig := range usdcConfig.Tokens {
		supportedPoolsBySelector[chainSelector] = tokenConfig.SourcePoolAddress
	}
	lggr.Infow("Created USDC Token Data Observer",
		"supportedTokenPools", supportedPoolsBySelector,
	)
	return &CCTPv2TokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		attestationClient:        attestationClient,
	}, nil
}

func InitCCTPv2TokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	attestationEncoder AttestationEncoder,
	attestationClient *CCTPv2AttestationClient,
) *CCTPv2TokenDataObserver {
	return &CCTPv2TokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		attestationClient:        attestationClient,
	}
}

// Observe TODO: doc
func (u *CCTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	tokenDataObservations := make(exectypes.TokenDataObservations)
	for chainSelector, chainMessages := range messages {
		tokenDataObservations[chainSelector] = getMessageTokenDataForSourceChain(
			ctx, u.lggr, chainSelector, chainMessages, u.supportedPoolsBySelector, u.attestationEncoder,
			u.attestationClient,
		)
	}

	return tokenDataObservations, nil
}

// IsTokenSupported TODO: doc
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

// TODO: doc
func getMessageTokenDataForSourceChain(
	ctx context.Context,
	lggr logger.Logger,
	sourceChain cciptypes.ChainSelector,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	attestationEncoder AttestationEncoder,
	attestationClient *CCTPv2AttestationClient,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	cctpV2EnabledTokenPoolAddress, ok := supportedPoolsBySelector[sourceChain]
	if !ok {
		return notSupportedMessageTokenData(ccipMessages)
	}

	sourceTokenDataPayloads := make(map[cciptypes.SeqNum]map[int]SourceTokenDataPayload)
	for seqNum, ccipMessage := range ccipMessages {
		sourceTokenDataPayloads[seqNum] = getSourceTokenDataPayloads(ccipMessage, cctpV2EnabledTokenPoolAddress)
	}

	sourceDomainId, err := getSourceDomainID(sourceChain, sourceTokenDataPayloads)
	if err != nil {
		return errorMessageTokenData(err, ccipMessages, sourceTokenDataPayloads)
	}

	var txHashes mapset.Set[string]
	txHashes = getTxHashes(sourceTokenDataPayloads, ccipMessages)

	var cctpV2Messages map[string]Message
	cctpV2Messages = getCCTPv2Messages(ctx, lggr, attestationClient, sourceDomainId, txHashes)

	var tokenIndexToCCTPv2Message map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError
	tokenIndexToCCTPv2Message = matchCCTPv2MessagesToSourceTokenDataPayloads(
		cctpV2Messages, sourceTokenDataPayloads, matchesCctpMessage)

	return convertCCTPv2MessagesToMessageTokenData(ctx, ccipMessages, tokenIndexToCCTPv2Message, attestationEncoder)
}

// TODO: doc
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

// TODO: doc
func errorMessageTokenData(
	err error,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
	sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := notSupportedMessageTokenData(ccipMessages)
	for seqNum, tokenPayloads := range sourceTokenDataPayloads {
		for tokenIndex := range tokenPayloads {
			result[seqNum].TokenData[tokenIndex] = exectypes.NewErrorTokenData(err)
		}
	}

	return result
}

type CCTPv2MessageOrError struct {
	message Message
	err     error
}

// TODO: doc
func convertCCTPv2MessagesToMessageTokenData(
	ctx context.Context,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
	tokenIndexToCCTPv2Message map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError,
	attestationEncoder AttestationEncoder,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
	for seqNum, ccipMessage := range ccipMessages {
		tokenDataList := make([]exectypes.TokenData, 0, len(ccipMessage.TokenAmounts))
		for tokenIndex := range ccipMessage.TokenAmounts {
			var tokenData exectypes.TokenData
			if tokenIndexToCCTPv2Message[seqNum] == nil {
				tokenData = exectypes.NotSupportedTokenData()
			} else if cctpMessageOrError, ok := tokenIndexToCCTPv2Message[seqNum][tokenIndex]; !ok {
				tokenData = exectypes.NotSupportedTokenData()
			} else if cctpMessageOrError.err != nil {
				tokenData = exectypes.NewErrorTokenData(cctpMessageOrError.err)
			} else {
				tokenData = cctpMessageOrError.message.TokenData(ctx, attestationEncoder)
			}

			tokenDataList = append(tokenDataList, tokenData)
		}
		result[seqNum] = exectypes.NewMessageTokenData(tokenDataList...)
	}

	return result
}

// TODO: doc
func getCCTPv2Messages(
	ctx context.Context,
	lggr logger.Logger,
	attestationClient *CCTPv2AttestationClient,
	sourceDomainId uint32,
	txHashes mapset.Set[string],
) map[string]Message {
	cctpV2Messages := make(map[string]Message)
	for txHash := range txHashes.Iter() {
		cctpResponse, err := attestationClient.GetMessages(ctx, sourceDomainId, txHash)
		if err != nil {
			lggr.Infow("Failed to get CCTPv2 messages",
				"sourceDomainId", sourceDomainId,
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

// TODO: doc
func getTxHashes(
	sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
	ccipMessages map[cciptypes.SeqNum]cciptypes.Message,
) mapset.Set[string] {
	txHashes := mapset.NewSet[string]()
	for seqNum, payloads := range sourceTokenDataPayloads {
		if len(payloads) > 0 {
			if message, exists := ccipMessages[seqNum]; exists {
				txHashes.Add(message.Header.TxHash)
			}
		}
	}

	return txHashes
}

// TODO: doc
func matchCCTPv2MessagesToSourceTokenDataPayloads(
	cctpV2Messages map[string]Message,
	sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
	isMatch func(SourceTokenDataPayload, Message) bool,
) map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError {
	matchedCCTPv2Messages := make(map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError)

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
					break
				}
			}
			if foundNonce == "" {
				matchedCCTPv2Messages[seqNum][tokenIndex] = CCTPv2MessageOrError{
					err: fmt.Errorf(
						"no CCTPv2 message found for source token data payload, seqNum: %d, tokenIndex: %d",
						seqNum, tokenIndex,
					),
				}
			} else {
				delete(cctpV2Messages, foundNonce)
			}
		}
	}

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
		} else {
			// TODO: debug log
		}
	}
	return sourceTokenDataPayloads
}

// getSourceDomainId returns the source domain ID for the provided source chain. All SourceTokenDataPayloads for the
// given source chain must have the same source domain ID. If no SourceTokenDataPayloads are found for the
func getSourceDomainID(
	sourceChain cciptypes.ChainSelector,
	seqNumToSourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload,
) (uint32, error) {
	sourceDomainIdSet := false
	var sourceDomainID uint32
	for seqNum, sourceTokenDataPayloads := range seqNumToSourceTokenDataPayloads {
		for _, sourceTokenDataPayload := range sourceTokenDataPayloads {
			if !sourceDomainIdSet {
				sourceDomainID = sourceTokenDataPayload.SourceDomain
				sourceDomainIdSet = true
			} else if sourceDomainID != sourceTokenDataPayload.SourceDomain {
				return 0, fmt.Errorf("multiple source domain IDs found for the same source chain: sourceChain %d, "+
					"sourceDomainIDs %d and %d, seqNum %d", sourceChain, sourceDomainID,
					sourceTokenDataPayload.SourceDomain, seqNum)
			}
		}
	}

	if !sourceDomainIdSet {
		return 0, fmt.Errorf("no source domain ID found for source chain %s", sourceChain)
	}

	return sourceDomainID, nil
}
