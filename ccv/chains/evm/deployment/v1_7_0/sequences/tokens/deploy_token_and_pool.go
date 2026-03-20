package tokens

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/burn_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/burn_with_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	token_pool_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
)

// DeployTokenAndPoolInput is the input for the DeployBurnMintTokenAndPool sequence.
type DeployTokenAndPoolInput struct {
	// Accounts is a map of account addresses to initial mint amounts.
	Accounts map[common.Address]*big.Int
	// DeployTokenPoolInput is the input for the DeployTokenPool sequence.
	DeployTokenPoolInput DeployTokenPoolInput
	// RegistryAddress is the TokenAdminRegistry address; when set and RateLimitAdmin is zero, RateLimitAdmin is imported from the active pool if it is < 2.0.0.
	RegistryAddress common.Address
}

func (c DeployTokenAndPoolInput) ChainSelector() uint64 {
	return c.DeployTokenPoolInput.ChainSel
}

var DeployTokenAndPool = cldf_ops.NewSequence(
	"deploy-token-and-pool",
	semver.MustParse("2.0.0"),
	"Deploys a token and its associated token pool to an EVM chain, granting rights to the token pool and minting initial supply",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenAndPoolInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]evm_contract.WriteOutput, 0)

		// Deploy burn mint token.
		deployTokenReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.Deploy, chain, evm_contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
			ChainSelector:  input.DeployTokenPoolInput.ChainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *burn_mint_erc20_with_drip.Version),
			Args: burn_mint_erc20_with_drip.ConstructorArgs{
				Name:   input.DeployTokenPoolInput.TokenSymbol,
				Symbol: input.DeployTokenPoolInput.TokenSymbol,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token to %s: %w", chain, err)
		}
		deployTokenReport.Output.Qualifier = input.DeployTokenPoolInput.TokenSymbol // Use the token symbol as the qualifier.
		addresses = append(addresses, deployTokenReport.Output)

		tokenAddress := common.HexToAddress(deployTokenReport.Output.Address)
		// When RegistryAddress is set and RateLimitAdmin is not provided, import from the active pool if it is < 2.0.0.
		if input.RegistryAddress != (common.Address{}) && input.DeployTokenPoolInput.RateLimitAdmin == (common.Address{}) {
			rla, err := getRateLimitAdminFromActivePool(b, chain, input.ChainSelector(), input.RegistryAddress, tokenAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to import rate limit admin from active pool: %w", err)
			}
			if rla != (common.Address{}) {
				input.DeployTokenPoolInput.RateLimitAdmin = rla
			}
		}

		// Deploy token pool.
		input.DeployTokenPoolInput.ConstructorArgs.Token = tokenAddress
		poolType := deployment.ContractType(input.DeployTokenPoolInput.TokenPoolType)
		poolVersion := input.DeployTokenPoolInput.TokenPoolVersion
		isBurnMint := (poolType == burn_mint_token_pool.ContractType && poolVersion != nil && poolVersion.Equal(burn_mint_token_pool.Version)) ||
			(poolType == burn_from_mint_token_pool.ContractType && poolVersion != nil && poolVersion.Equal(burn_from_mint_token_pool.Version)) ||
			(poolType == burn_with_from_mint_token_pool.ContractType && poolVersion != nil && poolVersion.Equal(burn_with_from_mint_token_pool.Version))
		isLockRelease := (poolType == lock_release_token_pool.ContractType && poolVersion != nil && poolVersion.Equal(lock_release_token_pool.Version)) ||
			(poolType == siloed_lock_release_token_pool.ContractType && poolVersion != nil && poolVersion.Equal(siloed_lock_release_token_pool.Version))
		switch {
		case isBurnMint:
			deployTokenPoolReport, err := cldf_ops.ExecuteSequence(b, DeployBurnMintTokenPool, chain, input.DeployTokenPoolInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token pool to %s: %w", chain, err)
			}
			addresses = append(addresses, deployTokenPoolReport.Output.Addresses...)
		case isLockRelease:
			deployTokenPoolReport, err := cldf_ops.ExecuteSequence(b, DeployLockReleaseTokenPool, chain, input.DeployTokenPoolInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy lock release token pool to %s: %w", chain, err)
			}
			addresses = append(addresses, deployTokenPoolReport.Output.Addresses...)
		default:
			return sequences.OnChainOutput{}, fmt.Errorf("token pool type %s and version %s is not supported", input.DeployTokenPoolInput.TokenPoolType, input.DeployTokenPoolInput.TokenPoolVersion)
		}

		var tokenPoolAddress common.Address
		for _, address := range addresses {
			if strings.Contains(string(address.Type), "TokenPool") {
				tokenPoolAddress = common.HexToAddress(address.Address)
				break
			}
		}
		if tokenPoolAddress == (common.Address{}) {
			return sequences.OnChainOutput{}, fmt.Errorf("token pool address not found")
		}

		// Grant mint and burn roles to the token pool.
		grantRolesReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.GrantMintAndBurnRoles, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          tokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint and burn roles to token pool on %s: %w", chain, err)
		}
		writes = append(writes, grantRolesReport.Output)

		// Grant roles to the deployer key so we can mint initial supply.
		grantMintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.GrantMintAndBurnRoles, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint role to deployer on %s: %w", chain, err)
		}
		writes = append(writes, grantMintReport.Output)

		// Mint initial supply to each account.
		for account, amount := range input.Accounts {
			mintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.Mint, chain, evm_contract.FunctionInput[burn_mint_erc20_with_drip.MintArgs]{
				ChainSelector: input.DeployTokenPoolInput.ChainSel,
				Address:       common.HexToAddress(deployTokenReport.Output.Address),
				Args: burn_mint_erc20_with_drip.MintArgs{
					Account: account,
					Amount:  amount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to mint %s tokens to %s on %s: %w", amount.String(), account.Hex(), chain, err)
			}
			writes = append(writes, mintReport.Output)
		}

		// Revoke mint role from the deployer key for safety.
		revokeMintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.RevokeMintRole, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke mint role from deployer on %s: %w", chain, err)
		}
		writes = append(writes, revokeMintReport.Output)

		// Revoke burn role from the deployer key for safety.
		revokeBurnReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20_with_drip.RevokeBurnRole, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke burn role from deployer on %s: %w", chain, err)
		}
		writes = append(writes, revokeBurnReport.Output)

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)

// getRateLimitAdminFromActivePool returns the rate limit admin from the active pool in the TokenAdminRegistry when that pool is < 2.0.0. Returns zero address otherwise.
func getRateLimitAdminFromActivePool(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	registryAddress, tokenAddress common.Address,
) (common.Address, error) {
	if registryAddress == (common.Address{}) || tokenAddress == (common.Address{}) {
		return common.Address{}, nil
	}
	tokenConfigReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.GetTokenConfig, chain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       registryAddress,
		Args:          tokenAddress,
	})
	if err != nil {
		return common.Address{}, err
	}
	activePool := tokenConfigReport.Output.TokenPool
	if activePool == (common.Address{}) {
		return common.Address{}, nil
	}
	typeAndVersionReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       activePool,
		Args:          struct{}{},
	})
	if err != nil {
		return common.Address{}, err
	}
	if typeAndVersionReport.Output.Version.GreaterThanEqual(semver.MustParse("2.0.0")) {
		// Configuration import from another 2.0.0 pool is not currently supported
		return common.Address{}, nil
	}
	rlaReport, err := cldf_ops.ExecuteOperation(b, token_pool_v161.GetRateLimitAdmin, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       activePool,
		Args:          struct{}{},
	})
	if err != nil {
		return common.Address{}, err
	}
	return rlaReport.Output, nil
}
