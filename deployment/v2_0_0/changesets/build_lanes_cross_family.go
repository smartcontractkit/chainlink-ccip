package changesets

import (
	"fmt"
	"slices"
	"sort"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

// ChainOverrides holds per-chain lane settings that differ from chain-family adapter defaults.
// Only set fields you need to override; omitted fields use adapter defaults at apply time.
type ChainOverrides struct {
	AllowTrafficFrom          *bool    `json:"allowTrafficFrom,omitempty" yaml:"allowTrafficFrom,omitempty"`
	MessageNetworkFeeUSDCents *uint16  `json:"messageNetworkFeeUSDCents,omitempty" yaml:"messageNetworkFeeUSDCents,omitempty"`
	TokenNetworkFeeUSDCents   *uint16  `json:"tokenNetworkFeeUSDCents,omitempty" yaml:"tokenNetworkFeeUSDCents,omitempty"`
	AllowlistEnabled          *bool    `json:"allowlistEnabled,omitempty" yaml:"allowlistEnabled,omitempty"`
	AllowList                 []string `json:"allowList,omitempty" yaml:"allowList,omitempty"`
}

// CrossFamilyLanePair defines a bidirectional lane with optional per-chain overrides.
// ChainAOverrides apply on chain A when configuring the leg to chain B (and committee verifier
// remote chain B). ChainBOverrides apply symmetrically on chain B toward chain A.
type CrossFamilyLanePair struct {
	ChainA          uint64          `json:"chainA" yaml:"chainA"`
	ChainB          uint64          `json:"chainB" yaml:"chainB"`
	ChainAOverrides *ChainOverrides `json:"chainAOverrides,omitempty" yaml:"chainAOverrides,omitempty"`
	ChainBOverrides *ChainOverrides `json:"chainBOverrides,omitempty" yaml:"chainBOverrides,omitempty"`
}

// BuildLanesCrossFamilyConfig is the durable-pipeline payload for build_lanes_cross_family
// on the CCV domain. Same top-level YAML shape as CCIP 1.6 ConnectChains lane list.
//
// Override priority (highest to lowest):
//  1. YAML ChainOverrides on CrossFamilyLanePair (per-lane user input)
//  2. Chain-family adapter defaults (applied inside ConfigureChainsForLanesFromTopology)
type BuildLanesCrossFamilyConfig struct {
	Lanes      []CrossFamilyLanePair `json:"lanes" yaml:"lanes"`
	MCMS       mcms.Input            `json:"mcms" yaml:"mcms"`
	TestRouter *bool                 `json:"testRouter,omitempty" yaml:"testRouter,omitempty"`
}

// UseTestRouter reports whether the test router should be used instead of the production router.
func (c BuildLanesCrossFamilyConfig) UseTestRouter() bool {
	return c.TestRouter != nil && *c.TestRouter
}

// ConfigureChainsForLanesFromTopologyConfig configures CCIP 2.0 lanes from topology plus
// bidirectional lane pairs. Contract addresses, CCVs, executors, and fee-quoter defaults are
// resolved automatically; only ChainOverrides fields need to be set when defaults are wrong.
type ConfigureChainsForLanesFromTopologyConfig struct {
	Topology *offchain.EnvironmentTopology `json:"topology" yaml:"topology"`
	BuildLanesCrossFamilyConfig
}

// expandLanesToPartialChainConfigs converts bidirectional lane pairs into the internal
// chain-centric representation used by the changeset apply path.
func expandLanesToPartialChainConfigs(lanes []CrossFamilyLanePair, topology *offchain.NOPTopology) ([]partialChainConfig, error) {
	if len(lanes) == 0 {
		return nil, fmt.Errorf("at least one lane must be specified")
	}

	qualifiers := make([]string, 0, len(topology.Committees))
	for qualifier := range topology.Committees {
		qualifiers = append(qualifiers, qualifier)
	}

	byChain := make(map[uint64]*partialChainConfig)
	for i, lane := range lanes {
		if lane.ChainA == 0 || lane.ChainB == 0 {
			return nil, fmt.Errorf("lane %d: chainA and chainB are required", i)
		}
		if lane.ChainA == lane.ChainB {
			return nil, fmt.Errorf("lane %d: chainA and chainB must differ", i)
		}

		mergeLaneLeg(byChain, lane.ChainA, lane.ChainB, qualifiers, lane.ChainAOverrides)
		mergeLaneLeg(byChain, lane.ChainB, lane.ChainA, qualifiers, lane.ChainBOverrides)
	}

	out := make([]partialChainConfig, 0, len(byChain))
	for _, sel := range sortedKeys(byChain) {
		out = append(out, *byChain[sel])
	}
	return out, nil
}

// filterPartialChainsToEnvironment keeps only chains deployed in the environment.
// Remote counterparts may be absent from BlockChains; they are still referenced on each
// local chain's RemoteChains map and resolved via the remote chain family adapter.
func filterPartialChainsToEnvironment(blockChains cldf_chain.BlockChains, chains []partialChainConfig) []partialChainConfig {
	available := blockChains.ListChainSelectors()
	filtered := make([]partialChainConfig, 0, len(chains))
	for _, chainCfg := range chains {
		if !slices.Contains(available, chainCfg.ChainSelector) {
			continue
		}
		filtered = append(filtered, chainCfg)
	}
	return filtered
}

func sortedKeys(m map[uint64]*partialChainConfig) []uint64 {
	keys := make([]uint64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}

func mergeLaneLeg(byChain map[uint64]*partialChainConfig, local, remote uint64, qualifiers[]string, o *ChainOverrides) {
	cfg, ok := byChain[local]
	if !ok {
		cvConfigs := make([]committeeVerifierInputConfig, 0, len(qualifiers))
		for _, qualifier := range qualifiers {
			cvConfigs = append(cvConfigs, committeeVerifierInputConfig{
				CommitteeQualifier: defaultQualifier,
				RemoteChains:       make(map[uint64]committeeVerifierRemoteChainInput),
				// TODO: AllowedFinalityConfig: local.AllowedFinalityConfig, sourced from where?
			}),
		}

		cfg = &partialChainConfig{
			ChainSelector: local,
			CommitteeVerifiers: cvConfigs,
			RemoteChains: make(map[uint64]partialRemoteChainConfig),
		}
		byChain[local] = cfg
	}

	cv := chainOverridesToCommitteeVerifierRemote(o)
	rc := chainOverridesToPartialRemote(o)

	if existing, ok := cfg.CommitteeVerifiers[0].RemoteChains[remote]; ok {
		cfg.CommitteeVerifiers[0].RemoteChains[remote] = mergeCommitteeVerifierRemoteInput(existing, cv)
	} else {
		cfg.CommitteeVerifiers[0].RemoteChains[remote] = cv
	}

	if existing, ok := cfg.RemoteChains[remote]; ok {
		cfg.RemoteChains[remote] = mergePartialRemoteInput(existing, rc)
	} else {
		cfg.RemoteChains[remote] = rc
	}
}

func chainOverridesToPartialRemote(o *ChainOverrides) partialRemoteChainConfig {
	if o == nil {
		return partialRemoteChainConfig{}
	}
	return partialRemoteChainConfig{
		AllowTrafficFrom:          o.AllowTrafficFrom,
		MessageNetworkFeeUSDCents: o.MessageNetworkFeeUSDCents,
		TokenNetworkFeeUSDCents:   o.TokenNetworkFeeUSDCents,
	}
}

func chainOverridesToCommitteeVerifierRemote(o *ChainOverrides) committeeVerifierRemoteChainInput {
	if o == nil {
		return committeeVerifierRemoteChainInput{}
	}
	return committeeVerifierRemoteChainInput{
		AllowlistEnabled:        o.AllowlistEnabled,
		AddedAllowlistedSenders: append([]string(nil), o.AllowList...),
	}
}

func mergePartialRemoteInput(base, overlay partialRemoteChainConfig) partialRemoteChainConfig {
	if overlay.AllowTrafficFrom != nil {
		base.AllowTrafficFrom = overlay.AllowTrafficFrom
	}
	if overlay.MessageNetworkFeeUSDCents != nil {
		base.MessageNetworkFeeUSDCents = overlay.MessageNetworkFeeUSDCents
	}
	if overlay.TokenNetworkFeeUSDCents != nil {
		base.TokenNetworkFeeUSDCents = overlay.TokenNetworkFeeUSDCents
	}
	return base
}

func mergeCommitteeVerifierRemoteInput(base, overlay committeeVerifierRemoteChainInput) committeeVerifierRemoteChainInput {
	if overlay.AllowlistEnabled != nil {
		base.AllowlistEnabled = overlay.AllowlistEnabled
	}
	if len(overlay.AddedAllowlistedSenders) > 0 {
		base.AddedAllowlistedSenders = overlay.AddedAllowlistedSenders
	}
	if len(overlay.RemovedAllowlistedSenders) > 0 {
		base.RemovedAllowlistedSenders = overlay.RemovedAllowlistedSenders
	}
	if overlay.FeeUSDCents != nil {
		base.FeeUSDCents = overlay.FeeUSDCents
	}
	if overlay.GasForVerification != nil {
		base.GasForVerification = overlay.GasForVerification
	}
	if overlay.PayloadSizeBytes != nil {
		base.PayloadSizeBytes = overlay.PayloadSizeBytes
	}
	return base
}
