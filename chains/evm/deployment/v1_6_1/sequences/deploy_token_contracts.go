package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc677"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/factory_burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *EVMAdapter) DeployTokenContracts() *cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return DeployTokenContracts
}

var DeployTokenContracts = cldf_ops.NewSequence(
	"deploy-token-contracts",
	semver.MustParse("1.0.0"),
	"Deploy given type of token contracts",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokens.DeployTokenInput) (sequences.OnChainOutput, error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		chain := chains.EVMChains()[input.ChainSelector]

		var tokenRef datastore.AddressRef
		var err error

		switch input.Type {
		case erc20.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, erc20.Deploy, chain, contract.DeployInput[erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chain.Selector,
				Args: erc20.ConstructorArgs{
					Name:   input.Name,
					Symbol: input.Symbol,
				},
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ERC20 token: %w", err)
			}

		case erc677.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, erc677.Deploy, chain, contract.DeployInput[erc677.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(erc677.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chain.Selector,
				Args: erc677.ConstructorArgs{
					Name:   input.Name,
					Symbol: input.Symbol,
				},
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ERC677 token: %w", err)
			}

		case burn_mint_erc20.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, burn_mint_erc20.Deploy, chain, contract.DeployInput[burn_mint_erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chain.Selector,
				Args: burn_mint_erc20.ConstructorArgs{
					Name:      input.Name,
					Symbol:    input.Symbol,
					Decimals:  input.Decimals,
					MaxSupply: input.Supply,
					PreMint:   input.PreMint, // pre-mint given amount to deployer address. Not advised to use against mainnet.
				},
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintERC20 token: %w", err)
			}

		case burn_mint_erc677.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, burn_mint_erc677.Deploy, chain, contract.DeployInput[burn_mint_erc677.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc677.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chain.Selector,
				Args: burn_mint_erc677.ConstructorArgs{
					Name:      input.Name,
					Symbol:    input.Symbol,
					Decimals:  input.Decimals,
					MaxSupply: input.Supply,
				},
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintERC677 token: %w", err)
			}

		case factory_burn_mint_erc20.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, factory_burn_mint_erc20.Deploy, chain, contract.DeployInput[factory_burn_mint_erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(factory_burn_mint_erc20.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chain.Selector,
				Args: factory_burn_mint_erc20.ConstructorArgs{
					Name:      input.Name,
					Symbol:    input.Symbol,
					Decimals:  input.Decimals,
					MaxSupply: input.Supply,
					PreMint:   input.PreMint,       // pre-mint given amount to deployer address. Not advised to use against mainnet.
					NewOwner:  input.ExternalAdmin, // Owner of the token contract
				},
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy FactoryBurnMintERC20 token: %w", err)
			}

		case burn_mint_erc20_with_drip.ContractType:
			tokenRef, err = contract.MaybeDeployContract(b, burn_mint_erc20_with_drip.Deploy, chain, contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chain.Selector,
				Args: burn_mint_erc20_with_drip.ConstructorArgs{
					Name:      input.Name,
					Symbol:    input.Symbol,
					Decimals:  input.Decimals,
					MaxSupply: input.Supply,
					PreMint:   input.PreMint, // pre-mint given amount to deployer address. Not advised to use against mainnet.
				},
			}, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintERC20WithDrip token: %w", err)
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported token type: %s", input.Type)
		}

		addresses = append(addresses, tokenRef)

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
