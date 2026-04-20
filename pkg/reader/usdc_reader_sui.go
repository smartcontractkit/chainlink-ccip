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

type suiUSDCMessageReader struct {
	lggr           logger.Logger
	contractReader contractreader.Extended
}

// getMessageTokenData extracts token data from the CCTP MessageSent event.
func (u suiUSDCMessageReader) getMessageTokenData(
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

// SuiCCTPUSDCMessageEvent matches the onChain MessageSent event for Sui.
type SuiCCTPUSDCMessageEvent struct {
	Version           uint32
	SourceDomain      uint32
	DestinationDomain uint32
	Nonce             uint64
	Sender            []byte
	Recipient         []byte
	DestinationCaller []byte
	MessageBody       []byte
}

// MessagesByTokenID retrieves the CCTP MessageSent events for Sui.
func (u suiUSDCMessageReader) MessagesByTokenID(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[MessageTokenID]cciptypes.RampTokenAmount,
) (map[MessageTokenID]cciptypes.Bytes, error) {
	if len(tokens) == 0 {
		return map[MessageTokenID]cciptypes.Bytes{}, nil
	}
	u.lggr.Debugw("Searching for Sui CCTP USDC logs",
		"numExpected", len(tokens))

	// Parse the extra data field to get the CCTP nonces and source domains.
	cctpData, err := u.getMessageTokenData(tokens)
	if err != nil {
		return nil, err
	}

	// Query the token pool contract for the MessageSent event data.
	// Similar to Solana, we need separate expressions for each nonce and source domain pair.
	expressions := []query.Expression{}
	for _, data := range cctpData {
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
		&SuiCCTPUSDCMessageEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("error querying contract reader for chain %d: %w", source, err)
	}

	u.lggr.Debugw("CR query results", "count", len(iter))

	// Read CR responses and store the results.
	out := make(map[MessageTokenID]cciptypes.Bytes)
	for _, item := range iter {
		u.lggr.Debugw("CR query item", "item", item)
		event, ok1 := item.Data.(*SuiCCTPUSDCMessageEvent)
		if !ok1 {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}
		u.lggr.Debugw("CR query event", "event", event)

		// This is O(n^2). We could optimize it by storing the cctpData in a map with a composite key.
		for tokenID, metadata := range cctpData {
			if metadata.Nonce == event.Nonce && metadata.SourceDomain == event.SourceDomain {
				out[tokenID] = event.MessageBody
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
