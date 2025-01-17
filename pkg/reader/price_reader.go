package reader

import (
	"context"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type PriceReader interface {
	// GetFeedPricesUSD returns the prices of the provided tokens in USD normalized to e18.
	//	1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
	//	1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
	//	1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
	// The order of the returned prices corresponds to the order of the provided tokens.
	GetFeedPricesUSD(ctx context.Context,
		tokens []ccipocr3.UnknownEncodedAddress) (map[ccipocr3.UnknownEncodedAddress]*big.Int, error)

	// GetFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the specified chain
	GetFeeQuoterTokenUpdates(
		ctx context.Context,
		tokens []ccipocr3.UnknownEncodedAddress,
		chain ccipocr3.ChainSelector,
	) (map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig, error)
}

type priceReader struct {
	lggr         logger.Logger
	chainReaders map[ccipocr3.ChainSelector]contractreader.ContractReaderFacade
	tokenInfo    map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo
	ccipReader   CCIPReader
	feedChain    ccipocr3.ChainSelector
}

func NewPriceReader(
	lggr logger.Logger,
	chainReaders map[ccipocr3.ChainSelector]contractreader.ContractReaderFacade,
	tokenInfo map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo,
	ccipReader CCIPReader,
	feedChain ccipocr3.ChainSelector,
) PriceReader {
	return &priceReader{
		lggr:         lggr,
		chainReaders: chainReaders,
		tokenInfo:    tokenInfo,
		ccipReader:   ccipReader,
		feedChain:    feedChain,
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

// ContractTokenMap maps contracts to their token indices
type ContractTokenMap map[commontypes.BoundContract][]ccipocr3.UnknownEncodedAddress

// Number of batch operations performed (getLatestRoundData and getDecimals)
const priceReaderOperationCount = 2

func (pr *priceReader) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
	chain ccipocr3.ChainSelector,
) (map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig, error) {
	lggr := logutil.WithContextValues(ctx, pr.lggr)
	updates := make([]plugintypes.TimestampedUnixBig, len(tokens))
	updateMap := make(map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig)

	feeQuoterAddress, err := pr.ccipReader.GetContractAddress(consts.ContractNameFeeQuoter, chain)
	if err != nil {
		lggr.Debugw("failed to get fee quoter address.", "chain", chain, "err", err)
		return updateMap, nil
	}

	lggr.Infow("getting fee quoter token updates",
		"tokens", tokens,
		"chain", chain,
		"feeQuoterAddress", typeconv.AddressBytesToString(feeQuoterAddress, uint64(chain)),
	)

	byteTokens := make([][]byte, 0, len(tokens))
	for _, token := range tokens {
		byteToken, err := typeconv.AddressStringToBytes(string(token), uint64(chain))
		if err != nil {
			lggr.Warnw("failed to convert token address to bytes", "token", token, "err", err)
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
		lggr.Warnw("contract reader not found", "chain", chain)
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
			lggr.Debugw("empty fee quoter update found",
				"chain", chain,
				"token", token,
			)
			continue
		}
		updateMap[token] = plugintypes.TimeStampedBigFromUnix(updates[i])
	}

	return updateMap, nil
}

// GetFeedPricesUSD gets USD prices for multiple tokens using batch requests
func (pr *priceReader) GetFeedPricesUSD(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
) (map[ccipocr3.UnknownEncodedAddress]*big.Int, error) {
	lggr := logutil.WithContextValues(ctx, pr.lggr)
	prices := make(map[ccipocr3.UnknownEncodedAddress]*big.Int, len(tokens))
	if pr.feedChainReader() == nil {
		lggr.Debug("node does not support feed chain")
		return prices, nil
	}

	// Create batch request grouped by contract
	batchRequest, contractTokenMap, err := pr.prepareBatchRequest(tokens)
	if err != nil {
		return nil, fmt.Errorf("prepare batch request: %w", err)
	}

	// Execute batch request
	results, err := pr.feedChainReader().BatchGetLatestValues(ctx, batchRequest)
	if err != nil {
		return nil, fmt.Errorf("batch request failed: %w", err)
	}

	// Process results by contract
	for boundContract, tokens := range contractTokenMap {
		contractResults, ok := results[boundContract]
		if !ok || len(contractResults) != priceReaderOperationCount {
			lggr.Errorf("invalid results for contract %s", boundContract.Address)
			continue
		}

		// Get price data
		latestRoundData, err := pr.getPriceData(contractResults[0], boundContract)
		if err != nil {
			lggr.Errorw("calling getPriceData", err)
			continue
		}

		// Get decimals
		decimals, err := pr.getDecimals(contractResults[1], boundContract)
		if err != nil {
			lggr.Errorw("calling getPriceData", err)
			continue
		}

		// Normalize price for this contract
		normalizedContractPrice := pr.normalizePrice(latestRoundData.Answer, *decimals)

		// Apply the normalized price to all tokens using this contract
		for _, token := range tokens {
			tokenInfo := pr.tokenInfo[token]
			price := calculateUsdPer1e18TokenAmount(normalizedContractPrice, tokenInfo.Decimals)
			if price == nil {
				lggr.Errorw("failed to calculate price", "token", token)
				continue
			}
			prices[token] = price
		}
	}

	return prices, nil
}

func (pr *priceReader) getPriceData(
	result commontypes.BatchReadResult,
	boundContract commontypes.BoundContract,
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
		return nil, fmt.Errorf("invalid price data type for contract %s", boundContract.Address)
	}
	return latestRoundData, nil
}

func (pr *priceReader) getDecimals(
	result commontypes.BatchReadResult,
	boundContract commontypes.BoundContract,
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
		return nil, fmt.Errorf("invalid decimals data type for contract %s", boundContract.Address)
	}
	return decimals, nil
}

// prepareBatchRequest creates a batch request grouped by contract and returns the mapping of contracts to token indices
func (pr *priceReader) prepareBatchRequest(
	tokens []ccipocr3.UnknownEncodedAddress,
) (commontypes.BatchGetLatestValuesRequest, ContractTokenMap, error) {
	batchRequest := make(commontypes.BatchGetLatestValuesRequest)
	contractTokenMap := make(ContractTokenMap)

	for _, token := range tokens {
		tokenInfo, ok := pr.tokenInfo[token]
		if !ok {
			return nil, nil, fmt.Errorf("get tokenInfo for %s: missing token info", token)
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

		// Track which tokens use this contract
		contractTokenMap[boundContract] = append(contractTokenMap[boundContract], token)
	}

	return batchRequest, contractTokenMap, nil
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
