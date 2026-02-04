package changesets

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

// CCTPChainConfig specifies configuration required for a chain to deploy CCTP.
type CCTPChainConfig struct {
	// TokenMessenger is the address of the TokenMessenger contract.
	TokenMessenger string
	// USDCToken is the address of the USDCToken contract.
	USDCToken string
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// FastFinalityBps are the basis points charged for fast finality.
	FastFinalityBps uint16
	// StorageLocations is the set of storage locations for the CCTPVerifier contract.
	StorageLocations []string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
	// RegisteredPoolRef is a reference to the pool that should be set on the registry on this chain.
	RegisteredPoolRef datastore.AddressRef
	// RemoteChains is the set of remote chains to configure.
	RemoteChains map[uint64]adapters.RemoteCCTPChainConfig
}

// DeployCCTPChainsConfig is the configuration for the DeployCCTPChains changeset.
type DeployCCTPChainsConfig struct {
	// Chains specifies the chains to deploy CCTP to.
	Chains map[uint64]CCTPChainConfig
	// MCMS configures the resulting proposal.
	MCMS *mcms.Input
}

// DeployCCTPChains returns a changeset that deploys CCTP contracts on chains.
func DeployCCTPChains(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[DeployCCTPChainsConfig] {
	return cldf.CreateChangeSet(makeApplyDeployCCTPChains(cctpChainRegistry, mcmsRegistry), makeVerifyDeployCCTPChains(cctpChainRegistry, mcmsRegistry))
}

func makeVerifyDeployCCTPChains(_ *adapters.CCTPChainRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployCCTPChainsConfig) error {
	return func(e cldf.Environment, cfg DeployCCTPChainsConfig) error {
		if cfg.MCMS != nil {
			err := cfg.MCMS.Validate()
			if err != nil {
				return fmt.Errorf("failed to validate MCMS input: %w", err)
			}
		}

		for chainSel, chainCfg := range cfg.Chains {
			if _, err := chain_selectors.GetSelectorFamily(chainSel); err != nil {
				return err
			}
			for remoteChainSelector := range chainCfg.RemoteChains {
				if _, err := chain_selectors.GetSelectorFamily(remoteChainSelector); err != nil {
					return err
				}
				if _, ok := cfg.Chains[remoteChainSelector]; !ok {
					return fmt.Errorf("remote chain selector %d not found in chains", remoteChainSelector)
				}
			}
		}

		return nil
	}
}

func makeApplyDeployCCTPChains(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployCCTPChainsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DeployCCTPChainsConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		adaptersByChain := make(map[uint64]adapters.CCTPChain)
		for chainSel := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(chainSel)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainSel, err)
			}
			adapter, ok := cctpChainRegistry.GetCCTPChain(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no CCTP adapter registered for chain family '%s'", family)
			}
			adaptersByChain[chainSel] = adapter
		}

		// Deploy across all chains.
		newDS := datastore.NewMemoryDataStore()
		for chainSel, chainCfg := range cfg.Chains {
			dep := adapters.DeployCCTPChainDeps{
				BlockChains: e.BlockChains,
				DataStore:   e.DataStore,
			}
			in := adapters.DeployCCTPInput{
				ChainSelector:    chainSel,
				TokenMessenger:   chainCfg.TokenMessenger,
				USDCToken:        chainCfg.USDCToken,
				DeployerContract: chainCfg.DeployerContract,
				FastFinalityBps:  chainCfg.FastFinalityBps,
				StorageLocations: chainCfg.StorageLocations,
				FeeAggregator:    chainCfg.FeeAggregator,
			}
			deployCCTPChainReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adaptersByChain[chainSel].DeployCCTPChain(), dep, in)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy CCTP on chain with selector %d: %w", chainSel, err)
			}

			batchOps = append(batchOps, deployCCTPChainReport.Output.BatchOps...)
			reports = append(reports, deployCCTPChainReport.ExecutionReports...)
			for _, r := range deployCCTPChainReport.Output.Addresses {
				if err := newDS.Addresses().Add(r); err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
		}

		// Configure across all chains.
		// Create a new datastore that merges the existing datastore with the new datastore.
		// Enables the configuration sequence to have all existing and new addresses.
		combinedDS := datastore.NewMemoryDataStore()
		err := combinedDS.Merge(e.DataStore)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to merge datastore: %w", err)
		}
		err = combinedDS.Merge(newDS.Seal())
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to merge datastore: %w", err)
		}
		for chainSel, chainCfg := range cfg.Chains {
			remoteChains := make(map[uint64]adapters.RemoteCCTPChain)
			remoteRegisteredPoolRefs := make(map[uint64]datastore.AddressRef)
			for remoteChainSelector := range chainCfg.RemoteChains {
				family, err := chain_selectors.GetSelectorFamily(remoteChainSelector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteChainSelector, err)
				}
				adapter, ok := cctpChainRegistry.GetCCTPChain(family)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no CCTP adapter registered for chain family '%s'", family)
				}
				remoteChains[remoteChainSelector] = adapter
				remoteRegisteredPoolRefs[remoteChainSelector] = cfg.Chains[remoteChainSelector].RegisteredPoolRef
			}
			dep := adapters.ConfigureCCTPChainForLanesDeps{
				BlockChains:  e.BlockChains,
				DataStore:    combinedDS.Seal(),
				RemoteChains: remoteChains,
			}
			in := adapters.ConfigureCCTPChainForLanesInput{
				ChainSelector:            chainSel,
				USDCToken:                chainCfg.USDCToken,
				RegisteredPoolRef:        chainCfg.RegisteredPoolRef,
				RemoteRegisteredPoolRefs: remoteRegisteredPoolRefs,
				RemoteChains:             chainCfg.RemoteChains,
			}
			configureCCTPChainForLanesReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adaptersByChain[chainSel].ConfigureCCTPChainForLanes(), dep, in)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure CCTP on chain with selector %d: %w", chainSel, err)
			}
			batchOps = append(batchOps, configureCCTPChainForLanesReport.Output.BatchOps...)
			reports = append(reports, configureCCTPChainForLanesReport.ExecutionReports...)
		}

		// Return the output.
		var mcmsInput mcms.Input
		if cfg.MCMS != nil {
			mcmsInput = *cfg.MCMS
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			WithDataStore(newDS). // We still only want to return the new addresses.
			Build(mcmsInput)
	}
}
