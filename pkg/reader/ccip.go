package reader

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"

	ocr3types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// TODO: unit test the implementation when the actual contract reader and writer interfaces are finalized and mocks
// can be generated.
type ccipChainReader struct {
	lggr            logger.Logger
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter
	destChain       cciptypes.ChainSelector
	offrampAddress  string
}

func newCCIPChainReaderInternal(
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
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
	}

	contracts := ContractAddresses{
		consts.ContractNameOffRamp: {
			destChain: offrampAddress,
		},
	}
	if err := reader.Sync(context.Background(), contracts); err != nil {
		lggr.Infow("failed to sync contracts", "err", err)
	}

	return reader
}

// WithExtendedContractReader sets the extended contract reader for the provided chain.
func (r *ccipChainReader) WithExtendedContractReader(
	ch cciptypes.ChainSelector, cr contractreader.Extended) *ccipChainReader {
	r.contractReaders[ch] = cr
	return r
}

func (r *ccipChainReader) CommitReportsGTETimestamp(
	ctx context.Context, dest cciptypes.ChainSelector, ts time.Time, limit int,
) ([]plugintypes2.CommitPluginReportWithMeta, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, dest); err != nil {
		return nil, err
	}

	// ---------------------------------------------------
	// The following types are used to decode the events
	// but should be replaced by chain-reader modifiers and use the base cciptypes.CommitReport type.

	type MerkleRoot struct {
		SourceChainSelector uint64
		MinSeqNr            uint64
		MaxSeqNr            uint64
		MerkleRoot          cciptypes.Bytes32
		OnRampAddress       []byte
	}

	type TokenPriceUpdate struct {
		SourceToken []byte
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

	type CommitReportAccepted struct {
		PriceUpdates PriceUpdates
		MerkleRoots  []MerkleRoot
	}

	type CommitReportAcceptedEvent struct {
		Report CommitReportAccepted
	}
	// ---------------------------------------------------

	ev := CommitReportAcceptedEvent{}

	iter, err := r.contractReaders[dest].ExtendedQueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameCommitReportAccepted,
			Expressions: []query.Expression{
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortByTimestamp(query.Asc)},
		},
		&ev,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}
	r.lggr.Debugw("queried commit reports", "numReports", len(iter),
		"destChain", dest,
		"ts", ts,
		"limit", limit)

	reports := make([]plugintypes2.CommitPluginReportWithMeta, 0)
	for _, item := range iter {
		ev, is := (item.Data).(*CommitReportAcceptedEvent)
		if !is {
			return nil, fmt.Errorf("unexpected type %T while expecting a commit report", item)
		}

		valid := item.Timestamp >= uint64(ts.Unix())
		if !valid {
			r.lggr.Debugw("commit report too old, skipping", "report", ev.Report, "item", item,
				"destChain", dest,
				"ts", ts,
				"limit", limit)
			continue
		}

		merkleRoots := make([]cciptypes.MerkleRootChain, 0, len(ev.Report.MerkleRoots))
		for _, mr := range ev.Report.MerkleRoots {
			merkleRoots = append(merkleRoots, cciptypes.MerkleRootChain{
				ChainSel: cciptypes.ChainSelector(mr.SourceChainSelector),
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

		for _, tokenPriceUpdate := range ev.Report.PriceUpdates.TokenPriceUpdates {
			priceUpdates.TokenPriceUpdates = append(priceUpdates.TokenPriceUpdates, cciptypes.TokenPrice{
				TokenID: ocr3types.Account(typeconv.AddressBytesToString(tokenPriceUpdate.SourceToken, uint64(r.destChain))),
				Price:   cciptypes.NewBigInt(tokenPriceUpdate.UsdPerToken),
			})
		}

		for _, gasPriceUpdate := range ev.Report.PriceUpdates.GasPriceUpdates {
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

	if len(reports) < limit {
		return reports, nil
	}
	return reports[:limit], nil
}

func (r *ccipChainReader) ExecutedMessageRanges(
	ctx context.Context, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.SeqNumRange, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, dest); err != nil {
		return nil, err
	}

	type ExecutionStateChangedEvent struct {
		SourceChainSelector cciptypes.ChainSelector
		SequenceNumber      cciptypes.SeqNum
		State               uint8
	}

	dataTyp := ExecutionStateChangedEvent{}

	iter, err := r.contractReaders[dest].ExtendedQueryKey(
		ctx,
		consts.ContractNameOffRamp,
		query.KeyFilter{
			Key: consts.EventNameExecutionStateChanged,
			Expressions: []query.Expression{
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{query.NewSortByTimestamp(query.Asc)},
		},
		&dataTyp,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query offRamp: %w", err)
	}

	executed := make([]cciptypes.SeqNumRange, 0)
	for _, item := range iter {
		stateChange, ok := item.Data.(*ExecutionStateChangedEvent)
		if !ok {
			return nil, fmt.Errorf("failed to cast %T to executionStateChangedEvent", item.Data)
		}

		// todo: filter via the query
		valid := stateChange.SourceChainSelector == source &&
			stateChange.SequenceNumber >= seqNumRange.Start() &&
			stateChange.SequenceNumber <= seqNumRange.End() &&
			stateChange.State > 0
		if !valid {
			r.lggr.Debugw("skipping invalid state change", "stateChange", stateChange)
			continue
		}

		executed = append(executed, cciptypes.NewSeqNumRange(stateChange.SequenceNumber, stateChange.SequenceNumber))
	}

	return executed, nil
}

func (r *ccipChainReader) MsgsBetweenSeqNums(
	ctx context.Context, sourceChainSelector cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, sourceChainSelector); err != nil {
		return nil, err
	}

	onRampAddress, err := r.GetContractAddress(consts.ContractNameOnRamp, sourceChainSelector)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address: %w", err)
	}

	type SendRequestedEvent struct {
		DestChainSelector cciptypes.ChainSelector
		Message           cciptypes.Message
	}

	seq, err := r.contractReaders[sourceChainSelector].ExtendedQueryKey(
		ctx,
		consts.ContractNameOnRamp,
		query.KeyFilter{
			Key: consts.EventNameCCIPMessageSent,
			Expressions: []query.Expression{
				query.Confidence(primitives.Finalized),
			},
		},
		query.LimitAndSort{
			SortBy: []query.SortBy{
				query.NewSortByTimestamp(query.Asc),
			},
		},
		&SendRequestedEvent{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query onRamp: %w", err)
	}

	r.lggr.Infow("queried messages between sequence numbers",
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

		// todo: filter via the query
		valid := msg.Message.Header.SourceChainSelector == sourceChainSelector &&
			msg.Message.Header.DestChainSelector == r.destChain &&
			msg.Message.Header.SequenceNumber >= seqNumRange.Start() &&
			msg.Message.Header.SequenceNumber <= seqNumRange.End()

		msg.Message.Header.OnRamp = onRampAddress

		if valid {
			msgs = append(msgs, msg.Message)
		}
	}

	r.lggr.Infow("decoded messages between sequence numbers", "msgs", msgs,
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String())

	return msgs, nil
}

// GetExpectedNextSequenceNumber implements CCIP.
func (r *ccipChainReader) GetExpectedNextSequenceNumber(
	ctx context.Context,
	sourceChainSelector, destChainSelector cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	if destChainSelector != r.destChain {
		return 0, fmt.Errorf("expected destination chain %d, got %d", r.destChain, destChainSelector)
	}

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
			"destChainSelector": destChainSelector,
		},
		&expectedNextSequenceNumber,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get expected next sequence number from onramp: %w", err)
	}

	return cciptypes.SeqNum(expectedNextSequenceNumber), nil
}

func (r *ccipChainReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) ([]cciptypes.SeqNum, error) {
	cfgs, err := r.getSourceChainsConfig(ctx, chains)
	if err != nil {
		return nil, fmt.Errorf("get source chains config: %w", err)
	}

	res := make([]cciptypes.SeqNum, 0, len(chains))
	for _, chain := range chains {
		cfg, exists := cfgs[chain]
		if !exists {
			return nil, fmt.Errorf("source chain config not found for chain %d", chain)
		}
		if cfg.MinSeqNr == 0 {
			return nil, fmt.Errorf("minSeqNr not found for chain %d", chain)
		}
		res = append(res, cciptypes.SeqNum(cfg.MinSeqNr))
	}

	return res, err
}

func (r *ccipChainReader) Nonces(
	ctx context.Context,
	sourceChainSelector, destChainSelector cciptypes.ChainSelector,
	addresses []string,
) (map[string]uint64, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, destChainSelector); err != nil {
		return nil, err
	}

	res := make(map[string]uint64)
	mu := new(sync.Mutex)
	eg := new(errgroup.Group)

	for _, address := range addresses {
		address := address
		eg.Go(func() error {
			sender, err := typeconv.AddressStringToBytes(address, uint64(destChainSelector))
			if err != nil {
				return fmt.Errorf("failed to convert address %s to bytes: %w", address, err)
			}

			var resp uint64
			err = r.contractReaders[destChainSelector].ExtendedGetLatestValue(
				ctx,
				consts.ContractNameNonceManager,
				consts.MethodNameGetInboundNonce,
				primitives.Unconfirmed,
				map[string]any{
					"sourceChainSelector": sourceChainSelector,
					"sender":              sender,
				},
				&resp,
			)
			if err != nil {
				return fmt.Errorf("failed to get nonce for address %s: %w", address, err)
			}
			mu.Lock()
			defer mu.Unlock()
			res[address] = resp
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ccipChainReader) GetAvailableChainsFeeComponents(
	ctx context.Context,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	feeComponents := make(map[cciptypes.ChainSelector]types.ChainFeeComponents, len(r.contractWriters))
	for chain, chainWriter := range r.contractWriters {
		feeComponent, err := chainWriter.GetFeeComponents(ctx)
		if err != nil {
			r.lggr.Errorw("failed to get chain fee components for chain %d: %w", chain, err)
			continue
		}
		feeComponents[chain] = *feeComponent
	}
	return feeComponents
}

func (r *ccipChainReader) GetWrappedNativeTokenPriceUSD(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	// 1. Call chain's router to get native token address https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/Router.sol#L189:L191
	// nolint:lll
	// 2. Call chain's FeeQuoter to get native tokens price  https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L229-L229
	prices := make(map[cciptypes.ChainSelector]cciptypes.BigInt)
	for _, chain := range selectors {
		reader, ok := r.contractReaders[chain]
		if !ok {
			r.lggr.Warnw("contract reader not found", "chain", chain)
			continue
		}

		//TODO: Use batching in the future
		var nativeTokenAddress ocr3types.Account
		err := reader.ExtendedGetLatestValue(
			ctx,
			consts.ContractNameRouter,
			consts.MethodNameRouterGetWrappedNative,
			primitives.Unconfirmed,
			nil,
			&nativeTokenAddress)

		if err != nil {
			r.lggr.Errorw("failed to get native token address", "chain", chain, "err", err)
			continue
		}

		if nativeTokenAddress == "" {
			r.lggr.Errorw("native token address is empty", "chain", chain)
			continue
		}

		var price *big.Int
		err = reader.ExtendedGetLatestValue(
			ctx,
			consts.ContractNameFeeQuoter,
			consts.MethodNameFeeQuoterGetTokenPrices,
			primitives.Unconfirmed,
			map[string]any{
				"token": nativeTokenAddress,
			},
			&price,
		)
		if err != nil {
			r.lggr.Errorw("failed to get native token price", "chain", chain, "err", err)
			continue
		}

		if price == nil {
			r.lggr.Errorw("native token price is nil", "chain", chain)
			continue
		}
		prices[chain] = cciptypes.NewBigInt(price)
	}
	return prices
}

// GetChainFeePriceUpdate Read from Destination chain FeeQuoter latest fee updates for the provided chains.
// It unpacks the packed fee into the ChainFeeUSDPrices struct.
// nolint:lll
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L263-L263
func (r *ccipChainReader) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) map[cciptypes.ChainSelector]plugintypes.TimestampedBig {
	feeUpdates := make(map[cciptypes.ChainSelector]plugintypes.TimestampedBig, len(selectors))
	for _, chain := range selectors {
		update := plugintypes.TimestampedBig{}
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
			r.lggr.Errorw("failed to get chain fee price update", "chain", chain, "err", err)
			continue
		}
		feeUpdates[chain] = update
		//feeUpdates[chain] = chainfee.ChainFeeUpdate{
		//	Timestamp: update.Timestamp,
		//	ChainFee:  chainfee.FromPackedFee(update.Value.Int),
		//}
	}

	return feeUpdates
}

func (r *ccipChainReader) DiscoverContracts(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
) (ContractAddresses, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, destChain); err != nil {
		return nil, err
	}

	chains := maps.Keys(r.contractReaders)

	// OnRamps are in the offramp SourceChainConfig.
	sourceConfigs, err := r.getSourceChainsConfig(ctx, chains)
	if err != nil {
		return nil, fmt.Errorf("unable to get SourceChainsConfig: %w", err)
	}

	// NonceManager is in the offramp static config.
	staticConfig, err := r.getOfframpStaticConfig(ctx, r.destChain)
	if err != nil {
		return nil, fmt.Errorf("unable to lookup nonce manager: %w", err)
	}

	// TODO: Lookup FeeQuoter (from onRamp DynamicConfig)
	// TODO: Lookup Router (from onRamp DestChainConfig)

	// Build response object.
	onramps := make(map[cciptypes.ChainSelector][]byte, len(chains))
	for chain, cfg := range sourceConfigs {
		onramps[chain] = cfg.OnRamp
	}
	resp := map[string]map[cciptypes.ChainSelector][]byte{
		consts.ContractNameOnRamp: onramps,
		consts.ContractNameNonceManager: {
			destChain: staticConfig.NonceManager,
		},
	}
	return resp, nil
}

// Sync goes through the input contracts and binds them to the contract reader.
func (r *ccipChainReader) Sync(ctx context.Context, contracts ContractAddresses) error {
	var errs []error
	for contractName, chainSelToAddress := range contracts {
		for chainSel, address := range chainSelToAddress {
			// defense in depth: don't bind if the address is empty.
			// callers should ensure this but we double check here.
			if len(address) == 0 {
				r.lggr.Warnw("skipping binding empty address for contract",
					"contractName", contractName,
					"chainSel", chainSel,
				)
				continue
			}

			// try to bind
			_, err := bindExtendedReaderContract(ctx, r.contractReaders, chainSel, contractName, address)
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

func (r *ccipChainReader) Close(ctx context.Context) error {
	return nil
}

func (r *ccipChainReader) GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error) {
	bindings := r.contractReaders[chain].GetBindings(contractName)
	if len(bindings) != 1 {
		return nil, fmt.Errorf("expected one binding for the %s contract, got %d", contractName, len(bindings))
	}

	addressBytes, err := typeconv.AddressStringToBytes(bindings[0].Binding.Address, uint64(chain))
	if err != nil {
		return nil, fmt.Errorf("convert address %s to bytes: %w", bindings[0].Binding.Address, err)
	}

	return addressBytes, nil
}

// getSourceChainsConfig returns the offRamp contract's source chain configurations for each supported source chain.
func (r *ccipChainReader) getSourceChainsConfig(
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
		chainSel := chainSel
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
				return fmt.Errorf("failed to get source chain config: %w", err)
			}

			// This may happen due to some sort of regression in the codec that unmarshals
			// chain data -> go struct.
			if len(resp.OnRamp) == 0 {
				return fmt.Errorf(
					"onRamp misconfigured/didn't unmarshal for chain %d: %x",
					chainSel,
					resp.OnRamp,
				)
			}

			if !resp.IsEnabled {
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

// sourceChainConfig is used to parse the response from the offRamp contract's getSourceChainConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L94
//
//nolint:lll // It's a URL.
type sourceChainConfig struct {
	IsEnabled bool
	OnRamp    []byte
	MinSeqNr  uint64
}

// getOfframpStaticConfig returns the destination offRamp contract's static configuration.
func (r *ccipChainReader) getOfframpStaticConfig(
	ctx context.Context,
	chain cciptypes.ChainSelector,
) (offrampStaticChainConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, chain); err != nil {
		return offrampStaticChainConfig{}, err
	}

	resp := offrampStaticChainConfig{}
	err := r.contractReaders[chain].ExtendedGetLatestValue(
		ctx,
		consts.ContractNameOffRamp,
		consts.MethodNameOfframpGetStaticConfig,
		primitives.Unconfirmed,
		map[string]any{},
		&resp,
	)
	if err != nil {
		return offrampStaticChainConfig{}, fmt.Errorf("failed to get source chain config: %w", err)
	}
	return resp, nil
}

// offrampStaticChainConfig is used to parse the response from the offRamp contract's getStaticConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L86
//
//nolint:lll // It's a URL.
type offrampStaticChainConfig struct {
	ChainSelector      cciptypes.ChainSelector `json:"chainSelector"`
	RmnProxy           []byte                  `json:"rmnProxy"`
	TokenAdminRegistry []byte                  `json:"tokenAdminRegistry"`
	NonceManager       []byte                  `json:"nonceManager"`
}

// Interface compliance check
var _ CCIPReader = (*ccipChainReader)(nil)
