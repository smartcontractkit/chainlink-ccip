package reader

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type PriceReader interface {
	// GetTokenFeedPricesUSD returns the prices of the provided tokens in USD normalized to e18.
	//	1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
	//	1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
	//	1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
	// The order of the returned prices corresponds to the order of the provided tokens.
	GetTokenFeedPricesUSD(ctx context.Context, tokens []ocr2types.Account) ([]*big.Int, error)

	// GetFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the dest chain.
	GetFeeQuoterTokenUpdates(
		ctx context.Context,
		tokens []ocr2types.Account,
	) (map[ocr2types.Account]plugintypes.TimestampedBig, error)
}

type OnchainTokenPricesReader struct {
	// Reader for the chain that will have the token prices on-chain
	ContractReader   contractreader.ContractReaderFacade
	TokenInfo        map[types.Account]pluginconfig.TokenInfo
	FeeQuoterAddress types.Account
	// FeeQuoterEnabled Flag until we discover feeQuoter address and bind it correctly
	feeQuoterEnabled bool
}

func NewOnchainTokenPricesReader(
	contractReader contractreader.ContractReaderFacade,
	tokenInfo map[types.Account]pluginconfig.TokenInfo,
) *OnchainTokenPricesReader {
	return &OnchainTokenPricesReader{
		ContractReader: contractReader,
		TokenInfo:      tokenInfo,
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

func (pr *OnchainTokenPricesReader) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokens []ocr2types.Account,
) (map[ocr2types.Account]plugintypes.TimestampedBig, error) {
	if !pr.feeQuoterEnabled {
		return make(map[ocr2types.Account]plugintypes.TimestampedBig), nil
	}
	var updates []plugintypes.TimestampedBig

	boundContract := commontypes.BoundContract{
		Address: string(pr.FeeQuoterAddress),
		Name:    consts.ContractNameFeeQuoter,
	}
	// MethodNameFeeQuoterGetTokenPrices returns an empty update with
	// a timestamp and price of 0 if the token is not found
	if err :=
		pr.ContractReader.GetLatestValue(
			ctx,
			boundContract.ReadIdentifier(consts.MethodNameFeeQuoterGetTokenPrices),
			primitives.Unconfirmed,
			tokens,
			&updates,
		); err != nil {
		return nil, fmt.Errorf("failed to get token prices: %w", err)
	}

	updateMap := make(map[ocr2types.Account]plugintypes.TimestampedBig)
	for i, token := range tokens {
		// token not available on fee quoter
		if updates[i].Timestamp == time.Unix(0, 0) {
			continue
		}
		updateMap[token] = updates[i]
	}

	return updateMap, nil
}

func (pr *OnchainTokenPricesReader) GetTokenFeedPricesUSD(
	ctx context.Context, tokens []ocr2types.Account,
) ([]*big.Int, error) {
	prices := make([]*big.Int, len(tokens))
	eg := new(errgroup.Group)
	for idx, token := range tokens {
		idx := idx
		token := token
		eg.Go(func() error {
			// TODO: Once chainreader new changes https://github.com/smartcontractkit/chainlink-common/pull/603
			// are merged we'll need to use the bound contract
			boundContract := commontypes.BoundContract{
				Address: pr.TokenInfo[token].AggregatorAddress,
				Name:    consts.ContractNamePriceAggregator,
			}
			rawTokenPrice, err := pr.getRawTokenPriceE18Normalized(ctx, token, boundContract)
			if err != nil {
				return fmt.Errorf("failed to get token price for %s: %w", token, err)
			}
			tokenInfo, ok := pr.TokenInfo[token]
			if !ok {
				return fmt.Errorf("failed to get tokenInfo for %s: %w", token, err)
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

func (pr *OnchainTokenPricesReader) getFeedDecimals(
	ctx context.Context,
	token ocr2types.Account,
	boundContract commontypes.BoundContract,
) (uint8, error) {
	var decimals uint8
	if err :=
		pr.ContractReader.GetLatestValue(
			ctx,
			boundContract.ReadIdentifier(consts.MethodNameGetDecimals),
			primitives.Unconfirmed,
			nil,
			&decimals,
		); err != nil {
		return 0, fmt.Errorf("decimals call failed for token %s: %w", token, err)
	}

	return decimals, nil
}

func (pr *OnchainTokenPricesReader) getRawTokenPriceE18Normalized(
	ctx context.Context,
	token ocr2types.Account,
	boundContract commontypes.BoundContract,
) (*big.Int, error) {
	var latestRoundData LatestRoundData
	identifier := boundContract.ReadIdentifier(consts.MethodNameGetLatestRoundData)
	if err :=
		pr.ContractReader.GetLatestValue(
			ctx,
			identifier,
			primitives.Unconfirmed,
			nil,
			&latestRoundData,
		); err != nil {
		return nil, fmt.Errorf("latestRoundData call failed for token %s: %w", token, err)
	}

	decimals, err1 := pr.getFeedDecimals(ctx, token, boundContract)
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

// Ensure OnchainTokenPricesReader implements PriceReader
var _ PriceReader = (*OnchainTokenPricesReader)(nil)
