package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	v1_5_1_burn_from_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_from_mint_token_pool"
	v1_5_1_burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_mint_token_pool"
	v1_5_1_burn_to_address_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_to_address_mint_token_pool"
	v1_5_1_burn_with_from_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_with_from_mint_token_pool"
	v1_5_1_lock_release_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/lock_release_token_pool"
	v1_6_0_burn_mint_with_external_minter_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/burn_mint_with_external_minter_token_pool"
	v1_6_0_hybrid_with_external_minter_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/hybrid_with_external_minter_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/token_governor"
	v1_6_1_burn_from_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_from_mint_token_pool"
	v1_6_1_burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_token_pool"
	v1_6_1_burn_mint_with_lock_release_flag_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	v1_6_1_burn_to_address_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_to_address_mint_token_pool"
	v1_6_1_burn_with_from_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_with_from_mint_token_pool"
	v1_6_1_lock_release_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	v1_6_1_siloed_lock_release_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var DeployTokenPool = cldf_ops.NewSequence(
	"deploy-token-pool",
	common_utils.Version_1_6_1,
	"Deploy given type of token pool contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
		if input.TokenPoolVersion == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("TokenPoolVersion is required")
		}

		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		chain := chains.EVMChains()[input.ChainSelector]

		tokenPoolAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(input.PoolType),
			Qualifier:     input.TokenPoolQualifier,
		}, input.ChainSelector, datastore_utils.FullRef)
		if err == nil {
			b.Logger.Info("Token pool already deployed at address:", tokenPoolAddr.Address)
			return sequences.OnChainOutput{}, nil
		}

		// find token address from the data store
		tokenAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Qualifier:     input.TokenSymbol,
		}, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("token with symbol '%s' is not found in datastore, %v", input.TokenSymbol, err)
		}

		// get token decimals
		token, err := burn_mint_erc20.NewBurnMintERC20(common.HexToAddress(tokenAddr.Address), chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token contract at address '%s': %w", tokenAddr.Address, err)
		}
		tokenDecimal, err := token.Decimals(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token decimals for token at address '%s': %w", tokenAddr.Address, err)
		}

		// find the router address from the data store
		routerAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(router.ContractType),
		}, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("token with symbol '%s' is not found in datastore, %v", input.TokenSymbol, err)
		}

		// find the rmnproxy address from the data store
		rmpProxyAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(rmnproxyops.ContractType),
		}, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("token with symbol '%s' is not found in datastore, %v", input.TokenSymbol, err)
		}

		// prepare allowlist
		var allowlist []common.Address
		if len(input.Allowlist) > 0 {
			allowlist = make([]common.Address, 0, len(input.Allowlist))
			for _, addr := range input.Allowlist {
				allowlist = append(allowlist, common.HexToAddress(addr))
			}
		}

		var poolRef datastore.AddressRef

		typeAndVersion := deployment.NewTypeAndVersion(deployment.ContractType(input.PoolType), *input.TokenPoolVersion).String()
		qualifier := input.TokenPoolQualifier

		switch typeAndVersion {
		// v1.6.1 pools
		case v1_6_1_burn_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RMNProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_from_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_from_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_from_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_from_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_from_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnFromMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_mint_with_lock_release_flag_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_mint_with_lock_release_flag_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_mint_with_lock_release_flag_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_mint_with_lock_release_flag_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_mint_with_lock_release_flag_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintWithLockReleaseFlagTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_to_address_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_to_address_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_to_address_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_to_address_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_to_address_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
					BurnAddress:        common.HexToAddress(input.BurnAddress),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnToAddressMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_with_from_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_with_from_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_with_from_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_with_from_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_with_from_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnWithFromMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_lock_release_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_lock_release_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_lock_release_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_lock_release_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_lock_release_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LockReleaseTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_siloed_lock_release_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_siloed_lock_release_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_siloed_lock_release_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_siloed_lock_release_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_siloed_lock_release_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy SiloedLockReleaseTokenPool v1.6.1: %w", err)
			}

		// 1.6.0 pools
		case v1_6_0_burn_mint_with_external_minter_token_pool.TypeAndVersion.String():
			tokenGovernor, err := fetchTokenGovernor(input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to fetch token governor address: %w", err)
			}
			poolRef, err = contract.MaybeDeployContract(b, v1_6_0_burn_mint_with_external_minter_token_pool.Deploy, chain, contract.DeployInput[v1_6_0_burn_mint_with_external_minter_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_0_burn_mint_with_external_minter_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_0_burn_mint_with_external_minter_token_pool.ConstructorArgs{
					Minter:             common.HexToAddress(tokenGovernor),
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintWithExternalMinterTokenPool v1.6.0: %w", err)
			}

		case v1_6_0_hybrid_with_external_minter_token_pool.TypeAndVersion.String():
			tokenGovernor, err := fetchTokenGovernor(input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to fetch token governor address: %w", err)
			}
			poolRef, err = contract.MaybeDeployContract(b, v1_6_0_hybrid_with_external_minter_token_pool.Deploy, chain, contract.DeployInput[v1_6_0_hybrid_with_external_minter_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_0_hybrid_with_external_minter_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_0_hybrid_with_external_minter_token_pool.ConstructorArgs{
					Minter:             common.HexToAddress(tokenGovernor),
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy HybridWithExternalMinterTokenPool v1.6.0: %w", err)
			}

		// v1.5.1 pools
		case v1_5_1_burn_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_5_1_burn_mint_token_pool.Deploy, chain, contract.DeployInput[v1_5_1_burn_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_5_1_burn_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_5_1_burn_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintTokenPool v1.5.1: %w", err)
			}

		case v1_5_1_burn_from_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_5_1_burn_from_mint_token_pool.Deploy, chain, contract.DeployInput[v1_5_1_burn_from_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_5_1_burn_from_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_5_1_burn_from_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnFromMintTokenPool v1.5.1: %w", err)
			}

		case v1_5_1_burn_to_address_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_5_1_burn_to_address_mint_token_pool.Deploy, chain, contract.DeployInput[v1_5_1_burn_to_address_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_5_1_burn_to_address_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_5_1_burn_to_address_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
					BurnAddress:        common.HexToAddress(input.BurnAddress),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnToAddressMintTokenPool v1.5.1: %w", err)
			}

		case v1_5_1_burn_with_from_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_5_1_burn_with_from_mint_token_pool.Deploy, chain, contract.DeployInput[v1_5_1_burn_with_from_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_5_1_burn_with_from_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_5_1_burn_with_from_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnWithFromMintTokenPool v1.5.1: %w", err)
			}

		case v1_5_1_lock_release_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_5_1_lock_release_token_pool.Deploy, chain, contract.DeployInput[v1_5_1_lock_release_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_5_1_lock_release_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_5_1_lock_release_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr.Address),
					LocalTokenDecimals: tokenDecimal,
					Allowlist:          allowlist,
					RmnProxy:           common.HexToAddress(rmpProxyAddr.Address),
					AcceptLiquidity:    *input.AcceptLiquidity,
					Router:             common.HexToAddress(routerAddr.Address),
				},
				Qualifier: &qualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LockReleaseTokenPool v1.5.1: %w", err)
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type and version: %s", typeAndVersion)
		}

		addresses = append(addresses, poolRef)

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)

func fetchTokenGovernor(input tokenapi.DeployTokenPoolInput) (string, error) {
	tokenGovernor := input.TokenGovernor
	if tokenGovernor == "" {
		// fetch token governor from the data store
		tokenGovernorAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(token_governor.ContractType),
			Qualifier:     input.TokenSymbol,
		}, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return "", fmt.Errorf("token governor for token with symbol '%s' is not found in datastore, %v", input.TokenSymbol, err)
		}
		tokenGovernor = tokenGovernorAddr.Address
	}
	return tokenGovernor, nil
}
