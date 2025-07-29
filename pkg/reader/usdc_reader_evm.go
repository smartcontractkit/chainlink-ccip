package reader

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

type evmUSDCMessageReader struct {
	lggr           logger.Logger
	contractReader contractreader.Extended
	cctpDestDomain map[uint64]uint32
}

func (u evmUSDCMessageReader) MessagesByTokenID(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]cciptypes.Bytes, error) {
	if len(tokens) == 0 {
		return map[MessageTokenID]cciptypes.Bytes{}, nil
	}

	// 1. Extract 3rd word from the MessageSent(bytes) - it's going to be our identifier
	eventIDsByMsgTokenID, err := u.getMessageTransmitterEventIDs(dest, tokens)
	if err != nil {
		return nil, err
	}

	// 2. Query the MessageTransmitter contract for the MessageSent events based on the 3rd words.
	// We need entire MessageSent payload to use that with the Attestation API
	expressions := []query.Expression{query.Confidence(primitives.Finalized)}
	if len(eventIDsByMsgTokenID) > 0 {
		eventIDs := make([]eventID, 0, len(eventIDsByMsgTokenID))
		for _, id := range eventIDsByMsgTokenID {
			eventIDs = append(eventIDs, id)
		}

		expressions = append(expressions, query.Comparator(
			consts.CCTPMessageSentValue,
			primitives.ValueComparator{
				Value:    primitives.Any(eventIDs),
				Operator: primitives.Eq,
			}))
	}

	keyFilter, err := query.Where(
		consts.EventNameCCTPMessageSent,
		expressions...,
	)
	if err != nil {
		return nil, err
	}

	iter, err := u.contractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameCCTPMessageTransmitter,
		keyFilter,
		query.NewLimitAndSort(
			query.Limit{Count: uint64(len(eventIDsByMsgTokenID))},
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
	for tokenID, messageID := range eventIDsByMsgTokenID {
		message, ok1 := messageSentEvents[messageID]
		if !ok1 {
			// Token not available in the source chain, it should never happen at this stage
			u.lggr.Warnw("Message not found in the source chain",
				"seqNr", tokenID.SeqNr,
				"tokenIndex", tokenID.Index,
				"chainSelector", source,
			)
			continue
		}
		out[tokenID] = message
	}

	return out, nil
}

// getMessageTransmitterEventIDs extracts an event ID to be used for fetching the CCTP MessageSent event.
func (u evmUSDCMessageReader) getMessageTransmitterEventIDs(
	destChainSelector cciptypes.ChainSelector,
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]eventID, error) {
	messageTransmitterEvents := make(map[MessageTokenID]eventID)

	for id, token := range tokens {
		sourceTokenPayload, err := extractABIPayload(token.ExtraData)
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
