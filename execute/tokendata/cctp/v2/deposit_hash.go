// Package v2 provides depositHash calculation and decoding functionality for CCTPv2.
// This file contains the core cryptographic functions for working with CCTPv2 depositHashes,
// which are content-addressable identifiers used for matching cross-chain transfers.
package v2

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// CCTP version tags for identifying V2 messages.
const (
	// CCTP_VERSION_2_TAG identifies standard CCTP V2 transfers (slow transfers).
	// Preimage: keccak256("CCTP_V2")
	CCTP_VERSION_2_TAG = 0xb148ea5f

	// CCTP_VERSION_2_CCV_TAG identifies CCTP V2 transfers with CCIP v1.7 fast transfer support.
	// CCV = Cross-Chain Verification. Enables fast transfers with verification infrastructure.
	// Preimage: keccak256("CCTP_V2_CCV")
	CCTP_VERSION_2_CCV_TAG = 0x3047587c
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
//
// Example:
//
//	extraData := []byte{...} // 40 bytes from CCIP message
//	payload, err := DecodeSourceTokenDataPayloadV2(extraData)
//	if err != nil {
//	    // Handle invalid payload
//	}
//	// Use payload.SourceDomain and payload.DepositHash for matching
func DecodeSourceTokenDataPayloadV2(extraData cciptypes.Bytes) (*SourceTokenDataPayloadV2, error) {
	// Validate length
	if len(extraData) != 40 {
		return nil, fmt.Errorf("invalid V2 source pool data length: expected 40 bytes, got %d", len(extraData))
	}

	// Extract and validate version tag (bytes 0-3)
	versionTag := binary.BigEndian.Uint32(extraData[0:4])
	if versionTag != CCTP_VERSION_2_TAG && versionTag != CCTP_VERSION_2_CCV_TAG {
		return nil, fmt.Errorf("invalid CCTPv2 version tag: expected 0x%x or 0x%x, got 0x%x",
			CCTP_VERSION_2_TAG, CCTP_VERSION_2_CCV_TAG, versionTag)
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

// depositHashParams contains all parsed parameters needed for depositHash calculation.
type depositHashParams struct {
	sourceDomain         uint32
	destinationDomain    uint32
	amount               *big.Int
	maxFee               *big.Int
	minFinalityThreshold uint32
	mintRecipient        [32]byte
	burnToken            [32]byte
	destinationCaller    [32]byte
}

// parseDepositHashParams extracts and validates all parameters from a CCTPv2DecodedMessage.
// This function handles all string parsing and hex decoding, returning a structured
// set of parameters ready for ABI encoding.
func parseDepositHashParams(msg CCTPv2DecodedMessage) (depositHashParams, error) {
	var params depositHashParams
	var err error

	// Parse sourceDomain
	sourceDomain64, err := strconv.ParseUint(msg.SourceDomain, 10, 32)
	if err != nil {
		return params, fmt.Errorf("parse sourceDomain: %w", err)
	}
	params.sourceDomain = uint32(sourceDomain64)

	// Parse destinationDomain
	destDomain64, err := strconv.ParseUint(msg.DestinationDomain, 10, 32)
	if err != nil {
		return params, fmt.Errorf("parse destinationDomain: %w", err)
	}
	params.destinationDomain = uint32(destDomain64)

	// Parse amount
	params.amount = new(big.Int)
	if _, ok := params.amount.SetString(msg.DecodedMessageBody.Amount, 10); !ok {
		return params, fmt.Errorf("parse amount: invalid number")
	}
	if params.amount.Sign() < 0 {
		return params, fmt.Errorf("parse amount: negative amount not allowed")
	}

	// Parse maxFee (defaults to 0 if empty)
	params.maxFee = new(big.Int)
	if msg.DecodedMessageBody.MaxFee != "" {
		if _, ok := params.maxFee.SetString(msg.DecodedMessageBody.MaxFee, 10); !ok {
			return params, fmt.Errorf("parse maxFee: invalid number")
		}
		if params.maxFee.Sign() < 0 {
			return params, fmt.Errorf("parse maxFee: negative fee not allowed")
		}
	}

	// Parse minFinalityThreshold
	minFinality64, err := strconv.ParseUint(msg.MinFinalityThreshold, 10, 32)
	if err != nil {
		return params, fmt.Errorf("parse minFinalityThreshold: %w", err)
	}
	params.minFinalityThreshold = uint32(minFinality64)

	// Parse hex addresses to bytes32
	params.mintRecipient, err = hexToBytes32(msg.DecodedMessageBody.MintRecipient)
	if err != nil {
		return params, fmt.Errorf("parse mintRecipient: %w", err)
	}

	params.burnToken, err = hexToBytes32(msg.DecodedMessageBody.BurnToken)
	if err != nil {
		return params, fmt.Errorf("parse burnToken: %w", err)
	}

	params.destinationCaller, err = hexToBytes32(msg.DestinationCaller)
	if err != nil {
		return params, fmt.Errorf("parse destinationCaller: %w", err)
	}

	return params, nil
}

// createDepositHashABIArguments creates the ABI arguments structure needed for encoding.
// The order of these arguments must match the Solidity implementation exactly.
func createDepositHashABIArguments() abi.Arguments {
	// Create ABI types for encoding
	uint32Type, _ := abi.NewType("uint32", "", nil)
	uint256Type, _ := abi.NewType("uint256", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)

	// Define ABI encoding arguments (matches Solidity's abi.encode order)
	return abi.Arguments{
		{Type: uint32Type},  // sourceDomain
		{Type: uint256Type}, // amount
		{Type: uint32Type},  // destinationDomain
		{Type: bytes32Type}, // mintRecipient
		{Type: bytes32Type}, // burnToken
		{Type: bytes32Type}, // destinationCaller
		{Type: uint256Type}, // maxFee
		{Type: uint32Type},  // minFinalityThreshold
	}
}

// CalculateDepositHash calculates the depositHash from CCTPv2 message fields.
// This must match Solidity's _calculateDepositHash in USDCSourcePoolDataCodec.sol.
// The hash is: keccak256(abi.encode(sourceDomain, amount, destinationDomain, mintRecipient,
// burnToken, destinationCaller, maxFee, minFinalityThreshold))
//
// This function is critical for matching CCIP messages with Circle's attestations.
//
// Parameters from msg are parsed and validated, then ABI-encoded in the exact order required
// by the Solidity implementation. Any mismatch in encoding will result in hash mismatches and
// failed attestation matching.
func CalculateDepositHash(msg CCTPv2DecodedMessage) ([32]byte, error) {
	var zero [32]byte

	// Parse all parameters from the message
	params, err := parseDepositHashParams(msg)
	if err != nil {
		return zero, err
	}

	// Create ABI arguments structure
	arguments := createDepositHashABIArguments()

	// Pack the arguments using ABI encoding
	encoded, err := arguments.Pack(
		params.sourceDomain,
		params.amount,
		params.destinationDomain,
		params.mintRecipient,
		params.burnToken,
		params.destinationCaller,
		params.maxFee,
		params.minFinalityThreshold,
	)
	if err != nil {
		return zero, fmt.Errorf("ABI encode: %w", err)
	}

	// Calculate Keccak256 hash
	hash := crypto.Keccak256Hash(encoded)
	return hash, nil
}

// hexToBytes32 converts a hex string (with or without 0x prefix) to a [32]byte array.
// This function handles addresses from different chains with varying byte lengths:
// - EVM addresses are 20 bytes and are left-padded with zeros to 32 bytes
// - Solana addresses are already 32 bytes
// The padding ensures consistent byte32 representation across different blockchain address formats.
func hexToBytes32(hexStr string) ([32]byte, error) {
	var result [32]byte

	// Remove 0x prefix if present
	hexStr = strings.TrimPrefix(hexStr, "0x")

	// Decode hex string
	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		return result, fmt.Errorf("decode hex: %w", err)
	}

	// Validate length
	if len(decoded) > 32 {
		return result, fmt.Errorf("hex string too long: %d bytes", len(decoded))
	}

	// Copy to bytes32 (left-padded with zeros for addresses, which are 20 bytes)
	copy(result[32-len(decoded):], decoded)

	return result, nil
}
