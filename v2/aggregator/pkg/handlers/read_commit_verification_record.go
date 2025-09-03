// Package handlers provides HTTP and gRPC request handlers for the aggregator service.
package handlers

import (
	"context"

	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pb/aggregator"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/common"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
)

// ReadCommitVerificationRecordHandler handles requests to read commit verification records.
type ReadCommitVerificationRecordHandler struct {
	storage common.CommitVerificationStore
}

// Handle processes the read request and retrieves the corresponding commit verification record.
func (h *ReadCommitVerificationRecordHandler) Handle(ctx context.Context, req *aggregator.ReadCommitVerificationRequest) (*aggregator.ReadCommitVerificationResponse, error) {
	id := model.CommitVerificationRecordIdentifier{
		ParticipantID: req.GetParticipantId(),
		CommitteeID:   req.GetCommitteeId(),
		MessageID:     req.GetMessageId(),
	}

	record, err := h.storage.GetCommitVerification(ctx, id)
	if err != nil {
		return nil, err
	}

	return &aggregator.ReadCommitVerificationResponse{
		CommitVerificationRecord: &record.CommitVerificationRecord,
	}, nil
}

// NewReadCommitVerificationRecordHandler creates a new instance of ReadCommitVerificationRecordHandler.
func NewReadCommitVerificationRecordHandler(store common.CommitVerificationStore) *ReadCommitVerificationRecordHandler {
	return &ReadCommitVerificationRecordHandler{
		storage: store,
	}
}
