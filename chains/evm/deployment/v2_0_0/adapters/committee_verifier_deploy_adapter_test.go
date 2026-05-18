package adapters_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	rmn_proxy "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

const (
	testChainSelector  = uint64(5009297550715157269)
	testFeeAggregator  = "0x000000000000000000000000000000000000FEED"
	testAllowlistAdmin = "0x000000000000000000000000000000000000ADD1"
	testRMNProxyAddr   = "0x1111111111111111111111111111111111111111"
)

func rmnProxyRef(selector uint64) datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType(rmn_proxy.ContractType),
		Version:       rmn_proxy.Version,
		Address:       testRMNProxyAddr,
	}
}

func baseInput(selector uint64) ccvdeploymentadapters.DeployCommitteeVerifierInput {
	return ccvdeploymentadapters.DeployCommitteeVerifierInput{
		ChainSelector:     selector,
		DeployerContract:  "0x0000000000000000000000000000000000000FAC",
		ExistingAddresses: []datastore.AddressRef{rmnProxyRef(selector)},
		Params: ccvdeploymentadapters.CommitteeVerifierDeployParams{
			Version:       committee_verifier.Version,
			FeeAggregator: testFeeAggregator,
			Qualifier:     "alpha",
		},
	}
}

// TestEVMCommitteeVerifierDeployAdapter_Validation exercises every error
// branch in the chain-agnostic-input → EVM-input conversion that runs before
// the underlying DeployCommitteeVerifier sequence executes. A simulated
// environment is required because the sequence type-checks the dependency
// against the registered EVMChain even for short-circuited error paths.
func TestEVMCommitteeVerifierDeployAdapter_Validation(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	adapter := &adapters.EVMCommitteeVerifierDeployAdapter{}

	tests := []struct {
		name       string
		mutate     func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput)
		wantErrSub string
	}{
		{
			name:       "chain not in BlockChains",
			mutate:     func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) { in.ChainSelector = 1 },
			wantErrSub: "EVM chain not found",
		},
		{
			name: "DeployerContract empty",
			mutate: func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) {
				in.DeployerContract = ""
			},
			wantErrSub: "DeployerContract is required",
		},
		{
			name: "DeployerContract not hex",
			mutate: func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) {
				in.DeployerContract = "not-a-hex"
			},
			wantErrSub: "is not a valid hex address",
		},
		{
			name: "RMNProxy missing from ExistingAddresses",
			mutate: func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) {
				in.ExistingAddresses = nil
			},
			wantErrSub: "RMNProxy",
		},
		{
			name: "FeeAggregator empty",
			mutate: func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) {
				in.Params.FeeAggregator = ""
			},
			wantErrSub: "FeeAggregator is required",
		},
		{
			name: "FeeAggregator zero address",
			mutate: func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) {
				in.Params.FeeAggregator = "0x0000000000000000000000000000000000000000"
			},
			wantErrSub: "cannot be zero address",
		},
		{
			name: "FeeAggregator malformed",
			mutate: func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) {
				in.Params.FeeAggregator = "not-a-hex"
			},
			wantErrSub: "is not a valid hex address",
		},
		{
			name: "AllowlistAdmin malformed",
			mutate: func(in *ccvdeploymentadapters.DeployCommitteeVerifierInput) {
				in.Params.AllowlistAdmin = "not-a-hex"
			},
			wantErrSub: "is not a valid hex address",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := baseInput(testChainSelector)
			tc.mutate(&in)

			_, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				adapter.DeployCommitteeVerifier(),
				e.BlockChains,
				in,
			)
			require.ErrorContains(t, err, tc.wantErrSub)
		})
	}
}

// TestEVMCommitteeVerifierDeployAdapter_HappyPath proves the chain-agnostic
// wrapper correctly delegates to the underlying EVM
// sequences.DeployCommitteeVerifier and returns its deployed addresses. It
// pre-deploys a CREATE2Factory and supplies a syntactically valid RMNProxy
// address ref (the CV constructor takes the address as data; no on-chain
// proxy call is made by the deploy path).
func TestEVMCommitteeVerifierDeployAdapter_HappyPath(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	evmChain := e.BlockChains.EVMChains()[testChainSelector]

	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, evmChain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *create2_factory.Version),
			ChainSelector:  testChainSelector,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{evmChain.DeployerKey.From},
			},
		}, nil,
	)
	require.NoError(t, err)

	adapter := &adapters.EVMCommitteeVerifierDeployAdapter{}

	in := ccvdeploymentadapters.DeployCommitteeVerifierInput{
		ChainSelector:     testChainSelector,
		DeployerContract:  create2FactoryRef.Address,
		ExistingAddresses: []datastore.AddressRef{rmnProxyRef(testChainSelector)},
		Params: ccvdeploymentadapters.CommitteeVerifierDeployParams{
			Version:          committee_verifier.Version,
			FeeAggregator:    testFeeAggregator,
			AllowlistAdmin:   testAllowlistAdmin,
			StorageLocations: []string{"https://test.chain.link.fake"},
			Qualifier:        "alpha",
		},
	}

	report, err := cldf_ops.ExecuteSequence(
		e.OperationsBundle, adapter.DeployCommitteeVerifier(), e.BlockChains, in,
	)
	require.NoError(t, err)

	// The wrapper must surface both the verifier and the resolver deployed by
	// the underlying sequence.
	wantTypes := map[datastore.ContractType]bool{
		datastore.ContractType(committee_verifier.ContractType):                   false,
		datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType): false,
	}
	for _, ref := range report.Output.Addresses {
		if _, ok := wantTypes[ref.Type]; ok {
			wantTypes[ref.Type] = true
			require.Equal(t, "alpha", ref.Qualifier)
			require.Equal(t, testChainSelector, ref.ChainSelector)
		}
	}
	for ctype, found := range wantTypes {
		require.True(t, found, "expected contract type %s in adapter output", ctype)
	}
}

// TestEVMCommitteeVerifierDeployAdapter_AllowlistAdminOptional confirms an
// empty AllowlistAdmin is accepted and defaults to the zero address rather
// than failing validation. This is the inverse of the malformed-admin
// validation case.
func TestEVMCommitteeVerifierDeployAdapter_AllowlistAdminOptional(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	evmChain := e.BlockChains.EVMChains()[testChainSelector]
	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, evmChain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *create2_factory.Version),
			ChainSelector:  testChainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{evmChain.DeployerKey.From}},
		}, nil,
	)
	require.NoError(t, err)

	adapter := &adapters.EVMCommitteeVerifierDeployAdapter{}

	in := ccvdeploymentadapters.DeployCommitteeVerifierInput{
		ChainSelector:     testChainSelector,
		DeployerContract:  create2FactoryRef.Address,
		ExistingAddresses: []datastore.AddressRef{rmnProxyRef(testChainSelector)},
		Params: ccvdeploymentadapters.CommitteeVerifierDeployParams{
			Version:       committee_verifier.Version,
			FeeAggregator: testFeeAggregator,
			// AllowlistAdmin intentionally omitted
			Qualifier: "alpha",
		},
	}

	_, err = cldf_ops.ExecuteSequence(
		e.OperationsBundle, adapter.DeployCommitteeVerifier(), e.BlockChains, in,
	)
	require.NoError(t, err)
}

