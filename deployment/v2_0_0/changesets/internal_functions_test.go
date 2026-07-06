package changesets

// Internal tests for package-private helpers that cannot be reached from the
// external _test package (signerFromJDIfMissing, fetchSigningKeysForNOPsByFamilies,
// deriveFamiliesFromSelectors).

import (
	"context"
	"encoding/hex"
	"strings"
	"testing"

	gethcrypto "github.com/ethereum/go-ethereum/crypto"
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

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain/operations/fetch_signing_keys"
)

// internalStubOffchain satisfies offchain.Client for tests that only need a
// non-nil client and will never have their methods invoked.
type internalStubOffchain struct {
	jobv1.JobServiceClient
	nodev1.NodeServiceClient
	csav1.CSAServiceClient
}

var _ cldf_offchain.Client = (*internalStubOffchain)(nil)

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

// ---- fetchSigningKeysForNOPsByFamilies ----

func TestFetchSigningKeysForNOPsByFamilies_NilOffchain_ReturnsNil(t *testing.T) {
	e := deployment.Environment{}
	result, _, err := fetchSigningKeysForNOPsByFamilies(e, []offchain.NOPConfig{{Alias: "nop1"}}, []string{chainsel.FamilyEVM})
	require.NoError(t, err)
	assert.Nil(t, result)
}

func TestFetchSigningKeysForNOPsByFamilies_AllSignersPresent_ReturnsNil(t *testing.T) {
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
	result, _, err := fetchSigningKeysForNOPsByFamilies(e, nops, []string{chainsel.FamilyEVM})
	require.NoError(t, err)
	assert.Nil(t, result)
}

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

func TestFetchSigningKeysForNOPsByFamilies_CallsJD_WhenSignerMissing(t *testing.T) {
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
		{Alias: "nop1"},
	}

	result, _, err := fetchSigningKeysForNOPsByFamilies(e, nops, []string{chainsel.FamilyEVM})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, result, "nop1")
	assert.NotEmpty(t, result["nop1"][chainsel.FamilyEVM])
}

func TestFetchSigningKeysForNOPsByFamilies_OnlyFetchesForNOPsMissingSigner(t *testing.T) {
	lggr := logger.Test(t)
	bundle := cldf_ops.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		cldf_ops.NewMemoryReporter(),
	)

	mock := &jdMockOffchain{
		listNodesFn: func(_ context.Context, _ *nodev1.ListNodesRequest, _ ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
			return &nodev1.ListNodesResponse{
				Nodes: []*nodev1.Node{
					{Id: "node-2", Name: "nop2"},
				},
			}, nil
		},
		listNodeChainConfigsFn: func(_ context.Context, in *nodev1.ListNodeChainConfigsRequest, _ ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
			require.ElementsMatch(t, []string{"node-2"}, in.Filter.NodeIds)
			return &nodev1.ListNodeChainConfigsResponse{
				ChainConfigs: []*nodev1.ChainConfig{
					{
						NodeId: "node-2",
						Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM},
						Ocr2Config: &nodev1.OCR2Config{
							OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x2222"},
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
		NodeIDs:          []string{"node-1", "node-2"},
	}

	result, _, err := fetchSigningKeysForNOPsByFamilies(e, []offchain.NOPConfig{
		{Alias: "nop1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xexisting"}},
		{Alias: "nop2"},
	}, []string{chainsel.FamilyEVM})
	require.NoError(t, err)

	require.NotNil(t, result)
	assert.Equal(t, "0x2222", result["nop2"][chainsel.FamilyEVM])
	_, hasNop1 := result["nop1"]
	assert.False(t, hasNop1, "nop1 already has EVM signer, should not be fetched")
}

// ---- signerAddressForNOPAlias precedence ----

func TestSignerAddressForNOPAlias_TopologySignerTakesPrecedenceOverJD(t *testing.T) {
	topologySigner := "0xTOPOLOGY_SIGNER"
	jdSigner := "0xJD_SIGNER"

	lggr := logger.Test(t)
	e := deployment.Environment{Logger: lggr}

	topology := &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{
					Alias:                 "nop1",
					SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: topologySigner},
				},
			},
		},
	}

	jdKeys := fetch_signing_keys.SigningKeysByNOP{
		"nop1": {chainsel.FamilyEVM: jdSigner},
	}

	result, err := signerAddressForNOPAlias(e, topology, "nop1", chainsel.FamilyEVM, "test-committee", 1, jdKeys, "")
	require.NoError(t, err)
	assert.Equal(t, topologySigner, result, "topology signer must take precedence over JD-fetched key")
}

func TestSignerAddressForNOPAlias_FallsBackToJDWhenTopologyMissing(t *testing.T) {
	jdSigner := "0xJD_SIGNER"

	lggr := logger.Test(t)
	e := deployment.Environment{Logger: lggr}

	topology := &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop1"},
			},
		},
	}

	jdKeys := fetch_signing_keys.SigningKeysByNOP{
		"nop1": {chainsel.FamilyEVM: jdSigner},
	}

	result, err := signerAddressForNOPAlias(e, topology, "nop1", chainsel.FamilyEVM, "test-committee", 1, jdKeys, "")
	require.NoError(t, err)
	assert.Equal(t, jdSigner, result, "should fall back to JD key when topology signer is absent")
}

// ---- translateSignerAddress ----

func TestTranslateSignerAddress_EVMToSolana(t *testing.T) {
	// EVM and Solana addresses encode the identical 20 bytes; only string formatting differs.
	translated, err := translateSignerAddress(chainsel.FamilyEVM, "0xAbC1230000000000000000000000000000000000", chainsel.FamilySolana)
	require.NoError(t, err)
	assert.Equal(t, "abc1230000000000000000000000000000000000", translated)
}

func TestTranslateSignerAddress_SolanaToEVM(t *testing.T) {
	translated, err := translateSignerAddress(chainsel.FamilySolana, "abc1230000000000000000000000000000000000", chainsel.FamilyEVM)
	require.NoError(t, err)
	assert.Equal(t, "0xAbc1230000000000000000000000000000000000", translated)
}

func testRawPubKeyHex(t *testing.T) (hexKey string, evmAddress string) {
	t.Helper()
	privKey, err := gethcrypto.HexToECDSA(strings.Repeat("01", 32))
	require.NoError(t, err)
	pubKeyBytes := gethcrypto.FromECDSAPub(&privKey.PublicKey)
	return hex.EncodeToString(pubKeyBytes), gethcrypto.PubkeyToAddress(privKey.PublicKey).Hex()
}

func TestTranslateSignerAddress_RawPubKeyFamiliesAreInterchangeable(t *testing.T) {
	rawPubKey, _ := testRawPubKeyHex(t)
	translated, err := translateSignerAddress(chainsel.FamilyAptos, rawPubKey, chainsel.FamilyStellar)
	require.NoError(t, err)
	assert.Equal(t, rawPubKey, translated)

	translated, err = translateSignerAddress(chainsel.FamilyCanton, "0x"+rawPubKey, chainsel.FamilyAptos)
	require.NoError(t, err)
	assert.Equal(t, rawPubKey, translated, "0x prefix on a raw-pubkey-class value should be tolerated")
}

func TestTranslateSignerAddress_RawPubKeyToAddress(t *testing.T) {
	// The derived address must match go-ethereum's own derivation so the destination
	// contract can actually verify signatures produced by this key.
	rawPubKey, wantAddress := testRawPubKeyHex(t)
	translated, err := translateSignerAddress(chainsel.FamilyAptos, rawPubKey, chainsel.FamilyEVM)
	require.NoError(t, err)
	assert.Equal(t, wantAddress, translated)
}

func TestTranslateSignerAddress_AddressToRawPubKeyIsRejected(t *testing.T) {
	_, err := translateSignerAddress(chainsel.FamilyEVM, "0xAbC1230000000000000000000000000000000000", chainsel.FamilyCanton)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not reversible")
}

// ---- translateFromKnownFamily / signerAddressForNOPAlias cross-family fallback ----

func TestSignerAddressForNOPAlias_TranslatesFromSiblingFamily(t *testing.T) {
	// Reproduces the reported bug: a standalone Solana verifier NOP only has a
	// "solana" signer address in JD, but the lane's destination (local) chain is EVM.
	lggr := logger.Test(t)
	e := deployment.Environment{Logger: lggr}

	topology := &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "solana-verifier-1"},
			},
		},
	}

	jdKeys := fetch_signing_keys.SigningKeysByNOP{
		"solana-verifier-1": {chainsel.FamilySolana: "abc1230000000000000000000000000000000000"},
	}

	result, err := signerAddressForNOPAlias(e, topology, "solana-verifier-1", chainsel.FamilyEVM, "default", 12463857294658392847, jdKeys, "")
	require.NoError(t, err)
	assert.Equal(t, "0xAbc1230000000000000000000000000000000000", result)
}

func TestSignerAddressForNOPAlias_TranslationFailureIsActionable(t *testing.T) {
	// An EVM-only NOP with no raw public key on file cannot sign into a Canton-destination
	// lane: an EVM address can't be un-hashed back into a raw public key.
	lggr := logger.Test(t)
	e := deployment.Environment{Logger: lggr}

	topology := &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{
					Alias:                 "evm-only-verifier",
					SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xAbC1230000000000000000000000000000000000"},
				},
			},
		},
	}

	_, err := signerAddressForNOPAlias(e, topology, "evm-only-verifier", chainsel.FamilyCanton, "default", 1, nil, "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "missing signer_address for family canton")
	assert.Contains(t, err.Error(), "not reversible")
}

func TestSignerAddressForNOPAlias_RawPubKeyUnblocksAddressToRawPubKeyDirection(t *testing.T) {
	// Same setup as above, but the NOP's raw public key is now available — the previously
	// unsupported EVM-address -> Canton-raw-pubkey direction is no longer needed because
	// the raw key itself is on hand, so the lookup succeeds directly from it.
	rawPubKey, _ := testRawPubKeyHex(t)
	lggr := logger.Test(t)
	e := deployment.Environment{Logger: lggr}

	topology := &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{
					Alias:                 "evm-only-verifier",
					SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xAbC1230000000000000000000000000000000000"},
				},
			},
		},
	}

	result, err := signerAddressForNOPAlias(e, topology, "evm-only-verifier", chainsel.FamilyCanton, "default", 1, nil, rawPubKey)
	require.NoError(t, err)
	assert.Equal(t, rawPubKey, result)
}

func TestSignerAddressForNOPAlias_RawPubKeyFromJDUnblocksSolanaToCanton(t *testing.T) {
	// The scenario the user flagged: a solana-only NOP whose OnchainSigningAddress is
	// EVM-style (un-hashable) needs to sign into a canton-destination lane. This only
	// works because JD also carries the NOP's raw public key (RawPubKeyByNOP), pushed
	// alongside every family bootstrap.go declares for the node.
	rawPubKey, _ := testRawPubKeyHex(t)
	lggr := logger.Test(t)
	e := deployment.Environment{Logger: lggr}

	topology := &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "solana-verifier-1"},
			},
		},
	}

	jdKeys := fetch_signing_keys.SigningKeysByNOP{
		"solana-verifier-1": {chainsel.FamilySolana: "abc1230000000000000000000000000000000000"},
	}

	result, err := signerAddressForNOPAlias(e, topology, "solana-verifier-1", chainsel.FamilyCanton, "default", 1, jdKeys, rawPubKey)
	require.NoError(t, err)
	assert.Equal(t, rawPubKey, result)
}

// ---- deriveFamiliesFromSelectors ----

func TestDeriveFamiliesFromSelectors_DeduplicatesAndIgnoresInvalid(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	families := deriveFamiliesFromSelectors([]uint64{sel1, sel2, 0xDEAD})
	assert.ElementsMatch(t, []string{chainsel.FamilyEVM}, families)
}

func TestDeriveFamiliesFromSelectors_EmptyInput(t *testing.T) {
	families := deriveFamiliesFromSelectors(nil)
	assert.Empty(t, families)
}
