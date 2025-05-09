package v1_test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	v1 "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestProtobufBackwardCompatability(t *testing.T) {
	inputObjectOld := exectypes.Outcome{
		State: "yee haw",
		CommitReports: []exectypes.CommitData{
			{
				SourceChain:         99,
				Timestamp:           time.UnixMilli(9000),
				BlockNum:            201,
				SequenceNumberRange: cciptypes.NewSeqNumRange(250, 300),
			},
		},
		Report: cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{
				{
					SourceChainSelector: 123,
				},
			},
		},
	}

	inputObjectReports := exectypes.Outcome{
		State: "yee haw",
		CommitReports: []exectypes.CommitData{
			{
				SourceChain:         99,
				Timestamp:           time.UnixMilli(9000),
				BlockNum:            201,
				SequenceNumberRange: cciptypes.NewSeqNumRange(250, 300),
			},
		},
		Reports: []cciptypes.ExecutePluginReport{
			{
				ChainReports: []cciptypes.ExecutePluginReportSingleChain{
					{
						SourceChainSelector: 123,
					},
				},
			},
		},
	}

	testcases := []struct {
		name           string
		encodedOutcome string // b64
		input          exectypes.Outcome
	}{
		{
			name:           "original outcome",
			encodedOutcome: "Cgd5ZWUgaGF3EjMIYxoCCAkgyQEqIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMgYI+gEQrAIaBAoCCHs=",
			input:          inputObjectOld,
		},
		{
			name:           "reports outcome",
			encodedOutcome: "Cgd5ZWUgaGF3EjMIYxoCCAkgyQEqIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMgYI+gEQrAIaBAoCCHs=",
			input:          inputObjectReports,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			enc := v1.NewExecCodecProto()
			encoded, err := enc.EncodeOutcome(testcase.input)
			require.NoError(t, err)
			b64Encoded := base64.StdEncoding.EncodeToString(encoded)
			require.Equal(t, testcase.encodedOutcome, b64Encoded)
		})
	}
}
