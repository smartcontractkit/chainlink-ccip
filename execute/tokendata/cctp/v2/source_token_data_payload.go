// Package v2 provides source token data payload decoding functionality for CCTPv2.
// This file contains the structures and functions for decoding CCTPv2 source token data payloads
// that are embedded in CCIP TokenAmount ExtraData fields.
package v2

import (
	"encoding/binary"
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// CCTP version tags for identifying V2 messages.
const (
	// CCTPVersion2Tag identifies standard CCTP V2 transfers (slow transfers).
	// Preimage: keccak256("CCTP_V2")
	CCTPVersion2Tag = 0xb148ea5f

	// CCTPVersion2CCVTag identifies CCTP V2 transfers with CCIP v1.7 fast transfer support.
	// CCV = Cross-Chain Verification. Enables fast transfers with verification infrastructure.
	// Preimage: keccak256("CCTP_V2_CCV")
	CCTPVersion2CCVTag = 0x3047587c
)

// SourceTokenDataPayloadV2 represents the CCTPv2 source pool data embedded in message token data.
// This payload is extracted from the ExtraData field of CCIP TokenAmount messages and contains
// the source domain and depositHash needed to match with Circle's attestation API responses.
type SourceTokenDataPayloadV2 struct {
	// SourceDomain is the Circle domain ID of the source chain
	SourceDomain uint32
	// DepositHash is the content-addressable hash that uniquely identifies a CCTPv2 transfer.
	// It's calculated as keccak256(abi.encode(sourceDomain, amount, destinationDomain,
	// mintRecipient, burnToken, destinationCaller, maxFee, minFinalityThreshold))
	DepositHash [32]byte
}

// DecodeSourceTokenDataPayloadV2 decodes SourceTokenDataPayloadV2 from ExtraData bytes.
// The payload is encoded as: bytes4(versionTag) + uint32(sourceDomain) + bytes32(depositHash)
// Total length: 40 bytes (4 + 4 + 32)
func DecodeSourceTokenDataPayloadV2(extraData cciptypes.Bytes) (*SourceTokenDataPayloadV2, error) {
	// Validate length
	if len(extraData) != 40 {
		return nil, fmt.Errorf("invalid V2 source pool data length: expected 40 bytes, got %d", len(extraData))
	}

	// Extract and validate version tag (bytes 0-3)
	versionTag := binary.BigEndian.Uint32(extraData[0:4])
	if versionTag != CCTPVersion2Tag && versionTag != CCTPVersion2CCVTag {
		return nil, fmt.Errorf("invalid CCTPv2 version tag: expected 0x%x or 0x%x, got 0x%x",
			CCTPVersion2Tag, CCTPVersion2CCVTag, versionTag)
	}

	// Extract sourceDomain (bytes 4-7, big-endian uint32)
	sourceDomain := binary.BigEndian.Uint32(extraData[4:8])

	// Extract depositHash (bytes 8-39)
	var depositHash [32]byte
	copy(depositHash[:], extraData[8:40])

	return &SourceTokenDataPayloadV2{
		SourceDomain: sourceDomain,
		DepositHash:  depositHash,
	}, nil
}
