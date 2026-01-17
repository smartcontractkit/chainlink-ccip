package v1_test

import (
	"encoding/base64"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1/ocrtypecodecpb"

	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	v1 "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
)

var inputObjectOld = exectypes.Outcome{
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

var inputObjectReports = exectypes.Outcome{
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

func TestProtobufBackwardCompatability(t *testing.T) {
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

// TestExistingEncodingCompatibility verifies existing encoded data can be decoded correctly
func TestExistingEncodingCompatibility(t *testing.T) {
	// Base64 encoded data from original test
	encodedB64 := "Cgd5ZWUgaGF3EjMIYxoCCAkgyQEqIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMgYI+gEQrAIaBAoCCHs="

	codec := v1.NewExecCodecProto()
	encoded, err := base64.StdEncoding.DecodeString(encodedB64)
	require.NoError(t, err)

	outcome, err := codec.DecodeOutcome(encoded)
	require.NoError(t, err)

	// Verify the outcome has both fields populated
	require.Equal(t, exectypes.PluginState("yee haw"), outcome.State)
	require.NotEmpty(t, outcome.Report.ChainReports)
	require.Equal(t, cciptypes.ChainSelector(123), outcome.Report.ChainReports[0].SourceChainSelector)
	require.Len(t, outcome.Reports, 1)
	require.Equal(t, outcome.Report, outcome.Reports[0])
}

// TestSingleReportBackwardCompatibility tests that single reports use the legacy field
func TestSingleReportBackwardCompatibility(t *testing.T) {
	codec := v1.NewExecCodecProto()

	// Create an outcome with a single report in Reports
	outcome := exectypes.Outcome{
		State: "test state",
		Reports: []cciptypes.ExecutePluginReport{
			{
				ChainReports: []cciptypes.ExecutePluginReportSingleChain{
					{SourceChainSelector: 42},
				},
			},
		},
	}

	encoded, err := codec.EncodeOutcome(outcome)
	require.NoError(t, err)

	// Manually decode the protobuf to verify structure
	pbOutcome := &ocrtypecodecpb.ExecOutcome{}
	err = proto.Unmarshal(encoded, pbOutcome)
	require.NoError(t, err)

	// With a single report, it should use the legacy field only
	require.NotNil(t, pbOutcome.ExecutePluginReport)
	require.Empty(t, pbOutcome.ExecutePluginReports)
	require.Equal(t, uint64(42), pbOutcome.ExecutePluginReport.ChainReports[0].SourceChainSelector)

	decodedOutcome, err := codec.DecodeOutcome(encoded)
	require.NoError(t, err)

	// Both fields should now be populated in the decoded outcome
	require.Equal(t, cciptypes.ChainSelector(42), decodedOutcome.Report.ChainReports[0].SourceChainSelector)
	require.Len(t, decodedOutcome.Reports, 1)
	require.Equal(t, cciptypes.ChainSelector(42), decodedOutcome.Reports[0].ChainReports[0].SourceChainSelector)
}

// TestMultipleReportEncoding verifies that multiple reports use the new field
func TestMultipleReportEncoding(t *testing.T) {
	codec := v1.NewExecCodecProto()

	// Create an outcome with multiple reports
	outcome := exectypes.Outcome{
		State: "test state",
		Reports: []cciptypes.ExecutePluginReport{
			{
				ChainReports: []cciptypes.ExecutePluginReportSingleChain{
					{SourceChainSelector: 42},
				},
			},
			{
				ChainReports: []cciptypes.ExecutePluginReportSingleChain{
					{SourceChainSelector: 43},
				},
			},
		},
	}

	encoded, err := codec.EncodeOutcome(outcome)
	require.NoError(t, err)

	// Manually decode the protobuf to verify structure
	pbOutcome := &ocrtypecodecpb.ExecOutcome{}
	err = proto.Unmarshal(encoded, pbOutcome)
	require.NoError(t, err)

	// With multiple reports, it should use the new field only
	require.Nil(t, pbOutcome.ExecutePluginReport)
	require.Len(t, pbOutcome.ExecutePluginReports, 2)
	require.Equal(t, uint64(42), pbOutcome.ExecutePluginReports[0].ChainReports[0].SourceChainSelector)
	require.Equal(t, uint64(43), pbOutcome.ExecutePluginReports[1].ChainReports[0].SourceChainSelector)

	decodedOutcome, err := codec.DecodeOutcome(encoded)
	require.NoError(t, err)
	require.Len(t, decodedOutcome.Reports, 2)
	require.Equal(t, cciptypes.ChainSelector(42), decodedOutcome.Reports[0].ChainReports[0].SourceChainSelector)
	require.Equal(t, cciptypes.ChainSelector(43), decodedOutcome.Reports[1].ChainReports[0].SourceChainSelector)
}

// TestEdgeCases tests various edge cases
func TestEdgeCases(t *testing.T) {
	codec := v1.NewExecCodecProto()

	// Test with empty Reports array
	t.Run("empty reports array", func(t *testing.T) {
		outcome := exectypes.Outcome{
			State:   "test state",
			Reports: []cciptypes.ExecutePluginReport{},
		}

		encoded, err := codec.EncodeOutcome(outcome)
		require.NoError(t, err)

		decoded, err := codec.DecodeOutcome(encoded)
		require.NoError(t, err)

		require.Equal(t, outcome.State, decoded.State)
		require.Empty(t, decoded.Reports)
		require.Nil(t, decoded.Report.ChainReports)
	})

	t.Run("report with zero chain reports", func(t *testing.T) {
		outcome := exectypes.Outcome{
			State:  "test state",
			Report: cciptypes.ExecutePluginReport{},
		}

		encoded, err := codec.EncodeOutcome(outcome)
		require.NoError(t, err)

		decoded, err := codec.DecodeOutcome(encoded)
		require.NoError(t, err)

		require.Equal(t, outcome.State, decoded.State)
		require.Empty(t, decoded.Reports)
		require.Nil(t, decoded.Report.ChainReports)
	})
}
