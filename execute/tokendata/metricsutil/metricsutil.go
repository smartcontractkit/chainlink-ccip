// Package metricsutil provides lazily-registered beholder instruments for the tokendata
// subpackages. Instruments must NOT be registered in package-level var initializers:
// beholder.GetClient() returns the noop client until the node installs the real one, so
// registration is deferred to first use via sync.Once.
package metricsutil

import (
	"context"
	"sync"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// LazyCounter is a beholder Int64Counter registered on first use.
type LazyCounter struct {
	name string
	once sync.Once
	inst metric.Int64Counter
}

// LazyGauge is a beholder Int64Gauge registered on first use.
type LazyGauge struct {
	name string
	once sync.Once
	inst metric.Int64Gauge
}

func NewLazyCounter(name string) *LazyCounter {
	return &LazyCounter{name: name}
}

func NewLazyGauge(name string) *LazyGauge {
	return &LazyGauge{name: name}
}

func (c *LazyCounter) init() {
	c.once.Do(func() {
		inst, err := beholder.GetClient().Meter.Int64Counter(c.name)
		if err != nil {
			// Leave the instrument nil; metrics are silently absent rather than
			// misreported if registration fails.
			return
		}
		c.inst = inst
	})
}

func (g *LazyGauge) init() {
	g.once.Do(func() {
		inst, err := beholder.GetClient().Meter.Int64Gauge(g.name)
		if err != nil {
			return
		}
		g.inst = inst
	})
}

// Add increments the counter, or is a no-op if registration failed.
func (c *LazyCounter) Add(ctx context.Context, inc int64, attrs ...attribute.KeyValue) {
	c.init()
	if c.inst == nil {
		return
	}
	c.inst.Add(ctx, inc, metric.WithAttributes(attrs...))
}

// Record sets the gauge value, or is a no-op if registration failed.
func (g *LazyGauge) Record(ctx context.Context, val int64, attrs ...attribute.KeyValue) {
	g.init()
	if g.inst == nil {
		return
	}
	g.inst.Record(ctx, val, metric.WithAttributes(attrs...))
}
