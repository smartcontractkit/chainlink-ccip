package metrics

import (
	"fmt"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	exectypes2 "github.com/smartcontractkit/chainlink-ccip/ocr3/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal/libs"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal/plugincommon"
)

const (
	chainID  = "EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG"
	selector = cciptypes.ChainSelector(16423721717087811551)
)

func Test_TrackingTokenReadiness(t *testing.T) {
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(cleanupMetrics(reporter))

	tcs := []struct {
		name                  string
		observation           exectypes2.TokenDataObservations
		state                 exectypes2.PluginState
		expectedReadyTokens   int
		expectedWaitingTokens int
	}{
		{
			name:                  "empty/missing structs should not report anything",
			observation:           exectypes2.TokenDataObservations{},
			state:                 exectypes2.GetMessages,
			expectedReadyTokens:   0,
			expectedWaitingTokens: 0,
		},
		{
			name: "single chain with some tokens ready and some waiting",
			observation: exectypes2.TokenDataObservations{
				cciptypes.ChainSelector(123): map[cciptypes.SeqNum]exectypes2.MessageTokenData{
					1: exectypes2.NewMessageTokenData(),
					2: exectypes2.NewMessageTokenData(exectypes2.NewSuccessTokenData([]byte("asdf"))),
				},
				cciptypes.ChainSelector(456): map[cciptypes.SeqNum]exectypes2.MessageTokenData{
					4: exectypes2.NewMessageTokenData(
						exectypes2.NewSuccessTokenData([]byte("asdf")),
						exectypes2.NewSuccessTokenData([]byte("qwer")),
						exectypes2.NewErrorTokenData(fmt.Errorf("error")),
					),
				},
			},
			expectedWaitingTokens: 1,
			expectedReadyTokens:   3,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackObservation(exectypes2.Observation{TokenData: tc.observation}, tc.state)

			readyTokens := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.ObservationMethod, string(tc.state), "tokenReady",
				),
			)
			require.Equal(t, tc.expectedReadyTokens, int(readyTokens))

			waitingTokens := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.ObservationMethod, string(tc.state), "tokenWaiting",
				),
			)
			require.Equal(t, tc.expectedWaitingTokens, int(waitingTokens))
		})
	}
}

func Test_TrackingObservations(t *testing.T) {
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(cleanupMetrics(reporter))

	tcs := []struct {
		name                  string
		observation           exectypes2.Observation
		state                 exectypes2.PluginState
		expectedMessageCount  int
		expectedCommitReports int
		expectedNonces        int
	}{
		{
			name: "empty/missing structs should not report anything",
			observation: exectypes2.Observation{
				CommitReports: make(exectypes2.CommitObservations),
				Messages:      make(exectypes2.MessageObservations),
			},
			state:                 exectypes2.GetCommitReports,
			expectedCommitReports: 0,
			expectedMessageCount:  0,
		},
		{
			name: "observation with nonces",
			observation: exectypes2.Observation{
				Nonces: exectypes2.NonceObservations{
					cciptypes.ChainSelector(123): map[string]uint64{
						"0x123": 1,
						"0x456": 2,
					},
					cciptypes.ChainSelector(456): map[string]uint64{
						"0x123": 3,
					},
					cciptypes.ChainSelector(789): nil,
				},
			},
			state:                 exectypes2.GetMessages,
			expectedCommitReports: 0,
			expectedMessageCount:  0,
			expectedNonces:        3,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackObservation(tc.observation, tc.state)

			nonces := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.ObservationMethod, string(tc.state), "nonces",
				),
			)
			require.Equal(t, tc.expectedNonces, int(nonces))

			commitReports := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.ObservationMethod, string(tc.state), "commitReports",
				),
			)
			require.Equal(t, tc.expectedCommitReports, int(commitReports))

			messages := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.ObservationMethod, string(tc.state), "messages",
				),
			)
			require.Equal(t, tc.expectedMessageCount, int(messages))
		})
	}
}

func Test_TrackingOutcomes(t *testing.T) {
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(cleanupMetrics(reporter))

	tcs := []struct {
		name                     string
		outcome                  exectypes2.Outcome
		state                    exectypes2.PluginState
		expectedMessagesCount    int
		expectedSourceChainCount int
		expectedTokenDataCount   int
	}{
		{
			name: "empty structs should not report anything",
			outcome: exectypes2.Outcome{
				Reports: []cciptypes.ExecutePluginReport{
					{
						ChainReports: []cciptypes.ExecutePluginReportSingleChain{},
					},
				},
			},
			state:                    exectypes2.GetCommitReports,
			expectedMessagesCount:    0,
			expectedSourceChainCount: 0,
			expectedTokenDataCount:   0,
		},
		{
			name: "single chain report should be properly tracked",
			outcome: exectypes2.Outcome{
				Reports: []cciptypes.ExecutePluginReport{
					{
						ChainReports: []cciptypes.ExecutePluginReportSingleChain{
							{
								SourceChainSelector: 123,
								Messages:            make([]cciptypes.Message, 2),
								OffchainTokenData:   make([][][]byte, 0),
							},
						},
					},
				},
			},
			state:                    exectypes2.Filter,
			expectedMessagesCount:    2,
			expectedTokenDataCount:   0,
			expectedSourceChainCount: 1,
		},
		{
			name: "multiple chain reports should be tracked",
			outcome: exectypes2.Outcome{
				Reports: []cciptypes.ExecutePluginReport{
					{
						ChainReports: []cciptypes.ExecutePluginReportSingleChain{
							{
								SourceChainSelector: 123,
								Messages:            make([]cciptypes.Message, 10),
								OffchainTokenData:   make([][][]byte, 20),
							},
							{
								SourceChainSelector: 250,
								Messages:            make([]cciptypes.Message, 5),
								OffchainTokenData:   make([][][]byte, 10),
							},
						},
					},
				},
			},
			state:                    exectypes2.GetCommitReports,
			expectedMessagesCount:    15,
			expectedTokenDataCount:   30,
			expectedSourceChainCount: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackOutcome(tc.outcome, tc.state)

			messages := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.OutcomeMethod, string(tc.state), "messages",
				),
			)
			require.Equal(t, tc.expectedMessagesCount, int(messages))

			sourceChains := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.OutcomeMethod, string(tc.state), "sourceChains",
				),
			)
			require.Equal(t, tc.expectedSourceChainCount, int(sourceChains))

			tokenData := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					"solana", chainID, plugincommon.OutcomeMethod, string(tc.state), "tokenData",
				),
			)
			require.Equal(t, tc.expectedTokenDataCount, int(tokenData))
		})
	}
}

func Test_SequenceNumbers(t *testing.T) {
	selector1 := cciptypes.ChainSelector(12922642891491394802)
	selector2 := cciptypes.ChainSelector(909606746561742123)

	tt := []struct {
		name   string
		obs    exectypes2.Observation
		out    exectypes2.Outcome
		method plugincommon.MethodType
		exp    map[cciptypes.ChainSelector]cciptypes.SeqNum
	}{
		{
			name:   "empty observation should not report anything",
			obs:    exectypes2.Observation{},
			method: plugincommon.ObservationMethod,
			exp:    map[cciptypes.ChainSelector]cciptypes.SeqNum{},
		},
		{
			name: "single chain observation with seq nr",
			obs: exectypes2.Observation{
				Messages: exectypes2.MessageObservations{
					selector1: {
						1: {},
						4: {},
					},
				},
			},
			method: plugincommon.ObservationMethod,
			exp: map[cciptypes.ChainSelector]cciptypes.SeqNum{
				selector1: 4,
			},
		},
		{
			name: "multiple chain observations with sequence numbers",
			obs: exectypes2.Observation{
				Messages: exectypes2.MessageObservations{
					selector1: {
						1: {},
						2: {},
					},
					selector2: {
						4: {},
					},
				},
			},
			method: plugincommon.ObservationMethod,
			exp: map[cciptypes.ChainSelector]cciptypes.SeqNum{
				selector1: 2,
				selector2: 4,
			},
		},
		{
			name: "single chain outcome with seq nr",
			out: exectypes2.Outcome{
				CommitReports: []exectypes2.CommitData{
					{
						SourceChain: selector1,
						Messages: []cciptypes.Message{
							{
								Header: cciptypes.RampMessageHeader{
									SequenceNumber: 2,
								},
							},
						},
					},
				},
			},
			method: plugincommon.OutcomeMethod,
			exp: map[cciptypes.ChainSelector]cciptypes.SeqNum{
				selector1: 2,
			},
		},
		{
			name: "multiple chain outcomes with sequence numbers",
			out: exectypes2.Outcome{
				CommitReports: []exectypes2.CommitData{
					{
						SourceChain: selector1,
						Messages: []cciptypes.Message{
							{
								Header: cciptypes.RampMessageHeader{
									SequenceNumber: 2,
								},
							},
						},
					},
					{
						SourceChain: selector2,
						Messages: []cciptypes.Message{
							{
								Header: cciptypes.RampMessageHeader{
									SequenceNumber: 4,
								},
							},
							{
								Header: cciptypes.RampMessageHeader{
									SequenceNumber: 2,
								},
							},
						},
					},
				},
			},
			method: plugincommon.OutcomeMethod,
			exp: map[cciptypes.ChainSelector]cciptypes.SeqNum{
				selector1: 2,
				selector2: 4,
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			reporter, err := NewPromReporter(logger.Test(t), selector)
			require.NoError(t, err)

			t.Cleanup(cleanupMetrics(reporter))

			switch tc.method {
			case plugincommon.ObservationMethod:
				reporter.TrackObservation(tc.obs, exectypes2.GetCommitReports)
			case plugincommon.OutcomeMethod:
				reporter.TrackOutcome(tc.out, exectypes2.GetCommitReports)
			}

			for sourceSelector, maxSeqNr := range tc.exp {
				sourceFamily, sourceID, ok := libs.GetChainInfoFromSelector(sourceSelector)
				require.True(t, ok)

				seqNum := testutil.ToFloat64(
					reporter.sequenceNumbers.WithLabelValues("solana", chainID, sourceFamily, sourceID, tc.method),
				)
				require.Equal(t, float64(maxSeqNr), seqNum)
			}
		})
	}
}

func Test_ExecLatency(t *testing.T) {
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Run("single latency observation", func(t *testing.T) {
		reporter.TrackLatency(exectypes2.GetCommitReports, plugincommon.ObservationMethod, 100, nil)
		l1 := internal.CounterFromHistogramByLabels(
			t, reporter.latencyHistogram, "solana", chainID, "observation", "GetCommitReports",
		)
		require.Equal(t, 1, l1)

		errs := testutil.ToFloat64(
			reporter.execErrors.WithLabelValues("solana", chainID, "observation", "GetCommitReports"),
		)
		require.Equal(t, float64(0), errs)
	})

	t.Run("multiple latency outcomes", func(t *testing.T) {
		passCounter := 10
		for i := 0; i < passCounter; i++ {
			reporter.TrackLatency(exectypes2.Filter, plugincommon.OutcomeMethod, time.Second, nil)
		}
		l2 := internal.CounterFromHistogramByLabels(t, reporter.latencyHistogram, "solana", chainID, "outcome", "Filter")
		require.Equal(t, passCounter, l2)
	})

	t.Run("multiple latency observation with errors", func(t *testing.T) {
		errCounter := 5
		for i := 0; i < errCounter; i++ {
			reporter.TrackLatency(exectypes2.GetMessages, plugincommon.ObservationMethod, time.Second, fmt.Errorf("error"))
		}
		errs := testutil.ToFloat64(
			reporter.execErrors.WithLabelValues("solana", chainID, "observation", "GetMessages"),
		)
		require.Equal(t, float64(errCounter), errs)
	})
}

func Test_LatencyAndErrors(t *testing.T) {
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Run("single latency metric", func(t *testing.T) {
		processor := "discovery1"
		method := "query"

		reporter.TrackProcessorLatency(processor, method, time.Second, nil)
		l1 := internal.CounterFromHistogramByLabels(
			t, reporter.processorLatencyHistogram, "solana", chainID, processor, method,
		)
		require.Equal(t, 1, l1)

		errs := testutil.ToFloat64(
			reporter.processorErrors.WithLabelValues("solana", chainID, processor, method),
		)
		require.Equal(t, float64(0), errs)
	})

	t.Run("multiple latency metrics", func(t *testing.T) {
		processor := "discovery2"
		method := "observation"

		passCounter := 10
		for i := 0; i < passCounter; i++ {
			reporter.TrackProcessorLatency(processor, method, time.Second, nil)
		}
		l2 := internal.CounterFromHistogramByLabels(
			t, reporter.processorLatencyHistogram, "solana", chainID, processor, method,
		)
		require.Equal(t, passCounter, l2)
	})

	t.Run("multiple error metrics", func(t *testing.T) {
		processor := "discovery3"
		method := "outcome"

		errCounter := 5
		for i := 0; i < errCounter; i++ {
			reporter.TrackProcessorLatency(processor, method, time.Second, fmt.Errorf("error"))
		}
		errs := testutil.ToFloat64(
			reporter.processorErrors.WithLabelValues("solana", chainID, processor, method),
		)
		require.Equal(t, float64(errCounter), errs)
	})
}

func cleanupMetrics(p *PromReporter) func() {
	return func() {
		p.sequenceNumbers.Reset()
		p.outputDetailsCounter.Reset()
		p.latencyHistogram.Reset()
		p.execErrors.Reset()
		p.processorLatencyHistogram.Reset()
		p.processorErrors.Reset()
	}
}
