package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type expected struct {
	counter *prometheus.CounterVec
	labels  []string
	count   int
}

func Test_TrackingObservations(t *testing.T) {

}

func Test_TrackingOutcomes(t *testing.T) {
	chainID := "2337"
	selector := cciptypes.ChainSelector(12922642891491394802)

	reporter, err := NewPromReporter(logger.Test(t), selector)
	require.NoError(t, err)

	t.Cleanup(func() {
		promOutcomeDetails.Reset()
	})

	tcs := []struct {
		name     string
		outcome  exectypes.Outcome
		state    exectypes.PluginState
		expected []expected
	}{
		{
			name: "empty structs shouldn not report anything",
			outcome: exectypes.Outcome{
				Report: cciptypes.ExecutePluginReport{
					ChainReports: []cciptypes.ExecutePluginReportSingleChain{},
				},
			},
			state: exectypes.GetCommitReports,
			expected: []expected{
				{
					counter: promOutcomeDetails,
					labels:  []string{"GetCommitReports", "messages"},
					count:   0,
				},
				{
					counter: promOutcomeDetails,
					labels:  []string{"GetCommitReports", "tokenData"},
					count:   0,
				},
				{
					counter: promOutcomeDetails,
					labels:  []string{"GetCommitReports", "sourceChains"},
					count:   0,
				},
			},
		},
		{
			name: "empty structs shouldn not report anything",
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
			state: exectypes.Filter,
			expected: []expected{
				{
					counter: promOutcomeDetails,
					labels:  []string{"Filter", "messages"},
					count:   2,
				},
				{
					counter: promOutcomeDetails,
					labels:  []string{"Filter", "tokenData"},
					count:   0,
				},
				{
					counter: promOutcomeDetails,
					labels:  []string{"Filter", "sourceChains"},
					count:   1,
				},
			},
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
			state: exectypes.GetCommitReports,
			expected: []expected{
				{
					counter: promOutcomeDetails,
					labels:  []string{"GetCommitReports", "messages"},
					count:   15,
				},
				{
					counter: promOutcomeDetails,
					labels:  []string{"GetCommitReports", "tokenData"},
					count:   30,
				},
				{
					counter: promOutcomeDetails,
					labels:  []string{"GetCommitReports", "sourceChains"},
					count:   2,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			reporter.TrackOutcome(tc.outcome, tc.state)

			for _, exp := range tc.expected {
				labels := make([]string, 0, len(exp.labels)+1)
				labels = append(labels, chainID)
				labels = append(labels, exp.labels...)

				actual := testutil.ToFloat64(exp.counter.WithLabelValues(labels...))
				require.Equal(t, exp.count, int(actual))
			}
		})
	}
}
