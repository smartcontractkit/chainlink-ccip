package parse_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/filter"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func mustParseTime(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, "2024-12-09T20:59:53.531Z")
	return t
}

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected parse.Data
	}{
		{
			name: "merkle root",
			line: `{"level":"info","ts":"2024-12-09T20:59:53.531Z","logger":"CCIPCommitPlugin.evm.1337.3379446385462418246.0xe6e340d132b5f46d1e472debcd681b2abc16e57e","caller":"merkleroot/outcome.go:37","msg":"Sending Outcome","version":"2.18.0@732cc15","plugin":"Commit","oracleID":3,"donID":1,"processor":"MerkleRoot","outcome":{"outcomeType":1,"rangesSelectedForReport":[],"rootsToReport":null,"offRampNextSeqNums":[{"chainSel":12922642891491394802,"seqNum":2}],"reportTransmissionCheckAttempts":0,"rmnReportSignatures":null,"rmnRemoteCfg":{"contractAddress":"0x322813fd9a801c5507c9de605d63cea4f2ce6c44","configDigest":"0x000be848c9e6eacda7ab37900ed1a6261fd78e7d53b9483cfb8e7a83e75c0193","signers":[{"onchainPublicKey":"0x0100000000000000000000000000000000000000","nodeIndex":0}],"f":0,"configVersion":1,"rmnReportVersion":"0x9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53"}},"nextState":1,"outcomeDuration":0.00010525}`,
			expected: parse.Data{
				FilterName:     "Merkle Root Observation",
				LoggerName:     "CCIPCommitPlugin.evm.1337.3379446385462418246.0xe6e340d132b5f46d1e472debcd681b2abc16e57e",
				Level:          "info",
				Timestamp:      mustParseTime("2024-12-09T20:59:53.531Z"),
				Message:        "Sending Outcome",
				Version:        "2.18.0@732cc15",
				Caller:         "merkleroot/outcome.go:37",
				Plugin:         "Commit",
				DONID:          1,
				OracleID:       3,
				SequenceNumber: 0,
				Component:      "MerkleRoot",
				Details:        "Sending Outcome",
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := parse.Filter(tc.line, parse.LogTypeJSON)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.expected, *result)
		})
	}
	fmt.Println(tests)
}

func Test_ParseLine_Mixed(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected map[string]any
	}{
		{
			name: "mixed log",
			line: `logger.go:146: 2025-01-17T13:59:55.521+0200	INFO	CCIPCommitPlugin.evm.1337.3379446385462418246.0x075f98f19ef9873523cde0267ab8b0253904363e	commit/plugin.go:482	closing commit plugin	{"version": "unset@unset", "plugin": "Commit", "oracleID": 1, "donID": 2, "configDigest": "000a7d1df8632e2b3479350dcca1ee46eeec889dc37eb2ab094e63a1820ba291", "component": "Plugin"}`,
			expected: map[string]any{
				"caller":     "commit/plugin.go:482",
				"component":  "Plugin",
				"level":      "INFO",
				"logger":     "CCIPCommitPlugin.evm.1337.3379446385462418246.0x075f98f19ef9873523cde0267ab8b0253904363e",
				"message":    "closing commit plugin",
				"timestamp":  "2025-01-17T13:59:55.521+0200",
				"jsonFields": `{"version": "unset@unset", "plugin": "Commit", "oracleID": 1, "donID": 2, "configDigest": "000a7d1df8632e2b3479350dcca1ee46eeec889dc37eb2ab094e63a1820ba291", "component": "Plugin"}`,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result, err := parse.ParseLine(tc.line, parse.LogTypeMixed)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.expected, result)
		})
	}
}
