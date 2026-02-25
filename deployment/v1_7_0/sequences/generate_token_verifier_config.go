package sequences

import (
	"fmt"
	"time"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/token_verifier_config"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type GenerateTokenVerifierConfigInput struct {
	// ChainSelectors is the list of chain selectors to consider. Defaults to all chain selectors in the environment.
	ChainSelectors []uint64
	// PyroscopeURL is the URL of the Pyroscope server for profiling (optional).
	PyroscopeURL string
	// Monitoring is the monitoring configuration containing beholder settings.
	Monitoring shared.MonitoringInput
	// Lombard contains the configuration for the Lombard verifier.
	Lombard LombardConfigInput
	// CCTP contains the configuration for the CCTP verifier.
	CCTP CCTPConfigInput
}

type LombardConfigInput struct {
	Qualifier               string
	VerifierID              string
	VerifierVersion         offchain.ByteSlice
	AttestationAPI          string
	AttestationAPITimeout   time.Duration
	AttestationAPInterval   time.Duration
	AttestationAPIBatchSize int
}

type CCTPConfigInput struct {
	Qualifier              string
	VerifierID             string
	VerifierVersion        offchain.ByteSlice
	AttestationAPI         string
	AttestationAPITimeout  time.Duration
	AttestationAPInterval  time.Duration
	AttestationAPICooldown time.Duration
}

type GenerateTokenVerifierConfigDeps struct {
	Env deployment.Environment
}

type GenerateTokenVerifierConfigOutput struct {
	Config offchain.TokenVerifierConfig
}

var GenerateTokenVerifierConfig = operations.NewSequence(
	"generate-token-verifier-config",
	semver.MustParse("1.0.0"),
	"Generates token verifier configuration from datastore contract addresses and explicit input",
	func(b operations.Bundle, deps GenerateTokenVerifierConfigDeps, input GenerateTokenVerifierConfigInput) (GenerateTokenVerifierConfigOutput, error) {
		buildResult, err := operations.ExecuteOperation(b, token_verifier_config.BuildConfig, token_verifier_config.BuildConfigDeps{
			Env: deps.Env,
		}, token_verifier_config.BuildConfigInput{
			ChainSelectors:   input.ChainSelectors,
			CCTPQualifier:    input.CCTP.Qualifier,
			LombardQualifier: input.Lombard.Qualifier,
		})
		if err != nil {
			return GenerateTokenVerifierConfigOutput{}, fmt.Errorf("failed to build token verifier config: %w", err)
		}

		// Build the verifier config from the build result and input
		config := offchain.TokenVerifierConfig{
			PyroscopeURL:       input.PyroscopeURL,
			OnRampAddresses:    buildResult.Output.Config.OnRampAddresses,
			RMNRemoteAddresses: buildResult.Output.Config.RMNRemoteAddresses,
			TokenVerifiers:     []offchain.TokenVerifierPluginConfig{},
			Monitoring: offchain.MonitoringConfig{
				Enabled: input.Monitoring.Enabled,
				Type:    input.Monitoring.Type,
				Beholder: offchain.BeholderConfig{
					InsecureConnection:       input.Monitoring.Beholder.InsecureConnection,
					CACertFile:               input.Monitoring.Beholder.CACertFile,
					OtelExporterGRPCEndpoint: input.Monitoring.Beholder.OtelExporterGRPCEndpoint,
					OtelExporterHTTPEndpoint: input.Monitoring.Beholder.OtelExporterHTTPEndpoint,
					LogStreamingEnabled:      input.Monitoring.Beholder.LogStreamingEnabled,
					MetricReaderInterval:     input.Monitoring.Beholder.MetricReaderInterval,
					TraceSampleRatio:         input.Monitoring.Beholder.TraceSampleRatio,
					TraceBatchTimeout:        input.Monitoring.Beholder.TraceBatchTimeout,
				},
			},
		}

		// Add CCTP verifier if CCTP addresses are present
		if len(buildResult.Output.Config.CCTPVerifierAddresses) > 0 {
			cctpVerifierID := input.CCTP.VerifierID
			if cctpVerifierID == "" {
				cctpVerifierID = fmt.Sprintf("cctp-%s", input.CCTP.Qualifier)
			}
			cctpVerifier := offchain.TokenVerifierPluginConfig{
				VerifierID: cctpVerifierID,
				Type:       "cctp",
				Version:    "2.0",
				CCTPConfig: &offchain.CCTPConfig{
					AttestationAPI:         input.CCTP.AttestationAPI,
					AttestationAPITimeout:  input.CCTP.AttestationAPITimeout,
					AttestationAPIInterval: input.CCTP.AttestationAPInterval,
					AttestationAPICooldown: input.CCTP.AttestationAPICooldown,
					VerifierVersion:        input.CCTP.VerifierVersion,
					Verifiers:              convertToMap(buildResult.Output.Config.CCTPVerifierAddresses),
					VerifierResolvers:      convertToMap(buildResult.Output.Config.CCTPVerifierResolverAddresses),
				},
			}
			config.TokenVerifiers = append(config.TokenVerifiers, cctpVerifier)
		}

		// Add Lombard verifier if Lombard addresses are present
		if len(buildResult.Output.Config.LombardVerifierResolverAddresses) > 0 {
			lombardVerifierID := input.Lombard.VerifierID
			if lombardVerifierID == "" {
				lombardVerifierID = fmt.Sprintf("lombard-%s", input.Lombard.Qualifier)
			}
			lombardVerifier := offchain.TokenVerifierPluginConfig{
				VerifierID: lombardVerifierID,
				Type:       "lombard",
				Version:    "1.0",
				LombardConfig: &offchain.LombardConfig{
					AttestationAPI:          input.Lombard.AttestationAPI,
					AttestationAPITimeout:   input.Lombard.AttestationAPITimeout,
					AttestationAPIInterval:  input.Lombard.AttestationAPInterval,
					AttestationAPIBatchSize: input.Lombard.AttestationAPIBatchSize,
					VerifierVersion:         input.Lombard.VerifierVersion,
					VerifierResolvers:       convertToMap(buildResult.Output.Config.LombardVerifierResolverAddresses),
				},
			}
			config.TokenVerifiers = append(config.TokenVerifiers, lombardVerifier)
		}

		return GenerateTokenVerifierConfigOutput{
			Config: config,
		}, nil
	},
)

// convertToMap converts map[string]string to map[string]any for TOML serialization.
func convertToMap(input map[string]string) map[string]any {
	result := make(map[string]any, len(input))
	for k, v := range input {
		result[k] = v
	}
	return result
}
