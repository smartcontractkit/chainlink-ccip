package reader

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"time"

	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"

	ocr3types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
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
	ctx context.Context,
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

	type CommitReportAcceptedEvent struct {
		PriceUpdates PriceUpdates
		MerkleRoots  []MerkleRoot
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
			r.lggr.Debugw("commit report too old, skipping", "report", ev, "item", item,
				"destChain", dest,
				"ts", ts,
				"limit", limit)
			continue
		}

		r.lggr.Debugw("processing commit report", "report", ev, "item", item)

		merkleRoots := make([]cciptypes.MerkleRootChain, 0, len(ev.MerkleRoots))
		for _, mr := range ev.MerkleRoots {
			onRampAddress, err := r.GetContractAddress(
				consts.ContractNameOnRamp,
				cciptypes.ChainSelector(mr.SourceChainSelector),
			)
			if err != nil {
				return nil, fmt.Errorf("get onRamp address for selector %d: %w", mr.SourceChainSelector, err)
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
				TokenID: ocr3types.Account(typeconv.AddressBytesToString(tokenPriceUpdate.SourceToken, uint64(r.destChain))),
				Price:   cciptypes.NewBigInt(tokenPriceUpdate.UsdPerToken),
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

	r.lggr.Debugw("decoded commit reports", "reports", reports)

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
	cfgs, err := r.getOffRampSourceChainsConfig(ctx, chains)
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
	// 2. Call chain's FeeQuoter to get native tokens price  https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L229-L229
	//
	//nolint:lll
	prices := make(map[cciptypes.ChainSelector]cciptypes.BigInt)
	for _, chain := range selectors {
		reader, ok := r.contractReaders[chain]
		if !ok {
			r.lggr.Warnw("contract reader not found", "chain", chain)
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
			r.lggr.Warnw("failed to get native token address", "chain", chain, "err", err)
			continue
		}

		if nativeTokenAddress.String() == "0x" {
			r.lggr.Errorw("native token address is empty", "chain", chain)
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
			r.lggr.Errorw("failed to get native token price", "chain", chain, "err", err)
			continue
		}

		if update == nil {
			r.lggr.Errorw("native token price is nil", "chain", chain)
			continue
		}
		prices[chain] = update.Value
	}
	return prices
}

// GetChainFeePriceUpdate Read from Destination chain FeeQuoter latest fee updates for the provided chains.
// It unpacks the packed fee into the ChainFeeUSDPrices struct.
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L263-L263
//
//nolint:lll
func (r *ccipChainReader) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) map[cciptypes.ChainSelector]plugintypes.TimestampedBig {
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
			r.lggr.Warnw("failed to get chain fee price update", "chain", chain, "err", err)
			continue
		}
		if update.Timestamp == 0 || update.Value.IsEmpty() {
			continue
		}
		feeUpdates[chain] = plugintypes.TimeStampedBigFromUnix(update)
	}

	return feeUpdates
}

func (r *ccipChainReader) GetRMNRemoteConfig(
	ctx context.Context,
	destChainSelector cciptypes.ChainSelector,
) (rmntypes.RemoteConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, destChainSelector); err != nil {
		return rmntypes.RemoteConfig{}, err
	}

	contractAddress, err := r.GetContractAddress(consts.ContractNameRMNRemote, destChainSelector)
	if err != nil {
		return rmntypes.RemoteConfig{}, fmt.Errorf("get RMNRemote contract address: %w", err)
	}

	// TODO: make the calls in parallel using errgroup
	var vc versionedConfig
	err = r.contractReaders[destChainSelector].ExtendedGetLatestValue(
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

	err = r.contractReaders[destChainSelector].ExtendedGetLatestValue(
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
	r.lggr.Infow("got RMNRemote report digest header", "digest", header.DigestHeader)

	signers := make([]rmntypes.RemoteSignerInfo, 0, len(vc.Config.Signers))
	for _, signer := range vc.Config.Signers {
		signers = append(signers, rmntypes.RemoteSignerInfo{
			OnchainPublicKey: signer.OnchainPublicKey,
			NodeIndex:        signer.NodeIndex,
		})
	}

	return rmntypes.RemoteConfig{
		ContractAddress:  contractAddress,
		ConfigDigest:     cciptypes.Bytes32(vc.Config.RMNHomeContractConfigDigest),
		Signers:          signers,
		MinSigners:       vc.Config.MinSigners,
		ConfigVersion:    vc.Version,
		RmnReportVersion: header.DigestHeader,
	}, nil
}

func (r *ccipChainReader) discoverDestinationContracts(
	ctx context.Context,
	allChains []cciptypes.ChainSelector,
) (ContractAddresses, error) {
	// Exit without an error if we cannot read the destination.
	if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
		return nil, fmt.Errorf("validate extended reader existence: %w", err)
	}

	// build up resp as we go.
	var resp ContractAddresses

	// OnRamps are in the offRamp SourceChainConfig.
	sourceConfigs, err := r.getOffRampSourceChainsConfig(ctx, allChains)
	if err != nil {
		return nil, fmt.Errorf("unable to get SourceChainsConfig: %w", err)
	}
	{
		// Iterate in chain selector order so that the router config is deterministic.
		keys := maps.Keys(sourceConfigs)
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, chain := range keys {
			cfg := sourceConfigs[chain]
			resp = resp.Append(consts.ContractNameOnRamp, chain, cfg.OnRamp)
			// The local router is located in each source chain config. Add it once.
			if len(resp[consts.ContractNameRouter][r.destChain]) == 0 {
				resp = resp.Append(consts.ContractNameRouter, r.destChain, cfg.Router)
			}
		}
	}

	// NonceManager and RMNRemote are in the offramp static config.
	var staticConfig offRampStaticChainConfig
	err = r.getDestinationData(
		ctx,
		r.destChain,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetStaticConfig,
		&staticConfig,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to lookup nonce manager (offramp static config): %w", err)
	}
	resp = resp.Append(consts.ContractNameNonceManager, r.destChain, staticConfig.NonceManager)
	resp = resp.Append(consts.ContractNameRMNRemote, r.destChain, staticConfig.RmnRemote)
	r.lggr.Infow("appending RMN remote contract address", "address", staticConfig.RmnRemote)

	// FeeQuoter from the offRamp dynamic config.
	var dynamicConfig offRampDynamicChainConfig
	err = r.getDestinationData(
		ctx,
		r.destChain,
		consts.ContractNameOffRamp,
		consts.MethodNameOffRampGetDynamicConfig,
		&dynamicConfig,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to lookup fee quoter (offramp dynamic config): %w", err)
	}
	resp = resp.Append(consts.ContractNameFeeQuoter, r.destChain, dynamicConfig.FeeQuoter)

	return resp, nil
}

func (r *ccipChainReader) DiscoverContracts(
	ctx context.Context,
	allChains []cciptypes.ChainSelector,
) (ContractAddresses, error) {
	// TODO: Remove nil handling once the discovery processor is able to discover source chain addresses.
	// At that point it will pass in the destChain on the first round, and in subsequent rounds it will pass in
	// the full list of observed sources.
	if allChains == nil {
		allChains = maps.Keys(r.contractReaders)
	}
	resp, err := r.discoverDestinationContracts(ctx, allChains)
	// Ignore the error if the destination chain is not available. We still want to continue
	// discovering contracts from any source chains that may be available.
	if err != nil {
		if !errors.Is(err, ErrContractReaderNotFound) {
			return nil, fmt.Errorf("discover destination contracts: %w", err)
		}
		// Make sure
		resp = nil
	}

	// The following calls are on dynamically configured chains which may not
	// be available when this function is called. Eventually they will be
	// configured through consensus when the Sync function is called, but until
	// that happens the ErrNoBindings case must be handled gracefully.

	myChains := maps.Keys(r.contractReaders)

	// Read onRamps for FeeQuoter in DynamicConfig.
	{
		dynamicConfigs, err := r.getOnRampDynamicConfigs(ctx, myChains)
		if errors.Is(err, contractreader.ErrNoBindings) {
			// ErrNoBindings is an allowable error.
			r.lggr.Infow("unable to lookup source fee quoters, this is expected during initialization", "err", err)
		} else if err != nil {
			return nil, fmt.Errorf("unable to lookup source fee quoters (onRamp dynamic config): %w", err)
		} else {
			for chain, cfg := range dynamicConfigs {
				resp = resp.Append(consts.ContractNameFeeQuoter, chain, cfg.FeeQuoter)
			}
		}
	}

	// Read onRamps for Router in DestChainConfig.
	{
		destChainConfig, err := r.getOnRampDestChainConfig(ctx, myChains)
		if errors.Is(err, contractreader.ErrNoBindings) {
			// ErrNoBindings is an allowable error.
			r.lggr.Infow("unable to lookup source routers, this is expected during initialization", "err", err)
		} else if err != nil {
			return nil, fmt.Errorf("unable to lookup source routers (onRamp dest chain config): %w", err)
		} else {
			for chain, cfg := range destChainConfig {
				resp = resp.Append(consts.ContractNameRouter, chain, cfg.Router)
			}
		}
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
		return cciptypes.BigInt{}, fmt.Errorf("failed to get LINK token price, addr: %v, err: %w", tokenAddr, err)
	}

	price := timestampedPrice.Value.Int

	if price.Cmp(big.NewInt(0)) == 0 {
		return cciptypes.BigInt{}, fmt.Errorf("LINK token price is 0, addr: %v", tokenAddr)
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
	OnRamp    []byte
	MinSeqNr  uint64
}

// getOffRampSourceChainsConfig returns the offRamp contract's source chain configurations for each supported source
// chain.
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

			if len(resp.OnRamp) == 0 {
				// This indicates that the source chain config is NOT set for this
				// chain selector.
				// This can happen if a node is set up to read from a chain that has
				// not yet been configured in the offramp.
				r.lggr.Debugw("source chain config not found", "chain", chainSel)
				return nil
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

// offRampStaticChainConfig is used to parse the response from the offRamp contract's getStaticConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/a3f61f7458e4499c2c62eb38581c60b4942b1160/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L86
//
//nolint:lll // It's a URL.
type offRampStaticChainConfig struct {
	ChainSelector      cciptypes.ChainSelector `json:"chainSelector"`
	RmnRemote          []byte                  `json:"rmnRemote"`
	TokenAdminRegistry []byte                  `json:"tokenAdminRegistry"`
	NonceManager       []byte                  `json:"nonceManager"`
}

// offRampDynamicChainConfig maps to DynamicChainConfig in OffRamp.sol
type offRampDynamicChainConfig struct {
	FeeQuoter                               []byte `json:"feeQuoter"`
	PermissionLessExecutionThresholdSeconds uint32 `json:"permissionLessExecutionThresholdSeconds"`
	MaxTokenTransferGas                     uint32 `json:"maxTokenTransferGas"`
	MaxPoolReleaseOrMintGas                 uint32 `json:"maxPoolReleaseOrMintGas"`
	MessageValidator                        []byte `json:"messageValidator"`
}

//nolint:unused // it will be used soon // TODO: Remove nolint
type offRampDestChainConfig struct {
	SequenceNumber   uint64 `json:"sequenceNumber"`
	AllowListEnabled bool   `json:"allowListEnabled"`
	Router           []byte `json:"router"`
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
type onRampDynamicChainConfig struct {
	FeeQuoter        []byte `json:"feeQuoter"`
	MessageValidator []byte `json:"messageValidator"`
	FeeAggregator    []byte `json:"feeAggregator"`
	AllowListAdmin   []byte `json:"allowListAdmin"`
}

//nolint:dupl // It's not quite duplicate code...
func (r *ccipChainReader) getOnRampDynamicConfigs(
	ctx context.Context,
	srcChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]onRampDynamicChainConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, srcChains...); err != nil {
		return nil, err
	}

	result := make(map[cciptypes.ChainSelector]onRampDynamicChainConfig)

	mu := new(sync.Mutex)
	eg := new(errgroup.Group)
	for _, chainSel := range srcChains {
		// no onramp for the destination chain
		if chainSel == r.destChain {
			continue
		}

		chainSel := chainSel
		eg.Go(func() error {
			// read onramp dynamic config
			resp := onRampDynamicChainConfig{}
			err := r.contractReaders[chainSel].ExtendedGetLatestValue(
				ctx,
				consts.ContractNameOnRamp,
				consts.MethodNameOnRampGetDynamicConfig,
				primitives.Unconfirmed,
				map[string]any{},
				&resp,
			)
			if err != nil {
				return fmt.Errorf("failed to get onramp dynamic config: %w", err)
			}
			mu.Lock()
			result[chainSel] = resp
			mu.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return result, nil
}

// See DestChainConfig in OnRamp.sol
type onRampDestChainConfig struct {
	SequenceNumber   uint64 `json:"sequenceNumber"`
	AllowListEnabled bool   `json:"allowListEnabled"`
	Router           []byte `json:"router"`
}

//nolint:dupl // It's not quite duplicate code...
func (r *ccipChainReader) getOnRampDestChainConfig(
	ctx context.Context,
	srcChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]onRampDestChainConfig, error) {
	if err := validateExtendedReaderExistence(r.contractReaders, srcChains...); err != nil {
		return nil, err
	}

	result := make(map[cciptypes.ChainSelector]onRampDestChainConfig)

	mu := new(sync.Mutex)
	eg := new(errgroup.Group)
	for _, chainSel := range srcChains {
		// no onramp for the destination chain
		if chainSel == r.destChain {
			continue
		}

		chainSel := chainSel
		eg.Go(func() error {
			// read onramp dynamic config
			resp := onRampDestChainConfig{}
			err := r.contractReaders[chainSel].ExtendedGetLatestValue(
				ctx,
				consts.ContractNameOnRamp,
				consts.MethodNameOnRampGetDestChainConfig,
				primitives.Unconfirmed,
				map[string]any{},
				&resp,
			)
			if err != nil {
				return fmt.Errorf("failed to get onramp dynamic config: %w", err)
			}
			mu.Lock()
			result[chainSel] = resp
			mu.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return result, nil
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
	RMNHomeContractConfigDigest []byte   `json:"rmnHomeContractConfigDigest"`
	Signers                     []signer `json:"signers"`
	MinSigners                  uint64   `json:"minSigners"`
}

// versionnedConfig is used to parse the response from the RMNRemote contract's getVersionedConfig method.
// See: https://github.com/smartcontractkit/ccip/blob/ccip-develop/contracts/src/v0.8/ccip/rmn/RMNRemote.sol#L167-L169
type versionedConfig struct {
	Version uint32 `json:"version"`
	Config  config `json:"config"`
}

// Interface compliance check
var _ CCIPReader = (*ccipChainReader)(nil)
