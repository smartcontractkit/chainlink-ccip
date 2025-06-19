package chainaccessor

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
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
	tokenAddress cciptypes.UnknownAddress,
) (cciptypes.TimestampedUnixBig, error) {
	var update cciptypes.TimestampedUnixBig
	err := l.contractReader.ExtendedGetLatestValue(
		ctx,
		consts.ContractNameFeeQuoter,
		consts.MethodNameFeeQuoterGetTokenPrice,
		primitives.Unconfirmed,
		map[string]any{
			"token": tokenAddress,
		},
		&update,
	)
	if err != nil {
		return cciptypes.TimestampedUnixBig{}, fmt.Errorf("failed to get token price from fee quoter: %w", err)
	}
	return update, nil
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
	confidence primitives.ConfidenceLevel,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	internalLimit := limit * 2
	iter, err := l.contractReader.ExtendedQueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameCommitReportAccepted,
			Expressions: []query.Expression{
				query.Timestamp(uint64(ts.Unix()), primitives.Gte),
				// We don't need to wait for the commit report accepted event to be finalized
				// before we can start optimistically processing it.
				query.Confidence(confidence),
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
		"destChain", l.chainSelector,
		"ts", ts,
		"limit", internalLimit)

	reports := l.processCommitReports(lggr, iter, ts, limit)
	return reports, nil
}

func (l *DefaultAccessor) ExecutedMessages(
	ctx context.Context,
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence cciptypes.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
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
	iter, err := l.contractReader.ExtendedQueryKey(
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
		for _, seqNrRange := range seqNumRanges {
			expr := query.Comparator(consts.EventAttributeSequenceNumber,
				primitives.ValueComparator{
					Value:    seqNrRange.Start(),
					Operator: primitives.Gte,
				},
				primitives.ValueComparator{
					Value:    seqNrRange.End(),
					Operator: primitives.Lte,
				})
			seqRangeExpressions = append(seqRangeExpressions, expr)
			countSqNrs += uint64(seqNrRange.End() - seqNrRange.Start() + 1)
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
				Value:    0, // > 0 corresponds to:  IN_PROGRESS, SUCCESS, FAILURE
				Operator: primitives.Gt,
			}),
			query.Confidence(confidence),
		},
	}
	return keyFilter, countSqNrs
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
	addressesByChain map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
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

	contractInput, responses, err := prepareNoncesInput(
		lggr,
		addressesByChain,
		addressCount,
		sortedChains,
		l.addrCodec,
	)
	if err != nil {
		return nil, err
	}

	request := contractreader.ExtendedBatchGetLatestValuesRequest{
		consts.ContractNameNonceManager: contractInput,
	}

	batchResult, _, err := l.contractReader.ExtendedBatchGetLatestValues(
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

func prepareNoncesInput(
	lggr logger.Logger,
	addressesByChain map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress,
	addressCount int,
	sortedChains []cciptypes.ChainSelector,
	addrCodec cciptypes.AddressCodec,
) ([]types.BatchRead, []chainAddressNonce, error) {
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
			responses[counter] = chainAddressNonce{chain: chain, address: string(address)}
			counter++
		}
	}
	return contractInput, responses, nil
}

func (l *DefaultAccessor) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	lggr := logutil.WithContextValues(ctx, l.lggr)

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
	batchResult, _, err := l.contractReader.ExtendedBatchGetLatestValues(
		ctx,
		contractreader.ExtendedBatchGetLatestValuesRequest{
			consts.ContractNameFeeQuoter: contractBatch,
		},
		false, // Don't allow stale reads for fee updates
	)

	if err != nil {
		lggr.Errorw("failed to batch get chain fee price updates", "err", err)
		return make(map[cciptypes.ChainSelector]cciptypes.TimestampedBig) // Return a new empty map
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
		return make(map[cciptypes.ChainSelector]cciptypes.TimestampedBig) // Return a new empty map
	}

	if len(feeQuoterResults) != len(selectors) {
		lggr.Errorw("Mismatch between requested selectors and results count",
			"selectors", len(selectors),
			"results", len(feeQuoterResults))
		// Continue processing the results we did get, but this might indicate an issue
	}

	// 4. Process Results using helper
	return l.processFeePriceUpdateResults(lggr, selectors, feeQuoterResults)
}

// processFeePriceUpdateResults iterates through batch results, validates them,
// and returns a new feeUpdates map.
func (l *DefaultAccessor) processFeePriceUpdateResults(
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

func (l *DefaultAccessor) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	var latestSeqNr uint64
	err := l.contractReader.ExtendedGetLatestValue(
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

// processCommitReports decodes the commit reports from the query results
// and returns the ones that can be properly parsed and validated.
func (l *DefaultAccessor) processCommitReports(
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
		blessedMerkleRoots, unblessedMerkleRoots := l.processMerkleRoots(allMerkleRoots, isBlessed)

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

func (l *DefaultAccessor) processMerkleRoots(
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

func (l *DefaultAccessor) processPriceUpdates(priceUpdates PriceUpdates) (cciptypes.PriceUpdates, error) {
	lggr := l.lggr
	updates := cciptypes.PriceUpdates{
		TokenPriceUpdates: make([]cciptypes.TokenPrice, 0),
		GasPriceUpdates:   make([]cciptypes.GasPriceChain, 0),
	}

	for _, tokenPriceUpdate := range priceUpdates.TokenPriceUpdates {
		sourceTokenAddrStr, err := l.addrCodec.AddressBytesToString(tokenPriceUpdate.SourceToken, l.chainSelector)
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
