// Package common provides shared interfaces
package common

import "github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"

// Sink defines an interface for submitting aggregated commit reports.
type Sink interface {
	SubmitReport(report *model.CommitAggregatedReport) error
}
