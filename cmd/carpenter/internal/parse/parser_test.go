package parse_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/filter"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func mustParseTime(t *testing.T, str string) time.Time {
	tm, err := time.Parse(time.RFC3339, str)
	require.NoError(t, err)
	return tm
}

//nolint:lll // long test data
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
				FilterName:     "CommitFilter",
				ProdLoggerName: "CCIPCommitPlugin.evm.1337.3379446385462418246.0xe6e340d132b5f46d1e472debcd681b2abc16e57e",
				ProdLevel:      "info",
				ProdTimestamp:  mustParseTime(t, "2024-12-09T20:59:53.531Z").String(),
				ProdMessage:    "Sending Outcome",
				ProdCaller:     "merkleroot/outcome.go:37",
				Version:        "2.18.0@732cc15",
				Plugin:         "Commit",
				DONID:          1,
				OracleID:       3,
				SequenceNumber: 0,
				Component:      "",
				Details:        "",
				RawLoggerFields: map[string]interface{}{
					"level":     "info",
					"ts":        "2024-12-09T20:59:53.531Z",
					"logger":    "CCIPCommitPlugin.evm.1337.3379446385462418246.0xe6e340d132b5f46d1e472debcd681b2abc16e57e",
					"caller":    "merkleroot/outcome.go:37",
					"msg":       "Sending Outcome",
					"version":   "2.18.0@732cc15",
					"plugin":    "Commit",
					"oracleID":  float64(3),
					"donID":     float64(1),
					"processor": "MerkleRoot",
					"outcome": map[string]interface{}{
						"outcomeType":                     float64(1),
						"rangesSelectedForReport":         []interface{}{},
						"rootsToReport":                   nil,
						"offRampNextSeqNums":              []interface{}{map[string]interface{}{"chainSel": float64(12922642891491394802), "seqNum": float64(2)}},
						"reportTransmissionCheckAttempts": float64(0),
						"rmnReportSignatures":             nil,
						"rmnRemoteCfg": map[string]interface{}{
							"contractAddress": "0x322813fd9a801c5507c9de605d63cea4f2ce6c44",
							"configDigest":    "0x000be848c9e6eacda7ab37900ed1a6261fd78e7d53b9483cfb8e7a83e75c0193",
							"signers": []interface{}{map[string]interface{}{
								"onchainPublicKey": "0x0100000000000000000000000000000000000000",
								"nodeIndex":        float64(0),
							}},
							"f":                float64(0),
							"configVersion":    float64(1),
							"rmnReportVersion": "0x9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53",
						},
					},
					"nextState":       float64(1),
					"outcomeDuration": 0.00010525,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := filter.Filter(tc.line, parse.LogTypeJSON, false)
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, tc.expected, *result)
		})
	}
	fmt.Println(tests)
}
