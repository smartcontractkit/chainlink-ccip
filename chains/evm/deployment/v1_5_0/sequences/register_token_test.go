package sequences_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	"github.com/stretchr/testify/require"
)

func TestRegisterToken(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(chainSel uint64, tokenAddress, tokenPoolAddress, tokenAdminRegistryAddress common.Address) sequences.RegisterTokenInput
		expectedErr string
	}{
		{
			desc: "happy path - no external admin",
			makeInput: func(chainSel uint64, tokenAddress, tokenPoolAddress, tokenAdminRegistryAddress common.Address) sequences.RegisterTokenInput {
				return sequences.RegisterTokenInput{
					ChainSelector:             chainSel,
					TokenAddress:              tokenAddress,
					TokenPoolAddress:          tokenPoolAddress,
					ExternalAdmin:             common.Address{},
					TokenAdminRegistryAddress: tokenAdminRegistryAddress,
				}
			},
			expectedErr: "",
		},
		{
			desc: "happy path - external admin",
			makeInput: func(chainSel uint64, tokenAddress, tokenPoolAddress, tokenAdminRegistryAddress common.Address) sequences.RegisterTokenInput {
				return sequences.RegisterTokenInput{
					ChainSelector:             chainSel,
					TokenAddress:              tokenAddress,
					TokenPoolAddress:          tokenPoolAddress,
					ExternalAdmin:             common.HexToAddress("0x07"),
					TokenAdminRegistryAddress: tokenAdminRegistryAddress,
				}
			},
			expectedErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSel := uint64(5009297550715157269)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSel}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")

			// Deploy token admin registry
			tokenAdminRegistryReport, err := operations.ExecuteOperation(
				e.OperationsBundle,
				token_admin_registry.Deploy,
				e.BlockChains.EVMChains()[chainSel],
				evm_contract.DeployInput[token_admin_registry.ConstructorArgs]{
					ChainSelector:  chainSel,
					TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			tokenAdminRegistryAddress := common.HexToAddress(tokenAdminRegistryReport.Output.Address)

			// Deploy token
			tokenAddress, tx, _, err := burn_mint_erc20.DeployBurnMintERC20(
				e.BlockChains.EVMChains()[chainSel].DeployerKey,
				e.BlockChains.EVMChains()[chainSel].Client,
				"Test Token",
				"TEST",
				18,
				big.NewInt(1000000000000000000),
				big.NewInt(0),
			)
			require.NoError(t, err, "DeployBurnMintERC20 should not error")
			_, err = e.BlockChains.EVMChains()[chainSel].Confirm(tx)
			require.NoError(t, err, "Confirm should not error")

			// Deploy token pool
			tokenPoolAddress, tx, _, err := burn_mint_token_pool.DeployBurnMintTokenPool(
				e.BlockChains.EVMChains()[chainSel].DeployerKey,
				e.BlockChains.EVMChains()[chainSel].Client,
				tokenAddress,
				18,
				[]common.Address{},
				common.HexToAddress("0x01"),
				common.HexToAddress("0x02"),
			)

			input := test.makeInput(
				chainSel,
				tokenAddress,
				tokenPoolAddress,
				tokenAdminRegistryAddress,
			)
			_, err = operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.RegisterToken,
				e.BlockChains.EVMChains()[chainSel],
				input,
			)
			if test.expectedErr != "" {
				require.Error(t, err, "ExecuteSequence should error")
				require.Contains(t, err.Error(), test.expectedErr)
				return
			}
			require.NoError(t, err, "ExecuteSequence should not error")

			// Checks
			tokenConfigReport, err := operations.ExecuteOperation(
				testsetup.BundleWithFreshReporter(e.OperationsBundle),
				token_admin_registry.GetTokenConfig,
				e.BlockChains.EVMChains()[chainSel],
				evm_contract.FunctionInput[common.Address]{
					ChainSelector: chainSel,
					Address:       tokenAdminRegistryAddress,
					Args:          tokenAddress,
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			if input.ExternalAdmin != (common.Address{}) {
				// We can propose an external admin, but we can't accept ownership or set the pool address since we don't control the admin.
				require.Equal(t, input.ExternalAdmin, tokenConfigReport.Output.PendingAdministrator)
				require.Equal(t, (common.Address{}), tokenConfigReport.Output.Administrator)
				require.Equal(t, (common.Address{}), tokenConfigReport.Output.TokenPool)
			} else {
				// No external admin means that the owner of the token admin registry will be proposed as the admin.
				// Since the deployer key is the owner of the token admin registry, it can accept admin rights and set the pool address.
				require.Equal(t, (common.Address{}), tokenConfigReport.Output.PendingAdministrator)
				require.Equal(t, e.BlockChains.EVMChains()[chainSel].DeployerKey.From, tokenConfigReport.Output.Administrator)
				require.Equal(t, input.TokenPoolAddress, tokenConfigReport.Output.TokenPool)
			}
		})
	}
}
