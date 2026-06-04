package changesets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
	committees := map[string]offchain.CommitteeConfig{
		"alpha": {},
		"beta":  {},
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
