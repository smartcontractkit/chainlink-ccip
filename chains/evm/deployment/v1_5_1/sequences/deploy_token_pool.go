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
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var DeployTokenPool = cldf_ops.NewSequence(
	"deploy-token-pool-v1.5.1",
	common_utils.Version_1_5_1,
	"Deploy v1.5.1 token pool contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
		if input.TokenPoolVersion == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("TokenPoolVersion is required")
		}

		chain := chains.EVMChains()[input.ChainSelector]
		qualifier := input.TokenPoolQualifier

		tokenAddr, err := resolveTokenAddress(input)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if qualifier == "" {
			qualifier = tokenAddr
		}

		matches := input.ExistingDataStore.Addresses().Filter(
			datastore.AddressRefByType(datastore.ContractType(input.PoolType)),
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByQualifier(qualifier),
			datastore.AddressRefByVersion(input.TokenPoolVersion),
		)
		if len(matches) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"multiple token pools found in datastore with type '%s', version '%s', qualifier '%s' on chain with selector %d",
				input.PoolType, input.TokenPoolVersion.String(), qualifier, input.ChainSelector,
			)
		}
		if len(matches) == 1 {
			b.Logger.Info("Token pool already deployed at address:", matches[0].Address)
			return sequences.OnChainOutput{}, nil
		}

		token, err := burn_mint_erc20.NewBurnMintERC20(common.HexToAddress(tokenAddr), chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token contract at address '%s': %w", tokenAddr, err)
		}
		tokenDecimal, err := token.Decimals(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token decimals for token at address '%s': %w", tokenAddr, err)
		}

		routerAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(router.ContractType),
		}, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find router address in datastore for chain with selector %d: %w", input.ChainSelector, err)
		}

		rmpProxyAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(rmnproxyops.ContractType),
		}, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find rmnproxy address in datastore for chain with selector %d: %w", input.ChainSelector, err)
		}

		var allowlist []common.Address
		if len(input.Allowlist) > 0 {
			allowlist = make([]common.Address, 0, len(input.Allowlist))
			for _, addr := range input.Allowlist {
				allowlist = append(allowlist, common.HexToAddress(addr))
			}
		}

		var poolRef datastore.AddressRef
		typeAndVersion := deployment.NewTypeAndVersion(deployment.ContractType(input.PoolType), *input.TokenPoolVersion).String()

		switch typeAndVersion {
		case v1_5_1_burn_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_5_1_burn_mint_token_pool.Deploy, chain, contract.DeployInput[v1_5_1_burn_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_5_1_burn_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_5_1_burn_mint_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(tokenAddr),
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
					Token:              common.HexToAddress(tokenAddr),
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
					Token:              common.HexToAddress(tokenAddr),
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
					Token:              common.HexToAddress(tokenAddr),
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
					Token:              common.HexToAddress(tokenAddr),
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
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported v1.5.1 token pool type and version: %s", typeAndVersion)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{poolRef},
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)

func resolveTokenAddress(input tokenapi.DeployTokenPoolInput) (string, error) {
	var tokenAddr string
	if input.TokenRef != nil && input.TokenRef.Address != "" {
		tokenAddr = input.TokenRef.Address
	}
	if input.TokenRef != nil && input.TokenRef.Qualifier != "" {
		storedAddr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, *input.TokenRef, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return "", fmt.Errorf("token with symbol '%s' is not found in datastore, %v", input.TokenRef.Qualifier, err)
		}
		if tokenAddr != "" && storedAddr.Address != tokenAddr {
			return "", fmt.Errorf("provided token address '%s' does not match address '%s' found in datastore for symbol '%s'", tokenAddr, storedAddr.Address, input.TokenRef.Qualifier)
		}
		if tokenAddr == "" {
			tokenAddr = storedAddr.Address
		}
	}
	if tokenAddr == "" {
		return "", fmt.Errorf("token address must be provided either directly or via a datastore reference")
	}
	return tokenAddr, nil
}
