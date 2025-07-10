package v2

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// SourceTokenDataPayload TODO: doc
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

// Returns a list of SourceTokenDataPayloads for the provided message. The length of the returned list is equal to
// the number of tokens in the message. If a token is not supported or cannot be decoded, the corresponding entry
// in the returned list will be nil.
func getCCTPv2SourceTokenDataPayloads(
	cctpV2EnabledTokenPoolAddress string,
	message cciptypes.Message,
) []*SourceTokenDataPayload {
	sourceTokenDataPayloads := make([]*SourceTokenDataPayload, 0, len(message.TokenAmounts))
	for _, tokenAmount := range message.TokenAmounts {
		sourceTokenDataPayload, err := getCCTPv2SourceTokenDataPayload(cctpV2EnabledTokenPoolAddress, tokenAmount)
		if err != nil {
			sourceTokenDataPayloads = append(sourceTokenDataPayloads, nil)
		} else {
			sourceTokenDataPayloads = append(sourceTokenDataPayloads, sourceTokenDataPayload)
		}
	}

	return sourceTokenDataPayloads
}

// TODO: doc
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
