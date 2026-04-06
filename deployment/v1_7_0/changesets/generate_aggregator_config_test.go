package changesets_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

var _ adapters.AggregatorConfigAdapter = (*mockOffchainConfigAdapter)(nil)

type mockOffchainConfigAdapter struct {
	committeeStatesByChain map[uint64][]*adapters.CommitteeState
	scanErr                error
	verifierAddresses      map[uint64]string
	resolveErr             error
}

func (m *mockOffchainConfigAdapter) ScanCommitteeStates(_ context.Context, _ deployment.Environment, chainSelector uint64) ([]*adapters.CommitteeState, error) {
	if m.scanErr != nil {
		return nil, m.scanErr
	}
	return m.committeeStatesByChain[chainSelector], nil
}

func (m *mockOffchainConfigAdapter) ResolveVerifierAddress(_ datastore.DataStore, chainSelector uint64, _ string) (string, error) {
	if m.resolveErr != nil {
		return "", m.resolveErr
	}
	addr, ok := m.verifierAddresses[chainSelector]
	if !ok {
		return "", fmt.Errorf("no verifier address for chain %d", chainSelector)
	}
	return addr, nil
}

func newTestBlockChains(selectors []uint64) cldf_chain.BlockChains {
	chains := make(map[uint64]cldf_chain.BlockChain, len(selectors))
	for _, sel := range selectors {
		chains[sel] = cldfevm.Chain{Selector: sel}
	}
	return cldf_chain.NewBlockChains(chains)
}

func newAggregatorConfigRegistry() *adapters.AggregatorConfigRegistry {
	return &adapters.AggregatorConfigRegistry{}
}

func newTopologyForCommittee(qualifier string, selectors []uint64) *offchain.EnvironmentTopology {
	chainConfigs := make(map[string]offchain.ChainCommitteeConfig, len(selectors))
	for _, sel := range selectors {
		chainConfigs[strconv.FormatUint(sel, 10)] = offchain.ChainCommitteeConfig{}
	}

	return &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			Committees: map[string]offchain.CommitteeConfig{
				qualifier: {
					Qualifier:    qualifier,
					ChainConfigs: chainConfigs,
				},
			},
		},
	}
}

func TestGenerateAggregatorConfig_Validation(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	registry := newAggregatorConfigRegistry()
	cs := changesets.GenerateAggregatorConfig(registry)

	tests := []struct {
		name        string
		input       changesets.GenerateAggregatorConfigInput
		errContains string
	}{
		{
			name: "missing service identifier",
			input: changesets.GenerateAggregatorConfigInput{
				CommitteeQualifier: "default",
			},
			errContains: "service identifier is required",
		},
		{
			name: "missing committee qualifier",
			input: changesets.GenerateAggregatorConfigInput{
				ServiceIdentifier: "agg",
			},
			errContains: "committee qualifier is required",
		},
		{
			name: "missing topology",
			input: changesets.GenerateAggregatorConfigInput{
				ServiceIdentifier:  "agg",
				CommitteeQualifier: "default",
			},
			errContains: "topology is required",
		},
		{
			name: "missing nop topology",
			input: changesets.GenerateAggregatorConfigInput{
				ServiceIdentifier:  "agg",
				CommitteeQualifier: "default",
				Topology:           &offchain.EnvironmentTopology{},
			},
			errContains: "NOP topology is required",
		},
		{
			name: "missing committee in topology",
			input: changesets.GenerateAggregatorConfigInput{
				ServiceIdentifier:  "agg",
				CommitteeQualifier: "default",
				Topology: &offchain.EnvironmentTopology{
					NOPTopology: &offchain.NOPTopology{Committees: map[string]offchain.CommitteeConfig{}},
				},
			},
			errContains: `committee "default" not found in topology`,
		},
		{
			name: "topology chain missing from environment",
			input: changesets.GenerateAggregatorConfigInput{
				ServiceIdentifier:  "agg",
				CommitteeQualifier: "default",
				Topology:           newTopologyForCommittee("default", []uint64{sel1, 9999}),
			},
			errContains: `committee "default" references chain selector 9999 which is not available in environment`,
		},
	}

	env := deployment.Environment{
		BlockChains: newTestBlockChains([]uint64{sel1, sel2}),
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(env, tc.input)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.errContains)
		})
	}
}

func TestGenerateAggregatorConfig_BuildsCorrectConfig(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel3 := chainsel.TEST_90000003.Selector

	sel1Str := strconv.FormatUint(sel1, 10)
	sel2Str := strconv.FormatUint(sel2, 10)

	qualifier := "default"
	addr1 := "0x1111111111111111111111111111111111111111"
	addr2 := "0x2222222222222222222222222222222222222222"
	signer1 := "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	signer2 := "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

	mock := &mockOffchainConfigAdapter{
		committeeStatesByChain: map[uint64][]*adapters.CommitteeState{
			sel1: {
				{
					Qualifier:     qualifier,
					ChainSelector: sel1,
					Address:       addr1,
					SignatureConfigs: []adapters.SignatureConfig{
						{SourceChainSelector: sel2, Signers: []string{signer1, signer2}, Threshold: 2},
					},
				},
			},
			sel2: {
				{
					Qualifier:     qualifier,
					ChainSelector: sel2,
					Address:       addr2,
					SignatureConfigs: []adapters.SignatureConfig{
						{SourceChainSelector: sel1, Signers: []string{signer1}, Threshold: 1},
					},
				},
			},
		},
		verifierAddresses: map[uint64]string{
			sel1: addr1,
			sel2: addr2,
		},
	}

	registry := newAggregatorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{sel1, sel2, sel3}),
	}

	cs := changesets.GenerateAggregatorConfig(registry)
	output, err := cs.Apply(env, changesets.GenerateAggregatorConfigInput{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: qualifier,
		Topology:           newTopologyForCommittee(qualifier, []uint64{sel1, sel2}),
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	cfg, err := offchain.GetAggregatorConfig(output.DataStore.Seal(), "test-aggregator")
	require.NoError(t, err)

	assert.Len(t, cfg.QuorumConfigs, 2)

	qc1 := cfg.QuorumConfigs[sel1Str]
	require.NotNil(t, qc1)
	assert.Equal(t, addr1, qc1.SourceVerifierAddress)
	assert.Len(t, qc1.Signers, 1)
	assert.Equal(t, signer1, qc1.Signers[0].Address)
	assert.Equal(t, uint8(1), qc1.Threshold)

	qc2 := cfg.QuorumConfigs[sel2Str]
	require.NotNil(t, qc2)
	assert.Equal(t, addr2, qc2.SourceVerifierAddress)
	assert.Len(t, qc2.Signers, 2)
	assert.Equal(t, uint8(2), qc2.Threshold)

	assert.Len(t, cfg.DestinationVerifiers, 2)
	assert.Equal(t, addr1, cfg.DestinationVerifiers[sel1Str])
	assert.Equal(t, addr2, cfg.DestinationVerifiers[sel2Str])
}

func TestGenerateAggregatorConfig_ScanError(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockOffchainConfigAdapter{
		scanErr: fmt.Errorf("scan failed"),
	}

	registry := newAggregatorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	env := deployment.Environment{
		BlockChains: newTestBlockChains([]uint64{sel1}),
	}

	cs := changesets.GenerateAggregatorConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateAggregatorConfigInput{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: "default",
		Topology:           newTopologyForCommittee("default", []uint64{sel1}),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "scan failed")
}

func TestGenerateAggregatorConfig_ResolveAddressError(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	qualifier := "default"
	mock := &mockOffchainConfigAdapter{
		committeeStatesByChain: map[uint64][]*adapters.CommitteeState{
			sel1: {
				{
					Qualifier:     qualifier,
					ChainSelector: sel1,
					Address:       "0x1111111111111111111111111111111111111111",
					SignatureConfigs: []adapters.SignatureConfig{
						{SourceChainSelector: sel1, Signers: []string{"0xaaa"}, Threshold: 1},
					},
				},
			},
		},
		resolveErr: fmt.Errorf("address not found"),
	}

	registry := newAggregatorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	env := deployment.Environment{
		BlockChains: newTestBlockChains([]uint64{sel1}),
	}

	cs := changesets.GenerateAggregatorConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateAggregatorConfigInput{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: qualifier,
		Topology:           newTopologyForCommittee(qualifier, []uint64{sel1}),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "address not found")
}

func TestGenerateAggregatorConfig_CommitteeNotFound(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockOffchainConfigAdapter{
		committeeStatesByChain: map[uint64][]*adapters.CommitteeState{},
	}

	registry := newAggregatorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	env := deployment.Environment{
		BlockChains: newTestBlockChains([]uint64{sel1}),
	}

	cs := changesets.GenerateAggregatorConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateAggregatorConfigInput{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: "default",
		Topology:           newTopologyForCommittee("default", []uint64{sel1}),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), `committee "default" not found in deployed verifier state`)
}

func TestGenerateAggregatorConfig_DuplicateSourceChainAcrossDestinationsSucceeds(t *testing.T) {
	destChainA := chainsel.TEST_90000001.Selector
	destChainB := chainsel.TEST_90000002.Selector
	sourceChain := chainsel.TEST_90000003.Selector

	sourceChainStr := strconv.FormatUint(sourceChain, 10)
	qualifier := "default"
	verifierAddrA := "0x1111111111111111111111111111111111111111"
	verifierAddrB := "0x2222222222222222222222222222222222222222"
	verifierAddrSrc := "0x3333333333333333333333333333333333333333"
	signer1 := "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	signer2 := "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

	// Each destination chain's verifier reports the same source chain signature config.
	// This is the normal case: when a committee is deployed across multiple destination
	// chains, each destination's verifier stores the same set of source chain configs.
	mock := &mockOffchainConfigAdapter{
		committeeStatesByChain: map[uint64][]*adapters.CommitteeState{
			destChainA: {
				{
					Qualifier:     qualifier,
					ChainSelector: destChainA,
					Address:       verifierAddrA,
					SignatureConfigs: []adapters.SignatureConfig{
						{SourceChainSelector: sourceChain, Signers: []string{signer1, signer2}, Threshold: 2},
					},
				},
			},
			destChainB: {
				{
					Qualifier:     qualifier,
					ChainSelector: destChainB,
					Address:       verifierAddrB,
					SignatureConfigs: []adapters.SignatureConfig{
						{SourceChainSelector: sourceChain, Signers: []string{signer1, signer2}, Threshold: 2},
					},
				},
			},
		},
		verifierAddresses: map[uint64]string{
			destChainA:  verifierAddrA,
			destChainB:  verifierAddrB,
			sourceChain: verifierAddrSrc,
		},
	}

	registry := newAggregatorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{destChainA, destChainB, sourceChain}),
	}

	cs := changesets.GenerateAggregatorConfig(registry)
	output, err := cs.Apply(env, changesets.GenerateAggregatorConfigInput{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: qualifier,
		Topology:           newTopologyForCommittee(qualifier, []uint64{destChainA, destChainB, sourceChain}),
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	cfg, err := offchain.GetAggregatorConfig(output.DataStore.Seal(), "test-aggregator")
	require.NoError(t, err)

	qc := cfg.QuorumConfigs[sourceChainStr]
	require.NotNil(t, qc, "quorum config for source chain must exist")
	assert.Equal(t, verifierAddrSrc, qc.SourceVerifierAddress)
	assert.Len(t, qc.Signers, 2)
	assert.Equal(t, uint8(2), qc.Threshold)
}

func TestGenerateAggregatorConfig_DuplicateQualifierOnSameChainFails(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	qualifier := "default"

	mock := &mockOffchainConfigAdapter{
		committeeStatesByChain: map[uint64][]*adapters.CommitteeState{
			sel1: {
				{
					Qualifier:     qualifier,
					ChainSelector: sel1,
					Address:       "0x1111111111111111111111111111111111111111",
					SignatureConfigs: []adapters.SignatureConfig{
						{SourceChainSelector: sel1, Signers: []string{"0xaaa"}, Threshold: 1},
					},
				},
				{
					Qualifier:     qualifier,
					ChainSelector: sel1,
					Address:       "0x2222222222222222222222222222222222222222",
					SignatureConfigs: []adapters.SignatureConfig{
						{SourceChainSelector: sel1, Signers: []string{"0xbbb"}, Threshold: 1},
					},
				},
			},
		},
	}

	registry := newAggregatorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	env := deployment.Environment{
		BlockChains: newTestBlockChains([]uint64{sel1}),
	}

	cs := changesets.GenerateAggregatorConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateAggregatorConfigInput{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: qualifier,
		Topology:           newTopologyForCommittee(qualifier, []uint64{sel1}),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "multiple committee verifiers with qualifier")
}

func TestGenerateAggregatorConfig_MissingAdapterForFamily(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	registry := newAggregatorConfigRegistry()

	env := deployment.Environment{
		BlockChains: newTestBlockChains([]uint64{sel1}),
	}

	cs := changesets.GenerateAggregatorConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateAggregatorConfigInput{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: "default",
		Topology:           newTopologyForCommittee("default", []uint64{sel1}),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no offchain config adapter registered")
}
