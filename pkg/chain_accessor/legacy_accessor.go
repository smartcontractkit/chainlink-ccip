package chain_accessor

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

type LegacyAccessor struct {
	lggr               logger.Logger
	srcChain           cciptypes.ChainSelector
	srcContractReader  contractreader.Extended
	srcContractWriter  types.ContractWriter
	destContractReader contractreader.Extended
	destContractWriter types.ContractWriter
	addrCodec          cciptypes.AddressCodec
	destChain          cciptypes.ChainSelector
}

type SendRequestedEvent struct {
	DestChainSelector cciptypes.ChainSelector
	SequenceNumber    cciptypes.SeqNum
	Message           cciptypes.Message
}

func NewLegacyAccessor(lgger logger.Logger, srcChain cciptypes.ChainSelector, destChain cciptypes.ChainSelector,
	srcContractReader contractreader.Extended, srcContractWriter types.ContractWriter, destContractReader contractreader.Extended,
	destContractWriter types.ContractWriter, addrCodec cciptypes.AddressCodec,
) *LegacyAccessor {
	// for now, all chains use the same cr/cw based legacy accessor
	return &LegacyAccessor{
		lggr:               lgger,
		srcChain:           srcChain,
		destChain:          destChain,
		srcContractReader:  srcContractReader,
		srcContractWriter:  srcContractWriter,
		destContractReader: destContractReader,
		destContractWriter: destContractWriter,
		addrCodec:          addrCodec,
	}
}

func (l LegacyAccessor) Metadata() cciptypes.AccessorMetadata {
	// contracts map need to be filled, but right now the srcContractReader doesn't support fetch all bindings contract, only allow fetch by contract name
	allBindings := l.srcContractReader.GetAllBindings()
	contracts := make(map[string]cciptypes.UnknownAddress, len(allBindings))
	for contractName, binding := range allBindings {
		addressBytes, err := l.addrCodec.AddressStringToBytes(binding[0].Binding.Address, l.srcChain)
		if err != nil {
			l.lggr.Errorf("failed to convert address to bytes : %v", err)
			continue
		}
		contracts[contractName] = addressBytes
	}

	return cciptypes.AccessorMetadata{
		ChainSelector: l.srcChain,
		Contracts:     contracts,
	}
}

func (l LegacyAccessor) GetContractAddress(contractName string) ([]byte, error) {
	bindings := l.srcContractReader.GetBindings(contractName)
	if len(bindings) != 1 {
		return nil, fmt.Errorf("expected one binding for the %s contract, got %d", contractName, len(bindings))
	}

	addressBytes, err := l.addrCodec.AddressStringToBytes(bindings[0].Binding.Address, l.srcChain)
	if err != nil {
		return nil, fmt.Errorf("convert address %s to bytes: %w", bindings[0].Binding.Address, err)
	}

	return addressBytes, nil
}

func (l LegacyAccessor) GetChainFeeComponents(ctx context.Context) (cciptypes.ChainFeeComponents, error) {
	fc, err := l.srcContractWriter.GetFeeComponents(ctx)
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
	seq, err := l.srcContractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPMessageSent,
			Expressions: []query.Expression{
				query.Comparator(consts.EventAttributeSourceChain, primitives.ValueComparator{
					Value:    l.srcChain,
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

	onRampAddress, err := l.GetContractAddress(consts.ContractNameOnRamp)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address: %w", err)
	}

	//Ensure the onRamp address hasn't changed during the query.
	if !bytes.Equal(onRampAddress, onRampAddressAfterQuery) {
		return nil, fmt.Errorf("onRamp address has changed from %s to %s", onRampAddress, onRampAddressAfterQuery)
	}

	l.lggr.Infow("queried messages between sequence numbers",
		"numMsgs", len(seq),
		"sourceChainSelector", l.srcChain,
		"seqNumRange", seqNumRange.String(),
	)

	msgs := make([]cciptypes.Message, 0)
	for _, item := range seq {
		msg, ok := item.Data.(*SendRequestedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}

		if err := validateSendRequestedEvent(msg, l.srcChain, dest, seqNumRange); err != nil {
			l.lggr.Errorw("validate send requested event", "err", err, "message", msg)
			continue
		}

		msg.Message.Header.OnRamp = onRampAddressAfterQuery
		msgs = append(msgs, msg.Message)
	}

	l.lggr.Infow("decoded messages between sequence numbers", "msgs", msgs,
		"sourceChainSelector", l.srcChain,
		"seqNumRange", seqNumRange.String())

	return msgs, nil
}

func (l LegacyAccessor) LatestMsgSeqNum(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)
	seq, err := l.srcContractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPMessageSent,
			Expressions: []query.Expression{
				query.Comparator(consts.EventAttributeSourceChain, primitives.ValueComparator{
					Value:    l.srcChain,
					Operator: primitives.Eq,
				}),
				query.Comparator(consts.EventAttributeDestChain, primitives.ValueComparator{
					Value:    dest,
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
		"numMsgs", len(seq), "sourceChainSelector", dest)
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

	if err := validateSendRequestedEvent(msg, l.srcChain, dest,
		cciptypes.NewSeqNumRange(msg.Message.Header.SequenceNumber, msg.Message.Header.SequenceNumber)); err != nil {
		return 0, fmt.Errorf("message invalid msg %v: %w", msg, err)
	}

	lggr.Infow("chain reader returning latest onramp sequence number",
		"seqNum", msg.Message.Header.SequenceNumber, "sourceChainSelector", l.srcChain)
	return msg.SequenceNumber, nil
}

func (l LegacyAccessor) GetExpectedNextSequenceNumber(ctx context.Context, dest cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	var expectedNextSequenceNumber uint64
	err := l.srcContractReader.ExtendedGetLatestValue(
		ctx,
		consts.ContractNameOnRamp,
		consts.MethodNameGetExpectedNextSequenceNumber,
		primitives.Unconfirmed,
		map[string]any{
			"destChainSelector": dest,
		},
		&expectedNextSequenceNumber,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get expected next sequence number from onramp, source chain: %d, dest chain: %d: %w",
			l.srcChain, dest, err)
	}

	if expectedNextSequenceNumber == 0 {
		return 0, fmt.Errorf("the returned expected next sequence num is 0, source chain: %d, dest chain: %d",
			l.srcChain, dest)
	}

	lggr.Debugw("chain reader returning expected next sequence number",
		"seqNum", expectedNextSequenceNumber, "sourceChainSelector", l.srcChain)
	return cciptypes.SeqNum(expectedNextSequenceNumber), nil
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
	lggr := logutil.WithContextValues(ctx, l.lggr)
	internalLimit := limit * 2
	iter, err := l.destContractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameCommitReportAccepted,
			Expressions: []query.Expression{
				query.Timestamp(uint64(ts.Unix()), primitives.Gte),
				// We don't need to wait for the commit report accepted event to be finalized
				// before we can start optimistically processing it.
				query.Confidence(primitives.Unconfirmed),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortBySequence(query.Asc)},
			Limit: query.Limit{
				Count: uint64(internalLimit),
			},
		},
		&CommitReportAcceptedEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}

	lggr.Debugw("queried commit reports", "numReports", len(iter),
		"confidence", primitives.Unconfirmed,
		"destChain", l.destChain,
		"ts", ts,
		"limit", internalLimit)

	reports := l.processCommitReports(lggr, iter, ts, limit)

	lggr.Debugw("decoded commit reports", "reports", reports)
	return reports, nil
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

// processCommitReports decodes the commit reports from the query results
// and returns the ones that can be properly parsed and validated.
func (l LegacyAccessor) processCommitReports(
	lggr logger.Logger, iter []types.Sequence, ts time.Time, limit int,
) []cciptypes.CommitPluginReportWithMeta {
	reports := make([]cciptypes.CommitPluginReportWithMeta, 0)
	for _, item := range iter {
		ev, err := validateCommitReportAcceptedEvent(item, ts)
		if err != nil {
			lggr.Errorw("validate commit report accepted event", "err", err, "ev", item.Data)
			continue
		}

		lggr.Debugw("processing commit report", "report", ev, "item", item)

		isBlessed := make(map[cciptypes.Bytes32]bool, len(ev.BlessedMerkleRoots))
		for _, mr := range ev.BlessedMerkleRoots {
			isBlessed[mr.MerkleRoot] = true
		}
		allMerkleRoots := append(ev.BlessedMerkleRoots, ev.UnblessedMerkleRoots...)
		blessedMerkleRoots, unblessedMerkleRoots := processMerkleRoots(allMerkleRoots, isBlessed)

		priceUpdates, err := l.processPriceUpdates(ev.PriceUpdates)
		if err != nil {
			lggr.Errorw("failed to process price updates", "err", err, "priceUpdates", ev.PriceUpdates)
			continue
		}

		blockNum, err := strconv.ParseUint(item.Head.Height, 10, 64)
		if err != nil {
			lggr.Errorw("failed to parse block number", "blockNum", item.Head.Height, "err", err)
			continue
		}

		reports = append(reports, cciptypes.CommitPluginReportWithMeta{
			Report: cciptypes.CommitPluginReport{
				BlessedMerkleRoots:   blessedMerkleRoots,
				UnblessedMerkleRoots: unblessedMerkleRoots,
				PriceUpdates:         priceUpdates,
			},
			Timestamp: time.Unix(int64(item.Timestamp), 0),
			BlockNum:  blockNum,
		})
	}

	lggr.Debugw("decoded commit reports", "reports", reports)

	if len(reports) < limit {
		return reports
	}
	return reports[:limit]
}

// ------------ helper functions migrated from reader/ccip.go ------------

type MerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       cciptypes.UnknownAddress
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          cciptypes.Bytes32
}

type CommitReportAcceptedEvent struct {
	BlessedMerkleRoots   []MerkleRoot
	UnblessedMerkleRoots []MerkleRoot
	PriceUpdates         PriceUpdates
}

type PriceUpdates struct {
	TokenPriceUpdates []TokenPriceUpdate
	GasPriceUpdates   []GasPriceUpdate
}

type TokenPriceUpdate struct {
	SourceToken cciptypes.UnknownAddress
	UsdPerToken *big.Int
}

type GasPriceUpdate struct {
	// DestChainSelector is the chain that the gas price is for (some plugin source chain).
	// Not the chain that the gas price is stored on.
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type rmnDigestHeader struct {
	DigestHeader cciptypes.Bytes32
}

func validateCommitReportAcceptedEvent(seq types.Sequence, gteTimestamp time.Time) (*CommitReportAcceptedEvent, error) {
	ev, is := (seq.Data).(*CommitReportAcceptedEvent)
	if !is {
		return nil, fmt.Errorf("unexpected type %T while expecting a commit report", seq)
	}

	if ev == nil {
		return nil, fmt.Errorf("commit report accepted event is nil")
	}

	if seq.Timestamp < uint64(gteTimestamp.Unix()) {
		return nil, fmt.Errorf("commit report accepted event timestamp is less than the minimum timestamp %v<%v",
			seq.Timestamp, gteTimestamp.Unix())
	}

	if err := validateMerkleRoots(append(ev.BlessedMerkleRoots, ev.UnblessedMerkleRoots...)); err != nil {
		return nil, fmt.Errorf("merkle roots: %w", err)
	}

	for _, tpus := range ev.PriceUpdates.TokenPriceUpdates {
		if tpus.SourceToken.IsZeroOrEmpty() {
			return nil, fmt.Errorf("invalid source token address: %s", tpus.SourceToken.String())
		}
		if tpus.UsdPerToken == nil || tpus.UsdPerToken.Cmp(big.NewInt(0)) <= 0 {
			return nil, fmt.Errorf("nil or non-positive usd per token")
		}
	}

	for _, gpus := range ev.PriceUpdates.GasPriceUpdates {
		if gpus.UsdPerUnitGas == nil || gpus.UsdPerUnitGas.Cmp(big.NewInt(0)) < 0 {
			return nil, fmt.Errorf("nil or negative usd per unit gas: %s", gpus.UsdPerUnitGas.String())
		}
	}

	return ev, nil
}

func validateMerkleRoots(merkleRoots []MerkleRoot) error {
	seenRoots := mapset.NewSet[cciptypes.Bytes32]()

	for _, mr := range merkleRoots {
		if seenRoots.Contains(mr.MerkleRoot) {
			return fmt.Errorf("duplicate merkle root: %s", mr.MerkleRoot.String())
		}
		seenRoots.Add(mr.MerkleRoot)

		if mr.SourceChainSelector == 0 {
			return fmt.Errorf("source chain is zero")
		}
		if mr.MinSeqNr == 0 {
			return fmt.Errorf("minSeqNr is zero")
		}
		if mr.MaxSeqNr == 0 {
			return fmt.Errorf("maxSeqNr is zero")
		}
		if mr.MinSeqNr > mr.MaxSeqNr {
			return fmt.Errorf("minSeqNr is greater than maxSeqNr")
		}
		if mr.MerkleRoot.IsEmpty() {
			return fmt.Errorf("empty merkle root")
		}
		if mr.OnRampAddress.IsZeroOrEmpty() {
			return fmt.Errorf("invalid onramp address: %s", mr.OnRampAddress.String())
		}
	}

	return nil
}

func processMerkleRoots(
	allMerkleRoots []MerkleRoot, isBlessed map[cciptypes.Bytes32]bool,
) (blessedMerkleRoots []cciptypes.MerkleRootChain, unblessedMerkleRoots []cciptypes.MerkleRootChain) {
	blessedMerkleRoots = make([]cciptypes.MerkleRootChain, 0, len(isBlessed))
	unblessedMerkleRoots = make([]cciptypes.MerkleRootChain, 0, len(allMerkleRoots)-len(isBlessed))
	for _, mr := range allMerkleRoots {
		mrc := cciptypes.MerkleRootChain{
			ChainSel:      cciptypes.ChainSelector(mr.SourceChainSelector),
			OnRampAddress: mr.OnRampAddress,
			SeqNumsRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(mr.MinSeqNr),
				cciptypes.SeqNum(mr.MaxSeqNr),
			),
			MerkleRoot: mr.MerkleRoot,
		}
		if isBlessed[mr.MerkleRoot] {
			blessedMerkleRoots = append(blessedMerkleRoots, mrc)
		} else {
			unblessedMerkleRoots = append(unblessedMerkleRoots, mrc)
		}
	}
	return blessedMerkleRoots, unblessedMerkleRoots
}

func (l LegacyAccessor) processPriceUpdates(
	priceUpdates PriceUpdates,
) (cciptypes.PriceUpdates, error) {
	lggr := l.lggr
	updates := cciptypes.PriceUpdates{
		TokenPriceUpdates: make([]cciptypes.TokenPrice, 0),
		GasPriceUpdates:   make([]cciptypes.GasPriceChain, 0),
	}

	for _, tokenPriceUpdate := range priceUpdates.TokenPriceUpdates {
		sourceTokenAddrStr, err := l.addrCodec.AddressBytesToString(tokenPriceUpdate.SourceToken, l.destChain)
		if err != nil {
			lggr.Errorw("failed to convert source token address to string", "err", err)
			return updates, err
		}
		updates.TokenPriceUpdates = append(updates.TokenPriceUpdates, cciptypes.TokenPrice{
			TokenID: cciptypes.UnknownEncodedAddress(sourceTokenAddrStr),
			Price:   cciptypes.NewBigInt(tokenPriceUpdate.UsdPerToken),
		})
	}

	for _, gasPriceUpdate := range priceUpdates.GasPriceUpdates {
		updates.GasPriceUpdates = append(updates.GasPriceUpdates, cciptypes.GasPriceChain{
			ChainSel: cciptypes.ChainSelector(gasPriceUpdate.DestChainSelector),
			GasPrice: cciptypes.NewBigInt(gasPriceUpdate.UsdPerUnitGas),
		})
	}

	return updates, nil
}
