package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// NewFromOnChainSequenceParams configures NewFromOnChainSequence.
type NewFromOnChainSequenceParams[IN any, DEP any, CFG any] struct {
	// Sequence is the operations.Sequence to execute.
	Sequence *operations.Sequence[IN, sequences.OnChainOutput, DEP]
	// ResolveInput resolves the input for the sequence based on the environment and changeset config.
	ResolveInput func(e deployment.Environment, cfg CFG) (IN, error)
	// ResolveDeps resolves the dependencies for the sequence based on the environment and changeset config.
	ResolveDep func(e deployment.Environment, cfg CFG) (DEP, error)
}

type WithMCMS[CFG any] struct {
	MCMS mcms.Input
	Cfg  CFG
}

// NewFromOnChainSequence creates a Changeset from an operations.Sequence that deploys contracts on-chain and performs write operations.
// It wraps sequence execution with DataStore and MCMS integration. The returned function takes an MCMSReaderRegistry and returns a Changeset.
func NewFromOnChainSequence[IN any, DEP any, CFG any](params NewFromOnChainSequenceParams[IN, DEP, CFG]) func(mcmsRegistry *MCMSReaderRegistry) deployment.ChangeSetV2[WithMCMS[CFG]] {
	return func(mcmsRegistry *MCMSReaderRegistry) deployment.ChangeSetV2[WithMCMS[CFG]] {
		resolve := func(e deployment.Environment, cfg CFG) (IN, DEP, error) {
			var in IN
			var dep DEP
			var err error
			in, err = params.ResolveInput(e, cfg)
			if err != nil {
				return in, dep, fmt.Errorf("failed to resolve input for sequence with ID %s: %w", params.Sequence.ID(), err)
			}
			dep, err = params.ResolveDep(e, cfg)
			if err != nil {
				return in, dep, fmt.Errorf("failed to resolve dependencies for sequence with ID %s: %w", params.Sequence.ID(), err)
			}
			return in, dep, nil
		}
		validate := func(e deployment.Environment, cfg WithMCMS[CFG]) error {
			_, _, err := resolve(e, cfg.Cfg)
			return err
		}
		apply := func(e deployment.Environment, cfg WithMCMS[CFG]) (deployment.ChangesetOutput, error) {
			input, deps, err := resolve(e, cfg.Cfg)
			if err != nil {
				return deployment.ChangesetOutput{}, err
			}
			report, err := operations.ExecuteSequence(e.OperationsBundle, params.Sequence, deps, input)
			if err != nil {
				return deployment.ChangesetOutput{Reports: report.ExecutionReports}, fmt.Errorf("failed to execute sequence with ID %s: %w", params.Sequence.ID(), err)
			}
			ds := datastore.NewMemoryDataStore()
			for _, r := range report.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return deployment.ChangesetOutput{Reports: report.ExecutionReports}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}

			return NewOutputBuilder(e, mcmsRegistry).
				WithReports(report.ExecutionReports).
				WithDataStore(ds).
				WithBatchOps(report.Output.BatchOps).
				Build(cfg.MCMS)
		}

		return deployment.CreateChangeSet(apply, validate)
	}
}
