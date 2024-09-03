package commit

import (
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
)

// ValidateObservation validates an observation to ensure it is well-formed
func (p *Plugin) ValidateObservation(
	outCtx ocr3types.OutcomeContext,
	_ types.Query,
	ao types.AttributedObservation,
) error {
	obs, err := DecodeCommitPluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("failed to decode commit plugin observation: %w", err)
	}

	prevOutcome := p.decodeOutcome(outCtx.PreviousOutcome)
	if err := validateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	merkleObs := MerkleRootObservation{
		OracleID:    ao.Observer,
		Observation: obs.MerkleRootObs,
	}

	err = p.merkleRootProcessor.ValidateObservation(prevOutcome.MerkleRootOutcome, merkleroot.Query{}, merkleObs)
	if err != nil {
		return fmt.Errorf("validate merkle roots observation: %w", err)
	}

	tokenObs := TokenPricesObservation{
		OracleID:    ao.Observer,
		Observation: obs.TokenPriceObs,
	}
	err = p.tokenPriceProcessor.ValidateObservation(prevOutcome.TokenPriceOutcome, tokenprice.Query{}, tokenObs)
	if err != nil {
		return fmt.Errorf("validate token prices observation: %w", err)
	}

	gasObs := ChainFeeObservation{
		OracleID:    ao.Observer,
		Observation: obs.ChainFeeObs,
	}
	if err := p.chainFeeProcessor.ValidateObservation(prevOutcome.ChainFeeOutcome, chainfee.Query{}, gasObs); err != nil {
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
