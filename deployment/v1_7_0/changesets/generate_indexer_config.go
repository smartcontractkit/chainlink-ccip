package changesets

import (
	"fmt"
	"slices"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	idxconfig "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/indexer_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/sequences"
)

// GenerateIndexerConfigCfg contains the configuration for the generate indexer config changeset.
type GenerateIndexerConfigCfg struct {
	// ServiceIdentifier is the identifier for this indexer service (e.g. "default-indexer")
	ServiceIdentifier string
	// VerifierNameToQualifier maps verifier names (matching VerifierConfig.Name) to qualifiers
	// used for looking up addresses in the datastore.
	CommitteeVerifierNameToQualifier map[string]string
	CCTPVerifierNameToQualifier      map[string]string
	LombardVerifierNameToQualifier   map[string]string
	// ChainSelectors are the source chains the indexer will monitor.
	// If empty, defaults to all chain selectors available in the environment.
	ChainSelectors []uint64
}

// GenerateIndexerConfig creates a changeset that generates the indexer configuration
// by scanning on-chain CommitteeVerifier contracts. It generates one entry per verifier name
// with all IssuerAddresses (resolver addresses) for that verifier across all chains.
func GenerateIndexerConfig() deployment.ChangeSetV2[GenerateIndexerConfigCfg] {
	validate := func(e deployment.Environment, cfg GenerateIndexerConfigCfg) error {
		if cfg.ServiceIdentifier == "" {
			return fmt.Errorf("service identifier is required")
		}
		if len(cfg.CommitteeVerifierNameToQualifier) == 0 &&
			len(cfg.CCTPVerifierNameToQualifier) == 0 &&
			len(cfg.LombardVerifierNameToQualifier) == 0 {
			return fmt.Errorf("at least one verifier name to qualifier mapping is required")
		}
		envSelectors := e.BlockChains.ListChainSelectors()
		for _, s := range cfg.ChainSelectors {
			if !slices.Contains(envSelectors, s) {
				return fmt.Errorf("selector %d is not available in environment", s)
			}
		}

		return nil
	}

	apply := func(e deployment.Environment, cfg GenerateIndexerConfigCfg) (deployment.ChangesetOutput, error) {
		selectors := cfg.ChainSelectors
		if len(selectors) == 0 {
			selectors = e.BlockChains.ListChainSelectors()
		}

		input := idxconfig.BuildConfigInput{
			ServiceIdentifier:                cfg.ServiceIdentifier,
			CommitteeVerifierNameToQualifier: cfg.CommitteeVerifierNameToQualifier,
			CCTPVerifierNameToQualifier:      cfg.CCTPVerifierNameToQualifier,
			LombardVerifierNameToQualifier:   cfg.LombardVerifierNameToQualifier,
			ChainSelectors:                   selectors,
		}

		deps := sequences.GenerateIndexerConfigDeps{
			Env: e,
		}

		report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.GenerateIndexerConfig, deps, input)
		if err != nil {
			return deployment.ChangesetOutput{
				Reports: report.ExecutionReports,
			}, fmt.Errorf("failed to generate indexer config: %w", err)
		}

		outputDS := datastore.NewMemoryDataStore()
		if e.DataStore != nil {
			if err := outputDS.Merge(e.DataStore); err != nil {
				return deployment.ChangesetOutput{
					Reports: report.ExecutionReports,
				}, fmt.Errorf("failed to merge existing datastore: %w", err)
			}
		}

		idxCfg := idxconfig.GeneratedVerifiersToGeneratedConfig(report.Output.Verifiers)

		if err := ccv.SaveIndexerConfig(outputDS, report.Output.ServiceIdentifier, idxCfg); err != nil {
			return deployment.ChangesetOutput{
				Reports: report.ExecutionReports,
			}, fmt.Errorf("failed to save indexer config: %w", err)
		}

		return deployment.ChangesetOutput{
			Reports:   report.ExecutionReports,
			DataStore: outputDS,
		}, nil
	}

	return deployment.CreateChangeSet(apply, validate)
}
