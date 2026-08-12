package tokenprice

import (
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

// MetricsReporter exposes the metrics methods used by the tokenprice processor.
// It embeds the generic processor metrics and adds tokenprice-specific reporters so
// future metrics can be added here without widening the shared plugincommon interface.
type MetricsReporter interface {
	plugincommon.MetricsReporter

	// TrackConsensusDropped reports a key that was dropped during shared consensus
	// aggregation (e.g. fChain, FeedTokenPrices, FeeQuoterUpdates). The reason
	// distinguishes "threshold not defined", "insufficient agreement", and "split".
	TrackConsensusDropped(objectName string, key string, reason string)
}

// NoopMetrics is a no-op MetricsReporter for tests.
type NoopMetrics struct{}

func (n NoopMetrics) TrackProcessorLatency(string, string, time.Duration, error) {}

func (n NoopMetrics) TrackProcessorOutput(string, plugincommon.MethodType, plugintypes.Trackable) {}

func (n NoopMetrics) TrackConsensusDropped(string, string, string) {}
