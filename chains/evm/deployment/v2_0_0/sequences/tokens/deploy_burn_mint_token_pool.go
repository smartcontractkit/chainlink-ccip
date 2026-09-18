package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_to_address_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_with_from_mint_token_pool"
)

var DeployBurnMintTokenPool = cldf_ops.NewSequence(
	"deploy-burn-mint-token-pool",
	semver.MustParse("2.0.0"),
	"Deploys a burn mint token pool to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)
		// AdvancedPoolHooks are not deployed by the token expansion flow. The pool is
		// constructed with the zero address so no hooks contract is wired in.
		constructorArgs := struct {
			Token              common.Address
			LocalTokenDecimals uint8
			AdvancedPoolHooks  common.Address
			RMNProxy           common.Address
			Router             common.Address
		}{
			Token:              input.ConstructorArgs.Token,
			LocalTokenDecimals: input.ConstructorArgs.Decimals,
			AdvancedPoolHooks:  common.Address{},
			RMNProxy:           input.ConstructorArgs.RMNProxy,
			Router:             input.ConstructorArgs.Router,
		}

		var tpDeployReport *datastore.AddressRef
		switch deployment.ContractType(input.TokenPoolType) {
		case burn_mint_token_pool.ContractType:
			report, deployErr := cldf_ops.ExecuteOperation(b, burn_mint_token_pool.Deploy, chain, evm_contract.DeployInput[burn_mint_token_pool.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: typeAndVersion,
				Args: burn_mint_token_pool.ConstructorArgs{
					Token:              constructorArgs.Token,
					LocalTokenDecimals: constructorArgs.LocalTokenDecimals,
					AdvancedPoolHooks:  constructorArgs.AdvancedPoolHooks,
					RmnProxy:           constructorArgs.RMNProxy,
					Router:             constructorArgs.Router,
				},
				Qualifier: &input.TokenSymbol,
			})
			tpDeployReport, err = &report.Output, deployErr
		case burn_from_mint_token_pool.ContractType:
			report, deployErr := cldf_ops.ExecuteOperation(b, burn_from_mint_token_pool.Deploy, chain, evm_contract.DeployInput[burn_from_mint_token_pool.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: typeAndVersion,
				Args: burn_from_mint_token_pool.ConstructorArgs{
					Token:              constructorArgs.Token,
					LocalTokenDecimals: constructorArgs.LocalTokenDecimals,
					AdvancedPoolHooks:  constructorArgs.AdvancedPoolHooks,
					RmnProxy:           constructorArgs.RMNProxy,
					Router:             constructorArgs.Router,
				},
				Qualifier: &input.TokenSymbol,
			})
			tpDeployReport, err = &report.Output, deployErr
		case burn_with_from_mint_token_pool.ContractType:
			report, deployErr := cldf_ops.ExecuteOperation(b, burn_with_from_mint_token_pool.Deploy, chain, evm_contract.DeployInput[burn_with_from_mint_token_pool.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: typeAndVersion,
				Args: burn_with_from_mint_token_pool.ConstructorArgs{
					Token:              constructorArgs.Token,
					LocalTokenDecimals: constructorArgs.LocalTokenDecimals,
					AdvancedPoolHooks:  constructorArgs.AdvancedPoolHooks,
					RmnProxy:           constructorArgs.RMNProxy,
					Router:             constructorArgs.Router,
				},
				Qualifier: &input.TokenSymbol,
			})
			tpDeployReport, err = &report.Output, deployErr
		case burn_to_address_mint_token_pool.ContractType:
			report, deployErr := cldf_ops.ExecuteOperation(b, burn_to_address_mint_token_pool.Deploy, chain, evm_contract.DeployInput[burn_to_address_mint_token_pool.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: typeAndVersion,
				Args: burn_to_address_mint_token_pool.ConstructorArgs{
					Token:              constructorArgs.Token,
					LocalTokenDecimals: constructorArgs.LocalTokenDecimals,
					AdvancedPoolHooks:  constructorArgs.AdvancedPoolHooks,
					RmnProxy:           constructorArgs.RMNProxy,
					Router:             constructorArgs.Router,
					BurnAddress:        input.ConstructorArgs.BurnAddress,
				},
				Qualifier: &input.TokenSymbol,
			})
			tpDeployReport, err = &report.Output, deployErr
		default:
			return sequences.OnChainOutput{}, fmt.Errorf("unsupported burn mint token pool type %s", input.TokenPoolType)
		}
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
		}

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:                    input.ChainSel,
			TokenPoolAddress:                 common.HexToAddress(tpDeployReport.Address),
			RateLimitAdmin:                   input.RateLimitAdmin,
			RouterAddress:                    input.ConstructorArgs.Router,
			ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
			FeeAdmin:                         input.FeeAdmin,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", tpDeployReport.Address, chain, err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{*tpDeployReport},
			BatchOps:  configureReport.Output.BatchOps,
		}, nil
	},
)
