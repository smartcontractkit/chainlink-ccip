package ccipocr3

import (
	"context"
)

type CommitPluginCodec interface {
	Encode(context.Context, CommitPluginReport) ([]byte, error)
	Decode(context.Context, []byte) (CommitPluginReport, error)
}

type ExecutePluginCodec interface {
	Encode(context.Context, ExecutePluginReport) ([]byte, error)
	Decode(context.Context, []byte) (ExecutePluginReport, error)
}

type MessageHasher interface {
	Hash(context.Context, Message) (Bytes32, error)
}

// RMNCrypto provides a chain-agnostic interface for verifying RMN signatures.
// For example, on EVM, RMN reports are abi-encoded prior to being signed.
// On Solana, they would be borsh encoded instead, etc.
type RMNCrypto interface {
	// VerifyReportSignatures verifies each provided signature against the provided report and the signer addresses.
	// If any signature is invalid (no matching signer address is found), an error is returned immediately.
	VerifyReportSignatures(
		ctx context.Context,
		sigs []RMNECDSASignature,
		report RMNReport,
		signerAddresses []Bytes,
	) error
}

// TokenDataEncoder is a generic interface for encoding offchain token data for different chain families.
// Right now it supports only USDC/CCTP, but every new token that requires offchain data processsing
// should be added to that interface and implemented in the downstream repositories (e.g. chainlink-ccip, chainlink).
type TokenDataEncoder interface {
	EncodeUSDC(ctx context.Context, message Bytes, attestation Bytes) (Bytes, error)
}
