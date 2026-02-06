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

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
)

var DeployBurnMintTokenPool = cldf_ops.NewSequence(
	"deploy-burn-mint-token-pool",
	semver.MustParse("1.7.0"),
	"Deploys a burn mint token pool to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		hooksDeployReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.Deploy, chain, evm_contract.DeployInput[advanced_pool_hooks.ConstructorArgs]{
			ChainSelector:  input.ChainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks.ContractType, *advanced_pool_hooks.Version),
			Args: advanced_pool_hooks.ConstructorArgs{
				Allowlist:                        input.AdvancedPoolHooksConfig.Allowlist,
				ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
				PolicyEngine:                     input.AdvancedPoolHooksConfig.PolicyEngine,
				AuthorizedCallers:                input.AdvancedPoolHooksConfig.AuthorizedCallers,
			},
			Qualifier: &input.TokenSymbol,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy advanced pool hooks to %s: %w", chain, err)
		}

		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)
		tpDeployReport, err := cldf_ops.ExecuteOperation(b, burn_mint_token_pool.Deploy, chain, evm_contract.DeployInput[burn_mint_token_pool.ConstructorArgs]{
			ChainSelector:  input.ChainSel,
			TypeAndVersion: typeAndVersion,
			Args: burn_mint_token_pool.ConstructorArgs{
				Token:              input.ConstructorArgs.Token,
				LocalTokenDecimals: input.ConstructorArgs.Decimals,
				AdvancedPoolHooks:  common.HexToAddress(hooksDeployReport.Output.Address),
				RMNProxy:           input.ConstructorArgs.RMNProxy,
				Router:             input.ConstructorArgs.Router,
			},
			Qualifier: &input.TokenSymbol,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
		}

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:                    input.ChainSel,
			TokenPoolAddress:                 common.HexToAddress(tpDeployReport.Output.Address),
			RateLimitAdmin:                   input.RateLimitAdmin,
			AdvancedPoolHooks:                common.HexToAddress(hooksDeployReport.Output.Address),
			RouterAddress:                    input.ConstructorArgs.Router,
			ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
			FeeAggregator:                    input.FeeAggregator,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", tpDeployReport.Output.Address, chain, err)
		}

		// If hooks authorized callers gating is enabled at deployment time, ensure the newly deployed token pool is authorized.
		// Otherwise, calls to preflightCheck/postflightCheck will revert when executed by the token pool.
		if len(input.AdvancedPoolHooksConfig.AuthorizedCallers) > 0 {
			poolAddr := common.HexToAddress(tpDeployReport.Output.Address)
			hooksAddr := common.HexToAddress(hooksDeployReport.Output.Address)

			// Check if the pool is already an authorized caller.
			getAuthorizedCallersReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetAllAuthorizedCallers, chain, evm_contract.FunctionInput[any]{
				ChainSelector: input.ChainSel,
				Address:       hooksAddr,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers from advanced pool hooks %s on %s: %w", hooksAddr, chain, err)
			}

			alreadyAuthorized := false
			for _, caller := range getAuthorizedCallersReport.Output {
				if caller == poolAddr {
					alreadyAuthorized = true
					break
				}
			}

			if !alreadyAuthorized {
				applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyAuthorizedCallerUpdates, chain, evm_contract.FunctionInput[advanced_pool_hooks.AuthorizedCallerArgs]{
					ChainSelector: input.ChainSel,
					Address:       hooksAddr,
					Args: advanced_pool_hooks.AuthorizedCallerArgs{
						AddedCallers: []common.Address{poolAddr},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize token pool %s on advanced pool hooks with address %s on %s: %w", poolAddr, hooksAddr, chain, err)
				}

				batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{applyAuthorizedCallerUpdatesReport.Output})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
				}
				configureReport.Output.BatchOps = append(configureReport.Output.BatchOps, batchOp)
			}
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{tpDeployReport.Output, hooksDeployReport.Output},
			BatchOps:  configureReport.Output.BatchOps,
		}, nil
	},
)
