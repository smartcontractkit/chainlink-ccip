package tokens

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	aphbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/advanced_pool_hooks"
	lockboxbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
)

var DeployLockReleaseTokenPool = cldf_ops.NewSequence(
	"deploy-lock-release-token-pool",
	semver.MustParse("2.0.0"),
	"Deploys a lock release token pool to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		lockBoxDeployRef, err := ops2contract.MaybeDeployContract(b, erc20_lock_box.Deploy, chain, ops2contract.DeployInput[erc20_lock_box.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
			Args: erc20_lock_box.ConstructorArgs{
				Token: input.ConstructorArgs.Token,
			},
			Qualifier: &input.TokenSymbol,
		}, nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ERC20 lock box to %s: %w", chain, err)
		}

		hooksDeployRef, err := ops2contract.MaybeDeployContract(b, advanced_pool_hooks.Deploy, chain, ops2contract.DeployInput[advanced_pool_hooks.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks.ContractType, *advanced_pool_hooks.Version),
			Args: advanced_pool_hooks.ConstructorArgs{
				Allowlist:                        input.AdvancedPoolHooksConfig.Allowlist,
				ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
				PolicyEngine:                     input.AdvancedPoolHooksConfig.PolicyEngine,
				AuthorizedCallers:                input.AdvancedPoolHooksConfig.AuthorizedCallers,
			},
			Qualifier: &input.TokenSymbol,
		}, nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy advanced pool hooks to %s: %w", chain, err)
		}

		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)
		tpDeployRef, err := ops2contract.MaybeDeployContract(b, lock_release_token_pool.Deploy, chain, ops2contract.DeployInput[lock_release_token_pool.ConstructorArgs]{
			TypeAndVersion: typeAndVersion,
			Args: lock_release_token_pool.ConstructorArgs{
				Token:              input.ConstructorArgs.Token,
				LocalTokenDecimals: input.ConstructorArgs.Decimals,
				AdvancedPoolHooks:  common.HexToAddress(hooksDeployRef.Address),
				RmnProxy:           input.ConstructorArgs.RMNProxy,
				Router:             input.ConstructorArgs.Router,
				LockBox:            common.HexToAddress(lockBoxDeployRef.Address),
			},
			Qualifier: &input.TokenSymbol,
		}, nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
		}

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:                    input.ChainSel,
			TokenPoolAddress:                 common.HexToAddress(tpDeployRef.Address),
			RateLimitAdmin:                   input.RateLimitAdmin,
			AdvancedPoolHooks:                common.HexToAddress(hooksDeployRef.Address),
			RouterAddress:                    input.ConstructorArgs.Router,
			ThresholdAmountForAdditionalCCVs: input.ThresholdAmountForAdditionalCCVs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", tpDeployRef.Address, chain, err)
		}

		{
			poolAddr := common.HexToAddress(tpDeployRef.Address)
			hooksAddr := common.HexToAddress(hooksDeployRef.Address)
			aph, err := bindAdvancedPoolHooks(hooksAddr, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			getAuthorizedCallersReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.NewReadGetAllAuthorizedCallers(aph), chain, ops2contract.FunctionInput[struct{}]{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers from advanced pool hooks %s on %s: %w", hooksAddr, chain, err)
			}

			if !slices.Contains(getAuthorizedCallersReport.Output, poolAddr) {
				applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.NewWriteApplyAuthorizedCallerUpdates(aph), chain, ops2contract.FunctionInput[aphbind.AuthorizedCallersAuthorizedCallerArgs]{
					Args: aphbind.AuthorizedCallersAuthorizedCallerArgs{
						AddedCallers: []common.Address{poolAddr},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize token pool %s on advanced pool hooks with address %s on %s: %w", poolAddr, hooksAddr, chain, err)
				}

				batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{writeOutputOps2ToLegacy(applyAuthorizedCallerUpdatesReport.Output)})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
				}
				configureReport.Output.BatchOps = append(configureReport.Output.BatchOps, batchOp)
			}
		}

		lockBoxAddr := common.HexToAddress(lockBoxDeployRef.Address)
		lockBox, err := lockboxbind.NewERC20LockBox(lockBoxAddr, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to bind lock box at %s: %w", lockBoxAddr.Hex(), err)
		}
		applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.NewWriteApplyAuthorizedCallerUpdates(lockBox), chain, ops2contract.FunctionInput[lockboxbind.AuthorizedCallersAuthorizedCallerArgs]{
			Args: lockboxbind.AuthorizedCallersAuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(tpDeployRef.Address),
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply authorized caller updates to lock box on %s: %w", chain, err)
		}
		batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{writeOutputOps2ToLegacy(applyAuthorizedCallerUpdatesReport.Output)})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		configureReport.Output.BatchOps = append(configureReport.Output.BatchOps, batchOp)

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{
				tpDeployRef,
				hooksDeployRef,
				lockBoxDeployRef,
			},
			BatchOps: configureReport.Output.BatchOps,
		}, nil
	},
)
