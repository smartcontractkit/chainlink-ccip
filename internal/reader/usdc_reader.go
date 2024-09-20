package reader

import (
	"context"
	"encoding/binary"
	"fmt"

	sel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type MessageHash []byte

type USDCMessageReader interface {
	MessageHashes(ctx context.Context,
		source, dest cciptypes.ChainSelector,
		tokens map[exectypes.MessageTokenID]cciptypes.RampTokenAmount,
	) (map[exectypes.MessageTokenID]MessageHash, error)
}

const (
	CCTPMessageVersion      = uint32(0)
	MessageTransmitter      = "MessageTransmitter"
	MessageTransmitterEvent = "MessageSent"
	MessageSentWordIndex    = 3
)

// TODO, this should be fetched from USDC Token Pool and cached
var CCTPDestDomains = map[uint64]uint32{
	sel.ETHEREUM_MAINNET.Selector:            0,
	sel.AVALANCHE_MAINNET.Selector:           1,
	sel.ETHEREUM_MAINNET_OPTIMISM_1.Selector: 2,
	sel.ETHEREUM_MAINNET_ARBITRUM_1.Selector: 3,
	sel.ETHEREUM_MAINNET_BASE_1.Selector:     6,
	sel.POLYGON_MAINNET.Selector:             7,
	sel.AVALANCHE_TESTNET_FUJI.Selector:      1,
}

type usdcMessageReader struct {
	config          pluginconfig.USDCCCTPObserverConfig
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended
	cctpDestDomain  map[uint64]uint32
	eventIndex      int
}

func NewUSDCMessageReader(
	config pluginconfig.USDCCCTPObserverConfig,
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended,
) (USDCMessageReader, error) {
	for chainSelector, token := range config.Tokens {
		err := bindMessageTransmitters(context.Background(), contractReaders, chainSelector, token.SourcePoolAddress)
		if err != nil {
			return nil, fmt.Errorf("unable to bind message transmitter for chain %d: %w", chainSelector, err)
		}
	}

	return usdcMessageReader{
		config:          config,
		contractReaders: contractReaders,
		cctpDestDomain:  CCTPDestDomains,
		eventIndex:      MessageSentWordIndex,
	}, nil
}

type MessageSentID [32]byte

type MessageSentEvent struct {
	payload []byte
}

func (m MessageSentEvent) UnpackID() (MessageSentID, error) {
	var result [32]byte

	// Check if the data slice has at least 96 bytes
	if len(m.payload) < 128 {
		return result, fmt.Errorf("data slice too short, must be at least 128 bytes")
	}

	// Slice the third 32-byte segment (from index 96 to 128)
	copy(result[:], m.payload[96:128])

	return result, nil
}

func (u usdcMessageReader) MessageHashes(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[exectypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[exectypes.MessageTokenID]MessageHash, error) {
	messageIDs, err := u.recreateMessageTransmitterEvents(dest, tokens)
	if err != nil {
		return nil, err
	}

	iter, err := u.contractReaders[source].ExtendedQueryKey(
		ctx,
		MessageTransmitter,
		query.KeyFilter{
			Key:         MessageTransmitterEvent,
			Expressions: []query.Expression{},
		},
		query.NewLimitAndSort(query.Limit{}, query.NewSortBySequence(query.Asc)),
		&MessageSentEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query onRamp: %w", err)
	}

	messageSentEvents := make(map[MessageSentID]MessageHash)
	for _, item := range iter {
		event, ok := item.Data.(*MessageSentEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}
		messageID, err1 := event.UnpackID()
		if err1 != nil {
			return nil, err1
		}
		messageSentEvents[messageID] = event.payload
	}

	out := make(map[exectypes.MessageTokenID]MessageHash)
	for tokenID, messageID := range messageIDs {
		messageHash, ok := messageSentEvents[messageID]
		if !ok {
			// Token not available in the source chain
			continue
		}
		out[tokenID] = messageHash
	}

	return out, nil
}

func (u usdcMessageReader) recreateMessageTransmitterEvents(
	destChainSelector cciptypes.ChainSelector,
	tokens map[exectypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[exectypes.MessageTokenID]MessageSentID, error) {
	messageTransmitterEvents := make(map[exectypes.MessageTokenID]MessageSentID)

	for id, token := range tokens {
		sourceTokenPayload, err := parseUSDCExtraData(token)
		if err != nil {
			return nil, err
		}

		// USDC message payload:
		// uint32 _msgVersion,
		// uint32 _msgSourceDomain,
		// uint32 _msgDestinationDomain,
		// uint64 _msgNonce,
		// bytes32 _msgSender,
		// Since it's packed, all of these values contribute to the first slot

		msgVersionBytes := [4]byte{}
		binary.BigEndian.PutUint32(msgVersionBytes[:], CCTPMessageVersion)

		sourceDomainBytes := [4]byte{}
		binary.BigEndian.PutUint32(sourceDomainBytes[:], sourceTokenPayload.SourceDomain)

		destDomain, ok := u.cctpDestDomain[uint64(destChainSelector)]
		if !ok {
			return nil, fmt.Errorf("destination domain not found for chain %s", destChainSelector)
		}

		destDomainBytes := [4]byte{}
		binary.BigEndian.PutUint32(destDomainBytes[:], destDomain)

		nonceBytes := [8]byte{}
		binary.BigEndian.PutUint64(nonceBytes[:], sourceTokenPayload.Nonce)

		senderBytes := [12]byte{}

		messageTransmittedEvent := [32]byte(
			append(
				append(
					append(
						append(
							msgVersionBytes[:],
							sourceDomainBytes[:]...),
						destDomainBytes[:]...),
					nonceBytes[:]...),
				senderBytes[:]...),
		)
		messageTransmitterEvents[id] = messageTransmittedEvent
	}
	return messageTransmitterEvents, nil
}

func bindMessageTransmitters(
	ctx context.Context,
	readers map[cciptypes.ChainSelector]contractreader.Extended,
	chainSel cciptypes.ChainSelector,
	address string,
) error {
	r, ok := readers[chainSel]
	if !ok {
		return fmt.Errorf("no contract reader found for chain %d", chainSel)
	}

	if err := r.Bind(ctx, []types.BoundContract{
		{
			Address: address,
			Name:    MessageTransmitter,
		},
	}); err != nil {
		return fmt.Errorf("unable to bind %s for chain %d: %w", MessageTransmitter, chainSel, err)
	}

	return nil
}

type sourceTokenDataPayload struct {
	Nonce        uint64
	SourceDomain uint32
}

// extractDetailsFromUSDCMessage extracts the nonce and source domain from the USDC message.
// Please see the Solidity code for USDCTokenPool to understand more details
//
//	struct SourceTokenDataPayload {
//		uint64 nonce;
//		uint32 sourceDomain;
//	}
//	return Pool.LockOrBurnOutV1({
//	   destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
//	   destPoolData: abi.encode(SourceTokenDataPayload({nonce: nonce, sourceDomain: i_localDomainIdentifier}))
//	 });
func parseUSDCExtraData(token cciptypes.RampTokenAmount) (*sourceTokenDataPayload, error) {
	if len(token.ExtraData) < 12 {
		return nil, fmt.Errorf("extraData is too short, expected at least 12 bytes")
	}

	// Extract the nonce (first 8 bytes)
	nonce := binary.BigEndian.Uint64(token.ExtraData[:8])
	// Extract the sourceDomain (next 4 bytes)
	sourceDomain := binary.BigEndian.Uint32(token.ExtraData[8:12])

	return &sourceTokenDataPayload{
		Nonce:        nonce,
		SourceDomain: sourceDomain,
	}, nil
}

type FakeUSDCMessageReader struct {
	Messages map[exectypes.MessageTokenID]MessageHash
}

func NewFakeUSDCMessageReader(messages map[cciptypes.SeqNum]map[int][]byte) FakeUSDCMessageReader {
	out := make(map[exectypes.MessageTokenID]MessageHash)
	for seqNum, message := range messages {
		for tokenID, hash := range message {
			out[exectypes.NewMessageTokenID(seqNum, tokenID)] = hash
		}
	}
	return FakeUSDCMessageReader{Messages: out}
}

func (f FakeUSDCMessageReader) MessageHashes(
	_ context.Context,
	_, _ cciptypes.ChainSelector,
	tokens map[exectypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[exectypes.MessageTokenID]MessageHash, error) {
	outcome := make(map[exectypes.MessageTokenID]MessageHash)
	for tokenID := range tokens {
		outcome[tokenID] = f.Messages[tokenID]
	}
	return outcome, nil
}
