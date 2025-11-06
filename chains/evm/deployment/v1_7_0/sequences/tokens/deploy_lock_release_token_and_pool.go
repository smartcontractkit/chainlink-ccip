package tokens

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
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

type DeployLockReleaseTokenAndPoolInput struct {
	Accounts             map[common.Address]*big.Int
	TokenInfo            TokenInfo
	DeployTokenPoolInput DeployTokenPoolInput
	PoolFundingAmount    *big.Int
}

func (c DeployLockReleaseTokenAndPoolInput) ChainSelector() uint64 {
	return c.DeployTokenPoolInput.ChainSel
}

var DeployLockReleaseTokenAndPool = cldf_ops.NewSequence(
	"deploy-lock-release-token-and-pool",
	semver.MustParse("1.7.0"),
	"Deploys a burn mint token and associated lock release token pool to an EVM chain, funding the pool and minting initial supply",
	func(b operations.Bundle, chain evm.Chain, input DeployLockReleaseTokenAndPoolInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0, 2)
		writes := make([]contract.WriteOutput, 0, 2+len(input.Accounts))

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
		deployTokenReport.Output.Qualifier = input.DeployTokenPoolInput.TokenSymbol
		addresses = append(addresses, deployTokenReport.Output)

		input.DeployTokenPoolInput.ConstructorArgs.Token = common.HexToAddress(deployTokenReport.Output.Address)
		deployTokenPoolReport, err := cldf_ops.ExecuteSequence(b, DeployTokenPool, chain, input.DeployTokenPoolInput)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy token pool to %s: %w", chain, err)
		}
		addresses = append(addresses, deployTokenPoolReport.Output.Addresses...)

		grantMintReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.GrantMintRole, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.DeployTokenPoolInput.ChainSel,
			Address:       common.HexToAddress(deployTokenReport.Output.Address),
			Args:          chain.DeployerKey.From,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to grant mint role to deployer on %s: %w", chain, err)
		}
		writes = append(writes, grantMintReport.Output)

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

		if input.PoolFundingAmount != nil && input.PoolFundingAmount.Cmp(big.NewInt(0)) > 0 {
			poolFundingReport, err := cldf_ops.ExecuteOperation(b, burn_mint_erc677.Mint, chain, evm_contract.FunctionInput[burn_mint_erc677.MintArgs]{
				ChainSelector: input.DeployTokenPoolInput.ChainSel,
				Address:       common.HexToAddress(deployTokenReport.Output.Address),
				Args: burn_mint_erc677.MintArgs{
					Account: common.HexToAddress(deployTokenPoolReport.Output.Addresses[0].Address),
					Amount:  input.PoolFundingAmount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to mint %s tokens to pool %s on %s: %w", input.PoolFundingAmount.String(), deployTokenPoolReport.Output.Addresses[0].Address, chain, err)
			}
			writes = append(writes, poolFundingReport.Output)
		}

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
