package attestation_indexer

import (
	"context"
	"fmt"

	"golang.org/x/sync/syncmap"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/internal/executor"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

// InMemoryAttestationIndexer is an in-memory implementation of an attestation indexer
// It stores all proofs in an in-memory map, indexed by messageID([32]byte) and then by verifierID([32]byte)
// This implementation is useful when assuming that attestations will be requested via messageID
type InMemoryAttestationIndexer struct {
	attestations syncmap.Map // concurrency-safe map of messageID to map of verifierID to attestation(proof)
}

// NewInMemoryAttestationIndexer creates a new AttestationIndexer instance
func NewInMemoryAttestationIndexer() *InMemoryAttestationIndexer {
	// Create a new InMemoryAttestationIndexer with an initialized attestations map
	a := &InMemoryAttestationIndexer{
		attestations: syncmap.Map{},
	}
	return a
}

// GetAttestationsForMessage retrieves attestations for a given message and converts it into the executor.Attestation type
func (at *InMemoryAttestationIndexer) GetAttestationsForMessage(_ context.Context, msg modsectypes.Message) ([]executor.Attestation, error) {
	// Retrieve the map of attestations for the given messageID
	verifierProofs, ok := at.attestations.Load(msg.Header.MessageID)
	if !ok {
		return nil, nil // No attestations found for this messageID
	}
	attestations := make([]executor.Attestation, 0)
	verifierProofs.(*syncmap.Map).Range(func(key, value interface{}) bool {
		attestations = append(attestations, executor.Attestation{
			Proof:      value.([]byte),
			VerifierId: key.([32]byte),
		})
		return true
	})

	return attestations, nil
}

// MessageHasQuorum checks if the message has a quorum of attestations
// Although this is less performant in an in-memory indexer, it will be much more performant in production with a real indexer
// TODO: do we even need this? Not currently used
func (at *InMemoryAttestationIndexer) MessageHasQuorum(_ context.Context, msg modsectypes.Message, required int) bool {
	// For in-memory indexer, we assume that if there are any attestations for the messageID, we have a quorum
	verifierProofs, ok := at.attestations.Load(msg.Header.MessageID)
	if !ok {
		return false // No attestations found for this messageID
	}
	count := 0
	verifierProofs.(*syncmap.Map).Range(func(_, _ interface{}) bool {
		count++
		if count >= required {
			return false // Stop iterating if we already have enough attestations
		}
		return true
	})
	return count >= required
}

// AddAttestation adds an attestation for a given messageID and verifierID
// This is meant to be called as a hook whenever a verifier pushes to their storage
func (at *InMemoryAttestationIndexer) AddAttestation(msgId [32]byte, verifierId [32]byte, proof []byte) error {
	verifierProofs, ok := at.attestations.Load(msgId)
	// If no attestations exist for this messageID, create a new map and store the attestation
	if !ok {
		newMap := syncmap.Map{}
		newMap.Store(verifierId, proof)
		at.attestations.Store(msgId, &newMap)
		return nil
	}

	// If previous proofs exist, add proof to the existing map
	// we use LoadOrStore to check if the verifierId already has a proof. If so, we should error
	_, loaded := verifierProofs.(*syncmap.Map).LoadOrStore(verifierId, proof)
	if loaded {
		return fmt.Errorf("tried adding an attestation for message %x and verifier %x, but it already exists", msgId, verifierId)
	}
	return nil
}
