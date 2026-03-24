package changesets_test

import (
	"context"
	"fmt"
	"math/big"
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
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

// lanesTest_MockReader was previously in configure_chains_for_lanes_test.go
type lanesTest_MockReader struct{}

func (m *lanesTest_MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, _ mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{
		MCMAddress:      "0xMCM",
		StartingOpCount: 42,
	}, nil
}

func (m *lanesTest_MockReader) GetTimelockRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0xTimelock",
		Type:          "Timelock",
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

func (m *lanesTest_MockReader) GetMCMSRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0xMCM",
		Type:          "MCM",
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

var _ adapters.CommitteeVerifierContractAdapter = (*mockCommitteeVerifierContractAdapter)(nil)

type mockCommitteeVerifierContractAdapter struct {
	contractsByChainAndQualifier map[string][]datastore.AddressRef
	resolveErr                   error
}

func (m *mockCommitteeVerifierContractAdapter) ResolveCommitteeVerifierContracts(
	_ datastore.DataStore,
	chainSelector uint64,
	qualifier string,
) ([]datastore.AddressRef, error) {
	if m.resolveErr != nil {
		return nil, m.resolveErr
	}
	key := fmt.Sprintf("%d:%s", chainSelector, qualifier)
	contracts, ok := m.contractsByChainAndQualifier[key]
	if !ok {
		return nil, fmt.Errorf("no contracts for chain %d qualifier %q", chainSelector, qualifier)
	}
	return contracts, nil
}

var _ lanes.LaneAdapter = (*mockLaneAdapter)(nil)

type mockLaneAdapter struct {
	sequenceErr     error
	routerCalls     int
	testRouterCalls int
	sourceInputs    []lanes.UpdateLanesInput
	destInputs      []lanes.UpdateLanesInput
}

func (m *mockLaneAdapter) ConfigureLaneLegAsSource() *cldf_ops.Sequence[lanes.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-lane-leg-as-source",
		semver.MustParse("2.0.0"),
		"Mock source sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
			m.sourceInputs = append(m.sourceInputs, input)
			if m.sequenceErr != nil {
				return sequences.OnChainOutput{}, m.sequenceErr
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{},
			}, nil
		},
	)
}

func (m *mockLaneAdapter) ConfigureLaneLegAsDest() *cldf_ops.Sequence[lanes.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-lane-leg-as-dest",
		semver.MustParse("2.0.0"),
		"Mock dest sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
			m.destInputs = append(m.destInputs, input)
			if m.sequenceErr != nil {
				return sequences.OnChainOutput{}, m.sequenceErr
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{},
			}, nil
		},
	)
}

func (m *mockLaneAdapter) GetOnRampAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return []byte("0xOnRamp"), nil
}

func (m *mockLaneAdapter) GetOffRampAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return []byte("0xOffRamp"), nil
}

func (m *mockLaneAdapter) GetRouterAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	m.routerCalls++
	return []byte("0xRouter"), nil
}

func (m *mockLaneAdapter) GetTestRouter(_ datastore.DataStore, _ uint64) ([]byte, error) {
	m.testRouterCalls++
	return []byte("0xTestRouter"), nil
}

func (m *mockLaneAdapter) GetFQAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return []byte("0xFeeQuoter"), nil
}

func (m *mockLaneAdapter) GetFeeQuoterDestChainConfig() lanes.FeeQuoterDestChainConfig {
	return lanes.DefaultFeeQuoterDestChainConfig(false, chainsel.ETHEREUM_MAINNET.Selector)
}

func (m *mockLaneAdapter) GetDefaultGasPrice() *big.Int {
	return big.NewInt(2e12)
}

func (m *mockLaneAdapter) DisableRemoteChain() *cldf_ops.Sequence[lanes.DisableRemoteChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func newCommitteeVerifierTestEnv(t *testing.T, selectors []uint64) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	chains := make(map[uint64]cldf_chain.BlockChain, len(selectors))
	for _, sel := range selectors {
		chains[sel] = cldfevm.Chain{Selector: sel}
	}

	ds := datastore.NewMemoryDataStore()
	for _, sel := range selectors {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel,
			Address:       "0xRouter" + strconv.FormatUint(sel, 10),
			Type:          "Router",
			Version:       semver.MustParse("1.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel,
			Address:       "0xOnRamp" + strconv.FormatUint(sel, 10),
			Type:          "OnRamp",
			Version:       semver.MustParse("1.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel,
			Address:       "0xFeeQuoter" + strconv.FormatUint(sel, 10),
			Type:          "FeeQuoter",
			Version:       semver.MustParse("1.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel,
			Address:       "0xOffRamp" + strconv.FormatUint(sel, 10),
			Type:          "OffRamp",
			Version:       semver.MustParse("1.0.0"),
		}))
	}

	return deployment.Environment{
		Name:        "test",
		BlockChains: cldf_chain.NewBlockChains(chains),
		DataStore:   ds.Seal(),
		Logger:      lggr,
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

func newCommitteeVerifierTestEnvWithContracts(t *testing.T, selectors []uint64, verifierAddr, resolverAddr string) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	chains := make(map[uint64]cldf_chain.BlockChain, len(selectors))
	for _, sel := range selectors {
		chains[sel] = cldfevm.Chain{Selector: sel}
	}

	ds := datastore.NewMemoryDataStore()
	for _, sel := range selectors {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: "0xRouter" + strconv.FormatUint(sel, 10),
			Type: "Router", Version: semver.MustParse("1.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: "0xOnRamp" + strconv.FormatUint(sel, 10),
			Type: "OnRamp", Version: semver.MustParse("1.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: "0xFeeQuoter" + strconv.FormatUint(sel, 10),
			Type: "FeeQuoter", Version: semver.MustParse("1.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: "0xOffRamp" + strconv.FormatUint(sel, 10),
			Type: "OffRamp", Version: semver.MustParse("1.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: verifierAddr, Qualifier: "default",
			Type: "CommitteeVerifier", Version: semver.MustParse("2.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: resolverAddr, Qualifier: "default",
			Type: "CommitteeVerifierResolver", Version: semver.MustParse("2.0.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: "0xExec",
			Type: "Executor", Version: semver.MustParse("1.0.0"),
		}))
	}

	return deployment.Environment{
		Name:        "test",
		BlockChains: cldf_chain.NewBlockChains(chains),
		DataStore:   ds.Seal(),
		Logger:      lggr,
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

func getOrRegisterMockLaneAdapter(t *testing.T) *mockLaneAdapter {
	t.Helper()

	laneRegistry := lanes.GetLaneAdapterRegistry()
	mockAdapter := &mockLaneAdapter{}
	if registeredAdapter, ok := laneRegistry.GetLaneAdapter(chainsel.FamilyEVM, semver.MustParse("2.0.0")); ok {
		var typeOK bool
		mockAdapter, typeOK = registeredAdapter.(*mockLaneAdapter)
		require.True(t, typeOK, "expected mock lane adapter in test registry")
		mockAdapter.routerCalls = 0
		mockAdapter.testRouterCalls = 0
		mockAdapter.sequenceErr = nil
		mockAdapter.sourceInputs = nil
		mockAdapter.destInputs = nil
	} else {
		laneRegistry.RegisterLaneAdapter(chainsel.FamilyEVM, semver.MustParse("2.0.0"), mockAdapter)
	}

	return mockAdapter
}

func fqOverride(isEnabled bool, maxDataBytes uint32) *lanes.FeeQuoterDestChainConfigOverride {
	o := lanes.FeeQuoterDestChainConfigOverride(func(cfg *lanes.FeeQuoterDestChainConfig) {
		cfg.IsEnabled = isEnabled
		cfg.MaxDataBytes = maxDataBytes
	})
	return &o
}

func TestConfigureChainsForLanesFromTopology_Validation(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	laneRegistry := lanes.GetLaneAdapterRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, laneRegistry, nil)

	tests := []struct {
		name        string
		input       changesets.ConfigureChainsForLanesFromTopologyConfig
		errContains string
	}{
		{
			name:        "nil topology returns error",
			input:       changesets.ConfigureChainsForLanesFromTopologyConfig{},
			errContains: "topology is required",
		},
		{
			name: "empty committees returns error",
			input: changesets.ConfigureChainsForLanesFromTopologyConfig{
				Topology: &offchain.EnvironmentTopology{
					NOPTopology: &offchain.NOPTopology{
						NOPs:       []offchain.NOPConfig{{Alias: "nop-1"}},
						Committees: map[string]offchain.CommitteeConfig{},
					},
				},
			},
			errContains: "no committees defined in topology",
		},
		{
			name: "unknown chain selector returns error",
			input: changesets.ConfigureChainsForLanesFromTopologyConfig{
				Topology: &offchain.EnvironmentTopology{
					NOPTopology: &offchain.NOPTopology{
						NOPs: []offchain.NOPConfig{{Alias: "nop-1"}},
						Committees: map[string]offchain.CommitteeConfig{
							"default": {Qualifier: "default"},
						},
					},
				},
				Chains: []changesets.PartialChainConfig{
					{ChainSelector: 9999},
				},
			},
			errContains: "chain selector 9999 is not available in environment",
		},
	}

	env := newCommitteeVerifierTestEnv(t, []uint64{sel1})

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(env, tc.input)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.errContains)
		})
	}
}

func TestConfigureChainsForLanesFromTopology_AdapterNotRegistered(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	laneRegistry := lanes.GetLaneAdapterRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, laneRegistry, nil)

	env := newCommitteeVerifierTestEnv(t, []uint64{sel1, sel2})

	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner1"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							strconv.FormatUint(sel2, 10): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: sel1,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							sel2: {GasForVerification: 100000},
						},
					},
				},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no committee verifier contract adapter registered")
}

func TestConfigureChainsForLanesFromTopology_AddressResolutionFails(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		resolveErr: fmt.Errorf("contract not deployed"),
	})
	laneRegistry := lanes.GetLaneAdapterRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, laneRegistry, nil)

	env := newCommitteeVerifierTestEnv(t, []uint64{sel1, sel2})

	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner1"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							strconv.FormatUint(sel2, 10): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: sel1,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							sel2: {GasForVerification: 100000},
						},
					},
				},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "contract not deployed")
}

func TestConfigureChainsForLanesFromTopology_MissingSignerForNOP(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", sel1): {
				{Address: "0xVerifier1", ChainSelector: sel1, Qualifier: "default"},
				{Address: "0xResolver1", ChainSelector: sel1, Qualifier: "default"},
			},
		},
	})
	laneRegistry := lanes.GetLaneAdapterRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, laneRegistry, nil)

	env := newCommitteeVerifierTestEnv(t, []uint64{sel1, sel2})

	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1"},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							strconv.FormatUint(sel2, 10): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: sel1,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							sel2: {GasForVerification: 100000},
						},
					},
				},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing signer_address")
}

func TestConfigureChainsForLanesFromTopology_ResolvesAndDelegates(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	verifierAddr := "0xVerifier1"
	resolverAddr := "0xResolver1"

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", sel1): {
				{Address: verifierAddr, ChainSelector: sel1, Qualifier: "default", Type: "CommitteeVerifier", Version: semver.MustParse("2.0.0")},
				{Address: resolverAddr, ChainSelector: sel1, Qualifier: "default", Type: "CommitteeVerifierResolver", Version: semver.MustParse("2.0.0")},
			},
		},
	})

	mockAdapter := getOrRegisterMockLaneAdapter(t)
	laneRegistry := lanes.GetLaneAdapterRegistry()

	mcmsReg := cs_core.GetRegistry()
	mcmsReg.RegisterMCMSReader("evm", &lanesTest_MockReader{})

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, laneRegistry, mcmsReg)

	env := newCommitteeVerifierTestEnvWithContracts(t, []uint64{sel1, sel2}, verifierAddr, resolverAddr)

	output, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner1"}},
					{Alias: "nop-2", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner2"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							strconv.FormatUint(sel2, 10): {NOPAliases: []string{"nop-1", "nop-2"}, Threshold: 2},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: sel1,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							sel2: {
								AllowlistEnabled:   true,
								GasForVerification: 100000,
								FeeUSDCents:        50,
								PayloadSizeBytes:   1024,
							},
						},
					},
				},
				DefaultExecutor: datastore.AddressRef{
					ChainSelector: sel1,
					Type:          "Executor",
					Version:       semver.MustParse("1.0.0"),
				},
				FeeQuoterDestChainConfigOverrides: fqOverride(true, 50000),
				ExecutorDestChainConfig: lanes.ExecutorDestChainConfig{
					Enabled: true,
				},
				AddressBytesLength:   20,
				BaseExecutionGasCost: 80000,
				RemoteChains: map[uint64]changesets.RemoteLaneConfig{
					sel2: {
						TestRouter: true,
						Chain: lanes.ChainDefinition{
							Selector: sel2,
							DefaultExecutor: datastore.AddressRef{
								ChainSelector: sel2,
								Type:          "Executor",
								Version:       semver.MustParse("1.0.0"),
							},
							FeeQuoterDestChainConfigOverrides: fqOverride(true, 50000),
							ExecutorDestChainConfig: lanes.ExecutorDestChainConfig{
								Enabled: true,
							},
							AddressBytesLength:   20,
							BaseExecutionGasCost: 80000,
						},
					},
				},
			},
		},
		MCMS: mcms.Input{},
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
	assert.Equal(t, 2, mockAdapter.routerCalls)
	assert.Equal(t, 2, mockAdapter.testRouterCalls)
}

func TestConfigureChainsForLanesFromTopology_UsesRouterAndTestRouter(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel3 := chainsel.TEST_90000003.Selector

	verifierAddr := "0xVerifier1"
	resolverAddr := "0xResolver1"

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", sel1): {
				{Address: verifierAddr, ChainSelector: sel1, Qualifier: "default", Type: "CommitteeVerifier", Version: semver.MustParse("2.0.0")},
				{Address: resolverAddr, ChainSelector: sel1, Qualifier: "default", Type: "CommitteeVerifierResolver", Version: semver.MustParse("2.0.0")},
			},
		},
	})

	mockAdapter := getOrRegisterMockLaneAdapter(t)
	laneRegistry := lanes.GetLaneAdapterRegistry()

	mcmsReg := cs_core.GetRegistry()
	mcmsReg.RegisterMCMSReader("evm", &lanesTest_MockReader{})

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, laneRegistry, mcmsReg)
	env := newCommitteeVerifierTestEnvWithContracts(t, []uint64{sel1, sel2, sel3}, verifierAddr, resolverAddr)

	remoteLane := func(selector uint64) changesets.RemoteLaneConfig {
		return changesets.RemoteLaneConfig{
			Chain: lanes.ChainDefinition{
				Selector: selector,
				DefaultExecutor: datastore.AddressRef{
					ChainSelector: selector,
					Type:          "Executor",
					Version:       semver.MustParse("1.0.0"),
				},
				FeeQuoterDestChainConfigOverrides: fqOverride(true, 50000),
				ExecutorDestChainConfig: lanes.ExecutorDestChainConfig{
					Enabled: true,
				},
				AddressBytesLength:   20,
				BaseExecutionGasCost: 80000,
			},
		}
	}

	output, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner1"}},
					{Alias: "nop-2", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner2"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							strconv.FormatUint(sel2, 10): {NOPAliases: []string{"nop-1", "nop-2"}, Threshold: 2},
							strconv.FormatUint(sel3, 10): {NOPAliases: []string{"nop-1", "nop-2"}, Threshold: 2},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: sel1,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							sel2: {
								AllowlistEnabled:   true,
								GasForVerification: 100000,
								FeeUSDCents:        50,
								PayloadSizeBytes:   1024,
							},
							sel3: {
								AllowlistEnabled:   true,
								GasForVerification: 100000,
								FeeUSDCents:        50,
								PayloadSizeBytes:   1024,
							},
						},
					},
				},
				DefaultExecutor: datastore.AddressRef{
					ChainSelector: sel1,
					Type:          "Executor",
					Version:       semver.MustParse("1.0.0"),
				},
				FeeQuoterDestChainConfigOverrides: fqOverride(true, 50000),
				ExecutorDestChainConfig: lanes.ExecutorDestChainConfig{
					Enabled: true,
				},
				AddressBytesLength:   20,
				BaseExecutionGasCost: 80000,
				RemoteChains: map[uint64]changesets.RemoteLaneConfig{
					sel2: remoteLane(sel2),
					sel3: func() changesets.RemoteLaneConfig {
						cfg := remoteLane(sel3)
						cfg.TestRouter = true
						return cfg
					}(),
				},
			},
		},
		MCMS: mcms.Input{},
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
	assert.Equal(t, 4, mockAdapter.routerCalls)
	assert.Equal(t, 2, mockAdapter.testRouterCalls)
	require.Len(t, mockAdapter.sourceInputs, 4)
	require.Len(t, mockAdapter.destInputs, 4)

	routerInputs := 0
	testRouterInputs := 0
	for _, input := range append(append([]lanes.UpdateLanesInput{}, mockAdapter.sourceInputs...), mockAdapter.destInputs...) {
		if input.TestRouter {
			testRouterInputs++
			assert.Equal(t, []byte("0xTestRouter"), input.Source.Router)
			assert.Equal(t, []byte("0xTestRouter"), input.Dest.Router)
		} else {
			routerInputs++
			assert.Equal(t, []byte("0xRouter"), input.Source.Router)
			assert.Equal(t, []byte("0xRouter"), input.Dest.Router)
		}
	}
	assert.Equal(t, 4, routerInputs)
	assert.Equal(t, 4, testRouterInputs)
}

func TestConfigureChainsForLanesFromTopology_CommitteeNotFound(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{},
	})
	laneRegistry := lanes.GetLaneAdapterRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, laneRegistry, nil)

	env := newCommitteeVerifierTestEnv(t, []uint64{sel1, sel2})

	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner1"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {Qualifier: "default"},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: sel1,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "nonexistent",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							sel2: {},
						},
					},
				},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), `committee "nonexistent" not found`)
}
