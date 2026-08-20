package rcmetrics

import "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

// NoOp implementations of the metrics interfaces. Returned by the New* constructors
// when no beholder meter is configured, so callers never hold a nil value and
// can call the record methods unconditionally. Each receiver is a no-op.

// NoopReaderMetrics is a no-op ReaderMetrics.
type NoopReaderMetrics struct{}

func (NoopReaderMetrics) RecordChainGap(string, ccipocr3.ChainSelector, string) {}
func (NoopReaderMetrics) RecordReadEmpty(string, ccipocr3.ChainSelector)        {}
func (NoopReaderMetrics) RecordMsgDropped(string, string)                       {}

// NoopObservedReaderMetrics is a no-op ObservedReaderMetrics.
type NoopObservedReaderMetrics struct{}

func (NoopObservedReaderMetrics) RecordReadOutcome(ccipocr3.ChainSelector, string, string) {}
func (NoopObservedReaderMetrics) RecordChainFee(ccipocr3.ChainSelector, string, int64)     {}

// NoopConfigPollerMetrics is a no-op ConfigPollerMetrics.
type NoopConfigPollerMetrics struct{}

func (NoopConfigPollerMetrics) RecordCacheAge(ccipocr3.ChainSelector, string, int64)  {}
func (NoopConfigPollerMetrics) RecordPollSuccess(ccipocr3.ChainSelector)              {}
func (NoopConfigPollerMetrics) RecordPollFailure(ccipocr3.ChainSelector, string)      {}
func (NoopConfigPollerMetrics) RecordOverwrittenEmpty(ccipocr3.ChainSelector, string) {}
func (NoopConfigPollerMetrics) RecordLastSuccess(ccipocr3.ChainSelector, int64)       {}

// NoopAccessorMetrics is a no-op AccessorMetrics.
type NoopAccessorMetrics struct{}

func (NoopAccessorMetrics) RecordBatchResult(string, ccipocr3.ChainSelector, string) {}
func (NoopAccessorMetrics) RecordEmptyRead(string, ccipocr3.ChainSelector)           {}
func (NoopAccessorMetrics) RecordRowDrop(string, ccipocr3.ChainSelector, string)     {}
func (NoopAccessorMetrics) RecordFinalityViolated(ccipocr3.ChainSelector)            {}
