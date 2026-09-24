package tokens

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

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
type RemoveRemotePoolsPerPool struct {
	ChainSelector       uint64               `yaml:"selector" json:"selector,string"`
	Pool                datastore.AddressRef `yaml:"pool" json:"pool"`
	RemotePoolsToRemove []RemotePoolToRemove `yaml:"remotePoolsToRemove" json:"remotePoolsToRemove"`
	// AllRemotes removes every remote the pool is connected to, discovered automatically.
	AllRemotes bool `yaml:"allRemotes" json:"allRemotes"`
	// Bidirectional also removes this pool from each peer's remote list (reverse pass).
	Bidirectional bool `yaml:"bidirectional" json:"bidirectional"`
	// Deactivate unregisters the pool from the TokenAdminRegistry and performs the full teardown.
	// Must be specified alone.
	Deactivate bool `yaml:"deactivate" json:"deactivate"`
}

// RemoveRemotePools returns a changeset that removes remote pool entries from existing token
// pools. The operation version is inferred from the token pool (via the datastore), so the
// top-level changeset does not require a version field.
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
			if _, err := chain_selectors.GetSelectorFamily(pool.ChainSelector); err != nil {
				return fmt.Errorf("invalid chain selector %d: %w", pool.ChainSelector, err)
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

			seenRemotes := make(map[uint64]struct{})
			for _, remote := range pool.RemotePoolsToRemove {
				if remote.Selector == pool.ChainSelector {
					return fmt.Errorf("remote chain selector %d must not equal the pool's own chain selector", remote.Selector)
				}

				if _, err := chain_selectors.GetSelectorFamily(remote.Selector); err != nil {
					return fmt.Errorf("invalid remote chain selector %d: %w", remote.Selector, err)
				}

				if datastore_utils.IsAddressRefEmpty(remote.Remote) {
					return fmt.Errorf("remote pool entry for chain selector %d has an empty remote ref", remote.Selector)
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

// pairingKey identifies a torn-down lane pairing by the pool whose forward pass removed the
// remote (localSelector/localPool) and the remote pool that was removed (remoteSelector/
// remotePool). Dedup must be scoped to the exact pool pair: two different pools on the same chain
// pair are distinct pairings and must not dedupe each other.
type pairingKey struct {
	localSelector  uint64
	localPool      string
	remoteSelector uint64
	remotePool     string
}

// normalizeAddr round-trips an address string through the chain family's address normalizer so
// that addresses from different sources (user input, datastore, on-chain reads) compare equal.
func normalizeAddr(selector uint64, addr string) (string, error) {
	b, err := deploy.StringToBytes(selector, addr)
	if err != nil {
		return "", err
	}
	return deploy.BytesToString(selector, b)
}

// forwardPairingKey builds the dedup key for a forward-pass removal of remote from the pool at
// localPool on localSelector. The remote address is normalized so it compares equal to the
// address the reverse pass reads back from the peer.
func forwardPairingKey(localSelector uint64, localPool string, remote RemotePoolToRemove) (pairingKey, error) {
	remotePool, err := normalizeAddr(remote.Selector, remote.Remote.Address)
	if err != nil {
		return pairingKey{}, err
	}
	return pairingKey{
		localSelector:  localSelector,
		localPool:      localPool,
		remoteSelector: remote.Selector,
		remotePool:     remotePool,
	}, nil
}

func removeRemotePoolsApply() func(cldf.Environment, RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
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
		removedPairings := make(map[pairingKey]struct{})

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

			for _, remote := range remotesToRemove {
				key, err := forwardPairingKey(selector, fullPoolRef.Address, remote)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to build pairing key for remote chain selector %d on chain selector %d: %w", remote.Selector, selector, err)
				}
				removedPairings[key] = struct{}{}
			}

			// Reverse pass: for bidirectional/deactivate, remove this pool from each peer's remote list.
			if pool.Bidirectional || pool.Deactivate {
				reverseBatchOps, reverseReports, err := removeRemotePoolsReverse(e, tokenRegistry, adapter, selector, fullPoolRef, fullTokenRef, remotesToRemove, removedPairings)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to reverse-remove pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
				}
				batchOps = append(batchOps, reverseBatchOps...)
				reports = append(reports, reverseReports...)
			}

			// TAR unregister: for deactivate, unregister the pool from the TokenAdminRegistry.
			// This runs AFTER the forward + reverse cleanup so a peer is never left pointing at an
			// already-unregistered pool mid-transaction.
			if pool.Deactivate {
				unregisterBatchOps, unregisterReports, err := unregisterToken(e, tokenRegistry, family, selector, fullPoolRef, fullTokenRef)
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

	// allRemotes/deactivate discover the remotes automatically.
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

// removeRemotePoolsReverse performs the reverse pass: for each remote removed from the pool on
// selector, resolve the peer's active pool and remove this pool from that peer's remote list.
// The local pool's address as the peer stores it is discovered by reading the peer's remote pool
// list for the local chain, so the reverse pass works uniformly across chain families (the peer
// may store a contract address or a program ID). The reverse pass is idempotent: pairings already
// absent are skipped (with a warn log), and cross-pool bidirectional entries are deduped via
// removedPairings.
func removeRemotePoolsReverse(
	e cldf.Environment,
	tokenRegistry *TokenAdapterRegistry,
	localAdapter TokenAdapter,
	selector uint64,
	fullPoolRef datastore.AddressRef,
	fullTokenRef datastore.AddressRef,
	remotesToRemove []RemotePoolToRemove,
	removedPairings map[pairingKey]struct{},
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)

	for _, remote := range remotesToRemove {
		remoteSelector := remote.Selector

		remoteFamily, err := chain_selectors.GetSelectorFamily(remoteSelector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteSelector, err)
		}

		// Resolve the peer's adapter and pool ref from the peer pool address given in the forward
		// input. The peer token ref is resolved separately below (from the local pool's remote
		// config) because a Solana pool program ID is shared across mints, so the peer adapter
		// cannot derive the token from the pool address alone.
		remoteAdapter, _, _, _, err := ResolveAdapterAndRefs(e, tokenRegistry, remoteSelector, remote.Remote, datastore.AddressRef{})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve peer adapter on remote chain selector %d: %w", remoteSelector, err)
		}

		// Resolve the peer's token from the local pool's remote config. This is the token the peer
		// serves for this lane, and is required to resolve the peer's active pool below.
		remoteTokenRef, err := resolvePeerTokenRef(e, tokenRegistry, localAdapter, selector, fullPoolRef, fullTokenRef, remoteSelector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve peer token on remote chain selector %d: %w", remoteSelector, err)
		}

		// Resolve the peer's active pool. A hard error (not a silent skip) when it cannot be
		// resolved — this is distinct from the "pairing absent" idempotency case.
		manager, ok := tokenRegistry.GetTokenAdminRegistryManager(remoteFamily)
		if !ok {
			return nil, nil, fmt.Errorf("no token admin registry manager for remote chain family %s", remoteFamily)
		}
		activePool, err := manager.GetActivePool(e, remoteSelector, remoteTokenRef)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve active pool for token on remote chain selector %d: %w", remoteSelector, err)
		}
		if len(activePool) == 0 {
			return nil, nil, fmt.Errorf("token on remote chain selector %d has no active pool registered; cannot resolve peer for reverse removal", remoteSelector)
		}

		// The peer's active pool is the pool that will receive the removal. Read its remote list
		// from the same pool so the read and the write target the same contract (during a peer
		// pool upgrade the explicitly configured remote address can differ from the active pool).
		activePoolAddr, err := deploy.BytesToString(remoteSelector, activePool)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to normalize active pool address on remote chain selector %d: %w", remoteSelector, err)
		}
		activePoolRef, err := ResolveTokenPoolRef(e, tokenRegistry, remoteSelector, datastore.AddressRef{Address: activePoolAddr})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve active pool ref on remote chain selector %d: %w", remoteSelector, err)
		}

		remoteRemover, ok := remoteAdapter.(RemotePoolRemover)
		if !ok {
			return nil, nil, fmt.Errorf("adapter for remote chain selector %d does not support remote pool removal", remoteSelector)
		}

		// Discover the local pool's address as the peer stores it by reading the peer's remote
		// pool list for the local chain. This is chain-agnostic: the peer may store the local
		// pool as a contract address (EVM) or a program ID (Solana), and the migrator returns the
		// raw on-chain bytes either way. The active pool bytes are used as the pool address, which
		// matches the TokenPoolMigrator contract (the config PDA for Solana, the contract for EVM).
		peerMigrator, ok := remoteAdapter.(TokenPoolMigrator)
		if !ok {
			return nil, nil, fmt.Errorf("adapter for remote chain selector %d does not support remote pool discovery", remoteSelector)
		}
		peerTokenBytes, err := remoteAdapter.AddressRefToBytes(remoteTokenRef)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert peer token ref to bytes on remote chain selector %d: %w", remoteSelector, err)
		}
		peerRemotes, err := peerMigrator.GetRemotePools(e, remoteSelector, activePool, peerTokenBytes, selector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read peer remote pools for chain selector %d on remote chain selector %d: %w", selector, remoteSelector, err)
		}

		// Filter the peer's remote list down to the exact local pool being torn down. The peer may
		// list several local pools for this chain (e.g. during an upgrade), and we must only remove
		// the one this entry targets — never unrelated pairings.
		//
		// The peer stores the local pool in its counterpart form for the local family (EVM: the pool
		// contract; Solana: the pool config PDA), which is NOT fullPoolRef.Address (Solana normalizes
		// that to the program ID). Derive the counterpart so the comparison below is like-for-like
		// across families.
		localPoolBytes, err := localAdapter.AddressRefToBytes(fullPoolRef)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert local pool ref to bytes on chain selector %d: %w", selector, err)
		}
		localTokenBytes, err := localAdapter.AddressRefToBytes(fullTokenRef)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert local token ref to bytes on chain selector %d: %w", selector, err)
		}
		localCounterpartBytes, err := localAdapter.DeriveTokenPoolCounterpart(e, selector, localPoolBytes, localTokenBytes)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to derive local pool counterpart on chain selector %d: %w", selector, err)
		}
		localPoolAddr, err := deploy.BytesToString(selector, localCounterpartBytes)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to normalize local pool address on chain selector %d: %w", selector, err)
		}
		remotesToRemoveFromPeer := make([]RemotePoolToRemove, 0, len(peerRemotes))
		for _, peerRemote := range peerRemotes {
			peerRemoteAddr, err := deploy.BytesToString(selector, peerRemote)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to normalize peer remote pool address for chain selector %d: %w", selector, err)
			}
			if peerRemoteAddr != localPoolAddr {
				continue
			}
			remotesToRemoveFromPeer = append(remotesToRemoveFromPeer, RemotePoolToRemove{
				Selector: selector,
				Remote:   datastore.AddressRef{Address: peerRemoteAddr},
			})
		}
		if len(remotesToRemoveFromPeer) == 0 {
			e.Logger.Warnf("skipping reverse removal of pool %s on chain %d from pool on chain %d: peer does not list this pool as a remote", fullPoolRef.Address, selector, remoteSelector)
			continue
		}

		key := pairingKey{
			localSelector:  remoteSelector,
			localPool:      activePoolRef.Address,
			remoteSelector: selector,
			remotePool:     localPoolAddr,
		}
		if _, alreadyRemoved := removedPairings[key]; alreadyRemoved {
			e.Logger.Warnf("skipping reverse removal of pool %s on chain %d from pool on chain %d: pairing already removed", fullPoolRef.Address, selector, remoteSelector)
			continue
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, remoteRemover.RemoveRemotePools(), e.BlockChains, RemoveRemotePoolsSequenceInput{
			Selector:            remoteSelector,
			TokenPoolRef:        activePoolRef,
			TokenRef:            remoteTokenRef,
			RemotePoolsToRemove: remotesToRemoveFromPeer,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to remove pool from peer on remote chain selector %d: %w", remoteSelector, err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		removedPairings[key] = struct{}{}
	}

	return batchOps, reports, nil
}

// resolvePeerTokenRef resolves the token ref the peer pool serves for the lane from the local
// pool's remote config. The local pool's GetRemoteToken returns the peer's token address, which
// is then resolved to a full ref on the peer chain. This is required because a Solana pool
// program ID is shared across mints, so the peer adapter cannot derive the token from the pool
// address alone.
func resolvePeerTokenRef(
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

	poolBytes, err := localAdapter.AddressRefToBytes(fullPoolRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to convert pool ref to bytes on chain selector %d: %w", selector, err)
	}
	tokenBytes, err := localAdapter.AddressRefToBytes(fullTokenRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to convert token ref to bytes on chain selector %d: %w", selector, err)
	}

	remoteTokenBytes, err := localMigrator.GetRemoteToken(e, selector, poolBytes, tokenBytes, remoteSelector)
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
// that has moved on is never clobbered.
func unregisterToken(
	e cldf.Environment,
	tokenRegistry *TokenAdapterRegistry,
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
	targetPool, err := poolTARAddress(e, selector, fullPoolRef)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to resolve pool address for chain selector %d: %w", selector, err)
	}
	if len(activePool) == 0 || !bytes.Equal(activePool, targetPool) {
		e.Logger.Warnf("skipping TAR unregister for token on chain %d: active pool does not match pool %s", selector, fullPoolRef.Address)
		return nil, nil, nil
	}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, manager.UnregisterToken(), e.BlockChains, UnregisterTokenSequenceInput{
		Selector:          selector,
		TokenRef:          fullTokenRef,
		TokenPoolRef:      fullPoolRef,
		ExistingDataStore: e.DataStore,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unregister token on chain selector %d: %w", selector, err)
	}

	return report.Output.BatchOps, report.ExecutionReports, nil
}

// poolTARAddress returns the pool address in the form the TokenAdminRegistry stores it, as raw
// bytes. For EVM this is the pool contract address. For Solana the registry stores the pool
// program ID, but the resolved pool ref carries the pool config PDA, so the program ID is read
// back from the config PDA's owner.
func poolTARAddress(e cldf.Environment, selector uint64, fullPoolRef datastore.AddressRef) ([]byte, error) {
	family, err := chain_selectors.GetSelectorFamily(selector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", selector, err)
	}
	if family != chain_selectors.FamilySolana {
		return deploy.StringToBytes(selector, fullPoolRef.Address)
	}

	chain, ok := e.BlockChains.SolanaChains()[selector]
	if !ok {
		return nil, fmt.Errorf("solana chain with selector %d not defined", selector)
	}
	addr, err := solana.PublicKeyFromBase58(fullPoolRef.Address)
	if err != nil {
		return nil, fmt.Errorf("invalid solana pool address %q on chain %d: %w", fullPoolRef.Address, selector, err)
	}
	resp, err := chain.Client.GetAccountInfoWithOpts(e.GetContext(), addr, &rpc.GetAccountInfoOpts{Commitment: cldf_solana.SolDefaultCommitment})
	if err != nil {
		return nil, fmt.Errorf("failed to get account info for pool address %s on chain %d: %w", addr, selector, err)
	}
	if resp == nil || resp.Value == nil {
		return nil, fmt.Errorf("pool address %s not found on chain %d", addr, selector)
	}
	// An executable account is the pool program ID itself; otherwise the address is the pool
	// config PDA and its owner is the program ID.
	if resp.Value.Executable {
		return addr.Bytes(), nil
	}
	return resp.Value.Owner.Bytes(), nil
}
