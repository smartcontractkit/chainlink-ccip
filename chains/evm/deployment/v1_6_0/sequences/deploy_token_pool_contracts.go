package sequences

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	adaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
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
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var DeployTokenPool = cldf_ops.NewSequence(
	"deploy-token-pool",
	utils.Version_1_6_1,
	"Deploy given type of token pool contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokenapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
		}

		// Validate required deployment inputs
		poolutil := adaptersV1_0_0.EVMTokenBase{}
		if input.TokenPoolVersion == nil {
			return sequences.OnChainOutput{}, errors.New("TokenPoolVersion is required")
		}
		if input.TokenRef == nil {
			return sequences.OnChainOutput{}, errors.New("TokenRef is required")
		}

		// Parse the token ref as an EVM address
		tokenAddress, err := poolutil.ParseNonZeroAddressRef(input.ExistingDataStore, input.TokenRef.Clone(), chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve token address from ref: %w", err)
		}

		// If no pool qualifier is provided, then fall back to using the token address
		poolQualifier := input.TokenPoolQualifier
		if poolQualifier == "" {
			poolQualifier = tokenAddress.Hex()
		}

		// NOTE: the datastore uses the type, selector, qualifier, and version of an address
		// ref to uniquely identify records, so the query below should only match one record
		// at most. If multiple records are returned, then this would indicate an issue with
		// the datastore's data integrity. If no matches are returned, then the ref does not
		// exist and we proceed with the deployment.
		matches := input.ExistingDataStore.Addresses().Filter(
			datastore.AddressRefByType(datastore.ContractType(input.PoolType)),
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByQualifier(poolQualifier),
			datastore.AddressRefByVersion(input.TokenPoolVersion),
		)
		if len(matches) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"multiple token pools found in datastore with type '%s', version '%s', qualifier '%s' on chain with selector %d",
				input.PoolType, input.TokenPoolVersion.String(), poolQualifier, chain.Selector,
			)
		}
		if len(matches) == 1 {
			b.Logger.Infof("Token pool already deployed: %s", datastore_utils.SprintRef(matches[0]))
			return sequences.OnChainOutput{Addresses: matches}, nil
		}

		// Infer pool deployment inputs
		tokenDecimals, err := poolutil.TokenInfo(b, input.ExistingDataStore, chain, tokenAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token decimals for token at address '%s': %w", tokenAddress, err)
		}
		rmnProxyAddr, err := poolutil.GetRMNProxyAddress(input.ExistingDataStore, chain.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve rmn proxy address for chain selector %d: %w", chain.Selector, err)
		}
		routerAddr, err := poolutil.ResolveRouterAddress(input.ExistingDataStore, chain.Selector, input.RouterRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve router address for chain selector %d: %w", chain.Selector, err)
		}
		allowlist, err := poolutil.ParseAddressStrings(input.Allowlist)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to parse allowlist: %w", err)
		}

		// Build type and version struct
		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.PoolType),
			*input.TokenPoolVersion,
		)

		// Deploy the desired pool contract
		var poolRef datastore.AddressRef
		switch typeAndVersion.String() {
		case v1_6_1_burn_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_mint_token_pool.ConstructorArgs{
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_from_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_from_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_from_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_from_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_from_mint_token_pool.ConstructorArgs{
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnFromMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_mint_with_lock_release_flag_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_mint_with_lock_release_flag_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_mint_with_lock_release_flag_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_mint_with_lock_release_flag_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_mint_with_lock_release_flag_token_pool.ConstructorArgs{
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintWithLockReleaseFlagTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_to_address_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_to_address_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_to_address_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_to_address_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_to_address_mint_token_pool.ConstructorArgs{
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
					BurnAddress:        common.HexToAddress(input.BurnAddress),
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnToAddressMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_burn_with_from_mint_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_burn_with_from_mint_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_burn_with_from_mint_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_burn_with_from_mint_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_burn_with_from_mint_token_pool.ConstructorArgs{
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnWithFromMintTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_lock_release_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_lock_release_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_lock_release_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_lock_release_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_lock_release_token_pool.ConstructorArgs{
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LockReleaseTokenPool v1.6.1: %w", err)
			}

		case v1_6_1_siloed_lock_release_token_pool.TypeAndVersion.String():
			poolRef, err = contract.MaybeDeployContract(b, v1_6_1_siloed_lock_release_token_pool.Deploy, chain, contract.DeployInput[v1_6_1_siloed_lock_release_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_1_siloed_lock_release_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_1_siloed_lock_release_token_pool.ConstructorArgs{
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy SiloedLockReleaseTokenPool v1.6.1: %w", err)
			}

		case v1_6_0_burn_mint_with_external_minter_token_pool.TypeAndVersion.String():
			tokenGovernor, err := fetchTokenGovernor(input)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to fetch token governor address: %w", err)
			}
			poolRef, err = contract.MaybeDeployContract(b, v1_6_0_burn_mint_with_external_minter_token_pool.Deploy, chain, contract.DeployInput[v1_6_0_burn_mint_with_external_minter_token_pool.ConstructorArgs]{
				TypeAndVersion: v1_6_0_burn_mint_with_external_minter_token_pool.TypeAndVersion,
				ChainSelector:  chain.Selector,
				Args: v1_6_0_burn_mint_with_external_minter_token_pool.ConstructorArgs{
					Minter:             tokenGovernor,
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
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
					Minter:             tokenGovernor,
					LocalTokenDecimals: tokenDecimals,
					Token:              tokenAddress,
					Allowlist:          allowlist,
					RmnProxy:           rmnProxyAddr,
					Router:             routerAddr,
				},
				Qualifier: &poolQualifier,
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy HybridWithExternalMinterTokenPool v1.6.0: %w", err)
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token pool type and version: %s", typeAndVersion)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{poolRef},
			BatchOps:  []mcms_types.BatchOperation{},
		}, nil
	},
)

func fetchTokenGovernor(input tokenapi.DeployTokenPoolInput) (common.Address, error) {
	// If the token governor address is provided directly, then
	// skip the daastore lookup and use the provided address.
	if input.TokenGovernor != "" {
		if !common.IsHexAddress(input.TokenGovernor) {
			return common.Address{}, fmt.Errorf("provided token governor address '%s' is not a valid hex address", input.TokenGovernor)
		} else {
			return common.HexToAddress(input.TokenGovernor), nil
		}
	}

	// If the token governor address isn't provided, then try
	// to find it in the datastore.
	tokenGovernorAddr, err := datastore_utils.FindAndFormatRef(
		input.ExistingDataStore,
		datastore.AddressRef{
			ChainSelector: input.ChainSelector,
			Type:          datastore.ContractType(token_governor.ContractType),
			Qualifier:     input.TokenRef.Qualifier,
		},
		input.ChainSelector,
		datastore_utils_evm.ToEVMAddress,
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token governor address in datastore: %w", err)
	}

	return tokenGovernorAddr, nil
}
