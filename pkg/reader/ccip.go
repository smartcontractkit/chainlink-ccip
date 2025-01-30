package reader

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// TODO: unit test the implementation when the actual contract reader and writer interfaces are finalized and mocks
// can be generated.
type ccipChainReader struct {
	lggr            logger.Logger
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter
	destChain       cciptypes.ChainSelector
	offrampAddress  string
	extraDataCodec  cciptypes.ExtraDataCodec
}

func newCCIPChainReaderInternal(
	ctx context.Context,
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
	extraDataCodec cciptypes.ExtraDataCodec,
) *ccipChainReader {
	var crs = make(map[cciptypes.ChainSelector]contractreader.Extended)
	for chainSelector, cr := range contractReaders {
		crs[chainSelector] = contractreader.NewExtendedContractReader(cr)
	}

	reader := &ccipChainReader{
		lggr:            lggr,
		contractReaders: crs,
		contractWriters: contractWriters,
		destChain:       destChain,
		offrampAddress:  typeconv.AddressBytesToString(offrampAddress, uint64(destChain)),
		extraDataCodec:  extraDataCodec,
	}

	contracts := ContractAddresses{
		consts.ContractNameOffRamp: {
			destChain: offrampAddress,
		},
	}
	if err := reader.Sync(ctx, contracts); err != nil {
		lggr.Errorw("failed to sync contracts", "err", err)
	}

	return reader
}

// WithExtendedContractReader sets the extended contract reader for the provided chain.
func (r *ccipChainReader) WithExtendedContractReader(
	ch cciptypes.ChainSelector, cr contractreader.Extended) *ccipChainReader {
	r.contractReaders[ch] = cr
	return r
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
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type PriceUpdates struct {
	TokenPriceUpdates []TokenPriceUpdate
	GasPriceUpdates   []GasPriceUpdate
}

type CommitReportAcceptedEvent struct {
	MerkleRoots  []MerkleRoot
	PriceUpdates PriceUpdates
}

// ---------------------------------------------------

func (r *ccipChainReader) CommitReportsGTETimestamp(
	ctx context.Context, ts time.Time, limit int,
) ([]plugintypes2.CommitPluginReportWithMeta, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, err
	}

	ev := CommitReportAcceptedEvent{}

	iter, err := r.contractReaders[r.destChain].ExtendedQueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameCommitReportAccepted,
			Expressions: []query.Expression{
				query.Timestamp(uint64(ts.Unix()), primitives.Gte),
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortBySequence(query.Asc)},
			Limit: query.Limit{
				Count: uint64(limit * 2),
			},
		},
		&ev,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}
	lggr.Debugw("queried commit reports", "numReports", len(iter),
		"destChain", r.destChain,
		"ts", ts,
		"limit", limit*2)

	reports := make([]plugintypes2.CommitPluginReportWithMeta, 0)
	for _, item := range iter {
		ev, is := (item.Data).(*CommitReportAcceptedEvent)
		if !is {
			return nil, fmt.Errorf("unexpected type %T while expecting a commit report", item)
		}

		if err := validateCommitReportAcceptedEvent(*ev, item, ts); err != nil {
			lggr.Errorw("validate commit report accepted event", "err", err, "ev", ev)
			continue
		}

		lggr.Debugw("processing commit report", "report", ev, "item", item)

		merkleRoots := make([]cciptypes.MerkleRootChain, 0, len(ev.MerkleRoots))
		for _, mr := range ev.MerkleRoots {
			onRampAddress, err := r.GetContractAddress(
				consts.ContractNameOnRamp,
				cciptypes.ChainSelector(mr.SourceChainSelector),
			)
			if err != nil {
				r.lggr.Errorw("get onRamp address for selector %d: %w", mr.SourceChainSelector, err)
				continue
			}
			merkleRoots = append(merkleRoots, cciptypes.MerkleRootChain{
				ChainSel:      cciptypes.ChainSelector(mr.SourceChainSelector),
				OnRampAddress: onRampAddress,
				SeqNumsRange: cciptypes.NewSeqNumRange(
					cciptypes.SeqNum(mr.MinSeqNr),
					cciptypes.SeqNum(mr.MaxSeqNr),
				),
				MerkleRoot: mr.MerkleRoot,
			})
		}

		priceUpdates := cciptypes.PriceUpdates{
			TokenPriceUpdates: make([]cciptypes.TokenPrice, 0),
			GasPriceUpdates:   make([]cciptypes.GasPriceChain, 0),
		}

		for _, tokenPriceUpdate := range ev.PriceUpdates.TokenPriceUpdates {
			priceUpdates.TokenPriceUpdates = append(priceUpdates.TokenPriceUpdates, cciptypes.TokenPrice{
				TokenID: cciptypes.UnknownEncodedAddress(
					typeconv.AddressBytesToString(tokenPriceUpdate.SourceToken, uint64(r.destChain))),
				Price: cciptypes.NewBigInt(tokenPriceUpdate.UsdPerToken),
			})
		}

		for _, gasPriceUpdate := range ev.PriceUpdates.GasPriceUpdates {
			priceUpdates.GasPriceUpdates = append(priceUpdates.GasPriceUpdates, cciptypes.GasPriceChain{
				ChainSel: cciptypes.ChainSelector(gasPriceUpdate.DestChainSelector),
				GasPrice: cciptypes.NewBigInt(gasPriceUpdate.UsdPerUnitGas),
			})
		}

		blockNum, err := strconv.ParseUint(item.Head.Height, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse block number %s: %w", item.Head.Height, err)
		}

		reports = append(reports, plugintypes2.CommitPluginReportWithMeta{
			Report: cciptypes.CommitPluginReport{
				MerkleRoots:  merkleRoots,
				PriceUpdates: priceUpdates,
			},
			Timestamp: time.Unix(int64(item.Timestamp), 0),
			BlockNum:  blockNum,
		})
	}

	lggr.Debugw("decoded commit reports", "reports", reports)

	return r.populateDisabledChainsInfo(ctx, reports, limit)
}

// populateDisabledChainsInfo will populate the DisabledSourceChains field of the reports.
// If all the chains of a report are disabled, the whole report will be skipped.
func (r *ccipChainReader) populateDisabledChainsInfo(
	ctx context.Context,
	reports []plugintypes2.CommitPluginReportWithMeta,
	limit int,
) ([]plugintypes2.CommitPluginReportWithMeta, error) {
	sourceChainConfigs, err := r.getAllOffRampSourceChainsConfig(ctx, r.lggr, r.destChain)
	if err != nil {
		return nil, fmt.Errorf("get all offRamp source chains config: %w", err)
	}

	disabledSourceChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain, cfg := range sourceChainConfigs {
		if !cfg.IsEnabled {
			disabledSourceChains.Add(chain)
		}
	}
	r.lggr.Debugw("disabled source chains", "chains", disabledSourceChains)

	reportsAfterRemovingDisabled := make([]plugintypes2.CommitPluginReportWithMeta, 0)
	for _, rep := range reports {
		chainsOfReport := mapset.NewSet[cciptypes.ChainSelector]()
		for _, mr := range rep.Report.MerkleRoots {
			chainsOfReport.Add(mr.ChainSel)
		}

		disabledChainsOfReportSet := chainsOfReport.Intersect(disabledSourceChains)
		disabledChainsOfReportSlice := disabledChainsOfReportSet.ToSlice()
		sort.Slice(disabledChainsOfReportSlice, func(i, j int) bool {
			return disabledChainsOfReportSlice[i] < disabledChainsOfReportSlice[j]
		})

		r.lggr.Debugw("disabled source chains of report",
			"report", rep, "chains", disabledChainsOfReportSlice)

		if chainsOfReport.Cardinality() == disabledChainsOfReportSet.Cardinality() {
			r.lggr.Warnw("all source chains of report are disabled, skipping report",
				"report", rep)
		} else {
			rep.DisabledSourceChains = disabledChainsOfReportSlice
			reportsAfterRemovingDisabled = append(reportsAfterRemovingDisabled, rep)
		}
	}

	if len(reportsAfterRemovingDisabled) < limit {
		return reportsAfterRemovingDisabled, nil
	}
	return reportsAfterRemovingDisabled[:limit], nil
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
	ctx context.Context, source cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.SeqNum, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, err
	}

	dataTyp := ExecutionStateChangedEvent{}

	iter, err := r.contractReaders[r.destChain].ExtendedQueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameExecutionStateChanged,
			Expressions: []query.Expression{
				query.Comparator(consts.EventAttributeSourceChain, primitives.ValueComparator{
					Value:    source,
					Operator: primitives.Eq,
				}),
				query.Comparator(consts.EventAttributeSequenceNumber, primitives.ValueComparator{
					Value:    seqNumRange.Start(),
					Operator: primitives.Gte,
				}, primitives.ValueComparator{
					Value:    seqNumRange.End(),
					Operator: primitives.Lte,
				}),
				query.Comparator(consts.EventAttributeState, primitives.ValueComparator{
					Value:    0,
					Operator: primitives.Gt,
				}),
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortBySequence(query.Asc)},
			Limit: query.Limit{
				Count: uint64(seqNumRange.End() - seqNumRange.Start() + 1),
			},
		},
		&dataTyp,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}

	executed := make([]cciptypes.SeqNum, 0)
	for _, item := range iter {
		stateChange, ok := item.Data.(*ExecutionStateChangedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %T to executionStateChangedEvent", item.Data)
		}

		if err := validateExecutionStateChangedEvent(stateChange, seqNumRange, source); err != nil {
			r.lggr.Errorw("validate execution state changed event",
				"err", err, "stateChange", stateChange)
			continue
		}

		executed = append(executed, stateChange.SequenceNumber)
	}

	return executed, nil
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
			lggr.Errorw("validate send requested event", "err", err, "msg", msg)
			continue
		}

		msg.Message.ExtraArgsDecoded, err = r.extraDataCodec.DecodeExtraArgs(msg.Message.ExtraArgs, sourceChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to decode the ExtraArgs: %w", err)
		}

		msg.Message.Header.OnRamp = onRampAddress

		for i, tokenAmount := range msg.Message.TokenAmounts {
			msg.Message.TokenAmounts[i].DestExecDataDecoded, err = r.extraDataCodec.DecodeTokenAmountDestExecData(
				tokenAmount.DestExecData,
				sourceChainSelector,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to decode token amount destExecData (%v): %w", tokenAmount.DestExecData, err)
			}
		}
		msgs = append(msgs, msg.Message)
	}

	lggr.Infow("decoded messages between sequence numbers", "msgs", msgs,
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String())

	return msgs, nil
}

// GetExpectedNextSequenceNumber implements CCIP.
func (r *ccipChainReader) GetExpectedNextSequenceNumber(
	ctx context.Context,
	sourceChainSelector cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
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

	return cciptypes.SeqNum(expectedNextSequenceNumber), nil
}

func (r *ccipChainReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	cfgs, err := r.getOffRampSourceChainsConfig(ctx, chains)
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

		if cfg.MinSeqNr == 0 {
			lggr.Errorf("minSeqNr not found for chain %d or is set to 0, chain is skipped.", chain)
			continue
		}

		res[chain] = cciptypes.SeqNum(cfg.MinSeqNr)
	}

	return res, err
}

func (r *ccipChainReader) Nonces(
	ctx context.Context,
	sourceChainSelector cciptypes.ChainSelector,
	addresses []string,
) (map[string]uint64, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, err
	}

	// Prepare the batch request
	contractBatch := make([]types.BatchRead, len(addresses))
	addressToIndex := make(map[string]int, len(addresses))
	responses := make([]uint64, len(addresses))

	for i, address := range addresses {
		sender, err := typeconv.AddressStringToBytes(address, uint64(sourceChainSelector))
		if err != nil {
			return nil, fmt.Errorf("failed to convert address %s to bytes: %w", address, err)
		}

		// TODO: evm only, need to make chain agnostic.
		// pad the sender slice to 32 bytes from the left
		sender = slicelib.LeftPadBytes(sender, 32)

		lggr.Infow("getting nonce for address",
			"address", address, "sender", hex.EncodeToString(sender))

		contractBatch[i] = types.BatchRead{
			ReadName: consts.MethodNameGetInboundNonce,
			Params: map[string]any{
				"sourceChainSelector": sourceChainSelector,
				"sender":              sender,
			},
			ReturnVal: &responses[i],
		}
		addressToIndex[address] = i
	}

	request := contractreader.ExtendedBatchGetLatestValuesRequest{
		consts.ContractNameNonceManager: contractBatch,
	}

	batchResult, err := r.contractReaders[r.destChain].ExtendedBatchGetLatestValues(
		ctx,
		request,
	)
	if err != nil {
		return nil, fmt.Errorf("batch get nonces failed: %w", err)
	}

	// Process results
	res := make(map[string]uint64, len(addresses))
	for _, results := range batchResult {
		for i, readResult := range results {
			address := getAddressByIndex(addressToIndex, i)

			returnVal, err := readResult.GetResult()
			if err != nil {
				r.lggr.Errorw("failed to get nonce for address", "address", address, "err", err)
				continue
			}

			val, ok := returnVal.(*uint64)
			if !ok || val == nil {
				r.lggr.Errorw("invalid nonce value returned", "address", address)
				continue
			}

			res[address] = *val
		}
	}

	return res, nil
}

func getAddressByIndex(addressToIndex map[string]int, index int) string {
	for address, idx := range addressToIndex {
		if idx == index {
			return address
		}
	}
	return ""
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

		if feeComponent.ExecutionFee == nil {
			lggr.Errorw("execution fee is nil", "chain", chain)
			continue
		}
		if feeComponent.DataAvailabilityFee == nil {
			lggr.Errorw("data availability fee is nil", "chain", chain)
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

		//TODO: Use batching in the future
		var nativeTokenAddress cciptypes.Bytes
		err := reader.ExtendedGetLatestValue(
			ctx,
			consts.ContractNameRouter,
			consts.MethodNameRouterGetWrappedNative,
			primitives.Unconfirmed,
			nil,
			&nativeTokenAddress)
		if err != nil {
			lggr.Warnw("failed to get native token address", "chain", chain, "err", err)
			continue
		}

		if nativeTokenAddress.String() == "0x" {
			lggr.Errorw("native token address is empty", "chain", chain)
			continue
		}

		var update *plugintypes.TimestampedUnixBig
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

		if update == nil || update.Timestamp == 0 {
			lggr.Warnw("no native token price available", "chain", chain)
			continue
		}
		if update.Value == nil {
			lggr.Errorw("native token price is nil", "chain", chain)
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
func (r *ccipChainReader) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) map[cciptypes.ChainSelector]plugintypes.TimestampedBig {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		lggr.Errorw("GetChainFeePriceUpdate dest chain extended reader not exist", "err", err)
		return nil
	}

	feeUpdates := make(map[cciptypes.ChainSelector]plugintypes.TimestampedBig, len(selectors))
	for _, chain := range selectors {
		update := plugintypes.TimestampedUnixBig{}
		// Read from dest chain
		err := r.contractReaders[r.destChain].ExtendedGetLatestValue(
			ctx,
			consts.ContractNameFeeQuoter,
			consts.MethodNameGetFeePriceUpdate,
			primitives.Unconfirmed,
			map[string]any{
				// That actually means that this selector is a source chain for the destChain
				"destChainSelector": chain,
			},
			&update,
		)
		if err != nil {
			lggr.Warnw("failed to get chain fee price update", "chain", chain, "err", err)
			continue
		}
		if update.Timestamp == 0 || update.Value == nil || update.Value.Cmp(big.NewInt(0)) == 0 {
			lggr.Debugw("chain fee price update is empty", "chain", chain)
			continue
		}
		feeUpdates[chain] = plugintypes.TimeStampedBigFromUnix(update)
	}

	return feeUpdates
}

func (r *ccipChainReader) GetRMNRemoteConfig(
	ctx context.Context,
) (rmntypes.RemoteConfig, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return rmntypes.RemoteConfig{}, err
	}

	// RMNRemote address stored in the offramp static config is actually the proxy contract address.
	// Here we will get the RMNRemote address from the proxy contract by calling the RMNProxy contract.
	proxyContractAddress, err := r.GetContractAddress(consts.ContractNameRMNRemote, r.destChain)
	if err != nil {
		return rmntypes.RemoteConfig{}, fmt.Errorf("get RMNRemote proxy contract address: %w", err)
	}

	rmnRemoteAddress, err := r.getRMNRemoteAddress(ctx, lggr, r.destChain, proxyContractAddress)
	if err != nil {
		return rmntypes.RemoteConfig{}, fmt.Errorf("get RMNRemote address: %w", err)
	}
	lggr.Debugw("got RMNRemote address", "address", rmnRemoteAddress)

	// TODO: make the calls in parallel using errgroup
	var vc versionedConfig
	err = r.contractReaders[r.destChain].ExtendedGetLatestValue(
		ctx,
		consts.ContractNameRMNRemote,
		consts.MethodNameGetVersionedConfig,
		primitives.Unconfirmed,
		map[string]any{},
		&vc,
	)
	if err != nil {
		return rmntypes.RemoteConfig{}, fmt.Errorf("get RMNRemote config: %w", err)
	}

	type ret struct {
		DigestHeader cciptypes.Bytes32
	}
	var header ret

	err = r.contractReaders[r.destChain].ExtendedGetLatestValue(
		ctx,
		consts.ContractNameRMNRemote,
		consts.MethodNameGetReportDigestHeader,
		primitives.Unconfirmed,
		map[string]any{},
		&header,
	)
	if err != nil {
		return rmntypes.RemoteConfig{}, fmt.Errorf("get RMNRemote report digest header: %w", err)
	}
	lggr.Infow("got RMNRemote report digest header", "digest", header.DigestHeader)

	signers := make([]rmntypes.RemoteSignerInfo, 0, len(vc.Config.Signers))
	for _, signer := range vc.Config.Signers {
		signers = append(signers, rmntypes.RemoteSignerInfo{
			OnchainPublicKey: signer.OnchainPublicKey,
			NodeIndex:        signer.NodeIndex,
		})
	}

	return rmntypes.RemoteConfig{
		ContractAddress:  rmnRemoteAddress,
		ConfigDigest:     vc.Config.RMNHomeContractConfigDigest,
		Signers:          signers,
		FSign:            vc.Config.F,
		ConfigVersion:    vc.Version,
		RmnReportVersion: header.DigestHeader,
	}, nil
}

// GetRmnCurseInfo returns rmn curse/pausing information about the provided chains
// from the destination chain RMN remote contract.
func (r *ccipChainReader) GetRmnCurseInfo(
	ctx context.Context,
	sourceChainSelectors []cciptypes.ChainSelector,
) (*CurseInfo, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, fmt.Errorf("validate dest=%d extended reader existence: %w", r.destChain, err)
	}

	type retTyp struct {
		CursedSubjects [][16]byte
	}
	var cursedSubjects retTyp

	err := r.contractReaders[r.destChain].ExtendedGetLatestValue(
		ctx,
		consts.ContractNameRMNRemote,
		consts.MethodNameGetCursedSubjects,
		primitives.Unconfirmed,
		map[string]any{},
		&cursedSubjects,
	)
	if err != nil {
		return nil, fmt.Errorf("get latest value of %s: %w", consts.MethodNameGetCursedSubjects, err)
	}

	lggr.Debugw("got cursed subjects", "cursedSubjects", cursedSubjects.CursedSubjects)

	return getCurseInfoFromCursedSubjects(
		lggr,
		mapset.NewSet(cursedSubjects.CursedSubjects...),
		r.destChain,
		sourceChainSelectors,
	), nil
}

func getCurseInfoFromCursedSubjects(
	lggr logger.Logger,
	cursedSubjectsSet mapset.Set[[16]byte],
	destChainSelector cciptypes.ChainSelector,
	sourceChainSelectors []cciptypes.ChainSelector,
) *CurseInfo {
	curseInfo := &CurseInfo{
		CursedSourceChains: make(map[cciptypes.ChainSelector]bool, len(sourceChainSelectors)),
		CursedDestination: cursedSubjectsSet.Contains(GlobalCurseSubject) ||
			cursedSubjectsSet.Contains(chainSelectorToBytes16(destChainSelector)),
		GlobalCurse: cursedSubjectsSet.Contains(GlobalCurseSubject),
	}

	for _, ch := range sourceChainSelectors {
		chainSelB16 := chainSelectorToBytes16(ch)
		lggr.Debugf("checking if chain %d is cursed after casting it to 16 bytes: %v", ch, chainSelB16)
		curseInfo.CursedSourceChains[ch] = cursedSubjectsSet.Contains(chainSelB16)
	}

	return curseInfo
}

func chainSelectorToBytes16(chainSel cciptypes.ChainSelector) [16]byte {
	var result [16]byte
	// Convert the uint64 to bytes and place it in the last 8 bytes of the array
	binary.BigEndian.PutUint64(result[8:], uint64(chainSel))
	return result
}

// discoverOffRampContracts uses the offRamp for a given chain to discover the addresses of other contracts.
func (r *ccipChainReader) discoverOffRampContracts(
	ctx context.Context,
	lggr logger.Logger,
	chain cciptypes.ChainSelector,
) (ContractAddresses, error) {
	// Exit without an error if we cannot read the destination.
	if err := validateExtendedReaderExistence(r.contractReaders, chain); err != nil {
		return nil, fmt.Errorf("validate extended reader existence: %w", err)
	}

	// build up resp as we go.
	resp := make(ContractAddresses)

	// OnRamps are in the offRamp SourceChainConfig.
	{
		sourceConfigs, err := r.getAllOffRampSourceChainsConfig(ctx, lggr, chain)
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
			if len(resp[consts.ContractNameRouter][chain]) == 0 {
				resp = resp.Append(consts.ContractNameRouter, chain, cfg.Router)
				lggr.Infow("appending router contract address", "address", cfg.Router)
			}
		}
	}

	// NonceManager and RMNRemote are in the offramp static config.
	{
		var staticConfig offRampStaticChainConfig
		err := r.getDestinationData(
			ctx,
			chain,
			consts.ContractNameOffRamp,
			consts.MethodNameOffRampGetStaticConfig,
			&staticConfig,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to lookup nonce manager and rmn proxy remote (offramp static config): %w", err)
		}
		resp = resp.Append(consts.ContractNameNonceManager, chain, staticConfig.NonceManager)
		resp = resp.Append(consts.ContractNameRMNRemote, chain, staticConfig.RmnRemote)
		lggr.Infow("appending RMN remote contract address", "address", staticConfig.RmnRemote)
	}

	// FeeQuoter from the offRamp dynamic config.
	{
		var dynamicConfig offRampDynamicChainConfig
		err := r.getDestinationData(
			ctx,
			chain,
			consts.ContractNameOffRamp,
			consts.MethodNameOffRampGetDynamicConfig,
			&dynamicConfig,
		)
		if err != nil {
			return nil, fmt.Errorf("unable to lookup fee quoter (offramp dynamic config): %w", err)
		}
		resp = resp.Append(consts.ContractNameFeeQuoter, chain, dynamicConfig.FeeQuoter)
		lggr.Infow("appending fee quoter contract address", "address", dynamicConfig.FeeQuoter)
	}

	return resp, nil
}

func (r *ccipChainReader) DiscoverContracts(ctx context.Context) (ContractAddresses, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	resp := make(ContractAddresses)

	// Discover destination contracts if the dest chain is supported.
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err == nil {
		resp, err = r.discoverOffRampContracts(ctx, lggr, r.destChain)
		if err != nil {
			return nil, fmt.Errorf("discover destination contracts: %w", err)
		}
	}

	// The following calls are on dynamically configured chains which may not
	// be available when this function is called. Eventually they will be
	// configured through consensus when the Sync function is called, but until
	// that happens the ErrNoBindings case must be handled gracefully.

	myChains := maps.Keys(r.contractReaders)

	// Read onRamps for FeeQuoter in DynamicConfig.
	dynamicConfigs := r.getOnRampDynamicConfigs(ctx, lggr, myChains)
	for chain, cfg := range dynamicConfigs {
		resp = resp.Append(consts.ContractNameFeeQuoter, chain, cfg.DynamicConfig.FeeQuoter)
	}

	// Read onRamps for Router in DestChainConfig.
	destChainConfig := r.getOnRampDestChainConfig(ctx, myChains)
	for chain, cfg := range destChainConfig {
		resp = resp.Append(consts.ContractNameRouter, chain, cfg.Router)
	}

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
			_, err := bindExtendedReaderContract(ctx, lggr, r.contractReaders, chainSel, contractName, address)
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

	addressBytes, err := typeconv.AddressStringToBytes(bindings[0].Binding.Address, uint64(chain))
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
	var staticConfig feeQuoterStaticConfig
	err := r.getDestinationData(
		ctx,
		r.destChain,
		consts.ContractNameFeeQuoter,
		consts.MethodNameFeeQuoterGetStaticConfig,
		&staticConfig,
	)

	if err != nil {
		return feeQuoterStaticConfig{}, fmt.Errorf("unable to lookup fee quoter (offramp static config): %w", err)
	}

	if len(staticConfig.LinkToken) == 0 {
		return feeQuoterStaticConfig{}, fmt.Errorf("link token address is empty")
	}

	return staticConfig, nil
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

	var timestampedPrice plugintypes.TimestampedUnixBig
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
type sourceChainConfig struct {
	Router    []byte // local router
	IsEnabled bool
	MinSeqNr  uint64
	OnRamp    cciptypes.UnknownAddress
}

func (scc sourceChainConfig) check() (bool /* enabled */, error) {
	// The chain may be set in CCIPHome's ChainConfig map but not hooked up yet in the offramp.
	if !scc.IsEnabled {
		return false, nil
	}
	// This may happen due to some sort of regression in the codec that unmarshals
	// chain data -> go struct.
	if len(scc.OnRamp) == 0 {
		return false, fmt.Errorf(
			"onRamp misconfigured/didn't unmarshal: %x",
			scc.OnRamp,
		)
	}
	return scc.IsEnabled, nil
}

// getOffRampSourceChainsConfig returns the offRamp contract's source chain configurations for each supported source
// chain. If some chain is disabled it is not included in the response.
func (r *ccipChainReader) getOffRampSourceChainsConfig(
	ctx context.Context, chains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]sourceChainConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, err
	}

	res := make(map[cciptypes.ChainSelector]sourceChainConfig)
	mu := new(sync.Mutex)

	eg := new(errgroup.Group)
	for _, chainSel := range chains {
		if chainSel == r.destChain {
			continue
		}

		// TODO: look into using BatchGetLatestValue instead to simplify concurrency?
		eg.Go(func() error {
			resp := sourceChainConfig{}
			err := r.contractReaders[r.destChain].ExtendedGetLatestValue(
				ctx,
				consts.ContractNameOffRamp,
				consts.MethodNameGetSourceChainConfig,
				primitives.Unconfirmed,
				map[string]any{
					"sourceChainSelector": chainSel,
				},
				&resp,
			)
			if err != nil {
				return fmt.Errorf("failed to get source chain config for source chain %d: %w",
					chainSel, err)
			}

			enabled, err := resp.check()
			if err != nil {
				return fmt.Errorf("source chain config check for chain %d failed: %w", chainSel, err)
			}
			if !enabled {
				// We don't want to process disabled chains prematurely.
				r.lggr.Debugw("source chain is disabled", "chain", chainSel)
				return nil
			}

			mu.Lock()
			res[chainSel] = resp
			mu.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

// selectorsAndConfigs wraps the return values from getAllSourceChainConfigs.
type selectorsAndConfigs struct {
	Selectors          []uint64            `mapstructure:"F0"`
	SourceChainConfigs []sourceChainConfig `mapstructure:"F1"`
}

// getAllOffRampSourceChainsConfig get all enabled source chain configs from the offRamp for the provided chain.
func (r *ccipChainReader) getAllOffRampSourceChainsConfig(
	ctx context.Context,
	lggr logger.Logger,
	chain cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]sourceChainConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, chain); err != nil {
		return nil, fmt.Errorf("validate extended reader existence: %w", err)
	}

	configs := make(map[cciptypes.ChainSelector]sourceChainConfig)

	var resp selectorsAndConfigs
	err := r.contractReaders[chain].ExtendedGetLatestValue(
		ctx,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetAllSourceChainConfigs,
		primitives.Unconfirmed,
		map[string]any{},
		&resp,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get source chain configs for source chain %d: %w",
			chain, err)
	}

	if len(resp.SourceChainConfigs) != len(resp.Selectors) {
		return nil, fmt.Errorf("selectors and source chain configs length mismatch: %v", resp)
	}

	lggr.Debugw("got source chain configs", "configs", resp)

	// Populate the map.
	for i := range resp.Selectors {
		chainSel := cciptypes.ChainSelector(resp.Selectors[i])
		cfg := resp.SourceChainConfigs[i]

		enabled, err := cfg.check()
		if err != nil {
			return nil, fmt.Errorf("source chain config check for chain %d failed: %w", chainSel, err)
		}
		if !enabled {
			// We don't want to process disabled chains prematurely.
			lggr.Debugw("source chain is disabled", "chain", chainSel)
			continue
		}

		configs[chainSel] = cfg
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

// getData returns data for a single reader.
func (r *ccipChainReader) getDestinationData(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	contract string,
	method string,
	returnVal any,
) error {
	if err := validateExtendedReaderExistence(r.contractReaders, destChain); err != nil {
		return err
	}

	if destChain != r.destChain {
		return fmt.Errorf("expected destination chain %d, got %d", r.destChain, destChain)
	}

	return r.contractReaders[destChain].ExtendedGetLatestValue(
		ctx,
		contract,
		method,
		primitives.Unconfirmed,
		map[string]any{},
		returnVal,
	)
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

func (r *ccipChainReader) getOnRampDynamicConfigs(
	ctx context.Context,
	lggr logger.Logger,
	srcChains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]getOnRampDynamicConfigResponse {
	result := make(map[cciptypes.ChainSelector]getOnRampDynamicConfigResponse)

	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for _, chainSel := range srcChains {
		// no onramp for the destination chain
		if chainSel == r.destChain {
			continue
		}
		if r.contractReaders[chainSel] == nil {
			r.lggr.Errorw("contract reader not found", "chain", chainSel)
			continue
		}

		wg.Add(1)
		go func(chainSel cciptypes.ChainSelector) {
			defer wg.Done()
			// read onramp dynamic config
			resp := getOnRampDynamicConfigResponse{}
			err := r.contractReaders[chainSel].ExtendedGetLatestValue(
				ctx,
				consts.ContractNameOnRamp,
				consts.MethodNameOnRampGetDynamicConfig,
				primitives.Unconfirmed,
				map[string]any{},
				&resp,
			)
			lggr.Debugw("got onramp dynamic config",
				"chain", chainSel,
				"resp", resp)
			if err != nil {
				if errors.Is(err, contractreader.ErrNoBindings) {
					// ErrNoBindings is an allowable error during initialization
					lggr.Infow(
						"unable to lookup source fee quoters (onRamp dynamic config), "+
							"this is expected during initialization", "err", err)
				} else {
					lggr.Errorw("unable to lookup source fee quoters (onRamp dynamic config)",
						"chain", chainSel, "err", err)
				}
				return
			}
			mu.Lock()
			result[chainSel] = resp
			mu.Unlock()
		}(chainSel)
	}

	wg.Wait()

	return result
}

// See DestChainConfig in OnRamp.sol
type onRampDestChainConfig struct {
	SequenceNumber   uint64 `json:"sequenceNumber"`
	AllowListEnabled bool   `json:"allowListEnabled"`
	Router           []byte `json:"router"`
}

func (r *ccipChainReader) getOnRampDestChainConfig(
	ctx context.Context,
	srcChains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]onRampDestChainConfig {
	result := make(map[cciptypes.ChainSelector]onRampDestChainConfig)

	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for _, chainSel := range srcChains {
		// no onramp for the destination chain
		if chainSel == r.destChain {
			continue
		}
		if r.contractReaders[chainSel] == nil {
			r.lggr.Errorw("contract reader not found", "chain", chainSel)
			continue
		}

		// For chain X, all DestChainConfigs will have one of 2 values for the Router address
		// 1. Chain X Test Router in case we're testing a new lane
		// 2. Chain X Router
		wg.Add(1)
		go func(chainSel cciptypes.ChainSelector) {
			defer wg.Done()
			resp := onRampDestChainConfig{}
			err := r.contractReaders[chainSel].ExtendedGetLatestValue(
				ctx,
				consts.ContractNameOnRamp,
				consts.MethodNameOnRampGetDestChainConfig,
				primitives.Unconfirmed,
				map[string]any{
					"destChainSelector": r.destChain,
				},
				&resp,
			)
			if err != nil {
				if errors.Is(err, contractreader.ErrNoBindings) {
					// ErrNoBindings is an allowable error during initialization
					r.lggr.Infow("unable to lookup source routers (onRamp dest chain config), "+
						"this is expected during initialization", "chain", chainSel, "err", err)
				} else {
					r.lggr.Errorw("unable to lookup source routers (onRamp dest chain config)",
						"chain", chainSel, "err", err)
				}
				return
			}

			if len(resp.Router) == 0 {
				r.lggr.Errorw("router address is empty", "chain", chainSel)
				return
			}

			mu.Lock()
			result[chainSel] = resp
			mu.Unlock()
		}(chainSel)
	}

	wg.Wait()
	return result
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
	F                           uint64            `json:"f"` // previously: MinSigners
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
	_, err := bindExtendedReaderContract(ctx, lggr, r.contractReaders, chain, consts.ContractNameRMNProxy, rmnRemoteProxyAddress)
	if err != nil {
		return nil, fmt.Errorf("bind RMN proxy contract: %w", err)
	}

	// get the RMN remote address from the proxy
	var rmnRemoteAddress []byte
	err = r.getDestinationData(
		ctx,
		chain,
		consts.ContractNameRMNProxy,
		consts.MethodNameGetARM,
		&rmnRemoteAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to lookup RMN remote address (RMN proxy): %w", err)
	}

	return rmnRemoteAddress, nil
}

// Get the DestChainConfig from the FeeQuoter contract on the given chain.
func (r *ccipChainReader) getFeeQuoterDestChainConfig(
	ctx context.Context,
	chainSelector cciptypes.ChainSelector,
) (cciptypes.FeeQuoterDestChainConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, chainSelector); err != nil {
		return cciptypes.FeeQuoterDestChainConfig{}, err
	}

	var destChainConfig cciptypes.FeeQuoterDestChainConfig
	srcReader := r.contractReaders[chainSelector]
	err := srcReader.ExtendedGetLatestValue(
		ctx,
		consts.ContractNameFeeQuoter,
		consts.MethodNameGetDestChainConfig,
		primitives.Unconfirmed,
		map[string]any{
			"destChainSelector": r.destChain,
		},
		&destChainConfig,
	)

	if err != nil {
		return cciptypes.FeeQuoterDestChainConfig{},
			fmt.Errorf("get dest chain config for source chain %d: %w",
				chainSelector, err)
	}

	return destChainConfig, nil
}

// GetMedianDataAvailabilityGasConfig returns the median of the DataAvailabilityGasConfig values from all FeeQuoters
// DA data lives in the FeeQuoter contract on the source chain. To get the config of the destination chain, we need to
// read the FeeQuoter contract on the source chain. As nodes are not required to have all chains configured, we need to
// read all FeeQuoter contracts to get the median.
func (r *ccipChainReader) GetMedianDataAvailabilityGasConfig(
	ctx context.Context,
) (cciptypes.DataAvailabilityGasConfig, error) {
	overheadGasValues := make([]uint32, 0)
	gasPerByteValues := make([]uint16, 0)
	multiplierBpsValues := make([]uint16, 0)

	// TODO: pay attention to performance here, as we are looping through all chains
	for chain := range r.contractReaders {
		config, err := r.getFeeQuoterDestChainConfig(ctx, chain)
		if err != nil {
			continue
		}
		if config.IsEnabled && config.HasNonEmptyDAGasParams() {
			overheadGasValues = append(overheadGasValues, config.DestDataAvailabilityOverheadGas)
			gasPerByteValues = append(gasPerByteValues, config.DestGasPerDataAvailabilityByte)
			multiplierBpsValues = append(multiplierBpsValues, config.DestDataAvailabilityMultiplierBps)
		}
	}

	// Calculate medians
	medianOverheadGas := consensus.Median(overheadGasValues, func(a, b uint32) bool { return a < b })
	medianGasPerByte := consensus.Median(gasPerByteValues, func(a, b uint16) bool { return a < b })
	medianMultiplierBps := consensus.Median(multiplierBpsValues, func(a, b uint16) bool { return a < b })

	daConfig := cciptypes.DataAvailabilityGasConfig{
		DestDataAvailabilityOverheadGas:   medianOverheadGas,
		DestGasPerDataAvailabilityByte:    medianGasPerByte,
		DestDataAvailabilityMultiplierBps: medianMultiplierBps,
	}

	return daConfig, nil
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
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return [32]byte{}, fmt.Errorf("validate dest=%d extended reader existence: %w", r.destChain, err)
	}

	type ConfigInfo struct {
		ConfigDigest                   [32]byte
		F                              uint8
		N                              uint8
		IsSignatureVerificationEnabled bool
	}

	type OCRConfig struct {
		ConfigInfo   ConfigInfo
		Signers      [][]byte
		Transmitters [][]byte
	}

	type OCRConfigResponse struct {
		OCRConfig OCRConfig
	}

	var resp OCRConfigResponse
	err := r.contractReaders[r.destChain].ExtendedGetLatestValue(
		ctx,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampLatestConfigDetails,
		primitives.Unconfirmed,
		map[string]any{
			"ocrPluginType": pluginType,
		},
		&resp,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("get latest config digest: %w", err)
	}

	return resp.OCRConfig.ConfigInfo.ConfigDigest, nil
}

func validateCommitReportAcceptedEvent(
	ev CommitReportAcceptedEvent, seq types.Sequence, gteTimestamp time.Time) error {
	if seq.Timestamp < uint64(gteTimestamp.Unix()) {
		return fmt.Errorf("commit report accepted event timestamp is less than the minimum timestamp %v<%v",
			seq.Timestamp, gteTimestamp.Unix())
	}

	if err := validateMerkleRoots(ev.MerkleRoots); err != nil {
		return fmt.Errorf("merkle roots: %w", err)
	}

	for _, tpus := range ev.PriceUpdates.TokenPriceUpdates {
		if len(tpus.SourceToken) == 0 {
			return fmt.Errorf("empty source token")
		}
		if tpus.UsdPerToken == nil {
			return fmt.Errorf("nil usd per token")
		}
	}

	for _, gpus := range ev.PriceUpdates.GasPriceUpdates {
		if gpus.DestChainSelector == 0 {
			return fmt.Errorf("dest chain is zero")
		}
		if gpus.UsdPerUnitGas == nil {
			return fmt.Errorf("nil usd per unit gas")
		}
	}

	return nil
}

func validateMerkleRoots(merkleRoots []MerkleRoot) error {
	for _, mr := range merkleRoots {
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
		if len(mr.OnRampAddress) == 0 {
			return fmt.Errorf("empty onramp address")
		}
	}

	return nil
}

func validateExecutionStateChangedEvent(
	ev *ExecutionStateChangedEvent, expRange cciptypes.SeqNumRange, sourceChain cciptypes.ChainSelector) error {
	if ev == nil {
		return fmt.Errorf("execution state changed event is nil")
	}

	if ev.SequenceNumber < expRange.Start() || ev.SequenceNumber > expRange.End() {
		return fmt.Errorf("execution state changed event sequence number is not in the expected range")
	}

	if ev.SourceChainSelector != sourceChain {
		return fmt.Errorf("source chain is not the expected queried one")
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

	if ev.DestChainSelector != dest {
		return fmt.Errorf("dest chain is not the expected queried one")
	}

	if ev.Message.Header.SourceChainSelector != source {
		return fmt.Errorf("source chain is not the expected queried one")
	}

	if ev.SequenceNumber < seqNumRange.Start() || ev.SequenceNumber > seqNumRange.End() {
		return fmt.Errorf("send requested event sequence number is not in the expected range")
	}

	if ev.Message.Header.MessageID.IsEmpty() {
		return fmt.Errorf("message ID is zero")
	}

	if len(ev.Message.Header.OnRamp) == 0 {
		return fmt.Errorf("empty onramp address")
	}

	if len(ev.Message.Receiver) == 0 {
		return fmt.Errorf("empty receiver address")
	}

	if len(ev.Message.Sender) == 0 {
		return fmt.Errorf("empty sender address")
	}

	if len(ev.Message.FeeToken) == 0 {
		return fmt.Errorf("empty fee token")
	}

	if ev.Message.FeeTokenAmount.IsEmpty() {
		return fmt.Errorf("fee token amount is zero")
	}

	if len(ev.Message.FeeToken) == 0 {
		return fmt.Errorf("empty fee token")
	}

	return nil
}

// Interface compliance check
var _ CCIPReader = (*ccipChainReader)(nil)
