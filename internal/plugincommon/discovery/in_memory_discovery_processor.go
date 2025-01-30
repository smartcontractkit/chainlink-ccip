package discovery

import (
	"context"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"sync"
)

// InMemoryDiscoveryProcessor is an in-memory implementation of PluginProcessor.
type InMemoryDiscoveryProcessor struct {
	lggr          logger.Logger
	validateCount int
	validateMutex sync.Mutex
}

// NewInMemoryDiscoveryProcessor creates a new InMemoryDiscoveryProcessor.
func NewInMemoryDiscoveryProcessor(lggr logger.Logger) *InMemoryDiscoveryProcessor {
	return &InMemoryDiscoveryProcessor{
		lggr:          lggr,
		validateMutex: sync.Mutex{},
	}
}

// Query returns an empty query.
func (p *InMemoryDiscoveryProcessor) Query(_ context.Context, _ dt.Outcome) (dt.Query, error) {
	return nil, nil
}

// Observation returns an empty observation.
func (p *InMemoryDiscoveryProcessor) Observation(_ context.Context, _ dt.Outcome, _ dt.Query) (dt.Observation, error) {
	return dt.Observation{}, nil
}

// ValidateObservation returns the count of how many times the function was called.
func (p *InMemoryDiscoveryProcessor) ValidateObservation(_ dt.Outcome, _ dt.Query, _ plugincommon.AttributedObservation[dt.Observation]) error {
	p.validateMutex.Lock()
	defer p.validateMutex.Unlock()
	p.validateCount++
	return nil
}

// Outcome returns an empty outcome.
func (p *InMemoryDiscoveryProcessor) Outcome(_ context.Context, _ dt.Outcome, _ dt.Query, _ []plugincommon.AttributedObservation[dt.Observation]) (dt.Outcome, error) {
	return dt.Outcome{}, nil
}

// ValidateObservationCalled returns true if ValidateObservation was called at least once.
func (p *InMemoryDiscoveryProcessor) ValidateObservationCalled() bool {
	p.validateMutex.Lock()
	defer p.validateMutex.Unlock()
	return p.validateCount > 0
}

// Close does nothing for the in-memory processor.
func (p *InMemoryDiscoveryProcessor) Close() error {
	return nil
}

// Ensure that InMemoryDiscoveryProcessor implements PluginProcessor.
var _ plugincommon.PluginProcessor[dt.Query, dt.Observation, dt.Outcome] = &InMemoryDiscoveryProcessor{}
