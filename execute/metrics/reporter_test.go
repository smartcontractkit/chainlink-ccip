package metrics

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
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
				reporter.tokenDataReadinessCounter.WithLabelValues(chainID, string(tc.state), "ready"),
			)
			require.Equal(t, tc.expectedReadyTokens, int(readyTokens))

			waitingTokens := testutil.ToFloat64(
				reporter.tokenDataReadinessCounter.WithLabelValues(chainID, string(tc.state), "waiting"),
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
		name                   string
		observation            exectypes.Observation
		state                  exectypes.PluginState
		expectedMessageCount   int
		expectedCommitReports  int
		expectedNonces         int
		expectedCostlyMessages int
	}{
		{
			name: "empty/missing structs should not report anything",
			observation: exectypes.Observation{
				CommitReports: make(exectypes.CommitObservations),
				Messages:      make(exectypes.MessageObservations),
			},
			state:                  exectypes.GetCommitReports,
			expectedCommitReports:  0,
			expectedMessageCount:   0,
			expectedNonces:         0,
			expectedCostlyMessages: 0,
		},
		{
			name: "observation with messages which some of them are costly",
			observation: exectypes.Observation{
				CommitReports: exectypes.CommitObservations{
					cciptypes.ChainSelector(123): make([]exectypes.CommitData, 2),
					cciptypes.ChainSelector(456): nil,
					cciptypes.ChainSelector(780): make([]exectypes.CommitData, 1),
				},
				Messages: exectypes.MessageObservations{
					cciptypes.ChainSelector(123): map[cciptypes.SeqNum]cciptypes.Message{
						1: {},
						2: {},
					},
				},
				CostlyMessages: make([]cciptypes.Bytes32, 1),
			},
			state:                  exectypes.Filter,
			expectedCommitReports:  3,
			expectedMessageCount:   2,
			expectedNonces:         0,
			expectedCostlyMessages: 1,
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
			state:                  exectypes.GetMessages,
			expectedCommitReports:  0,
			expectedMessageCount:   0,
			expectedNonces:         3,
			expectedCostlyMessages: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackObservation(tc.observation, tc.state)

			costlyMsgs := testutil.ToFloat64(reporter.costlyMessagesCounter.WithLabelValues(chainID, string(tc.state)))
			require.Equal(t, tc.expectedCostlyMessages, int(costlyMsgs))

			nonces := testutil.ToFloat64(
				reporter.observationDetailsCounter.WithLabelValues(chainID, string(tc.state), "nonces"),
			)
			require.Equal(t, tc.expectedNonces, int(nonces))

			commitReports := testutil.ToFloat64(
				reporter.observationDetailsCounter.WithLabelValues(chainID, string(tc.state), "commitReports"),
			)
			require.Equal(t, tc.expectedCommitReports, int(commitReports))

			messages := testutil.ToFloat64(
				reporter.observationDetailsCounter.WithLabelValues(chainID, string(tc.state), "messages"),
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
				reporter.outcomeDetailsCounter.WithLabelValues(chainID, string(tc.state), "messages"),
			)
			require.Equal(t, tc.expectedMessagesCount, int(messages))

			sourceChains := testutil.ToFloat64(
				reporter.outcomeDetailsCounter.WithLabelValues(chainID, string(tc.state), "sourceChains"),
			)
			require.Equal(t, tc.expectedSourceChainCount, int(sourceChains))

			tokenData := testutil.ToFloat64(
				reporter.outcomeDetailsCounter.WithLabelValues(chainID, string(tc.state), "tokenData"),
			)
			require.Equal(t, tc.expectedTokenDataCount, int(tokenData))
		})
	}
}

func cleanupMetrics(p *PromReporter) func() {
	return func() {
		p.tokenDataReadinessCounter.Reset()
		p.outcomeDetailsCounter.Reset()
		p.observationDetailsCounter.Reset()
		p.observationDetailsCounter.Reset()
	}
}
