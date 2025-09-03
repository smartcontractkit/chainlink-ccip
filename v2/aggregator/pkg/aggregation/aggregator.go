// Package aggregation provides commit report aggregation functionality for the aggregator service.
package aggregation

import (
	"context"

	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/common"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
)

// CommitReportAggregator is responsible for aggregating commit reports from multiple verifiers.
// It manages the verification and storage of commit reports through a configurable storage backend,
// processes aggregation requests via a message channel, and forwards verified reports to a sink.
type CommitReportAggregator struct {
	storage       common.CommitVerificationStore
	sink          common.Sink
	messageIDChan chan aggregationRequest
	config        model.AggregatorConfig
}

type aggregationRequest struct {
	// CommitteeID is the ID of the committee for the aggregation request.
	CommitteeID string
	MessageID   model.MessageID
}

// CheckAggregation enqueues a new aggregation request for the specified committee and message ID.
func (c *CommitReportAggregator) CheckAggregation(committeeID string, messageID model.MessageID) error {
	c.messageIDChan <- aggregationRequest{
		CommitteeID: committeeID,
		MessageID:   messageID,
	}
	return nil
}

func (c *CommitReportAggregator) checkAggregationAndSubmitComplete(committeeID string, messageID model.MessageID) (*model.CommitAggregatedReport, error) {
	verifications, err := c.storage.ListCommitVerificationByMessageID(context.Background(), committeeID, messageID)
	if err != nil {
		return nil, err
	}

	aggregatedReport := &model.CommitAggregatedReport{
		MessageID:     messageID,
		Verifications: verifications,
	}

	if ok, err := c.checkQuorum(aggregatedReport); err != nil {
		if err := c.sink.SubmitReport(aggregatedReport); err != nil {
			return nil, err
		}
	} else if !ok {
		return nil, nil
	}

	return nil, nil
}

func (c *CommitReportAggregator) checkQuorum(aggregatedReport *model.CommitAggregatedReport) (bool, error) {
	quorumConfig := c.config.Committees[aggregatedReport.CommitteeID].QuorumConfigs[aggregatedReport.GetDestinationSelector()]

	// Check if we have enough signatures to meet the quorum
	// TODO: Make the quorum check configurable, currently doing F + 1 to match contract
	if len(aggregatedReport.Verifications) != int(quorumConfig.F)+1 {
		return false, nil
	}

	return true, nil
}

// StartBackground begins processing aggregation requests in the background.
func (c *CommitReportAggregator) StartBackground(ctx context.Context) {
	go func() {
		select {
		case request := <-c.messageIDChan:
			go func() {
				_, _ = c.checkAggregationAndSubmitComplete(request.CommitteeID, request.MessageID)
			}()
		case <-ctx.Done():
			return
		}
	}()
}

// NewCommitReportAggregator creates a new instance of CommitReportAggregator with the provided dependencies.
// It initializes the aggregator with a storage backend for commit verifications, a sink for forwarding
// completed reports, and configuration settings for committee quorum requirements.
//
// Parameters:
//   - storage: Interface for storing and retrieving commit verifications
//   - sink: Interface for submitting aggregated reports
//   - config: Configuration containing committee and quorum settings
//
// Returns:
//   - *CommitReportAggregator: A new aggregator instance ready to process commit reports
//
// The returned aggregator must have StartBackground called to begin processing aggregation requests.
func NewCommitReportAggregator(storage common.CommitVerificationStore, sink common.Sink, config model.AggregatorConfig) *CommitReportAggregator {
	return &CommitReportAggregator{
		storage:       storage,
		sink:          sink,
		messageIDChan: make(chan aggregationRequest),
		config:        config,
	}
}
