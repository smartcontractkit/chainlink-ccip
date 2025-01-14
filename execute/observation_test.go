package execute

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
)

func Test_filterCommitReports(t *testing.T) {
	tests := []struct {
		name      string
		msgObs    exectypes.MessageObservations
		commitObs exectypes.CommitObservations
		want      exectypes.CommitObservations
	}{
		{
			name:   "no messages",
			msgObs: exectypes.MessageObservations{},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
					},
				},
			},
			want: exectypes.CommitObservations{},
		},
		{
			name: "missing messages",
			msgObs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
				},
			},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
			want: exectypes.CommitObservations{},
		},
		{
			name: "valid messages",
			msgObs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
					3: cciptypes.Message{},
				},
			},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
			want: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
		},
		{
			name: "mixed valid and invalid messages",
			msgObs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
					3: cciptypes.Message{},
				},
			},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(4, 5),
					},
				},
			},
			want: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := filterCommitReports(tt.msgObs, tt.commitObs)
			assert.Equal(t, tt.want, got)
		})
	}
}
