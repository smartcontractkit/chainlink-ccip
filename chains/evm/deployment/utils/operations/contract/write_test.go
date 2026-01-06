package contract

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"
)

func TestWriteOutput_Executed(t *testing.T) {
	tests := []struct {
		desc     string
		output   WriteOutput
		expected bool
	}{
		{
			desc: "not executed",
			output: WriteOutput{
				ExecInfo: nil,
			},
			expected: false,
		},
		{
			desc: "executed",
			output: WriteOutput{
				ExecInfo: &ExecInfo{
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

func TestWrite(t *testing.T) {
	address := common.HexToAddress("0x01")
	validChainSel := uint64(5009297550715157269)
	invalidChainSel := uint64(12345)

	contractABI := `[{
		"inputs": [{"name": "value", "type": "uint256"}],
		"name": "InvalidValue",
		"type": "error"
	}]`

	tests := []struct {
		desc            string
		input           FunctionInput[int]
		deployerAddress common.Address
		expectedErr     string
	}{
		{
			desc: "args validation failure",
			input: FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          3,
			},
			expectedErr: "invalid args for test-write: input must be even",
		},
		{
			desc: "revert from contract",
			input: FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          10,
			},
			deployerAddress: OwnerAddress,
			expectedErr:     "due to error -`InvalidValue` args [1]: 6072742c0000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			desc: "mismatched chain selector",
			input: FunctionInput[int]{
				ChainSelector: invalidChainSel,
				Address:       address,
				Args:          2,
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
		{
			desc: "called by owner",
			input: FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
			deployerAddress: OwnerAddress,
		},
		{
			desc: "not called by owner",
			input: FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
			deployerAddress: common.HexToAddress("0x03"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			write := NewWrite(WriteParams[int, *testContract]{
				Name:            "test-write",
				Version:         semver.MustParse("1.0.0"),
				Description:     "Test write operation",
				ContractType:    testContractType,
				ContractABI:     contractABI,
				NewContract:     newTestContract,
				IsAllowedCaller: OnlyOwner[*testContract, int],
				Validate: func(input int) error {
					if input%2 != 0 {
						return fmt.Errorf("input must be even")
					}
					return nil
				},
				CallContract: func(contract *testContract, opts *bind.TransactOpts, input int) (*types.Transaction, error) {
					return contract.Write(opts, input)
				},
			})

			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")

			bundle := operations.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				operations.NewMemoryReporter(),
			)

			var confirmed bool
			chain := evm.Chain{
				Selector: validChainSel,
				DeployerKey: &bind.TransactOpts{
					From: test.deployerAddress,
				},
				Confirm: func(tx *types.Transaction) (uint64, error) {
					confirmed = true
					return 1, nil
				},
			}

			report, err := operations.ExecuteOperation(bundle, write, chain, test.input)
			if test.expectedErr != "" {
				require.Error(t, err, "Expected ExecuteOperation error but got none")
				require.Contains(t, err.Error(), test.expectedErr)
			} else {
				require.NoError(t, err, "Unexpected ExecuteOperation error")
				if test.deployerAddress == OwnerAddress {
					require.True(t, confirmed, "Expected transaction to be confirmed when called by owner")
					require.True(t, report.Output.Executed(), "Expected Executed to be true when called by owner")
				} else {
					require.False(t, confirmed, "Expected transaction to not be confirmed when not called by owner")
					require.False(t, report.Output.Executed(), "Expected Executed to be false when not called by owner")
				}
				require.Equal(t, validChainSel, report.Output.ChainSelector, "Unexpected ChainSelector in output")
				require.Equal(t, []byte{0xDE, 0xAD, 0xBE, 0xEF}, report.Output.Tx.Data, "Unexpected tx data in output")
				require.Equal(t, address.Hex(), report.Output.Tx.To, "Unexpected to address in output")
				require.Equal(t, string(testContractType), report.Output.Tx.ContractType, "Unexpected ContractType in output")
			}
		})
	}
}

func TestBatchOperationFromWrites(t *testing.T) {
	tests := []struct {
		desc        string
		outputs     []WriteOutput
		expected    mcms_types.BatchOperation
		expectedErr string
	}{
		{
			desc: "single output",
			outputs: []WriteOutput{
				{
					ChainSelector: 5009297550715157269,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
			},
			expected: mcms_types.BatchOperation{
				ChainSelector: 5009297550715157269,
				Transactions: []mcms_types.Transaction{
					{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
			},
		},
		{
			desc: "multiple outputs same chain",
			outputs: []WriteOutput{
				{
					ChainSelector: 5009297550715157269,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
				{
					ChainSelector: 5009297550715157269,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x02").Hex(),
						Data:             common.Hex2Bytes("0xcafebabe"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
			},
			expected: mcms_types.BatchOperation{
				ChainSelector: 5009297550715157269,
				Transactions: []mcms_types.Transaction{
					{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
					{
						To:               common.HexToAddress("0x02").Hex(),
						Data:             common.Hex2Bytes("0xcafebabe"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
			},
		},
		{
			desc: "multiple outputs different chains",
			outputs: []WriteOutput{
				{
					ChainSelector: 5009297550715157269,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
				{
					ChainSelector: 4340886533089894000,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x02").Hex(),
						Data:             common.Hex2Bytes("0xcafebabe"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
				},
			},
			expected:    mcms_types.BatchOperation{},
			expectedErr: "writes target multiple chains",
		},
		{
			desc:     "no outputs",
			outputs:  []WriteOutput{},
			expected: mcms_types.BatchOperation{},
		},
		{
			desc: "all executed outputs",
			outputs: []WriteOutput{
				{
					ChainSelector: 5009297550715157269,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x01").Hex(),
						Data:             common.Hex2Bytes("0xdeadbeef"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
					ExecInfo: &ExecInfo{
						Hash: "0xabc123",
					},
				},
				{
					ChainSelector: 5009297550715157269,
					Tx: mcms_types.Transaction{
						To:               common.HexToAddress("0x02").Hex(),
						Data:             common.Hex2Bytes("0xcafebabe"),
						AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
					},
					ExecInfo: &ExecInfo{
						Hash: "0xdef456",
					},
				},
			},
			expected: mcms_types.BatchOperation{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			batchOp, err := NewBatchOperationFromWrites(test.outputs)
			if test.expectedErr != "" {
				require.Error(t, err, "Expected error from NewBatchOperationFromWrites")
				require.Contains(t, err.Error(), test.expectedErr, "Unexpected error message")
				return
			}
			require.NoError(t, err, "Unexpected error from NewBatchOperationFromWrites")
			require.Equal(t, test.expected, batchOp, "Unexpected BatchOperation result")
		})
	}
}
