package commit

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ValidateObservation validates an observation to ensure it is well-formed
func (p *Plugin) ValidateObservation(
	_ context.Context,
	outCtx ocr3types.OutcomeContext,
	q types.Query,
	ao types.AttributedObservation,
) error {
	decodedQ, err := DecodeCommitPluginQuery(q)
	if err != nil {
		return fmt.Errorf("decode query: %w", err)
	}

	obs, err := DecodeCommitPluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("failed to decode commit plugin observation: %w", err)
	}

	prevOutcome := p.decodeOutcome(outCtx.PreviousOutcome)
	if err := validateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	merkleObs := merkleRootObservation{
		OracleID:    ao.Observer,
		Observation: obs.MerkleRootObs,
	}

	err = p.merkleRootProcessor.ValidateObservation(prevOutcome.MerkleRootOutcome, decodedQ.MerkleRootQuery, merkleObs)
	if err != nil {
		return fmt.Errorf("validate merkle roots observation: %w", err)
	}

	tokenObs := tokenPricesObservation{
		OracleID:    ao.Observer,
		Observation: obs.TokenPriceObs,
	}
	err = p.tokenPriceProcessor.ValidateObservation(prevOutcome.TokenPriceOutcome, decodedQ.TokenPriceQuery, tokenObs)
	if err != nil {
		return fmt.Errorf("validate token prices observation: %w", err)
	}

	gasObs := chainFeeObservation{
		OracleID:    ao.Observer,
		Observation: obs.ChainFeeObs,
	}
	err = p.chainFeeProcessor.ValidateObservation(prevOutcome.ChainFeeOutcome, decodedQ.ChainFeeQuery, gasObs)
	if err != nil {
		return fmt.Errorf("validate chain fee observation: %w", err)
	}

	return nil
}

func validateFChain(fChain map[cciptypes.ChainSelector]int) error {
	for _, f := range fChain {
		if f < 0 {
			return fmt.Errorf("fChain %d is negative", f)
		}
	}

	return nil
}
