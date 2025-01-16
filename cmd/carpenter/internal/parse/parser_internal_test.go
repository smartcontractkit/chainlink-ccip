package parse

import "testing"

func Test_sanitizeString(t *testing.T) {
	type args struct {
		s string
	}
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
		{
			name: "Trim whitespace and escapes",
			lines: []string{
				`    {\"\"} `,
			},
			want: []string{`{""}`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			for i, line := range tt.lines {
				got = append(got, sanitizeString(line))
				if got[i] != tt.want[i] {
					t.Errorf("sanitizeString() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
