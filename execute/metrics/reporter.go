package metrics

import (
	"time"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

// Reporter is a simple interface used for tracking observations and outcomes of the execution plugin.
// Default implementation is based on the Prometheus metrics, but it can be extended to support other metrics systems.
// Main goal is to provide a simple way to track the performance of the execution plugin, for instance:
// - understand how efficiently we batch (number of messages, number of token data, number of source chains used etc.)
// - understand how many messages, reports, token data are observed by plugins
type Reporter interface {
	TrackObservation(obs exectypes.Observation, state exectypes.PluginState)
	TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState)
	TrackLatency(state exectypes.PluginState, method plugincommon.MethodType, latency time.Duration, err error)
	TrackProcessorOutput(string, plugincommon.MethodType, plugintypes.Trackable)
	TrackProcessorLatency(processor string, method plugincommon.MethodType, latency time.Duration, err error)
}

type Noop struct{}

func (n *Noop) TrackObservation(exectypes.Observation, exectypes.PluginState) {}

func (n *Noop) TrackOutcome(exectypes.Outcome, exectypes.PluginState) {}

func (n *Noop) TrackLatency(exectypes.PluginState, plugincommon.MethodType, time.Duration, error) {}

func (n *Noop) TrackProcessorOutput(string, plugincommon.MethodType, plugintypes.Trackable) {}

func (n *Noop) TrackProcessorLatency(string, plugincommon.MethodType, time.Duration, error) {}

var _ Reporter = &Noop{}
var _ Reporter = &PromReporter{}
