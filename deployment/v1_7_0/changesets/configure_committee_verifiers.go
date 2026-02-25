package changesets

import (
	"fmt"
	"slices"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/fetch_signing_keys"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// CommitteeVerifierRemoteChainConfig configures the CommitteeVerifier for a remote chain.
type CommitteeVerifierRemoteChainConfig struct {
	// Whether to allow traffic TO the remote chain.
	AllowlistEnabled bool
	// Addresses that are allowed to send messages TO the remote chain.
	AddedAllowlistedSenders []string
	// Addresses that are no longer allowed to send messages TO the remote chain.
	RemovedAllowlistedSenders []string
	// The fee in USD cents charged for verification on the remote chain.
	FeeUSDCents uint16
	// The gas required to execute the verification call on the destination chain (used for billing).
	GasForVerification uint32
	// The size of the CCV specific payload in bytes (used for billing).
	PayloadSizeBytes uint32
}

// CommitteeVerifierConfig configures a CommitteeVerifier contract.
type CommitteeVerifierConfig struct {
	CommitteeQualifier string
	// RemoteChains specifies the configuration for each remote chain supported by the committee verifier.
	RemoteChains map[uint64]CommitteeVerifierRemoteChainConfig
}

type PartialChainConfig struct {
	// The selector of the chain being configured.
	ChainSelector uint64
	// The Router on the chain being configured.
	// We assume that all connections defined will use the same router, either test or production.
	Router datastore.AddressRef
	// The OnRamp on the chain being configured.
	// Similarly, we assume that all connections will use the same OnRamp.
	OnRamp datastore.AddressRef
	// The CommitteeVerifiers on the chain being configured.
	// There can be multiple committee verifiers on a chain, each controlled by a different entity.
	CommitteeVerifiers []CommitteeVerifierConfig
	// The FeeQuoter on the chain being configured.
	FeeQuoter datastore.AddressRef
	// The OffRamp on the chain being configured
	OffRamp datastore.AddressRef
	// The configuration for each remote chain that we want to connect to.
	RemoteChains map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]
}

// ConfigureChainsForLanesConfig is the configuration for the ConfigureChainsForLanes changeset.
type ConfigureChainsForLanesFromTopologyConfig struct {
	Topology *topology.EnvironmentTopology
	// Chains specifies the chains to configure.
	Chains []PartialChainConfig
	// MCMS configures the resulting proposal.
	MCMS mcms.Input
}

// ConfigureCommitteeVerifiersFromTopology creates a changeset that configures
// CommitteeVerifier contracts with signers and thresholds from the topology.
func ConfigureChainsForLanesFromTopology(chainFamilyRegistry *adapters.ChainFamilyRegistry, mcmsRegistry *changesetscore.MCMSReaderRegistry) deployment.ChangeSetV2[ConfigureChainsForLanesFromTopologyConfig] {
	validate := func(e deployment.Environment, cfg ConfigureChainsForLanesFromTopologyConfig) error {
		if cfg.Topology == nil {
			return fmt.Errorf("topology is required")
		}

		if len(cfg.Topology.NOPTopology.Committees) == 0 {
			return fmt.Errorf("no committees defined in topology")
		}

		for _, chain := range cfg.Chains {
			if !slices.Contains(e.BlockChains.ListChainSelectors(), chain.ChainSelector) {
				return fmt.Errorf("chain selector %d is not available in environment", chain.ChainSelector)
			}
		}

		return nil
	}

	apply := func(e deployment.Environment, cfg ConfigureChainsForLanesFromTopologyConfig) (deployment.ChangesetOutput, error) {
		if cfg.Topology == nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("topology is required")
		}

		signingKeysByNOP := fetchAllSigningKeysForTopology(e, cfg.Topology)

		chains := make([]ChainConfig, 0, len(cfg.Chains))
		for _, chain := range cfg.Chains {
			committeeVerifiers := make([]adapters.CommitteeVerifierConfig[datastore.AddressRef], 0, len(chain.CommitteeVerifiers))
			for _, committeeVerifier := range chain.CommitteeVerifiers {
				remoteChains := make(map[uint64]adapters.CommitteeVerifierRemoteChainConfig, len(committeeVerifier.RemoteChains))
				for remoteChainSelector, remoteChainConfig := range committeeVerifier.RemoteChains {
					signatureConfig, err := getSignatureConfigForLane(e, cfg.Topology, committeeVerifier.CommitteeQualifier, chain.ChainSelector, remoteChainSelector, signingKeysByNOP)
					if err != nil {
						return deployment.ChangesetOutput{}, fmt.Errorf("failed to get signature config for source chain %d: %w", remoteChainSelector, err)
					}
					remoteChains[remoteChainSelector] = adapters.CommitteeVerifierRemoteChainConfig{
						AllowlistEnabled:          remoteChainConfig.AllowlistEnabled,
						AddedAllowlistedSenders:   remoteChainConfig.AddedAllowlistedSenders,
						RemovedAllowlistedSenders: remoteChainConfig.RemovedAllowlistedSenders,
						FeeUSDCents:               remoteChainConfig.FeeUSDCents,
						GasForVerification:        remoteChainConfig.GasForVerification,
						PayloadSizeBytes:          remoteChainConfig.PayloadSizeBytes,
						SignatureConfig:           *signatureConfig,
					}
				}

				committeeVerifierAddresses := e.DataStore.Addresses().Filter(
					datastore.AddressRefByChainSelector(chain.ChainSelector),
					datastore.AddressRefByType(datastore.ContractType(committee_verifier.ContractType)),
					datastore.AddressRefByQualifier(committeeVerifier.CommitteeQualifier),
				)
				if len(committeeVerifierAddresses) == 0 {
					return deployment.ChangesetOutput{}, fmt.Errorf("no committee verifier addresses found for chain %d and committee qualifier %q", chain.ChainSelector, committeeVerifier.CommitteeQualifier)
				}

				committeeVerifierResolverAddresses := e.DataStore.Addresses().Filter(
					datastore.AddressRefByChainSelector(chain.ChainSelector),
					datastore.AddressRefByType(datastore.ContractType(committee_verifier.ResolverType)),
					datastore.AddressRefByQualifier(committeeVerifier.CommitteeQualifier),
				)
				if len(committeeVerifierResolverAddresses) == 0 {
					return deployment.ChangesetOutput{}, fmt.Errorf("no committee verifier resolver addresses found for chain %d and committee qualifier %q", chain.ChainSelector, committeeVerifier.CommitteeQualifier)
				}

				committeeVerifiers = append(committeeVerifiers, adapters.CommitteeVerifierConfig[datastore.AddressRef]{
					CommitteeVerifier: []datastore.AddressRef{
						{
							Address:       committeeVerifierAddresses[0].Address,
							ChainSelector: chain.ChainSelector,
							Qualifier:     committeeVerifier.CommitteeQualifier,
						},
						{
							Address:       committeeVerifierResolverAddresses[0].Address,
							ChainSelector: chain.ChainSelector,
							Qualifier:     committeeVerifier.CommitteeQualifier,
						},
					},
					RemoteChains: remoteChains,
				})
			}
			chains = append(chains, ChainConfig{
				ChainSelector:      chain.ChainSelector,
				RemoteChains:       chain.RemoteChains,
				FeeQuoter:          chain.FeeQuoter,
				OnRamp:             chain.OnRamp,
				OffRamp:            chain.OffRamp,
				Router:             chain.Router,
				CommitteeVerifiers: committeeVerifiers,
			})
		}

		return ConfigureChainsForLanes(chainFamilyRegistry, mcmsRegistry).Apply(e, ConfigureChainsForLanesConfig{
			Chains: chains,
			MCMS:   cfg.MCMS,
		})
	}

	return deployment.CreateChangeSet(apply, validate)
}

func fetchAllSigningKeysForTopology(e deployment.Environment, topo *topology.EnvironmentTopology) fetch_signing_keys.SigningKeysByNOP {
	if e.Offchain == nil {
		return nil
	}

	aliasSet := make(map[string]bool)
	for _, nop := range topo.NOPTopology.NOPs {
		if nop.SignerAddressByFamily == nil || nop.SignerAddressByFamily[chainsel.FamilyEVM] == "" {
			aliasSet[nop.Alias] = true
		}
	}

	if len(aliasSet) == 0 {
		return nil
	}

	aliases := make([]string, 0, len(aliasSet))
	for alias := range aliasSet {
		aliases = append(aliases, alias)
	}

	if e.Offchain == nil {
		e.Logger.Debugw("Offchain client not available, skipping signing key fetch")
		return nil
	}

	report, err := operations.ExecuteOperation(
		e.OperationsBundle,
		fetch_signing_keys.FetchNOPSigningKeys,
		fetch_signing_keys.FetchSigningKeysDeps{
			JDClient: e.Offchain,
			Logger:   e.Logger,
			NodeIDs:  e.NodeIDs,
		},
		fetch_signing_keys.FetchSigningKeysInput{
			NOPAliases: aliases,
		},
	)
	if err != nil {
		e.Logger.Warnw("Failed to fetch signing keys from JD", "error", err)
		return nil
	}

	return report.Output.SigningKeysByNOP
}

func getSignatureConfigForLane(
	e deployment.Environment,
	topo *topology.EnvironmentTopology,
	committeeQualifier string,
	localSelector uint64,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (*adapters.CommitteeVerifierSignatureQuorumConfig, error) {
	committee, ok := topo.NOPTopology.Committees[committeeQualifier]
	if !ok {
		return nil, fmt.Errorf("committee %q not found", committeeQualifier)
	}

	chainCfg, ok := committee.ChainConfigs[strconv.FormatUint(remoteSelector, 10)]
	if !ok {
		return nil, fmt.Errorf("chain selector %d not found in committee %q", remoteSelector, committeeQualifier)
	}

	// Get the local chain family to determine which signer address to use for each NOP alias
	// The signer format can differ by chain family, e.g. EVM uses hex addresses while other chains may use public keys
	localFamily, err := chainsel.GetSelectorFamily(localSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get selector family for selector %d: %w", localSelector, err)
	}

	signers := make([]string, 0, len(chainCfg.NOPAliases))
	for _, alias := range chainCfg.NOPAliases {
		signer, err := signerAddressForNOPAlias(e, topo, alias, localFamily, committeeQualifier, remoteSelector, signingKeysByNOP)
		if err != nil {
			return nil, err
		}
		signers = append(signers, signer)
	}

	return &adapters.CommitteeVerifierSignatureQuorumConfig{
		Threshold: chainCfg.Threshold,
		Signers:   signers,
	}, nil
}

func signerAddressForNOPAlias(
	e deployment.Environment,
	topo *topology.EnvironmentTopology,
	alias string,
	localFamily string,
	committeeQualifier string,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (string, error) {
	nop, ok := topo.NOPTopology.GetNOP(alias)
	if !ok {
		return "", fmt.Errorf(
			"NOP alias %q not found for committee %q chain %d",
			alias, committeeQualifier, remoteSelector,
		)
	}

	// Config wins
	if nop.SignerAddressByFamily != nil {
		if addr := nop.SignerAddressByFamily[localFamily]; addr != "" {
			return addr, nil
		}
	}

	// JD fallback via shared helper
	if signer, ok := signerFromJDIfMissing(
		nop.SignerAddressByFamily,
		alias,
		localFamily,
		signingKeysByNOP,
	); ok {
		e.Logger.Debugw("Using signing address from JD",
			"nopAlias", alias,
			"chainFamily", localFamily,
			"signerAddress", signer,
		)
		return signer, nil
	}

	return "", fmt.Errorf(
		"NOP %q missing signer_address for family %s on committee %q chain %d",
		alias, localFamily, committeeQualifier, remoteSelector,
	)
}
