package reader

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/maps"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

var _ CCIPReader = (*CachedChainReader)(nil)

type cache struct {
	sync.RWMutex
	data          NogoResponse
	lastRefresh   time.Time
	refreshPeriod time.Duration
}

type CachedChainReader struct {
	ccipReader *ccipChainReader
	cache      *cache
}

func NewCachedChainReader(
	reader *ccipChainReader,
	refreshPeriod time.Duration,
) *CachedChainReader {
	reader.lggr.Infow("Creating new cached chain reader",
		"refreshPeriod", refreshPeriod)
	return &CachedChainReader{
		ccipReader: reader,
		cache: &cache{
			refreshPeriod: refreshPeriod,
		},
	}
}

func (r *CachedChainReader) getCachedResponse(ctx context.Context) (NogoResponse, error) {
	r.ccipReader.lggr.Infow("Getting cached response")
	r.cache.RLock()
	timeSinceLastRefresh := time.Since(r.cache.lastRefresh)
	if timeSinceLastRefresh < r.cache.refreshPeriod {
		defer r.cache.RUnlock()
		r.ccipReader.lggr.Infow("Cache hit",
			"timeSinceLastRefresh", timeSinceLastRefresh,
			"refreshPeriod", r.cache.refreshPeriod)
		return r.cache.data, nil
	}
	r.cache.RUnlock()

	r.ccipReader.lggr.Infow("Cache miss, refreshing",
		"timeSinceLastRefresh", timeSinceLastRefresh,
		"refreshPeriod", r.cache.refreshPeriod)
	return r.refreshCache(ctx)
}

func (r *CachedChainReader) refreshCache(ctx context.Context) (NogoResponse, error) {
	r.cache.Lock()
	defer r.cache.Unlock()

	timeSinceLastRefresh := time.Since(r.cache.lastRefresh)
	if timeSinceLastRefresh < r.cache.refreshPeriod {
		r.ccipReader.lggr.Infow("Cache was refreshed by another goroutine",
			"timeSinceLastRefresh", timeSinceLastRefresh)
		return r.cache.data, nil
	}

	r.ccipReader.lggr.Infow("Starting cache refresh",
		"lastRefresh", r.cache.lastRefresh)

	startTime := time.Now()
	newData, err := r.ccipReader.refresh(ctx)
	refreshDuration := time.Since(startTime)

	if err != nil {
		if !r.cache.lastRefresh.IsZero() {
			r.ccipReader.lggr.Warnw("Failed to refresh cache, using old data",
				"error", err,
				"lastRefresh", r.cache.lastRefresh,
				"refreshDuration", refreshDuration)
			return r.cache.data, nil
		}
		r.ccipReader.lggr.Errorw("Failed to refresh cache, no old data available",
			"error", err,
			"refreshDuration", refreshDuration)
		return NogoResponse{}, fmt.Errorf("failed to refresh cache: %w", err)
	}

	r.cache.data = newData
	r.cache.lastRefresh = time.Now()

	r.ccipReader.lggr.Infow("Cache refresh completed",
		"refreshDuration", refreshDuration,
		"newLastRefresh", r.cache.lastRefresh)

	return newData, nil
}

func (r *CachedChainReader) refresh(ctx context.Context) (NogoResponse, error) {
	return r.getCachedResponse(ctx)
}

// CCIPReader interface implementation
func (r *CachedChainReader) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	resp, err := r.refresh(ctx)
	if err != nil {
		return [32]byte{}, err
	}

	readerData, err := r.ccipReader.GetOffRampConfigDigest(ctx, pluginType)
	if err != nil {
		return [32]byte{}, err
	}

	var respFromBatch OCRConfigResponse
	if pluginType == consts.PluginTypeCommit {
		respFromBatch = resp.Offramp.CommitLatestOCRConfig
	} else {
		respFromBatch = resp.Offramp.ExecLatestOCRConfig
	}

	r.ccipReader.lggr.Infow("Getting offramp config digest",
		"pluginType", pluginType,
		"respfromcache", respFromBatch.OCRConfig.ConfigInfo.ConfigDigest,
		"readerData", readerData)

	return respFromBatch.OCRConfig.ConfigInfo.ConfigDigest, nil
}

func (r *CachedChainReader) GetRMNRemoteConfig(ctx context.Context) (rmntypes.RemoteConfig, error) {
	resp, err := r.refresh(ctx)
	if err != nil {
		return rmntypes.RemoteConfig{}, err
	}

	respFromReader, err := r.ccipReader.GetRMNRemoteConfig(ctx)
	if err != nil {
		return rmntypes.RemoteConfig{}, err
	}

	ret := rmntypes.RemoteConfig{
		ContractAddress:  resp.RMNProxy.RMNRemoteAddress,
		ConfigDigest:     resp.RMNRemote.RMNRemoteVersionedConfig.Config.RMNHomeContractConfigDigest,
		Signers:          r.buildSigners(resp.RMNRemote.RMNRemoteVersionedConfig.Config.Signers),
		FSign:            resp.RMNRemote.RMNRemoteVersionedConfig.Config.F,
		ConfigVersion:    resp.RMNRemote.RMNRemoteVersionedConfig.Version,
		RmnReportVersion: resp.RMNRemote.RMNRemoteDigestHeader.DigestHeader,
	}

	r.ccipReader.lggr.Infow("Getting RMN remote config",
		"respfromcache", ret,
		"readerData", respFromReader)

	// Here we need to construct the RMN config from our cached response
	return rmntypes.RemoteConfig{
		ContractAddress:  resp.RMNProxy.RMNRemoteAddress,
		ConfigDigest:     resp.RMNRemote.RMNRemoteVersionedConfig.Config.RMNHomeContractConfigDigest,
		Signers:          r.buildSigners(resp.RMNRemote.RMNRemoteVersionedConfig.Config.Signers),
		FSign:            resp.RMNRemote.RMNRemoteVersionedConfig.Config.F,
		ConfigVersion:    resp.RMNRemote.RMNRemoteVersionedConfig.Version,
		RmnReportVersion: resp.RMNRemote.RMNRemoteDigestHeader.DigestHeader,
	}, nil
}

func (r *CachedChainReader) buildSigners(signers []signer) []rmntypes.RemoteSignerInfo {
	result := make([]rmntypes.RemoteSignerInfo, 0, len(signers))
	for _, s := range signers {
		result = append(result, rmntypes.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		})
	}
	return result
}

// Forward other CCIPReader interface methods to the underlying reader
func (r *CachedChainReader) CommitReportsGTETimestamp(ctx context.Context, ts time.Time, limit int) ([]plugintypes2.CommitPluginReportWithMeta, error) {
	return r.ccipReader.CommitReportsGTETimestamp(ctx, ts, limit)
}

func (r *CachedChainReader) ExecutedMessages(ctx context.Context, source cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange) ([]cciptypes.SeqNum, error) {
	return r.ccipReader.ExecutedMessages(ctx, source, seqNumRange)
}

func (r *CachedChainReader) MsgsBetweenSeqNums(ctx context.Context, sourceChainSelector cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange) ([]cciptypes.Message, error) {
	return r.ccipReader.MsgsBetweenSeqNums(ctx, sourceChainSelector, seqNumRange)
}

func (r *CachedChainReader) GetExpectedNextSequenceNumber(ctx context.Context, sourceChainSelector cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
	return r.ccipReader.GetExpectedNextSequenceNumber(ctx, sourceChainSelector)
}

func (r *CachedChainReader) NextSeqNum(ctx context.Context, chains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
	return r.ccipReader.NextSeqNum(ctx, chains)
}

func (r *CachedChainReader) Nonces(ctx context.Context, sourceChainSelector cciptypes.ChainSelector, addresses []string) (map[string]uint64, error) {
	return r.ccipReader.Nonces(ctx, sourceChainSelector, addresses)
}

func (r *CachedChainReader) GetChainsFeeComponents(ctx context.Context, chains []cciptypes.ChainSelector) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	return r.ccipReader.GetChainsFeeComponents(ctx, chains)
}

func (r *CachedChainReader) GetDestChainFeeComponents(ctx context.Context) (types.ChainFeeComponents, error) {
	return r.ccipReader.GetDestChainFeeComponents(ctx)
}

func (r *CachedChainReader) GetWrappedNativeTokenPriceUSD(ctx context.Context, selectors []cciptypes.ChainSelector) map[cciptypes.ChainSelector]cciptypes.BigInt {
	return r.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, selectors)
}

func (r *CachedChainReader) GetChainFeePriceUpdate(ctx context.Context, selectors []cciptypes.ChainSelector) map[cciptypes.ChainSelector]plugintypes.TimestampedBig {
	return r.ccipReader.GetChainFeePriceUpdate(ctx, selectors)
}

func (r *CachedChainReader) GetRmnCurseInfo(ctx context.Context, sourceChainSelectors []cciptypes.ChainSelector) (*CurseInfo, error) {
	return r.ccipReader.GetRmnCurseInfo(ctx, sourceChainSelectors)
}

func (r *CachedChainReader) GetLatestPriceSeqNr(ctx context.Context) (uint64, error) {
	return r.ccipReader.GetLatestPriceSeqNr(ctx)
}

func (r *CachedChainReader) GetMedianDataAvailabilityGasConfig(ctx context.Context) (cciptypes.DataAvailabilityGasConfig, error) {
	return r.ccipReader.GetMedianDataAvailabilityGasConfig(ctx)
}

func (r *CachedChainReader) DiscoverContracts(ctx context.Context) (ContractAddresses, error) {
	lggr := logutil.WithContextValues(ctx, r.ccipReader.lggr)
	resp := make(ContractAddresses)

	// Get cached response
	cachedResp, err := r.refresh(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get cached response: %w", err)
	}

	// Discover destination contracts if the dest chain is supported.
	if err := validateExtendedReaderExistence(r.ccipReader.contractReaders, r.ccipReader.destChain); err == nil {
		// Use data from cache for static and dynamic configs
		resp = resp.Append(
			consts.ContractNameNonceManager,
			r.ccipReader.destChain,
			cachedResp.Offramp.StaticConfig.NonceManager)
		resp = resp.Append(consts.ContractNameRMNRemote, r.ccipReader.destChain, cachedResp.Offramp.StaticConfig.RmnRemote)
		resp = resp.Append(consts.ContractNameFeeQuoter, r.ccipReader.destChain, cachedResp.Offramp.DynamicConfig.FeeQuoter)

		// Process source chain configs from cache
		selAndConf := cachedResp.Offramp.SelectorsAndConf
		for i, sourceChain := range selAndConf.Selectors {
			cfg := selAndConf.SourceChainConfigs[i]
			if !cfg.IsEnabled {
				continue
			}
			resp = resp.Append(consts.ContractNameOnRamp, cciptypes.ChainSelector(sourceChain), cfg.OnRamp)
			if len(resp[consts.ContractNameRouter][r.ccipReader.destChain]) == 0 {
				resp = resp.Append(consts.ContractNameRouter, r.ccipReader.destChain, cfg.Router)
				lggr.Infow("appending router contract address", "address", cfg.Router)
			}
		}
	}

	lggr.Infow("discovered contracts from cache", "contracts", resp)

	// The following calls are on dynamically configured chains which may not
	// be available when this function is called.
	myChains := maps.Keys(r.ccipReader.contractReaders)

	// Read onRamps for FeeQuoter in DynamicConfig.
	dynamicConfigs := r.ccipReader.getOnRampDynamicConfigs(ctx, lggr, myChains)
	for chain, cfg := range dynamicConfigs {
		resp = resp.Append(consts.ContractNameFeeQuoter, chain, cfg.DynamicConfig.FeeQuoter)
	}

	// Read onRamps for Router in DestChainConfig.
	destChainConfig := r.ccipReader.getOnRampDestChainConfig(ctx, myChains)
	for chain, cfg := range destChainConfig {
		resp = resp.Append(consts.ContractNameRouter, chain, cfg.Router)
	}

	return resp, nil
}

func (r *CachedChainReader) Sync(ctx context.Context, contracts ContractAddresses) error {
	return r.ccipReader.Sync(ctx, contracts)
}

func (r *CachedChainReader) GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error) {
	return r.ccipReader.GetContractAddress(contractName, chain)
}

func (r *CachedChainReader) LinkPriceUSD(ctx context.Context) (cciptypes.BigInt, error) {
	return r.ccipReader.LinkPriceUSD(ctx)
}

// ForceRefresh forces a cache refresh regardless of the refresh period
func (r *CachedChainReader) ForceRefresh(ctx context.Context) error {
	r.ccipReader.lggr.Infow("Force refreshing cache")
	_, err := r.refreshCache(ctx)
	if err != nil {
		r.ccipReader.lggr.Errorw("Force refresh failed",
			"error", err)
	} else {
		r.ccipReader.lggr.Infow("Force refresh completed successfully")
	}
	return err
}
