package contract

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
)

func TestDeploy(t *testing.T) {
	address := common.HexToAddress("0x01")
	validChainSel := uint64(5009297550715157269)
	invalidChainSel := uint64(12345)

	type ConstructorArgs struct {
		Value int
	}

	tests := []struct {
		desc        string
		input       DeployInput[ConstructorArgs]
		isZkSyncVM  bool
		expectedErr string
	}{
		{
			desc: "args validation failure",
			input: DeployInput[ConstructorArgs]{
				ChainSelector: validChainSel,
				Args:          ConstructorArgs{Value: 3},
			},
			expectedErr: "invalid constructor args for test-deployment: input must be even",
		},
		{
			desc: "mismatched chain selector",
			input: DeployInput[ConstructorArgs]{
				ChainSelector: invalidChainSel,
				Args:          ConstructorArgs{Value: 2},
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
		{
			desc: "bytecode not defined for version",
			input: DeployInput[ConstructorArgs]{
				ChainSelector:  validChainSel,
				Args:           ConstructorArgs{Value: 2},
				TypeAndVersion: deployment.NewTypeAndVersion(testContractType, *semver.MustParse("1.1.0")),
			},
			expectedErr: fmt.Sprintf("no bytecode defined for %s %s", testContractType, "1.1.0"),
		},
		{
			desc: "revert from contract",
			input: DeployInput[ConstructorArgs]{
				ChainSelector:  validChainSel,
				Args:           ConstructorArgs{Value: 10},
				TypeAndVersion: deployment.NewTypeAndVersion(testContractType, *semver.MustParse("1.0.0")),
			},
			expectedErr: "due to error -`InvalidValue` args [1]: 6072742c0000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			desc: "zkSyncVM deployment",
			input: DeployInput[ConstructorArgs]{
				ChainSelector:  validChainSel,
				Args:           ConstructorArgs{Value: 2},
				TypeAndVersion: deployment.NewTypeAndVersion(testContractType, *semver.MustParse("1.0.0")),
			},
			isZkSyncVM: true,
		},
		{
			desc: "evm deployment",
			input: DeployInput[ConstructorArgs]{
				ChainSelector:  validChainSel,
				Args:           ConstructorArgs{Value: 2},
				TypeAndVersion: deployment.NewTypeAndVersion(testContractType, *semver.MustParse("1.0.0")),
			},
			isZkSyncVM: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			contractABI := `[{
				"inputs": [{"name": "value", "type": "uint256"}],
				"name": "InvalidValue",
				"type": "error"
			}]`

			op := NewDeploy(DeployParams[ConstructorArgs]{
				Name:        "test-deployment",
				Version:     semver.MustParse("1.0.0"),
				Description: "Test deployment operation",
				ContractMetadata: &bind.MetaData{
					ABI: contractABI,
				},
				BytecodeByTypeAndVersion: map[string]Bytecode{
					deployment.NewTypeAndVersion(testContractType, *semver.MustParse("1.0.0")).String(): {EVM: []byte{}},
				},
				Validate: func(input ConstructorArgs) error {
					if input.Value%2 != 0 {
						return fmt.Errorf("input must be even")
					}
					return nil
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
				Confirm: func(tx *types.Transaction) (uint64, error) {
					confirmed = true
					return 1, nil
				},
				IsZkSyncVM: test.isZkSyncVM,
			}

			deployZkContract = func(
				_ *accounts.TransactOpts,
				_ []byte,
				_ *clients.Client,
				_ *accounts.Wallet,
				_ *abi.ABI,
				_ ...interface{},
			) (common.Address, error) {
				return address, nil
			}
			deployEVMContract = func(
				_ *bind.TransactOpts,
				_ abi.ABI,
				_ []byte,
				_ bind.ContractBackend,
				params ...interface{},
			) (common.Address, *types.Transaction, *bind.BoundContract, error) {
				// Not caught by operation validation, revert reason should be surfaced
				if params[0] == 10 {
					return address, &types.Transaction{}, nil, &rpcError{
						Data: common.Bytes2Hex(append(
							crypto.Keccak256([]byte("InvalidValue(uint256)"))[:4],
							common.LeftPadBytes([]byte{1}, 32)...,
						)),
					}
				}

				return address, types.NewTx(&types.LegacyTx{
					To:   &address,
					Data: []byte{0xDE, 0xAD, 0xBE, 0xEF},
				}), nil, nil
			}

			report, err := operations.ExecuteOperation(bundle, op, chain, test.input)
			if test.expectedErr != "" {
				require.Error(t, err, "Expected ExecuteOperation error but got none")
				require.Contains(t, err.Error(), test.expectedErr)
			} else {
				require.NoError(t, err, "Unexpected ExecuteOperation error")
				if test.isZkSyncVM {
					require.False(t, confirmed, "Expected transaction to not be confirmed on ZkSyncVM")
				} else {
					require.True(t, confirmed, "Expected transaction to be confirmed on EVM")
				}
				require.Equal(t, validChainSel, report.Output.ChainSelector, "Unexpected ChainSelector in output")
				require.Equal(t, address.Hex(), report.Output.Address, "Unexpected address in output")
				require.Equal(t, datastore.ContractType(testContractType), report.Output.Type, "Unexpected contract type in output")
				require.Equal(t, semver.MustParse("1.0.0"), report.Output.Version, "Unexpected version in output")
			}
		})
	}
}
