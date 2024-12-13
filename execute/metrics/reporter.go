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
	TrackObservation(obs exectypes.Observation, state exectypes.PluginState)
	TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState)
}

const (
	sourceChainsLabel  = "sourceChains"
	messagesLabel      = "messages"
	tokenDataLabel     = "tokenData"
	commitReportsLabel = "commitReports"
	noncesLabel        = "nonces"

	tokenStateReady   = "ready"
	tokenStateWaiting = "waiting"
)

var (
	promObservationDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_observation_components_sizes",
			Help: "This metric tracks the number of different items in the observation of the execute plugin " +
				"(e.g. number of messages, number of token data etc.)",
		},
		[]string{"chainID", "state", "type"},
	)
	promOutcomeDetails = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_outcome_components_sizes",
			Help: "This metric tracks the number of different items in the outcome of the execute plugin " +
				"(e.g. number of messages, number of source chains used etc.)",
		},
		[]string{"chainID", "state", "type"},
	)
	promCostlyMessages = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_costly_messages",
			Help: "This metric tracks the number of costly messages observed by the execute plugin",
		},
		[]string{"chainID", "state"},
	)
	promTokenDataReadiness = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_token_data_readiness",
			Help: "This metric tracks the readiness of the token data observed by the execute plugin",
		},
		[]string{"chainID", "state", "status"},
	)
)

type PromReporter struct {
	lggr    logger.Logger
	chainID string

	// Prometheus reporters
	observationDetailsCounter *prometheus.CounterVec
	outcomeDetailsCounter     *prometheus.CounterVec
	costlyMessagesCounter     *prometheus.CounterVec
	tokenDataReadinessCounter *prometheus.CounterVec
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
		tokenDataReadinessCounter: promTokenDataReadiness,
	}, nil
}

func (p *PromReporter) TrackObservation(obs exectypes.Observation, state exectypes.PluginState) {
	castedState := string(state)
	p.trackCommitReports(obs.CommitReports, castedState)
	p.trackMessages(obs.Messages, castedState)
	p.trackTokenData(obs.TokenData, castedState)
	p.trackNonceData(obs.Nonces, castedState)
	p.trackCostlyMessages(obs.CostlyMessages, castedState)
}

func (p *PromReporter) TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState) {
	p.trackOutcomeComponents(outcome, string(state))
}

func (p *PromReporter) trackOutcomeComponents(outcome exectypes.Outcome, state string) {
	counters := map[string]int{
		messagesLabel:     0,
		tokenDataLabel:    0,
		sourceChainsLabel: 0,
	}

	for chainSelector, report := range outcome.Report.ChainReports {
		counters[sourceChainsLabel]++
		counters[messagesLabel] += len(report.Messages)
		counters[tokenDataLabel] += len(report.OffchainTokenData)

		p.lggr.Debugw("Execute plugin reporting outcome",
			"state", state,
			"sourceChainSelector", chainSelector,
			"destChainSelector", p.chainID,
			"messagesCount", len(report.Messages),
			"tokenDataCount", len(report.OffchainTokenData),
		)
	}

	for key, count := range counters {
		p.outcomeDetailsCounter.
			WithLabelValues(p.chainID, state, key).
			Add(float64(count))
	}
}

func (p *PromReporter) trackCommitReports(commits exectypes.CommitObservations, state string) {
	commitReportsCount := 0
	for _, commit := range commits {
		commitReportsCount += len(commit)
	}

	p.observationDetailsCounter.
		WithLabelValues(p.chainID, state, commitReportsLabel).
		Add(float64(commitReportsCount))
}

func (p *PromReporter) trackMessages(messages exectypes.MessageObservations, state string) {
	messagesCount := 0
	for _, chainMessages := range messages {
		messagesCount += len(chainMessages)
	}

	p.observationDetailsCounter.
		WithLabelValues(p.chainID, state, messagesLabel).
		Add(float64(messagesCount))
}

func (p *PromReporter) trackTokenData(tokens exectypes.TokenDataObservations, state string) {
	tokenCounters := map[string]int{
		tokenStateReady:   0,
		tokenStateWaiting: 0,
	}

	for _, chainTokens := range tokens {
		for _, tokenData := range chainTokens {
			for _, token := range tokenData.TokenData {
				counterKey := tokenStateWaiting
				if token.IsReady() {
					counterKey = tokenStateReady
				}
				tokenCounters[counterKey]++
			}
		}
	}

	for key, count := range tokenCounters {
		p.tokenDataReadinessCounter.
			WithLabelValues(p.chainID, state, key).
			Add(float64(count))
	}
}

func (p *PromReporter) trackNonceData(nonces exectypes.NonceObservations, state string) {
	noncesCount := 0
	for _, chainNonces := range nonces {
		noncesCount += len(chainNonces)
	}

	p.observationDetailsCounter.
		WithLabelValues(p.chainID, state, noncesLabel).
		Add(float64(noncesCount))
}

func (p *PromReporter) trackCostlyMessages(costlyMessages []cciptypes.Bytes32, state string) {
	p.costlyMessagesCounter.
		WithLabelValues(p.chainID, state).
		Add(float64(len(costlyMessages)))
}

type Noop struct{}

func (n *Noop) TrackObservation(exectypes.Observation, exectypes.PluginState) {}

func (n *Noop) TrackOutcome(exectypes.Outcome, exectypes.PluginState) {}

var _ Reporter = &Noop{}
var _ Reporter = &PromReporter{}
