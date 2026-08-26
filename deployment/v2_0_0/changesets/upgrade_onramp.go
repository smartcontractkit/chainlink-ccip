package changesets

import (
	"context"
	"encoding/hex"
	"fmt"

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
	DestSelectorsInScope                []uint64
	MCMS                                mcms.Input
}

func encodeAddressToHex(addr []byte) string {
	return "0x" + hex.EncodeToString(addr)
}

func validateUpgradeOnrampConfig(cfg UpgradeOnrampConfig) error {
	if cfg.ChainSelector == 0 {
		return fmt.Errorf("chain selector is required")
	}
	if len(cfg.DestSelectorsInScope) == 0 {
		return fmt.Errorf("at least one dest selector must be in scope")
	}
	return nil
}

func initUpgradeOnrampChangesets(upgradeOnrampRegistry *adapters.OnRampUpgraderRegistry, cfg *UpgradeOnrampConfig) (string, adapters.OnRampUpgrader, error) {
	if err := cfg.MCMS.PopulateDefaults(); err != nil {
		return "", nil, err
	}

	if err := cfg.MCMS.Validate(); err != nil {
		return "", nil, err
	}

	family, err := chainsel.GetSelectorFamily(cfg.ChainSelector)
	if err != nil {
		return "", nil, fmt.Errorf("get chain family: %w", err)
	}
	upgrader, ok := upgradeOnrampRegistry.Get(family)
	if !ok {
		return "", nil, fmt.Errorf("no OnRampUpgrader registered for family %q", family)
	}
	return family, upgrader, nil
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
		if err := validateUpgradeOnrampConfig(cfg); err != nil {
			return fmt.Errorf("invalid UpgradeOnrampConfig: %w", err)
		}

		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		family, upgrader, err := initUpgradeOnrampChangesets(upgradeOnrampRegistry, &cfg)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		localAdapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no chain family adapter for family %q", family)
		}

		// Determine whether Phase 1 has already deployed the replacement OnRamp.
		//
		// If it has, this is either:
		//   - a retry of Phase 1, or
		//   - a later destination-lane batch for the same source chain.
		//
		// In either case we must reuse the existing canonical/legacy pair.
		existingUpgrade, upgradeExists, err := upgrader.ExistingOnRampUpgrade(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("inspect existing OnRamp upgrade: %w", err)
		}

		var oldOnRampBytes []byte
		var existingNewOnRampBytes []byte

		if upgradeExists {
			oldOnRampBytes, err = wireEncodeOnRampRef(localAdapter, existingUpgrade.LegacyOnRampRef, cfg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("encode legacy OnRamp: %w", err)
			}

			existingNewOnRampBytes, err = wireEncodeOnRampRef(localAdapter, existingUpgrade.NewOnRampRef, cfg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("encode existing new OnRamp: %w", err)
			}
		} else {
			// First Phase 1 execution. The canonical OnRamp is still the old one.
			if err := upgrader.VerifyOnrampRequireUpgrade(e, cfg.ChainSelector); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("verify OnRamp requires upgrade: %w", err)
			}

			result, err := localAdapter.GetOnRampAddress(e.DataStore, cfg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("encode old OnRamp: %w", err)
			}

			oldOnRampBytes = result
		}

		// Resolve the source OnRamp's configured destinations before making any
		// mutations. Before the first Phase 1 run this reads the legacy/canonical
		// OnRamp. On a retry it reads the already-deployed canonical new OnRamp.
		destChainSelectors, err := upgrader.DestChainSelectors(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve dest chain selectors: %w", err)
		}

		scopedDestSelectors := scopeDestSelectors(destChainSelectors, cfg.DestSelectorsInScope)

		// Pre-flight the remote OffRamp state before deployment/reconciliation.
		//
		// First execution:
		//   current must contain legacy and may contain only legacy.
		//
		// Retry / later lane batch:
		//   current must contain legacy, and may contain either:
		//     [legacy]
		//   or
		//     [legacy, new]
		//
		// This means a lane whose Phase 1 update already executed is accepted,
		// while a new lane batch that still only knows legacy is also accepted.
		for _, remoteSel := range scopedDestSelectors {
			allowedCurrent := [][]byte{oldOnRampBytes}
			if upgradeExists {
				allowedCurrent = append(allowedCurrent, existingNewOnRampBytes)
			}

			if err := verifyOffRampSourceOnRamps(e, chainFamilyRegistry, remoteSel, cfg.ChainSelector, allowedCurrent, nil, [][]byte{oldOnRampBytes}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"pre-flight OffRamp whitelist check for upgrade: %w",
					err,
				)
			}
		}

		// DeployNewOnRamp is now a reconcile operation:
		//   - first execution deploys;
		//   - retry/later lane batch returns the existing canonical/legacy pair.
		result, err := upgrader.DeployNewOnRamp(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("upgrade OnRamp: %w", err)
		}

		batchOps := append([]mcms_types.BatchOperation{}, result.BatchOps...)

		ds := datastore.NewMemoryDataStore()

		// Always return both refs. On a retry these are the existing refs and the
		// datastore merge is an idempotent upsert.
		if err := ds.Addresses().Add(result.NewOnRampRef); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("add new OnRamp ref to datastore: %w", err)
		}

		if err := ds.Addresses().Add(result.LegacyOnRampRef); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("add legacy OnRamp ref to datastore: %w", err)
		}

		// Ownership transfer is only part of the initial deployment.
		//
		// A Phase 1 retry must not try to transfer ownership again, because the
		// first Phase 1 MCMS operation may already have transferred ownership to
		// the timelock.
		if !result.Reused && !cfg.DisableTransferAndVerifierOwnership {
			ownershipBatchOps, _, errOwnership := deploy.TransferToTimelock(
				cfg.ChainSelector, e, cfg.MCMS, []datastore.AddressRef{result.NewOnRampRef},
				mcmsReaderRegistry, transferOwnershipReg,
			)
			if errOwnership != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("transfer ownership: %w", errOwnership)
			}
			batchOps = append(batchOps, ownershipBatchOps...)
		}

		newOnRampHex, err := wireEncodeOnRampRef(localAdapter, result.NewOnRampRef, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("encode new OnRamp: %w", err)
		}

		// Reconcile every scoped destination to the Phase 1 desired state:
		//
		//   remote OffRamp source whitelist == [legacy, new]
		//
		// SetOffRampSourceOnRamps is expected to skip when this exact state is
		// already present.
		for _, remoteSel := range scopedDestSelectors {
			setter, err := offrampSourceOnRampSetter(chainFamilyRegistry, remoteSel)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			batchOp, skipped, err := setter.SetOffRampSourceOnRamps(e, adapters.OffRampSetSourceOnRampsEntry{
				LocalChainSelector:  remoteSel,
				SourceChainSelector: cfg.ChainSelector,
				OnRamps:             []string{encodeAddressToHex(oldOnRampBytes), encodeAddressToHex(newOnRampHex)},
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

// UpgradeOnrampPhase2 stages ProdRouter lanes behind the TestRouter so the
// replacement OnRamp can be smoke tested before going live. The OnRamp's
// DefaultCCVs remain unchanged (CommitteeVerifier); TESTVTR requires the
// test verifier through its AdvancedPoolHooks configuration.
//
// TestRouter-class lanes are untouched here because TestRouter is already
// their production path.
func UpgradeOnrampPhase2(
	onrampUpgraderRegistry *adapters.OnRampUpgraderRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if err := validateUpgradeOnrampConfig(cfg); err != nil {
			return fmt.Errorf("invalid UpgradeOnrampConfig: %w", err)
		}
		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		family, upgrader, err := initUpgradeOnrampChangesets(onrampUpgraderRegistry, &cfg)
		if err != nil {
			return cldf.ChangesetOutput{}, err
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

		e.OperationsBundle = operations.NewBundle(func() context.Context { return context.Background() }, e.Logger, operations.NewMemoryReporter())
		promoteOps, err := upgrader.PromoteOnrampToTestRouter(e, cfg.ChainSelector, scopedDestSelectors)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("stage on TestRouter: %w", err)
		}

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithBatchOps(promoteOps).
			Build(cfg.MCMS)
	}

	return cldf.CreateChangeSet(apply, validate)
}

// UpgradeOnrampPhase3 promotes the new OnRamp to production for both lane classes:
//   - ProdRouter dests promoted from the
//     TestRouter (where Phase 2 staged them) to the prod Router.
//   - TestRoutee dests are promoted directly to the TestRouter — this is their only
//     promotion, since their production path is the TestRouter itself and they never stage.
//
// Both promotions require the verifier jobs to already observe both OnRamps (Phase 1.5), so
// this phase never causes downtime on either lane class.
func UpgradeOnrampPhase3(
	onrampUpgrader *adapters.OnRampUpgraderRegistry,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if err := validateUpgradeOnrampConfig(cfg); err != nil {
			return fmt.Errorf("invalid UpgradeOnrampConfig: %w", err)
		}

		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		family, upgrader, err := initUpgradeOnrampChangesets(onrampUpgrader, &cfg)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		localAdapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no chain family adapter for family %q", family)
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

		newOnramp, err := localAdapter.GetOnRampAddress(e.DataStore, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("get current OnRamp address: %w", err)
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

		if len(scopedTestDestSelectors) > 0 {
			// Pre-flight: Verify that the offramp source whitelist the new onramp, maybe allowlist the legacy onramp but not needed for this state
			err = ensureNewOnrampIsAllowlistedAndMaybeLegacy(e, upgrader, localAdapter, chainFamilyRegistry, scopedTestDestSelectors, cfg, newOnramp)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight OffRamp whitelist check: %w", err)
			}

			promoteOps, err := upgrader.PromoteOnrampToTestRouter(e, cfg.ChainSelector, scopedTestDestSelectors)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("promote to TestRouter: %w", err)
			}
			batchOps = append(batchOps, promoteOps...)
		}

		if len(scopedProdDestSelectors) > 0 {
			// Pre-flight: Verify that the offramp source whitelist the new onramp, maybe allowlist the legacy onramp but not needed for this state
			err = ensureNewOnrampIsAllowlistedAndMaybeLegacy(e, upgrader, localAdapter, chainFamilyRegistry, scopedProdDestSelectors, cfg, newOnramp)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight OffRamp whitelist check: %w", err)
			}

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

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}

	return cldf.CreateChangeSet(apply, validate)
}

// UpgradeOnrampPhase3Rollback rolls production traffic back to the legacy
// OnRamp for the selected destination lanes.
//
// This is intentionally a traffic-only rollback:
//   - original ProdRouter lanes: ProdRouter -> legacy OnRamp
//   - original TestRouter lanes: TestRouter -> legacy OnRamp
//
// It does NOT:
//   - modify the new OnRamp's dest-chain configs;
//   - change remote OffRamp source-OnRamp allowlists;
//   - change datastore refs;
//   - change verifier jobs;
//   - undo Phase 1.
//
// Pre-flight:
//  1. Phase 3 must currently be active for every selected lane.
//  2. The legacy OnRamp must still be allowlisted on every destination OffRamp.
//
// Therefore this rollback MUST NOT be used after UpgradeOnrampCleanup has
// removed legacy from the remote OffRamps.
func UpgradeOnrampPhase3Rollback(
	onrampUpgraderRegistry *adapters.OnRampUpgraderRegistry,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if err := validateUpgradeOnrampConfig(cfg); err != nil {
			return fmt.Errorf("invalid UpgradeOnrampConfig: %w", err)
		}
		return nil
	}

	apply := func(
		e cldf.Environment,
		cfg UpgradeOnrampConfig,
	) (cldf.ChangesetOutput, error) {
		family, upgrader, err := initUpgradeOnrampChangesets(
			onrampUpgraderRegistry,
			&cfg,
		)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		localAdapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return cldf.ChangesetOutput{},
				fmt.Errorf("no chain family adapter for family %q", family)
		}

		// The legacy ref is both our rollback target and our proof that the
		// upgrade has not been fully cleaned up.
		legacyRef, err := upgrader.LegacyOnRampRef(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve legacy OnRamp: %w", err)
		}

		legacyOnRamp, err := wireEncodeOnRampRef(localAdapter, legacyRef, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("encode legacy OnRamp: %w", err)
		}

		newOnRamp, err := localAdapter.GetOnRampAddress(e.DataStore, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve new OnRamp: %w", err)
		}

		// Classification comes from the legacy OnRamp, whose configs preserve
		// each destination's original production router.
		laneClass, err := upgrader.ClassifyDestChains(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("classify dest chains: %w", err)
		}

		scopedProdDestSelectors := scopeDestSelectors(laneClass.ProdRouterDests, cfg.DestSelectorsInScope)
		scopedTestDestSelectors := scopeDestSelectors(laneClass.TestRouterDests, cfg.DestSelectorsInScope)

		if len(scopedProdDestSelectors) == 0 && len(scopedTestDestSelectors) == 0 {
			return cldf.ChangesetOutput{}, fmt.Errorf("none of the requested destination selectors belong to this OnRamp")
		}

		// Fail closed if Phase 3 is not currently active. This prevents an
		// emergency rollback proposal from overwriting some unexpected router
		// state.
		if err := upgrader.VerifyPromotedToRouters(e, cfg.ChainSelector, scopedProdDestSelectors, scopedTestDestSelectors); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight Phase 3 state check: %w", err)
		}

		scopedDestSelectors := append(
			append(
				make([]uint64, 0,
					len(scopedProdDestSelectors)+len(scopedTestDestSelectors)),
				scopedProdDestSelectors...,
			),
			scopedTestDestSelectors...,
		)

		// The rollback is only safe while remote destinations still accept
		// messages emitted by legacy.
		//
		// Allow only the known Phase-1 pair [legacy,new], and require legacy
		// to actually be present.
		for _, remoteSel := range scopedDestSelectors {
			if err := verifyOffRampSourceOnRamps(e, chainFamilyRegistry, remoteSel, cfg.ChainSelector, [][]byte{legacyOnRamp, newOnRamp}, nil, [][]byte{legacyOnRamp}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"pre-flight legacy OffRamp allowlist check for destination %d: %w",
					remoteSel,
					err,
				)
			}
		}

		// Use a fresh reporter for the rollback writes.
		e.OperationsBundle = operations.NewBundle(func() context.Context { return context.Background() }, e.Logger, operations.NewMemoryReporter())

		rollbackOps, err := upgrader.RollbackToLegacyRouters(e, cfg.ChainSelector, scopedProdDestSelectors, scopedTestDestSelectors)
		if err != nil {
			return cldf.ChangesetOutput{},
				fmt.Errorf("rollback routers to legacy OnRamp: %w", err)
		}

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithBatchOps(rollbackOps).
			Build(cfg.MCMS)
	}

	return cldf.CreateChangeSet(apply, validate)
}

// UpgradeOnrampPhase2Rollback undoes the Phase 2 staging by pointing the
// TestRouter back at the legacy OnRamp for the selected ProdRouter-class lanes.
//
// Phase 2 repoints the TestRouter at the new OnRamp so the replacement can be
// smoke tested. When the new OnRamp cannot deliver traffic — for example when
// its CommitteeVerifier is misconfigured — the TestRouter has no working ramp
// and the lanes on it stop. This changeset restores a working ramp without
// undoing Phase 1.
//
// It takes the same input as Phase 2 and, like UpgradeOnrampPhase3Rollback, it
// is a traffic-only rollback:
//   - original ProdRouter lanes: TestRouter -> legacy OnRamp
//
// It does NOT:
//   - touch the production Router;
//   - touch TestRouter-class lanes, whose production path is the TestRouter;
//   - rewrite the new OnRamp's dest-chain configs (their Router field keeps
//     pointing at the TestRouter, so Phase 2 can be re-applied later);
//   - change remote OffRamp source-OnRamp allowlists;
//   - change datastore refs or verifier jobs.
//
// Pre-flight (both fail closed):
//  1. The TestRouter must currently route every selected lane through the new
//     OnRamp. This rejects lanes that Phase 2 never staged, and lanes that
//     Phase 3 has already promoted to the production Router.
//  2. The legacy OnRamp must still be allowlisted on every destination OffRamp.
//
// Therefore this rollback MUST NOT be used after UpgradeOnrampCleanup has
// removed legacy from the remote OffRamps.
//
// NOTE: the TestRouter is often owned by the deployer key rather than the
// timelock. In that case the router write is signed and sent while this
// changeset applies, and the output carries no MCMS proposal.
func UpgradeOnrampPhase2Rollback(
	onrampUpgraderRegistry *adapters.OnRampUpgraderRegistry,
	chainFamilyRegistry *adapters.ChainFamilyRegistry,
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) cldf.ChangeSetV2[UpgradeOnrampConfig] {
	validate := func(e cldf.Environment, cfg UpgradeOnrampConfig) error {
		if err := validateUpgradeOnrampConfig(cfg); err != nil {
			return fmt.Errorf("invalid UpgradeOnrampConfig: %w", err)
		}
		return nil
	}

	apply := func(e cldf.Environment, cfg UpgradeOnrampConfig) (cldf.ChangesetOutput, error) {
		family, upgrader, err := initUpgradeOnrampChangesets(onrampUpgraderRegistry, &cfg)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		localAdapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no chain family adapter for family %q", family)
		}

		// Classification comes from the legacy OnRamp, whose dest configs keep
		// each destination's original production router. A lane that Phase 2
		// staged therefore still classifies as ProdRouter class.
		laneClass, err := upgrader.ClassifyDestChains(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("classify dest chains: %w", err)
		}

		scopedDestSelectors := scopeDestSelectors(laneClass.ProdRouterDests, cfg.DestSelectorsInScope)
		if len(scopedDestSelectors) == 0 {
			return cldf.ChangesetOutput{}, fmt.Errorf("none of the requested destination selectors is a ProdRouter-class lane of this OnRamp")
		}

		// The legacy ref is both our rollback target and our proof that the
		// upgrade has not been cleaned up.
		legacyRef, err := upgrader.LegacyOnRampRef(e, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve legacy OnRamp: %w", err)
		}

		legacyOnRamp, err := wireEncodeOnRampRef(localAdapter, legacyRef, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("encode legacy OnRamp: %w", err)
		}

		newOnRamp, err := localAdapter.GetOnRampAddress(e.DataStore, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve new OnRamp: %w", err)
		}

		// Fail closed unless Phase 2 is currently active on the TestRouter for
		// every selected lane. Passing no prod selectors makes the production
		// Router check a no-op.
		if err := upgrader.VerifyPromotedToRouters(e, cfg.ChainSelector, nil, scopedDestSelectors); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight Phase 2 state check: %w (Phase 2 must be active on the TestRouter and Phase 3 must not have run for these lanes)", err)
		}

		// The rollback is only safe while remote destinations still accept
		// messages that legacy emits.
		//
		// Allow only the known Phase 1 pair [legacy,new], and require legacy to
		// actually be present.
		for _, remoteSel := range scopedDestSelectors {
			if err := verifyOffRampSourceOnRamps(e, chainFamilyRegistry, remoteSel, cfg.ChainSelector, [][]byte{legacyOnRamp, newOnRamp}, nil, [][]byte{legacyOnRamp}); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"pre-flight legacy OffRamp allowlist check for destination %d: %w",
					remoteSel,
					err,
				)
			}
		}

		// Use a fresh reporter for the rollback writes.
		e.OperationsBundle = operations.NewBundle(func() context.Context { return context.Background() }, e.Logger, operations.NewMemoryReporter())

		// The scoped lanes are ProdRouter class, but they are passed as the test
		// selectors because the TestRouter is what Phase 2 repointed and what
		// this rollback must restore. No production Router write is produced.
		rollbackOps, err := upgrader.RollbackToLegacyRouters(e, cfg.ChainSelector, nil, scopedDestSelectors)
		if err != nil {
			return cldf.ChangesetOutput{},
				fmt.Errorf("rollback TestRouter to legacy OnRamp: %w", err)
		}

		// A deployer-owned TestRouter executes the write during apply, which
		// leaves no MCMS transaction to propose.
		if len(rollbackOps) == 0 {
			e.Logger.Infow("TestRouter rollback to the legacy OnRamp produced no MCMS transaction; the router write already executed with the deployer key",
				"chain", cfg.ChainSelector, "destChains", scopedDestSelectors)

			return cldf.ChangesetOutput{}, nil
		}

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithBatchOps(rollbackOps).
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
		family, upgrader, err := initUpgradeOnrampChangesets(onrampUpgraderRegistry, &cfg)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		chainAdapter, ok := chainFamilyRegistry.GetChainFamily(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no adapter for family %q", family)
		}

		newOnRampAddr, err := chainAdapter.GetOnRampAddress(e.DataStore, cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("resolve new OnRamp: %w", err)
		}

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
		scopedDestSelectors := scopeDestSelectors(destChainSelectors, cfg.DestSelectorsInScope)
		// Pre-flight: verify that the new OnRamp is already allowlisted on all remote OffRamps, and that the legacy OnRamp is either allowlisted or not.
		if err := ensureNewOnrampIsAllowlistedAndMaybeLegacy(e, upgrader, chainAdapter, chainFamilyRegistry, scopedDestSelectors, cfg, newOnRampAddr); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("pre-flight OffRamp whitelist check: %w", err)
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
				OnRamps:             []string{encodeAddressToHex(newOnRampAddr)},
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("disallow old OnRamp on OffRamp for remote %d: %w", remoteSel, err)
			}
			if !skipped && batchOp != nil {
				batchOps = append(batchOps, *batchOp)
			}
		}

		e.Logger.Infow("OnRamp upgrade cleanup complete. If and only if no other lanes are left for that chain operators must manually: (1) remove the legacy OnRamp (qualifier=\"legacy\") from the datastore, (2) tell NOPs to remove temp verifier jobs",
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

// ensureNewOnrampIsAllowlistedAndMaybeLegacy verifies that the new OnRamp is allowlisted on all remote OffRamps, and that the legacy OnRamp is either allowlisted or not. This guards against silently dropping an onramp that still serves traffic.
func ensureNewOnrampIsAllowlistedAndMaybeLegacy(e cldf.Environment, upgrader adapters.OnRampUpgrader, chainAdapter adapters.ChainFamily, chainFamilyRegistry *adapters.ChainFamilyRegistry, scopedDestSelectors []uint64, cfg UpgradeOnrampConfig, newOnRampAddr []byte) error {
	var expectedRemovals [][]byte
	legacyRef, err := upgrader.LegacyOnRampRef(e, cfg.ChainSelector)
	if err != nil {
		e.Logger.Infow("No legacy OnRamp ref found; only the new OnRamp is expected on remote OffRamps",
			"chain", cfg.ChainSelector, "err", err)
	} else {
		legacyOnRampHex, err := wireEncodeOnRampRef(chainAdapter, legacyRef, cfg.ChainSelector)
		if err != nil {
			return fmt.Errorf("encode legacy OnRamp: %w", err)
		}
		expectedRemovals = [][]byte{legacyOnRampHex}
	}

	// Pre-flight: verify no remote OffRamp whitelists an unknown onramp that the
	// update below would silently drop, and that the new OnRamp is actually
	// whitelisted before we promote traffic to it.
	for _, remoteSel := range scopedDestSelectors {
		if err := verifyOffRampSourceOnRamps(e, chainFamilyRegistry, remoteSel, cfg.ChainSelector,
			[][]byte{newOnRampAddr}, expectedRemovals, [][]byte{newOnRampAddr}); err != nil {
			return fmt.Errorf("pre-flight OffRamp whitelist check: %w", err)
		}
	}
	return nil
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

// wireEncodeOnRampRef returns the wire-encoded OnRamp ref, using a
// temp datastore so the family adapter resolves it as the canonical OnRamp.
func wireEncodeOnRampRef(adapter adapters.ChainFamily, ref datastore.AddressRef, chainSelector uint64) ([]byte, error) {
	ref.Qualifier = ""
	tmp := datastore.NewMemoryDataStore()
	if err := tmp.Addresses().Add(ref); err != nil {
		return nil, fmt.Errorf("add OnRamp ref to temp datastore: %w", err)
	}
	onRampBytes, err := adapter.GetOnRampAddress(tmp.Seal(), chainSelector)
	if err != nil {
		return nil, fmt.Errorf("encode OnRamp %s: %w", ref.Address, err)
	}
	return onRampBytes, nil
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
	desired [][]byte,
	expectedRemovals [][]byte,
	mustContain [][]byte,
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
		allowed[hex.EncodeToString(addr)] = struct{}{}
	}
	for _, addr := range expectedRemovals {
		allowed[hex.EncodeToString(addr)] = struct{}{}
	}

	currentSet := make(map[string]struct{}, len(current))
	var unexpected []string
	for _, addr := range current {
		key := hex.EncodeToString(addr)
		currentSet[key] = struct{}{}
		if _, ok := allowed[key]; !ok {
			unexpected = append(unexpected, key)
		}
	}
	if len(unexpected) > 0 {
		return fmt.Errorf("OffRamp on chain %d whitelists unexpected onramps %v for source chain %d; "+
			"refusing to override them — investigate and use the OffRampSetSourceOnRamps changeset to update deliberately",
			remoteSel, unexpected, sourceSel)
	}
	for _, addr := range mustContain {
		key := hex.EncodeToString(addr)
		if _, ok := currentSet[key]; !ok {
			return fmt.Errorf("OffRamp on chain %d does not whitelist expected onramp %s for source chain %d (current: %v)",
				remoteSel, key, sourceSel, current)
		}
	}
	return nil
}
