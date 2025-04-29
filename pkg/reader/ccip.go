package reader

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"sort"
	"strconv"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Default refresh period for cache if not specified
const defaultRefreshPeriod = 30 * time.Second

// TODO: unit test the implementation when the actual contract reader and writer interfaces are finalized and mocks
// can be generated.
type ccipChainReader struct {
	lggr            logger.Logger
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter
	destChain       cciptypes.ChainSelector
	offrampAddress  string
	configPoller    ConfigPoller
	addrCodec       cciptypes.AddressCodec
}

func newCCIPChainReaderInternal(
	ctx context.Context,
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
	addrCodec cciptypes.AddressCodec,
) *ccipChainReader {
	var crs = make(map[cciptypes.ChainSelector]contractreader.Extended)
	for chainSelector, cr := range contractReaders {
		crs[chainSelector] = contractreader.NewExtendedContractReader(cr)
	}

	offrampAddrStr, err := addrCodec.AddressBytesToString(offrampAddress, destChain)
	if err != nil {
		panic(fmt.Sprintf("failed to convert offramp address to string: %v", err))
	}

	reader := &ccipChainReader{
		lggr:            lggr,
		contractReaders: crs,
		contractWriters: contractWriters,
		destChain:       destChain,
		offrampAddress:  offrampAddrStr,
		addrCodec:       addrCodec,
	}

	// Initialize cache with readers
	reader.configPoller = newConfigPoller(lggr, reader, defaultRefreshPeriod)

	contracts := ContractAddresses{
		consts.ContractNameOffRamp: {
			destChain: offrampAddress,
		},
	}
	if err := reader.Sync(ctx, contracts); err != nil {
		lggr.Errorw("failed to sync contracts", "err", err)
	}

	// After contracts are synced, start the background polling
	lggr.Info("Starting config background polling")
	if err := reader.configPoller.Start(ctx); err != nil {
		// Log the error but don't fail - we can still function without background polling
		// by fetching configs on demand
		lggr.Errorw("failed to start config background polling", "err", err)
	}

	return reader
}

// WithExtendedContractReader sets the extended contract reader for the provided chain.
func (r *ccipChainReader) WithExtendedContractReader(
	ch cciptypes.ChainSelector, cr contractreader.Extended) *ccipChainReader {
	r.contractReaders[ch] = cr
	return r
}

func (r *ccipChainReader) Close() error {
	if err := r.configPoller.Close(); err != nil {
		r.lggr.Warnw("Error closing config poller", "err", err)
		// Continue with shutdown even if there's an error
	}
	r.lggr.Info("Stopped CCIP chain reader")
	return nil
}

// ---------------------------------------------------
// The following types are used to decode the events
// but should be replaced by chain-reader modifiers and use the base cciptypes.CommitReport type.

type MerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       cciptypes.UnknownAddress
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          cciptypes.Bytes32
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

type PriceUpdates struct {
	TokenPriceUpdates []TokenPriceUpdate
	GasPriceUpdates   []GasPriceUpdate
}

type CommitReportAcceptedEvent struct {
	BlessedMerkleRoots   []MerkleRoot
	UnblessedMerkleRoots []MerkleRoot
	PriceUpdates         PriceUpdates
}

type rmnDigestHeader struct {
	DigestHeader cciptypes.Bytes32
}

type OCRConfigResponse struct {
	OCRConfig OCRConfig
}

type OCRConfig struct {
	ConfigInfo   ConfigInfo
	Signers      [][]byte
	Transmitters [][]byte
}

type ConfigInfo struct {
	ConfigDigest                   [32]byte
	F                              uint8
	N                              uint8
	IsSignatureVerificationEnabled bool
}

type RMNCurseResponse struct {
	CursedSubjects [][16]byte
}

// ---------------------------------------------------

func (r *ccipChainReader) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	confidence primitives.ConfidenceLevel,
	limit int) ([]cciptypes.CommitPluginReportWithMeta, error) {

	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return []cciptypes.CommitPluginReportWithMeta{}, err
	}

	lggr := logutil.WithContextValues(ctx, r.lggr)
	internalLimit := limit * 2
	iter, err := r.contractReaders[r.destChain].ExtendedQueryKey(
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
		"confidence", confidence,
		"destChain", r.destChain,
		"ts", ts,
		"limit", internalLimit)

	reports := r.processCommitReports(lggr, iter, ts, limit)

	lggr.Debugw("decoded commit reports", "reports", reports)
	return reports, nil
}

// processCommitReports decodes the commit reports from the query results
// and returns the ones that can be properly parsed and validated.
func (r *ccipChainReader) processCommitReports(
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
		blessedMerkleRoots, unblessedMerkleRoots := r.processMerkleRoots(allMerkleRoots, isBlessed)

		priceUpdates, err := r.processPriceUpdates(ev.PriceUpdates)
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

func (r *ccipChainReader) processMerkleRoots(
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

func (r *ccipChainReader) processPriceUpdates(
	priceUpdates PriceUpdates,
) (cciptypes.PriceUpdates, error) {
	lggr := r.lggr
	updates := cciptypes.PriceUpdates{
		TokenPriceUpdates: make([]cciptypes.TokenPrice, 0),
		GasPriceUpdates:   make([]cciptypes.GasPriceChain, 0),
	}

	for _, tokenPriceUpdate := range priceUpdates.TokenPriceUpdates {
		sourceTokenAddrStr, err := r.addrCodec.AddressBytesToString(tokenPriceUpdate.SourceToken, r.destChain)
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

type ExecutionStateChangedEvent struct {
	SourceChainSelector cciptypes.ChainSelector
	SequenceNumber      cciptypes.SeqNum
	MessageID           cciptypes.Bytes32
	MessageHash         cciptypes.Bytes32
	State               uint8
	ReturnData          cciptypes.Bytes
	GasUsed             big.Int
}

func (r *ccipChainReader) ExecutedMessages(
	ctx context.Context,
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence primitives.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, err
	}

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
	iter, err := r.contractReaders[r.destChain].ExtendedQueryKey(
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

type SendRequestedEvent struct {
	DestChainSelector cciptypes.ChainSelector
	SequenceNumber    cciptypes.SeqNum
	Message           cciptypes.Message
}

func (r *ccipChainReader) MsgsBetweenSeqNums(
	ctx context.Context, sourceChainSelector cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, sourceChainSelector); err != nil {
		return nil, err
	}

	onRampAddress, err := r.GetContractAddress(consts.ContractNameOnRamp, sourceChainSelector)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address: %w", err)
	}

	seq, err := r.contractReaders[sourceChainSelector].ExtendedQueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPMessageSent,
			Expressions: []query.Expression{
				query.Comparator(consts.EventAttributeSourceChain, primitives.ValueComparator{
					Value:    sourceChainSelector,
					Operator: primitives.Eq,
				}),
				query.Comparator(consts.EventAttributeDestChain, primitives.ValueComparator{
					Value:    r.destChain,
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

	onRampAddressAfterQuery, err := r.GetContractAddress(consts.ContractNameOnRamp, sourceChainSelector)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address after query: %w", err)
	}

	// Ensure the onRamp address hasn't changed during the query.
	if !bytes.Equal(onRampAddress, onRampAddressAfterQuery) {
		return nil, fmt.Errorf("onRamp address has changed from %s to %s", onRampAddress, onRampAddressAfterQuery)
	}

	lggr.Infow("queried messages between sequence numbers",
		"numMsgs", len(seq),
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String(),
	)

	msgs := make([]cciptypes.Message, 0)
	for _, item := range seq {
		msg, ok := item.Data.(*SendRequestedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %v to Message", item.Data)
		}

		if err := validateSendRequestedEvent(msg, sourceChainSelector, r.destChain, seqNumRange); err != nil {
			lggr.Errorw("validate send requested event", "err", err, "message", msg)
			continue
		}

		msg.Message.Header.OnRamp = onRampAddress
		msgs = append(msgs, msg.Message)
	}

	lggr.Infow("decoded messages between sequence numbers", "msgs", msgs,
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String())

	return msgs, nil
}

// LatestMsgSeqNum reads the source chain and returns the latest finalized message sequence number.
func (r *ccipChainReader) LatestMsgSeqNum(
	ctx context.Context, chain cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, chain); err != nil {
		return 0, err
	}

	seq, err := r.contractReaders[chain].ExtendedQueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPMessageSent,
			Expressions: []query.Expression{
				query.Comparator(consts.EventAttributeSourceChain, primitives.ValueComparator{
					Value:    chain,
					Operator: primitives.Eq,
				}),
				query.Comparator(consts.EventAttributeDestChain, primitives.ValueComparator{
					Value:    r.destChain,
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
		"numMsgs", len(seq), "sourceChainSelector", chain)
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

	if err := validateSendRequestedEvent(msg, chain, r.destChain,
		cciptypes.NewSeqNumRange(msg.Message.Header.SequenceNumber, msg.Message.Header.SequenceNumber)); err != nil {
		return 0, fmt.Errorf("message invalid msg %v: %w", msg, err)
	}

	lggr.Infow("chain reader returning latest onramp sequence number",
		"seqNum", msg.Message.Header.SequenceNumber, "sourceChainSelector", chain)
	return msg.SequenceNumber, nil
}

// GetExpectedNextSequenceNumber implements CCIP.
func (r *ccipChainReader) GetExpectedNextSequenceNumber(
	ctx context.Context,
	sourceChainSelector cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	if err := validateExtendedReaderExistence(r.contractReaders, sourceChainSelector); err != nil {
		return 0, err
	}

	var expectedNextSequenceNumber uint64
	err := r.contractReaders[sourceChainSelector].ExtendedGetLatestValue(
		ctx,
		consts.ContractNameOnRamp,
		consts.MethodNameGetExpectedNextSequenceNumber,
		primitives.Unconfirmed,
		map[string]any{
			"destChainSelector": r.destChain,
		},
		&expectedNextSequenceNumber,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get expected next sequence number from onramp, source chain: %d, dest chain: %d: %w",
			sourceChainSelector, r.destChain, err)
	}

	if expectedNextSequenceNumber == 0 {
		return 0, fmt.Errorf("the returned expected next sequence num is 0, source chain: %d, dest chain: %d",
			sourceChainSelector, r.destChain)
	}

	lggr.Debugw("chain reader returning expected next sequence number",
		"seqNum", expectedNextSequenceNumber, "sourceChainSelector", sourceChainSelector)
	return cciptypes.SeqNum(expectedNextSequenceNumber), nil
}

// NextSeqNum returns the current sequence numbers for chains.
// This always fetches fresh data directly from contracts to ensure accuracy.
// Critical for proper message sequencing.
func (r *ccipChainReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	// Use our direct fetch method that doesn't affect the cache
	cfgs, err := r.fetchFreshSourceChainConfigs(ctx, r.destChain, chains)
	if err != nil {
		return nil, fmt.Errorf("get source chains config: %w", err)
	}

	res := make(map[cciptypes.ChainSelector]cciptypes.SeqNum, len(chains))
	for _, chain := range chains {
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

type chainAddressNonce struct {
	chain    cciptypes.ChainSelector
	address  string
	response uint64
}

func (r *ccipChainReader) Nonces(
	ctx context.Context,
	addressesByChain map[cciptypes.ChainSelector][]string,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, err
	}

	// sort the input to ensure deterministic results
	sortedChains := maps.Keys(addressesByChain)
	slices.Sort(sortedChains)

	// create the structure that will contain our result
	res := make(map[cciptypes.ChainSelector]map[string]uint64)
	var addressCount int
	for _, addresses := range addressesByChain {
		addressCount += len(addresses)
	}

	contractInput, responses, err := prepareNoncesInput(lggr, addressesByChain, addressCount, sortedChains, r.addrCodec)
	if err != nil {
		return nil, err
	}

	request := contractreader.ExtendedBatchGetLatestValuesRequest{
		consts.ContractNameNonceManager: contractInput,
	}

	batchResult, _, err := r.contractReaders[r.destChain].ExtendedBatchGetLatestValues(
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

			sender, err := addrCodec.AddressStringToBytes(address, chain)
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

func (r *ccipChainReader) GetChainsFeeComponents(
	ctx context.Context,
	chains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	feeComponents := make(map[cciptypes.ChainSelector]types.ChainFeeComponents, len(r.contractWriters))

	for _, chain := range chains {
		chainWriter, ok := r.contractWriters[chain]
		if !ok {
			lggr.Errorw("contract writer not found", "chain", chain)
			continue
		}
		feeComponent, err := chainWriter.GetFeeComponents(ctx)
		if err != nil {
			lggr.Errorw("failed to get chain fee components", "chain", chain, "err", err)
			continue
		}

		if feeComponent.ExecutionFee == nil || feeComponent.ExecutionFee.Cmp(big.NewInt(0)) <= 0 {
			lggr.Errorw("execution fee is nil or non positive", "chain", chain)
			continue
		}
		if feeComponent.DataAvailabilityFee == nil || feeComponent.DataAvailabilityFee.Cmp(big.NewInt(0)) < 0 {
			lggr.Errorw("data availability fee is nil or negative", "chain", chain)
			continue
		}

		feeComponents[chain] = *feeComponent
	}
	return feeComponents
}

func (r *ccipChainReader) GetDestChainFeeComponents(
	ctx context.Context,
) (types.ChainFeeComponents, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	feeComponents := r.GetChainsFeeComponents(ctx, []cciptypes.ChainSelector{r.destChain})
	components, ok := feeComponents[r.destChain]
	if !ok {
		lggr.Errorw("dest chain components not found", "chain", r.destChain)
		return types.ChainFeeComponents{}, errors.New("dest chain fee components not found")
	}

	return components, nil
}

func (r *ccipChainReader) GetWrappedNativeTokenPriceUSD(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	// 1. Call chain's router to get native token address https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/Router.sol#L189:L191
	// 2. Call chain's FeeQuoter to get native tokens price  https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L229-L229
	//
	//nolint:lll
	prices := make(map[cciptypes.ChainSelector]cciptypes.BigInt)
	for _, chain := range selectors {
		reader, ok := r.contractReaders[chain]
		if !ok {
			lggr.Warnw("contract reader not found", "chain", chain)
			continue
		}

		config, err := r.configPoller.GetChainConfig(ctx, chain)
		if err != nil {
			lggr.Warnw("failed to get chain config for native token address", "chain", chain, "err", err)
			continue
		}
		nativeTokenAddress := config.Router.WrappedNativeAddress

		if cciptypes.UnknownAddress(nativeTokenAddress).IsZeroOrEmpty() {
			lggr.Warnw("Native token address is zero or empty. Ignore for disabled chains otherwise "+
				"check for router misconfiguration", "chain", chain, "address", nativeTokenAddress.String())
			continue
		}

		var update cciptypes.TimestampedUnixBig
		err = reader.ExtendedGetLatestValue(
			ctx,
			consts.ContractNameFeeQuoter,
			consts.MethodNameFeeQuoterGetTokenPrice,
			primitives.Unconfirmed,
			map[string]any{
				"token": nativeTokenAddress,
			},
			&update,
		)
		if err != nil {
			lggr.Errorw("failed to get native token price", "chain", chain, "err", err)
			continue
		}

		if update.Timestamp == 0 {
			lggr.Warnw("no native token price available", "chain", chain)
			continue
		}
		if update.Value == nil || update.Value.Cmp(big.NewInt(0)) <= 0 {
			lggr.Errorw("native token price is nil or non-positive", "chain", chain)
			continue
		}
		prices[chain] = cciptypes.NewBigInt(update.Value)
	}
	return prices
}

// GetChainFeePriceUpdate Read from Destination chain FeeQuoter latest fee updates for the provided chains.
// It unpacks the packed fee into the ChainFeeUSDPrices struct.
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L263-L263
//
//nolint:lll
func (r *ccipChainReader) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		lggr.Errorw("GetChainFeePriceUpdate dest chain extended reader not exist", "err", err)
		return nil
	}

	if len(selectors) == 0 {
		return make(map[cciptypes.ChainSelector]cciptypes.TimestampedBig) // Return a new empty map
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
	batchResult, _, err := r.contractReaders[r.destChain].ExtendedBatchGetLatestValues(
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
	return r.processFeePriceUpdateResults(lggr, selectors, feeQuoterResults)
}

// processFeePriceUpdateResults iterates through batch results, validates them,
// and returns a new feeUpdates map.
func (r *ccipChainReader) processFeePriceUpdateResults(
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

// buildSigners converts internal signer representation to RMN signer info format
func (r *ccipChainReader) buildSigners(signers []signer) []cciptypes.RemoteSignerInfo {
	result := make([]cciptypes.RemoteSignerInfo, 0, len(signers))
	for _, s := range signers {
		result = append(result, cciptypes.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		})
	}
	return result
}

func (r *ccipChainReader) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return cciptypes.RemoteConfig{}, fmt.Errorf("get chain config: %w", err)
	}

	// RMNRemote address stored in the offramp static config is actually the proxy contract address.
	// Here we will get the RMNRemote address from the proxy contract by calling the RMNProxy contract.
	proxyContractAddress, err := r.GetContractAddress(consts.ContractNameRMNRemote, r.destChain)
	if err != nil {
		return cciptypes.RemoteConfig{}, fmt.Errorf("get RMNRemote proxy contract address: %w", err)
	}

	rmnRemoteAddress, err := r.getRMNRemoteAddress(ctx, lggr, r.destChain, proxyContractAddress)
	if err != nil {
		return cciptypes.RemoteConfig{}, fmt.Errorf("get RMNRemote address: %w", err)
	}

	return cciptypes.RemoteConfig{
		ContractAddress:  rmnRemoteAddress,
		ConfigDigest:     config.RMNRemote.VersionedConfig.Config.RMNHomeContractConfigDigest,
		Signers:          r.buildSigners(config.RMNRemote.VersionedConfig.Config.Signers),
		FSign:            config.RMNRemote.VersionedConfig.Config.FSign,
		ConfigVersion:    config.RMNRemote.VersionedConfig.Version,
		RmnReportVersion: config.RMNRemote.DigestHeader.DigestHeader,
	}, nil
}

// GetRmnCurseInfo returns rmn curse/pausing information about the provided chains
// from the destination chain RMN remote contract.
func (r *ccipChainReader) GetRmnCurseInfo(ctx context.Context) (CurseInfo, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return CurseInfo{}, fmt.Errorf("validate dest=%d extended reader existence: %w", r.destChain, err)
	}

	// TODO: Curse requires a dedicated cache, but for now fetching it in background,
	// together with the other configurations
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return CurseInfo{}, fmt.Errorf("get chain config: %w", err)
	}

	return config.CurseInfo, nil
}

func getCurseInfoFromCursedSubjects(
	cursedSubjectsSet mapset.Set[[16]byte],
	destChainSelector cciptypes.ChainSelector,
) *CurseInfo {
	curseInfo := &CurseInfo{
		CursedSourceChains: make(map[cciptypes.ChainSelector]bool, cursedSubjectsSet.Cardinality()),
		CursedDestination: cursedSubjectsSet.Contains(GlobalCurseSubject) ||
			cursedSubjectsSet.Contains(chainSelectorToBytes16(destChainSelector)),
		GlobalCurse: cursedSubjectsSet.Contains(GlobalCurseSubject),
	}

	for _, cursedSubject := range cursedSubjectsSet.ToSlice() {
		if cursedSubject == GlobalCurseSubject {
			continue
		}

		chainSelector := cciptypes.ChainSelector(binary.BigEndian.Uint64(cursedSubject[8:]))
		if chainSelector == destChainSelector {
			continue
		}

		curseInfo.CursedSourceChains[chainSelector] = true
	}
	return curseInfo
}

func chainSelectorToBytes16(chainSel cciptypes.ChainSelector) [16]byte {
	var result [16]byte
	// Convert the uint64 to bytes and place it in the last 8 bytes of the array
	binary.BigEndian.PutUint64(result[8:], uint64(chainSel))
	return result
}

// discoverOffRampContracts uses the offRamp for destChain to discover the addresses of other contracts.
func (r *ccipChainReader) discoverOffRampContracts(
	ctx context.Context,
	lggr logger.Logger,
	chains []cciptypes.ChainSelector,
) (ContractAddresses, error) {
	// Get from cache
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return nil, fmt.Errorf("unable to lookup RMN remote address (RMN proxy): %w", err)
	}

	resp := make(ContractAddresses)

	// OnRamps are in the offRamp SourceChainConfig.
	{
		sourceConfigs, err := r.getOffRampSourceChainsConfig(ctx, lggr, chains, false)

		if err != nil {
			return nil, fmt.Errorf("unable to get SourceChainsConfig: %w", err)
		}

		// Iterate results in sourceChain selector order so that the router config is deterministic.
		keys := maps.Keys(sourceConfigs)
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, sourceChain := range keys {
			cfg := sourceConfigs[sourceChain]
			resp = resp.Append(consts.ContractNameOnRamp, sourceChain, cfg.OnRamp)
			// The local router is located in each source sourceChain config. Add it once.
			if len(resp[consts.ContractNameRouter][r.destChain]) == 0 {
				resp = resp.Append(consts.ContractNameRouter, r.destChain, cfg.Router)
				lggr.Infow("appending router contract address", "address", cfg.Router)
			}
		}
	}

	// Add static config contracts
	if len(config.Offramp.StaticConfig.RmnRemote) > 0 {
		lggr.Infow("appending RMN remote contract address",
			"address", hex.EncodeToString(config.Offramp.StaticConfig.RmnRemote),
			"chain", r.destChain)
		resp = resp.Append(consts.ContractNameRMNRemote, r.destChain, config.Offramp.StaticConfig.RmnRemote)
	}

	if len(config.Offramp.StaticConfig.NonceManager) > 0 {
		resp = resp.Append(consts.ContractNameNonceManager, r.destChain, config.Offramp.StaticConfig.NonceManager)
	}

	// Add dynamic config contracts
	if len(config.Offramp.DynamicConfig.FeeQuoter) > 0 {
		lggr.Infow("appending fee quoter contract address",
			"address", hex.EncodeToString(config.Offramp.DynamicConfig.FeeQuoter),
			"chain", r.destChain)
		resp = resp.Append(consts.ContractNameFeeQuoter, r.destChain, config.Offramp.DynamicConfig.FeeQuoter)
	}

	return resp, nil
}

func (r *ccipChainReader) DiscoverContracts(ctx context.Context,
	chains []cciptypes.ChainSelector,
) (ContractAddresses, error) {
	var resp ContractAddresses
	lggr := logutil.WithContextValues(ctx, r.lggr)

	// Discover destination contracts if the dest chain is supported.
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err == nil {
		resp, err = r.discoverOffRampContracts(ctx, lggr, chains)
		// Can't continue with discovery if the destination chain is not available.
		// We read source chains OnRamps from there, and onRamps are essential for feeQuoter and Router discovery.
		if err != nil {
			return nil, fmt.Errorf("discover destination contracts: %w", err)
		}
	}

	// The following calls are on dynamically configured chains which may not
	// be available when this function is called. Eventually they will be
	// configured through consensus when the Sync function is called, but until
	// that happens the ErrNoBindings case must be handled gracefully.

	myChains := maps.Keys(r.contractReaders)

	// Use wait group for parallel processing
	var wg sync.WaitGroup
	mu := new(sync.Mutex)

	// Process each source chain's OnRamp configurations
	for _, chain := range myChains {
		if chain == r.destChain {
			continue
		}

		// Check if we have a reader for this chain
		if _, exists := r.contractReaders[chain]; !exists {
			lggr.Debugw("Contract reader not found for chain", "chain", chain)
			continue
		}

		chainCopy := chain
		wg.Add(1)
		go func(chainSel cciptypes.ChainSelector) {
			defer wg.Done()

			// Get cached OnRamp configurations
			config, err := r.configPoller.GetChainConfig(ctx, chainSel)
			if err != nil {
				lggr.Errorw("Failed to get chain config",
					"chain", chainSel,
					"err", err)
				return
			}

			// Use mutex to safely update the shared resp
			mu.Lock()
			defer mu.Unlock()

			// Add FeeQuoter from dynamic config
			if !cciptypes.UnknownAddress(config.OnRamp.DynamicConfig.DynamicConfig.FeeQuoter).IsZeroOrEmpty() {
				resp = resp.Append(
					consts.ContractNameFeeQuoter,
					chainSel,
					config.OnRamp.DynamicConfig.DynamicConfig.FeeQuoter)
			}

			// Add Router from dest chain config
			if !cciptypes.UnknownAddress(config.OnRamp.DestChainConfig.Router).IsZeroOrEmpty() {
				resp = resp.Append(
					consts.ContractNameRouter,
					chainSel,
					config.OnRamp.DestChainConfig.Router)
			}
		}(chainCopy)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return resp, nil
}

// Sync goes through the input contracts and binds them to the contract reader.
func (r *ccipChainReader) Sync(ctx context.Context, contracts ContractAddresses) error {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	var errs []error
	for contractName, chainSelToAddress := range contracts {
		for chainSel, address := range chainSelToAddress {
			// defense in depth: don't bind if the address is empty.
			// callers should ensure this but we double check here.
			if len(address) == 0 {
				lggr.Warnw("skipping binding empty address for contract",
					"contractName", contractName,
					"chainSel", chainSel,
				)
				continue
			}

			// try to bind
			_, err := bindExtendedReaderContract(ctx, lggr, r.contractReaders, chainSel, contractName, address, r.addrCodec)
			if err != nil {
				if errors.Is(err, ErrContractReaderNotFound) {
					// don't support this chain
					continue
				}
				// some other error, gather
				// TODO: maybe return early?
				errs = append(errs, err)
			}
		}
	}

	return errors.Join(errs...)
}

func (r *ccipChainReader) GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error) {
	extendedReader, ok := r.contractReaders[chain]
	if !ok {
		return nil, fmt.Errorf("contract reader not found for chain %d", chain)
	}

	bindings := extendedReader.GetBindings(contractName)
	if len(bindings) != 1 {
		return nil, fmt.Errorf("expected one binding for the %s contract, got %d", contractName, len(bindings))
	}

	addressBytes, err := r.addrCodec.AddressStringToBytes(bindings[0].Binding.Address, chain)
	if err != nil {
		return nil, fmt.Errorf("convert address %s to bytes: %w", bindings[0].Binding.Address, err)
	}

	return addressBytes, nil
}

// LinkPriceUSD gets the LINK price in 1e-18 USDs from the FeeQuoter contract on the destination chain.
// For example, if the price is 1 LINK = 10 USD, this function will return 10e18 (10 * 1e18). You can think of this
// function returning the price of LINK not in USD, but in a small denomination of USD, similar to returning
// the price of ETH not in ETH but in wei (1e-18 ETH).
func (r *ccipChainReader) LinkPriceUSD(ctx context.Context) (cciptypes.BigInt, error) {
	// Ensure we can read from the destination chain.
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return cciptypes.BigInt{}, fmt.Errorf("failed to validate dest chain reader existence: %w", err)
	}

	// TODO: consider caching this value.
	feeQuoterCfg, err := r.getDestFeeQuoterStaticConfig(ctx)
	if err != nil {
		return cciptypes.BigInt{}, fmt.Errorf("get destination fee quoter static config: %w", err)
	}

	linkPriceUSD, err := r.getFeeQuoterTokenPriceUSD(ctx, feeQuoterCfg.LinkToken)
	if err != nil {
		return cciptypes.BigInt{}, fmt.Errorf("get LINK price in USD: %w", err)
	}

	if linkPriceUSD.Int == nil {
		return cciptypes.BigInt{}, fmt.Errorf("LINK price is nil")
	}

	if linkPriceUSD.Int.Cmp(big.NewInt(0)) == 0 {
		return cciptypes.BigInt{}, fmt.Errorf("LINK price is 0")
	}

	return linkPriceUSD, nil
}

// feeQuoterStaticConfig is used to parse the response from the feeQuoter contract's getStaticConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/FeeQuoter.sol#L946
//
//nolint:lll // It's a URL.
type feeQuoterStaticConfig struct {
	MaxFeeJuelsPerMsg  cciptypes.BigInt `json:"maxFeeJuelsPerMsg"`
	LinkToken          []byte           `json:"linkToken"`
	StalenessThreshold uint32           `json:"stalenessThreshold"`
}

// getDestFeeQuoterStaticConfig returns the destination chain's Fee Quoter's StaticConfig
func (r *ccipChainReader) getDestFeeQuoterStaticConfig(ctx context.Context) (feeQuoterStaticConfig, error) {
	// Get from cache
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return feeQuoterStaticConfig{}, fmt.Errorf("get chain config: %w", err)
	}

	if len(config.FeeQuoter.StaticConfig.LinkToken) == 0 {
		return feeQuoterStaticConfig{}, fmt.Errorf("link token address is empty")
	}

	return config.FeeQuoter.StaticConfig, nil
}

// getFeeQuoterTokenPriceUSD gets the token price in USD of the given token address from the FeeQuoter contract on the
// destination chain.
func (r *ccipChainReader) getFeeQuoterTokenPriceUSD(ctx context.Context, tokenAddr []byte) (cciptypes.BigInt, error) {
	if len(tokenAddr) == 0 {
		return cciptypes.BigInt{}, fmt.Errorf("tokenAddr is empty")
	}

	reader, ok := r.contractReaders[r.destChain]
	if !ok {
		return cciptypes.BigInt{}, fmt.Errorf("contract reader not found for chain %d", r.destChain)
	}

	var timestampedPrice cciptypes.TimestampedUnixBig
	err := reader.ExtendedGetLatestValue(
		ctx,
		consts.ContractNameFeeQuoter,
		consts.MethodNameFeeQuoterGetTokenPrice,
		primitives.Unconfirmed,
		map[string]any{
			"token": tokenAddr,
		},
		&timestampedPrice,
	)
	if err != nil {
		return cciptypes.BigInt{}, fmt.Errorf("failed to get token price, addr: %v, err: %w", tokenAddr, err)
	}

	price := timestampedPrice.Value

	if price == nil {
		return cciptypes.BigInt{}, fmt.Errorf("token price is nil,  addr: %v", tokenAddr)
	}
	if price.Cmp(big.NewInt(0)) == 0 {
		return cciptypes.BigInt{}, fmt.Errorf("token price is 0, addr: %v", tokenAddr)
	}

	return cciptypes.NewBigInt(price), nil
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

// GetOffRampSourceChainsConfig returns the static source chain configs for all the provided source chains.
// This method returns configurations without the MinSeqNr field, which should be fetched separately when needed.
func (r *ccipChainReader) GetOffRampSourceChainsConfig(ctx context.Context, chains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	return r.getOffRampSourceChainsConfig(ctx, r.lggr, chains, true)
}

// getOffRampSourceChainsConfig gets static source chain configs from the configPoller cache.
// These configs deliberately exclude MinSeqNr to prevent usage of potentially stale sequence numbers.
// For obtaining fresh sequence numbers, use ccipChainReader.GetLatestMinSeqNrs.
//
//nolint:revive
func (r *ccipChainReader) getOffRampSourceChainsConfig(
	ctx context.Context,
	lggr logger.Logger,
	chains []cciptypes.ChainSelector,
	includeDisabled bool,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, fmt.Errorf("validate extended reader existence: %w", err)
	}

	// Use the ConfigPoller to handle caching
	configs, err := r.configPoller.GetOfframpSourceChainConfigs(ctx, r.destChain, chains)
	if err != nil {
		return nil, fmt.Errorf("get source chain configs: %w", err)
	}

	// Filter out disabled chains if needed
	if !includeDisabled {
		for chain, cfg := range configs {
			enabled, err := cfg.check()
			if err != nil {
				return nil, fmt.Errorf("source chain config check for chain %d failed: %w", chain, err)
			}
			if !enabled {
				lggr.Debugw("Filtering out disabled source chain",
					"chain", chain,
					"error", err,
					"enabled", enabled)
				delete(configs, chain)
			}
		}
	}

	return configs, nil
}

// fetchFreshSourceChainConfigs always fetches fresh source chain configs directly from contracts
// without using any cached values. Use this when up-to-date data is critical, especially
// for sequence number accuracy.
func (r *ccipChainReader) fetchFreshSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]SourceChainConfig, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	reader, exists := r.contractReaders[destChain]
	if !exists {
		return nil, fmt.Errorf("no contract reader for chain %d", destChain)
	}

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
	results, _, err := reader.ExtendedBatchGetLatestValues(
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

// offRampStaticChainConfig is used to parse the response from the offRamp contract's getStaticConfig method.
// See: <chainlink repo>/contracts/src/v0.8/ccip/offRamp/OffRamp.sol:StaticConfig
type offRampStaticChainConfig struct {
	ChainSelector        cciptypes.ChainSelector `json:"chainSelector"`
	GasForCallExactCheck uint16                  `json:"gasForCallExactCheck"`
	RmnRemote            []byte                  `json:"rmnRemote"`
	TokenAdminRegistry   []byte                  `json:"tokenAdminRegistry"`
	NonceManager         []byte                  `json:"nonceManager"`
}

// offRampDynamicChainConfig maps to DynamicConfig in OffRamp.sol
type offRampDynamicChainConfig struct {
	FeeQuoter                               []byte `json:"feeQuoter"`
	PermissionLessExecutionThresholdSeconds uint32 `json:"permissionLessExecutionThresholdSeconds"`
	IsRMNVerificationDisabled               bool   `json:"isRMNVerificationDisabled"`
	MessageInterceptor                      []byte `json:"messageInterceptor"`
}

// See DynamicChainConfig in OnRamp.sol
type onRampDynamicConfig struct {
	FeeQuoter              []byte `json:"feeQuoter"`
	ReentrancyGuardEntered bool   `json:"reentrancyGuardEntered"`
	MessageInterceptor     []byte `json:"messageInterceptor"`
	FeeAggregator          []byte `json:"feeAggregator"`
	AllowListAdmin         []byte `json:"allowListAdmin"`
}

// We're wrapping the onRampDynamicConfig this way to map to on-chain return type which is a named struct
// https://github.com/smartcontractkit/chainlink/blob/12af1de88238e0e918177d6b5622070417f48adf/contracts/src/v0.8/ccip/onRamp/OnRamp.sol#L328
//
//nolint:lll
type getOnRampDynamicConfigResponse struct {
	DynamicConfig onRampDynamicConfig `json:"dynamicConfig"`
}

// See DestChainConfig in OnRamp.sol
type onRampDestChainConfig struct {
	SequenceNumber   uint64 `json:"sequenceNumber"`
	AllowListEnabled bool   `json:"allowListEnabled"`
	Router           []byte `json:"router"`
}

// signer is used to parse the response from the RMNRemote contract's getVersionedConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/rmn/RMNRemote.sol#L42-L45
type signer struct {
	OnchainPublicKey []byte `json:"onchainPublicKey"`
	NodeIndex        uint64 `json:"nodeIndex"`
}

// config is used to parse the response from the RMNRemote contract's getVersionedConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/rmn/RMNRemote.sol#L49-L53
type config struct {
	RMNHomeContractConfigDigest cciptypes.Bytes32 `json:"rmnHomeContractConfigDigest"`
	Signers                     []signer          `json:"signers"`
	FSign                       uint64            `json:"fSign"` // previously: MinSigners
}

// versionedConfig is used to parse the response from the RMNRemote contract's getVersionedConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/rmn/RMNRemote.sol#L167-L169
type versionedConfig struct {
	Version uint32 `json:"version"`
	Config  config `json:"config"`
}

// getARM gets the RMN remote address from the RMN proxy address.
// See: https://github.com/smartcontractkit/chainlink/blob/3c7817c566c5d0aa14519c679fa85b227ac97cc5/contracts/src/v0.8/ccip/rmn/ARMProxy.sol#L40-L44
//
//nolint:lll
func (r *ccipChainReader) getRMNRemoteAddress(
	ctx context.Context,
	lggr logger.Logger,
	chain cciptypes.ChainSelector,
	rmnRemoteProxyAddress []byte) ([]byte, error) {
	_, err := bindExtendedReaderContract(ctx, lggr, r.contractReaders, chain, consts.ContractNameRMNProxy, rmnRemoteProxyAddress, r.addrCodec)
	if err != nil {
		return nil, fmt.Errorf("bind RMN proxy contract: %w", err)
	}

	// Get the address from cache instead of making a contract call
	config, err := r.configPoller.GetChainConfig(ctx, chain)
	if err != nil {
		return nil, fmt.Errorf("get chain config: %w", err)
	}

	return config.RMNProxy.RemoteAddress, nil
}

func (r *ccipChainReader) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return 0, fmt.Errorf("validate dest=%d extended reader existence: %w", r.destChain, err)
	}

	var latestSeqNr uint64
	err := r.contractReaders[r.destChain].ExtendedGetLatestValue(
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

func (r *ccipChainReader) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return [32]byte{}, fmt.Errorf("get chain config: %w", err)
	}

	var resp OCRConfigResponse
	if pluginType == consts.PluginTypeCommit {
		resp = config.Offramp.CommitLatestOCRConfig
	} else {
		resp = config.Offramp.ExecLatestOCRConfig
	}

	return resp.OCRConfig.ConfigInfo.ConfigDigest, nil
}

func (r *ccipChainReader) prepareBatchConfigRequests(
	chainSel cciptypes.ChainSelector) contractreader.ExtendedBatchGetLatestValuesRequest {

	var (
		commitLatestOCRConfig OCRConfigResponse
		execLatestOCRConfig   OCRConfigResponse
		staticConfig          offRampStaticChainConfig
		dynamicConfig         offRampDynamicChainConfig
		rmnRemoteAddress      []byte
		rmnDigestHeader       rmnDigestHeader
		rmnVersionConfig      versionedConfig
		feeQuoterConfig       feeQuoterStaticConfig
		onRampDynamicConfig   getOnRampDynamicConfigResponse
		onRampDestConfig      onRampDestChainConfig
		wrappedNativeAddress  []byte
		cursedSubjects        RMNCurseResponse
	)

	var requests contractreader.ExtendedBatchGetLatestValuesRequest

	// Only add OnRamp config requests if this is a source chain (not destination chain)
	if chainSel != r.destChain {
		requests = contractreader.ExtendedBatchGetLatestValuesRequest{
			consts.ContractNameOnRamp: {
				{
					ReadName:  consts.MethodNameOnRampGetDynamicConfig,
					Params:    map[string]any{},
					ReturnVal: &onRampDynamicConfig,
				},
				{
					ReadName: consts.MethodNameOnRampGetDestChainConfig,
					Params: map[string]any{
						"destChainSelector": r.destChain,
					},
					ReturnVal: &onRampDestConfig,
				},
			},
			consts.ContractNameRouter: {
				{
					ReadName:  consts.MethodNameRouterGetWrappedNative,
					Params:    map[string]any{},
					ReturnVal: &wrappedNativeAddress,
				},
			},
		}
	} else {
		// Add all other contract requests for the destination chain
		requests = contractreader.ExtendedBatchGetLatestValuesRequest{
			consts.ContractNameOffRamp: {
				{
					ReadName: consts.MethodNameOffRampLatestConfigDetails,
					Params: map[string]any{
						"ocrPluginType": consts.PluginTypeCommit,
					},
					ReturnVal: &commitLatestOCRConfig,
				},
				{
					ReadName: consts.MethodNameOffRampLatestConfigDetails,
					Params: map[string]any{
						"ocrPluginType": consts.PluginTypeExecute,
					},
					ReturnVal: &execLatestOCRConfig,
				},
				{
					ReadName:  consts.MethodNameOffRampGetStaticConfig,
					Params:    map[string]any{},
					ReturnVal: &staticConfig,
				},
				{
					ReadName:  consts.MethodNameOffRampGetDynamicConfig,
					Params:    map[string]any{},
					ReturnVal: &dynamicConfig,
				},
			},
			consts.ContractNameRMNProxy: {{
				ReadName:  consts.MethodNameGetARM,
				Params:    map[string]any{},
				ReturnVal: &rmnRemoteAddress,
			}},
			consts.ContractNameRMNRemote: {
				{
					ReadName:  consts.MethodNameGetReportDigestHeader,
					Params:    map[string]any{},
					ReturnVal: &rmnDigestHeader,
				},
				{
					ReadName:  consts.MethodNameGetVersionedConfig,
					Params:    map[string]any{},
					ReturnVal: &rmnVersionConfig,
				},
				{
					ReadName:  consts.MethodNameGetCursedSubjects,
					Params:    map[string]any{},
					ReturnVal: &cursedSubjects,
				},
			},
			consts.ContractNameFeeQuoter: {{
				ReadName:  consts.MethodNameFeeQuoterGetStaticConfig,
				Params:    map[string]any{},
				ReturnVal: &feeQuoterConfig,
			}},
		}
	}

	return requests
}

func (r *ccipChainReader) processConfigResults(
	chainSel cciptypes.ChainSelector,
	batchResult types.BatchGetLatestValuesResult) (ChainConfigSnapshot, error) {
	config := ChainConfigSnapshot{}

	for contract, results := range batchResult {
		var err error
		switch contract.Name {
		case consts.ContractNameOffRamp:
			config.Offramp, err = r.processOfframpResults(results)
		case consts.ContractNameRMNProxy:
			config.RMNProxy, err = r.processRMNProxyResults(results)
		case consts.ContractNameRMNRemote:
			config.RMNRemote, config.CurseInfo, err = r.processRMNRemoteResults(results)
		case consts.ContractNameFeeQuoter:
			config.FeeQuoter, err = r.processFeeQuoterResults(results)
		case consts.ContractNameOnRamp:
			// Only process OnRamp results for source chains
			if chainSel != r.destChain {
				config.OnRamp, err = r.processOnRampResults(results)
			}
		case consts.ContractNameRouter:
			// Only process Router results for source chains
			if chainSel != r.destChain {
				config.Router, err = r.processRouterResults(results)
			}
		default:
			r.lggr.Warnw("Unhandled contract in batch results", "contract", contract.Name)
		}
		if err != nil {
			return ChainConfigSnapshot{}, fmt.Errorf("process %s results: %w", contract.Name, err)
		}
	}

	return config, nil
}

func (r *ccipChainReader) processRouterResults(results []types.BatchReadResult) (RouterConfig, error) {
	if len(results) != 1 {
		return RouterConfig{}, fmt.Errorf("expected 1 router result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return RouterConfig{}, fmt.Errorf("get router wrapped native result: %w", err)
	}

	if bytes, ok := val.(*[]byte); ok {
		return RouterConfig{
			WrappedNativeAddress: cciptypes.Bytes(*bytes),
		}, nil
	}

	return RouterConfig{}, fmt.Errorf("invalid type for router wrapped native address: %T", val)
}

func (r *ccipChainReader) processOnRampResults(results []types.BatchReadResult) (OnRampConfig, error) {
	if len(results) != 2 {
		return OnRampConfig{}, fmt.Errorf("expected 2 OnRamp results, got %d", len(results))
	}

	var config OnRampConfig

	// Process DynamicConfig
	val, err := results[0].GetResult()
	if err != nil {
		return OnRampConfig{}, fmt.Errorf("get OnRamp dynamic config result: %w", err)
	}

	dynamicConfig, ok := val.(*getOnRampDynamicConfigResponse)
	if !ok {
		return OnRampConfig{}, fmt.Errorf("invalid type for OnRamp dynamic config: %T", val)
	}
	config.DynamicConfig = *dynamicConfig

	// Process DestChainConfig
	val, err = results[1].GetResult()
	if err != nil {
		return OnRampConfig{}, fmt.Errorf("get OnRamp dest chain config result: %w", err)
	}

	destConfig, ok := val.(*onRampDestChainConfig)
	if !ok {
		return OnRampConfig{}, fmt.Errorf("invalid type for OnRamp dest chain config: %T", val)
	}
	config.DestChainConfig = *destConfig

	return config, nil
}

// GetOnRampConfig returns the cached OnRamp configurations for a source chain
func (r *ccipChainReader) GetOnRampConfig(ctx context.Context, srcChain cciptypes.ChainSelector) (OnRampConfig, error) {
	if srcChain == r.destChain {
		return OnRampConfig{}, fmt.Errorf("cannot get OnRamp configs for destination chain %d", srcChain)
	}

	config, err := r.configPoller.GetChainConfig(ctx, srcChain)
	if err != nil {
		return OnRampConfig{}, fmt.Errorf("get chain config: %w", err)
	}

	return config.OnRamp, nil
}

func (r *ccipChainReader) processOfframpResults(
	results []types.BatchReadResult) (OfframpConfig, error) {

	if len(results) != 4 {
		return OfframpConfig{}, fmt.Errorf("expected 4 offramp results, got %d", len(results))
	}

	config := OfframpConfig{}

	// Define processors for each expected result
	processors := []resultProcessor{
		// CommitLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for CommitLatestOCRConfig: %T", val)
			}
			config.CommitLatestOCRConfig = *typed
			return nil
		},
		// ExecLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for ExecLatestOCRConfig: %T", val)
			}
			config.ExecLatestOCRConfig = *typed
			return nil
		},
		// StaticConfig
		func(val interface{}) error {
			typed, ok := val.(*offRampStaticChainConfig)
			if !ok {
				return fmt.Errorf("invalid type for StaticConfig: %T", val)
			}
			config.StaticConfig = *typed
			return nil
		},
		// DynamicConfig
		func(val interface{}) error {
			typed, ok := val.(*offRampDynamicChainConfig)
			if !ok {
				return fmt.Errorf("invalid type for DynamicConfig: %T", val)
			}
			config.DynamicConfig = *typed
			return nil
		},
	}

	// Process each result with its corresponding processor
	for i, result := range results {
		val, err := result.GetResult()
		if err != nil {
			return OfframpConfig{}, fmt.Errorf("get offramp result %d: %w", i, err)
		}

		if err := processors[i](val); err != nil {
			return OfframpConfig{}, fmt.Errorf("process result %d: %w", i, err)
		}
	}

	return config, nil
}

func (r *ccipChainReader) processRMNProxyResults(results []types.BatchReadResult) (RMNProxyConfig, error) {
	if len(results) != 1 {
		return RMNProxyConfig{}, fmt.Errorf("expected 1 RMN proxy result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return RMNProxyConfig{}, fmt.Errorf("get RMN proxy result: %w", err)
	}

	if bytes, ok := val.(*[]byte); ok {
		return RMNProxyConfig{
			RemoteAddress: *bytes,
		}, nil
	}

	return RMNProxyConfig{}, fmt.Errorf("invalid type for RMN proxy remote address: %T", val)
}

func (r *ccipChainReader) processRMNRemoteResults(results []types.BatchReadResult) (
	RMNRemoteConfig,
	CurseInfo,
	error,
) {
	config := RMNRemoteConfig{}

	if len(results) != 3 {
		return RMNRemoteConfig{}, CurseInfo{}, fmt.Errorf("expected 3 RMN remote results, got %d", len(results))
	}

	// Process DigestHeader
	val, err := results[0].GetResult()
	if err != nil {
		return RMNRemoteConfig{}, CurseInfo{}, fmt.Errorf("get RMN remote digest header result: %w", err)
	}

	typed, ok := val.(*rmnDigestHeader)
	if !ok {
		return RMNRemoteConfig{}, CurseInfo{}, fmt.Errorf("invalid type for RMN remote digest header: %T", val)
	}
	config.DigestHeader = *typed

	// Process VersionedConfig
	val, err = results[1].GetResult()
	if err != nil {
		return RMNRemoteConfig{}, CurseInfo{}, fmt.Errorf("get RMN remote versioned config result: %w", err)
	}

	vconf, ok := val.(*versionedConfig)
	if !ok {
		return RMNRemoteConfig{}, CurseInfo{}, fmt.Errorf("invalid type for RMN remote versioned config: %T", val)
	}
	config.VersionedConfig = *vconf

	// Process CursedSubjects
	val, err = results[2].GetResult()
	if err != nil {
		return RMNRemoteConfig{}, CurseInfo{}, fmt.Errorf("get RMN remote cursed subjects result: %w", err)
	}

	c, ok := val.(*RMNCurseResponse)
	if !ok {
		return RMNRemoteConfig{}, CurseInfo{}, fmt.Errorf("invalid type for RMN remote cursed subjects: %T", val)
	}
	curseInfo := *getCurseInfoFromCursedSubjects(
		mapset.NewSet(c.CursedSubjects...),
		r.destChain,
	)

	return config, curseInfo, nil
}

func (r *ccipChainReader) processFeeQuoterResults(results []types.BatchReadResult) (FeeQuoterConfig, error) {
	if len(results) != 1 {
		return FeeQuoterConfig{}, fmt.Errorf("expected 1 fee quoter result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return FeeQuoterConfig{}, fmt.Errorf("get fee quoter result: %w", err)
	}

	if typed, ok := val.(*feeQuoterStaticConfig); ok {
		return FeeQuoterConfig{
			StaticConfig: *typed,
		}, nil
	}

	return FeeQuoterConfig{}, fmt.Errorf("invalid type for fee quoter static config: %T", val)
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

// ccipReaderInternal defines the interface that ConfigPoller needs from the ccipChainReader
// This allows for better encapsulation and easier testing through mocking
type ccipReaderInternal interface {
	// getDestChain returns the destination chain selector
	getDestChain() cciptypes.ChainSelector

	// getContractReader returns the contract reader for the specified chain
	getContractReader(chain cciptypes.ChainSelector) (contractreader.Extended, bool)

	// prepareBatchConfigRequests prepares the batch requests for fetching chain configuration
	prepareBatchConfigRequests(chainSel cciptypes.ChainSelector) contractreader.ExtendedBatchGetLatestValuesRequest

	// processConfigResults processes the batch results into a ChainConfigSnapshot
	processConfigResults(
		chainSel cciptypes.ChainSelector,
		batchResult types.BatchGetLatestValuesResult) (ChainConfigSnapshot, error)

	// fetchFreshSourceChainConfigs fetches source chain configurations from the specified destination chain
	fetchFreshSourceChainConfigs(
		ctx context.Context, destChain cciptypes.ChainSelector,
		sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]SourceChainConfig, error)
}

// getDestChain returns the destination chain selector
func (r *ccipChainReader) getDestChain() cciptypes.ChainSelector {
	return r.destChain
}

// getContractReader returns the contract reader for the specified chain
func (r *ccipChainReader) getContractReader(chain cciptypes.ChainSelector) (contractreader.Extended, bool) {
	reader, exists := r.contractReaders[chain]
	return reader, exists
}

// Interface compliance check
var _ CCIPReader = (*ccipChainReader)(nil)
