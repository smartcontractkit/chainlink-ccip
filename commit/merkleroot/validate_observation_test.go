package merkleroot

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

func Test_validateRMNRemoteConfig(t *testing.T) {
	testCases := []struct {
		name              string
		observer          commontypes.OracleID
		supportsDestChain bool
		rmnRemoteConfig   cciptypes.RemoteConfig
		expectedError     bool
	}{
		{
			name:              "is valid",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
		},
		{
			name:              "does not support destination chain",
			observer:          1,
			supportsDestChain: false, // <--
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty contract address",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{}, // <--
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty config digest",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{}, // <---
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "not enough signers to cover F+1 threshold",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1}, // <----
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty rmn report version",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{},
			},
			expectedError: true,
		},
		{
			name:              "duplicate signers",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 1}, // <---------
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty signer onchain public key",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: cciptypes.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []cciptypes.RemoteSignerInfo{
					{OnchainPublicKey: []byte{}, NodeIndex: 0}, // <-----
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				FSign:            1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			err := validateRMNRemoteConfig(tt.observer, tt.supportsDestChain, tt.rmnRemoteConfig)
			if tt.expectedError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
