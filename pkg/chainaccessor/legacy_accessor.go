package chainaccessor

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// DefaultAccessor is an implementation of cciptypes.ChainAccessor that allows the CCIPReader
// to cutover and migrate away from depending directly on contract reader and contract writer.
type DefaultAccessor struct {
	lggr           logger.Logger
	chainSelector  cciptypes.ChainSelector
	contractReader contractreader.Extended
	contractWriter types.ContractWriter
	addrCodec      cciptypes.AddressCodec
}

var _ cciptypes.ChainAccessor = (*DefaultAccessor)(nil)

func NewDefaultAccessor(
	lggr logger.Logger,
	chainSelector cciptypes.ChainSelector,
	contractReader contractreader.Extended,
	contractWriter types.ContractWriter,
	addrCodec cciptypes.AddressCodec,
) cciptypes.ChainAccessor {
	return &DefaultAccessor{
		lggr:           lggr,
		chainSelector:  chainSelector,
		contractReader: contractReader,
		contractWriter: contractWriter,
		addrCodec:      addrCodec,
	}
}

func (l *DefaultAccessor) Metadata() cciptypes.AccessorMetadata {
	// TODO(NONEVM-1865): implement or remove from CAL interface
	panic("implement me")
}

func (l *DefaultAccessor) GetContractAddress(contractName string) ([]byte, error) {
	bindings := l.contractReader.GetBindings(contractName)
	if len(bindings) != 1 {
		return nil, fmt.Errorf("expected one binding for the %s contract, got %d", contractName, len(bindings))
	}

	addressBytes, err := l.addrCodec.AddressStringToBytes(bindings[0].Binding.Address, l.chainSelector)
	if err != nil {
		return nil, fmt.Errorf("convert address %s to bytes: %w", bindings[0].Binding.Address, err)
	}

	return addressBytes, nil
}

func (l *DefaultAccessor) GetChainFeeComponents(ctx context.Context) (cciptypes.ChainFeeComponents, error) {
	fc, err := l.contractWriter.GetFeeComponents(ctx)
	if err != nil {
		return cciptypes.ChainFeeComponents{}, fmt.Errorf("get fee components: %w", err)
	}

	return *fc, nil
}

func (l *DefaultAccessor) Sync(ctx context.Context, contracts cciptypes.ContractAddresses) error {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) MsgsBetweenSeqNums(
	ctx context.Context,
	destChainSelector cciptypes.ChainSelector,
	seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)
	seq, err := l.contractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPMessageSent,
			Expressions: []query.Expression{
				query.Comparator(consts.EventAttributeDestChain, primitives.ValueComparator{
					Value:    destChainSelector,
					Operator: primitives.Eq,
				}),
				query.Comparator(consts.EventAttributeSequenceNumber, primitives.ValueComparator{
					Value:    seqNumRange.Start(),
					Operator: primitives.Gte,
				}, primitives.ValueComparator{
					Value:    seqNumRange.End(),
					Operator: primitives.Lte,
				}),
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{
				query.NewSortBySequence(query.Asc),
			},
			Limit: query.Limit{
				Count: uint64(seqNumRange.End() - seqNumRange.Start() + 1),
			},
		},
		&SendRequestedEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query onRamp: %w", err)
	}

	lggr.Infow("queried messages between sequence numbers",
		"numMsgs", len(seq),
		"sourceChainSelector", l.chainSelector,
		"seqNumRange", seqNumRange.String(),
	)

	onRampAddress, err := l.GetContractAddress(consts.ContractNameOnRamp)
	if err != nil {
		return nil, fmt.Errorf("failed to get onRamp contract address: %w", err)
	}

	msgs := make([]cciptypes.Message, 0)
	for _, item := range seq {
		msg, ok := item.Data.(*SendRequestedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}

		if err := ValidateSendRequestedEvent(msg, l.chainSelector, destChainSelector, seqNumRange); err != nil {
			lggr.Errorw("validate send requested event", "err", err, "message", msg)
			continue
		}

		msg.Message.Header.OnRamp = onRampAddress
		msgs = append(msgs, msg.Message)
	}

	lggr.Infow("decoded messages between sequence numbers",
		"msgs", msgs,
		"sourceChainSelector", l.chainSelector,
		"seqNumRange", seqNumRange.String(),
	)

	return msgs, nil
}

func (l *DefaultAccessor) LatestMsgSeqNum(
	ctx context.Context,
	destChainSelector cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	seq, err := l.contractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPMessageSent,
			Expressions: []query.Expression{
				query.Comparator(consts.EventAttributeDestChain, primitives.ValueComparator{
					Value:    destChainSelector,
					Operator: primitives.Eq,
				}),
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{
				query.NewSortBySequence(query.Desc),
			},
			Limit: query.Limit{Count: 1},
		},
		&SendRequestedEvent{},
	)
	if err != nil {
		return 0, fmt.Errorf("failed to query onRamp: %w", err)
	}

	lggr.Debugw("queried latest message from source",
		"numMsgs", len(seq),
		"sourceChainSelector", l.chainSelector,
	)
	if len(seq) > 1 {
		return 0, fmt.Errorf("more than one message found for the latest message query")
	}
	if len(seq) == 0 {
		return 0, nil
	}

	item := seq[0]
	msg, ok := item.Data.(*SendRequestedEvent)
	if !ok {
		return 0, fmt.Errorf("failed to cast %v to SendRequestedEvent", item.Data)
	}

	if err := ValidateSendRequestedEvent(msg, l.chainSelector, destChainSelector,
		cciptypes.NewSeqNumRange(msg.Message.Header.SequenceNumber, msg.Message.Header.SequenceNumber)); err != nil {
		return 0, fmt.Errorf("message invalid msg %v: %w", msg, err)
	}

	return msg.SequenceNumber, nil
}

func (l *DefaultAccessor) GetExpectedNextSequenceNumber(
	ctx context.Context,
	destChainSelector cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	var expectedNextSequenceNumber uint64
	err := l.contractReader.ExtendedGetLatestValue(
		ctx,
		consts.ContractNameOnRamp,
		consts.MethodNameGetExpectedNextSequenceNumber,
		primitives.Unconfirmed,
		map[string]any{
			"destChainSelector": destChainSelector,
		},
		&expectedNextSequenceNumber,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get expected next sequence number from onramp, source chain: %d, dest chain: %d: %w",
			l.chainSelector, destChainSelector, err)
	}

	if expectedNextSequenceNumber == 0 {
		return 0, fmt.Errorf("the returned expected next sequence num is 0, source chain: %d, dest chain: %d",
			l.chainSelector, destChainSelector)
	}

	return cciptypes.SeqNum(expectedNextSequenceNumber), nil
}

func (l *DefaultAccessor) GetTokenPriceUSD(
	ctx context.Context,
	address cciptypes.UnknownAddress,
) (cciptypes.BigInt, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) GetFeeQuoterDestChainConfig(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.FeeQuoterDestChainConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) ExecutedMessages(
	ctx context.Context,
	ranges map[cciptypes.ChainSelector]cciptypes.SeqNumRange,
	confidence cciptypes.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) NextSeqNum(
	ctx context.Context,
	sources []cciptypes.ChainSelector,
) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) Nonces(
	ctx context.Context,
	addresses map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) GetOffRampSourceChainsConfig(
	ctx context.Context,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *DefaultAccessor) GetRmnCurseInfo(ctx context.Context) (cciptypes.CurseInfo, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}
