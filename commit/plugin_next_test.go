package commit

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	v1 "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_outcomeNext_prices(t *testing.T) {
	destChainSel := chainsel.ABSTRACT_TESTNET.Selector
	const priceInflightChecks = 10
	testCases := []struct {
		name                         string
		currentRoundOcrSeqNum        uint64
		inflightOcrSeqNum            uint64
		remainingChecks              int
		observedPrices               bool
		onchainOcrSeqNumObservations []uint64
		fDestChainObservations       []int
		fRoleDon                     int
		expMainOutcome               committypes.MainOutcome
	}{
		{
			name:                         "no prices reported no inflight prices exist",
			currentRoundOcrSeqNum:        1,
			inflightOcrSeqNum:            0,
			remainingChecks:              0,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{0, 0, 0, 0},
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 0,
				RemainingPriceChecks:           0,
			},
		},
		{
			name:                         "prices reported no inflight prices exist",
			currentRoundOcrSeqNum:        1,
			inflightOcrSeqNum:            0,
			remainingChecks:              0,
			observedPrices:               true,
			onchainOcrSeqNumObservations: []uint64{0, 0, 0, 0},
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 1,
				RemainingPriceChecks:           priceInflightChecks,
			},
		},
		{
			name:                         "prices reported while inflight prices exist (that should log an error)",
			currentRoundOcrSeqNum:        2,
			inflightOcrSeqNum:            1,
			remainingChecks:              10,
			observedPrices:               true,
			onchainOcrSeqNumObservations: []uint64{1, 1, 1, 1},
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 2,
				RemainingPriceChecks:           priceInflightChecks,
			},
		},
		{
			name:                         "prices still inflight",
			currentRoundOcrSeqNum:        123,
			inflightOcrSeqNum:            120,
			remainingChecks:              5,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{110, 110, 110, 110},
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 120,
				RemainingPriceChecks:           4,
			},
		},
		{
			name:                         "prices transmitted on this round",
			currentRoundOcrSeqNum:        123,
			inflightOcrSeqNum:            120,
			remainingChecks:              5,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{120, 120, 120, 120},
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 0,
				RemainingPriceChecks:           0,
			},
		},
		{
			name:                         "prices transmitted on this round f+1 oracles did not see this",
			currentRoundOcrSeqNum:        123,
			inflightOcrSeqNum:            120,
			remainingChecks:              5,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{110, 120, 110, 120}, // <----
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 120,
				RemainingPriceChecks:           4,
			},
		},
		{
			name:                         "prices transmitted on this round f oracles did not see this",
			currentRoundOcrSeqNum:        123,
			inflightOcrSeqNum:            120,
			remainingChecks:              5,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{110, 120, 120, 120}, // <----
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 0,
				RemainingPriceChecks:           0,
			},
		},
		{
			name:                         "prices still inflight but no 2F+1 fDestChain observations",
			currentRoundOcrSeqNum:        123,
			inflightOcrSeqNum:            120,
			remainingChecks:              5,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{110, 110, 110, 110},
			fDestChainObservations:       []int{1, 1, 2, 2}, // <--
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 0,
				RemainingPriceChecks:           0,
			},
		},
		{
			name:                         "prices still inflight but no 2F+1 fDestChain observations",
			currentRoundOcrSeqNum:        123,
			inflightOcrSeqNum:            120,
			remainingChecks:              5,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{110, 110, 110, 110},
			fDestChainObservations:       []int{1, 3, 2, 4}, // <--
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 0,
				RemainingPriceChecks:           0,
			},
		},
		{
			name:                         "not enough observed onchain ocr seq nums",
			currentRoundOcrSeqNum:        123,
			inflightOcrSeqNum:            120,
			remainingChecks:              5,
			observedPrices:               false,
			onchainOcrSeqNumObservations: []uint64{110, 0, 0, 0},
			fDestChainObservations:       []int{1, 1, 1, 1},
			fRoleDon:                     1,
			expMainOutcome: committypes.MainOutcome{
				InflightPriceOcrSequenceNumber: 0,
				RemainingPriceChecks:           0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup mock dependencies
			ctx := tests.Context(t)
			mrp := plugincommon.NewMockPluginProcessor[merkleroot.Query, merkleroot.Observation, merkleroot.Outcome](t)
			tpp := plugincommon.NewMockPluginProcessor[tokenprice.Query, tokenprice.Observation, tokenprice.Outcome](t)
			cfp := plugincommon.NewMockPluginProcessor[chainfee.Query, chainfee.Observation, chainfee.Outcome](t)

			// initialize plugin
			p := &Plugin{
				lggr:                logger.Test(t),
				ocrTypeCodec:        v1.NewCommitCodecProto(),
				merkleRootProcessor: mrp,
				tokenPriceProcessor: tpp,
				chainFeeProcessor:   cfp,
				metricsReporter:     &metrics.Noop{},
				destChain:           ccipocr3.ChainSelector(destChainSel),
				reportingCfg:        ocr3types.ReportingPluginConfig{F: tc.fRoleDon},
				offchainCfg: pluginconfig.CommitOffchainConfig{
					InflightPriceCheckRetries: priceInflightChecks,
				},
			}

			// define mock expectations
			mrp.EXPECT().Outcome(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(merkleroot.Outcome{}, nil)

			if tc.observedPrices {
				tpp.EXPECT().Outcome(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(tokenprice.Outcome{
						TokenPrices: map[ccipocr3.UnknownEncodedAddress]ccipocr3.BigInt{
							"123": ccipocr3.NewBigIntFromInt64(2),
						},
					}, nil)
			} else {
				tpp.EXPECT().Outcome(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(tokenprice.Outcome{}, nil)
			}

			cfp.EXPECT().Outcome(mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(chainfee.Outcome{}, nil)

			// define function inputs
			prevOutcome, err := p.ocrTypeCodec.EncodeOutcome(committypes.Outcome{
				MainOutcome: committypes.MainOutcome{
					InflightPriceOcrSequenceNumber: ccipocr3.SeqNum(tc.inflightOcrSeqNum),
					RemainingPriceChecks:           tc.remainingChecks,
				},
			})
			require.NoError(t, err)

			require.Equal(t, len(tc.fDestChainObservations), len(tc.onchainOcrSeqNumObservations), "wrong tc")
			aos := make([]types.AttributedObservation, 0, len(tc.fDestChainObservations))
			for i := range tc.fDestChainObservations {
				obs, err := p.ocrTypeCodec.EncodeObservation(committypes.Observation{
					FChain: map[ccipocr3.ChainSelector]int{
						ccipocr3.ChainSelector(destChainSel): tc.fDestChainObservations[i],
					},
					OnChainPriceOcrSeqNum: tc.onchainOcrSeqNumObservations[i],
				})
				require.NoError(t, err)
				aos = append(aos, types.AttributedObservation{Observation: obs, Observer: commontypes.OracleID(i + 1)})
			}

			// call the function
			outcomeBytes, err := p.outcomeNext(ctx, ocr3types.OutcomeContext{
				SeqNr:           tc.currentRoundOcrSeqNum,
				PreviousOutcome: prevOutcome,
			}, types.Query{}, aos)
			require.NoError(t, err)
			outcome, err := p.ocrTypeCodec.DecodeOutcome(outcomeBytes)
			require.NoError(t, err)

			// make assertions
			require.Equal(t, tc.expMainOutcome, outcome.MainOutcome)
		})
	}
}
