package tokens

import (
	"bytes"
	"fmt"
	"slices"

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
// runs. Writes are queued during planning and only run from execute, so an input that cannot be
// fully resolved fails without touching any chain.
//
// removedPairings holds the pairings already queued for removal, so overlapping entries (e.g.
// A and B both bidirectional) never queue the same removal twice. Planning reads the initial
// on-chain state, as every read does under MCMS, so on-chain reads cannot dedupe for us.
//
// Removals are grouped per target pool, so each pool gets a single removal call. Some families
// rewrite a remote chain's whole pool list on every removal (Solana), and under MCMS every
// rewrite is built from the initial state, so two calls on one pool would undo each other.
// execute runs the grouped removals in phases: pools that are only peers of an entry first, then
// pools that are some entry's own pool (in the order they were first queued), then the
// chain-config wipes, then the TAR unregisters. An entry queues its peers before its own pool, so
// its peers are cleaned before its own pool (whose list records the remaining work). The exception:
// if an entry's own pool was already queued as an earlier entry's peer, it can run before one of
// its peers that is itself some entry's own pool.
type removeRemotePoolsPlanner struct {
	env             cldf.Environment
	registry        *TokenAdapterRegistry
	removedPairings map[removeRemotePoolsPairingKey]struct{}
	removalsByPool  map[removeRemotePoolsRemovalKey]*removeRemotePoolsRemoval
	wipesByPool     map[removeRemotePoolsRemovalKey]*removeRemotePoolsWipe
	unregisters     []func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error)
	removals        []*removeRemotePoolsRemoval       // in first-queued order
	wipes           []*removeRemotePoolsWipe          // in first-queued order
	wipeCandidates  []*removeRemotePoolsWipeCandidate // in registration order; resolved by resolveWipes
}

// removeRemotePoolsRemovalKey identifies a target pool. The token is part of the key because a
// Solana pool program is shared across mints, and each mint has its own remote configs.
type removeRemotePoolsRemovalKey struct {
	selector uint64
	pool     string
	token    string
}

// removeRemotePoolsRemoval is the single removal call for one target pool, combining every
// removal queued for it.
type removeRemotePoolsRemoval struct {
	remover RemotePoolRemover
	input   RemoveRemotePoolsSequenceInput
	ownPool bool // the target is an entry's own pool (a forward removal was queued for it)
	wiped   bool // every queued removal falls on a wiped lane, so the call is dropped
}

// removeRemotePoolsWipe is the single chain-config removal call for one target pool, combining
// every selector queued for it.
type removeRemotePoolsWipe struct {
	remover RemoteChainConfigRemover
	input   RemoveRemoteChainConfigsSequenceInput
}

// removeRemotePoolsWipeCandidate is a pool whose lanes may be wiped. Planning only records the lanes
// and the remote pools each lists (read at plan time, normalized); resolveWipes decides which lanes
// to wipe once every entry is planned, against the merged removals queued for the pool. Deciding per
// entry would miss a lane emptied by removals from several entries (two entries resolving to the
// same pool, or another entry's reverse pass).
type removeRemotePoolsWipeCandidate struct {
	key     removeRemotePoolsRemovalKey
	remover RemoteChainConfigRemover
	input   RemoveRemoteChainConfigsSequenceInput // RemoteChainSelectors is filled by resolveWipes
	lanes   []removeRemotePoolsLane
	// ownPool marks an entry's own pool, whose lanes are every supported chain. A peer candidate only
	// records the one lane it was created for, so only an own-pool candidate can tell that a removal
	// targets an unsupported selector.
	ownPool bool
}

// removeRemotePoolsLane is one supported remote chain of a wipe candidate and the normalized remote
// pools the candidate lists for it.
type removeRemotePoolsLane struct {
	selector uint64
	listed   []string
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
	pool         RemoveRemotePoolsPerPool
	label        string
	selector     uint64
	family       string
	adapter      TokenAdapter
	remover      RemotePoolRemover
	chainRemover RemoteChainConfigRemover // nil when the adapter cannot remove chain configs
	migrator     TokenPoolMigrator        // nil when the adapter cannot read remote config; see requireMigrator
	poolRef      datastore.AddressRef
	tokenRef     datastore.AddressRef
	poolBytes    []byte
	tokenBytes   []byte
}

func newRemoveRemotePoolsPlanner(e cldf.Environment, registry *TokenAdapterRegistry) *removeRemotePoolsPlanner {
	return &removeRemotePoolsPlanner{
		env:             e,
		registry:        registry,
		removedPairings: make(map[removeRemotePoolsPairingKey]struct{}),
		removalsByPool:  make(map[removeRemotePoolsRemovalKey]*removeRemotePoolsRemoval),
		wipesByPool:     make(map[removeRemotePoolsRemovalKey]*removeRemotePoolsWipe),
	}
}

// newRemovalKey builds the normalized key identifying a target pool (see removeRemotePoolsRemovalKey).
func newRemovalKey(selector uint64, poolRef, tokenRef datastore.AddressRef) (removeRemotePoolsRemovalKey, error) {
	pool, err := deploy.RoundTripAddress(selector, poolRef.Address)
	if err != nil {
		return removeRemotePoolsRemovalKey{}, fmt.Errorf("failed to normalize pool address %s on chain selector %d: %w", poolRef.Address, selector, err)
	}
	token, err := deploy.RoundTripAddress(selector, tokenRef.Address)
	if err != nil {
		return removeRemotePoolsRemovalKey{}, fmt.Errorf("failed to normalize token address %s on chain selector %d: %w", tokenRef.Address, selector, err)
	}
	return removeRemotePoolsRemovalKey{selector: selector, pool: pool, token: token}, nil
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

	if pool.WipeChainConfigs && (entry.chainRemover == nil || !entry.chainRemover.SupportsRemoteChainConfigRemoval()) {
		return fmt.Errorf("adapter for chain selector %d (family %s, version %s) does not support removing remote chain configs", entry.selector, entry.family, entry.poolRef.Version)
	}

	if pool.Bidirectional || pool.Deactivate {
		// Full teardown is performed by cleaning up the neighboring pools first, then the target pool itself.
		// The target pool remote pool list is read to discover the neighbors, so if we clean up the target pool
		// first, then we won't be able to discover the neighbors to clean up if the changeset partially fails.
		if err := p.planReverse(entry, remotes); err != nil {
			return fmt.Errorf("failed to plan reverse removal of pool %s on chain selector %d: %w", entry.label, entry.selector, err)
		}
	}
	// Forward removals are queued even for lanes that end up wiped: resolveWipes decides the wipes
	// against the merged removals, then drops the removals the wipes make redundant.
	if len(remotes) > 0 {
		if err := p.planForward(entry, remotes); err != nil {
			return err
		}
	}
	if pool.WipeChainConfigs {
		if err := p.addOwnWipeCandidate(entry); err != nil {
			return fmt.Errorf("failed to resolve remote chains to wipe for pool %s on chain selector %d: %w", entry.label, entry.selector, err)
		}
	}
	if pool.Deactivate {
		if err := p.planUnregister(entry); err != nil {
			return fmt.Errorf("failed to plan TAR unregister of pool %s on chain selector %d: %w", entry.label, entry.selector, err)
		}
	}

	return nil
}

// execute runs the queued writes in phases (see removeRemotePoolsPlanner) and collects their batch
// operations and reports.
func (p *removeRemotePoolsPlanner) execute() ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	if err := p.resolveWipes(); err != nil {
		return nil, nil, err
	}

	// Drop queued per-pool removals for any selector whose chain config is being wiped on the same
	// pool: the wipe deletes the whole config, so the removal would be redundant.
	for key, wipe := range p.wipesByPool {
		removal, ok := p.removalsByPool[key]
		if !ok {
			continue
		}
		kept := removal.input.RemotePoolsToRemove[:0]
		for _, remote := range removal.input.RemotePoolsToRemove {
			if !slices.Contains(wipe.input.RemoteChainSelectors, remote.Selector) {
				kept = append(kept, remote)
			}
		}
		removal.wiped = len(kept) == 0 && len(removal.input.RemotePoolsToRemove) > 0
		removal.input.RemotePoolsToRemove = kept
	}

	executors := make([]func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error), 0, len(p.removals)+len(p.wipes)+len(p.unregisters))
	for _, removal := range p.removals {
		if !removal.ownPool && !removal.wiped {
			executors = append(executors, p.getRemoveRemotePoolExecutor(removal))
		}
	}
	for _, removal := range p.removals {
		if removal.ownPool && !removal.wiped {
			executors = append(executors, p.getRemoveRemotePoolExecutor(removal))
		}
	}
	for _, wipe := range p.wipes {
		executors = append(executors, p.getRemoveRemoteChainConfigsExecutor(wipe))
	}
	executors = append(executors, p.unregisters...)

	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)
	for _, execute := range executors {
		output, summary, err := execute()
		if err != nil {
			return nil, nil, err
		}
		batchOps = append(batchOps, output.BatchOps...)
		reports = append(reports, summary...)
	}

	return batchOps, reports, nil
}

func (p *removeRemotePoolsPlanner) getRemoveRemoteChainConfigsExecutor(wipe *removeRemotePoolsWipe) func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
	return func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
		report, err := cldf_ops.ExecuteSequence(p.env.OperationsBundle, wipe.remover.RemoveRemoteChainConfigs(), p.env.BlockChains, wipe.input)
		if err != nil {
			return sequences.OnChainOutput{}, nil, fmt.Errorf("failed to remove remote chain configs from pool %s on chain selector %d: %w", wipe.input.TokenPoolRef.Address, wipe.input.Selector, err)
		}
		return report.Output, report.ExecutionReports, nil
	}
}

func (p *removeRemotePoolsPlanner) getRemoveRemotePoolExecutor(removal *removeRemotePoolsRemoval) func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
	return func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
		report, err := cldf_ops.ExecuteSequence(p.env.OperationsBundle, removal.remover.RemoveRemotePools(), p.env.BlockChains, removal.input)
		if err != nil {
			return sequences.OnChainOutput{}, nil, fmt.Errorf("failed to remove remote pools from pool %s on chain selector %d: %w", removal.input.TokenPoolRef.Address, removal.input.Selector, err)
		}
		return report.Output, report.ExecutionReports, nil
	}
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

	// NOTE: unlike RemotePoolRemover, the TokenPoolMigrator and RemoteChainConfigRemover interfaces
	// may not be used by the changeset so we shouldn't fail hard here. Instead, the presence of the
	// interface is checked when needed. An error is returned if the user requests an operation that
	// requires it.
	entry.chainRemover, _ = entry.adapter.(RemoteChainConfigRemover)
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

// handleUnsupportedPeer handles a peer this tooling cannot process (reason says why): with
// skipUnsupportedPeers it logs a warning and returns nil so the caller skips the peer, otherwise it
// returns an error. The decision rests only on local lookups (loaded chains, registered readers and
// adapters), never on a failed read, so real failures are never skipped.
func (p *removeRemotePoolsPlanner) handleUnsupportedPeer(entry *removeRemotePoolsEntry, remoteSelector uint64, reason string) error {
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

// queueRemoval adds input's removals to the single removal call for its target pool. ownPool marks
// a forward removal, i.e. the target is the queuing entry's own pool.
func (p *removeRemotePoolsPlanner) queueRemoval(remover RemotePoolRemover, input RemoveRemotePoolsSequenceInput, ownPool bool) error {
	key, err := newRemovalKey(input.Selector, input.TokenPoolRef, input.TokenRef)
	if err != nil {
		return err
	}

	removal, exists := p.removalsByPool[key]
	if !exists {
		removal = &removeRemotePoolsRemoval{remover: remover, input: input}
		p.removals = append(p.removals, removal)
		p.removalsByPool[key] = removal
	} else {
		removal.input.RemotePoolsToRemove = append(removal.input.RemotePoolsToRemove, input.RemotePoolsToRemove...)
	}

	removal.ownPool = removal.ownPool || ownPool

	return nil
}

// queueWipe adds input's selectors to the single chain-config removal call for its target pool.
// Duplicate selectors are collapsed so two entries naming the same pool never queue a selector
// twice (applyChainUpdates reverts on an already-removed chain).
func (p *removeRemotePoolsPlanner) queueWipe(remover RemoteChainConfigRemover, input RemoveRemoteChainConfigsSequenceInput) error {
	key, err := newRemovalKey(input.Selector, input.TokenPoolRef, input.TokenRef)
	if err != nil {
		return err
	}

	wipe, exists := p.wipesByPool[key]
	if !exists {
		wipe = &removeRemotePoolsWipe{remover: remover, input: input}
		p.wipes = append(p.wipes, wipe)
		p.wipesByPool[key] = wipe
		return nil
	}
	for _, selector := range input.RemoteChainSelectors {
		if !slices.Contains(wipe.input.RemoteChainSelectors, selector) {
			wipe.input.RemoteChainSelectors = append(wipe.input.RemoteChainSelectors, selector)
		}
	}

	return nil
}

func (p *removeRemotePoolsPlanner) queueUnregister(entry *removeRemotePoolsEntry, writer TokenAdminRegistryWriter, input UnregisterTokenSequenceInput) {
	p.unregisters = append(p.unregisters, func() (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
		report, err := cldf_ops.ExecuteSequence(p.env.OperationsBundle, writer.UnregisterToken(), p.env.BlockChains, input)
		if err != nil {
			return sequences.OnChainOutput{}, nil, fmt.Errorf("failed to unregister token for pool %s on chain selector %d: %w",
				entry.label, entry.selector, err)
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

// addOwnWipeCandidate records the entry's own pool as a wipe candidate: every lane the pool
// currently supports, with the remote pools it lists. Only currently-supported selectors are
// recorded, so a re-run never re-emits a wipe for a chain that is already gone (applyChainUpdates
// reverts NonExistentChain); resolveWipes warns about removals that target any other selector.
func (p *removeRemotePoolsPlanner) addOwnWipeCandidate(entry *removeRemotePoolsEntry) error {
	migrator, err := entry.requireMigrator()
	if err != nil {
		return err
	}
	supportedChains, err := migrator.GetSupportedChains(p.env, entry.selector, entry.poolBytes, entry.tokenBytes)
	if err != nil {
		return fmt.Errorf("failed to discover supported chains: %w", err)
	}

	lanes := make([]removeRemotePoolsLane, 0, len(supportedChains))
	for _, remoteSelector := range supportedChains {
		stored, err := migrator.GetRemotePools(p.env, entry.selector, entry.poolBytes, entry.tokenBytes, remoteSelector)
		if err != nil {
			return fmt.Errorf("failed to get remote pools for remote chain selector %d: %w", remoteSelector, err)
		}
		lane := removeRemotePoolsLane{selector: remoteSelector, listed: make([]string, 0, len(stored))}
		for _, remotePool := range stored {
			normalized, err := deploy.BytesToString(remoteSelector, remotePool)
			if err != nil {
				return fmt.Errorf("failed to normalize remote pool address for remote chain selector %d: %w", remoteSelector, err)
			}
			lane.listed = append(lane.listed, normalized)
		}
		lanes = append(lanes, lane)
	}

	return p.addWipeCandidate(entry.chainRemover, RemoveRemoteChainConfigsSequenceInput{
		Selector:     entry.selector,
		TokenPoolRef: entry.poolRef,
		TokenRef:     entry.tokenRef,
	}, lanes, true)
}

func (p *removeRemotePoolsPlanner) addWipeCandidate(remover RemoteChainConfigRemover, input RemoveRemoteChainConfigsSequenceInput, lanes []removeRemotePoolsLane, ownPool bool) error {
	key, err := newRemovalKey(input.Selector, input.TokenPoolRef, input.TokenRef)
	if err != nil {
		return err
	}
	p.wipeCandidates = append(p.wipeCandidates, &removeRemotePoolsWipeCandidate{key: key, remover: remover, input: input, lanes: lanes, ownPool: ownPool})
	return nil
}

// resolveWipes queues a wipe for every candidate lane that the merged removals queued for the
// candidate's pool leave empty. Candidates resolve in registration order, so an entry's peers are
// still wiped before its own pool. For an own-pool candidate, a merged removal whose selector the
// pool does not report as supported is warned about once per pool and selector, however many
// entries or remote pools name it.
func (p *removeRemotePoolsPlanner) resolveWipes() error {
	warned := make(map[removeRemotePoolsRemovalKey]map[uint64]struct{})
	for _, candidate := range p.wipeCandidates {
		removing := make(map[uint64]map[string]struct{})
		if removal, ok := p.removalsByPool[candidate.key]; ok {
			for _, remote := range removal.input.RemotePoolsToRemove {
				if candidate.ownPool && !slices.ContainsFunc(candidate.lanes, func(lane removeRemotePoolsLane) bool { return lane.selector == remote.Selector }) {
					if warned[candidate.key] == nil {
						warned[candidate.key] = make(map[uint64]struct{})
					}
					if _, done := warned[candidate.key][remote.Selector]; !done {
						warned[candidate.key][remote.Selector] = struct{}{}
						p.env.Logger.Warnf("not wiping the chain config for remote chain %d on pool %s (chain %d): the pool does not report it as supported (already removed, or not discoverable on this chain family)", remote.Selector, candidate.input.TokenPoolRef.Address, candidate.input.Selector)
					}
				}
				normalized, err := deploy.RoundTripAddress(remote.Selector, remote.Remote.Address)
				if err != nil {
					return fmt.Errorf("failed to normalize remote pool address %s for remote chain selector %d: %w", remote.Remote.Address, remote.Selector, err)
				}
				if removing[remote.Selector] == nil {
					removing[remote.Selector] = make(map[string]struct{})
				}
				removing[remote.Selector][normalized] = struct{}{}
			}
		}

		var wipes []uint64
		for _, lane := range candidate.lanes {
			emptied := true
			for _, listed := range lane.listed {
				if _, beingRemoved := removing[lane.selector][listed]; !beingRemoved {
					emptied = false
					break
				}
			}
			if emptied {
				wipes = append(wipes, lane.selector)
			}
		}
		if len(wipes) == 0 {
			continue
		}

		input := candidate.input
		input.RemoteChainSelectors = wipes
		if err := p.queueWipe(candidate.remover, input); err != nil {
			return err
		}
	}
	return nil
}

// planReverse queues the removal of the entry's pool from its peers. See RemoveRemotePools for why
// each peer chain can have up to two target pools.
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
	return p.queueRemoval(entry.remover, RemoveRemotePoolsSequenceInput{
		Selector:            entry.selector,
		TokenPoolRef:        entry.poolRef,
		TokenRef:            entry.tokenRef,
		RemotePoolsToRemove: forwardRemotes,
	}, true)
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
			// if BytesToString fails, we still want to log the raw bytes so the user can inspect them
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

// reverseTargets returns the peer pools to clean for one remote: the peer's TAR-active pool (when
// it has one), plus the pool the remote entry names when that is a different pool. It returns no
// targets for remotes the reverse pass skips.
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
		return nil, datastore.AddressRef{}, p.handleUnsupportedPeer(entry, remoteSelector, "chain is not loaded in the environment")
	}

	// Get the peer's TAR reader, which is needed to read the peer's TAR-active pool.
	remoteFamily, err := chainsel.GetSelectorFamily(remoteSelector)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteSelector, err)
	}
	tar, ok := p.registry.GetTokenAdminRegistryReader(remoteFamily)
	if !ok {
		return nil, datastore.AddressRef{}, p.handleUnsupportedPeer(entry, remoteSelector, fmt.Sprintf("no token admin registry reader for chain family %s", remoteFamily))
	}
	remoteTokenRef, err := p.remoteTokenRef(entry, remoteSelector)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve peer token on remote chain selector %d: %w", remoteSelector, err)
	}
	activePoolBytes, err := tar.GetActivePool(p.env, remoteSelector, remoteTokenRef)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve active pool for token on remote chain selector %d: %w", remoteSelector, err)
	}

	// Clean up target 1: the pool that the remote entry names
	var namedPoolRef datastore.AddressRef
	var remotePoolAddr string
	isZero, err := deploy.IsZeroAddress(remoteSelector, remote.Remote.Address)
	if err != nil {
		return nil, datastore.AddressRef{}, fmt.Errorf("failed to decode peer pool address %s on remote chain selector %d: %w", remote.Remote.Address, remoteSelector, err)
	}
	if !isZero {
		namedPoolRef, err = ResolveTokenPoolRef(p.env, p.registry, remoteSelector, remote.Remote)
		if err != nil {
			return nil, datastore.AddressRef{}, fmt.Errorf(
				"failed to resolve peer pool ref %s on remote chain selector %d (if this remote entry is not a real pool, "+
					"remove it from this pool with a lane-only removal, i.e. remotePoolsToRemove without bidirectional, then re-run): %w",
				datastore_utils.SprintRef(remote.Remote), remoteSelector, err,
			)
		}
		remotePoolAddr, err = deploy.RoundTripAddress(remoteSelector, namedPoolRef.Address)
		if err != nil {
			return nil, datastore.AddressRef{}, fmt.Errorf("failed to normalize peer pool address on remote chain selector %d: %w", remoteSelector, err)
		}
	}

	// Clean up target 2: the peer's TAR-active pool, which may be absent (the peer may not have registered one)
	var activePoolRef datastore.AddressRef
	var activePoolAddr string
	if len(activePoolBytes) != 0 {
		activePoolString, err := deploy.BytesToString(remoteSelector, activePoolBytes)
		if err != nil {
			return nil, datastore.AddressRef{}, fmt.Errorf("failed to normalize active pool address on remote chain selector %d: %w", remoteSelector, err)
		}
		activePoolRef, err = ResolveTokenPoolRef(p.env, p.registry, remoteSelector, datastore.AddressRef{Address: activePoolString})
		if err != nil {
			return nil, datastore.AddressRef{}, fmt.Errorf("failed to resolve active pool ref on remote chain selector %d: %w", remoteSelector, err)
		}
		activePoolAddr, err = deploy.RoundTripAddress(remoteSelector, activePoolRef.Address)
		if err != nil {
			return nil, datastore.AddressRef{}, fmt.Errorf("failed to normalize active pool address on remote chain selector %d: %w", remoteSelector, err)
		}
	}

	// Select the targets: both when they differ, one when only one exists or they coincide, none
	// when neither exists (the remote entry is the zero address and the peer has no active pool).
	// The cases that clean only one side or nothing warn once, so a zero-address entry or a
	// missing active pool is reported exactly once rather than at each step that observes it.
	switch {
	case remotePoolAddr != "" && activePoolAddr != "" && remotePoolAddr != activePoolAddr:
		return []datastore.AddressRef{activePoolRef, namedPoolRef}, remoteTokenRef, nil
	case remotePoolAddr != "" && activePoolAddr != "" && remotePoolAddr == activePoolAddr:
		return []datastore.AddressRef{namedPoolRef}, remoteTokenRef, nil
	case remotePoolAddr != "" && activePoolAddr == "":
		p.env.Logger.Warnf("token on remote chain selector %d has no active pool registered; cleaning only the pool named by the remote entry", remoteSelector)
		return []datastore.AddressRef{namedPoolRef}, remoteTokenRef, nil
	case remotePoolAddr == "" && activePoolAddr != "":
		p.env.Logger.Warnf("remote entry for chain %d on pool %s (chain %d) is the zero address; only the peer's active pool is targeted by the reverse pass", remoteSelector, entry.poolRef.Address, entry.selector)
		return []datastore.AddressRef{activePoolRef}, remoteTokenRef, nil
	default:
		p.env.Logger.Warnf("remote entry for chain %d on pool %s (chain %d) is the zero address and the peer has no active pool; nothing to clean", remoteSelector, entry.poolRef.Address, entry.selector)
		return nil, remoteTokenRef, nil
	}
}

// planReverseTarget queues the removal of the local pool (localPoolAddr) from one peer pool, if
// that peer pool lists it.
func (p *removeRemotePoolsPlanner) planReverseTarget(entry *removeRemotePoolsEntry, localPoolAddr string, remoteSelector uint64, remoteTokenRef, target datastore.AddressRef) error {
	key := removeRemotePoolsPairingKey{localSelector: remoteSelector, localPool: target.Address, remoteSelector: entry.selector, remotePool: localPoolAddr}
	// A pairing an earlier entry already queued is not removed twice. With wipeChainConfigs the peer
	// is still read so its lane is registered as a wipe candidate (and validated); returning here would
	// make whether the peer's emptied lane is wiped depend on entry order.
	_, alreadyQueued := p.removedPairings[key]
	if alreadyQueued && !entry.pool.WipeChainConfigs {
		p.env.Logger.Warnf("skipping reverse removal of pool %s on chain %d from peer pool %s on chain %d: pairing already removed", entry.poolRef.Address, entry.selector, target.Address, remoteSelector)
		return nil
	}

	// Resolve the adapter from the target pool's own version: a peer's pools can differ in version.
	remoteAdapter, _, err := ResolveAdapter(p.registry, remoteSelector, target.Version)
	if err != nil {
		return p.handleUnsupportedPeer(entry, remoteSelector, fmt.Sprintf("no adapter for peer pool %s: %v", target.Address, err))
	}
	remoteMigrator, isMigrator := remoteAdapter.(TokenPoolMigrator)
	remoteRemover, isRemover := remoteAdapter.(RemotePoolRemover)
	if !isMigrator || !isRemover {
		return p.handleUnsupportedPeer(entry, remoteSelector, fmt.Sprintf("adapter for peer pool %s does not support remote pool discovery and removal", target.Address))
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
	// One match is enough even if the peer stores the pool in several encodings: the remover removes
	// every stored encoding of the address, and a duplicate entry would queue duplicate removals.
	// allListLocalPool additionally tracks whether every entry is the local pool, i.e. removing it
	// alone leaves the peer with no remote pools for this chain.
	allListLocalPool := len(listed) > 0
	listsLocalPool := false
	listedAddrs := make([]string, 0, len(listed))
	for _, stored := range listed {
		storedAddr, err := deploy.BytesToString(entry.selector, stored)
		if err != nil {
			return fmt.Errorf("failed to normalize peer remote pool address for chain selector %d: %w", entry.selector, err)
		}
		listedAddrs = append(listedAddrs, storedAddr)
		if storedAddr == localPoolAddr {
			listsLocalPool = true
		} else {
			allListLocalPool = false
		}
	}
	if !listsLocalPool {
		p.env.Logger.Warnf("skipping reverse removal of pool %s on chain %d from peer pool %s on chain %d: peer pool does not list this pool as a remote", entry.poolRef.Address, entry.selector, target.Address, remoteSelector)
		return nil
	}

	// With wipeChainConfigs, the peer's lane for this chain is a wipe candidate: resolveWipes wipes it
	// if the merged removals empty it. A peer adapter that cannot wipe is an unsupported peer when this
	// removal alone empties the lane: fail, or with skipUnsupportedPeers warn and remove only the pool
	// entry below. Known limitation: such a peer is not recorded as a candidate, so a lane emptied only
	// by removals merged from several entries keeps its chain config without a warning. No current
	// adapter reaches this: the EVM versions that cannot wipe (1.2.0, 1.4.0) also fail every remote
	// pool removal, so it only applies to a future family that removes pools but cannot wipe.
	if entry.pool.WipeChainConfigs {
		chainRemover, canWipe := remoteAdapter.(RemoteChainConfigRemover)
		canWipe = canWipe && chainRemover.SupportsRemoteChainConfigRemoval()
		switch {
		case canWipe:
			if err := p.addWipeCandidate(chainRemover, RemoveRemoteChainConfigsSequenceInput{
				Selector:     remoteSelector,
				TokenPoolRef: target,
				TokenRef:     remoteTokenRef,
			}, []removeRemotePoolsLane{{selector: entry.selector, listed: listedAddrs}}, false); err != nil {
				return err
			}
		case allListLocalPool && !entry.pool.SkipUnsupportedPeers:
			return fmt.Errorf("peer on remote chain selector %d is not supported: adapter for peer pool %s does not support removing remote chain configs (set skipUnsupportedPeers to remove the peer's pool entry while leaving its chain config)", remoteSelector, target.Address)
		case allListLocalPool:
			p.env.Logger.Warnf("removing pool %s on chain %d from peer pool %s on chain %d, but leaving the peer's chain config for chain %d in place: the peer adapter does not support removing remote chain configs; clean it up on chain %d with its own tooling", entry.poolRef.Address, entry.selector, target.Address, remoteSelector, entry.selector, remoteSelector)
		}
	}

	if alreadyQueued {
		p.env.Logger.Warnf("skipping reverse removal of pool %s on chain %d from peer pool %s on chain %d: pairing already removed", entry.poolRef.Address, entry.selector, target.Address, remoteSelector)
		return nil
	}
	p.removedPairings[key] = struct{}{}

	return p.queueRemoval(remoteRemover, RemoveRemotePoolsSequenceInput{
		Selector:            remoteSelector,
		TokenPoolRef:        target,
		TokenRef:            remoteTokenRef,
		RemotePoolsToRemove: []RemotePoolToRemove{{Selector: entry.selector, Remote: datastore.AddressRef{Address: localPoolAddr}}},
	}, false)
}
