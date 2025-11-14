// Package v2
// Implements depositHash calculation for CCTPv2 Messages, which is needed to match CCTPv2
// messages to SourceTokenDataPayloadV2 instances (which have a DepositHash field).
package v2

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

// depositHashABIArguments is the cached ABI arguments structure for deposit hash encoding.
var depositHashABIArguments = initDepositHashABIArguments()

// CalculateDepositHash calculates the depositHash from CCTPv2 message fields.
// This must match Solidity's _calculateDepositHash in USDCSourcePoolDataCodec.sol.
// The hash is: keccak256(abi.encode(sourceDomain, amount, destinationDomain, mintRecipient,
// burnToken, destinationCaller, maxFee, minFinalityThreshold))
//
// This function is critical for matching CCIP token transfers with Circle's attestations.
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

	// Use the cached ABI arguments structure
	arguments := depositHashABIArguments

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
		return zero, fmt.Errorf("failed to ABI encode deposit hash parameters: %w", err)
	}

	// Calculate Keccak256 hash
	hash := crypto.Keccak256Hash(encoded)
	return hash, nil
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
	params.sourceDomain, err = parseUint32Param(msg.SourceDomain, "source domain")
	if err != nil {
		return params, err
	}

	// Parse destinationDomain
	params.destinationDomain, err = parseUint32Param(msg.DestinationDomain, "destination domain")
	if err != nil {
		return params, err
	}

	// Parse amount
	params.amount = new(big.Int)
	if _, ok := params.amount.SetString(msg.DecodedMessageBody.Amount, 10); !ok {
		return params, fmt.Errorf("failed to parse amount: invalid number")
	}
	if params.amount.Sign() < 0 {
		return params, fmt.Errorf("failed to parse amount: negative amount not allowed")
	}

	// Parse maxFee (defaults to 0 if empty)
	params.maxFee = new(big.Int)
	if msg.DecodedMessageBody.MaxFee != "" {
		if _, ok := params.maxFee.SetString(msg.DecodedMessageBody.MaxFee, 10); !ok {
			return params, fmt.Errorf("failed to parse max fee: invalid number")
		}
		if params.maxFee.Sign() < 0 {
			return params, fmt.Errorf("failed to parse max fee: negative fee not allowed")
		}
	}

	// Parse minFinalityThreshold
	params.minFinalityThreshold, err = parseUint32Param(msg.MinFinalityThreshold, "min finality threshold")
	if err != nil {
		return params, err
	}

	// Parse hex addresses to bytes32
	params.mintRecipient, err = hexToBytes32(msg.DecodedMessageBody.MintRecipient)
	if err != nil {
		return params, fmt.Errorf("failed to parse mint recipient: %w", err)
	}

	params.burnToken, err = hexToBytes32(msg.DecodedMessageBody.BurnToken)
	if err != nil {
		return params, fmt.Errorf("failed to parse burn token: %w", err)
	}

	params.destinationCaller, err = hexToBytes32(msg.DestinationCaller)
	if err != nil {
		return params, fmt.Errorf("failed to parse destination caller: %w", err)
	}

	return params, nil
}

// initDepositHashABIArguments creates the ABI arguments structure needed for encoding.
// The order of these arguments must match the Solidity implementation exactly.
func initDepositHashABIArguments() abi.Arguments {
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
		return result, fmt.Errorf("failed to decode hex string: %w", err)
	}

	// Validate length
	if len(decoded) > 32 {
		return result, fmt.Errorf("failed to convert hex to bytes32: hex string too long (%d bytes)", len(decoded))
	}

	// Copy to bytes32 (left-pads with zeros for EVM addresses, which are 20 bytes)
	copy(result[32-len(decoded):], decoded)

	return result, nil
}

// parseUint32Param converts the given string to a uint32
func parseUint32Param(param string, paramName string) (uint32, error) {
	parsed, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s: %w", paramName, err)
	}
	return uint32(parsed), nil
}
