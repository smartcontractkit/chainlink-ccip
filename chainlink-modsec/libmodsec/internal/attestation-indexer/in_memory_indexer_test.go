package attestation_indexer

import (
	"context"
	"testing"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryAttestationIndexer(t *testing.T) {
	type testCase struct {
		name  string
		setup func(*InMemoryAttestationIndexer)
		test  func(*testing.T, *InMemoryAttestationIndexer)
	}

	// Test data
	verifierID1 := [32]byte{7, 8, 9}
	verifierID2 := [32]byte{10, 11, 12}
	verifierID3 := [32]byte{13, 14, 15}
	msgID1 := [32]byte{1}
	msgID2 := [32]byte{2}
	proof1_1, proof1_2 := []byte("proof_V1_Msg1"), []byte("proof_V1_Msg2")
	proof2_1, proof2_2 := []byte("proof_V2_Msg1"), []byte("proof_V2_Msg2")
	proof3_1, proof3_2 := []byte("proof_V3_Msg1"), []byte("proof_V3_Msg2")

	msg1 := modsectypes.Message{
		Header: modsectypes.Header{
			MessageID: msgID1,
		},
	}
	msg2 := modsectypes.Message{
		Header: modsectypes.Header{
			MessageID: msgID2,
		},
	}

	testCases := []testCase{
		{
			name:  "new indexer is empty",
			setup: func(indexer *InMemoryAttestationIndexer) {},
			test: func(t *testing.T, indexer *InMemoryAttestationIndexer) {
				attestations, err := indexer.GetAttestationsForMessage(context.Background(), msg1)
				require.NoError(t, err)
				assert.Nil(t, attestations)

				hasQuorum := indexer.MessageHasQuorum(context.Background(), msg1, 1)
				assert.False(t, hasQuorum)
			},
		},
		{
			name: "add single attestation",
			setup: func(indexer *InMemoryAttestationIndexer) {
				err := indexer.AddAttestation(msgID1, verifierID1, proof1_1)
				require.NoError(t, err)
			},
			test: func(t *testing.T, indexer *InMemoryAttestationIndexer) {
				attestations, err := indexer.GetAttestationsForMessage(context.Background(), msg1)
				require.NoError(t, err)
				require.Len(t, attestations, 1)
				assert.Equal(t, proof1_1, attestations[0].Proof)
				assert.Equal(t, verifierID1, attestations[0].VerifierId)

				hasQuorum := indexer.MessageHasQuorum(context.Background(), msg1, 1)
				assert.True(t, hasQuorum)
			},
		},
		{
			name: "add multiple attestations for same message",
			setup: func(indexer *InMemoryAttestationIndexer) {
				err := indexer.AddAttestation(msgID1, verifierID1, proof1_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID1, verifierID2, proof2_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID1, verifierID3, proof3_1)
				require.NoError(t, err)
			},
			test: func(t *testing.T, indexer *InMemoryAttestationIndexer) {
				attestations, err := indexer.GetAttestationsForMessage(context.Background(), msg1)
				require.NoError(t, err)
				require.Len(t, attestations, 3)

				// Check all attestations are present (order may vary)
				proofs := make(map[string]bool)
				verifiers := make(map[[32]byte]bool)
				for _, att := range attestations {
					proofs[string(att.Proof)] = true
					verifiers[att.VerifierId] = true
				}
				assert.True(t, proofs[string(proof1_1)])
				assert.True(t, proofs[string(proof2_1)])
				assert.True(t, proofs[string(proof3_1)])
				assert.True(t, verifiers[verifierID1])
				assert.True(t, verifiers[verifierID2])
				assert.True(t, verifiers[verifierID3])

				hasQuorum := indexer.MessageHasQuorum(context.Background(), msg1, 3)
				assert.True(t, hasQuorum)
				hasQuorum = indexer.MessageHasQuorum(context.Background(), msg1, 4)
				assert.False(t, hasQuorum)
			},
		},
		{
			name: "overwrite attestation for same verifier should error",
			setup: func(indexer *InMemoryAttestationIndexer) {
				err := indexer.AddAttestation(msgID1, verifierID1, proof1_1)
				require.NoError(t, err)

			},
			test: func(t *testing.T, indexer *InMemoryAttestationIndexer) {
				err := indexer.AddAttestation(msgID1, verifierID1, proof1_2)
				require.Error(t, err)

				attestations, err := indexer.GetAttestationsForMessage(context.Background(), msg1)
				require.NoError(t, err)
				require.Len(t, attestations, 1)
				assert.Equal(t, proof1_1, attestations[0].Proof)
				assert.Equal(t, verifierID1, attestations[0].VerifierId)
			},
		},
		{
			name: "multiple messages with multiple attestations",
			setup: func(indexer *InMemoryAttestationIndexer) {
				err := indexer.AddAttestation(msgID1, verifierID1, proof1_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID1, verifierID2, proof2_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID1, verifierID3, proof3_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID2, verifierID2, proof2_2)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID2, verifierID3, proof3_2)
				require.NoError(t, err)
			},
			test: func(t *testing.T, indexer *InMemoryAttestationIndexer) {
				attestations1, err := indexer.GetAttestationsForMessage(context.Background(), msg1)
				require.NoError(t, err)
				require.Len(t, attestations1, 3)

				attestations2, err := indexer.GetAttestationsForMessage(context.Background(), msg2)
				require.NoError(t, err)
				require.Len(t, attestations2, 2)

				hasQuorum1 := indexer.MessageHasQuorum(context.Background(), msg1, 3)
				assert.True(t, hasQuorum1)
				hasQuorum2 := indexer.MessageHasQuorum(context.Background(), msg2, 2)
				assert.True(t, hasQuorum2)
			},
		},
		{
			name: "quorum check with different requirements",
			setup: func(indexer *InMemoryAttestationIndexer) {
				err := indexer.AddAttestation(msgID1, verifierID1, proof1_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID1, verifierID2, proof2_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID1, verifierID3, proof3_1)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID2, verifierID2, proof2_2)
				require.NoError(t, err)
				err = indexer.AddAttestation(msgID2, verifierID3, proof3_2)
				require.NoError(t, err)
			},
			test: func(t *testing.T, indexer *InMemoryAttestationIndexer) {
				assert.True(t, indexer.MessageHasQuorum(context.Background(), msg1, 2))
				assert.True(t, indexer.MessageHasQuorum(context.Background(), msg1, 3))
				assert.False(t, indexer.MessageHasQuorum(context.Background(), msg1, 4))
				assert.True(t, indexer.MessageHasQuorum(context.Background(), msg2, 1))
				assert.True(t, indexer.MessageHasQuorum(context.Background(), msg2, 2))
				assert.False(t, indexer.MessageHasQuorum(context.Background(), msg2, 3))
			},
		},
		{
			name: "quorum check short circuits",
			setup: func(indexer *InMemoryAttestationIndexer) {
				// Add many attestations to test short-circuiting
				for i := 0; i < 10; i++ {
					verifierID := [32]byte{byte(i)}
					proof := []byte{byte(i)}
					err := indexer.AddAttestation(msgID1, verifierID, proof)
					require.NoError(t, err)
				}
			},
			test: func(t *testing.T, indexer *InMemoryAttestationIndexer) {
				// Should return true quickly when required count is met
				assert.True(t, indexer.MessageHasQuorum(context.Background(), msg1, 5))
				assert.True(t, indexer.MessageHasQuorum(context.Background(), msg1, 10))
				assert.False(t, indexer.MessageHasQuorum(context.Background(), msg1, 11))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			indexer := NewInMemoryAttestationIndexer()
			tc.setup(indexer)
			tc.test(t, indexer)
		})
	}
}

func TestNewInMemoryAttestationIndexer(t *testing.T) {
	indexer := NewInMemoryAttestationIndexer()
	assert.NotNil(t, indexer)

	// Verify it's properly initialized
	msgID := [32]byte{1, 2, 3}
	msg := modsectypes.Message{
		Header: modsectypes.Header{
			MessageID: msgID,
		},
	}

	attestations, err := indexer.GetAttestationsForMessage(context.Background(), msg)
	require.NoError(t, err)
	assert.Nil(t, attestations)
}
