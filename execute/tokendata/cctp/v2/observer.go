package v2

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"

	//"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
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

//// Observe TODO: doc
//func (u *CTTPv2TokenDataObserver) Observe(
//	ctx context.Context,
//	messages exectypes.MessageObservations,
//) (exectypes.TokenDataObservations, error) {
//	lggr := logutil.WithContextValues(ctx, u.lggr)
//
//	tokenDataObservations := make(exectypes.TokenDataObservations)
//	for chainSelector, chainMessages := range messages {
//		seqNum2TokenData := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
//
//		for seqNum, message := range chainMessages {
//			seqNum2TokenData[seqNum] = u.getMessageTokenData(ctx, lggr, chainSelector, message)
//		}
//
//		tokenDataObservations[chainSelector] = seqNum2TokenData
//	}
//
//	return tokenDataObservations, nil
//}

// Observe TODO: doc
func (u *CTTPv2TokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	//lggr := logutil.WithContextValues(ctx, u.lggr)

	tokenDataObservations := make(exectypes.TokenDataObservations)
	for chainSelector, chainMessages := range messages {
		tokenDataObservations[chainSelector] = u.getTokenDataForSourceChain(ctx, chainSelector, u.attestationEncoder, chainMessages)
	}

	return tokenDataObservations, nil
}

// IsTokenSupported TODO: doc
func (u *CTTPv2TokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	_, err := u.getValidatedSourceTokenDataPayload(sourceChain, msgToken)
	return err == nil
}

// TODO: doc
func (u *CTTPv2TokenDataObserver) getValidatedSourceTokenDataPayload(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) (*SourceTokenDataPayload, error) {
	if !strings.EqualFold(u.supportedPoolsBySelector[sourceChain], msgToken.SourcePoolAddress.String()) {
		return nil, fmt.Errorf("unsupported token pool address")
	}

	tokenData, err := DecodeSourceTokenDataPayload(msgToken.ExtraData)
	if err != nil {
		return nil, err
	}

	if tokenData.CCTPVersion != reader.CttpVersion2 {
		return nil, fmt.Errorf("unsupported CCTP version: %d", tokenData.CCTPVersion)
	}

	return tokenData, nil
}

func (u *CTTPv2TokenDataObserver) Close() error {
	return nil
}

//// TODO: doc, impl
//// Only call CCTP v2 API if the message has CCTP v2 tokens
//func (u *CTTPv2TokenDataObserver) getMessageTokenData(
//	ctx context.Context,
//	lggr logger.Logger,
//	sourceChain cciptypes.ChainSelector,
//	message cciptypes.Message,
//) exectypes.MessageTokenData {
//	cctpV2TokenIndexes := u.getCctpV2TokenIndexes(sourceChain, message)
//	// If there are no CCTP v2 tokens, we can return an empty token data
//	if len(cctpV2TokenIndexes) == 0 {
//		tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
//		return exectypes.NewMessageTokenData(tokenData...)
//	}
//
//	sourceDomainId, err := u.getSourceDomainId(message, cctpV2TokenIndexes)
//	if err != nil {
//		err = fmt.Errorf("failed to get source domain ID for CCIP message %s: %w", message.Header.MessageID, err)
//		return errorTokenData(lggr, message, cctpV2TokenIndexes, err)
//	}
//
//	cctpV2Messages, err := u.getRelevantCctpV2Messages(ctx, message, sourceDomainId)
//	if err != nil {
//		err = fmt.Errorf(
//			"failed to get relevant CCTPv2 messages for: CCIP message ID %s, source domain ID %d, "+
//				"transaction hash %s, error %w",
//			message.Header.MessageID, sourceDomainId, message.Header.TxHash, err,
//		)
//		return errorTokenData(lggr, message, cctpV2TokenIndexes, err)
//	}
//
//	return u.associateTokensToAttestations(ctx, cctpV2Messages, message, sourceDomainId, cctpV2TokenIndexes)
//}
//
//// TODO: doc
//func errorTokenData(
//	lggr logger.Logger,
//	message cciptypes.Message,
//	cctpV2TokenIndexes []int,
//	err error,
//) exectypes.MessageTokenData {
//	lggr.Warnf(err.Error())
//	tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
//	for _, tokenIndex := range cctpV2TokenIndexes {
//		tokenData[tokenIndex] = exectypes.NewErrorTokenData(err)
//	}
//	return exectypes.NewMessageTokenData(tokenData...)
//}
//
//// TODO: doc
//func (u *CTTPv2TokenDataObserver) associateTokensToAttestations(
//	ctx context.Context,
//	cctpV2Messages Messages,
//	message cciptypes.Message,
//	sourceDomainId uint32,
//	cctpV2TokenIndexes []int,
//) exectypes.MessageTokenData {
//	tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
//	assignedCctpMessages := make([]bool, len(cctpV2Messages.Messages))
//
//	for _, tokenIndex := range cctpV2TokenIndexes {
//		amount := message.TokenAmounts[tokenIndex].Amount.String()
//		found := false
//		for i, cctpV2msg := range cctpV2Messages.Messages {
//			if cctpV2msg.DecodedMessage.DecodedMessageBody.Amount == amount && !assignedCctpMessages[i] {
//				tokenData[tokenIndex] = cctpV2msg.TokenData(ctx, u.attestationEncoder)
//				assignedCctpMessages[i] = true
//				found = true
//				break
//			}
//		}
//
//		if !found {
//			err := missingCctpV2MessageError(u.lggr, message, sourceDomainId, tokenIndex)
//			tokenData[tokenIndex] = exectypes.NewErrorTokenData(err)
//		}
//	}
//
//	return exectypes.NewMessageTokenData(tokenData...)
//}
//
//// TODO: doc
//func missingCctpV2MessageError(
//	lggr logger.Logger,
//	message cciptypes.Message,
//	sourceDomainId uint32,
//	tokenIndex int,
//) error {
//	msg := fmt.Sprintf(
//		"A CCTPv2 USDC token transfer was made for message ID %s with token index %d, however, when fetching "+
//			"the CCTP v2 messages from the CCTPv2 HTTP API for the CCIP message's transaction hash %s and source "+
//			"domain ID %d, no CCTP v2 message was found that had the same token amount as the CCIP token transfer. "+
//			"Either we incorrectly filtered out the correct CCTPv2 message in an earlier step (which is a bug), or "+
//			"the CCTP v2 API is not returning the expected message (which is a bug in the CCTP v2 API). Call the "+
//			"CCTP v2 API with souceDomainId %d and transactionHash %s to see if there is a message with amount %s, "+
//			"to confirm if the issue is with the CCTP v2 API or with the filtering logic. One potential cause is "+
//			"that we filter CCTP v2 messages by the CCIP message's decoded receiver address, so we may have "+
//			"decoded it incorrectly. The raw/undecoded CCIP message receiver address is %s",
//		message.Header.MessageID.String(), tokenIndex, message.Header.TxHash, sourceDomainId, sourceDomainId,
//		message.Header.TxHash, message.TokenAmounts[tokenIndex].Amount.String(), message.Receiver.String())
//
//	lggr.Warnf(msg)
//	return fmt.Errorf(msg)
//}
//
//// TODO: doc, note about EVM ABI decoding
//func (u *CTTPv2TokenDataObserver) getRelevantCctpV2Messages(
//	ctx context.Context,
//	message cciptypes.Message,
//	sourceDomainId uint32,
//) (Messages, error) {
//
//	cctpV2Messages, err := u.attestationClient.GetMessages(ctx, sourceDomainId, message.Header.TxHash)
//	if err != nil {
//		return Messages{}, err
//	}
//
//	decodedReceiver, err := decodeEvmReceiver(message.Receiver)
//	if err != nil {
//		return Messages{}, fmt.Errorf("failed to EVM ABI decode receiver address: %w", err)
//	}
//
//	// filter for CCTP v2 messages that have correct receiver and are v2
//	return filterForRelevantCctpV2Messages(cctpV2Messages, decodedReceiver.String()), nil
//}
//
//// TODO: doc
//func filterForRelevantCctpV2Messages(cctpV2Messages Messages, receiver string) Messages {
//	relevantCctpV2Messages := make([]Message, 0)
//
//	for _, msg := range cctpV2Messages.Messages {
//		if msg.CCTPVersion != int(reader.CttpVersion2) {
//			continue
//		}
//
//		if msg.DecodedMessage.DecodedMessageBody.MintRecipient == receiver {
//			relevantCctpV2Messages = append(relevantCctpV2Messages, msg)
//		}
//	}
//
//	return Messages{Messages: relevantCctpV2Messages}
//}
//
//// TODO: doc
//func decodeEvmReceiver(receiver cciptypes.UnknownAddress) (cciptypes.Bytes, error) {
//	argType, err := abi.NewType("address", "", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	args := abi.Arguments{{Type: argType}}
//	unpacked, err := args.Unpack(receiver)
//	if err != nil {
//		return nil, err
//	}
//
//	return unpacked[0].([]byte), nil
//}
//
//// TODO: doc
//func (u *CTTPv2TokenDataObserver) getCctpV2TokenIndexes(
//	sourceChain cciptypes.ChainSelector,
//	message cciptypes.Message,
//) []int {
//	tokenIndexes := make([]int, 0)
//	for i, tokenAmount := range message.TokenAmounts {
//		if u.IsTokenSupported(sourceChain, tokenAmount) {
//			tokenIndexes = append(tokenIndexes, i)
//		}
//	}
//
//	return tokenIndexes
//}
//
//// TODO: doc
//func (u *CTTPv2TokenDataObserver) getSourceDomainId(
//	message cciptypes.Message,
//	cctpV2TokenIndexes []int,
//) (uint32, error) {
//	if len(cctpV2TokenIndexes) == 0 {
//		return 0, fmt.Errorf("getSourceDomainId was called with an empty cctpV2TokenIndexes")
//	}
//
//	var sourceDomainId uint32
//
//	for _, tokenIndex := range cctpV2TokenIndexes {
//		tokenAmount := message.TokenAmounts[tokenIndex]
//		tokenData, err := reader.NewSourceTokenDataPayloadFromBytes(tokenAmount.ExtraData)
//		if err != nil {
//			return 0, fmt.Errorf("failed to decode token data at index %d, error: %w", tokenIndex, err)
//		}
//
//		if sourceDomainId == 0 {
//			sourceDomainId = tokenData.SourceDomain
//		} else if sourceDomainId != tokenData.SourceDomain {
//			// A message originates from a single source domain, so we should not have multiple source domain IDs
//			return 0, fmt.Errorf(
//				"multiple source domain IDs for two different CCTPv2 transfers found in the same CCIP message:"+
//					"sourceDomainIDs %d and %d",
//				sourceDomainId, tokenData.SourceDomain,
//			)
//		}
//	}
//
//	if sourceDomainId == 0 {
//		return 0, fmt.Errorf("no token data had a non-zero source domain ID")
//	}
//
//	return sourceDomainId, nil
//}

type CCTPv2Response struct {
	cctpMessages map[string]Message
	err          error
}

func (u *CTTPv2TokenDataObserver) getTokenDataForSourceChain(
	ctx context.Context,
	sourceChain cciptypes.ChainSelector,
	attestationEncoder AttestationEncoder,
	messages map[cciptypes.SeqNum]cciptypes.Message,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

	var seqNumToSourceTokenDataPayloads = make(map[cciptypes.SeqNum][]*SourceTokenDataPayload)
	for seqNum, message := range messages {
		seqNumToSourceTokenDataPayloads[seqNum] = u.getSourceTokenDataPayloads(sourceChain, message)
	}
	// If >1 source domain IDs are found, return errored TokenData
	// If no source domain ID is found, return empty TokenData
	sourceDomainId, err := getSourceDomainId(sourceChain, seqNumToSourceTokenDataPayloads)
	if err != nil {
		// TODO: impl
	}
	txHashes := mapset.NewSet[string]()
	// Only add tx hashes that have CCTPv2 tokens
	for _, message := range messages {
		txHashes.Add(message.Header.TxHash)
	}

	// Only call this for tx hashes that have CCTPv2 tokens
	responses := u.callCCTPv2API(ctx, sourceDomainId, txHashes)

	// Now associate SourceTokenDataPayload with CCTPv2Message
	for seqNum, sourceTokenDataPayloads := range seqNumToSourceTokenDataPayloads {
		txHash := messages[seqNum].Header.TxHash
		result[seqNum] = associate(ctx, attestationEncoder, sourceTokenDataPayloads, responses[txHash])
	}

	return result
}

// TODO: doc
func (u *CTTPv2TokenDataObserver) callCCTPv2API(
	ctx context.Context,
	sourceDomainId uint32,
	txHashes mapset.Set[string],
) map[string]CCTPv2Response {
	nonceToResponse := make(map[string]CCTPv2Response)
	for txHash := range txHashes.Iter() {
		cctpResponse, err := u.attestationClient.GetMessages(ctx, sourceDomainId, txHash)
		if err != nil {
			nonceToResponse[txHash] = CCTPv2Response{err: err}
			continue
		}

		if len(cctpResponse.Messages) == 0 {
			nonceToResponse[txHash] = CCTPv2Response{err: fmt.Errorf("no CCTPv2 messages found for txHash %s", txHash)}
			continue
		}

		cctpMessages := make(map[string]Message)
		for _, msg := range cctpResponse.Messages {
			cctpMessages[msg.EventNonce] = msg
		}
		nonceToResponse[txHash] = CCTPv2Response{cctpMessages: cctpMessages}
	}

	return nonceToResponse
}

// TODO: doc
func getSourceDomainId(
	sourceChain cciptypes.ChainSelector,
	seqNumToSourceTokenDataPayloads map[cciptypes.SeqNum][]*SourceTokenDataPayload,
) (uint32, error) {
	sourceDomainIdSet := false
	var sourceDomainId uint32
	for seqNum, sourceTokenDataPayloads := range seqNumToSourceTokenDataPayloads {
		for _, sourceTokenDataPayload := range sourceTokenDataPayloads {
			if sourceTokenDataPayload == nil {
				continue
			}

			if !sourceDomainIdSet {
				sourceDomainId = sourceTokenDataPayload.SourceDomain
				sourceDomainIdSet = true
			} else if sourceDomainId != sourceTokenDataPayload.SourceDomain {
				return 0, fmt.Errorf("multiple source domain IDs found for the same source chain: sourceChain %d, "+
					"sourceDomainIDs %d and %d, seqNum %d", sourceChain, sourceDomainId,
					sourceTokenDataPayload.SourceDomain, seqNum)
			}
		}
	}

	if !sourceDomainIdSet {
		return 0, fmt.Errorf("no source domain ID found for source chain %s", sourceChain)
	}

	return sourceDomainId, nil
}

// TODO: doc
func (u *CTTPv2TokenDataObserver) getSourceTokenDataPayloads(
	sourceChain cciptypes.ChainSelector,
	message cciptypes.Message,
) []*SourceTokenDataPayload {
	sourceTokenDataPayloads := make([]*SourceTokenDataPayload, 0, len(message.TokenAmounts))
	for _, tokenAmount := range message.TokenAmounts {
		sourceTokenDataPayload, err := u.getValidatedSourceTokenDataPayload(sourceChain, tokenAmount)
		if err != nil {
			sourceTokenDataPayloads = append(sourceTokenDataPayloads, nil)
		} else {
			sourceTokenDataPayloads = append(sourceTokenDataPayloads, sourceTokenDataPayload)
		}
	}

	return sourceTokenDataPayloads
}

// TODO: better name, doc
// handle cctpResponse.err != nil
func associate(
	ctx context.Context,
	attestationEncoder AttestationEncoder,
	sourceTokenDataPayloads []*SourceTokenDataPayload,
	cctpResponse CCTPv2Response,
) exectypes.MessageTokenData {
	result := make([]exectypes.TokenData, 0, len(sourceTokenDataPayloads))
	for _, sourceTokenData := range sourceTokenDataPayloads {
		var tokenData exectypes.TokenData

		if sourceTokenData == nil {
			tokenData = exectypes.NotSupportedTokenData()
		} else {
			cctpMessage, err := findCctpMessage(sourceTokenData, cctpResponse)
			if err != nil {
				tokenData = exectypes.NewErrorTokenData(err)
			} else {
				tokenData = cctpMessage.TokenData(ctx, attestationEncoder)
				// TODO: doc
				delete(cctpResponse.cctpMessages, cctpMessage.EventNonce)
			}
		}

		result = append(result, tokenData)
	}

	return exectypes.NewMessageTokenData(result...)
}

func findCctpMessage(
	sourceTokenData *SourceTokenDataPayload,
	cctpResponse CCTPv2Response,
) (Message, error) {
	if cctpResponse.err != nil {
		return Message{}, cctpResponse.err
	}

	for _, cctpMessage := range cctpResponse.cctpMessages {
		if sourceTokenData.matchesCctpMessage(cctpMessage) {
			return cctpMessage, nil
		}
	}

	return Message{}, nil
}

type SourceTokenDataPayload struct {
	Nonce                uint64
	SourceDomain         uint32
	CCTPVersion          reader.CCTPVersion
	Amount               cciptypes.BigInt
	DestinationDomain    uint32
	MintRecipient        cciptypes.Bytes32
	BurnToken            cciptypes.Bytes32
	DestinationCaller    cciptypes.Bytes32
	MaxFee               cciptypes.BigInt
	MinFinalityThreshold uint32
}

func DecodeSourceTokenDataPayload(data []byte) (*SourceTokenDataPayload, error) {
	argTypes := abi.Arguments{
		{Type: mustABIType("uint64")},
		{Type: mustABIType("uint32")},
		{Type: mustABIType("uint8")},
		{Type: mustABIType("uint256")},
		{Type: mustABIType("uint32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("uint256")},
		{Type: mustABIType("uint32")},
	}

	vals, err := argTypes.Unpack(data)
	if err != nil {
		return nil, err
	}
	if len(vals) != 10 {
		return nil, fmt.Errorf("unexpected number of unpacked values")
	}

	return &SourceTokenDataPayload{
		Nonce:                vals[0].(uint64),
		SourceDomain:         vals[1].(uint32),
		CCTPVersion:          reader.CCTPVersion(vals[2].(uint8)),
		Amount:               cciptypes.NewBigInt(vals[3].(*big.Int)),
		DestinationDomain:    vals[4].(uint32),
		MintRecipient:        vals[5].([32]byte),
		BurnToken:            vals[6].([32]byte),
		DestinationCaller:    vals[7].([32]byte),
		MaxFee:               cciptypes.NewBigInt(vals[8].(*big.Int)),
		MinFinalityThreshold: vals[9].(uint32),
	}, nil
}

func mustABIType(t string) abi.Type {
	typ, err := abi.NewType(t, "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}

// TODO: doc, rename
func (s *SourceTokenDataPayload) matchesCctpMessage(
	cctpMessage Message,
) bool {
	if int(s.CCTPVersion) != cctpMessage.CCTPVersion {
		return false
	}
	if fmt.Sprintf("%d", s.SourceDomain) != cctpMessage.DecodedMessage.SourceDomain {
		return false
	}
	if fmt.Sprintf("%d", s.DestinationDomain) != cctpMessage.DecodedMessage.DestinationDomain {
		return false
	}

	// MinFinalityThreshold is optional
	if cctpMessage.DecodedMessage.MinFinalityThreshold != "" &&
		fmt.Sprintf("%d", s.MinFinalityThreshold) != cctpMessage.DecodedMessage.MinFinalityThreshold {
		return false
	}

	if s.Amount.String() != cctpMessage.DecodedMessage.DecodedMessageBody.Amount {
		return false
	}

	if cctpMessage.DecodedMessage.DecodedMessageBody.MaxFee != "" &&
		s.MaxFee.String() != cctpMessage.DecodedMessage.DecodedMessageBody.MaxFee {
		return false
	}

	if !addressMatch(cctpMessage.DecodedMessage.DestinationCaller, s.DestinationCaller) {
		return false
	}

	if !addressMatch(cctpMessage.DecodedMessage.DecodedMessageBody.BurnToken, s.BurnToken) {
		return false
	}

	if !addressMatch(cctpMessage.DecodedMessage.DecodedMessageBody.MintRecipient, s.MintRecipient) {
		return false
	}

	return true
}

// TODO: doc, rename
func addressMatch(cctpAddress string, sourceAddress cciptypes.Bytes32) bool {
	cctpAddressBytes, err := hex.DecodeString(strings.TrimPrefix(cctpAddress, "0x"))
	if err != nil {
		return false
	}
	if len(cctpAddressBytes) > 32 {
		return false
	}

	return bytes.Equal(sourceAddress[32-len(cctpAddressBytes):], cctpAddressBytes)
}
