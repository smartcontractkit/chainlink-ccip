package token_verifier_config

import (
	"fmt"
	"strconv"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	onrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	dsutil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type VerifierGeneratedConfig struct {
	OnRampAddresses                  map[string]string `json:"on_ramp_addresses"`
	RMNRemoteAddresses               map[string]string `json:"rmn_remote_addresses"`
	CCTPVerifierAddresses            map[string]string `json:"cctp_verifier_addresses"`
	CCTPVerifierResolverAddresses    map[string]string `json:"cctp_verifier_resolver_addresses"`
	LombardVerifierResolverAddresses map[string]string `json:"lombard_verifier_resolver_addresses"`
}

type BuildConfigInput struct {
	CCTPQualifier    string
	LombardQualifier string
	ChainSelectors   []uint64
}

// BuildConfigOutput contains the generated token verifier configuration.
type BuildConfigOutput struct {
	Config *VerifierGeneratedConfig
}

// BuildConfigDeps contains the dependencies for building the token verifier config.
type BuildConfigDeps struct {
	Env deployment.Environment
}

// BuildConfig is an operation that generates the token verifier configuration
// by querying the datastore for contract addresses.
var BuildConfig = operations.NewOperation(
	"build-token-verifier-config",
	semver.MustParse("1.0.0"),
	"Builds the token verifier configuration from datastore contract addresses",
	func(b operations.Bundle, deps BuildConfigDeps, input BuildConfigInput) (BuildConfigOutput, error) {
		ds := deps.Env.DataStore
		toAddress := func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }

		onRampAddresses := make(map[string]string)
		rmnRemoteAddresses := make(map[string]string)
		cctpVerifierAddresses := make(map[string]string)
		cctpVerifierResolverAddresses := make(map[string]string)
		lombardVerifierResolverAddresses := make(map[string]string)

		for _, chainSelector := range input.ChainSelectors {
			chainSelectorStr := strconv.FormatUint(chainSelector, 10)

			// Get OnRamp address (required)
			onRampAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
				Type: datastore.ContractType(onrampoperations.ContractType),
			}, chainSelector, toAddress)
			if err != nil {
				return BuildConfigOutput{}, fmt.Errorf("failed to get on ramp address for chain %d: %w", chainSelector, err)
			}
			onRampAddresses[chainSelectorStr] = onRampAddr

			// Get RMN Remote address (required)
			rmnRemoteAddr, err := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
				Type: datastore.ContractType(rmn_remote.ContractType),
			}, chainSelector, toAddress)
			if err != nil {
				return BuildConfigOutput{}, fmt.Errorf("failed to get rmn remote address for chain %d: %w", chainSelector, err)
			}
			rmnRemoteAddresses[chainSelectorStr] = rmnRemoteAddr

			// Get CCTP Verifier address (ContractType) - optional
			cctpVerifierAddr, cctpVerifierErr := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ContractType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			// Get CCTP Verifier Resolver address (ResolverType) - optional
			cctpVerifierResolverAddr, cctpResolverErr := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
				Type:      datastore.ContractType(cctp_verifier.ResolverType),
				Qualifier: input.CCTPQualifier,
			}, chainSelector, toAddress)

			// Both CCTP addresses must be present or both absent
			if (cctpVerifierErr == nil) != (cctpResolverErr == nil) {
				return BuildConfigOutput{}, fmt.Errorf(
					"chain %d: cctp verifier and resolver must both exist or both be absent (verifier error: %v, resolver error: %v)",
					chainSelector, cctpVerifierErr, cctpResolverErr)
			}

			// If both CCTP addresses exist, add them to the maps
			if cctpVerifierErr == nil && cctpResolverErr == nil {
				cctpVerifierAddresses[chainSelectorStr] = cctpVerifierAddr
				cctpVerifierResolverAddresses[chainSelectorStr] = cctpVerifierResolverAddr
			}

			// Get Lombard Verifier Resolver address (ResolverType) - optional
			lombardVerifierResolverAddr, lombardResolverErr := dsutil.FindAndFormatRef(ds, datastore.AddressRef{
				Type:      datastore.ContractType(lombard_verifier.ResolverType),
				Qualifier: input.LombardQualifier,
			}, chainSelector, toAddress)

			// If Lombard address exists, add it to the map
			if lombardResolverErr == nil {
				lombardVerifierResolverAddresses[chainSelectorStr] = lombardVerifierResolverAddr
			}
		}

		return BuildConfigOutput{
			Config: &VerifierGeneratedConfig{
				OnRampAddresses:                  onRampAddresses,
				RMNRemoteAddresses:               rmnRemoteAddresses,
				CCTPVerifierAddresses:            cctpVerifierAddresses,
				CCTPVerifierResolverAddresses:    cctpVerifierResolverAddresses,
				LombardVerifierResolverAddresses: lombardVerifierResolverAddresses,
			},
		}, nil
	},
)
