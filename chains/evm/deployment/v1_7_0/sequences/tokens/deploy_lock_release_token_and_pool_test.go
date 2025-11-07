package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/lock_release_token_pool"
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

func basicDeployLockReleaseTokenAndPoolInput(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
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
	return tokens.DeployLockReleaseTokenAndPoolInput{
		Accounts: map[common.Address]*big.Int{
			common.HexToAddress("0x01"): big.NewInt(300_000),
			common.HexToAddress("0x02"): big.NewInt(300_000),
		},
		TokenInfo: tokens.TokenInfo{
			Decimals:  18,
			MaxSupply: big.NewInt(2_000_000),
			Name:      "Test Token",
		},
		DeployTokenPoolInput: tokens.DeployTokenPoolInput{
			ChainSel:         chainReport.Input.ChainSelector,
			TokenPoolType:    datastore.ContractType(lock_release_token_pool.ContractType),
			TokenPoolVersion: semver.MustParse("1.7.0"),
			TokenSymbol:      "TEST",
			RateLimitAdmin:   common.Address{},
			ConstructorArgs: token_pool.ConstructorArgs{
				Token:              common.Address{},
				LocalTokenDecimals: 18,
				Allowlist:          []common.Address{},
				RMNProxy:           rmnProxyAddress,
				Router:             routerAddress,
			},
		},
		PoolFundingAmount: big.NewInt(100_000),
	}
}

func TestDeployLockReleaseTokenAndPool(t *testing.T) {
	tests := []struct {
		name      string
		makeInput func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput
	}{
		{
			name: "basic deployment",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
				return basicDeployLockReleaseTokenAndPoolInput(chainReport)
			},
		},
		{
			name: "deployment without pool funding",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
				input := basicDeployLockReleaseTokenAndPoolInput(chainReport)
				input.PoolFundingAmount = nil
				return input
			},
		},
		{
			name: "deployment with zero pool funding",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
				input := basicDeployLockReleaseTokenAndPoolInput(chainReport)
				input.PoolFundingAmount = big.NewInt(0)
				return input
			},
		},
		{
			name: "deployment with large pool funding",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
				input := basicDeployLockReleaseTokenAndPoolInput(chainReport)
				input.PoolFundingAmount = big.NewInt(500_000)
				return input
			},
		},
		{
			name: "deployment with allowlist",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
				input := basicDeployLockReleaseTokenAndPoolInput(chainReport)
				input.DeployTokenPoolInput.ConstructorArgs.Allowlist = []common.Address{
					common.HexToAddress("0x03"),
					common.HexToAddress("0x04"),
				}
				return input
			},
		},
		{
			name: "deployment with rate limit admin",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
				input := basicDeployLockReleaseTokenAndPoolInput(chainReport)
				input.DeployTokenPoolInput.RateLimitAdmin = common.HexToAddress("0x05")
				return input
			},
		},
		{
			name: "deployment with custom token parameters",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) tokens.DeployLockReleaseTokenAndPoolInput {
				input := basicDeployLockReleaseTokenAndPoolInput(chainReport)
				input.TokenInfo = tokens.TokenInfo{
					Decimals:  6,
					MaxSupply: big.NewInt(10_000_000),
					Name:      "Custom Token",
				}
				input.DeployTokenPoolInput.TokenSymbol = "CUSTOM"
				input.DeployTokenPoolInput.ConstructorArgs.LocalTokenDecimals = 6
				return input
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chainSel := uint64(5009297550715157269)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSel}),
			)
			require.NoError(t, err)

			chainReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSel],
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSel,
					ContractParams: testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err)

			input := tt.makeInput(chainReport)

			report, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployLockReleaseTokenAndPool,
				e.BlockChains.EVMChains()[chainSel],
				input,
			)
			require.NoError(t, err)
			require.Len(t, report.Output.Addresses, 2)

			tokenAddress := report.Output.Addresses[0].Address
			poolAddress := report.Output.Addresses[1].Address

			token, err := token_bindings.NewBurnMintERC677(common.HexToAddress(tokenAddress), e.BlockChains.EVMChains()[chainSel].Client)
			require.NoError(t, err)

			name, err := token.Name(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err)
			require.Equal(t, input.TokenInfo.Name, name)

			symbol, err := token.Symbol(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err)
			require.Equal(t, input.DeployTokenPoolInput.TokenSymbol, symbol)

			decimals, err := token.Decimals(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err)
			require.Equal(t, input.TokenInfo.Decimals, decimals)

			maxSupply, err := token.MaxSupply(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err)
			require.Equal(t, input.TokenInfo.MaxSupply, maxSupply)

			for account, expectedAmount := range input.Accounts {
				balance, err := token.BalanceOf(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, account)
				require.NoError(t, err)
				require.Equal(t, expectedAmount, balance)
			}

			if input.PoolFundingAmount != nil && input.PoolFundingAmount.Cmp(big.NewInt(0)) > 0 {
				poolBalance, err := token.BalanceOf(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, common.HexToAddress(poolAddress))
				require.NoError(t, err)
				require.Equal(t, input.PoolFundingAmount, poolBalance)
			}

			totalSupply, err := token.TotalSupply(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err)
			expectedTotal := big.NewInt(0)
			for _, amount := range input.Accounts {
				expectedTotal = new(big.Int).Add(expectedTotal, amount)
			}
			if input.PoolFundingAmount != nil && input.PoolFundingAmount.Cmp(big.NewInt(0)) > 0 {
				expectedTotal = new(big.Int).Add(expectedTotal, input.PoolFundingAmount)
			}
			require.Equal(t, expectedTotal, totalSupply)
		})
	}
}
