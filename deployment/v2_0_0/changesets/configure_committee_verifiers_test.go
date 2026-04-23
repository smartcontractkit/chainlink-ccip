package changesets_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
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

func (m *mockCommitteeVerifierContractAdapter) GetCommitteeVerifierResolver(
	ds datastore.DataStore,
	chainSelector uint64,
	qualifier string,
) ([]datastore.AddressRef, error) {
	refs, err := m.ResolveCommitteeVerifierContracts(ds, chainSelector, qualifier)
	if err != nil {
		return nil, err
	}
	resolverType := datastore.ContractType("CommitteeVerifierResolver")
	var out []datastore.AddressRef
	for _, ref := range refs {
		if ref.Type == resolverType {
			out = append(out, ref)
		}
	}
	if len(out) > 0 {
		return out, nil
	}
	if len(refs) == 1 {
		return refs, nil
	}
	return nil, fmt.Errorf("no CommitteeVerifierResolver ref for default lane CCVs (chain %d qualifier %q)", chainSelector, qualifier)
}

func newResolverTestEnv(t *testing.T, selectors []uint64) deployment.Environment {
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

func TestTopologyCommitteePopulator_AdapterNotRegistered(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()

	populator := changesets.NewTopologyCommitteePopulator(committeeRegistry, &offchain.EnvironmentTopology{
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
	})

	env := newResolverTestEnv(t, []uint64{sel1, sel2})

	_, err := populator.PopulateCommitteeConfig(env, sel1, []lanes.CommitteeVerifierInput{
		{
			CommitteeQualifier: "default",
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainInput{
				sel2: {GasForVerification: 100000},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no committee verifier contract adapter")
}

func TestTopologyCommitteePopulator_AddressResolutionFails(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		resolveErr: fmt.Errorf("contract not deployed"),
	})

	populator := changesets.NewTopologyCommitteePopulator(committeeRegistry, &offchain.EnvironmentTopology{
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
	})

	env := newResolverTestEnv(t, []uint64{sel1, sel2})

	_, err := populator.PopulateCommitteeConfig(env, sel1, []lanes.CommitteeVerifierInput{
		{
			CommitteeQualifier: "default",
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainInput{
				sel2: {GasForVerification: 100000},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "contract not deployed")
}

func TestTopologyCommitteePopulator_MissingSignerForNOP(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", sel1): {
				{Address: "0xVerifier1", ChainSelector: sel1, Qualifier: "default"},
			},
		},
	})

	populator := changesets.NewTopologyCommitteePopulator(committeeRegistry, &offchain.EnvironmentTopology{
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
	})

	env := newResolverTestEnv(t, []uint64{sel1, sel2})

	_, err := populator.PopulateCommitteeConfig(env, sel1, []lanes.CommitteeVerifierInput{
		{
			CommitteeQualifier: "default",
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainInput{
				sel2: {GasForVerification: 100000},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing signer_address")
}

func TestTopologyCommitteePopulator_CommitteeNotFound(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{},
	})

	populator := changesets.NewTopologyCommitteePopulator(committeeRegistry, &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xSigner1"}},
			},
			Committees: map[string]offchain.CommitteeConfig{
				"default": {Qualifier: "default"},
			},
		},
	})

	env := newResolverTestEnv(t, []uint64{sel1, sel2})

	_, err := populator.PopulateCommitteeConfig(env, sel1, []lanes.CommitteeVerifierInput{
		{
			CommitteeQualifier: "nonexistent",
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainInput{
				sel2: {},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), `committee "nonexistent" not found`)
}

func TestTopologyCommitteePopulator_PopulatesSuccessfully(t *testing.T) {
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

	populator := changesets.NewTopologyCommitteePopulator(committeeRegistry, &offchain.EnvironmentTopology{
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
	})

	env := newResolverTestEnv(t, []uint64{sel1, sel2})

	result, err := populator.PopulateCommitteeConfig(env, sel1, []lanes.CommitteeVerifierInput{
		{
			CommitteeQualifier: "default",
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainInput{
				sel2: {
					AllowlistEnabled:   true,
					GasForVerification: 100000,
					FeeUSDCents:        50,
					PayloadSizeBytes:   1024,
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, result, 1)

	cv := result[0]
	require.Len(t, cv.CommitteeVerifier, 2)
	assert.Equal(t, verifierAddr, cv.CommitteeVerifier[0].Address)
	assert.Equal(t, resolverAddr, cv.CommitteeVerifier[1].Address)

	rc, ok := cv.RemoteChains[sel2]
	require.True(t, ok)
	assert.True(t, rc.AllowlistEnabled)
	assert.Equal(t, uint32(100000), rc.GasForVerification)
	assert.Equal(t, uint16(50), rc.FeeUSDCents)
	assert.Equal(t, uint16(1024), rc.PayloadSizeBytes)
	assert.Equal(t, uint8(2), rc.SignatureConfig.Threshold)
	assert.ElementsMatch(t, []string{"0xSigner1", "0xSigner2"}, rc.SignatureConfig.Signers)
}

func TestTopologyCommitteePopulator_SigningKeyErrorPersistedAcrossCalls(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", sel1): {
				{Address: "0xVerifier1", ChainSelector: sel1, Qualifier: "default"},
			},
		},
	})

	mockOffchain := &jdMockOffchain{
		listNodesFn: func(_ context.Context, _ *nodev1.ListNodesRequest, _ ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
			return nil, fmt.Errorf("JD unavailable")
		},
		listNodeChainConfigsFn: func(_ context.Context, _ *nodev1.ListNodeChainConfigsRequest, _ ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
			return nil, fmt.Errorf("JD unavailable")
		},
	}

	populator := changesets.NewTopologyCommitteePopulator(committeeRegistry, &offchain.EnvironmentTopology{
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
	})

	env := newResolverTestEnv(t, []uint64{sel1, sel2})
	env.Offchain = mockOffchain
	env.NodeIDs = []string{"node-1"}

	inputs := []lanes.CommitteeVerifierInput{
		{
			CommitteeQualifier: "default",
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainInput{
				sel2: {GasForVerification: 100000},
			},
		},
	}

	_, err1 := populator.PopulateCommitteeConfig(env, sel1, inputs)
	require.Error(t, err1, "first call should fail because JD is unavailable")
	assert.Contains(t, err1.Error(), "failed to fetch signing keys")

	_, err2 := populator.PopulateCommitteeConfig(env, sel1, inputs)
	require.Error(t, err2, "second call must also fail with the persisted error")
	assert.Contains(t, err2.Error(), "failed to fetch signing keys")
}
