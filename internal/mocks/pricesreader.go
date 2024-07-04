package mocks

import (
	"context"
	"math/big"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/mock"
)

type TokenPricesReader struct {
	*mock.Mock
}

func NewTokenPricesReader() *TokenPricesReader {
	return &TokenPricesReader{
		Mock: &mock.Mock{},
	}
}

func (t TokenPricesReader) GetTokenPricesUSD(ctx context.Context, tokens []ocrtypes.Account) ([]*big.Int, error) {
	args := t.Called(ctx, tokens)
	return args.Get(0).([]*big.Int), args.Error(1)
}
