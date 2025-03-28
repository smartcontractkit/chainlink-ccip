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
			line: `logger.go:146: 2025-01-20T13:44:15.033+0200	INFO	CCIPCommitPlugin.evm.90000001.909606746561742123.0x4c0c911d212a458359ff6064c6e843351bc10a5b	merkleroot/outcome.go:47	Sending Outcome	{"version": "unset@unset", "plugin": "Commit", "oracleID": 2, "donID": 1, "configDigest": "000a31c3bde664a6bbca191145d2b03781578a8d6b72793be73f6dc8025a4ff6", "component": "MerkleRoot", "ocrSeqNr": 3, "ocrPhase": "otcm", "outcome": {"outcomeType":3,"rangesSelectedForReport":null,"rootsToReport":[],"offRampNextSeqNums":[],"reportTransmissionCheckAttempts":0,"rmnReportSignatures":[],"rmnRemoteCfg":{"contractAddress":"0x06ff4addc18d262fea3e8d1d7c4ea9422fb27c2d","configDigest":"0x000bf4797ecb8e030cc32f81ec5eb543ea956e0e4ad36d9037854cd409593b6b","signers":[{"onchainPublicKey":"0x0100000000000000000000000000000000000000","nodeIndex":0}],"fSign":0,"configVersion":1,"rmnReportVersion":"0x9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53"}}, "nextState": "buildingReport", "outcomeDuration": "136.791µs"}`,
			expected: &Data{
				ProdLoggerName: "CCIPCommitPlugin.evm.90000001.909606746561742123.0x4c0c911d212a458359ff6064c6e843351bc10a5b",
				ProdLevel:      "INFO",
				ProdTimestamp:  mustParseCustomLayout(t, "2025-01-20T13:44:15.033+0200").String(),
				ProdCaller:     "merkleroot/outcome.go:47",
				ProdMessage:    "Sending Outcome",
				SequenceNumber: 3,
				OCRPhase:       "otcm",
				Component:      "MerkleRoot",
				DONID:          1,
				OracleID:       2,
				Version:        "unset@unset",
				ConfigDigest:   "000a31c3bde664a6bbca191145d2b03781578a8d6b72793be73f6dc8025a4ff6",
				Plugin:         "Commit",
				RawLoggerFields: map[string]any{
					"version":      "unset@unset",
					"plugin":       "Commit",
					"oracleID":     float64(2),
					"donID":        float64(1),
					"configDigest": "000a31c3bde664a6bbca191145d2b03781578a8d6b72793be73f6dc8025a4ff6",
					"component":    "MerkleRoot",
					"ocrSeqNr":     float64(3),
					"ocrPhase":     "otcm",
					"nextState":    "buildingReport",
					"outcome": map[string]any{
						"outcomeType":                     float64(3),
						"rangesSelectedForReport":         nil,
						"rootsToReport":                   []any{},
						"offRampNextSeqNums":              []any{},
						"reportTransmissionCheckAttempts": float64(0),
						"rmnReportSignatures":             []any{},
						"rmnRemoteCfg": map[string]any{
							"contractAddress":  "0x06ff4addc18d262fea3e8d1d7c4ea9422fb27c2d",
							"configDigest":     "0x000bf4797ecb8e030cc32f81ec5eb543ea956e0e4ad36d9037854cd409593b6b",
							"configVersion":    float64(1),
							"fSign":            float64(0),
							"rmnReportVersion": "0x9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53",
							"signers": []any{
								map[string]any{
									"onchainPublicKey": "0x0100000000000000000000000000000000000000",
									"nodeIndex":        float64(0),
								},
							},
						},
					},
					"outcomeDuration": "136.791µs",
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
