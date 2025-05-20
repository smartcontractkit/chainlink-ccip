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
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-ccip/pkg/chain_accessor"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

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
	accessors       map[cciptypes.ChainSelector]cciptypes.ChainAccessor

	destChain      cciptypes.ChainSelector
	offrampAddress string
	configPoller   ConfigPoller
	addrCodec      cciptypes.AddressCodec
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
	var cas = make(map[cciptypes.ChainSelector]cciptypes.ChainAccessor)
	for chainSelector, cr := range contractReaders {
		crs[chainSelector] = contractreader.NewExtendedContractReader(cr)

		if contractWriters[chainSelector] == nil {
			panic(fmt.Sprintf("contract writer not found for chain %d", chainSelector))
		}
		cas[chainSelector] = chain_accessor.NewLegacyAccessor(lggr, chainSelector, crs[chainSelector], contractWriters[chainSelector], addrCodec)
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
		accessors:       cas,
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
	ctx context.Context, ts time.Time, limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	reports, err := r.accessors[r.destChain].CommitReportsGTETimestamp(ctx, ts, limit)
	if err != nil {
		return nil, fmt.Errorf("err in CommitReportsGTETimestamp: %w", err)
	}

	/*
		// r.processCommitReports
		for _, report := range reports {
			verifyReportData(report)
		}
	*/

	lggr.Debugw("decoded commit reports", "reports", reports)

	if len(reports) < limit {
		return reports, nil
	}
	return reports[:limit], nil
}

func (r *ccipChainReader) ExecutedMessages(
	ctx context.Context,
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence primitives.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	//lggr := logutil.WithContextValues(ctx, r.lggr)

	/*
		if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
			return nil, err
		}
	*/

	// trim empty ranges from rangesPerChain
	// otherwise we may get SQL errors from the chainreader.
	nonEmptyRangesPerChain := make(map[cciptypes.ChainSelector][]cciptypes.SeqNumRange)
	for chain, ranges := range rangesPerChain {
		if len(ranges) > 0 {
			nonEmptyRangesPerChain[chain] = ranges
		}
	}

	executedMessages, err := r.accessors[r.destChain].ExecutedMessages(ctx, nonEmptyRangesPerChain, confidence)
	if err != nil {
		return nil, fmt.Errorf("err in ExecutedMessages: %w", err)
	}

	return executedMessages, nil
}

func (r *ccipChainReader) MsgsBetweenSeqNums(
	ctx context.Context, sourceChainSelector cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	if err := validateExtendedReaderExistence(r.contractReaders, sourceChainSelector); err != nil {
		return nil, err
	}

	onRampAddress, err := r.accessors[sourceChainSelector].GetContractAddress(consts.ContractNameOnRamp)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address: %w", err)
	}

	msgs, err := r.accessors[sourceChainSelector].MsgsBetweenSeqNums(ctx, r.destChain, seqNumRange)
	if err != nil {
		return nil, fmt.Errorf("err in MsgsBetweenSeqNums: %w", err)
	}

	onRampAddressAfterQuery, err := r.accessors[sourceChainSelector].GetContractAddress(consts.ContractNameOnRamp)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address: %w", err)
	}

	// Ensure the onRamp address hasn't changed during the query.
	if !bytes.Equal(onRampAddress, onRampAddressAfterQuery) {
		return nil, fmt.Errorf("onRamp address has changed from %s to %s", onRampAddress, onRampAddressAfterQuery)
	}

	lggr.Infow("queried messages between sequence numbers",
		"numMsgs", len(msgs),
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String(),
	)

	for i, _ := range msgs {
		msg := msgs[i]

		/*
			if err := validateSendRequestedEvent(msg, sourceChainSelector, r.destChain, seqNumRange); err != nil {
				lggr.Errorw("validate send requested event", "err", err, "message", msg)
				continue
			}
		*/

		msg.Header.OnRamp = onRampAddress
		msgs[i] = msg
	}

	lggr.Infow("decoded messages between sequence numbers", "msgs", msgs,
		"sourceChainSelector", sourceChainSelector,
		"seqNumRange", seqNumRange.String())

	return msgs, nil
}

// LatestMsgSeqNum reads the source chain and returns the latest finalized message sequence number.
func (r *ccipChainReader) LatestMsgSeqNum(
	ctx context.Context, chain cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	//lggr := logutil.WithContextValues(ctx, r.lggr)

	/*
		if err := validateExtendedReaderExistence(r.contractReaders, chain); err != nil {
			return 0, err
		}
	*/

	sn, err := r.accessors[chain].LatestMsgSeqNum(ctx, r.destChain)
	if err != nil {
		return 0, fmt.Errorf("err in LatestMsgSeqNum: %w", err)
	}

	return sn, nil
}

// GetExpectedNextSequenceNumber implements CCIP.
func (r *ccipChainReader) GetExpectedNextSequenceNumber(
	ctx context.Context,
	sourceChainSelector cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	/*
		if err := validateExtendedReaderExistence(r.contractReaders, sourceChainSelector); err != nil {
			return 0, err
		}
	*/
	expectedNextSequenceNumber, err := r.accessors[sourceChainSelector].GetExpectedNextSequenceNumber(ctx, r.destChain)
	if err != nil {
		return 0, fmt.Errorf("err in GetExpectedNextSequenceNumber: %w", err)
	}

	if expectedNextSequenceNumber == 0 {
		return 0, fmt.Errorf("the returned expected next sequence num is 0, source chain: %d, dest chain: %d",
			sourceChainSelector, r.destChain)
	}

	lggr.Debugw("chain reader returning expected next sequence number",
		"seqNum", expectedNextSequenceNumber, "sourceChainSelector", sourceChainSelector)
	return expectedNextSequenceNumber, nil
}

// NextSeqNum returns the current sequence numbers for chains.
// This always fetches fresh data directly from contracts to ensure accuracy.
// Critical for proper message sequencing.
func (r *ccipChainReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
	//lggr := logutil.WithContextValues(ctx, r.lggr)

	nums, err := r.accessors[r.destChain].NextSeqNum(ctx, chains)
	if err != nil {
		return nil, fmt.Errorf("err in NextSeqNum: %w", err)
	}
	return nums, nil
}

func (r *ccipChainReader) Nonces(
	ctx context.Context,
	addressesByChain map[cciptypes.ChainSelector][]string,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	/*
		lggr := logutil.WithContextValues(ctx, r.lggr)
		if err := validateExtendedReaderExistence(r.contractReaders, r.destChain); err != nil {
			return nil, err
		}
	*/

	// sort the input to ensure deterministic results
	sortedChains := maps.Keys(addressesByChain)
	slices.Sort(sortedChains)

	castString := make(map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress)
	for chain, addresses := range addressesByChain {
		for _, address := range addresses {
			castString[chain] = append(castString[chain], cciptypes.UnknownEncodedAddress(address))
		}
	}

	results, err := r.accessors[r.destChain].Nonces(ctx, castString)
	if err != nil {
		return nil, fmt.Errorf("err in Nonces: %w", err)
	}
	return results, nil
}

func (r *ccipChainReader) GetChainsFeeComponents(
	ctx context.Context,
	chains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	feeComponents := make(map[cciptypes.ChainSelector]types.ChainFeeComponents, len(r.contractWriters))

	for _, chain := range chains {
		accessor, ok := r.accessors[chain]
		if !ok {
			lggr.Errorw("accessor not found", "chain", chain)
			continue
		}

		feeComponent, err := accessor.GetChainFeeComponents(ctx)
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

		feeComponents[chain] = feeComponent
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
		//reader, ok := r.contractReaders[chain]
		accessor, ok := r.accessors[chain]
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

		if nativeTokenAddress.String() == "0x" {
			lggr.Errorw("native token address is empty", "chain", chain)
			continue
		}

		update, err := accessor.GetTokenPriceUSD(ctx, cciptypes.UnknownAddress(nativeTokenAddress))
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

	feeUpdates, err := r.accessors[r.destChain].GetChainFeePriceUpdate(ctx, selectors)
	if err != nil {
		lggr.Errorw("failed to get chain fee price update", "err", err)
		return nil
	}
	return feeUpdates
}

func (r *ccipChainReader) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	// TODO: some funny business was happening here. GetRMNRemoteConfig may need more inputs and/or updates to
	//       the discovery & binding logic. Alternatively, the logic might end up as part of the config poller
	//       queries.
	cfg, err := r.accessors[r.destChain].GetRMNRemoteConfig(ctx)
	if err != nil {
		return cciptypes.RemoteConfig{}, fmt.Errorf("get RMNRemote config: %w", err)
	}
	return cfg, nil
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

func (r *ccipChainReader) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	latestSeqNr, err := r.accessors[r.destChain].GetLatestPriceSeqNr(ctx)
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

// ccipReaderInternal defines the interface that ConfigPoller needs from the ccipChainReader
// This allows for better encapsulation and easier testing through mocking
type ccipReaderInternal interface {
	// getDestChain returns the destination chain selector
	getDestChain() cciptypes.ChainSelector

	// getContractReader returns the contract reader for the specified chain
	getContractReader(chain cciptypes.ChainSelector) (contractreader.Extended, bool)
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
