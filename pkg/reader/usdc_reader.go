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
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
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
	CCTPMessageVersion = uint32(0)
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
	contractReaders map[cciptypes.ChainSelector]types.ContractReader
	cctpDestDomain  map[uint64]uint32
	boundContracts  map[cciptypes.ChainSelector]types.BoundContract
}

type eventID [32]byte

type MessageSentEvent struct {
	Arg0 []byte
}

func (m MessageSentEvent) unpackID() (eventID, error) {
	var result [32]byte

	// Check if the data slice has at least 32 bytes
	if len(m.Arg0) < 32 {
		return result, fmt.Errorf("data slice too short, must be at least 128 bytes")
	}

	// Slice the first 32-byte segment
	copy(result[:], m.Arg0[:32])

	return result, nil
}

func NewUSDCMessageReader(
	tokensConfig map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig,
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
) (USDCMessageReader, error) {
	boundContracts := make(map[cciptypes.ChainSelector]types.BoundContract)
	for chainSelector, token := range tokensConfig {
		contract, err := bindMessageTransmitters(
			context.Background(),
			contractReaders,
			chainSelector,
			token.SourceMessageTransmitterAddr,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to bind message transmitter for chain %d: %w", chainSelector, err)
		}
		boundContracts[chainSelector] = contract
	}

	return usdcMessageReader{
		contractReaders: contractReaders,
		cctpDestDomain:  CCTPDestDomains,
		boundContracts:  boundContracts,
	}, nil
}

func (u usdcMessageReader) MessageHashes(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[exectypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[exectypes.MessageTokenID]MessageHash, error) {
	// 1. Extract 3rd word from the MessageSent(bytes) - it's going to be our identifier
	eventIDs, err := u.recreateMessageTransmitterEvents(dest, tokens)
	if err != nil {
		return nil, err
	}

	// 2. Query the MessageTransmitter contract for the MessageSent events based on the 3rd words.
	// We need entire MessageSent payload to use that with the Attestation API
	cr, ok := u.boundContracts[source]
	if !ok {
		return nil, fmt.Errorf("no contract bound for chain %d", source)
	}

	iter, err := u.contractReaders[source].QueryKey(
		ctx,
		cr,
		query.KeyFilter{
			Key:         consts.EventNameCCTPMessageSent,
			Expressions: []query.Expression{},
		},
		query.NewLimitAndSort(query.Limit{}, query.NewSortBySequence(query.Asc)),
		&MessageSentEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query onRamp: %w", err)
	}

	messageSentEvents := make(map[eventID]MessageHash)
	for _, item := range iter {
		event, ok1 := item.Data.(*MessageSentEvent)
		if !ok1 {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}
		e, err1 := event.unpackID()
		if err1 != nil {
			return nil, err1
		}
		messageSentEvents[e] = event.Arg0
	}

	// 3. This should be done by ChainReader - picking only events matching eventIDs
	out := make(map[exectypes.MessageTokenID]MessageHash)
	for tokenID, messageID := range eventIDs {
		messageHash, ok1 := messageSentEvents[messageID]
		if !ok1 {
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
) (map[exectypes.MessageTokenID]eventID, error) {
	messageTransmitterEvents := make(map[exectypes.MessageTokenID]eventID)

	for id, token := range tokens {
		sourceTokenPayload, err := NewSourceTokenDataPayloadFromBytes(token.ExtraData)
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
	readers map[cciptypes.ChainSelector]types.ContractReader,
	chainSel cciptypes.ChainSelector,
	address string,
) (types.BoundContract, error) {
	var empty types.BoundContract
	r, ok := readers[chainSel]
	if !ok {
		return empty, fmt.Errorf("no contract reader found for chain %d", chainSel)
	}

	contract := types.BoundContract{
		Address: address,
		Name:    consts.ContractNameCCTPMessageTransmitter,
	}
	if err := r.Bind(ctx, []types.BoundContract{contract}); err != nil {
		return empty,
			fmt.Errorf(
				"unable to bind %s for chain %d: %w",
				consts.ContractNameCCTPMessageTransmitter, chainSel, err,
			)
	}

	return contract, nil
}

// SourceTokenDataPayload extracts the nonce and source domain from the USDC message.
// Please see the Solidity code in USDCTokenPool to understand more details
//
//	struct SourceTokenDataPayload {
//		uint64 nonce;
//		uint32 sourceDomain;
//	}
//	return Pool.LockOrBurnOutV1({
//	   destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
//	   destPoolData: abi.encode(SourceTokenDataPayload({nonce: nonce, sourceDomain: i_localDomainIdentifier}))
//	 });
type SourceTokenDataPayload struct {
	Nonce        uint64
	SourceDomain uint32
}

func NewSourceTokenDataPayload(nonce uint64, sourceDomain uint32) *SourceTokenDataPayload {
	return &SourceTokenDataPayload{
		Nonce:        nonce,
		SourceDomain: sourceDomain,
	}
}

func NewSourceTokenDataPayloadFromBytes(extraData cciptypes.Bytes) (*SourceTokenDataPayload, error) {
	if len(extraData) < 12 {
		return nil, fmt.Errorf("extraData is too short, expected at least 12 bytes")
	}

	// Extract the nonce (first 8 bytes)
	nonce := binary.BigEndian.Uint64(extraData[:8])
	// Extract the sourceDomain (next 4 bytes)
	sourceDomain := binary.BigEndian.Uint32(extraData[8:12])

	return &SourceTokenDataPayload{
		Nonce:        nonce,
		SourceDomain: sourceDomain,
	}, nil
}

func (s SourceTokenDataPayload) ToBytes() cciptypes.Bytes {
	nonceBytes := [8]byte{}
	binary.BigEndian.PutUint64(nonceBytes[:], s.Nonce)

	sourceDomainBytes := [4]byte{}
	binary.BigEndian.PutUint32(sourceDomainBytes[:], s.SourceDomain)

	return append(nonceBytes[:], sourceDomainBytes[:]...)
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
