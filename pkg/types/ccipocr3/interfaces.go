package ccipocr3

import (
	"context"
)

// TODO: Consolidate CommitPluginCodec, ExecutePluginCodec, MessageHasher, ExtraDataCodec into a single Codec interface.

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

type ExtraDataCodec interface {
	// Deprecated: DecodeExtraData reformat bytes into a chain agnostic map[string]any representation for extra args
	DecodeExtraData(extraArgs Bytes, sourceChainSelector ChainSelector) (map[string]any, error)
	// DecodeExtraArgs reformat bytes into a chain agnostic map[string]any representation for extra args
	DecodeExtraArgs(extraArgs Bytes, sourceChainSelector ChainSelector) (map[string]any, error)
	// DecodeTokenAmountDestExecData reformat bytes to chain-agnostic map[string]any for tokenAmount DestExecData field
	DecodeTokenAmountDestExecData(destExecData Bytes, sourceChainSelector ChainSelector) (map[string]any, error)
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
		signerAddresses []UnknownAddress,
	) error
}

// TokenDataEncoder is a generic interface for encoding offchain token data for different chain families.
// Right now it supports only USDC/CCTP, but every new token that requires offchain data processsing
// should be added to that interface and implemented in the downstream repositories (e.g. chainlink-ccip, chainlink).
type TokenDataEncoder interface {
	EncodeUSDC(ctx context.Context, message Bytes, attestation Bytes) (Bytes, error)
}

// EstimateProvider is used to estimate the gas cost of a message or a merkle tree.
type EstimateProvider interface {
	CalculateMerkleTreeGas(numRequests int) uint64
	CalculateMessageMaxGas(msg Message) uint64
}
