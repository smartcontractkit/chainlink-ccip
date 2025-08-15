package executor

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

type Executor struct {
	stopCh    chan struct{}
	messageCh chan modsectypes.Message

	// timedMessageCh is an optional timed message channel for configurable tick-based message delivery
	timedMessageCh modsectypes.TimedMessageChannel

	// List of sources which can potentially provide messages to be executed. We use an array here in case
	// custom executors want to subscribe directly to a verifier storage
	// Example could be a kafka topic, CL commit verifier, etc.
	MessageSources []MessageReader

	// List of sources which can potentially provide attestations for messages
	// This could be changed to a map of verifierID to attestation reader
	AttestationReader AttestationReader

	// Map of destination chains to writers. A writer is some abstraction on the chain access layer
	Transmitters map[uint64]modsectypes.ContractTransmitter

	// Used for encoding/decoding messages
	messageCodec modsectypes.MessageCodec

	// Node's peerId used to determine whether it's this executor's turn to process a message
	peerId [32]byte

	// leaderElection is used to determine if it's this executor's turn to process a message
	leaderElection LeaderElection
}

// Option is the Executor functional option type
type Option func(*Executor) error

// NewExecutor creates a new Executor and applies the provided options
func NewExecutor(opts ...Option) *Executor {
	// Apply all options
	e := &Executor{}
	for _, opt := range opts {
		err := opt(e)
		if err != nil {
			// Handle error, could be logging or panic based on use case
			fmt.Printf("Error applying option: %v\n", err)
			return nil // or handle error appropriately
		}
	}
	return e
}

// WithMessageSources adds message readers to the executor. An example message source is
// the CL commit verifier that tells us about all messages pushed to the chain
func WithMessageSources(readers []MessageReader) Option {
	return func(e *Executor) error {
		if e.MessageSources == nil {
			e.MessageSources = make([]MessageReader, 0)
		}
		for _, reader := range readers {
			// we could add a health check here or validation that this executor is allowlisted on the reader
			if reader == nil {
				return fmt.Errorf("cannot add nil reader to executor")
			}
			e.MessageSources = append(e.MessageSources, reader)
		}
		return nil
	}
}

// WithAttestationReader adds an attestation reader to the executor
// // An attestation reader is an abstraction that allows us to fetch attestations for messages from different sources
// handling multiple attestation sources is up to the implementation of the reader
func WithAttestationReader(r AttestationReader) Option {
	return func(e *Executor) error {
		if r == nil {
			return fmt.Errorf("cannot add nil attestation reader to executor")
		}

		e.AttestationReader = r
		return nil
	}
}

// WithDestChainTransmitter adds a Contract Transmitter for a specific destination chain to the executor
func WithDestChainTransmitter(chain uint64, writer modsectypes.ContractTransmitter) Option {
	return func(e *Executor) error {
		if writer == nil {
			return fmt.Errorf("cannot add nil writer to executor")
		}
		if e.Transmitters == nil {
			e.Transmitters = make(map[uint64]modsectypes.ContractTransmitter)
		}

		if _, exists := e.Transmitters[chain]; exists {
			return fmt.Errorf("writer for chain %d already exists", chain)
		}
		e.Transmitters[chain] = writer
		return nil
	}
}

// WithLeaderElection adds a leader election algorithm to the executor
func WithLeaderElection(leaderElection LeaderElection) Option {
	return func(e *Executor) error {
		if leaderElection == nil {
			return fmt.Errorf("cannot add nil leader election to executor")
		}
		e.leaderElection = leaderElection
		return nil
	}
}

// WithTimedMessageChannel adds a timed message channel to the executor
// This allows messages to be sent with configurable tick delays
func WithTimedMessageChannel(tmc modsectypes.TimedMessageChannel) Option {
	return func(e *Executor) error {
		if tmc == nil {
			return fmt.Errorf("cannot add nil timed message channel to executor")
		}
		e.timedMessageCh = tmc
		return nil
	}
}

func (e *Executor) Start() {
	go e.run()
}

func (e *Executor) Stop() {
	close(e.stopCh)
}

func (e *Executor) isMyTurn(msg modsectypes.Message) (bool, time.Duration) {
	// Check if the message is the leader for the destination chain
	return e.leaderElection.IsLeader(msg.Header.MessageID, msg.Header.DestChainSelector)
}
