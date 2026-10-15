package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// ValidateLockReleaseMigrationPoolPair enforces: old is a legacy lock-release pool (incl. *AndProxy),
// old version < new version, and new is a lock-release pool at exactly v2.0.0.
func (a *EVMTokenBase) ValidateLockReleaseMigrationPoolPair(b cldf_ops.Bundle, chain evm.Chain, oldPool, newPool common.Address) (type_and_version.TypeAndVersion, type_and_version.TypeAndVersion, error) {
	oldTV, err := a.GetTypeAndVersion(b, chain, oldPool)
	if err != nil {
		return type_and_version.TypeAndVersion{}, type_and_version.TypeAndVersion{}, err
	}
	newTV, err := a.GetTypeAndVersion(b, chain, newPool)
	if err != nil {
		return type_and_version.TypeAndVersion{}, type_and_version.TypeAndVersion{}, err
	}
	if !a.IsLegacyLockReleasePoolType(string(oldTV.Type)) {
		return type_and_version.TypeAndVersion{}, type_and_version.TypeAndVersion{}, fmt.Errorf(
			"old pool %s is a %s %s, not a lock-release pool; the migration source must be a LockReleaseTokenPool, SiloedLockReleaseTokenPool, or LockReleaseTokenPoolAndProxy",
			oldPool, oldTV.Type, oldTV.Version)
	}
	if !oldTV.Version.LessThan(newTV.Version) {
		return type_and_version.TypeAndVersion{}, type_and_version.TypeAndVersion{}, fmt.Errorf(
			"old pool %s is %s %s and new pool %s is %s %s; the old pool must be strictly older than the new pool (are the refs reversed?)",
			oldPool, oldTV.Type, oldTV.Version, newPool, newTV.Type, newTV.Version)
	}
	if !a.IsLockReleasePoolType(string(newTV.Type)) || !newTV.Version.Equal(cciputils.Version_2_0_0) {
		return type_and_version.TypeAndVersion{}, type_and_version.TypeAndVersion{}, fmt.Errorf(
			"new pool %s is a %s %s, not a %s lock-release pool; the migration destination must be a v2.0.0 lockbox-based pool",
			newPool, newTV.Type, newTV.Version, cciputils.Version_2_0_0)
	}
	// Siloed liquidity can only land in a siloed pool and vice versa: the sink funds per-silo
	// lockboxes for a siloed pool and a single lockbox otherwise. Checked here, before the source
	// emits any write, so a mismatched pair cannot strand extracted liquidity in the timelock.
	oldSiloed := string(oldTV.Type) == cciputils.SiloedLockReleaseTokenPool.String()
	newSiloed := string(newTV.Type) == cciputils.SiloedLockReleaseTokenPool.String()
	if oldSiloed != newSiloed {
		return type_and_version.TypeAndVersion{}, type_and_version.TypeAndVersion{}, fmt.Errorf(
			"old pool %s is a %s but new pool %s is a %s; a siloed pool can only migrate to a siloed pool and a non-siloed pool to a non-siloed pool",
			oldPool, oldTV.Type, newPool, newTV.Type)
	}
	return oldTV, newTV, nil
}

// RequireTimelockOwnsOldPool requires the old pool to be owned by the timelock. The source's
// owner-gated writes then always land in the timelock batch (never executed directly by the
// deployer), so the rebalancer handover, withdrawal, and restore run in order and in the same
// proposal as the sink's funding. A deployer-owned pool would otherwise have the handover and
// restore executed immediately while the withdrawal waits in the proposal.
func (a *EVMTokenBase) RequireTimelockOwnsOldPool(b cldf_ops.Bundle, chain evm.Chain, oldPool, timelock common.Address) error {
	ownerReport, err := cldf_ops.ExecuteOperation(b, token_pool.Owner, chain, contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       oldPool,
	})
	if err != nil {
		return fmt.Errorf("failed to get owner of old pool %s: %w", oldPool, err)
	}
	if ownerReport.Output != timelock {
		return fmt.Errorf(
			"old pool %s is owned by %s, not the timelock %s; transfer ownership of the old pool to the timelock before migrating its liquidity",
			oldPool, ownerReport.Output, timelock)
	}
	return nil
}

// ResolveMigrationToken reads getToken() from both pools and requires them to match. Every lock-release
// pool version exposes getToken() with the same ABI, so the v1.5.0 read binding serves all versions.
func (a *EVMTokenBase) ResolveMigrationToken(b cldf_ops.Bundle, chain evm.Chain, oldPool, newPool common.Address) (common.Address, error) {
	newTokenReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       newPool,
	})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get token address from new pool %s: %w", newPool, err)
	}
	oldTokenReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       oldPool,
	})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get token address from old pool %s: %w", oldPool, err)
	}
	if oldTokenReport.Output != newTokenReport.Output {
		return common.Address{}, fmt.Errorf(
			"old pool %s manages token %s but new pool %s manages token %s; both pools must manage the same token",
			oldPool, oldTokenReport.Output, newPool, newTokenReport.Output)
	}
	return newTokenReport.Output, nil
}
