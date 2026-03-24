package changesets_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

var _ adapters.DeployChainContractsAdapter = (*mockDeployAdapter)(nil)

type mockDeployAdapter struct {
	outputByChain map[uint64]sequences.OnChainOutput
	errByChain    map[uint64]error
}

func (m *mockDeployAdapter) IsSupportedChain(_ deployment.Environment, _ uint64) bool {
	return false
}

func (m *mockDeployAdapter) DeriveLaneVersionsForChain(_ deployment.Environment, _ uint64) (map[uint64]*semver.Version, []*semver.Version, error) {
	return map[uint64]*semver.Version{}, []*semver.Version{}, nil
}

func (m *mockDeployAdapter) InitializeAdapter(_ deployment.Environment, _ uint64) error {
	return nil
}

func (m *mockDeployAdapter) ConnectedChains(_ deployment.Environment, _ uint64) ([]uint64, error) {
	return []uint64{}, nil
}

func (m *mockDeployAdapter) SupportedTokensPerRemoteChain(_ deployment.Environment, _ uint64) (map[uint64][]common.Address, error) {
	return map[uint64][]common.Address{}, nil
}

func (m *mockDeployAdapter) SequenceImportConfig() *cldf_ops.Sequence[deploy.ImportConfigPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-import-config",
		semver.MustParse("2.0.0"),
		"Mock sequence for importing config",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input deploy.ImportConfigPerChainInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, nil
		})
}

func (m *mockDeployAdapter) SetContractParamsFromImportedConfig() *cldf_ops.Sequence[adapters.DeployChainConfigCreatorInput, adapters.DeployContractParams, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-set-contract-params",
		semver.MustParse("2.0.0"),
		"Mock sequence for setting contract params from imported config",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.DeployChainConfigCreatorInput) (adapters.DeployContractParams, error) {
			return adapters.DeployContractParams{
				OnRamp: adapters.OnRampDeployParams{
					Version:       semver.MustParse("2.0.0"),
					FeeAggregator: "0xDummyOnRampFeeAgg",
				},
			}, nil
		})
}

func (m *mockDeployAdapter) DeployChainContracts() *cldf_ops.Sequence[adapters.DeployChainContractsInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-deploy-chain-contracts",
		semver.MustParse("2.0.0"),
		"Mock deploy sequence for testing",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.DeployChainContractsInput) (sequences.OnChainOutput, error) {
			if m.errByChain != nil {
				if err, ok := m.errByChain[input.ChainSelector]; ok {
					return sequences.OnChainOutput{}, err
				}
			}
			if m.outputByChain != nil {
				if out, ok := m.outputByChain[input.ChainSelector]; ok {
					return out, nil
				}
			}
			return sequences.OnChainOutput{}, nil
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

func newDefaultPerChainCfg() changesets.DeployChainContractsPerChainCfg {
	return changesets.DeployChainContractsPerChainCfg{
		DeployerContract: "0x0000000000000000000000000000000000001234",
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
				ChainSelectors:                          []uint64{sel1},
				DefaultCfg:                              newDefaultPerChainCfg(),
				IgnoreImportedConfigFromPreviousVersion: true,
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
				ChainSelectors:                          []uint64{sel1},
				DefaultCfg:                              newDefaultPerChainCfg(),
				IgnoreImportedConfigFromPreviousVersion: true,
			},
			expectedErr: "no committees defined in topology",
		},
		{
			name: "rejects empty chain selectors",
			cfg: changesets.DeployChainContractsCfg{
				Topology:                                validTopology,
				ChainSelectors:                          []uint64{},
				DefaultCfg:                              newDefaultPerChainCfg(),
				IgnoreImportedConfigFromPreviousVersion: true,
			},
			expectedErr: "at least one chain selector is required",
		},
		{
			name: "rejects unknown chain selector",
			cfg: changesets.DeployChainContractsCfg{
				Topology:                                validTopology,
				ChainSelectors:                          []uint64{99999},
				DefaultCfg:                              newDefaultPerChainCfg(),
				IgnoreImportedConfigFromPreviousVersion: true,
			},
			expectedErr: "chain selector 99999 is not available in environment",
		},
		{
			name: "rejects empty DeployerContract",
			cfg: changesets.DeployChainContractsCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				DefaultCfg: changesets.DeployChainContractsPerChainCfg{
					DeployerContract: "",
				},
				IgnoreImportedConfigFromPreviousVersion: true,
			},
			expectedErr: "DeployerContract is required",
		},
		{
			name: "rejects empty DeployerContract in chain override",
			cfg: changesets.DeployChainContractsCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				DefaultCfg:     newDefaultPerChainCfg(),
				ChainCfgs: map[uint64]changesets.DeployChainContractsPerChainCfg{
					sel1: {DeployerContract: ""},
				},
				IgnoreImportedConfigFromPreviousVersion: true,
			},
			expectedErr: "DeployerContract is required",
		},
		{
			name: "rejects duplicate chain selectors",
			cfg: changesets.DeployChainContractsCfg{
				Topology:                                validTopology,
				ChainSelectors:                          []uint64{sel1, sel1},
				DefaultCfg:                              newDefaultPerChainCfg(),
				IgnoreImportedConfigFromPreviousVersion: true,
			},
			expectedErr: "duplicate chain selector",
		},
		{
			name: "rejects ChainCfgs key not in ChainSelectors",
			cfg: changesets.DeployChainContractsCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				DefaultCfg:     newDefaultPerChainCfg(),
				ChainCfgs: map[uint64]changesets.DeployChainContractsPerChainCfg{
					77777: newDefaultPerChainCfg(),
				},
				IgnoreImportedConfigFromPreviousVersion: true,
			},
			expectedErr: "ChainCfgs contains selector 77777 which is not in ChainSelectors",
		},
		{
			name: "accepts valid config",
			cfg: changesets.DeployChainContractsCfg{
				Topology:                                validTopology,
				ChainSelectors:                          []uint64{sel1},
				DefaultCfg:                              newDefaultPerChainCfg(),
				IgnoreImportedConfigFromPreviousVersion: true,
			},
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

func TestDeployChainContracts_With_DummyConfigImport(t *testing.T) {
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
	registry.RegisterLaneVersionResolver(chainsel.FamilyEVM, mock)

	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsCfg{
			Topology:       newDeployTestTopology(sel1),
			ChainSelectors: []uint64{sel1},
			DefaultCfg:     newDefaultPerChainCfg(),
		},
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
		Cfg: changesets.DeployChainContractsCfg{
			Topology:                                newDeployTestTopology(sel1),
			ChainSelectors:                          []uint64{sel1},
			DefaultCfg:                              newDefaultPerChainCfg(),
			IgnoreImportedConfigFromPreviousVersion: true,
		},
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
		Cfg: changesets.DeployChainContractsCfg{
			Topology:                                newDeployTestTopology(sel1, sel2),
			ChainSelectors:                          []uint64{sel1, sel2},
			DefaultCfg:                              newDefaultPerChainCfg(),
			IgnoreImportedConfigFromPreviousVersion: true,
		},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	assert.Len(t, addrs, 2)
}

func TestDeployChainContracts_Apply_PerChainOverrideIsUsed(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	env := newDeployTestEnv(t, []uint64{sel1, sel2})

	var capturedInputs []adapters.DeployChainContractsInput

	captureAdapter := &capturingDeployAdapter{
		captured: &capturedInputs,
	}

	registry := adapters.NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, captureAdapter)

	overrideCfg := newDefaultPerChainCfg()
	overrideCfg.DeployerContract = "0x0000000000000000000000000000000000005678"

	cs := changesets.DeployChainContracts(registry)
	_, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsCfg{
			Topology:       newDeployTestTopology(sel1, sel2),
			ChainSelectors: []uint64{sel1, sel2},
			DefaultCfg:     newDefaultPerChainCfg(),
			ChainCfgs: map[uint64]changesets.DeployChainContractsPerChainCfg{
				sel2: overrideCfg,
			},
			IgnoreImportedConfigFromPreviousVersion: true,
		},
	})
	require.NoError(t, err)

	require.Len(t, capturedInputs, 2)
	assert.Equal(t, "0x0000000000000000000000000000000000001234", capturedInputs[0].DeployerContract)
	assert.Equal(t, "0x0000000000000000000000000000000000005678", capturedInputs[1].DeployerContract)
}

type capturingDeployAdapter struct {
	captured *[]adapters.DeployChainContractsInput
}

func (c *capturingDeployAdapter) SetContractParamsFromImportedConfig() *cldf_ops.Sequence[adapters.DeployChainConfigCreatorInput, adapters.DeployContractParams, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"capturing-set-contract-params",
		semver.MustParse("2.0.0"),
		"Captures inputs for testing",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.DeployChainConfigCreatorInput) (adapters.DeployContractParams, error) {
			return adapters.DeployContractParams{}, nil
		})
}

var _ adapters.DeployChainContractsAdapter = (*capturingDeployAdapter)(nil)

func (c *capturingDeployAdapter) DeployChainContracts() *cldf_ops.Sequence[adapters.DeployChainContractsInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"capturing-deploy",
		semver.MustParse("2.0.0"),
		"Captures inputs for testing",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.DeployChainContractsInput) (sequences.OnChainOutput, error) {
			*c.captured = append(*c.captured, input)
			return sequences.OnChainOutput{}, nil
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
		Cfg: changesets.DeployChainContractsCfg{
			Topology:                                newDeployTestTopology(sel1),
			ChainSelectors:                          []uint64{sel1},
			DefaultCfg:                              newDefaultPerChainCfg(),
			IgnoreImportedConfigFromPreviousVersion: true,
		},
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
			Topology:                                newDeployTestTopology(sel2),
			ChainSelectors:                          []uint64{sel1},
			DefaultCfg:                              newDefaultPerChainCfg(),
			IgnoreImportedConfigFromPreviousVersion: true,
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
		Cfg: changesets.DeployChainContractsCfg{
			Topology:                                newDeployTestTopology(sel1),
			ChainSelectors:                          []uint64{sel1},
			DefaultCfg:                              newDefaultPerChainCfg(),
			IgnoreImportedConfigFromPreviousVersion: true,
		},
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
