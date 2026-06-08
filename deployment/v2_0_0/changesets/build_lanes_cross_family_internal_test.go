package changesets

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

func TestExpandLanesToPartialChainConfigs_MinimalLane(t *testing.T) {
	chainA := uint64(90000001)
	chainB := uint64(90000002)

	chains, err := expandLanesToPartialChainConfigs([]CrossFamilyLanePair{
		{ChainA: chainA, ChainB: chainB},
	}, nil)
	require.NoError(t, err)
	require.Len(t, chains, 2)

	bySel := make(map[uint64]partialChainConfig, len(chains))
	for _, c := range chains {
		bySel[c.ChainSelector] = c
	}

	cfgA := bySel[chainA]
	require.Len(t, cfgA.CommitteeVerifiers, 1)
	assert.Equal(t, defaultQualifier, cfgA.CommitteeVerifiers[0].CommitteeQualifier)
	assert.Contains(t, cfgA.CommitteeVerifiers[0].RemoteChains, chainB)
	assert.Contains(t, cfgA.RemoteChains, chainB)

	cfgB := bySel[chainB]
	assert.Contains(t, cfgB.CommitteeVerifiers[0].RemoteChains, chainA)
	assert.Contains(t, cfgB.RemoteChains, chainA)
}

func TestExpandLanesToPartialChainConfigs_AppliesChainOverrides(t *testing.T) {
	chainA := uint64(1)
	chainB := uint64(2)
	fee := uint16(99)

	chains, err := expandLanesToPartialChainConfigs([]CrossFamilyLanePair{
		{
			ChainA: chainA,
			ChainB: chainB,
			ChainAOverrides: &ChainOverrides{
				RemoteChainCfg: PartialRemoteChainConfig{
					MessageNetworkFeeUSDCents: &fee,
				},
			},
		},
	}, nil)
	require.NoError(t, err)

	var cfgA partialChainConfig
	for _, c := range chains {
		if c.ChainSelector == chainA {
			cfgA = c
		}
	}
	require.Equal(t, fee, *cfgA.RemoteChains[chainB].MessageNetworkFeeUSDCents)
}

func TestExpandLanesToPartialChainConfigs_MergesMultipleLanesOnSameChain(t *testing.T) {
	chainA := uint64(1)
	chainB := uint64(2)
	chainC := uint64(3)

	chains, err := expandLanesToPartialChainConfigs([]CrossFamilyLanePair{
		{ChainA: chainA, ChainB: chainB},
		{ChainA: chainA, ChainB: chainC},
	}, nil)
	require.NoError(t, err)

	var cfgA partialChainConfig
	for _, c := range chains {
		if c.ChainSelector == chainA {
			cfgA = c
		}
	}
	assert.Len(t, cfgA.RemoteChains, 2)
	assert.Contains(t, cfgA.RemoteChains, chainB)
	assert.Contains(t, cfgA.RemoteChains, chainC)
}

func TestExpandLanesToPartialChainConfigs_MultipleCommittees(t *testing.T) {
	chainA := uint64(1)
	chainB := uint64(2)
	chainConfigsBoth := map[string]offchain.ChainCommitteeConfig{
		strconv.FormatUint(chainA, 10): {},
		strconv.FormatUint(chainB, 10): {},
	}
	committees := map[string]offchain.CommitteeConfig{
		"alpha": {Qualifier: "alpha", ChainConfigs: chainConfigsBoth},
		"beta":  {Qualifier: "beta", ChainConfigs: chainConfigsBoth},
	}

	chains, err := expandLanesToPartialChainConfigs([]CrossFamilyLanePair{
		{ChainA: chainA, ChainB: chainB},
	}, committees)
	require.NoError(t, err)
	require.Len(t, chains, 2)

	bySel := make(map[uint64]partialChainConfig, len(chains))
	for _, c := range chains {
		bySel[c.ChainSelector] = c
	}

	cfgA := bySel[chainA]
	require.Len(t, cfgA.CommitteeVerifiers, 2)

	// Qualifiers must be sorted deterministically.
	assert.Equal(t, "alpha", cfgA.CommitteeVerifiers[0].CommitteeQualifier)
	assert.Equal(t, "beta", cfgA.CommitteeVerifiers[1].CommitteeQualifier)

	// Each verifier must have the remote chain entry.
	assert.Contains(t, cfgA.CommitteeVerifiers[0].RemoteChains, chainB)
	assert.Contains(t, cfgA.CommitteeVerifiers[1].RemoteChains, chainB)

	cfgB := bySel[chainB]
	require.Len(t, cfgB.CommitteeVerifiers, 2)
	assert.Equal(t, "alpha", cfgB.CommitteeVerifiers[0].CommitteeQualifier)
	assert.Equal(t, "beta", cfgB.CommitteeVerifiers[1].CommitteeQualifier)
	assert.Contains(t, cfgB.CommitteeVerifiers[0].RemoteChains, chainA)
	assert.Contains(t, cfgB.CommitteeVerifiers[1].RemoteChains, chainA)
}

func TestExpandLanesToPartialChainConfigs_FiltersCommitteeQualifiersPerRemoteChain(t *testing.T) {
	chainA := uint64(100)
	chainB := uint64(200)

	// alpha is on both chains; beta only on chainA; gamma only on chainB.
	committees := map[string]offchain.CommitteeConfig{
		"alpha": {
			Qualifier: "alpha",
			ChainConfigs: map[string]offchain.ChainCommitteeConfig{
				strconv.FormatUint(chainA, 10): {},
				strconv.FormatUint(chainB, 10): {},
			},
		},
		"beta": {
			Qualifier: "beta",
			ChainConfigs: map[string]offchain.ChainCommitteeConfig{
				strconv.FormatUint(chainA, 10): {},
			},
		},
		"gamma": {
			Qualifier: "gamma",
			ChainConfigs: map[string]offchain.ChainCommitteeConfig{
				strconv.FormatUint(chainB, 10): {},
			},
		},
	}

	chains, err := expandLanesToPartialChainConfigs([]CrossFamilyLanePair{
		{ChainA: chainA, ChainB: chainB},
	}, committees)
	require.NoError(t, err)

	bySel := make(map[uint64]partialChainConfig, len(chains))
	for _, c := range chains {
		bySel[c.ChainSelector] = c
	}

	cfgA := bySel[chainA]
	require.Len(t, cfgA.CommitteeVerifiers, 2)
	assert.Equal(t, "alpha", cfgA.CommitteeVerifiers[0].CommitteeQualifier)
	assert.Equal(t, "gamma", cfgA.CommitteeVerifiers[1].CommitteeQualifier)
	for _, cv := range cfgA.CommitteeVerifiers {
		assert.Contains(t, cv.RemoteChains, chainB)
	}

	cfgB := bySel[chainB]
	require.Len(t, cfgB.CommitteeVerifiers, 2)
	assert.Equal(t, "alpha", cfgB.CommitteeVerifiers[0].CommitteeQualifier)
	assert.Equal(t, "beta", cfgB.CommitteeVerifiers[1].CommitteeQualifier)
	for _, cv := range cfgB.CommitteeVerifiers {
		assert.Contains(t, cv.RemoteChains, chainA)
	}
}

func TestExpandLanesToPartialChainConfigs_ChainOverridesCommitteeVerifierFinalityConfig(t *testing.T) {
	finalityCfg := &finality.Config{WaitForSafe: true, BlockDepth: 100}

	chains, err := expandLanesToPartialChainConfigs([]CrossFamilyLanePair{
		{
			ChainA: 10,
			ChainB: 20,
			ChainAOverrides: &ChainOverrides{
				CommitteeVerifierFinalityConfig: finalityCfg,
			},
		},
	}, nil)
	require.NoError(t, err)

	var cfg partialChainConfig
	for _, c := range chains {
		if c.ChainSelector == 10 {
			cfg = c
		}
	}
	require.Len(t, cfg.CommitteeVerifiers, 1)
	require.NotNil(t, cfg.CommitteeVerifiers[0].AllowedFinalityConfig)
	assert.Equal(t, *finalityCfg, *cfg.CommitteeVerifiers[0].AllowedFinalityConfig)
}

func TestExpandLanesToPartialChainConfigs_ChainOverridesToCommitteeVerifier(t *testing.T) {
	enabled := true
	allow := []string{"0xabc"}

	chains, err := expandLanesToPartialChainConfigs([]CrossFamilyLanePair{
		{
			ChainA: 10,
			ChainB: 20,
			ChainAOverrides: &ChainOverrides{
				AllowlistEnabled: &enabled,
				AllowList:        allow,
			},
		},
	}, nil)
	require.NoError(t, err)

	var cfg partialChainConfig
	for _, c := range chains {
		if c.ChainSelector == 10 {
			cfg = c
		}
	}
	rc := cfg.CommitteeVerifiers[0].RemoteChains[20]
	require.NotNil(t, rc.AllowlistEnabled)
	assert.True(t, *rc.AllowlistEnabled)
	assert.Equal(t, allow, rc.AddedAllowlistedSenders)
}

func TestMergeLaneLeg(t *testing.T) {
	local := uint64(10)
	remoteB := uint64(20)
	remoteC := uint64(30)

	t.Run("multiple remote same qualifier", func(t *testing.T) {
		byChain := make(map[uint64]*partialChainConfig)
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha"}, nil)
		mergeLaneLeg(byChain, local, remoteC, []string{"alpha"}, nil)

		cfg := byChain[local]
		require.NotNil(t, cfg)
		assert.Equal(t, local, cfg.ChainSelector)
		require.Len(t, cfg.CommitteeVerifiers, 1)
		assert.Equal(t, "alpha", cfg.CommitteeVerifiers[0].CommitteeQualifier)
		assert.Contains(t, cfg.CommitteeVerifiers[0].RemoteChains, remoteB)
		assert.Contains(t, cfg.CommitteeVerifiers[0].RemoteChains, remoteC)
		assert.Contains(t, cfg.RemoteChains, remoteB)
		assert.Contains(t, cfg.RemoteChains, remoteC)
	})

	t.Run("initializes local chain with qualifiers and remote", func(t *testing.T) {
		byChain := make(map[uint64]*partialChainConfig)
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha", "beta"}, nil)

		cfg := byChain[local]
		require.NotNil(t, cfg)
		assert.Equal(t, local, cfg.ChainSelector)
		require.Len(t, cfg.CommitteeVerifiers, 2)
		assert.Equal(t, "alpha", cfg.CommitteeVerifiers[0].CommitteeQualifier)
		assert.Equal(t, "beta", cfg.CommitteeVerifiers[1].CommitteeQualifier)
		assert.Contains(t, cfg.CommitteeVerifiers[0].RemoteChains, remoteB)
		assert.Contains(t, cfg.CommitteeVerifiers[1].RemoteChains, remoteB)
		assert.Contains(t, cfg.RemoteChains, remoteB)
	})

	t.Run("adds remote only to qualifiers in second merge", func(t *testing.T) {
		byChain := make(map[uint64]*partialChainConfig)
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha", "beta"}, nil)
		mergeLaneLeg(byChain, local, remoteC, []string{"alpha"}, nil)

		cfg := byChain[local]
		assert.Contains(t, cfg.CommitteeVerifiers[0].RemoteChains, remoteB)
		assert.Contains(t, cfg.CommitteeVerifiers[0].RemoteChains, remoteC)
		assert.Contains(t, cfg.CommitteeVerifiers[1].RemoteChains, remoteB)
		assert.NotContains(t, cfg.CommitteeVerifiers[1].RemoteChains, remoteC)
		assert.Contains(t, cfg.RemoteChains, remoteC)
		assert.Contains(t, cfg.RemoteChains, remoteB)
	})

	t.Run("adds remote only to new qualifiers in second merge", func(t *testing.T) {
		byChain := make(map[uint64]*partialChainConfig)
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha"}, nil)
		mergeLaneLeg(byChain, local, remoteC, []string{"beta"}, nil)

		cfg := byChain[local]
		assert.Contains(t, cfg.CommitteeVerifiers[0].RemoteChains, remoteB)
		assert.NotContains(t, cfg.CommitteeVerifiers[0].RemoteChains, remoteC)
		assert.NotContains(t, cfg.CommitteeVerifiers[1].RemoteChains, remoteB)
		assert.Contains(t, cfg.CommitteeVerifiers[1].RemoteChains, remoteC)
		assert.Contains(t, cfg.RemoteChains, remoteC)
		assert.Contains(t, cfg.RemoteChains, remoteB)
	})

	t.Run("appends new qualifier on existing local chain", func(t *testing.T) {
		byChain := make(map[uint64]*partialChainConfig)
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha"}, nil)
		mergeLaneLeg(byChain, local, remoteB, []string{"beta"}, nil)

		cfg := byChain[local]
		require.Len(t, cfg.CommitteeVerifiers, 2)
		assert.Equal(t, "alpha", cfg.CommitteeVerifiers[0].CommitteeQualifier)
		assert.Equal(t, "beta", cfg.CommitteeVerifiers[1].CommitteeQualifier)
	})

	t.Run("applies committee verifier finality config from chain overrides", func(t *testing.T) {
		byChain := make(map[uint64]*partialChainConfig)
		finalityCfg := &finality.Config{WaitForSafe: true, BlockDepth: 50}
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha"}, &ChainOverrides{
			CommitteeVerifierFinalityConfig: finalityCfg,
		})

		cfg := byChain[local]
		require.NotNil(t, cfg.CommitteeVerifiers[0].AllowedFinalityConfig)
		assert.Equal(t, *finalityCfg, *cfg.CommitteeVerifiers[0].AllowedFinalityConfig)
	})

	t.Run("merges chain overrides into committee verifier and remote config", func(t *testing.T) {
		enabled := true
		allow := []string{"0xabc"}
		feeFirst := uint16(10)
		feeSecond := uint16(20)

		byChain := make(map[uint64]*partialChainConfig)
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha"}, &ChainOverrides{
			AllowlistEnabled: &enabled,
			AllowList:        allow,
			RemoteChainCfg: PartialRemoteChainConfig{
				MessageNetworkFeeUSDCents: &feeFirst,
			},
		})
		mergeLaneLeg(byChain, local, remoteB, []string{"alpha"}, &ChainOverrides{
			RemoteChainCfg: PartialRemoteChainConfig{
				MessageNetworkFeeUSDCents: &feeSecond,
			},
		})

		cfg := byChain[local]
		cv := cfg.CommitteeVerifiers[0].RemoteChains[remoteB]
		require.NotNil(t, cv.AllowlistEnabled)
		assert.True(t, *cv.AllowlistEnabled)
		assert.Equal(t, allow, cv.AddedAllowlistedSenders)
		require.Equal(t, feeSecond, *cfg.RemoteChains[remoteB].MessageNetworkFeeUSDCents)
	})
}
