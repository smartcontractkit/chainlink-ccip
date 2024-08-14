package reader

import (
	"context"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

type TokenPrices interface {
	// GetTokenPricesUSD returns the prices of the provided tokens in USD.
	// The order of the returned prices corresponds to the order of the provided tokens.
	GetTokenPricesUSD(ctx context.Context, tokens []ocr2types.Account) ([]*big.Int, error)
}

type OnchainTokenPricesReader struct {
	// Reader for the chain that will have the token prices on-chain
	// TODO: Using internal Extended Reader until we have the new changes for ChainReader that will allow us to choose the bound contract address
	ContractReader contractreader.Extended
	PriceSources   map[types.Account]pluginconfig.ArbitrumPriceSource
}

func NewOnchainTokenPricesReader(
	contractReader contractreader.Extended,
	priceSources map[types.Account]pluginconfig.ArbitrumPriceSource,
) *OnchainTokenPricesReader {
	return &OnchainTokenPricesReader{
		ContractReader: contractReader,
		PriceSources:   priceSources,
	}
}

type LatestRoundData struct {
	RoundID         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}

func (pr *OnchainTokenPricesReader) GetTokenPricesUSD(
	ctx context.Context, tokens []ocr2types.Account,
) ([]*big.Int, error) {
	prices := make([]*big.Int, len(tokens))
	eg := new(errgroup.Group)
	for idx, token := range tokens {
		idx := idx
		token := token
		eg.Go(func() error {
			//TODO: Once chainreader new changes are merged we'll need to use the bound contract
			//boundContract := commontypes.BoundContract{
			//	Address: pr.PriceSources[token].AggregatorAddress,
			//	Name: consts.ContractNamePriceAggregator,
			//}
			latestRoundData := LatestRoundData{}
			if err :=
				pr.ContractReader.GetLatestValue(
					ctx,
					consts.ContractNamePriceAggregator,
					consts.MethodNameGetLatestRoundData,
					primitives.Finalized,
					nil,
					latestRoundData); err != nil {
				return fmt.Errorf("failed to get token price for %s: %w", token, err)
			}
			prices[idx] = latestRoundData.Answer
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

// TODO: Remove this once we have the new changes for ChainReader that will allow us to choose the bound contract address
func getContractName(token ocr2types.Account) string {
	return fmt.Sprintf("AggregatorV3Interface_%s", token)
}

// Ensure OnchainTokenPricesReader implements TokenPrices
var _ TokenPrices = (*OnchainTokenPricesReader)(nil)
