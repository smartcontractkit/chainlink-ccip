package sequences_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	changesetadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

var evmFamilySelector = [4]byte{0x28, 0x12, 0xd5, 0x2c}

func boolPtr(v bool) *bool { return &v }

type deployedContracts struct {
	router                    string
	onRamp                    string
	feeQuoter                 string
	offRamp                   string
	committeeVerifier         string
	committeeVerifierResolver string
	executor                  string
}

func deployChain(
	t *testing.T,
	e *deployment.Environment,
	chainSelector uint64,
) deployedContracts {
	t.Helper()
	evmChain := e.BlockChains.EVMChains()[chainSelector]

	create2Ref, err := contract.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, evmChain,
		contract.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{evmChain.DeployerKey.From}},
		}, nil,
	)
	require.NoError(t, err)

	report, err := operations.ExecuteSequence(
		e.OperationsBundle, sequences.DeployChainContracts, evmChain,
		sequences.DeployChainContractsInput{
			ChainSelector:    chainSelector,
			CREATE2Factory:   common.HexToAddress(create2Ref.Address),
			ContractParams:   testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err)

	var out deployedContracts
	for _, addr := range report.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(router.ContractType):
			out.router = addr.Address
		case datastore.ContractType(onramp.ContractType):
			out.onRamp = addr.Address
		case datastore.ContractType(fee_quoter.ContractType):
			out.feeQuoter = addr.Address
		case datastore.ContractType(offramp.ContractType):
			out.offRamp = addr.Address
		case datastore.ContractType(committee_verifier.ContractType):
			out.committeeVerifier = addr.Address
		case datastore.ContractType(sequences.ExecutorProxyType):
			if addr.Qualifier == "default" {
				out.executor = addr.Address
			}
		case datastore.ContractType(sequences.CommitteeVerifierResolverType):
			out.committeeVerifierResolver = addr.Address
		}
	}
	return out
}

func buildConfigureChainForLanesInput(
	local deployedContracts,
	localSelector uint64,
	remote deployedContracts,
	remoteSelector uint64,
) changesetadapters.ConfigureChainForLanesInput {
	return changesetadapters.ConfigureChainForLanesInput{
		ChainSelector: localSelector,
		Router:        local.router,
		OnRamp:        local.onRamp,
		FeeQuoter:     local.feeQuoter,
		OffRamp:       local.offRamp,
		CommitteeVerifiers: []changesetadapters.CommitteeVerifierConfig[datastore.AddressRef]{
			{
				CommitteeVerifier: []datastore.AddressRef{
					{Address: local.committeeVerifier, Type: datastore.ContractType(committee_verifier.ContractType), Version: committee_verifier.Version},
					{Address: local.committeeVerifierResolver, Type: datastore.ContractType(sequences.CommitteeVerifierResolverType), Version: semver.MustParse("2.0.0")},
				},
				RemoteChains: map[uint64]changesetadapters.CommitteeVerifierRemoteChainConfig{
					remoteSelector: testsetup.AdapterCommitteeVerifierRemoteChainConfig(),
				},
			},
		},
		RemoteChains: map[uint64]changesetadapters.RemoteChainConfig[[]byte, string]{
			remoteSelector: {
				AllowTrafficFrom:    boolPtr(true),
				OnRamps:             [][]byte{common.HexToAddress(remote.onRamp).Bytes()},
				OffRamp:             common.HexToAddress(remote.offRamp).Bytes(),
				DefaultExecutor:     local.executor,
				DefaultInboundCCVs:  []string{local.committeeVerifier},
				DefaultOutboundCCVs: []string{local.committeeVerifier},
				FeeQuoterDestChainConfig: changesetadapters.FeeQuoterDestChainConfig{
					IsEnabled:                   true,
					MaxDataBytes:                30_000,
					MaxPerMsgGasLimit:           3_000_000,
					DestGasOverhead:             300_000,
					DefaultTokenFeeUSDCents:     25,
					DestGasPerPayloadByteBase:   16,
					DefaultTokenDestGasOverhead: 90_000,
					DefaultTxGasLimit:           200_000,
					NetworkFeeUSDCents:          10,
					ChainFamilySelector:         evmFamilySelector,
					LinkFeeMultiplierPercent:     90,
				},
				ExecutorDestChainConfig: changesetadapters.ExecutorDestChainConfig{
					USDCentsFee: 50,
					Enabled:     true,
				},
				AddressBytesLength:   20,
				BaseExecutionGasCost: 80_000,
			},
		},
	}
}

func assertOnChainState(
	t *testing.T,
	b operations.Bundle,
	evmChain cldfevm.Chain,
	local deployedContracts,
	remote deployedContracts,
	remoteChainSelector uint64,
) {
	t.Helper()

	onRampOnRouter, err := operations.ExecuteOperation(b, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.router), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, local.onRamp, onRampOnRouter.Output.Hex())

	offRampsOnRouter, err := operations.ExecuteOperation(b, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.router),
	})
	require.NoError(t, err)
	require.Len(t, offRampsOnRouter.Output, 1)
	assert.Equal(t, local.offRamp, offRampsOnRouter.Output[0].OffRamp.Hex())

	srcCfg, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.True(t, srcCfg.Output.IsEnabled)
	assert.Equal(t, local.router, srcCfg.Output.Router.Hex())
	assert.Len(t, srcCfg.Output.OnRamps, 1)
	assert.Equal(t, common.LeftPadBytes(common.HexToAddress(remote.onRamp).Bytes(), 32), srcCfg.Output.OnRamps[0])
	assert.Len(t, srcCfg.Output.DefaultCCVs, 1)
	assert.Equal(t, local.committeeVerifier, srcCfg.Output.DefaultCCVs[0].Hex())

	destCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.onRamp), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, local.router, destCfg.Output.Router.Hex())
	assert.Equal(t, common.HexToAddress(remote.offRamp).Bytes(), destCfg.Output.OffRamp)
	assert.Equal(t, local.executor, destCfg.Output.DefaultExecutor.Hex())
	assert.Len(t, destCfg.Output.DefaultCCVs, 1)
	assert.Equal(t, local.committeeVerifier, destCfg.Output.DefaultCCVs[0].Hex())

	executorDestChains, err := operations.ExecuteOperation(b, executor.GetDestChains, evmChain, contract.FunctionInput[struct{}]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.executor),
	})
	require.NoError(t, err)
	require.Len(t, executorDestChains.Output, 1)
	assert.Equal(t, remoteChainSelector, executorDestChains.Output[0].DestChainSelector)
	assert.Equal(t, uint16(50), executorDestChains.Output[0].Config.UsdCentsFee)
	assert.True(t, executorDestChains.Output[0].Config.Enabled)

	fqDestCfg, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.True(t, fqDestCfg.Output.IsEnabled)
	assert.Equal(t, uint32(30_000), fqDestCfg.Output.MaxDataBytes)
	assert.Equal(t, uint32(3_000_000), fqDestCfg.Output.MaxPerMsgGasLimit)
	assert.Equal(t, uint32(300_000), fqDestCfg.Output.DestGasOverhead)
	assert.Equal(t, evmFamilySelector, fqDestCfg.Output.ChainFamilySelector)

	verifierRemoteCfg, err := operations.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, local.router, verifierRemoteCfg.Output.Router.Hex())
	assert.False(t, verifierRemoteCfg.Output.AllowlistEnabled)
}

func TestConfigureChainForLanes_ConfiguresSingleRemoteChainEndToEnd(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteChainSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteChainSelector)

	input := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteChainSelector)

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	assertOnChainState(t, testsetup.BundleWithFreshReporter(e.OperationsBundle), evmChain, local, remote, remoteChainSelector)
}

func TestConfigureChainForLanes_Idempotent(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteChainSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteChainSelector)

	input := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteChainSelector)

	firstReport, err := operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	assertOnChainState(t, testsetup.BundleWithFreshReporter(e.OperationsBundle), evmChain, local, remote, remoteChainSelector)

	secondReport, err := operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	assertOnChainState(t, testsetup.BundleWithFreshReporter(e.OperationsBundle), evmChain, local, remote, remoteChainSelector)
	assert.Less(t, len(secondReport.ExecutionReports), len(firstReport.ExecutionReports),
		"second invocation should execute fewer operations (idempotent skip)")
}

func TestConfigureChainForLanes_AddingNewLaneDoesNotAffectExistingLane(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteA := chainsel.TEST_90000002.Selector
	remoteB := chainsel.TEST_90000003.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteA, remoteB}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remoteAContracts := deployChain(t, e, remoteA)
	remoteBContracts := deployChain(t, e, remoteB)

	inputA := buildConfigureChainForLanesInput(local, chainSelector, remoteAContracts, remoteA)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, inputA,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	snapshotLaneA := captureLaneState(t, b, evmChain, local, remoteAContracts, remoteA)

	inputB := buildConfigureChainForLanesInput(local, chainSelector, remoteBContracts, remoteB)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, inputB,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)

	snapshotLaneAAfter := captureLaneA(t, b, evmChain, local, remoteAContracts, remoteA)
	assert.Equal(t, snapshotLaneA, snapshotLaneAAfter, "lane A config must be identical after adding lane B")

	snapshotLaneB := captureLaneState(t, b, evmChain, local, remoteBContracts, remoteB)
	assert.True(t, snapshotLaneB.offRampSourceEnabled, "lane B OffRamp source should be enabled")
	assert.Equal(t, local.router, snapshotLaneB.onRampDestRouter, "lane B OnRamp should reference the router")
	assert.Equal(t, local.router, snapshotLaneB.verifierRouter, "lane B verifier should reference the router")
	assert.True(t, snapshotLaneB.feeQuoterEnabled, "lane B FeeQuoter should be enabled")
	assert.True(t, snapshotLaneB.executorEnabled, "lane B executor should be enabled")

	offRampsOnRouter, err := operations.ExecuteOperation(b, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.router),
	})
	require.NoError(t, err)
	assert.Len(t, offRampsOnRouter.Output, 2, "router should have offramps for both lanes")
}

type laneSnapshot struct {
	offRampSourceEnabled bool
	offRampRouter        string
	offRampOnRamps       [][]byte
	offRampDefaultCCVs   []common.Address

	onRampDestRouter    string
	onRampOffRamp       []byte
	onRampDefaultExec   string
	onRampDefaultCCVs   []common.Address
	onRampAddrBytes     uint8
	onRampBaseGas       uint32
	onRampMsgFeeUSD     uint16
	onRampTokenFeeUSD   uint16
	onRampTokenReceiver bool

	feeQuoterEnabled    bool
	feeQuoterMaxData    uint32
	feeQuoterMaxGas     uint32
	feeQuoterGasOvhd    uint32
	feeQuoterFamily     [4]byte
	feeQuoterNetFee     uint16
	feeQuoterLinkMult   uint8
	feeQuoterDefTokFee  uint16
	feeQuoterDefTokGas  uint32
	feeQuoterDefTxGas   uint32
	feeQuoterGasPerByte uint8

	executorEnabled bool
	executorFee     uint16

	verifierRouter           string
	verifierAllowlistEnabled bool
	verifierAllowedSenders   []common.Address
	verifierFeeUSDCents      uint16
	verifierGasForVerify     uint32
	verifierPayloadSize      uint32
	verifierSigners          []common.Address
	verifierThreshold        uint8
}

func captureLaneState(
	t *testing.T,
	b operations.Bundle,
	evmChain cldfevm.Chain,
	local deployedContracts,
	remote deployedContracts,
	remoteSelector uint64,
) laneSnapshot {
	t.Helper()
	var s laneSnapshot

	srcCfg, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	s.offRampSourceEnabled = srcCfg.Output.IsEnabled
	s.offRampRouter = srcCfg.Output.Router.Hex()
	s.offRampOnRamps = srcCfg.Output.OnRamps
	s.offRampDefaultCCVs = srcCfg.Output.DefaultCCVs

	destCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.onRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	s.onRampDestRouter = destCfg.Output.Router.Hex()
	s.onRampOffRamp = destCfg.Output.OffRamp
	s.onRampDefaultExec = destCfg.Output.DefaultExecutor.Hex()
	s.onRampDefaultCCVs = destCfg.Output.DefaultCCVs
	s.onRampAddrBytes = destCfg.Output.AddressBytesLength
	s.onRampBaseGas = destCfg.Output.BaseExecutionGasCost
	s.onRampMsgFeeUSD = destCfg.Output.MessageNetworkFeeUSDCents
	s.onRampTokenFeeUSD = destCfg.Output.TokenNetworkFeeUSDCents
	s.onRampTokenReceiver = destCfg.Output.TokenReceiverAllowed

	fqCfg, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteSelector,
	})
	require.NoError(t, err)
	s.feeQuoterEnabled = fqCfg.Output.IsEnabled
	s.feeQuoterMaxData = fqCfg.Output.MaxDataBytes
	s.feeQuoterMaxGas = fqCfg.Output.MaxPerMsgGasLimit
	s.feeQuoterGasOvhd = fqCfg.Output.DestGasOverhead
	s.feeQuoterFamily = fqCfg.Output.ChainFamilySelector
	s.feeQuoterNetFee = fqCfg.Output.NetworkFeeUSDCents
	s.feeQuoterLinkMult = fqCfg.Output.LinkFeeMultiplierPercent
	s.feeQuoterDefTokFee = fqCfg.Output.DefaultTokenFeeUSDCents
	s.feeQuoterDefTokGas = fqCfg.Output.DefaultTokenDestGasOverhead
	s.feeQuoterDefTxGas = fqCfg.Output.DefaultTxGasLimit
	s.feeQuoterGasPerByte = fqCfg.Output.DestGasPerPayloadByteBase

	executorDestChains, err := operations.ExecuteOperation(b, executor.GetDestChains, evmChain, contract.FunctionInput[struct{}]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.executor),
	})
	require.NoError(t, err)
	for _, dc := range executorDestChains.Output {
		if dc.DestChainSelector == remoteSelector {
			s.executorEnabled = dc.Config.Enabled
			s.executorFee = dc.Config.UsdCentsFee
			break
		}
	}

	verifierCfg, err := operations.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier), Args: remoteSelector,
	})
	require.NoError(t, err)
	s.verifierRouter = verifierCfg.Output.Router.Hex()
	s.verifierAllowlistEnabled = verifierCfg.Output.AllowlistEnabled
	s.verifierAllowedSenders = verifierCfg.Output.AllowedSendersList

	verifierFee, err := operations.ExecuteOperation(b, committee_verifier.GetFee, evmChain, contract.FunctionInput[committee_verifier.GetFeeArgs]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier),
		Args: committee_verifier.GetFeeArgs{DestChainSelector: remoteSelector},
	})
	require.NoError(t, err)
	s.verifierFeeUSDCents = verifierFee.Output.FeeUSDCents
	s.verifierGasForVerify = verifierFee.Output.GasForVerification
	s.verifierPayloadSize = verifierFee.Output.PayloadSizeBytes

	sigCfg, err := operations.ExecuteOperation(b, committee_verifier.GetSignatureConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier), Args: remoteSelector,
	})
	require.NoError(t, err)
	s.verifierSigners = sigCfg.Output.Signers
	s.verifierThreshold = sigCfg.Output.Threshold

	return s
}

// captureLaneA is an alias that makes the test read naturally.
var captureLaneA = captureLaneState

type deployedContractsWithTestRouter struct {
	deployedContracts
	testRouter string
}

func deployChainWithTestRouter(
	t *testing.T,
	e *deployment.Environment,
	chainSelector uint64,
) deployedContractsWithTestRouter {
	t.Helper()
	evmChain := e.BlockChains.EVMChains()[chainSelector]

	create2Ref, err := contract.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, evmChain,
		contract.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{evmChain.DeployerKey.From}},
		}, nil,
	)
	require.NoError(t, err)

	report, err := operations.ExecuteSequence(
		e.OperationsBundle, sequences.DeployChainContracts, evmChain,
		sequences.DeployChainContractsInput{
			ChainSelector:    chainSelector,
			CREATE2Factory:   common.HexToAddress(create2Ref.Address),
			ContractParams:   testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
			DeployTestRouter: true,
		},
	)
	require.NoError(t, err)

	var out deployedContractsWithTestRouter
	for _, addr := range report.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(router.ContractType):
			out.router = addr.Address
		case datastore.ContractType(router.TestRouterContractType):
			out.testRouter = addr.Address
		case datastore.ContractType(onramp.ContractType):
			out.onRamp = addr.Address
		case datastore.ContractType(fee_quoter.ContractType):
			out.feeQuoter = addr.Address
		case datastore.ContractType(offramp.ContractType):
			out.offRamp = addr.Address
		case datastore.ContractType(committee_verifier.ContractType):
			out.committeeVerifier = addr.Address
		case datastore.ContractType(sequences.ExecutorProxyType):
			if addr.Qualifier == "default" {
				out.executor = addr.Address
			}
		case datastore.ContractType(sequences.CommitteeVerifierResolverType):
			out.committeeVerifierResolver = addr.Address
		}
	}
	return out
}

func TestConfigureChainForLanes_ConfiguresMultipleRemoteChainsInSingleCall(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteA := chainsel.TEST_90000002.Selector
	remoteB := chainsel.TEST_90000003.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteA, remoteB}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remoteAContracts := deployChain(t, e, remoteA)
	remoteBContracts := deployChain(t, e, remoteB)

	input := buildConfigureChainForLanesInput(local, chainSelector, remoteAContracts, remoteA)
	input.RemoteChains[remoteB] = changesetadapters.RemoteChainConfig[[]byte, string]{
		AllowTrafficFrom:    boolPtr(true),
		OnRamps:             [][]byte{common.HexToAddress(remoteBContracts.onRamp).Bytes()},
		OffRamp:             common.HexToAddress(remoteBContracts.offRamp).Bytes(),
		DefaultExecutor:     local.executor,
		DefaultInboundCCVs:  []string{local.committeeVerifier},
		DefaultOutboundCCVs: []string{local.committeeVerifier},
		FeeQuoterDestChainConfig: changesetadapters.FeeQuoterDestChainConfig{
			IsEnabled:                   true,
			MaxDataBytes:                50_000,
			MaxPerMsgGasLimit:           5_000_000,
			DestGasOverhead:             400_000,
			DefaultTokenFeeUSDCents:     30,
			DestGasPerPayloadByteBase:   16,
			DefaultTokenDestGasOverhead: 100_000,
			DefaultTxGasLimit:           250_000,
			NetworkFeeUSDCents:          20,
			ChainFamilySelector:      evmFamilySelector,
			LinkFeeMultiplierPercent: 90,
		},
		ExecutorDestChainConfig: changesetadapters.ExecutorDestChainConfig{
			USDCentsFee: 100,
			Enabled:     true,
		},
		AddressBytesLength:   20,
		BaseExecutionGasCost: 80_000,
	}
	input.CommitteeVerifiers[0].RemoteChains[remoteB] = testsetup.AdapterCommitteeVerifierRemoteChainConfig()

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	offRampsOnRouter, err := operations.ExecuteOperation(b, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.router),
	})
	require.NoError(t, err)
	assert.Len(t, offRampsOnRouter.Output, 2)

	for _, remoteSelector := range []uint64{remoteA, remoteB} {
		srcCfg, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
			ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteSelector,
		})
		require.NoError(t, err)
		assert.True(t, srcCfg.Output.IsEnabled, "source chain config for %d should be enabled", remoteSelector)
	}

	fqDestA, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteA,
	})
	require.NoError(t, err)
	assert.Equal(t, uint32(30_000), fqDestA.Output.MaxDataBytes)

	fqDestB, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteB,
	})
	require.NoError(t, err)
	assert.Equal(t, uint32(50_000), fqDestB.Output.MaxDataBytes)

	executorDestChains, err := operations.ExecuteOperation(b, executor.GetDestChains, evmChain, contract.FunctionInput[struct{}]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.executor),
	})
	require.NoError(t, err)
	assert.Len(t, executorDestChains.Output, 2)
}

func TestConfigureChainForLanes_CommitteeVerifierSignatureConfigUpdate(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteA := chainsel.TEST_90000002.Selector
	remoteB := chainsel.TEST_90000003.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteA, remoteB}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remoteAContracts := deployChain(t, e, remoteA)
	remoteBContracts := deployChain(t, e, remoteB)

	signerA := common.HexToAddress("0xA1")
	signerB := common.HexToAddress("0xB1")
	signerB2 := common.HexToAddress("0xB2")

	inputA := buildConfigureChainForLanesInput(local, chainSelector, remoteAContracts, remoteA)
	cvCfgA := testsetup.AdapterCommitteeVerifierRemoteChainConfig()
	cvCfgA.SignatureConfig.Signers = []string{signerA.Hex()}
	cvCfgA.SignatureConfig.Threshold = 1
	inputA.CommitteeVerifiers[0].RemoteChains[remoteA] = cvCfgA

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, inputA,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	snapshotA := captureLaneState(t, b, evmChain, local, remoteAContracts, remoteA)
	require.Len(t, snapshotA.verifierSigners, 1)
	assert.Equal(t, signerA, snapshotA.verifierSigners[0])
	assert.Equal(t, uint8(1), snapshotA.verifierThreshold)

	inputB := buildConfigureChainForLanesInput(local, chainSelector, remoteBContracts, remoteB)
	cvCfgB := testsetup.AdapterCommitteeVerifierRemoteChainConfig()
	cvCfgB.SignatureConfig.Signers = []string{signerB.Hex(), signerB2.Hex()}
	cvCfgB.SignatureConfig.Threshold = 2
	inputB.CommitteeVerifiers[0].RemoteChains[remoteB] = cvCfgB

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, inputB,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	snapshotAAfter := captureLaneState(t, b, evmChain, local, remoteAContracts, remoteA)
	assert.Equal(t, snapshotA, snapshotAAfter, "lane A signature config must be unchanged after adding lane B")

	snapshotBAfter := captureLaneState(t, b, evmChain, local, remoteBContracts, remoteB)
	assert.Len(t, snapshotBAfter.verifierSigners, 2)
	assert.Equal(t, uint8(2), snapshotBAfter.verifierThreshold)
}

func TestConfigureChainForLanes_CommitteeVerifierAllowlistUpdates(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	sender1 := common.HexToAddress("0xAAA1")
	sender2 := common.HexToAddress("0xAAA2")

	input := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	cvCfg := testsetup.AdapterCommitteeVerifierRemoteChainConfig()
	cvCfg.AllowlistEnabled = true
	cvCfg.AddedAllowlistedSenders = []string{sender1.Hex(), sender2.Hex()}
	input.CommitteeVerifiers[0].RemoteChains[remoteSelector] = cvCfg

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	snap1 := captureLaneState(t, b, evmChain, local, remote, remoteSelector)
	assert.True(t, snap1.verifierAllowlistEnabled)
	assert.Len(t, snap1.verifierAllowedSenders, 2)

	input2 := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	cvCfg2 := testsetup.AdapterCommitteeVerifierRemoteChainConfig()
	cvCfg2.AllowlistEnabled = true
	cvCfg2.RemovedAllowlistedSenders = []string{sender1.Hex()}
	input2.CommitteeVerifiers[0].RemoteChains[remoteSelector] = cvCfg2

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input2,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	snap2 := captureLaneState(t, b, evmChain, local, remote, remoteSelector)
	assert.True(t, snap2.verifierAllowlistEnabled)
	require.Len(t, snap2.verifierAllowedSenders, 1)
	assert.Equal(t, sender2, snap2.verifierAllowedSenders[0])
}

func TestConfigureChainForLanes_PartialUpdatePreservesExistingFields(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	fullInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, fullInput,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	snapshotBefore := captureLaneState(t, b, evmChain, local, remote, remoteSelector)

	partialInput := changesetadapters.ConfigureChainForLanesInput{
		ChainSelector: chainSelector,
		Router:        local.router,
		OnRamp:        local.onRamp,
		FeeQuoter:     local.feeQuoter,
		OffRamp:       local.offRamp,
		RemoteChains: map[uint64]changesetadapters.RemoteChainConfig[[]byte, string]{
			remoteSelector: {
				AllowTrafficFrom:    boolPtr(true),
				OnRamps:             [][]byte{common.HexToAddress(remote.onRamp).Bytes()},
				OffRamp:             common.HexToAddress(remote.offRamp).Bytes(),
				DefaultInboundCCVs:  []string{local.committeeVerifier},
				DefaultOutboundCCVs: []string{local.committeeVerifier},
				BaseExecutionGasCost: 120_000,
				FeeQuoterDestChainConfig: changesetadapters.FeeQuoterDestChainConfig{
					OverrideExistingConfig: true,
					DestGasOverhead:        500_000,
				},
				DefaultExecutor: local.executor,
				ExecutorDestChainConfig: changesetadapters.ExecutorDestChainConfig{
					USDCentsFee: 50,
					Enabled:     true,
				},
			},
		},
	}

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, partialInput,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	snapshotAfter := captureLaneState(t, b, evmChain, local, remote, remoteSelector)

	assert.Equal(t, uint32(120_000), snapshotAfter.onRampBaseGas, "BaseExecutionGasCost should be updated")
	assert.Equal(t, uint32(500_000), snapshotAfter.feeQuoterGasOvhd, "FeeQuoter DestGasOverhead should be updated")

	assert.Equal(t, snapshotBefore.offRampSourceEnabled, snapshotAfter.offRampSourceEnabled)
	assert.Equal(t, snapshotBefore.offRampRouter, snapshotAfter.offRampRouter)
	assert.Equal(t, snapshotBefore.onRampDestRouter, snapshotAfter.onRampDestRouter)
	assert.Equal(t, snapshotBefore.onRampAddrBytes, snapshotAfter.onRampAddrBytes)
	assert.Equal(t, snapshotBefore.feeQuoterMaxData, snapshotAfter.feeQuoterMaxData)
	assert.Equal(t, snapshotBefore.feeQuoterMaxGas, snapshotAfter.feeQuoterMaxGas)
	assert.Equal(t, snapshotBefore.feeQuoterFamily, snapshotAfter.feeQuoterFamily)
	assert.Equal(t, snapshotBefore.feeQuoterNetFee, snapshotAfter.feeQuoterNetFee)
	assert.Equal(t, snapshotBefore.feeQuoterDefTokFee, snapshotAfter.feeQuoterDefTokFee)
	assert.Equal(t, snapshotBefore.feeQuoterDefTokGas, snapshotAfter.feeQuoterDefTokGas)
	assert.Equal(t, snapshotBefore.feeQuoterDefTxGas, snapshotAfter.feeQuoterDefTxGas)
	assert.Equal(t, snapshotBefore.executorEnabled, snapshotAfter.executorEnabled)
	assert.Equal(t, snapshotBefore.executorFee, snapshotAfter.executorFee)
}

func TestConfigureChainForLanes_OverrideExistingFeeQuoterConfig(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	initialInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, initialInput,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	fqBefore, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, uint32(30_000), fqBefore.Output.MaxDataBytes)

	noOverrideInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	rc := noOverrideInput.RemoteChains[remoteSelector]
	rc.FeeQuoterDestChainConfig.OverrideExistingConfig = false
	rc.FeeQuoterDestChainConfig.MaxDataBytes = 99_000
	noOverrideInput.RemoteChains[remoteSelector] = rc

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, noOverrideInput,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	fqAfterNoOverride, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, uint32(30_000), fqAfterNoOverride.Output.MaxDataBytes,
		"with OverrideExistingConfig=false, an already-enabled FeeQuoter config should not be updated")

	overrideInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	rc = overrideInput.RemoteChains[remoteSelector]
	rc.FeeQuoterDestChainConfig.OverrideExistingConfig = true
	rc.FeeQuoterDestChainConfig.MaxDataBytes = 99_000
	overrideInput.RemoteChains[remoteSelector] = rc

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, overrideInput,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	fqAfterOverride, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, uint32(99_000), fqAfterOverride.Output.MaxDataBytes,
		"with OverrideExistingConfig=true, FeeQuoter config should be updated")
}

func TestConfigureChainForLanes_TestRouterSetup(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	localWithTest := deployChainWithTestRouter(t, e, chainSelector)
	require.NotEmpty(t, localWithTest.testRouter, "TestRouter should be deployed")
	require.NotEqual(t, localWithTest.router, localWithTest.testRouter,
		"TestRouter and Router should be different addresses")

	remote := deployChain(t, e, remoteSelector)

	input := buildConfigureChainForLanesInput(localWithTest.deployedContracts, chainSelector, remote, remoteSelector)
	input.Router = localWithTest.testRouter

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	onRampOnTestRouter, err := operations.ExecuteOperation(b, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(localWithTest.testRouter), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, localWithTest.onRamp, onRampOnTestRouter.Output.Hex(),
		"TestRouter should have OnRamp wired")

	offRampsOnTestRouter, err := operations.ExecuteOperation(b, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(localWithTest.testRouter),
	})
	require.NoError(t, err)
	require.Len(t, offRampsOnTestRouter.Output, 1)
	assert.Equal(t, localWithTest.offRamp, offRampsOnTestRouter.Output[0].OffRamp.Hex(),
		"TestRouter should have OffRamp wired")

	onRampOnProdRouter, err := operations.ExecuteOperation(b, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(localWithTest.router), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, common.Address{}, onRampOnProdRouter.Output,
		"production Router should NOT have OnRamp wired")

	offRampsOnProdRouter, err := operations.ExecuteOperation(b, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(localWithTest.router),
	})
	require.NoError(t, err)
	assert.Empty(t, offRampsOnProdRouter.Output,
		"production Router should NOT have OffRamps wired")

	srcCfg, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(localWithTest.offRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, localWithTest.testRouter, srcCfg.Output.Router.Hex(),
		"OffRamp source config should reference the TestRouter")

	destCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(localWithTest.onRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, localWithTest.testRouter, destCfg.Output.Router.Hex(),
		"OnRamp dest config should reference the TestRouter")

	verifierCfg, err := operations.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(localWithTest.committeeVerifier), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, localWithTest.testRouter, verifierCfg.Output.Router.Hex(),
		"CommitteeVerifier remote config should reference the TestRouter")
}

func TestConfigureChainForLanes_AllowTrafficFromFalseDisablesTraffic(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteChainSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteChainSelector)

	enabledInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteChainSelector)
	require.NotNil(t, enabledInput.RemoteChains[remoteChainSelector].AllowTrafficFrom)
	require.True(t, *enabledInput.RemoteChains[remoteChainSelector].AllowTrafficFrom,
		"precondition: buildConfigureChainForLanesInput sets AllowTrafficFrom=true")

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, enabledInput,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	srcCfgEnabled, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	require.True(t, srcCfgEnabled.Output.IsEnabled, "traffic should be enabled after first run")

	disabledInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteChainSelector)
	rc := disabledInput.RemoteChains[remoteChainSelector]
	rc.AllowTrafficFrom = boolPtr(false)
	disabledInput.RemoteChains[remoteChainSelector] = rc

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, disabledInput,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	srcCfgDisabled, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.False(t, srcCfgDisabled.Output.IsEnabled,
		"AllowTrafficFrom=false should disable traffic on the OffRamp source config")
}

func TestConfigureChainForLanes_AllowTrafficFromFalseToTrueEnablesTraffic(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteChainSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteChainSelector)

	disabledInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteChainSelector)
	rc := disabledInput.RemoteChains[remoteChainSelector]
	rc.AllowTrafficFrom = boolPtr(false)
	disabledInput.RemoteChains[remoteChainSelector] = rc

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, disabledInput,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	srcCfgDisabled, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	require.False(t, srcCfgDisabled.Output.IsEnabled, "traffic should be disabled after first run")

	enabledInput := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteChainSelector)

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, enabledInput,
	)
	require.NoError(t, err)

	b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	srcCfgEnabled, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.True(t, srcCfgEnabled.Output.IsEnabled,
		"AllowTrafficFrom=true should re-enable traffic on the OffRamp source config")
}

func TestConfigureChainForLanes_RejectsOnRampOverwriteWhenNotAllowed(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	dummyOnRamp := common.HexToAddress(local.router)
	_, err = operations.ExecuteOperation(b, router.ApplyRampUpdates, evmChain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(local.router),
		Args: router.ApplyRampsUpdatesArgs{
			OnRampUpdates: []router.OnRamp{{DestChainSelector: remoteSelector, OnRamp: dummyOnRamp}},
		},
	})
	require.NoError(t, err)

	input := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	input.AllowOnrampOverride = false

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "refusing to overwrite")
	assert.Contains(t, err.Error(), "AllowOnrampOverride is false")

	b2 := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	onRampReport, err := operations.ExecuteOperation(b2, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector, Address: common.HexToAddress(local.router), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, dummyOnRamp, onRampReport.Output,
		"router onRamp should be unchanged after rejected overwrite")
}

func TestConfigureChainForLanes_OverwritesOnRampWhenAllowed(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	localWithTest := deployChainWithTestRouter(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	dummyOnRamp := common.HexToAddress(localWithTest.testRouter)
	_, err = operations.ExecuteOperation(b, router.ApplyRampUpdates, evmChain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(localWithTest.testRouter),
		Args: router.ApplyRampsUpdatesArgs{
			OnRampUpdates: []router.OnRamp{{DestChainSelector: remoteSelector, OnRamp: dummyOnRamp}},
		},
	})
	require.NoError(t, err)

	input := buildConfigureChainForLanesInput(localWithTest.deployedContracts, chainSelector, remote, remoteSelector)
	input.Router = localWithTest.testRouter
	input.AllowOnrampOverride = true

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	b2 := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	onRampReport, err := operations.ExecuteOperation(b2, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector, Address: common.HexToAddress(localWithTest.testRouter), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, localWithTest.onRamp, onRampReport.Output.Hex(),
		"test router onRamp should be overwritten to the 2.0 onRamp when AllowOnrampOverride is true")
}

func TestConfigureChainForLanes_AllowsFirstTimeWriteWithoutOverrideFlag(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	input := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteSelector)
	input.AllowOnrampOverride = false

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	onRampReport, err := operations.ExecuteOperation(b, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector, Address: common.HexToAddress(local.router), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, local.onRamp, onRampReport.Output.Hex(),
		"router onRamp should be set on first write even with AllowOnrampOverride=false")
}

func TestConfigureChainForLanes_GasPriceUpdateSkippedWhenAlreadySet(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteChainSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
	)
	require.NoError(t, err)

	local := deployChain(t, e, chainSelector)
	remote := deployChain(t, e, remoteChainSelector)

	input := buildConfigureChainForLanesInput(local, chainSelector, remote, remoteChainSelector)
	rc := input.RemoteChains[remoteChainSelector]
	rc.FeeQuoterDestChainConfig.USDPerUnitGas = big.NewInt(42_000)
	input.RemoteChains[remoteChainSelector] = rc

	report1, err := operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)

	gasPrice, err := operations.ExecuteOperation(b, fee_quoter.GetDestinationChainGasPrice, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteChainSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(42_000), gasPrice.Output.Value, "gas price should be 42000 after first run")

	report2, err := operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err)
	assert.Less(t, len(report2.ExecutionReports), len(report1.ExecutionReports),
		"second run should skip gas price update since value is already set")
}

func TestConfigureChainForLanes_RejectsOnRampRouterDowngrade(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	localWithTest := deployChainWithTestRouter(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	input := buildConfigureChainForLanesInput(localWithTest.deployedContracts, chainSelector, remote, remoteSelector)
	input.AllowOnrampOverride = true

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, input,
	)
	require.NoError(t, err, "first run with prod router should succeed")

	downgradeInput := buildConfigureChainForLanesInput(localWithTest.deployedContracts, chainSelector, remote, remoteSelector)
	downgradeInput.Router = localWithTest.testRouter
	downgradeInput.AllowOnrampOverride = false

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, downgradeInput,
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OnRamp")
	assert.Contains(t, err.Error(), "refusing to overwrite")

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	onRampCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector, Address: common.HexToAddress(localWithTest.onRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, common.HexToAddress(localWithTest.router), onRampCfg.Output.Router,
		"OnRamp router should still point to the prod router after rejected downgrade")
}

func TestConfigureChainForLanes_RejectsOffRampRouterDowngrade(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteSelector}),
	)
	require.NoError(t, err)

	localWithTest := deployChainWithTestRouter(t, e, chainSelector)
	remote := deployChain(t, e, remoteSelector)

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	prodRouter := common.HexToAddress(localWithTest.router)

	// Pre-configure only the OffRamp source chain config with the prod router.
	// The OnRamp is left unconfigured so its router is zero — the OnRamp check
	// passes, and the OffRamp check is what triggers the guard.
	_, err = operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		offramp.ApplySourceChainConfigUpdates, evmChain,
		contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
			ChainSelector: chainSelector,
			Address:       common.HexToAddress(localWithTest.offRamp),
			Args: []offramp.SourceChainConfigArgs{{
				Router:              prodRouter,
				SourceChainSelector: remoteSelector,
				IsEnabled:           true,
				OnRamps:             [][]byte{common.LeftPadBytes(common.HexToAddress(remote.onRamp).Bytes(), 32)},
			}},
		},
	)
	require.NoError(t, err)

	downgradeInput := buildConfigureChainForLanesInput(localWithTest.deployedContracts, chainSelector, remote, remoteSelector)
	downgradeInput.Router = localWithTest.testRouter
	downgradeInput.AllowOnrampOverride = false

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		sequences.ConfigureChainForLanes, e.BlockChains, downgradeInput,
	)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "OffRamp")
	assert.Contains(t, err.Error(), "refusing to overwrite")

	b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	offRampCfg, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector, Address: common.HexToAddress(localWithTest.offRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, prodRouter, offRampCfg.Output.Router,
		"OffRamp router should still point to the prod router after rejected downgrade")
}

