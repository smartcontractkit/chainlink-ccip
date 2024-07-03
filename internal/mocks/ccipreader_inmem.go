package mocks

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type MessagesWithMetadata struct {
	cciptypes.CCIPMsg
	Executed    bool
	Destination cciptypes.ChainSelector
}

type InMemoryCCIPReader struct {
	// Reports that may be returned.
	Reports []cciptypes.CommitPluginReportWithMeta

	// Messages that may be returned.
	Messages map[cciptypes.ChainSelector][]MessagesWithMetadata

	// Dest is used implicitly in some functions
	Dest cciptypes.ChainSelector
}

func (r InMemoryCCIPReader) CommitReportsGTETimestamp(
	_ context.Context, _ cciptypes.ChainSelector, ts time.Time, limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	results := slicelib.Filter(r.Reports, func(report cciptypes.CommitPluginReportWithMeta) bool {
		return report.Timestamp.After(ts) || report.Timestamp.Equal(ts)
	})
	if len(results) > limit {
		return results[:limit], nil
	}
	return results, nil
}

func (r InMemoryCCIPReader) ExecutedMessageRanges(
	ctx context.Context, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.SeqNumRange, error) {
	msgs, ok := r.Messages[source]
	// no messages for chain
	if !ok {
		return nil, nil
	}
	filtered := slicelib.Filter(msgs, func(msg MessagesWithMetadata) bool {
		return seqNumRange.Contains(msg.SeqNum) && msg.Destination == dest && msg.Executed
	})

	// Build executed ranges
	var ranges []cciptypes.SeqNumRange
	var currentRange *cciptypes.SeqNumRange
	for _, msg := range filtered {
		if currentRange != nil && currentRange.End()+1 == msg.SeqNum {
			// expand current range
			currentRange.SetEnd(msg.SeqNum)
		} else {
			if currentRange != nil {
				ranges = append(ranges, *currentRange)
			}
			// initialize new range and add it to the list
			newRange := cciptypes.NewSeqNumRange(msg.SeqNum, msg.SeqNum)
			currentRange = &newRange
		}
	}
	if currentRange != nil {
		ranges = append(ranges, *currentRange)
	}
	return ranges, nil
}

func (r InMemoryCCIPReader) MsgsBetweenSeqNums(
	_ context.Context, chain cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.CCIPMsg, error) {
	msgs, ok := r.Messages[chain]
	if !ok {
		// no messages for chain
		return nil, nil
	}

	filtered := slicelib.Filter(msgs, func(msg MessagesWithMetadata) bool {
		return seqNumRange.Contains(msg.SeqNum) && msg.Destination == r.Dest
	})
	return slicelib.Map(filtered, func(msg MessagesWithMetadata) cciptypes.CCIPMsg {
		return msg.CCIPMsg
	}), nil
}

func (r InMemoryCCIPReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) (seqNum []cciptypes.SeqNum, err error) {
	panic("implement me")
}

func (r InMemoryCCIPReader) GasPrices(
	ctx context.Context, chains []cciptypes.ChainSelector,
) ([]cciptypes.BigInt, error) {
	panic("implement me")
}

func (r InMemoryCCIPReader) Close(ctx context.Context) error {
	return nil
}

// Interface compatibility check.
var _ cciptypes.CCIPReader = InMemoryCCIPReader{}
