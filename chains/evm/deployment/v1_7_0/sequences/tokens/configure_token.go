package tokens

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// ConfigureTokenInput is the input for the ConfigureTokens sequence.
type ConfigureTokenInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolRef describes the desired token pool.
	// If the ref includes an address, that address will be used.
	// Otherwise, a new token pool will be deployed with the desired type and version.
	TokenPoolRef datastore.AddressRef
	// TokenAddress is the address of the token to be configured.
	TokenAddress common.Address
	// LocalTokenDecimals is the number of decimals the local token uses.
	LocalTokenDecimals uint8
	// AllowList is the list of addresses allowed to transfer tokens.
	// If empty upon deployment, an allow-list can never be set.
	// Likewise, if populated upon deployment, the allow-list can never be disabled.
	AllowList []common.Address
	// RMNProxyAddress is the address of the RMNProxy contract on this chain.
	RMNProxyAddress common.Address
	// RouterAddress is the address of the Router contract on this chain.
	// If left empty, setRouter will not be attempted.
	RouterAddress common.Address
	// RateLimitAdmin is an additional address allowed to set rate limiters.
	// If left empty, setRateLimitAdmin will not be attempted.
	RateLimitAdmin common.Address
	// RemoteChains specifies the remote chains to configure on the token pool.
	RemoteChains map[uint64]RemoteChainConfig
	// AdminAddress is the address of the desired token admin.
	// If left empty, the owner of the token admin registry will be used.
	AdminAddress common.Address
	// OnlyPropose indicates whether to only propose the admin, and not accept the role or set the pool.
	OnlyPropose bool
	// TokenAdminRegistryAddress is the address of the TokenAdminRegistry contract.
	TokenAdminRegistryAddress common.Address
}

func (c ConfigureTokenInput) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	if c.TokenAddress == (common.Address{}) {
		return errors.New("token address must be defined")
	}
	requiredDeploymentInfoDefined := c.TokenPoolRef.Version != nil &&
		c.TokenPoolRef.Type != "" &&
		c.RMNProxyAddress != (common.Address{}) &&
		c.RouterAddress != (common.Address{})
	if c.TokenPoolRef.Address == "" && !requiredDeploymentInfoDefined {
		return errors.New("must define token pool address OR type, version, rmn proxy address, and router address")
	}
	if c.TokenPoolRef.Address != "" && !common.IsHexAddress(c.TokenPoolRef.Address) {
		return fmt.Errorf("tokenPoolRef address is not a valid hex address: %s", c.TokenPoolRef.Address)
	}
	return nil
}

var ConfigureToken = cldf_ops.NewSequence(
	"configure-token",
	semver.MustParse("1.7.0"),
	"Configures a token on an EVM chain for usage with CCIP",
	func(b operations.Bundle, chain evm.Chain, input ConfigureTokenInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		writes := make([]contract.WriteOutput, 0)
		refs := make([]datastore.AddressRef, 0, 1) // At most 1 new ref (the token pool)

		// Deploy a token pool (if one doesn't already exist)
		var tokenPoolRef datastore.AddressRef
		if input.TokenPoolRef.Address != "" {
			tokenPoolRef = input.TokenPoolRef
		} else {
			typeAndVersion := deployment.NewTypeAndVersion(
				deployment.ContractType(input.TokenPoolRef.Type),
				*input.TokenPoolRef.Version,
			)
			deployReport, err := cldf_ops.ExecuteOperation(b, token_pool.Deploy, chain, contract.DeployInput[token_pool.ConstructorArgs]{
				ChainSelector:  input.ChainSelector,
				TypeAndVersion: typeAndVersion,
				Args: token_pool.ConstructorArgs{
					Token:              input.TokenAddress,
					LocalTokenDecimals: input.LocalTokenDecimals,
					Allowlist:          input.AllowList,
					RMNProxy:           input.RMNProxyAddress,
					Router:             input.RouterAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
			}
			tokenPoolRef = deployReport.Output
		}
		refs = append(refs, tokenPoolRef)

		// Apply allow-list updates (if necessary)
		// First, check if the allow-list is enabled
		if len(input.AllowList) != 0 {
			allowListEnabledReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetAllowListEnabled, chain, contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(tokenPoolRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allow-list status from token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
			}
			if allowListEnabledReport.Output {
				// Allow-list is enabled, so we first check the current allow-list
				currentAllowListReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetAllowList, chain, contract.FunctionInput[any]{
					ChainSelector: input.ChainSelector,
					Address:       common.HexToAddress(tokenPoolRef.Address),
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get current allow-list from token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
				}
				adds, removes := makeAllowListUpdates(currentAllowListReport.Output, input.AllowList)

				// Apply any updates to the allow-list if they exist
				if len(adds) != 0 || len(removes) != 0 {
					applyAllowListUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyAllowListUpdates, chain, contract.FunctionInput[token_pool.ApplyAllowListUpdatesArgs]{
						ChainSelector: input.ChainSelector,
						Address:       common.HexToAddress(tokenPoolRef.Address),
						Args: token_pool.ApplyAllowListUpdatesArgs{
							Adds:    adds,
							Removes: removes,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to apply allow-list updates to token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
					}
					writes = append(writes, applyAllowListUpdatesReport.Output)
				}
			}
		}

		// Set router (if necessary)
		// Check the router currently set on the token pool
		if input.RouterAddress != (common.Address{}) {
			currentRouterReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRouter, chain, contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(tokenPoolRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get current router from token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
			}
			if currentRouterReport.Output != input.RouterAddress {
				// Router is not set to desired, so update it
				setRouterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRouter, chain, contract.FunctionInput[common.Address]{
					ChainSelector: input.ChainSelector,
					Address:       common.HexToAddress(tokenPoolRef.Address),
					Args:          input.RouterAddress,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set router on token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
				}
				writes = append(writes, setRouterReport.Output)
			}
		}

		// Set rate limit admin (if necessary)
		// Check the rate limit admin currently set on the token pool
		if input.RateLimitAdmin != (common.Address{}) {
			currentRateLimitAdminReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRateLimitAdmin, chain, contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(tokenPoolRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get current rate limit admin from token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
			}
			if currentRateLimitAdminReport.Output != input.RateLimitAdmin {
				// Rate limit admin is not set to desired, so update it
				setRateLimitAdminReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitAdmin, chain, contract.FunctionInput[common.Address]{
					ChainSelector: input.ChainSelector,
					Address:       common.HexToAddress(tokenPoolRef.Address),
					Args:          input.RateLimitAdmin,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limit admin on token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
				}
				writes = append(writes, setRateLimitAdminReport.Output)
			}
		}

		// Configure remote chains on the token pool as specified
		// This means adding any remote chains not already added, removing any remote chains that are no longer desired,
		// or modifying rate limiters on any existing remote chains.
		// TODO: Change to ConfigureTokenPoolForRemoteChains (plural)?
		for destChainSelector, remoteChainConfig := range input.RemoteChains {
			configureTokenPoolForRemoteChainReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChain, chain, ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    common.HexToAddress(tokenPoolRef.Address),
				RemoteChainSelector: destChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s for remote chain with selector %d: %w", tokenPoolRef.Address, chain, destChainSelector, err)
			}
			writes = append(writes, configureTokenPoolForRemoteChainReport.Output.Writes...)
		}

		// Register token on the token admin registry
		// If no admin address is specified, use the owner of the token admin registry
		var adminAddress common.Address
		if input.AdminAddress == (common.Address{}) {
			ownerReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.Owner, chain, contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenAdminRegistryAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get owner of token admin registry at address %s on %s: %w", input.TokenAdminRegistryAddress, chain, err)
			}
			adminAddress = ownerReport.Output
		}
		registerTokenReport, err := cldf_ops.ExecuteSequence(b, RegisterToken, chain, RegisterTokenInput{
			ChainSelector:             input.ChainSelector,
			TokenPoolAddress:          common.HexToAddress(tokenPoolRef.Address),
			AdminAddress:              adminAddress,
			OnlyPropose:               input.OnlyPropose,
			TokenAdminRegistryAddress: input.TokenAdminRegistryAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to register token pool with address %s on %s: %w", tokenPoolRef.Address, chain, err)
		}
		writes = append(writes, registerTokenReport.Output.Writes...)

		return sequences.OnChainOutput{
			Writes:    writes,
			Addresses: refs,
		}, nil
	},
)

// makeAllowListUpdates compares the current and desired allow-lists and returns the addresses to add and remove.
func makeAllowListUpdates(current, desired []common.Address) (adds, removes []common.Address) {
	currentSet := make(map[common.Address]struct{}, len(current))
	for _, addr := range current {
		currentSet[addr] = struct{}{}
	}
	desiredSet := make(map[common.Address]struct{}, len(desired))
	for _, addr := range desired {
		desiredSet[addr] = struct{}{}
	}

	for addr := range desiredSet {
		if _, exists := currentSet[addr]; !exists {
			adds = append(adds, addr)
		}
	}
	for addr := range currentSet {
		if _, exists := desiredSet[addr]; !exists {
			removes = append(removes, addr)
		}
	}
	return adds, removes
}
