package commitrmnocb

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func (p *Plugin) ObservationQuorum(_ ocr3types.OutcomeContext, _ types.Query) (ocr3types.Quorum, error) {
	// Across all chains we require at least 2F+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
}

// Observation TODO: doc
func (p *Plugin) Observation(
	ctx context.Context, outCtx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)

	switch nextState {
	case SelectingRangesForReport:
		return CommitPluginObservation{
			OnRampMaxSeqNums:  p.ObserveOnRampMaxSeqNums(),
			OffRampMaxSeqNums: p.ObserveOffRampMaxSeqNums(),
			FChain:            p.ObserveFChain(),
		}.Encode()

	case BuildingReport:
		return CommitPluginObservation{
			MerkleRoots: p.ObserveMerkleRoots(previousOutcome.RangesSelectedForReport),
			GasPrices:   p.ObserveGasPrices(ctx),
			TokenPrices: p.ObserveTokenPrices(ctx),
			FChain:      p.ObserveFChain(),
		}.Encode()

	case WaitingForReportTransmission:
		return CommitPluginObservation{
			OffRampMaxSeqNums: p.ObserveOffRampMaxSeqNums(),
			FChain:            p.ObserveFChain(),
		}.Encode()

	default:
		p.log.Warnw("Unexpected state", "state", nextState)
		return types.Observation{}, nil
	}
}

// ObserveOnRampMaxSeqNums TODO: doc
func (p *Plugin) ObserveOnRampMaxSeqNums() []plugintypes.SeqNumChain {
	onRampMaxSeqNums, err := p.onChain.GetOnRampMaxSeqNums()
	if err != nil {
		p.log.Warnw("call to GetOnRampMaxSeqNums failed", "err", err)
	}

	return onRampMaxSeqNums
}

// ObserveOffRampMaxSeqNums TODO: doc
func (p *Plugin) ObserveOffRampMaxSeqNums() []plugintypes.SeqNumChain {
	offRampMaxSeqNums, err := p.onChain.GetOffRampMaxSeqNums()
	if err != nil {
		p.log.Warnw("call to GetOffRampMaxSeqNums failed", "err", err)
	}

	return offRampMaxSeqNums
}

// ObserveMerkleRoots TODO: doc
// Return empty array on error
func (p *Plugin) ObserveMerkleRoots(ranges []ChainRange) []MerkleRootAndChain {
	roots, err := p.onChain.GetMerkleRoots(ranges)
	if err != nil {
		p.log.Warnw("call to GetMerkleRoots failed", "err", err)
		return nil
	}

	return roots
}

// ObserveGasPrices TODO: doc
// Return empty array on error
func (p *Plugin) ObserveGasPrices(ctx context.Context) []cciptypes.GasPriceChain {
	// TODO: Should this be sourceChains or supportedChains?
	chains := p.sourceChains()
	if len(chains) == 0 {
		return []cciptypes.GasPriceChain{}
	}

	gasPrices, err := p.ccipReader.GasPrices(ctx, chains)
	if err != nil {
		p.log.Warnw("failed to get gas prices", "err", err)
		return []cciptypes.GasPriceChain{}
	}

	if len(gasPrices) != len(chains) {
		p.log.Warnw(
			"gas prices length mismatch",
			"len(gasPrices)", len(gasPrices),
			"len(chains)", len(chains),
		)
		return []cciptypes.GasPriceChain{}
	}

	gasPricesGwei := make([]cciptypes.GasPriceChain, 0, len(chains))
	for i, chain := range chains {
		gasPricesGwei = append(gasPricesGwei, cciptypes.NewGasPriceChain(gasPrices[i].Int, chain))
	}

	return gasPricesGwei
}

// ObserveTokenPrices TODO: doc
// Return empty array on error
func (p *Plugin) ObserveTokenPrices(ctx context.Context) []cciptypes.TokenPrice {
	tokenPrices, err := p.observeTokenPricesHelper(ctx)
	if err != nil {
		p.log.Warnw("call to ObserveTokenPrices failed", "err", err)
	}
	return tokenPrices
}

// ObserveTokenPricesHelper TODO: doc
// Return empty array on error
func (p *Plugin) observeTokenPricesHelper(ctx context.Context) ([]cciptypes.TokenPrice, error) {
	if p.cfg.TokenPricesObserver {
		tokenPrices, err := p.tokenPricesReader.GetTokenPricesUSD(ctx, p.cfg.PricedTokens)
		if err != nil {
			return nil, err
		}

		if len(tokenPrices) != len(p.cfg.PricedTokens) {
			return nil, fmt.Errorf("token prices length mismatch: got %d, expected %d",
				len(tokenPrices), len(p.cfg.PricedTokens))
		}

		tokenPricesUSD := make([]cciptypes.TokenPrice, 0, len(p.cfg.PricedTokens))
		for i, token := range p.cfg.PricedTokens {
			tokenPricesUSD = append(tokenPricesUSD, cciptypes.NewTokenPrice(token, tokenPrices[i]))
		}

		return tokenPricesUSD, nil
	}

	return nil, nil
}

// ObserveFChain TODO: doc
// Return empty array on error
func (p *Plugin) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		p.log.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}
