package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func Test_CommitReportValidator_ExecutePluginCommitData(t *testing.T) {
	tests := []struct {
		name    string
		min     int
		reports []plugintypes.ExecutePluginCommitData
		valid   []plugintypes.ExecutePluginCommitData
	}{
		{
			name:  "empty",
			valid: nil,
		},
		{
			name: "single report, enough observations",
			min:  1,
			reports: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
			},
			valid: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
			},
		},
		{
			name: "single report, not enough observations",
			min:  2,
			reports: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
			},
			valid: nil,
		},
		{
			name: "multiple reports, partial observations",
			min:  2,
			reports: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{3}},
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
			},
			valid: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
			},
		},
		{
			name: "multiple reports for same root",
			min:  2,
			reports: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
				{MerkleRoot: [32]byte{1}, BlockNum: 2},
				{MerkleRoot: [32]byte{1}, BlockNum: 3},
				{MerkleRoot: [32]byte{1}, BlockNum: 4},
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
			},
			valid: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
			},
		},
		{
			name: "different executed messages same root",
			min:  2,
			reports: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{2, 3}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{3, 4}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{4, 5}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{5, 6}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
			},
			valid: []plugintypes.ExecutePluginCommitData{
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Initialize the minObservationValidator
			idFunc := func(data plugintypes.ExecutePluginCommitData) [32]byte {
				return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
			}
			validator := NewMinObservationValidator[plugintypes.ExecutePluginCommitData](tt.min, idFunc)
			for _, report := range tt.reports {
				validator.Add(report)
			}

			// Test the results
			got := validator.GetValid()
			if !assert.ElementsMatch(t, got, tt.valid) {
				t.Errorf("GetValid() = %v, valid %v", got, tt.valid)
			}
		})
	}
}

func Test_CommitReportValidator_Generics(t *testing.T) {
	type Generic struct {
		number int
	}

	// Initialize the minObservationValidator
	idFunc := func(data Generic) [32]byte {
		return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
	}
	validator := NewMinObservationValidator[Generic](2, idFunc)

	wantValue := Generic{number: 1}
	otherValue := Generic{number: 2}

	validator.Add(wantValue)
	validator.Add(wantValue)
	validator.Add(otherValue)

	// Test the results

	wantValid := []Generic{wantValue}
	got := validator.GetValid()
	if !assert.ElementsMatch(t, got, wantValid) {
		t.Errorf("GetValid() = %v, valid %v", got, wantValid)
	}
}
