package metrics

import (
	"fmt"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	chainID  = "2337"
	selector = cciptypes.ChainSelector(12922642891491394802)
)

func Test_TrackingTokenReadiness(t *testing.T) {
	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(cleanupMetrics(reporter))

	tcs := []struct {
		name                  string
		observation           exectypes.TokenDataObservations
		state                 exectypes.PluginState
		expectedReadyTokens   int
		expectedWaitingTokens int
	}{
		{
			name:                  "empty/missing structs should not report anything",
			observation:           exectypes.TokenDataObservations{},
			state:                 exectypes.GetMessages,
			expectedReadyTokens:   0,
			expectedWaitingTokens: 0,
		},
		{
			name: "single chain with some tokens ready and some waiting",
			observation: exectypes.TokenDataObservations{
				cciptypes.ChainSelector(123): map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: exectypes.NewMessageTokenData(),
					2: exectypes.NewMessageTokenData(exectypes.NewSuccessTokenData([]byte("asdf"))),
				},
				cciptypes.ChainSelector(456): map[cciptypes.SeqNum]exectypes.MessageTokenData{
					4: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("asdf")),
						exectypes.NewSuccessTokenData([]byte("qwer")),
						exectypes.NewErrorTokenData(fmt.Errorf("error")),
					),
				},
			},
			expectedWaitingTokens: 1,
			expectedReadyTokens:   3,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackObservation(exectypes.Observation{TokenData: tc.observation}, tc.state)

			readyTokens := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					chainID, plugincommon.ObservationMethod, string(tc.state), "tokenReady",
				),
			)
			require.Equal(t, tc.expectedReadyTokens, int(readyTokens))

			waitingTokens := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					chainID, plugincommon.ObservationMethod, string(tc.state), "tokenWaiting",
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
		observation           exectypes.Observation
		state                 exectypes.PluginState
		expectedMessageCount  int
		expectedCommitReports int
		expectedNonces        int
	}{
		{
			name: "empty/missing structs should not report anything",
			observation: exectypes.Observation{
				CommitReports: make(exectypes.CommitObservations),
				Messages:      make(exectypes.MessageObservations),
			},
			state:                 exectypes.GetCommitReports,
			expectedCommitReports: 0,
			expectedMessageCount:  0,
		},
		{
			name: "observation with nonces",
			observation: exectypes.Observation{
				Nonces: exectypes.NonceObservations{
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
			state:                 exectypes.GetMessages,
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
					chainID, plugincommon.ObservationMethod, string(tc.state), "nonces",
				),
			)
			require.Equal(t, tc.expectedNonces, int(nonces))

			commitReports := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					chainID, plugincommon.ObservationMethod, string(tc.state), "commitReports",
				),
			)
			require.Equal(t, tc.expectedCommitReports, int(commitReports))

			messages := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					chainID, plugincommon.ObservationMethod, string(tc.state), "messages",
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
		outcome                  exectypes.Outcome
		state                    exectypes.PluginState
		expectedMessagesCount    int
		expectedSourceChainCount int
		expectedTokenDataCount   int
	}{
		{
			name: "empty structs should not report anything",
			outcome: exectypes.Outcome{
				Report: cciptypes.ExecutePluginReport{
					ChainReports: []cciptypes.ExecutePluginReportSingleChain{},
				},
			},
			state:                    exectypes.GetCommitReports,
			expectedMessagesCount:    0,
			expectedSourceChainCount: 0,
			expectedTokenDataCount:   0,
		},
		{
			name: "single chain report should be properly tracked",
			outcome: exectypes.Outcome{
				Report: cciptypes.ExecutePluginReport{
					ChainReports: []cciptypes.ExecutePluginReportSingleChain{
						{
							SourceChainSelector: 123,
							Messages:            make([]cciptypes.Message, 2),
							OffchainTokenData:   make([][][]byte, 0),
						},
					},
				},
			},
			state:                    exectypes.Filter,
			expectedMessagesCount:    2,
			expectedTokenDataCount:   0,
			expectedSourceChainCount: 1,
		},
		{
			name: "multiple chain reports should be tracked",
			outcome: exectypes.Outcome{
				Report: cciptypes.ExecutePluginReport{
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
			state:                    exectypes.GetCommitReports,
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
					chainID, plugincommon.OutcomeMethod, string(tc.state), "messages",
				),
			)
			require.Equal(t, tc.expectedMessagesCount, int(messages))

			sourceChains := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					chainID, plugincommon.OutcomeMethod, string(tc.state), "sourceChains",
				),
			)
			require.Equal(t, tc.expectedSourceChainCount, int(sourceChains))

			tokenData := testutil.ToFloat64(
				reporter.outputDetailsCounter.WithLabelValues(
					chainID, plugincommon.OutcomeMethod, string(tc.state), "tokenData",
				),
			)
			require.Equal(t, tc.expectedTokenDataCount, int(tokenData))
		})
	}
}

func Test_SequenceNumbers(t *testing.T) {
	chain1 := "2337"
	selector1 := cciptypes.ChainSelector(12922642891491394802)
	chain2 := "3337"
	selector2 := cciptypes.ChainSelector(4793464827907405086)

	tt := []struct {
		name   string
		obs    exectypes.Observation
		out    exectypes.Outcome
		method plugincommon.MethodType
		exp    map[string]cciptypes.SeqNum
	}{
		{
			name:   "empty observation should not report anything",
			obs:    exectypes.Observation{},
			method: plugincommon.ObservationMethod,
			exp:    map[string]cciptypes.SeqNum{},
		},
		{
			name: "single chain observation with seq nr",
			obs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					selector1: {
						1: {},
						4: {},
					},
				},
			},
			method: plugincommon.ObservationMethod,
			exp: map[string]cciptypes.SeqNum{
				chain1: 4,
			},
		},
		{
			name: "multiple chain observations with sequence numbers",
			obs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
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
			exp: map[string]cciptypes.SeqNum{
				chain1: 2,
				chain2: 4,
			},
		},
		{
			name: "single chain outcome with seq nr",
			out: exectypes.Outcome{
				CommitReports: []exectypes.CommitData{
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
			exp: map[string]cciptypes.SeqNum{
				chain1: 2,
			},
		},
		{
			name: "multiple chain outcomes with sequence numbers",
			out: exectypes.Outcome{
				CommitReports: []exectypes.CommitData{
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
			exp: map[string]cciptypes.SeqNum{
				chain1: 2,
				chain2: 4,
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
				reporter.TrackObservation(tc.obs, exectypes.GetCommitReports)
			case plugincommon.OutcomeMethod:
				reporter.TrackOutcome(tc.out, exectypes.GetCommitReports)
			}

			for sourceChain, maxSeqNr := range tc.exp {
				seqNum := testutil.ToFloat64(
					reporter.sequenceNumbers.WithLabelValues(chainID, sourceChain, tc.method),
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
		reporter.TrackLatency(exectypes.GetCommitReports, plugincommon.ObservationMethod, 100, nil)
		l1 := internal.CounterFromHistogramByLabels(t, reporter.latencyHistogram, chainID, "observation", "GetCommitReports")
		require.Equal(t, 1, l1)

		errs := testutil.ToFloat64(
			reporter.execErrors.WithLabelValues(chainID, "observation", "GetCommitReports"),
		)
		require.Equal(t, float64(0), errs)
	})

	t.Run("multiple latency outcomes", func(t *testing.T) {
		passCounter := 10
		for i := 0; i < passCounter; i++ {
			reporter.TrackLatency(exectypes.Filter, plugincommon.OutcomeMethod, time.Second, nil)
		}
		l2 := internal.CounterFromHistogramByLabels(t, reporter.latencyHistogram, chainID, "outcome", "Filter")
		require.Equal(t, passCounter, l2)
	})

	t.Run("multiple latency observation with errors", func(t *testing.T) {
		errCounter := 5
		for i := 0; i < errCounter; i++ {
			reporter.TrackLatency(exectypes.GetMessages, plugincommon.ObservationMethod, time.Second, fmt.Errorf("error"))
		}
		errs := testutil.ToFloat64(
			reporter.execErrors.WithLabelValues(chainID, "observation", "GetMessages"),
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
		l1 := internal.CounterFromHistogramByLabels(t, reporter.processorLatencyHistogram, chainID, processor, method)
		require.Equal(t, 1, l1)

		errs := testutil.ToFloat64(
			reporter.processorErrors.WithLabelValues(chainID, processor, method),
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
		l2 := internal.CounterFromHistogramByLabels(t, reporter.processorLatencyHistogram, chainID, processor, method)
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
			reporter.processorErrors.WithLabelValues(chainID, processor, method),
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
