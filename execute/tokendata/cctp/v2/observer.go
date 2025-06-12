package v2

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"

	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

const (
	attestationStatusComplete string = "complete"
)

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

type CTTPv2TokenDataObserver struct {
	lggr                     logger.Logger
	destChainSelector        cciptypes.ChainSelector
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
	attestationEncoder       AttestationEncoder
	attestationClient        *CTTPv2AttestationClient
}

func NewCTTPv2TokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	usdcConfig pluginconfig.USDCCCTPObserverConfig,
	attestationEncoder AttestationEncoder,
) (*CTTPv2TokenDataObserver, error) {
	attestationClient, err := NewCTTPv2AttestationClient(lggr, usdcConfig)
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
	return &CTTPv2TokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		attestationClient:        attestationClient,
	}, nil
}

func InitCTTPv2TokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	attestationEncoder AttestationEncoder,
	attestationClient *CTTPv2AttestationClient,
) *CTTPv2TokenDataObserver {
	return &CTTPv2TokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		attestationClient:        attestationClient,
	}
}

// Observe TODO: doc
func (u *CTTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	lggr := logutil.WithContextValues(ctx, u.lggr)

	tokenDataObservations := make(exectypes.TokenDataObservations)
	for chainSelector, chainMessages := range messages {
		seqNum2TokenData := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seqNum, message := range chainMessages {
			seqNum2TokenData[seqNum] = u.getMessageTokenData(ctx, lggr, chainSelector, message)
		}

		tokenDataObservations[chainSelector] = seqNum2TokenData
	}

	return tokenDataObservations, nil
}

// IsTokenSupported TODO: doc
func (u *CTTPv2TokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	if !strings.EqualFold(u.supportedPoolsBySelector[sourceChain], msgToken.SourcePoolAddress.String()) {
		return false
	}

	tokenData, err := reader.NewSourceTokenDataPayloadFromBytes(msgToken.ExtraData)
	if err != nil {
		return false
	}

	return tokenData.CCTPVersion == reader.CttpVersion2
}

func (u *CTTPv2TokenDataObserver) Close() error {
	return nil
}

// TODO: doc, impl
// Only call CCTP v2 API if the message has CCTP v2 tokens
func (u *CTTPv2TokenDataObserver) getMessageTokenData(
	ctx context.Context,
	lggr logger.Logger,
	sourceChain cciptypes.ChainSelector,
	message cciptypes.Message,
) exectypes.MessageTokenData {
	cctpV2TokenIndexes := u.getCctpV2TokenIndexes(sourceChain, message)
	// If there are no CCTP v2 tokens, we can return an empty token data
	if len(cctpV2TokenIndexes) == 0 {
		tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
		return exectypes.NewMessageTokenData(tokenData...)
	}

	sourceDomainId, err := u.getSourceDomainId(message, cctpV2TokenIndexes)
	if err != nil {
		err = fmt.Errorf("failed to get source domain ID for CCIP message %s: %w", message.Header.MessageID, err)
		return errorTokenData(lggr, message, cctpV2TokenIndexes, err)
	}

	cctpV2Messages, err := u.getRelevantCctpV2Messages(ctx, message, sourceDomainId)
	if err != nil {
		err = fmt.Errorf(
			"failed to get relevant CCTPv2 messages for: CCIP message ID %s, source domain ID %d, "+
				"transaction hash %s, error %w",
			message.Header.MessageID, sourceDomainId, message.Header.TxHash, err,
		)
		return errorTokenData(lggr, message, cctpV2TokenIndexes, err)
	}

	return u.associateTokensToAttestations(ctx, cctpV2Messages, message, sourceDomainId, cctpV2TokenIndexes)
}

// TODO: doc
func errorTokenData(
	lggr logger.Logger,
	message cciptypes.Message,
	cctpV2TokenIndexes []int,
	err error,
) exectypes.MessageTokenData {
	lggr.Warnf(err.Error())
	tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
	for _, tokenIndex := range cctpV2TokenIndexes {
		tokenData[tokenIndex] = exectypes.NewErrorTokenData(err)
	}
	return exectypes.NewMessageTokenData(tokenData...)
}

// TODO: doc
func (u *CTTPv2TokenDataObserver) associateTokensToAttestations(
	ctx context.Context,
	cctpV2Messages Messages,
	message cciptypes.Message,
	sourceDomainId uint32,
	cctpV2TokenIndexes []int,
) exectypes.MessageTokenData {
	tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
	assignedCctpMessages := make([]bool, len(cctpV2Messages.Messages))

	for _, tokenIndex := range cctpV2TokenIndexes {
		amount := message.TokenAmounts[tokenIndex].Amount.String()
		found := false
		for i, cctpV2msg := range cctpV2Messages.Messages {
			if cctpV2msg.DecodedMessage.DecodedMessageBody.Amount == amount && !assignedCctpMessages[i] {
				tokenData[tokenIndex] = cctpV2msg.TokenData(ctx, u.attestationEncoder)
				assignedCctpMessages[i] = true
				found = true
				break
			}
		}

		if !found {
			err := missingCctpV2MessageError(u.lggr, message, sourceDomainId, tokenIndex)
			tokenData[tokenIndex] = exectypes.NewErrorTokenData(err)
		}
	}

	return exectypes.NewMessageTokenData(tokenData...)
}

// TODO: doc
func missingCctpV2MessageError(
	lggr logger.Logger,
	message cciptypes.Message,
	sourceDomainId uint32,
	tokenIndex int,
) error {
	msg := fmt.Sprintf(
		"A CCTPv2 USDC token transfer was made for message ID %s with token index %d, however, when fetching "+
			"the CCTP v2 messages from the CCTPv2 HTTP API for the CCIP message's transaction hash %s and source "+
			"domain ID %d, no CCTP v2 message was found that had the same token amount as the CCIP token transfer. "+
			"Either we incorrectly filtered out the correct CCTPv2 message in an earlier step (which is a bug), or "+
			"the CCTP v2 API is not returning the expected message (which is a bug in the CCTP v2 API). Call the "+
			"CCTP v2 API with souceDomainId %d and transactionHash %s to see if there is a message with amount %s, "+
			"to confirm if the issue is with the CCTP v2 API or with the filtering logic. One potential cause is "+
			"that we filter CCTP v2 messages by the CCIP message's decoded receiver address, so we may have "+
			"decoded it incorrectly. The raw/undecoded CCIP message receiver address is %s",
		message.Header.MessageID.String(), tokenIndex, message.Header.TxHash, sourceDomainId, sourceDomainId,
		message.Header.TxHash, message.TokenAmounts[tokenIndex].Amount.String(), message.Receiver.String())

	lggr.Warnf(msg)
	return fmt.Errorf(msg)
}

// TODO: doc, note about EVM ABI decoding
func (u *CTTPv2TokenDataObserver) getRelevantCctpV2Messages(
	ctx context.Context,
	message cciptypes.Message,
	sourceDomainId uint32,
) (Messages, error) {

	cctpV2Messages, err := u.attestationClient.GetMessages(ctx, fmt.Sprintf("%d", sourceDomainId), message.Header.TxHash)
	if err != nil {
		return Messages{}, err
	}

	decodedReceiver, err := decodeEvmReceiver(message.Receiver)
	if err != nil {
		return Messages{}, fmt.Errorf("failed to EVM ABI decode receiver address: %w", err)
	}

	// filter for CCTP v2 messages that have correct receiver and are v2
	return filterForRelevantCctpV2Messages(cctpV2Messages, decodedReceiver.String()), nil
}

// TODO: doc
func filterForRelevantCctpV2Messages(cctpV2Messages Messages, receiver string) Messages {
	relevantCctpV2Messages := make([]Message, 0)

	for _, msg := range cctpV2Messages.Messages {
		if msg.CCTPVersion != int(reader.CttpVersion2) {
			continue
		}

		if msg.DecodedMessage.DecodedMessageBody.MintRecipient == receiver {
			relevantCctpV2Messages = append(relevantCctpV2Messages, msg)
		}
	}

	return Messages{Messages: relevantCctpV2Messages}
}

// TODO: doc
func decodeEvmReceiver(receiver cciptypes.UnknownAddress) (cciptypes.Bytes, error) {
	argType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}

	args := abi.Arguments{{Type: argType}}
	unpacked, err := args.Unpack(receiver)
	if err != nil {
		return nil, err
	}

	return unpacked[0].([]byte), nil
}

// TODO: doc
func (u *CTTPv2TokenDataObserver) getCctpV2TokenIndexes(
	sourceChain cciptypes.ChainSelector,
	message cciptypes.Message,
) []int {
	tokenIndexes := make([]int, 0)
	for i, tokenAmount := range message.TokenAmounts {
		if u.IsTokenSupported(sourceChain, tokenAmount) {
			tokenIndexes = append(tokenIndexes, i)
		}
	}

	return tokenIndexes
}

// TODO: doc
func (u *CTTPv2TokenDataObserver) getSourceDomainId(
	message cciptypes.Message,
	cctpV2TokenIndexes []int,
) (uint32, error) {
	if len(cctpV2TokenIndexes) == 0 {
		return 0, fmt.Errorf("getSourceDomainId was called with an empty cctpV2TokenIndexes")
	}

	var sourceDomainId uint32

	for _, tokenIndex := range cctpV2TokenIndexes {
		tokenAmount := message.TokenAmounts[tokenIndex]
		tokenData, err := reader.NewSourceTokenDataPayloadFromBytes(tokenAmount.ExtraData)
		if err != nil {
			return 0, fmt.Errorf("failed to decode token data at index %d, error: %w", tokenIndex, err)
		}

		if sourceDomainId == 0 {
			sourceDomainId = tokenData.SourceDomain
		} else if sourceDomainId != tokenData.SourceDomain {
			// A message originates from a single source domain, so we should not have multiple source domain IDs
			return 0, fmt.Errorf(
				"multiple source domain IDs for two different CCTPv2 transfers found in the same CCIP message:"+
					"sourceDomainIDs %d and %d",
				sourceDomainId, tokenData.SourceDomain,
			)
		}
	}

	if sourceDomainId == 0 {
		return 0, fmt.Errorf("no token data had a non-zero source domain ID")
	}

	return sourceDomainId, nil
}
