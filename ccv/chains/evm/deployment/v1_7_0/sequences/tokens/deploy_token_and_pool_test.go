package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	token_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
)

var thresholdAmountForAdditionalCCVs = big.NewInt(1e18)

func basicDeployTokenAndPoolInput(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput], isLockRelease bool) tokens.DeployTokenAndPoolInput {
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

	var tokenPoolType datastore.ContractType
	var tokenPoolVersion *semver.Version
	if isLockRelease {
		tokenPoolType = datastore.ContractType(lock_release_token_pool.ContractType)
		tokenPoolVersion = lock_release_token_pool.Version
	} else {
		tokenPoolType = datastore.ContractType(burn_mint_token_pool.BurnMintContractType)
		tokenPoolVersion = burn_mint_token_pool.Version
	}

	return tokens.DeployTokenAndPoolInput{
		Accounts: map[common.Address]*big.Int{
			common.HexToAddress("0x01"): big.NewInt(500_000),
			common.HexToAddress("0x02"): big.NewInt(500_000),
		},
		DeployTokenPoolInput: tokens.DeployTokenPoolInput{
			ChainSel:                         chainReport.Input.ChainSelector,
			TokenPoolType:                    tokenPoolType,
			TokenPoolVersion:                 tokenPoolVersion,
			TokenSymbol:                      "TEST",
			RateLimitAdmin:                   common.HexToAddress("0x01"),
			ThresholdAmountForAdditionalCCVs: thresholdAmountForAdditionalCCVs,
			FeeAggregator:                    common.HexToAddress("0x03"),
			ConstructorArgs: tokens.ConstructorArgs{
				Decimals: 18,
				RMNProxy: rmnProxyAddress,
				Router:   routerAddress,
			},
			AdvancedPoolHooksConfig: tokens.AdvancedPoolHooksConfig{
				Allowlist:         nil,              // disabled
				PolicyEngine:      common.Address{}, // disabled
				AuthorizedCallers: nil,              // no initial callers; deploy sequence will authorize the deployed pool
			},
		},
	}
}

func TestDeployTokenAndPool(t *testing.T) {
	tests := []struct {
		desc          string
		isLockRelease bool
		expectedErr   string
	}{
		{
			desc:          "happy path - burn mint token pool",
			isLockRelease: false,
			expectedErr:   "",
		},
		{
			desc:          "happy path - lock release token pool",
			isLockRelease: true,
			expectedErr:   "",
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
			input := basicDeployTokenAndPoolInput(chainReport, test.isLockRelease)
			poolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployTokenAndPool,
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

			tokenAddress := poolReport.Output.Addresses[0].Address
			poolAddress := poolReport.Output.Addresses[1].Address

			if test.isLockRelease {
				// Lock release token pool returns 4 addresses: token, pool, hooks, lockBox
				require.Len(t, poolReport.Output.Addresses, 4, "Expected 4 addresses in output (token, pool, hooks, lockBox)")
			} else {
				// Burn mint token pool returns 3 addresses: token, pool, hooks
				require.Len(t, poolReport.Output.Addresses, 3, "Expected 3 addresses in output (token, pool, hooks)")
			}

			// Check token metadata
			token, err := token_bindings.NewBurnMintERC20WithDrip(common.HexToAddress(tokenAddress), e.BlockChains.EVMChains()[chainSel].Client)
			require.NoError(t, err, "NewBurnMintERC677 should not error")
			name, err := token.Name(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "Name should not error")
			require.Equal(t, input.DeployTokenPoolInput.TokenSymbol, name, "Expected token name to be the same as the deployed token")
			symbol, err := token.Symbol(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "Symbol should not error")
			require.Equal(t, input.DeployTokenPoolInput.TokenSymbol, symbol, "Expected token symbol to be the same as the deployed token")
			decimals, err := token.Decimals(&bind.CallOpts{Context: e.OperationsBundle.GetContext()})
			require.NoError(t, err, "Decimals should not error")
			require.Equal(t, uint8(18), decimals, "Expected token decimals to be 18")

			// For burn mint token pools, check mint and burn roles
			// Check token minters
			hasMintRole, err := token.HasRole(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, burn_mint_erc20_with_drip.MintRole, common.HexToAddress(poolAddress))
			require.NoError(t, err, "HasRole should not error")
			require.True(t, hasMintRole, "Expected token pool to have the mint role")

			// Check token burners
			hasBurnRole, err := token.HasRole(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, burn_mint_erc20_with_drip.BurnRole, common.HexToAddress(poolAddress))
			require.NoError(t, err, "HasRole should not error")
			require.True(t, hasBurnRole, "Expected token pool to have the burn role")

			// Ensure that the deployer key is not a minter or burner
			hasMintRole, err = token.HasRole(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, burn_mint_erc20_with_drip.MintRole, e.BlockChains.EVMChains()[chainSel].DeployerKey.From)
			require.NoError(t, err, "HasRole should not error")
			require.False(t, hasMintRole, "Expected deployer key to not have the mint role")
			hasBurnRole, err = token.HasRole(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, burn_mint_erc20_with_drip.BurnRole, e.BlockChains.EVMChains()[chainSel].DeployerKey.From)
			require.NoError(t, err, "HasRole should not error")
			require.False(t, hasBurnRole, "Expected deployer key to not have the burn role")

			// Check balance of each account
			for addr, amount := range input.Accounts {
				balance, err := token.BalanceOf(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, addr)
				require.NoError(t, err, "BalanceOf should not error")
				require.Equal(t, amount, balance, "Expected balance of account %s to be the same as the deployed token", addr)
			}
		})
	}
}
