package common

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Constants for CCIP v1.7
const (
	MaxNumTokens = 1
	MaxDataSize  = 1024 // 1kb
)

var (
	// Domain separators and message type hashes
	EVM2AnyMessageHash          = crypto.Keccak256([]byte("EVM_2_ANY_MESSAGE_HASH"))
	Any2EVMMessageHash          = crypto.Keccak256([]byte("ANY_2_EVM_MESSAGE_HASH"))
	LeafDomainSeparator         = make([]byte, 32)
	InternalNodeDomainSeparator = append(make([]byte, 31), byte(1))
)

// UnknownAddress represents an address on an unknown chain.
type UnknownAddress []byte

// NewUnknownAddressFromHex creates an UnknownAddress from a hex string
func NewUnknownAddressFromHex(s string) (UnknownAddress, error) {
	if s == "" {
		return UnknownAddress{}, nil
	}

	// Remove 0x prefix if present
	if len(s) >= 2 && s[:2] == "0x" {
		s = s[2:]
	}

	bytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string: %w", err)
	}

	return UnknownAddress(bytes), nil
}

// String returns the hex representation of the address
func (a UnknownAddress) String() string {
	if len(a) == 0 {
		return ""
	}
	return "0x" + hex.EncodeToString(a)
}

// Bytes returns the raw bytes of the address
func (a UnknownAddress) Bytes() []byte {
	return []byte(a)
}

// MessageHeader represents the common header for all CCIP messages
type MessageHeader struct {
	MessageID           cciptypes.Bytes32       `json:"message_id"`
	SourceChainSelector cciptypes.ChainSelector `json:"source_chain_selector"`
	DestChainSelector   cciptypes.ChainSelector `json:"dest_chain_selector"`
	SequenceNumber      cciptypes.SeqNum        `json:"sequence_number"`
}

// TokenTransfer represents a token transfer in the CCIP protocol
type TokenTransfer struct {
	SourceTokenAddress UnknownAddress `json:"source_token_address"`
	DestTokenAddress   UnknownAddress `json:"dest_token_address"`
	ExtraData          []byte         `json:"extra_data"`
	Amount             *big.Int       `json:"amount"`
}

// Receipt represents return data from a verifier or executor
type Receipt struct {
	Issuer            UnknownAddress `json:"issuer"`
	FeeTokenAmount    *big.Int       `json:"fee_token_amount"`
	DestGasLimit      uint64         `json:"dest_gas_limit"`
	DestBytesOverhead uint32         `json:"dest_bytes_overhead"`
	ExtraArgs         []byte         `json:"extra_args"`
}

// EVM2AnyVerifierMessage represents a message sent from an EVM chain
type EVM2AnyVerifierMessage struct {
	Header           MessageHeader  `json:"header"`
	Sender           UnknownAddress `json:"sender"`
	Data             []byte         `json:"data"`
	Receiver         UnknownAddress `json:"receiver"`
	FeeToken         UnknownAddress `json:"fee_token"`
	FeeTokenAmount   *big.Int       `json:"fee_token_amount"`
	FeeValueJuels    *big.Int       `json:"fee_value_juels"`
	TokenTransfer    TokenTransfer  `json:"token_transfer"`
	VerifierReceipts []Receipt      `json:"verifier_receipts"`
	ExecutorReceipt  *Receipt       `json:"executor_receipt"`
	TokenReceipt     *Receipt       `json:"token_receipt"`
	ExtraArgs        []byte         `json:"extra_args"`
}

// Any2AnyVerifierMessage represents a chain-agnostic CCIP message
type Any2AnyVerifierMessage struct {
	Header           MessageHeader  `json:"header"`
	Sender           UnknownAddress `json:"sender"`
	OnRampAddress    UnknownAddress `json:"onramp_address"` // CCVProxy address
	Data             []byte         `json:"data"`
	Receiver         UnknownAddress `json:"receiver"`
	FeeToken         UnknownAddress `json:"fee_token"`
	FeeTokenAmount   *big.Int       `json:"fee_token_amount"`
	FeeValueJuels    *big.Int       `json:"fee_value_juels"`
	TokenTransfer    TokenTransfer  `json:"token_transfer"`
	VerifierReceipts []Receipt      `json:"verifier_receipts"`
	ExecutorReceipt  *Receipt       `json:"executor_receipt"`
	TokenReceipt     *Receipt       `json:"token_receipt"`
	ExtraArgs        []byte         `json:"extra_args"`
}

// Any2EVMMessageMetadata represents metadata for Any2EVM messages
type Any2EVMMessageMetadata struct {
	SourceChainSelector cciptypes.ChainSelector `json:"source_chain_selector"`
	DestChainSelector   cciptypes.ChainSelector `json:"dest_chain_selector"`
	OnRampAddress       UnknownAddress          `json:"onramp_address"`
}

// Any2EVMVerifierMessage represents a message to be executed on an EVM chain
type Any2EVMVerifierMessage struct {
	Header        MessageHeader  `json:"header"`
	Sender        UnknownAddress `json:"sender"`
	Data          []byte         `json:"data"`
	Receiver      UnknownAddress `json:"receiver"`
	TokenTransfer TokenTransfer  `json:"token_transfer"`
	GasLimit      uint32         `json:"gas_limit"`
	ExtraArgs     []byte         `json:"extra_args"`
	OnRampAddress UnknownAddress `json:"onramp_address"`
}

// VerificationTask represents the complete CCIPMessageSent event data from the onRamp/proxy
// This struct wraps the Any2AnyVerifierMessage with additional event fields that are important
// for verification and processing
type VerificationTask struct {
	Message      Any2AnyVerifierMessage `json:"message"`       // the complete message
	ReceiptBlobs [][]byte               `json:"receipt_blobs"` // receipt blobs from event
}

// CCVData represents Cross-Chain Verification data
type CCVData struct {
	MessageID             cciptypes.Bytes32       `json:"message_id"`
	SequenceNumber        cciptypes.SeqNum        `json:"sequence_number"`
	SourceChainSelector   cciptypes.ChainSelector `json:"source_chain_selector"`
	DestChainSelector     cciptypes.ChainSelector `json:"dest_chain_selector"`
	SourceVerifierAddress UnknownAddress          `json:"source_verifier_address"`
	DestVerifierAddress   UnknownAddress          `json:"dest_verifier_address"`
	CCVData               []byte                  `json:"ccv_data"`  // The actual proof/signature
	BlobData              []byte                  `json:"blob_data"` // Additional verifier-specific data
	Timestamp             int64                   `json:"timestamp"` // Unix timestamp when verification completed (in microseconds)
	Message               Any2AnyVerifierMessage  `json:"message"`   // Complete message event being verified
}

// TimestampQueryResponse represents the response from timestamp-based CCV data queries.
// Contains the queried data organized by destination chain along with
// pagination metadata for efficient executor polling workflows.
type TimestampQueryResponse struct {
	// Data organized by destination chain selector
	Data map[cciptypes.ChainSelector][]CCVData `json:"data"`
	// Next timestamp to query (nil if no more data)
	NextTimestamp *int64 `json:"next_timestamp,omitempty"`
	// Whether more data exists at current timestamp
	HasMore bool `json:"has_more"`
	// Total number of items returned
	TotalCount int `json:"total_count"`
}

// OffchainStorageWriter defines the interface for verifiers to store CCV data.
// This interface is used by CCIP verifiers to store their CCV data
// after verification. Each verifier has write access to its own storage instance.
type OffchainStorageWriter interface {
	// StoreCCVData stores multiple CCV data entries in the offchain storage
	StoreCCVData(ctx context.Context, ccvDataList []CCVData) error
}

// OffchainStorageReader defines the interface for executors to query CCV data by timestamp.
// This interface is used by CCIP executors to poll for new CCV data using
// timestamp-based queries with offset pagination. Designed for efficient
// executor polling workflows.
type OffchainStorageReader interface {
	// GetCCVDataByTimestamp queries CCV data by timestamp with offset-based pagination
	GetCCVDataByTimestamp(
		ctx context.Context,
		destChainSelectors []cciptypes.ChainSelector,
		startTimestamp int64,
		sourceChainSelectors []cciptypes.ChainSelector,
		limit int,
		offset int,
	) (*TimestampQueryResponse, error)
}
