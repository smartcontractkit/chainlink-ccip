package inmem

import (
	"context"
	"math/big"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type MessagesWithMetadata struct {
	cciptypes.Message
	Executed    bool
	Destination cciptypes.ChainSelector
}

type InMemoryCCIPReader struct {
	// UnfinalizedReports that may be returned.
	UnfinalizedReports []cciptypes.CommitPluginReportWithMeta
	FinalizedReports   []cciptypes.CommitPluginReportWithMeta

	// Messages that may be returned.
	Messages map[cciptypes.ChainSelector][]MessagesWithMetadata

	// Dest is used implicitly in some functions
	Dest cciptypes.ChainSelector

	ConfigDigest [32]byte
}

func (r InMemoryCCIPReader) GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error) {
	panic("not implemented")
}

// GetExpectedNextSequenceNumber implements reader.CCIP.
func (r InMemoryCCIPReader) GetExpectedNextSequenceNumber(
	ctx context.Context,
	sourceChainSelector cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	panic("unimplemented")
}

func (r InMemoryCCIPReader) CommitReportsGTETimestamp(ctx context.Context,
	ts time.Time,
	confidence primitives.ConfidenceLevel,
	limit int) ([]cciptypes.CommitPluginReportWithMeta, error) {
	unfinalized := slicelib.Filter(r.UnfinalizedReports, func(report cciptypes.CommitPluginReportWithMeta) bool {
		return report.Timestamp.After(ts) || report.Timestamp.Equal(ts)
	})

	finalized := slicelib.Filter(r.FinalizedReports, func(report cciptypes.CommitPluginReportWithMeta) bool {
		return report.Timestamp.After(ts) || report.Timestamp.Equal(ts)
	})

	if len(unfinalized) > limit {
		unfinalized = unfinalized[:limit]
	}
	if len(finalized) > limit {
		finalized = finalized[:limit]
	}

	if confidence == primitives.Unconfirmed {
		return unfinalized, nil
	}
	return finalized, nil
}

func (r InMemoryCCIPReader) ExecutedMessages(
	ctx context.Context,
	rangesByChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	_ primitives.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	var ret = make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
	for source, seqNumRanges := range rangesByChain {
		msgs, ok := r.Messages[source]
		// no messages for chain
		if !ok {
			return nil, nil
		}
		filtered := slicelib.Filter(msgs, func(msg MessagesWithMetadata) bool {
			if msg.Destination != r.Dest || !msg.Executed {
				return false
			}
			for _, r := range seqNumRanges {
				if r.Contains(msg.Header.SequenceNumber) {
					return true
				}
			}
			return false
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

		seqNums := make([]cciptypes.SeqNum, 0, len(ranges))
		for _, r := range ranges {
			seqNums = append(seqNums, r.ToSlice()...)
		}

		unqSeqNums := mapset.NewSet(seqNums...).ToSlice()
		sort.Slice(unqSeqNums, func(i, j int) bool { return unqSeqNums[i] < unqSeqNums[j] })
		ret[source] = unqSeqNums
	}
	return ret, nil
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

func (r InMemoryCCIPReader) LatestMsgSeqNum(
	ctx context.Context, chain cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	return 0, nil
}

func (r InMemoryCCIPReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	panic("implement me")
}

func (r InMemoryCCIPReader) Nonces(
	ctx context.Context,
	addressesByChain map[cciptypes.ChainSelector][]string,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	return nil, nil
}

func (r InMemoryCCIPReader) GetChainsFeeComponents(
	ctx context.Context,
	chains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	panic("implement me")
}

func (r InMemoryCCIPReader) GetDestChainFeeComponents(_ context.Context) (types.ChainFeeComponents, error) {
	feeComponents := types.ChainFeeComponents{
		ExecutionFee:        big.NewInt(0),
		DataAvailabilityFee: big.NewInt(0),
	}
	return feeComponents, nil
}

func (r InMemoryCCIPReader) GetWrappedNativeTokenPriceUSD(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	return nil
}

func (r InMemoryCCIPReader) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	return nil
}

func (r InMemoryCCIPReader) DiscoverContracts(
	ctx context.Context,
	allChains []cciptypes.ChainSelector) (reader.ContractAddresses, error) {
	return nil, nil
}

func (r InMemoryCCIPReader) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	return cciptypes.RemoteConfig{}, nil
}

func (r InMemoryCCIPReader) GetRmnCurseInfo(ctx context.Context) (reader.CurseInfo, error) {
	return reader.CurseInfo{
		CursedSourceChains: map[cciptypes.ChainSelector]bool{},
		CursedDestination:  false,
		GlobalCurse:        false,
	}, nil
}

func (r InMemoryCCIPReader) LinkPriceUSD(ctx context.Context) (cciptypes.BigInt, error) {
	return cciptypes.NewBigIntFromInt64(100), nil
}

// Sync can be used to perform frequent syncing operations inside the reader implementation.
// Returns a bool indicating whether something was updated.
func (r InMemoryCCIPReader) Sync(_ context.Context, _ reader.ContractAddresses) error {
	return nil
}

func (r InMemoryCCIPReader) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	return 0, nil
}

func (r InMemoryCCIPReader) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	return r.ConfigDigest, nil
}

func (r InMemoryCCIPReader) GetOffRampSourceChainsConfig(ctx context.Context, chains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]reader.StaticSourceChainConfig, error) {
	return nil, nil
}

// Close implements the reader.CCIPReader interface
func (r InMemoryCCIPReader) Close() error {
	// Since this is an in-memory implementation with no persistent connections
	// or resources to clean up, we can simply return nil
	return nil
}

// Interface compatibility check.
var _ reader.CCIPReader = InMemoryCCIPReader{}
