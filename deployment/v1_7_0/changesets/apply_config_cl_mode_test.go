package changesets_test

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	chainsel "github.com/smartcontractkit/chain-selectors"

	jobpb "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/mocks"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

type proposeJobCapture struct {
	mu       sync.Mutex
	requests []*jobpb.ProposeJobRequest
}

func (c *proposeJobCapture) capture(req *jobpb.ProposeJobRequest) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.requests = append(c.requests, req)
}

func (c *proposeJobCapture) get() []*jobpb.ProposeJobRequest {
	c.mu.Lock()
	defer c.mu.Unlock()
	copied := make([]*jobpb.ProposeJobRequest, len(c.requests))
	copy(copied, c.requests)
	return copied
}

func (c *proposeJobCapture) findByNodeID(nodeID string) *jobpb.ProposeJobRequest {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, r := range c.requests {
		if r.NodeId == nodeID {
			return r
		}
	}
	return nil
}

func chainConfig(nodeID, chainID string) *nodev1.ChainConfig {
	return &nodev1.ChainConfig{
		NodeId: nodeID,
		Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: chainID},
		Ocr2Config: &nodev1.OCR2Config{
			OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x123"},
		},
	}
}

func TestApplyVerifierConfig_FailsWhenNOPMissingChainSupport(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1, sel2}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-2", "90000001"),
				chainConfig("node-2", "90000002"),
			},
		}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "chain support validation failed")
	assert.Contains(t, err.Error(), "nop1")
}

func TestApplyVerifierConfig_PassesWhenNOPSupportsExtraChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1, sel2}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-1", "90000002"),
				chainConfig("node-1", "90000003"),
				chainConfig("node-2", "90000001"),
				chainConfig("node-2", "90000002"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "job-1"}}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyVerifierConfig_SkipsChainSupportValidationForStandaloneNOPs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

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

func TestApplyVerifierConfig_PassesWhenNonTargetNOPMissingChainSupport(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}

	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1, sel2}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-1", "90000002"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "job-1"}}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

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

func TestApplyExecutorConfig_FailsWhenNOPMissingChainSupport(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	executorAdapter := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
			sel2: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, executorAdapter)

	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)
	topo := newTopologyWithChainConfigs(
		[]string{"nop1", "nop2"},
		"pool1",
		shared.NOPModeCL,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
			sel2Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
		},
	)
	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-2", "90000001"),
				chainConfig("node-2", "90000002"),
			},
		}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyExecutorConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "chain support validation failed")
	assert.Contains(t, err.Error(), "nop1")
}

func TestApplyExecutorConfig_PassesWhenNOPSupportsExtraChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	executorAdapter := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
			sel2: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, executorAdapter)

	topo := newMinimalTopology([]string{"nop1", "nop2"}, "pool1", shared.NOPModeCL)
	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-1", "90000002"),
				chainConfig("node-1", "90000003"),
				chainConfig("node-2", "90000001"),
				chainConfig("node-2", "90000002"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "job-1"}}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyExecutorConfig_SkipsChainSupportValidationForStandaloneNOPs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	executorAdapter := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, executorAdapter)

	topo := newMinimalTopology([]string{"nop1", "nop2"}, "pool1", shared.NOPModeStandalone)
	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyExecutorConfig_PassesWhenNonTargetNOPMissingChainSupport(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	executorAdapter := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
			sel2: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, executorAdapter)

	topo := newMinimalTopology([]string{"nop1", "nop2"}, "pool1", shared.NOPModeCL)
	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-1", "90000002"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "job-1"}}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
		TargetNOPs:        []shared.NOPAlias{"nop1"},
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyVerifierConfig_ProposesCorrectSpecToCorrectNode(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1})

	captured := &proposeJobCapture{}

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-2", "90000001"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(_ context.Context, req *jobpb.ProposeJobRequest, _ ...grpc.CallOption) (*jobpb.ProposeJobResponse, error) {
			captured.capture(req)
			return &jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "prop-" + req.NodeId}}, nil
		},
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	requests := captured.get()
	require.Len(t, requests, 2, "expected exactly 2 ProposeJob calls for 2 NOPs")

	for _, req := range requests {
		assert.Contains(t, []string{"node-1", "node-2"}, req.NodeId, "proposed to unexpected node")
		assert.Contains(t, req.Spec, `type = "ccvcommitteeverifier"`, "spec should be a verifier job")
		assert.Contains(t, req.Spec, "0xCommitteeVerifier", "spec should contain committee verifier address")
	}

	nop1Req := captured.findByNodeID("node-1")
	require.NotNil(t, nop1Req, "nop1 should have a proposal on node-1")
	assert.Contains(t, nop1Req.Spec, "nop1-", "nop1's spec should reference nop1 in the job name")
	assert.NotContains(t, nop1Req.Spec, "nop2-", "nop1's spec must not contain nop2's job name")

	nop2Req := captured.findByNodeID("node-2")
	require.NotNil(t, nop2Req, "nop2 should have a proposal on node-2")
	assert.Contains(t, nop2Req.Spec, "nop2-", "nop2's spec should reference nop2 in the job name")
	assert.NotContains(t, nop2Req.Spec, "nop1-", "nop2's spec must not contain nop1's job name")
}

func TestApplyExecutorConfig_ProposesCorrectSpecToCorrectNode(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	executorAdapter := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, executorAdapter)

	topo := newMinimalTopology([]string{"nop1", "nop2"}, "pool1", shared.NOPModeCL)
	env := newTestExecutorEnv(t, []uint64{sel1})

	captured := &proposeJobCapture{}

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-2", "90000001"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(_ context.Context, req *jobpb.ProposeJobRequest, _ ...grpc.CallOption) (*jobpb.ProposeJobResponse, error) {
			captured.capture(req)
			return &jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "prop-" + req.NodeId}}, nil
		},
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyExecutorConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	requests := captured.get()
	require.Len(t, requests, 2, "expected exactly 2 ProposeJob calls for 2 NOPs")

	for _, req := range requests {
		assert.Contains(t, []string{"node-1", "node-2"}, req.NodeId, "proposed to unexpected node")
		assert.Contains(t, req.Spec, `type = "ccvexecutor"`, "spec should be an executor job")
		assert.Contains(t, req.Spec, "0xOffRamp", "spec should contain offramp address")
	}

	nop1Req := captured.findByNodeID("node-1")
	require.NotNil(t, nop1Req, "nop1 should have a proposal on node-1")
	assert.Contains(t, nop1Req.Spec, "nop1-pool1-executor", "nop1's spec should have correct job name")
	assert.NotContains(t, nop1Req.Spec, "nop2-pool1-executor", "nop1's spec must not contain nop2's job name")

	nop2Req := captured.findByNodeID("node-2")
	require.NotNil(t, nop2Req, "nop2 should have a proposal on node-2")
	assert.Contains(t, nop2Req.Spec, "nop2-pool1-executor", "nop2's spec should have correct job name")
	assert.NotContains(t, nop2Req.Spec, "nop1-pool1-executor", "nop2's spec must not contain nop1's job name")
}

func TestApplyVerifierConfig_ExtraNodeIDsDoNotReceiveProposals(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1})

	captured := &proposeJobCapture{}

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.MatchedBy(func(req *nodev1.ListNodesRequest) bool {
		return len(req.Filter.Ids) == 3
	})).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
				{Id: "node-extra", Name: "extra-nop"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
				chainConfig("node-2", "90000001"),
				chainConfig("node-extra", "90000001"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(_ context.Context, req *jobpb.ProposeJobRequest, _ ...grpc.CallOption) (*jobpb.ProposeJobResponse, error) {
			captured.capture(req)
			return &jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "prop-1"}}, nil
		},
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2", "node-extra"}

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	requests := captured.get()
	require.Len(t, requests, 1, "only 1 NOP in topology, so exactly 1 proposal expected")
	assert.Equal(t, "node-1", requests[0].NodeId, "proposal should target node-1 (nop1), not extra nodes")
}

func TestApplyVerifierConfig_FailsWhenTopologyNOPNotInNodeIDs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1})

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
			},
		}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1"}

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nop2", "error should reference the NOP not covered by NodeIDs")
}

func TestApplyVerifierConfig_ListNodesFilteredByNodeIDs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1})

	var capturedListNodesReq *nodev1.ListNodesRequest

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).RunAndReturn(
		func(_ context.Context, req *nodev1.ListNodesRequest, _ ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
			capturedListNodesReq = req
			return &nodev1.ListNodesResponse{
				Nodes: []*nodev1.Node{
					{Id: "node-1", Name: "nop1"},
				},
			}, nil
		},
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "job-1"}}, nil,
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	require.NotNil(t, capturedListNodesReq)
	require.NotNil(t, capturedListNodesReq.Filter)
	assert.ElementsMatch(t, []string{"node-1", "node-2"}, capturedListNodesReq.Filter.Ids,
		"ListNodes should be filtered to exactly env.NodeIDs")
}

func TestApplyVerifierConfig_TargetNOPsScopesProposals(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1", "nop2", "nop3"}, "c1", []uint64{sel1}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1})

	captured := &proposeJobCapture{}

	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop1"},
				{Id: "node-2", Name: "nop2"},
				{Id: "node-3", Name: "nop3"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				chainConfig("node-1", "90000001"),
			},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(_ context.Context, req *jobpb.ProposeJobRequest, _ ...grpc.CallOption) (*jobpb.ProposeJobResponse, error) {
			captured.capture(req)
			return &jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "prop-" + req.NodeId}}, nil
		},
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2", "node-3"}

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		TargetNOPs:               []shared.NOPAlias{"nop1"},
	})
	require.NoError(t, err)

	requests := captured.get()
	require.Len(t, requests, 1, "only nop1 is targeted, so exactly 1 proposal expected")
	assert.Equal(t, "node-1", requests[0].NodeId, "proposal should target node-1 (nop1) only")

	specLower := strings.ToLower(requests[0].Spec)
	assert.NotContains(t, specLower, "nop2", "nop2's data should not appear in nop1's spec")
	assert.NotContains(t, specLower, "nop3", "nop3's data should not appear in nop1's spec")
}

func TestApplyVerifierConfig_RevokeOrphanedJobsDoesNotAffectOtherCommittees(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	ds := datastore.NewMemoryDataStore()
	executorJob := shared.JobInfo{
		JobID:    shared.JobID("nop1-pool1-executor"),
		NOPAlias: shared.NOPAlias("nop1"),
		Proposals: map[string]shared.ProposalRevision{
			"p1": {Status: shared.JobProposalStatusApproved},
		},
	}
	c2VerifierJob := shared.JobInfo{
		JobID:    shared.JobID("nop1-agg-c2-verifier"),
		NOPAlias: shared.NOPAlias("nop1"),
		Proposals: map[string]shared.ProposalRevision{
			"p2": {Status: shared.JobProposalStatusApproved},
		},
	}
	require.NoError(t, offchain.SaveJobs(ds, []shared.JobInfo{executorJob, c2VerifierJob}))

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1})
	env.DataStore = ds.Seal()
	env.Offchain = mocks.NewMockClient(t)
	env.NodeIDs = []string{"node-1"}
	env.Offchain.(*mocks.MockClient).EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop1"}}}, nil,
	)
	env.Offchain.(*mocks.MockClient).EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{chainConfig("node-1", "90000001")},
		}, nil,
	)
	env.Offchain.(*mocks.MockClient).EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "prop-1"}}, nil,
	)

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		RevokeOrphanedJobs:       true,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	allJobs, err := offchain.GetAllJobs(output.DataStore.Seal())
	require.NoError(t, err)
	_, hasExecutor := allJobs[shared.NOPAlias("nop1")][executorJob.JobID]
	assert.True(t, hasExecutor, "executor pool job must remain after verifier c1 revoke orphaned")
	nopJobs := allJobs[shared.NOPAlias("nop1")]
	require.NotNil(t, nopJobs)
	_, hasC2 := nopJobs[c2VerifierJob.JobID]
	assert.True(t, hasC2, "other committee (c2) verifier job must remain after verifier c1 revoke orphaned")
}

func TestApplyExecutorConfig_RevokeOrphanedJobsDoesNotAffectVerifierJobs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	executorAdapter := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{"pool1": {sel1}},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}
	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, executorAdapter)

	ds := datastore.NewMemoryDataStore()
	verifierJob := shared.JobInfo{
		JobID:    shared.JobID("nop1-agg-c1-verifier"),
		NOPAlias: shared.NOPAlias("nop1"),
		Proposals: map[string]shared.ProposalRevision{
			"p1": {Status: shared.JobProposalStatusApproved},
		},
	}
	require.NoError(t, offchain.SaveJobs(ds, []shared.JobInfo{verifierJob}))

	topo := newMinimalTopology([]string{"nop1"}, "pool1", shared.NOPModeCL)
	env := newTestExecutorEnv(t, []uint64{sel1})
	env.DataStore = ds.Seal()
	env.Offchain = mocks.NewMockClient(t)
	env.NodeIDs = []string{"node-1"}
	env.Offchain.(*mocks.MockClient).EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop1"}}}, nil,
	)
	env.Offchain.(*mocks.MockClient).EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{chainConfig("node-1", "90000001")},
		}, nil,
	)
	env.Offchain.(*mocks.MockClient).EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "prop-1"}}, nil,
	)

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:           topo,
		ExecutorQualifier:  "pool1",
		RevokeOrphanedJobs: true,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	allJobs, err := offchain.GetAllJobs(output.DataStore.Seal())
	require.NoError(t, err)
	nopJobs := allJobs[shared.NOPAlias("nop1")]
	require.NotNil(t, nopJobs)
	_, hasVerifier := nopJobs[verifierJob.JobID]
	assert.True(t, hasVerifier, "verifier job must remain after executor revoke orphaned")
}

func TestApplyVerifierConfig_WrapperPassesTargetNOPsAndRevokeOrphanedJobsToSequence(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	verifierAdapter := &mockVerifierJobConfigAdapter{
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
	registry.Register(chainsel.FamilyEVM, verifierAdapter)

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1}, shared.NOPModeCL)
	env := newVerifierTestEnv(t, []uint64{sel1})
	captured := &proposeJobCapture{}
	mockJD := mocks.NewMockClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop1"}, {Id: "node-2", Name: "nop2"}},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{chainConfig("node-1", "90000001"), chainConfig("node-2", "90000001")},
		}, nil,
	)
	mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(_ context.Context, req *jobpb.ProposeJobRequest, _ ...grpc.CallOption) (*jobpb.ProposeJobResponse, error) {
			captured.capture(req)
			return &jobpb.ProposeJobResponse{Proposal: &jobpb.Proposal{Id: "prop-" + req.NodeId}}, nil
		},
	)
	env.Offchain = mockJD
	env.NodeIDs = []string{"node-1", "node-2"}

	cs := changesets.ApplyVerifierConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		TargetNOPs:               []shared.NOPAlias{"nop1"},
		RevokeOrphanedJobs:       true,
	})
	require.NoError(t, err)

	requests := captured.get()
	require.Len(t, requests, 1, "wrapper must pass TargetNOPs to ManageJobProposals; only nop1 targeted so exactly 1 proposal")
	assert.Equal(t, "node-1", requests[0].NodeId, "proposal must be for nop1 (node-1) only")
}
