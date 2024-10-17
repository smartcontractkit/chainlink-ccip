package ccipocr3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitPluginReport(t *testing.T) {
	t.Run("is empty", func(t *testing.T) {
		r := CommitPluginReport{}
		assert.True(t, r.IsEmpty())
	})

	t.Run("is not empty", func(t *testing.T) {
		r := CommitPluginReport{
			MerkleRoots: make([]MerkleRootChain, 1),
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
			RMNSignatures: make([]RMNECDSASignature, 1),
		}
		assert.False(t, r.IsEmpty())

		r = CommitPluginReport{
			MerkleRoots: make([]MerkleRootChain, 1),
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
