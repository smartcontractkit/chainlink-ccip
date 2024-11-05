package commit

import (
	"context"
	"fmt"
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
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type merkleRootObservation = plugincommon.AttributedObservation[merkleroot.Observation]
type tokenPricesObservation = plugincommon.AttributedObservation[tokenprice.Observation]
type chainFeeObservation = plugincommon.AttributedObservation[chainfee.Observation]

type Plugin struct {
	donID               plugintypes.DonID
	oracleID            commontypes.OracleID
	oracleIDToP2PID     map[commontypes.OracleID]libocrtypes.PeerID
	offchainCfg         pluginconfig.CommitOffchainConfig
	ccipReader          readerpkg.CCIPReader
	tokenPricesReader   readerpkg.PriceReader
	reportCodec         cciptypes.CommitPluginCodec
	lggr                logger.Logger
	homeChain           reader.HomeChain
	rmnHomeReader       readerpkg.RMNHome
	reportingCfg        ocr3types.ReportingPluginConfig
	chainSupport        plugincommon.ChainSupport
	merkleRootProcessor plugincommon.PluginProcessor[merkleroot.Query, merkleroot.Observation, merkleroot.Outcome]
	tokenPriceProcessor plugincommon.PluginProcessor[tokenprice.Query, tokenprice.Observation, tokenprice.Outcome]
	chainFeeProcessor   plugincommon.PluginProcessor[chainfee.Query, chainfee.Observation, chainfee.Outcome]
	discoveryProcessor  *discovery.ContractDiscoveryProcessor

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
) *Plugin {
	lggr = logger.Named(lggr, "CommitPlugin")
	lggr = logger.With(lggr, "donID", donID, "oracleID", reportingCfg.OracleID)
	lggr.Infow("creating new plugin instance", "p2pID", oracleIDToP2pID[reportingCfg.OracleID])

	if offchainCfg.MaxMerkleTreeSize == 0 {
		lggr.Warnw("MaxMerkleTreeSize not set, using default value which is for EVM",
			"default", merklemulti.MaxNumberTreeLeaves)
		offchainCfg.MaxMerkleTreeSize = merklemulti.MaxNumberTreeLeaves
	}

	chainSupport := plugincommon.NewCCIPChainSupport(
		logger.Named(lggr, "CCIPChainSupport"),
		homeChain,
		oracleIDToP2pID,
		reportingCfg.OracleID,
		destChain,
	)

	rmnController := rmn.NewController(
		logger.Named(lggr, "RMNController"),
		rmnCrypto,
		offchainCfg.SignObservationPrefix,
		rmnPeerClient,
		rmnHomeReader,
		observationsInitialRequestTimerDuration(reportingCfg.MaxDurationQuery),
		reportsInitialRequestTimerDuration(reportingCfg.MaxDurationQuery),
	)

	merkleRootProcessor := merkleroot.NewProcessor(
		reportingCfg.OracleID,
		oracleIDToP2pID,
		logger.Named(lggr, "MerkleRootProcessor"),
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
	)

	tokenPriceProcessor := tokenprice.NewProcessor(
		reportingCfg.OracleID,
		lggr,
		offchainCfg,
		destChain,
		chainSupport,
		tokenPricesReader,
		homeChain,
		reportingCfg.F,
	)

	discoveryProcessor := discovery.NewContractDiscoveryProcessor(
		lggr,
		&ccipReader,
		homeChain,
		destChain,
		reportingCfg.F,
		oracleIDToP2pID,
	)

	chainFeeProcessr := chainfee.NewProcessor(
		lggr,
		reportingCfg.OracleID,
		destChain,
		homeChain,
		ccipReader,
		offchainCfg,
		chainSupport,
		reportingCfg.F,
	)

	return &Plugin{
		donID:               donID,
		oracleID:            reportingCfg.OracleID,
		oracleIDToP2PID:     oracleIDToP2pID,
		lggr:                lggr,
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
	}
}

func (p *Plugin) Query(ctx context.Context, outCtx ocr3types.OutcomeContext) (types.Query, error) {
	var err error
	var q Query

	prevOutcome := p.decodeOutcome(outCtx.PreviousOutcome)

	q.MerkleRootQuery, err = p.merkleRootProcessor.Query(ctx, prevOutcome.MerkleRootOutcome)
	if err != nil {
		p.lggr.Errorw("get merkle roots query", "err", err)
	}

	q.TokenPriceQuery, err = p.tokenPriceProcessor.Query(ctx, prevOutcome.TokenPriceOutcome)
	if err != nil {
		p.lggr.Errorw("get token prices query", "err", err)
	}

	q.ChainFeeQuery, err = p.chainFeeProcessor.Query(ctx, prevOutcome.ChainFeeOutcome)
	if err != nil {
		p.lggr.Errorw("get chain fee query", "err", err)
	}

	return q.Encode()
}

func (p *Plugin) ObservationQuorum(
	ctx context.Context, _ ocr3types.OutcomeContext, _ types.Query, aos []types.AttributedObservation,
) (bool, error) {
	// Across all chains we require at least 2F+1 observations.
	return quorumhelper.ObservationCountReachesObservationQuorum(
		quorumhelper.QuorumTwoFPlusOne, p.reportingCfg.N, p.reportingCfg.F, aos), nil
}

func (p *Plugin) Observation(
	ctx context.Context, outCtx ocr3types.OutcomeContext, q types.Query,
) (types.Observation, error) {
	prevOutcome := p.decodeOutcome(outCtx.PreviousOutcome)
	fChain := p.ObserveFChain()

	decodedQ, err := DecodeCommitPluginQuery(q)
	if err != nil {
		return nil, fmt.Errorf("decode query: %w", err)
	}

	// If the contracts are not initialized then only submit contracts discovery related observations.
	if !p.contractsInitialized.Load() && p.discoveryProcessor != nil {
		discoveryObs, err := p.discoveryProcessor.Observation(ctx, dt.Outcome{}, dt.Query{})
		if err != nil {
			p.lggr.Errorw("failed to discover contracts", "err", err)
		}

		obs := Observation{DiscoveryObs: discoveryObs}
		encoded, err := obs.Encode()
		if err != nil {
			return nil, fmt.Errorf("encode discovery observation: %w, observation: %+v", err, obs)
		}

		p.lggr.Infow("contracts not initialized, only making discovery observations", "discoveryObs", discoveryObs)
		p.lggr.Debugw("commit plugin making observation", "encodedObservation", encoded, "observation", obs)

		return encoded, nil
	}

	merkleRootObs, err := p.merkleRootProcessor.Observation(ctx, prevOutcome.MerkleRootOutcome, decodedQ.MerkleRootQuery)
	if err != nil {
		p.lggr.Errorw("failed to get merkle observation", "err", err)
	}

	tokenPriceObs, err := p.tokenPriceProcessor.Observation(ctx, prevOutcome.TokenPriceOutcome, decodedQ.TokenPriceQuery)
	if err != nil {
		p.lggr.Errorw("failed to get token prices", "err", err)
	}

	chainFeeObs, err := p.chainFeeProcessor.Observation(ctx, prevOutcome.ChainFeeOutcome, decodedQ.ChainFeeQuery)
	if err != nil {
		p.lggr.Errorw("failed to get gas prices", "err", err)
	}

	obs := Observation{
		MerkleRootObs: merkleRootObs,
		TokenPriceObs: tokenPriceObs,
		ChainFeeObs:   chainFeeObs,
		FChain:        fChain,
	}
	encoded, err := obs.Encode()
	if err != nil {
		return nil, fmt.Errorf("failed to encode observation: %w, observation: %+v", err, obs)
	}

	p.lggr.Debugw("Commit plugin making observation", "encodedObservation", encoded, "observation", obs)
	return encoded, nil
}

func (p *Plugin) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		p.lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

// Outcome depending on the current state, either:
// - chooses the seq num ranges for the next round
// - builds a report
// - checks for the transmission of a previous report
func (p *Plugin) Outcome(
	ctx context.Context, outCtx ocr3types.OutcomeContext, q types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	p.lggr.Debugw("Commit plugin performing outcome",
		"outctx", outCtx,
		"query", q,
		"attributedObservations", aos)

	prevOutcome := p.decodeOutcome(outCtx.PreviousOutcome)

	decodedQ, err := DecodeCommitPluginQuery(q)
	if err != nil {
		return nil, fmt.Errorf("decode query: %w", err)
	}

	var merkleObservations []merkleRootObservation
	var tokensObservations []tokenPricesObservation
	var feeObservations []chainFeeObservation
	var discoveryObservations []plugincommon.AttributedObservation[dt.Observation]

	for _, ao := range aos {
		obs, err := DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			p.lggr.Errorw("failed to decode observation", "err", err)
			continue
		}
		p.lggr.Debugw("Commit plugin outcome decoded observation", "observation", obs)
		merkleObservations = append(merkleObservations,
			merkleRootObservation{
				OracleID:    ao.Observer,
				Observation: obs.MerkleRootObs,
			},
		)

		tokensObservations = append(tokensObservations,
			tokenPricesObservation{
				OracleID:    ao.Observer,
				Observation: obs.TokenPriceObs,
			},
		)

		feeObservations = append(feeObservations,
			chainFeeObservation{
				OracleID:    ao.Observer,
				Observation: obs.ChainFeeObs,
			},
		)

		discoveryObservations = append(discoveryObservations,
			plugincommon.AttributedObservation[dt.Observation]{
				OracleID:    ao.Observer,
				Observation: obs.DiscoveryObs,
			})
	}

	if p.discoveryProcessor != nil {
		p.lggr.Infow("Processing discovery observations", "discoveryObservations", discoveryObservations)
		_, err = p.discoveryProcessor.Outcome(ctx, dt.Outcome{}, dt.Query{}, discoveryObservations)
		if err != nil {
			return nil, fmt.Errorf("unable to process outcome of discovery processor: %w", err)
		}
		p.contractsInitialized.Store(true)
	}

	merkleRootOutcome, err := p.merkleRootProcessor.Outcome(
		ctx,
		prevOutcome.MerkleRootOutcome,
		decodedQ.MerkleRootQuery,
		merkleObservations,
	)
	if err != nil {
		p.lggr.Errorw("failed to get merkle outcome", "err", err)
	}

	tokenPriceOutcome, err := p.tokenPriceProcessor.Outcome(
		ctx,
		prevOutcome.TokenPriceOutcome,
		decodedQ.TokenPriceQuery,
		tokensObservations,
	)
	if err != nil {
		p.lggr.Warnw("failed to get token prices outcome", "err", err)
	}

	chainFeeOutcome, err := p.chainFeeProcessor.Outcome(
		ctx,
		prevOutcome.ChainFeeOutcome,
		decodedQ.ChainFeeQuery,
		feeObservations,
	)
	if err != nil {
		p.lggr.Warnw("failed to get gas prices outcome", "err", err)
	}

	return Outcome{
		MerkleRootOutcome: merkleRootOutcome,
		TokenPriceOutcome: tokenPriceOutcome,
		ChainFeeOutcome:   chainFeeOutcome,
	}.Encode()
}

func (p *Plugin) Close() error {
	return services.CloseAll(
		p.merkleRootProcessor,
		p.tokenPriceProcessor,
		p.chainFeeProcessor,
		p.discoveryProcessor,
	)
}

func (p *Plugin) decodeOutcome(outcome ocr3types.Outcome) Outcome {
	if len(outcome) == 0 {
		return Outcome{}
	}

	decodedOutcome, err := DecodeOutcome(outcome)
	if err != nil {
		p.lggr.Errorw("Failed to decode Outcome", "outcome", outcome, "err", err)
		return Outcome{}
	}

	return decodedOutcome
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
