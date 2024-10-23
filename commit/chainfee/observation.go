package chainfee

import (
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"context"

	mapset "github.com/deckarep/golang-set/v2"

	"time"
)

func (p *processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {
	// Get the fee components for all available chains that we can read from
	feeComponents := p.ccipReader.GetAvailableChainsFeeComponents(ctx)
	feeComponentsChains, err := p.chainSupport.SupportedChains(p.oracleID)
	if err != nil {
		return Observation{}, err
	}

	availableChains := feeComponentsChains.Intersect(mapset.NewSetFromMapKeys(feeComponents)).ToSlice()
	// Get the native token prices for all available chains that we can read from
	nativeTokenPrices := p.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, availableChains)
	// Get the latest chain fee price updates for the source chains
	timestampedPriceUpdates := p.ccipReader.GetChainFeePriceUpdate(ctx, availableChains)
	// Convert the timestamped price updates to a map of chain fee updates
	chainFeeUpdates := FeeUpdatesFromTimestampedBig(timestampedPriceUpdates)

	fChain := p.ObserveFChain()

	p.lggr.Infow("observed fee components",
		"availableChains", availableChains,
		"feeComponents", feeComponents,
		"nativeTokenPrices", nativeTokenPrices,
		"chainFeeUpdates", chainFeeUpdates,
		"fChain", fChain,
	)

	return Observation{
		FChain:            fChain,
		FeeComponents:     feeComponents,
		NativeTokenPrices: nativeTokenPrices,
		ChainFeeUpdates:   chainFeeUpdates,
		TimestampNow:      time.Now().UTC(),
	}, nil
}

func (p *processor) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		p.lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}
