package v2

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// SourceTokenDataPayload represents the ABI-encoded token data payload for CCTP v2 transfers.
// This struct contains all the parameters needed to perform a cross-chain USDC transfer,
// including burn parameters, destination details, and fees. It's extracted from the
// ExtraData field of CCIP token amounts.
// Nonce is expected to be 0 for CCTP v2 transfers, as the CCTP v2 protocol does not return a nonce on-chain. This
// field exists for backwards compatibility.
type SourceTokenDataPayload struct {
	Nonce                uint64             `abi:"nonce"`
	SourceDomain         uint32             `abi:"sourceDomain"`
	CCTPVersion          reader.CCTPVersion `abi:"cctpVersion"`
	Amount               cciptypes.BigInt   `abi:"amount"`
	DestinationDomain    uint32             `abi:"destinationDomain"`
	MintRecipient        cciptypes.Bytes32  `abi:"mintRecipient"`
	BurnToken            cciptypes.Bytes32  `abi:"burnToken"`
	DestinationCaller    cciptypes.Bytes32  `abi:"destinationCaller"`
	MaxFee               cciptypes.BigInt   `abi:"maxFee"`
	MinFinalityThreshold uint32             `abi:"minFinalityThreshold"`
}

// matchesCctpMessage checks if the SourceTokenDataPayload matches the provided CCTPv2 Message.
func matchesCctpMessage(
	s SourceTokenDataPayload,
	m Message,
) bool {
	if int(s.CCTPVersion) != m.CCTPVersion {
		return false
	}
	if fmt.Sprintf("%d", s.SourceDomain) != m.DecodedMessage.SourceDomain {
		return false
	}
	if fmt.Sprintf("%d", s.DestinationDomain) != m.DecodedMessage.DestinationDomain {
		return false
	}

	// MinFinalityThreshold is optional
	if m.DecodedMessage.MinFinalityThreshold != "" &&
		fmt.Sprintf("%d", s.MinFinalityThreshold) != m.DecodedMessage.MinFinalityThreshold {
		return false
	}

	if s.Amount.String() != m.DecodedMessage.DecodedMessageBody.Amount {
		return false
	}

	// MaxFee is optional
	if m.DecodedMessage.DecodedMessageBody.MaxFee != "" &&
		s.MaxFee.String() != m.DecodedMessage.DecodedMessageBody.MaxFee {
		return false
	}

	if !addressMatch(m.DecodedMessage.DestinationCaller, s.DestinationCaller) {
		return false
	}

	if !addressMatch(m.DecodedMessage.DecodedMessageBody.BurnToken, s.BurnToken) {
		return false
	}

	if !addressMatch(m.DecodedMessage.DecodedMessageBody.MintRecipient, s.MintRecipient) {
		return false
	}

	return true
}

// addressMatch returns true if the provided cctpAddress matches the right-aligned bytes of sourceAddress
// This is needed because the address returned by the CCTP v2 API is not padded (e.g. EVM addresses will be 20 bytes,
// Solana addresses will be 32 bytes, etc) however sourceAddress is always 32 bytes, and for EVM addresses (which are
// 20 bytes) the leftmost 12 bytes will be zero due to ABI encoding.
func addressMatch(cctpAddress string, sourceAddress cciptypes.Bytes32) bool {
	cctpAddressBytes, err := hex.DecodeString(strings.TrimPrefix(cctpAddress, "0x"))
	if err != nil {
		return false
	}
	if len(cctpAddressBytes) > 32 {
		return false
	}

	return bytes.Equal(sourceAddress[32-len(cctpAddressBytes):], cctpAddressBytes)
}

// getCCTPv2SourceTokenDataPayload extracts and validates a CCTP v2 source token data payload
// from a CCIP token amount. It verifies that the token's source pool address matches the
// configured CCTP v2 pool and that the payload contains valid CCTP v2 data.
func getCCTPv2SourceTokenDataPayload(
	cctpV2EnabledTokenPoolAddress string,
	msgToken cciptypes.RampTokenAmount,
) (*SourceTokenDataPayload, error) {
	if !strings.EqualFold(cctpV2EnabledTokenPoolAddress, msgToken.SourcePoolAddress.String()) {
		return nil, fmt.Errorf("unsupported token pool address")
	}

	tokenData, err := DecodeSourceTokenDataPayload(msgToken.ExtraData)
	if err != nil {
		return nil, err
	}

	if tokenData.CCTPVersion != reader.CctpVersion2 {
		return nil, fmt.Errorf("unsupported CCTP version: %d", tokenData.CCTPVersion)
	}

	return tokenData, nil
}

func DecodeSourceTokenDataPayload(data []byte) (*SourceTokenDataPayload, error) {
	argTypes := abi.Arguments{
		{Type: mustABIType("uint64")},
		{Type: mustABIType("uint32")},
		{Type: mustABIType("uint8")},
		{Type: mustABIType("uint256")},
		{Type: mustABIType("uint32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("bytes32")},
		{Type: mustABIType("uint256")},
		{Type: mustABIType("uint32")},
	}

	vals, err := argTypes.Unpack(data)
	if err != nil {
		return nil, err
	}
	if len(vals) != 10 {
		return nil, fmt.Errorf("unexpected number of unpacked values")
	}

	return &SourceTokenDataPayload{
		Nonce:                vals[0].(uint64),
		SourceDomain:         vals[1].(uint32),
		CCTPVersion:          reader.CCTPVersion(vals[2].(uint8)),
		Amount:               cciptypes.NewBigInt(vals[3].(*big.Int)),
		DestinationDomain:    vals[4].(uint32),
		MintRecipient:        vals[5].([32]byte),
		BurnToken:            vals[6].([32]byte),
		DestinationCaller:    vals[7].([32]byte),
		MaxFee:               cciptypes.NewBigInt(vals[8].(*big.Int)),
		MinFinalityThreshold: vals[9].(uint32),
	}, nil
}

func mustABIType(t string) abi.Type {
	typ, err := abi.NewType(t, "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}
