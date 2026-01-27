package changesets

import (
	"fmt"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type DeployLombardChainsConfig struct {
	Chains []adapters.DeployLombardInput[datastore.AddressRef, datastore.AddressRef]
	// MCMS configures the resulting proposal.
	MCMS *mcms.Input
}

func DeployLombardChains(lombardChainRegistry *adapters.LombardChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[DeployLombardChainsConfig] {
	return cldf.CreateChangeSet(
		makeApplyDeployLombardChains(lombardChainRegistry, mcmsRegistry),
		makeVerifyDeployLombardChains(lombardChainRegistry, mcmsRegistry),
	)
}

func makeVerifyDeployLombardChains(_ *adapters.LombardChainRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployLombardChainsConfig) error {
	return func(e cldf.Environment, cfg DeployLombardChainsConfig) error {
		if cfg.MCMS != nil {
			err := cfg.MCMS.Validate()
			if err != nil {
				return fmt.Errorf("failed to validate MCMS input: %w", err)
			}
		}

		for _, chainCfg := range cfg.Chains {
			if _, err := chain_selectors.GetSelectorFamily(chainCfg.ChainSelector); err != nil {
				return err
			}
			if chainCfg.TokenPool.Address == "" {
				return fmt.Errorf("token pool is empty for chain with selector %d", chainCfg.ChainSelector)
			}
			for remoteChainSelector := range chainCfg.RemoteChains {
				if _, err := chain_selectors.GetSelectorFamily(remoteChainSelector); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func makeApplyDeployLombardChains(lombardChainRegistry *adapters.LombardChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployLombardChainsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DeployLombardChainsConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

		for _, chainCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(chainCfg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainCfg.ChainSelector, err)
			}
			adapter, ok := lombardChainRegistry.GetLombardChain(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no CCTP adapter registered for chain family '%s'", family)
			}

			// Resolve AddressRefs in the adapter input
			resolvedInput, err := resolveDeployLombardChainInput(e, chainCfg.ChainSelector, chainCfg, adapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve DeployCCTPInput for chain selector %d: %w", chainCfg.ChainSelector, err)
			}

			// Call into DeployCCTPChain sequence
			deployLombardChainReport, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle, adapter.DeployLombardChain(), e.BlockChains, resolvedInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy CCTP on chain with selector %d: %w", chainCfg.ChainSelector, err)
			}

			batchOps = append(batchOps, deployLombardChainReport.Output.BatchOps...)
			reports = append(reports, deployLombardChainReport.ExecutionReports...)
			for _, r := range deployLombardChainReport.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
		}

		var mcmsInput mcms.Input
		if cfg.MCMS != nil {
			mcmsInput = *cfg.MCMS
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			WithDataStore(ds).
			Build(mcmsInput)
	}
}

func resolveDeployLombardChainInput(
	e cldf.Environment,
	selector uint64,
	cfg adapters.DeployLombardInput[datastore.AddressRef, datastore.AddressRef],
	adapter adapters.LombardChain,
) (adapters.DeployLombardInput[string, []byte], error) {
	out := adapters.DeployLombardInput[string, []byte]{
		ChainSelector: selector,
	}
	return out, nil
}
