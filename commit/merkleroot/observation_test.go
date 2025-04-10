package merkleroot

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	types2 "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/ragep2p/types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/mocks/commit/merkleroot"
	rmn_mock "github.com/smartcontractkit/chainlink-ccip/mocks/commit/merkleroot/rmn"
	common_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func TestObservation(t *testing.T) {
	mockObserver := merkleroot.NewMockObserver(t)
	mockCCIPReader := readerpkg_mock.NewMockCCIPReader(t)
	chainSupport := common_mock.NewMockChainSupport(t)

	destChain := cciptypes.ChainSelector(909606746561742123)

	offchainAddress := []byte(rand.RandomAddress())

	p := &Processor{
		lggr:            logger.Test(t),
		observer:        mockObserver,
		rmnCrypto:       signatureVerifierAlwaysTrue{},
		ccipReader:      mockCCIPReader,
		destChain:       destChain,
		offchainCfg:     pluginconfig.CommitOffchainConfig{RMNEnabled: true},
		chainSupport:    chainSupport,
		metricsReporter: NoopMetrics{},
	}

	thirtyTwoBytes := [32]byte{1, 2, 3}
	ctx := context.Background()

	testCases := []struct {
		name        string
		prevOutcome Outcome
		query       Query
		setupMocks  func()
		expectedObs Observation
		expectedErr string
	}{
		{
			name: "SelectingRangesForReport",
			prevOutcome: Outcome{
				OutcomeType: ReportTransmitted,
			},
			query: Query{},
			setupMocks: func() {
				mockObserver.EXPECT().ObserveOffRampNextSeqNums(mock.Anything).Return(
					[]plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 10}}).Once()
				mockObserver.EXPECT().ObserveLatestOnRampSeqNums(mock.Anything).Return(
					[]plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 15}})
				mockObserver.EXPECT().ObserveRMNRemoteCfg(mock.Anything).Return(cciptypes.RemoteConfig{})
				mockObserver.EXPECT().ObserveFChain(mock.Anything).Return(map[cciptypes.ChainSelector]int{1: 3})
			},
			expectedObs: Observation{
				OffRampNextSeqNums: []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 10}},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 15}},
				RMNRemoteConfig:    cciptypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{1: 3},
			},
		},
		{
			name: "BuildingReport",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: destChain, SeqNumRange: cciptypes.SeqNumRange{5, 10}},
				},
				RMNRemoteCfg: testhelpers.CreateRMNRemoteCfg(),
			},
			query: Query{
				RMNSignatures: &rmn.ReportSignatures{
					Signatures: []*serialization.EcdsaSignature{{R: thirtyTwoBytes[:], S: thirtyTwoBytes[:]}},
				},
			},
			setupMocks: func() {
				mockObserver.EXPECT().ObserveMerkleRoots(mock.Anything, mock.Anything).Return([]cciptypes.MerkleRootChain{
					{
						ChainSel:     1,
						SeqNumsRange: [2]cciptypes.SeqNum{5, 10},
						MerkleRoot:   [32]byte{1},
					}})
				mockObserver.EXPECT().ObserveFChain(mock.Anything).Return(map[cciptypes.ChainSelector]int{1: 3})
				mockCCIPReader.EXPECT().GetContractAddress(mock.Anything, mock.Anything).Return(offchainAddress, nil)
			},
			expectedObs: Observation{
				MerkleRoots: []cciptypes.MerkleRootChain{
					{
						ChainSel:     1,
						SeqNumsRange: [2]cciptypes.SeqNum{5, 10},
						MerkleRoot:   [32]byte{1}},
				},
				RMNEnabledChains: map[cciptypes.ChainSelector]bool{1: true},
				FChain:           map[cciptypes.ChainSelector]int{1: 3},
			},
		},
		{
			name: "WaitingForReportTransmission",
			prevOutcome: Outcome{
				OutcomeType:  ReportInFlight,
				RMNRemoteCfg: testhelpers.CreateRMNRemoteCfg(),
			},
			query: Query{},
			setupMocks: func() {
				mockObserver.EXPECT().ObserveOffRampNextSeqNums(mock.Anything).Return(
					[]plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 20}}).Once()
				mockObserver.EXPECT().ObserveFChain(mock.Anything).Return(map[cciptypes.ChainSelector]int{1: 3})
			},
			expectedObs: Observation{
				OffRampNextSeqNums: []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 20}},
				FChain:             map[cciptypes.ChainSelector]int{1: 3},
			},
		},
		{
			name: "BuildingReport with RetryRMNSignatures",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
			},
			query: Query{
				RetryRMNSignatures: true,
			},
			setupMocks: func() {
				mockObserver.EXPECT().ObserveFChain(mock.Anything).Return(map[cciptypes.ChainSelector]int{1: 3})
			},
			expectedObs: Observation{
				FChain: map[cciptypes.ChainSelector]int{1: 3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks()

			rmnHomeReader := readerpkg_mock.NewMockRMNHome(t)
			rmnHomeReader.EXPECT().GetRMNEnabledSourceChains(tc.prevOutcome.RMNRemoteCfg.ConfigDigest).
				Return(map[cciptypes.ChainSelector]bool{1: true}, nil).Maybe()

			p.rmnHomeReader = rmnHomeReader
			p.rmnControllerCfgDigest = tc.prevOutcome.RMNRemoteCfg.ConfigDigest // skip rmn controller setup
			obs, err := p.Observation(ctx, tc.prevOutcome, tc.query)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedObs, obs)
			}

			mockObserver.AssertExpectations(t)
			mockCCIPReader.AssertExpectations(t)
		})
	}
}

func Test_ObserveOffRampNextSeqNums(t *testing.T) {
	const nodeID commontypes.OracleID = 1
	knownSourceChains := []cciptypes.ChainSelector{4, 7, 19}
	nextSeqNums := map[cciptypes.ChainSelector]cciptypes.SeqNum{4: 345, 7: 608, 19: 7713}

	testCases := []struct {
		name      string
		expResult []plugintypes.SeqNumChain
		getDeps   func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader)
	}{
		{
			name: "Happy path",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)
				ccipReader.EXPECT().NextSeqNum(mock.Anything, knownSourceChains).Return(nextSeqNums, nil)
				ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).Return(reader.CurseInfo{}, nil)
				return chainSupport, ccipReader
			},
			expResult: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(4, 345),
				plugintypes.NewSeqNumChain(7, 608),
				plugintypes.NewSeqNumChain(19, 7713),
			},
		},
		{
			name: "nil is returned when supportsDestChain is false",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(false, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)
				return chainSupport, ccipReader
			},
			expResult: nil,
		},
		{
			name: "nil is returned when supportsDestChain errors",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(false, errors.New("some error"))
				ccipReader := reader_mock.NewMockCCIPReader(t)
				return chainSupport, ccipReader
			},
			expResult: nil,
		},
		{
			name: "nil is returned when knownSourceChains errors",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(nil, errors.New("some error"))
				ccipReader := reader_mock.NewMockCCIPReader(t)
				return chainSupport, ccipReader
			},
			expResult: nil,
		},
		{
			name: "nextSeqNums returns incorrect number of seq nums, other chains should be processed correctly",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)
				// return a smaller slice, should trigger validation condition

				nextSeqNumsCp := maps.Clone(nextSeqNums)
				delete(nextSeqNumsCp, cciptypes.ChainSelector(4))

				ccipReader.EXPECT().NextSeqNum(mock.Anything, knownSourceChains).Return(nextSeqNumsCp, nil)
				ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).
					Return(reader.CurseInfo{}, nil)
				return chainSupport, ccipReader
			},
			expResult: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(7, 608),
				plugintypes.NewSeqNumChain(19, 7713),
			},
		},
		{
			name: "dest chain is cursed sequence numbers not observed",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)
				ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).Return(reader.CurseInfo{
					CursedSourceChains: nil,
					CursedDestination:  true,
					GlobalCurse:        false,
				}, nil)
				return chainSupport, ccipReader
			},
		},
		{
			name: "global curse sequence numbers not observed",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)
				ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).Return(reader.CurseInfo{
					CursedSourceChains: nil,
					CursedDestination:  false,
					GlobalCurse:        true,
				}, nil)
				return chainSupport, ccipReader
			},
		},
		{
			name: "one source chain is cursed sequence numbers not observed for that chain",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				knownSourceChains := []cciptypes.ChainSelector{4, 7, 19}
				cursedSourceChains := map[cciptypes.ChainSelector]bool{7: true, 4: false}
				knownSourceChainsExcludingCursed := []cciptypes.ChainSelector{4, 19}
				nextSeqNumsExcludingCursed := map[cciptypes.ChainSelector]cciptypes.SeqNum{4: 345, 19: 7713}

				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)

				ccipReader.EXPECT().NextSeqNum(mock.Anything, knownSourceChainsExcludingCursed).
					Return(nextSeqNumsExcludingCursed, nil)

				ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).Return(reader.CurseInfo{
					CursedSourceChains: cursedSourceChains,
					CursedDestination:  false,
					GlobalCurse:        false,
				}, nil)
				return chainSupport, ccipReader
			},
			expResult: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(4, 345),
				plugintypes.NewSeqNumChain(19, 7713),
			},
		},
		{
			name: "all source chains are cursed",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				knownSourceChains := []cciptypes.ChainSelector{4, 7, 19}
				cursedSourceChains := map[cciptypes.ChainSelector]bool{7: true, 4: true, 19: true}

				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)

				ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).Return(reader.CurseInfo{
					CursedSourceChains: cursedSourceChains,
					CursedDestination:  false,
					GlobalCurse:        false,
				}, nil)
				return chainSupport, ccipReader
			},
		},
		{
			name: "ccip reader error while fetching curse info",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)

				ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).
					Return(reader.CurseInfo{}, fmt.Errorf("some error"))
				return chainSupport, ccipReader
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tests.Context(t)

			chainSupport, ccipReader := tc.getDeps(t)
			defer chainSupport.AssertExpectations(t)
			defer ccipReader.AssertExpectations(t)

			o := newObserverImpl(
				logger.Test(t),
				nil,
				nodeID,
				chainSupport,
				ccipReader,
				mocks.NewMessageHasher(),
			)

			assert.Equal(t, tc.expResult, o.ObserveOffRampNextSeqNums(ctx))
		})
	}
}

func Test_ObserveMerkleRoots(t *testing.T) {
	testCases := []struct {
		name                     string
		ranges                   []plugintypes.ChainRange
		supportedChains          mapset.Set[cciptypes.ChainSelector]
		supportedChainsFails     bool
		msgsBetweenSeqNums       map[cciptypes.ChainSelector][]cciptypes.Message
		msgsBetweenSeqNumsErrors map[cciptypes.ChainSelector]error
		expMerkleRoots           map[cciptypes.ChainSelector]string
	}{
		{
			name: "Success single chain",
			ranges: []plugintypes.ChainRange{
				{
					ChainSel:    8,
					SeqNumRange: cciptypes.SeqNumRange{10, 11},
				},
			},
			supportedChains:      mapset.NewSet[cciptypes.ChainSelector](8),
			supportedChainsFails: false,
			msgsBetweenSeqNums: map[cciptypes.ChainSelector][]cciptypes.Message{
				8: {{
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1a"),
						SequenceNumber: 10},
				}, {
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1b"),
						SequenceNumber: 11},
				},
				},
			},
			msgsBetweenSeqNumsErrors: map[cciptypes.ChainSelector]error{},
			expMerkleRoots: map[cciptypes.ChainSelector]string{
				8: "5b81aaf37240df67f3ab0e845f30e29f35fdf9169e2517c436c1c0c11224c97b",
			},
		},
		{
			name: "Success multiple chains",
			ranges: []plugintypes.ChainRange{
				{
					ChainSel:    8,
					SeqNumRange: cciptypes.SeqNumRange{10, 11},
				},
				{
					ChainSel:    15,
					SeqNumRange: cciptypes.SeqNumRange{53, 55},
				},
			},
			supportedChains:      mapset.NewSet[cciptypes.ChainSelector](8, 15),
			supportedChainsFails: false,
			msgsBetweenSeqNums: map[cciptypes.ChainSelector][]cciptypes.Message{
				8: {{
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1a"),
						SequenceNumber: 10},
				}, {
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1b"),
						SequenceNumber: 11}},
				},
				15: {{
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x2a"),
						SequenceNumber: 53},
				}, {
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x2b"),
						SequenceNumber: 54},
				}, {
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x2c"),
						SequenceNumber: 55}},
				},
			},
			msgsBetweenSeqNumsErrors: map[cciptypes.ChainSelector]error{},
			expMerkleRoots: map[cciptypes.ChainSelector]string{
				8:  "5b81aaf37240df67f3ab0e845f30e29f35fdf9169e2517c436c1c0c11224c97b",
				15: "c7685b1be19745f244da890574cf554d75a3feeaf0e1181541c594d77ac1d6c4",
			},
		},
		{
			name: "Unsupported chain does not return a merkle root",
			ranges: []plugintypes.ChainRange{
				{
					ChainSel:    8,
					SeqNumRange: cciptypes.SeqNumRange{10, 11},
				},
				{
					// Unsupported chain
					ChainSel:    12,
					SeqNumRange: cciptypes.SeqNumRange{50, 60},
				},
			},
			supportedChains:      mapset.NewSet[cciptypes.ChainSelector](8),
			supportedChainsFails: false,
			msgsBetweenSeqNums: map[cciptypes.ChainSelector][]cciptypes.Message{
				8: {{
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1a"),
						SequenceNumber: 10},
				}, {
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1b"),
						SequenceNumber: 11},
				},
				},
			},
			msgsBetweenSeqNumsErrors: map[cciptypes.ChainSelector]error{},
			expMerkleRoots: map[cciptypes.ChainSelector]string{
				8: "5b81aaf37240df67f3ab0e845f30e29f35fdf9169e2517c436c1c0c11224c97b",
			},
		},
		{
			name: "Call to supportedChains fails",
			ranges: []plugintypes.ChainRange{
				{
					ChainSel:    8,
					SeqNumRange: cciptypes.SeqNumRange{10, 11},
				},
			},
			supportedChains:      mapset.NewSet[cciptypes.ChainSelector](8),
			supportedChainsFails: true,
			msgsBetweenSeqNums: map[cciptypes.ChainSelector][]cciptypes.Message{
				8: {{
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1a"),
						SequenceNumber: 10},
				}, {
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1b"),
						SequenceNumber: 11},
				},
				},
			},
			msgsBetweenSeqNumsErrors: map[cciptypes.ChainSelector]error{},
			expMerkleRoots:           nil,
		},
		{
			name: "msgsBetweenSeqNums fails for a chain",
			ranges: []plugintypes.ChainRange{
				{
					ChainSel:    8,
					SeqNumRange: cciptypes.SeqNumRange{10, 11},
				},
				{
					ChainSel:    12,
					SeqNumRange: cciptypes.SeqNumRange{50, 60},
				},
			},
			supportedChains:      mapset.NewSet[cciptypes.ChainSelector](8),
			supportedChainsFails: false,
			msgsBetweenSeqNums: map[cciptypes.ChainSelector][]cciptypes.Message{
				8: {{
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1a"),
						SequenceNumber: 10},
				}, {
					Header: cciptypes.RampMessageHeader{
						MessageID:      mustNewMessageID("0x1b"),
						SequenceNumber: 11}},
				},
				12: {},
			},
			msgsBetweenSeqNumsErrors: map[cciptypes.ChainSelector]error{
				12: fmt.Errorf("error"),
			},
			expMerkleRoots: map[cciptypes.ChainSelector]string{
				8: "5b81aaf37240df67f3ab0e845f30e29f35fdf9169e2517c436c1c0c11224c97b",
			},
		},
		{
			name: "multiple chains, some of them have missing messages within the range",
			ranges: []plugintypes.ChainRange{
				{ChainSel: 8, SeqNumRange: cciptypes.SeqNumRange{10, 11}},
				{ChainSel: 15, SeqNumRange: cciptypes.SeqNumRange{53, 55}},
				{ChainSel: 16, SeqNumRange: cciptypes.SeqNumRange{63, 65}},
				{ChainSel: 9, SeqNumRange: cciptypes.SeqNumRange{93, 95}},
				{ChainSel: 17, SeqNumRange: cciptypes.SeqNumRange{73, 75}},
				{ChainSel: 18, SeqNumRange: cciptypes.SeqNumRange{83, 85}},
			},
			supportedChains:      mapset.NewSet[cciptypes.ChainSelector](8, 15, 16, 9, 17, 18),
			supportedChainsFails: false,
			msgsBetweenSeqNums: map[cciptypes.ChainSelector][]cciptypes.Message{
				// 8: valid messages
				8: {{Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x1a"), SequenceNumber: 10}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x1b"), SequenceNumber: 11}}},
				// 15: missing middle message of the range
				15: {{Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x2a"), SequenceNumber: 53}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x2c"), SequenceNumber: 55}}},
				// 16: missing first message of the range
				16: {{Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x3a"), SequenceNumber: 64}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x3c"), SequenceNumber: 65}}},
				// 17: missing last message of the range
				17: {{Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x4a"), SequenceNumber: 73}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x4c"), SequenceNumber: 74}}},
				// 18: length of msgs is correct but sequence numbers are not
				18: {{Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x5a"), SequenceNumber: 84}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x5b"), SequenceNumber: 85}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0x5c"), SequenceNumber: 86}},
				},
				// 9: valid messages
				9: {{Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0xa1"), SequenceNumber: 93}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0xa2"), SequenceNumber: 94}}, {
					Header: cciptypes.RampMessageHeader{MessageID: mustNewMessageID("0xa3"), SequenceNumber: 95}}},
			},
			msgsBetweenSeqNumsErrors: map[cciptypes.ChainSelector]error{},
			expMerkleRoots: map[cciptypes.ChainSelector]string{
				8: "5b81aaf37240df67f3ab0e845f30e29f35fdf9169e2517c436c1c0c11224c97b",
				9: "f1b02d28559f60a67b431e2c580ac1d6b3e0fd7319ff055c6c67408aa31788e4",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			var nodeID commontypes.OracleID = 1
			mockCCIPReader := reader_mock.NewMockCCIPReader(t)
			for _, r := range tc.ranges {
				// Skip unexpected calls.
				if tc.supportedChainsFails || !tc.supportedChains.Contains(r.ChainSel) {
					continue
				}

				var err error
				if e, exists := tc.msgsBetweenSeqNumsErrors[r.ChainSel]; exists {
					err = e
				}
				mockCCIPReader.On(
					"MsgsBetweenSeqNums", ctx, r.ChainSel, r.SeqNumRange,
				).Return(tc.msgsBetweenSeqNums[r.ChainSel], err)
			}

			mockCCIPReader.EXPECT().
				GetContractAddress(mock.Anything, mock.Anything).
				Return(cciptypes.Bytes{}, nil).Maybe()

			chainSupport := common_mock.NewMockChainSupport(t)
			if tc.supportedChainsFails {
				chainSupport.On("SupportedChains", nodeID).Return(
					mapset.NewSet[cciptypes.ChainSelector](), fmt.Errorf("error"),
				)
			} else {
				chainSupport.On("SupportedChains", nodeID).Return(tc.supportedChains, nil)
			}

			o := newObserverImpl(
				logger.Test(t),
				nil,
				nodeID,
				chainSupport,
				mockCCIPReader,
				mocks.NewMessageHasher(),
			)

			roots := o.ObserveMerkleRoots(ctx, tc.ranges)
			if tc.expMerkleRoots == nil {
				assert.Nil(t, roots)
			} else {
				assert.Len(t, roots, len(tc.expMerkleRoots))
				for _, root := range roots {
					assert.Equal(t, tc.expMerkleRoots[root.ChainSel], hex.EncodeToString(root.MerkleRoot[:]))
				}
			}
		})
	}
}

func Test_computeMerkleRoot(t *testing.T) {
	testCases := []struct {
		name           string
		messageHeaders []cciptypes.RampMessageHeader
		messageHasher  cciptypes.MessageHasher
		expMerkleRoot  string
		expErr         bool
	}{
		{
			name: "Single message success",
			messageHeaders: []cciptypes.RampMessageHeader{
				{
					MessageID:      mustNewMessageID("0x1a"),
					SequenceNumber: 112,
				}},
			messageHasher: mocks.NewMessageHasher(),
			expMerkleRoot: "1a00000000000000000000000000000000000000000000000000000000000000",
			expErr:        false,
		},
		{
			name: "Multiple messages success",
			messageHeaders: []cciptypes.RampMessageHeader{
				{
					MessageID:      mustNewMessageID("0x1a"),
					SequenceNumber: 112,
				},
				{
					MessageID:      mustNewMessageID("0x23"),
					SequenceNumber: 113,
				},
				{
					MessageID:      mustNewMessageID("0x87"),
					SequenceNumber: 114,
				}},
			messageHasher: mocks.NewMessageHasher(),
			expMerkleRoot: "94c7e711e6f2acf41dca598ced55b6925e55aaed83520dc5ea6cbc054344564b",
			expErr:        false,
		},
		{
			name: "Sequence number gap",
			messageHeaders: []cciptypes.RampMessageHeader{
				{
					MessageID:      mustNewMessageID("0x10"),
					SequenceNumber: 34,
				},
				{
					MessageID:      mustNewMessageID("0x12"),
					SequenceNumber: 36,
				}},
			messageHasher: mocks.NewMessageHasher(),
			expMerkleRoot: "",
			expErr:        true,
		},
		{
			name:           "Empty messages",
			messageHeaders: []cciptypes.RampMessageHeader{},
			messageHasher:  mocks.NewMessageHasher(),
			expMerkleRoot:  "",
			expErr:         true,
		},
		{
			name: "Bad hasher",
			messageHeaders: []cciptypes.RampMessageHeader{
				{
					MessageID:      mustNewMessageID("0x1a"),
					SequenceNumber: 112,
				},
				{
					MessageID:      mustNewMessageID("0x23"),
					SequenceNumber: 113,
				},
				{
					MessageID:      mustNewMessageID("0x87"),
					SequenceNumber: 114,
				}},
			messageHasher: NewBadMessageHasher(),
			expMerkleRoot: "94c7e711e6f2acf41dca598ced55b6925e55aaed83520dc5ea6cbc054344564b",
			expErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := observerImpl{
				lggr:      logger.Test(t),
				msgHasher: tc.messageHasher,
			}

			msgs := make([]cciptypes.Message, 0)
			for _, h := range tc.messageHeaders {
				msgs = append(msgs, cciptypes.Message{Header: h})
			}

			rootBytes, err := p.computeMerkleRoot(context.Background(), p.lggr, msgs)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			rootString := hex.EncodeToString(rootBytes[:])
			assert.Equal(t, tc.expMerkleRoot, rootString)
		})
	}
}

func Test_Processor_initializeRMNController(t *testing.T) {
	ctx := tests.Context(t)

	p := &Processor{
		lggr:        logger.Test(t),
		offchainCfg: pluginconfig.CommitOffchainConfig{RMNEnabled: false},
	}

	err := p.prepareRMNController(ctx, p.lggr, Outcome{})
	assert.NoError(t, err, "rmn is not enabled")

	p.offchainCfg.RMNEnabled = true
	p.rmnControllerCfgDigest = cciptypes.Bytes32{1}
	err = p.prepareRMNController(ctx, p.lggr, Outcome{})
	assert.NoError(t, err, "rmn enabled but controller already initialized")

	p.rmnControllerCfgDigest = cciptypes.Bytes32{1}
	err = p.prepareRMNController(ctx, p.lggr, Outcome{})
	assert.NoError(t, err, "previous outcome does not contain remote config digest")

	rmnHomeReader := readerpkg_mock.NewMockRMNHome(t)
	rmnController := rmn_mock.NewMockController(t)
	p.rmnHomeReader = rmnHomeReader
	p.rmnController = rmnController

	cfg := testhelpers.CreateRMNRemoteCfg()
	rmnNodes := []rmntypes.HomeNodeInfo{
		{ID: 1, PeerID: types.PeerID{1, 2, 3}},
		{ID: 10, PeerID: types.PeerID{1, 2, 31}},
	}
	oracleIDs := []ragep2ptypes.PeerID{}
	rmnHomeReader.EXPECT().GetRMNNodesInfo(cfg.ConfigDigest).Return(rmnNodes, nil)

	rmnController.EXPECT().InitConnection(
		ctx,
		cciptypes.Bytes32(p.reportingCfg.ConfigDigest),
		cfg.ConfigDigest,
		oracleIDs,
		rmnNodes,
	).Return(nil)

	err = p.prepareRMNController(ctx, p.lggr, Outcome{RMNRemoteCfg: cfg})
	assert.NoError(t, err, "rmn controller initialized")
	assert.Equal(t, cfg.ConfigDigest, p.rmnControllerCfgDigest)
}

func Test_Processor_ObservationQuorum(t *testing.T) {
	testCases := []struct {
		name                      string
		numOracles                int
		bigF                      int
		numAttributedObservations int
		expectedQuorum            bool
		expErr                    bool
	}{
		{
			name:                      "all empty no quorum",
			numOracles:                0,
			bigF:                      0,
			numAttributedObservations: 0,
			expectedQuorum:            false,
			expErr:                    false,
		},
		{
			name:                      "happy path 2F+1 observations",
			numOracles:                8,
			bigF:                      3,
			numAttributedObservations: 7,
			expectedQuorum:            true,
			expErr:                    false,
		},
		{
			name:                      "no quorum path less than 2F+1 observations",
			numOracles:                8,
			bigF:                      3,
			numAttributedObservations: 6,
			expectedQuorum:            false,
			expErr:                    false,
		},
		{
			name:                      "zero observations case",
			numOracles:                8,
			bigF:                      3,
			numAttributedObservations: 0,
			expectedQuorum:            false,
			expErr:                    false,
		},
		{
			name:                      "even with zero oracles quorum not affected",
			numOracles:                0,
			bigF:                      3,
			numAttributedObservations: 7,
			expectedQuorum:            true,
			expErr:                    false,
		},
	}

	ctx := tests.Context(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := &Processor{
				lggr: logger.Test(t),
				reportingCfg: ocr3types.ReportingPluginConfig{
					N: tc.numOracles,
					F: tc.bigF,
				},
			}

			quorum, err := p.ObservationQuorum(
				ctx,
				ocr3types.OutcomeContext{},
				types2.Query{},
				make([]types2.AttributedObservation, tc.numAttributedObservations),
			)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedQuorum, quorum)
		})
	}
}

func Test_shouldSkipRMNVerification(t *testing.T) {
	testCases := []struct {
		name                       string
		nextProcessorState         processorState
		queryContainsRmnSigs       bool
		queryIndicatesSigsRetrying bool
		rmnRemoteConfigEmpty       bool
		expErr                     bool
		expSkip                    bool
	}{
		{
			name:    "all empty should skip rmn verification",
			expSkip: true,
		},
		{
			name:                 "happy path proceed with verification",
			nextProcessorState:   buildingReport,
			queryContainsRmnSigs: true,
		},
		{
			name:               "rmn sigs missing but error is not expected since chains might be rmn-disabled",
			nextProcessorState: buildingReport,
			expErr:             false,
		},
		{
			name:                       "rmn sigs are present while we retry sigs in the next round this is invalid",
			nextProcessorState:         buildingReport,
			queryContainsRmnSigs:       true,
			queryIndicatesSigsRetrying: true,
			expErr:                     true,
		},
		{
			name:                       "retrying sigs in the next round sig verification should be skipped",
			nextProcessorState:         buildingReport,
			queryIndicatesSigsRetrying: true,
			expSkip:                    true,
		},
		{
			name:                 "rmn remote config from previous outcome is empty error is expected",
			nextProcessorState:   buildingReport,
			queryContainsRmnSigs: true,
			rmnRemoteConfigEmpty: true,
			expErr:               true,
		},
		{
			name:                 "signatures were provided but we are not in the right state",
			nextProcessorState:   selectingRangesForReport,
			queryContainsRmnSigs: true,
			expErr:               true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := Query{}

			if tc.queryContainsRmnSigs {
				q.RMNSignatures = &rmn.ReportSignatures{
					Signatures: make([]*serialization.EcdsaSignature, 1),
				}
			}

			if tc.queryIndicatesSigsRetrying {
				q.RetryRMNSignatures = true
			}

			prevOutcome := Outcome{}
			if !tc.rmnRemoteConfigEmpty {
				prevOutcome.RMNRemoteCfg = cciptypes.RemoteConfig{FSign: 1}
			}

			shouldSkip, err := shouldSkipRMNVerification(tc.nextProcessorState, q, prevOutcome)
			if tc.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expSkip, shouldSkip)
		})
	}
}

func mustNewMessageID(msgIDHex string) cciptypes.Bytes32 {
	msgID, err := cciptypes.NewBytes32FromString(msgIDHex)
	if err != nil {
		panic(err)
	}
	return msgID
}

type BadMessageHasher struct{}

func NewBadMessageHasher() *BadMessageHasher {
	return &BadMessageHasher{}
}

// Always returns an error
func (m *BadMessageHasher) Hash(ctx context.Context, msg cciptypes.Message) (cciptypes.Bytes32, error) {
	return cciptypes.Bytes32{}, fmt.Errorf("failed to hash")
}

// signatureVerifierAlwaysTrue is a signature verifier that always returns true.
type signatureVerifierAlwaysTrue struct{}

func (a signatureVerifierAlwaysTrue) Verify(_ ed25519.PublicKey, _, _ []byte) bool {
	return true
}

func (a signatureVerifierAlwaysTrue) VerifyReportSignatures(
	_ context.Context, _ []cciptypes.RMNECDSASignature, _ cciptypes.RMNReport, _ []cciptypes.UnknownAddress) error {
	return nil
}
