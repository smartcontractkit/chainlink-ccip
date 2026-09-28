package tokens

import (
	"bytes"
	"fmt"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// removeRemotePoolsPlanner resolves every write of a RemoveRemotePools input before any of them
// runs. Writes are queued in execution order (per entry: reverse, forward, unregister) and only run
// from execute, so an input that cannot be fully resolved fails without touching any chain.
//
// removedPairings holds the pairings already queued for removal, so overlapping entries (e.g.
// A and B both bidirectional) never queue the same removal twice. Planning reads the initial
// on-chain state, as every read does under MCMS, so on-chain reads cannot dedupe for us.
type removeRemotePoolsPlanner struct {
	env             cldf.Environment
	registry        *TokenAdapterRegistry
	removedPairings map[removeRemotePoolsPairingKey]struct{}
	writes          []func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error)
}

// removeRemotePoolsPairingKey identifies a pairing by the pool that lists the remote
// (localSelector/localPool) and the listed remote pool (remoteSelector/remotePool). It is scoped
// to the exact pool pair, so two pools on the same chain pair never dedupe each other.
type removeRemotePoolsPairingKey struct {
	localSelector  uint64
	localPool      string
	remoteSelector uint64
	remotePool     string
}

// removeRemotePoolsEntry is the resolved local pool of one input entry.
type removeRemotePoolsEntry struct {
	pool       RemoveRemotePoolsPerPool
	label      string
	selector   uint64
	family     string
	adapter    TokenAdapter
	remover    RemotePoolRemover
	migrator   TokenPoolMigrator // nil when the adapter cannot read remote config; see requireMigrator
	poolRef    datastore.AddressRef
	tokenRef   datastore.AddressRef
	poolBytes  []byte
	tokenBytes []byte
}

func newRemoveRemotePoolsPlanner(e cldf.Environment, registry *TokenAdapterRegistry) *removeRemotePoolsPlanner {
	return &removeRemotePoolsPlanner{
		env:             e,
		registry:        registry,
		removedPairings: make(map[removeRemotePoolsPairingKey]struct{}),
	}
}

// planPool queues every write for one input entry.
func (p *removeRemotePoolsPlanner) planPool(pool RemoveRemotePoolsPerPool) error {
	entry, err := p.createRemoveRemotePoolsEntry(pool)
	if err != nil {
		return err
	}
	remotes, err := p.resolveRemotesToRemove(entry)
	if err != nil {
		return fmt.Errorf("failed to resolve remotes to remove for pool %s on chain selector %d: %w", entry.label, entry.selector, err)
	}
	if pool.Bidirectional || pool.Deactivate {
		// Full teardown is performed by cleaning up the neighboring pools first, then the target pool itself.
		// The target pool remote pool list is read to discover the neighbors, so if we clean up the target pool
		// first, then we won't be able to discover the neighbors to clean up if the changeset partially fails.
		if err := p.planReverse(entry, remotes); err != nil {
			return fmt.Errorf("failed to plan reverse removal of pool %s on chain selector %d: %w", entry.label, entry.selector, err)
		}
	}
	if err := p.planForward(entry, remotes); err != nil {
		return err
	}
	if pool.Deactivate {
		if err := p.planUnregister(entry); err != nil {
			return fmt.Errorf("failed to plan TAR unregister of pool %s on chain selector %d: %w", entry.label, entry.selector, err)
		}
	}
	return nil
}

// execute runs the queued writes in order and collects their batch operations and reports.
func (p *removeRemotePoolsPlanner) execute() ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	for _, write := range p.writes {
		output, writeReports, err := write()
		if err != nil {
			return nil, nil, err
		}
		batchOps = append(batchOps, output.BatchOps...)
		reports = append(reports, writeReports...)
	}
	return batchOps, reports, nil
}

func (p *removeRemotePoolsPlanner) createRemoveRemotePoolsEntry(pool RemoveRemotePoolsPerPool) (*removeRemotePoolsEntry, error) {
	entry := &removeRemotePoolsEntry{pool: pool, label: datastore_utils.SprintRef(pool.Pool), selector: pool.ChainSelector}
	poolRef := pool.Pool
	poolRef.ChainSelector = pool.ChainSelector

	var err error
	entry.adapter, entry.family, entry.poolRef, entry.tokenRef, err = ResolveAdapterAndRefs(p.env, p.registry, entry.selector, poolRef, datastore.AddressRef{})
	if err != nil {
		return nil, fmt.Errorf("failed to resolve pool %s on chain selector %d: %w", entry.label, entry.selector, err)
	}

	var ok bool
	if entry.remover, ok = entry.adapter.(RemotePoolRemover); !ok {
		return nil, fmt.Errorf("adapter for chain selector %d (family %s, version %s) does not support remote pool removal", entry.selector, entry.family, entry.poolRef.Version)
	}

	// NOTE: unlike RemotePoolRemove, the TokenPoolMigrator interface may not be used by the changeset, so
	// we should not fail hard here. Instead, the presence of the interface is checked when needed, and an
	// error is returned if the user requests an operation that requires it.
	entry.migrator, _ = entry.adapter.(TokenPoolMigrator)

	if entry.poolBytes, err = entry.adapter.AddressRefToBytes(entry.poolRef); err != nil {
		return nil, fmt.Errorf("failed to convert pool ref to bytes on chain selector %d: %w", entry.selector, err)
	}
	if entry.tokenBytes, err = entry.adapter.AddressRefToBytes(entry.tokenRef); err != nil {
		return nil, fmt.Errorf("failed to convert token ref to bytes on chain selector %d: %w", entry.selector, err)
	}

	return entry, nil
}

// requireMigrator returns the entry's remote-config reader, needed by discovery and the reverse pass.
func (entry *removeRemotePoolsEntry) requireMigrator() (TokenPoolMigrator, error) {
	if entry.migrator == nil {
		return nil, fmt.Errorf("adapter for chain selector %d (family %s) does not support remote pool discovery", entry.selector, entry.family)
	}
	return entry.migrator, nil
}

// unsupportedPeer handles a peer this tooling cannot process (reason says why): with
// skipUnsupportedPeers it logs a warning and returns nil so the caller skips the peer, otherwise it
// returns an error. The decision rests only on local lookups (loaded chains, registered readers and
// adapters), never on a failed read, so real failures are never skipped.
func (p *removeRemotePoolsPlanner) unsupportedPeer(entry *removeRemotePoolsEntry, remoteSelector uint64, reason string) error {
	if entry.pool.SkipUnsupportedPeers {
		p.env.Logger.Warnf("skipping reverse removal of pool %s on chain %d from remote chain %d: %s; the peer keeps listing this pool, clean it up on that chain with its own tooling", entry.poolRef.Address, entry.selector, remoteSelector, reason)
		return nil
	}
	return fmt.Errorf("peer on remote chain selector %d is not supported: %s (set skipUnsupportedPeers to skip it and clean it up on that chain with its own tooling)", remoteSelector, reason)
}

// remoteTokenRef resolves the peer's token from the local pool's remote config. The peer pool
// address alone is not enough: a Solana pool program is shared across mints.
func (p *removeRemotePoolsPlanner) remoteTokenRef(entry *removeRemotePoolsEntry, remoteSelector uint64) (datastore.AddressRef, error) {
	remoteTokenBytes, err := entry.migrator.GetRemoteToken(p.env, entry.selector, entry.poolBytes, entry.tokenBytes, remoteSelector)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to read remote token for remote chain selector %d: %w", remoteSelector, err)
	}
	remoteTokenAddr, err := deploy.BytesToString(remoteSelector, remoteTokenBytes)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to normalize remote token address for remote chain selector %d: %w", remoteSelector, err)
	}
	return ResolveTokenRef(p.env, p.registry, remoteSelector, datastore.AddressRef{Address: remoteTokenAddr})
}

func (p *removeRemotePoolsPlanner) queueRemoval(entry *removeRemotePoolsEntry, remover RemotePoolRemover, input RemoveRemotePoolsSequenceInput) {
	p.writes = append(p.writes, func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
		report, err := cldf_ops.ExecuteSequence(p.env.OperationsBundle, remover.RemoveRemotePools(), p.env.BlockChains, input)
		if err != nil {
			return sequences.OnChainOutput{}, nil, fmt.Errorf("failed to execute removal for pool %s on chain selector %d: failed to remove remote pools from pool %s on chain selector %d: %w",
				entry.label, entry.selector, input.TokenPoolRef.Address, input.Selector, err)
		}
		return report.Output, report.ExecutionReports, nil
	})
}

func (p *removeRemotePoolsPlanner) queueUnregister(entry *removeRemotePoolsEntry, writer TokenAdminRegistryWriter, input UnregisterTokenSequenceInput) {
	p.writes = append(p.writes, func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
		report, err := cldf_ops.ExecuteSequence(p.env.OperationsBundle, writer.UnregisterToken(), p.env.BlockChains, input)
		if err != nil {
			return sequences.OnChainOutput{}, nil, fmt.Errorf("failed to execute removal for pool %s on chain selector %d: failed to unregister token on chain selector %d: %w",
				entry.label, entry.selector, input.Selector, err)
		}
		return report.Output, report.ExecutionReports, nil
	})
}

// resolveRemotesToRemove returns the explicit remotePoolsToRemove, or for allRemotes/deactivate
// every remote pool the local pool lists.
func (p *removeRemotePoolsPlanner) resolveRemotesToRemove(entry *removeRemotePoolsEntry) ([]RemotePoolToRemove, error) {
	if !entry.pool.AllRemotes && !entry.pool.Deactivate {
		return entry.pool.RemotePoolsToRemove, nil
	}
	migrator, err := entry.requireMigrator()
	if err != nil {
		return nil, err
	}
	supportedChains, err := migrator.GetSupportedChains(p.env, entry.selector, entry.poolBytes, entry.tokenBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to discover supported chains: %w", err)
	}
	remotes := make([]RemotePoolToRemove, 0, len(supportedChains))
	for _, remoteSelector := range supportedChains {
		remotePools, err := migrator.GetRemotePools(p.env, entry.selector, entry.poolBytes, entry.tokenBytes, remoteSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get remote pools for remote chain selector %d: %w", remoteSelector, err)
		}
		for _, remotePool := range remotePools {
			remotePoolAddr, err := deploy.BytesToString(remoteSelector, remotePool)
			if err != nil {
				return nil, fmt.Errorf("failed to normalize remote pool address for remote chain selector %d: %w", remoteSelector, err)
			}
			remotes = append(remotes, RemotePoolToRemove{Selector: remoteSelector, Remote: datastore.AddressRef{Address: remotePoolAddr}})
		}
	}
	return remotes, nil
}

// planReverse queues the removal of the entry's pool from its peers. See RemoveRemotePools for why
// each peer chain can have two target pools.
func (p *removeRemotePoolsPlanner) planReverse(entry *removeRemotePoolsEntry, remotes []RemotePoolToRemove) error {
	migrator, err := entry.requireMigrator()
	if err != nil {
		return err
	}

	// Peers store the local pool in its counterpart form (EVM: the pool contract; Solana: the config
	// PDA, not the program ID that poolRef normalizes to).
	counterpart, err := entry.adapter.DeriveTokenPoolCounterpart(p.env, entry.selector, entry.poolBytes, entry.tokenBytes)
	if err != nil {
		return fmt.Errorf("failed to derive local pool counterpart on chain selector %d: %w", entry.selector, err)
	}
	localPoolAddr, err := deploy.BytesToString(entry.selector, counterpart)
	if err != nil {
		return fmt.Errorf("failed to normalize local pool address on chain selector %d: %w", entry.selector, err)
	}

	// Chain support survives a remote pool removal, so it tells a torn-down lane (still supported)
	// from a never-configured one (skipped).
	supportedChains, err := migrator.GetSupportedChains(p.env, entry.selector, entry.poolBytes, entry.tokenBytes)
	if err != nil {
		return fmt.Errorf("failed to get supported chains for pool on chain selector %d: %w", entry.selector, err)
	}
	localSupported := make(map[uint64]struct{}, len(supportedChains))
	for _, s := range supportedChains {
		localSupported[s] = struct{}{}
	}

	for _, remote := range remotes {
		targets, remoteTokenRef, err := p.reverseTargets(entry, localSupported, remote)
		if err != nil {
			return err
		}
		for _, target := range targets {
			if err := p.planReverseTarget(entry, localPoolAddr, remote.Selector, remoteTokenRef, target); err != nil {
				return err
			}
		}
	}
	return nil
}

// planForward queues the removal of remotes from the entry's own pool, skipping pairings an
// earlier entry's reverse pass already queued (a duplicate would revert under MCMS).
func (p *removeRemotePoolsPlanner) planForward(entry *removeRemotePoolsEntry, remotes []RemotePoolToRemove) error {
	forwardRemotes := make([]RemotePoolToRemove, 0, len(remotes))
	for _, remote := range remotes {
		// Normalized so it compares equal to the address a peer's reverse pass reads back.
		remotePool, err := deploy.RoundTripAddress(remote.Selector, remote.Remote.Address)
		if err != nil {
			return fmt.Errorf("failed to build pairing key for remote chain selector %d on chain selector %d: %w", remote.Selector, entry.selector, err)
		}
		key := removeRemotePoolsPairingKey{localSelector: entry.selector, localPool: entry.poolRef.Address, remoteSelector: remote.Selector, remotePool: remotePool}
		if _, already := p.removedPairings[key]; already {
			p.env.Logger.Warnf("skipping forward removal of remote %d from pool %s on chain %d: pairing already removed", remote.Selector, entry.poolRef.Address, entry.selector)
			continue
		}
		p.removedPairings[key] = struct{}{}
		forwardRemotes = append(forwardRemotes, remote)
	}
	p.queueRemoval(entry, entry.remover, RemoveRemotePoolsSequenceInput{
		Selector:            entry.selector,
		TokenPoolRef:        entry.poolRef,
		TokenRef:            entry.tokenRef,
		RemotePoolsToRemove: forwardRemotes,
	})
	return nil
}

// planUnregister queues the TAR unregister only when the TAR still points at the entry's pool; an
// empty entry, or one that has moved on to another pool, is left untouched (with a warning).
// GetActivePool returns the same byte form as AddressRefToBytes, so the two compare directly.
func (p *removeRemotePoolsPlanner) planUnregister(entry *removeRemotePoolsEntry) error {
	tar, ok := p.registry.GetTokenAdminRegistryReader(entry.family)
	if !ok {
		return fmt.Errorf("no token admin registry reader for chain family %s", entry.family)
	}
	activePool, err := tar.GetActivePool(p.env, entry.selector, entry.tokenRef)
	if err != nil {
		return fmt.Errorf("failed to read active pool for token on chain selector %d: %w", entry.selector, err)
	}
	if len(activePool) == 0 {
		p.env.Logger.Warnf("skipping TAR unregister for token on chain %d: already unregistered (TAR entry is empty)", entry.selector)
		return nil
	}
	if !bytes.Equal(activePool, entry.poolBytes) {
		activePoolAddr, err := deploy.BytesToString(entry.selector, activePool)
		if err != nil {
			activePoolAddr = fmt.Sprintf("%x", activePool)
		}
		p.env.Logger.Warnf("skipping TAR unregister for token on chain %d: TAR points at a different pool (%s), not pool %s; leaving it untouched", entry.selector, activePoolAddr, entry.poolRef.Address)
		return nil
	}
	// Unregistering is a write, so it needs a manager; a reader-only family can still deactivate a
	// pool that is no longer registered (the skips above).
	manager, ok := p.registry.GetTokenAdminRegistryManager(entry.family)
	if !ok {
		return fmt.Errorf("token admin registry for chain family %s does not support unregistering tokens", entry.family)
	}
	p.queueUnregister(entry, manager, UnregisterTokenSequenceInput{
		Selector:          entry.selector,
		TokenRef:          entry.tokenRef,
		ExistingDataStore: p.env.DataStore,
	})
	return nil
}

// reverseTargets returns the peer pools to clean for one remote: the peer's TAR-active pool, plus
// the pool the remote entry names when that is a different pool. It returns no targets for remotes
// the reverse pass skips.
func (p *removeRemotePoolsPlanner) reverseTargets(entry *removeRemotePoolsEntry, localSupported map[uint64]struct{}, remote RemotePoolToRemove) ([]datastore.AddressRef, datastore.AddressRef, error) {
	remoteSelector := remote.Selector

	// A deprecated chain has no reachable RPC; the forward pass still strips it locally.
	isDeprecated, err := chainsel.IsDeprecated(remoteSelector)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to check if remote chain selector %d is deprecated: %w", remoteSelector, err)
	}
	if isDeprecated {
		p.env.Logger.Infof("skipping reverse removal for deprecated remote chain selector %d", remoteSelector)
		return nil, datastore.AddressRef{}, nil
	}
	if _, supported := localSupported[remoteSelector]; !supported {
		p.env.Logger.Warnf("skipping reverse removal of pool %s on chain %d from pool on chain %d: local pool does not have this lane configured", entry.poolRef.Address, entry.selector, remoteSelector)
		return nil, datastore.AddressRef{}, nil
	}
	if !p.env.BlockChains.Exists(remoteSelector) {
		return nil, datastore.AddressRef{}, p.unsupportedPeer(entry, remoteSelector, "chain is not loaded in the environment")
	}

	// Get the peer's TAR reader, which is needed to read the peer's TAR-active pool.
	remoteFamily, err := chainsel.GetSelectorFamily(remoteSelector)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteSelector, err)
	}
	tar, ok := p.registry.GetTokenAdminRegistryReader(remoteFamily)
	if !ok {
		return nil, datastore.AddressRef{}, p.unsupportedPeer(entry, remoteSelector, fmt.Sprintf("no token admin registry reader for chain family %s", remoteFamily))
	}

	// Fetch the peer's TAR-active pool.
	remoteTokenRef, err := p.remoteTokenRef(entry, remoteSelector)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve peer token on remote chain selector %d: %w", remoteSelector, err)
	}
	activePoolBytes, err := tar.GetActivePool(p.env, remoteSelector, remoteTokenRef)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve active pool for token on remote chain selector %d: %w", remoteSelector, err)
	}
	if len(activePoolBytes) == 0 {
		return nil, datastore.AddressRef{}, fmt.Errorf("token on remote chain selector %d has no active pool registered; cannot resolve peer for reverse removal", remoteSelector)
	}
	activePoolString, err := deploy.BytesToString(remoteSelector, activePoolBytes)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to normalize active pool address on remote chain selector %d: %w", remoteSelector, err)
	}
	activePoolRef, err := ResolveTokenPoolRef(p.env, p.registry, remoteSelector, datastore.AddressRef{Address: activePoolString})
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve active pool ref on remote chain selector %d: %w", remoteSelector, err)
	}
	activePoolAddr, err := deploy.RoundTripAddress(remoteSelector, activePoolRef.Address)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to normalize active pool address on remote chain selector %d: %w", remoteSelector, err)
	}

	// If we encounter a remote pool with the zero address, then the only neighbor we can clean is the peer's TAR-active pool
	isZero, err := deploy.IsZeroAddress(remoteSelector, remote.Remote.Address)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to decode peer pool address %s on remote chain selector %d: %w", remote.Remote.Address, remoteSelector, err)
	}
	if isZero {
		p.env.Logger.Warnf("remote entry for chain %d on pool %s (chain %d) is the zero address; only the peer's active pool is targeted by the reverse pass", remoteSelector, entry.poolRef.Address, entry.selector)
		return []datastore.AddressRef{activePoolRef}, remoteTokenRef, nil
	}

	// If the remote pool isn't the zero address AND it's not the active pool, then we have two neighboring pools to clean
	namedPoolRef, err := ResolveTokenPoolRef(p.env, p.registry, remoteSelector, remote.Remote)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf(
			"failed to resolve peer pool ref %s on remote chain selector %d (if this remote entry is not a real pool, "+
				"remove it from this pool with a lane-only removal, i.e. remotePoolsToRemove without bidirectional, then re-run): %w",
			datastore_utils.SprintRef(remote.Remote), remoteSelector, err,
		)
	}
	namedPoolAddr, err := deploy.RoundTripAddress(remoteSelector, namedPoolRef.Address)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to normalize peer pool address on remote chain selector %d: %w", remoteSelector, err)
	}
	targets := []datastore.AddressRef{activePoolRef}
	if namedPoolAddr != activePoolAddr {
		targets = append(targets, namedPoolRef)
	}

	return targets, remoteTokenRef, nil
}

// planReverseTarget queues the removal of the local pool (localPoolAddr) from one peer pool, if
// that peer pool lists it.
func (p *removeRemotePoolsPlanner) planReverseTarget(entry *removeRemotePoolsEntry, localPoolAddr string, remoteSelector uint64, remoteTokenRef, target datastore.AddressRef) error {
	key := removeRemotePoolsPairingKey{localSelector: remoteSelector, localPool: target.Address, remoteSelector: entry.selector, remotePool: localPoolAddr}
	if _, already := p.removedPairings[key]; already {
		p.env.Logger.Warnf("skipping reverse removal of pool %s on chain %d from peer pool %s on chain %d: pairing already removed", entry.poolRef.Address, entry.selector, target.Address, remoteSelector)
		return nil
	}

	// Resolve the adapter from the target pool's own version: a peer's pools can differ in version.
	remoteAdapter, _, err := ResolveAdapter(p.registry, remoteSelector, target.Version)
	if err != nil {
		return p.unsupportedPeer(entry, remoteSelector, fmt.Sprintf("no adapter for peer pool %s: %v", target.Address, err))
	}
	remoteMigrator, isMigrator := remoteAdapter.(TokenPoolMigrator)
	remoteRemover, isRemover := remoteAdapter.(RemotePoolRemover)
	if !isMigrator || !isRemover {
		return p.unsupportedPeer(entry, remoteSelector, fmt.Sprintf("adapter for peer pool %s does not support remote pool discovery and removal", target.Address))
	}
	remoteTokenBytes, err := remoteAdapter.AddressRefToBytes(remoteTokenRef)
	if err != nil {
		return fmt.Errorf("failed to convert peer token ref to bytes on remote chain selector %d: %w", remoteSelector, err)
	}
	targetPoolBytes, err := remoteAdapter.AddressRefToBytes(target)
	if err != nil {
		return fmt.Errorf("failed to convert peer pool ref to bytes on remote chain selector %d: %w", remoteSelector, err)
	}
	listed, err := remoteMigrator.GetRemotePools(p.env, remoteSelector, targetPoolBytes, remoteTokenBytes, entry.selector)
	if err != nil {
		return fmt.Errorf("failed to read remote pools of peer pool %s for chain selector %d on remote chain selector %d: %w", target.Address, entry.selector, remoteSelector, err)
	}

	// Match the peer's list against the local pool only; the peer may list other pools on this chain.
	toRemove := make([]RemotePoolToRemove, 0, 1)
	for _, stored := range listed {
		storedAddr, err := deploy.BytesToString(entry.selector, stored)
		if err != nil {
			return fmt.Errorf("failed to normalize peer remote pool address for chain selector %d: %w", entry.selector, err)
		}
		if storedAddr == localPoolAddr {
			toRemove = append(toRemove, RemotePoolToRemove{Selector: entry.selector, Remote: datastore.AddressRef{Address: storedAddr}})
		}
	}
	if len(toRemove) == 0 {
		p.env.Logger.Warnf("skipping reverse removal of pool %s on chain %d from peer pool %s on chain %d: peer pool does not list this pool as a remote", entry.poolRef.Address, entry.selector, target.Address, remoteSelector)
		return nil
	}

	p.removedPairings[key] = struct{}{}
	p.queueRemoval(entry, remoteRemover, RemoveRemotePoolsSequenceInput{
		Selector:            remoteSelector,
		TokenPoolRef:        target,
		TokenRef:            remoteTokenRef,
		RemotePoolsToRemove: toRemove,
	})

	return nil
}
