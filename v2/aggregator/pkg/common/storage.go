// Package common provides shared interfaces
package common

import (
	"context"

	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
)

// CommitVerificationStore defines an interface for storing and retrieving commit verification records.
// SaveCommitVerification persists a commit verification record.
// GetCommitVerification retrieves a commit verification record by its identifier.
// ListCommitVerificationByMessageID retrieves all commit verification records for a specific message ID and committee ID.
type CommitVerificationStore interface {
	SaveCommitVerification(ctx context.Context, record *model.CommitVerificationRecord) error
	GetCommitVerification(ctx context.Context, id model.CommitVerificationRecordIdentifier) (*model.CommitVerificationRecord, error)
	ListCommitVerificationByMessageID(ctx context.Context, committeeID string, messageID model.MessageID) ([]*model.CommitVerificationRecord, error)
}
