// Package rcmetrics holds the beholder instruments for the CCIP reader's
// data-source layer (the CcipReader, the chain accessor wrapper, and the
// config poller cache). This layer is shared by both the commit and execute
// plugins, so every instrument is plugin-generic (no ccip_commit_* prefix).
//
// Each metrics set is an interface with a NoOp implementation (noop.go) so
// callers never hold a nil value: the New* constructors return NoOp when no
// beholder meter is available (e.g. in tests). This keeps call sites free of
// nil-guards and makes the receivers easy to mock in future tests.
//
// The constructors return an error if an instrument fails to register; callers
// must propagate it rather than ignore it (a registration failure would silently
// drop the metric otherwise).
package rcmetrics

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/metricsutil"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ReaderMetrics instruments the CcipReader's concrete methods (per-chain reads).
type ReaderMetrics interface {
	// RecordChainGap records, per read (query), how many source chains ended up in
	// each state once the requested set is reconciled against what was returned.
	RecordChainGap(query string, chainSelector ccipocr3.ChainSelector, state string)
	// RecordReadEmpty records a query that returned nothing with no error.
	RecordReadEmpty(query string, chainSelector ccipocr3.ChainSelector)
	// RecordMsgDropped records a read record dropped on validation/cast.
	RecordMsgDropped(query, reason string)
}

// readerMetrics implements ReaderMetrics against a beholder meter.
type readerMetrics struct {
	chainGap   metric.Int64Counter
	readEmpty  metric.Int64Counter
	msgDropped metric.Int64Counter
	destAttrs  []attribute.KeyValue
}

// NewReaderMetrics registers the reader instruments against m, scoped to destChainSelector (the
// destination chain the owning ccipChainReader instance serves -- see destChainAttrs). Returns a
// NoOp implementation (and no error) when m is nil.
func NewReaderMetrics(m metric.Meter, destChainSelector ccipocr3.ChainSelector) (ReaderMetrics, error) {
	if m == nil {
		return NoopReaderMetrics{}, nil
	}
	rm := &readerMetrics{destAttrs: metricsutil.DestChainAttrs(destChainSelector)}
	var err error
	if rm.chainGap, err = m.Int64Counter("ccip_reader_chain_gap"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_chain_gap: %w", err)
	}
	if rm.readEmpty, err = m.Int64Counter("ccip_reader_read_empty"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_read_empty: %w", err)
	}
	if rm.msgDropped, err = m.Int64Counter("ccip_reader_msg_dropped"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_msg_dropped: %w", err)
	}
	return rm, nil
}

func (r *readerMetrics) RecordChainGap(query string, chainSelector ccipocr3.ChainSelector, state string) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), r.destAttrs...)
	attrs = append(attrs,
		attribute.String("query", query),
		attribute.String("state", state),
	)
	r.chainGap.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (r *readerMetrics) RecordReadEmpty(query string, chainSelector ccipocr3.ChainSelector) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), r.destAttrs...)
	attrs = append(attrs, attribute.String("query", query))
	r.readEmpty.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (r *readerMetrics) RecordMsgDropped(query, reason string) {
	// Copy destAttrs (shared across every call) before appending -- don't grow it in place.
	attrs := append(append([]attribute.KeyValue{}, r.destAttrs...),
		attribute.String("query", query),
		attribute.String("reason", reason),
	)
	r.msgDropped.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// ObservedReaderMetrics instruments the observedCCIPReader (per-query outcomes).
type ObservedReaderMetrics interface {
	// RecordReadOutcome records a reader query classified as ok/empty/error.
	RecordReadOutcome(chainSelector ccipocr3.ChainSelector, query, outcome string)
	// RecordChainFee records an observed chain fee component (exec/da) value.
	RecordChainFee(chainSelector ccipocr3.ChainSelector, feeType string, value int64)
}

// observedReaderMetrics implements ObservedReaderMetrics against a meter.
type observedReaderMetrics struct {
	readOutcome metric.Int64Counter
	bhChainFee  metric.Int64Gauge
	destAttrs   []attribute.KeyValue
}

// NewObservedReaderMetrics registers the observed-reader instruments against m, scoped to
// destChainSelector (see destChainAttrs). Returns a NoOp implementation (and no error) when m is
// nil.
func NewObservedReaderMetrics(
	m metric.Meter, destChainSelector ccipocr3.ChainSelector,
) (ObservedReaderMetrics, error) {
	if m == nil {
		return NoopObservedReaderMetrics{}, nil
	}
	om := &observedReaderMetrics{destAttrs: metricsutil.DestChainAttrs(destChainSelector)}
	var err error
	if om.readOutcome, err = m.Int64Counter("ccip_reader_read_outcome"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_read_outcome: %w", err)
	}
	if om.bhChainFee, err = m.Int64Gauge("ccip_reader_chain_fee_components"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_chain_fee_components: %w", err)
	}
	return om, nil
}

func (o *observedReaderMetrics) RecordReadOutcome(chainSelector ccipocr3.ChainSelector, query, outcome string) {
	// NOTE: for this metric metricsutil.ChainAttrs(chainSelector) and o.destAttrs are always the same
	// destination chain (recordReadOutcome is always called with the reader's own destChain) --
	// see the "one semantic exception" note in docs/metrics/reader-metrics.md. Both are still
	// attached for schema consistency with every other metric in this package.
	attrs := append(metricsutil.ChainAttrs(chainSelector), o.destAttrs...)
	attrs = append(attrs,
		attribute.String("query", query),
		attribute.String("outcome", outcome),
	)
	o.readOutcome.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (o *observedReaderMetrics) RecordChainFee(chainSelector ccipocr3.ChainSelector, feeType string, value int64) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), o.destAttrs...)
	attrs = append(attrs, attribute.String("feeType", feeType))
	o.bhChainFee.Record(context.Background(), value, metric.WithAttributes(attrs...))
}

// ConfigPollerMetrics instruments the config poller background cache.
type ConfigPollerMetrics interface {
	// RecordCacheAge gauges how stale the served cache is, per chain/kind.
	RecordCacheAge(chainSelector ccipocr3.ChainSelector, kind string, ageSeconds int64)
	// RecordPollSuccess / RecordPollFailure count background refresh outcomes per chain.
	RecordPollSuccess(chainSelector ccipocr3.ChainSelector)
	RecordPollFailure(chainSelector ccipocr3.ChainSelector)
	// RecordOverwrittenEmpty counts refreshes that time-stamped empty data as fresh, per chain/kind.
	RecordOverwrittenEmpty(chainSelector ccipocr3.ChainSelector, kind string)
	// RecordLastSuccess records the last successful refresh (epoch seconds) per chain.
	RecordLastSuccess(chainSelector ccipocr3.ChainSelector, epochSeconds int64)
}

// configPollerMetrics implements ConfigPollerMetrics against a meter.
type configPollerMetrics struct {
	cacheAgeSeconds      metric.Int64Gauge
	pollSuccess          metric.Int64Counter
	pollFailure          metric.Int64Counter
	overwrittenEmpty     metric.Int64Counter
	lastSuccessTimestamp metric.Int64Gauge
	destAttrs            []attribute.KeyValue
}

// NewConfigPollerMetrics registers the config poller instruments against m, scoped to
// destChainSelector -- the destination chain the owning configPollerV2 instance serves. A node runs
// one configPollerV2 per destination chain, and several of those instances can independently poll
// the SAME source chain (e.g. two dest-chain lanes both listing Solana as a known source chain);
// without destChainSelector here their series would collide into one per (chain, node_id),
// indistinguishable from each other -- see destChainAttrs's doc comment. Returns a NoOp
// implementation (and no error) when m is nil.
func NewConfigPollerMetrics(m metric.Meter, destChainSelector ccipocr3.ChainSelector) (ConfigPollerMetrics, error) {
	if m == nil {
		return NoopConfigPollerMetrics{}, nil
	}
	cm := &configPollerMetrics{destAttrs: metricsutil.DestChainAttrs(destChainSelector)}
	var err error
	if cm.cacheAgeSeconds, err = m.Int64Gauge("ccip_reader_config_cache_age_seconds"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_config_cache_age_seconds: %w", err)
	}
	if cm.pollSuccess, err = m.Int64Counter("ccip_reader_config_poll_success"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_config_poll_success: %w", err)
	}
	if cm.pollFailure, err = m.Int64Counter("ccip_reader_config_poll_failure"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_config_poll_failure: %w", err)
	}
	if cm.overwrittenEmpty, err = m.Int64Counter("ccip_reader_config_cache_overwritten_empty"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_config_cache_overwritten_empty: %w", err)
	}
	if cm.lastSuccessTimestamp, err = m.Int64Gauge("ccip_reader_config_poller_last_success_timestamp"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_config_poller_last_success_timestamp: %w", err)
	}
	return cm, nil
}

func (c *configPollerMetrics) RecordCacheAge(chainSelector ccipocr3.ChainSelector, kind string, ageSeconds int64) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), c.destAttrs...)
	attrs = append(attrs, attribute.String("kind", kind))
	c.cacheAgeSeconds.Record(context.Background(), ageSeconds, metric.WithAttributes(attrs...))
}

func (c *configPollerMetrics) RecordPollSuccess(chainSelector ccipocr3.ChainSelector) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), c.destAttrs...)
	c.pollSuccess.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (c *configPollerMetrics) RecordPollFailure(chainSelector ccipocr3.ChainSelector) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), c.destAttrs...)
	c.pollFailure.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (c *configPollerMetrics) RecordOverwrittenEmpty(chainSelector ccipocr3.ChainSelector, kind string) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), c.destAttrs...)
	attrs = append(attrs, attribute.String("kind", kind))
	c.overwrittenEmpty.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (c *configPollerMetrics) RecordLastSuccess(chainSelector ccipocr3.ChainSelector, epochSeconds int64) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), c.destAttrs...)
	c.lastSuccessTimestamp.Record(context.Background(), epochSeconds, metric.WithAttributes(attrs...))
}

// AccessorMetrics instruments the chain accessor wrapper (per-operation reads).
// Retained for the interface-based pattern even though the accessor wrapper is
// not yet wired; the NoOp fallback keeps callers nil-free when it lands.
type AccessorMetrics interface {
	RecordBatchResult(operation string, chainSelector ccipocr3.ChainSelector, outcome string)
	RecordEmptyRead(operation string, chainSelector ccipocr3.ChainSelector)
	RecordRowDrop(operation string, chainSelector ccipocr3.ChainSelector, reason string)
	RecordFinalityViolated(chainSelector ccipocr3.ChainSelector)
}

// accessorMetrics implements AccessorMetrics against a meter.
type accessorMetrics struct {
	batchResult      metric.Int64Counter
	emptyRead        metric.Int64Counter
	rowDrops         metric.Int64Counter
	finalityViolated metric.Int64Counter
	destAttrs        []attribute.KeyValue
}

// NewAccessorMetrics registers the accessor instruments against m, scoped to destChainSelector (see
// destChainAttrs). Returns a NoOp implementation (and no error) when m is nil.
func NewAccessorMetrics(m metric.Meter, destChainSelector ccipocr3.ChainSelector) (AccessorMetrics, error) {
	if m == nil {
		return NoopAccessorMetrics{}, nil
	}
	am := &accessorMetrics{destAttrs: metricsutil.DestChainAttrs(destChainSelector)}
	var err error
	if am.batchResult, err = m.Int64Counter("ccip_reader_batch_result"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_batch_result: %w", err)
	}
	if am.emptyRead, err = m.Int64Counter("ccip_reader_accessor_empty_read"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_accessor_empty_read: %w", err)
	}
	if am.rowDrops, err = m.Int64Counter("ccip_reader_accessor_row_drops"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_accessor_row_drops: %w", err)
	}
	if am.finalityViolated, err = m.Int64Counter("ccip_reader_accessor_finality_violated"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_accessor_finality_violated: %w", err)
	}
	return am, nil
}

func (a *accessorMetrics) RecordBatchResult(operation string, chainSelector ccipocr3.ChainSelector, outcome string) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), a.destAttrs...)
	attrs = append(attrs,
		attribute.String("operation", operation),
		attribute.String("outcome", outcome),
	)
	a.batchResult.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (a *accessorMetrics) RecordEmptyRead(operation string, chainSelector ccipocr3.ChainSelector) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), a.destAttrs...)
	attrs = append(attrs, attribute.String("operation", operation))
	a.emptyRead.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (a *accessorMetrics) RecordRowDrop(operation string, chainSelector ccipocr3.ChainSelector, reason string) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), a.destAttrs...)
	attrs = append(attrs,
		attribute.String("operation", operation),
		attribute.String("reason", reason),
	)
	a.rowDrops.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (a *accessorMetrics) RecordFinalityViolated(chainSelector ccipocr3.ChainSelector) {
	attrs := append(metricsutil.ChainAttrs(chainSelector), a.destAttrs...)
	a.finalityViolated.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}
