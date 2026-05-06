package authorizedcallers

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// Config is the input to ConfigureAuthorizedCallersChangeset.
// Each entry in Updates targets a specific (chain, contract type, version) triple.
// Multiple entries for different contract types on the same chain are supported and
// will produce separate batch operations — one per target contract.
type Config struct {
	Updates []ApplyInput `json:"updates" yaml:"updates"`
	// Force skips the pre-write idempotency filter when true. When false, adds that
	// are already present and removes that are already absent are silently dropped.
	Force  bool       `json:"force"  yaml:"force"`
	Reason string     `json:"reason" yaml:"reason"`
	MCMS   mcms.Input `json:"mcms"   yaml:"mcms"`
}

// ConfigureAuthorizedCallersChangeset returns a changeset that applies
// applyAuthorizedCallerUpdates on one or more AuthorizedCallers-inheriting contracts
// across any registered chain family. It is the counterpart of fastcurse.CurseChangeset.
func ConfigureAuthorizedCallersChangeset(
	reg *AuthorizedCallersRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) cldf.ChangeSetV2[Config] {
	return cldf.CreateChangeSet(applyAuthorizedCallers(reg, mcmsRegistry), verifyAuthorizedCallersInput(reg))
}

func verifyAuthorizedCallersInput(reg *AuthorizedCallersRegistry) func(cldf.Environment, Config) error {
	return func(e cldf.Environment, cfg Config) error {
		for _, in := range cfg.Updates {
			family, err := chain_selectors.GetSelectorFamily(in.ChainSelector)
			if err != nil {
				return fmt.Errorf("failed to get chain family for selector %d: %w", in.ChainSelector, err)
			}
			if _, ok := reg.GetAdapter(family, in.ContractType, in.Version); !ok {
				return fmt.Errorf("no authorized callers adapter registered for chain family %q, contract type %q, version %s",
					family, in.ContractType, in.Version.String())
			}
		}
		return nil
	}
}

func applyAuthorizedCallers(
	reg *AuthorizedCallersRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, Config) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg Config) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Group updates by (family, contractType, version, chainSelector). Multiple
		// AuthorizedCallers-inheriting contracts can exist on the same chain (e.g. RMN
		// and FeeQuoter), so the grouping key must include the contract identity.
		type groupKey struct {
			adapterKey    string
			chainSelector uint64
		}
		type groupDetail struct {
			adapter AuthorizedCallersAdapter
			input   ApplyInput
		}

		grouped := make(map[groupKey]groupDetail)
		// Preserve insertion order so the sequence of batch ops is deterministic.
		order := make([]groupKey, 0, len(cfg.Updates))

		for _, in := range cfg.Updates {
			family, err := chain_selectors.GetSelectorFamily(in.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", in.ChainSelector, err)
			}
			adapter, ok := reg.GetAdapter(family, in.ContractType, in.Version)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"no authorized callers adapter registered for chain family %q, contract type %q, version %s",
					family, in.ContractType, in.Version.String())
			}
			if err := adapter.Initialize(e, in); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"failed to initialize authorized callers adapter for chain %d, contract %q v%s: %w",
					in.ChainSelector, in.ContractType, in.Version.String(), err)
			}

			key := groupKey{
				adapterKey:    adapterKey(family, in.ContractType, in.Version),
				chainSelector: in.ChainSelector,
			}
			if _, exists := grouped[key]; !exists {
				order = append(order, key)
			}
			grouped[key] = groupDetail{adapter: adapter, input: in}
		}

		for _, key := range order {
			detail := grouped[key]
			effectiveInput := detail.input

			if !cfg.Force {
				filtered, err := filterCallerUpdate(e, detail.adapter, detail.input)
				if err != nil {
					return cldf.ChangesetOutput{}, err
				}
				if len(filtered.AddedCallers) == 0 && len(filtered.RemovedCallers) == 0 {
					e.Logger.Infof(
						"No-op: authorized caller updates already applied on chain %d for contract %q v%s, skipping",
						key.chainSelector, detail.input.ContractType, detail.input.Version.String(),
					)
					continue
				}
				effectiveInput.Update = filtered
			}

			e.Logger.Infof(
				"Applying authorized caller updates on chain %d for contract %q v%s: +%d -%d callers",
				key.chainSelector, effectiveInput.ContractType, effectiveInput.Version.String(),
				len(effectiveInput.Update.AddedCallers), len(effectiveInput.Update.RemovedCallers),
			)

			seqReport, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				detail.adapter.ApplyAuthorizedCallerUpdates(),
				e.BlockChains,
				effectiveInput,
			)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf(
					"failed to apply authorized caller updates on chain %d for contract %q v%s: %w",
					key.chainSelector, effectiveInput.ContractType, effectiveInput.Version.String(), err)
			}
			batchOps = append(batchOps, seqReport.Output.BatchOps...)
			reports = append(reports, seqReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

// filterCallerUpdate reads the current on-chain caller set and removes
// already-present adds and already-absent removes to keep the apply path idempotent.
func filterCallerUpdate(e cldf.Environment, adapter AuthorizedCallersAdapter, in ApplyInput) (CallerUpdate, error) {
	current, err := adapter.GetAllAuthorizedCallers(e, in.ChainSelector, in.ContractType, in.Version)
	if err != nil {
		return CallerUpdate{}, fmt.Errorf(
			"failed to read current authorized callers on chain %d for contract %q v%s: %w",
			in.ChainSelector, in.ContractType, in.Version.String(), err)
	}

	currentSet := make(map[string]struct{}, len(current))
	for _, c := range current {
		currentSet[string(c)] = struct{}{}
	}

	added := make([]Caller, 0, len(in.Update.AddedCallers))
	for _, c := range in.Update.AddedCallers {
		if _, exists := currentSet[string(c)]; !exists {
			added = append(added, c)
		}
	}

	removed := make([]Caller, 0, len(in.Update.RemovedCallers))
	for _, c := range in.Update.RemovedCallers {
		if _, exists := currentSet[string(c)]; exists {
			removed = append(removed, c)
		}
	}

	return CallerUpdate{AddedCallers: added, RemovedCallers: removed}, nil
}
