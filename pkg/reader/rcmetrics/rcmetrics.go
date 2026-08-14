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
	"strconv"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ReaderMetrics instruments the CcipReader's concrete methods (per-chain reads).
type ReaderMetrics interface {
	// RecordChainGap records, per read (query), how many source chains ended up in
	// each state once the requested set is reconciled against what was returned.
	RecordChainGap(query, chain, state string)
	// RecordReadEmpty records a query that returned nothing with no error.
	RecordReadEmpty(query, chain string)
	// RecordMsgDropped records a read record dropped on validation/cast.
	RecordMsgDropped(query, reason string)
}

// readerMetrics implements ReaderMetrics against a beholder meter.
type readerMetrics struct {
	chainGap   metric.Int64Counter
	readEmpty  metric.Int64Counter
	msgDropped metric.Int64Counter
}

// NewReaderMetrics registers the reader instruments against m. Returns a NoOp
// implementation (and no error) when m is nil.
func NewReaderMetrics(m metric.Meter) (ReaderMetrics, error) {
	if m == nil {
		return NoopReaderMetrics{}, nil
	}
	rm := &readerMetrics{}
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

func (r *readerMetrics) RecordChainGap(query, chain, state string) {
	r.chainGap.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("query", query),
		attribute.String("chain", chain),
		attribute.String("state", state),
	))
}

func (r *readerMetrics) RecordReadEmpty(query, chain string) {
	r.readEmpty.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("query", query),
		attribute.String("chain", chain),
	))
}

func (r *readerMetrics) RecordMsgDropped(query, reason string) {
	r.msgDropped.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("query", query),
		attribute.String("reason", reason),
	))
}

// ObservedReaderMetrics instruments the observedCCIPReader (per-query outcomes).
type ObservedReaderMetrics interface {
	// RecordReadOutcome records a reader query classified as ok/empty/error.
	RecordReadOutcome(chainID, query, outcome string)
	// RecordChainFee records an observed chain fee component (exec/da) value.
	RecordChainFee(chainFamily, chainID, feeType string, value int64)
}

// observedReaderMetrics implements ObservedReaderMetrics against a meter.
type observedReaderMetrics struct {
	readOutcome metric.Int64Counter
	bhChainFee  metric.Int64Gauge
}

// NewObservedReaderMetrics registers the observed-reader instruments against m.
// Returns a NoOp implementation (and no error) when m is nil.
func NewObservedReaderMetrics(m metric.Meter) (ObservedReaderMetrics, error) {
	if m == nil {
		return NoopObservedReaderMetrics{}, nil
	}
	om := &observedReaderMetrics{}
	var err error
	if om.readOutcome, err = m.Int64Counter("ccip_reader_read_outcome"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_read_outcome: %w", err)
	}
	if om.bhChainFee, err = m.Int64Gauge("ccip_reader_chain_fee_components"); err != nil {
		return nil, fmt.Errorf("register ccip_reader_chain_fee_components: %w", err)
	}
	return om, nil
}

func (o *observedReaderMetrics) RecordReadOutcome(chainID, query, outcome string) {
	o.readOutcome.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chainID", chainID),
		attribute.String("query", query),
		attribute.String("outcome", outcome),
	))
}

func (o *observedReaderMetrics) RecordChainFee(chainFamily, chainID, feeType string, value int64) {
	o.bhChainFee.Record(context.Background(), value, metric.WithAttributes(
		attribute.String("chainFamily", chainFamily),
		attribute.String("chainID", chainID),
		attribute.String("feeType", feeType),
	))
}

// ConfigPollerMetrics instruments the config poller background cache.
type ConfigPollerMetrics interface {
	// RecordCacheAge gauges how stale the served cache is, per chain/kind.
	RecordCacheAge(chain, kind string, ageSeconds int64)
	// RecordPollSuccess / RecordPollFailure count background refresh outcomes per chain.
	RecordPollSuccess(chain string)
	RecordPollFailure(chain string)
	// RecordOverwrittenEmpty counts refreshes that time-stamped empty data as fresh.
	RecordOverwrittenEmpty(kind string)
	// RecordLastSuccess records the last successful refresh (epoch seconds) per chain.
	RecordLastSuccess(chain string, epochSeconds int64)
}

// configPollerMetrics implements ConfigPollerMetrics against a meter.
type configPollerMetrics struct {
	cacheAgeSeconds      metric.Int64Gauge
	pollSuccess          metric.Int64Counter
	pollFailure          metric.Int64Counter
	overwrittenEmpty     metric.Int64Counter
	lastSuccessTimestamp metric.Int64Gauge
}

// NewConfigPollerMetrics registers the config poller instruments against m.
// Returns a NoOp implementation (and no error) when m is nil.
func NewConfigPollerMetrics(m metric.Meter) (ConfigPollerMetrics, error) {
	if m == nil {
		return NoopConfigPollerMetrics{}, nil
	}
	cm := &configPollerMetrics{}
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

func (c *configPollerMetrics) RecordCacheAge(chain, kind string, ageSeconds int64) {
	c.cacheAgeSeconds.Record(context.Background(), ageSeconds, metric.WithAttributes(
		attribute.String("chain", chain),
		attribute.String("kind", kind),
	))
}

func (c *configPollerMetrics) RecordPollSuccess(chain string) {
	c.pollSuccess.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chain", chain),
	))
}

func (c *configPollerMetrics) RecordPollFailure(chain string) {
	c.pollFailure.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chain", chain),
	))
}

func (c *configPollerMetrics) RecordOverwrittenEmpty(kind string) {
	c.overwrittenEmpty.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("kind", kind),
	))
}

func (c *configPollerMetrics) RecordLastSuccess(chain string, epochSeconds int64) {
	c.lastSuccessTimestamp.Record(context.Background(), epochSeconds, metric.WithAttributes(
		attribute.String("chain", chain),
	))
}

// AccessorMetrics instruments the chain accessor wrapper (per-operation reads).
// Retained for the interface-based pattern even though the accessor wrapper is
// not yet wired; the NoOp fallback keeps callers nil-free when it lands.
type AccessorMetrics interface {
	RecordBatchResult(operation, chain, outcome string)
	RecordEmptyRead(operation, chain string)
	RecordRowDrop(operation, chain, reason string)
	RecordFinalityViolated(chain string)
}

// accessorMetrics implements AccessorMetrics against a meter.
type accessorMetrics struct {
	batchResult      metric.Int64Counter
	emptyRead        metric.Int64Counter
	rowDrops         metric.Int64Counter
	finalityViolated metric.Int64Counter
}

// NewAccessorMetrics registers the accessor instruments against m.
// Returns a NoOp implementation (and no error) when m is nil.
func NewAccessorMetrics(m metric.Meter) (AccessorMetrics, error) {
	if m == nil {
		return NoopAccessorMetrics{}, nil
	}
	am := &accessorMetrics{}
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

func (a *accessorMetrics) RecordBatchResult(operation, chain, outcome string) {
	a.batchResult.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("operation", operation),
		attribute.String("chain", chain),
		attribute.String("outcome", outcome),
	))
}

func (a *accessorMetrics) RecordEmptyRead(operation, chain string) {
	a.emptyRead.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("operation", operation),
		attribute.String("chain", chain),
	))
}

func (a *accessorMetrics) RecordRowDrop(operation, chain, reason string) {
	a.rowDrops.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("operation", operation),
		attribute.String("chain", chain),
		attribute.String("reason", reason),
	))
}

func (a *accessorMetrics) RecordFinalityViolated(chain string) {
	a.finalityViolated.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chain", chain),
	))
}

// ChainLabel resolves a chain selector to a human-readable label for metric
// cardinality, falling back to the numeric selector when unknown.
func ChainLabel(sel ccipocr3.ChainSelector) string {
	return strconv.FormatUint(uint64(sel), 10)
}
