package inmem

import (
	"context"
	"math/big"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	internaltypes "github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
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

func (r InMemoryCCIPReader) CommitReportsGTETimestamp(
	_ context.Context, ts time.Time, limit int,
) ([]plugintypes.CommitPluginReportWithMeta, error) {
	results := slicelib.Filter(r.Reports, func(report plugintypes.CommitPluginReportWithMeta) bool {
		return report.Timestamp.After(ts) || report.Timestamp.Equal(ts)
	})
	if len(results) > limit {
		return results[:limit], nil
	}
	return results, nil
}

func (r InMemoryCCIPReader) ExecutedMessages(
	ctx context.Context, source cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.SeqNum, error) {
	msgs, ok := r.Messages[source]
	// no messages for chain
	if !ok {
		return nil, nil
	}
	filtered := slicelib.Filter(msgs, func(msg MessagesWithMetadata) bool {
		return seqNumRange.Contains(msg.Header.SequenceNumber) && msg.Destination == r.Dest && msg.Executed
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
	return unqSeqNums, nil
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
	source cciptypes.ChainSelector,
	addresses []string,
) (map[string]uint64, error) {
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
) map[cciptypes.ChainSelector]internaltypes.TimestampedBig {
	return nil
}

func (r InMemoryCCIPReader) DiscoverContracts(
	ctx context.Context,
	allChains []cciptypes.ChainSelector) (reader.ContractAddresses, error) {
	return nil, nil
}

func (r InMemoryCCIPReader) GetRMNRemoteConfig(ctx context.Context) (rmntypes.RemoteConfig, error) {
	return rmntypes.RemoteConfig{}, nil
}

func (r InMemoryCCIPReader) GetRmnCurseInfo(
	ctx context.Context, sourceChainSelectors []cciptypes.ChainSelector,
) (*reader.CurseInfo, error) {
	return &reader.CurseInfo{
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

func (r InMemoryCCIPReader) GetMedianDataAvailabilityGasConfig(
	ctx context.Context,
) (cciptypes.DataAvailabilityGasConfig, error) {
	return cciptypes.DataAvailabilityGasConfig{}, nil
}

func (r InMemoryCCIPReader) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	return 0, nil
}

func (r InMemoryCCIPReader) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	return r.ConfigDigest, nil
}

func (r InMemoryCCIPReader) GetOffRampSourceChainsConfig(ctx context.Context, chains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]reader.SourceChainConfig, error) {
	return nil, nil
}

// Interface compatibility check.
var _ reader.CCIPReader = InMemoryCCIPReader{}
