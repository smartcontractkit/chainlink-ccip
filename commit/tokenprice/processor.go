package tokenprice

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	pkgreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type processor struct {
	oracleID commontypes.OracleID
	// Don't use this logger directly but rather through logutil\.WithContextValues where possible
	lggr             logger.Logger
	offChainCfg      pluginconfig.CommitOffchainConfig
	destChain        cciptypes.ChainSelector
	chainSupport     plugincommon.ChainSupport
	tokenPriceReader pkgreader.PriceReader
	homeChain        reader.HomeChain
	metricsReporter  plugincommon.MetricsReporter
	fRoleDON         int
	obs              observer
}

func NewProcessor(
	oracleID commontypes.OracleID,
	lggr logger.Logger,
	offChainCfg pluginconfig.CommitOffchainConfig,
	destChain cciptypes.ChainSelector,
	chainSupport plugincommon.ChainSupport,
	tokenPriceReader pkgreader.PriceReader,
	homeChain reader.HomeChain,
	fRoleDON int,
	metricsReporter plugincommon.MetricsReporter,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
	var obs observer
	baseObs := newBaseObserver(
		tokenPriceReader,
		destChain,
		oracleID,
		chainSupport,
		offChainCfg,
	)
	if !offChainCfg.TokenPriceAsyncObserverDisabled {
		obs = newAsyncObserver(
			lggr,
			baseObs,
			offChainCfg.TokenPriceAsyncObserverSyncFreq.Duration(),
			offChainCfg.TokenPriceAsyncObserverSyncTimeout.Duration(),
		)
	} else {
		obs = baseObs
	}
	p := &processor{
		oracleID:         oracleID,
		lggr:             lggr,
		offChainCfg:      offChainCfg,
		destChain:        destChain,
		chainSupport:     chainSupport,
		tokenPriceReader: tokenPriceReader,
		homeChain:        homeChain,
		fRoleDON:         fRoleDON,
		metricsReporter:  metricsReporter,
		obs:              obs,
	}
	return plugincommon.NewTrackedProcessor(lggr, p, processorsLabel, metricsReporter)
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

func (p *processor) Outcome(
	ctx context.Context,
	prevOutcome Outcome,
	_ Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	lggr.Infow("processing token price outcome")
	// If set to zero, no prices will be reported (i.e keystone feeds would be active).
	if p.offChainCfg.TokenPriceBatchWriteFrequency.Duration() == 0 {
		lggr.Debugw("TokenPriceBatchWriteFrequency is set to zero, no prices will be reported")
		return newEmptyOutcome(), nil
	}

	inflightTokenPricesOutcome := newInflightPricesOutcome(
		prevOutcome.InflightTokenPriceUpdates, prevOutcome.InflightRemainingChecks-1)

	consensusObservation, err := p.getConsensusObservation(lggr, aos)
	if err != nil {
		return inflightTokenPricesOutcome, fmt.Errorf("get consensus observation: %w", err)
	}

	tokenPriceOutcome := p.selectTokensForUpdate(lggr, consensusObservation)
	lggr.Infow("outcome token prices", "tokenPrices", tokenPriceOutcome)

	if len(tokenPriceOutcome) == 0 {
		lggr.Debugw("No token prices to report")
		return inflightTokenPricesOutcome, nil
	}

	// Check if we have inflight token price updates.
	if prevOutcome.HasInflightTokenPriceUpdates() {
		return p.computeInflightPricesOutcome(lggr, consensusObservation, prevOutcome), nil
	}

	inflightTokenPriceUpdates := make(map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig)
	for tokenAddr := range tokenPriceOutcome {
		oldTokenPriceUpdate, exists := consensusObservation.FeeQuoterTokenUpdates[tokenAddr]
		inflightTokenPriceUpdates[tokenAddr] = cciptypes.NewTimestampedBig(0, time.Time{})
		if exists {
			inflightTokenPriceUpdates[tokenAddr] = oldTokenPriceUpdate
		}
	}

	return newTokenPricesOutcome(
		tokenPriceOutcome,
		inflightTokenPriceUpdates,
		int64(p.offChainCfg.InflightPriceCheckRetries),
	), nil
}

// computeInflightPricesOutcome is called if in this round we wait for prices to appear OnChain.
// If we still wait for some prices it will decrement the number of available retries.
// If all prices appeared OnChain or no retries left it sends an empty outcome so we can transmit fresh prices in the
// next round.
func (p *processor) computeInflightPricesOutcome(
	lggr logger.Logger, consensusObservation ConsensusObservation, prevOutcome Outcome,
) Outcome {
	lggr.Debugw("checking for previously transmitted token price price updates to appear on-chain",
		"prevUpdate", prevOutcome.InflightTokenPriceUpdates,
		"currUpdates", consensusObservation.FeeQuoterTokenUpdates,
		"remRetries", prevOutcome.InflightRemainingChecks,
	)

	for chainSel, inflightUpdate := range prevOutcome.InflightTokenPriceUpdates {
		lggr2 := logger.With(lggr, "chainSel", chainSel, "prevUpdate", inflightUpdate,
			"currUpdates", consensusObservation.FeeQuoterTokenUpdates)

		currUpdate, exists := consensusObservation.FeeQuoterTokenUpdates[chainSel]
		priceAppearedOnChain := exists && currUpdate.Timestamp.After(inflightUpdate.Timestamp)
		if !priceAppearedOnChain {
			lggr2.Infow("waiting for previously transmitted token price updates to appear on-chain")
			return newInflightPricesOutcome(
				prevOutcome.InflightTokenPriceUpdates, prevOutcome.InflightRemainingChecks-1)
		}

		lggr2.Debugw("previously transmitted token price update appeared on-chain")
	}

	// we don't want to transmit the current prices in this round because they might have been recorded onChain
	// in-between Observation and Outcome ocr3 phases, and we might be reporting duplicates. We instead want to send
	// an empty outcome so that in the next round we can properly send new prices.
	lggr.Infow("all inflight token prices appeared OnChain")
	return newEmptyOutcome()
}

func (p *processor) Close() error {
	p.obs.close()
	return nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}
