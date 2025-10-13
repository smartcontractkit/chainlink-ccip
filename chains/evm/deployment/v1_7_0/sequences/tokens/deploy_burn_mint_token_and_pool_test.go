package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	token_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc677"
	"github.com/stretchr/testify/require"
)

func basicDeployBurnMintTokenAndPoolInput(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployBurnMintTokenAndPoolInput {
	var rmnProxyAddress common.Address
	var routerAddress common.Address
	for _, addr := range chainReport.Output.Addresses {
		if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
			rmnProxyAddress = common.HexToAddress(addr.Address)
		}
		if addr.Type == datastore.ContractType(router.ContractType) {
			routerAddress = common.HexToAddress(addr.Address)
		}
	}
	return tokens.DeployBurnMintTokenAndPoolInput{
		Accounts: map[common.Address]*big.Int{
			common.HexToAddress("0x01"): big.NewInt(500_000),
			common.HexToAddress("0x02"): big.NewInt(500_000),
		},
		TokenInfo: tokens.TokenInfo{
			Decimals:  18,
			MaxSupply: big.NewInt(1_000_000),
			Name:      "Test Token",
		},
		DeployTokenPoolInput: tokens.DeployTokenPoolInput{
			ChainSel:         chainReport.Input.ChainSelector,
			TokenPoolType:    datastore.ContractType(burn_mint_token_pool.ContractType),
			TokenPoolVersion: semver.MustParse("1.7.0"),
			TokenSymbol:      "TEST",
			RateLimitAdmin:   common.HexToAddress("0x01"),
			ConstructorArgs: token_pool.ConstructorArgs{
				LocalTokenDecimals: 18,
				Allowlist: []common.Address{
					common.HexToAddress("0x02"),
				},
				RMNProxy: rmnProxyAddress,
				Router:   routerAddress,
			},
		},
	}
}

func TestDeployBurnMintTokenAndPool(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployBurnMintTokenAndPoolInput
		expectedErr string
	}{
		{
			desc:        "happy path",
			makeInput:   basicDeployBurnMintTokenAndPoolInput,
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
			input := test.makeInput(chainReport)
			poolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployBurnMintTokenAndPool,
				e.BlockChains.EVMChains()[chainSel],
				input,
			)
			if test.expectedErr != "" {
				require.Error(t, err, "ExecuteSequence should error")
				require.Contains(t, err.Error(), test.expectedErr)
				return
			}
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, poolReport.Output.BatchOps, 1, "Expected 1 batch operation in output")
			require.Len(t, poolReport.Output.BatchOps[0].Transactions, 0, "Expected 0 transactions in batch operation")
			require.Len(t, poolReport.Output.Addresses, 2, "Expected 2 addresses in output")
			tokenAddress := poolReport.Output.Addresses[0].Address
			poolAddress := poolReport.Output.Addresses[1].Address

			// Check token metadata
			token, err := token_bindings.NewBurnMintERC677(common.HexToAddress(tokenAddress), e.BlockChains.EVMChains()[chainSel].Client)
			require.NoError(t, err, "NewBurnMintERC677 should not error")
			name, err := token.Name(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "Name should not error")
			require.Equal(t, input.TokenInfo.Name, name, "Expected token name to be the same as the deployed token")
			symbol, err := token.Symbol(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "Symbol should not error")
			require.Equal(t, input.DeployTokenPoolInput.TokenSymbol, symbol, "Expected token symbol to be the same as the deployed token")
			decimals, err := token.Decimals(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "Decimals should not error")
			require.Equal(t, input.TokenInfo.Decimals, decimals, "Expected token decimals to be the same as the deployed token")
			totalSupply, err := token.TotalSupply(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "TotalSupply should not error")
			require.Equal(t, input.TokenInfo.MaxSupply, totalSupply, "Expected token total supply to be the same as the deployed token")

			// Check token minters
			minters, err := token.GetMinters(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "GetMinters should not error")
			require.Equal(t, []common.Address{common.HexToAddress(poolAddress)}, minters, "Expected token pool to be the minter of the token")

			// Check token burners
			burners, err := token.GetBurners(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "GetBurners should not error")
			require.Equal(t, []common.Address{common.HexToAddress(poolAddress)}, burners, "Expected token pool to be the burner of the token")
		})
	}
}
