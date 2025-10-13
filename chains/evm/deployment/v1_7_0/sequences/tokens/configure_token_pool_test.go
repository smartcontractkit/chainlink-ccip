package tokens_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestConfigurePool(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(tokenAndPoolReport operations.SequenceReport[tokens.DeployBurnMintTokenAndPoolInput, seq_core.OnChainOutput]) tokens.ConfigureTokenPoolInput
		expectedErr string
	}{
		{
			desc: "happy path",
			makeInput: func(tokenAndPoolReport operations.SequenceReport[tokens.DeployBurnMintTokenAndPoolInput, seq_core.OnChainOutput]) tokens.ConfigureTokenPoolInput {
				threshold := big.NewInt(123)
				return tokens.ConfigureTokenPoolInput{
					ChainSelector:    tokenAndPoolReport.Input.DeployTokenPoolInput.ChainSel,
					TokenPoolAddress: common.HexToAddress(tokenAndPoolReport.Output.Addresses[1].Address),
					AllowList: []common.Address{
						common.HexToAddress("0x07"),
						common.HexToAddress("0x08"),
					},
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
			require.Zero(t, getDynamicConfigReport.Output.ThresholdAmountForAdditionalCCVs.Cmp(input.ThresholdAmountForAdditionalCCVs))

			// Check allowlist
			getAllowlistReport, err := operations.ExecuteOperation(
				testsetup.BundleWithFreshReporter(e.OperationsBundle),
				token_pool.GetAllowList,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       input.TokenPoolAddress,
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.AllowList, getAllowlistReport.Output, "Expected allowlist address to be the same as the deployed allowlist")

			// Check rate limit admin
			getRateLimitAdminReport, err := operations.ExecuteOperation(
				testsetup.BundleWithFreshReporter(e.OperationsBundle),
				token_pool.GetRateLimitAdmin,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[any]{
					ChainSelector: chainSel,
					Address:       input.TokenPoolAddress,
				},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, input.RateLimitAdmin, getRateLimitAdminReport.Output, "Expected rate limit admin address to be the same as the deployed rate limit admin")
		})
	}
}
