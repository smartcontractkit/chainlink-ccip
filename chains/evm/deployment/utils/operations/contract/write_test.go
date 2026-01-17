package contract_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

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
		input           contract.FunctionInput[int]
		deployerAddress common.Address
		expectedErr     string
	}{
		{
			desc: "args validation failure",
			input: contract.FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          3,
			},
			expectedErr: "invalid args for test-write: input must be even",
		},
		{
			desc: "revert from contract",
			input: contract.FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          10,
			},
			deployerAddress: OwnerAddress,
			expectedErr:     "due to error -`InvalidValue` args [1]: 6072742c0000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			desc: "mismatched chain selector",
			input: contract.FunctionInput[int]{
				ChainSelector: invalidChainSel,
				Address:       address,
				Args:          2,
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
		{
			desc: "called by owner",
			input: contract.FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
			deployerAddress: OwnerAddress,
		},
		{
			desc: "not called by owner",
			input: contract.FunctionInput[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
			deployerAddress: common.HexToAddress("0x03"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			write := contract.NewWrite(
				"test-write",
				semver.MustParse("1.0.0"),
				"Test write operation",
				testContractType,
				contractABI,
				newTestContract,
				contract.OnlyOwner,
				func(input int) error {
					if input%2 != 0 {
						return fmt.Errorf("input must be even")
					}
					return nil
				},
				func(contract *testContract, opts *bind.TransactOpts, input int) (*types.Transaction, error) {
					return contract.Write(opts, input)
				},
			)

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
					require.True(t, report.Output.Executed, "Expected Executed to be true when called by owner")
				} else {
					require.False(t, confirmed, "Expected transaction to not be confirmed when not called by owner")
					require.False(t, report.Output.Executed, "Expected Executed to be false when not called by owner")
				}
				require.Equal(t, validChainSel, report.Output.ChainSelector, "Unexpected ChainSelector in output")
				require.Equal(t, []byte{0xDE, 0xAD, 0xBE, 0xEF}, report.Output.Tx.Data, "Unexpected tx data in output")
				require.Equal(t, address.Hex(), report.Output.Tx.To, "Unexpected to address in output")
				require.Equal(t, string(testContractType), report.Output.Tx.ContractType, "Unexpected ContractType in output")
			}
		})
	}
}
