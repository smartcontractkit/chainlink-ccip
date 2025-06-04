package commit

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

	inflightPricesExist := prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber > 0
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
//
//nolint:gocyclo
func (p *Plugin) outcomeNext(
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
	observedOnChainOcrSeqNums := make([]uint64, 0, len(aos))
	fChainObservations := make(map[cciptypes.ChainSelector][]int)

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

		if obs.OnChainPriceOcrSeqNum > 0 {
			observedOnChainOcrSeqNums = append(observedOnChainOcrSeqNums, obs.OnChainPriceOcrSeqNum)
		}

		for chain, f := range obs.FChain {
			if f > 0 {
				if _, ok := fChainObservations[chain]; !ok {
					fChainObservations[chain] = make([]int, 0, f)
				}
				fChainObservations[chain] = append(fChainObservations[chain], f)
			}
		}
	}

	if p.discoveryProcessor != nil {
		lggr.Infow("Processing discovery observations", "discoveryObservations", discoveryObservations)

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

	mainOutcome, invalidatePriceCache, err := p.getMainOutcomeAndCacheInvalidation(
		lggr, prevOutcome, observedOnChainOcrSeqNums, fChainObservations)
	if err != nil {
		lggr.Errorw("failed to get main outcome and cache invalidation", "err", err)
	}
	ctx = context.WithValue(ctx, consts.InvalidateCacheKey, invalidatePriceCache)

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
		lggr.Warnw("failed to get chain fee prices outcome", "err", err)
	}

	if len(tokenPriceOutcome.TokenPrices) > 0 || len(chainFeeOutcome.GasPrices) > 0 {
		mainOutcome.InflightPriceOcrSequenceNumber = cciptypes.SeqNum(outCtx.SeqNr)
		mainOutcome.RemainingPriceChecks = p.offchainCfg.InflightPriceCheckRetries
	}

	out := committypes.Outcome{
		MerkleRootOutcome: merkleRootOutcome,
		TokenPriceOutcome: tokenPriceOutcome,
		ChainFeeOutcome:   chainFeeOutcome,
		MainOutcome:       mainOutcome,
	}
	p.metricsReporter.TrackOutcome(out)

	lggr.Infow("Commit plugin finished outcome", "outcome", out)
	return p.ocrTypeCodec.EncodeOutcome(out)
}

func (p *Plugin) getMainOutcomeAndCacheInvalidation(
	lggr logger.Logger,
	prevOutcome committypes.Outcome,
	observedOnChainOcrSeqNums []uint64,
	fChainObservations map[cciptypes.ChainSelector][]int,
) (committypes.MainOutcome, bool, error) {
	mainOutcome := committypes.MainOutcome{}
	invalidatePriceCache := false

	if prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber == 0 {
		return mainOutcome, false, nil
	}

	for _, v := range observedOnChainOcrSeqNums {
		if v == 0 {
			return mainOutcome, false, fmt.Errorf("observed ocr seq num cannot be zero at this point")
		}
	}

	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(p.reportingCfg.F))
	fChainConsensus := consensus.GetConsensusMap(lggr, "mainFChain", fChainObservations, donThresh)
	fDestChain, ok := fChainConsensus[p.destChain]
	if !ok {
		return mainOutcome, false, fmt.Errorf("destChain=%d no f consensus for %v", p.destChain, fChainObservations)
	}

	if len(observedOnChainOcrSeqNums) < 2*fDestChain+1 {
		return mainOutcome, false, fmt.Errorf("onChainOcrSeqNums no consensus required=%d got=%d %v",
			fDestChain, len(observedOnChainOcrSeqNums), observedOnChainOcrSeqNums)
	}

	sort.Slice(observedOnChainOcrSeqNums, func(i, j int) bool {
		return observedOnChainOcrSeqNums[i] < observedOnChainOcrSeqNums[j]
	})
	consensusOnChainOcrSeqNum := observedOnChainOcrSeqNums[fDestChain]

	pricesTransmitted := consensusOnChainOcrSeqNum >= uint64(prevOutcome.MainOutcome.InflightPriceOcrSequenceNumber)
	if pricesTransmitted || prevOutcome.MainOutcome.RemainingPriceChecks == 0 {
		invalidatePriceCache = true
		mainOutcome.InflightPriceOcrSequenceNumber = 0
		mainOutcome.RemainingPriceChecks = 0
	} else if prevOutcome.MainOutcome.RemainingPriceChecks > 0 {
		mainOutcome.RemainingPriceChecks--
	}

	return mainOutcome, invalidatePriceCache, nil
}
