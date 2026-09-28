package tokens

import (
	"errors"
	"fmt"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
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

// RemoveRemotePoolsPerPool groups remote pool removals for a single token pool. It supports
// several modes, which are mutually exclusive in the following way:
//
//   - remotePoolsToRemove alone: lane-only removal (forward pass only).
//   - allRemotes: forward pass only, removing every remote the pool is connected to (discovered
//     automatically via the pool's remote-config reader).
//   - bidirectional (with allRemotes or remotePoolsToRemove): also performs a reverse pass,
//     removing this pool from each peer's remote list.
//   - deactivate: must be specified alone (no other mode flags and no explicit
//     remotePoolsToRemove). Unregisters the pool from the TokenAdminRegistry (setPool to zero)
//     AND performs the allRemotes forward cleanup AND the bidirectional reverse cleanup,
//     completing the full teardown. The remotes are always discovered automatically. Unregistering
//     requires the chain family to register a TokenAdminRegistryManager; a family registered only
//     as a reader fails at planning when an unregister is needed.
//
// The reverse pass removes this pool from up to two pools per peer chain: the peer's TAR-active
// pool and the peer pool this pool is paired with (they differ once the peer has been upgraded,
// since remote-pool lists are append-only across upgrades). Retired pools on this chain are not
// swept automatically: to retire several pools (e.g. the old pools left behind by upgrades), list
// each one as its own entry.
//
// Every entry is resolved before anything is written: if any pool, remote entry, or peer cannot be
// resolved, the changeset fails without touching any chain. A remote entry holding the zero
// address names no peer pool, so the reverse pass skips it as a target (the forward pass still
// removes it). Any other remote entry that is not a real pool (an EOA, a typo, a misencoded
// address) fails resolution; remove it with a lane-only removal (remotePoolsToRemove without
// bidirectional, which never resolves the peer) and re-run.
//
// Partial failures and re-runs: a write can still fail partway (e.g. a reverted transaction), so
// the reverse pass runs before the forward pass. For allRemotes/deactivate (whose remotes are
// discovered from this pool's own remote list) a re-run after a partial failure rediscovers the
// same peers and completes the teardown. Caveat: when ownership is mixed, execution order follows
// ownership rather than code order. Operations the deployer key owns land immediately, while MCMS-
// owned ones land only when the proposal executes. If this pool is deployer-owned and a peer pool
// is MCMS-owned, the forward removal can land while the peer's removal is still pending; if that
// proposal is then dropped, a discovery-based re-run no longer sees the peer. To recover, re-run
// with bidirectional and an explicit remotePoolsToRemove listing the peer pools, which does not
// depend on this pool's remote list.
//
// The SkipUnsupportedPeers flag makes the reverse pass skip (with a warning) peers this tooling
// cannot process instead of failing: a peer chain not loaded in the environment, a peer family with
// no registered TAR reader, or a peer pool whose adapter does not support remote pool discovery and
// removal. A skipped peer keeps listing this pool; clean it up on that chain with its own
// tooling. Other failures (RPC errors, unresolvable remote entries, a peer with no active pool)
// still fail. Requires bidirectional or deactivate.
type RemoveRemotePoolsPerPool struct {
	ChainSelector        uint64               `yaml:"selector" json:"selector,string"`
	Pool                 datastore.AddressRef `yaml:"pool" json:"pool"`
	RemotePoolsToRemove  []RemotePoolToRemove `yaml:"remotePoolsToRemove" json:"remotePoolsToRemove"`
	Bidirectional        bool                 `yaml:"bidirectional" json:"bidirectional"`
	AllRemotes           bool                 `yaml:"allRemotes" json:"allRemotes"`
	Deactivate           bool                 `yaml:"deactivate" json:"deactivate"`
	SkipUnsupportedPeers bool                 `yaml:"skipUnsupportedPeers" json:"skipUnsupportedPeers"`
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
			// The pool's own chain must be reachable. Deprecated remotes are allowed: the forward pass
			// still strips them locally and the reverse pass skips them.
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

			// Mode rules (see RemoveRemotePoolsPerPool).
			if pool.Deactivate {
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
				if pool.SkipUnsupportedPeers && !pool.Bidirectional {
					return fmt.Errorf("pool entry %s on chain selector %d: skipUnsupportedPeers requires bidirectional or deactivate", datastore_utils.SprintRef(pool.Pool), pool.ChainSelector)
				}
				if !pool.AllRemotes && len(pool.RemotePoolsToRemove) == 0 {
					return fmt.Errorf("pool entry %s on chain selector %d has no remote pools to remove", datastore_utils.SprintRef(pool.Pool), pool.ChainSelector)
				}
			}

			// Several remotes may share a selector (a pool can list several pools per chain), but the
			// same remote pool twice would queue a duplicate removal.
			type remotePoolKey struct {
				selector uint64
				address  string
			}
			seenRemotes := make(map[remotePoolKey]struct{})
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
				remoteAddr, err := deploy.RoundTripAddress(remote.Selector, remote.Remote.Address)
				if err != nil {
					return fmt.Errorf("invalid remote pool address %q for chain selector %d: %w", remote.Remote.Address, remote.Selector, err)
				}
				key := remotePoolKey{selector: remote.Selector, address: remoteAddr}
				if _, dup := seenRemotes[key]; dup {
					return fmt.Errorf("duplicate remote pool %s for remote chain selector %d on pool on chain selector %d", remoteAddr, remote.Selector, pool.ChainSelector)
				}
				seenRemotes[key] = struct{}{}
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

func removeRemotePoolsApply() func(cldf.Environment, RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
		planner := newRemoveRemotePoolsPlanner(e, GetTokenAdapterRegistry())
		for _, pool := range cfg.Pools {
			if err := planner.planPool(pool); err != nil {
				return cldf.ChangesetOutput{}, err
			}
		}
		batchOps, reports, err := planner.execute()
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		return changesets.NewOutputBuilder(e, changesets.GetRegistry()).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
