package commit

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/shared"
)

type MerkleRootObservation = shared.AttributedObservation[merkleroot.Observation]
type TokenPricesObservation = shared.AttributedObservation[tokenprice.Observation]
type ChainFeeObservation = shared.AttributedObservation[chainfee.Observation]

type Plugin struct {
	nodeID              commontypes.OracleID
	oracleIDToP2pID     map[commontypes.OracleID]libocrtypes.PeerID
	cfg                 pluginconfig.CommitPluginConfig
	ccipReader          readerpkg.CCIPReader
	readerSyncer        *plugincommon.BackgroundReaderSyncer
	tokenPricesReader   reader.PriceReader
	reportCodec         cciptypes.CommitPluginCodec
	lggr                logger.Logger
	homeChain           reader.HomeChain
	reportingCfg        ocr3types.ReportingPluginConfig
	chainSupport        plugincommon.ChainSupport
	merkleRootProcessor shared.PluginProcessor[merkleroot.Query, merkleroot.Observation, merkleroot.Outcome]
	tokenPriceProcessor shared.PluginProcessor[tokenprice.Query, tokenprice.Observation, tokenprice.Outcome]
	chainFeeProcessor   shared.PluginProcessor[chainfee.Query, chainfee.Observation, chainfee.Outcome]
}

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	cfg pluginconfig.CommitPluginConfig,
	ccipReader readerpkg.CCIPReader,
	tokenPricesReader reader.PriceReader,
	reportCodec cciptypes.CommitPluginCodec,
	msgHasher cciptypes.MessageHasher,
	lggr logger.Logger,
	homeChain reader.HomeChain,
	reportingCfg ocr3types.ReportingPluginConfig,
) *Plugin {
	readerSyncer := plugincommon.NewBackgroundReaderSyncer(
		lggr,
		ccipReader,
		syncTimeout(cfg.SyncTimeout),
		syncFrequency(cfg.SyncFrequency),
	)
	if err := readerSyncer.Start(context.Background()); err != nil {
		lggr.Errorw("error starting background reader syncer", "err", err)
	}

	chainSupport := plugincommon.NewCCIPChainSupport(
		lggr,
		homeChain,
		oracleIDToP2pID,
		nodeID,
		cfg.DestChain,
	)

	merkleRootProcessor := merkleroot.NewProcessor(
		nodeID,
		lggr,
		cfg,
		homeChain,
		ccipReader,
		msgHasher,
		reportingCfg,
		chainSupport,
	)
	tokenPriceProcessor := tokenprice.NewProcessor(
		nodeID,
		lggr,
		cfg,
		chainSupport,
		tokenPricesReader,
		homeChain,
	)

	return &Plugin{
		nodeID:              nodeID,
		oracleIDToP2pID:     oracleIDToP2pID,
		lggr:                lggr,
		cfg:                 cfg,
		tokenPricesReader:   tokenPricesReader,
		ccipReader:          ccipReader,
		homeChain:           homeChain,
		readerSyncer:        readerSyncer,
		reportCodec:         reportCodec,
		reportingCfg:        reportingCfg,
		chainSupport:        chainSupport,
		merkleRootProcessor: merkleRootProcessor,
		tokenPriceProcessor: tokenPriceProcessor,
		chainFeeProcessor:   chainfee.NewProcessor(),
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

func (p *Plugin) ObservationQuorum(_ ocr3types.OutcomeContext, _ types.Query) (ocr3types.Quorum, error) {
	// Across all chains we require at least 2F+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
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
	return obs.Encode()
}

func (p *Plugin) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		p.lggr.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

// Outcome depending on the current state, either:
// - chooses the seq num ranges for the next round
// - builds a report
// - checks for the transmission of a previous report
func (p *Plugin) Outcome(
	outCtx ocr3types.OutcomeContext, q types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	prevOutcome := p.decodeOutcome(outCtx.PreviousOutcome)

	decodedQ, err := DecodeCommitPluginQuery(q)
	if err != nil {
		return nil, fmt.Errorf("decode query: %w", err)
	}

	var merkleObservations []MerkleRootObservation
	var tokensObservations []TokenPricesObservation
	var feeObservations []ChainFeeObservation

	for _, ao := range aos {
		obs, err := DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			p.lggr.Errorw("failed to decode observation", "err", err)
			continue
		}
		merkleObservations = append(merkleObservations,
			MerkleRootObservation{
				OracleID:    ao.Observer,
				Observation: obs.MerkleRootObs,
			},
		)

		tokensObservations = append(tokensObservations,
			TokenPricesObservation{
				OracleID:    ao.Observer,
				Observation: obs.TokenPriceObs,
			},
		)

		feeObservations = append(feeObservations,
			ChainFeeObservation{
				OracleID:    ao.Observer,
				Observation: obs.ChainFeeObs,
			},
		)
	}

	merkleRootOutcome, err := p.merkleRootProcessor.Outcome(
		prevOutcome.MerkleRootOutcome,
		decodedQ.MerkleRootQuery,
		merkleObservations,
	)
	if err != nil {
		p.lggr.Errorw("failed to get merkle outcome", "err", err)
	}

	tokenPriceOutcome, err := p.tokenPriceProcessor.Outcome(
		prevOutcome.TokenPriceOutcome,
		decodedQ.TokenPriceQuery,
		tokensObservations,
	)

	if err != nil {
		p.lggr.Errorw("failed to get token prices outcome", "err", err)
	}

	chainFeeOutcome, err := p.chainFeeProcessor.Outcome(
		prevOutcome.ChainFeeOutcome,
		decodedQ.ChainFeeQuery,
		feeObservations,
	)
	if err != nil {
		p.lggr.Errorw("failed to get gas prices outcome", "err", err)
	}

	return Outcome{
		MerkleRootOutcome: merkleRootOutcome,
		TokenPriceOutcome: tokenPriceOutcome,
		ChainFeeOutcome:   chainFeeOutcome,
	}.Encode()
}

func (p *Plugin) Close() error {
	timeout := 10 * time.Second
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	defer cf()

	if err := p.readerSyncer.Close(); err != nil {
		p.lggr.Errorw("error closing reader syncer", "err", err)
	}

	if err := p.ccipReader.Close(ctx); err != nil {
		return fmt.Errorf("close ccip reader: %w", err)
	}

	return nil
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

func syncFrequency(configuredValue time.Duration) time.Duration {
	if configuredValue.Milliseconds() == 0 {
		return 10 * time.Second
	}
	return configuredValue
}

func syncTimeout(configuredValue time.Duration) time.Duration {
	if configuredValue.Milliseconds() == 0 {
		return 3 * time.Second
	}
	return configuredValue
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
