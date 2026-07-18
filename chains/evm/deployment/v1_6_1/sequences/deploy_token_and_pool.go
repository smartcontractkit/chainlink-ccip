package sequences

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_token_pool"
	bmBindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// DeployTokenAndPoolInput is the input for the DeployBurnMintTokenAndPool sequence.
type DeployTokenAndPoolInput struct {
	// Accounts is a map of account addresses to initial mint amounts.
	Accounts map[common.Address]*big.Int
	// DeployTokenPoolInput is the input for the DeployTokenPool sequence.
	DeployTokenPoolInput DeployTokenPoolInput
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
		writes := make([]contract.WriteOutput, 0)

		// Deploy burn mint token.
		deployTokenReport, err := evmops.ExecuteDeploy(b, burn_mint_erc20_with_drip.Deploy, chain, contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
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

		// Deploy token pool.
		input.DeployTokenPoolInput.ConstructorArgs.Token = common.HexToAddress(deployTokenReport.Output.Address) // Set the token address to the deployed token.
		switch {
		case input.DeployTokenPoolInput.TokenPoolType == datastore.ContractType(burn_mint_token_pool.ContractType) && input.DeployTokenPoolInput.TokenPoolVersion.Equal(burn_mint_token_pool.Version):
			deployTokenPoolReport, err := cldf_ops.ExecuteSequence(b, DeployBurnMintTokenPool, chain, input.DeployTokenPoolInput)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy burn mint token pool to %s: %w", chain, err)
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
		grantRolesReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(deployTokenReport.Output.Address), bmBindings.NewBurnMintERC20WithDrip, burn_mint_erc20_with_drip.NewWriteGrantMintAndBurnRoles, tokenPoolAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint and burn roles to token pool on %s: %w", chain, err)
		}
		writes = append(writes, grantRolesReport.Output)

		// Grant roles to the deployer key so we can mint initial supply.
		grantMintReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(deployTokenReport.Output.Address), bmBindings.NewBurnMintERC20WithDrip, burn_mint_erc20_with_drip.NewWriteGrantMintAndBurnRoles, chain.DeployerKey.From)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint role to deployer on %s: %w", chain, err)
		}
		writes = append(writes, grantMintReport.Output)

		// Mint initial supply to each account.
		for account, amount := range input.Accounts {
			mintReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(deployTokenReport.Output.Address), bmBindings.NewBurnMintERC20WithDrip, burn_mint_erc20_with_drip.NewWriteMint, burn_mint_erc20_with_drip.MintArgs{
				Account: account,
				Amount:  amount,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to mint %s tokens to %s on %s: %w", amount.String(), account.Hex(), chain, err)
			}
			writes = append(writes, mintReport.Output)
		}

		// Revoke mint role from the deployer key for safety.
		revokeMintReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(deployTokenReport.Output.Address), bmBindings.NewBurnMintERC20WithDrip, burn_mint_erc20_with_drip.NewWriteRevokeMintRole, chain.DeployerKey.From)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke mint role from deployer on %s: %w", chain, err)
		}
		writes = append(writes, revokeMintReport.Output)

		// Revoke burn role from the deployer key for safety.
		revokeBurnReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(deployTokenReport.Output.Address), bmBindings.NewBurnMintERC20WithDrip, burn_mint_erc20_with_drip.NewWriteRevokeBurnRole, chain.DeployerKey.From)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to revoke burn role from deployer on %s: %w", chain, err)
		}
		writes = append(writes, revokeBurnReport.Output)

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
