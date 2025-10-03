package contract_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
	"github.com/stretchr/testify/require"
)

func TestWriteOutput_Executed(t *testing.T) {
	tests := []struct {
		desc     string
		output   contract.WriteOutput
		expected bool
	}{
		{
			desc: "not executed",
			output: contract.WriteOutput{
				ExecInfo: nil,
			},
			expected: false,
		},
		{
			desc: "executed",
			output: contract.WriteOutput{
				ExecInfo: &contract.ExecInfo{
					Hash: "0xabc123",
				},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			result := test.output.Executed()
			require.Equal(t, test.expected, result)
		})
	}
}
