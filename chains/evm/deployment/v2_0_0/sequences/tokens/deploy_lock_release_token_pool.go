package tokens

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
)

var DeployLockReleaseTokenPool = cldf_ops.NewSequence(
	"deploy-lock-release-token-pool",
	semver.MustParse("2.0.0"),
	"Deploys a lock release token pool to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)

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

		// Check if this is a siloed pool type
		isSiloedPool := input.TokenPoolType == datastore.ContractType(siloed_lock_release_token_pool.ContractType)

		var tpDeployReport *datastore.AddressRef
		var addresses []datastore.AddressRef
		var batchOps []mcms_types.BatchOperation

		if isSiloedPool {
			// Validate lock box groups for siloed pool
			if err := input.ValidateLockBoxGroups(); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid lock box groups: %w", err)
			}

			// Deploy siloed pool without lockboxes (they'll be configured later)
			siloedPoolReport, err := cldf_ops.ExecuteOperation(b, siloed_lock_release_token_pool.Deploy, chain, evm_contract.DeployInput[siloed_lock_release_token_pool.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: typeAndVersion,
				Args: siloed_lock_release_token_pool.ConstructorArgs{
					Token:              input.ConstructorArgs.Token,
					LocalTokenDecimals: input.ConstructorArgs.Decimals,
					AdvancedPoolHooks:  common.HexToAddress(hooksDeployReport.Output.Address),
					RmnProxy:           input.ConstructorArgs.RMNProxy,
					Router:             input.ConstructorArgs.Router,
				},
				Qualifier: &input.TokenSymbol,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
			}
			tpDeployReport = &siloedPoolReport.Output

			// Deploy one lockbox per group and configure them on the pool
			lockBoxConfigs := make([]siloed_lock_release_token_pool.LockBoxConfig, 0)
			poolAddr := common.HexToAddress(tpDeployReport.Address)

			for groupIdx, group := range input.LockBoxGroups {
				// Deploy lockbox for this group
				lbQualifier := fmt.Sprintf("group-%d", groupIdx)
				lockBoxDeployReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.Deploy, chain, evm_contract.DeployInput[erc20_lock_box.ConstructorArgs]{
					ChainSelector:  input.ChainSel,
					TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
					Args: erc20_lock_box.ConstructorArgs{
						Token: input.ConstructorArgs.Token,
					},
					Qualifier: &lbQualifier,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ERC20 lock box for group %d to %s: %w", groupIdx, chain, err)
				}

				lockBoxAddr := common.HexToAddress(lockBoxDeployReport.Output.Address)
				addresses = append(addresses, lockBoxDeployReport.Output)

				// Authorize pool on this lockbox
				authorizeReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, evm_contract.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
					ChainSelector: input.ChainSel,
					Address:       lockBoxAddr,
					Args: erc20_lock_box.AuthorizedCallerArgs{
						AddedCallers: []common.Address{poolAddr},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize pool on lock box for group %d on %s: %w", groupIdx, chain, err)
				}
				batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{authorizeReport.Output})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
				}
				batchOps = append(batchOps, batchOp)

				// Map all remote chains in this group to this lockbox
				for _, remoteChainSel := range group {
					lockBoxConfigs = append(lockBoxConfigs, siloed_lock_release_token_pool.LockBoxConfig{
						RemoteChainSelector: remoteChainSel,
						LockBox:             lockBoxAddr,
					})
				}
			}

			// Configure all lockboxes on the pool
			configLBReport, err := cldf_ops.ExecuteOperation(b, siloed_lock_release_token_pool.ConfigureLockBoxes, chain, evm_contract.FunctionInput[[]siloed_lock_release_token_pool.LockBoxConfig]{
				ChainSelector: input.ChainSel,
				Address:       poolAddr,
				Args:          lockBoxConfigs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure lock boxes on siloed pool on %s: %w", chain, err)
			}
			batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{configLBReport.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			batchOps = append(batchOps, batchOp)

		} else {
			// Standard lock release pool deployment
			lockBoxDeployReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.Deploy, chain, evm_contract.DeployInput[erc20_lock_box.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
				Args: erc20_lock_box.ConstructorArgs{
					Token: input.ConstructorArgs.Token,
				},
				Qualifier: &input.TokenSymbol,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ERC20 lock box to %s: %w", chain, err)
			}

			standardPoolReport, err := cldf_ops.ExecuteOperation(b, lock_release_token_pool.Deploy, chain, evm_contract.DeployInput[lock_release_token_pool.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: typeAndVersion,
				Args: lock_release_token_pool.ConstructorArgs{
					Token:              input.ConstructorArgs.Token,
					LocalTokenDecimals: input.ConstructorArgs.Decimals,
					AdvancedPoolHooks:  common.HexToAddress(hooksDeployReport.Output.Address),
					RmnProxy:           input.ConstructorArgs.RMNProxy,
					Router:             input.ConstructorArgs.Router,
					LockBox:            common.HexToAddress(lockBoxDeployReport.Output.Address),
				},
				Qualifier: &input.TokenSymbol,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
			}
			tpDeployReport = &standardPoolReport.Output
			addresses = append(addresses, lockBoxDeployReport.Output)

			// Add lock release token pool to the authorized callers of the lock box.
			applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, evm_contract.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
				ChainSelector: input.ChainSel,
				Address:       common.HexToAddress(lockBoxDeployReport.Output.Address),
				Args: erc20_lock_box.AuthorizedCallerArgs{
					AddedCallers: []common.Address{
						common.HexToAddress(standardPoolReport.Output.Address),
					},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to lock box on %s: %w", chain, err)
			}
			batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{applyAuthorizedCallerUpdatesReport.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			batchOps = append(batchOps, batchOp)
		}

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:                    input.ChainSel,
			TokenPoolAddress:                 common.HexToAddress(tpDeployReport.Address),
			RateLimitAdmin:                   input.RateLimitAdmin,
			AdvancedPoolHooks:                common.HexToAddress(hooksDeployReport.Output.Address),
			RouterAddress:                    input.ConstructorArgs.Router,
			ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
			FeeAdmin:                    input.FeeAdmin,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", tpDeployReport.Address, chain, err)
		}

		// Add the newly deployed token pool as an authorized caller on the hooks.
		{
			poolAddr := common.HexToAddress(tpDeployReport.Address)
			hooksAddr := common.HexToAddress(hooksDeployReport.Output.Address)

			getAuthorizedCallersReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetAllAuthorizedCallers, chain, evm_contract.FunctionInput[struct{}]{
				ChainSelector: input.ChainSel,
				Address:       hooksAddr,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers from advanced pool hooks %s on %s: %w", hooksAddr, chain, err)
			}

			if !slices.Contains(getAuthorizedCallersReport.Output, poolAddr) {
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
				batchOps = append(batchOps, batchOp)
			}
		}

		return sequences.OnChainOutput{
			Addresses: append([]datastore.AddressRef{*tpDeployReport, hooksDeployReport.Output}, addresses...),
			BatchOps:  append(batchOps, configureReport.Output.BatchOps...),
		}, nil
	},
)
