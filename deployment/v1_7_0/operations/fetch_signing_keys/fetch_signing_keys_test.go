package fetch_signing_keys

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
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

func TestFetchNOPSigningKeys_EmptyNOPAliases_ReturnsEmpty(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)

	input := FetchSigningKeysInput{
		NOPAliases: []string{},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.NoError(t, err)
	assert.Empty(t, report.Output.SigningKeysByNOP)
}

func TestFetchNOPSigningKeys_SingleNOP_EVMSigningKey_ReturnsKey(t *testing.T) {
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
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
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

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	require.Contains(t, report.Output.SigningKeysByNOP, "nop-1")
	assert.Equal(t, "0xabcd1234", report.Output.SigningKeysByNOP["nop-1"][chainsel.FamilyEVM])
}

func TestFetchNOPSigningKeys_MultipleNOPs_DifferentChains_ReturnsAllKeys(t *testing.T) {
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
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "evm-key-1",
						},
					},
				},
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_SOLANA},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "solana-key-1",
						},
					},
				},
				{
					NodeId: "node-2",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "evm-key-2",
						},
					},
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1", "nop-2"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)

	require.NoError(t, err)
	require.Len(t, report.Output.SigningKeysByNOP, 2)

	assert.Equal(t, "0xevm-key-1", report.Output.SigningKeysByNOP["nop-1"][chainsel.FamilyEVM])
	assert.Equal(t, "0xsolana-key-1", report.Output.SigningKeysByNOP["nop-1"][chainsel.FamilySolana])
	assert.Equal(t, "0xevm-key-2", report.Output.SigningKeysByNOP["nop-2"][chainsel.FamilyEVM])
}

func TestFetchNOPSigningKeys_NOPNotFound_ContinuesWithOtherNOPs(t *testing.T) {
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
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
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

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1", "non-existent-nop"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	require.Len(t, report.Output.SigningKeysByNOP, 1)
	assert.Contains(t, report.Output.SigningKeysByNOP, "nop-1")
}

func TestFetchNOPSigningKeys_ListNodesError_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		nil, fmt.Errorf("connection refused"),
	)

	bundle := newTestBundle(t)

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1"},
	}

	_, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list nodes")
}

func TestFetchNOPSigningKeys_ListChainConfigsError_ReturnsError(t *testing.T) {
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

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1"},
	}

	_, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to list chain configs")
}

func TestFetchNOPSigningKeys_NilOCR2Config_Skipped(t *testing.T) {
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
					NodeId:     "node-1",
					Chain:      &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
					Ocr2Config: nil,
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	assert.Empty(t, report.Output.SigningKeysByNOP)
}

func TestFetchNOPSigningKeys_EmptySigningAddress_Skipped(t *testing.T) {
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
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
							OnchainSigningAddress: "",
						},
					},
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	assert.Empty(t, report.Output.SigningKeysByNOP)
}

func TestFetchNOPSigningKeys_UnsupportedChainType_Skipped(t *testing.T) {
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
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_UNSPECIFIED},
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

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	assert.Empty(t, report.Output.SigningKeysByNOP)
}

func TestFetchNOPSigningKeys_AllNOPsNotFound_ReturnsEmpty(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{{Id: "node-1", Name: "other-nop"}},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := FetchSigningKeysInput{
		NOPAliases: []string{"nop-1", "nop-2"},
	}

	report, err := operations.ExecuteOperation(bundle, FetchNOPSigningKeys, FetchSigningKeysDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)

	require.NoError(t, err)
	assert.Empty(t, report.Output.SigningKeysByNOP)
}
