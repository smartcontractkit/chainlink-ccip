package tokens

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// TokenInfo is the information about the token to be deployed.
type TokenInfo struct {
	// Decimals is the number of decimals for the token.
	Decimals uint8
	// MaxSupply is the maximum supply of the token.
	MaxSupply *big.Int
	// Name is the name of the token.
	Name string
}

// DeployTokenAndPoolInput is the input for the DeployBurnMintTokenAndPool sequence.
type DeployTokenAndPoolInput struct {
	// Accounts is a map of account addresses to initial mint amounts.
	Accounts map[common.Address]*big.Int
	// TokenInfo is the information about the token to be deployed.
	// Token symbol will be taken from DeployTokenPoolInput.TokenSymbol.
	TokenInfo TokenInfo
	// DeployTokenPoolInput is the input for the DeployTokenPool sequence.
	DeployTokenPoolInput DeployTokenPoolInput
}

func (c DeployTokenAndPoolInput) ChainSelector() uint64 {
	return c.DeployTokenPoolInput.ChainSel
}

var DeployTokenAndPool = cldf_ops.NewSequence(
	"deploy-token-and-pool",
	semver.MustParse("1.7.0"),
	"Deploys a token and its associated token pool to an EVM chain, granting rights to the token pool and minting initial supply",
	func(b operations.Bundle, chain evm.Chain, input DeployTokenAndPoolInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)

		// Deploy burn mint token.
		deployTokenReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.Deploy, chain, evm_contract.DeployInput[burn_mint_erc677.ConstructorArgs]{
			ChainSelector:  input.DeployTokenPoolInput.ChainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc677.ContractType, *semver.MustParse("1.0.0")),
			Args: burn_mint_erc677.ConstructorArgs{
				Name:      input.TokenInfo.Name,
				Symbol:    input.DeployTokenPoolInput.TokenSymbol,
				Decimals:  input.TokenInfo.Decimals,
				MaxSupply: input.TokenInfo.MaxSupply,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token to %s: %w", chain, err)
		}
		deployTokenReport.Output.Qualifier = input.DeployTokenPoolInput.TokenSymbol // Use the token symbol as the qualifier.
		addresses = append(addresses, deployTokenReport.Output)

		// Deploy token pool.
		input.DeployTokenPoolInput.ConstructorArgs.Token = common.HexToAddress(deployTokenReport.Output.Address) // Set the token address to the deployed token.
		switch {
		case burn_mint_token_pool.IsSupported(deployment.ContractType(input.DeployTokenPoolInput.TokenPoolType), input.DeployTokenPoolInput.TokenPoolVersion):
			deployTokenPoolReport, err := cldf_ops.ExecuteSequence(b, DeployBurnMintTokenPool, chain, input.DeployTokenPoolInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token pool to %s: %w", chain, err)
			}
			addresses = append(addresses, deployTokenPoolReport.Output.Addresses...)
		/* TODO: Enable when the lockbox is finalized
		case lock_release_token_pool.IsSupported(deployment.ContractType(input.DeployTokenPoolInput.TokenPoolType), input.DeployTokenPoolInput.TokenPoolVersion):
			deployTokenPoolReport, err := cldf_ops.ExecuteSequence(b, DeployLockReleaseTokenPool, chain, input.DeployTokenPoolInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy lock release token pool to %s: %w", chain, err)
			}
			addresses = append(addresses, deployTokenPoolReport.Output.Addresses...)
		*/
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
		grantRolesReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.GrantMintAndBurnRoles, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          tokenPoolAddress,
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
		revokeMintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.RevokeMintRole, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke mint role from deployer on %s: %w", chain, err)
		}
		writes = append(writes, revokeMintReport.Output)

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
