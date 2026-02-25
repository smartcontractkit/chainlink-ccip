package aggregator_config

import (
	"fmt"
	"strconv"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	dsutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
)

// Signer represents a participant in the commit verification process.
// This is a serializable intermediate type used for operation output.
type Signer struct {
	Address string
}

// QuorumConfig represents the configuration for a quorum of signers.
// This is a serializable intermediate type used for operation output.
type QuorumConfig struct {
	SourceVerifierAddress string
	Signers               []Signer
	Threshold             uint8
}

// Committee represents a group of signers participating in the commit verification process.
// This is a serializable intermediate type used for operation output.
type Committee struct {
	QuorumConfigs        map[string]*QuorumConfig
	DestinationVerifiers map[string]string
}

// BuildConfigInput contains the input parameters for building the aggregator config.
type BuildConfigInput struct {
	// ServiceIdentifier is the identifier for this aggregator service (e.g. "default-aggregator")
	ServiceIdentifier string
	// CommitteeQualifier is the unique identifier for this committee.
	CommitteeQualifier string
	// ChainSelectors are the chain selectors that will be considered. Defaults to all chain selectors in the environment.
	ChainSelectors []uint64
}

// BuildConfigOutput contains the generated aggregator committee configuration.
type BuildConfigOutput struct {
	ServiceIdentifier string
	Committee         *Committee
}

// BuildConfigDeps contains the dependencies for building the aggregator config.
// Now uses deployment.Environment to access chains for on-chain scanning.
type BuildConfigDeps struct {
	Env deployment.Environment
}

// BuildConfig is an operation that generates the aggregator committee configuration
// by scanning the on-chain state of deployed CommitteeVerifier contracts.
var BuildConfig = operations.NewOperation(
	"build-aggregator-config",
	semver.MustParse("1.0.0"),
	"Builds the aggregator committee configuration from on-chain state",
	func(b operations.Bundle, deps BuildConfigDeps, input BuildConfigInput) (BuildConfigOutput, error) {
		ctx := b.GetContext()

		// Scan the on-chain topology to get actual signers and thresholds from contracts
		onChainTopology, err := ccv.ScanOnChainTopology(ctx, deps.Env)
		if err != nil {
			return BuildConfigOutput{}, fmt.Errorf("failed to scan on-chain topology: %w", err)
		}

		// Get the committee states for the specified qualifier
		committeeStates, ok := onChainTopology.Committees[input.CommitteeQualifier]
		if !ok || len(committeeStates) == 0 {
			return BuildConfigOutput{}, fmt.Errorf("committee %q not found in on-chain topology", input.CommitteeQualifier)
		}

		quorumConfigs, err := buildQuorumConfigsFromOnChain(deps.Env.DataStore, committeeStates, input.CommitteeQualifier, input.ChainSelectors)
		if err != nil {
			return BuildConfigOutput{}, fmt.Errorf("failed to build quorum configs: %w", err)
		}

		destVerifiers, err := buildDestinationVerifiers(deps.Env.DataStore, input.CommitteeQualifier, input.ChainSelectors)
		if err != nil {
			return BuildConfigOutput{}, fmt.Errorf("failed to build destination verifiers: %w", err)
		}

		return BuildConfigOutput{
			ServiceIdentifier: input.ServiceIdentifier,
			Committee: &Committee{
				QuorumConfigs:        quorumConfigs,
				DestinationVerifiers: destVerifiers,
			},
		}, nil
	},
)

// buildQuorumConfigsFromOnChain builds the quorum configuration from on-chain state.
// Only chains in chainSelectors are included in the output.
func buildQuorumConfigsFromOnChain(
	ds datastore.DataStore,
	committeeStates []*ccv.OnChainCommitteeState,
	committeeQualifier string,
	chainSelectors []uint64,
) (map[string]*QuorumConfig, error) {
	supportedChains := make(map[uint64]bool)
	for _, sel := range chainSelectors {
		supportedChains[sel] = true
	}

	quorumConfigs := make(map[string]*QuorumConfig)

	for _, state := range committeeStates {
		for _, sigConfig := range state.SignatureConfigs {
			if !supportedChains[sigConfig.SourceChainSelector] {
				continue
			}

			chainSelectorStr := strconv.FormatUint(sigConfig.SourceChainSelector, 10)

			if _, exists := quorumConfigs[chainSelectorStr]; exists {
				continue
			}

			sourceVerifierAddr, err := dsutils.FindAndFormatFirstRef(ds, sigConfig.SourceChainSelector, func(r datastore.AddressRef) (string, error) { return r.Address, nil },
				datastore.AddressRef{
					Type:      datastore.ContractType(committee_verifier.ResolverType),
					Qualifier: committeeQualifier,
				},
				datastore.AddressRef{
					Type:      datastore.ContractType(committee_verifier.ContractType),
					Qualifier: committeeQualifier,
				},
			)
			if err != nil {
				return nil, fmt.Errorf("failed to resolve source verifier for chain %d: %w", sigConfig.SourceChainSelector, err)
			}

			configSigners := make([]Signer, 0, len(sigConfig.Signers))
			for _, signer := range sigConfig.Signers {
				configSigners = append(configSigners, Signer{
					Address: signer.Hex(),
				})
			}

			quorumConfigs[chainSelectorStr] = &QuorumConfig{
				SourceVerifierAddress: sourceVerifierAddr,
				Signers:               configSigners,
				Threshold:             sigConfig.Threshold,
			}
		}
	}

	return quorumConfigs, nil
}

func buildDestinationVerifiers(
	ds datastore.DataStore,
	committeeQualifier string,
	destChainSelectors []uint64,
) (map[string]string, error) {
	destVerifiers := make(map[string]string)

	for _, chainSelector := range destChainSelectors {
		addr, err := dsutils.FindAndFormatFirstRef(ds, chainSelector, func(r datastore.AddressRef) (string, error) { return r.Address, nil },
			datastore.AddressRef{
				Type:      datastore.ContractType(committee_verifier.ResolverType),
				Qualifier: committeeQualifier,
			},
			datastore.AddressRef{
				Type:      datastore.ContractType(committee_verifier.ContractType),
				Qualifier: committeeQualifier,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve destination verifier for chain %d: %w", chainSelector, err)
		}
		destVerifiers[strconv.FormatUint(chainSelector, 10)] = addr
	}

	return destVerifiers, nil
}
