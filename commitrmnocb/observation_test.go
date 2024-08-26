package commitrmnocb

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"

	observer_mock "github.com/smartcontractkit/chainlink-ccip/mocks/commitrmnocb"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

var (
	merkleRoots = []cciptypes.MerkleRootChain{
		{
			ChainSel:     1,
			SeqNumsRange: [2]cciptypes.SeqNum{5, 78},
			MerkleRoot:   [32]byte{1},
		},
	}

	gasPrices = []cciptypes.GasPriceChain{
		{
			GasPrice: cciptypes.NewBigIntFromInt64(99),
			ChainSel: 8,
		},
	}

	tokenPrices = []cciptypes.TokenPrice{
		{
			TokenID: "token23",
			Price:   cciptypes.NewBigIntFromInt64(80761),
		},
	}

	feedPrices = []cciptypes.TokenPrice{
		{
			TokenID: "token23",
			Price:   cciptypes.NewBigIntFromInt64(80761),
		},
	}

	registryUpdates = []cciptypes.TokenPrice{
		{
			TokenID: "token23",
			Price:   cciptypes.NewBigIntFromInt64(80761),
		},
	}

	offRampNextSeqNums = []plugintypes.SeqNumChain{
		{
			ChainSel: 456,
			SeqNum:   9987,
		},
	}

	fChain = map[cciptypes.ChainSelector]int{
		872: 3,
	}
)

func Test_Observation_SelectingRangesForReport(t *testing.T) {
	ctx := context.Background()

	observer := observer_mock.NewMockObserver(t)
	observer.On(
		"ObserveOffRampNextSeqNums", ctx,
	).Return(offRampNextSeqNums)
	observer.On("ObserveFChain").Return(fChain)

	previousOutcome := Outcome{
		OutcomeType: ReportTransmitted,
	}

	actualObs := observe(t, ctx, observer, previousOutcome)

	// We don't need to worry about comparing timestamps
	// Observation will always return the current time
	expectedObs := Observation{
		OnRampMaxSeqNums:   offRampNextSeqNums,
		OffRampNextSeqNums: offRampNextSeqNums,
		FChain:             fChain,
		Timestamp:          actualObs.Timestamp,
	}

	assert.Equal(t, expectedObs, actualObs)
}

func Test_Observation_BuildingReport(t *testing.T) {
	ctx := context.Background()

	observer := observer_mock.NewMockObserver(t)
	observer.On(
		"ObserveMerkleRoots", ctx, mock.Anything,
	).Return(merkleRoots)
	observer.On(
		"ObserveFeedTokenPrices", ctx,
	).Return(tokenPrices)
	observer.On(
		"ObservePriceRegistryTokenUpdates", ctx,
	).Return(registryUpdates)
	observer.On(
		"ObserveGasPrices", ctx,
	).Return(gasPrices)
	observer.On("ObserveFChain").Return(fChain)

	previousOutcome := Outcome{
		OutcomeType: ReportIntervalsSelected,
	}

	actualObs := observe(t, ctx, observer, previousOutcome)

	expectedObs := Observation{
		MerkleRoots:               merkleRoots,
		GasPrices:                 gasPrices,
		FeedTokenPrices:           feedPrices,
		PriceRegistryTokenUpdates: registryUpdates,
		FChain:                    fChain,
		Timestamp:                 actualObs.Timestamp,
	}

	assert.Equal(t, expectedObs, actualObs)
}

func Test_Observation_WaitingForReportTransmission(t *testing.T) {
	ctx := context.Background()

	observer := observer_mock.NewMockObserver(t)
	observer.On(
		"ObserveOffRampNextSeqNums", ctx,
	).Return(offRampNextSeqNums)
	observer.On("ObserveFChain").Return(fChain)

	previousOutcome := Outcome{
		OutcomeType: ReportInFlight,
	}

	actualObs := observe(t, ctx, observer, previousOutcome)
	expectedObs := Observation{
		OffRampNextSeqNums: offRampNextSeqNums,
		FChain:             fChain,
		Timestamp:          actualObs.Timestamp,
	}

	assert.Equal(t, expectedObs, actualObs)
}

func observe(t *testing.T, ctx context.Context, observer *observer_mock.MockObserver, previousOutcome Outcome) Observation {
	p := Plugin{
		lggr:     logger.Test(t),
		observer: observer,
	}

	previousOutcomeEncoded, err := previousOutcome.Encode()
	assert.NoError(t, err)

	result, err := p.Observation(
		ctx,
		ocr3types.OutcomeContext{PreviousOutcome: previousOutcomeEncoded},
		types.Query{},
	)
	assert.NoError(t, err)

	actualObs, err := DecodeCommitPluginObservation(result)
	assert.NoError(t, err)

	return actualObs
}

// Additional test functions can be written similarly by reusing the global variables.

func Test_ObserveOffRampNextSeqNums(t *testing.T) {
	testCases := []struct {
		name                   string
		supportsDestChain      bool
		supportsDestChainError error
		knownSourceChains      []cciptypes.ChainSelector
		knownSourceChainsError error
		nextSeqNums            []cciptypes.SeqNum
		nextSeqNumsError       error
		expResult              []plugintypes.SeqNumChain
	}{
		{
			name:                   "Happy path",
			supportsDestChain:      true,
			supportsDestChainError: nil,
			knownSourceChains:      []cciptypes.ChainSelector{4, 7, 19},
			knownSourceChainsError: nil,
			nextSeqNums:            []cciptypes.SeqNum{345, 608, 7713},
			nextSeqNumsError:       nil,
			expResult: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(4, 345),
				plugintypes.NewSeqNumChain(7, 608),
				plugintypes.NewSeqNumChain(19, 7713),
			},
		},
		{
			name:                   "nil is returned when supportsDestChain is false",
			supportsDestChain:      false,
			supportsDestChainError: nil,
			knownSourceChains:      []cciptypes.ChainSelector{4, 7, 19},
			knownSourceChainsError: nil,
			nextSeqNums:            []cciptypes.SeqNum{345, 608, 7713},
			nextSeqNumsError:       nil,
			expResult:              nil,
		},
		{
			name:                   "nil is returned when supportsDestChain errors",
			supportsDestChain:      true,
			supportsDestChainError: fmt.Errorf("error"),
			knownSourceChains:      []cciptypes.ChainSelector{4, 7, 19},
			knownSourceChainsError: nil,
			nextSeqNums:            []cciptypes.SeqNum{345, 608, 7713},
			nextSeqNumsError:       nil,
			expResult:              nil,
		},
		{
			name:                   "nil is returned when knownSourceChains errors",
			supportsDestChain:      true,
			supportsDestChainError: nil,
			knownSourceChains:      []cciptypes.ChainSelector{4, 7, 19},
			knownSourceChainsError: fmt.Errorf("error"),
			nextSeqNums:            []cciptypes.SeqNum{345, 608, 7713},
			nextSeqNumsError:       nil,
			expResult:              nil,
		},
		{
			name:                   "nil is returned when nextSeqNums returns incorrect number of seq nums",
			supportsDestChain:      true,
			supportsDestChainError: nil,
			knownSourceChains:      []cciptypes.ChainSelector{4, 7, 19},
			knownSourceChainsError: nil,
			nextSeqNums:            []cciptypes.SeqNum{345, 608},
			nextSeqNumsError:       nil,
			expResult:              nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			var nodeID commontypes.OracleID = 1
			reader := reader_mock.NewMockCCIP(t)
			if tc.supportsDestChain && tc.supportsDestChainError == nil && tc.knownSourceChainsError == nil {
				reader.On(
					"NextSeqNum", ctx, tc.knownSourceChains,
				).Return(tc.nextSeqNums, tc.nextSeqNumsError)
			}

			chainSupport := mocks.NewChainSupport()
			chainSupport.On(
				"SupportsDestChain", nodeID,
			).Return(tc.supportsDestChain, tc.supportsDestChainError)
			chainSupport.On(
				"KnownSourceChainsSlice",
			).Return(tc.knownSourceChains, tc.knownSourceChainsError)

			o := ObserverImpl{
				nodeID:       nodeID,
				lggr:         logger.Test(t),
				msgHasher:    mocks.NewMessageHasher(),
				ccipReader:   reader,
				chainSupport: chainSupport,
			}

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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			var nodeID commontypes.OracleID = 1
			reader := reader_mock.NewMockCCIP(t)
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

			chainSupport := mocks.NewChainSupport()
			if tc.supportedChainsFails {
				chainSupport.On("SupportedChains", nodeID).Return(
					mapset.NewSet[cciptypes.ChainSelector](), fmt.Errorf("error"),
				)
			} else {
				chainSupport.On("SupportedChains", nodeID).Return(tc.supportedChains, nil)
			}

			o := ObserverImpl{
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
			p := ObserverImpl{
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
