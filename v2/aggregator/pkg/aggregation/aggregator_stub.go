package aggregation

import "github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"

// AggregatorSinkStub is a stub implementation of the Sink interface for testing purposes.
type AggregatorSinkStub struct {
}

// SubmitReport is a stub implementation of the SubmitReport method for testing purposes.
func (s *AggregatorSinkStub) SubmitReport(_ *model.CommitAggregatedReport) error {
	return nil
}
