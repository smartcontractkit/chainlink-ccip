// Package v2 provides shared types for CCTPv2 token data observation.
package v2

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// AttestationEncoder encodes CCTP message and attestation into format expected by USDC token pool.
// The encoder is responsible for combining the CCTP message bytes with the attestation signature
// in the format required by the destination chain's USDC token pool contract.
type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

// TxKey uniquely identifies a transaction on a specific source chain.
// It's used as a composite key for grouping CCTP messages that belong to the same transaction.
type TxKey struct {
	// SourceDomain is the Circle domain ID of the source chain
	SourceDomain uint32
	// TxHash is the transaction hash on the source chain
	TxHash string
}

// depositHashResult stores the result of a depositHash calculation.
// This is used to pre-calculate and cache depositHash values for all Circle messages
// in a transaction to avoid redundant calculations during matching.
type depositHashResult struct {
	// hash is the calculated depositHash, or zero value if error occurred
	hash [32]byte
	// err is non-nil if the depositHash calculation failed
	err error
}