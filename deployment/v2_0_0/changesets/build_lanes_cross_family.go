package changesets

import (
	"fmt"
	"slices"
	"sort"
	"strconv"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

// ChainOverrides holds per-chain lane settings that differ from chain-family adapter defaults.
// Only set fields you need to override; omitted fields use adapter defaults at apply time.
type ChainOverrides struct {
	AllowlistEnabled                *bool                    `json:"allowlistEnabled,omitempty" yaml:"allowlistEnabled,omitempty"`
	AllowList                       []string                 `json:"allowList,omitempty" yaml:"allowList,omitempty"`
	CommitteeVerifierFinalityConfig *finality.Config         `json:"committeeVerifierFinalityConfig,omitempty" yaml:"committeeVerifierFinalityConfig,omitempty"`
	RemoteChainCfg                  PartialRemoteChainConfig `json:"remoteChainCfg,omitempty" yaml:"remoteChainCfg,omitempty"`
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
func expandLanesToPartialChainConfigs(lanes []CrossFamilyLanePair, committees map[string]offchain.CommitteeConfig) ([]partialChainConfig, error) {
	if len(lanes) == 0 {
		return nil, fmt.Errorf("at least one lane must be specified")
	}

	byChain := make(map[uint64]*partialChainConfig)
	for i, lane := range lanes {
		if lane.ChainA == 0 || lane.ChainB == 0 {
			return nil, fmt.Errorf("lane %d: chainA and chainB are required", i)
		}
		if lane.ChainA == lane.ChainB {
			return nil, fmt.Errorf("lane %d: chainA and chainB must differ", i)
		}
		var qualifiers []string
		remoteKey := strconv.FormatUint(lane.ChainB, 10)
		for _, cvCfg := range committees {
			if _, exists := cvCfg.ChainConfigs[remoteKey]; exists {
				qualifiers = append(qualifiers, cvCfg.Qualifier)
			}
		}
		sort.Strings(qualifiers)
		if len(qualifiers) == 0 {
			if len(committees) == 0 {
				qualifiers = []string{defaultQualifier}
			} else {
				return nil, fmt.Errorf("lane %d: no committees have chain_config for remote chain %d", i, lane.ChainB)
			}
		}
		mergeLaneLeg(byChain, lane.ChainA, lane.ChainB, qualifiers, lane.ChainAOverrides)
		qualifiers = nil
		remoteKey = strconv.FormatUint(lane.ChainA, 10)
		for _, cvCfg := range committees {
			if _, exists := cvCfg.ChainConfigs[remoteKey]; exists {
				qualifiers = append(qualifiers, cvCfg.Qualifier)
			}
		}
		sort.Strings(qualifiers)
		if len(qualifiers) == 0 {
			if len(committees) == 0 {
				qualifiers = []string{defaultQualifier}
			} else {
				return nil, fmt.Errorf("lane %d: no committees have chain_config for remote chain %d", i, lane.ChainA)
			}
		}
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

func mergeLaneLeg(byChain map[uint64]*partialChainConfig, local, remote uint64, qualifiers []string, o *ChainOverrides) {
	cfg, ok := byChain[local]
	if !ok {
		cvConfigs := make([]committeeVerifierInputConfig, 0, len(qualifiers))
		for _, q := range qualifiers {
			cv := committeeVerifierInputConfig{
				CommitteeQualifier: q,
				RemoteChains:       make(map[uint64]committeeVerifierRemoteChainInput),
			}
			if o != nil {
				cv.AllowedFinalityConfig = o.CommitteeVerifierFinalityConfig
			}
			cvConfigs = append(cvConfigs, cv)
		}

		cfg = &partialChainConfig{
			ChainSelector:      local,
			CommitteeVerifiers: cvConfigs,
			RemoteChains:       make(map[uint64]PartialRemoteChainConfig),
		}
		byChain[local] = cfg
	}

	cv := chainOverridesToCommitteeVerifierRemote(o)
	rc := chainOverridesToPartialRemote(o)

	// Ensure we have a verifier entry per qualifier, and only attach this remote chain to those qualifiers.
	byQualifier := make(map[string]int, len(cfg.CommitteeVerifiers))
	for i := range cfg.CommitteeVerifiers {
		byQualifier[cfg.CommitteeVerifiers[i].CommitteeQualifier] = i
	}
	for _, q := range qualifiers {
		idx, ok := byQualifier[q]
		if !ok {
			newCv := committeeVerifierInputConfig{
				CommitteeQualifier: q,
				RemoteChains:       make(map[uint64]committeeVerifierRemoteChainInput),
			}
			if o != nil {
				newCv.AllowedFinalityConfig = o.CommitteeVerifierFinalityConfig
			}
			cfg.CommitteeVerifiers = append(cfg.CommitteeVerifiers, newCv)
			idx = len(cfg.CommitteeVerifiers) - 1
			byQualifier[q] = idx
		}
		cfgCv := cfg.CommitteeVerifiers[idx]
		if o != nil && o.CommitteeVerifierFinalityConfig != nil {
			cfgCv.AllowedFinalityConfig = o.CommitteeVerifierFinalityConfig
		}
		if existingRemote, ok := cfgCv.RemoteChains[remote]; ok {
			cfgCv.RemoteChains[remote] = mergeCommitteeVerifierRemoteInput(existingRemote, cv)
		} else {
			cfgCv.RemoteChains[remote] = cv
		}
		cfg.CommitteeVerifiers[idx] = cfgCv
	}
	sort.Slice(cfg.CommitteeVerifiers, func(i, j int) bool {
		return cfg.CommitteeVerifiers[i].CommitteeQualifier < cfg.CommitteeVerifiers[j].CommitteeQualifier
	})

	if existing, ok := cfg.RemoteChains[remote]; ok {
		cfg.RemoteChains[remote] = mergePartialRemoteInput(existing, rc)
	} else {
		cfg.RemoteChains[remote] = rc
	}
}

func chainOverridesToPartialRemote(o *ChainOverrides) PartialRemoteChainConfig {
	if o == nil {
		return PartialRemoteChainConfig{}
	}
	return o.RemoteChainCfg
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

func mergePartialRemoteInput(base, overlay PartialRemoteChainConfig) PartialRemoteChainConfig {
	if overlay.AllowTrafficFrom != nil {
		base.AllowTrafficFrom = overlay.AllowTrafficFrom
	}
	if overlay.MessageNetworkFeeUSDCents != nil {
		base.MessageNetworkFeeUSDCents = overlay.MessageNetworkFeeUSDCents
	}
	if overlay.TokenNetworkFeeUSDCents != nil {
		base.TokenNetworkFeeUSDCents = overlay.TokenNetworkFeeUSDCents
	}
	if overlay.TokenReceiverAllowed != nil {
		base.TokenReceiverAllowed = overlay.TokenReceiverAllowed
	}
	if overlay.DefaultExecutorQualifier != nil {
		base.DefaultExecutorQualifier = overlay.DefaultExecutorQualifier
	}
	if overlay.BaseExecutionGasCost != nil {
		base.BaseExecutionGasCost = overlay.BaseExecutionGasCost
	}
	base.FeeQuoterDestChainConfig = mergeFeeQuoterDestChainConfig(base.FeeQuoterDestChainConfig, overlay.FeeQuoterDestChainConfig)
	base.ExecutorDestChainConfig = utils.CoalescePtr(overlay.ExecutorDestChainConfig, base.ExecutorDestChainConfig)
	if overlay.DefaultInboundCCVs != nil {
		base.DefaultInboundCCVs = overlay.DefaultInboundCCVs
	}
	if overlay.DefaultOutboundCCVs != nil {
		base.DefaultOutboundCCVs = overlay.DefaultOutboundCCVs
	}
	if overlay.LaneMandatedOutboundCCVs != nil {
		base.LaneMandatedOutboundCCVs = overlay.LaneMandatedOutboundCCVs
	}
	if overlay.LaneMandatedInboundCCVs != nil {
		base.LaneMandatedInboundCCVs = overlay.LaneMandatedInboundCCVs
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
