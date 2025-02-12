package metrics

import (
	"time"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

// Reporter is a simple interface used for tracking observations and outcomes of the commit plugin.
// Default implementation is based on the Prometheus metrics, but it can be extended to support other metrics systems.
// It allows you to track observation/outcome on the processor level as well as on the individual plugin level.
// That gives us more flexibility and granularity in tracking the performance of the commit plugin.
// Processors have a dedicated sub-interfaces covering only the relevant methods for reporting, please see:
// - chainfee.MetricsReporter
// - merkleroot.MetricsReporter
// - tokenprice.MetricsReporter
// - CommitPluginReporter
// This split is required to define the reporting logic in one place but inject only relevant dependencies to
// plugins/processors. Also, it solves the problem of cyclic dependencies between the plugins/processors.
type Reporter interface {
	TrackObservation(obs committypes.Observation)
	TrackOutcome(outcome committypes.Outcome)

	TrackRmnReport(latency float64, success bool)
	TrackRmnRequest(method string, latency float64, nodeID uint64, err string)

	TrackProcessorLatency(processor string, method string, latency time.Duration, err error)
	TrackProcessorObservation(processor string, obs plugintypes.Trackable)
	TrackProcessorOutcome(processor string, out plugintypes.Trackable)
}

type CommitPluginReporter interface {
	TrackObservation(obs committypes.Observation)
	TrackOutcome(outcome committypes.Outcome)
}

type Noop struct{}

func (n *Noop) TrackObservation(committypes.Observation) {}

func (n *Noop) TrackOutcome(committypes.Outcome) {}

func (n *Noop) TrackMerkleObservation(merkleroot.Observation, string) {}

func (n *Noop) TrackMerkleOutcome(merkleroot.Outcome, string) {}

func (n *Noop) TrackRmnReport(float64, bool) {}

func (n *Noop) TrackRmnRequest(string, float64, uint64, string) {}

func (n *Noop) TrackTokenPricesObservation(tokenprice.Observation) {}

func (n *Noop) TrackTokenPricesOutcome(tokenprice.Outcome) {}

func (n *Noop) TrackProcessorLatency(string, string, time.Duration, error) {}

func (n *Noop) TrackProcessorObservation(string, plugintypes.Trackable) {}

func (n *Noop) TrackProcessorOutcome(string, plugintypes.Trackable) {}

var _ Reporter = &PromReporter{}
var _ CommitPluginReporter = &PromReporter{}
var _ merkleroot.MetricsReporter = &PromReporter{}
