package tokens

import (
	"bytes"
	"errors"
	"fmt"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// RemoveRemotePoolsInput is the input for the RemoveRemotePools changeset. It removes remote
// pool entries from existing token pools.
type RemoveRemotePoolsInput struct {
	Pools []RemoveRemotePoolsPerPool `yaml:"pools" json:"pools"`
	MCMS  mcms.Input                 `yaml:"mcms,omitempty" json:"mcms"`
}

// RemoveRemotePoolsPerPool groups remote pool removals for a single token pool. The pool is
// referenced by an AddressRef so operators can identify it by qualifier, by address, or by any
// other unique combination of ref fields.
//
// The modes are mutually exclusive in the following way:
//   - remotePoolsToRemove alone: lane-only removal (forward pass only).
//   - allRemotes: forward pass only, removing every remote the pool is connected to (discovered
//     automatically via the pool's remote-config reader).
//   - bidirectional (with allRemotes or remotePoolsToRemove): also performs a reverse pass,
//     removing this pool from each peer's remote list.
//   - deactivate: must be specified alone (no other mode flags and no explicit
//     remotePoolsToRemove). Unregisters the pool from the TokenAdminRegistry (setPool to zero)
//     AND performs the allRemotes forward cleanup AND the bidirectional reverse cleanup,
//     completing the full teardown. The remotes are always discovered automatically.
//
// The reverse pass removes this pool from up to two pools per peer chain: the peer's TAR-active
// pool and the peer pool this pool is paired with (they differ once the peer has been upgraded,
// since remote-pool lists are append-only across upgrades). Retired pools on this chain are not
// swept automatically: to retire several pools (e.g. the old pools left behind by upgrades), list
// each one as its own entry.
//
// Partial failures and re-runs: the reverse pass runs before the forward pass, so for
// allRemotes/deactivate (whose remotes are discovered from this pool's own remote list) a re-run
// after a partial failure rediscovers the same peers and completes the teardown. Caveat: when
// ownership is mixed, execution order follows ownership rather than code order. Operations the
// deployer key owns land immediately, while MCMS-owned ones land only when the proposal executes.
// If this pool is deployer-owned and a peer pool is MCMS-owned, the forward removal can land while
// the peer's removal is still pending; if that proposal is then dropped, a discovery-based re-run
// no longer sees the peer. To recover, re-run with bidirectional and an explicit
// remotePoolsToRemove listing the peer pools, which does not depend on this pool's remote list.
type RemoveRemotePoolsPerPool struct {
	ChainSelector       uint64               `yaml:"selector" json:"selector,string"`
	Pool                datastore.AddressRef `yaml:"pool" json:"pool"`
	RemotePoolsToRemove []RemotePoolToRemove `yaml:"remotePoolsToRemove" json:"remotePoolsToRemove"`
	Bidirectional       bool                 `yaml:"bidirectional" json:"bidirectional"`
	AllRemotes          bool                 `yaml:"allRemotes" json:"allRemotes"`
	Deactivate          bool                 `yaml:"deactivate" json:"deactivate"`
}

// RemoveRemotePools returns a changeset that removes remote pool entries from existing token
// pools. The operation version is inferred from the token pool (via the datastore), so the
// top-level changeset does not require a version field.
//
// This changeset tidies up pool topology; it is not a safe decommission and does not protect
// message flow. Use it on pools that have already been migrated away from (the TAR points at a
// newer pool), once in-flight messages have settled on the new pools: removing a remote pool
// makes the peer reject messages still in flight from it. Deactivating a pool that is still the
// active pool also leaves peers able to send to this chain (removing remote pools does not remove
// chain support), and those transfers fail on arrival.
//
// Example: retiring an old pool after upgrades. Each list below is a pool's remote pools, and TAR
// shows each chain's registered (active) pool. Start with a mesh of A, B and C:
//
//	TAR   -> A_old, B_old, C_old
//	A_old = [B_old, C_old]
//	B_old = [A_old, C_old]
//	C_old = [A_old, B_old]
//
// Upgrade A (A_old -> A_new). The new pool copies the old pool's remote list, the new pool is
// added to each peer's active pool (remote lists are append-only, so the old pool stays listed
// for in-flight protection), and the TAR switches to the new pool:
//
//	TAR   -> A_new, B_old, C_old
//	A_new = [B_old, C_old]
//
//	A_old = [B_old, C_old]
//	B_old = [A_new, A_old, C_old]
//	C_old = [A_new, A_old, B_old]
//
// Upgrade B (B_old -> B_new). B_new copies B_old's list, which includes A_old. B_new is added to
// the peers' active pools (A_new and C_old), so A_old is never told about B_new:
//
//	TAR   -> A_new, B_new, C_old
//	A_new = [B_old, B_new, C_old]
//	B_new = [A_new, A_old, C_old]
//
//	A_old = [B_old, C_old]
//	B_old = [A_new, A_old, C_old]
//	C_old = [A_new, A_old, B_old, B_new]
//
// deactivate(A_old). A_old is now listed on two pools of chain B: B_new (B's active pool) and
// B_old (the pool A_old is paired with). The reverse pass targets both, since cleaning only one
// would leave A_old listed on the other. On chain C (never upgraded) both targets are C_old, so it
// is processed once. The TAR unregister is skipped because A_old is not the active pool:
//
//	TAR   -> A_new, B_new, C_old
//	A_new = [B_old, B_new, C_old]
//	B_new = [A_new, C_old]
//
//	A_old = []
//	B_old = [A_new, C_old]
//	C_old = [A_new, B_old, B_new]
//
// B_old stays wired until it is retired too. Adding a deactivate(B_old) entry (in the same input
// or a later run) leaves only the active pools connected:
//
//	TAR   -> A_new, B_new, C_old
//	A_new = [B_new, C_old]
//	B_new = [A_new, C_old]
//
//	A_old = []
//	B_old = []
//	C_old = [A_new, B_new]
func RemoveRemotePools() cldf.ChangeSetV2[RemoveRemotePoolsInput] {
	return cldf.CreateChangeSet(removeRemotePoolsApply(), removeRemotePoolsVerify())
}

func removeRemotePoolsVerify() func(cldf.Environment, RemoveRemotePoolsInput) error {
	return func(_ cldf.Environment, cfg RemoveRemotePoolsInput) error {
		if len(cfg.Pools) == 0 {
			return errors.New("input must contain at least one pool entry")
		}

		type poolKey struct {
			selector uint64
			ref      string
		}

		seenPools := make(map[poolKey]struct{})
		for _, pool := range cfg.Pools {
			// The pool's own chain must be live: every operation for this entry (forward removal and
			// the TAR unregister) executes on it. A deprecated (decommissioned) chain has no reachable
			// RPC, so reject it here rather than failing deep in apply. Deprecated remotes are allowed
			// — the forward pass still strips them from the local pool, and the reverse pass skips them.
			isDeprecated, err := chainsel.IsDeprecated(pool.ChainSelector)
			if err != nil {
				return fmt.Errorf("failed to check if chain selector %d is deprecated: %w", pool.ChainSelector, err)
			}
			if isDeprecated {
				return fmt.Errorf("pool entry on chain selector %d: chain is deprecated/decommissioned", pool.ChainSelector)
			}
			if datastore_utils.IsAddressRefEmpty(pool.Pool) {
				return fmt.Errorf("pool entry on chain selector %d has an empty pool ref", pool.ChainSelector)
			}

			// Mode mutual-exclusivity rules.
			if pool.Deactivate {
				// deactivate performs the full teardown (allRemotes forward cleanup + bidirectional
				// reverse pass + TAR unregister) and must be specified alone: no other mode flags and
				// no explicit remotePoolsToRemove. The remotes are always discovered automatically.
				if pool.AllRemotes || pool.Bidirectional || len(pool.RemotePoolsToRemove) > 0 {
					return fmt.Errorf("pool entry %s on chain selector %d: deactivate must be specified alone", datastore_utils.SprintRef(pool.Pool), pool.ChainSelector)
				}
			} else {
				if pool.AllRemotes && len(pool.RemotePoolsToRemove) > 0 {
					return fmt.Errorf("pool entry %s on chain selector %d: allRemotes and remotePoolsToRemove are mutually exclusive", datastore_utils.SprintRef(pool.Pool), pool.ChainSelector)
				}
				if pool.Bidirectional && !pool.AllRemotes && len(pool.RemotePoolsToRemove) == 0 {
					return fmt.Errorf("pool entry %s on chain selector %d: bidirectional requires allRemotes or remotePoolsToRemove", datastore_utils.SprintRef(pool.Pool), pool.ChainSelector)
				}
				if !pool.AllRemotes && len(pool.RemotePoolsToRemove) == 0 {
					return fmt.Errorf("pool entry %s on chain selector %d has no remote pools to remove", datastore_utils.SprintRef(pool.Pool), pool.ChainSelector)
				}
			}

			// Block duplicate remote selectors for this pool entry, and validate each remote's selector and ref.
			seenRemotes := make(map[uint64]struct{})
			for _, remote := range pool.RemotePoolsToRemove {
				if remote.Selector == pool.ChainSelector {
					return fmt.Errorf("remote chain selector %d must not equal the pool's own chain selector", remote.Selector)
				}
				if _, err := chainsel.GetSelectorFamily(remote.Selector); err != nil {
					return fmt.Errorf("invalid remote chain selector %d: %w", remote.Selector, err)
				}
				if datastore_utils.IsAddressRefEmpty(remote.Remote) {
					return fmt.Errorf("remote pool entry for chain selector %d has an empty remote ref", remote.Selector)
				}
				if remote.Remote.Address == "" {
					return fmt.Errorf("remote pool entry for chain selector %d must set remote.address", remote.Selector)
				}
				if _, dup := seenRemotes[remote.Selector]; dup {
					return fmt.Errorf("duplicate remote chain selector %d for pool on chain selector %d", remote.Selector, pool.ChainSelector)
				}
				seenRemotes[remote.Selector] = struct{}{}
			}

			key := poolKey{selector: pool.ChainSelector, ref: datastore_utils.SprintRef(pool.Pool)}
			if _, dup := seenPools[key]; dup {
				return fmt.Errorf("duplicate pool entry for chain selector %d and ref %s", pool.ChainSelector, datastore_utils.SprintRef(pool.Pool))
			}

			seenPools[key] = struct{}{}
		}
		return nil
	}
}

// removeRemotePoolsPairingKey identifies a torn-down lane pairing by the pool whose forward pass
// removed the remote (localSelector/localPool) and the remote pool that was removed
// (remoteSelector/remotePool). Dedup must be scoped to the exact pool pair: two different pools on
// the same chain pair are distinct pairings and must not dedupe each other.
type removeRemotePoolsPairingKey struct {
	localSelector  uint64
	localPool      string
	remoteSelector uint64
	remotePool     string
}

func removeRemotePoolsApply() func(cldf.Environment, RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
	// forwardPairingKey builds the dedup key for a forward-pass removal of remote from the pool at
	// localPool on localSelector. The remote address is normalized so it compares equal to the
	// address the reverse pass reads back from the peer.
	forwardPairingKey := func(localSelector uint64, localPool string, remote RemotePoolToRemove) (removeRemotePoolsPairingKey, error) {
		remotePool, err := deploy.RoundTripAddress(remote.Selector, remote.Remote.Address)
		if err != nil {
			return removeRemotePoolsPairingKey{}, err
		}
		return removeRemotePoolsPairingKey{
			localSelector:  localSelector,
			localPool:      localPool,
			remoteSelector: remote.Selector,
			remotePool:     remotePool,
		}, nil
	}

	return func(e cldf.Environment, cfg RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		tokenRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		// removedPairings tracks lane pairings that have already been torn down, keyed by the exact
		// pool pair (local pool + remote pool), so a bidirectional removal never double-removes a
		// pairing when A and B are both specified for a bidirectional removal of each other. The
		// key includes the pool identity so two distinct pools on the same chain pair do not
		// dedupe each other.
		removedPairings := make(map[removeRemotePoolsPairingKey]struct{})
		for _, pool := range cfg.Pools {
			selector := pool.ChainSelector
			poolRef := pool.Pool
			poolRef.ChainSelector = selector

			adapter, family, fullPoolRef, fullTokenRef, err := ResolveAdapterAndRefs(e, tokenRegistry, selector, poolRef, datastore.AddressRef{})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
			}
			remover, ok := adapter.(RemotePoolRemover)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"adapter for chain selector %d (family %s, version %s) does not support remote pool removal",
					selector, family, fullPoolRef.Version,
				)
			}

			// Resolve the set of remotes to remove in the forward pass.
			remotesToRemove, err := resolveRemotesToRemove(e, adapter, family, selector, fullPoolRef, fullTokenRef, pool)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve remotes to remove for pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
			}

			// Reverse pass first: for bidirectional/deactivate, remove this pool from each peer's remote
			// list BEFORE the forward pass empties this pool's own list. For allRemotes/deactivate the
			// remotes are discovered from this pool's list, so that list doubles as the record of
			// remaining work: if a run fails partway (e.g. deployer-key operations that land
			// immediately), a re-run rediscovers the same peers and finishes the teardown, skipping
			// peers already cleaned. If the forward pass ran first and the reverse pass then failed,
			// a re-run would discover nothing and leave the peer side configured.
			if pool.Bidirectional || pool.Deactivate {
				reverseBatchOps, reverseReports, err := removeRemotePoolsReverse(e, tokenRegistry, adapter, selector, fullPoolRef, fullTokenRef, remotesToRemove, removedPairings)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to reverse-remove pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
				}
				batchOps = append(batchOps, reverseBatchOps...)
				reports = append(reports, reverseReports...)
			}

			// Skip pairings a prior entry's reverse pass already tore down, so we never emit a
			// duplicate removal. Under MCMS the batched reads all see the initial state, so the
			// sequence's own read-check can't dedupe across entries and the duplicate would revert
			// on-chain when the proposal executes.
			forwardRemotes := make([]RemotePoolToRemove, 0, len(remotesToRemove))
			for _, remote := range remotesToRemove {
				key, err := forwardPairingKey(selector, fullPoolRef.Address, remote)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to build pairing key for remote chain selector %d on chain selector %d: %w", remote.Selector, selector, err)
				}
				if _, already := removedPairings[key]; already {
					e.Logger.Warnf("skipping forward removal of remote %d from pool %s on chain %d: pairing already removed", remote.Selector, fullPoolRef.Address, selector)
					continue
				}
				forwardRemotes = append(forwardRemotes, remote)
			}

			// Forward pass: remove the specified remotes from this pool.
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, remover.RemoveRemotePools(), e.BlockChains, RemoveRemotePoolsSequenceInput{
				Selector:            selector,
				TokenPoolRef:        fullPoolRef,
				TokenRef:            fullTokenRef,
				RemotePoolsToRemove: forwardRemotes,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to remove remote pools from pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)

			// Record the forward-pass removals so a later entry's reverse pass never double-removes a pairing.
			for _, remote := range remotesToRemove {
				key, err := forwardPairingKey(selector, fullPoolRef.Address, remote)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to build pairing key for remote chain selector %d on chain selector %d: %w", remote.Selector, selector, err)
				}
				removedPairings[key] = struct{}{}
			}

			// TAR unregister: for deactivate, unregister the pool from the TokenAdminRegistry. It is
			// emitted after the reverse and forward cleanup to keep the topology tidy. This order
			// does not protect message flow, and across chains it only holds when operations land
			// immediately: under MCMS each chain's proposal executes independently.
			if pool.Deactivate {
				unregisterBatchOps, unregisterReports, err := unregisterToken(e, tokenRegistry, adapter, family, selector, fullPoolRef, fullTokenRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to unregister pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
				}
				batchOps = append(batchOps, unregisterBatchOps...)
				reports = append(reports, unregisterReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

// resolveRemotesToRemove determines the set of remotes to remove in the forward pass for a pool.
// For allRemotes/deactivate, it discovers the pool's supported chains via the remote-config
// reader (TokenPoolMigrator). For explicit remotePoolsToRemove, it returns them as-is.
func resolveRemotesToRemove(
	e cldf.Environment,
	adapter TokenAdapter,
	family string,
	selector uint64,
	fullPoolRef datastore.AddressRef,
	fullTokenRef datastore.AddressRef,
	pool RemoveRemotePoolsPerPool,
) ([]RemotePoolToRemove, error) {
	if !pool.AllRemotes && !pool.Deactivate {
		return pool.RemotePoolsToRemove, nil
	}

	migrator, ok := adapter.(TokenPoolMigrator)
	if !ok {
		return nil, fmt.Errorf("adapter for chain selector %d (family %s) does not support remote pool discovery", selector, family)
	}
	poolBytes, err := adapter.AddressRefToBytes(fullPoolRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert pool ref to bytes on chain selector %d: %w", selector, err)
	}
	tokenBytes, err := adapter.AddressRefToBytes(fullTokenRef)
	if err != nil {
		return nil, fmt.Errorf("failed to convert token ref to bytes on chain selector %d: %w", selector, err)
	}
	supportedChains, err := migrator.GetSupportedChains(e, selector, poolBytes, tokenBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to discover supported chains for pool on chain selector %d: %w", selector, err)
	}

	remotes := make([]RemotePoolToRemove, 0, len(supportedChains))
	for _, remoteSelector := range supportedChains {
		remotePools, err := migrator.GetRemotePools(e, selector, poolBytes, tokenBytes, remoteSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get remote pools for remote chain selector %d on chain selector %d: %w", remoteSelector, selector, err)
		}
		for _, remotePool := range remotePools {
			remotePoolAddr, err := deploy.BytesToString(remoteSelector, remotePool)
			if err != nil {
				return nil, fmt.Errorf("failed to normalize remote pool address for remote chain selector %d: %w", remoteSelector, err)
			}
			remotes = append(remotes, RemotePoolToRemove{
				Selector: remoteSelector,
				Remote:   datastore.AddressRef{Address: remotePoolAddr},
			})
		}
	}

	return remotes, nil
}

// removeRemotePoolsReverse performs the reverse pass: for each remote of the pool on selector,
// remove this pool from the peer's TAR-active pool and from the peer pool named by the remote
// entry (deduped when they are the same pool). Each target pool's remote list is read and
// filtered to this pool's counterpart address, so the reverse pass works uniformly across chain
// families and never removes unrelated pairings. The reverse pass is idempotent: pairings already
// absent are skipped (with a warn log), and cross-pool bidirectional entries are deduped via
// removedPairings. See RemoveRemotePools for a worked example of why both targets are needed.
func removeRemotePoolsReverse(
	e cldf.Environment,
	tokenRegistry *TokenAdapterRegistry,
	localAdapter TokenAdapter,
	selector uint64,
	fullPoolRef datastore.AddressRef,
	fullTokenRef datastore.AddressRef,
	remotesToRemove []RemotePoolToRemove,
	removedPairings map[removeRemotePoolsPairingKey]struct{},
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)

	// Precompute local-pool facts used for every remote:
	//   - poolToRemoveAddr: the local pool's address as a peer stores it (its counterpart form: EVM the
	//     pool contract, Solana the config PDA), NOT fullPoolRef.Address (Solana normalizes that to
	//     the program ID). Used to filter the peer's remote list to the exact local pool below.
	//   - localSupported: the chains the local pool actually has configured. This is chain-config
	//     based (not the remote-pool list), so it survives a forward pool removal and still lets us
	//     distinguish a never-configured lane (the documented warn-and-skip idempotency case).
	localMigrator, ok := localAdapter.(TokenPoolMigrator)
	if !ok {
		return nil, nil, fmt.Errorf("adapter for chain selector %d does not support remote pool discovery", selector)
	}
	localTokenBytes, err := localAdapter.AddressRefToBytes(fullTokenRef)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert local token ref to bytes on chain selector %d: %w", selector, err)
	}
	localPoolBytes, err := localAdapter.AddressRefToBytes(fullPoolRef)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert local pool ref to bytes on chain selector %d: %w", selector, err)
	}
	poolToRemoveBytes, err := localAdapter.DeriveTokenPoolCounterpart(e, selector, localPoolBytes, localTokenBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to derive local pool counterpart on chain selector %d: %w", selector, err)
	}
	poolToRemoveAddr, err := deploy.BytesToString(selector, poolToRemoveBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to normalize local pool address on chain selector %d: %w", selector, err)
	}
	supportedChains, err := localMigrator.GetSupportedChains(e, selector, localPoolBytes, localTokenBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get supported chains for pool on chain selector %d: %w", selector, err)
	}
	localSupported := make(map[uint64]struct{}, len(supportedChains))
	for _, s := range supportedChains {
		localSupported[s] = struct{}{}
	}

	for _, remote := range remotesToRemove {
		remoteSelector := remote.Selector

		// Skip deprecated (decommissioned) remote chains: their RPC is unreachable, so the reverse
		// pass (which reads the peer's active pool and remote list, then removes on the peer) cannot
		// run against them. The forward pass has already stripped the deprecated remote from the
		// local pool, which is the reachable, meaningful cleanup. A lookup failure is a hard error
		// rather than a silent skip.
		isDeprecated, err := chainsel.IsDeprecated(remoteSelector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to check if remote chain selector %d is deprecated: %w", remoteSelector, err)
		}
		if isDeprecated {
			e.Logger.Infof("skipping reverse removal for deprecated remote chain selector %d", remoteSelector)
			continue
		}

		// Skip lanes the local pool no longer has configured — the documented warn-and-skip
		// idempotent case (a never-configured pairing, or a re-run after the pairing was torn down).
		// Removing a remote pool leaves the local chain config in place, so a configured lane still
		// reports as supported here even after its forward removal, while a never-configured lane
		// does not. Gating here also avoids a hard error from the peer-token resolution below (which
		// reads the local pool's remote config and would error for an absent lane).
		if _, supported := localSupported[remoteSelector]; !supported {
			e.Logger.Warnf("skipping reverse removal of pool %s on chain %d from pool on chain %d: local pool does not have this lane configured", fullPoolRef.Address, selector, remoteSelector)
			continue
		}

		// Get the remote chain's family and token admin registry manager, which is used to resolve
		// the peer's active pool for this lane.
		remoteFamily, err := chainsel.GetSelectorFamily(remoteSelector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteSelector, err)
		}
		manager, ok := tokenRegistry.GetTokenAdminRegistryManager(remoteFamily)
		if !ok {
			return nil, nil, fmt.Errorf("no token admin registry manager for remote chain family %s", remoteFamily)
		}

		// Target 1: remove the pool named by the remote entry (remote.Remote)
		configuredPoolRef, err := ResolveTokenPoolRef(e, tokenRegistry, remoteSelector, remote.Remote)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve peer pool ref %s on remote chain selector %d: %w", datastore_utils.SprintRef(remote.Remote), remoteSelector, err)
		}
		configuredAddr, err := deploy.RoundTripAddress(remoteSelector, configuredPoolRef.Address)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to normalize peer pool address on remote chain selector %d: %w", remoteSelector, err)
		}

		// Target 2: query the TAR for this remote chain and get the active remote pool
		remoteTokenRef, err := getRemoteTokenRef(e, tokenRegistry, localAdapter, selector, fullPoolRef, fullTokenRef, remoteSelector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve peer token on remote chain selector %d: %w", remoteSelector, err)
		}
		remoteActivePoolBytes, err := manager.GetActivePool(e, remoteSelector, remoteTokenRef)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve active pool for token on remote chain selector %d: %w", remoteSelector, err)
		}
		if len(remoteActivePoolBytes) == 0 {
			// A peer with no active pool is a hard error (not a silent skip); this is distinct from
			// the "pairing absent" idempotency case.
			return nil, nil, fmt.Errorf("token on remote chain selector %d has no active pool registered; cannot resolve peer for reverse removal", remoteSelector)
		}
		remoteActivePoolString, err := deploy.BytesToString(remoteSelector, remoteActivePoolBytes)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to normalize active pool address on remote chain selector %d: %w", remoteSelector, err)
		}
		remoteActivePoolRef, err := ResolveTokenPoolRef(e, tokenRegistry, remoteSelector, datastore.AddressRef{Address: remoteActivePoolString})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve active pool ref on remote chain selector %d: %w", remoteSelector, err)
		}
		remoteActivePoolAddr, err := deploy.RoundTripAddress(remoteSelector, remoteActivePoolRef.Address)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to normalize active pool address on remote chain selector %d: %w", remoteSelector, err)
		}

		// Targets: the active pool always; the paired pool too when it is a different pool (deduped so
		// a peer that was never upgraded is processed once).
		targets := []datastore.AddressRef{remoteActivePoolRef}
		if configuredAddr != remoteActivePoolAddr {
			// v1.5.0 pools have no removeRemotePool, so a retired v1.5.0 peer pool cannot be cleaned
			// here. Keep going (the active pool is still cleaned) and say so. For the active pool the
			// remover's own error is surfaced, as before.
			if configuredPoolRef.Version != nil && configuredPoolRef.Version.Equal(utils.Version_1_5_0) {
				e.Logger.Warnf("skipping v1.5.0 peer pool %s on chain %d: v1.5.0 pools cannot remove individual remote pools, so pool %s on chain %d may remain listed on it", configuredPoolRef.Address, remoteSelector, fullPoolRef.Address, selector)
			} else {
				targets = append(targets, configuredPoolRef)
			}
		}

		// Now iterate over each target and perform the reverse removal
		for _, targetRef := range targets {
			key := removeRemotePoolsPairingKey{
				localSelector:  remoteSelector,
				localPool:      targetRef.Address,
				remoteSelector: selector,
				remotePool:     poolToRemoveAddr,
			}
			if _, alreadyRemoved := removedPairings[key]; alreadyRemoved {
				e.Logger.Warnf("skipping reverse removal of pool %s on chain %d from peer pool %s on chain %d: pairing already removed", fullPoolRef.Address, selector, targetRef.Address, remoteSelector)
				continue
			}

			// Resolve the adapter from the target pool itself so the read, the write, and the ABI
			// all match the contract being modified (the peer's pools can be on different versions).
			remoteAdapter, _, err := ResolveAdapter(tokenRegistry, remoteSelector, targetRef.Version)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to resolve adapter for peer pool %s on remote chain selector %d: %w", targetRef.Address, remoteSelector, err)
			}
			remoteMigrator, ok := remoteAdapter.(TokenPoolMigrator)
			if !ok {
				return nil, nil, fmt.Errorf("adapter for remote chain selector %d does not support remote pool discovery", remoteSelector)
			}
			remoteRemover, ok := remoteAdapter.(RemotePoolRemover)
			if !ok {
				return nil, nil, fmt.Errorf("adapter for remote chain selector %d does not support remote pool removal", remoteSelector)
			}

			// Read the target pool's remote list for the local chain (the same pool that will receive
			// the removal). The migrator returns raw on-chain bytes, so this works uniformly across
			// families (EVM contract address, Solana program ID / config PDA).
			remoteTokenBytes, err := remoteAdapter.AddressRefToBytes(remoteTokenRef)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert peer token ref to bytes on remote chain selector %d: %w", remoteSelector, err)
			}
			targetPoolBytes, err := remoteAdapter.AddressRefToBytes(targetRef)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to convert peer pool ref to bytes on remote chain selector %d: %w", remoteSelector, err)
			}
			targetPoolRemotes, err := remoteMigrator.GetRemotePools(e, remoteSelector, targetPoolBytes, remoteTokenBytes, selector)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to read remote pools of peer pool %s for chain selector %d on remote chain selector %d: %w", targetRef.Address, selector, remoteSelector, err)
			}

			// Filter the target pool's remote list down to the exact local pool being torn down
			// (poolToRemoveAddr, precomputed above). The peer may list several local pools for this
			// chain (e.g. after an upgrade), and we must only remove the one this entry targets —
			// never unrelated pairings.
			remotesToRemoveFromTargetPool := make([]RemotePoolToRemove, 0, len(targetPoolRemotes))
			for _, remotePoolOnTargetPool := range targetPoolRemotes {
				remotePoolAddressOnTargetPool, err := deploy.BytesToString(selector, remotePoolOnTargetPool)
				if err != nil {
					return nil, nil, fmt.Errorf("failed to normalize peer remote pool address for chain selector %d: %w", selector, err)
				}
				if remotePoolAddressOnTargetPool != poolToRemoveAddr {
					continue
				}
				remotesToRemoveFromTargetPool = append(remotesToRemoveFromTargetPool, RemotePoolToRemove{
					Selector: selector,
					Remote:   datastore.AddressRef{Address: remotePoolAddressOnTargetPool},
				})
			}
			if len(remotesToRemoveFromTargetPool) == 0 {
				e.Logger.Warnf("skipping reverse removal of pool %s on chain %d from peer pool %s on chain %d: peer pool does not list this pool as a remote", fullPoolRef.Address, selector, targetRef.Address, remoteSelector)
				continue
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, remoteRemover.RemoveRemotePools(), e.BlockChains, RemoveRemotePoolsSequenceInput{
				Selector:            remoteSelector,
				TokenPoolRef:        targetRef,
				TokenRef:            remoteTokenRef,
				RemotePoolsToRemove: remotesToRemoveFromTargetPool,
			})
			if err != nil {
				return nil, nil, fmt.Errorf("failed to remove pool from peer pool %s on remote chain selector %d: %w", targetRef.Address, remoteSelector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
			removedPairings[key] = struct{}{}
		}
	}

	return batchOps, reports, nil
}

// getRemoteTokenRef resolves the token ref the peer pool serves for the lane from the local
// pool's remote config. The local pool's GetRemoteToken returns the peer's token address, which
// is then resolved to a full ref on the peer chain. This is required because a Solana pool
// program ID is shared across mints, so the peer adapter cannot derive the token from the pool
// address alone.
func getRemoteTokenRef(
	e cldf.Environment,
	tokenRegistry *TokenAdapterRegistry,
	localAdapter TokenAdapter,
	selector uint64,
	fullPoolRef datastore.AddressRef,
	fullTokenRef datastore.AddressRef,
	remoteSelector uint64,
) (datastore.AddressRef, error) {
	localMigrator, ok := localAdapter.(TokenPoolMigrator)
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("adapter for chain selector %d does not support remote pool discovery", selector)
	}

	localTokenBytes, err := localAdapter.AddressRefToBytes(fullTokenRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to convert token ref to bytes on chain selector %d: %w", selector, err)
	}
	localPoolBytes, err := localAdapter.AddressRefToBytes(fullPoolRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to convert pool ref to bytes on chain selector %d: %w", selector, err)
	}

	remoteTokenBytes, err := localMigrator.GetRemoteToken(e, selector, localPoolBytes, localTokenBytes, remoteSelector)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to read remote token for remote chain selector %d: %w", remoteSelector, err)
	}
	remoteTokenAddr, err := deploy.BytesToString(remoteSelector, remoteTokenBytes)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to normalize remote token address for remote chain selector %d: %w", remoteSelector, err)
	}

	return ResolveTokenRef(e, tokenRegistry, remoteSelector, datastore.AddressRef{Address: remoteTokenAddr})
}

// unregisterToken unregisters the pool from the TokenAdminRegistry via the family's
// TokenAdminRegistryManager. It first reads the token's current active pool and only emits the
// unregister when that pool is the pool being removed; if the entry now points at a different
// pool or is already empty, the unregister is skipped (with a warn log) so a live registration
// that has moved on is never clobbered. The pool being removed is compared in the byte form
// returned by the family adapter's AddressRefToBytes, which GetActivePool is required to match.
func unregisterToken(
	e cldf.Environment,
	tokenRegistry *TokenAdapterRegistry,
	adapter TokenAdapter,
	family string,
	selector uint64,
	fullPoolRef datastore.AddressRef,
	fullTokenRef datastore.AddressRef,
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	manager, ok := tokenRegistry.GetTokenAdminRegistryManager(family)
	if !ok {
		return nil, nil, fmt.Errorf("no token admin registry manager for chain family %s", family)
	}
	activePool, err := manager.GetActivePool(e, selector, fullTokenRef)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read active pool for token on chain selector %d: %w", selector, err)
	}
	targetPool, err := adapter.AddressRefToBytes(fullPoolRef)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to convert pool ref to bytes on chain selector %d: %w", selector, err)
	}
	if len(activePool) == 0 {
		e.Logger.Warnf("skipping TAR unregister for token on chain %d: already unregistered (TAR entry is empty)", selector)
		return nil, nil, nil
	}
	if !bytes.Equal(activePool, targetPool) {
		activePoolAddr, err := deploy.BytesToString(selector, activePool)
		if err != nil {
			activePoolAddr = fmt.Sprintf("%x", activePool)
		}
		e.Logger.Warnf("skipping TAR unregister for token on chain %d: TAR points at a different pool (%s), not pool %s; leaving it untouched", selector, activePoolAddr, fullPoolRef.Address)
		return nil, nil, nil
	}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, manager.UnregisterToken(), e.BlockChains, UnregisterTokenSequenceInput{
		Selector:          selector,
		TokenRef:          fullTokenRef,
		ExistingDataStore: e.DataStore,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unregister token on chain selector %d: %w", selector, err)
	}

	return report.Output.BatchOps, report.ExecutionReports, nil
}
