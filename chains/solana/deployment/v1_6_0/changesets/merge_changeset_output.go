package changesets

import (
	"cmp"
	"errors"
	"fmt"
	"maps"
	"reflect"
	"slices"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
)

// MergeChangesetOutput merges the source ChangesetOutput into the destination ChangesetOutput.
//
// The source's address book and data store are also published to the environment, so a
// sub-changeset merged later can resolve contracts that an earlier one deployed.
//
// The address-book and data-store merge is validated in full before any of it is applied. If the
// source cannot be merged, nothing is changed and the caller can fix the source and merge again
// from the same starting point, rather than discovering the problem with the destination already
// partly updated.
func MergeChangesetOutput(e cldf.Environment, dest *cldf.ChangesetOutput, src cldf.ChangesetOutput) error {
	if dest == nil {
		return nil
	}

	plan, err := planChangesetMerge(e, *dest, src)
	if err != nil {
		return err
	}
	if err := plan.applyAddressBookAndDataStore(e, dest, src); err != nil {
		return err
	}

	// Reports are merged here.
	if dest.Reports == nil {
		dest.Reports = src.Reports
	} else if src.Reports != nil {
		dest.Reports = append(dest.Reports, src.Reports...)
	}

	src.AddressBook = nil
	src.DataStore = nil
	if err := cldf.MergeChangesetOutput(e, dest, src); err != nil {
		return fmt.Errorf("failed to merge changeset output: %w", err)
	}

	return nil
}

// changesetMergePlan is a validated merge of the address book and data store, ready to apply.
// Building it mutates nothing, so a merge that turns out to be illegal is rejected before the
// destination or the environment has been touched: the caller is never left holding a half merged
// output.
type changesetMergePlan struct {
	// addressBook replaces dest.AddressBook. Nil when the source has no address book.
	addressBook cldf.AddressBook
	// envAddresses are the source's address book entries the environment does not have yet.
	envAddresses []envAddress
	// dataStore replaces dest.DataStore. Nil when the source has no data store.
	dataStore datastore.MutableDataStore
	// envRefs are the source's address refs to publish to the environment.
	envRefs []datastore.AddressRef
	// superseded are environment refs that envRefs replace, reported so a redeploy is visible
	// rather than silent.
	superseded []datastore.AddressRef
}

// envAddress is one address book entry bound for the environment.
type envAddress struct {
	chainSelector  uint64
	address        string
	typeAndVersion cldf.TypeAndVersion
}

// mergeAddressBookEntries merges address-book entries using the address rules of the chain
// family. AddressBookMap keys are strings, so its native Merge method treats differently-cased EVM
// spellings as different contracts even though they identify the same address.
func mergeAddressBookEntries(dst cldf.AddressBook, entries map[uint64]map[string]cldf.TypeAndVersion) error {
	for _, chainSelector := range sortedKeys(entries) {
		for _, address := range sortedKeys(entries[chainSelector]) {
			tv := entries[chainSelector][address]
			knownAddresses, err := dst.AddressesForChain(chainSelector)
			if err != nil && !errors.Is(err, cldf.ErrChainNotFound) {
				return err
			}

			alreadyRecorded := false
			for knownAddress, knownTV := range knownAddresses {
				if !shared.AddressesEqual(chainSelector, knownAddress, address) {
					continue
				}
				if !knownTV.Equal(tv) {
					return fmt.Errorf("address %s is already recorded as %s, cannot merge it as %s", knownAddress, knownTV.String(), tv.String())
				}
				alreadyRecorded = true
				break
			}
			if alreadyRecorded {
				continue
			}
			if err := dst.Save(chainSelector, address, tv); err != nil {
				return err
			}
		}
	}
	return nil
}

// planChangesetMerge validates the address-book and data-store merge against both the destination
// and the environment and returns what applying it would change. It performs no mutation.
func planChangesetMerge(env cldf.Environment, dest, src cldf.ChangesetOutput) (*changesetMergePlan, error) {
	plan := &changesetMergePlan{}

	if err := plan.planAddressBook(env, dest.AddressBook, src.AddressBook); err != nil {
		return nil, err
	}

	if err := plan.planDataStore(env, dest.DataStore, src.DataStore); err != nil {
		return nil, err
	}

	return plan, nil
}

// planAddressBook stages the merged address book in a fresh book, so the destination is only
// replaced once the whole merge is known to be legal, and works out which entries the
// environment is still missing.
// The merged book consists of addresses in dest, src, and env.ExistingAddresses
func (p *changesetMergePlan) planAddressBook(env cldf.Environment, dest, src cldf.AddressBook) error {
	if isNilStore(src) {
		return nil
	}

	srcAddrs, err := src.Addresses()
	if err != nil {
		return fmt.Errorf("failed to read source address book: %w", err)
	}

	merged := cldf.NewMemoryAddressBook()
	if !isNilStore(dest) {
		destAddrs, readErr := dest.Addresses()
		if readErr != nil {
			return fmt.Errorf("failed to read destination address book: %w", readErr)
		}
		if err = mergeAddressBookEntries(merged, destAddrs); err != nil {
			return fmt.Errorf("failed to stage destination address book: %w", err)
		}
	}
	if err = mergeAddressBookEntries(merged, srcAddrs); err != nil {
		return fmt.Errorf("failed to merge address book: %w", err)
	}
	p.addressBook = merged

	if isNilStore(env.ExistingAddresses) {
		return nil
	}

	envAddrs, err := env.ExistingAddresses.Addresses()
	if err != nil {
		return fmt.Errorf("failed to read environment address book: %w", err)
	}

	// An address the environment already records under the same type and version is not
	// re-saved: AddressBookMap.Save rejects a repeated address, and a changeset that registers
	// contracts the environment already knows about is doing so legitimately. An address
	// recorded under a *different* type and version is a genuine disagreement and is rejected
	// here, before anything is merged.
	for _, chainSelector := range sortedKeys(srcAddrs) {
		for _, address := range sortedKeys(srcAddrs[chainSelector]) {
			tv := srcAddrs[chainSelector][address]

			var known cldf.TypeAndVersion
			found := false
			for knownAddress, knownTV := range envAddrs[chainSelector] {
				if shared.AddressesEqual(chainSelector, knownAddress, address) {
					known = knownTV
					found = true
					break
				}
			}
			if found {
				if !known.Equal(tv) {
					return fmt.Errorf(
						"failed to merge address book: the environment records %s on chain %d as %s, the changeset as %s",
						address, chainSelector, known.String(), tv.String(),
					)
				}

				continue
			}

			p.envAddresses = append(p.envAddresses, envAddress{
				chainSelector: chainSelector, address: address, typeAndVersion: tv,
			})
		}
	}

	return nil
}

// planDataStore stages the merged data store in a fresh store and validates every source record
// against what the destination and the environment already hold.
func (p *changesetMergePlan) planDataStore(env cldf.Environment, dest, src datastore.MutableDataStore) error {
	if isNilStore(src) {
		return nil
	}

	srcRefs, err := src.Addresses().Fetch()
	if err != nil {
		return fmt.Errorf("failed to read source data store: %w", err)
	}

	// Every key is built from a version, and semver's String has a value receiver, so a record
	// with no version would panic rather than report itself. Reject it as the stores do.
	for _, ref := range srcRefs {
		if ref.Version == nil {
			return fmt.Errorf("failed to merge data store: address %s on chain %d has no version: %w",
				ref.Address, ref.ChainSelector, datastore.ErrAddressRefVersionRequired)
		}
		// An empty address is not a record of a deployed contract: it is a reservation
		// (ReserveRefs) that the changeset never filled. Merged, it would persist as a ref that
		// resolves to nothing and is mistaken for progress.
		if ref.Address == "" {
			return fmt.Errorf("failed to merge data store: %s has no address; a reserved key was never filled",
				ref.Key().String())
		}
	}

	// MemoryDataStore keeps a record alongside a staged deletion until Merge applies it. Do not
	// validate or publish those records: they are not part of the resulting store.
	if srcRefsStore, ok := src.Addresses().(*datastore.MemoryAddressRefStore); ok && len(srcRefsStore.DeletedRemoteKeys) > 0 {
		deleted := make(map[string]struct{}, len(srcRefsStore.DeletedRemoteKeys))
		for _, key := range srcRefsStore.DeletedRemoteKeys {
			deleted[key] = struct{}{}
		}

		activeRefs := srcRefs[:0]
		for _, ref := range srcRefs {
			if _, isDeleted := deleted[ref.Key().String()]; !isDeleted {
				activeRefs = append(activeRefs, ref)
			}
		}
		srcRefs = activeRefs
	}

	// Two sub-changesets of one run claiming the same key is a mistake, not a redeploy: only
	// one of the two addresses can survive, and MutableDataStore.Merge upserts, so without
	// this check the later one silently wins and an address is lost.
	claimed := make(map[string]datastore.AddressRef, len(srcRefs))
	for _, ref := range srcRefs {
		key := ref.Key().String()
		if earlier, taken := claimed[key]; taken {
			if !shared.AddressesEqual(ref.ChainSelector, earlier.Address, ref.Address) {
				return fmt.Errorf(
					"failed to merge data store: addresses %s and %s both claim %s; give them distinct qualifiers",
					earlier.Address, ref.Address, key,
				)
			}

			continue
		}
		claimed[key] = ref
	}

	if !isNilStore(dest) {
		destRefs, fetchErr := dest.Addresses().Fetch()
		if fetchErr != nil {
			return fmt.Errorf("failed to read destination data store: %w", fetchErr)
		}

		heldBy := make(map[string]string, len(destRefs))
		for _, ref := range destRefs {
			if ref.Version == nil {
				continue
			}
			heldBy[ref.Key().String()] = ref.Address
		}

		for _, ref := range srcRefs {
			if held, taken := heldBy[ref.Key().String()]; taken && !shared.AddressesEqual(ref.ChainSelector, held, ref.Address) {
				return fmt.Errorf(
					"failed to merge data store: addresses %s and %s both claim %s; give them distinct qualifiers",
					held, ref.Address, ref.Key().String(),
				)
			}
		}
	}

	if err = p.planMetadata(dest, src); err != nil {
		return err
	}

	merged := datastore.NewMemoryDataStore()
	if !isNilStore(dest) {
		if err = merged.Merge(dest.Seal()); err != nil {
			return fmt.Errorf("failed to stage destination data store: %w", err)
		}
	}
	if err = merged.Merge(src.Seal()); err != nil {
		return fmt.Errorf("failed to merge data store: %w", err)
	}
	p.dataStore = merged

	if isNilStore(env.DataStore) {
		return nil
	}

	// The environment is the state before this run, so a source ref taking over a key it holds
	// is a redeploy superseding what was there, not a collision. It is recorded so the
	// replacement is reported rather than silent.
	envRefs, err := env.DataStore.Addresses().Fetch()
	if err != nil {
		return fmt.Errorf("failed to read environment data store: %w", err)
	}

	envHeldBy := make(map[string]datastore.AddressRef, len(envRefs))
	for _, ref := range envRefs {
		if ref.Version == nil {
			continue
		}
		envHeldBy[ref.Key().String()] = ref
	}

	for _, ref := range srcRefs {
		if held, taken := envHeldBy[ref.Key().String()]; taken && !shared.AddressesEqual(ref.ChainSelector, held.Address, ref.Address) {
			p.superseded = append(p.superseded, held)
		}
	}
	p.envRefs = srcRefs

	return nil
}

// planMetadata rejects metadata the merge would overwrite with something different. The metadata
// stores upsert, and the environment metadata is a single record that Merge simply re-sets, so two
// sub-changesets disagreeing about a record would silently keep the later one. Identical records
// merge as before.
func (p *changesetMergePlan) planMetadata(dest, src datastore.MutableDataStore) error {
	if isNilStore(dest) {
		return nil
	}

	destChain, err := dest.ChainMetadata().Fetch()
	if err != nil {
		return fmt.Errorf("failed to read destination chain metadata: %w", err)
	}

	held := make(map[string]any, len(destChain))
	for _, record := range destChain {
		held[record.Key().String()] = record.Metadata
	}

	srcChain, err := src.ChainMetadata().Fetch()
	if err != nil {
		return fmt.Errorf("failed to read source chain metadata: %w", err)
	}

	for _, record := range srcChain {
		if existing, taken := held[record.Key().String()]; taken && !reflect.DeepEqual(existing, record.Metadata) {
			return fmt.Errorf("failed to merge data store: conflicting chain metadata for chain %d",
				record.ChainSelector)
		}
	}

	destContract, err := dest.ContractMetadata().Fetch()
	if err != nil {
		return fmt.Errorf("failed to read destination contract metadata: %w", err)
	}

	held = make(map[string]any, len(destContract))
	for _, record := range destContract {
		held[record.Key().String()] = record.Metadata
	}

	srcContract, err := src.ContractMetadata().Fetch()
	if err != nil {
		return fmt.Errorf("failed to read source contract metadata: %w", err)
	}

	for _, record := range srcContract {
		if existing, taken := held[record.Key().String()]; taken && !reflect.DeepEqual(existing, record.Metadata) {
			return fmt.Errorf("failed to merge data store: conflicting contract metadata for %s on chain %d",
				record.Address, record.ChainSelector)
		}
	}

	srcEnv, err := src.EnvMetadata().Get()
	if err != nil {
		if errors.Is(err, datastore.ErrEnvMetadataNotSet) {
			return nil
		}

		return fmt.Errorf("failed to read source environment metadata: %w", err)
	}

	destEnv, err := dest.EnvMetadata().Get()
	if err != nil {
		if errors.Is(err, datastore.ErrEnvMetadataNotSet) {
			return nil
		}

		return fmt.Errorf("failed to read destination environment metadata: %w", err)
	}

	if !reflect.DeepEqual(destEnv.Metadata, srcEnv.Metadata) {
		return errors.New("failed to merge data store: conflicting environment metadata")
	}

	return nil
}

// applyAddressBookAndDataStore commits the validated address-book and data-store merge. Everything
// it does was validated while the plan was built, so the remaining error paths are reported as
// what they are: a store that changed under us.
func (p *changesetMergePlan) applyAddressBookAndDataStore(env cldf.Environment, dest *cldf.ChangesetOutput, src cldf.ChangesetOutput) error {
	for _, entry := range p.envAddresses {
		if err := env.ExistingAddresses.Save(entry.chainSelector, entry.address, entry.typeAndVersion); err != nil {
			return fmt.Errorf("failed to merge existing addresses to environment after validation: %w", err)
		}
	}

	if len(p.envRefs) > 0 {
		mutable, ok := env.DataStore.Addresses().(datastore.MutableAddressRefStore)
		if !ok {
			if env.Logger != nil {
				env.Logger.Warnw("Environment data store is not writable; "+
					"contracts deployed by this changeset will not be resolvable through env.DataStore",
					"refs", len(p.envRefs))
			}
		} else {
			for _, ref := range p.envRefs {
				if err := mutable.Upsert(ref); err != nil {
					return fmt.Errorf("failed to merge data store into environment after validation: %w", err)
				}
			}
		}
	}

	// The environment also has to see the source's staged deletions and metadata, or a later
	// sub-changeset would resolve stale records through env.DataStore.
	if err := propagateSrcDataStoreToEnv(env, src.DataStore); err != nil {
		return err
	}

	if env.Logger != nil {
		for _, ref := range p.superseded {
			env.Logger.Infow("Changeset output supersedes an existing address ref",
				"key", ref.Key().String(), "previousAddress", ref.Address)
		}
	}

	if p.addressBook != nil {
		dest.AddressBook = p.addressBook
	}
	if p.dataStore != nil {
		dest.DataStore = p.dataStore
	}

	return nil
}

// propagateSrcDataStoreToEnv applies the source's staged deletions and metadata to the
// environment's data store, so a sub-changeset merged later resolves the same state through
// env.DataStore that the final ChangesetOutput carries. It mirrors MemoryDataStore.Merge's
// deletion and metadata handling, writing through the environment store's mutable interfaces.
// A store that is not memory-backed (and so cannot be mutated) is skipped with a warning, as in
// the address-ref propagation above.
func propagateSrcDataStoreToEnv(env cldf.Environment, src datastore.MutableDataStore) error {
	if isNilStore(env.DataStore) || isNilStore(src) {
		return nil
	}

	if srcAddr, ok := src.Addresses().(*datastore.MemoryAddressRefStore); ok {
		if envAddr, ok := env.DataStore.Addresses().(datastore.MutableAddressRefStore); ok {
			for _, dk := range srcAddr.DeletedRemoteKeys {
				key, keyErr := datastore.NewAddressRefKeyFromString(dk)
				if keyErr != nil {
					return fmt.Errorf("failed to parse address ref deletion key %q: %w", dk, keyErr)
				}
				if err := envAddr.RemoteDelete(key); err != nil {
					return fmt.Errorf("failed to propagate address ref deletion to environment: %w", err)
				}
				if err := envAddr.Delete(key); err != nil && !errors.Is(err, datastore.ErrAddressRefNotFound) {
					return fmt.Errorf("failed to propagate address ref deletion to environment: %w", err)
				}
			}
		} else if env.Logger != nil {
			env.Logger.Warnw("Environment data store is not writable; " +
				"address ref deletions from this changeset will not be visible through env.DataStore")
		}
	}

	if srcChain, ok := src.ChainMetadata().(*datastore.MemoryChainMetadataStore); ok {
		if envChain, ok := env.DataStore.ChainMetadata().(datastore.MutableChainMetadataStore); ok {
			for _, record := range srcChain.Records {
				if err := envChain.Upsert(record); err != nil {
					return fmt.Errorf("failed to propagate chain metadata to environment: %w", err)
				}
			}
			for _, dk := range srcChain.DeletedRemoteKeys {
				key, keyErr := datastore.NewChainMetadataKeyFromString(dk)
				if keyErr != nil {
					return fmt.Errorf("failed to parse chain metadata deletion key %q: %w", dk, keyErr)
				}
				if err := envChain.RemoteDelete(key); err != nil {
					return fmt.Errorf("failed to propagate chain metadata deletion to environment: %w", err)
				}
				if err := envChain.Delete(key); err != nil && !errors.Is(err, datastore.ErrChainMetadataNotFound) {
					return fmt.Errorf("failed to propagate chain metadata deletion to environment: %w", err)
				}
			}
		}
	}

	if srcContract, ok := src.ContractMetadata().(*datastore.MemoryContractMetadataStore); ok {
		if envContract, ok := env.DataStore.ContractMetadata().(datastore.MutableContractMetadataStore); ok {
			for _, record := range srcContract.Records {
				if err := envContract.Upsert(record); err != nil {
					return fmt.Errorf("failed to propagate contract metadata to environment: %w", err)
				}
			}
			for _, dk := range srcContract.DeletedRemoteKeys {
				key, keyErr := datastore.NewContractMetadataKeyFromString(dk)
				if keyErr != nil {
					return fmt.Errorf("failed to parse contract metadata deletion key %q: %w", dk, keyErr)
				}
				if err := envContract.RemoteDelete(key); err != nil {
					return fmt.Errorf("failed to propagate contract metadata deletion to environment: %w", err)
				}
				if err := envContract.Delete(key); err != nil && !errors.Is(err, datastore.ErrContractMetadataNotFound) {
					return fmt.Errorf("failed to propagate contract metadata deletion to environment: %w", err)
				}
			}
		}
	}

	if srcEnv, ok := env.DataStore.EnvMetadata().(datastore.MutableEnvMetadataStore); ok {
		envMeta, err := src.EnvMetadata().Get()
		if err == nil {
			if err := srcEnv.Set(envMeta); err != nil {
				return fmt.Errorf("failed to propagate environment metadata to environment: %w", err)
			}
		} else if !errors.Is(err, datastore.ErrEnvMetadataNotSet) {
			return fmt.Errorf("failed to read source environment metadata: %w", err)
		}
	}

	return nil
}

// isNilStore reports whether v is unset, including the case of a non-nil interface holding a
// nil pointer. A changeset returning ChangesetOutput{DataStore: (*datastore.MemoryDataStore)(nil)}
// passes an ordinary != nil check and then panics on first use.
func isNilStore(v any) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)

	return rv.Kind() == reflect.Pointer && rv.IsNil()
}

// sortedKeys returns m's keys in order, so a merge reports the same problem first on every run.
func sortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	return slices.Sorted(maps.Keys(m))
}
