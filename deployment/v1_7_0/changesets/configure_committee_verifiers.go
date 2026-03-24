package changesets

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/Masterminds/semver/v3"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/operations/fetch_signing_keys"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type CommitteeVerifierRemoteChainConfig struct {
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []string
	RemovedAllowlistedSenders []string
	FeeUSDCents               uint16
	GasForVerification        uint32
	PayloadSizeBytes          uint32
}

type CommitteeVerifierInputConfig struct {
	CommitteeQualifier string
	RemoteChains       map[uint64]CommitteeVerifierRemoteChainConfig
}

type PartialChainConfig struct {
	ChainSelector            uint64
	CommitteeVerifiers       []CommitteeVerifierInputConfig
	DefaultInboundCCVs       []datastore.AddressRef
	LaneMandatedInboundCCVs  []datastore.AddressRef
	DefaultOutboundCCVs      []datastore.AddressRef
	LaneMandatedOutboundCCVs []datastore.AddressRef
	DefaultExecutor          datastore.AddressRef
	FeeQuoterDestChainConfigOverrides *lanes.FeeQuoterDestChainConfigOverride
	ExecutorDestChainConfig  lanes.ExecutorDestChainConfig
	AddressBytesLength       uint8
	BaseExecutionGasCost     uint32
	RemoteChains             map[uint64]RemoteLaneConfig
}

type RemoteLaneConfig struct {
	Chain      lanes.ChainDefinition
	TestRouter bool
}

type ConfigureChainsForLanesFromTopologyConfig struct {
	Topology *offchain.EnvironmentTopology
	Chains   []PartialChainConfig
	MCMS     mcms.Input
}

func ConfigureChainsForLanesFromTopology(
	committeeVerifierContractRegistry *adapters.CommitteeVerifierContractRegistry,
	laneAdapterRegistry *lanes.LaneAdapterRegistry,
	mcmsRegistry *changesetscore.MCMSReaderRegistry,
) deployment.ChangeSetV2[ConfigureChainsForLanesFromTopologyConfig] {
	validate := func(e deployment.Environment, cfg ConfigureChainsForLanesFromTopologyConfig) error {
		if cfg.Topology == nil {
			return fmt.Errorf("topology is required")
		}

		if cfg.Topology.NOPTopology == nil || len(cfg.Topology.NOPTopology.Committees) == 0 {
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

		signingKeysByNOP := fetchSigningKeysForNOPs(e, cfg.Topology.NOPTopology.NOPs)

		laneConfigs := make([]lanes.LaneConfig, 0)
		for _, chain := range cfg.Chains {
			committeeVerifiers := make([]lanes.CommitteeVerifierConfig[datastore.AddressRef], 0, len(chain.CommitteeVerifiers))
			for _, cv := range chain.CommitteeVerifiers {
				remoteChains := make(map[uint64]lanes.CommitteeVerifierRemoteChainConfig, len(cv.RemoteChains))
				for remoteChainSelector, remoteChainConfig := range cv.RemoteChains {
					signatureConfig, err := getSignatureConfigForLane(e, cfg.Topology, cv.CommitteeQualifier, chain.ChainSelector, remoteChainSelector, signingKeysByNOP)
					if err != nil {
						return deployment.ChangesetOutput{}, fmt.Errorf("failed to get signature config for lane local chain %d -> remote chain %d: %w", chain.ChainSelector, remoteChainSelector, err)
					}
					remoteChains[remoteChainSelector] = lanes.CommitteeVerifierRemoteChainConfig{
						AllowlistEnabled:          remoteChainConfig.AllowlistEnabled,
						AddedAllowlistedSenders:   remoteChainConfig.AddedAllowlistedSenders,
						RemovedAllowlistedSenders: remoteChainConfig.RemovedAllowlistedSenders,
						FeeUSDCents:               remoteChainConfig.FeeUSDCents,
						GasForVerification:        remoteChainConfig.GasForVerification,
						PayloadSizeBytes:          remoteChainConfig.PayloadSizeBytes,
						SignatureConfig:           *signatureConfig,
					}
				}

				adapter, err := committeeVerifierContractRegistry.GetByChain(chain.ChainSelector)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("no committee verifier contract adapter for chain %d: %w", chain.ChainSelector, err)
				}

				contracts, err := adapter.ResolveCommitteeVerifierContracts(e.DataStore, chain.ChainSelector, cv.CommitteeQualifier)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve committee verifier contracts for chain %d qualifier %q: %w", chain.ChainSelector, cv.CommitteeQualifier, err)
				}

				committeeVerifiers = append(committeeVerifiers, lanes.CommitteeVerifierConfig[datastore.AddressRef]{
					CommitteeVerifier: contracts,
					RemoteChains:      remoteChains,
				})
			}

			for remoteChainSelector, remoteLaneConfig := range chain.RemoteChains {
				remoteChainDef := remoteLaneConfig.Chain
				remoteChainDef.Selector = remoteChainSelector
				laneConfigs = append(laneConfigs, lanes.LaneConfig{
					ChainA: lanes.ChainDefinition{
						Selector:                 chain.ChainSelector,
						CommitteeVerifiers:       committeeVerifiers,
						DefaultInboundCCVs:       chain.DefaultInboundCCVs,
						LaneMandatedInboundCCVs:  chain.LaneMandatedInboundCCVs,
						DefaultOutboundCCVs:      chain.DefaultOutboundCCVs,
						LaneMandatedOutboundCCVs: chain.LaneMandatedOutboundCCVs,
						DefaultExecutor:          chain.DefaultExecutor,
						FeeQuoterDestChainConfigOverrides: chain.FeeQuoterDestChainConfigOverrides,
						ExecutorDestChainConfig:  chain.ExecutorDestChainConfig,
						AddressBytesLength:       chain.AddressBytesLength,
						BaseExecutionGasCost:     chain.BaseExecutionGasCost,
					},
					ChainB:     remoteChainDef,
					Version:    semver.MustParse("2.0.0"),
					TestRouter: remoteLaneConfig.TestRouter,
				})
			}
		}

		connectChainsCS := lanes.ConnectChains(laneAdapterRegistry, mcmsRegistry)
		connectChainsCfg := lanes.ConnectChainsConfig{
			Lanes: laneConfigs,
			MCMS:  cfg.MCMS,
		}
		if err := connectChainsCS.VerifyPreconditions(e, connectChainsCfg); err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("connect chains precondition check failed: %w", err)
		}
		return connectChainsCS.Apply(e, connectChainsCfg)
	}

	return deployment.CreateChangeSet(apply, validate)
}

func getSignatureConfigForLane(
	e deployment.Environment,
	topology *offchain.EnvironmentTopology,
	committeeQualifier string,
	localSelector uint64,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (*lanes.CommitteeVerifierSignatureQuorumConfig, error) {
	committee, ok := topology.NOPTopology.Committees[committeeQualifier]
	if !ok {
		return nil, fmt.Errorf("committee %q not found", committeeQualifier)
	}

	chainCfg, ok := committee.ChainConfigs[strconv.FormatUint(remoteSelector, 10)]
	if !ok {
		return nil, fmt.Errorf("chain selector %d not found in committee %q", remoteSelector, committeeQualifier)
	}

	localFamily, err := chainsel.GetSelectorFamily(localSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get selector family for selector %d: %w", localSelector, err)
	}

	signers := make([]string, 0, len(chainCfg.NOPAliases))
	for _, alias := range chainCfg.NOPAliases {
		signer, err := signerAddressForNOPAlias(e, topology, alias, localFamily, committeeQualifier, remoteSelector, signingKeysByNOP)
		if err != nil {
			return nil, err
		}
		signers = append(signers, signer)
	}

	return &lanes.CommitteeVerifierSignatureQuorumConfig{
		Threshold: chainCfg.Threshold,
		Signers:   signers,
	}, nil
}

func signerAddressForNOPAlias(
	e deployment.Environment,
	topology *offchain.EnvironmentTopology,
	alias string,
	localFamily string,
	committeeQualifier string,
	remoteSelector uint64,
	signingKeysByNOP fetch_signing_keys.SigningKeysByNOP,
) (string, error) {
	nop, ok := topology.NOPTopology.GetNOP(alias)
	if !ok {
		return "", fmt.Errorf(
			"NOP alias %q not found for committee %q chain %d",
			alias, committeeQualifier, remoteSelector,
		)
	}

	if nop.SignerAddressByFamily != nil {
		if addr := nop.SignerAddressByFamily[localFamily]; addr != "" {
			return addr, nil
		}
	}

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
