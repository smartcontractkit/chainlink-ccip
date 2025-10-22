package fastcurse

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type RMNCurseConfig struct {
	CurseActions []CurseActionInput
	// Use this if you want to include curse subject even when they are already cursed (CurseChangeset) or already uncursed (UncurseChangeset)
	Force  bool
	Reason string
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

func CurseChangeset(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[RMNCurseConfig] {
	return cldf.CreateChangeSet(applyCurse(cr, mcmsRegistry), verifyCurseInput(cr, mcmsRegistry))
}

func UncurseChangeset(cr *CurseRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[RMNCurseConfig] {
	return cldf.CreateChangeSet(applyUncurse(cr, mcmsRegistry), verifyCurseInput(cr, mcmsRegistry))
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
		for _, adapter := range cr.CurseAdapters {
			err := adapter.Initialize(e)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to initialize necessary curse adapter addresses: %w", err)
			}
		}
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
		for _, adapter := range cr.CurseAdapters {
			err := adapter.Initialize(e)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to initialize necessary curse adapter addresses: %w", err)
			}
		}
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
				// Only curse the subjects that are cursed
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
