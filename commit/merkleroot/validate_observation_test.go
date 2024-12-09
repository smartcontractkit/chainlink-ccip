package merkleroot

import (
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
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
				{ChainSel: 1, SeqNumsRange: [2]cciptypes.SeqNum{10, 20}, MerkleRoot: [32]byte{1, 2, 3}},
				{ChainSel: 2, SeqNumsRange: [2]cciptypes.SeqNum{24, 45}, MerkleRoot: [32]byte{1, 2, 3}},
			},
			observer:                10,
			observerSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2),
			expErr:                  false,
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
		rmnRemoteConfig   types.RemoteConfig
		expectedError     bool
	}{
		{
			name:              "is valid",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				F:                1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
		},
		{
			name:              "does not support destination chain",
			observer:          1,
			supportsDestChain: false, // <--
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				F:                1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty contract address",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{}, // <--
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				F:                1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty config digest",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{}, // <---
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				F:                1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "not enough signers to cover F+1 threshold",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1}, // <----
				},
				F:                1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty rmn report version",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				F:                1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{},
			},
			expectedError: true,
		},
		{
			name:              "duplicate signers",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{0, 0, 1}, NodeIndex: 0},
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 1}, // <---------
				},
				F:                1,
				ConfigVersion:    1,
				RmnReportVersion: cciptypes.Bytes32{0, 0, 0, 1},
			},
			expectedError: true,
		},
		{
			name:              "empty signer onchain public key",
			observer:          1,
			supportsDestChain: true,
			rmnRemoteConfig: types.RemoteConfig{
				ContractAddress: []byte{1, 2, 3},
				ConfigDigest:    cciptypes.Bytes32{1, 2, 3},
				Signers: []types.RemoteSignerInfo{
					{OnchainPublicKey: []byte{}, NodeIndex: 0}, // <-----
					{OnchainPublicKey: []byte{0, 0, 2}, NodeIndex: 1},
					{OnchainPublicKey: []byte{0, 0, 3}, NodeIndex: 2},
				},
				F:                1,
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

func Test_validateFChain(t *testing.T) {
	testCases := []struct {
		name   string
		fChain map[cciptypes.ChainSelector]int
		expErr bool
	}{
		{
			name: "FChain contains negative values",
			fChain: map[cciptypes.ChainSelector]int{
				1: 11,
				2: -4,
			},
			expErr: true,
		},
		{
			name: "FChain valid",
			fChain: map[cciptypes.ChainSelector]int{
				12: 6,
				7:  9,
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateFChain(tc.fChain)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateMerkleRootsState(t *testing.T) {
	testCases := []struct {
		name                 string
		onRampNextSeqNum     []plugintypes.SeqNumChain
		offRampExpNextSeqNum []cciptypes.SeqNum
		readerErr            error
		expErr               bool
	}{
		{
			name: "happy path",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: []cciptypes.SeqNum{100, 200},
			expErr:               false,
		},
		{
			name: "one root is stale",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: []cciptypes.SeqNum{100, 201}, // <- 200 is already on chain
			expErr:               true,
		},
		{
			name: "one root has gap",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 101), // <- onchain 99 but we submit 101 instead of 100
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: []cciptypes.SeqNum{100, 200},
			expErr:               true,
		},
		{
			name: "reader returned wrong number of seq nums",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: []cciptypes.SeqNum{100, 200, 300},
			expErr:               true,
		},
		{
			name: "reader error",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: []cciptypes.SeqNum{100, 200},
			readerErr:            fmt.Errorf("reader error"),
			expErr:               true,
		},
	}

	ctx := tests.Context(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := readermock.NewMockCCIPReader(t)
			rep := cciptypes.CommitPluginReport{}
			chains := make([]cciptypes.ChainSelector, 0, len(tc.onRampNextSeqNum))
			for _, snc := range tc.onRampNextSeqNum {
				rep.MerkleRoots = append(rep.MerkleRoots, cciptypes.MerkleRootChain{
					ChainSel:     snc.ChainSel,
					SeqNumsRange: cciptypes.NewSeqNumRange(snc.SeqNum, snc.SeqNum+10),
				})
				chains = append(chains, snc.ChainSel)
			}
			reader.EXPECT().NextSeqNum(ctx, chains).Return(tc.offRampExpNextSeqNum, tc.readerErr)

			err := ValidateMerkleRootsState(ctx, rep.MerkleRoots, reader)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
