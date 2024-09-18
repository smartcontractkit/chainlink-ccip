package reader

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	types2 "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"

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
	contractReaders map[cciptypes.ChainSelector]types.ContractReader,
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

	/*
		contracts := ContractAddresses{
			consts.ContractNameOffRamp: {
				destChain: offrampAddress,
			},
		}
		if err := reader.Sync(context.Background(), contracts); err != nil {
			lggr.Infow("failed to sync contracts", "err", err)
		}
	*/

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
	if err := r.validateReaderExistence(dest); err != nil {
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

	extendedBindings := r.contractReaders[dest].GetBindings(consts.ContractNameOffRamp)
	if len(extendedBindings) != 1 {
		return nil, fmt.Errorf("expected one binding for offRamp contract, got %d", len(extendedBindings))
	}
	contractBinding := extendedBindings[0].Binding
	iter, err := r.contractReaders[dest].QueryKey(
		ctx,
		contractBinding,
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
				TokenID: types2.Account(typeconv.AddressBytesToString(tokenPriceUpdate.SourceToken, uint64(r.destChain))),
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
	if err := r.validateReaderExistence(dest); err != nil {
		return nil, err
	}

	type ExecutionStateChangedEvent struct {
		SourceChainSelector cciptypes.ChainSelector
		SequenceNumber      cciptypes.SeqNum
		State               uint8
	}

	dataTyp := ExecutionStateChangedEvent{}

	extendedBindings := r.contractReaders[dest].GetBindings(consts.ContractNameOffRamp)
	if len(extendedBindings) != 1 {
		return nil, fmt.Errorf("expected one binding for offRamp contract, got %d", len(extendedBindings))
	}
	contractBinding := extendedBindings[0].Binding
	iter, err := r.contractReaders[dest].QueryKey(
		ctx,
		contractBinding,
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
	if err := r.validateReaderExistence(sourceChainSelector); err != nil {
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

	bindings := r.contractReaders[sourceChainSelector].GetBindings(consts.ContractNameOnRamp)
	if len(bindings) != 1 {
		return nil, fmt.Errorf("expected one binding for the OnRamp contract, got %d", len(bindings))
	}

	seq, err := r.contractReaders[sourceChainSelector].QueryKey(
		ctx,
		bindings[0].Binding,
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

	if err := r.validateReaderExistence(sourceChainSelector); err != nil {
		return 0, err
	}

	extendedBindings := r.contractReaders[sourceChainSelector].GetBindings(consts.ContractNameOnRamp)
	if len(extendedBindings) != 1 {
		return 0, fmt.Errorf("expected one binding for the OnRamp contract, got %d", len(extendedBindings))
	}
	contractBinding := extendedBindings[0].Binding

	var expectedNextSequenceNumber uint64
	err := r.contractReaders[sourceChainSelector].GetLatestValue(
		ctx,
		contractBinding.ReadIdentifier(consts.MethodNameGetExpectedNextSequenceNumber),
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
	if err := r.validateReaderExistence(destChainSelector); err != nil {
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

			extendedBindings := r.contractReaders[destChainSelector].GetBindings(consts.ContractNameNonceManager)
			if len(extendedBindings) != 1 {
				return fmt.Errorf("expected one binding for the NonceManager contract, got %d", len(extendedBindings))
			}
			contractBinding := extendedBindings[0].Binding

			var resp uint64
			err = r.contractReaders[destChainSelector].GetLatestValue(
				ctx,
				contractBinding.ReadIdentifier(consts.MethodNameGetInboundNonce),
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

func (r *ccipChainReader) GasPrices(ctx context.Context, chains []cciptypes.ChainSelector) ([]cciptypes.BigInt, error) {
	if err := r.validateWriterExistence(chains...); err != nil {
		return nil, err
	}

	eg := new(errgroup.Group)
	gasPrices := make([]cciptypes.BigInt, len(chains))
	for i, chain := range chains {
		i, chain := i, chain
		eg.Go(func() error {
			gasPrice, err := r.contractWriters[chain].GetFeeComponents(ctx)
			if err != nil {
				return fmt.Errorf("failed to get gas price: %w", err)
			}
			gasPrices[i] = cciptypes.NewBigInt(gasPrice.ExecutionFee)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return gasPrices, nil
}

func (r *ccipChainReader) DiscoverContracts(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
) (ContractAddresses, error) {
	if err := r.validateReaderExistence(destChain); err != nil {
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

	// TODO: Loookup fee quoter?

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

// bindReaderContract is a generic helper for binding contracts to readers, the addresses input is the same object
// returned by DiscoverContracts.
//
// No error is returned if contractName is not found in the contracts. This allows calling the function before all
// contracts are discovered.
//
//nolint:unused // it will be used soon.
func (r *ccipChainReader) bindReaderContract(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
	contractName string,
	address []byte,
) error {
	if err := r.validateReaderExistence(chainSel); err != nil {
		return fmt.Errorf("validate reader existence: %w", err)
	}

	encAddress := typeconv.AddressBytesToString(address, uint64(chainSel))

	// Bind the contract address to the reader.
	// If the same address exists -> no-op
	// If the address is changed -> updates the address, overwrites the existing one
	// If the contract not bound -> binds to the new address
	if err := r.contractReaders[chainSel].Bind(ctx, []types.BoundContract{
		{
			Address: encAddress,
			Name:    contractName,
		},
	}); err != nil {
		return fmt.Errorf("unable to bind %s for chain %d: %w", contractName, chainSel, err)
	}

	return nil
}

// newSync goes through the input contracts and binds them to the contract reader.
//
//nolint:unused // it will be used soon.
func (r *ccipChainReader) newSync(ctx context.Context, contracts ContractAddresses) error {
	var errs []error
	for contractName, chainSelToAddress := range contracts {
		for chainSel, address := range chainSelToAddress {
			// try to bind
			err := r.bindReaderContract(ctx, chainSel, contractName, address)
			if err != nil {
				if errors.Is(err, ErrContractReaderNotFound) {
					// don't support this chain
					continue
				}
				// some other error, gather
				// TODO: maybe return early?
				errs = append(errs, err)
			}
			// error is nil, nothing to do
		}
	}
	return errors.Join(errs...)

	// OffRamp
	/*
		offrampBytes, err := typeconv.AddressStringToBytes(r.offrampAddress, uint64(r.destChain))
		if err != nil {
			return err
		}
		contracts[consts.ContractNameOffRamp] = map[cciptypes.ChainSelector][]byte{
			r.destChain: offrampBytes,
		}
	*/
	/*
		err = r.bindReaderContract(
			ctx,
			r.destChain,
			consts.ContractNameOffRamp,
			contracts,
		)
		if err != nil {
			return fmt.Errorf("sync error (offramp): %w", err)
		}

		// OnRamps
		err = r.bindReaderContracts(
			ctx,
			maps.Keys(r.contractReaders),
			consts.ContractNameOnRamp,
			contracts,
		)
		if err != nil {
			return fmt.Errorf("sync error (onramp): %w", err)
		}

		// Nonce manager
		err = r.bindReaderContract(
			ctx,
			r.destChain,
			consts.ContractNameNonceManager,
			contracts,
		)
		if err != nil {
			return fmt.Errorf("sync error (nonce manager): %w", err)
		}
	*/
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
	if err := r.validateReaderExistence(r.destChain); err != nil {
		return nil, err
	}

	res := make(map[cciptypes.ChainSelector]sourceChainConfig)
	mu := new(sync.Mutex)

	eg := new(errgroup.Group)
	for _, chainSel := range chains {
		if chainSel == r.destChain {
			continue
		}

		chainSel := chainSel
		eg.Go(func() error {
			resp := sourceChainConfig{}
			extendedBindings := r.contractReaders[r.destChain].GetBindings(consts.ContractNameOffRamp)
			if len(extendedBindings) != 1 {
				return fmt.Errorf("expected one binding for offRamp contract, got %d", len(extendedBindings))
			}
			contractBinding := extendedBindings[0].Binding
			err := r.contractReaders[r.destChain].GetLatestValue(
				ctx,
				contractBinding.ReadIdentifier(consts.MethodNameGetSourceChainConfig),
				primitives.Unconfirmed,
				map[string]any{
					"sourceChainSelector": chainSel,
				},
				&resp,
			)
			if err != nil {
				return fmt.Errorf("failed to get source chain config: %w", err)
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

func (r *ccipChainReader) validateReaderExistence(chains ...cciptypes.ChainSelector) error {
	for _, ch := range chains {
		_, exists := r.contractReaders[ch]
		if !exists {
			return fmt.Errorf("chain %d: %w", ch, ErrContractReaderNotFound)
		}
	}
	return nil
}

func (r *ccipChainReader) validateWriterExistence(chains ...cciptypes.ChainSelector) error {
	for _, ch := range chains {
		_, exists := r.contractWriters[ch]
		if !exists {
			return fmt.Errorf("chain %d: %w", ch, ErrContractWriterNotFound)
		}
	}
	return nil
}

// getSourceChainsConfig returns the destination offRamp contract's static chain configuration.
func (r *ccipChainReader) getOfframpStaticConfig(
	ctx context.Context,
	chain cciptypes.ChainSelector,
) (offrampStaticChainConfig, error) {
	if err := r.validateReaderExistence(chain); err != nil {
		return offrampStaticChainConfig{}, err
	}

	extendedBindings := r.contractReaders[chain].GetBindings(consts.ContractNameOffRamp)
	if len(extendedBindings) != 1 {
		return offrampStaticChainConfig{},
			fmt.Errorf("expected one binding for Offramp contract, got %d", len(extendedBindings))
	}
	contractBinding := extendedBindings[0].Binding

	resp := offrampStaticChainConfig{}
	err := r.contractReaders[chain].GetLatestValue(
		ctx,
		contractBinding.ReadIdentifier(consts.MethodNameOfframpGetStaticConfig),
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
