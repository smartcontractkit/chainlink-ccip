package verifier

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/common"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// SourceReaderConfig contains configuration for the EVM source reader
type SourceReaderConfig struct {
	ChainSelector       cciptypes.ChainSelector `json:"chain_selector"`
	OnRampAddress       common.UnknownAddress   `json:"onramp_address"`
	PollInterval        time.Duration           `json:"poll_interval"`
	StartBlock          uint64                  `json:"start_block,omitempty"`
	MessagesChannelSize int                     `json:"messages_channel_size"`
}

// SourceConfig contains configuration for a single source chain
type SourceConfig struct {
	VerifierAddress common.UnknownAddress `json:"verifier_address"`
}

// CoordinatorConfig contains configuration for the verification coordinator
type CoordinatorConfig struct {
	VerifierID            string                                   `json:"verifier_id"`
	SourceConfigs         map[cciptypes.ChainSelector]SourceConfig `json:"source_configs"`
	ProcessingChannelSize int                                      `json:"processing_channel_size"`
	ProcessingTimeout     time.Duration                            `json:"processing_timeout"`
	MaxBatchSize          int                                      `json:"max_batch_size"`
}

// SourceReader defines the interface for reading CCIP messages from source chains
type SourceReader interface {
	// Start begins reading messages and pushing them to the messages channel
	Start(ctx context.Context) error

	// Stop stops the reader and closes the messages channel
	Stop() error

	// VerificationTaskChannel returns the channel where new message events are delivered
	VerificationTaskChannel() <-chan common.VerificationTask

	// HealthCheck returns the current health status of the reader
	HealthCheck(ctx context.Context) error
}

// MessageSigner defines the interface for signing messages
type MessageSigner interface {
	// SignMessage signs a message event and returns the signature
	SignMessage(ctx context.Context, verificationTask common.VerificationTask) ([]byte, error)

	// GetSignerAddress returns the address of the signer
	GetSignerAddress() common.UnknownAddress
}

// ECDSASigner implements MessageSigner using ECDSA
type ECDSASigner struct {
	privateKey []byte
	address    common.UnknownAddress
}

// TOOD: implement properly
// NewECDSAMessageSigner creates a new ECDSA message signer
func NewECDSAMessageSigner(privateKey []byte) (*ECDSASigner, error) {
	if len(privateKey) == 0 {
		return nil, fmt.Errorf("private key cannot be empty")
	}

	// For simplicity, use the first 20 bytes of the private key as address
	// In a real implementation, this would derive the public key and address
	addressBytes := make([]byte, 20)
	copy(addressBytes, privateKey[:20])

	return &ECDSASigner{
		privateKey: privateKey,
		address:    common.UnknownAddress(addressBytes),
	}, nil
}

// TOOD: implement properly
// SignMessage signs a message event using ECDSA
func (s *ECDSASigner) SignMessage(ctx context.Context, messageEvent common.VerificationTask) ([]byte, error) {
	// Create a hash of the message for signing
	messageHash := make([]byte, 32)
	copy(messageHash, messageEvent.Message.Header.MessageID[:])

	// For simplicity, return a mock signature
	// In a real implementation, this would use crypto/ecdsa
	signature := make([]byte, 65) // Standard ECDSA signature length
	copy(signature, s.privateKey[:32])
	copy(signature[32:], messageHash[:32])
	signature[64] = 0x1b // Recovery ID

	return signature, nil
}

// GetSignerAddress returns the address of the signer
func (s *ECDSASigner) GetSignerAddress() common.UnknownAddress {
	return s.address
}
