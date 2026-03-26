package changesets

// Internal tests for package-private helpers that cannot be reached from the
// external _test package (convertTopologyMonitoring, mustDecodeHex,
// signerFromJDIfMissing, fetchSigningKeysForNOPs).

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_offchain "github.com/smartcontractkit/chainlink-deployments-framework/offchain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	csav1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/csa"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/operations/fetch_signing_keys"
)

// internalStubOffchain satisfies offchain.Client for tests that only need a
// non-nil client and will never have their methods invoked.
type internalStubOffchain struct {
	jobv1.JobServiceClient
	nodev1.NodeServiceClient
	csav1.CSAServiceClient
}

var _ cldf_offchain.Client = (*internalStubOffchain)(nil)

// ---- convertTopologyMonitoring ----

func TestConvertTopologyMonitoring_NilReturnsEmpty(t *testing.T) {
	result := convertTopologyMonitoring(nil)
	assert.False(t, result.Enabled)
	assert.Empty(t, result.Type)
}

// ---- mustDecodeHex ----

func TestMustDecodeHex_ValidHex(t *testing.T) {
	b := mustDecodeHex("deadbeef")
	require.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, b)
}

func TestMustDecodeHex_InvalidHexPanics(t *testing.T) {
	require.Panics(t, func() { mustDecodeHex("ZZZZ") })
}

// ---- signerFromJDIfMissing ----

func TestSignerFromJDIfMissing_AlreadyPresent(t *testing.T) {
	signer, ok := signerFromJDIfMissing(
		map[string]string{chainsel.FamilyEVM: "0xabc"},
		"nop1", chainsel.FamilyEVM, nil,
	)
	assert.Empty(t, signer)
	assert.False(t, ok)
}

func TestSignerFromJDIfMissing_NilSignersNilKeys(t *testing.T) {
	signer, ok := signerFromJDIfMissing(nil, "nop1", chainsel.FamilyEVM, nil)
	assert.Empty(t, signer)
	assert.False(t, ok)
}

func TestSignerFromJDIfMissing_FoundInJD(t *testing.T) {
	keys := fetch_signing_keys.SigningKeysByNOP{
		"nop1": {chainsel.FamilyEVM: "0xjd-signer"},
	}
	signer, ok := signerFromJDIfMissing(nil, "nop1", chainsel.FamilyEVM, keys)
	assert.Equal(t, "0xjd-signer", signer)
	assert.True(t, ok)
}

func TestSignerFromJDIfMissing_NotFoundInJD(t *testing.T) {
	keys := fetch_signing_keys.SigningKeysByNOP{
		"nop1": {},
	}
	signer, ok := signerFromJDIfMissing(nil, "nop1", chainsel.FamilyEVM, keys)
	assert.Empty(t, signer)
	assert.False(t, ok)
}

// ---- fetchSigningKeysForNOPs ----

func TestFetchSigningKeysForNOPs_NilOffchain_ReturnsNil(t *testing.T) {
	e := deployment.Environment{}
	result := fetchSigningKeysForNOPs(e, []offchain.NOPConfig{{Alias: "nop1"}})
	assert.Nil(t, result)
}

func TestFetchSigningKeysForNOPs_AllSignersPresent_ReturnsNil(t *testing.T) {
	// All NOPs already have an EVM signer → aliases list is empty → return nil.
	lggr := logger.Test(t)
	bundle := cldf_ops.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		cldf_ops.NewMemoryReporter(),
	)
	e := deployment.Environment{
		Logger:           lggr,
		OperationsBundle: bundle,
		Offchain:         &internalStubOffchain{},
	}
	nops := []offchain.NOPConfig{
		{Alias: "nop1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xabc"}},
		{Alias: "nop2", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xdef"}},
	}
	result := fetchSigningKeysForNOPs(e, nops)
	assert.Nil(t, result)
}

// jdMockOffchain overrides ListNodes and ListNodeChainConfigs so the JD code path
// in fetchSigningKeysForNOPs can be exercised without a real JD server.
type jdMockOffchain struct {
	internalStubOffchain
	listNodesFn            func(context.Context, *nodev1.ListNodesRequest, ...grpc.CallOption) (*nodev1.ListNodesResponse, error)
	listNodeChainConfigsFn func(context.Context, *nodev1.ListNodeChainConfigsRequest, ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error)
}

func (m *jdMockOffchain) ListNodes(ctx context.Context, in *nodev1.ListNodesRequest, opts ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
	return m.listNodesFn(ctx, in, opts...)
}

func (m *jdMockOffchain) ListNodeChainConfigs(ctx context.Context, in *nodev1.ListNodeChainConfigsRequest, opts ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
	return m.listNodeChainConfigsFn(ctx, in, opts...)
}

func TestFetchSigningKeysForNOPs_CallsJD_WhenSignerMissing(t *testing.T) {
	// NOP without an EVM signer triggers the JD lookup path.
	lggr := logger.Test(t)
	bundle := cldf_ops.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		cldf_ops.NewMemoryReporter(),
	)

	mock := &jdMockOffchain{
		listNodesFn: func(_ context.Context, _ *nodev1.ListNodesRequest, _ ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
			return &nodev1.ListNodesResponse{
				Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop1"}},
			}, nil
		},
		listNodeChainConfigsFn: func(_ context.Context, _ *nodev1.ListNodeChainConfigsRequest, _ ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
			return &nodev1.ListNodeChainConfigsResponse{
				ChainConfigs: []*nodev1.ChainConfig{
					{
						NodeId: "node-1",
						Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
						Ocr2Config: &nodev1.OCR2Config{
							OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
								OnchainSigningAddress: "0xdeadbeef",
							},
						},
					},
				},
			}, nil
		},
	}

	e := deployment.Environment{
		Logger:           lggr,
		OperationsBundle: bundle,
		Offchain:         mock,
		NodeIDs:          []string{"node-1"},
	}

	nops := []offchain.NOPConfig{
		{Alias: "nop1"}, // no EVM signer → triggers JD fetch
	}

	result := fetchSigningKeysForNOPs(e, nops)
	require.NotNil(t, result)
	require.Contains(t, result, "nop1")
	assert.NotEmpty(t, result["nop1"][chainsel.FamilyEVM])
}
