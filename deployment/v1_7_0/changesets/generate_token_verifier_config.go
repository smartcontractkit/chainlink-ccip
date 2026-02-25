package changesets

import (
	"fmt"
	"slices"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
)

const (
	// TestnetCCTPAttestationAPI is the Circle attestation API endpoint for testnet/non-production environments.
	TestnetCCTPAttestationAPI = "https://iris-api-sandbox.circle.com"
	// TestnetLombardAttestationAPI is the Lombard attestation API endpoint for testnet/non-production environments.
	TestnetLombardAttestationAPI = "https://gastald-testnet.prod.lombard.finance/api/"
	// MainnetCCTPAttestationAPI is the Circle attestation API endpoint for mainnet/production environments.
	MainnetCCTPAttestationAPI = "https://iris-api.circle.com"
	// MainnetLombardAttestationAPI is the Lombard attestation API endpoint for mainnet/production environments.
	MainnetLombardAttestationAPI = "https://mainnet.prod.lombard.finance/api/"
)

// GenerateTokenVerifierConfigCfg contains the configuration for the generate token verifier config changeset.
type GenerateTokenVerifierConfigCfg struct {
	// ServiceIdentifier is the identifier for this token verifier service (e.g. "default-token-verifier").
	ServiceIdentifier string
	// ChainSelectors are the chains to generate configuration for.
	// If empty, defaults to all chain selectors available in the environment.
	ChainSelectors []uint64
	// PyroscopeURL is the URL of the Pyroscope server for profiling (optional).
	PyroscopeURL string
	// Monitoring is the monitoring configuration containing beholder settings.
	Monitoring shared.MonitoringInput
	// Lombard contains the Lombard verifier configuration.
	Lombard sequences.LombardConfigInput
	// CCTP contains the CCTP verifier configuration.
	CCTP sequences.CCTPConfigInput
}

// GenerateTokenVerifierConfig creates a changeset that generates the token verifier configuration
// by scanning on-chain token verifier contracts (CCTP and Lombard).
func GenerateTokenVerifierConfig() deployment.ChangeSetV2[GenerateTokenVerifierConfigCfg] {
	validate := func(e deployment.Environment, cfg GenerateTokenVerifierConfigCfg) error {
		if cfg.ServiceIdentifier == "" {
			return fmt.Errorf("service identifier is required")
		}

		// Validate chain selectors
		envSelectors := e.BlockChains.ListChainSelectors()
		for _, s := range cfg.ChainSelectors {
			if !slices.Contains(envSelectors, s) {
				return fmt.Errorf("selector %d is not available in environment", s)
			}
		}

		return nil
	}

	apply := func(e deployment.Environment, cfg GenerateTokenVerifierConfigCfg) (deployment.ChangesetOutput, error) {
		selectors := cfg.ChainSelectors
		if len(selectors) == 0 {
			selectors = e.BlockChains.ListChainSelectors()
		}

		// Determine if this is a production environment
		isProd := shared.IsProductionEnvironment(e.Name)

		// Apply defaults for both verifier configurations
		lombardConfig := applyLombardDefaults(cfg.Lombard, isProd)
		cctpConfig := applyCCTPDefaults(cfg.CCTP, isProd)

		input := sequences.GenerateTokenVerifierConfigInput{
			ChainSelectors: selectors,
			PyroscopeURL:   cfg.PyroscopeURL,
			Monitoring:     cfg.Monitoring,
			Lombard:        lombardConfig,
			CCTP:           cctpConfig,
		}

		deps := sequences.GenerateTokenVerifierConfigDeps{
			Env: e,
		}

		report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.GenerateTokenVerifierConfig, deps, input)
		if err != nil {
			return deployment.ChangesetOutput{
				Reports: report.ExecutionReports,
			}, fmt.Errorf("failed to generate token verifier config: %w", err)
		}

		outputDS := datastore.NewMemoryDataStore()
		if e.DataStore != nil {
			if err := outputDS.Merge(e.DataStore); err != nil {
				return deployment.ChangesetOutput{
					Reports: report.ExecutionReports,
				}, fmt.Errorf("failed to merge existing datastore: %w", err)
			}
		}

		if err := ccv.SaveTokenVerifierConfig(outputDS, cfg.ServiceIdentifier, &report.Output.Config); err != nil {
			return deployment.ChangesetOutput{
				Reports: report.ExecutionReports,
			}, fmt.Errorf("failed to save token verifier config: %w", err)
		}

		return deployment.ChangesetOutput{
			Reports:   report.ExecutionReports,
			DataStore: outputDS,
		}, nil
	}

	return deployment.CreateChangeSet(apply, validate)
}

// applyLombardDefaults applies default values to Lombard configuration based on the environment.
func applyLombardDefaults(cfg sequences.LombardConfigInput, isProd bool) sequences.LombardConfigInput {
	attestationAPI := cfg.AttestationAPI
	if attestationAPI == "" {
		if isProd {
			attestationAPI = MainnetLombardAttestationAPI
		} else {
			attestationAPI = TestnetLombardAttestationAPI
		}
	}

	attestationAPITimeout := cfg.AttestationAPITimeout
	if attestationAPITimeout == 0 {
		attestationAPITimeout = 1 * time.Second
	}

	attestationAPIInterval := cfg.AttestationAPInterval
	if attestationAPIInterval == 0 {
		attestationAPIInterval = 100 * time.Millisecond
	}

	attestationAPIBatchSize := cfg.AttestationAPIBatchSize
	if attestationAPIBatchSize == 0 {
		attestationAPIBatchSize = 20
	}

	verifierVersion := cfg.VerifierVersion
	if len(verifierVersion) == 0 {
		verifierVersion = offchain.LombardDefaultVerifierVersion
	}

	return sequences.LombardConfigInput{
		Qualifier:               cfg.Qualifier,
		VerifierID:              cfg.VerifierID,
		VerifierVersion:         verifierVersion,
		AttestationAPI:          attestationAPI,
		AttestationAPITimeout:   attestationAPITimeout,
		AttestationAPInterval:   attestationAPIInterval,
		AttestationAPIBatchSize: attestationAPIBatchSize,
	}
}

// applyCCTPDefaults applies default values to CCTP configuration based on the environment.
func applyCCTPDefaults(cfg sequences.CCTPConfigInput, isProd bool) sequences.CCTPConfigInput {
	attestationAPI := cfg.AttestationAPI
	if attestationAPI == "" {
		if isProd {
			attestationAPI = MainnetCCTPAttestationAPI
		} else {
			attestationAPI = TestnetCCTPAttestationAPI
		}
	}

	attestationAPITimeout := cfg.AttestationAPITimeout
	if attestationAPITimeout == 0 {
		attestationAPITimeout = 1 * time.Second
	}

	attestationAPIInterval := cfg.AttestationAPInterval
	if attestationAPIInterval == 0 {
		attestationAPIInterval = 100 * time.Millisecond
	}

	attestationAPICooldown := cfg.AttestationAPICooldown
	if attestationAPICooldown == 0 {
		attestationAPICooldown = 5 * time.Minute
	}

	verifierVersion := cfg.VerifierVersion
	if len(verifierVersion) == 0 {
		verifierVersion = offchain.CCTPDefaultVerifierVersion
	}

	return sequences.CCTPConfigInput{
		Qualifier:              cfg.Qualifier,
		VerifierID:             cfg.VerifierID,
		VerifierVersion:        verifierVersion,
		AttestationAPI:         attestationAPI,
		AttestationAPITimeout:  attestationAPITimeout,
		AttestationAPInterval:  attestationAPIInterval,
		AttestationAPICooldown: attestationAPICooldown,
	}
}
