package consensus_test

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_CommitReportValidator_ExecutePluginCommitData(t *testing.T) {
	tests := []struct {
		name    string
		min     consensus.Threshold
		reports []exectypes.CommitData
		valid   []exectypes.CommitData
	}{
		{
			name:  "empty",
			valid: nil,
		},
		{
			name: "single report, enough observations",
			min:  1,
			reports: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}},
			},
			valid: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}},
			},
		},
		{
			name: "single report, not enough observations",
			min:  2,
			reports: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}},
			},
			valid: nil,
		},
		{
			name: "multiple reports, partial observations",
			min:  2,
			reports: []exectypes.CommitData{
				{MerkleRoot: [32]byte{3}},
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
			},
			valid: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}},
				{MerkleRoot: [32]byte{2}},
			},
		},
		{
			name: "multiple reports for same root",
			min:  2,
			reports: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
				{MerkleRoot: [32]byte{1}, BlockNum: 2},
				{MerkleRoot: [32]byte{1}, BlockNum: 3},
				{MerkleRoot: [32]byte{1}, BlockNum: 4},
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
			},
			valid: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}, BlockNum: 1},
			},
		},
		{
			name: "different executed messages same root",
			min:  2,
			reports: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{2, 3}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{3, 4}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{4, 5}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{5, 6}},
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
			},
			valid: []exectypes.CommitData{
				{MerkleRoot: [32]byte{1}, ExecutedMessages: []cciptypes.SeqNum{1, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Initialize the minObservation
			idFunc := func(data exectypes.CommitData) [32]byte {
				return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
			}
			validator := consensus.NewMinObservation[exectypes.CommitData](tt.min, idFunc)
			for _, report := range tt.reports {
				validator.Add(report)
			}

			// Test the results
			got := validator.GetValid()
			if !assert.ElementsMatch(t, got, tt.valid) {
				t.Errorf("GetValid() = %v, valid %v", got, tt.valid)
			}

			oracleValidator := consensus.NewOracleMinObservation[exectypes.CommitData](tt.min, idFunc)
			for i, report := range tt.reports {
				oracleValidator.Add(report, commontypes.OracleID(i))
			}

			// Test the results
			got = validator.GetValid()
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

	// Initialize the minObservation
	idFunc := func(data Generic) [32]byte {
		return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
	}
	validator := consensus.NewMinObservation[Generic](2, idFunc)

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

	oracleValidator := consensus.NewOracleMinObservation[Generic](2, idFunc)

	oracleValidator.Add(wantValue, 1)
	oracleValidator.Add(wantValue, 2)
	oracleValidator.Add(otherValue, 3)

	// Test the results

	got = validator.GetValid()
	if !assert.ElementsMatch(t, got, wantValid) {
		t.Errorf("GetValid() = %v, valid %v", got, wantValid)
	}
}
