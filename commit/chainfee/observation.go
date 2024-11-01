package chainfee

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"time"
)

func (p *processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {
	supportedChains, err := p.chainSupport.SupportedChains(p.oracleID)
	if err != nil {
		return Observation{}, err
	}

	// Get the fee components for all available chains that we can read from
	feeComponents := p.ccipReader.GetAvailableChainsFeeComponents(ctx, supportedChains.ToSlice())
	// Get the native token prices for all available chains that we can read from
	nativeTokenPrices := p.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, supportedChains.ToSlice())
	// Get the latest chain fee price updates for the source chains
	timestampedPriceUpdates := p.ccipReader.GetChainFeePriceUpdate(ctx, supportedChains.ToSlice())
	// Convert the timestamped price updates to a map of chain fee updates
	chainFeeUpdates := FeeUpdatesFromTimestampedBig(timestampedPriceUpdates)

	fChain := p.ObserveFChain()
	now := time.Now().UTC()

	p.lggr.Infow("observed fee components",
		"supportedChains", supportedChains.ToSlice(),
		"feeComponents", feeComponents,
		"nativeTokenPrices", nativeTokenPrices,
		"chainFeeUpdates", chainFeeUpdates,
		"fChain", fChain,
		"timestampNow", now,
	)

	return Observation{
		FChain:            fChain,
		FeeComponents:     feeComponents,
		NativeTokenPrices: nativeTokenPrices,
		ChainFeeUpdates:   chainFeeUpdates,
		TimestampNow:      now,
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
