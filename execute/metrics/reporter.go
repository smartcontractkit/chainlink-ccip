package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Reporter is a simple interface used for tracking observations and outcomes of the execution plugin.
// Default implementation is based on the Prometheus metrics, but it can be extended to support other metrics systems.
// Main goal is to provide a simple way to track the performance of the execution plugin, for instance:
// - understand how efficiently we batch (number of messages, number of token data, number of source chains used etc.)
// - understand how many messages, reports, token data are observed by plugins
type Reporter interface {
	TrackObservation(obs exectypes.Observation)
	TrackOutcome(outcome exectypes.Outcome)
}

const (
	sourceChainsLabel  = "sourceChains"
	messagesLabel      = "messages"
	tokenDataLabel     = "tokenData"
	commitReportsLabel = "commitReports"
	noncesLabel        = "nonces"
)

var (
	promObservationDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_observation_components",
			Help: "This metric tracks the number of different items in the observation of the execute plugin " +
				"(e.g. number of messages, number of token data etc.)",
		},
		[]string{"chainID", "type"},
	)
	promOutcomeDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_outcome_components",
			Help: "This metric tracks the number of different items in the outcome of the execute plugin " +
				"(e.g. number of messages, number of source chains used etc.)",
		},
		[]string{"chainID", "type"},
	)
	promCostlyMessages = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_costly_messages",
			Help: "This metric tracks the number of costly messages observed by the execute plugin",
		},
		[]string{"chainID"},
	)
)

type PromReporter struct {
	lggr    logger.Logger
	chainID string

	// Prometheus components
	observationDetailsCounter *prometheus.CounterVec
	outcomeDetailsCounter     *prometheus.CounterVec
	costlyMessagesCounter     *prometheus.CounterVec
}

func NewPromReporter(lggr logger.Logger, selector cciptypes.ChainSelector) (*PromReporter, error) {
	chainID, err := sel.GetChainIDFromSelector(uint64(selector))
	if err != nil {
		return nil, err
	}

	return &PromReporter{
		lggr:    lggr,
		chainID: chainID,

		observationDetailsCounter: promObservationDetails,
		outcomeDetailsCounter:     promOutcomeDetails,
		costlyMessagesCounter:     promCostlyMessages,
	}, nil
}

func (p *PromReporter) TrackObservation(obs exectypes.Observation) {
	p.trackCommitReports(obs.CommitReports)
	p.trackMessages(obs.Messages)
	p.trackTokenData(obs.TokenData)
	p.trackNonceData(obs.Nonces)
	p.trackCostlyMessages(obs.CostlyMessages)
}

func (p *PromReporter) TrackOutcome(outcome exectypes.Outcome) {
	p.trackOutcomeComponents(outcome)
}

func (p *PromReporter) trackOutcomeComponents(outcome exectypes.Outcome) {
	messagesCount := 0
	tokenDataCount := 0
	sources := 0

	for chainSelector, report := range outcome.Report.ChainReports {
		sources++
		messagesCount += len(report.Messages)
		tokenDataCount += len(report.OffchainTokenData)

		p.lggr.Debugw("Execute plugin reporting outcome",
			"sourceChainSelector", chainSelector,
			"destChainSelector", p.chainID,
			"messagesCount", len(report.Messages),
			"tokenDataCount", len(report.OffchainTokenData),
		)
	}

	p.outcomeDetailsCounter.
		WithLabelValues(p.chainID, messagesLabel).
		Add(float64(messagesCount))
	p.outcomeDetailsCounter.
		WithLabelValues(p.chainID, tokenDataLabel).
		Add(float64(tokenDataCount))
	p.outcomeDetailsCounter.
		WithLabelValues(p.chainID, sourceChainsLabel).
		Add(float64(sources))
}

func (p *PromReporter) trackCommitReports(commits exectypes.CommitObservations) {
	commitReportsCount := 0
	for _, commit := range commits {
		commitReportsCount += len(commit)
	}

	p.observationDetailsCounter.
		WithLabelValues(p.chainID, commitReportsLabel).
		Add(float64(commitReportsCount))
}

func (p *PromReporter) trackMessages(messages exectypes.MessageObservations) {
	messagesCount := 0
	for _, chainMessages := range messages {
		messagesCount += len(chainMessages)
	}

	p.observationDetailsCounter.
		WithLabelValues(p.chainID, messagesLabel).
		Add(float64(messagesCount))
}

func (p *PromReporter) trackTokenData(tokens exectypes.TokenDataObservations) {
	tokensCount := 0
	for _, chainTokens := range tokens {
		tokensCount += len(chainTokens)
	}

	p.observationDetailsCounter.
		WithLabelValues(p.chainID, tokenDataLabel).
		Add(float64(tokensCount))
}

func (p *PromReporter) trackNonceData(nonces exectypes.NonceObservations) {
	noncesCount := 0
	for _, chainNonces := range nonces {
		noncesCount += len(chainNonces)
	}

	p.observationDetailsCounter.
		WithLabelValues(p.chainID, noncesLabel).
		Add(float64(noncesCount))
}

func (p *PromReporter) trackCostlyMessages(costlyMessages []cciptypes.Bytes32) {
	p.costlyMessagesCounter.
		WithLabelValues(p.chainID).
		Add(float64(len(costlyMessages)))
}

type Noop struct{}

func (n *Noop) TrackObservation(exectypes.Observation) {
}

func (n *Noop) TrackOutcome(exectypes.Outcome) {}
