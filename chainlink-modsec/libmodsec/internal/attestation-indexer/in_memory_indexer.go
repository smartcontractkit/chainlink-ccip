package attestation_indexer

import (
	"context"
	"fmt"
	"slices"

	"golang.org/x/sync/syncmap"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

// InMemoryAttestationIndexer is an in-memory implementation of an attestation indexer
// It stores all proofs in an in-memory map, indexed by messageID([32]byte) and then by veriiferType(enum) and then verifierID([32]byte)
// This implementation is useful when assuming that attestations will be requested via messageID
// we only organize by verifierType here for the sake of "VerifierTypeMeetsQuorum" function
// Normally, we can just get verifierType from verifierID, but that would require loading all verifiers and iterating over them
// There are significantly more optimizations we can make in a real indexer solution
/**
at.attestations: {
	[messageID1]: {
		[verifierType1]: {
			[verifierID1]: proof1,
			[verifierID2]: proof2,
		},
		[verifierType2]: {
			[verifierID3]: proof3,
			[verifierID4]: proof4,
		},
	},
}
*/
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

// GetAttestationsForMessage retrieves attestations for a given message and converts it into the modsectypes.Attestation type
func (at *InMemoryAttestationIndexer) GetAttestationsForMessage(_ context.Context, msg modsectypes.Message) ([]modsectypes.Attestation, error) {
	// Retrieve the map of attestations for the given messageID
	verifierTypesToIDs, ok := at.attestations.Load(msg.Header.MessageID)
	if !ok {
		return nil, nil // No attestations found for this messageID
	}
	attestations := make([]modsectypes.Attestation, 0)

	// Iterate over the verifier types for this messageID
	verifierTypesToIDs.(*syncmap.Map).Range(func(verifierType, verifierIdsMap interface{}) bool {
		verifierIdsMap.(*syncmap.Map).Range(func(verifierId, proof interface{}) bool {
			// Append the attestation to the list
			attestations = append(attestations, modsectypes.Attestation{
				Proof:      proof.([]byte),
				VerifierId: verifierId.([32]byte),
			})
			return true // Continue iterating over verifierIdsMap
		})
		return true
	})

	return attestations, nil
}

func (at *InMemoryAttestationIndexer) MessageIsReady(ctx context.Context, msg modsectypes.Message) bool {
	verifierTypesReady := make([]bool, 0)
	// we should be looping over all verifier types in the message, but that functionality isn't available yet
	for _, verifierType := range []modsectypes.VerifierType{modsectypes.VerifierTypeCLCommit} {
		verifierTypesReady = append(verifierTypesReady, at.VerifierTypeMeetsQuorum(ctx, msg.Header.MessageID, verifierType, 1))
	}

	// perform an && on all results
	return !slices.Contains(verifierTypesReady, false)
}

// AddAttestation adds an attestation for a given messageID and verifierID
// This is meant to be called as a hook whenever a verifier pushes to their storage
func (at *InMemoryAttestationIndexer) AddAttestation(_ context.Context, msgId [32]byte, verifierId [32]byte, proof []byte) error {
	verifierType, err := at.GetVerifierType(verifierId)
	if err != nil {
		return fmt.Errorf("failed to get verifier type for verifier %x: %w", verifierId, err)
	}
	// If this is the first attestation for this messageID, we need to create a new map for verifier types
	msgToVerifierTypes, _ := at.attestations.LoadOrStore(msgId, &syncmap.Map{})

	// if this is the first attestation for this verifiertype, we need to create a new map for verifierId
	verifierTypesToVerifierIds, _ := msgToVerifierTypes.(*syncmap.Map).LoadOrStore(verifierType, &syncmap.Map{})

	verifierTypesToVerifierIds.(*syncmap.Map).Store(verifierId, proof)
	return nil
}

func (at *InMemoryAttestationIndexer) VerifierTypeMeetsQuorum(_ context.Context, msgId [32]byte, verifierType modsectypes.VerifierType, required int) bool {
	messageIdMap, ok := at.attestations.Load(msgId)
	if !ok {
		return false
	}

	verifierTypeMap, ok := messageIdMap.(*syncmap.Map).Load(verifierType)
	if !ok {
		return false
	}

	verifierTypeMap.(*syncmap.Map).Range(func(_, _ interface{}) bool {
		required--
		if required <= 0 {
			return false // Stop iterating if we already have enough attestations
		}
		return true
	})
	return required <= 0
}

// GetVerifierType returns the verifier type for a given verifierId. It should eventually live in a different place
// This is currently a placeholder that always returns CLCommit
func (at *InMemoryAttestationIndexer) GetVerifierType(_ [32]byte) (modsectypes.VerifierType, error) {
	// In a real implementation, this would look up the verifierId in a database or configuration
	return modsectypes.VerifierTypeCLCommit, nil
}
