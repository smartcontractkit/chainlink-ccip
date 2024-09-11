package commit

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	ct "github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/shared"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// ValidateObservation validates an observation to ensure it is well-formed
func (p *Plugin) ValidateObservation(
	outCtx ocr3types.OutcomeContext,
	q types.Query,
	ao types.AttributedObservation,
) error {
	decodedQ, err := committypes.DecodeCommitPluginQuery(q)
	if err != nil {
		return fmt.Errorf("decode query: %w", err)
	}

	obs, err := committypes.DecodeCommitPluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("failed to decode commit plugin observation: %w", err)
	}

	prevOutcome := p.decodeOutcome(outCtx.PreviousOutcome)
	if err := validateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	merkleObs := shared.AttributedObservation[ct.Observation]{
		OracleID:    ao.Observer,
		Observation: obs,
	}

	err = p.merkleRootProcessor.ValidateObservation(prevOutcome, decodedQ, merkleObs)
	if err != nil {
		return fmt.Errorf("validate merkle roots observation: %w", err)
	}

	tokenObs := shared.AttributedObservation[ct.TokenPriceObservation]{
		OracleID:    ao.Observer,
		Observation: obs.TokenPriceObs,
	}
	err = p.tokenPriceProcessor.ValidateObservation(prevOutcome.TokenPriceOutcome, decodedQ, tokenObs)
	if err != nil {
		return fmt.Errorf("validate token prices observation: %w", err)
	}

	//gasObs := ChainFeeObservation{
	//	OracleID:    ao.Observer,
	//	Observation: obs.ChainFeeObs,
	//}
	//err = p.chainFeeProcessor.ValidateObservation(prevOutcome.ChainFeeOutcome, decodedQ.ChainFeeQuery, gasObs)
	//if err != nil {
	//	return fmt.Errorf("validate chain fee observation: %w", err)
	//}

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
