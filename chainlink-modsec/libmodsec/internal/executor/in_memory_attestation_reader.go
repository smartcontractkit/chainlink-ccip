package executor

import (
	"context"
	"fmt"
	indexer "github.com/smartcontractkit/chainlink-modsec/libmodsec/internal/attestation-indexer"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
	"time"
)

type InMemoryAttestationReader struct {
	InMemoryIndexer indexer.InMemoryAttestationIndexer
}

// NewInMemoryAttestationReader creates a new InMemoryAttestationReader instance
func NewInMemoryAttestationReader() *InMemoryAttestationReader {
	return &InMemoryAttestationReader{}
}

func (at *InMemoryAttestationReader) GetAttestations(ctx context.Context, msg modsectypes.Message) ([]Attestation, error) {
	startTime := time.Now()
	timeoutCh := make(chan struct{})
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context deadline exceeded")
			return nil, nil
		case <-timeoutCh:
			return nil, fmt.Errorf("timed out waiting for attestations for message %s", msg.Header.MessageID)
		default:
			if at.hasTimedOut(startTime) {
				timeoutCh <- struct{}{} // Signal that we have timed out
				continue
			}
			// check if we have enough attestations
			// TODO: get required number of attestations from the message header or config
			if !at.InMemoryIndexer.MessageHasQuorum(ctx, msg, 1) {
				time.Sleep(2 * time.Second) // Sleep for a short duration before checking again
				continue
			}

			return at.InMemoryIndexer.GetAttestationsForMessage(ctx, msg)

		}
	}
}

func (at *InMemoryAttestationReader) hasTimedOut(startTime time.Time) bool {
	// Check if the operation has timed out after 30 seconds
	return time.Since(startTime) > 30*time.Second
}
