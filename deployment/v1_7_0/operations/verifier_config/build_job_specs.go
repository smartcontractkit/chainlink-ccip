package verifier_config

import (
	"fmt"
	"slices"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

// NOPInput defines the configuration for a single NOP (Node Operator).
type NOPInput struct {
	// Alias is the unique identifier for this NOP.
	Alias shared.NOPAlias
	// SignerAddress is the address used by this NOP for signing attestations.
	SignerAddressByFamily map[string]string
	// Mode specifies how the NOP runs its jobs (cl or standalone).
	Mode shared.NOPMode
}

// AggregatorInput defines the configuration for an aggregator instance.
type AggregatorInput struct {
	// Name is the unique identifier for this aggregator instance.
	Name string
	// Address is the network endpoint of the aggregator service.
	Address string
	// InsecureAggregatorConnection disables TLS verification when connecting to the aggregator.
	InsecureAggregatorConnection bool
}

// CommitteeInput defines the configuration for a verifier committee.
type CommitteeInput struct {
	// Qualifier is the unique identifier for this committee.
	Qualifier string
	// Aggregators is the list of aggregator instances that serve this committee.
	Aggregators []AggregatorInput
	// NOPAliases is the list of NOP aliases that are members of this committee (across all chains).
	NOPAliases []shared.NOPAlias
	// ChainNOPAliases maps chain selector (as string) to the list of NOP aliases that are members on that chain.
	ChainNOPAliases map[string][]shared.NOPAlias
}

type BuildJobSpecsInput struct {
	GeneratedConfig *VerifierGeneratedConfig
	// TargetNOPs limits which NOPs will have their job specs updated. Defaults to all NOPs in the committee when empty.
	TargetNOPs      []shared.NOPAlias
	EnvironmentNOPs []NOPInput
	Committee       CommitteeInput
	PyroscopeURL    string
	Monitoring      shared.MonitoringInput
	// DisableFinalityCheckers is a list of chain selectors (as strings) for which
	// the finality violation checker should be disabled.
	DisableFinalityCheckers []string
}

type BuildJobSpecsOutput struct {
	JobSpecs      shared.NOPJobSpecs
	AffectedScope shared.VerifierJobScope
}

var BuildJobSpecs = operations.NewOperation(
	"build-verifier-job-specs",
	semver.MustParse("1.0.0"),
	"Builds verifier job specs from generated config and explicit input",
	func(b operations.Bundle, deps struct{}, input BuildJobSpecsInput) (BuildJobSpecsOutput, error) {
		nopByAlias := make(map[shared.NOPAlias]NOPInput, len(input.EnvironmentNOPs))
		for _, nop := range input.EnvironmentNOPs {
			nopByAlias[nop.Alias] = nop
		}

		jobSpecs := make(shared.NOPJobSpecs)
		scope := shared.VerifierJobScope{
			CommitteeQualifier: input.Committee.Qualifier,
		}

		targetNOPs := input.TargetNOPs
		if len(targetNOPs) == 0 {
			targetNOPs = input.Committee.NOPAliases
		}

		for _, nopAlias := range targetNOPs {
			nop, ok := nopByAlias[nopAlias]
			if !ok {
				return BuildJobSpecsOutput{}, fmt.Errorf("NOP %q not found in input", nopAlias)
			}

			nopChains := getNOPChainMembership(nopAlias, input.Committee.ChainNOPAliases)

			if len(input.Committee.ChainNOPAliases) > 0 && len(nopChains) == 0 {
				continue
			}

			for _, agg := range input.Committee.Aggregators {
				verifierJobID := shared.NewVerifierJobID(nopAlias, agg.Name, scope)

				// TODO: This need to be updated once verifier support multiple families
				signerAddress := nop.SignerAddressByFamily[chainsel.FamilyEVM]
				if signerAddress == "" {
					return BuildJobSpecsOutput{}, fmt.Errorf("NOP %q missing signer address for family %s - ensure chain is configured on node", nop.Alias, chainsel.FamilyEVM)
				}

				verifierCfg := offchain.VerifierCommitConfig{
					VerifierID:                     verifierJobID.GetVerifierID(),
					AggregatorAddress:              agg.Address,
					InsecureAggregatorConnection:   agg.InsecureAggregatorConnection,
					SignerAddress:                  signerAddress,
					PyroscopeURL:                   input.PyroscopeURL,
					CommitteeVerifierAddresses:     filterAddressesByChains(input.GeneratedConfig.CommitteeVerifierAddresses, nopChains),
					OnRampAddresses:                filterAddressesByChains(input.GeneratedConfig.OnRampAddresses, nopChains),
					DefaultExecutorOnRampAddresses: filterAddressesByChains(input.GeneratedConfig.DefaultExecutorOnRampAddresses, nopChains),
					RMNRemoteAddresses:             filterAddressesByChains(input.GeneratedConfig.RMNRemoteAddresses, nopChains),
					DisableFinalityCheckers:        input.DisableFinalityCheckers,
					Monitoring:                     convertMonitoringInput(input.Monitoring),
				}

				configBytes, err := toml.Marshal(verifierCfg)
				if err != nil {
					return BuildJobSpecsOutput{}, fmt.Errorf("failed to marshal verifier config to TOML for NOP %q aggregator %q: %w", nopAlias, agg.Name, err)
				}

				jobID := verifierJobID.ToJobID()
				jobSpec := fmt.Sprintf(`schemaVersion = 1
type = "ccvcommitteeverifier"
name = "%s"
externalJobID = "%s"
committeeVerifierConfig = """
%s"""
`, string(jobID), jobID.ToExternalJobID(), string(configBytes))

				if jobSpecs[nopAlias] == nil {
					jobSpecs[nopAlias] = make(map[shared.JobID]string)
				}
				jobSpecs[nopAlias][jobID] = jobSpec
			}
		}

		return BuildJobSpecsOutput{
			JobSpecs:      jobSpecs,
			AffectedScope: scope,
		}, nil
	},
)

func convertMonitoringInput(cfg shared.MonitoringInput) offchain.MonitoringConfig {
	return offchain.MonitoringConfig{
		Enabled: cfg.Enabled,
		Type:    cfg.Type,
		Beholder: offchain.BeholderConfig{
			InsecureConnection:       cfg.Beholder.InsecureConnection,
			CACertFile:               cfg.Beholder.CACertFile,
			OtelExporterGRPCEndpoint: cfg.Beholder.OtelExporterGRPCEndpoint,
			OtelExporterHTTPEndpoint: cfg.Beholder.OtelExporterHTTPEndpoint,
			LogStreamingEnabled:      cfg.Beholder.LogStreamingEnabled,
			MetricReaderInterval:     cfg.Beholder.MetricReaderInterval,
			TraceSampleRatio:         cfg.Beholder.TraceSampleRatio,
			TraceBatchTimeout:        cfg.Beholder.TraceBatchTimeout,
		},
	}
}

func getNOPChainMembership(nopAlias shared.NOPAlias, chainNOPAliases map[string][]shared.NOPAlias) map[string]bool {
	chains := make(map[string]bool)
	if chainNOPAliases == nil {
		return chains
	}
	for chainSelector, nops := range chainNOPAliases {
		if slices.Contains(nops, nopAlias) {
			chains[chainSelector] = true
		}
	}
	return chains
}

func filterAddressesByChains(addresses map[string]string, nopChains map[string]bool) map[string]string {
	if len(nopChains) == 0 {
		return addresses
	}
	filtered := make(map[string]string, len(nopChains))
	for chainSelector, addr := range addresses {
		if nopChains[chainSelector] {
			filtered[chainSelector] = addr
		}
	}
	return filtered
}
