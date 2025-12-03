package reader

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

type PriceReader interface {
	// GetFeedPricesUSD returns the prices of the provided tokens in USD normalized to e18.
	//	1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
	//	1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
	//	1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
	// The order of the returned prices corresponds to the order of the provided tokens.
	GetFeedPricesUSD(ctx context.Context,
		tokens []ccipocr3.UnknownEncodedAddress) (ccipocr3.TokenPriceMap, error)

	// GetFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the specified chain
	GetFeeQuoterTokenUpdates(
		ctx context.Context,
		tokens []ccipocr3.UnknownEncodedAddress,
		chain ccipocr3.ChainSelector,
	) (map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedBig, error)
}

type priceReader struct {
	lggr           logger.Logger
	chainAccessors map[ccipocr3.ChainSelector]ccipocr3.ChainAccessor
	tokenInfo      map[ccipocr3.UnknownEncodedAddress]ccipocr3.TokenInfo
	ccipReader     CCIPReader
	feedChain      ccipocr3.ChainSelector
	addressCodec   ccipocr3.AddressCodec
}

func NewPriceReader(
	lggr logger.Logger,
	chainAccessors map[ccipocr3.ChainSelector]ccipocr3.ChainAccessor,
	tokenInfo map[ccipocr3.UnknownEncodedAddress]ccipocr3.TokenInfo,
	ccipReader CCIPReader,
	feedChain ccipocr3.ChainSelector,
	addressCodec ccipocr3.AddressCodec,
) PriceReader {
	return &priceReader{
		lggr:           lggr,
		chainAccessors: chainAccessors,
		tokenInfo:      tokenInfo,
		ccipReader:     ccipReader,
		feedChain:      feedChain,
		addressCodec:   addressCodec,
	}
}

func (pr *priceReader) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
	chain ccipocr3.ChainSelector,
) (map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedBig, error) {
	lggr := logutil.WithContextValues(ctx, pr.lggr)

	tokensBytes := make([]ccipocr3.UnknownAddress, 0, len(tokens))
	for _, token := range tokens {
		tokenAddressBytes, err := pr.addressCodec.AddressStringToBytes(string(token), chain)
		if err != nil {
			lggr.Warnw("failed to convert token address to bytes", "token", token, "err", err)
			continue
		}

		tokensBytes = append(tokensBytes, tokenAddressBytes)
	}

	accessor, err := getChainAccessor(pr.chainAccessors, chain)
	if err != nil {
		// Don't return an error if the chain accessor is not found, just log warning and return nil
		lggr.Warnw("chain accessor not found", "chain", chain, "err", err)
		return nil, nil
	}

	updates, err := accessor.GetFeeQuoterTokenUpdates(ctx, tokensBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to get fee quoter token updates: %w", err)
	}

	updateMap := make(map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedBig, len(updates))
	for k, v := range updates {
		updateMap[k] = ccipocr3.TimeStampedBigFromUnix(v)
	}

	return updateMap, nil
}

// GetFeedPricesUSD gets USD prices for multiple tokens using batch requests
func (pr *priceReader) GetFeedPricesUSD(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
) (ccipocr3.TokenPriceMap, error) {
	lggr := logutil.WithContextValues(ctx, pr.lggr)

	accessor, err := getChainAccessor(pr.chainAccessors, pr.feedChain)
	if err != nil {
		// Don't return an error if the chain accessor is not found, just log and return empty map
		lggr.Debugw("chain accessor not found on node for feed chain", "chain", pr.feedChain, "err", err)
		return make(ccipocr3.TokenPriceMap), nil
	}

	prices, err := accessor.GetFeedPricesUSD(ctx, tokens, pr.tokenInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get feed prices USD: %w", err)
	}

	return prices, nil
}

// Ensure priceReader implements PriceReader
var _ PriceReader = (*priceReader)(nil)
