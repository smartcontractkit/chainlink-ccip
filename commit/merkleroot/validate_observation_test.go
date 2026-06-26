package merkleroot

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/commontypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

func Test_validateObservedMerkleRoots(t *testing.T) {
	testCases := []struct {
		name                    string
		merkleRoots             []cciptypes.MerkleRootChain
		observer                commontypes.OracleID
		observerSupportedChains mapset.Set[cciptypes.ChainSelector]
		expErr                  bool
	}{
		{
			name: "Chain not supported",
			merkleRoots: []cciptypes.MerkleRootChain{
				{ChainSel: 1, SeqNumsRange: [2]cciptypes.SeqNum{10, 20}, MerkleRoot: [32]byte{1, 2, 3}},
				{ChainSel: 2, SeqNumsRange: [2]cciptypes.SeqNum{24, 45}, MerkleRoot: [32]byte{1, 2, 3}},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 3, 4, 5),
			expErr:                  true,
		},
		{
			name: "Duplicate chains",
			merkleRoots: []cciptypes.MerkleRootChain{
				{ChainSel: 1, SeqNumsRange: [2]cciptypes.SeqNum{10, 20}, MerkleRoot: [32]byte{1, 2, 3}},
				{ChainSel: 2, SeqNumsRange: [2]cciptypes.SeqNum{24, 45}, MerkleRoot: [32]byte{1, 2, 3}},
				{ChainSel: 2, SeqNumsRange: [2]cciptypes.SeqNum{3, 7}, MerkleRoot: [32]byte{1, 2, 3}},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  true,
		},
		{
			name: "Valid offRampMaxSeqNums",
			merkleRoots: []cciptypes.MerkleRootChain{
				{ChainSel: 1, SeqNumsRange: [2]cciptypes.SeqNum{10, 20}, MerkleRoot: [32]byte{1, 2, 3}, OnRampAddress: []byte{1}},
				{ChainSel: 2, SeqNumsRange: [2]cciptypes.SeqNum{24, 24}, MerkleRoot: [32]byte{1, 2, 3}, OnRampAddress: []byte{2}},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  false,
		},
		{
			name: "Missing OnRampAddress",
			merkleRoots: []cciptypes.MerkleRootChain{
				{ChainSel: 1, SeqNumsRange: [2]cciptypes.SeqNum{10, 20}, MerkleRoot: [32]byte{1, 2, 3}, OnRampAddress: []byte{1}},
				{ChainSel: 2, SeqNumsRange: [2]cciptypes.SeqNum{24, 45}, MerkleRoot: [32]byte{1, 2, 3}, OnRampAddress: []byte{}},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  true,
		},
		{
			name: "SeqNums range invalid",
			merkleRoots: []cciptypes.MerkleRootChain{
				{ChainSel: 1, SeqNumsRange: [2]cciptypes.SeqNum{10, 9}, MerkleRoot: [32]byte{1, 2, 3}, OnRampAddress: []byte{1}},
				{ChainSel: 2, SeqNumsRange: [2]cciptypes.SeqNum{24, 45}, MerkleRoot: [32]byte{1, 2, 3}, OnRampAddress: []byte{2}},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedMerkleRoots(tc.merkleRoots, tc.observer, tc.observerSupportedChains)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateObservedOnRampMaxSeqNums(t *testing.T) {
	testCases := []struct {
		name                    string
		onRampMaxSeqNums        []plugintypes.SeqNumChain
		observer                commontypes.OracleID
		observerSupportedChains mapset.Set[cciptypes.ChainSelector]
		expErr                  bool
	}{
		{
			name: "Chain not supported",
			onRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 3, 4, 5),
			expErr:                  true,
		},
		{
			name: "Duplicate chains",
			onRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
				{ChainSel: 2, SeqNum: 33},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  true,
		},
		{
			name: "Valid offRampMaxSeqNums",
			onRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  false,
		},
		{
			name: "Valid if SeqNum is 0",
			onRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 0},
				{ChainSel: 2, SeqNum: 20},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  false,
		},
		{
			name: "Invalid if chain is 0",
			onRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 0, SeqNum: 123},
				{ChainSel: 2, SeqNum: 20},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedOnRampMaxSeqNums(tc.onRampMaxSeqNums, tc.observer, tc.observerSupportedChains)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateObservedOffRampMaxSeqNums(t *testing.T) {
	testCases := []struct {
		name              string
		offRampMaxSeqNums []plugintypes.SeqNumChain
		observer          commontypes.OracleID
		supportsDestChain bool
		expErr            bool
	}{
		{
			name: "Dest chain not supported",
			offRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			observer:          10,
			supportsDestChain: false,
			expErr:            true,
		},
		{
			name: "Duplicate chains",
			offRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
				{ChainSel: 2, SeqNum: 33},
			},
			observer:          10,
			supportsDestChain: false,
			expErr:            true,
		},
		{
			name: "Valid offRampMaxSeqNums",
			offRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			observer:          10,
			supportsDestChain: true,
			expErr:            false,
		},
		{
			name: "Invalid if SeqNum is 0",
			offRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 0},
			},
			observer:          10,
			supportsDestChain: true,
			expErr:            true,
		},
		{
			name: "Invalid if Chain is 0",
			offRampMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 0, SeqNum: 123},
			},
			observer:          10,
			supportsDestChain: true,
			expErr:            true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedOffRampMaxSeqNums(tc.offRampMaxSeqNums, tc.observer, tc.supportsDestChain)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
