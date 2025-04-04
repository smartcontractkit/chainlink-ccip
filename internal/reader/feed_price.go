package reader

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

// LatestRoundData is what AggregatorV3Interface returns for price feed
type LatestRoundData struct {
	RoundID         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}

// ContractTokenMap maps contracts to their token indices
type ContractTokenMap map[types.BoundContract][]cciptypes.UnknownEncodedAddress

// Number of batch operations performed (getLatestRoundData and getDecimals)
const priceReaderOperationCount = 2

// const MaxFailedPolls = 10

// FeedPricePoller fetches and caches token feed prices on a schedule
type FeedPricePoller interface {
	// GetFeedPricesUSD returns the cached feed token prices
	GetFeedPricesUSD() cciptypes.TokenPriceMap

	// RegisterTokens registers tokens to be observed
	RegisterTokens(tokens []cciptypes.UnknownEncodedAddress)

	// RegisterTokensWithInfo registers tokens with their token info
	RegisterTokensWithInfo(tokenInfo map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo)

	services.Service
}

type feedPricePoller struct {
	// Common service fields
	wg          sync.WaitGroup
	stopCh      services.StopChan
	sync        services.StateMachine
	lggr        logger.Logger
	failedPolls atomic.Uint32

	// Price reading dependencies
	chainReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade
	tokenInfo    map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo
	feedChain    cciptypes.ChainSelector
	addressCodec cciptypes.AddressCodec

	// Thread-safe state
	mu              sync.RWMutex
	feedTokenPrices cciptypes.TokenPriceMap
	lastUpdated     time.Time

	// Registry of tokens to observe
	tokensMu sync.RWMutex

	// How frequently the poller fetches prices
	pollingInterval time.Duration
}

// NewFeedPricePoller creates a new feed price poller
func NewFeedPricePoller(
	lggr logger.Logger,
	chainReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	initialTokenInfo map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo,
	feedChain cciptypes.ChainSelector,
	addressCodec cciptypes.AddressCodec,
	pollingInterval time.Duration,
) FeedPricePoller {
	// Create a copy of the initial token info to avoid potential race conditions
	tokenInfo := make(map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo)
	for token, info := range initialTokenInfo {
		tokenInfo[token] = info
	}

	return &feedPricePoller{
		stopCh:          make(chan struct{}),
		lggr:            logutil.WithComponent(lggr, "FeedPricePoller"),
		chainReaders:    chainReaders,
		tokenInfo:       tokenInfo,
		feedChain:       feedChain,
		addressCodec:    addressCodec,
		feedTokenPrices: make(cciptypes.TokenPriceMap),
		pollingInterval: pollingInterval,
		failedPolls:     atomic.Uint32{},
	}
}

// Start begins polling for token prices
func (p *feedPricePoller) Start(ctx context.Context) error {
	p.tokensMu.RLock()
	tokenCount := len(p.tokenInfo)
	p.tokensMu.RUnlock()

	p.lggr.Infow("Start Polling Feed Prices", "tokenCount", tokenCount, "interval", p.pollingInterval)
	return p.sync.StartOnce(p.Name(), func() error {
		p.wg.Add(1)
		go p.poll()
		return nil
	})
}

// Close stops the polling
func (p *feedPricePoller) Close() error {
	err := p.sync.StopOnce(p.Name(), func() error {
		defer p.wg.Wait()
		close(p.stopCh)
		return nil
	})

	if errors.Is(err, services.ErrAlreadyStopped) {
		return nil
	}
	return err
}

// Ready checks if the service is ready
func (p *feedPricePoller) Ready() error {
	return p.sync.Ready()
}

// HealthReport returns the health status
func (p *feedPricePoller) HealthReport() map[string]error {
	f := p.failedPolls.Load()
	if f >= MaxFailedPolls {
		p.sync.SvcErrBuffer.Append(fmt.Errorf("polling failed %d times in a row", MaxFailedPolls))
	}
	return map[string]error{p.Name(): p.sync.Healthy()}
}

// Name returns the service name
func (p *feedPricePoller) Name() string {
	return p.lggr.Name()
}

// poll periodically fetches token prices
func (p *feedPricePoller) poll() {
	defer p.wg.Done()
	ctx, cancel := p.stopCh.NewCtx()
	defer cancel()

	// Initial fetch once poll is called before any ticks
	if err := p.fetchPrices(ctx); err != nil {
		// Just log, don't return error as we want to keep polling
		p.lggr.Errorw("Initial fetch of feed prices failed", "err", err)
		p.failedPolls.Add(1)
	} else {
		p.failedPolls.Store(0)
	}

	ticker := time.NewTicker(p.pollingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			p.failedPolls.Store(0)
			return
		case <-ticker.C:
			if err := p.fetchPrices(ctx); err != nil {
				p.failedPolls.Add(1)
				p.lggr.Errorw("Failed to fetch feed prices", "err", err)
			} else {
				p.failedPolls.Store(0)
			}
		}
	}
}

// GetFeedPricesUSD returns the cached feed token prices
func (p *feedPricePoller) GetFeedPricesUSD() cciptypes.TokenPriceMap {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// Create a copy to avoid race conditions
	result := make(cciptypes.TokenPriceMap, len(p.feedTokenPrices))
	for token, price := range p.feedTokenPrices {
		result[token] = price
	}

	return result
}

// RegisterTokens registers tokens to be observed
func (p *feedPricePoller) RegisterTokens(tokens []cciptypes.UnknownEncodedAddress) {
	p.tokensMu.Lock()
	defer p.tokensMu.Unlock()

	// Only register tokens that have token info
	anyNewTokens := false
	for _, token := range tokens {
		_, exists := p.tokenInfo[token]
		if !exists {
			p.lggr.Warnw("Attempted to register token without token info",
				"token", token)
		} else {
			anyNewTokens = true
		}
	}

	if anyNewTokens {
		p.lggr.Debugw("Registered tokens for feed price polling",
			"tokenCount", len(tokens))
	}
}

// RegisterTokensWithInfo registers tokens with their token info
func (p *feedPricePoller) RegisterTokensWithInfo(newTokenInfo map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo) {
	if len(newTokenInfo) == 0 {
		return
	}

	p.tokensMu.Lock()
	defer p.tokensMu.Unlock()

	// Add new token info
	anyNewTokens := false
	for token, info := range newTokenInfo {
		if _, exists := p.tokenInfo[token]; !exists {
			p.tokenInfo[token] = info
			anyNewTokens = true
		}
	}

	if anyNewTokens {
		p.lggr.Infow("Registered new tokens with info for feed price polling",
			"newTokenCount", len(newTokenInfo),
			"totalTokenCount", len(p.tokenInfo))
	}
}

// fetchPrices fetches token feed prices and updates the cache
func (p *feedPricePoller) fetchPrices(ctx context.Context) error {
	// Get list of tokens from tokenInfo
	p.tokensMu.RLock()
	tokens := make([]cciptypes.UnknownEncodedAddress, 0, len(p.tokenInfo))
	for token := range p.tokenInfo {
		tokens = append(tokens, token)
	}
	p.tokensMu.RUnlock()

	if len(tokens) == 0 {
		p.lggr.Debug("No tokens to observe")
		return nil
	}

	// Fetch token prices using batch request
	tokenPrices, err := p.getFeedPricesUSD(ctx, tokens)
	if err != nil {
		return fmt.Errorf("failed to fetch feed token prices: %w", err)
	}

	p.mu.Lock()
	p.feedTokenPrices = tokenPrices
	p.lastUpdated = time.Now()
	p.mu.Unlock()

	p.lggr.Debugw("Updated feed token prices",
		"count", len(tokenPrices),
		"timestamp", p.lastUpdated,
	)

	return nil
}

// getFeedPricesUSD gets USD prices for multiple tokens using batch requests
func (p *feedPricePoller) getFeedPricesUSD(
	ctx context.Context,
	tokens []cciptypes.UnknownEncodedAddress,
) (cciptypes.TokenPriceMap, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)
	prices := make(cciptypes.TokenPriceMap)

	if p.feedChainReader() == nil {
		lggr.Debug("node does not support feed chain")
		return prices, nil
	}

	// We need to hold the lock when preparing the batch request
	// to ensure we have the latest token info
	p.tokensMu.RLock()
	batchRequest, contractTokenMap, err := p.prepareBatchRequest(tokens)
	p.tokensMu.RUnlock()

	if err != nil {
		return nil, fmt.Errorf("prepare batch request: %w", err)
	}

	// Execute batch request
	results, err := p.feedChainReader().BatchGetLatestValues(ctx, batchRequest)
	if err != nil {
		return nil, fmt.Errorf("batch request failed: %w", err)
	}

	// Process results by contract
	for boundContract, contractTokens := range contractTokenMap {
		contractResults, ok := results[boundContract]
		if !ok || len(contractResults) != priceReaderOperationCount {
			lggr.Errorw("invalid results for contract", "contract", boundContract.Address)
			continue
		}

		// Get price data
		latestRoundData, err := p.getPriceData(contractResults[0], boundContract)
		if err != nil {
			lggr.Errorw("calling getPriceData", "err", err, "contract", boundContract.Address)
			continue
		}

		// Get decimals
		decimals, err := p.getDecimals(contractResults[1], boundContract)
		if err != nil {
			lggr.Errorw("calling getDecimals", "err", err, "contract", boundContract.Address)
			continue
		}

		if latestRoundData.Answer == nil || latestRoundData.Answer.Cmp(big.NewInt(0)) <= 0 {
			lggr.Errorw("latestRoundData.Answer is nil or non positive", "contract", boundContract.Address)
			continue
		}

		// Normalize price for this contract
		normalizedContractPrice := p.normalizePrice(latestRoundData.Answer, *decimals)

		// Apply the normalized price to all tokens using this contract
		p.tokensMu.RLock()
		for _, token := range contractTokens {
			tokenInfo, exists := p.tokenInfo[token]
			if !exists {
				lggr.Warnw("token info not found when processing prices", "token", token)
				continue
			}

			price := calculateUsdPer1e18TokenAmount(normalizedContractPrice, tokenInfo.Decimals)
			if price == nil {
				lggr.Errorw("failed to calculate price", "token", token)
				continue
			}
			prices[token] = cciptypes.BigInt{Int: price}
		}
		p.tokensMu.RUnlock()
	}

	return prices, nil
}

// Input price is USD per full token, with 18 decimal precision
// Result price is USD per 1e18 of smallest token denomination, with 18 decimal precision
// Examples:
//
//	1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
//	1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
//	1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
func calculateUsdPer1e18TokenAmount(price *big.Int, decimals uint8) *big.Int {
	tmp := big.NewInt(0).Mul(price, big.NewInt(1e18))
	return tmp.Div(tmp, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}

// getPriceData extracts price data from batch results
func (p *feedPricePoller) getPriceData(
	result types.BatchReadResult,
	boundContract types.BoundContract,
) (*LatestRoundData, error) {
	priceResult, err := result.GetResult()
	if err != nil {
		return nil, fmt.Errorf("get price for contract %s: %w", boundContract.Address, err)
	}
	if priceResult == nil {
		return nil, fmt.Errorf("priceResult value is nil for contract %s", boundContract.Address)
	}
	latestRoundData, ok := priceResult.(*LatestRoundData)
	if !ok {
		return nil, fmt.Errorf("invalid price data type for contract %s: got %T", boundContract.Address, priceResult)
	}
	return latestRoundData, nil
}

// getDecimals extracts decimals from batch results
func (p *feedPricePoller) getDecimals(
	result types.BatchReadResult,
	boundContract types.BoundContract,
) (*uint8, error) {
	decimalResult, err := result.GetResult()
	if err != nil {
		return nil, fmt.Errorf("get decimals for contract %s: %w", boundContract.Address, err)
	}
	if decimalResult == nil {
		return nil, fmt.Errorf("decimalResult value is nil for contract %s", boundContract.Address)
	}
	decimals, ok := decimalResult.(*uint8)
	if !ok {
		return nil, fmt.Errorf("invalid decimals data type for contract %s: got %T", boundContract.Address, decimalResult)
	}
	return decimals, nil
}

// prepareBatchRequest creates a batch request grouped by contract
func (p *feedPricePoller) prepareBatchRequest(
	tokens []cciptypes.UnknownEncodedAddress,
) (types.BatchGetLatestValuesRequest, ContractTokenMap, error) {
	batchRequest := make(types.BatchGetLatestValuesRequest)
	contractTokenMap := make(ContractTokenMap)

	for _, token := range tokens {
		tokenInfo, ok := p.tokenInfo[token]
		if !ok {
			return nil, nil, fmt.Errorf("get tokenInfo for %s: missing token info", token)
		}

		boundContract := types.BoundContract{
			Address: string(tokenInfo.AggregatorAddress),
			Name:    consts.ContractNamePriceAggregator,
		}

		// Initialize contract batch if it doesn't exist
		if _, exists := batchRequest[boundContract]; !exists {
			batchRequest[boundContract] = make(types.ContractBatch, priceReaderOperationCount)
			batchRequest[boundContract][0] = types.BatchRead{
				ReadName:  consts.MethodNameGetLatestRoundData,
				Params:    nil,
				ReturnVal: &LatestRoundData{},
			}
			batchRequest[boundContract][1] = types.BatchRead{
				ReadName:  consts.MethodNameGetDecimals,
				Params:    nil,
				ReturnVal: new(uint8),
			}
		}

		// Track which tokens use this contract
		contractTokenMap[boundContract] = append(contractTokenMap[boundContract], token)
	}

	return batchRequest, contractTokenMap, nil
}

// normalizePrice normalizes price based on decimals
func (p *feedPricePoller) normalizePrice(price *big.Int, decimals uint8) *big.Int {
	answer := new(big.Int).Set(price)
	if decimals < 18 {
		return answer.Mul(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-int64(decimals)), nil))
	}
	if decimals > 18 {
		return answer.Div(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)-18), nil))
	}
	return answer
}

// feedChainReader returns the contract reader for the feed chain
func (p *feedPricePoller) feedChainReader() contractreader.ContractReaderFacade {
	return p.chainReaders[p.feedChain]
}

// Ensure feedPricePoller implements FeedPricePoller
var _ FeedPricePoller = (*feedPricePoller)(nil)
