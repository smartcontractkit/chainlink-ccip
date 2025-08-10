package verifier

import (
	"context"
)

// Reader is an interface for reading the blockchain for new work.
type Reader interface {
	// Next returns a channel that will yield the next piece of work to be processed.
	Next(ctx context.Context) <-chan Work

	// Watch returns a channel that will yield all work to be processed.
	Watch(ctx context.Context) <-chan Work
}

// Transformer is an interface that defines how to transform a Work item into
// the generic modsec.Message format and encoded into a destination specific
// payload. This interface encapsulates most of the chain agnostic services
// required by the 1.6 implementation.
type Transformer interface {
	Transform(work Work) HandlerPayload
}

// Writer is an interface that defines how to write attestations and messages to a storage layer.
// The writer is responsible for deciding how to store the messages and attestations, including
// how to avoid duplicates, and how to structure the storage for efficient scanning.
//
// TODO: standard io.Writer interface could be used to simplify testing.
// Modex extensions would look like "modsec.NewS3Writer(path)"
type Writer interface {
	// WriteMessage stores the message that is being attested.
	// The writer implementation is responsible for deciding how to store the message.
	// It should consider things like how it will be retrieved later, and how to avoid duplicates.
	// For example, it may decide to store messages in a hierarchy based on the time, block number, sequence number, etc.
	WriteMessage(ctx context.Context, msg HandlerPayload) error

	// WriteAttestation stores the attestation for a message.
	// The writer implementation is responsible for deciding how to store the attestation.
	// It should consider things like how it will be retrieved later, and how to avoid duplicates.
	WriteAttestation(ctx context.Context, msg Attestation) error
}

// Handler is a function that processes incoming work from the work channel and sends results to the result channel.
// There may be no results for a particular piece of work, in which case the result channel should not be written to.
// For example, the verifier may have a rule to skip processing until a certain block depth is reached. In that case,
// the work may be received, but no result will be produced until later.
//
// The verifier should respect the context for cancellation. If the context is cancelled, the verifier should stop.
//
// TODO: should verifier be a full service with start/stop/cancel, a cache, maybe a db connection, etc?
type Handler func(ctx context.Context, payload HandlerPayload, result chan<- Attestation)
