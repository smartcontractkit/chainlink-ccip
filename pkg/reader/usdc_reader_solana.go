package reader

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

type solanaUSDCMessageReader struct {
	lggr           logger.Logger
	contractReader contractreader.Extended
}

func NewSolanaUSDCReader(
	ctx context.Context,
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended,
	addrCodec cciptypes.AddressCodec,
	chainSelector cciptypes.ChainSelector,
	bytesAddress cciptypes.UnknownAddress,
) (USDCMessageReader, error) {
	// Bind the 3rd party MessageTransmitter contract, this is where CCTP MessageSent events are emitted.
	_, err := bindReaderContract(
		ctx,
		lggr,
		contractReaders,
		chainSelector,
		consts.ContractNameUSDCTokenPool,
		bytesAddress,
		addrCodec,
	)
	if err != nil {
		return nil, err
	}

	return solanaUSDCMessageReader{
		lggr:           lggr,
		contractReader: contractReaders[chainSelector],
	}, nil
}

// getMessageTokenData extracts token data from the CCTP MessageSent event.
func (u solanaUSDCMessageReader) getMessageTokenData(
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]SourceTokenDataPayload, error) {
	messageTransmitterEvents := make(map[MessageTokenID]SourceTokenDataPayload)

	for id, token := range tokens {
		sourceTokenPayload, err := extractABIPayload(token.ExtraData)
		if err != nil {
			return nil, err
		}
		messageTransmitterEvents[id] = *sourceTokenPayload
	}
	return messageTransmitterEvents, nil
}

// SolanaCCTPUSDCMessageEvent matches the onChain CcipCctpMessageSentEvent.
type SolanaCCTPUSDCMessageEvent struct {
	Discriminator       [8]byte
	OriginalSender      [32]byte
	RemoteChainSelector uint64
	MsgTotalNonce       uint64
	EventAddress        [32]byte
	SourceDomain        uint32
	CctpNonce           uint64
	MessageSentBytes    []byte
}

// MessagesByTokenID retrieves the CCTP MessageSent events.
func (u solanaUSDCMessageReader) MessagesByTokenID(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]cciptypes.Bytes, error) {
	if len(tokens) == 0 {
		return map[MessageTokenID]cciptypes.Bytes{}, nil
	}
	u.lggr.Debugw("Searching for Solana CCTP USDC logs",
		"numExpected", len(tokens))

	// Parse the extra data field to get the CCTP nonces and source domains.
	cctpData, err := u.getMessageTokenData(tokens)
	if err != nil {
		return nil, err
	}

	// Query the token pool contract for the MessageSent event data.
	expressions := []query.Expression{}
	for _, data := range cctpData {
		// This is much more expensive than the EVM version. Rather than a
		// single ANY expression, we have separate expressions for each
		// nonce and source domain pair. This is because Solana doesn't have
		// a combined ID like EVM does.
		// TODO: optimize. CR modifier or a new field in our event.
		expressions = append(expressions, query.And(
			query.Comparator(
				consts.EventAttributeCCTPNonce,
				primitives.ValueComparator{
					Value:    data.Nonce,
					Operator: primitives.Eq,
				},
			),
			query.Comparator(
				consts.EventAttributeSourceDomain,
				primitives.ValueComparator{
					Value:    data.SourceDomain,
					Operator: primitives.Eq,
				},
			),
		))
	}

	// Parent expressions for the query.
	keyFilter, err := query.Where(
		// Using same as EVM for consistency. Only used by off-chain components so name does not have to align with on-chain
		consts.EventNameCCTPMessageSent,
		query.Confidence(primitives.Finalized),
		query.Or(expressions...),
	)
	if err != nil {
		return nil, err
	}

	iter, err := u.contractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameUSDCTokenPool,
		keyFilter,
		query.NewLimitAndSort(
			query.Limit{Count: uint64(len(cctpData))},
			query.NewSortBySequence(query.Asc),
		),
		&SolanaCCTPUSDCMessageEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("error querying contract reader for chain %d: %w", source, err)
	}

	u.lggr.Debugw("CR query results", "count", len(iter))

	// 3. Read CR responses and store the results.
	out := make(map[MessageTokenID]cciptypes.Bytes)
	for _, item := range iter {
		u.lggr.Debugw("CR query item", "item", item)
		event, ok1 := item.Data.(*SolanaCCTPUSDCMessageEvent)
		if !ok1 {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}
		u.lggr.Debugw("CR query event", "event", event)

		// This is O(n^2). We could optimize it by storing the cctpData in a map with a composite key.
		for tokenID, metadata := range cctpData {
			if metadata.Nonce == event.CctpNonce && metadata.SourceDomain == event.SourceDomain {
				out[tokenID] = event.MessageSentBytes
				u.lggr.Infow("Found CCTP event", "tokenID", tokenID, "event", event)
				break
			}
		}

		u.lggr.Warnw("Found unexpected CCTP event", "event", event)
	}

	// Check if any were missed.
	for tokenID := range tokens {
		if _, ok := out[tokenID]; !ok {
			// Token is not available in the source chain, it should never happen at this stage
			u.lggr.Warnw("Message not found in the source chain",
				"seqNr", tokenID.SeqNr,
				"tokenIndex", tokenID.Index,
				"chainSelector", source,
				"data", cctpData[tokenID],
			)
		}
	}

	return out, nil
}
