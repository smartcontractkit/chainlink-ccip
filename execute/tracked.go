package execute

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/metrics"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
)

// TrackedPlugin tracks latency of the basic ReportingPlugin's methods. The special ingredient (compared to OCR3)
// metrics is additional tracking of the state machine state. Please see metrics.Reporter for more details and the
// exact prometheus metrics that are exposed.
type TrackedPlugin struct {
	ocr3types.ReportingPlugin[[]byte]
	lggr     logger.Logger
	reporter metrics.Reporter
	codec    ocrtypecodec.ExecCodec
}

func NewTrackedPlugin(
	plugin ocr3types.ReportingPlugin[[]byte],
	lggr logger.Logger,
	reporter metrics.Reporter,
	codec ocrtypecodec.ExecCodec,
) ocr3types.ReportingPlugin[[]byte] {
	return &TrackedPlugin{
		ReportingPlugin: plugin,
		lggr:            lggr,
		reporter:        reporter,
		codec:           codec,
	}
}

func (p *TrackedPlugin) Query(
	ctx context.Context, outctx ocr3types.OutcomeContext,
) (types.Query, error) {
	return withTrackedMethod(
		p,
		plugincommon.QueryMethod,
		outctx,
		func() (types.Query, error) { return p.ReportingPlugin.Query(ctx, outctx) },
	)
}

func (p *TrackedPlugin) Observation(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query,
) (types.Observation, error) {
	return withTrackedMethod(
		p,
		plugincommon.ObservationMethod,
		outctx,
		func() (types.Observation, error) { return p.ReportingPlugin.Observation(ctx, outctx, query) },
	)
}

func (p *TrackedPlugin) Outcome(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	return withTrackedMethod(
		p,
		plugincommon.OutcomeMethod,
		outctx,
		func() (ocr3types.Outcome, error) { return p.ReportingPlugin.Outcome(ctx, outctx, query, aos) },
	)
}

func withTrackedMethod[T any](
	p *TrackedPlugin,
	method plugincommon.MethodType,
	outctx ocr3types.OutcomeContext,
	f func() (T, error),
) (T, error) {
	queryStarted := time.Now()
	resp, err := f()

	latency := time.Since(queryStarted)
	state := currentState(p, outctx)

	p.reporter.TrackLatency(state, method, latency, err)
	p.lggr.Debugw("tracking exec latency",
		"state", state,
		"method", method,
		"latency", latency,
	)

	return resp, err
}

func currentState(p *TrackedPlugin, outctx ocr3types.OutcomeContext) exectypes.PluginState {
	// It's already decoded by the plugin, but we can't easily access it on the interface level.
	// Assuming for now that CPU consumption for this extra decoding is negligible.
	out, err := p.codec.DecodeOutcome(outctx.PreviousOutcome)
	if err != nil {
		p.lggr.Errorw("unable to get state", "error", err)
		return exectypes.Unknown
	}
	return out.State.Next()
}
