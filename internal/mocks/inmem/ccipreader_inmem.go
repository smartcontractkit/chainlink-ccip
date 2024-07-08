package inmem

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

type MessagesWithMetadata struct {
	cciptypes.Message
	Executed    bool
	Destination cciptypes.ChainSelector
}

type InMemoryCCIPReader struct {
	// Reports that may be returned.
	Reports []plugintypes.CommitPluginReportWithMeta

	// Messages that may be returned.
	Messages map[cciptypes.ChainSelector][]MessagesWithMetadata

	// Dest is used implicitly in some functions
	Dest cciptypes.ChainSelector
}

func (r InMemoryCCIPReader) CommitReportsGTETimestamp(
	_ context.Context, _ cciptypes.ChainSelector, ts time.Time, limit int,
) ([]plugintypes.CommitPluginReportWithMeta, error) {
	results := slicelib.Filter(r.Reports, func(report plugintypes.CommitPluginReportWithMeta) bool {
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
		return seqNumRange.Contains(msg.Header.SequenceNumber) && msg.Destination == dest && msg.Executed
	})

	// Build executed ranges
	var ranges []cciptypes.SeqNumRange
	var currentRange *cciptypes.SeqNumRange
	for _, msg := range filtered {
		if currentRange != nil && currentRange.End()+1 == msg.Header.SequenceNumber {
			// expand current range
			currentRange.SetEnd(msg.Header.SequenceNumber)
		} else {
			if currentRange != nil {
				ranges = append(ranges, *currentRange)
			}
			// initialize new range and add it to the list
			newRange := cciptypes.NewSeqNumRange(msg.Header.SequenceNumber, msg.Header.SequenceNumber)
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
) ([]cciptypes.Message, error) {
	msgs, ok := r.Messages[chain]
	if !ok {
		// no messages for chain
		return nil, nil
	}

	filtered := slicelib.Filter(msgs, func(msg MessagesWithMetadata) bool {
		return seqNumRange.Contains(msg.Header.SequenceNumber) && msg.Destination == r.Dest
	})
	return slicelib.Map(filtered, func(msg MessagesWithMetadata) cciptypes.Message {
		return msg.Message
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

// Sync can be used to perform frequent syncing operations inside the reader implementation.
// Returns a bool indicating whether something was updated.
func (r InMemoryCCIPReader) Sync(ctx context.Context) (bool, error) {
	return false, nil
}

// Interface compatibility check.
var _ reader.CCIP = InMemoryCCIPReader{}
