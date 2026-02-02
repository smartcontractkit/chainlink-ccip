package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
)

func TestConfigurePool(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(tokenAndPoolReport operations.SequenceReport[tokens.DeployTokenAndPoolInput, seq_core.OnChainOutput]) tokens.ConfigureTokenPoolInput
		expectedErr string
	}{
		{
			desc: "happy path",
			makeInput: func(tokenAndPoolReport operations.SequenceReport[tokens.DeployTokenAndPoolInput, seq_core.OnChainOutput]) tokens.ConfigureTokenPoolInput {
				threshold := big.NewInt(123)
				return tokens.ConfigureTokenPoolInput{
					ChainSelector:                    tokenAndPoolReport.Input.DeployTokenPoolInput.ChainSel,
					TokenPoolAddress:                 common.HexToAddress(tokenAndPoolReport.Output.Addresses[1].Address),
					AdvancedPoolHooks:                common.HexToAddress(tokenAndPoolReport.Output.Addresses[2].Address),
					RouterAddress:                    common.HexToAddress("0x09"),
					ThresholdAmountForAdditionalCCVs: threshold,
					RateLimitAdmin:                   common.HexToAddress("0x10"),
				}
			},
			expectedErr: "",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSel := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSel}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")

			// Deploy chain
			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chainSel,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")
			chainReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSel],
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSel,
					ContractParams: testsetup.CreateBasicContractParams(),
					CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Deploy token and token pool
			tokenAndPoolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployTokenAndPool,
				e.BlockChains.EVMChains()[chainSel],
				basicDeployTokenAndPoolInput(chainReport, false),
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Configure token pool
			input := test.makeInput(tokenAndPoolReport)
			configureReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.ConfigureTokenPool,
				e.BlockChains.EVMChains()[chainSel],
				input,
			)
			if test.expectedErr != "" {
				require.Error(t, err, "ExecuteSequence should error")
				require.Contains(t, err.Error(), test.expectedErr)
				return
			}
			require.NoError(t, err, "ExecuteSequence should not error")

			require.Len(t, configureReport.Output.BatchOps, 1, "Expected 1 batch operation in output")
			require.Len(t, configureReport.Output.BatchOps[0].Transactions, 0, "Expected 0 transactions in batch operation")
			require.Len(t, configureReport.Output.Addresses, 0, "Expected 0 addresses in output")

			// Check dynamic config
			getDynamicConfigReport, err := operations.ExecuteOperation(
				testsetup.BundleWithFreshReporter(e.OperationsBundle),
				token_pool.GetDynamicConfig,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       input.TokenPoolAddress,
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.RouterAddress, getDynamicConfigReport.Output.Router, "Expected router address to be the same as the deployed router")
			require.Equal(t, input.RateLimitAdmin, getDynamicConfigReport.Output.RateLimitAdmin, "Expected rate limit admin address to be the same as the inputted rate limit admin")
		})
	}
}
