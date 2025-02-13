package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/exp/maps"

	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	promProcessorOutputCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_output_sizes",
			Help: "This metric tracks the number of different items in the commit processor",
		},
		[]string{"chainID", "method", "state", "type"},
	)
	promSequenceNumbers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_exec_max_sequence_number",
			Help: "This metric tracks the max sequence number observed by the commit processor",
		},
		[]string{"chainID", "sourceChainSelector", "method"},
	)
)

type PromReporter struct {
	lggr    logger.Logger
	chainID string

	// Prometheus reporters
	outputDetailsCounter *prometheus.CounterVec
	sequenceNumbers      *prometheus.GaugeVec
}

func NewPromReporter(lggr logger.Logger, selector cciptypes.ChainSelector) (*PromReporter, error) {
	chainID, err := sel.GetChainIDFromSelector(uint64(selector))
	if err != nil {
		return nil, err
	}

	return &PromReporter{
		lggr:    lggr,
		chainID: chainID,

		outputDetailsCounter: promProcessorOutputCounter,
		sequenceNumbers:      promSequenceNumbers,
	}, nil
}

func (p *PromReporter) TrackObservation(obs exectypes.Observation, state exectypes.PluginState) {
	p.trackOutputStats(obs, state, plugincommon.ObservationMethod)

	for sourceChainSelector, cr := range obs.Messages {
		maxSeqNr := pickHighestSeqNr(maps.Keys(cr))
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.ObservationMethod)
	}
}

func (p *PromReporter) TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState) {
	p.trackOutputStats(&outcome, state, plugincommon.OutcomeMethod)

	for _, cr := range outcome.CommitReports {
		sourceChainSelector := cr.SourceChain
		maxSeqNr := pickHighestSeqNrInMessages(cr.Messages)
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.OutcomeMethod)
	}
}

func (p *PromReporter) trackMaxSequenceNumber(
	sourceChainSelector cciptypes.ChainSelector,
	maxSeqNr int,
	method plugincommon.MethodType,
) {
	if maxSeqNr == 0 {
		return
	}

	p.sequenceNumbers.
		WithLabelValues(p.chainID, sourceChainSelector.String(), method).
		Set(float64(maxSeqNr))
}

func (p *PromReporter) trackOutputStats(
	output plugintypes.Trackable,
	state exectypes.PluginState,
	method plugincommon.MethodType,
) {
	stringState := string(state)
	for key, val := range output.Stats() {
		p.outputDetailsCounter.
			WithLabelValues(p.chainID, method, stringState, key).
			Add(float64(val))
	}
}

func pickHighestSeqNrInMessages(messages []cciptypes.Message) int {
	seqNrs := make([]cciptypes.SeqNum, len(messages))
	for i, m := range messages {
		seqNrs[i] = m.Header.SequenceNumber
	}
	return pickHighestSeqNr(seqNrs)
}

func pickHighestSeqNr(seqNrs []cciptypes.SeqNum) int {
	seqNr := cciptypes.SeqNum(0)
	for _, s := range seqNrs {
		if s > seqNr {
			seqNr = s
		}
	}
	return int(seqNr)
}
