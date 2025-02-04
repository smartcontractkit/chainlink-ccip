package commit

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"

	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
)

// ValidateObservation validates an observation to ensure it is well-formed
func (p *Plugin) ValidateObservation(
	_ context.Context,
	outCtx ocr3types.OutcomeContext,
	q types.Query,
	ao types.AttributedObservation,
) error {
	decodedQ, err := p.ocrTypeCodec.DecodeQuery(q)
	if err != nil {
		return fmt.Errorf("decode query: %w", err)
	}

	obs, err := p.ocrTypeCodec.DecodeObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("failed to decode commit plugin observation: %w", err)
	}

	prevOutcome, err := p.ocrTypeCodec.DecodeOutcome(outCtx.PreviousOutcome)
	if err != nil {
		return fmt.Errorf("decode previous outcome: %w", err)
	}

	if err := plugincommon.ValidateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	merkleObs := attributedMerkleRootObservation{
		OracleID:    ao.Observer,
		Observation: obs.MerkleRootObs,
	}

	err = p.merkleRootProcessor.ValidateObservation(prevOutcome.MerkleRootOutcome, decodedQ.MerkleRootQuery, merkleObs)
	if err != nil {
		return fmt.Errorf("validate merkle roots observation: %w", err)
	}

	tokenObs := attributedTokenPricesObservation{
		OracleID:    ao.Observer,
		Observation: obs.TokenPriceObs,
	}
	err = p.tokenPriceProcessor.ValidateObservation(prevOutcome.TokenPriceOutcome, decodedQ.TokenPriceQuery, tokenObs)
	if err != nil {
		return fmt.Errorf("validate token prices observation: %w", err)
	}

	gasObs := attributedChainFeeObservation{
		OracleID:    ao.Observer,
		Observation: obs.ChainFeeObs,
	}
	err = p.chainFeeProcessor.ValidateObservation(prevOutcome.ChainFeeOutcome, decodedQ.ChainFeeQuery, gasObs)
	if err != nil {
		return fmt.Errorf("validate chain fee observation: %w", err)
	}

	if p.discoveryProcessor != nil {
		discoveryObs := plugincommon.AttributedObservation[dt.Observation]{
			OracleID:    ao.Observer,
			Observation: obs.DiscoveryObs,
		}
		err = p.discoveryProcessor.ValidateObservation(dt.Outcome{}, dt.Query{}, discoveryObs)
		if err != nil {
			return fmt.Errorf("validate discovery observation: %w", err)
		}
	}

	return nil
}
