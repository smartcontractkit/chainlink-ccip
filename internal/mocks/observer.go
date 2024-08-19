package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

type Observer struct {
	*mock.Mock
}

func NewObserver() *Observer {
	return &Observer{
		Mock: &mock.Mock{},
	}
}

func (o Observer) ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain {
	args := o.Called(ctx)
	return args.Get(0).([]plugintypes.SeqNumChain)
}

func (o Observer) ObserveMerkleRoots(
	ctx context.Context,
	ranges []plugintypes.ChainRange,
) []cciptypes.MerkleRootChain {
	args := o.Called(ctx, ranges)
	return args.Get(0).([]cciptypes.MerkleRootChain)
}

func (o Observer) ObserveTokenPrices(ctx context.Context) []cciptypes.TokenPrice {
	args := o.Called(ctx)
	return args.Get(0).([]cciptypes.TokenPrice)
}

func (o Observer) ObserveGasPrices(ctx context.Context) []cciptypes.GasPriceChain {
	args := o.Called(ctx)
	return args.Get(0).([]cciptypes.GasPriceChain)
}

func (o Observer) ObserveFChain() map[cciptypes.ChainSelector]int {
	args := o.Called()
	return args.Get(0).(map[cciptypes.ChainSelector]int)
}
