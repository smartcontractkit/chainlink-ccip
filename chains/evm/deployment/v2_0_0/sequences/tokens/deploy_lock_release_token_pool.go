package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
)

var DeployLockReleaseTokenPool = cldf_ops.NewSequence(
	"deploy-lock-release-token-pool",
	semver.MustParse("2.0.0"),
	"Deploys a lock release token pool to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

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

		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)
		// AdvancedPoolHooks are not deployed by the token expansion flow. The pool is
		// constructed with the zero address so no hooks contract is wired in.
		tpDeployReport, err := cldf_ops.ExecuteOperation(b, lock_release_token_pool.Deploy, chain, evm_contract.DeployInput[lock_release_token_pool.ConstructorArgs]{
			ChainSelector:  input.ChainSel,
			TypeAndVersion: typeAndVersion,
			Args: lock_release_token_pool.ConstructorArgs{
				Token:              input.ConstructorArgs.Token,
				LocalTokenDecimals: input.ConstructorArgs.Decimals,
				AdvancedPoolHooks:  common.Address{},
				RmnProxy:           input.ConstructorArgs.RMNProxy,
				Router:             input.ConstructorArgs.Router,
				LockBox:            common.HexToAddress(lockBoxDeployReport.Output.Address),
			},
			Qualifier: &input.TokenSymbol,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
		}

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSel,
			TokenPoolAddress: common.HexToAddress(tpDeployReport.Output.Address),
			RateLimitAdmin:   input.RateLimitAdmin,
			RouterAddress:    input.ConstructorArgs.Router,
			FeeAdmin:         input.FeeAdmin,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", tpDeployReport.Output.Address, chain, err)
		}

		// Add lock release token pool to the authorized callers of the lock box.
		applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, evm_contract.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
			ChainSelector: input.ChainSel,
			Address:       common.HexToAddress(lockBoxDeployReport.Output.Address),
			Args: erc20_lock_box.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(tpDeployReport.Output.Address),
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
		configureReport.Output.BatchOps = append(configureReport.Output.BatchOps, []mcms_types.BatchOperation{batchOp}...)

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{
				tpDeployReport.Output,
				lockBoxDeployReport.Output,
			},
			BatchOps: configureReport.Output.BatchOps,
		}, nil
	},
)
