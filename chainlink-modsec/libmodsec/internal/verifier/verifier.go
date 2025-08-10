package verifier

import (
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

// Verifier is the main verifier service. It manages the lifecycle of data
// fetching, handling, and writing a resulting attestation.
type Verifier struct {
	workCh chan Work
	stopCh chan struct{}

	// configurable components
	handlers    []Handler
	signer      modsectypes.Signer
	reader      Reader
	transformer Transformer
	writer      []Writer
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

// WithReader sets a reader for the verifier
func WithReader(reader Reader) Option {
	return func(v *Verifier) {
		v.reader = reader
	}
}

// WithTransformer sets the transformer for the verifier
func WithTransformer(transformer Transformer) Option {
	return func(v *Verifier) {
		v.transformer = transformer
	}
}

// WithWriter adds a writer for the verifier
func WithWriter(writer Writer) Option {
	return func(v *Verifier) {
		v.writer = append(v.writer, writer)
	}
}

// WithHandler adds a handler function, there can be more than one.
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
