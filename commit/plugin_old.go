package commit

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

// observationOld contains old Observation logic that is still used for backwards compatibility purposes.
func (p *Plugin) observationOld(
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

		lggr.Infow("contracts not initialized, only making discovery observations")
		logutil.LogWhenExceedFrequency(&p.lastStateLog, stateLoggingFrequency, func() {
			lggr.Infow("Commit plugin making observation", "encodedObservation", encoded, "observation", obs)
		})

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

	tokenPriceObs, chainFeeObs := p.getPriceRelatedObservations(ctx, lggr, prevOutcome, decodedQ)

	obs := committypes.Observation{
		MerkleRootObs: merkleRootObs,
		TokenPriceObs: tokenPriceObs,
		DiscoveryObs:  discoveryObs,
		ChainFeeObs:   chainFeeObs,
		FChain:        p.ObserveFChain(lggr),
	}
	p.metricsReporter.TrackObservation(
		obs, outCtx.Round) //nolint:staticcheck // we rely on Round for OTI metrics compatibility

	encoded, err := p.ocrTypeCodec.EncodeObservation(obs)
	if err != nil {
		return nil, fmt.Errorf("encode observation: %w, observation: %+v, seq nr: %d", err, obs, outCtx.SeqNr)
	}

	lggr.Infow("Commit plugin making observation", "encodedObservation", encoded, "observation", obs)
	return encoded, nil
}

func (p *Plugin) getPriceRelatedObservations(
	ctx context.Context,
	lggr logger.Logger,
	prevOutcome committypes.Outcome,
	decodedQ committypes.Query,
) (tokenprice.Observation, chainfee.Observation) {
	invalidatePriceCache := false
	waitingForPriceUpdatesToMakeItOnchain := prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber > 0

	// If we are waiting for price updates to make it onchain, but we have no more checks remaining, stop waiting.
	if waitingForPriceUpdatesToMakeItOnchain && prevOutcome.MainOutcome.RemainingPriceChecks == 0 {
		lggr.Warnw(
			"no more price checks remaining, prices of previous outcome did not make it through",
			"inflightPriceOcrSequenceNumber", prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber,
		)
		waitingForPriceUpdatesToMakeItOnchain = false
	}

	// If we still wait for price updates to make it onchain, check if the latest price report made it through.
	if waitingForPriceUpdatesToMakeItOnchain {
		latestPriceOcrSeqNum, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
		if err != nil {
			lggr.Errorw("get latest price sequence number", "err", err)
			// Observe fChain so we don't get cryptic fChain errors in the outcome phase.
			return tokenprice.Observation{
					FChain:    p.ObserveFChain(lggr),
					Timestamp: time.Now().UTC(),
				}, chainfee.Observation{
					FChain:       p.ObserveFChain(lggr),
					TimestampNow: time.Now().UTC(),
				}
		}

		if cciptypes.SeqNum(latestPriceOcrSeqNum) >= prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber {
			lggr.Infow("previous price report made it through", "ocrSeqNum", latestPriceOcrSeqNum)
			invalidatePriceCache = true
			waitingForPriceUpdatesToMakeItOnchain = false
		}
	}

	// If we are still waiting for price updates to make it onchain, don't make any price observations.
	if waitingForPriceUpdatesToMakeItOnchain {
		lggr.Infow("waiting for price updates to make it onchain, no prices observed in this round",
			"inflightPriceOcrSequenceNumber", prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber,
			"remainingPriceChecks", prevOutcome.MainOutcome.RemainingPriceChecks,
		)
		// Observe fChain so we don't get cryptic fChain errors in the outcome phase.
		return tokenprice.Observation{
				FChain:    p.ObserveFChain(lggr),
				Timestamp: time.Now().UTC(),
			}, chainfee.Observation{
				FChain:       p.ObserveFChain(lggr),
				TimestampNow: time.Now().UTC(),
			}
	}

	var tokenPriceObs tokenprice.Observation
	var chainFeeObs chainfee.Observation

	wg := sync.WaitGroup{}
	wg.Add(2)

	ctx = context.WithValue(ctx, consts.InvalidateCacheKey, invalidatePriceCache)

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

// outcomeOld contains old Outcome logic that is still used for backwards compatibility purposes.
func (p *Plugin) outcomeOld(
	ctx context.Context, outCtx ocr3types.OutcomeContext, q types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {

	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, outCtx.SeqNr, logutil.PhaseOutcome)
	lggr.Debugw("commit plugin performing outcome", "attributedObservations", aos)

	prevOutcome, err := p.ocrTypeCodec.DecodeOutcome(outCtx.PreviousOutcome)
	if err != nil {
		return nil, fmt.Errorf("decode previous outcome: %w", err)
	}

	decodedQ, err := p.ocrTypeCodec.DecodeQuery(q)
	if err != nil {
		return nil, fmt.Errorf("decode query: %w", err)
	}

	merkleRootObservations := make([]attributedMerkleRootObservation, 0, len(aos))
	tokenPricesObservations := make([]attributedTokenPricesObservation, 0, len(aos))
	chainFeeObservations := make([]attributedChainFeeObservation, 0, len(aos))
	discoveryObservations := make([]plugincommon.AttributedObservation[dt.Observation], 0, len(aos))

	for _, ao := range aos {
		obs, err := p.ocrTypeCodec.DecodeObservation(ao.Observation)
		if err != nil {
			lggr.Warnw("failed to decode observation, observation skipped",
				"err", err, "observer", ao.Observer, "observation", ao.Observation)
			continue
		}

		lggr.Debugw("Commit plugin outcome decoded observation", "observation", obs, "observer", ao.Observer)

		merkleRootObservations = append(merkleRootObservations, attributedMerkleRootObservation{
			OracleID: ao.Observer, Observation: obs.MerkleRootObs})

		tokenPricesObservations = append(tokenPricesObservations, attributedTokenPricesObservation{
			OracleID: ao.Observer, Observation: obs.TokenPriceObs})

		chainFeeObservations = append(chainFeeObservations, attributedChainFeeObservation{
			OracleID: ao.Observer, Observation: obs.ChainFeeObs})

		discoveryObservations = append(discoveryObservations, plugincommon.AttributedObservation[dt.Observation]{
			OracleID: ao.Observer, Observation: obs.DiscoveryObs})
	}

	if p.discoveryProcessor != nil {
		logutil.LogWhenExceedFrequency(&p.lastStateLog, stateLoggingFrequency, func() {
			lggr.Debugw("Processing discovery observations", "discoveryObservations", discoveryObservations)
		})
		// The outcome phase of the discovery processor is binding contracts to the chain reader. This is the reason
		// we ignore the outcome of the discovery processor.
		_, err = p.discoveryProcessor.Outcome(ctx, dt.Outcome{}, dt.Query{}, discoveryObservations)
		if err != nil {
			lggr.Errorw("failed to get discovery processor outcome", "err", err)
		} else {
			p.contractsInitialized.Store(true)
		}
	}

	merkleRootOutcome, err := p.merkleRootProcessor.Outcome(
		ctx,
		prevOutcome.MerkleRootOutcome,
		decodedQ.MerkleRootQuery,
		merkleRootObservations,
	)
	if err != nil {
		lggr.Errorw("failed to get merkle roots outcome", "err", err)
	}

	tokenPriceOutcome, err := p.tokenPriceProcessor.Outcome(
		ctx,
		prevOutcome.TokenPriceOutcome,
		decodedQ.TokenPriceQuery,
		tokenPricesObservations,
	)
	if err != nil {
		lggr.Warnw("failed to get token prices outcome", "err", err)
	}

	chainFeeOutcome, err := p.chainFeeProcessor.Outcome(
		ctx,
		prevOutcome.ChainFeeOutcome,
		decodedQ.ChainFeeQuery,
		chainFeeObservations,
	)
	if err != nil {
		lggr.Warnw("failed to get gas prices outcome", "err", err)
	}

	out := committypes.Outcome{
		MerkleRootOutcome: merkleRootOutcome,
		TokenPriceOutcome: tokenPriceOutcome,
		ChainFeeOutcome:   chainFeeOutcome,
		MainOutcome:       p.getMainOutcome(outCtx, prevOutcome, tokenPriceOutcome, chainFeeOutcome),
	}
	p.metricsReporter.TrackOutcome(
		out, outCtx.Round) //nolint:staticcheck // we rely on Round for OTI metrics compatibility

	lggr.Infow("Commit plugin finished outcome", "outcome", out)

	return p.ocrTypeCodec.EncodeOutcome(out)
}

func (p *Plugin) getMainOutcome(
	outCtx ocr3types.OutcomeContext,
	prevOutcome committypes.Outcome,
	tokenPriceOutcome tokenprice.Outcome,
	chainFeeOutcome chainfee.Outcome,
) committypes.MainOutcome {
	pricesObservedInThisRound := len(tokenPriceOutcome.TokenPrices) > 0 || len(chainFeeOutcome.GasPrices) > 0
	if pricesObservedInThisRound {
		return committypes.MainOutcome{
			InflightPriceOcrSequenceNumber: cciptypes.SeqNum(outCtx.SeqNr),
			RemainingPriceChecks:           p.offchainCfg.InflightPriceCheckRetries,
		}
	}

	waitingForPriceUpdatesToMakeItOnchain := prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber > 0
	if waitingForPriceUpdatesToMakeItOnchain {
		remainingPriceChecks := prevOutcome.MainOutcome.RemainingPriceChecks - 1
		inflightOcrSeqNum := prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber

		if remainingPriceChecks < 0 {
			remainingPriceChecks = 0
			inflightOcrSeqNum = 0
		}

		return committypes.MainOutcome{
			InflightPriceOcrSequenceNumber: inflightOcrSeqNum,
			RemainingPriceChecks:           remainingPriceChecks,
		}
	}

	return committypes.MainOutcome{
		InflightPriceOcrSequenceNumber: 0,
		RemainingPriceChecks:           0,
	}
}
