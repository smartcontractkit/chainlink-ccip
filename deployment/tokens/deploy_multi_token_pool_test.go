package tokens_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

const multiTokenPoolTestChainSelector = uint64(5009297550715157269)

type mockMultiTokenPoolAdapter struct {
	sequence *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

func (m *mockMultiTokenPoolAdapter) DeployMultiTokenPool() *cldf_ops.Sequence[tokens.DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return m.sequence
}

type mockMultiTokenPoolMCMSReader struct {
	timelockRef datastore.AddressRef
}

func (m *mockMultiTokenPoolMCMSReader) GetChainMetadata(cldf.Environment, uint64, mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{}, nil
}

func (m *mockMultiTokenPoolMCMSReader) GetTimelockRef(cldf.Environment, uint64, mcms.Input) (datastore.AddressRef, error) {
	return m.timelockRef, nil
}

func (m *mockMultiTokenPoolMCMSReader) GetMCMSRef(cldf.Environment, uint64, mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{}, nil
}

func newMockMultiTokenPoolAdapter(captured *tokens.DeployMultiTokenPoolInput) *mockMultiTokenPoolAdapter {
	return &mockMultiTokenPoolAdapter{sequence: cldf_ops.NewSequence(
		"mock-deploy-multi-token-pool",
		semver.MustParse("2.1.0"),
		"Captures multi-token pool deployment input",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input tokens.DeployMultiTokenPoolInput) (sequences.OnChainOutput, error) {
			*captured = input
			return sequences.OnChainOutput{Addresses: []datastore.AddressRef{{
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(input.PoolType),
				Version:       input.PoolVersion,
				Qualifier:     input.PoolQualifier,
				Address:       "0x0000000000000000000000000000000000000001",
			}}}, nil
		},
	)}
}

func TestMultiTokenPoolAdapterRegistry(t *testing.T) {
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	first := &mockMultiTokenPoolAdapter{}
	second := &mockMultiTokenPoolAdapter{}
	version := semver.MustParse("2.1.0")

	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, version, first)
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, version, second)

	got, ok := registry.GetMultiTokenPoolAdapter(chain_selectors.FamilyEVM, version)
	require.True(t, ok)
	require.Same(t, first, got, "the first registration must win")

	got, ok = registry.GetMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.1"))
	require.True(t, ok)
	require.Same(t, first, got)

	_, ok = registry.GetMultiTokenPoolAdapter(chain_selectors.FamilyEVM, nil)
	require.False(t, ok)
	_, ok = registry.GetMultiTokenPoolAdapter(chain_selectors.FamilySolana, version)
	require.False(t, ok)
}

func TestDeployMultiTokenPoolsVerifyPreconditions(t *testing.T) {
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), &mockMultiTokenPoolAdapter{})
	validChainCfg := tokens.DeployMultiTokenPoolChainConfig{
		PoolType:      "BurnMintMultiTokenPool",
		PoolVersion:   semver.MustParse("2.1.0"),
		PoolQualifier: "test-pool",
		Admin:         "0x0000000000000000000000000000000000000002",
	}

	tests := []struct {
		name        string
		cfg         tokens.DeployMultiTokenPoolsConfig
		expectedErr string
	}{
		{name: "valid", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: validChainCfg}}},
		{name: "no chains", cfg: tokens.DeployMultiTokenPoolsConfig{}, expectedErr: "at least one chain is required"},
		{name: "invalid MCMS", cfg: tokens.DeployMultiTokenPoolsConfig{MCMS: mcms.Input{TimelockAction: "invalid"}, Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: validChainCfg}}, expectedErr: "invalid MCMS input"},
		{name: "missing MCMS reader for default admin", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolVersion: validChainCfg.PoolVersion, PoolQualifier: validChainCfg.PoolQualifier}}}, expectedErr: "no MCMS reader registered"},
		{name: "missing pool type", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolVersion: validChainCfg.PoolVersion, PoolQualifier: validChainCfg.PoolQualifier}}}, expectedErr: "poolType is required"},
		{name: "missing pool version", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolQualifier: validChainCfg.PoolQualifier}}}, expectedErr: "poolVersion is required"},
		{name: "missing qualifier", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolVersion: validChainCfg.PoolVersion}}}, expectedErr: "poolQualifier is required"},
		{name: "invalid selector", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{1: validChainCfg}}, expectedErr: "invalid chain selector 1"},
		{name: "missing adapter", cfg: tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{multiTokenPoolTestChainSelector: {PoolType: validChainCfg.PoolType, PoolVersion: semver.MustParse("3.0.0"), PoolQualifier: validChainCfg.PoolQualifier}}}, expectedErr: "no multi-token pool adapter registered"},
	}

	changeset := tokens.DeployMultiTokenPools(registry, changesets.NewMCMSReaderRegistry())
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := changeset.VerifyPreconditions(cldf.Environment{}, test.cfg)
			if test.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, test.expectedErr)
			}
		})
	}
}

func TestDeployMultiTokenPoolsApply(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{multiTokenPoolTestChainSelector}))
	require.NoError(t, err)

	var captured tokens.DeployMultiTokenPoolInput
	adapter := newMockMultiTokenPoolAdapter(&captured)
	registry := tokens.NewMultiTokenPoolAdapterRegistry()
	registry.RegisterMultiTokenPoolAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.1.0"), adapter)

	timelockAddress := "0x0000000000000000000000000000000000000002"
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, &mockMultiTokenPoolMCMSReader{timelockRef: datastore.AddressRef{
		ChainSelector: multiTokenPoolTestChainSelector,
		Address:       timelockAddress,
	}})

	changeset := tokens.DeployMultiTokenPools(registry, mcmsRegistry)
	cfg := tokens.DeployMultiTokenPoolsConfig{Chains: map[uint64]tokens.DeployMultiTokenPoolChainConfig{
		multiTokenPoolTestChainSelector: {
			PoolType:      "BurnMintMultiTokenPool",
			PoolVersion:   semver.MustParse("2.1.0"),
			PoolQualifier: "test-pool",
		},
	}}
	require.NoError(t, changeset.VerifyPreconditions(*e, cfg))
	out, err := changeset.Apply(*e, cfg)
	require.NoError(t, err)

	require.Equal(t, multiTokenPoolTestChainSelector, captured.ChainSelector)
	require.Equal(t, timelockAddress, captured.Admin, "an omitted admin defaults to the selected timelock")
	require.Equal(t, e.DataStore, captured.ExistingDataStore)
	refs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, refs, 1)
	require.Equal(t, "test-pool", refs[0].Qualifier)
}
