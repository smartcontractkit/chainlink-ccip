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

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

type mockChainFamilyAdapter struct {
	addressPrefix string
	inputs        []adapters.ConfigureChainForLanesInput
	err           error
	addresses     map[uint64]map[string][]byte
	executors     map[uint64]map[string]string
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

func (m *mockChainFamilyAdapter) GetOnRampAddress(_ datastore.DataStore, chainSelector uint64) ([]byte, error) {
	return m.getAddress(chainSelector, "OnRamp")
}

func (m *mockChainFamilyAdapter) GetOffRampAddress(_ datastore.DataStore, chainSelector uint64) ([]byte, error) {
	return m.getAddress(chainSelector, "OffRamp")
}

func (m *mockChainFamilyAdapter) GetFQAddress(_ datastore.DataStore, chainSelector uint64) ([]byte, error) {
	return m.getAddress(chainSelector, "FeeQuoter")
}

func (m *mockChainFamilyAdapter) GetRouterAddress(_ datastore.DataStore, chainSelector uint64) ([]byte, error) {
	return m.getAddress(chainSelector, "Router")
}

func (m *mockChainFamilyAdapter) GetTestRouter(_ datastore.DataStore, chainSelector uint64) ([]byte, error) {
	return m.getAddress(chainSelector, "TestRouter")
}

func (m *mockChainFamilyAdapter) ResolveExecutor(_ datastore.DataStore, chainSelector uint64, qualifier string) (string, error) {
	if m.executors != nil {
		if byQualifier, ok := m.executors[chainSelector]; ok {
			if addr, ok := byQualifier[qualifier]; ok {
				return addr, nil
			}
		}
	}
	return "", fmt.Errorf("executor not found for chain %d qualifier %q", chainSelector, qualifier)
}

func (m *mockChainFamilyAdapter) GetAddressBytesLength() uint8 {
	return 20
}

func (m *mockChainFamilyAdapter) GetChainFamilySelector() [4]byte {
	return [4]byte{0x28, 0x12, 0xd5, 0x2c}
}

func (m *mockChainFamilyAdapter) GetDefaultFeeQuoterDestChainConfig() adapters.FeeQuoterDestChainConfig {
	return adapters.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                32_000,
		MaxPerMsgGasLimit:           8_000_000,
		DestGasPerPayloadByteBase:   16,
		DefaultTokenFeeUSDCents:     25,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          10,
		LinkFeeMultiplierPercent:    90,
	}
}

func (m *mockChainFamilyAdapter) GetDefaultRemoteChainConfig() adapters.RemoteChainDefaults {
	return adapters.RemoteChainDefaults{
		AllowTrafficFrom:          true,
		ExecutorDestChainConfig:   adapters.ExecutorDestChainConfig{USDCentsFee: 0, Enabled: true},
		BaseExecutionGasCost:      175_000,
		TokenReceiverAllowed:      false,
		MessageNetworkFeeUSDCents: 10,
		TokenNetworkFeeUSDCents:   25,
	}
}

func (m *mockChainFamilyAdapter) GetDefaultCommitteeVerifierRemoteChainConfig() adapters.CommitteeVerifierRemoteChainDefaults {
	return adapters.CommitteeVerifierRemoteChainDefaults{
		AllowlistEnabled:   false,
		FeeUSDCents:        0,
		GasForVerification: 60_000,
		PayloadSizeBytes:   390,
	}
}

func (m *mockChainFamilyAdapter) GetDefaultFinalityConfig() finality.Config {
	return finality.Config{
		WaitForFinality: true,
		WaitForSafe:     true,
		BlockDepth:      1,
	}
}

func (m *mockChainFamilyAdapter) getAddress(chainSelector uint64, contractType string) ([]byte, error) {
	if m.addresses != nil {
		if byType, ok := m.addresses[chainSelector]; ok {
			if addr, ok := byType[contractType]; ok {
				return addr, nil
			}
		}
	}
	return nil, fmt.Errorf("address not found for chain %d type %s", chainSelector, contractType)
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

func ptrTo[T any](v T) *T { return &v }

func newMockAdapter(prefix string, chainAddresses map[uint64]map[string][]byte, executors map[uint64]map[string]string) *mockChainFamilyAdapter {
	return &mockChainFamilyAdapter{
		addressPrefix: prefix,
		addresses:     chainAddresses,
		executors:     executors,
	}
}

func TestConfigureChainsForLanesFromTopology_HappyPathAndCrossFamily(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteEVM := chainsel.TEST_90000002.Selector
	remoteSolana := chainsel.SOLANA_DEVNET.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
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

	localAdapter := newMockAdapter("local:", map[uint64]map[string][]byte{
		localSelector: {
			"Router":    {0xaa, 0x01},
			"OnRamp":    {0xbb, 0x02},
			"FeeQuoter": {0xcc, 0x03},
			"OffRamp":   {0xdd, 0x04},
		},
		remoteEVM: {
			"OnRamp":  {0xee, 0x11},
			"OffRamp": {0xee, 0x22},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
	remoteSolanaAdapter := newMockAdapter("sol:", map[uint64]map[string][]byte{
		remoteSolana: {
			"OnRamp":  {0xff, 0x11},
			"OffRamp": {0xff, 0x22},
		},
	}, nil)
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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
					RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
						remoteEVM:    {FeeUSDCents: ptrTo[uint16](10), GasForVerification: ptrTo[uint32](20), PayloadSizeBytes: ptrTo[uint16](30)},
						remoteSolana: {FeeUSDCents: ptrTo[uint16](40), GasForVerification: ptrTo[uint32](50), PayloadSizeBytes: ptrTo[uint16](60)},
					},
					},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteEVM: {
						DefaultExecutorQualifier: "default",
						DefaultInboundCCVs:       []datastore.AddressRef{testRef(localSelector, "0xverifier", "CommitteeVerifier")},
						LaneMandatedInboundCCVs:  []datastore.AddressRef{testRef(localSelector, "0xverifier", "CommitteeVerifier")},
						DefaultOutboundCCVs:      []datastore.AddressRef{testRef(localSelector, "0xverifier", "CommitteeVerifier")},
						LaneMandatedOutboundCCVs: []datastore.AddressRef{testRef(localSelector, "0xverifier", "CommitteeVerifier")},
					},
					remoteSolana: {
						DefaultExecutorQualifier: "default",
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
	assert.Equal(t, []byte{0xaa, 0x01}, input.Router)
	assert.Equal(t, []byte{0xbb, 0x02}, input.OnRamp)
	assert.Equal(t, []byte{0xcc, 0x03}, input.FeeQuoter)
	assert.Equal(t, []byte{0xdd, 0x04}, input.OffRamp)
	require.Len(t, input.RemoteChains, 2)
	assert.Equal(t, []byte{0xee, 0x22}, input.RemoteChains[remoteEVM].OffRamp)
	assert.Equal(t, []byte{0xff, 0x22}, input.RemoteChains[remoteSolana].OffRamp)
	assert.Equal(t, []byte{0xee, 0x11}, input.RemoteChains[remoteEVM].OnRamps[0])
	assert.Equal(t, []byte{0xff, 0x11}, input.RemoteChains[remoteSolana].OnRamps[0])
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
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
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

	localAdapter := newMockAdapter("local:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0x01}, "OnRamp": {0x02}, "FeeQuoter": {0x03}, "OffRamp": {0x04},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
	remoteAdapter := newMockAdapter("sol:", map[uint64]map[string][]byte{
		remoteSelector: {"OnRamp": {0x11}, "OffRamp": {0x22}},
	}, nil)
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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							remoteSelector: {},
						},
					},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {
						DefaultExecutorQualifier: "default",
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

	emptyAdapter := newMockAdapter("", nil, nil)
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, emptyAdapter)
	registry.RegisterChainFamily(chainsel.FamilySolana, emptyAdapter)

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
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
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

	localAdapter := newMockAdapter("", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0x01}, "OnRamp": {0x02}, "FeeQuoter": {0x03}, "OffRamp": {0x04},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, localAdapter)

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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {
						DefaultExecutorQualifier: "default",
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
		addAddress(t, ds, testRef(sel, fmt.Sprintf("0xverifier-%d", sel), "CommitteeVerifier"))
	}
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", chainA): {testRef(chainA, fmt.Sprintf("0xverifier-%d", chainA), "CommitteeVerifier")},
			fmt.Sprintf("%d:default", chainB): {testRef(chainB, fmt.Sprintf("0xverifier-%d", chainB), "CommitteeVerifier")},
		},
	})

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		chainA: {
			"Router": {0xa1}, "OnRamp": {0xa2}, "FeeQuoter": {0xa3}, "OffRamp": {0xa4},
		},
		chainB: {
			"Router": {0xb1}, "OnRamp": {0xb2}, "FeeQuoter": {0xb3}, "OffRamp": {0xb4},
		},
		sharedDest: {
			"OnRamp": {0xd1}, "OffRamp": {0xd2},
		},
	}, map[uint64]map[string]string{
		chainA: {"default": fmt.Sprintf("0xexecutor-%d", chainA)},
		chainB: {"default": fmt.Sprintf("0xexecutor-%d", chainB)},
	})
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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{sharedDest: {}}},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					sharedDest: {
						FeeQuoterDestChainConfig: changesets.FeeQuoterDestChainConfigOverrides{MaxDataBytes: ptrTo[uint32](1000)},
						ExecutorDestChainConfig:  &adapters.ExecutorDestChainConfig{USDCentsFee: 100, Enabled: true},
					},
				},
			},
			{
				ChainSelector: chainB,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{sharedDest: {}}},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					sharedDest: {
						FeeQuoterDestChainConfig: changesets.FeeQuoterDestChainConfigOverrides{MaxDataBytes: ptrTo[uint32](2000)},
						ExecutorDestChainConfig:  &adapters.ExecutorDestChainConfig{USDCentsFee: 200, Enabled: true},
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

func TestConfigureChainsForLanesFromTopology_UsesTestRouterWhenFlagIsSet(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {testRef(localSelector, "0xverifier", "CommitteeVerifier")},
		},
	})

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0xaa, 0x01}, "TestRouter": {0xbb, 0x99},
			"OnRamp": {0xaa, 0x02}, "FeeQuoter": {0xaa, 0x03}, "OffRamp": {0xaa, 0x04},
		},
		remoteSelector: {
			"OnRamp": {0xcc, 0x01}, "OffRamp": {0xcc, 0x02},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
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
		UseTestRouter: true,
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {
						DefaultExecutorQualifier: "default",
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, evmAdapter.inputs, 1)
	assert.Equal(t, []byte{0xbb, 0x99}, evmAdapter.inputs[0].Router)
	assert.True(t, evmAdapter.inputs[0].AllowOnrampOverride,
		"UseTestRouter=true should set AllowOnrampOverride=true on sequence input")
}

func TestConfigureChainsForLanesFromTopology_SelectsStandardRouterWhenBothExist(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {testRef(localSelector, "0xverifier", "CommitteeVerifier")},
		},
	})

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0xaa, 0x01}, "TestRouter": {0xbb, 0x99},
			"OnRamp": {0xaa, 0x02}, "FeeQuoter": {0xaa, 0x03}, "OffRamp": {0xaa, 0x04},
		},
		remoteSelector: {
			"OnRamp": {0xcc, 0x01}, "OffRamp": {0xcc, 0x02},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
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
		UseTestRouter: false,
		Chains: []changesets.PartialChainConfig{
			{
				ChainSelector: localSelector,
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {
						DefaultExecutorQualifier: "default",
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, evmAdapter.inputs, 1)
	assert.Equal(t, []byte{0xaa, 0x01}, evmAdapter.inputs[0].Router,
		"with UseTestRouter=false and both routers in the datastore, the standard Router should be selected")
	assert.False(t, evmAdapter.inputs[0].AllowOnrampOverride,
		"UseTestRouter=false should set AllowOnrampOverride=false on sequence input")
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
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
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

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0x01}, "OnRamp": {0x02}, "FeeQuoter": {0x03}, "OffRamp": {0x04},
		},
		remoteSelector: {
			"OnRamp": {0x11}, "OffRamp": {0x22},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							remoteSelector: {},
						},
					},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {
						DefaultExecutorQualifier: "default",
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
	remoteSelector := chainsel.TEST_90000002.Selector

	t.Run("missing topology", func(t *testing.T) {
		env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
		cs := changesets.ConfigureChainsForLanesFromTopology(
			adapters.NewCommitteeVerifierContractRegistry(),
			adapters.NewChainFamilyRegistry(),
			changesetscore.GetRegistry(),
		)
		err := cs.VerifyPreconditions(env, changesets.ConfigureChainsForLanesFromTopologyConfig{})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "topology is required")
	})

	t.Run("empty DefaultExecutorQualifier defaults to 'default'", func(t *testing.T) {
		env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
		ds := datastore.NewMemoryDataStore()
		addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
		env.DataStore = ds.Seal()

		committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
		committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
			contractsByChainAndQualifier: map[string][]datastore.AddressRef{
				fmt.Sprintf("%d:default", localSelector): {testRef(localSelector, "0xverifier", "CommitteeVerifier")},
			},
		})
		cs := changesets.ConfigureChainsForLanesFromTopology(
			committeeRegistry,
			adapters.NewChainFamilyRegistry(),
			changesetscore.GetRegistry(),
		)
		err := cs.VerifyPreconditions(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
			Topology: &offchain.EnvironmentTopology{
				NOPTopology: &offchain.NOPTopology{
					NOPs: []offchain.NOPConfig{{Alias: "nop-1"}},
					Committees: map[string]offchain.CommitteeConfig{
						"default": {Qualifier: "default", ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteSelector): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						}},
					},
				},
			},
			Chains: []changesets.PartialChainConfig{
				{
					ChainSelector: localSelector,
					RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
						remoteSelector: {},
					},
				},
			},
		})
		require.NoError(t, err)
	})

	t.Run("missing CCV adapter rejects early", func(t *testing.T) {
		env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
		cs := changesets.ConfigureChainsForLanesFromTopology(
			adapters.NewCommitteeVerifierContractRegistry(),
			adapters.NewChainFamilyRegistry(),
			changesetscore.GetRegistry(),
		)
		err := cs.VerifyPreconditions(env, changesets.ConfigureChainsForLanesFromTopologyConfig{
			Topology: &offchain.EnvironmentTopology{
				NOPTopology: &offchain.NOPTopology{
					NOPs: []offchain.NOPConfig{{Alias: "nop-1"}},
					Committees: map[string]offchain.CommitteeConfig{
						"default": {Qualifier: "default", ChainConfigs: map[string]offchain.ChainCommitteeConfig{
							fmt.Sprintf("%d", remoteSelector): {NOPAliases: []string{"nop-1"}, Threshold: 1},
						}},
					},
				},
			},
			Chains: []changesets.PartialChainConfig{
				{
					ChainSelector: localSelector,
					RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
						remoteSelector: {},
					},
				},
			},
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "auto-resolved CCVs")
	})
}

func TestConfigureChainsForLanesFromTopology_EmptyConfigUsesAdapterDefaults(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {testRef(localSelector, "0xverifier", "CommitteeVerifier")},
		},
	})

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0x01}, "OnRamp": {0x02}, "FeeQuoter": {0x03}, "OffRamp": {0x04},
		},
		remoteSelector: {
			"OnRamp": {0x11}, "OffRamp": {0x22},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, evmAdapter.inputs, 1)

	remote := evmAdapter.inputs[0].RemoteChains[remoteSelector]
	assert.Equal(t, "0xexecutor", evmAdapter.inputs[0].RemoteChains[remoteSelector].DefaultExecutor,
		"empty qualifier should resolve to 'default' executor")
	assert.NotNil(t, remote.AllowTrafficFrom)
	assert.True(t, *remote.AllowTrafficFrom, "adapter default AllowTrafficFrom is true")
	assert.True(t, remote.ExecutorDestChainConfig.Enabled, "adapter default ExecutorDestChainConfig.Enabled is true")
	assert.Equal(t, uint32(175_000), remote.BaseExecutionGasCost, "adapter default BaseExecutionGasCost")
	assert.NotNil(t, remote.TokenReceiverAllowed)
	assert.False(t, *remote.TokenReceiverAllowed, "adapter default TokenReceiverAllowed is false")
	assert.Equal(t, uint16(10), remote.MessageNetworkFeeUSDCents, "adapter default MessageNetworkFeeUSDCents")
	assert.Equal(t, uint16(25), remote.TokenNetworkFeeUSDCents, "adapter default TokenNetworkFeeUSDCents")
	assert.Equal(t, []string{"0xverifier"}, remote.DefaultInboundCCVs, "auto-resolved inbound CCVs from default qualifier")
	assert.Equal(t, []string{"0xverifier"}, remote.DefaultOutboundCCVs, "auto-resolved outbound CCVs from default qualifier")
}

func TestConfigureChainsForLanesFromTopology_PointerOverridesReplaceDefaults(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	env.DataStore = ds.Seal()

	committeeRegistry := adapters.NewCommitteeVerifierContractRegistry()
	committeeRegistry.Register(chainsel.FamilyEVM, &mockCommitteeVerifierContractAdapter{
		contractsByChainAndQualifier: map[string][]datastore.AddressRef{
			fmt.Sprintf("%d:default", localSelector): {testRef(localSelector, "0xverifier", "CommitteeVerifier")},
		},
	})

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0x01}, "OnRamp": {0x02}, "FeeQuoter": {0x03}, "OffRamp": {0x04},
		},
		remoteSelector: {
			"OnRamp": {0x11}, "OffRamp": {0x22},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{CommitteeQualifier: "default", RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{remoteSelector: {}}},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {
						AllowTrafficFrom:          ptrTo(false),
						BaseExecutionGasCost:      ptrTo[uint32](250_000),
						TokenReceiverAllowed:      ptrTo(true),
						MessageNetworkFeeUSDCents: ptrTo[uint16](50),
						TokenNetworkFeeUSDCents:   ptrTo[uint16](75),
						ExecutorDestChainConfig:   &adapters.ExecutorDestChainConfig{USDCentsFee: 500, Enabled: false},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, evmAdapter.inputs, 1)

	remote := evmAdapter.inputs[0].RemoteChains[remoteSelector]
	assert.NotNil(t, remote.AllowTrafficFrom)
	assert.False(t, *remote.AllowTrafficFrom, "override AllowTrafficFrom=false")
	assert.Equal(t, uint32(250_000), remote.BaseExecutionGasCost, "override BaseExecutionGasCost")
	assert.NotNil(t, remote.TokenReceiverAllowed)
	assert.True(t, *remote.TokenReceiverAllowed, "override TokenReceiverAllowed=true")
	assert.Equal(t, uint16(50), remote.MessageNetworkFeeUSDCents, "override MessageNetworkFeeUSDCents")
	assert.Equal(t, uint16(75), remote.TokenNetworkFeeUSDCents, "override TokenNetworkFeeUSDCents")
	assert.Equal(t, uint16(500), remote.ExecutorDestChainConfig.USDCentsFee, "override ExecutorDestChainConfig.USDCentsFee")
	assert.False(t, remote.ExecutorDestChainConfig.Enabled, "override ExecutorDestChainConfig.Enabled=false")
}

func TestConfigureChainsForLanesFromTopology_EmptyCVConfigUsesAdapterDefaults(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
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

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0x01}, "OnRamp": {0x02}, "FeeQuoter": {0x03}, "OffRamp": {0x04},
		},
		remoteSelector: {
			"OnRamp": {0x11}, "OffRamp": {0x22},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier: "default",
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							remoteSelector: {},
						},
					},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, evmAdapter.inputs, 1)

	cv := evmAdapter.inputs[0].CommitteeVerifiers[0]
	remoteCfg := cv.RemoteChains[remoteSelector]
	assert.False(t, remoteCfg.AllowlistEnabled, "adapter default AllowlistEnabled")
	assert.Equal(t, uint16(0), remoteCfg.FeeUSDCents, "adapter default FeeUSDCents")
	assert.Equal(t, uint32(60_000), remoteCfg.GasForVerification, "adapter default GasForVerification")
	assert.Equal(t, uint16(390), remoteCfg.PayloadSizeBytes, "adapter default PayloadSizeBytes")

	assert.True(t, cv.AllowedFinalityConfig.WaitForFinality, "adapter default WaitForFinality")
	assert.True(t, cv.AllowedFinalityConfig.WaitForSafe, "adapter default WaitForSafe")
	assert.Equal(t, uint16(1), cv.AllowedFinalityConfig.BlockDepth, "adapter default BlockDepth")
}

func TestConfigureChainsForLanesFromTopology_CVPointerOverridesReplaceDefaults(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	env := newConfigureChainsTestEnv(t, []uint64{localSelector}, nil)
	ds := datastore.NewMemoryDataStore()
	addAddress(t, ds, testRef(localSelector, "0xverifier", "CommitteeVerifier"))
	addAddress(t, ds, testRef(localSelector, "0xresolver", "CommitteeVerifierResolver"))
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

	evmAdapter := newMockAdapter("evm:", map[uint64]map[string][]byte{
		localSelector: {
			"Router": {0x01}, "OnRamp": {0x02}, "FeeQuoter": {0x03}, "OffRamp": {0x04},
		},
		remoteSelector: {
			"OnRamp": {0x11}, "OffRamp": {0x22},
		},
	}, map[uint64]map[string]string{
		localSelector: {"default": "0xexecutor"},
	})
	registry := adapters.NewChainFamilyRegistry()
	registry.RegisterChainFamily(chainsel.FamilyEVM, evmAdapter)

	customFinality := finality.Config{WaitForFinality: true}

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
				CommitteeVerifiers: []changesets.CommitteeVerifierInputConfig{
					{
						CommitteeQualifier:    "default",
						AllowedFinalityConfig: &customFinality,
						RemoteChains: map[uint64]changesets.CommitteeVerifierRemoteChainConfig{
							remoteSelector: {
								AllowlistEnabled:   ptrTo(true),
								FeeUSDCents:        ptrTo[uint16](99),
								GasForVerification: ptrTo[uint32](80_000),
								PayloadSizeBytes:   ptrTo[uint16](500),
							},
						},
					},
				},
				RemoteChains: map[uint64]changesets.PartialRemoteChainConfig{
					remoteSelector: {},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, evmAdapter.inputs, 1)

	cv := evmAdapter.inputs[0].CommitteeVerifiers[0]
	remoteCfg := cv.RemoteChains[remoteSelector]
	assert.True(t, remoteCfg.AllowlistEnabled, "override AllowlistEnabled=true")
	assert.Equal(t, uint16(99), remoteCfg.FeeUSDCents, "override FeeUSDCents")
	assert.Equal(t, uint32(80_000), remoteCfg.GasForVerification, "override GasForVerification")
	assert.Equal(t, uint16(500), remoteCfg.PayloadSizeBytes, "override PayloadSizeBytes")

	assert.Equal(t, customFinality, cv.AllowedFinalityConfig, "override AllowedFinalityConfig")
	assert.False(t, cv.AllowedFinalityConfig.WaitForSafe, "explicit finality should not have WaitForSafe")
	assert.Equal(t, uint16(0), cv.AllowedFinalityConfig.BlockDepth, "explicit finality should not have BlockDepth")
}
