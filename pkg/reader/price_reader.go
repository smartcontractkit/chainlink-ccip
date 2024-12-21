package reader

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/cache"
	"github.com/smartcontractkit/chainlink-ccip/internal/cache/cachekeys"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	defaultCacheExpiration = 1 * time.Minute
	cleanupInterval        = 10 * time.Minute
)

type PriceReader interface {
	// GetFeedPricesUSD returns the prices of the provided tokens in USD normalized to e18.
	//	1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
	//	1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
	//	1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
	// The order of the returned prices corresponds to the order of the provided tokens.
	GetFeedPricesUSD(ctx context.Context, tokens []ccipocr3.UnknownEncodedAddress) ([]*big.Int, error)

	// GetFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the specified chain
	GetFeeQuoterTokenUpdates(
		ctx context.Context,
		tokens []ccipocr3.UnknownEncodedAddress,
		chain ccipocr3.ChainSelector,
	) (map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig, error)
}

type priceReader struct {
	lggr           logger.Logger
	chainReaders   map[ccipocr3.ChainSelector]contractreader.ContractReaderFacade
	tokenInfo      map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo
	ccipReader     CCIPReader
	feedChain      ccipocr3.ChainSelector
	priceCache     cache.Cache[*big.Int]
	feeQuoterCache cache.Cache[plugintypes.TimestampedBig]
}

func NewPriceReader(
	lggr logger.Logger,
	chainReaders map[ccipocr3.ChainSelector]contractreader.ContractReaderFacade,
	tokenInfo map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo,
	ccipReader CCIPReader,
	feedChain ccipocr3.ChainSelector,
) PriceReader {
	return &priceReader{
		lggr:           lggr,
		chainReaders:   chainReaders,
		tokenInfo:      tokenInfo,
		ccipReader:     ccipReader,
		feedChain:      feedChain,
		priceCache:     cache.NewCustomCache[*big.Int](defaultCacheExpiration, cleanupInterval, nil),
		feeQuoterCache: cache.NewCustomCache[plugintypes.TimestampedBig](defaultCacheExpiration, cleanupInterval, nil),
	}
}

// LatestRoundData is what AggregatorV3Interface returns for price feed
// https://github.com/smartcontractkit/ccip/blob/8f3486ced41a414f724e6b12b1528db80b72346c/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol#L19
//
//nolint:lll
type LatestRoundData struct {
	RoundID         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}

// Number of batch operations performed (getLatestRoundData and getDecimals)
const priceReaderOperationCount = 2

func (pr *priceReader) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
	chain ccipocr3.ChainSelector,
) (map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig, error) {
	updateMap := make(map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig)

	// Check cache first and collect uncached tokens
	var uncachedTokens []ccipocr3.UnknownEncodedAddress
	for _, token := range tokens {
		cacheKey := cachekeys.FeeQuoterTokenUpdate(token, chain)
		if cached, found := pr.feeQuoterCache.Get(cacheKey); found {
			updateMap[token] = cached
			continue
		}
		uncachedTokens = append(uncachedTokens, token)
	}

	// If all tokens were in cache, return early
	if len(uncachedTokens) == 0 {
		return updateMap, nil
	}

	// Get fee quoter address for uncached tokens
	feeQuoterAddress, err := pr.ccipReader.GetContractAddress(consts.ContractNameFeeQuoter, chain)
	if err != nil {
		pr.lggr.Debugw("failed to get fee quoter address", "chain", chain, "err", err)
		return updateMap, nil
	}

	pr.lggr.Infow("getting fee quoter token updates for uncached tokens",
		"tokens", uncachedTokens,
		"chain", chain,
		"feeQuoterAddress", typeconv.AddressBytesToString(feeQuoterAddress, uint64(chain)),
	)

	// Convert uncached tokens to byte format
	byteTokens := make([][]byte, 0, len(uncachedTokens))
	for _, token := range uncachedTokens {
		byteToken, err := typeconv.AddressStringToBytes(string(token), uint64(chain))
		if err != nil {
			pr.lggr.Warnw("failed to convert token address to bytes", "token", token, "err", err)
			continue
		}

		byteTokens = append(byteTokens, byteToken)
	}

	boundContract := commontypes.BoundContract{
		Address: typeconv.AddressBytesToString(feeQuoterAddress[:], uint64(chain)),
		Name:    consts.ContractNameFeeQuoter,
	}

	cr, ok := pr.chainReaders[chain]
	if !ok {
		pr.lggr.Warnw("contract reader not found", "chain", chain)
		return updateMap, nil
	}

	// Get updates for uncached tokens
	updates := make([]plugintypes.TimestampedUnixBig, len(byteTokens))
	if err := cr.GetLatestValue(
		ctx,
		boundContract.ReadIdentifier(consts.MethodNameFeeQuoterGetTokenPrices),
		primitives.Unconfirmed,
		map[string]any{
			"tokens": byteTokens,
		},
		&updates,
	); err != nil {
		return nil, fmt.Errorf("failed to get fee quoter token updates: %w", err)
	}

	// Process results and update cache
	for i, token := range uncachedTokens {
		// Skip empty updates
		if updates[i].Timestamp == 0 || updates[i].Value == nil || updates[i].Value.Cmp(big.NewInt(0)) == 0 {
			pr.lggr.Debugw("empty fee quoter update found",
				"chain", chain,
				"token", token,
			)
			continue
		}

		// Convert and store update
		update := plugintypes.TimeStampedBigFromUnix(updates[i])
		updateMap[token] = update

		// Cache the result
		cacheKey := cachekeys.FeeQuoterTokenUpdate(token, chain)
		pr.feeQuoterCache.Set(cacheKey, update, cache.NoExpiration) // Use default expiration
	}

	return updateMap, nil
}

// GetFeedPricesUSD gets USD prices for multiple tokens using batch requests
func (pr *priceReader) GetFeedPricesUSD(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
) ([]*big.Int, error) {
	prices := make([]*big.Int, len(tokens))
	if pr.feedChainReader() == nil {
		pr.lggr.Debug("node does not support feed chain")
		return prices, nil
	}

	uncachedTokens, uncachedIndices := pr.collectUncachedTokens(tokens, prices)

	// If all tokens were in cache, return early
	if len(uncachedTokens) == 0 {
		return prices, nil
	}

	if err := pr.fetchAndProcessPrices(ctx, uncachedTokens, uncachedIndices, prices); err != nil {
		return nil, err
	}

	return prices, nil
}

func (pr *priceReader) collectUncachedTokens(
	tokens []ccipocr3.UnknownEncodedAddress,
	prices []*big.Int,
) ([]ccipocr3.UnknownEncodedAddress, map[int]int) {
	var uncachedTokens []ccipocr3.UnknownEncodedAddress
	uncachedIndices := make(map[int]int)

	for i, token := range tokens {
		cacheKey := cachekeys.FeedPricesUSD(token)
		if cached, found := pr.priceCache.Get(cacheKey); found {
			prices[i] = cached
		} else {
			uncachedIndices[len(uncachedTokens)] = i
			uncachedTokens = append(uncachedTokens, token)
		}
	}

	return uncachedTokens, uncachedIndices
}

// prepareBatchRequest creates a batch request grouped by contract and returns the mapping of contracts to token indices
func (pr *priceReader) prepareBatchRequest(
	tokens []ccipocr3.UnknownEncodedAddress,
) (commontypes.BatchGetLatestValuesRequest, error) {
	batchRequest := make(commontypes.BatchGetLatestValuesRequest)

	for _, token := range tokens {
		tokenInfo, ok := pr.tokenInfo[token]
		if !ok {
			return nil, fmt.Errorf("get tokenInfo for %s: missing token info", token)
		}

		boundContract := commontypes.BoundContract{
			Address: string(tokenInfo.AggregatorAddress),
			Name:    consts.ContractNamePriceAggregator,
		}

		// Initialize contract batch if it doesn't exist
		if _, exists := batchRequest[boundContract]; !exists {
			batchRequest[boundContract] = make(commontypes.ContractBatch, priceReaderOperationCount)
			batchRequest[boundContract][0] = commontypes.BatchRead{
				ReadName:  consts.MethodNameGetLatestRoundData,
				Params:    nil,
				ReturnVal: &LatestRoundData{},
			}
			batchRequest[boundContract][1] = commontypes.BatchRead{
				ReadName:  consts.MethodNameGetDecimals,
				Params:    nil,
				ReturnVal: new(uint8),
			}
		}
	}

	return batchRequest, nil
}

func (pr *priceReader) fetchAndProcessPrices(
	ctx context.Context,
	uncachedTokens []ccipocr3.UnknownEncodedAddress,
	uncachedIndices map[int]int,
	prices []*big.Int,
) error {
	batchRequest, err := pr.prepareBatchRequest(uncachedTokens)
	if err != nil {
		return fmt.Errorf("prepare batch request: %w", err)
	}

	results, err := pr.feedChainReader().BatchGetLatestValues(ctx, batchRequest)
	if err != nil {
		return fmt.Errorf("batch request failed: %w", err)
	}

	if err := pr.processPriceResults(uncachedTokens, uncachedIndices, results, prices); err != nil {
		return err
	}

	return pr.validatePrices(prices)
}

func (pr *priceReader) processPriceResults(
	uncachedTokens []ccipocr3.UnknownEncodedAddress,
	uncachedIndices map[int]int,
	results commontypes.BatchGetLatestValuesResult,
	prices []*big.Int,
) error {
	for i, token := range uncachedTokens {
		price, err := pr.getPriceFromResult(token, results)
		if err != nil {
			return err
		}

		originalIdx := uncachedIndices[i]
		prices[originalIdx] = price

		cacheKey := cachekeys.FeedPricesUSD(token)
		pr.priceCache.Set(cacheKey, price, cache.NoExpiration)
	}
	return nil
}

func (pr *priceReader) getPriceFromResult(
	token ccipocr3.UnknownEncodedAddress,
	results commontypes.BatchGetLatestValuesResult,
) (*big.Int, error) {
	tokenInfo := pr.tokenInfo[token]
	boundContract := commontypes.BoundContract{
		Address: string(tokenInfo.AggregatorAddress),
		Name:    consts.ContractNamePriceAggregator,
	}

	contractResults, ok := results[boundContract]
	if !ok || len(contractResults) != priceReaderOperationCount {
		return nil, fmt.Errorf("invalid results for contract %s", boundContract.Address)
	}

	// Get price data
	priceResult, err := contractResults[0].GetResult()
	if err != nil {
		return nil, fmt.Errorf("get price for contract %s: %w", boundContract.Address, err)
	}
	latestRoundData, ok := priceResult.(*LatestRoundData)
	if !ok {
		return nil, fmt.Errorf("invalid price data type for contract %s", boundContract.Address)
	}

	// Get decimals
	decimalResult, err := contractResults[1].GetResult()
	if err != nil {
		return nil, fmt.Errorf("get decimals for contract %s: %w", boundContract.Address, err)
	}
	decimals, ok := decimalResult.(*uint8)
	if !ok {
		return nil, fmt.Errorf("invalid decimals data type for contract %s", boundContract.Address)
	}

	normalizedContractPrice := pr.normalizePrice(latestRoundData.Answer, *decimals)
	return calculateUsdPer1e18TokenAmount(normalizedContractPrice, tokenInfo.Decimals), nil
}

func (pr *priceReader) normalizePrice(price *big.Int, decimals uint8) *big.Int {
	answer := new(big.Int).Set(price)
	if decimals < 18 {
		return answer.Mul(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-int64(decimals)), nil))
	}
	if decimals > 18 {
		return answer.Div(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)-18), nil))
	}
	return answer
}

func (pr *priceReader) validatePrices(prices []*big.Int) error {
	for _, price := range prices {
		if price == nil {
			return fmt.Errorf("failed to get all token prices successfully, some prices are nil")
		}
	}
	return nil
}

func (pr *priceReader) feedChainReader() contractreader.ContractReaderFacade {
	return pr.chainReaders[pr.feedChain]
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

// Ensure priceReader implements PriceReader
var _ PriceReader = (*priceReader)(nil)
