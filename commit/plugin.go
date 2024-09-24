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
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type MerkleRootObservation = plugincommon.AttributedObservation[merkleroot.Observation]
type TokenPricesObservation = plugincommon.AttributedObservation[tokenprice.Observation]
type ChainFeeObservation = plugincommon.AttributedObservation[chainfee.Observation]

type Plugin struct {
	donID               plugintypes.DonID
	nodeID              commontypes.OracleID
	oracleIDToP2pID     map[commontypes.OracleID]libocrtypes.PeerID
	cfg                 pluginconfig.CommitPluginConfig
	ccipReader          readerpkg.CCIPReader
	tokenPricesReader   reader.PriceReader
	reportCodec         cciptypes.CommitPluginCodec
	lggr                logger.Logger
	homeChain           reader.HomeChain
	reportingCfg        ocr3types.ReportingPluginConfig
	chainSupport        plugincommon.ChainSupport
	merkleRootProcessor plugincommon.PluginProcessor[merkleroot.Query, merkleroot.Observation, merkleroot.Outcome]
	tokenPriceProcessor plugincommon.PluginProcessor[tokenprice.Query, tokenprice.Observation, tokenprice.Outcome]
	chainFeeProcessor   plugincommon.PluginProcessor[chainfee.Query, chainfee.Observation, chainfee.Outcome]
	discoveryProcessor  *discovery.ContractDiscoveryProcessor
	rmnConfig           rmn.Config

	// state
	contractsInitialized bool
}

func NewPlugin(
	_ context.Context,
	donID plugintypes.DonID,
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
	rmnConfig rmn.Config,
) *Plugin {
	if cfg.MaxMerkleTreeSize == 0 {
		lggr.Warnw("MaxMerkleTreeSize not set, using default value which is for EVM",
			"default", pluginconfig.EvmDefaultMaxMerkleTreeSize)
		cfg.MaxMerkleTreeSize = pluginconfig.EvmDefaultMaxMerkleTreeSize
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
		rmn.Controller(nil),      // todo
		cciptypes.RMNCrypto(nil), // todo
		rmnConfig,
	)

	tokenPriceProcessor := tokenprice.NewProcessor(
		nodeID,
		lggr,
		cfg,
		chainSupport,
		tokenPricesReader,
		homeChain,
		reportingCfg.F,
	)

	discoveryProcessor := discovery.NewContractDiscoveryProcessor(
		lggr,
		&ccipReader,
		homeChain,
		cfg.DestChain,
		reportingCfg.F,
	)

	return &Plugin{
		donID:               donID,
		nodeID:              nodeID,
		oracleIDToP2pID:     oracleIDToP2pID,
		lggr:                lggr,
		cfg:                 cfg,
		tokenPricesReader:   tokenPricesReader,
		ccipReader:          ccipReader,
		homeChain:           homeChain,
		reportCodec:         reportCodec,
		reportingCfg:        reportingCfg,
		chainSupport:        chainSupport,
		merkleRootProcessor: merkleRootProcessor,
		tokenPriceProcessor: tokenPriceProcessor,
		chainFeeProcessor:   chainfee.NewProcessor(),
		discoveryProcessor:  discoveryProcessor,
		rmnConfig:           rmnConfig,
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

	var discoveryObs dt.Observation
	if p.discoveryProcessor != nil {
		discoveryObs, err = p.discoveryProcessor.Observation(ctx, dt.Outcome{}, dt.Query{})
		if err != nil {
			p.lggr.Errorw("failed to discover contracts", "err", err)
		}
		if !p.contractsInitialized {
			p.lggr.Infow("contracts not initialized, only making discovery observations")
			return Observation{DiscoveryObs: discoveryObs}.Encode()
		}
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
		DiscoveryObs:  discoveryObs,
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
	var discoveryObservations []plugincommon.AttributedObservation[dt.Observation]

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

		discoveryObservations = append(discoveryObservations,
			plugincommon.AttributedObservation[dt.Observation]{
				OracleID:    ao.Observer,
				Observation: obs.DiscoveryObs,
			})
	}

	if p.discoveryProcessor != nil {
		_, err = p.discoveryProcessor.Outcome(dt.Outcome{}, dt.Query{}, discoveryObservations)
		if err != nil {
			return nil, fmt.Errorf("unable to process outcome of discovery processor: %w", err)
		}
		p.contractsInitialized = true
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
		p.lggr.Warnw("failed to get token prices outcome", "err", err)
	}

	chainFeeOutcome, err := p.chainFeeProcessor.Outcome(
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
	timeout := 10 * time.Second
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	defer cf()

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

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
