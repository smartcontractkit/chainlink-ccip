package changesets

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// UpgradeOnrampConfig is the input for the OnRamp upgrade changesets.
type UpgradeOnrampConfig struct {
	ChainSelector                       uint64
	DisableTransferAndVerifierOwnership bool
	TestVerifierAllowedSenders          []string
	DestSelectorsInScope                []uint64
	MCMS                                mcms.Input
}

// UpgradeOnrampPhase1 upgrades the OnRamp with RMNProxy address, copies dest
// chain configs verbatim (including each dest's current Router), allows the new OnRamp on
// all remote OffRamps alongside the legacy one, and transfers ownership of the new OnRamp to
// the CLLCCIP timelock. Neither router is repointed at the new OnRamp here — that only happens
// once the verifier jobs observe both OnRamps (Phase 1.5, outside this changeset) — so this
// phase can never affect live traffic on any lane, ProdRouter or TestRouter.
//
// Pre-flight: for ProdRouter lane classes, the test verifier infrastructure (TESTVTR drip
// token, matching pool, test verifier, and resolver) must already be deployed and configured
// via the DeployTestVerifierChains changeset. This changeset validates the prerequisite but
// does not deploy test infrastructure itself. TestRouter lanes are production traffic for
// that lane and never stage.
func UpgradeOnrampPhase1(
	upgradeOnrampRegistry *adapters.OnRampUpgraderRegistry,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
	transferOwnershipReg *deploy.TransferOwnershipAdapterRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if cfg.ChainSelector == 0 {
			return fmt.Errorf("chain selector is required")
		}

		if len(cfg.DestSelectorsInScope) == 0 {
			return fmt.Errorf("at least one dest selector must be in scope")
		}

		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		if err := cfg.MCMS.PopulateDefaults(); err != nil {
			return cldf.ChangesetOutput{}, err
		}

		if err := cfg.MCMS.Validate(); err != nil {
			return cldf.ChangesetOutput{}, err
		}

		family, err := chainsel.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("get chain family: %w", err)
		}
		upgrader, ok := upgradeOnrampRegistry.Get(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no OnRampUpgrader registered for family %q", family)
		}
		localAdapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no chain family adapter for family %q", family)
		}

		// Pre-flight: Ask the adapter if the Onramp is already upgraded to the new version. If it is, we should not run this changeset.
		if err := upgrader.VerifyOnrampRequireUpgrade(e, cfg.ChainSelector); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("verify OnRamp require upgrade: %w", err)
		}

		// The old OnRamp is still canonical in the environment datastore (the merge of the
		// new refs only happens after this changeset returns). GetOnRampAddress returns the
		// source chain's wire format — for EVM, abi.encode(address), 32 bytes left-padded.
		oldOnRampBytes, err := localAdapter.GetOnRampAddress(e.DataStore, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("encode old OnRamp: %w", err)
		}
		oldOnRampHex := "0x" + hex.EncodeToString(oldOnRampBytes)

		// Pre-flight: the whitelist update below replaces each remote OffRamp's onramp
		// list, so verify it currently holds exactly the old OnRamp. Anything else would
		// be dropped silently — fail before any mutation instead.
		destChainSelectors, err := upgrader.DestChainSelectors(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve dest chain selectors: %w", err)
		}

		scopedDestSelectors := scopeDestSelectors(destChainSelectors, cfg.DestSelectorsInScope)

		for _, remoteSel := range scopedDestSelectors {
			if err := verifyOffRampSourceOnRamps(e, chainFamilyRegistry, remoteSel, cfg.ChainSelector,
				[]string{oldOnRampHex}, nil, []string{oldOnRampHex}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight OffRamp whitelist check for upgrade: %w", err)
			}
		}

		result, err := upgrader.DeployNewOnRamp(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("upgrade OnRamp: %w", err)
		}

		batchOps := append([]mcms_types.BatchOperation{}, result.BatchOps...)

		ds := datastore.NewMemoryDataStore()

		// DS merge at the end of the changeset is expected to upsert the new onramp ref and replace the legacy onramp
		if addErr := ds.Addresses().Add(result.NewOnRampRef); addErr != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("add new OnRamp ref to datastore: %w", addErr)
		}
		if addErr := ds.Addresses().Add(result.LegacyOnRampRef); addErr != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("add legacy OnRamp ref to datastore: %w", addErr)
		}

		if !cfg.DisableTransferAndVerifierOwnership {
			ownershipBatchOps, _, errOwnership := deploy.TransferToTimelock(
				cfg.ChainSelector, e, cfg.MCMS, []datastore.AddressRef{result.NewOnRampRef},
				mcmsReaderRegistry, transferOwnershipReg,
			)
			if errOwnership != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("transfer ownership: %w", errOwnership)
			}
			batchOps = append(batchOps, ownershipBatchOps...)
		}

		// The new OnRamp is only in our output datastore — wire-encode it in isolation.
		newOnRampHex, err := wireEncodeOnRampRef(localAdapter, result.NewOnRampRef, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("encode new OnRamp: %w", err)
		}

		// Allow the new OnRamp on every remote OffRamp, regardless of lane class — the OffRamp
		// natively supports multiple onramps per source chain, so this never affects live traffic.
		for _, remoteSel := range scopedDestSelectors {
			setter, err := offrampSourceOnRampSetter(chainFamilyRegistry, remoteSel)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			batchOp, skipped, err := setter.SetOffRampSourceOnRamps(e, adapters.OffRampSetSourceOnRampsEntry{
				LocalChainSelector:  remoteSel,
				SourceChainSelector: cfg.ChainSelector,
				OnRamps:             []string{oldOnRampHex, newOnRampHex},
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("allow new OnRamp on OffRamp for remote %d: %w", remoteSel, err)
			}
			if !skipped && batchOp != nil {
				batchOps = append(batchOps, *batchOp)
			}
		}

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}

	return cldf.CreateChangeSet(apply, validate)
}

// UpgradeOnrampPhase2 stages the ProdRouter lane class behind the TestRouter so it can be
// smoke tested before going live: it points the new OnRamp's DefaultCCVs at the test verifier
// resolver for those dest chains, then promotes them to the TestRouter. It is a no-op for a
// chain with no ProdRouter dests (e.g. purely TestRouter chains skip this phase
// entirely and go straight to Phase 3) — in that case the ownership pre-flight below is also
// skipped, since there is nothing to write. TestRouter dests are untouched here — their
// only promotion happens in Phase 3, once the verifier jobs observe both OnRamps.
func UpgradeOnrampPhase2(
	onrampUpgraderRegistry *adapters.OnRampUpgraderRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if cfg.ChainSelector == 0 {
			return fmt.Errorf("chain selector is required")
		}

		if len(cfg.DestSelectorsInScope) == 0 {
			return fmt.Errorf("at least one dest selector must be in scope")
		}
		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		if err := cfg.MCMS.PopulateDefaults(); err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if err := cfg.MCMS.Validate(); err != nil {
			return cldf.ChangesetOutput{}, err
		}

		family, err := chainsel.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("get chain family: %w", err)
		}
		upgrader, ok := onrampUpgraderRegistry.Get(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no OnRampUpgrader registered for family %q", family)
		}

		laneClass, err := upgrader.ClassifyDestChains(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("classify dest chains: %w", err)
		}
		if len(laneClass.ProdRouterDests) == 0 {
			return cldf.ChangesetOutput{}, nil
		}

		scopedDestSelectors := scopeDestSelectors(laneClass.ProdRouterDests, cfg.DestSelectorsInScope)

		if !cfg.DisableTransferAndVerifierOwnership {
			// Pre-flight: the new OnRamp must be timelock-owned by now; otherwise the
			// staging writes below would execute directly with the deployer key instead
			// of going through MCMS.
			mcmsReader, ok := mcmsReaderRegistry.GetMCMSReader(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no MCMS reader registered for chain family %q", family)
			}
			timelockRef, err := mcmsReader.GetTimelockRef(e, cfg.ChainSelector, cfg.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("resolve timelock: %w", err)
			}
			if err := upgrader.VerifyNewOnRampOwner(e, cfg.ChainSelector, timelockRef.Address); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight ownership check: %w", err)
			}
		}

		batchOps := make([]mcms_types.BatchOperation, 0)
		operationsBundle := operations.NewBundle(func() context.Context { return context.Background() }, e.Logger, operations.NewMemoryReporter())
		e.OperationsBundle = operationsBundle
		promoteOps, err := upgrader.PromoteOnrampToTestRouter(e, cfg.ChainSelector, scopedDestSelectors)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("stage on TestRouter: %w", err)
		}
		batchOps = append(batchOps, promoteOps...)

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}

	return cldf.CreateChangeSet(apply, validate)
}

// UpgradeOnrampPhase3 promotes the new OnRamp to production for both lane classes:
//   - ProdRouter dests are restored to the real committee CCVs and promoted from the
//     TestRouter (where Phase 2 staged them) to the prod Router.
//   - TestRoutee dests are promoted directly to the TestRouter — this is their only
//     promotion, since their production path is the TestRouter itself and they never stage.
//
// Both promotions require the verifier jobs to already observe both OnRamps (Phase 1.5), so
// this phase never causes downtime on either lane class.
func UpgradeOnrampPhase3(
	onrampUpgrader *adapters.OnRampUpgraderRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if cfg.ChainSelector == 0 {
			return fmt.Errorf("chain selector is required")
		}

		if len(cfg.DestSelectorsInScope) == 0 {
			return fmt.Errorf("at least one dest selector must be in scope")
		}

		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		if err := cfg.MCMS.PopulateDefaults(); err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if err := cfg.MCMS.Validate(); err != nil {
			return cldf.ChangesetOutput{}, err
		}

		family, err := chainsel.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("get chain family: %w", err)
		}
		upgrader, ok := onrampUpgrader.Get(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no OnRampUpgrader registered for family %q", family)
		}

		if !cfg.DisableTransferAndVerifierOwnership {
			// Pre-flight: the new OnRamp must be timelock-owned by now; otherwise the
			// promotion writes would execute directly with the deployer key instead of
			// going through MCMS.
			mcmsReader, ok := mcmsReaderRegistry.GetMCMSReader(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no MCMS reader registered for chain family %q", family)
			}
			timelockRef, err := mcmsReader.GetTimelockRef(e, cfg.ChainSelector, cfg.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("resolve timelock: %w", err)
			}
			if err := upgrader.VerifyNewOnRampOwner(e, cfg.ChainSelector, timelockRef.Address); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight ownership check: %w", err)
			}
		}

		laneClass, err := upgrader.ClassifyDestChains(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("classify dest chains: %w", err)
		}

		scopedProdDestSelectors := scopeDestSelectors(laneClass.ProdRouterDests, cfg.DestSelectorsInScope)

		scopedTestDestSelectors := scopeDestSelectors(laneClass.TestRouterDests, cfg.DestSelectorsInScope)

		batchOps := make([]mcms_types.BatchOperation, 0)
		operationsBundle := operations.NewBundle(func() context.Context { return context.Background() }, e.Logger, operations.NewMemoryReporter())
		e.OperationsBundle = operationsBundle

		if len(scopedProdDestSelectors) > 0 {
			// Pre-flight: the prod Router must still route through the legacy OnRamp for these
			// dest chains; anything else means Phase 3 already ran or an unexpected OnRamp is live.
			if err := upgrader.VerifyLegacyOnRampOnProdRouter(e, cfg.ChainSelector, scopedProdDestSelectors); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight prod Router check: %w", err)
			}

			// Fresh bundle: PromoteOnrampToProdRouter's own dest-config read must see the
			// write CopyDefaultCCVsFromLegacyOnRamp just made, not a memoized pre-write report.
			e.OperationsBundle = operations.NewBundle(func() context.Context { return context.Background() }, e.Logger, operations.NewMemoryReporter())
			promoteOps, err := upgrader.PromoteOnrampToProdRouter(e, cfg.ChainSelector, scopedProdDestSelectors)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("promote to prod Router: %w", err)
			}
			batchOps = append(batchOps, promoteOps...)
		}

		if len(scopedTestDestSelectors) > 0 {
			promoteOps, err := upgrader.PromoteOnrampToTestRouter(e, cfg.ChainSelector, scopedTestDestSelectors)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("promote to TestRouter: %w", err)
			}
			batchOps = append(batchOps, promoteOps...)
		}

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}

	return cldf.CreateChangeSet(apply, validate)
}

// Filter the destSelectors to only those in scope. Return an error if none are in scope.
func scopeDestSelectors(destSelectors []uint64, scope []uint64) []uint64 {
	scoped := make([]uint64, 0)
	scopeMap := make(map[uint64]struct{}, len(scope))
	for _, s := range scope {
		scopeMap[s] = struct{}{}
	}
	for _, d := range destSelectors {
		if _, ok := scopeMap[d]; ok {
			scoped = append(scoped, d)
		}
	}
	return scoped
}

// UpgradeOnrampCleanup disallows the old (legacy) OnRamp from all remote OffRamps.
// After this, only the new OnRamp remains allowed.
//
// The CLD framework's datastore merge is upsert-only and cannot delete entries, so operators
// must manually remove the legacy OnRamp (qualifier="legacy") from the environment datastore
// after this changeset executes. NOPs should also manually remove the temp verifier jobs from
// their CL nodes.
func UpgradeOnrampCleanup(
	onrampUpgraderRegistry *adapters.OnRampUpgraderRegistry,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if cfg.ChainSelector == 0 {
			return fmt.Errorf("chain selector is required")
		}

		if len(cfg.DestSelectorsInScope) == 0 {
			return fmt.Errorf("at least one dest selector must be in scope")
		}

		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		if err := cfg.MCMS.PopulateDefaults(); err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if err := cfg.MCMS.Validate(); err != nil {
			return cldf.ChangesetOutput{}, err
		}

		family, err := chainsel.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("get chain family: %w", err)
		}
		chainAdapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no adapter for family %q", family)
		}
		upgrader, ok := onrampUpgraderRegistry.Get(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no OnRampUpgrader registered for family %q", family)
		}

		newOnRampAddr, err := chainAdapter.GetOnRampAddress(e.DataStore, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve new OnRamp: %w", err)
		}
		newOnRampHex := "0x" + hex.EncodeToString(newOnRampAddr)

		destChainSelectors, err := upgrader.DestChainSelectors(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve dest chain selectors: %w", err)
		}

		laneClass, err := upgrader.ClassifyDestChains(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("classify dest chains: %w", err)
		}

		scopedProdDestSelectors := scopeDestSelectors(laneClass.ProdRouterDests, cfg.DestSelectorsInScope)
		scopedTestDestSelectors := scopeDestSelectors(laneClass.TestRouterDests, cfg.DestSelectorsInScope)

		// Pre-flight: removing the legacy OnRamp from remote OffRamp whitelists is only
		// safe once every dest chain has been promoted to its final router (prod Router for
		// the ProdRouter class, TestRouter for the TestRouter class).
		if err := upgrader.VerifyPromotedToRouters(e, cfg.ChainSelector, scopedProdDestSelectors, scopedTestDestSelectors); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight promotion check: %w", err)
		}

		// The whitelist update below replaces each remote OffRamp's onramp list with
		// [new]; the only onramp it may remove is the legacy one.
		var expectedRemovals []string
		legacyRef, err := upgrader.LegacyOnRampRef(e, cfg.ChainSelector)
		if err != nil {
			e.Logger.Infow("No legacy OnRamp ref found; only the new OnRamp is expected on remote OffRamps",
				"chain", cfg.ChainSelector, "err", err)
		} else {
			legacyOnRampHex, err := wireEncodeOnRampRef(chainAdapter, legacyRef, cfg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("encode legacy OnRamp: %w", err)
			}
			expectedRemovals = []string{legacyOnRampHex}
		}

		scopedDestSelectors := scopeDestSelectors(destChainSelectors, cfg.DestSelectorsInScope)

		// Pre-flight: verify no remote OffRamp whitelists an unknown onramp that the
		// update below would silently drop.
		for _, remoteSel := range scopedDestSelectors {
			if err := verifyOffRampSourceOnRamps(e, chainFamilyRegistry, remoteSel, cfg.ChainSelector,
				[]string{newOnRampHex}, expectedRemovals, nil); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight OffRamp whitelist check: %w", err)
			}
		}

		batchOps := make([]mcms_types.BatchOperation, 0)

		for _, remoteSel := range scopedDestSelectors {
			setter, err := offrampSourceOnRampSetter(chainFamilyRegistry, remoteSel)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			batchOp, skipped, err := setter.SetOffRampSourceOnRamps(e, adapters.OffRampSetSourceOnRampsEntry{
				LocalChainSelector:  remoteSel,
				SourceChainSelector: cfg.ChainSelector,
				OnRamps:             []string{newOnRampHex},
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("disallow old OnRamp on OffRamp for remote %d: %w", remoteSel, err)
			}
			if !skipped && batchOp != nil {
				batchOps = append(batchOps, *batchOp)
			}
		}

		e.Logger.Infow("OnRamp upgrade cleanup complete. Operators must manually: (1) remove the legacy OnRamp (qualifier=\"legacy\") from the datastore, (2) tell NOPs to remove temp verifier jobs",
			"chain", cfg.ChainSelector)

		if len(batchOps) == 0 {
			return cldf.ChangesetOutput{}, nil
		}

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}

	return cldf.CreateChangeSet(apply, validate)
}

// remoteChainFamilyAdapter resolves the chain family adapter for a remote chain.
func remoteChainFamilyAdapter(reg *adapters.ChainFamilyRegistry, remoteSel uint64) (adapters.ChainFamily, string, error) {
	remoteFamily, err := chainsel.GetSelectorFamily(remoteSel)
	if err != nil {
		return nil, "", fmt.Errorf("get remote chain family for %d: %w", remoteSel, err)
	}
	remoteAdapter, ok := reg.GetChainFamily(remoteFamily)
	if !ok {
		return nil, "", fmt.Errorf("no adapter for remote family %q", remoteFamily)
	}
	return remoteAdapter, remoteFamily, nil
}

// offrampSourceOnRampSetter resolves the remote chain's family adapter and asserts it
// supports OffRamp source onramp updates.
func offrampSourceOnRampSetter(reg *adapters.ChainFamilyRegistry, remoteSel uint64) (adapters.OffRampSourceOnRampSetter, error) {
	remoteAdapter, remoteFamily, err := remoteChainFamilyAdapter(reg, remoteSel)
	if err != nil {
		return nil, err
	}
	setter, ok := remoteAdapter.(adapters.OffRampSourceOnRampSetter)
	if !ok {
		return nil, fmt.Errorf("remote family %q does not support OffRamp source onramp updates", remoteFamily)
	}
	return setter, nil
}

// wireEncodeOnRampRef returns the 0x-prefixed wire-format hex of an OnRamp ref, using a
// temp datastore so the family adapter resolves it as the canonical OnRamp.
func wireEncodeOnRampRef(adapter adapters.ChainFamily, ref datastore.AddressRef, chainSelector uint64) (string, error) {
	ref.Qualifier = ""
	tmp := datastore.NewMemoryDataStore()
	if err := tmp.Addresses().Add(ref); err != nil {
		return "", fmt.Errorf("add OnRamp ref to temp datastore: %w", err)
	}
	onRampBytes, err := adapter.GetOnRampAddress(tmp.Seal(), chainSelector)
	if err != nil {
		return "", fmt.Errorf("encode OnRamp %s: %w", ref.Address, err)
	}
	return "0x" + hex.EncodeToString(onRampBytes), nil
}

// verifyOffRampSourceOnRamps reads the current onramp whitelist for sourceChainSelector
// on the remote chain's OffRamp and fails if replacing it with desired would remove an
// onramp not listed in expectedRemovals, or if any of mustContain is not currently
// whitelisted. This guards against silently dropping an onramp that still serves traffic.
func verifyOffRampSourceOnRamps(
	e cldf.Environment,
	reg *adapters.ChainFamilyRegistry,
	remoteSel uint64,
	sourceSel uint64,
	desired []string,
	expectedRemovals []string,
	mustContain []string,
) error {
	remoteAdapter, remoteFamily, err := remoteChainFamilyAdapter(reg, remoteSel)
	if err != nil {
		return err
	}
	reader, ok := remoteAdapter.(adapters.OffRampSourceOnRampReader)
	if !ok {
		return fmt.Errorf("remote family %q does not support OffRamp source onramp reads", remoteFamily)
	}

	current, err := reader.GetOffRampSourceOnRamps(e, remoteSel, sourceSel)
	if err != nil {
		return fmt.Errorf("read current onramps on OffRamp on chain %d: %w", remoteSel, err)
	}

	allowed := make(map[string]struct{}, len(desired)+len(expectedRemovals))
	for _, addr := range desired {
		allowed[strings.ToLower(addr)] = struct{}{}
	}
	for _, addr := range expectedRemovals {
		allowed[strings.ToLower(addr)] = struct{}{}
	}

	currentSet := make(map[string]struct{}, len(current))
	var unexpected []string
	for _, addr := range current {
		key := strings.ToLower(addr)
		currentSet[key] = struct{}{}
		if _, ok := allowed[key]; !ok {
			unexpected = append(unexpected, addr)
		}
	}
	if len(unexpected) > 0 {
		return fmt.Errorf("OffRamp on chain %d whitelists unexpected onramps %v for source chain %d; "+
			"refusing to override them — investigate and use the OffRampSetSourceOnRamps changeset to update deliberately",
			remoteSel, unexpected, sourceSel)
	}
	for _, addr := range mustContain {
		if _, ok := currentSet[strings.ToLower(addr)]; !ok {
			return fmt.Errorf("OffRamp on chain %d does not whitelist expected onramp %s for source chain %d (current: %v)",
				remoteSel, addr, sourceSel, current)
		}
	}
	return nil
}
