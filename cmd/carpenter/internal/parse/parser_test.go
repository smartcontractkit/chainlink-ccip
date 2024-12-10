package parse

import (
	"fmt"
	"testing"
	"time"
)

func mustParseTime(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, "2024-12-09T20:59:53.531Z")
	return t
}

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected Data
	}{
		{
			name: "merkle root",
			line: `{"level":"info","ts":"2024-12-09T20:59:53.531Z","logger":"CCIPCommitPlugin.evm.1337.3379446385462418246.0xe6e340d132b5f46d1e472debcd681b2abc16e57e","caller":"merkleroot/outcome.go:37","msg":"Sending Outcome","version":"2.18.0@732cc15","plugin":"Commit","oracleID":3,"donID":1,"plugin":"Commit","oracleID":3,"donID":1,"processor":"MerkleRoot","outcome":{"outcomeType":1,"rangesSelectedForReport":[],"rootsToReport":null,"offRampNextSeqNums":[{"chainSel":12922642891491394802,"seqNum":2}],"reportTransmissionCheckAttempts":0,"rmnReportSignatures":null,"rmnRemoteCfg":{"contractAddress":"0x322813fd9a801c5507c9de605d63cea4f2ce6c44","configDigest":"0x000be848c9e6eacda7ab37900ed1a6261fd78e7d53b9483cfb8e7a83e75c0193","signers":[{"onchainPublicKey":"0x0100000000000000000000000000000000000000","nodeIndex":0}],"f":0,"configVersion":1,"rmnReportVersion":"0x9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53"}},"nextState":1,"outcomeDuration":0.00010525}`,
			expected: Data{
				Level:     "info",
				Timestamp: mustParseTime("2024-12-09T20:59:53.531Z"),
				//DONID:
			},
		},
	}
	fmt.Println(tests)
}