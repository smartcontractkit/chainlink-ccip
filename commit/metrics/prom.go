package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	// Prometheus labels for merkle processor
	rootsLabel        = "roots"
	messagesLabel     = "messages"
	rmnSignatureLabel = "rmnSignatures"

	// Prometheus labels for token price processor
	tokenPricesLabel           = "tokenPrices"
	feedTokenPricesLabel       = "feedTokenPrices"
	feeQuoterTokenUpdatesLabel = "feeQuoterTokenUpdates"

	// Prometheus labels for chain fee processor
	gasPricesLabel         = "gasPrices"
	feeComponentsLabel     = "feeComponents"
	nativeTokenPricesLabel = "nativeTokenPrices"
	chainFeeUpdatesLabel   = "chainFeeUpdates"
)

var (
	promMerkleProcessorObservationDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_merkle_processor_observation_components_sizes",
			Help: "This metric tracks the number of different items in the merkle observation of the commit plugin",
		},
		[]string{"chainID", "state", "type"},
	)
	promMerkleOutcomeDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_merkle_processor_outcome_components_sizes",
			Help: "This metric tracks the number of different items in the merkle outcome of the commit plugin",
		},
		[]string{"chainID", "state", "type"},
	)
	promTokenPriceObservationDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_token_processor_observation_components_sizes",
			Help: "This metric tracks the number of different items in the token prices observation of the commit plugin",
		},
		[]string{"chainID", "type"},
	)
	promTokenPriceOutcomeDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_token_processor_outcome_components_sizes",
			Help: "This metric tracks the number of different items in the token prices outcome of the commit plugin",
		},
		[]string{"chainID", "type"},
	)
	promChainFeeObservationDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_chain_fee_processor_observation_components_sizes",
			Help: "This metric tracks the number of different items in the chain fee observation of the commit plugin",
		},
		[]string{"chainID", "type"},
	)
	promChainFeeOutcomeDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_chain_fee_processor_outcome_components_sizes",
			Help: "This metric tracks the number of different items in the chain fee outcome of the commit plugin",
		},
		[]string{"chainID", "type"},
	)
)

type PromReporter struct {
	lggr    logger.Logger
	chainID string
	// Prometheus components
	merkleProcessorObservationCounter   *prometheus.CounterVec
	merkleProcessorOutcomeCounter       *prometheus.CounterVec
	tokenProcessorObservationCounter    *prometheus.CounterVec
	tokenProcessorOutcomeCounter        *prometheus.CounterVec
	chainFeeProcessorObservationCounter *prometheus.CounterVec
	chainFeeProcessorOutcomeCounter     *prometheus.CounterVec
}

func NewPromReporter(lggr logger.Logger, selector cciptypes.ChainSelector) (*PromReporter, error) {
	chainID, err := sel.GetChainIDFromSelector(uint64(selector))
	if err != nil {
		return nil, err
	}

	return &PromReporter{
		lggr:    lggr,
		chainID: chainID,

		merkleProcessorObservationCounter:   promMerkleProcessorObservationDetails,
		merkleProcessorOutcomeCounter:       promMerkleOutcomeDetails,
		tokenProcessorObservationCounter:    promTokenPriceObservationDetails,
		tokenProcessorOutcomeCounter:        promTokenPriceOutcomeDetails,
		chainFeeProcessorObservationCounter: promChainFeeObservationDetails,
		chainFeeProcessorOutcomeCounter:     promChainFeeOutcomeDetails,
	}, nil
}

func (p *PromReporter) TrackObservation(obs committypes.Observation) {
}

func (p *PromReporter) TrackOutcome(outcome committypes.Outcome) {
}

func (p *PromReporter) TrackChainFeeObservation(obs chainfee.Observation) {
	counts := chainFeeObservationMetrics(obs)

	for key, count := range counts {
		p.chainFeeProcessorObservationCounter.
			WithLabelValues(p.chainID, key).
			Add(float64(count))
	}
}

func (p *PromReporter) TrackChainFeeOutcome(outcome chainfee.Outcome) {
	counts := chainFeeOutcomeMetrics(outcome)

	for key, count := range counts {
		p.chainFeeProcessorOutcomeCounter.
			WithLabelValues(p.chainID, key).
			Add(float64(count))
	}
}

func (p *PromReporter) TrackMerkleObservation(obs merkleroot.Observation, state string) {
	counts := merkleRootObservationMetrics(obs)

	for key, count := range counts {
		p.merkleProcessorObservationCounter.
			WithLabelValues(p.chainID, state, key).
			Add(float64(count))
	}
}

func (p *PromReporter) TrackMerkleOutcome(outcome merkleroot.Outcome, state string) {
	counts := merkleRootOutcomeMetrics(outcome)

	for key, count := range counts {
		p.merkleProcessorOutcomeCounter.
			WithLabelValues(p.chainID, state, key).
			Add(float64(count))
	}
}

func (p *PromReporter) TrackTokenPricesObservation(obs tokenprice.Observation) {
	counts := tokenPricesObservationMetrics(obs)

	for key, count := range counts {
		p.tokenProcessorObservationCounter.
			WithLabelValues(p.chainID, key).
			Add(float64(count))
	}
}

func (p *PromReporter) TrackTokenPricesOutcome(outcome tokenprice.Outcome) {
	counts := tokenPricesOutcomeMetrics(outcome)

	for key, count := range counts {
		p.tokenProcessorOutcomeCounter.
			WithLabelValues(p.chainID, key).
			Add(float64(count))
	}
}

func merkleRootObservationMetrics(obs merkleroot.Observation) map[string]int {
	counts := map[string]int{
		rootsLabel:    len(obs.MerkleRoots),
		messagesLabel: 0,
	}
	for _, root := range obs.MerkleRoots {
		counts[messagesLabel] += root.SeqNumsRange.Length()
	}
	return counts
}

func merkleRootOutcomeMetrics(outcome merkleroot.Outcome) map[string]int {
	counts := map[string]int{
		rootsLabel:        len(outcome.RootsToReport),
		rmnSignatureLabel: len(outcome.RMNReportSignatures),
		messagesLabel:     0,
	}
	for _, root := range outcome.RootsToReport {
		counts[messagesLabel] += root.SeqNumsRange.Length()
	}
	return counts
}

func tokenPricesObservationMetrics(obs tokenprice.Observation) map[string]int {
	return map[string]int{
		feedTokenPricesLabel:       len(obs.FeedTokenPrices),
		feeQuoterTokenUpdatesLabel: len(obs.FeeQuoterTokenUpdates),
	}
}

func tokenPricesOutcomeMetrics(outcome tokenprice.Outcome) map[string]int {
	return map[string]int{
		tokenPricesLabel: len(outcome.TokenPrices),
	}
}

func chainFeeObservationMetrics(obs chainfee.Observation) map[string]int {
	return map[string]int{
		feeComponentsLabel:     len(obs.FeeComponents),
		nativeTokenPricesLabel: len(obs.NativeTokenPrices),
		chainFeeUpdatesLabel:   len(obs.ChainFeeUpdates),
	}
}

func chainFeeOutcomeMetrics(outcome chainfee.Outcome) map[string]int {
	return map[string]int{
		gasPricesLabel: len(outcome.GasPrices),
	}
}
