package exectypes

import (
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"github.com/stretchr/testify/require"
)

func TestPluginState_Next(t *testing.T) {
	tests := []struct {
		name    string
		p       PluginState
		want    PluginState
		isPanic bool
	}{
		{
			name: "Zero value",
			p:    Unknown,
			want: GetCommitReports,
		},
		{
			name: "Initialized",
			p:    Initialized,
			want: GetCommitReports,
		},
		{
			name: "Phase 1 to 2",
			p:    GetCommitReports,
			want: GetMessages,
		},
		{
			name: "Phase 2 to 3",
			p:    GetMessages,
			want: Filter,
		},
		{
			name: "Phase 3 to 1",
			p:    Filter,
			want: GetCommitReports,
		},
		{
			name:    "panic",
			p:       PluginState("ElToroLoco"),
			isPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.isPanic {
				require.Panics(t, func() {
					tt.p.Next()
				})
				return
			}

			if got := tt.p.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSortedOutcome(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name           string
		pendingCommits []CommitData
		report         cciptypes.ExecutePluginReport
		wantCommits    []CommitData
		wantReports    []cciptypes.ExecutePluginReportSingleChain
	}{
		{
			name: "sorts by timestamp",
			pendingCommits: []CommitData{
				{Timestamp: now.Add(1 * time.Hour), SourceChain: 1, SequenceNumberRange: seqRange(1, 2)},
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(3, 4)},
			},
			wantCommits: []CommitData{
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(3, 4)},
				{Timestamp: now.Add(1 * time.Hour), SourceChain: 1, SequenceNumberRange: seqRange(1, 2)},
			},
		},
		{
			name: "equal timestamps sort by source chain",
			pendingCommits: []CommitData{
				{Timestamp: now, SourceChain: 2, SequenceNumberRange: seqRange(1, 2)},
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(3, 4)},
			},
			wantCommits: []CommitData{
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(3, 4)},
				{Timestamp: now, SourceChain: 2, SequenceNumberRange: seqRange(1, 2)},
			},
		},
		{
			name: "equal timestamps and chains sort by sequence",
			pendingCommits: []CommitData{
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(3, 4)},
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(1, 2)},
			},
			wantCommits: []CommitData{
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(1, 2)},
				{Timestamp: now, SourceChain: 1, SequenceNumberRange: seqRange(3, 4)},
			},
		},
		{
			name: "sorts chain reports by source chain selector",
			report: cciptypes.ExecutePluginReport{
				ChainReports: []cciptypes.ExecutePluginReportSingleChain{
					{SourceChainSelector: 2},
					{SourceChainSelector: 1},
				},
			},
			wantReports: []cciptypes.ExecutePluginReportSingleChain{
				{SourceChainSelector: 1},
				{SourceChainSelector: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSortedOutcome(Unknown, tt.pendingCommits, tt.report)

			if len(tt.wantCommits) > 0 {
				require.Equal(t, tt.wantCommits, got.CommitReports)
			}
			if len(tt.wantReports) > 0 {
				require.Equal(t, tt.wantReports, got.Report.ChainReports)
			}
		})
	}
}

// seqRange is a helper to create a SequenceNumberRange
func seqRange(start, end uint64) cciptypes.SeqNumRange {
	return cciptypes.NewSeqNumRange(cciptypes.SeqNum(start), cciptypes.SeqNum(end))
}

func TestOutcome_ToLogFormat(t *testing.T) {
	now := time.Now()

	testData1 := cciptypes.Bytes("test message data")
	testData2 := cciptypes.Bytes("test message")
	// Create a test outcome with some data
	testOutcome := Outcome{
		State: GetMessages,
		CommitReports: []CommitData{
			{
				Timestamp:           now,
				SourceChain:         1,
				SequenceNumberRange: seqRange(1, 2),
				// Add some message data that should be removed
				Messages: []cciptypes.Message{
					{Data: testData1,
						Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 2,
						}},
				},
			},
		},
		Report: cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{
				{
					SourceChainSelector: 1,
					// Add some message data that should be removed
					Messages: []cciptypes.Message{
						{
							Data: testData2,
							Header: cciptypes.RampMessageHeader{
								SourceChainSelector: 2,
							},
						}},
				},
			},
		},
	}

	// Get the log format
	logFormatOutcome := testOutcome.ToLogFormat()

	lggr, _ := logger.New()
	lggr.Infow("struct",
		"outcome", testOutcome)
	lggr.Infow("logFrmt",
		"outcome", logFormatOutcome)

	// Verify the original object is not modified
	require.Equal(t, testData1, testOutcome.CommitReports[0].Messages[0].Data)
	require.Equal(t, testData2, testOutcome.Report.ChainReports[0].Messages[0].Data)

	require.Equal(t, cciptypes.Bytes{}, logFormatOutcome.CommitReports[0].Messages[0].Data)
	require.Equal(t, cciptypes.Bytes{}, logFormatOutcome.Report.ChainReports[0].Messages[0].Data)
}
