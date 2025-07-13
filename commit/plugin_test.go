package commit

import (
	"errors"
	"math/big"
	"testing"
	"time"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	reader2 "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	v1 "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_queryPhaseRmnRelatedTimers(t *testing.T) {
	testCases := []struct {
		maxQueryDuration           time.Duration
		expInitialObservationTimer time.Duration
		expInitialReportTimer      time.Duration
	}{
		{
			maxQueryDuration:           1 * time.Second,
			expInitialObservationTimer: 550 * time.Millisecond,
			expInitialReportTimer:      200 * time.Millisecond,
		},
		{
			maxQueryDuration:           3 * time.Second,
			expInitialObservationTimer: 1650 * time.Millisecond,
			expInitialReportTimer:      600 * time.Millisecond,
		},
		{
			maxQueryDuration:           5 * time.Second,
			expInitialObservationTimer: 2750 * time.Millisecond,
			expInitialReportTimer:      1000 * time.Millisecond,
		},
		{
			maxQueryDuration:           15 * time.Second,
			expInitialObservationTimer: 3 * time.Second,
			expInitialReportTimer:      2 * time.Second,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.maxQueryDuration.String(), func(t *testing.T) {
			obsTimer := observationsInitialRequestTimerDuration(tc.maxQueryDuration).Round(time.Millisecond)
			sigTimer := reportsInitialRequestTimerDuration(tc.maxQueryDuration).Round(time.Millisecond)
			assert.Equal(t, tc.expInitialObservationTimer, obsTimer)
			assert.Equal(t, tc.expInitialReportTimer, sigTimer)
		})
	}
}

func TestObservation_prices(t *testing.T) {
	destChainSel := chainsel.ABSTRACT_TESTNET.Selector
	fDestChain := 2
	testCases := []struct {
		name                        string
		oracleSupportsDestChain     bool
		currentRoundOcrSeqNum       uint64
		inflightOcrSeqNum           uint64
		remainingChecks             int
		onchainOcrSeqNum            uint64
		rpcErr                      error
		expObservedOnChainOcrSeqNum uint64
		expObservedPrices           bool
	}{
		{
			name:                        "no inflight prices, no onchain prices, initial state",
			oracleSupportsDestChain:     true,
			currentRoundOcrSeqNum:       1,
			inflightOcrSeqNum:           0,
			remainingChecks:             0,
			onchainOcrSeqNum:            0,
			expObservedOnChainOcrSeqNum: 0,
			expObservedPrices:           false,
		},
		{
			name:                        "initial prices observed on this round",
			oracleSupportsDestChain:     true,
			currentRoundOcrSeqNum:       2,
			inflightOcrSeqNum:           0,
			remainingChecks:             0,
			onchainOcrSeqNum:            0,
			expObservedOnChainOcrSeqNum: 0,
			expObservedPrices:           true,
		},
		{
			name:                        "inflight prices from previous round not seen on-chain",
			oracleSupportsDestChain:     true,
			currentRoundOcrSeqNum:       3,
			inflightOcrSeqNum:           2,
			remainingChecks:             10,
			onchainOcrSeqNum:            0,
			expObservedOnChainOcrSeqNum: 0,
			expObservedPrices:           false,
		},
		{
			name:                        "inflight prices from previous round are seen on-chain",
			oracleSupportsDestChain:     true,
			currentRoundOcrSeqNum:       5,
			inflightOcrSeqNum:           2,
			remainingChecks:             10,
			onchainOcrSeqNum:            2, // <--
			expObservedOnChainOcrSeqNum: 2,
			expObservedPrices:           false,
		},
		{
			name:                        "inflight prices seen on-chain but oracle des not support dest",
			oracleSupportsDestChain:     false,
			currentRoundOcrSeqNum:       5,
			inflightOcrSeqNum:           2,
			remainingChecks:             10,
			onchainOcrSeqNum:            2,
			expObservedOnChainOcrSeqNum: 0, // <-- observes 0
			expObservedPrices:           false,
		},
		{
			name:                        "no inflight prices, no onchain prices, later state",
			oracleSupportsDestChain:     true,
			currentRoundOcrSeqNum:       5678,
			inflightOcrSeqNum:           0,
			remainingChecks:             0,
			onchainOcrSeqNum:            5600,
			expObservedOnChainOcrSeqNum: 0,
			expObservedPrices:           true,
		},
		{
			name:                        "inflight prices, later state",
			oracleSupportsDestChain:     true,
			currentRoundOcrSeqNum:       5678,
			inflightOcrSeqNum:           5670, // <--
			remainingChecks:             50,
			onchainOcrSeqNum:            5600,
			expObservedOnChainOcrSeqNum: 5600,
			expObservedPrices:           false,
		},
		{
			name:                        "rpc error while getting on chain seq num",
			oracleSupportsDestChain:     true,
			currentRoundOcrSeqNum:       5678,
			inflightOcrSeqNum:           5670,
			remainingChecks:             50,
			onchainOcrSeqNum:            5600,
			rpcErr:                      errors.New("some rpc error"),
			expObservedOnChainOcrSeqNum: 0, // <----
			expObservedPrices:           false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup mock dependencies
			ctx := t.Context()
			cs := plugincommon.NewMockChainSupport(t)
			ccr := reader.NewMockCCIPReader(t)
			mrp := plugincommon.NewMockPluginProcessor[merkleroot.Query, merkleroot.Observation, merkleroot.Outcome](t)
			tpp := plugincommon.NewMockPluginProcessor[tokenprice.Query, tokenprice.Observation, tokenprice.Outcome](t)
			cfp := plugincommon.NewMockPluginProcessor[chainfee.Query, chainfee.Observation, chainfee.Outcome](t)
			home := reader2.NewMockHomeChain(t)

			// initialize plugin
			p := &Plugin{
				lggr:                logger.Test(t),
				ocrTypeCodec:        v1.NewCommitCodecProto(),
				merkleRootProcessor: mrp,
				tokenPriceProcessor: tpp,
				chainFeeProcessor:   cfp,
				metricsReporter:     &metrics.Noop{},
				chainSupport:        cs,
				ccipReader:          ccr,
				destChain:           ccipocr3.ChainSelector(destChainSel),
				homeChain:           home,
				oracleID:            commontypes.OracleID(9),
				offchainCfg: pluginconfig.CommitOffchainConfig{
					DonBreakingChangesVersion: pluginconfig.DonBreakingChangesVersion1RoleDonSupport,
				},
			}

			// set expectations
			home.EXPECT().GetFChain().Return(map[ccipocr3.ChainSelector]int{
				ccipocr3.ChainSelector(destChainSel): fDestChain,
			}, nil)

			cs.EXPECT().SupportsDestChain(p.oracleID).Return(tc.oracleSupportsDestChain, nil).Maybe()

			mrp.EXPECT().Observation(mock.Anything, mock.Anything, mock.Anything).Return(merkleroot.Observation{}, nil)

			ccr.EXPECT().GetLatestPriceSeqNr(mock.Anything).Return(tc.onchainOcrSeqNum, tc.rpcErr).Maybe()

			tokenPriceObs := tokenprice.Observation{}
			if tc.expObservedPrices {
				tokenPriceObs.FeedTokenPrices = map[ccipocr3.UnknownEncodedAddress]ccipocr3.BigInt{
					"123": ccipocr3.NewBigIntFromInt64(2),
				}
			}
			tpp.EXPECT().Observation(mock.Anything, mock.Anything, mock.Anything).Return(tokenPriceObs, nil).Maybe()

			chainFeeObs := chainfee.Observation{}
			if tc.expObservedPrices {
				chainFeeObs.ChainFeeUpdates = map[ccipocr3.ChainSelector]chainfee.Update{
					ccipocr3.ChainSelector(destChainSel): {
						Timestamp: time.Now(),
						ChainFee: chainfee.ComponentsUSDPrices{
							ExecutionFeePriceUSD: big.NewInt(2),
							DataAvFeePriceUSD:    big.NewInt(3),
						},
					},
				}
			}
			cfp.EXPECT().Observation(mock.Anything, mock.Anything, mock.Anything).Return(chainFeeObs, nil).Maybe()

			// encode previous outcome and call observation function
			prevOutcome, err := p.ocrTypeCodec.EncodeOutcome(committypes.Outcome{
				MainOutcome: committypes.MainOutcome{
					InflightPriceOcrSequenceNumber: ccipocr3.SeqNum(tc.inflightOcrSeqNum),
					RemainingPriceChecks:           tc.remainingChecks,
				},
			})
			require.NoError(t, err)

			obsBytes, err := p.Observation(ctx, ocr3types.OutcomeContext{
				SeqNr:           tc.currentRoundOcrSeqNum,
				PreviousOutcome: prevOutcome,
			}, types.Query{})
			require.NoError(t, err)

			// assert results are the expected
			obs, err := p.ocrTypeCodec.DecodeObservation(obsBytes)
			require.NoError(t, err)

			require.Equal(t, tc.expObservedOnChainOcrSeqNum, obs.OnChainPriceOcrSeqNum)
			if tc.expObservedPrices {
				require.NotEmpty(t, obs.TokenPriceObs.FeedTokenPrices)
				require.NotEmpty(t, obs.ChainFeeObs.ChainFeeUpdates)
			}
		})
	}
}

func Test_Outcome_prices(t *testing.T) {
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
			ctx := t.Context()
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
					DonBreakingChangesVersion: pluginconfig.DonBreakingChangesVersion1RoleDonSupport,
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
			outcomeBytes, err := p.Outcome(ctx, ocr3types.OutcomeContext{
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
