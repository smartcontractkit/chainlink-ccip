package call_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/call"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

var TestContractType = cldf_deployment.ContractType("TestContract")

var OwnerAddress = common.HexToAddress("0x02")

type TestContract struct {
	address common.Address
	owner   common.Address
	value   int
}

func NewTestContract(address common.Address, backend bind.ContractBackend) (*TestContract, error) {
	return &TestContract{
		address: address,
		value:   0,
	}, nil
}

func (t *TestContract) Read(opts *bind.CallOpts, value int) (string, error) {
	if value%2 == 0 {
		return "even", nil
	}
	return "", fmt.Errorf("odd value: %d", value)
}

func (t *TestContract) Write(opts *bind.TransactOpts, value int) (*types.Transaction, error) {
	if value%2 == 0 {
		t.value = value
		return types.NewTx(&types.LegacyTx{
			To:   &t.address,
			Data: []byte{0xDE, 0xAD, 0xBE, 0xEF},
		}), nil
	}
	return &types.Transaction{}, fmt.Errorf("odd value: %d", value)
}

func (t *TestContract) Owner(opts *bind.CallOpts) (common.Address, error) {
	return OwnerAddress, nil
}

func (t *TestContract) Address() common.Address {
	return t.address
}

func TestRead(t *testing.T) {
	address := common.HexToAddress("0x01")
	validChainSel := uint64(5009297550715157269)
	invalidChainSel := uint64(12345)

	tests := []struct {
		desc        string
		input       call.Input[int]
		expectedErr string
	}{
		{
			desc: "valid even input",
			input: call.Input[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
		},
		{
			desc: "invalid odd input",
			input: call.Input[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          3,
			},
			expectedErr: "odd value: 3",
		},
		{
			desc: "mismatched chain selector",
			input: call.Input[int]{
				ChainSelector: invalidChainSel,
				Address:       address,
				Args:          2,
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			read := call.NewRead(
				"test-read",
				semver.MustParse("1.0.0"),
				"Test read operation",
				TestContractType,
				NewTestContract,
				func(contract *TestContract, opts *bind.CallOpts, input int) (string, error) {
					return contract.Read(opts, input)
				},
			)

			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")

			bundle := operations.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				operations.NewMemoryReporter(),
			)

			chain := evm.Chain{
				Selector: validChainSel,
			}

			report, err := operations.ExecuteOperation(bundle, read, chain, test.input)
			if test.expectedErr != "" {
				require.Error(t, err, "Expected ExecuteOperation error but got none")
				require.Contains(t, test.expectedErr, err.Error())
			} else {
				require.NoError(t, err, "Unexpected ExecuteOperation error")
				require.Equal(t, report.Output, "even")
			}
		})
	}
}

func TestWrite(t *testing.T) {
	address := common.HexToAddress("0x01")
	validChainSel := uint64(5009297550715157269)
	invalidChainSel := uint64(12345)

	tests := []struct {
		desc            string
		input           call.Input[int]
		deployerAddress common.Address
		expectedErr     string
	}{
		{
			desc: "args validation failure",
			input: call.Input[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          3,
			},
			expectedErr: "invalid args for test-write: input must be even",
		},
		{
			desc: "mismatched chain selector",
			input: call.Input[int]{
				ChainSelector: invalidChainSel,
				Address:       address,
				Args:          2,
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
		{
			desc: "called by owner",
			input: call.Input[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
			deployerAddress: OwnerAddress,
		},
		{
			desc: "not called by owner",
			input: call.Input[int]{
				ChainSelector: validChainSel,
				Address:       address,
				Args:          2,
			},
			deployerAddress: common.HexToAddress("0x03"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			write := call.NewWrite(
				"test-write",
				semver.MustParse("1.0.0"),
				"Test write operation",
				TestContractType,
				"", // ABI not used in test
				NewTestContract,
				call.OnlyOwner,
				func(input int) error {
					if input%2 != 0 {
						return fmt.Errorf("input must be even")
					}
					return nil
				},
				func(contract *TestContract, opts *bind.TransactOpts, input int) (*types.Transaction, error) {
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

			chain := evm.Chain{
				Selector: validChainSel,
				DeployerKey: &bind.TransactOpts{
					From: test.deployerAddress,
				},
				Confirm: func(tx *types.Transaction) (uint64, error) {
					return 1, nil
				},
			}

			report, err := operations.ExecuteOperation(bundle, write, chain, test.input)
			if test.expectedErr != "" {
				require.Error(t, err, "Expected ExecuteOperation error but got none")
				require.Contains(t, test.expectedErr, err.Error())
			} else {
				require.NoError(t, err, "Unexpected ExecuteOperation error")
				if test.deployerAddress == OwnerAddress {
					require.True(t, report.Output.Executed, "Expected Executed to be true when called by owner")
				} else {
					require.False(t, report.Output.Executed, "Expected Executed to be false when not called by owner")
				}
				require.Equal(t, validChainSel, report.Output.ChainSelector, "Unexpected ChainSelector in output")
				require.Equal(t, []byte{0xDE, 0xAD, 0xBE, 0xEF}, report.Output.Tx.Data, "Unexpected tx data in output")
				require.Equal(t, address.Hex(), report.Output.Tx.To, "Unexpected to address in output")
				require.Equal(t, string(TestContractType), report.Output.Tx.ContractType, "Unexpected ContractType in output")
			}
		})
	}
}
