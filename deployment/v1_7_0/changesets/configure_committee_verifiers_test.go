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
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

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

var _ adapters.ChainFamily = (*mockChainFamilyForVerifiers)(nil)

type mockChainFamilyForVerifiers struct {
	sequenceErr error
}

func (m *mockChainFamilyForVerifiers) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return []byte(ref.Address), nil
}

func (m *mockChainFamilyForVerifiers) ConfigureChainForLanes() *cldf_ops.Sequence[adapters.ConfigureChainForLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-chain-for-lanes",
		semver.MustParse("1.0.0"),
		"Mock sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.ConfigureChainForLanesInput) (sequences.OnChainOutput, error) {
			if m.sequenceErr != nil {
				return sequences.OnChainOutput{}, m.sequenceErr
			}
			return sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{
					{
						ChainSelector: input.ChainSelector,
						Address:       input.Router,
						Type:          "Router",
						Version:       semver.MustParse("1.0.0"),
					},
				},
				BatchOps: []mcms_types.BatchOperation{},
			}, nil
		},
	)
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
			Type: "CommitteeVerifier", Version: semver.MustParse("1.7.0"),
		}))
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel, Address: resolverAddr, Qualifier: "default",
			Type: "CommitteeVerifierResolver", Version: semver.MustParse("1.7.0"),
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

func TestConfigureChainsForLanesFromTopology_Validation(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	chainFamilyRegistry := adapters.NewChainFamilyRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, chainFamilyRegistry, nil)

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
	chainFamilyRegistry := adapters.NewChainFamilyRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, chainFamilyRegistry, nil)

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
				Router:        datastore.AddressRef{Type: "Router", Version: semver.MustParse("1.0.0")},
				OnRamp:        datastore.AddressRef{Type: "OnRamp", Version: semver.MustParse("1.0.0")},
				FeeQuoter:     datastore.AddressRef{Type: "FeeQuoter", Version: semver.MustParse("1.0.0")},
				OffRamp:       datastore.AddressRef{Type: "OffRamp", Version: semver.MustParse("1.0.0")},
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
	chainFamilyRegistry := adapters.NewChainFamilyRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, chainFamilyRegistry, nil)

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
				Router:        datastore.AddressRef{Type: "Router", Version: semver.MustParse("1.0.0")},
				OnRamp:        datastore.AddressRef{Type: "OnRamp", Version: semver.MustParse("1.0.0")},
				FeeQuoter:     datastore.AddressRef{Type: "FeeQuoter", Version: semver.MustParse("1.0.0")},
				OffRamp:       datastore.AddressRef{Type: "OffRamp", Version: semver.MustParse("1.0.0")},
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
	chainFamilyRegistry := adapters.NewChainFamilyRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, chainFamilyRegistry, nil)

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
				Router:        datastore.AddressRef{Type: "Router", Version: semver.MustParse("1.0.0")},
				OnRamp:        datastore.AddressRef{Type: "OnRamp", Version: semver.MustParse("1.0.0")},
				FeeQuoter:     datastore.AddressRef{Type: "FeeQuoter", Version: semver.MustParse("1.0.0")},
				OffRamp:       datastore.AddressRef{Type: "OffRamp", Version: semver.MustParse("1.0.0")},
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
				{Address: verifierAddr, ChainSelector: sel1, Qualifier: "default", Type: "CommitteeVerifier", Version: semver.MustParse("1.7.0")},
				{Address: resolverAddr, ChainSelector: sel1, Qualifier: "default", Type: "CommitteeVerifierResolver", Version: semver.MustParse("1.7.0")},
			},
		},
	})

	chainFamilyRegistry := adapters.NewChainFamilyRegistry()
	chainFamilyRegistry.RegisterChainFamily(chainsel.FamilyEVM, &mockChainFamilyForVerifiers{})

	mcmsReg := cs_core.GetRegistry()
	mcmsReg.RegisterMCMSReader("evm", &lanesTest_MockReader{})

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, chainFamilyRegistry, mcmsReg)

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
				Router:        datastore.AddressRef{Type: "Router", Version: semver.MustParse("1.0.0")},
				OnRamp:        datastore.AddressRef{Type: "OnRamp", Version: semver.MustParse("1.0.0")},
				FeeQuoter:     datastore.AddressRef{Type: "FeeQuoter", Version: semver.MustParse("1.0.0")},
				OffRamp:       datastore.AddressRef{Type: "OffRamp", Version: semver.MustParse("1.0.0")},
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
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					sel2: {
						AllowTrafficFrom: true,
						OnRamps: []datastore.AddressRef{
							{Type: "OnRamp", Version: semver.MustParse("1.0.0")},
						},
						OffRamp:         datastore.AddressRef{Type: "OffRamp", Version: semver.MustParse("1.0.0")},
						DefaultExecutor: datastore.AddressRef{Type: "Executor", Version: semver.MustParse("1.0.0"), Address: "0xExec"},
						FeeQuoterDestChainConfig: adapters.FeeQuoterDestChainConfig{
							IsEnabled:    true,
							MaxDataBytes: 50000,
						},
					},
				},
			},
		},
		MCMS: mcms.Input{},
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestConfigureChainsForLanesFromTopology_CommitteeNotFound(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{},
	})
	chainFamilyRegistry := adapters.NewChainFamilyRegistry()
	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, chainFamilyRegistry, nil)

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
