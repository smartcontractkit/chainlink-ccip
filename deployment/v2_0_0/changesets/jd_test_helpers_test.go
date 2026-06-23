package changesets_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	jobpb "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/mocks"
)

type mockJDNode struct {
	nodeID   string
	nopAlias string
	chainIDs []string
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

// expectJDInteractions sets up the expected interactions with the mock Job Distributor client for the given nodes.
func expectJDInteractions(t *testing.T, mockJD *mocks.MockClient, nodes []mockJDNode, expectRevoke bool) []string {
	t.Helper()
	return expectJDInteractionsWithProposeJob(t, mockJD, nodes, expectRevoke, nil)
}

// expectJDInteractionsWithProposeJob sets up the expected interactions with the mock Job Distributor client for the given nodes, allowing for a custom function to be called on each ProposeJob request.
func expectJDInteractionsWithProposeJob(
	t *testing.T,
	mockJD *mocks.MockClient,
	nodes []mockJDNode,
	expectRevoke bool,
	onPropose func(*jobpb.ProposeJobRequest),
) []string {
	t.Helper()

	nodeIDs := make([]string, 0, len(nodes))
	jdNodes := make([]*nodev1.Node, 0, len(nodes))
	chainConfigs := make([]*nodev1.ChainConfig, 0)
	for _, node := range nodes {
		nodeIDs = append(nodeIDs, node.nodeID)
		jdNodes = append(jdNodes, &nodev1.Node{Id: node.nodeID, Name: node.nopAlias})
		for _, chainID := range node.chainIDs {
			chainConfigs = append(chainConfigs, chainConfig(node.nodeID, chainID))
		}
	}

	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: jdNodes}, nil,
	)
	if len(chainConfigs) > 0 {
		mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
			&nodev1.ListNodeChainConfigsResponse{ChainConfigs: chainConfigs}, nil,
		)
	}
	proposeExpectation := mockJD.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(_ context.Context, req *jobpb.ProposeJobRequest, _ ...grpc.CallOption) (*jobpb.ProposeJobResponse, error) {
			if onPropose != nil {
				onPropose(req)
			}
			return &jobpb.ProposeJobResponse{
				Proposal: &jobpb.Proposal{Id: "proposal-" + req.NodeId, JobId: "job-" + req.NodeId, Spec: req.Spec, Revision: 1},
			}, nil
		},
	)
	if onPropose == nil {
		proposeExpectation.Maybe()
	}
	if expectRevoke {
		mockJD.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(&jobpb.RevokeJobResponse{}, nil)
	}

	return nodeIDs
}
