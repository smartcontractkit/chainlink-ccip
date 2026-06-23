package adapters_test

import (
	"math"
	"math/big"
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
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"

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
		{
			name: "FamilyExtras executorBlockDepth wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsExecutorBlockDepthExtra: "nope"}
			},
			wantErrSub: "must be an integer",
		},
		{
			name: "FamilyExtras executorWaitForSafe wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsExecutorWaitForSafeExtra: "nope"}
			},
			wantErrSub: "must be a bool",
		},
		{
			name: "FamilyExtras executorMaxCCVsPerMsg wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsExecutorMaxCCVsPerMsgExtra: "nope"}
			},
			wantErrSub: "must be an integer",
		},
		{
			name: "FamilyExtras executorMaxCCVsPerMsg out of range",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsExecutorMaxCCVsPerMsgExtra: int64(256)}
			},
			wantErrSub: "must be in [0, 255]",
		},
		{
			name: "FamilyExtras executorCcvAllowlistEnabled wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsExecutorCcvAllowlistEnabledExtra: "nope"}
			},
			wantErrSub: "must be a bool",
		},
		{
			name: "FamilyExtras offRampGasForCallExactCheck wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsOffRampGasForCallExactCheckExtra: "nope"}
			},
			wantErrSub: "must be an integer",
		},
		{
			name: "FamilyExtras onRampMaxUsdCentsPerMessage out of range",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsOnRampMaxUSDCentsPerMessageExtra: int64(math.MaxUint32) + 1}
			},
			wantErrSub: "must be in [0,",
		},
		{
			name: "FamilyExtras feeQuoterMaxFeeJuelsPerMsg wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsFeeQuoterMaxFeeJuelsPerMsgExtra: int64(5)}
			},
			wantErrSub: "must be a base-10 integer string",
		},
		{
			name: "FamilyExtras feeQuoterUsdPerLink unparseable string",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsFeeQuoterUSDPerLINKExtra: "not-a-number"}
			},
			wantErrSub: "must be a base-10 integer string",
		},
		{
			name: "FamilyExtras rmnRemoteLegacyRmn wrong type",
			mutate: func(in *ccvdeploymentadapters.ProtocolContractsDeployInput) {
				in.FamilyExtras = map[string]any{adapters.ProtocolContractsRMNRemoteLegacyRMNExtra: 42}
			},
			wantErrSub: "must be a string",
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

// TestEVMProtocolContractsDeployAdapter_ExecutorOverrides proves the executor's
// deploy-time config (allowed finality, max CCVs per message, CCV allowlist) is
// set from the FamilyExtras overrides. Executors validate the requested finality
// in getFee and are never reconfigured by the lane-configuration changesets, so
// deploy time is the only opportunity to set these (e.g. wait-for-safe).
func TestEVMProtocolContractsDeployAdapter_ExecutorOverrides(t *testing.T) {
	tests := []struct {
		name          string
		extras        map[string]any
		wantFinality  [4]byte
		wantMaxCCVs   uint8
		wantAllowlist bool
	}{
		{
			name:          "defaults when unset",
			extras:        map[string]any{adapters.ProtocolContractsFeeAggregatorExtra: testFeeAggregator},
			wantFinality:  finality.Config{BlockDepth: 1}.Raw(),
			wantMaxCCVs:   10,
			wantAllowlist: false,
		},
		{
			name: "all overrides applied",
			extras: map[string]any{
				adapters.ProtocolContractsFeeAggregatorExtra:               testFeeAggregator,
				adapters.ProtocolContractsExecutorBlockDepthExtra:          int64(1),
				adapters.ProtocolContractsExecutorWaitForSafeExtra:         true,
				adapters.ProtocolContractsExecutorMaxCCVsPerMsgExtra:       int64(5),
				adapters.ProtocolContractsExecutorCcvAllowlistEnabledExtra: true,
			},
			wantFinality:  finality.Config{BlockDepth: 1, WaitForSafe: true}.Raw(),
			wantMaxCCVs:   5,
			wantAllowlist: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
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
				DeployerKeyOwned: true,
				Executors: []ccvdeploymentadapters.ExecutorDeployParams{
					{Version: executor.Version, Qualifier: "default"},
				},
				FamilyExtras: tc.extras,
			}

			report, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle, adapter.DeployProtocolContracts(), e.BlockChains, in,
			)
			require.NoError(t, err)

			var executorAddr string
			for _, ref := range report.Output.Addresses {
				if ref.Type == datastore.ContractType(executor.ContractType) {
					executorAddr = ref.Address
					break
				}
			}
			require.NotEmpty(t, executorAddr, "executor must be deployed")

			finalityCfg, err := cldf_ops.ExecuteOperation(
				e.OperationsBundle, executor.GetAllowedFinalityConfig, evmChain,
				contract_utils.FunctionInput[struct{}]{
					ChainSelector: evmChain.Selector,
					Address:       common.HexToAddress(executorAddr),
				},
			)
			require.NoError(t, err)
			require.Equal(t, tc.wantFinality, finalityCfg.Output)

			dynamicCfg, err := cldf_ops.ExecuteOperation(
				e.OperationsBundle, executor.GetDynamicConfig, evmChain,
				contract_utils.FunctionInput[struct{}]{
					ChainSelector: evmChain.Selector,
					Address:       common.HexToAddress(executorAddr),
				},
			)
			require.NoError(t, err)
			require.Equal(t, tc.wantAllowlist, dynamicCfg.Output.CcvAllowlistEnabled)

			maxCCVs, err := cldf_ops.ExecuteOperation(
				e.OperationsBundle, executor.GetMaxCCVsPerMessage, evmChain,
				contract_utils.FunctionInput[struct{}]{
					ChainSelector: evmChain.Selector,
					Address:       common.HexToAddress(executorAddr),
				},
			)
			require.NoError(t, err)
			require.Equal(t, tc.wantMaxCCVs, maxCCVs.Output)
		})
	}
}

// TestEVMProtocolContractsDeployAdapter_ContractParamOverrides proves the
// optional per-contract deploy params (OffRamp gas, OnRamp max message fee,
// FeeQuoter prices/multipliers, RMNRemote legacy address) are applied from
// FamilyExtras. It sets every override and reads back the on-chain static
// configs that expose them — including the big.Int MaxFeeJuelsPerMsg, which must
// be passed as a base-10 string.
func TestEVMProtocolContractsDeployAdapter_ContractParamOverrides(t *testing.T) {
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

	wantMaxFeeJuels, ok := new(big.Int).SetString("300000000000000000000", 10)
	require.True(t, ok)

	in := ccvdeploymentadapters.ProtocolContractsDeployInput{
		ChainSelector:    testChainSelector,
		DeployerContract: create2FactoryRef.Address,
		DeployerKeyOwned: true,
		Executors: []ccvdeploymentadapters.ExecutorDeployParams{
			{Version: executor.Version, Qualifier: "default"},
		},
		FamilyExtras: map[string]any{
			adapters.ProtocolContractsFeeAggregatorExtra:                           testFeeAggregator,
			adapters.ProtocolContractsOffRampGasForCallExactCheckExtra:             int64(7000),
			adapters.ProtocolContractsOffRampMaxGasBufferToUpdateStateExtra:        int64(15000),
			adapters.ProtocolContractsOnRampMaxUSDCentsPerMessageExtra:             int64(50000),
			adapters.ProtocolContractsFeeQuoterMaxFeeJuelsPerMsgExtra:              "300000000000000000000",
			adapters.ProtocolContractsFeeQuoterLINKPremiumMultiplierWeiPerEthExtra: int64(800000000000000000),
			adapters.ProtocolContractsFeeQuoterWETHPremiumMultiplierWeiPerEthExtra: int64(1100000000000000000),
			adapters.ProtocolContractsFeeQuoterUSDPerLINKExtra:                     "16000000000000000000",
			adapters.ProtocolContractsFeeQuoterUSDPerWETHExtra:                     "2500000000000000000000",
			adapters.ProtocolContractsRMNRemoteLegacyRMNExtra:                      "0x000000000000000000000000000000000000bEEF",
		},
	}

	report, err := cldf_ops.ExecuteSequence(
		e.OperationsBundle, adapter.DeployProtocolContracts(), e.BlockChains, in,
	)
	require.NoError(t, err)

	addrByType := make(map[datastore.ContractType]string)
	for _, ref := range report.Output.Addresses {
		addrByType[ref.Type] = ref.Address
	}

	offRampStatic, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle, offramp.GetStaticConfig, evmChain,
		contract_utils.FunctionInput[struct{}]{
			ChainSelector: evmChain.Selector,
			Address:       common.HexToAddress(addrByType[datastore.ContractType(offramp.ContractType)]),
		},
	)
	require.NoError(t, err)
	require.Equal(t, uint16(7000), offRampStatic.Output.GasForCallExactCheck)
	require.Equal(t, uint32(15000), offRampStatic.Output.MaxGasBufferToUpdateState)

	onRampStatic, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle, onramp.GetStaticConfig, evmChain,
		contract_utils.FunctionInput[struct{}]{
			ChainSelector: evmChain.Selector,
			Address:       common.HexToAddress(addrByType[datastore.ContractType(onramp.ContractType)]),
		},
	)
	require.NoError(t, err)
	require.Equal(t, uint32(50000), onRampStatic.Output.MaxUSDCentsPerMessage)

	feeQuoterStatic, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle, fee_quoter.GetStaticConfig, evmChain,
		contract_utils.FunctionInput[struct{}]{
			ChainSelector: evmChain.Selector,
			Address:       common.HexToAddress(addrByType[datastore.ContractType(fee_quoter.ContractType)]),
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, wantMaxFeeJuels.Cmp(feeQuoterStatic.Output.MaxFeeJuelsPerMsg))
}
