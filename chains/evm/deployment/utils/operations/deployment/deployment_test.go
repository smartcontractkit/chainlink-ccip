package deployment_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/deployment"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
)

var TestContractType = cldf_deployment.ContractType("TestContract")

func TestDeploy(t *testing.T) {
	address := common.HexToAddress("0x01")
	validChainSel := uint64(5009297550715157269)
	invalidChainSel := uint64(12345)

	tests := []struct {
		desc        string
		input       deployment.Input[int]
		isZkSyncVM  bool
		expectedErr string
	}{
		{
			desc: "args validation failure",
			input: deployment.Input[int]{
				ChainSelector: validChainSel,
				Args:          3,
			},
			expectedErr: "invalid constructor args for test-deployment: input must be even",
		},
		{
			desc: "mismatched chain selector",
			input: deployment.Input[int]{
				ChainSelector: invalidChainSel,
				Args:          2,
			},
			expectedErr: fmt.Sprintf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", invalidChainSel, validChainSel),
		},
		{
			desc: "zkSyncVM deployment",
			input: deployment.Input[int]{
				ChainSelector: validChainSel,
				Args:          2,
			},
			isZkSyncVM: true,
		},
		{
			desc: "evm deployment",
			input: deployment.Input[int]{
				ChainSelector: validChainSel,
				Args:          2,
			},
			isZkSyncVM: false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			op := deployment.New(
				"test-deployment",
				semver.MustParse("1.0.0"),
				"Test deployment operation",
				TestContractType,
				func(input int) error {
					if input%2 != 0 {
						return fmt.Errorf("input must be even")
					}
					return nil
				},
				deployment.VMDeployers[int]{
					DeployEVM: func(auth *bind.TransactOpts, client bind.ContractBackend, args int) (common.Address, *types.Transaction, error) {
						return address, nil, nil
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
				require.Contains(t, test.expectedErr, err.Error())
			} else {
				require.NoError(t, err, "Unexpected ExecuteOperation error")
				if test.isZkSyncVM {
					require.False(t, confirmed, "Expected transaction to not be confirmed on ZkSyncVM")
				} else {
					require.True(t, confirmed, "Expected transaction to be confirmed on EVM")
				}
				require.Equal(t, validChainSel, report.Output.ChainSelector, "Unexpected ChainSelector in output")
				require.Equal(t, address.Hex(), report.Output.Address, "Unexpected address in output")
				require.Equal(t, TestContractType, report.Output.Type, "Unexpected contract type in output")
				require.Equal(t, semver.MustParse("1.0.0").String(), report.Output.Version, "Unexpected version in output")
			}
		})
	}
}
