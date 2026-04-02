package tokens_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var thresholdAmountForAdditionalCCVs = big.NewInt(1e18)

type deployedTokenAndPool struct {
	TokenAddress         common.Address
	TokenPoolAddress     common.Address
	AdvancedHooksAddress common.Address
	Addresses            []datastore.AddressRef
}

// deployTokenAndPoolForTest deploys a token and its associated pool using individual
// operations/sequences, replicating the state that the composite DeployTokenAndPool
// sequence used to produce: token deployed, pool deployed, mint/burn roles granted
// to the pool, initial supply minted, deployer roles revoked.
func deployTokenAndPoolForTest(
	t *testing.T,
	bundle operations.Bundle,
	chain evm.Chain,
	chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput],
	isLockRelease bool,
) deployedTokenAndPool {
	t.Helper()

	chainSel := chainReport.Input.ChainSelector

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

	// 1. Deploy token
	deployTokenReport, err := operations.ExecuteOperation(bundle, burn_mint_erc20_with_drip.Deploy, chain, contract_utils.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *burn_mint_erc20_with_drip.Version),
		Args: burn_mint_erc20_with_drip.ConstructorArgs{
			Name:   "TEST",
			Symbol: "TEST",
		},
	})
	require.NoError(t, err, "Failed to deploy token")
	tokenAddress := common.HexToAddress(deployTokenReport.Output.Address)
	tokenRef := deployTokenReport.Output
	tokenRef.Qualifier = "TEST"

	// 2. Deploy pool
	poolInput := tokens.DeployTokenPoolInput{
		ChainSel:                         chainSel,
		TokenSymbol:                      "TEST",
		RateLimitAdmin:                   common.HexToAddress("0x01"),
		ThresholdAmountForAdditionalCCVs: thresholdAmountForAdditionalCCVs,
		FeeAggregator:                    common.HexToAddress("0x03"),
		ConstructorArgs: tokens.ConstructorArgs{
			Token:    tokenAddress,
			Decimals: 18,
			RMNProxy: rmnProxyAddress,
			Router:   routerAddress,
		},
		AdvancedPoolHooksConfig: tokens.AdvancedPoolHooksConfig{
			Allowlist:         nil,
			PolicyEngine:      common.Address{},
			AuthorizedCallers: nil,
		},
	}
	if isLockRelease {
		poolInput.TokenPoolType = datastore.ContractType(lock_release_token_pool.ContractType)
		poolInput.TokenPoolVersion = lock_release_token_pool.Version
	} else {
		poolInput.TokenPoolType = datastore.ContractType(burn_mint_token_pool.ContractType)
		poolInput.TokenPoolVersion = burn_mint_token_pool.Version
	}

	var poolReport operations.SequenceReport[tokens.DeployTokenPoolInput, seq_core.OnChainOutput]
	if isLockRelease {
		poolReport, err = operations.ExecuteSequence(bundle, tokens.DeployLockReleaseTokenPool, chain, poolInput)
	} else {
		poolReport, err = operations.ExecuteSequence(bundle, tokens.DeployBurnMintTokenPool, chain, poolInput)
	}
	require.NoError(t, err, "Failed to deploy token pool")

	// Find the token pool address
	var tokenPoolAddress common.Address
	for _, addr := range poolReport.Output.Addresses {
		if strings.Contains(string(addr.Type), "TokenPool") {
			tokenPoolAddress = common.HexToAddress(addr.Address)
			break
		}
	}
	require.NotEqual(t, common.Address{}, tokenPoolAddress, "Token pool address should be found")

	// 3. Grant mint/burn roles to the pool
	_, err = operations.ExecuteOperation(bundle, burn_mint_erc20_with_drip.GrantMintAndBurnRoles, chain, contract_utils.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       tokenAddress,
		Args:          tokenPoolAddress,
	})
	require.NoError(t, err, "Failed to grant mint and burn roles to token pool")

	// 4. Grant mint/burn to deployer, mint initial supply, then revoke
	_, err = operations.ExecuteOperation(bundle, burn_mint_erc20_with_drip.GrantMintAndBurnRoles, chain, contract_utils.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       tokenAddress,
		Args:          chain.DeployerKey.From,
	})
	require.NoError(t, err, "Failed to grant mint and burn roles to deployer")

	accounts := map[common.Address]*big.Int{
		common.HexToAddress("0x01"): big.NewInt(500_000),
		common.HexToAddress("0x02"): big.NewInt(500_000),
	}
	for account, amount := range accounts {
		_, err = operations.ExecuteOperation(bundle, burn_mint_erc20_with_drip.Mint, chain, contract_utils.FunctionInput[burn_mint_erc20_with_drip.MintArgs]{
			ChainSelector: chainSel,
			Address:       tokenAddress,
			Args: burn_mint_erc20_with_drip.MintArgs{
				Account: account,
				Amount:  amount,
			},
		})
		require.NoError(t, err, "Failed to mint tokens")
	}

	_, err = operations.ExecuteOperation(bundle, burn_mint_erc20_with_drip.RevokeMintRole, chain, contract_utils.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       tokenAddress,
		Args:          chain.DeployerKey.From,
	})
	require.NoError(t, err, "Failed to revoke mint role from deployer")

	_, err = operations.ExecuteOperation(bundle, burn_mint_erc20_with_drip.RevokeBurnRole, chain, contract_utils.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       tokenAddress,
		Args:          chain.DeployerKey.From,
	})
	require.NoError(t, err, "Failed to revoke burn role from deployer")

	// Build return: addresses in same order as old DeployTokenAndPool (token=0, pool=1, hooks=2, ...)
	allAddresses := []datastore.AddressRef{tokenRef}
	allAddresses = append(allAddresses, poolReport.Output.Addresses...)

	var advancedHooksAddress common.Address
	if len(poolReport.Output.Addresses) >= 2 {
		advancedHooksAddress = common.HexToAddress(poolReport.Output.Addresses[1].Address)
	}

	return deployedTokenAndPool{
		TokenAddress:         tokenAddress,
		TokenPoolAddress:     tokenPoolAddress,
		AdvancedHooksAddress: advancedHooksAddress,
		Addresses:            allAddresses,
	}
}
