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
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/addressbook"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

// Default refresh period for cache if not specified
const defaultRefreshPeriod = 30 * time.Second

// TODO: unit test the implementation when the actual contract reader and writer interfaces are finalized and mocks
// can be generated.
type ccipChainReader struct {
	lggr            logger.Logger
	accessors       map[cciptypes.ChainSelector]cciptypes.ChainAccessor
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter

	destChain      cciptypes.ChainSelector
	offrampAddress string
	configPoller   ConfigPoller
	addrCodec      cciptypes.AddressCodec
	donAddressBook *addressbook.Book
}

func newCCIPChainReaderInternal(
	ctx context.Context,
	lggr logger.Logger,
	chainAccessors map[cciptypes.ChainSelector]cciptypes.ChainAccessor,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
	addrCodec cciptypes.AddressCodec,
) (*ccipChainReader, error) {
	return newCCIPChainReaderWithConfigPollerInternal(
		ctx,
		lggr,
		chainAccessors,
		contractReaders,
		contractWriters,
		destChain,
		offrampAddress,
		addrCodec,
		nil,
	)
}

func newCCIPChainReaderWithConfigPollerInternal(
	ctx context.Context,
	lggr logger.Logger,
	chainAccessors map[cciptypes.ChainSelector]cciptypes.ChainAccessor,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	contractWriters map[cciptypes.ChainSelector]types.ContractWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
	addrCodec cciptypes.AddressCodec,
	configPoller ConfigPoller,
) (*ccipChainReader, error) {
	var crs = make(map[cciptypes.ChainSelector]contractreader.Extended)
	for chainSelector, cr := range contractReaders {
		crs[chainSelector] = contractreader.NewExtendedContractReader(cr)
	}

	offrampAddrStr, err := addrCodec.AddressBytesToString(offrampAddress, destChain)
	if err != nil {
		// Panic here since the entire discovery process relies on the offramp address being valid.
		panic(fmt.Sprintf("failed to convert offramp address to string: %v", err))
	}

	reader := &ccipChainReader{
		lggr:            lggr,
		contractReaders: crs,
		contractWriters: contractWriters,
		accessors:       chainAccessors,
		destChain:       destChain,
		offrampAddress:  offrampAddrStr,
		addrCodec:       addrCodec,
		donAddressBook:  addressbook.NewBook(),
	}

	// Initialize cache with readers
	if configPoller != nil {
		reader.configPoller = configPoller
	} else {
		// TODO: maybe add config check to switch between v1 and v2 pollers?
		reader.configPoller = newConfigPollerV2(lggr, chainAccessors, destChain, defaultRefreshPeriod)
	}

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

	return reader, nil
}

// WithExtendedContractReaderTESTONLY sets the extended contract reader for the provided chain and updates the address
// book. This should only be called from tests!
func (r *ccipChainReader) WithExtendedContractReaderTESTONLY(
	ch cciptypes.ChainSelector, cr contractreader.Extended) *ccipChainReader {
	r.contractReaders[ch] = cr

	// Register the bound addresses in the address book
	for _, contractName := range consts.AllContractNames() {
		bindings := cr.GetBindings(contractName)
		if len(bindings) == 0 {
			continue
		}
		lastBinding := bindings[len(bindings)-1]

		addressBytes, err := r.addrCodec.AddressStringToBytes(lastBinding.Binding.Address, ch)
		if err != nil {
			r.lggr.Errorw("failed to convert address", "err", err)
			continue
		}

		err = r.donAddressBook.InsertOrUpdate(
			map[addressbook.ContractName]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
				addressbook.ContractName(contractName): {ch: addressBytes},
			})
		if err != nil {
			r.lggr.Errorw("failed to insert or update contract", "err", err)
			continue
		}
	}

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

// printReports is used to trim the size of the printed report. There can be a
// large number of reports, especially on Solana where gas and token
// price updates are split into separate reports. This function removes price
// only reports, and removes price data from combined reports in an attempt to
// ensure merkle roots are easier to find if logs are truncated.
func printReports(lggr logger.Logger, reports []cciptypes.CommitPluginReportWithMeta) {
	tokenPriceUpdates := 0
	gasPriceUpdates := 0
	var reportsWithRoots []cciptypes.CommitPluginReportWithMeta
	for _, report := range reports {
		gasPriceUpdates += len(report.Report.PriceUpdates.GasPriceUpdates)
		tokenPriceUpdates += len(report.Report.PriceUpdates.TokenPriceUpdates)

		if !report.Report.HasNoRoots() {
			// remove price updates from the report.
			cp := report
			cp.Report.PriceUpdates = cciptypes.PriceUpdates{}
			reportsWithRoots = append(reportsWithRoots, cp)
		}
	}
	lggr.Debugw("decoded commit reports",
		"numTokenPriceUpdates", tokenPriceUpdates,
		"numGasPriceUpdates", gasPriceUpdates,
		"reportsWithRoots", reportsWithRoots)
}

func (r *ccipChainReader) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	confidence primitives.ConfidenceLevel,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	destChainAccessor, err := getChainAccessor(r.accessors, r.destChain)
	if err != nil {
		return nil, fmt.Errorf("unable to getChainAccessor: %w", err)
	}

	reports, err := destChainAccessor.CommitReportsGTETimestamp(ctx, ts, confidence, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get commit reports from accessor: %w", err)
	}

	printReports(lggr, reports)

	return reports, nil
}

func (r *ccipChainReader) ExecutedMessages(
	ctx context.Context,
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence primitives.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	destChainAccessor, err := getChainAccessor(r.accessors, r.destChain)
	if err != nil {
		return nil, fmt.Errorf("unable to getChainAccessor: %w", err)
	}

	executedMessageSeqNumbersByChain, err := destChainAccessor.ExecutedMessages(ctx, rangesPerChain, confidence)
	if err != nil {
		return nil, fmt.Errorf("failed to get executed messages from accessor: %w", err)
	}

	return executedMessageSeqNumbersByChain, nil
}

func (r *ccipChainReader) MsgsBetweenSeqNums(
	ctx context.Context, sourceChainSelector cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	sourceChainAccessor, err := getChainAccessor(r.accessors, sourceChainSelector)
	if err != nil {
		return nil, fmt.Errorf("unable to getChainAccessor: %w", err)
	}

	onRampAddressBeforeQuery, err := sourceChainAccessor.GetContractAddress(consts.ContractNameOnRamp)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address: %w", err)
	}

	messages, err := sourceChainAccessor.MsgsBetweenSeqNums(ctx, r.destChain, seqNumRange)
	if err != nil {
		return nil, fmt.Errorf("failed to call MsgsBetweenSeqNums on sourceChainAccessor: %w", err)
	}

	onRampAddressAfterQuery, err := sourceChainAccessor.GetContractAddress(consts.ContractNameOnRamp)
	if err != nil {
		return nil, fmt.Errorf("get onRamp address after query: %w", err)
	}

	// Ensure the onRamp address hasn't changed during the query.
	if !bytes.Equal(onRampAddressBeforeQuery, onRampAddressAfterQuery) {
		return nil, fmt.Errorf("onRamp address has changed from %s to %s", onRampAddressBeforeQuery, onRampAddressAfterQuery)
	}

	return messages, nil
}

// LatestMsgSeqNum reads the source chain and returns the latest finalized message sequence number.
func (r *ccipChainReader) LatestMsgSeqNum(
	ctx context.Context, chain cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	chainAccessor, err := getChainAccessor(r.accessors, chain)
	if err != nil {
		return 0, fmt.Errorf("unable to getChainAccessor: %w", err)
	}

	seqNum, err := chainAccessor.LatestMessageTo(ctx, r.destChain)
	if err != nil {
		return 0, fmt.Errorf("failed to call accessor LatestMsgSeqNum, source chain: %d, dest chain: %d: %w",
			chain, r.destChain, err)
	}

	lggr.Debugw("chain reader returning latest onramp sequence number",
		"seqNum", seqNum, "sourceChainSelector", chain)
	return seqNum, nil
}

// GetExpectedNextSequenceNumber queries the next expected sequence number from the source
// chain OnRamp
func (r *ccipChainReader) GetExpectedNextSequenceNumber(
	ctx context.Context,
	sourceChainSelector cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	sourceChainAccessor, err := getChainAccessor(r.accessors, sourceChainSelector)
	if err != nil {
		return 0, fmt.Errorf("unable to getChainAccessor: %w", err)
	}
	expectedNextSeqNum, err := sourceChainAccessor.GetExpectedNextSequenceNumber(ctx, r.destChain)
	if err != nil {
		return 0, fmt.Errorf("failed to call accessor GetExpectedNextSequenceNumber, source chain: %d, dest chain: %d: %w",
			sourceChainSelector, r.destChain, err)
	}

	lggr.Debugw("chain accessor returning expected next sequence number",
		"seqNum", expectedNextSeqNum, "sourceChainSelector", sourceChainSelector)
	return expectedNextSeqNum, nil
}

// NextSeqNum returns the current sequence numbers for chains.
// This always fetches fresh data directly from contracts to ensure accuracy.
// Critical for proper message sequencing.
func (r *ccipChainReader) NextSeqNum(
	ctx context.Context, chains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	// Use our direct fetch method that doesn't affect the cache
	cfgs, err := r.fetchFreshSourceChainConfigsViaAccessor(ctx, r.destChain, chains)
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

func (r *ccipChainReader) Nonces(
	ctx context.Context,
	addressesByChain map[cciptypes.ChainSelector][]string,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	destChainAccessor, err := getChainAccessor(r.accessors, r.destChain)
	if err != nil {
		return nil, fmt.Errorf("unable to getChainAccessor: %w", err)
	}

	// Convert addresses to UnknownEncodedAddress
	addressesByUnknownEncodedAddress := make(map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress)
	for chain, addresses := range addressesByChain {
		for _, address := range addresses {
			unknownAddr := cciptypes.UnknownEncodedAddress(address)
			addressesByUnknownEncodedAddress[chain] = append(addressesByUnknownEncodedAddress[chain], unknownAddr)
		}
	}

	noncesByAddressByChain, err := destChainAccessor.Nonces(ctx, addressesByUnknownEncodedAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonces for addresses: %w", err)
	}

	return noncesByAddressByChain, nil
}

func (r *ccipChainReader) GetChainsFeeComponents(
	ctx context.Context,
	chains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	feeComponents := make(map[cciptypes.ChainSelector]types.ChainFeeComponents, len(r.contractWriters))

	for _, chain := range chains {
		chainAccessor, err := getChainAccessor(r.accessors, chain)
		if err != nil {
			lggr.Errorw("failed to get chain accessor", "chain", chain, "err", err)
			continue
		}

		feeComponent, err := chainAccessor.GetChainFeeComponents(ctx)
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

		feeComponents[chain] = types.ChainFeeComponents{
			ExecutionFee:        feeComponent.ExecutionFee,
			DataAvailabilityFee: feeComponent.DataAvailabilityFee,
		}
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

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, chainSelector := range selectors {
		// Capture loop variable
		chain := chainSelector

		wg.Add(1)
		go func() {
			defer wg.Done()

			chainAccessor, err := getChainAccessor(r.accessors, chain)
			if err != nil {
				lggr.Errorw("chain accessor not found, chain native price skipped", "chain", chain, "err", err)
				return
			}

			config, err := r.configPoller.GetChainConfig(ctx, chain)
			if err != nil {
				lggr.Warnw("failed to get chain config for native token address", "chain", chain, "err", err)
				return
			}
			nativeTokenAddress := config.Router.WrappedNativeAddress

			if cciptypes.UnknownAddress(nativeTokenAddress).IsZeroOrEmpty() {
				lggr.Warnw("Native token address is zero or empty. Ignore for disabled chains otherwise "+
					"check for router misconfiguration", "chain", chain, "address", nativeTokenAddress.String())
				return
			}

			price, err := chainAccessor.GetTokenPriceUSD(ctx, cciptypes.UnknownAddress(nativeTokenAddress))
			if err != nil {
				lggr.Errorw("failed to get native token price", "chain", chain, "address", nativeTokenAddress.String(), "err", err)
				return
			}

			if price.Timestamp == 0 {
				lggr.Warnw("no native token price available", "chain", chain)
				return
			}
			if price.Value == nil || price.Value.Cmp(big.NewInt(0)) <= 0 {
				lggr.Errorw("native token price is nil or non-positive", "chain", chain)
				return
			}

			mu.Lock()
			prices[chain] = cciptypes.NewBigInt(price.Value)
			mu.Unlock()
		}()
	}

	wg.Wait()

	return prices
}

// GetChainFeePriceUpdate Read from Destination chain FeeQuoter latest fee updates for the provided chains.
// It unpacks the packed fee into the ChainFeeUSDPrices struct.
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L263-L263
//
//nolint:lll
func (r *ccipChainReader) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	lggr := logutil.WithContextValues(ctx, r.lggr)
	destChainAccessor, err := getChainAccessor(r.accessors, r.destChain)
	if err != nil {
		lggr.Errorw("failed to getChainAccessor", "chain", r.destChain, "err", err)
		return nil
	}

	if len(selectors) == 0 {
		return make(map[cciptypes.ChainSelector]cciptypes.TimestampedBig) // Return a new empty map
	}

	updates, err := destChainAccessor.GetChainFeePriceUpdate(ctx, selectors)
	if err != nil {
		lggr.Errorw("failed to get chain fee price updates", "chain", r.destChain, "err", err)
		return nil
	}

	var result = make(map[cciptypes.ChainSelector]cciptypes.TimestampedBig, len(updates))
	for chain, update := range updates {
		result[chain] = cciptypes.TimeStampedBigFromUnix(update)
	}

	return result
}

// buildSigners converts internal signer representation to RMN signer info format
func (r *ccipChainReader) buildSigners(signers []cciptypes.Signer) []cciptypes.RemoteSignerInfo {
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
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return cciptypes.RemoteConfig{}, fmt.Errorf("get chain config: %w", err)
	}

	// RMNRemote address stored in the offramp static config is actually the proxy contract address.
	// Here we will get the RMNRemote address from the proxy contract by calling the RMNProxy contract.
	destChainAccessor, err := getChainAccessor(r.accessors, r.destChain)
	if err != nil {
		return cciptypes.RemoteConfig{}, fmt.Errorf("unable to getChainAccessor: %w", err)
	}
	proxyContractAddress, err := destChainAccessor.GetContractAddress(consts.ContractNameRMNRemote)
	if err != nil {
		return cciptypes.RemoteConfig{}, fmt.Errorf("get RMNRemote proxy contract address: %w", err)
	}

	rmnRemoteAddress, err := r.getRMNRemoteAddress(ctx, r.destChain, proxyContractAddress)
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
func (r *ccipChainReader) GetRmnCurseInfo(ctx context.Context) (cciptypes.CurseInfo, error) {
	// TODO: Curse requires a dedicated cache, but for now fetching it in background,
	// together with the other configurations
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return cciptypes.CurseInfo{}, fmt.Errorf("get chain config: %w", err)
	}

	return config.CurseInfo, nil
}

func getCurseInfoFromCursedSubjects(
	cursedSubjectsSet mapset.Set[[16]byte],
	destChainSelector cciptypes.ChainSelector,
) *cciptypes.CurseInfo {
	curseInfo := &cciptypes.CurseInfo{
		CursedSourceChains: make(map[cciptypes.ChainSelector]bool, cursedSubjectsSet.Cardinality()),
		CursedDestination: cursedSubjectsSet.Contains(cciptypes.GlobalCurseSubject) ||
			cursedSubjectsSet.Contains(chainSelectorToBytes16(destChainSelector)),
		GlobalCurse: cursedSubjectsSet.Contains(cciptypes.GlobalCurseSubject),
	}

	for _, cursedSubject := range cursedSubjectsSet.ToSlice() {
		if cursedSubject == cciptypes.GlobalCurseSubject {
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
	supportedChains, allChains []cciptypes.ChainSelector,
) (ContractAddresses, error) {
	var resp ContractAddresses
	var err error
	lggr := logutil.WithContextValues(ctx, r.lggr)

	if slices.Contains(supportedChains, r.destChain) {
		resp, err = r.discoverOffRampContracts(ctx, lggr, allChains)
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

	// Use wait group for parallel processing
	var wg sync.WaitGroup
	mu := new(sync.Mutex)

	// Process each source chain's OnRamp configurations
	for _, chain := range supportedChains {
		if chain == r.destChain {
			continue
		}
		// Sanity check that we have a reader for this chain
		_, crExists := r.contractReaders[chain]
		_, caExists := r.accessors[chain]
		if !crExists && !caExists {
			lggr.Errorw("Both Contract reader and chain accessor not found for a supported chain", "chain", chain)
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
// It also updates the addressbook with the new addresses.
// NOTE: You should ensure that Sync is called deterministically for every oracle in the DON to guarantee
// a consistent shared addressbook state.
func (r *ccipChainReader) Sync(ctx context.Context, contracts ContractAddresses) error {
	addressBookEntries := make(addressbook.ContractAddresses)
	for name, addrs := range contracts {
		addressBookEntries[addressbook.ContractName(name)] = addrs
	}
	if err := r.donAddressBook.InsertOrUpdate(addressBookEntries); err != nil {
		return fmt.Errorf("set address book state: %w", err)
	}

	// construct a map of chain selector to bound contract to be able to bind in parallel.
	type boundContract struct {
		name    string
		address cciptypes.UnknownAddress
	}

	var chainToContractBinding = make(map[cciptypes.ChainSelector]boundContract)
	for contractName, chainSelToAddress := range contracts {
		for chainSel, address := range chainSelToAddress {
			chainToContractBinding[chainSel] = boundContract{
				name:    contractName,
				address: address,
			}
		}
	}

	lggr := logutil.WithContextValues(ctx, r.lggr)
	var errGroup errgroup.Group
	for chainSelector, boundContract := range chainToContractBinding {
		errGroup.Go(func() error {
			// defense in depth: don't bind if the address is empty.
			// callers should ensure this but we double check here.
			if len(boundContract.address) == 0 {
				lggr.Warnw("skipping binding empty address for contract",
					"contractName", boundContract.name,
					"chainSel", chainSelector,
				)
				return nil
			}

			chainAccessor, err := getChainAccessor(r.accessors, chainSelector)
			if err != nil && errors.Is(err, ErrChainAccessorNotFound) {
				// don't support this chain, nothing to do.
				return nil
			}

			err = chainAccessor.Sync(ctx, boundContract.name, boundContract.address)
			if err != nil {
				return fmt.Errorf("bind reader contract %s on chain %s: %w", boundContract.name, chainSelector, err)
			}

			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("bind reader contracts: %w", err)
	}

	return nil
}

func (r *ccipChainReader) GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error) {
	return r.donAddressBook.GetContractAddress(addressbook.ContractName(contractName), chain)
}

// LinkPriceUSD gets the LINK price in 1e-18 USDs from the FeeQuoter contract on the destination chain.
// For example, if the price is 1 LINK = 10 USD, this function will return 10e18 (10 * 1e18). You can think of this
// function returning the price of LINK not in USD, but in a small denomination of USD, similar to returning
// the price of ETH not in ETH but in wei (1e-18 ETH).
func (r *ccipChainReader) LinkPriceUSD(ctx context.Context) (cciptypes.BigInt, error) {
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

// getDestFeeQuoterStaticConfig returns the destination chain's Fee Quoter's StaticConfig
func (r *ccipChainReader) getDestFeeQuoterStaticConfig(ctx context.Context) (cciptypes.FeeQuoterStaticConfig, error) {
	// Get from cache
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return cciptypes.FeeQuoterStaticConfig{}, fmt.Errorf("get chain config: %w", err)
	}

	if len(config.FeeQuoter.StaticConfig.LinkToken) == 0 {
		return cciptypes.FeeQuoterStaticConfig{}, fmt.Errorf("link token address is empty")
	}

	return config.FeeQuoter.StaticConfig, nil
}

// getFeeQuoterTokenPriceUSD gets the token price in USD of the given token address from the FeeQuoter contract on the
// destination chain.
func (r *ccipChainReader) getFeeQuoterTokenPriceUSD(ctx context.Context, tokenAddr []byte) (cciptypes.BigInt, error) {
	if len(tokenAddr) == 0 {
		return cciptypes.BigInt{}, fmt.Errorf("tokenAddr is empty")
	}

	accessor, err := getChainAccessor(r.accessors, r.destChain)
	if err != nil {
		return cciptypes.BigInt{}, fmt.Errorf("unable to get chain accessor for dest chain %d: %w", r.destChain, err)
	}

	timestampedPrice, err := accessor.GetTokenPriceUSD(ctx, tokenAddr)
	if err != nil {
		return cciptypes.BigInt{}, fmt.Errorf("failed to get token price from accessor, addr: %v, err: %w", tokenAddr, err)
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

func (r *ccipChainReader) fetchFreshSourceChainConfigsViaAccessor(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	chainAccessor, err := getChainAccessor(r.accessors, destChain)
	if err != nil {
		return nil, fmt.Errorf("unable to get chain accessor for dest chain %d: %w", destChain, err)
	}

	// Filter out destination chain
	filteredSourceChains := filterOutChainSelector(sourceChains, destChain)
	if len(filteredSourceChains) == 0 {
		return make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig), nil
	}

	_, sourceChainConfigsMap, err := chainAccessor.GetAllConfigsLegacy(
		ctx,
		destChain,
		filteredSourceChains,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fresh source chain configs via accessor: %w", err)
	}

	if len(sourceChainConfigsMap) == 0 {
		lggr.Infow("No fresh source chain configs found for destination chain", "destChain", destChain)
	}

	if len(sourceChainConfigsMap) != len(filteredSourceChains) {
		lggr.Infow("fetched source chain configs count mismatch",
			"expected", len(filteredSourceChains),
			"actual", len(sourceChainConfigsMap),
			"destChain", destChain)
	}

	return sourceChainConfigsMap, nil
}

// fetchFreshSourceChainConfigs always fetches fresh source chain configs directly from contracts
// without using any cached values. Use this when up-to-date data is critical, especially
// for sequence number accuracy.
// TODO: remove in favor of fetchFreshSourceChainConfigsViaAccessor() once migration to accessors is complete and
//
//	config_poller.go is removed.
func (r *ccipChainReader) fetchFreshSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	lggr := logutil.WithContextValues(ctx, r.lggr)

	reader, exists := r.contractReaders[destChain]
	if !exists {
		return nil, fmt.Errorf("no contract reader for chain %d", destChain)
	}

	// Filter out destination chain
	filteredSourceChains := filterOutChainSelector(sourceChains, destChain)
	if len(filteredSourceChains) == 0 {
		return make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig), nil
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
			ReturnVal: new(cciptypes.SourceChainConfig),
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
	configs := make(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig)

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

			cfg, ok := v.(*cciptypes.SourceChainConfig)
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

// getARM gets the RMN remote address from the RMN proxy address.
// See: https://github.com/smartcontractkit/chainlink/blob/3c7817c566c5d0aa14519c679fa85b227ac97cc5/contracts/src/v0.8/ccip/rmn/ARMProxy.sol#L40-L44
//
//nolint:lll
func (r *ccipChainReader) getRMNRemoteAddress(
	ctx context.Context,
	chain cciptypes.ChainSelector,
	rmnRemoteProxyAddress []byte) ([]byte, error) {
	chainAccessor, err := getChainAccessor(r.accessors, chain)
	if err != nil {
		return nil, fmt.Errorf("unable to getChainAccessor: %w", err)
	}
	err = chainAccessor.Sync(ctx, consts.ContractNameRMNProxy, rmnRemoteProxyAddress)
	if err != nil {
		return nil, fmt.Errorf("sync RMN proxy contract: %w", err)
	}

	// Get the address from cache instead of making a contract call
	config, err := r.configPoller.GetChainConfig(ctx, chain)
	if err != nil {
		return nil, fmt.Errorf("get chain config: %w", err)
	}

	return config.RMNProxy.RemoteAddress, nil
}

func (r *ccipChainReader) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	destChainAccessor, err := getChainAccessor(r.accessors, r.destChain)
	if err != nil {
		return 0, fmt.Errorf("unable to getChainAccessor: %w", err)
	}

	latestPriceSeqNum, err := destChainAccessor.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return 0, fmt.Errorf("get latest price sequence number from accessor: %w", err)
	}

	return uint64(latestPriceSeqNum), nil
}

func (r *ccipChainReader) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	config, err := r.configPoller.GetChainConfig(ctx, r.destChain)
	if err != nil {
		return [32]byte{}, fmt.Errorf("get chain config: %w", err)
	}

	var resp cciptypes.OCRConfigResponse
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
		commitLatestOCRConfig cciptypes.OCRConfigResponse
		execLatestOCRConfig   cciptypes.OCRConfigResponse
		staticConfig          cciptypes.OffRampStaticChainConfig
		dynamicConfig         cciptypes.OffRampDynamicChainConfig
		rmnRemoteAddress      []byte
		rmnDigestHeader       cciptypes.RMNDigestHeader
		rmnVersionConfig      cciptypes.VersionedConfig
		feeQuoterConfig       cciptypes.FeeQuoterStaticConfig
		onRampDynamicConfig   cciptypes.GetOnRampDynamicConfigResponse
		onRampDestConfig      cciptypes.OnRampDestChainConfig
		wrappedNativeAddress  []byte
		cursedSubjects        cciptypes.RMNCurseResponse
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

// Deprecated: these process...() functions have now been replaced with pkg/chainaccessor/config_processors.go to
// support only the DefaultAccessor. TODO: remove once the migration to DefaultAccessor is completed.
func (r *ccipChainReader) processConfigResults(
	chainSel cciptypes.ChainSelector,
	batchResult types.BatchGetLatestValuesResult) (cciptypes.ChainConfigSnapshot, error) {
	config := cciptypes.ChainConfigSnapshot{}

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
			return cciptypes.ChainConfigSnapshot{}, fmt.Errorf("process %s results: %w", contract.Name, err)
		}
	}

	return config, nil
}

func (r *ccipChainReader) processRouterResults(results []types.BatchReadResult) (cciptypes.RouterConfig, error) {
	if len(results) != 1 {
		return cciptypes.RouterConfig{}, fmt.Errorf("expected 1 router result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.RouterConfig{}, fmt.Errorf("get router wrapped native result: %w", err)
	}

	if bytes, ok := val.(*[]byte); ok {
		return cciptypes.RouterConfig{
			WrappedNativeAddress: cciptypes.Bytes(*bytes),
		}, nil
	}

	return cciptypes.RouterConfig{}, fmt.Errorf("invalid type for router wrapped native address: %T", val)
}

func (r *ccipChainReader) processOnRampResults(results []types.BatchReadResult) (cciptypes.OnRampConfig, error) {
	if len(results) != 2 {
		return cciptypes.OnRampConfig{}, fmt.Errorf("expected 2 OnRamp results, got %d", len(results))
	}

	var config cciptypes.OnRampConfig

	// Process DynamicConfig
	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.OnRampConfig{}, fmt.Errorf("get OnRamp dynamic config result: %w", err)
	}

	dynamicConfig, ok := val.(*cciptypes.GetOnRampDynamicConfigResponse)
	if !ok {
		return cciptypes.OnRampConfig{}, fmt.Errorf("invalid type for OnRamp dynamic config: %T", val)
	}
	config.DynamicConfig = *dynamicConfig

	// Process DestChainConfig
	val, err = results[1].GetResult()
	if err != nil {
		return cciptypes.OnRampConfig{}, fmt.Errorf("get OnRamp dest chain config result: %w", err)
	}

	destConfig, ok := val.(*cciptypes.OnRampDestChainConfig)
	if !ok {
		return cciptypes.OnRampConfig{}, fmt.Errorf("invalid type for OnRamp dest chain config: %T", val)
	}
	config.DestChainConfig = *destConfig

	return config, nil
}

// GetOnRampConfig returns the cached OnRamp configurations for a source chain
func (r *ccipChainReader) GetOnRampConfig(
	ctx context.Context,
	srcChain cciptypes.ChainSelector,
) (cciptypes.OnRampConfig, error) {
	if srcChain == r.destChain {
		return cciptypes.OnRampConfig{}, fmt.Errorf("cannot get OnRamp configs for destination chain %d", srcChain)
	}

	config, err := r.configPoller.GetChainConfig(ctx, srcChain)
	if err != nil {
		return cciptypes.OnRampConfig{}, fmt.Errorf("get chain config: %w", err)
	}

	return config.OnRamp, nil
}

func (r *ccipChainReader) processOfframpResults(
	results []types.BatchReadResult) (cciptypes.OfframpConfig, error) {

	if len(results) != 4 {
		return cciptypes.OfframpConfig{}, fmt.Errorf("expected 4 offramp results, got %d", len(results))
	}

	config := cciptypes.OfframpConfig{}

	// Define processors for each expected result
	processors := []resultProcessor{
		// CommitLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for CommitLatestOCRConfig: %T", val)
			}
			config.CommitLatestOCRConfig = *typed
			return nil
		},
		// ExecLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for ExecLatestOCRConfig: %T", val)
			}
			config.ExecLatestOCRConfig = *typed
			return nil
		},
		// StaticConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OffRampStaticChainConfig)
			if !ok {
				return fmt.Errorf("invalid type for StaticConfig: %T", val)
			}
			config.StaticConfig = *typed
			return nil
		},
		// DynamicConfig
		func(val interface{}) error {
			typed, ok := val.(*cciptypes.OffRampDynamicChainConfig)
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
			return cciptypes.OfframpConfig{}, fmt.Errorf("get offramp result %d: %w", i, err)
		}

		if err := processors[i](val); err != nil {
			return cciptypes.OfframpConfig{}, fmt.Errorf("process result %d: %w", i, err)
		}
	}

	return config, nil
}

func (r *ccipChainReader) processRMNProxyResults(results []types.BatchReadResult) (cciptypes.RMNProxyConfig, error) {
	if len(results) != 1 {
		return cciptypes.RMNProxyConfig{}, fmt.Errorf("expected 1 RMN proxy result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.RMNProxyConfig{}, fmt.Errorf("get RMN proxy result: %w", err)
	}

	if bytes, ok := val.(*[]byte); ok {
		return cciptypes.RMNProxyConfig{
			RemoteAddress: *bytes,
		}, nil
	}

	return cciptypes.RMNProxyConfig{}, fmt.Errorf("invalid type for RMN proxy remote address: %T", val)
}

func (r *ccipChainReader) processRMNRemoteResults(results []types.BatchReadResult) (
	cciptypes.RMNRemoteConfig,
	cciptypes.CurseInfo,
	error,
) {
	config := cciptypes.RMNRemoteConfig{}

	if len(results) != 3 {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("expected 3 RMN remote results, got %d", len(results))
	}

	// Process DigestHeader
	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("get RMN remote digest header result: %w", err)
	}

	typed, ok := val.(*cciptypes.RMNDigestHeader)
	if !ok {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("invalid type for RMN remote digest header: %T", val)
	}
	config.DigestHeader = *typed

	// Process VersionedConfig
	val, err = results[1].GetResult()
	if err != nil {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("get RMN remote versioned config result: %w", err)
	}

	vconf, ok := val.(*cciptypes.VersionedConfig)
	if !ok {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("invalid type for RMN remote versioned config: %T", val)
	}
	config.VersionedConfig = *vconf

	// Process CursedSubjects
	val, err = results[2].GetResult()
	if err != nil {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("get RMN remote cursed subjects result: %w", err)
	}

	c, ok := val.(*cciptypes.RMNCurseResponse)
	if !ok {
		return cciptypes.RMNRemoteConfig{}, cciptypes.CurseInfo{},
			fmt.Errorf("invalid type for RMN remote cursed subjects: %T", val)
	}
	curseInfo := *getCurseInfoFromCursedSubjects(
		mapset.NewSet(c.CursedSubjects...),
		r.destChain,
	)

	return config, curseInfo, nil
}

func (r *ccipChainReader) processFeeQuoterResults(
	results []types.BatchReadResult,
) (cciptypes.FeeQuoterConfig, error) {
	if len(results) != 1 {
		return cciptypes.FeeQuoterConfig{}, fmt.Errorf("expected 1 fee quoter result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return cciptypes.FeeQuoterConfig{}, fmt.Errorf("get fee quoter result: %w", err)
	}

	if typed, ok := val.(*cciptypes.FeeQuoterStaticConfig); ok {
		return cciptypes.FeeQuoterConfig{
			StaticConfig: *typed,
		}, nil
	}

	return cciptypes.FeeQuoterConfig{}, fmt.Errorf("invalid type for fee quoter static config: %T", val)
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
		batchResult types.BatchGetLatestValuesResult) (cciptypes.ChainConfigSnapshot, error)

	// fetchFreshSourceChainConfigs fetches source chain configurations from the specified destination chain
	fetchFreshSourceChainConfigs(
		ctx context.Context, destChain cciptypes.ChainSelector,
		sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error)
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
