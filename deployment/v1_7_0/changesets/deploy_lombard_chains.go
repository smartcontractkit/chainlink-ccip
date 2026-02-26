package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

type LombardChainConfig struct {
	// Bridge is the address of the Bridge contract provided by Lombard
	Bridge string
	// Token is the address of the token to be used in the LombardTokenPool.
	Token string
	// Adapter is the optional adapter address used by LombardTokenPool.
	Adapter        string
	TokenQualifier string
	// DeployerContract is a contract that can be used to deploy other contracts.
	// i.e. A CREATE2Factory contract on Ethereum can enable consistent deployments.
	DeployerContract string
	// StorageLocations is the set of storage locations for the LombardVerifier contract.
	StorageLocations []string
	// FeeAggregator is the address to which fees are withdrawn.
	FeeAggregator string
	// RateLimitAdmin is the address allowed to update token pool rate limits.
	RateLimitAdmin string
	// RemoteChains is the set of remote chains to configure.
	RemoteChains map[uint64]adapters.RemoteLombardChainConfig
}

type DeployLombardChainsConfig struct {
	Chains map[uint64]LombardChainConfig
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

		for chainSel, chainCfg := range cfg.Chains {
			if _, err := chain_selectors.GetSelectorFamily(chainSel); err != nil {
				return err
			}
			if chainCfg.Adapter != "" && !common.IsHexAddress(chainCfg.Adapter) {
				return fmt.Errorf("invalid adapter address for chain %d: %s", chainSel, chainCfg.Adapter)
			}
			for remoteChainSelector := range chainCfg.RemoteChains {
				if _, err := chain_selectors.GetSelectorFamily(remoteChainSelector); err != nil {
					return err
				}
				remoteCfg := chainCfg.RemoteChains[remoteChainSelector]
				if remoteCfg.RemoteAdapter != "" && !common.IsHexAddress(remoteCfg.RemoteAdapter) {
					return fmt.Errorf("invalid remote adapter for chain %d remote %d: %s", chainSel, remoteChainSelector, remoteCfg.RemoteAdapter)
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

		adaptersByChain := make(map[uint64]adapters.LombardChain)
		for chainSel := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(chainSel)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainSel, err)
			}
			adapter, ok := lombardChainRegistry.GetLombardChain(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no CCTP adapter registered for chain family '%s'", family)
			}
			adaptersByChain[chainSel] = adapter
		}

		// Deploy across all chains.
		newDS := datastore.NewMemoryDataStore()
		for chainSel, chainCfg := range cfg.Chains {
			dep := adapters.DeployLombardChainDeps{
				BlockChains: e.BlockChains,
				DataStore:   e.DataStore,
			}

			in := adapters.DeployLombardInput{
				ChainSelector:    chainSel,
				Bridge:           chainCfg.Bridge,
				Token:            chainCfg.Token,
				Adapter:          chainCfg.Adapter,
				TokenQualifier:   chainCfg.TokenQualifier,
				DeployerContract: chainCfg.DeployerContract,
				StorageLocations: chainCfg.StorageLocations,
				FeeAggregator:    chainCfg.FeeAggregator,
				RateLimitAdmin:   chainCfg.RateLimitAdmin,
			}
			deployLombardChainReport, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle, adaptersByChain[chainSel].DeployLombardChain(), dep, in)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy Lombard on chain with selector %d: %w", chainSel, err)
			}

			batchOps = append(batchOps, deployLombardChainReport.Output.BatchOps...)
			reports = append(reports, deployLombardChainReport.ExecutionReports...)
			for _, r := range deployLombardChainReport.Output.Addresses {
				if err := newDS.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
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
			remoteChains := make(map[uint64]adapters.RemoteLombardChain)
			for remoteChainSelector := range chainCfg.RemoteChains {
				family, err := chain_selectors.GetSelectorFamily(remoteChainSelector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteChainSelector, err)
				}
				adapter, ok := lombardChainRegistry.GetLombardChain(family)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no CCTP adapter registered for chain family '%s'", family)
				}
				remoteChains[remoteChainSelector] = adapter
			}

			dep := adapters.ConfigureLombardChainForLanesDeps{
				BlockChains:  e.BlockChains,
				DataStore:    combinedDS.Seal(),
				RemoteChains: remoteChains,
			}
			in := adapters.ConfigureLombardChainForLanesInput{
				ChainSelector:  chainSel,
				Token:          chainCfg.Token,
				TokenQualifier: chainCfg.TokenQualifier,
				RemoteChains:   chainCfg.RemoteChains,
			}
			configureChainForLanesReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adaptersByChain[chainSel].ConfigureLombardChainForLanes(), dep, in)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure Lombard on chain with selector %d: %w", chainSel, err)
			}
			batchOps = append(batchOps, configureChainForLanesReport.Output.BatchOps...)
			reports = append(reports, configureChainForLanesReport.ExecutionReports...)
		}

		var mcmsInput mcms.Input
		if cfg.MCMS != nil {
			mcmsInput = *cfg.MCMS
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			WithDataStore(newDS).
			Build(mcmsInput)
	}
}
