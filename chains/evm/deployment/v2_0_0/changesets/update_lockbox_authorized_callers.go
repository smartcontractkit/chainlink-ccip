package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	lbops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

// LockboxCallerUpdate is one lockbox's authorized-caller change. The chain it lives on comes
// from the enclosing ChainLockboxUpdate.
//
// The lockbox is named by Address rather than by datastore qualifier. Qualifiers are a
// side-effect of whichever sequence deployed the lockbox - the pool qualifier for a plain
// lock-release pool, "<poolQualifier>-silo(<minRemoteSelector>)" per silo group, or
// "remoteChainSelector(<selector>)" for siloed USDC - so they are not a stable addressing
// scheme. Address is unambiguous. A wrong address fails on-chain when the read or write
// reverts, rather than being pre-checked against the datastore - datastore refs go missing
// routinely, and blocking on one would force operators to hand-add refs just to proceed.
type LockboxCallerUpdate struct {
	Address common.Address   `json:"address"  yaml:"address"`
	Appends []common.Address `json:"appends,omitempty"  yaml:"appends,omitempty"`
	Removes []common.Address `json:"removes,omitempty"  yaml:"removes,omitempty"`
}

// ChainLockboxUpdate groups every lockbox update for one chain under that chain's selector.
// A chain legitimately holds several lockboxes - one per token, per silo group, or per
// remote chain - and nesting them keeps the selector written once per chain.
type ChainLockboxUpdate struct {
	Selector  uint64                `json:"selector"  yaml:"selector"`
	Lockboxes []LockboxCallerUpdate `json:"lockboxes" yaml:"lockboxes"`
}

type UpdateLockboxAuthorizedCallersCfg struct {
	// Version is not a field: this package is v2_0_0, so the target lockbox version is
	// already pinned by the changeset the operator picked.
	Input []ChainLockboxUpdate `json:"input" yaml:"input"`
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

	evmChains := e.BlockChains.EVMChains()
	seenChains := common_utils.NewSet[uint64]()

	for _, chainUpdate := range cfg.Input {
		sel := chainUpdate.Selector
		if _, err := chain_selectors.GetSelectorFamily(sel); err != nil {
			return fmt.Errorf("invalid chain selector %d: %w", sel, err)
		}
		if _, ok := evmChains[sel]; !ok {
			return fmt.Errorf("chain selector %d not found in environment EVM chains", sel)
		}
		// Each chain appears once and carries all of its lockboxes, so the nesting - and
		// therefore the single batch operation built per chain in apply - is unambiguous.
		if seenChains.Add(sel) {
			return fmt.Errorf(
				"duplicate entry for chain %d: merge its lockboxes into a single entry", sel)
		}
		if len(chainUpdate.Lockboxes) == 0 {
			return fmt.Errorf("no lockboxes listed for chain %d", sel)
		}

		// Scoped to the chain: the same lockbox address twice on one chain is a mistake,
		// but the same address on two chains is a coincidence, not an error.
		seenLockboxes := common_utils.NewSet[common.Address]()

		for _, update := range chainUpdate.Lockboxes {
			if update.Address == (common.Address{}) {
				return fmt.Errorf("zero lockbox address for chain %d", sel)
			}
			if seenLockboxes.Add(update.Address) {
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
	}

	return nil
}

func applyUpdateLockboxAuthorizedCallers(
	mcmsRegistry *cs_changesets.MCMSReaderRegistry,
) func(cldf_deployment.Environment, cs_changesets.WithMCMS[UpdateLockboxAuthorizedCallersCfg]) (cldf_deployment.ChangesetOutput, error) {
	return func(
		e cldf_deployment.Environment,
		input cs_changesets.WithMCMS[UpdateLockboxAuthorizedCallersCfg],
	) (cldf_deployment.ChangesetOutput, error) {
		cfg := input.Cfg

		batchOps := make([]mcms_types.BatchOperation, 0, len(cfg.Input))
		reports := make([]cldf_ops.Report[any, any], 0, len(cfg.Input))
		evmChains := e.BlockChains.EVMChains()

		// Input order is the operator's order, so batch operation order - and therefore the
		// MCMS proposal's merkle root - is reproducible for a given config.
		for _, chainUpdate := range cfg.Input {
			sel := chainUpdate.Selector
			chain, ok := evmChains[sel]
			if !ok {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"chain selector %d not found in environment EVM chains", sel)
			}

			writes := make([]contract.WriteOutput, 0, len(chainUpdate.Lockboxes))

			for _, update := range chainUpdate.Lockboxes {
				lockboxAddr := update.Address

				// WithForceExecute so the idempotency read is never served from a cached
				// report: a re-run after an out-of-band change has to see current state.
				currentReport, err := cldf_ops.ExecuteOperation(
					e.OperationsBundle, lbops.GetAllAuthorizedCallers, chain,
					contract.FunctionInput[struct{}]{
						ChainSelector: chain.Selector,
						Address:       lockboxAddr,
						Args:          struct{}{},
					},
					cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
				)
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
					// Same reason as the read above: the shared bundle would replay a cached
					// report for an identical operation and input, so a re-apply after an
					// out-of-band change would return the stale executed write and report
					// success having done nothing. applyAuthorizedCallerUpdates is a set
					// operation on-chain, so re-running it is harmless.
					cldf_ops.WithForceExecute[contract.FunctionInput[lbops.AuthorizedCallerArgs], evm.Chain](),
				)
				if err != nil {
					return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
						"failed to apply authorized caller updates on chain %d: %w", sel, err)
				}

				reports = append(reports, report.ToGenericReport())
				writes = append(writes, report.Output)
			}

			// One batch operation per chain, built from that chain's writes: MCMS executes a
			// batch operation atomically, so a chain whose lockboxes are updated together
			// cannot be left half-applied. NewBatchOperationFromWrites rejects writes that
			// span chains, which is exactly why it is called once per chain here.
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf(
					"failed to create batch from writes on chain %d: %w", sel, err)
			}
			batchOps = append(batchOps, batch)
		}

		// BatchOps, not SingleBatchOpPerChain: the per-chain merge is structural now, since
		// batchOps already holds at most one batch operation per chain. Empty batches - a
		// chain whose updates were all filtered as no-ops, or whose writes executed
		// directly - are dropped by WithBatchOps.
		return cs_changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}
