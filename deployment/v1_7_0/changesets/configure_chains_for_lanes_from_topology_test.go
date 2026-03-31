package changesets_test

import (
	"context"
	"fmt"
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
	cldf_offchain "github.com/smartcontractkit/chainlink-deployments-framework/offchain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	csav1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/csa"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

type mockChainFamilyAdapter struct {
	addressPrefix string
	inputs        []adapters.ConfigureChainForLanesInput
	err           error
}

func (m *mockChainFamilyAdapter) ConfigureChainForLanes() *cldf_ops.Sequence[adapters.ConfigureChainForLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-chain-for-lanes",
		semver.MustParse("1.0.0"),
		"mock",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input adapters.ConfigureChainForLanesInput) (sequences.OnChainOutput, error) {
			if m.err != nil {
				return sequences.OnChainOutput{}, m.err
			}
			m.inputs = append(m.inputs, input)
			return sequences.OnChainOutput{}, nil
		},
	)
}

func (m *mockChainFamilyAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return []byte(m.addressPrefix + ref.Address), nil
}

type stubOffchain struct {
	jobv1.JobServiceClient
	nodev1.NodeServiceClient
	csav1.CSAServiceClient
}

var _ cldf_offchain.Client = (*stubOffchain)(nil)

type jdMockOffchain struct {
	stubOffchain
	listNodesFn            func(context.Context, *nodev1.ListNodesRequest, ...grpc.CallOption) (*nodev1.ListNodesResponse, error)
	listNodeChainConfigsFn func(context.Context, *nodev1.ListNodeChainConfigsRequest, ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error)
}

func (m *jdMockOffchain) ListNodes(ctx context.Context, in *nodev1.ListNodesRequest, opts ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
	return m.listNodesFn(ctx, in, opts...)
}

func (m *jdMockOffchain) ListNodeChainConfigs(ctx context.Context, in *nodev1.ListNodeChainConfigsRequest, opts ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
	return m.listNodeChainConfigsFn(ctx, in, opts...)
}

func newConfigureChainsTestEnv(t *testing.T, localSelectors []uint64, offchainClient cldf_offchain.Client) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)

	chains := make(map[uint64]cldf_chain.BlockChain, len(localSelectors))
	for _, selector := range localSelectors {
		chains[selector] = cldfevm.Chain{Selector: selector}
	}

	return deployment.Environment{
		Name:        "test",
		BlockChains: cldf_chain.NewBlockChains(chains),
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		Logger:      lggr,
		Offchain:    offchainClient,
		NodeIDs:     []string{"node-1"},
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

func addAddress(t *testing.T, ds datastore.MutableDataStore, ref datastore.AddressRef) {
	t.Helper()
	require.NoError(t, ds.Addresses().Add(ref))
}

func testRef(chainSelector uint64, address string, contractType datastore.ContractType) datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: chainSelector,
		Address:       address,
		Type:          contractType,
		Version:       semver.MustParse("1.0.0"),
	}
}

func TestConfigureChainsForLanesFromTopology_HappyPathAndCrossFamily(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteEVM := chainsel.TEST_90000002.Selector
	remoteSolana := chainsel.SOLANA_DEVNET.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xrouter", "Router"))
	addAddress(t, ds, testRef(localSelector, "0xonramp", "OnRamp"))
	addAddress(t, ds, testRef(localSelector, "0xfeequoter", "FeeQuoter"))
	addAddress(t, ds, testRef(localSelector, "0xofframp", "OffRamp"))
	addAddress(t, ds, testRef(localSelector, "0xexecutor", "Executor"))
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
	addAddress(t, ds, testRef(remoteEVM, "0xremote-evm-onramp", "OnRamp"))
	addAddress(t, ds, testRef(remoteEVM, "0xremote-evm-offramp", "OffRamp"))
	addAddress(t, ds, testRef(remoteSolana, "solana-onramp", "OnRamp"))
	addAddress(t, ds, testRef(remoteSolana, "solana-offramp", "OffRamp"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {
				testRef(localSelector, "0xverifier", "CommitteeVerifier"),
				testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"),
			},
		},
	})

	localAdapter := &mockChainFamilyAdapter{addressPrefix: "local:"}
	remoteSolanaAdapter := &mockChainFamilyAdapter{addressPrefix: "sol:"}
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, localAdapter)
	registry.RegisterChainFamily(chainsel.FamilySolana, remoteSolanaAdapter)

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, registry, changesetscore.GetRegistry())
	output, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xsigner-1"}},
					{Alias: "nop-2", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xsigner-2"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteEVM):    {NOPAliases: []string{"nop-1"}, Threshold: 1},
							fmt.Sprintf("%d", remoteSolana): {NOPAliases: []string{"nop-1", "nop-2"}, Threshold: 2},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				Router:        testRef(localSelector, "0xrouter", "Router"),
				OnRamp:        testRef(localSelector, "0xonramp", "OnRamp"),
				FeeQuoter:     testRef(localSelector, "0xfeequoter", "FeeQuoter"),
				OffRamp:       testRef(localSelector, "0xofframp", "OffRamp"),
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							remoteEVM:    {FeeUSDCents: 10, GasForVerification: 20, PayloadSizeBytes: 30},
							remoteSolana: {FeeUSDCents: 40, GasForVerification: 50, PayloadSizeBytes: 60},
						},
					},
				},
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					remoteEVM: {
						OnRamps:                 []datastore.AddressRef{testRef(remoteEVM, "0xremote-evm-onramp", "OnRamp")},
						OffRamp:                 testRef(remoteEVM, "0xremote-evm-offramp", "OffRamp"),
						DefaultExecutor:         testRef(localSelector, "0xexecutor", "Executor"),
						DefaultInboundCCVs:      []datastore.AddressRef{testRef(localSelector, "0xverifier", "CommitteeVerifier")},
						LaneMandatedInboundCCVs: []datastore.AddressRef{testRef(localSelector, "0xverifier", "CommitteeVerifier")},
						DefaultOutboundCCVs:     []datastore.AddressRef{testRef(localSelector, "0xverifier", "CommitteeVerifier")},
						LaneMandatedOutboundCCVs: []datastore.AddressRef{
							testRef(localSelector, "0xverifier", "CommitteeVerifier"),
						},
					},
					remoteSolana: {
						OnRamps:         []datastore.AddressRef{testRef(remoteSolana, "solana-onramp", "OnRamp")},
						OffRamp:         testRef(remoteSolana, "solana-offramp", "OffRamp"),
						DefaultExecutor: testRef(localSelector, "0xexecutor", "Executor"),
					},
				},
			},
		},
	})
	require.NoError(t, err)
	assert.Empty(t, output.MCMSTimelockProposals)
	require.Len(t, localAdapter.inputs, 1)
	assert.Empty(t, remoteSolanaAdapter.inputs, "remote Solana adapter should not be called directly; it is only used for address conversion")

	input := localAdapter.inputs[0]
	assert.Equal(t, "0xrouter", input.Router)
	assert.Equal(t, "0xonramp", input.OnRamp)
	assert.Equal(t, "0xfeequoter", input.FeeQuoter)
	assert.Equal(t, "0xofframp", input.OffRamp)
	require.Len(t, input.RemoteChains, 2)
	assert.Equal(t, []byte("local:0xremote-evm-offramp"), input.RemoteChains[remoteEVM].OffRamp)
	assert.Equal(t, []byte("sol:solana-offramp"), input.RemoteChains[remoteSolana].OffRamp)
	assert.Equal(t, []byte("local:0xremote-evm-onramp"), input.RemoteChains[remoteEVM].OnRamps[0])
	assert.Equal(t, []byte("sol:solana-onramp"), input.RemoteChains[remoteSolana].OnRamps[0])
	assert.Equal(t, "0xexecutor", input.RemoteChains[remoteEVM].DefaultExecutor)
	assert.Equal(t, []string{"0xverifier"}, input.RemoteChains[remoteEVM].DefaultInboundCCVs)
	require.Len(t, input.CommitteeVerifiers, 1)
	assert.ElementsMatch(t, []string{"0xsigner-1"}, input.CommitteeVerifiers[0].RemoteChains[remoteEVM].SignatureConfig.Signers)
	assert.ElementsMatch(t, []string{"0xsigner-1", "0xsigner-2"}, input.CommitteeVerifiers[0].RemoteChains[remoteSolana].SignatureConfig.Signers)
}

func TestConfigureChainsForLanesFromTopology_JDFallback(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.SOLANA_DEVNET.Selector

	mockOffchain := &jdMockOffchain{
		listNodesFn: func(_ context.Context, _ *nodev1.ListNodesRequest, _ ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
			return &nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}}}, nil
		},
		listNodeChainConfigsFn: func(_ context.Context, _ *nodev1.ListNodeChainConfigsRequest, _ ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
			return &nodev1.ListNodeChainConfigsResponse{
				ChainConfigs: []*nodev1.ChainConfig{
					{
						NodeId: "node-1",
						Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
						Ocr2Config: &nodev1.OCR2Config{
							OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0xjd-signer"},
						},
					},
				},
			}, nil
		},
	}

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, mockOffchain)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xrouter", "Router"))
	addAddress(t, ds, testRef(localSelector, "0xonramp", "OnRamp"))
	addAddress(t, ds, testRef(localSelector, "0xfeequoter", "FeeQuoter"))
	addAddress(t, ds, testRef(localSelector, "0xofframp", "OffRamp"))
	addAddress(t, ds, testRef(localSelector, "0xexecutor", "Executor"))
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
	addAddress(t, ds, testRef(remoteSelector, "remote-onramp", "OnRamp"))
	addAddress(t, ds, testRef(remoteSelector, "remote-offramp", "OffRamp"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {
				testRef(localSelector, "0xverifier", "CommitteeVerifier"),
				testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"),
			},
		},
	})

	localAdapter := &mockChainFamilyAdapter{addressPrefix: "local:"}
	remoteAdapter := &mockChainFamilyAdapter{addressPrefix: "sol:"}
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, localAdapter)
	registry.RegisterChainFamily(chainsel.FamilySolana, remoteAdapter)

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, registry, changesetscore.GetRegistry())
	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilySolana: "sol-signer"}}},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteSelector): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				Router:        testRef(localSelector, "0xrouter", "Router"),
				OnRamp:        testRef(localSelector, "0xonramp", "OnRamp"),
				FeeQuoter:     testRef(localSelector, "0xfeequoter", "FeeQuoter"),
				OffRamp:       testRef(localSelector, "0xofframp", "OffRamp"),
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							remoteSelector: {},
						},
					},
				},
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					remoteSelector: {
						OnRamps:         []datastore.AddressRef{testRef(remoteSelector, "remote-onramp", "OnRamp")},
						OffRamp:         testRef(remoteSelector, "remote-offramp", "OffRamp"),
						DefaultExecutor: testRef(localSelector, "0xexecutor", "Executor"),
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, localAdapter.inputs, 1)
	assert.Equal(t, []string{"0xjd-signer"}, localAdapter.inputs[0].CommitteeVerifiers[0].RemoteChains[remoteSelector].SignatureConfig.Signers)
}

func TestConfigureChainsForLanesFromTopology_MissingSignerAfterJDFallback(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.SOLANA_DEVNET.Selector
	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {
				testRef(localSelector, "0xverifier", "CommitteeVerifier"),
				testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"),
			},
		},
	})

	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, &mockChainFamilyAdapter{})
	registry.RegisterChainFamily(chainsel.FamilySolana, &mockChainFamilyAdapter{})

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, registry, changesetscore.GetRegistry())
	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{{Alias: "nop-1"}},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteSelector): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing signer_address")
}

func TestConfigureChainsForLanesFromTopology_MissingRemoteAdapter(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.SOLANA_DEVNET.Selector
	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xrouter", "Router"))
	addAddress(t, ds, testRef(localSelector, "0xonramp", "OnRamp"))
	addAddress(t, ds, testRef(localSelector, "0xfeequoter", "FeeQuoter"))
	addAddress(t, ds, testRef(localSelector, "0xofframp", "OffRamp"))
	addAddress(t, ds, testRef(localSelector, "0xexecutor", "Executor"))
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
	addAddress(t, ds, testRef(remoteSelector, "remote-onramp", "OnRamp"))
	addAddress(t, ds, testRef(remoteSelector, "remote-offramp", "OffRamp"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {
				testRef(localSelector, "0xverifier", "CommitteeVerifier"),
				testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"),
			},
		},
	})

	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, &mockChainFamilyAdapter{})

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, registry, changesetscore.GetRegistry())
	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xsigner"}}},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteSelector): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				Router:        testRef(localSelector, "0xrouter", "Router"),
				OnRamp:        testRef(localSelector, "0xonramp", "OnRamp"),
				FeeQuoter:     testRef(localSelector, "0xfeequoter", "FeeQuoter"),
				OffRamp:       testRef(localSelector, "0xofframp", "OffRamp"),
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					remoteSelector: {
						OnRamps:         []datastore.AddressRef{testRef(remoteSelector, "remote-onramp", "OnRamp")},
						OffRamp:         testRef(remoteSelector, "remote-offramp", "OffRamp"),
						DefaultExecutor: testRef(localSelector, "0xexecutor", "Executor"),
					},
				},
			},
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no adapter registered for remote chain family")
}

func TestConfigureChainsForLanesFromTopology_PerSourceDestinationConfig(t *testing.T) {
	chainA := chainsel.TEST_90000001.Selector
	chainB := chainsel.TEST_90000002.Selector
	sharedDest := chainsel.TEST_90000003.Selector

	env := newConfigureChainsTestEnv(t, []uint64{chainA, chainB}, nil)
	ds := datastore.NewMemoryDataStore()
	for _, sel := range []uint64{chainA, chainB} {
		addAddress(t, ds, testRef(sel, fmt.Sprintf("0xrouter-%d", sel), "Router"))
		addAddress(t, ds, testRef(sel, fmt.Sprintf("0xonramp-%d", sel), "OnRamp"))
		addAddress(t, ds, testRef(sel, fmt.Sprintf("0xfeequoter-%d", sel), "FeeQuoter"))
		addAddress(t, ds, testRef(sel, fmt.Sprintf("0xofframp-%d", sel), "OffRamp"))
		addAddress(t, ds, testRef(sel, fmt.Sprintf("0xexecutor-%d", sel), "Executor"))
		addAddress(t, ds, testRef(sel, fmt.Sprintf("0xverifier-%d", sel), "CommitteeVerifier"))
	}
	addAddress(t, ds, testRef(sharedDest, "0xdest-onramp", "OnRamp"))
	addAddress(t, ds, testRef(sharedDest, "0xdest-offramp", "OffRamp"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", chainA): {testRef(chainA, fmt.Sprintf("0xverifier-%d", chainA), "CommitteeVerifier")},
			fmt.Sprintf("%d:default", chainB): {testRef(chainB, fmt.Sprintf("0xverifier-%d", chainB), "CommitteeVerifier")},
		},
	})

	evmAdapter := &mockChainFamilyAdapter{addressPrefix: "evm:"}
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, evmAdapter)

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, registry, changesetscore.GetRegistry())
	output, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xsigner"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", sharedDest): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: chainA,
				Router:        testRef(chainA, fmt.Sprintf("0xrouter-%d", chainA), "Router"),
				OnRamp:        testRef(chainA, fmt.Sprintf("0xonramp-%d", chainA), "OnRamp"),
				FeeQuoter:     testRef(chainA, fmt.Sprintf("0xfeequoter-%d", chainA), "FeeQuoter"),
				OffRamp:       testRef(chainA, fmt.Sprintf("0xofframp-%d", chainA), "OffRamp"),
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{sharedDest: {}}},
				},
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					sharedDest: {
						OnRamps:                []datastore.AddressRef{testRef(sharedDest, "0xdest-onramp", "OnRamp")},
						OffRamp:                testRef(sharedDest, "0xdest-offramp", "OffRamp"),
						DefaultExecutor:        testRef(chainA, fmt.Sprintf("0xexecutor-%d", chainA), "Executor"),
						FeeQuoterDestChainConfig: adapters.FeeQuoterDestChainConfig{MaxDataBytes: 1000},
						ExecutorDestChainConfig:  adapters.ExecutorDestChainConfig{USDCentsFee: 100, Enabled: true},
					},
				},
			},
			{
				ChainSelector: chainB,
				Router:        testRef(chainB, fmt.Sprintf("0xrouter-%d", chainB), "Router"),
				OnRamp:        testRef(chainB, fmt.Sprintf("0xonramp-%d", chainB), "OnRamp"),
				FeeQuoter:     testRef(chainB, fmt.Sprintf("0xfeequoter-%d", chainB), "FeeQuoter"),
				OffRamp:       testRef(chainB, fmt.Sprintf("0xofframp-%d", chainB), "OffRamp"),
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{sharedDest: {}}},
				},
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					sharedDest: {
						OnRamps:                []datastore.AddressRef{testRef(sharedDest, "0xdest-onramp", "OnRamp")},
						OffRamp:                testRef(sharedDest, "0xdest-offramp", "OffRamp"),
						DefaultExecutor:        testRef(chainB, fmt.Sprintf("0xexecutor-%d", chainB), "Executor"),
						FeeQuoterDestChainConfig: adapters.FeeQuoterDestChainConfig{MaxDataBytes: 2000},
						ExecutorDestChainConfig:  adapters.ExecutorDestChainConfig{USDCentsFee: 200, Enabled: true},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	assert.Empty(t, output.MCMSTimelockProposals)
	require.Len(t, evmAdapter.inputs, 2)

	inputsByChain := make(map[uint64]adapters.ConfigureChainForLanesInput, 2)
	for _, inp := range evmAdapter.inputs {
		inputsByChain[inp.ChainSelector] = inp
	}

	inputA := inputsByChain[chainA]
	assert.Equal(t, uint32(1000), inputA.RemoteChains[sharedDest].FeeQuoterDestChainConfig.MaxDataBytes)
	assert.Equal(t, uint16(100), inputA.RemoteChains[sharedDest].ExecutorDestChainConfig.USDCentsFee)

	inputB := inputsByChain[chainB]
	assert.Equal(t, uint32(2000), inputB.RemoteChains[sharedDest].FeeQuoterDestChainConfig.MaxDataBytes)
	assert.Equal(t, uint16(200), inputB.RemoteChains[sharedDest].ExecutorDestChainConfig.USDCentsFee)
}

func TestConfigureChainsForLanesFromTopology_AcceptsAnyRouterAddress(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xtest-router", "Router"))
	addAddress(t, ds, testRef(localSelector, "0xonramp", "OnRamp"))
	addAddress(t, ds, testRef(localSelector, "0xfeequoter", "FeeQuoter"))
	addAddress(t, ds, testRef(localSelector, "0xofframp", "OffRamp"))
	addAddress(t, ds, testRef(localSelector, "0xexecutor", "Executor"))
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(remoteSelector, "0xremote-onramp", "OnRamp"))
	addAddress(t, ds, testRef(remoteSelector, "0xremote-offramp", "OffRamp"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {testRef(localSelector, "0xverifier", "CommitteeVerifier")},
		},
	})

	evmAdapter := &mockChainFamilyAdapter{addressPrefix: "evm:"}
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, evmAdapter)

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, registry, changesetscore.GetRegistry())
	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xsigner"}},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteSelector): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				Router:        testRef(localSelector, "0xtest-router", "Router"),
				OnRamp:        testRef(localSelector, "0xonramp", "OnRamp"),
				FeeQuoter:     testRef(localSelector, "0xfeequoter", "FeeQuoter"),
				OffRamp:       testRef(localSelector, "0xofframp", "OffRamp"),
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					remoteSelector: {
						OnRamps:         []datastore.AddressRef{testRef(remoteSelector, "0xremote-onramp", "OnRamp")},
						OffRamp:         testRef(remoteSelector, "0xremote-offramp", "OffRamp"),
						DefaultExecutor: testRef(localSelector, "0xexecutor", "Executor"),
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, evmAdapter.inputs, 1)
	assert.Equal(t, "0xtest-router", evmAdapter.inputs[0].Router)
}

func TestConfigureChainsForLanesFromTopology_OnlyFetchesSigningKeysForCommitteeNOPs(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	mockOffchain := &jdMockOffchain{
		listNodesFn: func(_ context.Context, _ *nodev1.ListNodesRequest, _ ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
			return &nodev1.ListNodesResponse{Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "committee-nop-1"},
				{Id: "node-2", Name: "committee-nop-2"},
			}}, nil
		},
		listNodeChainConfigsFn: func(_ context.Context, _ *nodev1.ListNodeChainConfigsRequest, _ ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
			return &nodev1.ListNodeChainConfigsResponse{
				ChainConfigs: []*nodev1.ChainConfig{
					{
						NodeId: "node-1",
						Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
						Ocr2Config: &nodev1.OCR2Config{
							OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0xsigner-1"},
						},
					},
					{
						NodeId: "node-2",
						Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
						Ocr2Config: &nodev1.OCR2Config{
							OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0xsigner-2"},
						},
					},
				},
			}, nil
		},
	}

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, mockOffchain)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xrouter", "Router"))
	addAddress(t, ds, testRef(localSelector, "0xonramp", "OnRamp"))
	addAddress(t, ds, testRef(localSelector, "0xfeequoter", "FeeQuoter"))
	addAddress(t, ds, testRef(localSelector, "0xofframp", "OffRamp"))
	addAddress(t, ds, testRef(localSelector, "0xexecutor", "Executor"))
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
	addAddress(t, ds, testRef(remoteSelector, "0xremote-onramp", "OnRamp"))
	addAddress(t, ds, testRef(remoteSelector, "0xremote-offramp", "OffRamp"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {
				testRef(localSelector, "0xverifier", "CommitteeVerifier"),
				testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"),
			},
		},
	})

	evmAdapter := &mockChainFamilyAdapter{addressPrefix: "evm:"}
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, evmAdapter)

	cs := changesets.ConfigureChainsForLanesFromTopology(committeeRegistry, registry, changesetscore.GetRegistry())
	_, err := cs.Apply(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: &offchain.EnvironmentTopology{
			NOPTopology: &offchain.NOPTopology{
				NOPs: []offchain.NOPConfig{
					{Alias: "committee-nop-1"},
					{Alias: "committee-nop-2"},
					{Alias: "executor-only-nop"},
				},
				Committees: map[string]offchain.CommitteeConfig{
					"default": {
						Qualifier: "default",
						ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteSelector): {
								NOPAliases: []string{"committee-nop-1", "committee-nop-2"},
								Threshold:  2,
							},
						},
					},
				},
			},
		},
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				Router:        testRef(localSelector, "0xrouter", "Router"),
				OnRamp:        testRef(localSelector, "0xonramp", "OnRamp"),
				FeeQuoter:     testRef(localSelector, "0xfeequoter", "FeeQuoter"),
				OffRamp:       testRef(localSelector, "0xofframp", "OffRamp"),
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							remoteSelector: {},
						},
					},
				},
				RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
					remoteSelector: {
						OnRamps:         []datastore.AddressRef{testRef(remoteSelector, "0xremote-onramp", "OnRamp")},
						OffRamp:         testRef(remoteSelector, "0xremote-offramp", "OffRamp"),
						DefaultExecutor: testRef(localSelector, "0xexecutor", "Executor"),
					},
				},
			},
		},
	})
	require.NoError(t, err,
		"should succeed because JD only needs to know about committee NOPs, not executor-only NOPs")
	require.Len(t, evmAdapter.inputs, 1)
	assert.ElementsMatch(t,
		[]string{"0xsigner-1", "0xsigner-2"},
		evmAdapter.inputs[0].CommitteeVerifiers[0].RemoteChains[remoteSelector].SignatureConfig.Signers,
	)
}

func TestConfigureChainsForLanesFromTopology_VerifyPreconditions(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	cs := changesets.ConfigureChainsForLanesFromTopology(
		adapters.NewCommitteeVerifierContractRegistry(),
		adapters.NewChainFamilyRegistry(),
		changesetscore.GetRegistry(),
	)

	err := cs.VerifyPreconditions(env, changesets.ConfigureChainsForLanesFromTopologyConfig{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "topology is required")
}
