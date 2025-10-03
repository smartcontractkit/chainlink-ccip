package tokens

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeployBurnMintTokenAndPoolInput struct {
	// Accounts is a map of account addresses to initial mint amounts.
	Accounts             map[common.Address]*big.Int
	DeployTokenPoolInput DeployTokenPoolInput
}

func (c DeployBurnMintTokenAndPoolInput) ChainSelector() uint64 {
	return c.DeployTokenPoolInput.ChainSel
}

var DeployBurnMintTokenAndPool = cldf_ops.NewSequence(
	"deploy-burn-mint-token-and-pool",
	semver.MustParse("1.7.0"),
	"Deploys a burn mint token and associated token pool to an EVM chain, granting burn mint rights to the token pool and minting initial supply",
	func(b operations.Bundle, chain evm.Chain, input DeployBurnMintTokenAndPoolInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0, 2)                  // We expect to deploy 2 contracts, token and pool.
		writes := make([]contract.WriteOutput, 0, 1+len(input.Accounts)) // One write for granting roles, one for each mint.

		// Deploy burn mint token.
		deployTokenReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.Deploy, chain, evm_contract.DeployInput[burn_mint_erc677.ConstructorArgs]{
			ChainSelector:  input.DeployTokenPoolInput.ChainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc677.ContractType, *semver.MustParse("1.0.0")),
			Args: burn_mint_erc677.ConstructorArgs{
				Name:   input.DeployTokenPoolInput.TokenSymbol,
				Symbol: input.DeployTokenPoolInput.TokenSymbol,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token to %s: %w", chain, err)
		}
		deployTokenReport.Output.Qualifier = input.DeployTokenPoolInput.TokenSymbol // Use the token symbol as the qualifier.
		addresses = append(addresses, deployTokenReport.Output)

		// Deploy token pool.
		input.DeployTokenPoolInput.ConstructorArgs.Token = common.HexToAddress(deployTokenReport.Output.Address) // Set the token address to the deployed token.
		deployTokenPoolReport, err := cldf_ops.ExecuteSequence(b, DeployTokenPool, chain, input.DeployTokenPoolInput)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token pool to %s: %w", chain, err)
		}
		addresses = append(addresses, deployTokenPoolReport.Output.Addresses...)

		// Grant mint and burn roles to the token pool.
		grantRolesReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.GrantMintAndBurnRoles, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          common.HexToAddress(deployTokenPoolReport.Output.Addresses[0].Address), // The first address returned is the token pool.
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint and burn roles to token pool on %s: %w", chain, err)
		}
		writes = append(writes, grantRolesReport.Output)

		// Grant mint role to the deployer key so we can mint initial supply.
		grantMintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.GrantMintRole, chain, evm_contract.FunctionInput[common.Address]{
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
			mintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.Mint, chain, evm_contract.FunctionInput[burn_mint_erc677.MintArgs]{
				ChainSelector: input.DeployTokenPoolInput.ChainSel,
				Address:       common.HexToAddress(deployTokenReport.Output.Address),
				Args: burn_mint_erc677.MintArgs{
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
		revokeMintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.RevokeBurnRole, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke mint role from deployer on %s: %w", chain, err)
		}
		writes = append(writes, revokeMintReport.Output)

		return sequences.OnChainOutput{
			Addresses: addresses,
			Writes:    writes,
		}, nil
	},
)
