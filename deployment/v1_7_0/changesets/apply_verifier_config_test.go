package changesets_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

var _ adapters.VerifierConfigAdapter = (*mockVerifierJobConfigAdapter)(nil)

type mockVerifierJobConfigAdapter struct {
	chainConfigs map[uint64]*adapters.VerifierContractAddresses
	resolveErr   error
}

func (m *mockVerifierJobConfigAdapter) ResolveVerifierContractAddresses(
	_ datastore.DataStore,
	chainSelector uint64,
	_ string,
	_ string,
) (*adapters.VerifierContractAddresses, error) {
	if m.resolveErr != nil {
		return nil, m.resolveErr
	}
	cfg, ok := m.chainConfigs[chainSelector]
	if !ok {
		return nil, fmt.Errorf("no config for chain %d", chainSelector)
	}
	return cfg, nil
}

func newVerifierTopology(
	nopAliases []string,
	committeeQualifier string,
	selectors []uint64,
	mode shared.NOPMode,
) *offchain.EnvironmentTopology {
	nops := make([]offchain.NOPConfig, len(nopAliases))
	for i, alias := range nopAliases {
		nops[i] = offchain.NOPConfig{
			Alias:                 alias,
			Name:                  alias + "-name",
			Mode:                  mode,
			SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xabc123"},
		}
	}

	chainConfigs := make(map[string]offchain.ChainCommitteeConfig, len(selectors))
	for _, sel := range selectors {
		chainConfigs[fmt.Sprintf("%d", sel)] = offchain.ChainCommitteeConfig{
			NOPAliases: nopAliases,
			Threshold:  1,
		}
	}

	return &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: nops,
			Committees: map[string]offchain.CommitteeConfig{
				committeeQualifier: {
					Qualifier:    committeeQualifier,
					ChainConfigs: chainConfigs,
					Aggregators: []offchain.AggregatorConfig{
						{
							Name:    "agg-1",
							Address: "ws://agg:9090",
						},
					},
				},
			},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}
}

func newVerifierTestEnv(t *testing.T, selectors []uint64) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	chains := make(map[uint64]cldf_chain.BlockChain, len(selectors))
	for _, sel := range selectors {
		chains[sel] = cldfevm.Chain{Selector: sel}
	}
	return deployment.Environment{
		Name:        "test",
		BlockChains: cldf_chain.NewBlockChains(chains),
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		Logger:      lggr,
		OperationsBundle: operations.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			operations.NewMemoryReporter(),
		),
	}
}

func TestApplyVerifierConfig_Validation(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, &mockVerifierJobConfigAdapter{})

	tests := []struct {
		name    string
		env     deployment.Environment
		input   changesets.ApplyVerifierConfigInput
		wantErr string
	}{
		{
			name:    "missing topology returns error",
			env:     deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input:   changesets.ApplyVerifierConfigInput{},
			wantErr: "topology is required",
		},
		{
			name: "missing committee qualifier returns error",
			env:  deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology:                 newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, ""),
				DefaultExecutorQualifier: "pool1",
			},
			wantErr: "committee qualifier is required",
		},
		{
			name: "missing executor qualifier returns error",
			env:  deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology:           newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, ""),
				CommitteeQualifier: "c1",
			},
			wantErr: "default executor qualifier is required",
		},
		{
			name: "committee not found in topology returns error",
			env:  deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology:                 newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, ""),
				CommitteeQualifier:       "nonexistent",
				DefaultExecutorQualifier: "pool1",
			},
			wantErr: "committee \"nonexistent\" not found in topology",
		},
		{
			name: "target NOP not in topology returns error",
			env:  deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology:                 newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, ""),
				CommitteeQualifier:       "c1",
				DefaultExecutorQualifier: "pool1",
				TargetNOPs:               []shared.NOPAlias{"unknown-nop"},
			},
			wantErr: "NOP alias \"unknown-nop\" not found in topology",
		},
		{
			name: "chain selector not in environment returns error",
			env:  deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology:                 newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1, sel2}, ""),
				CommitteeQualifier:       "c1",
				DefaultExecutorQualifier: "pool1",
				ChainSelectors:           []uint64{sel2},
			},
			wantErr: "is not available in environment",
		},
		{
			name: "chain selector not in committee returns error",
			env:  newVerifierTestEnv(t, []uint64{sel1, sel2}),
			input: changesets.ApplyVerifierConfigInput{
				Topology:                 newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, ""),
				CommitteeQualifier:       "c1",
				DefaultExecutorQualifier: "pool1",
				ChainSelectors:           []uint64{sel2},
			},
			wantErr: "not configured in committee",
		},
		{
			name: "nil NOP topology returns error",
			env:  deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology: &offchain.EnvironmentTopology{
					IndexerAddress: []string{"http://indexer:8080"},
				},
				CommitteeQualifier:       "c1",
				DefaultExecutorQualifier: "pool1",
			},
			wantErr: "NOP topology with at least one NOP is required",
		},
		{
			name: "committee with no aggregators returns error",
			env:  deployment.Environment{BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology: &offchain.EnvironmentTopology{
					IndexerAddress: []string{"http://indexer:8080"},
					NOPTopology: &offchain.NOPTopology{
						NOPs: []offchain.NOPConfig{{Alias: "nop1", Name: "nop1-name"}},
						Committees: map[string]offchain.CommitteeConfig{
							"c1": {
								Qualifier:   "c1",
								Aggregators: []offchain.AggregatorConfig{},
								ChainConfigs: map[string]offchain.ChainCommitteeConfig{
									fmt.Sprintf("%d", sel1): {NOPAliases: []string{"nop1"}, Threshold: 1},
								},
							},
						},
					},
				},
				CommitteeQualifier:       "c1",
				DefaultExecutorQualifier: "pool1",
			},
			wantErr: "at least one aggregator is required",
		},
		{
			name: "pyroscope URL in production returns error",
			env:  deployment.Environment{Name: "mainnet", BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{sel1: cldfevm.Chain{Selector: sel1}})},
			input: changesets.ApplyVerifierConfigInput{
				Topology: func() *offchain.EnvironmentTopology {
					topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, "")
					topo.PyroscopeURL = "http://pyroscope"
					return topo
				}(),
				CommitteeQualifier:       "c1",
				DefaultExecutorQualifier: "pool1",
			},
			wantErr: "pyroscope URL is not supported for production environments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := changesets.ApplyVerifierConfig(registry)
			err := cs.VerifyPreconditions(tt.env, tt.input)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestApplyVerifierConfig_HappyPathBuildsJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyVerifierConfig_AdapterErrorPropagates(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		resolveErr: assert.AnError,
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to resolve contract addresses")
}

func TestApplyVerifierConfig_TargetNOPsFiltersJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1", "nop2", "nop3"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		TargetNOPs:               []shared.NOPAlias{"nop1"},
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyVerifierConfig_MissingSignerAddressReturnsError(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	topo.NOPTopology.NOPs[0].SignerAddressByFamily = nil

	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing signer address")
}

func TestApplyVerifierConfig_ChainSelectorsFilteredByCommitteeChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}
