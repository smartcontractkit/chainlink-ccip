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
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func TestObservation(t *testing.T) {
	mockCCIPReader := readerpkg_mock.NewMockCCIPReader(t)
	chainSupport := common_mock.NewMockChainSupport(t)

	destChain := cciptypes.ChainSelector(909606746561742123)

	offchainAddress := []byte(rand.RandomAddress())

	p := &Processor{
		lggr:         logger.Test(t),
		rmnCrypto:    signatureVerifierAlwaysTrue{},
		ccipReader:   mockCCIPReader,
		destChain:    destChain,
		offchainCfg:  pluginconfig.CommitOffchainConfig{RMNEnabled: true},
		chainSupport: chainSupport,
	}

	ctx := context.Background()

	testCases := []struct {
		name        string
		prevOutcome Outcome
		query       Query
		setupMocks  func() *merkleroot.MockObserver
		expectedObs Observation
		expectedErr string
	}{
		{
			name: "SelectingRangesForReport",
			prevOutcome: Outcome{
				OutcomeType: ReportTransmitted,
			},
			query: Query{},
			setupMocks: func() *merkleroot.MockObserver {
				mockObserver := merkleroot.NewMockObserver(t)
				mockObserver.On("ObserveCursedChains", mock.Anything, destChain).Return([]cciptypes.ChainSelector{}, nil)
				mockObserver.On("ObserveOffRampNextSeqNums", mock.Anything, []cciptypes.ChainSelector{}).Return(
					[]plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 10}}).Once()
				mockObserver.On("ObserveLatestOnRampSeqNums", mock.Anything, destChain, []cciptypes.ChainSelector{}).Return(
					[]plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 15}})
				mockObserver.On("ObserveRMNRemoteCfg", mock.Anything, destChain).Return(rmntypes.RemoteConfig{})
				mockObserver.On("ObserveFChain").Return(map[cciptypes.ChainSelector]int{1: 3})
				return mockObserver
			},
			expectedObs: Observation{
				OffRampNextSeqNums: []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 10}},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 15}},
				RMNRemoteConfig:    rmntypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{1: 3},
			},
		},
		{
			name: "SelectingRangesForReportWithCursedChains",
			prevOutcome: Outcome{
				OutcomeType: ReportTransmitted,
			},
			query: Query{},
			setupMocks: func() *merkleroot.MockObserver {
				mockObserver := merkleroot.NewMockObserver(t)
				mockObserver.On("ObserveCursedChains", mock.Anything, destChain).Return([]cciptypes.ChainSelector{1}, nil)
				mockObserver.On("ObserveOffRampNextSeqNums", mock.Anything, []cciptypes.ChainSelector{1}).Return(
					[]plugintypes.SeqNumChain{}).Once()
				mockObserver.On("ObserveLatestOnRampSeqNums", mock.Anything, destChain, []cciptypes.ChainSelector{1}).Return(
					[]plugintypes.SeqNumChain{})
				mockObserver.On("ObserveRMNRemoteCfg", mock.Anything, destChain).Return(rmntypes.RemoteConfig{})
				mockObserver.On("ObserveFChain").Return(map[cciptypes.ChainSelector]int{1: 3})
				return mockObserver
			},
			expectedObs: Observation{
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{},
				RMNRemoteConfig:    rmntypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{1: 3},
			},
		},
		{
			name: "SelectingRangesForReportWithCursedDest",
			prevOutcome: Outcome{
				OutcomeType: ReportTransmitted,
			},
			query: Query{},
			setupMocks: func() *merkleroot.MockObserver {
				mockObserver := merkleroot.NewMockObserver(t)
				mockObserver.On("ObserveCursedChains", mock.Anything, destChain).
					Return([]cciptypes.ChainSelector{destChain}, nil)
				return mockObserver
			},
			expectedObs: Observation{},
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
				RMNSignatures: &rmn.ReportSignatures{},
			},
			setupMocks: func() *merkleroot.MockObserver {
				mockObserver := merkleroot.NewMockObserver(t)
				mockObserver.On("ObserveMerkleRoots", mock.Anything, mock.Anything).Return([]cciptypes.MerkleRootChain{
					{
						ChainSel:     1,
						SeqNumsRange: [2]cciptypes.SeqNum{5, 10},
						MerkleRoot:   [32]byte{1},
					}})
				mockObserver.On("ObserveFChain").Return(map[cciptypes.ChainSelector]int{1: 3})
				mockCCIPReader.On("GetContractAddress", mock.Anything, mock.Anything).Return(offchainAddress, nil)
				return mockObserver
			},
			expectedObs: Observation{
				MerkleRoots: []cciptypes.MerkleRootChain{
					{
						ChainSel:     1,
						SeqNumsRange: [2]cciptypes.SeqNum{5, 10},
						MerkleRoot:   [32]byte{1}},
				},
				FChain: map[cciptypes.ChainSelector]int{1: 3},
			},
		},
		{
			name: "WaitingForReportTransmission",
			prevOutcome: Outcome{
				OutcomeType:  ReportInFlight,
				RMNRemoteCfg: testhelpers.CreateRMNRemoteCfg(),
			},
			query: Query{},
			setupMocks: func() *merkleroot.MockObserver {
				mockObserver := merkleroot.NewMockObserver(t)
				mockObserver.On("ObserveOffRampNextSeqNums", mock.Anything, []cciptypes.ChainSelector{}).Return(
					[]plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 20}}).Once()
				mockObserver.On("ObserveFChain").Return(map[cciptypes.ChainSelector]int{1: 3})
				return mockObserver
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
			setupMocks: func() *merkleroot.MockObserver {
				// No mocks needed for this case
				return merkleroot.NewMockObserver(t)
			},
			expectedObs: Observation{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockObserver := tc.setupMocks()
			p.observer = mockObserver
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
	nextSeqNums := []cciptypes.SeqNum{345, 608, 7713}

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
			name: "nil is returned when nextSeqNums returns incorrect number of seq nums",
			getDeps: func(t *testing.T) (*common_mock.MockChainSupport, *reader_mock.MockCCIPReader) {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportsDestChain(nodeID).Return(true, nil)
				chainSupport.EXPECT().KnownSourceChainsSlice().Return(knownSourceChains, nil)
				ccipReader := reader_mock.NewMockCCIPReader(t)
				// return a smaller slice, should trigger validation condition
				ccipReader.EXPECT().NextSeqNum(mock.Anything, knownSourceChains).Return(nextSeqNums[1:], nil)
				return chainSupport, ccipReader
			},
			expResult: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tests.Context(t)

			chainSupport, ccipReader := tc.getDeps(t)
			defer chainSupport.AssertExpectations(t)
			defer ccipReader.AssertExpectations(t)

			o := observerImpl{
				nodeID:       nodeID,
				lggr:         logger.Test(t),
				msgHasher:    mocks.NewMessageHasher(),
				ccipReader:   ccipReader,
				chainSupport: chainSupport,
			}

			assert.Equal(t, tc.expResult, o.ObserveOffRampNextSeqNums(ctx, nil))
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			var nodeID commontypes.OracleID = 1
			reader := reader_mock.NewMockCCIPReader(t)
			for _, r := range tc.ranges {
				// Skip unexpected calls.
				if tc.supportedChainsFails || !tc.supportedChains.Contains(r.ChainSel) {
					continue
				}

				var err error
				if e, exists := tc.msgsBetweenSeqNumsErrors[r.ChainSel]; exists {
					err = e
				}
				reader.On(
					"MsgsBetweenSeqNums", ctx, r.ChainSel, r.SeqNumRange,
				).Return(tc.msgsBetweenSeqNums[r.ChainSel], err)
			}

			reader.EXPECT().
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

			o := observerImpl{
				nodeID:       nodeID,
				lggr:         logger.Test(t),
				msgHasher:    mocks.NewMessageHasher(),
				ccipReader:   reader,
				chainSupport: chainSupport,
			}

			roots := o.ObserveMerkleRoots(ctx, tc.ranges)
			if tc.expMerkleRoots == nil {
				assert.Nil(t, roots)
			} else {
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

			rootBytes, err := p.computeMerkleRoot(context.Background(), msgs)

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

	err := p.initializeRMNController(ctx, Outcome{})
	assert.NoError(t, err, "rmn is not enabled")

	p.offchainCfg.RMNEnabled = true
	p.rmnControllerCfgDigest = cciptypes.Bytes32{1}
	err = p.initializeRMNController(ctx, Outcome{})
	assert.NoError(t, err, "rmn enabled but controller already initialized")

	p.rmnControllerCfgDigest = cciptypes.Bytes32{1}
	err = p.initializeRMNController(ctx, Outcome{})
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

	err = p.initializeRMNController(ctx, Outcome{RMNRemoteCfg: cfg})
	assert.NoError(t, err, "rmn controller initialized")
	assert.Equal(t, cfg.ConfigDigest, p.rmnControllerCfgDigest)
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
