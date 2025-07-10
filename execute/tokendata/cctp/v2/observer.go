package v2

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// exectypes and cciptypes
// type TokenDataObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData
// type ChainSelector uint64
// type SeqNum uint64
//type MessageTokenData struct {
//	TokenData []TokenData
//}
//type TokenData struct {
//	Ready bool            `json:"ready"`
//	Data  cciptypes.Bytes `json:"data"`
//	// Error and Supported are used only for internal processing, we don't want nodes to gossip about the
//	// internals they see during processing
//	Error     error `json:"-"`
//	Supported bool  `json:"-"`
//}
// type MessageObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message
//type Message struct {
//	// Header is the family-agnostic header for OnRamp and OffRamp messages.
//	// This is always set on all CCIP messages.
//	Header RampMessageHeader `json:"header"`
//	// Sender address on the source chain.
//	// i.e if the source chain is EVM, this is an abi-encoded EVM address.
//	Sender UnknownAddress `json:"sender"`
//	// Data is the arbitrary data payload supplied by the message sender.
//	Data Bytes `json:"data"`
//	// Receiver is the receiver address on the destination chain.
//	// This is encoded in the destination chain family specific encoding.
//	// i.e if the destination is EVM, this is abi.encode(receiver).
//	Receiver UnknownAddress `json:"receiver"`
//	// ExtraArgs is destination-chain specific extra args,
//	// such as the gasLimit for EVM chains.
//	// This field is encoded in the source chain encoding scheme.
//	ExtraArgs Bytes `json:"extraArgs"`
//	// FeeToken is the fee token address.
//	// i.e if the source chain is EVM, len(FeeToken) == 20 (i.e, is not abi-encoded).
//	FeeToken UnknownAddress `json:"feeToken"`
//	// FeeTokenAmount is the amount of fee tokens paid.
//	FeeTokenAmount BigInt `json:"feeTokenAmount"`
//	// FeeValueJuels is the fee amount in Juels
//	FeeValueJuels BigInt `json:"feeValueJuels"`
//	// TokenAmounts is the array of tokens and amounts to transfer.
//	TokenAmounts []RampTokenAmount `json:"tokenAmounts"`
//}
//type RampMessageHeader struct {
//	// MessageID is a unique identifier for the message, it should be unique across all chains.
//	// It is generated on the chain that the CCIP send is requested (i.e. the source chain of a message).
//	MessageID Bytes32 `json:"messageId"`
//	// SourceChainSelector is the chain selector of the chain that the message originated from.
//	SourceChainSelector ChainSelector `json:"sourceChainSelector,string"`
//	// DestChainSelector is the chain selector of the chain that the message is destined for.
//	DestChainSelector ChainSelector `json:"destChainSelector,string"`
//	// SequenceNumber is an auto-incrementing sequence number for the message.
//	// Not unique across lanes.
//	SequenceNumber SeqNum `json:"seqNum,string"`
//	// Nonce is the nonce for this lane for this sender, not unique across senders/lanes
//	Nonce uint64 `json:"nonce"`
//
//	// MsgHash is the hash of all the message fields.
//	// NOTE: The field is expected to be empty, and will be populated by the plugin using the MsgHasher interface.
//	MsgHash Bytes32 `json:"msgHash"` // populated
//
//	// OnRamp is the address of the onramp that sent the message.
//	// NOTE: This is populated by the ccip reader. Not emitted explicitly onchain.
//	OnRamp UnknownAddress `json:"onRamp"`
//
//	// TxHash is the hash of the transaction that emitted this message.
//	TxHash string `json:"txHash"`
//}

const (
	attestationStatusComplete string = "complete"
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
		tokenDataObservations[chainSelector] = u.getTokenDataForSourceChain(ctx, chainSelector, u.attestationEncoder, chainMessages)
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

type CCTPv2Response struct {
	cctpMessages map[string]Message
	err          error
}

func initMessageTokenData(message cciptypes.Message) exectypes.MessageTokenData {
	result := make([]exectypes.TokenData, 0, len(message.TokenAmounts))
	for range message.TokenAmounts {
		result = append(result, exectypes.NotSupportedTokenData())
	}
	return exectypes.NewMessageTokenData(result...)
}

//type CCTPv2MessageOrError struct {
//	message Message
//	err     error
//}
//
//func (c CCTPv2MessageOrError) TokenData(ctx context.Context, attestationEncoder AttestationEncoder) exectypes.TokenData {
//	if c.err != nil {
//		return exectypes.NewErrorTokenData(c.err)
//	}
//	return c.message.TokenData(ctx, attestationEncoder)
//}
//
//func blah() {
//	var attestationEncoder AttestationEncoder
//	var ctx context.Context
//	var ccipMessages map[cciptypes.SeqNum]cciptypes.Message
//	var sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload
//	// convert SourceTokenDataPayload to CCTPv2 Message
//	var cctpV2Responses map[string]CCTPv2Response
//	var matchedCCTPv2Messages map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError
//
//	for seqNum, tokenPayloads := range sourceTokenDataPayloads {
//		cctpV2Response := cctpV2Responses[ccipMessages[seqNum].Header.TxHash]
//		for tokenIndex, sourceTokenData := range tokenPayloads {
//			cctpMessageOrError := findAndConsumeCCTPv2Message(cctpV2Response, sourceTokenData)
//			// TODO: enrich error
//			matchedCCTPv2Messages[seqNum][tokenIndex] = cctpMessageOrError
//		}
//	}
//
//	// Convert CCTPv2 Message to TokenData
//	var result map[cciptypes.SeqNum]exectypes.MessageTokenData
//	for seqNum, ccipMessage := range ccipMessages {
//		tokenDataList := make([]exectypes.TokenData, 0, len(ccipMessage.TokenAmounts))
//		for tokenIndex := range ccipMessage.TokenAmounts {
//			var tokenData exectypes.TokenData
//			cctpMessageOrError, ok := matchedCCTPv2Messages[seqNum][tokenIndex]
//			if !ok {
//				tokenData = exectypes.NotSupportedTokenData()
//			} else if cctpMessageOrError.err != nil {
//				tokenData = exectypes.NewErrorTokenData(cctpMessageOrError.err)
//			} else {
//				tokenData = cctpMessageOrError.message.TokenData(ctx, attestationEncoder)
//			}
//
//			tokenDataList = append(tokenDataList, tokenData)
//		}
//		result[seqNum] = exectypes.NewMessageTokenData(tokenDataList...)
//	}
//}
//
//func findAndConsumeCCTPv2Message(
//	cctpV2Response CCTPv2Response,
//	tokenPayload SourceTokenDataPayload,
//) CCTPv2MessageOrError {
//	return CCTPv2MessageOrError{}
//}

func (u *CCTPv2TokenDataObserver) getTokenDataForSourceChain(
	ctx context.Context,
	sourceChain cciptypes.ChainSelector,
	attestationEncoder AttestationEncoder,
	messages map[cciptypes.SeqNum]cciptypes.Message,
) map[cciptypes.SeqNum]exectypes.MessageTokenData {
	result := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
	for seqNum, message := range messages {
		result[seqNum] = initMessageTokenData(message)
	}

	cctpV2EnabledTokenPoolAddress, ok := u.supportedPoolsBySelector[sourceChain]
	if !ok {
		return result
	}

	var seqNumToSourceTokenDataPayloads = make(map[cciptypes.SeqNum][]*SourceTokenDataPayload)
	for seqNum, message := range messages {
		seqNumToSourceTokenDataPayloads[seqNum] =
			getCCTPv2SourceTokenDataPayloads(cctpV2EnabledTokenPoolAddress, message)
	}

	sourceDomainId, err := getSourceDomainId(sourceChain, seqNumToSourceTokenDataPayloads)
	if err != nil {
		// TODO: impl, convert non-nil sourceTokenDataPayloads to error TokenData, nil to NotSupportedTokenData
	}

	txHashes := mapset.NewSet[string]()
	// Only add tx hashes of messages that have CCTPv2 tokens
	for seqNum, payloads := range seqNumToSourceTokenDataPayloads {
		for _, payload := range payloads {
			if payload != nil {
				txHashes.Add(messages[seqNum].Header.TxHash)
				break
			}
		}
	}

	// Only call this for tx hashes that have CCTPv2 tokens
	responses := u.callCCTPv2API(ctx, sourceDomainId, txHashes)

	for seqNum, sourceTokenDataPayloads := range seqNumToSourceTokenDataPayloads {
		txHash := messages[seqNum].Header.TxHash
		result[seqNum] = convertSourceTokenPayloadsToTokenData(ctx, messages[seqNum], attestationEncoder, sourceTokenDataPayloads, responses[txHash])
	}

	return result
}

// TODO: doc
func (u *CCTPv2TokenDataObserver) callCCTPv2API(
	ctx context.Context,
	sourceDomainId uint32,
	txHashes mapset.Set[string],
) map[string]CCTPv2Response {
	// TODO: fix, this is actually txHash to Response
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

// getSourceDomainId returns the source domain ID for the provided source chain. All SourceTokenDataPayloads for the
// given source chain must have the same source domain ID. If no SourceTokenDataPayloads are found for the
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

// Iterates over the provided sourceTokenDataPayloads and for each attempt to find a matching CCTPv2 message in the
// provided cctpResponse. If a match is found, the corresponding TokenData is created using the CCTPv2 message.
// Basically this function transforms SourceTokenDataPayload -> CTTPv2 Message -> TokenData
func convertSourceTokenPayloadsToTokenData(
	ctx context.Context,
	message cciptypes.Message,
	attestationEncoder AttestationEncoder,
	sourceTokenDataPayloads []*SourceTokenDataPayload,
	cctpResponse CCTPv2Response,
) exectypes.MessageTokenData {
	result := make([]exectypes.TokenData, 0, len(sourceTokenDataPayloads))
	for tokenIndex, sourceTokenData := range sourceTokenDataPayloads {
		var tokenData exectypes.TokenData

		if sourceTokenData == nil {
			tokenData = exectypes.NotSupportedTokenData()
		} else {
			cctpMessage, err := findCctpMessage(message, tokenIndex, sourceTokenData, cctpResponse)
			if err != nil {
				tokenData = exectypes.NewErrorTokenData(err)
			} else {
				tokenData = cctpMessage.TokenData(ctx, attestationEncoder)
				// If a CCTPv2 message was found, we need to delete it from the cctpResponse to avoid
				// reusing it for other SourceTokenDataPayloads.
				delete(cctpResponse.cctpMessages, cctpMessage.EventNonce)
			}
		}

		result = append(result, tokenData)
	}

	return exectypes.NewMessageTokenData(result...)
}

// findCctpMessage searches for a matching CCTPv2 Message in cctpResponse based on the provided SourceTokenDataPayload
func findCctpMessage(
	ccipMessage cciptypes.Message,
	tokenIndex int,
	sourceTokenData *SourceTokenDataPayload,
	cctpResponse CCTPv2Response,
) (Message, error) {
	if cctpResponse.err != nil {
		return Message{}, cctpResponse.err
	}
	if sourceTokenData == nil {
		return Message{}, fmt.Errorf("sourceTokenData is nil")
	}

	for _, cctpMessage := range cctpResponse.cctpMessages {
		if sourceTokenData.matchesCctpMessage(cctpMessage) {
			return cctpMessage, nil
		}
	}

	errMsg := fmt.Sprintf(
		"A CCTPv2 USDC token transfer was made for message ID %s with token index %d, however, when fetching "+
			"the CCTP v2 messages from the CCTPv2 HTTP API for the CCIP message's transaction hash %s and source "+
			"domain ID %d, no CCTP v2 message was found that matches the CCIP token transfer. "+
			"Either we incorrectly filtered out the associated CCTPv2 message in an earlier step (which is a bug), or "+
			"the CCTP v2 API is not returning the expected message. Call the "+
			"CCTP v2 API with sourceDomainID %d and transactionHash %s to verify if there is a message with fields: %v",
		ccipMessage.Header.MessageID.String(), tokenIndex, ccipMessage.Header.TxHash, sourceTokenData.SourceDomain,
		sourceTokenData.SourceDomain, ccipMessage.Header.TxHash, sourceTokenData,
	)

	return Message{}, fmt.Errorf(errMsg)
}

// matchesCctpMessage checks if the SourceTokenDataPayload matches the provided CCTPv2 Message.
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

	// MaxFee is optional
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

// addressMatch returns true if the provided cctpAddress matches the right-aligned bytes of sourceAddress
// This is needed because the address returned by the CCTP v2 API is not padded (e.g. EVM addresses will be 20 bytes,
// Solana addresses will be 32 bytes, etc) however sourceAddress is always 32 bytes, and for EVM addresses (which are
// 20 bytes) the leftmost 12 bytes will be zero due to ABI encoding.
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
