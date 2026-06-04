package changesets_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

var _ adapters.DeployChainContractsAdapter = (*mockDeployAdapter)(nil)

type mockDeployAdapter struct {
	outputByChain map[uint64]sequences.OnChainOutput
	errByChain    map[uint64]error
}

func (m *mockDeployAdapter) GetDefaultDeployContractParams(_ uint64) adapters.DeployContractParams {
	return adapters.DeployContractParams{}
}

func (m *mockDeployAdapter) ResolveDeployAddresses(_ deployment.Environment, _ uint64) (adapters.DeployChainResolvedAddresses, error) {
	return adapters.DeployChainResolvedAddresses{DeployerContract: "0x0000000000000000000000000000000000001234"}, nil
}

func (m *mockDeployAdapter) BuildDeployContractParams(in adapters.BuildDeployContractParamsInput) (adapters.DeployContractParams, error) {
	return adapters.DeployContractParams{CommitteeVerifiers: in.CommitteeVerifiers}, nil
}

func (m *mockDeployAdapter) DeployChainContracts() *cldf_ops.Sequence[adapters.DeployChainContractsInput, adapters.DeployChainContractsOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-deploy-chain-contracts",
		semver.MustParse("2.0.0"),
		"Mock deploy sequence for testing",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.DeployChainContractsInput) (adapters.DeployChainContractsOutput, error) {
			if m.errByChain != nil {
				if err, ok := m.errByChain[input.ChainSelector]; ok {
					return adapters.DeployChainContractsOutput{}, err
				}
			}
			if m.outputByChain != nil {
				if out, ok := m.outputByChain[input.ChainSelector]; ok {
					return adapters.DeployChainContractsOutput{
						OnChainOutput: out,
					}, nil
				}
			}
			return adapters.DeployChainContractsOutput{}, nil
		},
	)
}

func newDeployTestEnv(t *testing.T, selectors []uint64) deployment.Environment {
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
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

func newDeployTestTopology(chainSelectors ...uint64) *offchain.EnvironmentTopology {
	committees := map[string]offchain.CommitteeConfig{
		"default": {
			Qualifier:        "default",
			VerifierVersion:  semver.MustParse("2.0.0"),
			StorageLocations: []string{"https://store.test"},
			ChainConfigs:     make(map[string]offchain.ChainCommitteeConfig),
			Aggregators: []offchain.AggregatorConfig{
				{Name: "agg-1", Address: "localhost:8080"},
			},
		},
	}
	for _, sel := range chainSelectors {
		committees["default"].ChainConfigs[strconv.FormatUint(sel, 10)] = offchain.ChainCommitteeConfig{
			NOPAliases:    []string{"nop-1"},
			Threshold:     1,
			FeeAggregator: "0x0000000000000000000000000000000000000001",
		}
	}
	return &offchain.EnvironmentTopology{
		IndexerAddress: []string{"localhost:9090"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop-1", Name: "NOP One"},
			},
			Committees: committees,
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}
}

func newDeployChainContractsCfg(selectors ...uint64) changesets.DeployChainContractsCfg {
	chainOverrides := make(map[uint64]changesets.DeployChainContractsPerChainCfg, len(selectors))
	for _, sel := range selectors {
		chainOverrides[sel] = changesets.DeployChainContractsPerChainCfg{
			DeployerKeyOwned: true,
		}
	}
	return changesets.DeployChainContractsCfg{
		Topology:       newDeployTestTopology(selectors...),
		ChainSelectors: selectors,
		ChainOverrides: chainOverrides,
	}
}

func TestDeployChainContracts_Validate(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	validTopology := newDeployTestTopology(sel1)

	tests := []struct {
		name        string
		cfg         changesets.DeployChainContractsCfg
		expectedErr string
	}{
		{
			name: "rejects nil topology",
			cfg: changesets.DeployChainContractsCfg{
				ChainSelectors: []uint64{sel1},
			},
			expectedErr: "topology is required",
		},
		{
			name: "rejects empty committees",
			cfg: changesets.DeployChainContractsCfg{
				Topology: &offchain.EnvironmentTopology{
					IndexerAddress: []string{"localhost:9090"},
					NOPTopology: &offchain.NOPTopology{
						NOPs:       []offchain.NOPConfig{{Alias: "a", Name: "A"}},
						Committees: map[string]offchain.CommitteeConfig{},
					},
					ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
				},
				ChainSelectors: []uint64{sel1},
			},
			expectedErr: "no committees defined in topology",
		},
		{
			name: "rejects empty chain selectors",
			cfg: changesets.DeployChainContractsCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{},
			},
			expectedErr: "at least one chain selector is required",
		},
		{
			name: "rejects unknown chain selector",
			cfg: changesets.DeployChainContractsCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{99999},
			},
			expectedErr: "chain selector 99999 is not available in environment",
		},
		{
			name: "rejects duplicate chain selectors",
			cfg: changesets.DeployChainContractsCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1, sel1},
			},
			expectedErr: "duplicate chain selector",
		},
		{
			name: "rejects chainOverrides key not in chainSelectors",
			cfg: changesets.DeployChainContractsCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				ChainOverrides: map[uint64]changesets.DeployChainContractsPerChainCfg{
					77777: {},
				},
			},
			expectedErr: "chainOverrides contains selector 77777 which is not in chainSelectors",
		},
		{
			name: "accepts valid config",
			cfg:  newDeployChainContractsCfg(sel1),
		},
	}

	env := newDeployTestEnv(t, []uint64{sel1})
	registry := adapters.NewDeployChainContractsRegistry()
	cs := changesets.DeployChainContracts(registry)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
				MCMS: mcms.Input{},
				Cfg:  tc.cfg,
			})
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDeployChainContracts_Apply_MinimalConfigWithoutOverrides(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	env := newDeployTestEnv(t, []uint64{sel1})
	mock := &mockDeployAdapter{
		outputByChain: map[uint64]sequences.OnChainOutput{
			sel1: {
				Addresses: []datastore.AddressRef{
					{ChainSelector: sel1, Type: "TestContract", Version: semver.MustParse("2.0.0"), Address: "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"},
				},
			},
		},
	}
	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg:  newDeployChainContractsCfg(sel1),
	})
	require.NoError(t, err)
}

func TestDeployChainContracts_Apply_SingleChainSuccess(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	env := newDeployTestEnv(t, []uint64{sel1})

	expectedAddr := datastore.AddressRef{
		ChainSelector: sel1,
		Type:          "TestContract",
		Version:       semver.MustParse("2.0.0"),
		Address:       "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
	}

	mock := &mockDeployAdapter{
		outputByChain: map[uint64]sequences.OnChainOutput{
			sel1: {
				Addresses: []datastore.AddressRef{expectedAddr},
			},
		},
	}

	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	cs := changesets.DeployChainContracts(registry)
	out, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg:  newDeployChainContractsCfg(sel1),
	})
	require.NoError(t, err)
	require.NotNil(t, out.DataStore)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	assert.Len(t, addrs, 1)
	assert.Equal(t, expectedAddr.Address, addrs[0].Address)
	assert.Equal(t, expectedAddr.Type, addrs[0].Type)
}

func TestDeployChainContracts_Apply_MultiChainSuccess(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	env := newDeployTestEnv(t, []uint64{sel1, sel2})

	mock := &mockDeployAdapter{
		outputByChain: map[uint64]sequences.OnChainOutput{
			sel1: {
				Addresses: []datastore.AddressRef{
					{ChainSelector: sel1, Type: "Contract1", Version: semver.MustParse("2.0.0"), Address: "0xAAAA"},
				},
			},
			sel2: {
				Addresses: []datastore.AddressRef{
					{ChainSelector: sel2, Type: "Contract2", Version: semver.MustParse("2.0.0"), Address: "0xBBBB"},
				},
			},
		},
	}

	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	cs := changesets.DeployChainContracts(registry)
	out, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg:  newDeployChainContractsCfg(sel1, sel2),
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	assert.Len(t, addrs, 2)
}

func TestDeployChainContracts_Apply_PerChainContractParamsOverrideIsUsed(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	env := newDeployTestEnv(t, []uint64{sel1, sel2})

	maxUSDCents := uint32(50_00)
	captureAdapter := &capturingDeployAdapter{
		buildInputsByChain: make(map[uint64]adapters.BuildDeployContractParamsInput),
	}

	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, captureAdapter)

	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsCfg{
			Topology:       newDeployTestTopology(sel1, sel2),
			ChainSelectors: []uint64{sel1, sel2},
			ChainOverrides: map[uint64]changesets.DeployChainContractsPerChainCfg{
				sel1: {DeployerKeyOwned: true},
				sel2: {
					DeployerKeyOwned: true,
					ContractParams: &adapters.DeployContractParamsOverrides{
						OnRamp: &adapters.OnRampDeployParamsOverrides{
							MaxUSDCentsPerMessage: &maxUSDCents,
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)

	_, ok := captureAdapter.buildInputsByChain[sel1]
	require.True(t, ok, "chain %d should be deployed", sel1)
	sel2Input, ok := captureAdapter.buildInputsByChain[sel2]
	require.True(t, ok, "chain %d should be deployed", sel2)
	require.NotNil(t, sel2Input.Overrides)
	require.NotNil(t, sel2Input.Overrides.OnRamp)
	require.NotNil(t, sel2Input.Overrides.OnRamp.MaxUSDCentsPerMessage)
	assert.Equal(t, maxUSDCents, *sel2Input.Overrides.OnRamp.MaxUSDCentsPerMessage)
}

func TestDeployChainContracts_Apply_PerChainDeployFlagsAreUsed(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	env := newDeployTestEnv(t, []uint64{sel1, sel2})

	captureAdapter := &capturingDeployAdapter{
		buildInputsByChain:  make(map[uint64]adapters.BuildDeployContractParamsInput),
		deployInputsByChain: make(map[uint64]adapters.DeployChainContractsInput),
	}

	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, captureAdapter)

	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsCfg{
			Topology:       newDeployTestTopology(sel1, sel2),
			ChainSelectors: []uint64{sel1, sel2},
			ChainOverrides: map[uint64]changesets.DeployChainContractsPerChainCfg{
				sel1: {DeployerKeyOwned: true},
				sel2: {
					DeployTestRouter: true,
					DeployerKeyOwned: true,
				},
			},
		},
	})
	require.NoError(t, err)

	sel1Input, ok := captureAdapter.deployInputsByChain[sel1]
	require.True(t, ok, "chain %d should be deployed", sel1)
	assert.False(t, sel1Input.DeployTestRouter)
	assert.True(t, sel1Input.DeployerKeyOwned)

	sel2Input, ok := captureAdapter.deployInputsByChain[sel2]
	require.True(t, ok, "chain %d should be deployed", sel2)
	assert.True(t, sel2Input.DeployTestRouter)
	assert.True(t, sel2Input.DeployerKeyOwned)
}

type capturingDeployAdapter struct {
	buildInputsByChain  map[uint64]adapters.BuildDeployContractParamsInput
	deployInputsByChain map[uint64]adapters.DeployChainContractsInput
}

func (c *capturingDeployAdapter) GetDefaultDeployContractParams(_ uint64) adapters.DeployContractParams {
	return adapters.DeployContractParams{}
}

func (c *capturingDeployAdapter) ResolveDeployAddresses(_ deployment.Environment, _ uint64) (adapters.DeployChainResolvedAddresses, error) {
	return adapters.DeployChainResolvedAddresses{DeployerContract: "0x0000000000000000000000000000000000001234"}, nil
}

func (c *capturingDeployAdapter) BuildDeployContractParams(in adapters.BuildDeployContractParamsInput) (adapters.DeployContractParams, error) {
	c.buildInputsByChain[in.ChainSelector] = in
	return adapters.DeployContractParams{CommitteeVerifiers: in.CommitteeVerifiers}, nil
}

var _ adapters.DeployChainContractsAdapter = (*capturingDeployAdapter)(nil)

func (c *capturingDeployAdapter) DeployChainContracts() *cldf_ops.Sequence[adapters.DeployChainContractsInput, adapters.DeployChainContractsOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"capturing-deploy",
		semver.MustParse("2.0.0"),
		"Captures inputs for testing",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.DeployChainContractsInput) (adapters.DeployChainContractsOutput, error) {
			if c.deployInputsByChain != nil {
				c.deployInputsByChain[input.ChainSelector] = input
			}
			return adapters.DeployChainContractsOutput{}, nil
		},
	)
}

func TestDeployChainContracts_Apply_AdapterErrorPropagated(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	env := newDeployTestEnv(t, []uint64{sel1})

	mock := &mockDeployAdapter{
		errByChain: map[uint64]error{
			sel1: fmt.Errorf("deploy failed: out of gas"),
		},
	}

	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg:  newDeployChainContractsCfg(sel1),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "deploy failed: out of gas")
}

func TestDeployChainContracts_Apply_ReturnsError_WhenNoCommitteesHaveChainConfig(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	env := newDeployTestEnv(t, []uint64{sel1, sel2})

	mock := &mockDeployAdapter{}
	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsCfg{
			Topology:       newDeployTestTopology(sel2),
			ChainSelectors: []uint64{sel1},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no committees have chain_config for this selector")
}

func TestDeployChainContracts_Apply_NoAdapterRegistered(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	env := newDeployTestEnv(t, []uint64{sel1})

	registry := adapters.NewDeployChainContractsRegistry()
	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg:  newDeployChainContractsCfg(sel1),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no deploy chain contracts adapter registered")
}

func TestBuildCommitteeVerifierParams_MapsAllCommittees(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)
	topology := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"localhost:9090"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{{Alias: "nop-1", Name: "NOP One"}},
			Committees: map[string]offchain.CommitteeConfig{
				"alpha": {
					Qualifier:        "alpha",
					VerifierVersion:  semver.MustParse("2.0.0"),
					StorageLocations: []string{"https://store1.test"},
					ChainConfigs: map[string]offchain.ChainCommitteeConfig{
						selStr: {
							NOPAliases:     []string{"nop-1"},
							Threshold:      1,
							FeeAggregator:  "0x0000000000000000000000000000000000000001",
							AllowlistAdmin: "0x0000000000000000000000000000000000000002",
						},
					},
					Aggregators: []offchain.AggregatorConfig{{Name: "a", Address: "localhost:1"}},
				},
				"beta": {
					Qualifier:        "beta",
					VerifierVersion:  semver.MustParse("2.0.0"),
					StorageLocations: []string{"https://store2.test"},
					ChainConfigs: map[string]offchain.ChainCommitteeConfig{
						selStr: {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "0x0000000000000000000000000000000000000003",
						},
					},
					Aggregators: []offchain.AggregatorConfig{{Name: "b", Address: "localhost:2"}},
				},
			},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}

	params, err := changesets.BuildCommitteeVerifierParams(topology, sel)
	require.NoError(t, err)
	assert.Len(t, params, 2)

	qualifiers := map[string]adapters.CommitteeVerifierDeployParams{}
	for _, p := range params {
		qualifiers[p.Qualifier] = p
	}

	alpha, ok := qualifiers["alpha"]
	require.True(t, ok)
	assert.Equal(t, "0x0000000000000000000000000000000000000001", alpha.FeeAggregator)
	assert.Equal(t, "0x0000000000000000000000000000000000000002", alpha.AllowlistAdmin)
	assert.Equal(t, []string{"https://store1.test"}, alpha.StorageLocations)

	beta, ok := qualifiers["beta"]
	require.True(t, ok)
	assert.Equal(t, "0x0000000000000000000000000000000000000003", beta.FeeAggregator)
	assert.Empty(t, beta.AllowlistAdmin)
	assert.Equal(t, []string{"https://store2.test"}, beta.StorageLocations)
}

func TestBuildCommitteeVerifierParams_SkipsCommitteesWithoutChainConfig(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)
	otherSel := chainsel.TEST_90000002.Selector
	topology := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"localhost:9090"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{{Alias: "nop-1", Name: "NOP One"}},
			Committees: map[string]offchain.CommitteeConfig{
				"present": {
					Qualifier:       "present",
					VerifierVersion: semver.MustParse("2.0.0"),
					ChainConfigs: map[string]offchain.ChainCommitteeConfig{
						selStr: {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "0x0000000000000000000000000000000000000001",
						},
					},
					Aggregators: []offchain.AggregatorConfig{{Name: "a", Address: "localhost:1"}},
				},
				"absent": {
					Qualifier:       "absent",
					VerifierVersion: semver.MustParse("2.0.0"),
					ChainConfigs: map[string]offchain.ChainCommitteeConfig{
						strconv.FormatUint(otherSel, 10): {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "0x0000000000000000000000000000000000000002",
						},
					},
					Aggregators: []offchain.AggregatorConfig{{Name: "b", Address: "localhost:2"}},
				},
			},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}

	params, err := changesets.BuildCommitteeVerifierParams(topology, sel)
	require.NoError(t, err)
	assert.Len(t, params, 1)
	assert.Equal(t, "present", params[0].Qualifier)
}

func TestBuildCommitteeVerifierParams_RejectsNilNOPTopology(t *testing.T) {
	topology := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"localhost:9090"},
		ExecutorPools:  map[string]offchain.ExecutorPoolConfig{},
	}
	_, err := changesets.BuildCommitteeVerifierParams(topology, 1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "NOPTopology is nil")
}

func TestBuildCommitteeVerifierParams_RejectsNilVerifierVersion(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)
	topology := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"localhost:9090"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{{Alias: "nop-1", Name: "NOP One"}},
			Committees: map[string]offchain.CommitteeConfig{
				"broken": {
					Qualifier:       "broken",
					VerifierVersion: nil,
					ChainConfigs: map[string]offchain.ChainCommitteeConfig{
						selStr: {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "0x0000000000000000000000000000000000000001",
						},
					},
					Aggregators: []offchain.AggregatorConfig{{Name: "a", Address: "localhost:1"}},
				},
			},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}
	_, err := changesets.BuildCommitteeVerifierParams(topology, sel)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nil VerifierVersion")
}

func TestBuildCommitteeVerifierParams_ReturnsEmptyForEmptyCommittees(t *testing.T) {
	topology := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"localhost:9090"},
		NOPTopology: &offchain.NOPTopology{
			NOPs:       []offchain.NOPConfig{{Alias: "nop-1", Name: "NOP One"}},
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}
	params, err := changesets.BuildCommitteeVerifierParams(topology, 1)
	require.NoError(t, err)
	assert.Empty(t, params)
}

func TestBuildCommitteeVerifierParams_RejectsEmptyFeeAggregator(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)
	topology := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"localhost:9090"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{{Alias: "nop-1", Name: "NOP One"}},
			Committees: map[string]offchain.CommitteeConfig{
				"test": {
					Qualifier:       "test",
					VerifierVersion: semver.MustParse("2.0.0"),
					ChainConfigs: map[string]offchain.ChainCommitteeConfig{
						selStr: {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "",
						},
					},
					Aggregators: []offchain.AggregatorConfig{{Name: "a", Address: "localhost:1"}},
				},
			},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}
	_, err := changesets.BuildCommitteeVerifierParams(topology, sel)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "FeeAggregator is required")
}
