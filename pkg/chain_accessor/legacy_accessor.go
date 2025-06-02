package chain_accessor

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"golang.org/x/exp/maps"
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

func (l LegacyAccessor) ExecutedMessages(ctx context.Context, rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange, confidence cciptypes.ConfidenceLevel) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	// trim empty ranges from rangesPerChain
	// otherwise we may get SQL errors from the chainreader.
	nonEmptyRangesPerChain := make(map[cciptypes.ChainSelector][]cciptypes.SeqNumRange)
	for chain, ranges := range rangesPerChain {
		if len(ranges) > 0 {
			nonEmptyRangesPerChain[chain] = ranges
		}
	}

	dataTyp := ExecutionStateChangedEvent{}
	keyFilter, countSqNrs := createExecutedMessagesKeyFilter(nonEmptyRangesPerChain, confidence)
	if countSqNrs == 0 {
		lggr.Debugw("no sequence numbers to query", "nonEmptyRangesPerChain", nonEmptyRangesPerChain)
		return nil, nil
	}
	iter, err := l.destContractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameOffRamp,
		keyFilter,
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortBySequence(query.Asc)},
			Limit: query.Limit{
				Count: countSqNrs,
			},
		},
		&dataTyp,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}

	executed := make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
	for _, item := range iter {
		stateChange, ok := item.Data.(*ExecutionStateChangedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %T to ExecutionStateChangedEvent", item.Data)
		}

		if err := validateExecutionStateChangedEvent(stateChange, nonEmptyRangesPerChain); err != nil {
			lggr.Errorw("validate execution state changed event",
				"err", err, "stateChange", stateChange)
			continue
		}

		executed[stateChange.SourceChainSelector] =
			append(executed[stateChange.SourceChainSelector], stateChange.SequenceNumber)
	}

	return executed, nil
}

func (l LegacyAccessor) NextSeqNum(ctx context.Context, sources []cciptypes.ChainSelector) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	// Use our direct fetch method that doesn't affect the cache
	cfgs, err := l.fetchFreshSourceChainConfigs(ctx, l.destChain, sources)
	if err != nil {
		return nil, fmt.Errorf("get source chains config: %w", err)
	}

	res := make(map[cciptypes.ChainSelector]cciptypes.SeqNum, len(sources))
	for _, chain := range sources {
		cfg, exists := cfgs[chain]
		if !exists {
			lggr.Warnf("source chain config not found for chain %d, chain is skipped.", chain)
			continue
		}

		if !cfg.IsEnabled {
			lggr.Infof("source chain %d is disabled, chain is skipped.", chain)
			continue
		}

		if len(cfg.OnRamp) == 0 {
			lggr.Errorf("onRamp misconfigured for chain %d, chain is skipped: %x", chain, cfg.OnRamp)
			continue
		}

		if len(cfg.Router) == 0 {
			lggr.Errorf("router is empty for chain %d, chain is skipped: %v", chain, cfg.Router)
			continue
		}

		if cfg.MinSeqNr == 0 {
			lggr.Errorf("minSeqNr not found for chain %d or is set to 0, chain is skipped.", chain)
			continue
		}

		res[chain] = cciptypes.SeqNum(cfg.MinSeqNr)
	}

	return res, err
}

func (l LegacyAccessor) Nonces(ctx context.Context, addressesByChain map[cciptypes.ChainSelector][]string) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	// sort the input to ensure deterministic results
	sortedChains := maps.Keys(addressesByChain)
	slices.Sort(sortedChains)

	// create the structure that will contain our result
	res := make(map[cciptypes.ChainSelector]map[string]uint64)
	var addressCount int
	for _, addresses := range addressesByChain {
		addressCount += len(addresses)
	}

	contractInput, responses, err := prepareNoncesInput(lggr, addressesByChain, addressCount, sortedChains, l.addrCodec)
	if err != nil {
		return nil, err
	}

	request := contractreader.ExtendedBatchGetLatestValuesRequest{
		consts.ContractNameNonceManager: contractInput,
	}

	batchResult, _, err := l.destContractReader.ExtendedBatchGetLatestValues(
		ctx,
		request,
		false,
	)
	if err != nil {
		return nil, fmt.Errorf("batch get nonces failed: %w", err)
	}

	// Process results, we range over batchResults, but there should only be result for nonce manager
	for _, results := range batchResult {
		if len(results) != len(responses) {
			lggr.Errorw("unexpected number of nonces",
				"expected", len(responses), "got", len(results))
			continue
		}
		for i, readResult := range results {
			key := responses[i]

			returnVal, err := readResult.GetResult()
			if err != nil {
				lggr.Errorw("failed to get nonce for address", "address", key.address, "err", err)
				continue
			}

			val, ok := returnVal.(*uint64)
			if !ok || val == nil {
				lggr.Errorw("invalid nonce value returned", "address", key.address)
				continue
			}
			if _, ok := res[key.chain]; !ok {
				res[key.chain] = make(map[string]uint64)
			}
			res[key.chain][key.address] = *val
		}
	}

	return res, nil
}

func (l LegacyAccessor) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]cciptypes.TimestampedBig, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)
	rets := make(map[cciptypes.ChainSelector]cciptypes.TimestampedBig)

	if len(selectors) == 0 {
		return rets, nil // Return a new empty map
	}

	// 1. Build Batch Request
	contractBatch := make([]types.BatchRead, 0, len(selectors))
	for _, chain := range selectors {
		contractBatch = append(contractBatch, types.BatchRead{
			ReadName: consts.MethodNameGetFeePriceUpdate,
			Params: map[string]any{
				// That actually means that this selector is a source chain for the destChain
				"destChainSelector": chain,
			},
			// Pass a new pointer directly for type inference by the reader
			ReturnVal: new(cciptypes.TimestampedUnixBig),
		})
	}

	// 2. Execute Batch Request
	batchResult, _, err := l.destContractReader.ExtendedBatchGetLatestValues(
		ctx,
		contractreader.ExtendedBatchGetLatestValuesRequest{
			consts.ContractNameFeeQuoter: contractBatch,
		},
		false, // Don't allow stale reads for fee updates
	)

	if err != nil {
		lggr.Errorw("failed to batch get chain fee price updates", "err", err)
		return rets, nil // Return a new empty map
	}

	// 3. Find FeeQuoter Results
	var feeQuoterResults []types.BatchReadResult
	found := false
	for contract, results := range batchResult {
		if contract.Name == consts.ContractNameFeeQuoter {
			feeQuoterResults = results
			found = true
			break // Found the results, exit loop
		}
	}

	if !found {
		lggr.Errorw("FeeQuoter results missing from batch response")
		return rets, nil // Return a new empty map
	}

	if len(feeQuoterResults) != len(selectors) {
		lggr.Errorw("Mismatch between requested selectors and results count",
			"selectors", len(selectors),
			"results", len(feeQuoterResults))
		// Continue processing the results we did get, but this might indicate an issue
	}

	// 4. Process Results using helper
	return l.processFeePriceUpdateResults(lggr, selectors, feeQuoterResults), nil

}

// processFeePriceUpdateResults iterates through batch results, validates them,
// and returns a new feeUpdates map.
func (r LegacyAccessor) processFeePriceUpdateResults(
	lggr logger.Logger,
	selectors []cciptypes.ChainSelector,
	results []types.BatchReadResult,
) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	feeUpdates := make(map[cciptypes.ChainSelector]cciptypes.TimestampedBig)

	for i, chain := range selectors {
		if i >= len(results) {
			// Log error if we have fewer results than requested selectors
			lggr.Errorw("Skipping selector due to missing result",
				"selectorIndex", i,
				"chain", chain,
				"lenFeeQuoterResults", len(results))
			continue
		}

		readResult := results[i]
		val, err := readResult.GetResult()
		if err != nil {
			lggr.Warnw("failed to get chain fee price update from batch result",
				"chain", chain,
				"err", err)
			continue
		}

		// Type assert the result
		update, ok := val.(*cciptypes.TimestampedUnixBig)
		if !ok || update == nil {
			lggr.Warnw("Invalid type or nil value received for chain fee price update",
				"chain", chain,
				"type", fmt.Sprintf("%T", val),
				"ok", ok)
			continue
		}

		// Check if the update is empty
		if update.Timestamp == 0 || update.Value == nil {
			lggr.Debugw("chain fee price update is empty",
				"chain", chain,
				"update", update)
			continue
		}

		// Add valid update to the map
		feeUpdates[chain] = cciptypes.TimeStampedBigFromUnix(*update)
	}

	return feeUpdates
}

func (l LegacyAccessor) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	var latestSeqNr uint64
	err := l.destContractReader.ExtendedGetLatestValue(
		ctx,
		consts.ContractNameOffRamp,
		consts.MethodNameGetLatestPriceSequenceNumber,
		primitives.Unconfirmed,
		map[string]any{},
		&latestSeqNr,
	)
	if err != nil {
		return 0, fmt.Errorf("get latest price sequence number: %w", err)
	}

	return latestSeqNr, nil
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

type chainAddressNonce struct {
	chain    cciptypes.ChainSelector
	address  string
	response uint64
}

type ExecutionStateChangedEvent struct {
	SourceChainSelector cciptypes.ChainSelector
	SequenceNumber      cciptypes.SeqNum
	MessageID           cciptypes.Bytes32
	MessageHash         cciptypes.Bytes32
	State               uint8
	ReturnData          cciptypes.Bytes
	GasUsed             big.Int
}

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

// sourceChainConfig is used to parse the response from the offRamp contract's getSourceChainConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L94
//
//nolint:lll // It's a URL.
type SourceChainConfig struct {
	Router                    []byte // local router
	IsEnabled                 bool
	IsRMNVerificationDisabled bool
	MinSeqNr                  uint64
	OnRamp                    cciptypes.UnknownAddress
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

func createExecutedMessagesKeyFilter(
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence primitives.ConfidenceLevel) (query.KeyFilter, uint64) {

	var chainExpressions []query.Expression
	var countSqNrs uint64
	// final query should look like
	// (chainA && (sqRange1 || sqRange2 || ...)) || (chainB && (sqRange1 || sqRange2 || ...))
	sortedChains := maps.Keys(rangesPerChain)
	slices.Sort(sortedChains)
	for _, srcChain := range sortedChains {
		seqNumRanges := rangesPerChain[srcChain]
		var seqRangeExpressions []query.Expression
		for _, seqNr := range seqNumRanges {
			expr := query.Comparator(consts.EventAttributeSequenceNumber,
				primitives.ValueComparator{
					Value:    seqNr.Start(),
					Operator: primitives.Gte,
				},
				primitives.ValueComparator{
					Value:    seqNr.End(),
					Operator: primitives.Lte,
				})
			seqRangeExpressions = append(seqRangeExpressions, expr)
			countSqNrs += uint64(seqNr.End() - seqNr.Start() + 1)
		}
		combinedSeqNrs := query.Or(seqRangeExpressions...)

		chainExpressions = append(chainExpressions, query.And(
			combinedSeqNrs,
			query.Comparator(consts.EventAttributeSourceChain, primitives.ValueComparator{
				Value:    srcChain,
				Operator: primitives.Eq,
			}),
		))
	}
	extendedQuery := query.Or(chainExpressions...)

	keyFilter := query.KeyFilter{
		Key: consts.EventNameExecutionStateChanged,
		Expressions: []query.Expression{
			extendedQuery,
			// We don't need to wait for an execute state changed event to be finalized
			// before we optimistically mark a message as executed.
			query.Comparator(consts.EventAttributeState, primitives.ValueComparator{
				Value:    0,
				Operator: primitives.Gt,
			}),
			query.Confidence(confidence),
		},
	}
	return keyFilter, countSqNrs
}

func validateExecutionStateChangedEvent(
	ev *ExecutionStateChangedEvent, rangesByChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange) error {
	if ev == nil {
		return fmt.Errorf("execution state changed event is nil")
	}

	if _, ok := rangesByChain[ev.SourceChainSelector]; !ok {
		return fmt.Errorf("source chain of messages was not queries")
	}

	if !ev.SequenceNumber.IsWithinRanges(rangesByChain[ev.SourceChainSelector]) {
		return fmt.Errorf("execution state changed event sequence number is not in the expected range")
	}

	if ev.MessageHash.IsEmpty() {
		return fmt.Errorf("nil message hash")
	}

	if ev.MessageID.IsEmpty() {
		return fmt.Errorf("message ID is zero")
	}

	if ev.State == 0 {
		return fmt.Errorf("state is zero")
	}

	return nil
}

// fetchFreshSourceChainConfigs always fetches fresh source chain configs directly from contracts
// without using any cached values. Use this when up-to-date data is critical, especially
// for sequence number accuracy.
func (l LegacyAccessor) fetchFreshSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]SourceChainConfig, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	// Filter out destination chain
	filteredSourceChains := filterOutChainSelector(sourceChains, destChain)
	if len(filteredSourceChains) == 0 {
		return make(map[cciptypes.ChainSelector]SourceChainConfig), nil
	}

	// Prepare batch requests for the sourceChains to fetch the latest Unfinalized config values.
	contractBatch := make([]types.BatchRead, 0, len(filteredSourceChains))
	validSourceChains := make([]cciptypes.ChainSelector, 0, len(filteredSourceChains))

	for _, chain := range filteredSourceChains {
		validSourceChains = append(validSourceChains, chain)
		contractBatch = append(contractBatch, types.BatchRead{
			ReadName: consts.MethodNameGetSourceChainConfig,
			Params: map[string]any{
				"sourceChainSelector": chain,
			},
			ReturnVal: new(SourceChainConfig),
		})
	}

	// Execute batch request
	results, _, err := l.destContractReader.ExtendedBatchGetLatestValues(
		ctx,
		contractreader.ExtendedBatchGetLatestValuesRequest{
			consts.ContractNameOffRamp: contractBatch,
		},
		false,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get source chain configs: %w", err)
	}

	if len(results) != 1 {
		return nil, fmt.Errorf("unexpected number of results: %d", len(results))
	}

	// Process results
	configs := make(map[cciptypes.ChainSelector]SourceChainConfig)

	for _, readResult := range results {
		if len(readResult) != len(validSourceChains) {
			return nil, fmt.Errorf("selectors and source chain configs length mismatch: sourceChains=%v, results=%v",
				validSourceChains, results)
		}

		for i, chain := range validSourceChains {
			v, err := readResult[i].GetResult()
			if err != nil {
				lggr.Errorw("Failed to get source chain config",
					"chain", chain,
					"error", err)
				return nil, fmt.Errorf("GetSourceChainConfig for chainSelector=%d failed: %w", chain, err)
			}

			cfg, ok := v.(*SourceChainConfig)
			if !ok {
				lggr.Errorw("Invalid result type from GetSourceChainConfig",
					"chain", chain,
					"type", fmt.Sprintf("%T", v))
				return nil, fmt.Errorf(
					"invalid result type (%T) from GetSourceChainConfig for chainSelector=%d, expected *SourceChainConfig", v, chain)
			}

			configs[chain] = *cfg
		}
	}

	return configs, nil
}

// filterOutChainSelector removes a specified chain selector from a slice of chain selectors
func filterOutChainSelector(
	chains []cciptypes.ChainSelector,
	chainToFilter cciptypes.ChainSelector) []cciptypes.ChainSelector {
	if len(chains) == 0 {
		return nil
	}

	filtered := make([]cciptypes.ChainSelector, 0, len(chains))
	for _, chain := range chains {
		if chain != chainToFilter {
			filtered = append(filtered, chain)
		}
	}
	return filtered
}

func prepareNoncesInput(
	lggr logger.Logger,
	addressesByChain map[cciptypes.ChainSelector][]string,
	addressCount int,
	sortedChains []cciptypes.ChainSelector,
	addrCodec cciptypes.AddressCodec) ([]types.BatchRead, []chainAddressNonce, error) {

	contractInput := make([]types.BatchRead, addressCount)
	responses := make([]chainAddressNonce, addressCount)
	var counter int
	for _, chain := range sortedChains {
		addresses := addressesByChain[chain]
		// no addresses on this chain, no need to make requests
		if len(addresses) == 0 {
			continue
		}
		for _, address := range addresses {
			lggr.Infow("getting nonce for address",
				"address", address, "chain", chain)

			sender, err := addrCodec.AddressStringToBytes(string(address), chain)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert address %s to bytes: %w", address, err)
			}
			// TODO: evm only, need to make chain agnostic.
			// pad the sender slice to 32 bytes from the left
			sender = slicelib.LeftPadBytes(sender, 32)
			contractInput[counter] = types.BatchRead{
				ReadName: consts.MethodNameGetInboundNonce,
				Params: map[string]any{
					"sourceChainSelector": chain,
					"sender":              sender,
				},
				ReturnVal: &responses[counter].response,
			}
			responses[counter] = chainAddressNonce{chain: chain, address: address}
			counter++
		}
	}
	return contractInput, responses, nil
}
