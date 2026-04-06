package changesets

import (
	"errors"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

type GenerateIndexerConfigInput struct {
	ServiceIdentifier                string
	CommitteeVerifierNameToQualifier map[string]string
	CCTPVerifierNameToQualifier      map[string]string
	LombardVerifierNameToQualifier   map[string]string
}

func GenerateIndexerConfig(registry *adapters.IndexerConfigRegistry) deployment.ChangeSetV2[GenerateIndexerConfigInput] {
	validate := func(e deployment.Environment, cfg GenerateIndexerConfigInput) error {
		if cfg.ServiceIdentifier == "" {
			return fmt.Errorf("service identifier is required")
		}
		if len(cfg.CommitteeVerifierNameToQualifier) == 0 &&
			len(cfg.CCTPVerifierNameToQualifier) == 0 &&
			len(cfg.LombardVerifierNameToQualifier) == 0 {
			return fmt.Errorf("at least one verifier name to qualifier mapping is required")
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg GenerateIndexerConfigInput) (deployment.ChangesetOutput, error) {
		verifierMap, err := buildIndexerVerifierMap(e.DataStore, registry, e.BlockChains.ListChainSelectors(), cfg)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}

		idxCfg := toIndexerGeneratedConfig(verifierMap)

		outputDS := datastore.NewMemoryDataStore()
		if e.DataStore != nil {
			if err := outputDS.Merge(e.DataStore); err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to merge existing datastore: %w", err)
			}
		}

		if err := offchain.SaveIndexerConfig(outputDS, cfg.ServiceIdentifier, idxCfg); err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to save indexer config: %w", err)
		}

		return deployment.ChangesetOutput{
			DataStore: outputDS,
		}, nil
	}

	return deployment.CreateChangeSet(apply, validate)
}

func buildIndexerVerifierMap(
	ds datastore.DataStore,
	registry *adapters.IndexerConfigRegistry,
	selectors []uint64,
	cfg GenerateIndexerConfigInput,
) (map[string][]string, error) {
	verifierMap := make(map[string][]string)

	kindMappings := []struct {
		kind            adapters.VerifierKind
		nameToQualifier map[string]string
	}{
		{adapters.CommitteeVerifierKind, cfg.CommitteeVerifierNameToQualifier},
		{adapters.CCTPVerifierKind, cfg.CCTPVerifierNameToQualifier},
		{adapters.LombardVerifierKind, cfg.LombardVerifierNameToQualifier},
	}

	for _, km := range kindMappings {
		for name, qualifier := range km.nameToQualifier {
			addresses, err := collectVerifierAddresses(ds, registry, selectors, qualifier, km.kind)
			if err != nil {
				return nil, fmt.Errorf("failed to resolve addresses for verifier %q (qualifier %q): %w", name, qualifier, err)
			}
			verifierMap[name] = append(verifierMap[name], addresses...)
		}
	}

	return verifierMap, nil
}

func collectVerifierAddresses(
	ds datastore.DataStore,
	registry *adapters.IndexerConfigRegistry,
	selectors []uint64,
	qualifier string,
	kind adapters.VerifierKind,
) ([]string, error) {
	if !registry.HasAdapters() {
		return nil, fmt.Errorf("no indexer config adapter registered")
	}

	seen := make(map[string]bool)
	var addresses []string

	for _, sel := range selectors {
		adapter, err := registry.GetByChain(sel)
		if err != nil {
			return nil, err
		}

		addrs, err := adapter.ResolveVerifierAddresses(ds, sel, qualifier, kind)
		if err != nil {
			var missingErr *adapters.MissingIndexerVerifierAddressesError
			if errors.As(err, &missingErr) {
				continue
			}
			return nil, fmt.Errorf("chain %d: %w", sel, err)
		}

		for _, addr := range addrs {
			if !seen[addr] {
				seen[addr] = true
				addresses = append(addresses, addr)
			}
		}
	}

	if len(addresses) == 0 {
		return nil, fmt.Errorf("no deployed %s verifier addresses found for qualifier %q", kind, qualifier)
	}

	return addresses, nil
}

func toIndexerGeneratedConfig(verifierMap map[string][]string) *offchain.IndexerGeneratedConfig {
	verifiers := make([]offchain.IndexerVerifierConfig, 0, len(verifierMap))

	names := make([]string, 0, len(verifierMap))
	for name := range verifierMap {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		addresses := verifierMap[name]
		seen := make(map[string]bool)
		unique := make([]string, 0, len(addresses))
		for _, addr := range addresses {
			if !seen[addr] {
				seen[addr] = true
				unique = append(unique, addr)
			}
		}
		verifiers = append(verifiers, offchain.IndexerVerifierConfig{
			Name:            name,
			IssuerAddresses: unique,
		})
	}

	return &offchain.IndexerGeneratedConfig{
		Verifiers: verifiers,
	}
}
