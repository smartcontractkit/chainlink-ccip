package v2

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const sourceTokenDataPayloadHex = "0x" +
	"000000000000000000000000000000000000000000000000000000000000007b" + // uint64  nonce
	"000000000000000000000000000000000000000000000000000000000000006f" + // uint32  sourceDomain
	"0000000000000000000000000000000000000000000000000000000000000001" + // uint8   cctpVersion
	"00000000000000000000000000000000000000000000000000000000000003e8" + // uint256 amount
	"0000000000000000000000000000000000000000000000000000000012345678" + // uint32  destinationDomain
	"0000000000000000000000001234567890abcdef1234567890abcdef12345678" + // bytes32 mintRecipient
	"2222222222222222222222222222222222222222222222222222222222222222" + // bytes32 burnToken
	"3333333333333333333333333333333333333333333333333333333333333333" + // bytes32 destinationCaller
	"0000000000000000000000000000000000000000000000000000000000000032" + // uint256 maxFee
	"0000000000000000000000000000000000000000000000000000000000000005" // uint32  minFinalityThreshold

func TestDecodeSourceTokenDataPayload(t *testing.T) {
	tests := []struct {
		name      string
		bytes     []byte
		expectErr bool
		expected  *SourceTokenDataPayload
	}{
		{
			name:      "Decode valid payload",
			bytes:     mustBytes(sourceTokenDataPayloadHex),
			expectErr: false,
			expected: &SourceTokenDataPayload{
				Nonce:                123,
				SourceDomain:         111,
				CCTPVersion:          reader.CCTPVersion(1),
				Amount:               cciptypes.NewBigInt(big.NewInt(1000)),
				DestinationDomain:    0x12345678,
				MintRecipient:        mustBytes32("0x0000000000000000000000001234567890abcdef1234567890abcdef12345678"),
				BurnToken:            mustBytes32("0x2222222222222222222222222222222222222222222222222222222222222222"),
				DestinationCaller:    mustBytes32("0x3333333333333333333333333333333333333333333333333333333333333333"),
				MaxFee:               cciptypes.NewBigInt(big.NewInt(50)),
				MinFinalityThreshold: 5,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := DecodeSourceTokenDataPayload(tc.bytes)
			if tc.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("failed to decode SourceTokenDataPayload: %v", err)
			}

			// compare scalar fields
			if got.Nonce != tc.expected.Nonce ||
				got.SourceDomain != tc.expected.SourceDomain ||
				got.DestinationDomain != tc.expected.DestinationDomain ||
				got.MinFinalityThreshold != tc.expected.MinFinalityThreshold ||
				uint8(got.CCTPVersion) != uint8(tc.expected.CCTPVersion) {
				t.Fatalf("scalar fields mismatch\nwant %+v\ngot  %+v", tc.expected, got)
			}

			// compare big.Int-wrapped values via String()
			if got.Amount.String() != tc.expected.Amount.String() ||
				got.MaxFee.String() != tc.expected.MaxFee.String() {
				t.Fatalf("big.Int fields mismatch\nwant %v / %v\ngot  %v / %v",
					tc.expected.Amount, tc.expected.MaxFee, got.Amount, got.MaxFee)
			}

			// compare bytes32
			if !bytes.Equal(got.MintRecipient[:], tc.expected.MintRecipient[:]) ||
				!bytes.Equal(got.BurnToken[:], tc.expected.BurnToken[:]) ||
				!bytes.Equal(got.DestinationCaller[:], tc.expected.DestinationCaller[:]) {
				t.Fatalf("bytes32 fields mismatch\nwant %+v\ngot  %+v", tc.expected, got)
			}

			// fall-back reflect check to catch anything missed
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("DecodeSourceTokenDataPayload mismatch\nwant %+v\ngot  %+v", tc.expected, got)
			}
		})
	}
}

// mustBytes32 converts a 0x-prefixed hex string to a [32]byte.
func mustBytes32(h string) (out [32]byte) {
	b, err := hex.DecodeString(strings.TrimPrefix(h, "0x"))
	if err != nil {
		panic(err)
	}
	copy(out[32-len(b):], b) // right-align like EVM abi-encoding
	return
}

func mustBytes(h string) []byte {
	b, err := hex.DecodeString(strings.TrimPrefix(h, "0x"))
	if err != nil {
		panic(err)
	}
	return b
}

func happyPathPayload() SourceTokenDataPayload {
	return SourceTokenDataPayload{
		CCTPVersion:          1,
		SourceDomain:         111,
		DestinationDomain:    222,
		MinFinalityThreshold: 5,
		Amount:               cciptypes.NewBigIntFromInt64(1_000),
		MaxFee:               cciptypes.NewBigIntFromInt64(10),
		BurnToken:            mustBytes32("0x1111"),
		MintRecipient:        mustBytes32("0x2222"),
		DestinationCaller:    mustBytes32("0x3333"),
	}
}

func happyPathMessage() Message {
	return Message{
		CCTPVersion: 1,
		DecodedMessage: DecodedMessage{
			SourceDomain:         "111",
			DestinationDomain:    "222",
			MinFinalityThreshold: "5",
			DestinationCaller:    "0x3333",
			DecodedMessageBody: DecodedMessageBody{
				BurnToken:     "0x1111",
				MintRecipient: "0x2222",
				Amount:        "1000",
				MaxFee:        "10",
			},
		},
	}
}

func TestMatchesCctpMessage(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*SourceTokenDataPayload, *Message) // introduce mismatch
		wantOK bool
	}{
		{
			name:   "all fields match (happy path)",
			mutate: nil,
			wantOK: true,
		},
		{
			name: "CCTP version mismatch",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				p.CCTPVersion = 2
			},
			wantOK: false,
		},
		{
			name: "source domain mismatch",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				m.DecodedMessage.SourceDomain = "999"
			},
			wantOK: false,
		},
		{
			name: "destination domain mismatch",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				p.DestinationDomain = 888
			},
			wantOK: false,
		},
		{
			name: "min-finality mismatch (field present)",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				m.DecodedMessage.MinFinalityThreshold = "42"
			},
			wantOK: false,
		},
		{
			name: "min-finality absent on message (should still match)",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				m.DecodedMessage.MinFinalityThreshold = ""
			},
			wantOK: true,
		},
		{
			name: "amount mismatch",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				m.DecodedMessage.DecodedMessageBody.Amount = "999"
			},
			wantOK: false,
		},
		{
			name: "max-fee mismatch (field present)",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				p.MaxFee = cciptypes.NewBigIntFromInt64(55)
			},
			wantOK: false,
		},
		{
			name: "max-fee absent on message (should still match)",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				m.DecodedMessage.DecodedMessageBody.MaxFee = ""
			},
			wantOK: true,
		},
		{
			name: "destination caller mismatch",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				m.DecodedMessage.DestinationCaller = "0xdeadbeef"
			},
			wantOK: false,
		},
		{
			name: "burn token mismatch",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				m.DecodedMessage.DecodedMessageBody.BurnToken = "0x9999"
			},
			wantOK: false,
		},
		{
			name: "mint recipient mismatch",
			mutate: func(p *SourceTokenDataPayload, m *Message) {
				p.MintRecipient = mustBytes32("0xAAAA")
			},
			wantOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := happyPathPayload()
			m := happyPathMessage()
			if tc.mutate != nil {
				tc.mutate(&p, &m)
			}
			got := matchesCctpMessage(p, m)
			if got != tc.wantOK {
				t.Fatalf("expected %v, got %v", tc.wantOK, got)
			}
		})
	}
}

func TestAddressMatch(t *testing.T) {
	tests := []struct {
		name          string
		cctpAddress   string
		sourceAddress cciptypes.Bytes32
		expectMatch   bool
	}{
		{
			name:        "Exact match (20-byte address)",
			cctpAddress: "0x8fe6b999dc680ccfdd5bf7eb0974218be2542daa",
			sourceAddress: func() cciptypes.Bytes32 {
				var b cciptypes.Bytes32
				addr, _ := hex.DecodeString("0000000000000000000000008fe6b999dc680ccfdd5bf7eb0974218be2542daa")
				copy(b[:], addr)
				return b
			}(),
			expectMatch: true,
		},
		{
			name:        "Exact match (32-byte address)",
			cctpAddress: "0x111122223333444455556666777788889999aaaabbbbccccddddeeeeffff0000",
			sourceAddress: func() cciptypes.Bytes32 {
				var b cciptypes.Bytes32
				addr, _ := hex.DecodeString("111122223333444455556666777788889999aaaabbbbccccddddeeeeffff0000")
				copy(b[:], addr)
				return b
			}(),
			expectMatch: true,
		},
		{
			name:        "Mismatch address",
			cctpAddress: "0x1111111111111111111111111111111111111111",
			sourceAddress: func() cciptypes.Bytes32 {
				var b cciptypes.Bytes32
				addr, _ := hex.DecodeString("8fe6b999dc680ccfdd5bf7eb0974218be2542daa")
				copy(b[12:], addr)
				return b
			}(),
			expectMatch: false,
		},
		{
			name:          "Too long address (>32 bytes)",
			cctpAddress:   "0x" + strings.Repeat("aa", 33),
			sourceAddress: cciptypes.Bytes32{},
			expectMatch:   false,
		},
		{
			name:          "Malformed hex",
			cctpAddress:   "0x8fe6b999dc680ccfzzzzzzzz",
			sourceAddress: cciptypes.Bytes32{},
			expectMatch:   false,
		},
		{
			name:        "Shorter byte slice match (e.g. last 4 bytes)",
			cctpAddress: "0xdeadbeef",
			sourceAddress: func() cciptypes.Bytes32 {
				var b cciptypes.Bytes32
				b[28] = 0xde
				b[29] = 0xad
				b[30] = 0xbe
				b[31] = 0xef
				return b
			}(),
			expectMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addressMatch(tt.cctpAddress, tt.sourceAddress)
			require.Equal(t, tt.expectMatch, got)
		})
	}
}
