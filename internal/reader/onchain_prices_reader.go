package reader

import (
	"context"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	commontyps "github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

type TokenPrices interface {
	// GetTokenPricesUSD returns the prices of the provided tokens in USD.
	// The order of the returned prices corresponds to the order of the provided tokens.
	GetTokenPricesUSD(ctx context.Context, tokens []ocr2types.Account) ([]*big.Int, error)
}

type OnchainTokenPricesReader struct {
	// Reader for the chain that will have the token prices on-chain
	ContractReader commontyps.ContractReader
	PriceSources   map[types.Account]pluginconfig.ArbitrumPriceSource
}

func NewOnchainTokenPricesReader(
	contractReader commontyps.ContractReader,
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
			//TODO: Once chainreader new changes https://github.com/smartcontractkit/chainlink-common/pull/603
			// are merged we'll need to use the bound contract
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
					latestRoundData,
					//boundContract,
				); err != nil {
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

// Ensure OnchainTokenPricesReader implements TokenPrices
var _ TokenPrices = (*OnchainTokenPricesReader)(nil)
