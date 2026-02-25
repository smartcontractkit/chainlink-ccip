package fetch_node_chain_support

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	deploymocks "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/internal/mocks"
)

func newTestBundle(t *testing.T) operations.Bundle {
	t.Helper()
	lggr, err := logger.New()
	require.NoError(t, err)
	return operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)
}

func newTestLogger(t *testing.T) logger.Logger {
	t.Helper()
	lggr, err := logger.New()
	require.NoError(t, err)
	return lggr
}

func TestFetchNodeChainSupport_EmptyNOPAliases_ReturnsEmpty(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.NoError(t, err)
	assert.Empty(t, report.Output.SupportedChains)
}

func TestFetchNodeChainSupport_SingleNOP_EVMChain_ReturnsChainSelector(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}},
		}, nil,
	)
	mockClient.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "11155111"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "abcd1234",
						},
					},
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	require.Contains(t, report.Output.SupportedChains, "nop-1")
	require.Len(t, report.Output.SupportedChains["nop-1"], 1)
	assert.Equal(t, uint64(16015286601757825753), report.Output.SupportedChains["nop-1"][0])
}

func TestFetchNodeChainSupport_MultipleNOPs_MultipleChains_ReturnsAllSelectors(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop-1"},
				{Id: "node-2", Name: "nop-2"},
			},
		}, nil,
	)
	mockClient.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "11155111"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "key-1",
						},
					},
				},
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "421614"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "key-1",
						},
					},
				},
				{
					NodeId: "node-2",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "11155111"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "key-2",
						},
					},
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1", "nop-2"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)

	require.NoError(t, err)
	require.Len(t, report.Output.SupportedChains, 2)

	require.Contains(t, report.Output.SupportedChains, "nop-1")
	assert.Len(t, report.Output.SupportedChains["nop-1"], 2)

	require.Contains(t, report.Output.SupportedChains, "nop-2")
	assert.Len(t, report.Output.SupportedChains["nop-2"], 1)
}

func TestFetchNodeChainSupport_NOPNotFound_ContinuesWithOtherNOPs(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}},
		}, nil,
	)
	mockClient.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "11155111"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "key-1",
						},
					},
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1", "non-existent-nop"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	require.Len(t, report.Output.SupportedChains, 1)
	assert.Contains(t, report.Output.SupportedChains, "nop-1")
}

func TestFetchNodeChainSupport_ListNodesError_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		nil, fmt.Errorf("connection refused"),
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1"},
	}

	_, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list nodes")
}

func TestFetchNodeChainSupport_ListChainConfigsError_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}},
		}, nil,
	)
	mockClient.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		nil, fmt.Errorf("timeout"),
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1"},
	}

	_, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list chain configs")
}

func TestFetchNodeChainSupport_UnsupportedChainType_Skipped(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}},
		}, nil,
	)
	mockClient.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_UNSPECIFIED, Id: "123"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "some-key",
						},
					},
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	assert.Empty(t, report.Output.SupportedChains)
}

func TestFetchNodeChainSupport_InvalidChainId_Skipped(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}},
		}, nil,
	)
	mockClient.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "invalid-chain-id"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "some-key",
						},
					},
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	assert.Empty(t, report.Output.SupportedChains)
}

func TestFetchNodeChainSupport_AllNOPsNotFound_ReturnsEmpty(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "other-nop"}},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchNodeChainSupportInput{
		NOPAliases: []string{"nop-1", "nop-2"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNodeChainSupport, FetchNodeChainSupportDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	assert.Empty(t, report.Output.SupportedChains)
}
