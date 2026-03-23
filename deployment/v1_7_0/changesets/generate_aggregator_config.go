package changesets

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

type GenerateAggregatorConfigInput struct {
	ServiceIdentifier  string
	CommitteeQualifier string
	Topology           *offchain.EnvironmentTopology
}

func GenerateAggregatorConfig(registry *adapters.AggregatorConfigRegistry) deployment.ChangeSetV2[GenerateAggregatorConfigInput] {
	validate := func(e deployment.Environment, cfg GenerateAggregatorConfigInput) error {
		if cfg.ServiceIdentifier == "" {
			return fmt.Errorf("service identifier is required")
		}
		if cfg.CommitteeQualifier == "" {
			return fmt.Errorf("committee qualifier is required")
		}
		_, err := resolveAggregatorChainSelectors(e, cfg)
		return err
	}

	apply := func(e deployment.Environment, cfg GenerateAggregatorConfigInput) (deployment.ChangesetOutput, error) {
		chainSelectors, err := resolveAggregatorChainSelectors(e, cfg)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}

		committee, err := buildAggregatorCommittee(e, registry, cfg.CommitteeQualifier, chainSelectors)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to build aggregator config: %w", err)
		}

		outputDS := datastore.NewMemoryDataStore()
		if e.DataStore != nil {
			if err := outputDS.Merge(e.DataStore); err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to merge existing datastore: %w", err)
			}
		}

		if err := offchain.SaveAggregatorConfig(outputDS, cfg.ServiceIdentifier, committee); err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to save aggregator config: %w", err)
		}

		return deployment.ChangesetOutput{
			DataStore: outputDS,
		}, nil
	}

	return deployment.CreateChangeSet(apply, validate)
}

func resolveAggregatorChainSelectors(e deployment.Environment, cfg GenerateAggregatorConfigInput) ([]uint64, error) {
	if cfg.Topology == nil {
		return nil, fmt.Errorf("topology is required")
	}
	if cfg.Topology.NOPTopology == nil {
		return nil, fmt.Errorf("NOP topology is required")
	}

	committee, ok := cfg.Topology.NOPTopology.Committees[cfg.CommitteeQualifier]
	if !ok {
		return nil, fmt.Errorf("committee %q not found in topology", cfg.CommitteeQualifier)
	}

	chainSelectors, err := getCommitteeChainSelectors(committee)
	if err != nil {
		return nil, err
	}

	envSelectors := e.BlockChains.ListChainSelectors()
	for _, s := range chainSelectors {
		if !slices.Contains(envSelectors, s) {
			return nil, fmt.Errorf("committee %q references chain selector %d which is not available in environment", cfg.CommitteeQualifier, s)
		}
	}

	return chainSelectors, nil
}

func buildAggregatorCommittee(
	e deployment.Environment,
	registry *adapters.AggregatorConfigRegistry,
	committeeQualifier string,
	chainSelectors []uint64,
) (*offchain.Committee, error) {
	ctx := context.Background()

	type chainQualifier struct {
		chainSelector uint64
		qualifier     string
	}
	seen := make(map[chainQualifier]bool)
	allCommittees := make(map[string][]*adapters.CommitteeState)
	for _, sel := range chainSelectors {
		adapter, err := registry.GetByChain(sel)
		if err != nil {
			return nil, err
		}

		states, err := adapter.ScanCommitteeStates(ctx, e, sel)
		if err != nil {
			return nil, fmt.Errorf("failed to scan committee states on chain %d: %w", sel, err)
		}
		for _, state := range states {
			key := chainQualifier{chainSelector: sel, qualifier: state.Qualifier}
			if seen[key] {
				return nil, fmt.Errorf(
					"chain %d has multiple committee verifiers with qualifier %q",
					sel, state.Qualifier,
				)
			}
			seen[key] = true
			allCommittees[state.Qualifier] = append(allCommittees[state.Qualifier], state)
		}
	}

	committeeStates, ok := allCommittees[committeeQualifier]
	if !ok || len(committeeStates) == 0 {
		return nil, fmt.Errorf("committee %q not found in deployed verifier state", committeeQualifier)
	}

	quorumConfigs, err := buildQuorumConfigs(e.DataStore, registry, committeeStates, committeeQualifier, chainSelectors)
	if err != nil {
		return nil, fmt.Errorf("failed to build quorum configs: %w", err)
	}

	destVerifiers, err := buildDestinationVerifiers(e.DataStore, registry, committeeQualifier, chainSelectors)
	if err != nil {
		return nil, fmt.Errorf("failed to build destination verifiers: %w", err)
	}

	return &offchain.Committee{
		QuorumConfigs:        quorumConfigs,
		DestinationVerifiers: destVerifiers,
	}, nil
}

func buildQuorumConfigs(
	ds datastore.DataStore,
	registry *adapters.AggregatorConfigRegistry,
	committeeStates []*adapters.CommitteeState,
	committeeQualifier string,
	chainSelectors []uint64,
) (map[string]*offchain.QuorumConfig, error) {
	supportedChains := make(map[uint64]bool, len(chainSelectors))
	for _, sel := range chainSelectors {
		supportedChains[sel] = true
	}

	quorumConfigs := make(map[string]*offchain.QuorumConfig)

	for _, state := range committeeStates {
		for _, sigConfig := range state.SignatureConfigs {
			if !supportedChains[sigConfig.SourceChainSelector] {
				continue
			}

			chainSelectorStr := strconv.FormatUint(sigConfig.SourceChainSelector, 10)
			if _, exists := quorumConfigs[chainSelectorStr]; exists {
				continue
			}

			adapter, err := registry.GetByChain(sigConfig.SourceChainSelector)
			if err != nil {
				return nil, err
			}

			sourceVerifierAddr, err := adapter.ResolveVerifierAddress(ds, sigConfig.SourceChainSelector, committeeQualifier)
			if err != nil {
				return nil, fmt.Errorf("failed to resolve source verifier for chain %d: %w", sigConfig.SourceChainSelector, err)
			}

			signers := make([]offchain.Signer, 0, len(sigConfig.Signers))
			for _, addr := range sigConfig.Signers {
				signers = append(signers, offchain.Signer{Address: addr})
			}

			quorumConfigs[chainSelectorStr] = &offchain.QuorumConfig{
				SourceVerifierAddress: sourceVerifierAddr,
				Signers:               signers,
				Threshold:             sigConfig.Threshold,
			}
		}
	}

	return quorumConfigs, nil
}

func buildDestinationVerifiers(
	ds datastore.DataStore,
	registry *adapters.AggregatorConfigRegistry,
	committeeQualifier string,
	destChainSelectors []uint64,
) (map[string]string, error) {
	destVerifiers := make(map[string]string, len(destChainSelectors))

	for _, chainSelector := range destChainSelectors {
		adapter, err := registry.GetByChain(chainSelector)
		if err != nil {
			return nil, err
		}

		addr, err := adapter.ResolveVerifierAddress(ds, chainSelector, committeeQualifier)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve destination verifier for chain %d: %w", chainSelector, err)
		}
		destVerifiers[strconv.FormatUint(chainSelector, 10)] = addr
	}

	return destVerifiers, nil
}
