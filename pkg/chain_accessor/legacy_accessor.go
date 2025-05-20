package chain_accessor

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

type LegacyAccessor struct {
	lggr           logger.Logger
	chainSelector  cciptypes.ChainSelector
	contractReader contractreader.Extended
	contractWriter types.ContractWriter
	addrCodec      cciptypes.AddressCodec
}

type SendRequestedEvent struct {
	DestChainSelector cciptypes.ChainSelector
	SequenceNumber    cciptypes.SeqNum
	Message           cciptypes.Message
}

func NewLegacyAccessor(lgger logger.Logger, chainSelector cciptypes.ChainSelector, contractReader contractreader.Extended, contractWriter types.ContractWriter, addrCodec cciptypes.AddressCodec) *LegacyAccessor {
	// for now, all chains use the same cr/cw based legacy accessor
	return &LegacyAccessor{
		lggr:           lgger,
		chainSelector:  chainSelector,
		contractReader: contractReader,
		contractWriter: contractWriter,
		addrCodec:      addrCodec,
	}
}

func (l LegacyAccessor) Metadata() cciptypes.AccessorMetadata {
	// contracts map need to be filled, but right now the contractReader doesn't support fetch all bindings contract, only allow fetch by contract name
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

func (l LegacyAccessor) GetContractAddress(contractName string) ([]byte, error) {
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

func (l LegacyAccessor) GetChainFeeComponents(ctx context.Context) (cciptypes.ChainFeeComponents, error) {
	fc, err := l.contractWriter.GetFeeComponents(ctx)
	if err != nil {
		return cciptypes.ChainFeeComponents{}, fmt.Errorf("get fee components: %w", err)
	}

	return *fc, nil
}

func (l LegacyAccessor) Sync(ctx context.Context, contracts cciptypes.ContractAddresses) error {
	//Noop
	return nil
}

func (l LegacyAccessor) MsgsBetweenSeqNums(ctx context.Context, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange) ([]cciptypes.Message, error) {
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
					Value:    dest,
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

	onRampAddressAfterQuery, err := l.GetContractAddress(consts.ContractNameOnRamp)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address after query: %w", err)
	}

	// TODO keep or delete this check? Need to modify interface and add a param to support this
	// Ensure the onRamp address hasn't changed during the query.
	//if !bytes.Equal(onRampAddress, onRampAddressAfterQuery) {
	//	return nil, fmt.Errorf("onRamp address has changed from %s to %s", onRampAddress, onRampAddressAfterQuery)
	//}

	l.lggr.Infow("queried messages between sequence numbers",
		"numMsgs", len(seq),
		"sourceChainSelector", l.chainSelector,
		"seqNumRange", seqNumRange.String(),
	)

	msgs := make([]cciptypes.Message, 0)
	for _, item := range seq {
		msg, ok := item.Data.(*SendRequestedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}

		if err := validateSendRequestedEvent(msg, l.chainSelector, dest, seqNumRange); err != nil {
			l.lggr.Errorw("validate send requested event", "err", err, "message", msg)
			continue
		}

		msg.Message.Header.OnRamp = onRampAddressAfterQuery
		msgs = append(msgs, msg.Message)
	}

	l.lggr.Infow("decoded messages between sequence numbers", "msgs", msgs,
		"sourceChainSelector", l.chainSelector,
		"seqNumRange", seqNumRange.String())

	return msgs, nil
}

func (l LegacyAccessor) LatestMsgSeqNum(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetExpectedNextSequenceNumber(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetTokenPriceUSD(ctx context.Context, address cciptypes.UnknownAddress) (cciptypes.TimestampedUnixBig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetFeeQuoterDestChainConfig(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.FeeQuoterDestChainConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) CommitReportsGTETimestamp(ctx context.Context, ts time.Time, limit int) ([]cciptypes.CommitPluginReportWithMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) ExecutedMessages(ctx context.Context, ranges map[cciptypes.ChainSelector][]cciptypes.SeqNumRange, confidence cciptypes.ConfidenceLevel) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) NextSeqNum(ctx context.Context, sources []cciptypes.ChainSelector) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) Nonces(ctx context.Context, addresses map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]cciptypes.TimestampedBig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetOffRampSourceChainsConfig(ctx context.Context, sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (l LegacyAccessor) GetRmnCurseInfo(ctx context.Context) (cciptypes.CurseInfo, error) {
	//TODO implement me
	panic("implement me")
}

func validateSendRequestedEvent(
	ev *SendRequestedEvent, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange) error {
	if ev == nil {
		return fmt.Errorf("send requested event is nil")
	}

	if ev.Message.Header.DestChainSelector != dest {
		return fmt.Errorf("msg dest chain is not the expected queried one")
	}
	if ev.DestChainSelector != dest {
		return fmt.Errorf("dest chain is not the expected queried one")
	}

	if ev.Message.Header.SourceChainSelector != source {
		return fmt.Errorf("source chain is not the expected queried one")
	}

	if ev.SequenceNumber != ev.Message.Header.SequenceNumber {
		return fmt.Errorf("event sequence number does not match the message sequence number %d != %d",
			ev.SequenceNumber, ev.Message.Header.SequenceNumber)
	}

	if ev.SequenceNumber < seqNumRange.Start() || ev.SequenceNumber > seqNumRange.End() {
		return fmt.Errorf("send requested event sequence number is not in the expected range")
	}

	if ev.Message.Header.MessageID.IsEmpty() {
		return fmt.Errorf("message ID is zero")
	}

	if len(ev.Message.Receiver) == 0 {
		return fmt.Errorf("empty receiver address: %s", ev.Message.Receiver.String())
	}

	if ev.Message.Sender.IsZeroOrEmpty() {
		return fmt.Errorf("invalid sender address: %s", ev.Message.Sender.String())
	}

	if ev.Message.FeeTokenAmount.IsEmpty() {
		return fmt.Errorf("fee token amount is zero")
	}

	if ev.Message.FeeToken.IsZeroOrEmpty() {
		return fmt.Errorf("invalid fee token: %s", ev.Message.FeeToken.String())
	}

	return nil
}
