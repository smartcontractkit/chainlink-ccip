package commit

import (
	"context"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/smartcontractkit/libocr/quorumhelper"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type attributedMerkleRootObservation = plugincommon.AttributedObservation[merkleroot.Observation]
type attributedTokenPricesObservation = plugincommon.AttributedObservation[tokenprice.Observation]
type attributedChainFeeObservation = plugincommon.AttributedObservation[chainfee.Observation]

type Plugin struct {
	donID             plugintypes.DonID
	oracleID          commontypes.OracleID
	oracleIDToP2PID   map[commontypes.OracleID]libocrtypes.PeerID
	offchainCfg       pluginconfig.CommitOffchainConfig
	ccipReader        readerpkg.CCIPReader
	tokenPricesReader readerpkg.PriceReader
	reportCodec       cciptypes.CommitPluginCodec
	reportBuilder     builder.ReportBuilderFunc
	// Don't use this logger directly but rather through logutil\.WithContextValues where possible
	lggr                logger.Logger
	homeChain           reader.HomeChain
	rmnHomeReader       readerpkg.RMNHome
	reportingCfg        ocr3types.ReportingPluginConfig
	chainSupport        plugincommon.ChainSupport
	merkleRootProcessor plugincommon.PluginProcessor[merkleroot.Query, merkleroot.Observation, merkleroot.Outcome]
	tokenPriceProcessor plugincommon.PluginProcessor[tokenprice.Query, tokenprice.Observation, tokenprice.Outcome]
	chainFeeProcessor   plugincommon.PluginProcessor[chainfee.Query, chainfee.Observation, chainfee.Outcome]
	discoveryProcessor  plugincommon.PluginProcessor[dt.Query, dt.Observation, dt.Outcome]
	metricsReporter     metrics.CommitPluginReporter
	ocrTypeCodec        ocrtypecodec.CommitCodec

	// state
	contractsInitialized atomic.Bool
}

func NewPlugin(
	donID plugintypes.DonID,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	offchainCfg pluginconfig.CommitOffchainConfig,
	destChain cciptypes.ChainSelector,
	ccipReader readerpkg.CCIPReader,
	tokenPricesReader readerpkg.PriceReader,
	reportCodec cciptypes.CommitPluginCodec,
	msgHasher cciptypes.MessageHasher,
	lggr logger.Logger,
	homeChain reader.HomeChain,
	rmnHomeReader readerpkg.RMNHome,
	rmnCrypto cciptypes.RMNCrypto,
	rmnPeerClient rmn.PeerClient,
	reportingCfg ocr3types.ReportingPluginConfig,
	reporter metrics.Reporter,
	addressCodec cciptypes.AddressCodec,
	reportBuilder builder.ReportBuilderFunc,
) *Plugin {
	lggr.Infow("creating new plugin instance", "p2pID", oracleIDToP2pID[reportingCfg.OracleID])

	if offchainCfg.MaxMerkleTreeSize == 0 {
		lggr.Warnw("MaxMerkleTreeSize not set, using default value which is for EVM",
			"default", merklemulti.MaxNumberTreeLeaves)
		offchainCfg.MaxMerkleTreeSize = merklemulti.MaxNumberTreeLeaves
	}

	chainSupport := plugincommon.NewChainSupport(
		logutil.WithComponent(lggr, "ChainSupport"),
		homeChain,
		oracleIDToP2pID,
		reportingCfg.OracleID,
		destChain,
	)

	rmnController := rmn.NewController(
		logutil.WithComponent(lggr, "RMNController"),
		rmnCrypto,
		offchainCfg.SignObservationPrefix,
		rmnPeerClient,
		rmnHomeReader,
		observationsInitialRequestTimerDuration(reportingCfg.MaxDurationQuery),
		reportsInitialRequestTimerDuration(reportingCfg.MaxDurationQuery),
		reporter,
	)

	merkleRootProcessor := merkleroot.NewProcessor(
		reportingCfg.OracleID,
		oracleIDToP2pID,
		logutil.WithComponent(lggr, "MerkleRoot"),
		offchainCfg,
		destChain,
		homeChain,
		ccipReader,
		msgHasher,
		reportingCfg,
		chainSupport,
		rmnController,
		rmnCrypto,
		rmnHomeReader,
		reporter,
		addressCodec,
	)

	tokenPriceProcessor := tokenprice.NewProcessor(
		reportingCfg.OracleID,
		logutil.WithComponent(lggr, "TokenPrice"),
		offchainCfg,
		destChain,
		chainSupport,
		tokenPricesReader,
		homeChain,
		reportingCfg.F,
		reporter,
	)

	discoveryProcessor := discovery.NewContractDiscoveryProcessor(
		logutil.WithComponent(lggr, "Discovery"),
		&ccipReader,
		homeChain,
		destChain,
		reportingCfg.F,
		oracleIDToP2pID,
		reporter,
	)

	chainFeeProcessr := chainfee.NewProcessor(
		logutil.WithComponent(lggr, "ChainFee"),
		reportingCfg.OracleID,
		destChain,
		homeChain,
		ccipReader,
		offchainCfg,
		chainSupport,
		reportingCfg.F,
		reporter,
	)

	return &Plugin{
		donID:               donID,
		oracleID:            reportingCfg.OracleID,
		oracleIDToP2PID:     oracleIDToP2pID,
		lggr:                logutil.WithComponent(lggr, "CommitPlugin"),
		offchainCfg:         offchainCfg,
		tokenPricesReader:   tokenPricesReader,
		ccipReader:          ccipReader,
		homeChain:           homeChain,
		rmnHomeReader:       rmnHomeReader,
		reportCodec:         reportCodec,
		reportingCfg:        reportingCfg,
		chainSupport:        chainSupport,
		merkleRootProcessor: merkleRootProcessor,
		tokenPriceProcessor: tokenPriceProcessor,
		chainFeeProcessor:   chainFeeProcessr,
		discoveryProcessor:  discoveryProcessor,
		metricsReporter:     reporter,
		ocrTypeCodec:        ocrtypecodec.DefaultCommitCodec,
		reportBuilder:       reportBuilder,
	}
}

// Query returns the query for the next round.
// NOTE: In most cases the Query phase should not return an error based on outCtx to prevent infinite retries.
func (p *Plugin) Query(ctx context.Context, outCtx ocr3types.OutcomeContext) (types.Query, error) {
	// Ensure that sequence number is in the context for consumption by all
	// downstream processors and the ccip reader.
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, outCtx.SeqNr, logutil.PhaseQuery)

	var err error
	var q committypes.Query

	prevOutcome, err := p.ocrTypeCodec.DecodeOutcome(outCtx.PreviousOutcome)
	if err != nil {
		return nil, fmt.Errorf("decode previous outcome: %w", err)
	}

	q.MerkleRootQuery, err = p.merkleRootProcessor.Query(ctx, prevOutcome.MerkleRootOutcome)
	if err != nil {
		lggr.Errorw("get merkle roots query", "err", err)
	}

	q.TokenPriceQuery, err = p.tokenPriceProcessor.Query(ctx, prevOutcome.TokenPriceOutcome)
	if err != nil {
		lggr.Errorw("get token prices query", "err", err)
	}

	q.ChainFeeQuery, err = p.chainFeeProcessor.Query(ctx, prevOutcome.ChainFeeOutcome)
	if err != nil {
		lggr.Errorw("get chain fee query", "err", err)
	}

	return p.ocrTypeCodec.EncodeQuery(q)
}

func (p *Plugin) ObservationQuorum(
	ctx context.Context, _ ocr3types.OutcomeContext, _ types.Query, aos []types.AttributedObservation,
) (bool, error) {
	// Across all chains we require at least 2F+1 observations.
	return quorumhelper.ObservationCountReachesObservationQuorum(
		quorumhelper.QuorumTwoFPlusOne,
		p.reportingCfg.N,
		p.reportingCfg.F,
		aos,
	), nil
}

// Observation returns the observation for this round.
// NOTE: In most cases the Observation phase should not return an error based on outCtx to prevent infinite retries.
func (p *Plugin) Observation(
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

	tokenPriceObs, chainFeeObs := p.getPriceRelatedObservations(ctx, lggr, prevOutcome, decodedQ)

	obs := committypes.Observation{
		MerkleRootObs: merkleRootObs,
		TokenPriceObs: tokenPriceObs,
		DiscoveryObs:  discoveryObs,
		ChainFeeObs:   chainFeeObs,
		FChain:        p.ObserveFChain(lggr),
	}

	p.metricsReporter.TrackObservation(obs)

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

func (p *Plugin) ObserveFChain(lggr logger.Logger) map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

func (p *Plugin) Outcome(
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
	p.metricsReporter.TrackOutcome(out)

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

func (p *Plugin) Close() error {
	p.lggr.Infow("closing commit plugin")

	closeable := []io.Closer{
		p.merkleRootProcessor,
		p.tokenPriceProcessor,
		p.chainFeeProcessor,
		p.discoveryProcessor,
		p.ccipReader,
	}

	// Chains without RMN don't initialize the RMNHomeReader
	// TODO Consider initializing rmnHomeReader anyway but using some noop implementation
	if p.rmnHomeReader != nil {
		closeable = append(closeable, p.rmnHomeReader)
	}

	return services.CloseAll(closeable...)
}

// Assuming that we have to delegate a specific amount of time to the observation requests and the report requests.
// We define some percentages in order to help us calculate the time we have to delegate to each request timer.
const (
	observationDurationPercentage = 0.55
	reportDurationPercentage      = 0.4
	// remaining 5% for other query processing

	maxAllowedObservationTimeout = 3 * time.Second
	maxAllowedReportTimeout      = 2 * time.Second
)

func observationsInitialRequestTimerDuration(maxQueryDuration time.Duration) time.Duration {
	// we have queryCapacityForObservations to make the initial observation request and potentially a secondary request
	queryCapacityForObservations := time.Duration(observationDurationPercentage * float64(maxQueryDuration))

	// we divide in two parts one for the initial observation and one for the retry
	queryCapacityForInitialObservations := queryCapacityForObservations / 2

	// if the capacity is greater than the maximum allowed we return the max allowed
	if queryCapacityForInitialObservations < maxAllowedObservationTimeout {
		return queryCapacityForObservations
	}

	return maxAllowedObservationTimeout
}

func reportsInitialRequestTimerDuration(maxQueryDuration time.Duration) time.Duration {
	// we have queryCapacityForReports to make the initial reports request and potentially a secondary request
	queryCapacityForReports := time.Duration(reportDurationPercentage * float64(maxQueryDuration))

	// we divide in two parts one for the initial signatures request and one for the retry
	queryCapacityForInitialObservations := queryCapacityForReports / 2

	// if the capacity is greater than the maximum allowed we return the max allowed
	if queryCapacityForInitialObservations < maxAllowedReportTimeout {
		return queryCapacityForInitialObservations
	}

	return maxAllowedReportTimeout
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
