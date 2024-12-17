package reader

import (
	"context"
	"fmt"
	"math/big"

	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cachekeys "github.com/smartcontractkit/chainlink-ccip/internal/cache/keys"

	"github.com/smartcontractkit/chainlink-ccip/internal/cache"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
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
	lggr          logger.Logger
	chainReaders  map[ccipocr3.ChainSelector]contractreader.ContractReaderFacade
	tokenInfo     map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo
	ccipReader    CCIPReader
	feedChain     ccipocr3.ChainSelector
	decimalsCache cache.Cache[string, uint8]
}

func NewPriceReader(
	lggr logger.Logger,
	chainReaders map[ccipocr3.ChainSelector]contractreader.ContractReaderFacade,
	tokenInfo map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo,
	ccipReader CCIPReader,
	feedChain ccipocr3.ChainSelector,
	decimalsCache cache.Cache[string, uint8],
) PriceReader {
	return &priceReader{
		lggr:          lggr,
		chainReaders:  chainReaders,
		tokenInfo:     tokenInfo,
		ccipReader:    ccipReader,
		feedChain:     feedChain,
		decimalsCache: decimalsCache,
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

func (pr *priceReader) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
	chain ccipocr3.ChainSelector,
) (map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig, error) {
	updates := make([]plugintypes.TimestampedUnixBig, len(tokens))
	updateMap := make(map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig)

	feeQuoterAddress, err := pr.ccipReader.GetContractAddress(consts.ContractNameFeeQuoter, chain)
	if err != nil {
		pr.lggr.Debugw("failed to get fee quoter address.", "chain", chain, "err", err)
		return updateMap, nil
	}

	pr.lggr.Infow("getting fee quoter token updates",
		"tokens", tokens,
		"chain", chain,
		"feeQuoterAddress", typeconv.AddressBytesToString(feeQuoterAddress, uint64(chain)),
	)

	byteTokens := make([][]byte, 0, len(tokens))
	for _, token := range tokens {
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
		return nil, nil
	}
	// MethodNameFeeQuoterGetTokenPrices returns an empty update with
	// a timestamp and price of 0 if the token is not found
	if err :=
		cr.GetLatestValue(
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

	for i, token := range tokens {
		// token not available on fee quoter
		if updates[i].Timestamp == 0 || updates[i].Value == nil || updates[i].Value.Cmp(big.NewInt(0)) == 0 {
			pr.lggr.Debugw("empty fee quoter update found",
				"chain", chain,
				"token", token,
			)
			continue
		}
		updateMap[token] = plugintypes.TimeStampedBigFromUnix(updates[i])
	}

	return updateMap, nil
}

func (pr *priceReader) GetFeedPricesUSD(
	ctx context.Context, tokens []ccipocr3.UnknownEncodedAddress,
) ([]*big.Int, error) {
	prices := make([]*big.Int, len(tokens))
	if pr.feedChainReader() == nil {
		pr.lggr.Debug("node does not support feed chain")
		return prices, nil
	}
	eg := new(errgroup.Group)
	for idx, token := range tokens {
		eg.Go(func() error {
			boundContract := commontypes.BoundContract{
				Address: string(pr.tokenInfo[token].AggregatorAddress),
				Name:    consts.ContractNamePriceAggregator,
			}
			rawTokenPrice, err := pr.getRawTokenPriceE18Normalized(ctx, token, boundContract, pr.feedChainReader())
			if err != nil {
				return fmt.Errorf("token price for %s: %w", token, err)
			}
			tokenInfo, ok := pr.tokenInfo[token]
			if !ok {
				return fmt.Errorf("get tokenInfo for %s: %w", token, err)
			}

			prices[idx] = calculateUsdPer1e18TokenAmount(rawTokenPrice, tokenInfo.Decimals)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, fmt.Errorf("failed to get all token prices successfully: %w", err)
	}

	for _, price := range prices {
		if price == nil {
			return nil, fmt.Errorf("failed to get all token prices successfully, some prices are nil")
		}
	}

	return prices, nil
}

func (pr *priceReader) getFeedDecimals(
	ctx context.Context,
	token ccipocr3.UnknownEncodedAddress,
	boundContract commontypes.BoundContract,
	feedChainReader contractreader.ContractReaderFacade,
) (uint8, error) {
	// Create cache key using token and contract address
	cacheKey := cachekeys.TokenDecimals(token, boundContract.Address)

	// Try to get from cache first
	if decimals, found := pr.decimalsCache.Get(cacheKey); found {
		pr.lggr.Debugw("Cache hit for token decimals",
			"token", token,
			"contract", boundContract.Address,
			"decimals", decimals)
		return decimals, nil
	}

	// If not in cache, get from contract
	var decimals uint8
	if err :=
		feedChainReader.GetLatestValue(
			ctx,
			boundContract.ReadIdentifier(consts.MethodNameGetDecimals),
			primitives.Unconfirmed,
			nil,
			&decimals,
		); err != nil {
		return 0, fmt.Errorf("decimals call failed for token %s: %w", token, err)
	}

	// Store in cache and log whether it was a new entry
	isNew := pr.decimalsCache.Set(cacheKey, decimals)
	pr.lggr.Debugw("Cached token decimals",
		"token", token,
		"contract", boundContract.Address,
		"decimals", decimals,
		"isNewEntry", isNew)

	return decimals, nil
}

func (pr *priceReader) getRawTokenPriceE18Normalized(
	ctx context.Context,
	token ccipocr3.UnknownEncodedAddress,
	boundContract commontypes.BoundContract,
	feedChainReader contractreader.ContractReaderFacade,
) (*big.Int, error) {
	var latestRoundData LatestRoundData
	identifier := boundContract.ReadIdentifier(consts.MethodNameGetLatestRoundData)
	if err :=
		feedChainReader.GetLatestValue(
			ctx,
			identifier,
			primitives.Unconfirmed,
			nil,
			&latestRoundData,
		); err != nil {
		return nil, fmt.Errorf("latestRoundData call failed for token %s: %w", token, err)
	}

	decimals, err1 := pr.getFeedDecimals(ctx, token, boundContract, feedChainReader)
	if err1 != nil {
		return nil, fmt.Errorf("failed to get decimals for token %s: %w", token, err1)
	}
	answer := latestRoundData.Answer
	if decimals < 18 {
		answer.Mul(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-int64(decimals)), nil))
	} else if decimals > 18 {
		answer.Div(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)-18), nil))
	}
	return answer, nil
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
