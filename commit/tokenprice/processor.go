package tokenprice

import (
	"context"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/shared"
)

type Processor struct {
}

func NewProcessor() *Processor {
	return &Processor{}
}

func (w *Processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

func (w *Processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {
	return Observation{}, nil
}

func (w *Processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []shared.AttributedObservation[Observation],
) (Outcome, error) {
	return Outcome{}, nil
}

func (w *Processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao shared.AttributedObservation[Observation],
) error {
	//TODO: Validate token prices
	return nil
}

func validateObservedTokenPrices(tokenPrices []cciptypes.TokenPrice) error {
	tokensWithPrice := mapset.NewSet[types.Account]()
	for _, t := range tokenPrices {
		if tokensWithPrice.Contains(t.TokenID) {
			return fmt.Errorf("duplicate token price for token: %s", t.TokenID)
		}
		tokensWithPrice.Add(t.TokenID)

		if t.Price.IsEmpty() {
			return fmt.Errorf("token price must not be empty")
		}
	}
	return nil
}

var _ shared.PluginProcessor[Query, Observation, Outcome] = &Processor{}
