package executor

import (
	"context"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

// MessageReader is an interface that defines methods for reading messages to be executed
// by the executor. It is used to abstract the source of messages, allowing for
// implementations to read directly from verifier storage or from message indexer
type MessageReader interface {
	// SubscribeMessages returns a channel that will yield messages to be processed.
	// It's up to the implementation of the reader to decide how to fetch messages (if we want to base on sequence numbers, etc.)
	SubscribeMessages(ctx context.Context) <-chan modsectypes.Message
}

// AttestationReader is an interface that defines methods for reading attestations of a message
// It's up to the implementation to decide how to fetch attestations from different sources.
type AttestationReader interface {
	// GetAttestations retrieves attestations for a given message.
	// This function signature assumes that we can fetch attestations based on messageID
	// we will need this at minimum for manual execution
	// TODO: do returned attestations still need to be associated with its verifierID?
	// If not, we can just return []byte instead of creating the Attestation struct
	GetAttestations(ctx context.Context, msg modsectypes.Message) ([]Attestation, error)
}
