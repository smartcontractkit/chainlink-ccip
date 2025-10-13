package tokens_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestRegisterToken(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(chainSel uint64, tokenAddress, tokenPoolAddress, tokenAdminRegistryAddress common.Address) tokens.RegisterTokenInput
		expectedErr string
	}{
		{
			desc: "happy path - no external admin",
			makeInput: func(chainSel uint64, tokenAddress, tokenPoolAddress, tokenAdminRegistryAddress common.Address) tokens.RegisterTokenInput {
				return tokens.RegisterTokenInput{
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
			makeInput: func(chainSel uint64, tokenAddress, tokenPoolAddress, tokenAdminRegistryAddress common.Address) tokens.RegisterTokenInput {
				return tokens.RegisterTokenInput{
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

			// Deploy chain
			chainReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSel],
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSel,
					ContractParams: testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Deploy token and token pool
			tokenAndPoolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployBurnMintTokenAndPool,
				e.BlockChains.EVMChains()[chainSel],
				basicDeployBurnMintTokenAndPoolInput(chainReport),
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			tokenAddress := tokenAndPoolReport.Output.Addresses[0].Address
			tokenPoolAddress := tokenAndPoolReport.Output.Addresses[1].Address

			// Register token
			var tokenAdminRegistryAddress string
			for _, addr := range chainReport.Output.Addresses {
				if addr.Type == datastore.ContractType(token_admin_registry.ContractType) {
					tokenAdminRegistryAddress = addr.Address
				}
			}
			require.NotEmpty(t, tokenAdminRegistryAddress, "Token admin registry address should be found")

			input := test.makeInput(
				chainSel,
				common.HexToAddress(tokenAddress),
				common.HexToAddress(tokenPoolAddress),
				common.HexToAddress(tokenAdminRegistryAddress),
			)
			_, err = operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.RegisterToken,
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
					Address:       common.HexToAddress(tokenAdminRegistryAddress),
					Args:          common.HexToAddress(tokenAddress),
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
