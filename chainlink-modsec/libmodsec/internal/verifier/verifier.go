package verifier

import (
	"context"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

// Poller is an interface that defines how to poll for new work to be processed.
type Poller interface {
	// Next returns a channel that will yield the next piece of work to be processed.
	Next(ctx context.Context) <-chan Work

	// Watch returns a channel that will yield all work to be processed.
	Watch(ctx context.Context) <-chan Work
}

// Transformer is an interface that defines how to transform a Work item into a modsec message.
// TODO: tranform from source to verifier and again from verifier to destination?
type Transformer interface {
	Transform(work Work) HandlerPayload
}

// StoredMessage represents a message that will be attested. It is the object written to the storage layer.
type StoredMessage struct {
	ID      string
	Message modsectypes.Message
	Encoded []byte
	// block info?
}

// MessageAttestation represents the verifier's attestation of a message.
type MessageAttestation struct {
	Attestor string
	Sig      []byte
}

// AttestationWriter is an interface that defines how to write attestations and messages to a storage layer.
// The writer is responsible for deciding how to store the messages and attestations, including
// how to avoid duplicates, and how to structure the storage for efficient scanning.
type AttestationWriter interface {
	// StoreMessage stores the message that is being attested.
	// The writer implementation is responsible for deciding how to store the message.
	// It should consider things like how it will be retrieved later, and how to avoid duplicates.
	// For example, it may decide to store messages in a hierarchy based on the time, block number, sequence number, etc.
	StoreMessage(ctx context.Context, msg StoredMessage) error

	// StoreAttestation stores the attestation for a message.
	// The writer implementation is responsible for deciding how to store the attestation.
	// It should consider things like how it will be retrieved later, and how to avoid duplicates.
	StoreAttestation(ctx context.Context, msg MessageAttestation) error
}

// Verifier is the main verifier service. It manages the lifecycle of data
// fetching, handling, and writing a resulting attestation.
type Verifier struct {
	workCh chan Work
	stopCh chan struct{}

	// handlers are registered by name.
	handlers []Handler

	// state
	started bool

	// configurable services
	signer      modsectypes.Signer
	poller      Poller
	transformer Transformer
	writer      []AttestationWriter
	// Add more configurable fields as needed
}

// Option is the Verifier functional option type
type Option func(*Verifier)

// WithSigner sets a custom signer for the verifier
func WithSigner(signer modsectypes.Signer) Option {
	return func(v *Verifier) {
		v.signer = signer
	}
}

// WithPoller sets a custom poller for the verifier
func WithPoller(poller Poller) Option {
	return func(v *Verifier) {
		v.poller = poller
	}
}

// WithTransformer sets the transformer for the verifier
func WithTransformer(transformer Transformer) Option {
	return func(v *Verifier) {
		v.transformer = transformer
	}
}

// WithWriter adds a writer for the verifier
func WithWriter(writer AttestationWriter) Option {
	return func(v *Verifier) {
		v.writer = append(v.writer, writer)
	}
}

// WithHandler adds a handler function for a given name. Only one handler can be registered per name.
func WithHandler(handler Handler) Option {
	return func(v *Verifier) {
		v.handlers = append(v.handlers, handler)
	}
}

func NewVerifier(opts ...Option) *Verifier {
	v := &Verifier{
		stopCh: make(chan struct{}),
	}
	// Apply all options
	for _, opt := range opts {
		opt(v)
	}
	return v
}

func (s *Verifier) Start() {
	go s.run()
}

func (s *Verifier) Stop() {
	close(s.stopCh)
}
