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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
)

// LockBoxQualifier builds the datastore qualifier for the lockbox backing one silo group. The
// group's lowest remote chain selector identifies the group: a chain belongs to at most one group
// (enforced by DeployTokenPoolInput.ValidateLockBoxGroups), so the minimum is unique across groups
// and stays stable regardless of the order groups are listed in.
func LockBoxQualifier(poolQualifier string, group []uint64) string {
	return fmt.Sprintf("%s-silo(%d)", poolQualifier, slices.Min(group))
}

// DeploySiloedLockReleaseTokenPool deploys a SiloedLockReleaseTokenPool along with one ERC20LockBox
// per silo group.
//
// Unlike the plain LockReleaseTokenPool - which takes a single lockbox in its constructor - the
// siloed pool holds no liquidity itself and no internal accounting: lockOrBurn deposits into, and
// releaseOrMint withdraws from, the lockbox mapped to that remote chain. Siloing is therefore
// expressed purely by which chains share a lockbox, via configureLockBoxes.
var DeploySiloedLockReleaseTokenPool = cldf_ops.NewSequence(
	"deploy-siloed-lock-release-token-pool",
	semver.MustParse("2.0.0"),
	"Deploys a siloed lock release token pool and its per-silo lock boxes to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}
		if err := input.ValidateLockBoxGroups(); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid lock box groups: %w", err)
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
		hooksAddr := common.HexToAddress(hooksDeployReport.Output.Address)

		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)
		tpDeployReport, err := cldf_ops.ExecuteOperation(b, siloed_lock_release_token_pool.Deploy, chain, evm_contract.DeployInput[siloed_lock_release_token_pool.ConstructorArgs]{
			ChainSelector:  input.ChainSel,
			TypeAndVersion: typeAndVersion,
			Args: siloed_lock_release_token_pool.ConstructorArgs{
				Token:              input.ConstructorArgs.Token,
				LocalTokenDecimals: input.ConstructorArgs.Decimals,
				AdvancedPoolHooks:  hooksAddr,
				RmnProxy:           input.ConstructorArgs.RMNProxy,
				Router:             input.ConstructorArgs.Router,
			},
			Qualifier: &input.TokenSymbol,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
		}
		poolAddr := common.HexToAddress(tpDeployReport.Output.Address)

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:                    input.ChainSel,
			TokenPoolAddress:                 poolAddr,
			RateLimitAdmin:                   input.RateLimitAdmin,
			AdvancedPoolHooks:                hooksAddr,
			RouterAddress:                    input.ConstructorArgs.Router,
			ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", poolAddr, chain, err)
		}

		var writes []evm_contract.WriteOutput

		// Add the newly deployed token pool as an authorized caller on the hooks.
		getAuthorizedCallersReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetAllAuthorizedCallers, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSel,
			Address:       hooksAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers from advanced pool hooks %s on %s: %w", hooksAddr, chain, err)
		}
		if !slices.Contains(getAuthorizedCallersReport.Output, poolAddr) {
			applyHooksCallersReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyAuthorizedCallerUpdates, chain, evm_contract.FunctionInput[advanced_pool_hooks.AuthorizedCallerArgs]{
				ChainSelector: input.ChainSel,
				Address:       hooksAddr,
				Args: advanced_pool_hooks.AuthorizedCallerArgs{
					AddedCallers: []common.Address{poolAddr},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize token pool %s on advanced pool hooks with address %s on %s: %w", poolAddr, hooksAddr, chain, err)
			}
			writes = append(writes, applyHooksCallersReport.Output)
		}

		// Deploy one lockbox per silo group and authorize the pool on each. Every remote chain in a
		// group maps to that group's lockbox, so chains within a group share liquidity while chains
		// in different groups are isolated.
		// NOTE: Addresses[0] must remain the pool ref - DeployTokenPool relies on that convention.
		addresses := []datastore.AddressRef{tpDeployReport.Output, hooksDeployReport.Output}
		lockBoxConfigs := make([]siloed_lock_release_token_pool.LockBoxConfig, 0, len(input.LockBoxGroups))
		for _, group := range input.LockBoxGroups {
			qualifier := LockBoxQualifier(input.TokenSymbol, group)
			lbDeployReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.Deploy, chain, evm_contract.DeployInput[erc20_lock_box.ConstructorArgs]{
				ChainSelector:  input.ChainSel,
				TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
				Args: erc20_lock_box.ConstructorArgs{
					Token: input.ConstructorArgs.Token,
				},
				Qualifier: &qualifier,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ERC20 lock box %q to %s: %w", qualifier, chain, err)
			}
			addresses = append(addresses, lbDeployReport.Output)
			lockBoxAddr := common.HexToAddress(lbDeployReport.Output.Address)

			for _, remoteChainSelector := range group {
				lockBoxConfigs = append(lockBoxConfigs, siloed_lock_release_token_pool.LockBoxConfig{
					RemoteChainSelector: remoteChainSelector,
					LockBox:             lockBoxAddr,
				})
			}

			applyLockBoxCallersReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, evm_contract.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
				ChainSelector: input.ChainSel,
				Address:       lockBoxAddr,
				Args: erc20_lock_box.AuthorizedCallerArgs{
					AddedCallers: []common.Address{poolAddr},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize token pool %s on lock box %s on %s: %w", poolAddr, lockBoxAddr, chain, err)
			}
			writes = append(writes, applyLockBoxCallersReport.Output)
		}

		// Map every remote chain to its lockbox in a single call.
		configureLockBoxesReport, err := cldf_ops.ExecuteOperation(b, siloed_lock_release_token_pool.ConfigureLockBoxes, chain, evm_contract.FunctionInput[[]siloed_lock_release_token_pool.LockBoxConfig]{
			ChainSelector: input.ChainSel,
			Address:       poolAddr,
			Args:          lockBoxConfigs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure lock boxes on token pool %s on %s: %w", poolAddr, chain, err)
		}
		writes = append(writes, configureLockBoxesReport.Output)

		batchOps := configureReport.Output.BatchOps
		if len(writes) > 0 {
			batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			batchOps = append(batchOps, []mcms_types.BatchOperation{batchOp}...)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	},
)
