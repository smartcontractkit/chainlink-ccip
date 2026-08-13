package rcmetrics

// NoOp implementations of the metrics interfaces. Returned by the New* constructors
// when no beholder meter is configured, so callers never hold a nil value and
// can call the record methods unconditionally. Each receiver is a no-op.

// NoopReaderMetrics is a no-op ReaderMetrics.
type NoopReaderMetrics struct{}

func (NoopReaderMetrics) RecordChainGap(string, string, string) {}
func (NoopReaderMetrics) RecordReadEmpty(string, string)        {}
func (NoopReaderMetrics) RecordReadPartial(string, string)      {}
func (NoopReaderMetrics) RecordMsgDropped(string, string)       {}

// NoopObservedReaderMetrics is a no-op ObservedReaderMetrics.
type NoopObservedReaderMetrics struct{}

func (NoopObservedReaderMetrics) RecordReadOutcome(string, string, string)     {}
func (NoopObservedReaderMetrics) RecordChainFee(string, string, string, int64) {}

// NoopConfigPollerMetrics is a no-op ConfigPollerMetrics.
type NoopConfigPollerMetrics struct{}

func (NoopConfigPollerMetrics) RecordCacheAge(string, string, int64) {}
func (NoopConfigPollerMetrics) RecordPollSuccess(string)             {}
func (NoopConfigPollerMetrics) RecordPollFailure(string)             {}
func (NoopConfigPollerMetrics) RecordOverwrittenEmpty(string)        {}
func (NoopConfigPollerMetrics) RecordLastSuccess(string, int64)      {}

// NoopAccessorMetrics is a no-op AccessorMetrics.
type NoopAccessorMetrics struct{}

func (NoopAccessorMetrics) RecordBatchResult(string, string, string) {}
func (NoopAccessorMetrics) RecordEmptyRead(string, string)           {}
func (NoopAccessorMetrics) RecordRowDrop(string, string, string)     {}
func (NoopAccessorMetrics) RecordFinalityViolated(string)            {}
