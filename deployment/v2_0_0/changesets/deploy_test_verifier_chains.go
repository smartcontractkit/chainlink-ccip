package changesets

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// TestVerifierChainConfig specifies configuration for a chain to deploy test verifier infrastructure.
type TestVerifierChainConfig struct {
	// TokenName is the name of the drip token. Defaults to "TESTVTR".
	TokenName string
	// TokenSymbol is the symbol of the drip token. Defaults to "TESTVTR".
	TokenSymbol string
	// TokenDecimals is the number of decimals for the token. Defaults to 18.
	TokenDecimals uint8
	// PreMintAccounts maps addresses to initial mint amounts (as decimal strings in base units).
	PreMintAccounts map[string]string
	// AllowedSenders are addresses allowlisted on the test verifier.
	AllowedSenders []string
	// StorageLocations are passed to the VerifierTestHelper constructor.
	// The verifier's getFee delegates to RMN, which needs valid storage locations.
	StorageLocations []string
	// FeeAggregator is the fee aggregator address for the token pool.
	FeeAggregator string
	// RemoteChains is the set of remote chains to configure lanes for.
	RemoteChains map[uint64]adapters.RemoteTestVerifierChainConfig
}

// DeployTestVerifierChainsConfig is the configuration for the DeployTestVerifierChains changeset.
type DeployTestVerifierChainsConfig struct {
	// Chains specifies the chains to deploy test verifier infrastructure on.
	Chains map[uint64]TestVerifierChainConfig
	// MCMS configures the resulting proposal.
	MCMS *mcms.Input
}

const (
	defaultTokenName     = "TESTVTR"
	defaultTokenSymbol   = "TESTVTR"
	defaultTokenDecimals = 18
)

// DeployTestVerifierChains returns a changeset that deploys test verifier infrastructure
// (TESTVTR drip token, matching pool, test verifier, and resolver) on chains and
// configures lanes between them.
func DeployTestVerifierChains(
	testVerifierChainRegistry *adapters.TestVerifierChainRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) cldf.ChangeSetV2[DeployTestVerifierChainsConfig] {
	return cldf.CreateChangeSet(
		makeApplyDeployTestVerifierChains(testVerifierChainRegistry, mcmsRegistry),
		makeVerifyDeployTestVerifierChains(),
	)
}

func makeVerifyDeployTestVerifierChains() func(cldf.Environment, DeployTestVerifierChainsConfig) error {
	return func(e cldf.Environment, cfg DeployTestVerifierChainsConfig) error {
		if cfg.MCMS != nil {
			if err := cfg.MCMS.Validate(); err != nil {
				return fmt.Errorf("failed to validate MCMS input: %w", err)
			}
		}
		for chainSel, chainCfg := range cfg.Chains {
			if _, err := chain_selectors.GetSelectorFamily(chainSel); err != nil {
				return fmt.Errorf("invalid chain selector %d: %w", chainSel, err)
			}
			for _, addr := range chainCfg.AllowedSenders {
				if !common.IsHexAddress(addr) {
					return fmt.Errorf("invalid allowed sender address %q for chain %d", addr, chainSel)
				}
			}
			for remoteChainSelector := range chainCfg.RemoteChains {
				if _, err := chain_selectors.GetSelectorFamily(remoteChainSelector); err != nil {
					return fmt.Errorf("invalid remote chain selector %d: %w", remoteChainSelector, err)
				}
				remoteChainCfg, ok := cfg.Chains[remoteChainSelector]
				if !ok {
					return fmt.Errorf("remote chain selector %d not found in chains", remoteChainSelector)
				}
				if _, ok := remoteChainCfg.RemoteChains[chainSel]; !ok {
					return fmt.Errorf(
						"chain %d has remote %d but chain %d does not define a remote chain config for %d (each remote must point back to the current chain)",
						chainSel, remoteChainSelector, remoteChainSelector, chainSel,
					)
				}
			}
		}
		return nil
	}
}

func makeApplyDeployTestVerifierChains(
	testVerifierChainRegistry *adapters.TestVerifierChainRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, DeployTestVerifierChainsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DeployTestVerifierChainsConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Resolve adapters for all chains upfront.
		adaptersByChain := make(map[uint64]adapters.TestVerifierChainAdapter)
		for chainSel := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(chainSel)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", chainSel, err)
			}
			adapter, ok := testVerifierChainRegistry.Get(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no test verifier chain adapter registered for chain family %q", family)
			}
			adaptersByChain[chainSel] = adapter
		}

		// Phase 1: Deploy test verifier infrastructure on each chain.
		newDS := datastore.NewMemoryDataStore()
		for chainSel, chainCfg := range cfg.Chains {
			// Apply defaults.
			tokenName := chainCfg.TokenName
			if tokenName == "" {
				tokenName = defaultTokenName
			}
			tokenSymbol := chainCfg.TokenSymbol
			if tokenSymbol == "" {
				tokenSymbol = defaultTokenSymbol
			}
			tokenDecimals := chainCfg.TokenDecimals
			if tokenDecimals == 0 {
				tokenDecimals = defaultTokenDecimals
			}

			dep := adapters.DeployTestVerifierChainDeps{
				BlockChains: e.BlockChains,
				DataStore:   e.DataStore,
			}
			in := adapters.DeployTestVerifierChainInput{
				ChainSelector:    chainSel,
				TokenName:        tokenName,
				TokenSymbol:      tokenSymbol,
				TokenDecimals:    tokenDecimals,
				PreMintAccounts:  chainCfg.PreMintAccounts,
				AllowedSenders:   chainCfg.AllowedSenders,
				StorageLocations: chainCfg.StorageLocations,
				FeeAggregator:    chainCfg.FeeAggregator,
			}
			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adaptersByChain[chainSel].DeployTestVerifierChain(), dep, in)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy test verifier chain on chain %d: %w", chainSel, err)
			}
			batchOps = append(batchOps, deployReport.Output.BatchOps...)
			reports = append(reports, deployReport.ExecutionReports...)
			for _, r := range deployReport.Output.Addresses {
				if err := newDS.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf(
						"failed to add %s %s with address %s on chain %d to datastore: %w",
						r.Type, r.Version, r.Address, r.ChainSelector, err,
					)
				}
			}
		}

		// Merge datastores so configuration can resolve all addresses.
		combinedDS := datastore.NewMemoryDataStore()
		if err := combinedDS.Merge(e.DataStore); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to merge environment datastore: %w", err)
		}
		if err := combinedDS.Merge(newDS.Seal()); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to merge new datastore: %w", err)
		}

		// Phase 2: Configure lanes between all pairs.
		for chainSel, chainCfg := range cfg.Chains {
			remoteChains := make(map[uint64]adapters.RemoteTestVerifierChain)
			remoteChainConfigs := make(map[uint64]adapters.RemoteTestVerifierChainConfig, len(chainCfg.RemoteChains))
			for remoteChainSelector, remoteCfg := range chainCfg.RemoteChains {
				remoteFamily, err := chain_selectors.GetSelectorFamily(remoteChainSelector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for remote chain selector %d: %w", remoteChainSelector, err)
				}
				remoteAdapter, ok := testVerifierChainRegistry.Get(remoteFamily)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no test verifier chain adapter registered for remote chain family %q", remoteFamily)
				}
				remoteChains[remoteChainSelector] = remoteAdapter
				remoteChainConfigs[remoteChainSelector] = remoteCfg
			}

			dep := adapters.ConfigureTestVerifierForLanesDeps{
				BlockChains:  e.BlockChains,
				DataStore:    combinedDS.Seal(),
				RemoteChains: remoteChains,
			}
			in := adapters.ConfigureTestVerifierForLanesInput{
				ChainSelector:  chainSel,
				AllowedSenders: chainCfg.AllowedSenders,
				RemoteChains:   remoteChainConfigs,
			}
			configureReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adaptersByChain[chainSel].ConfigureTestVerifierChainForLanes(), dep, in)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure test verifier chain for lanes on chain %d: %w", chainSel, err)
			}
			batchOps = append(batchOps, configureReport.Output.BatchOps...)
			reports = append(reports, configureReport.ExecutionReports...)
			for _, r := range configureReport.Output.Addresses {
				if err := newDS.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf(
						"failed to add %s %s with address %s on chain %d to datastore: %w",
						r.Type, r.Version, r.Address, r.ChainSelector, err,
					)
				}
			}
		}

		// Return the output.
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

// PreMintToBigInt converts a decimal string amount to *big.Int.
func PreMintToBigInt(amount string) (*big.Int, error) {
	val, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid pre-mint amount: %s", amount)
	}
	return val, nil
}
