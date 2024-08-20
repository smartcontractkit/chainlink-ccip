package reader

import (
	"context"
	"math/big"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type IPriceObserver interface {
	ObserveTokensForReport(ctx context.Context, tokens []ocr2types.Account) ([]*big.Int, error)
}

type PriceObserver struct {
	priceReader PriceReader
	// The time that a token was last updated
	lastUpdate map[ocr2types.Account]time.Time
	// DeviationPPB is the deviation in parts per billion that the price feed is allowed to deviate
	// from the last written price on-chain before we write a new price.
	DeviationPPB map[ocr2types.Account]cciptypes.BigInt
}

func NewPriceObserver(priceReader PriceReader, deviationPPB map[ocr2types.Account]cciptypes.BigInt) *PriceObserver {
	return &PriceObserver{
		priceReader:  priceReader,
		lastUpdate:   make(map[ocr2types.Account]time.Time),
		DeviationPPB: deviationPPB,
	}
}

func (p PriceObserver) ObserveTokensForReport(ctx context.Context, tokens []ocr2types.Account) ([]*big.Int, error) {
	prices, err := p.priceReader.GetTokenPricesUSD(ctx, tokens)
	if err != nil {
		return nil, err
	}

	for i, price := range prices {
		if p.lastUpdate[tokens[i]].IsZero() || Deviates(price, p.lastUpdate[tokens[i]], p.DeviationPPB[tokens[i]].Int64()) {
			p.lastUpdate[tokens[i]] = time.Now()
		}
	}

	return prices, nil

}

// Deviates checks if x1 and x2 deviates based on the provided ppb (parts per billion)
// ppb is calculated based on the smaller value of the two
// e.g, if x1 > x2, deviation_parts_per_billion = ((x1 - x2) / x2) * 1e9
func Deviates(x1, x2 *big.Int, ppb int64) bool {
	// if x1 == 0 or x2 == 0, deviates if x2 != x1, to avoid the relative division by 0 error
	if x1.BitLen() == 0 || x2.BitLen() == 0 {
		return x1.Cmp(x2) != 0
	}
	diff := big.NewInt(0).Sub(x1, x2) // diff = x1-x2
	diff.Mul(diff, big.NewInt(1e9))   // diff = diff * 1e9
	// dividing by the smaller value gives consistent ppb regardless of input order, and supports >100% deviation.
	if x1.Cmp(x2) > 0 {
		diff.Div(diff, x2)
	} else {
		diff.Div(diff, x1)
	}
	return diff.CmpAbs(big.NewInt(ppb)) > 0 // abs(diff) > ppb
}
