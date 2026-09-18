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
//   - deactivate: must be specified alone. Unregisters the pool from the TokenAdminRegistry
//     (setPool to zero) AND performs the allRemotes forward cleanup AND the bidirectional reverse
//     cleanup, completing the full teardown.
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
				// deactivate implies allRemotes + bidirectional + TAR unregister, so the
				// allRemotes/bidirectional flags must not be set alongside it. remotePoolsToRemove
				// is allowed to provide the explicit remotes (required on Solana, where supported
				// chains cannot be enumerated on-chain).
				if pool.AllRemotes || pool.Bidirectional {
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

func removeRemotePoolsApply() func(cldf.Environment, RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RemoveRemotePoolsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		tokenRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		// removedPairings tracks (poolSelector, remoteSelector) pairs that have already been torn
		// down in the forward pass, so a bidirectional removal never double-removes a pairing when
		// A and B are both specified for a bidirectional removal of each other.
		removedPairings := make(map[uint64]map[uint64]struct{})

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
			remotesToRemove, err := resolveRemotesToRemove(e, adapter, family, selector, fullPoolRef, pool)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve remotes to remove for pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
			}

			// Forward pass: remove the specified remotes from this pool.
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, remover.RemoveRemotePools(), e.BlockChains, RemoveRemotePoolsSequenceInput{
				Selector:            selector,
				TokenPoolRef:        fullPoolRef,
				TokenRef:            fullTokenRef,
				RemotePoolsToRemove: remotesToRemove,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to remove remote pools from pool %s on chain selector %d: %w", datastore_utils.SprintRef(pool.Pool), selector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)

			if removedPairings[selector] == nil {
				removedPairings[selector] = make(map[uint64]struct{})
			}
			for _, remote := range remotesToRemove {
				removedPairings[selector][remote.Selector] = struct{}{}
			}

			// Reverse pass: for bidirectional/deactivate, remove this pool from each peer's remote list.
			if pool.Bidirectional || pool.Deactivate {
				reverseBatchOps, reverseReports, err := removeRemotePoolsReverse(e, tokenRegistry, selector, fullPoolRef, fullTokenRef, remotesToRemove, removedPairings)
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

	supportedChains, err := migrator.GetSupportedChains(e, selector, poolBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to discover supported chains for pool on chain selector %d: %w", selector, err)
	}

	remotes := make([]RemotePoolToRemove, 0, len(supportedChains))
	for _, remoteSelector := range supportedChains {
		remotePools, err := migrator.GetRemotePools(e, selector, poolBytes, remoteSelector)
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
	selector uint64,
	fullPoolRef datastore.AddressRef,
	fullTokenRef datastore.AddressRef,
	remotesToRemove []RemotePoolToRemove,
	removedPairings map[uint64]map[uint64]struct{},
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)

	for _, remote := range remotesToRemove {
		remoteSelector := remote.Selector

		if _, alreadyRemoved := removedPairings[remoteSelector][selector]; alreadyRemoved {
			e.Logger.Warnf("skipping reverse removal of pool %s on chain %d from pool on chain %d: pairing already removed", fullPoolRef.Address, selector, remoteSelector)
			continue
		}

		remoteFamily, err := chain_selectors.GetSelectorFamily(remoteSelector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteSelector, err)
		}

		// Resolve the peer's adapter, pool ref, and token ref directly from the peer pool address
		// given in the forward input. This avoids needing the TokenPoolMigrator (which cannot
		// resolve the mint from a Solana pool program ID alone).
		remoteAdapter, _, remotePoolRef, remoteTokenRef, err := ResolveAdapterAndRefs(e, tokenRegistry, remoteSelector, remote.Remote, datastore.AddressRef{})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to resolve peer adapter on remote chain selector %d: %w", remoteSelector, err)
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
		peerRemotes, err := peerMigrator.GetRemotePools(e, remoteSelector, activePool, selector)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read peer remote pools for chain selector %d on remote chain selector %d: %w", selector, remoteSelector, err)
		}
		if len(peerRemotes) == 0 {
			e.Logger.Warnf("skipping reverse removal of pool %s on chain %d from pool on chain %d: peer has no remote pool for this chain", fullPoolRef.Address, selector, remoteSelector)
			continue
		}

		remotesToRemoveFromPeer := make([]RemotePoolToRemove, 0, len(peerRemotes))
		for _, peerRemote := range peerRemotes {
			peerRemoteAddr, err := deploy.BytesToString(selector, peerRemote)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to normalize peer remote pool address for chain selector %d: %w", selector, err)
			}
			remotesToRemoveFromPeer = append(remotesToRemoveFromPeer, RemotePoolToRemove{
				Selector: selector,
				Remote:   datastore.AddressRef{Address: peerRemoteAddr},
			})
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, remoteRemover.RemoveRemotePools(), e.BlockChains, RemoveRemotePoolsSequenceInput{
			Selector:            remoteSelector,
			TokenPoolRef:        remotePoolRef,
			TokenRef:            remoteTokenRef,
			RemotePoolsToRemove: remotesToRemoveFromPeer,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to remove pool from peer on remote chain selector %d: %w", remoteSelector, err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		if removedPairings[remoteSelector] == nil {
			removedPairings[remoteSelector] = make(map[uint64]struct{})
		}
		removedPairings[remoteSelector][selector] = struct{}{}
	}

	return batchOps, reports, nil
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
