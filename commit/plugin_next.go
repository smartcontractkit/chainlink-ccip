package commit

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

// observationNext contains new Observation logic with breaking changes that should replace the existing
// Observation in the next release.
func (p *Plugin) observationNext(
	ctx context.Context, outCtx ocr3types.OutcomeContext, q types.Query,
) (types.Observation, error) {
	// Ensure that sequence number is in the context for consumption by all
	// downstream processors and the ccip reader.
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, outCtx.SeqNr, logutil.PhaseObservation)

	var discoveryObs dt.Observation
	var err error

	if p.discoveryProcessor != nil {
		tStart := time.Now()
		discoveryObs, err = p.discoveryProcessor.Observation(ctx, dt.Outcome{}, dt.Query{})
		lggr.Debugw("commit discovery observation finished",
			"duration", time.Since(tStart), "err", err)
		if err != nil {
			lggr.Errorw("failed to discover contracts", "err", err)
		}
	}

	// If the contracts are not initialized then only submit contracts discovery related observation.
	if !p.contractsInitialized.Load() && p.discoveryProcessor != nil {
		obs := committypes.Observation{DiscoveryObs: discoveryObs}
		encoded, err := p.ocrTypeCodec.EncodeObservation(obs)
		if err != nil {
			return nil, fmt.Errorf("encode discovery observation: %w, observation: %+v", err, obs)
		}

		lggr.Infow("contracts not initialized, only making discovery observations", "discoveryObs", discoveryObs)
		lggr.Infow("commit plugin making observation", "encodedObservation", encoded, "observation", obs)

		return encoded, nil
	}

	prevOutcome, err := p.ocrTypeCodec.DecodeOutcome(outCtx.PreviousOutcome)
	if err != nil {
		return nil, fmt.Errorf("decode previous outcome: %w", err)
	}

	decodedQ, err := p.ocrTypeCodec.DecodeQuery(q)
	if err != nil {
		return nil, fmt.Errorf("decode query: %w", err)
	}

	merkleRootObs, err := p.merkleRootProcessor.Observation(ctx, prevOutcome.MerkleRootOutcome, decodedQ.MerkleRootQuery)
	if err != nil {
		lggr.Errorw("get merkle root processor observation",
			"err", err,
			"prevMerkleRootOutcome", prevOutcome.MerkleRootOutcome,
			"decodedQ", decodedQ.MerkleRootQuery,
		)
	}

	obs := committypes.Observation{
		MerkleRootObs:         merkleRootObs,
		TokenPriceObs:         tokenprice.Observation{},
		DiscoveryObs:          discoveryObs,
		ChainFeeObs:           chainfee.Observation{},
		FChain:                p.ObserveFChain(lggr),
		OnChainPriceOcrSeqNum: 0,
	}

	inflightPricesExist := prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber > 0 &&
		prevOutcome.MainOutcome.RemainingPriceChecks > 0

	switch inflightPricesExist {
	case true:
		// If we have inflight prices destination chain supporting oracles only observe the onchain price ocr seq num.
		// We use this observation to check if prices are still inflight within the Outcome.
		obs.OnChainPriceOcrSeqNum, err = p.observeOnChainPriceOcrSeqNum(ctx)
		if err != nil {
			lggr.Errorw("failed to observe on-chain price seq number", "err", err)
		}
	default:
		// If we don't have inflight prices we can proceed with new price observations.
		obs.TokenPriceObs, obs.ChainFeeObs = p.getPriceObservations(ctx, lggr, prevOutcome, decodedQ)
	}

	p.metricsReporter.TrackObservation(obs)

	encoded, err := p.ocrTypeCodec.EncodeObservation(obs)
	if err != nil {
		return nil, fmt.Errorf("encode observation: %w, observation: %+v, seq nr: %d", err, obs, outCtx.SeqNr)
	}

	lggr.Infow("Commit plugin making observation", "encodedObservation", encoded, "observation", obs)
	return encoded, nil
}

func (p *Plugin) observeOnChainPriceOcrSeqNum(ctx context.Context) (uint64, error) {
	supportsDest, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		return 0, fmt.Errorf("check if oracle %d supports the destination chain: %w", p.oracleID, err)
	}

	if !supportsDest {
		return 0, nil
	}

	onChainPriceOcrSeqNum, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return 0, fmt.Errorf("get latest on-chain price seq number: %w", err)
	}

	return onChainPriceOcrSeqNum, nil
}

func (p *Plugin) getPriceObservations(
	ctx context.Context,
	lggr logger.Logger,
	prevOutcome committypes.Outcome,
	decodedQ committypes.Query,
) (tokenprice.Observation, chainfee.Observation) {
	var tokenPriceObs tokenprice.Observation
	var chainFeeObs chainfee.Observation

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		var err error
		tStart := time.Now()
		tokenPriceObs, err = p.tokenPriceProcessor.Observation(ctx, prevOutcome.TokenPriceOutcome, decodedQ.TokenPriceQuery)
		lggr.Debugw("token price observation finished", "duration", time.Since(tStart), "err", err)
		if err != nil {
			lggr.Errorw("get token price processor observation", "err", err,
				"prevTokenPriceOutcome", prevOutcome.TokenPriceOutcome)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		tStart := time.Now()
		chainFeeObs, err = p.chainFeeProcessor.Observation(ctx, prevOutcome.ChainFeeOutcome, decodedQ.ChainFeeQuery)
		lggr.Debugw("chain fee observation finished", "duration", time.Since(tStart), "err", err)
		if err != nil {
			lggr.Errorw("get gas prices processor observation",
				"err", err, "prevChainFeeOutcome", prevOutcome.ChainFeeOutcome)
		}
	}()

	wg.Wait()
	return tokenPriceObs, chainFeeObs
}

// OutcomeNext contains new Outcome logic with breaking changes that should replace the existing
// Outcome in the next release.
//
// NOTE: If you are building a feature make sure to include your changes here.
func (p *Plugin) outcomeNext(
	ctx context.Context, outCtx ocr3types.OutcomeContext, q types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	return ocr3types.Outcome{}, errors.New("not implemented")
}
