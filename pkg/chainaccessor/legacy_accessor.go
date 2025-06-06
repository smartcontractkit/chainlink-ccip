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

// LegacyAccessor is an implementation of cciptypes.ChainAccessor that allows the CCIPReader
// to cutover and migrate away from depending directly on contract reader and contract writer.
type LegacyAccessor struct {
	lggr           logger.Logger
	chainSelector  cciptypes.ChainSelector
	contractReader contractreader.Extended
	contractWriter types.ContractWriter
	addrCodec      cciptypes.AddressCodec
}

var _ cciptypes.ChainAccessor = (*LegacyAccessor)(nil)

func NewLegacyAccessor(
	lggr logger.Logger,
	chainSelector cciptypes.ChainSelector,
	contractReader contractreader.Extended,
	contractWriter types.ContractWriter,
	addrCodec cciptypes.AddressCodec,
) cciptypes.ChainAccessor {
	return &LegacyAccessor{
		lggr:           lggr,
		chainSelector:  chainSelector,
		contractReader: contractReader,
		contractWriter: contractWriter,
		addrCodec:      addrCodec,
	}
}

func (l *LegacyAccessor) Metadata() cciptypes.AccessorMetadata {
	allBindings := l.contractReader.GetAllBindings()
	contracts := make(map[string]cciptypes.UnknownAddress, len(allBindings))
	for contractName, binding := range allBindings {
		addressBytes, err := l.addrCodec.AddressStringToBytes(binding[0].Binding.Address, l.chainSelector)
		if err != nil {
			l.lggr.Errorf("failed to convert address to bytes : %v", err)
			continue
		}
		contracts[contractName] = addressBytes
	}

	return cciptypes.AccessorMetadata{
		ChainSelector: l.chainSelector,
		Contracts:     contracts,
	}
}

func (l *LegacyAccessor) GetContractAddress(contractName string) ([]byte, error) {
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

func (l *LegacyAccessor) GetChainFeeComponents(ctx context.Context) (cciptypes.ChainFeeComponents, error) {
	fc, err := l.contractWriter.GetFeeComponents(ctx)
	if err != nil {
		return cciptypes.ChainFeeComponents{}, fmt.Errorf("get fee components: %w", err)
	}

	return *fc, nil
}

func (l *LegacyAccessor) Sync(ctx context.Context, contracts cciptypes.ContractAddresses) error {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) MsgsBetweenSeqNums(
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
				query.Comparator(consts.EventAttributeSourceChain, primitives.ValueComparator{
					Value:    l.chainSelector,
					Operator: primitives.Eq,
				}),
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
		&cciptypes.SendRequestedEvent{},
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
		msg, ok := item.Data.(*cciptypes.SendRequestedEvent)
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

func (l *LegacyAccessor) LatestMsgSeqNum(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetExpectedNextSequenceNumber(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetTokenPriceUSD(
	ctx context.Context,
	address cciptypes.UnknownAddress,
) (cciptypes.BigInt, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetFeeQuoterDestChainConfig(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.FeeQuoterDestChainConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) ExecutedMessages(
	ctx context.Context,
	ranges map[cciptypes.ChainSelector]cciptypes.SeqNumRange,
	confidence cciptypes.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) NextSeqNum(
	ctx context.Context,
	sources []cciptypes.ChainSelector,
) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) Nonces(
	ctx context.Context,
	addresses map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetOffRampSourceChainsConfig(
	ctx context.Context,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}

func (l *LegacyAccessor) GetRmnCurseInfo(ctx context.Context) (cciptypes.CurseInfo, error) {
	// TODO(NONEVM-1865): implement
	panic("implement me")
}
