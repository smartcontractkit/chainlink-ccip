// Package functionaltests contains functional tests for the aggregator service.
package functionaltests

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pb/aggregator"
	"github.com/stretchr/testify/require"
)

func TestReadWriteCommitVerification(t *testing.T) {
	client, cleanup, err := CreateServerAndClient(t)
	if err != nil {
		t.Fatalf("failed to create server and client: %v", err)
	}
	t.Cleanup(cleanup)

	writeResp, err := client.WriteCommitVerification(t.Context(), &aggregator.WriteCommitVerificationRequest{
		ParticipantId: "participant1",
		CommitteeId:   "committee1",
		CommitVerificationRecord: &aggregator.CommitVerificationRecord{
			MessageId: []byte("message1"),
		},
	})

	require.NoError(t, err, "WriteCommitVerification failed")
	require.Equal(t, aggregator.WriteStatus_SUCCESS, writeResp.Status, "expected WriteStatus_SUCCESS")

	readResp, err := client.ReadCommitVerification(t.Context(), &aggregator.ReadCommitVerificationRequest{
		ParticipantId: "participant1",
		CommitteeId:   "committee1",
		MessageId:     []byte("message1"),
	})

	require.NoError(t, err, "ReadCommitVerification failed")
	require.NotNil(t, readResp, "expected non-nil response")
	require.Equal(t, []byte("message1"), readResp.CommitVerificationRecord.MessageId, "expected MessageId to match")
}
