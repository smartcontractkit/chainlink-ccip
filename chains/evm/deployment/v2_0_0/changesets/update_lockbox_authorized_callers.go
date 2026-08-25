package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	lbops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// LockboxCallerUpdate is one chain's authorized-caller change for one lockbox.
//
// The lockbox is named by Address rather than by datastore qualifier. Qualifiers are a
// side-effect of whichever sequence deployed the lockbox - the pool qualifier for a plain
// lock-release pool, "<poolQualifier>-silo(<minRemoteSelector>)" per silo group, or
// "remoteChainSelector(<selector>)" for siloed USDC - so they are not a stable addressing
// scheme. Address is unambiguous, and verify confirms it really is a lockbox.
type LockboxCallerUpdate struct {
	Selector uint64           `json:"selector" yaml:"selector"`
	Address  common.Address   `json:"address"  yaml:"address"`
	Appends  []common.Address `json:"appends,omitempty"  yaml:"appends,omitempty"`
	Removes  []common.Address `json:"removes,omitempty"  yaml:"removes,omitempty"`
}

type UpdateLockboxAuthorizedCallersCfg struct {
	// Version is cross-checked against the datastore ref for every Address, so a config
	// written for one lockbox version cannot silently act on a newer one.
	Version *semver.Version       `json:"version" yaml:"version"`
	Input   []LockboxCallerUpdate `json:"input"   yaml:"input"`
}

var UpdateLockboxAuthorizedCallers = func(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
) cldf_deployment.ChangeSetV2[cs_changesets.WithMCMS[UpdateLockboxAuthorizedCallersCfg]] {
	return cldf_deployment.CreateChangeSet(
		applyUpdateLockboxAuthorizedCallers(mcmsRegistry),
		verifyUpdateLockboxAuthorizedCallers,
	)
}

func verifyUpdateLockboxAuthorizedCallers(
	e cldf_deployment.Environment,
	input cs_changesets.WithMCMS[UpdateLockboxAuthorizedCallersCfg],
) error {
	cfg := input.Cfg
	if len(cfg.Input) == 0 {
		return fmt.Errorf("at least one entry is required in input")
	}
	if cfg.Version == nil {
		return fmt.Errorf("version is required so the target lockbox version can be cross-checked")
	}

	evmChains := e.BlockChains.EVMChains()
	// Keyed on chain AND address, not chain alone: a chain legitimately holds several
	// lockboxes (one per token, per silo group, or per remote chain), so several entries
	// may share a selector. Only the same lockbox twice is a mistake.
	type lockboxKey struct {
		selector uint64
		address  common.Address
	}
	seen := common_utils.NewSet[lockboxKey]()

	for _, update := range cfg.Input {
		sel := update.Selector
		if _, err := chain_selectors.GetSelectorFamily(sel); err != nil {
			return fmt.Errorf("invalid chain selector %d: %w", sel, err)
		}
		if _, ok := evmChains[sel]; !ok {
			return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
		}
		if update.Address == (common.Address{}) {
			return fmt.Errorf("zero lockbox address for chain %d", sel)
		}
		if seen.Add(lockboxKey{selector: sel, address: update.Address}) {
			return fmt.Errorf(
				"duplicate entry for lockbox %s on chain %d: merge them into a single entry",
				update.Address, sel)
		}

		appends := common_utils.NewSet[common.Address]()
		for _, addr := range update.Appends {
			if addr == (common.Address{}) {
				return fmt.Errorf("zero address in appends for chain %d", sel)
			}
			if appends.Add(addr) {
				return fmt.Errorf("duplicate address %s in appends for chain %d", addr, sel)
			}
		}

		removes := common_utils.NewSet[common.Address]()
		for _, addr := range update.Removes {
			if addr == (common.Address{}) {
				return fmt.Errorf("zero address in removes for chain %d", sel)
			}
			if removes.Add(addr) {
				return fmt.Errorf("duplicate address %s in removes for chain %d", addr, sel)
			}
			// An address in both lists never converges: the idempotency filter drops
			// whichever side matches current on-chain state, so each apply flips it.
			if appends.Has(addr) {
				return fmt.Errorf(
					"address %s is in both appends and removes for chain %d", addr, sel)
			}
		}
	}

	// Every target is confirmed before any chain is touched, so a bad entry cannot leave
	// earlier chains half-applied.
	if err := checkLockboxRefs(e.DataStore, cfg); err != nil {
		return err
	}

	// Validated against a populated copy so ValidUntil stays optional while a value the
	// operator did supply is still rejected if it is unusable. The qualifier needs no
	// default here: EVMMCMSReader.GetTimelockRef and GetMCMSRef already fall back to
	// CLLQualifier when it is empty.
	mcmsInput := input.MCMS
	if err := mcmsInput.PopulateDefaults(); err != nil {
		return fmt.Errorf("failed to populate MCMS defaults: %w", err)
	}
	if err := mcmsInput.Validate(); err != nil {
		return fmt.Errorf("invalid MCMS input: %w", err)
	}

	return nil
}

// checkLockboxRefs confirms every Address resolves in the datastore to an ERC20LockBox at
// the configured version on the configured chain.
func checkLockboxRefs(ds datastore.DataStore, cfg UpdateLockboxAuthorizedCallersCfg) error {
	for _, update := range cfg.Input {
		refs := ds.Addresses().Filter(datastore_utils.AddressRefToFilters(datastore.AddressRef{
			ChainSelector: update.Selector,
			Address:       update.Address.Hex(),
		})...)
		if len(refs) != 1 {
			return fmt.Errorf(
				"expected exactly 1 datastore ref for address %s on chain %d, found %d",
				update.Address, update.Selector, len(refs))
		}

		ref := refs[0]
		if ref.Type != datastore.ContractType(lbops.ContractType) {
			return fmt.Errorf(
				"address %s on chain %d is a %q, not an %s",
				update.Address, update.Selector, ref.Type, lbops.ContractType)
		}
		if ref.Version == nil || !ref.Version.Equal(cfg.Version) {
			return fmt.Errorf(
				"lockbox %s on chain %d is version %s, but config specifies %s",
				update.Address, update.Selector, versionString(ref.Version), cfg.Version)
		}
	}

	return nil
}

func versionString(v *semver.Version) string {
	if v == nil {
		return "<nil>"
	}

	return v.String()
}

func applyUpdateLockboxAuthorizedCallers(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
) func(cldf_deployment.Environment, cs_changesets.WithMCMS[UpdateLockboxAuthorizedCallersCfg]) (cldf_deployment.ChangesetOutput, error) {
	return func(
		e cldf_deployment.Environment,
		input cs_changesets.WithMCMS[UpdateLockboxAuthorizedCallersCfg],
	) (cldf_deployment.ChangesetOutput, error) {
		cfg := input.Cfg
		if cfg.Version == nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("version is required")
		}

		// Re-checked here because Apply is callable without VerifyPreconditions, and every
		// target must be known good before the first write goes out.
		if err := checkLockboxRefs(e.DataStore, cfg); err != nil {
			return cldf_deployment.ChangesetOutput{}, err
		}

		batchOps := make([]mcms_types.BatchOperation, 0, len(cfg.Input))
		reports := make([]cldf_ops.Report[any, any], 0, len(cfg.Input))
		evmChains := e.BlockChains.EVMChains()

		// Input order is the operator's order, so batch operation order - and therefore the
		// MCMS proposal's merkle root - is reproducible for a given config.
		for _, update := range cfg.Input {
			sel := update.Selector
			chain, ok := evmChains[sel]
			if !ok {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"chain selector %d not found in environment EVM chains", sel)
			}
			lockboxAddr := update.Address

			// Read current authorized callers with a fresh bundle so idempotency
			// reads are not served from a cached report.
			readBundle := cldf_ops.NewBundle(
				e.GetContext, e.Logger, cldf_ops.NewMemoryReporter())
			currentReport, err := cldf_ops.ExecuteOperation(
				readBundle, lbops.GetAllAuthorizedCallers, chain,
				contract.FunctionInput[struct{}]{
					ChainSelector: chain.Selector,
					Address:       lockboxAddr,
					Args:          struct{}{},
				})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"failed to read authorized callers on chain %d: %w", sel, err)
			}

			currentSet := common_utils.NewSet[common.Address]()
			for _, c := range currentReport.Output {
				currentSet.Add(c)
			}

			filteredAppends := make([]common.Address, 0, len(update.Appends))
			for _, a := range update.Appends {
				if !currentSet.Has(a) {
					filteredAppends = append(filteredAppends, a)
				}
			}

			filteredRemoves := make([]common.Address, 0, len(update.Removes))
			for _, r := range update.Removes {
				if currentSet.Has(r) {
					filteredRemoves = append(filteredRemoves, r)
				}
			}

			if len(filteredAppends) == 0 && len(filteredRemoves) == 0 {
				e.Logger.Infof(
					"No-op: authorized caller updates already applied on lockbox %s on chain %d, skipping",
					lockboxAddr, sel)

				continue
			}

			e.Logger.Infof(
				"Applying authorized caller updates on lockbox %s on chain %d: +%d -%d callers",
				lockboxAddr, sel, len(filteredAppends), len(filteredRemoves))

			report, err := cldf_ops.ExecuteOperation(
				e.OperationsBundle, lbops.ApplyAuthorizedCallerUpdates, chain,
				contract.FunctionInput[lbops.AuthorizedCallerArgs]{
					ChainSelector: chain.Selector,
					Address:       lockboxAddr,
					Args: lbops.AuthorizedCallerArgs{
						AddedCallers:   filteredAppends,
						RemovedCallers: filteredRemoves,
					},
				},
				// The shared bundle replays a cached report for an identical operation and
				// input. The reads above dodge that with a fresh bundle; the write needs
				// the same guarantee, or a re-apply after an out-of-band change would
				// return the stale executed write and report success having done nothing.
				// applyAuthorizedCallerUpdates is a set operation on-chain, so re-running
				// it is harmless.
				cldf_ops.WithForceExecute[contract.FunctionInput[lbops.AuthorizedCallerArgs], evm.Chain](),
			)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"failed to apply authorized caller updates on chain %d: %w", sel, err)
			}

			reports = append(reports, report.ToGenericReport())

			batch, err := contract.NewBatchOperationFromWrites(
				[]contract.WriteOutput{report.Output})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"failed to create batch from writes on chain %d: %w", sel, err)
			}
			batchOps = append(batchOps, batch)
		}

		// SingleBatchOpPerChain, not BatchOps: a chain can have several lockboxes updated
		// in one run, and MCMS expects all operations for a chain in one batch operation so
		// they execute atomically rather than as separately-executable units.
		return cs_changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithSingleBatchOpPerChain(batchOps).
			Build(input.MCMS)
	}
}
