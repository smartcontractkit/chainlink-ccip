package chainaccessor

import (
	"context"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

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

// ContractTokenMap maps contracts to their token indices
type ContractTokenMap map[types.BoundContract][]ccipocr3.UnknownEncodedAddress

func (l *DefaultAccessor) GetFeedPricesUSD(
	ctx context.Context,
	tokens []ccipocr3.UnknownEncodedAddress,
	tokenInfoMap map[ccipocr3.UnknownEncodedAddress]ccipocr3.TokenInfo,
) (ccipocr3.TokenPriceMap, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)

	prices := make(ccipocr3.TokenPriceMap)

	// Create batch request grouped by contract
	batchRequest, contractTokenMap := l.prepareBatchRequest(tokens, tokenInfoMap)

	// Execute batch request
	results, err := l.contractReader.BatchGetLatestValues(ctx, batchRequest)
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
		latestRoundData, err := l.getPriceData(contractResults[0], boundContract)
		if err != nil {
			lggr.Errorw("calling getPriceData", err)
			continue
		}

		// Get decimals
		decimals, err := l.getDecimals(contractResults[1], boundContract)
		if err != nil {
			lggr.Errorw("calling getPriceData", err)
			continue
		}

		if latestRoundData.Answer == nil || latestRoundData.Answer.Cmp(big.NewInt(0)) <= 0 {
			lggr.Errorw("latestRoundData.Answer is nil or non positive", "contract", boundContract.Address)
			continue
		}

		// Normalize price for this contract
		normalizedContractPrice := normalizePrice(latestRoundData.Answer, *decimals)

		// Apply the normalized price to all tokens using this contract
		for _, token := range tokens {
			tokenInfo := tokenInfoMap[token]
			price := calculateUsdPer1e18TokenAmount(normalizedContractPrice, tokenInfo.Decimals)
			if price == nil {
				lggr.Errorw("failed to calculate price", "token", token)
				continue
			}
			prices[token] = ccipocr3.NewBigInt(price)
		}
	}

	return prices, nil
}

func normalizePrice(price *big.Int, decimals uint8) *big.Int {
	answer := new(big.Int).Set(price)
	if decimals < 18 {
		return new(big.Int).Mul(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18-int64(decimals)), nil))
	}
	if decimals > 18 {
		return new(big.Int).Div(answer, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)-18), nil))
	}
	return answer
}

func (l *DefaultAccessor) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokenBytes []ccipocr3.UnknownAddress,
) (map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedUnixBig, error) {
	lggr := logutil.WithContextValues(ctx, l.lggr)
	feeQuoterAddress, err := l.GetContractAddress(consts.ContractNameFeeQuoter)
	if err != nil {
		return nil, fmt.Errorf("failed to get fee quoter address: %w", err)
	}

	feeQuoterAddressStr, err := l.addrCodec.AddressBytesToString(feeQuoterAddress[:], l.chainSelector)
	if err != nil {
		lggr.Warnw("failed to convert fee quoter address to string", "chain", l.chainSelector, "err", err)
		return make(map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedUnixBig), nil
	}

	updates := make([]ccipocr3.TimestampedUnixBig, len(tokenBytes))
	boundContract := types.BoundContract{
		Address: feeQuoterAddressStr,
		Name:    consts.ContractNameFeeQuoter,
	}

	// MethodNameFeeQuoterGetTokenPrices returns an empty update with
	// a timestamp and price of 0 if the token is not found
	if err :=
		l.contractReader.GetLatestValue(
			ctx,
			boundContract.ReadIdentifier(consts.MethodNameFeeQuoterGetTokenPrices),
			primitives.Unconfirmed,
			map[string]any{
				"tokens": tokenBytes,
			},
			&updates,
		); err != nil {
		return nil, fmt.Errorf("failed to get fee quoter token updates: %w", err)
	}

	updateMap := make(map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedUnixBig)
	for i, token := range tokenBytes {
		tokenAddressStr, err := l.addrCodec.AddressBytesToString(token[:], l.chainSelector)
		if err != nil {
			lggr.Errorw("failed to convert token address to string", "token", token, "chain", l.chainSelector, "err", err)
			continue
		}

		// token not available on fee quoter
		if updates[i].Timestamp == 0 || updates[i].Value == nil || updates[i].Value.Cmp(big.NewInt(0)) == 0 {
			lggr.Debugw("empty fee quoter update found", "chain", l.chainSelector, "token", tokenAddressStr)
			continue
		}
		updateMap[ccipocr3.UnknownEncodedAddress(tokenAddressStr)] = updates[i]
	}

	return updateMap, nil
}

// prepareBatchRequest creates a batch request grouped by contract and returns the mapping of contracts to token indices
func (l *DefaultAccessor) prepareBatchRequest(
	tokens []ccipocr3.UnknownEncodedAddress,
	tokenInfoMap map[ccipocr3.UnknownEncodedAddress]ccipocr3.TokenInfo,
) (types.BatchGetLatestValuesRequest, ContractTokenMap) {
	batchRequest := make(types.BatchGetLatestValuesRequest)
	contractTokenMap := make(ContractTokenMap)

	for _, token := range tokens {
		tokenInfo, ok := tokenInfoMap[token]
		if !ok {
			l.lggr.Errorw("get tokenInfo for %s: missing token info, token skipped", token)
			continue
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

	return batchRequest, contractTokenMap
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
	return new(big.Int).Div(tmp, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}

func (l *DefaultAccessor) getPriceData(
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
		return nil, fmt.Errorf("invalid price data type for contract %s", boundContract.Address)
	}
	return latestRoundData, nil
}

func (l *DefaultAccessor) getDecimals(
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
		return nil, fmt.Errorf("invalid decimals data type for contract %s", boundContract.Address)
	}
	return decimals, nil
}
