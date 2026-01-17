package contract_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
)

func TestDeploy(t *testing.T) {
	address := common.HexToAddress("0x01")
	validChainSel := uint64(5009297550715157269)
	invalidChainSel := uint64(12345)

	tests := []struct {
		desc        string
		input       contract.DeployInput[int]
		isZkSyncVM  bool
		expectedErr string
	}{
		{
			desc: "args validation failure",
			input: contract.DeployInput[int]{
				ChainSelector: validChainSel,
				Args:          3,
			},
			expectedErr: "invalid constructor args for test-deployment: input must be even",
		},
		{
			desc: "revert from contract",
			input: contract.DeployInput[int]{
				ChainSelector: validChainSel,
				Args:          10,
			},
			expectedErr: "due to error -`InvalidValue` args [1]: 6072742c0000000000000000000000000000000000000000000000000000000000000001",
		},
		{
			desc: "mismatched chain selector",
			input: contract.DeployInput[int]{
				ChainSelector: invalidChainSel,
				Args:          2,
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
		{
			desc: "zkSyncVM deployment",
			input: contract.DeployInput[int]{
				ChainSelector: validChainSel,
				Args:          2,
			},
			isZkSyncVM: true,
		},
		{
			desc: "evm deployment",
			input: contract.DeployInput[int]{
				ChainSelector: validChainSel,
				Args:          2,
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

			op := contract.NewDeploy(
				"test-deployment",
				semver.MustParse("1.0.0"),
				"Test deployment operation",
				testContractType,
				contractABI,
				func(input int) error {
					if input%2 != 0 {
						return fmt.Errorf("input must be even")
					}
					return nil
				},
				contract.VMDeployers[int]{
					DeployEVM: func(auth *bind.TransactOpts, client bind.ContractBackend, args int) (common.Address, *types.Transaction, error) {
						// Not caught by operation validation, revert reason should be surfaced
						if args == 10 {
							return address, &types.Transaction{}, &rpcError{
								Data: common.Bytes2Hex(append(
									crypto.Keccak256([]byte("InvalidValue(uint256)"))[:4],
									common.LeftPadBytes([]byte{1}, 32)...,
								)),
							}
						}

						return address, types.NewTx(&types.LegacyTx{
							To:   &address,
							Data: []byte{0xDE, 0xAD, 0xBE, 0xEF},
						}), nil
					},
					DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args int) (common.Address, error) {
						return address, nil
					},
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
				Confirm: func(tx *types.Transaction) (uint64, error) {
					confirmed = true
					return 1, nil
				},
				IsZkSyncVM: test.isZkSyncVM,
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
