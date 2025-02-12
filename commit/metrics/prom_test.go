package metrics

import (
	"math/big"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	chainID  = "2337"
	selector = cciptypes.ChainSelector(12922642891491394802)
)

func Test_TrackingTokenPrices(t *testing.T) {
	tokenPricesProcessor := "tokenprices"
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(cleanupMetrics(reporter))

	obsTcs := []struct {
		name                   string
		observation            tokenprice.Observation
		expectedFeedToken      int
		expectedFeeQuotedToken int
	}{
		{
			name: "empty/missing structs should not report anything",
			observation: tokenprice.Observation{
				FeedTokenPrices:       nil,
				FeeQuoterTokenUpdates: nil,
			},
			expectedFeedToken:      0,
			expectedFeeQuotedToken: 0,
		},
		{
			name: "data is properly reported",
			observation: tokenprice.Observation{
				FeedTokenPrices: cciptypes.TokenPriceMap{
					cciptypes.UnknownEncodedAddress("0x123"): {},
					cciptypes.UnknownEncodedAddress("0x456"): {},
				},
				FeeQuoterTokenUpdates: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
					cciptypes.UnknownEncodedAddress("0x123"): {},
					cciptypes.UnknownEncodedAddress("0x456"): plugintypes.NewTimestampedBig(1, time.Now()),
					cciptypes.UnknownEncodedAddress("0x789"): plugintypes.NewTimestampedBig(2, time.Now()),
				},
			},
			expectedFeedToken:      2,
			expectedFeeQuotedToken: 3,
		},
	}

	for _, tc := range obsTcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackProcessorObservation(tokenPricesProcessor, tc.observation, nil)

			feedTokens := int(testutil.ToFloat64(
				reporter.processorObservationCounter.WithLabelValues(chainID, tokenPricesProcessor, "feedTokenPrices")),
			)
			require.Equal(t, tc.expectedFeedToken, feedTokens)
			feeQuoted := int(testutil.ToFloat64(
				reporter.processorObservationCounter.WithLabelValues(chainID, tokenPricesProcessor, "feeQuoterTokenUpdates")),
			)
			require.Equal(t, tc.expectedFeeQuotedToken, feeQuoted)
		})
	}

	outTcs := []struct {
		name                string
		outcome             tokenprice.Outcome
		expectedTokenPrices int
	}{
		{
			name: "empty/missing structs should not report anything",
			outcome: tokenprice.Outcome{
				TokenPrices: cciptypes.TokenPriceMap{},
			},
			expectedTokenPrices: 0,
		},
		{
			name: "null structs should not report anything",
			outcome: tokenprice.Outcome{
				TokenPrices: nil,
			},
			expectedTokenPrices: 0,
		},
		{
			name: "data is properly reported",
			outcome: tokenprice.Outcome{
				TokenPrices: cciptypes.TokenPriceMap{
					cciptypes.UnknownEncodedAddress("0x123"): cciptypes.NewBigIntFromInt64(1),
					cciptypes.UnknownEncodedAddress("0x234"): cciptypes.NewBigIntFromInt64(2),
					cciptypes.UnknownEncodedAddress("0x125"): cciptypes.NewBigIntFromInt64(3),
				},
			},
			expectedTokenPrices: 3,
		},
	}

	for _, tc := range outTcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackProcessorOutcome(tokenPricesProcessor, tc.outcome, err)

			tokenPrices := int(testutil.ToFloat64(
				reporter.processorOutcomeCounter.WithLabelValues(chainID, tokenPricesProcessor, "tokenPrices")),
			)
			require.Equal(t, tc.expectedTokenPrices, tokenPrices)
		})
	}
}

func Test_TrackingChainFees(t *testing.T) {
	chainFeeProcessor := "chainfee"
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(cleanupMetrics(reporter))

	obsTcs := []struct {
		name                    string
		observation             chainfee.Observation
		expectedFeeComponents   int
		expectedNativePrices    int
		expectedCHainFeeUpdates int
	}{
		{
			name: "empty/missing structs should not report anything",
			observation: chainfee.Observation{
				FeeComponents:     nil,
				NativeTokenPrices: nil,
				ChainFeeUpdates:   map[cciptypes.ChainSelector]chainfee.Update{},
			},
			expectedFeeComponents:   0,
			expectedNativePrices:    0,
			expectedCHainFeeUpdates: 0,
		},
		{
			name: "data is properly reported",
			observation: chainfee.Observation{
				FeeComponents: map[cciptypes.ChainSelector]types.ChainFeeComponents{
					cciptypes.ChainSelector(1): {},
				},
				NativeTokenPrices: map[cciptypes.ChainSelector]cciptypes.BigInt{
					cciptypes.ChainSelector(2): {},
				},
				ChainFeeUpdates: map[cciptypes.ChainSelector]chainfee.Update{
					cciptypes.ChainSelector(2): {},
					cciptypes.ChainSelector(3): {},
				},
			},
			expectedFeeComponents:   1,
			expectedNativePrices:    1,
			expectedCHainFeeUpdates: 2,
		},
	}

	for _, tc := range obsTcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackProcessorObservation(chainFeeProcessor, tc.observation, nil)

			feeComponents := int(testutil.ToFloat64(
				reporter.processorObservationCounter.WithLabelValues(chainID, chainFeeProcessor, "feeComponents")),
			)
			require.Equal(t, tc.expectedFeeComponents, feeComponents)
			nativePrices := int(testutil.ToFloat64(
				reporter.processorObservationCounter.WithLabelValues(chainID, chainFeeProcessor, "nativeTokenPrices")),
			)
			require.Equal(t, tc.expectedNativePrices, nativePrices)
			chainFeeUpdates := int(testutil.ToFloat64(
				reporter.processorObservationCounter.WithLabelValues(chainID, chainFeeProcessor, "chainFeeUpdates")),
			)
			require.Equal(t, tc.expectedCHainFeeUpdates, chainFeeUpdates)
		})
	}

	outTcs := []struct {
		name              string
		outcome           chainfee.Outcome
		expectedGasPrices int
	}{
		{
			name: "empty/missing structs should not report anything",
			outcome: chainfee.Outcome{
				GasPrices: nil,
			},
			expectedGasPrices: 0,
		},
		{
			name: "data is properly reported",
			outcome: chainfee.Outcome{
				GasPrices: []cciptypes.GasPriceChain{
					cciptypes.NewGasPriceChain(big.NewInt(2), cciptypes.ChainSelector(2)),
					cciptypes.NewGasPriceChain(big.NewInt(3), cciptypes.ChainSelector(2)),
					cciptypes.NewGasPriceChain(big.NewInt(4), cciptypes.ChainSelector(2)),
					cciptypes.NewGasPriceChain(big.NewInt(5), cciptypes.ChainSelector(2)),
				},
			},
			expectedGasPrices: 4,
		},
	}

	for _, tc := range outTcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackProcessorOutcome(chainFeeProcessor, tc.outcome, nil)

			gasPrices := int(testutil.ToFloat64(
				reporter.processorOutcomeCounter.WithLabelValues(chainID, chainFeeProcessor, "gasPrices")),
			)
			require.Equal(t, tc.expectedGasPrices, gasPrices)
		})
	}
}

func Test_MerkleRoots(t *testing.T) {
	processor := "merkleroot"
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(cleanupMetrics(reporter))

	obsTcs := []struct {
		name             string
		observation      merkleroot.Observation
		state            string
		expectedRoots    int
		expectedMessages int
	}{
		{
			name: "empty/missing structs should not report anything",
			observation: merkleroot.Observation{
				MerkleRoots: nil,
			},
			state:            "state",
			expectedRoots:    0,
			expectedMessages: 0,
		},
		{
			name: "data is properly reported",
			observation: merkleroot.Observation{
				MerkleRoots: []cciptypes.MerkleRootChain{
					{
						ChainSel:      cciptypes.ChainSelector(1),
						OnRampAddress: rand.RandomBytes(32),
						SeqNumsRange:  cciptypes.NewSeqNumRange(1, 10),
						MerkleRoot:    rand.RandomBytes32(),
					},
					{
						ChainSel:      cciptypes.ChainSelector(2),
						OnRampAddress: rand.RandomBytes(32),
						SeqNumsRange:  cciptypes.NewSeqNumRange(2, 3),
						MerkleRoot:    rand.RandomBytes32(),
					},
				},
			},
			expectedRoots:    2,
			expectedMessages: 12,
		},
	}

	for _, tc := range obsTcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackProcessorObservation(processor, tc.observation, nil)

			roots := int(testutil.ToFloat64(
				reporter.processorObservationCounter.WithLabelValues(chainID, processor, "roots")),
			)
			require.Equal(t, tc.expectedRoots, roots)
			messages := int(testutil.ToFloat64(
				reporter.processorObservationCounter.WithLabelValues(chainID, processor, "messages")),
			)
			require.Equal(t, tc.expectedMessages, messages)
		})
	}

	outTcs := []struct {
		name                  string
		outcome               merkleroot.Outcome
		state                 string
		expectedRoots         int
		expectedMessages      int
		expectedRMNSignatures int
	}{
		{
			name: "empty/missing structs should not report anything",
			outcome: merkleroot.Outcome{
				RootsToReport: nil,
			},
			state:                 "state",
			expectedRoots:         0,
			expectedMessages:      0,
			expectedRMNSignatures: 0,
		},
		{
			name: "data is properly reported",
			outcome: merkleroot.Outcome{
				RootsToReport: []cciptypes.MerkleRootChain{
					{
						ChainSel:      cciptypes.ChainSelector(1),
						OnRampAddress: rand.RandomBytes(32),
						SeqNumsRange:  cciptypes.NewSeqNumRange(1, 2),
						MerkleRoot:    rand.RandomBytes32(),
					},
					{
						ChainSel:      cciptypes.ChainSelector(2),
						OnRampAddress: rand.RandomBytes(32),
						SeqNumsRange:  cciptypes.NewSeqNumRange(2, 5),
						MerkleRoot:    rand.RandomBytes32(),
					},
				},
				RMNReportSignatures: make([]cciptypes.RMNECDSASignature, 5),
			},
			state:                 "state",
			expectedRoots:         2,
			expectedMessages:      6,
			expectedRMNSignatures: 5,
		},
	}

	for _, tc := range outTcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackProcessorOutcome(processor, tc.outcome, nil)

			roots := int(testutil.ToFloat64(
				reporter.processorOutcomeCounter.WithLabelValues(chainID, processor, "roots")),
			)
			require.Equal(t, tc.expectedRoots, roots)
			messages := int(testutil.ToFloat64(
				reporter.processorOutcomeCounter.WithLabelValues(chainID, processor, "messages")),
			)
			require.Equal(t, tc.expectedMessages, messages)
			rmns := int(testutil.ToFloat64(
				reporter.processorOutcomeCounter.WithLabelValues(chainID, processor, "rmnSignatures")),
			)
			require.Equal(t, tc.expectedRMNSignatures, rmns)
		})
	}
}

func cleanupMetrics(reporter *PromReporter) func() {
	return func() {
		reporter.processorOutcomeErrors.Reset()
		reporter.processorObservationErrors.Reset()
		reporter.processorOutcomeCounter.Reset()
		reporter.processorObservationCounter.Reset()
	}
}
