// Package model defines the data structures and types used throughout the aggregator service.
package model

// CommitAggregatedReport represents a report of aggregated commit verifications.
type CommitAggregatedReport struct {
	MessageID     MessageID
	CommitteeID   string
	Verifications []*CommitVerificationRecord
}

// GetDestinationSelector retrieves the destination chain selector from the first verification record.
func (c *CommitAggregatedReport) GetDestinationSelector() uint64 {
	return c.Verifications[0].DestChainSelector
}
