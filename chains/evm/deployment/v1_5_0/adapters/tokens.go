package adapters

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tokseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var _ tokensapi.TokenAdapter = &TokenAdapter{}

type TokenAdapter struct{}

////////////////////
// Helper methods //
////////////////////

func (a *TokenAdapter) getTokenAdminRegistryAddress(addresses datastore.AddressRefStore, selector uint64) (common.Address, error) {
	refs := addresses.Filter(
		datastore.AddressRefByType(datastore.ContractType(tarops.ContractType)),
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByVersion(tarops.Version),
	)
	if len(refs) != 1 {
		return common.Address{}, fmt.Errorf("unexpectedly found %d matches for %q with version %q on chain %d", len(refs), tarops.ContractType, tarops.Version.String(), selector)
	}

	ref := refs[0]
	if !common.IsHexAddress(ref.Address) {
		return common.Address{}, fmt.Errorf("token admin registry address %q is not a valid hex address", ref.Address)
	}

	return common.HexToAddress(ref.Address), nil
}

func (a *TokenAdapter) getTokenPoolAddress(addresses datastore.AddressRefStore, selector uint64, qualifier string, poolType string) (common.Address, error) {
	// Define the version range for 1.5.x
	minVersion := semver.MustParse("1.5.0") // inclusive
	maxVersion := semver.MustParse("1.6.0") // exclusive

	// Get all matching token pool addresses
	refs := addresses.Filter(
		datastore.AddressRefByType(datastore.ContractType(poolType)),
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByQualifier(qualifier),
	)

	// Use the latest version in the 1.5.x series
	var latestRef *datastore.AddressRef
	latestVer := minVersion
	for _, ref := range refs {
		v := ref.Version
		if v.GreaterThanEqual(latestVer) && v.LessThan(maxVersion) {
			latestRef = &ref
			latestVer = v
		}
	}

	// If no matching reference was found, then return an error
	if latestRef == nil {
		return common.Address{}, fmt.Errorf("no token pool found for type %q with qualifier %q on chain %d", poolType, qualifier, selector)
	}

	// Double-check that the address is valid
	if !common.IsHexAddress(latestRef.Address) {
		return common.Address{}, fmt.Errorf("token pool address %q is not a valid hex address", latestRef.Address)
	}

	// Return the EVM token pool address
	return common.HexToAddress(latestRef.Address), nil
}

func (a *TokenAdapter) getTokenPool(addresses datastore.AddressRefStore, selector uint64, qualifier string, poolType string) (*token_pool.TokenPool, error) {
	addr, err := a.getTokenPoolAddress(addresses, selector, qualifier, poolType)
	if err != nil {
		return nil, fmt.Errorf("failed to get token pool address: %w", err)
	}

	tp, err := token_pool.NewTokenPool(addr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token pool contract at %q: %w", addr.Hex(), err)
	}

	return tp, nil
}

//////////////////////////////////////////////
// Implementation of TokenAdapter interface //
//////////////////////////////////////////////

func (a *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"token-adapter:configure-token-for-transfers",
		semver.MustParse("1.5.0"),
		"Configure a token for cross-chain transfers across multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			// TODO: implement me
			return sequences.OnChainOutput{}, errors.New("not implemented")
		})
}

func (a *TokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if !common.IsHexAddress(ref.Address) {
		return nil, fmt.Errorf("address %q is not a valid hex address", ref.Address)
	}

	return common.HexToAddress(ref.Address).Bytes(), nil
}

func (a *TokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not defined", chainSelector)
	}

	if !common.IsHexAddress(poolRef.Address) {
		return nil, fmt.Errorf("pool address %q is not a valid hex address", poolRef.Address)
	}

	tp, err := token_pool.NewTokenPool(common.HexToAddress(poolRef.Address), chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token pool contract at %q: %w", poolRef.Address, err)
	}

	tokenAddress, err := tp.GetToken(&bind.CallOpts{Context: e.GetContext()})
	if err != nil {
		return nil, fmt.Errorf("failed to get token address from token pool at %q: %w", poolRef.Address, err)
	}

	return tokenAddress.Bytes(), nil
}

func (a *TokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokensapi.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"token-adapter:manual-registration",
		semver.MustParse("1.5.0"),
		"Manually register a token and token pool on multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ManualRegistrationInput) (sequences.OnChainOutput, error) {
			ds := datastore.NewMemoryAddressRefStore()
			for _, addr := range input.ExistingAddresses {
				if err := ds.Upsert(addr); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to upsert address %v: %w", addr, err)
				}
			}

			tokenAdminRegistryAddress, err := a.getTokenAdminRegistryAddress(ds, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			tokenPool, err := a.getTokenPool(ds, input.ChainSelector, input.RegisterTokenConfigs.TokenPoolQualifier, input.RegisterTokenConfigs.PoolType)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token pool with qualifier %q on chain %d: %w", input.RegisterTokenConfigs.TokenPoolQualifier, input.ChainSelector, err)
			}

			tokenAddress, err := tokenPool.GetToken(&bind.CallOpts{Context: b.GetContext()})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool at %q: %w", tokenPool.Address().Hex(), err)
			}

			proposedOwner := input.RegisterTokenConfigs.ProposedOwner
			if !common.IsHexAddress(proposedOwner) {
				return sequences.OnChainOutput{}, fmt.Errorf("proposed owner address %q is not a valid hex address", proposedOwner)
			}

			var result sequences.OnChainOutput
			result, err = sequences.RunAndMergeSequence(b, chains,
				tarseq.ManualRegistrationSequence,
				tarseq.ManualRegistrationSequenceInput{
					AdminAddress:  common.HexToAddress(proposedOwner),
					ChainSelector: input.ChainSelector,
					TokenAddress:  tokenAddress,
					Address:       tokenAdminRegistryAddress,
				},
				result,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to manually register token on chain %d: %w", input.ChainSelector, err)
			}

			return result, nil
		})
}

func (a *TokenAdapter) DeployToken() *cldf_ops.Sequence[tokensapi.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return tokseq.DeployToken
}

func (a *TokenAdapter) DeployTokenVerify(in tokensapi.IN) error {
	// TODO: implement me
	return errors.New("not implemented")
}

func (a *TokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokensapi.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"token-adapter:deploy-token-pool-for-token",
		semver.MustParse("1.5.0"),
		"Deploy a token pool for an existing token on multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.DeployTokenPoolInput) (sequences.OnChainOutput, error) {
			// TODO: implement me
			return sequences.OnChainOutput{}, errors.New("not implemented")
		})
}

func (a *TokenAdapter) RegisterToken() *cldf_ops.Sequence[tokensapi.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"token-adapter:register-token",
		semver.MustParse("1.5.0"),
		"Register a token and its pool on multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.RegisterTokenInput) (sequences.OnChainOutput, error) {
			// TODO: implement me
			return sequences.OnChainOutput{}, errors.New("not implemented")
		})
}

func (a *TokenAdapter) SetPool() *cldf_ops.Sequence[tokensapi.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"token-adapter:set-pool",
		semver.MustParse("1.5.0"),
		"Set the pool for a token across multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.SetPoolInput) (sequences.OnChainOutput, error) {
			addresses := input.ExistingDataStore.Addresses()

			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}

			tokenAdminRegistryAddress, err := a.getTokenAdminRegistryAddress(addresses, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			tokenPool, err := a.getTokenPool(addresses, input.ChainSelector, input.TokenPoolQualifier, input.PoolType)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token pool with qualifier %q on chain %d: %w", input.TokenPoolQualifier, input.ChainSelector, err)
			}

			tokenAddress, err := tokenPool.GetToken(&bind.CallOpts{Context: b.GetContext()})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool at %q: %w", tokenPool.Address().Hex(), err)
			}

			report, err := cldf_ops.ExecuteOperation(b,
				tarops.SetPool,
				chain,
				contract.FunctionInput[tarops.SetPoolArgs]{
					Address:       tokenAdminRegistryAddress,
					ChainSelector: input.ChainSelector,
					Args: tarops.SetPoolArgs{
						TokenPoolAddress: tokenPool.Address(),
						TokenAddress:     tokenAddress,
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set pool for token %q on chain selector %d: %w", input.TokenSymbol, input.ChainSelector, err)
			}

			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}

			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batchOp},
			}, nil
		},
	)
}

func (a *TokenAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokensapi.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"token-adapter:update-authorities",
		semver.MustParse("1.5.0"),
		"Update the authorities for a token on multiple EVM chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.UpdateAuthoritiesInput) (sequences.OnChainOutput, error) {
			// TODO: implement me
			return sequences.OnChainOutput{}, errors.New("not implemented")
		})
}
