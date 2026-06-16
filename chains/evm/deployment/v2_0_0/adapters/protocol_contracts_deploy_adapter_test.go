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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

// TestEVMProtocolContractsDeployAdapter_Validation exercises the error branches
// in the chain-agnostic-input → EVM-input conversion that runs before the
// underlying DeployChainContracts sequence executes.
func TestEVMProtocolContractsDeployAdapter_Validation(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	adapter := &adapters.EVMProtocolContractsDeployAdapter{}

	tests := []struct {
		name       string
		mutate     func(in *ccvdeploymentadapters.ProtocolContractsDeployInput)
		wantErrSub string
	}{
		{
			name:       "chain not in BlockChains",
			mutate:     func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) { in.ChainSelector = 1 },
			wantErrSub: "EVM chain not found",
		},
		{
			name: "DeployerContract empty",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.DeployerContract = ""
			},
			wantErrSub: "DeployerContract is required",
		},
		{
			name: "DeployerContract not hex",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.DeployerContract = "not-a-hex"
			},
			wantErrSub: "is not a valid hex address",
		},
		{
			name: "FamilyExtras feeAggregator wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsFeeAggregatorExtra: 42}
			},
			wantErrSub: "must be a string hex address",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := ccvdeploymentadapters.ProtocolContractsDeployInput{
				ChainSelector:    testChainSelector,
				DeployerContract: "0x0000000000000000000000000000000000000FAC",
				DeployerKeyOwned: true,
			}
			tc.mutate(&in)

			_, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				adapter.DeployProtocolContracts(),
				e.BlockChains,
				in,
			)
			require.ErrorContains(t, err, tc.wantErrSub)
		})
	}
}

// TestEVMProtocolContractsDeployAdapter_HappyPath proves the adapter deploys the
// core CCIP protocol contracts (Router, FeeQuoter, OnRamp, OffRamp, Executor)
// via the underlying EVM DeployChainContracts sequence — and crucially deploys
// NO committee verifier, since those are deployed separately in the
// topology-free flow.
func TestEVMProtocolContractsDeployAdapter_HappyPath(t *testing.T) {
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

	adapter := &adapters.EVMProtocolContractsDeployAdapter{}

	in := ccvdeploymentadapters.ProtocolContractsDeployInput{
		ChainSelector:    testChainSelector,
		DeployerContract: create2FactoryRef.Address,
		// DeployerKeyOwned skips the MCMS timelock ownership-transfer step, which
		// would otherwise require pre-deployed timelock contracts.
		DeployerKeyOwned: true,
		Executors: []ccvdeploymentadapters.ExecutorDeployParams{
			{Version: executor.Version, Qualifier: "default"},
		},
		FamilyExtras: map[string]any{
			adapters.ProtocolContractsFeeAggregatorExtra: testFeeAggregator,
		},
	}

	report, err := cldf_ops.ExecuteSequence(
		e.OperationsBundle, adapter.DeployProtocolContracts(), e.BlockChains, in,
	)
	require.NoError(t, err)

	deployed := make(map[datastore.ContractType]bool)
	for _, ref := range report.Output.Addresses {
		deployed[ref.Type] = true
	}

	// Core protocol contracts must be present.
	for _, want := range []datastore.ContractType{
		datastore.ContractType(router.ContractType),
		datastore.ContractType(fee_quoter.ContractType),
		datastore.ContractType(onramp.ContractType),
		datastore.ContractType(offramp.ContractType),
		datastore.ContractType(executor.ContractType),
	} {
		require.True(t, deployed[want], "expected protocol contract %s in adapter output", want)
	}

	// Committee verifiers are deployed separately — none should appear here.
	require.False(t, deployed[datastore.ContractType(committee_verifier.ContractType)],
		"committee verifier must not be deployed by the protocol-only adapter")
}
