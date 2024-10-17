package reader

import (
	"context"
	"encoding/binary"
	"fmt"

	sel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"

	typconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type USDCMessageReader interface {
	MessageHashes(ctx context.Context,
		source, dest cciptypes.ChainSelector,
		tokens map[MessageTokenID]cciptypes.RampTokenAmount,
	) (map[MessageTokenID]cciptypes.Bytes, error)
}

const (
	CCTPMessageVersion = uint32(0)
)

// TODO, this should be fetched from USDC Token Pool and cached
var CCTPDestDomains = map[uint64]uint32{
	sel.ETHEREUM_MAINNET.Selector:                    0,
	sel.AVALANCHE_MAINNET.Selector:                   1,
	sel.ETHEREUM_MAINNET_OPTIMISM_1.Selector:         2,
	sel.ETHEREUM_MAINNET_ARBITRUM_1.Selector:         3,
	sel.ETHEREUM_MAINNET_BASE_1.Selector:             6,
	sel.POLYGON_MAINNET.Selector:                     7,
	sel.ETHEREUM_TESTNET_SEPOLIA.Selector:            0,
	sel.AVALANCHE_TESTNET_FUJI.Selector:              1,
	sel.ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1.Selector: 2,
	sel.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector: 3,
	sel.ETHEREUM_TESTNET_SEPOLIA_BASE_1.Selector:     6,
	sel.POLYGON_TESTNET_AMOY.Selector:                7,
}

type usdcMessageReader struct {
	lggr            logger.Logger
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade
	cctpDestDomain  map[uint64]uint32
	boundContracts  map[cciptypes.ChainSelector]types.BoundContract
}

type eventID [32]byte

// MessageSentEvent represents `MessageSent(bytes)` event emitted by the MessageTransmitter contract
type MessageSentEvent struct {
	Arg0 []byte
}

func (m MessageSentEvent) unpackID() (eventID, error) {
	var result [32]byte

	// Check if the data slice has at least 32 bytes
	if len(m.Arg0) < 32 {
		return result, fmt.Errorf("data slice too short, must be at least 32 bytes")
	}

	// Slice the first 32-byte segment
	copy(result[:], m.Arg0[:32])

	return result, nil
}

func NewUSDCMessageReader(
	ctx context.Context,
	lggr logger.Logger,
	tokensConfig map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
) (USDCMessageReader, error) {
	boundContracts := make(map[cciptypes.ChainSelector]types.BoundContract)
	for chainSelector, token := range tokensConfig {
		bytesAddress, err := typconv.AddressStringToBytes(token.SourceMessageTransmitterAddr, uint64(chainSelector))
		if err != nil {
			return nil, err
		}

		contract, err := bindFacadeReaderContract(
			ctx,
			contractReaders,
			chainSelector,
			consts.ContractNameCCTPMessageTransmitter,
			bytesAddress,
		)
		if err != nil {
			return nil, err
		}
		boundContracts[chainSelector] = contract
	}

	return usdcMessageReader{
		lggr:            lggr,
		contractReaders: contractReaders,
		cctpDestDomain:  CCTPDestDomains,
		boundContracts:  boundContracts,
	}, nil
}

func (u usdcMessageReader) MessageHashes(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]cciptypes.Bytes, error) {
	if len(tokens) == 0 {
		return map[MessageTokenID]cciptypes.Bytes{}, nil
	}

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

	eventFilter := make([]query.Expression, 0, len(eventIDs))
	for _, id := range eventIDs {
		eventFilter = append(
			eventFilter,
			query.Comparator(
				consts.CCTPMessageSentValue,
				primitives.ValueComparator{
					Value:    id,
					Operator: primitives.Eq,
				}),
		)
	}

	keyFilter, err := query.Where(
		consts.EventNameCCTPMessageSent,
		query.Or(eventFilter...),
		query.Confidence(primitives.Finalized),
	)
	if err != nil {
		return nil, err
	}

	iter, err := u.contractReaders[source].QueryKey(
		ctx,
		cr,
		keyFilter,
		query.NewLimitAndSort(
			query.Limit{Count: uint64(len(eventIDs))},
			query.NewSortBySequence(query.Asc),
		),
		&MessageSentEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("error querying contract reader for chain %d: %w", source, err)
	}

	messageSentEvents := make(map[eventID]cciptypes.Bytes)
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

	// 3. Remapping database events to the proper MessageTokenID
	out := make(map[MessageTokenID]cciptypes.Bytes)
	for tokenID, messageID := range eventIDs {
		messageHash, ok1 := messageSentEvents[messageID]
		if !ok1 {
			// Token not available in the source chain, it should never happen at this stage
			u.lggr.Warnw("Message not found in the source chain",
				"seqNr", tokenID.SeqNr,
				"tokenIndex", tokenID.Index,
				"chainSelector", source,
			)
			continue
		}
		out[tokenID] = messageHash
	}

	return out, nil
}

func (u usdcMessageReader) recreateMessageTransmitterEvents(
	destChainSelector cciptypes.ChainSelector,
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]eventID, error) {
	messageTransmitterEvents := make(map[MessageTokenID]eventID)

	for id, token := range tokens {
		sourceTokenPayload, err := NewSourceTokenDataPayloadFromBytes(token.ExtraData)
		if err != nil {
			return nil, err
		}

		destDomain, ok := u.cctpDestDomain[uint64(destChainSelector)]
		if !ok {
			return nil, fmt.Errorf("destination domain not found for chain %d", destChainSelector)
		}

		//nolint:lll
		// USDC message payload:
		// uint32 _msgVersion,
		// uint32 _msgSourceDomain,
		// uint32 _msgDestinationDomain,
		// uint64 _msgNonce,
		// bytes32 _msgSender,
		// Since it's packed, all of these values contribute to the first slot
		// https://github.com/circlefin/evm-cctp-contracts/blob/377c9bd813fb86a42d900ae4003599d82aef635a/src/MessageTransmitter.sol#L41
		// https://github.com/circlefin/evm-cctp-contracts/blob/377c9bd813fb86a42d900ae4003599d82aef635a/src/MessageTransmitter.sol#L365
		var buf []byte
		buf = binary.BigEndian.AppendUint32(buf, CCTPMessageVersion)
		buf = binary.BigEndian.AppendUint32(buf, sourceTokenPayload.SourceDomain)
		buf = binary.BigEndian.AppendUint32(buf, destDomain)
		buf = binary.BigEndian.AppendUint64(buf, sourceTokenPayload.Nonce)
		// First 12 bytes of the sender address are always empty for EVM
		senderBytes := [12]byte{}
		buf = append(buf, senderBytes[:]...)

		messageTransmitterEvents[id] = [32]byte(buf[:32])
	}
	return messageTransmitterEvents, nil
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
	Messages map[MessageTokenID]cciptypes.Bytes
}

func NewFakeUSDCMessageReader(messages map[MessageTokenID]cciptypes.Bytes) FakeUSDCMessageReader {
	return FakeUSDCMessageReader{Messages: messages}
}

func (f FakeUSDCMessageReader) MessageHashes(
	_ context.Context,
	_, _ cciptypes.ChainSelector,
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]cciptypes.Bytes, error) {
	outcome := make(map[MessageTokenID]cciptypes.Bytes)
	for tokenID := range tokens {
		outcome[tokenID] = f.Messages[tokenID]
	}
	return outcome, nil
}
