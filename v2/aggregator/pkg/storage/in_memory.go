// Package storage provides storage implementations for the aggregator service.
package storage

import (
	"bytes"
	"context"
	"errors"

	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
)

// InMemoryStorage is an in-memory implementation of the CommitVerificationStore interface.
type InMemoryStorage struct {
	records map[string]*model.CommitVerificationRecord
}

// SaveCommitVerification persists a commit verification record.
func (s *InMemoryStorage) SaveCommitVerification(_ context.Context, record *model.CommitVerificationRecord) error {
	id := record.GetID()
	s.records[id.ToIdentifier()] = record
	return nil
}

// GetCommitVerification retrieves a commit verification record by its identifier.
func (s *InMemoryStorage) GetCommitVerification(_ context.Context, id model.CommitVerificationRecordIdentifier) (*model.CommitVerificationRecord, error) {
	record, exists := s.records[id.ToIdentifier()]
	if !exists {
		return nil, errors.New("record not found")
	}
	return record, nil
}

// ListCommitVerificationByMessageID retrieves all commit verification records for a specific message ID and committee ID.
func (s *InMemoryStorage) ListCommitVerificationByMessageID(_ context.Context, committeeID string, messageID model.MessageID) ([]*model.CommitVerificationRecord, error) {
	var results []*model.CommitVerificationRecord
	for _, record := range s.records {
		if record.CommitteeID == committeeID && bytes.Equal(record.MessageId, messageID) {
			results = append(results, record)
		}
	}
	return results, nil
}

// NewInMemoryStorage creates a new instance of InMemoryStorage.
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		records: make(map[string]*model.CommitVerificationRecord),
	}
}
