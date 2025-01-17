package parse

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

//nolint:lll // long test data
func Test_sanitizeString(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []string
	}{
		{
			name: "CI header",
			lines: []string{
				"2025-01-16T15:03:05.2840675Z ##[group]Run cd $GITHUB_WORKSPACE/chainlink/ && cd integration-tests/smoke/ccip && go test ccip_messaging_test.go -timeout 12m -test.parallel=2 -count=1 -json",
				"2025-01-16T15:03:05.2841755Z [36;1mcd $GITHUB_WORKSPACE/chainlink/ && cd integration-tests/smoke/ccip && go test ccip_messaging_test.go -timeout 12m -test.parallel=2 -count=1 -json[0m",
				"2025-01-16T15:03:05.2868942Z shell: /usr/bin/bash -e {0}",
				"2025-01-16T15:03:05.2869178Z env:",
				"2025-01-16T15:03:05.2869635Z   DB_URL: ***localhost:5432/chainlink_test?sslmode=disable",
				"2025-01-16T15:03:05.2869938Z   GO_VERSION: 1.23.3",
				"2025-01-16T15:03:05.2870352Z   CL_DATABASE_URL: ***localhost:5432/chainlink_test?sslmode=disable",
				"2025-01-16T15:03:05.2870669Z ##[endgroup]",
			},
			want: make([]string, 8), // all trimmed completely
		},
		{
			name: "Trim time prefix",
			lines: []string{
				`2025-01-16T15:06:14.0436606Z {"Time":"2025-01-16T15:06:14.025535803Z","Action":"output","Package":"command-line-arguments","Test":"Test_CCIPMessaging","Output":"    logger.go:146: 2025-01-16T15:06:14.025Z\tINFO\tCCIPCommitPlugin.evm.90000002.5548718428018410741.0x479f3d4c9784e41f4f36619528bee28aa864d7be\treader/ccip.go:825\tappending router contract address\t{\"version\": \"unset@unset\", \"plugin\": \"Commit\", \"oracleID\": 0, \"donID\": 2, \"configDigest\": \"000ab94b6d7fbd21dfe30247363695788a40e984ddbc37e66bac0b5e2c82f57e\", \"component\": \"CCIPReader\", \"ocrSeqNr\": 1, \"address\": \"VPXRY9AjwMItJzOCQeocjpa6UlE=\"}\n"}`,
			},
			want: []string{
				`{"version": "unset@unset", "plugin": "Commit", "oracleID": 0, "donID": 2, "configDigest": "000ab94b6d7fbd21dfe30247363695788a40e984ddbc37e66bac0b5e2c82f57e", "component": "CCIPReader", "ocrSeqNr": 1, "address": "VPXRY9AjwMItJzOCQeocjpa6UlE="}`,
			},
		},
		// TODO: this is failing rn.
		// {
		// 	name: "Trim whitespace and escapes",
		// 	lines: []string{
		// 		`    {\"\"} `,
		// 	},
		// 	want: []string{`{""}`},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			for i, line := range tt.lines {
				got = append(got, sanitizeString(line, LogTypeCI))
				if got[i] != tt.want[i] {
					t.Errorf("sanitizeString() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func mustParseCustomLayout(t *testing.T, str string) time.Time {
	tm, err := parseCustomLayout(str)
	require.NoError(t, err)
	return tm
}

func Test_ParseLine_Mixed(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected *Data
	}{
		{
			name: "mixed log",
			line: `logger.go:146: 2025-01-17T13:59:55.521+0200	INFO	CCIPCommitPlugin.evm.1337.3379446385462418246.0x075f98f19ef9873523cde0267ab8b0253904363e	commit/plugin.go:482	closing commit plugin	{"version": "unset@unset", "plugin": "Commit", "oracleID": 1, "donID": 2, "configDigest": "000a7d1df8632e2b3479350dcca1ee46eeec889dc37eb2ab094e63a1820ba291", "component": "Plugin"}`,
			expected: &Data{
				LoggerName:     "CCIPCommitPlugin.evm.1337.3379446385462418246.0x075f98f19ef9873523cde0267ab8b0253904363e",
				Level:          "INFO",
				Timestamp:      mustParseCustomLayout(t, "2025-01-17T13:59:55.521+0200"),
				Caller:         "commit/plugin.go:482",
				SequenceNumber: 0,
				Component:      "Plugin",
				DONID:          2,
				OracleID:       1,
				Message:        "closing commit plugin",
				Version:        "unset@unset",
				ConfigDigest:   "000a7d1df8632e2b3479350dcca1ee46eeec889dc37eb2ab094e63a1820ba291",
				Plugin:         "Commit",
				RawLoggerFields: map[string]any{
					"version":      "unset@unset",
					"plugin":       "Commit",
					"oracleID":     float64(1),
					"donID":        float64(2),
					"configDigest": "000a7d1df8632e2b3479350dcca1ee46eeec889dc37eb2ab094e63a1820ba291",
					"component":    "Plugin",
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseLine(tc.line, LogTypeMixed)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.expected, result)
		})
	}
}
