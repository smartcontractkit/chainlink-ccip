package fastcurse

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type GlobalCurseOnNetworkInput struct {
	ChainSelectors map[uint64]*semver.Version
	MCMS           mcms.Input
}

type RMNCurseConfig struct {
	CurseActions []CurseActionInput
	// Use this if you want to include curse subject even when they are already cursed (CurseChangeset) or already uncursed (UncurseChangeset)
	Force bool
	// MCMS configures the resulting proposal.
	MCMS mcms.Input
}

// CurseActionInput represent a curse action to be applied on a chain (ChainSelector) with a specific SubjectToCurse derived from the SubjectChainSelector
// The curse action will by applied by calling the Curse method on the RMNRemote contract on the chain (ChainSelector)
type CurseActionInput struct {
	IsGlobalCurse        bool
	ChainSelector        uint64
	SubjectChainSelector uint64
	Version              *semver.Version
}

type curseActionDetails struct {
	curseAdapter CurseAdapter
	subjects     []Subject
}

func GloballyCurseChainChangeset(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[GlobalCurseOnNetworkInput] {
	return cldf.CreateChangeSet(applyGlobalCurseOnNetwork(cr, mcmsRegistry, true), verifyGlobalCurseOnNetworkInput(cr, mcmsRegistry))
}

func GloballyUncurseChainChangeset(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[GlobalCurseOnNetworkInput] {
	return cldf.CreateChangeSet(applyGlobalCurseOnNetwork(cr, mcmsRegistry, false), verifyGlobalCurseOnNetworkInput(cr, mcmsRegistry))
}

func CurseChangeset(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[RMNCurseConfig] {
	return cldf.CreateChangeSet(applyCurse(cr, mcmsRegistry), verifyCurseInput(cr, mcmsRegistry))
}

func UncurseChangeset(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[RMNCurseConfig] {
	return cldf.CreateChangeSet(applyUncurse(cr, mcmsRegistry), verifyCurseInput(cr, mcmsRegistry))
}

func verifyGlobalCurseOnNetworkInput(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, GlobalCurseOnNetworkInput) error {
	return func(e cldf.Environment, cfg GlobalCurseOnNetworkInput) error {
		return nil
	}
}

func formCurseConfigForGlobalCurse(e cldf.Environment, cr *CurseRegistry, cfg GlobalCurseOnNetworkInput) (RMNCurseConfig, error) {
	// form the curse input for each chain selector
	curseCfg := RMNCurseConfig{
		CurseActions: make([]CurseActionInput, 0),
		MCMS:         cfg.MCMS,
	}
	for chainSelector, version := range cfg.ChainSelectors {
		curseAction := CurseActionInput{
			IsGlobalCurse: true,
			ChainSelector: chainSelector,
			Version:       version,
		}
		curseCfg.CurseActions = append(curseCfg.CurseActions, curseAction)
		// get adapter
		family, err := chain_selectors.GetSelectorFamily(chainSelector)
		if err != nil {
			return curseCfg, err
		}
		adapter, ok := cr.GetCurseAdapter(family, version)
		if !ok {
			return curseCfg, fmt.Errorf("no curse adapter registered for chain family '%s' and RMN version '%s'",
				family, version.String())
		}
		err = adapter.Initialize(e, chainSelector)
		if err != nil {
			return RMNCurseConfig{}, err
		}
		connectedChains, err := adapter.ListConnectedChains(e, chainSelector)
		if err != nil {
			return curseCfg, fmt.Errorf("failed to list connected chains for chain selector %d: %w", chainSelector, err)
		}
		for _, connectedChainSelector := range connectedChains {
			connectedChainFamily, err := chain_selectors.GetSelectorFamily(connectedChainSelector)
			if err != nil {
				return curseCfg, err
			}
			connectedAdapter, ok := cr.GetCurseSubjectAdapter(connectedChainFamily)
			if !ok {
				return curseCfg, fmt.Errorf("no curse subject adapter registered for chain family '%s'",
					connectedChainFamily)
			}
			connectedVersion, err := connectedAdapter.DeriveCurseAdapterVersion(e, connectedChainSelector)
			if err != nil {
				return curseCfg, fmt.Errorf("failed to derive curse adapter version for chain selector %d: %w", connectedChainSelector, err)
			}
			curseCfg.CurseActions = append(curseCfg.CurseActions, CurseActionInput{
				ChainSelector:        connectedChainSelector,
				Version:              connectedVersion,
				SubjectChainSelector: chainSelector,
			})
		}
	}
	return curseCfg, nil
}

func applyGlobalCurseOnNetwork(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry, curse bool) func(cldf.Environment, GlobalCurseOnNetworkInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg GlobalCurseOnNetworkInput) (cldf.ChangesetOutput, error) {
		// form the curse input for each chain selector
		curseCfg, err := formCurseConfigForGlobalCurse(e, cr, cfg)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to form curse config for global curse: %w", err)
		}
		if curse {
			return applyCurse(cr, mcmsRegistry)(e, curseCfg)
		}
		return applyUncurse(cr, mcmsRegistry)(e, curseCfg)
	}
}

func verifyCurseInput(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, RMNCurseConfig) error {
	return func(e cldf.Environment, cfg RMNCurseConfig) error {
		return nil
	}
}

func applyCurse(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, RMNCurseConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RMNCurseConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		// Group curse actions by chain selector
		grouped, err := cr.groupRMNSubjectBySelector(e, cfg.CurseActions)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to group curse actions: %w", err)
		}
		for selector, curseDetail := range grouped {
			adapter := curseDetail.curseAdapter
			subjects := curseDetail.subjects
			notAlreadyCursedSubjects := make([]Subject, 0)
			for _, subject := range subjects {
				// Only curse the subjects that are not actually cursed
				cursed, err := adapter.IsSubjectCursedOnChain(e, selector, subject)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to check if subject %x is cursed on chain with selector %d: %w", subject, selector, err)
				}
				if cursed && !cfg.Force {
					e.Logger.Infof("Subject %x is already cursed on chain with selector %d, skipping", subject, selector)
					continue
				}
				notAlreadyCursedSubjects = append(notAlreadyCursedSubjects, subject)
			}
			if len(notAlreadyCursedSubjects) == 0 {
				e.Logger.Infof("No new subjects to curse on chain with selector %d, skipping", selector)
				continue
			}
			e.Logger.Infof("Cursing %d subjects on chain with selector %d", len(notAlreadyCursedSubjects), selector)
			curseReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.Curse(), e.BlockChains, CurseInput{
				Subjects:      notAlreadyCursedSubjects,
				ChainSelector: selector,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to curse subjects on chain with selector %d: %w", selector, err)
			}
			batchOps = append(batchOps, curseReport.Output.BatchOps...)
			reports = append(reports, curseReport.ExecutionReports...)
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

func applyUncurse(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, RMNCurseConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RMNCurseConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		// Group curse actions by chain selector
		grouped, err := cr.groupRMNSubjectBySelector(e, cfg.CurseActions)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to group curse actions: %w", err)
		}

		for selector, curseDetail := range grouped {
			adapter := curseDetail.curseAdapter
			subjects := curseDetail.subjects
			alreadyCursedSubjects := make([]Subject, 0)
			for _, subject := range subjects {
				cursed, err := adapter.IsSubjectCursedOnChain(e, selector, subject)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to check if subject %x is cursed on chain with selector %d: %w", subject, selector, err)
				}
				if !cursed && !cfg.Force {
					e.Logger.Infof("Subject %x is not cursed on chain with selector %d, skipping", subject, selector)
					continue
				}
				alreadyCursedSubjects = append(alreadyCursedSubjects, subject)
			}
			if len(alreadyCursedSubjects) == 0 {
				e.Logger.Infof("No new subjects to uncurse on chain with selector %d, skipping", selector)
				continue
			}
			e.Logger.Infof("Uncursing %d subjects on chain with selector %d", len(alreadyCursedSubjects), selector)
			unCurseReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.Uncurse(), e.BlockChains, CurseInput{
				Subjects:      alreadyCursedSubjects,
				ChainSelector: selector,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to curse subjects on chain with selector %d: %w", selector, err)
			}
			batchOps = append(batchOps, unCurseReport.Output.BatchOps...)
			reports = append(reports, unCurseReport.ExecutionReports...)
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
