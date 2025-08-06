package modsectypes

import (
	"context"
	"math/big"
	"time"
)

// Service is an interface that provides a way to start and stop a service.
type Service interface {
	Start(ctx context.Context) error
	Close() error
}

// ReceiptTypeVerifier is the type of receipt for a verifier.
const ReceiptTypeVerifier uint8 = 0

// ReceiptTypeExecutor is the type of receipt for an executor.
const ReceiptTypeExecutor uint8 = 1

// Header is the header of a CCIP message.
type Header struct {
	MessageID           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
}

// TokenTransfer is a token transfer in a CCIP message.
type TokenTransfer struct {
	SourceTokenAddress []byte
	SourcePoolAddress  []byte
	DestTokenAddress   []byte
	ExtraData          []byte
	Amount             *big.Int
	DestExecData       []byte
	RequiredVerifierID [32]byte
}

type Receipt struct {
	ReceiptType       uint8
	Issuer            []byte
	FeeTokenAmount    *big.Int
	DestGasLimit      uint64
	DestBytesOverhead uint32
	ExtraArgs         []byte
}

// TODO: implement the type.
type Message struct {
	Header         Header
	Sender         []byte
	Data           []byte
	Receiver       []byte
	FeeToken       []byte
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenTransfer  TokenTransfer
	Receipts       []Receipt
}

// EventCodec is an interface that provides a way to decode chain data into a message.
// Messages can be emitted alongside other data, e.g. on EVM the message is emitted alongside
// other indexed fields for ease of querying with plain ethereum clients.
type EventCodec interface {
	// Decode decodes chain data into a message.
	Decode(ctx context.Context, data []byte) (Message, error)
}

// MessageCodec is an interface that provides a way to encode and decode messages for specific
// chain families.
type MessageCodec interface {
	// Encode encodes a message into the chain-native encoding.
	Encode(ctx context.Context, message Message) ([]byte, error)

	// Decode decodes a chain-native encoding into a message.
	Decode(ctx context.Context, data []byte) (Message, error)
}

// StorageReader is an interface that provides a way to read data from an offchain storage
// endpoint.
// This is used by executors to read the proofs from the offchain storage.
type StorageReader interface {
	Get(ctx context.Context, key string) ([]byte, error)
	List(ctx context.Context, prefix string, startTime time.Time) ([]string, error)
}

// StorageWriter is an interface that provides a way to write data to an offchain storage.
// This is used by verifiers to write their proofs.
type StorageWriter interface {
	Put(ctx context.Context, key string, value []byte) error
}

// MessageQueryArgs are query arguments to read messages from the source chain onramp.
type MessageQueryArgs struct {
	// DestChainSelectors specifies the destination chain selectors to which the messages
	// are being sent to.
	DestChainSelectors []uint64

	// StartSeqNums specifies the sequence numbers to start reading from for each
	// destination chain selector.
	StartSeqNums []uint64

	// Limit specifies the maximum number of messages to read.
	Limit int
}

// SourceReader is an interface that provides a way to read messages from the source chain onramp.
// This is expected to be implemented per chain family.
type SourceReader interface {
	Service

	// Messages returns a channel of messages that are read from the source chain onramp.
	Messages() <-chan Message

	// GetMessages returns a slice of messages from the source chain onramp starting
	// from the given sequence number.
	GetMessages(ctx context.Context, query MessageQueryArgs) ([]Message, error)
}

// DestReader is an interface that provides a way to read data from the destination chain.
// This is expected to be implemented per chain family.
type DestReader interface {
	Service

	// GetNonce returns the nonce for the given account on the destination chain.
	GetNonce(ctx context.Context, sourceChainSelector uint64, account []byte) (uint64, error)

	// IsExecuted returns true if the given message has been executed on the destination chain.
	IsExecuted(ctx context.Context, message Message) (bool, error)
}

type WorkerType string

const (
	WorkerTypeExecutor WorkerType = "executor"
	WorkerTypeVerifier WorkerType = "verifier"
)

// WorkerRecord is a record of a worker that has been created onchain.
type WorkerRecord struct {
	// WorkerType is the type of worker that was created.
	WorkerType WorkerType

	// Created is true if the worker was created onchain, false if it was deleted.
	Created bool

	// Config is the configuration of the worker, JSON encoded.
	Config []byte
}

// HomeReader is an interface that provides a way to read data from the home chain.
// This is expected to be implemented only once, for the home chain family (Ethereum).
type HomeReader interface {
	Service

	// GetAllWorkers returns all the workers that have been created onchain.
	GetAllWorkers(ctx context.Context) ([]WorkerRecord, error)

	// Workers returns a channel of worker records that are read from the home chain.
	Workers() <-chan WorkerRecord
}

// ContractTransmitter is an interface that provides a way to execute a message on the destination
// chain.
// This is expected to be implemented per chain family.
type ContractTransmitter interface {
	// Transmit transmits a message to the destination chain.
	// extraPayload is any extra data required by the transmitter to form the final transaction,
	// expected to be used on chains like Solana.
	Transmit(ctx context.Context, encodedMessage []byte, proofs [][]byte, extraPayload []byte) error
}

// Signer is an interface that provides a way to sign a message.
// This is expected to be implemented per chain family.
type Signer interface {
	// Sign signs the provided digest. For EVM the digest is 32 bytes long though for different
	// hashing functions it may be longer, hence the []byte rather than [32]byte.
	Sign(ctx context.Context, digest []byte) ([]byte, error)
}

// Hasher is an interface that provides a way to hash a message.
// This is expected to be implemented per chain family.
type Hasher interface {
	// Hash hashes the provided data.
	// data must be appropriately formatted prior to the call - e.g. if a config digest
	// needs to be included alongside message data.
	Hash(ctx context.Context, data []byte) ([]byte, error)
}
