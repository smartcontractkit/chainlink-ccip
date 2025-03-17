package ccipocr3

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommitPluginReport(t *testing.T) {
	t.Run("is empty", func(t *testing.T) {
		r := CommitPluginReport{}
		assert.True(t, r.IsEmpty())

		// If a report only contains signatures it is still considered empty.
		r = CommitPluginReport{
			RMNSignatures: make([]RMNECDSASignature, 1),
		}
		assert.True(t, r.IsEmpty())
	})

	t.Run("is not empty", func(t *testing.T) {
		r := CommitPluginReport{
			BlessedMerkleRoots: make([]MerkleRootChain, 1),
		}
		assert.False(t, r.IsEmpty())

		r = CommitPluginReport{
			PriceUpdates: PriceUpdates{
				TokenPriceUpdates: make([]TokenPrice, 1),
				GasPriceUpdates:   make([]GasPriceChain, 1),
			},
		}
		assert.False(t, r.IsEmpty())

		r = CommitPluginReport{
			BlessedMerkleRoots: make([]MerkleRootChain, 1),
			PriceUpdates: PriceUpdates{
				TokenPriceUpdates: make([]TokenPrice, 1),
				GasPriceUpdates:   make([]GasPriceChain, 1),
			},
			RMNSignatures: make([]RMNECDSASignature, 1),
		}
		assert.False(t, r.IsEmpty())
	})
}

func TestMerkleRootChain_Equals_Structs(t *testing.T) {
	tests := []struct {
		name     string
		m1       MerkleRootChain
		m2       MerkleRootChain
		expected bool
	}{
		{
			name: "equal MerkleRootChains",
			m1: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			m2: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			expected: true,
		},
		{
			name: "different ChainSel",
			m1: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			m2: MerkleRootChain{
				ChainSel:      ChainSelector(2),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			expected: false,
		},
		{
			name: "different OnRampAddress",
			m1: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			m2: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x03, 0x04},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			expected: false,
		},
		{
			name: "different SeqNumsRange",
			m1: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			m2: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{2, 20},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			expected: false,
		},
		{
			name: "different MerkleRoot",
			m1: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x01, 0x02, 0x03},
			},
			m2: MerkleRootChain{
				ChainSel:      ChainSelector(1),
				OnRampAddress: []byte{0x01, 0x02},
				SeqNumsRange:  SeqNumRange{1, 10},
				MerkleRoot:    Bytes32{0x04, 0x05, 0x06},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.m1.Equals(tt.m2))
		})
	}
}

func TestMerkleRootChain_String(t *testing.T) {
	mrc := MerkleRootChain{
		ChainSel:      123,
		OnRampAddress: []byte{0x01, 0x02},
		SeqNumsRange:  SeqNumRange{123, 456},
		MerkleRoot:    Bytes32{1, 2, 3},
	}

	s := mrc.String()
	assert.Equal(t, "MerkleRoot(chain:123, seqNumsRange:[123 -> 456], "+
		"merkleRoot:0x0102030000000000000000000000000000000000000000000000000000000000, onRamp:0x0102)", s)
}

func TestDecodeCommitReportInfo(t *testing.T) {
	// Empty input
	{
		info, err := DecodeCommitReportInfo(nil)
		require.NoError(t, err)
		require.Equal(t, CommitReportInfo{}, info)
	}

	// Unsupported version
	{
		data := append([]byte{2}, []byte("{}")...)
		_, err := DecodeCommitReportInfo(data)
		require.ErrorContains(t, err, "unknown execute report info version (2)")
	}

	// Invalid field
	{
		data := append([]byte{1}, []byte(`{"InvalidField": 123}`)...)
		_, err := DecodeCommitReportInfo(data)
		require.ErrorContains(t, err, "unknown field")
	}

	// Valid input
	{
		validReport := CommitReportInfo{
			RemoteF:     1,
			MerkleRoots: []MerkleRootChain{},
			PriceUpdates: PriceUpdates{
				TokenPriceUpdates: []TokenPrice{
					{
						TokenID: "0x1234",
						Price:   NewBigInt(big.NewInt(123)),
					},
				},
			},
		}
		encoded, err := json.Marshal(validReport)
		require.NoError(t, err)

		data := append([]byte{1}, encoded...)
		decoded, err := DecodeCommitReportInfo(data)
		require.NoError(t, err)
		require.Equal(t, validReport, decoded)
	}

	// Non-object input
	{
		data := append([]byte{1}, []byte(`["unexpected array"]`)...)
		_, err := DecodeCommitReportInfo(data)
		require.ErrorContains(t, err, "cannot unmarshal array")
	}
}
