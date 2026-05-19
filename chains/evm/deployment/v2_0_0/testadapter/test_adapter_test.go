package testadapter

import (
	"errors"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/message_hasher"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"

	op_router "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	changesetadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

func TestInit_RegistersForkCCIPSendAndFamilyAdapters(t *testing.T) {
	version := semver.MustParse("2.0.0")
	registry := testadapters.GetTestAdapterRegistry()

	_, foundForkAdapter := registry.GetForkCCIPSendTestAdapter(chain_selectors.FamilyEVM, version)
	require.True(t, foundForkAdapter)

	_, foundFamilyAdapter := registry.GetTestAdapterForFamily(chain_selectors.FamilyEVM, version)
	require.True(t, foundFamilyAdapter)
}

func TestEVMAdapter_BuildMessage(t *testing.T) {
	adapter := &EVMAdapter{}

	feeToken := "0x0000000000000000000000000000000000000001"
	input := testadapters.MessageComponents{
		Receiver: []byte{0x12, 0x34},
		Data:     []byte("hello"),
		FeeToken: feeToken,
		ExtraArgs: []byte{
			0xaa, 0xbb,
		},
		TokenAmounts: []testadapters.TokenAmount{
			{
				Token:  "0x0000000000000000000000000000000000000002",
				Amount: big.NewInt(42),
			},
		},
	}

	msgAny, err := adapter.BuildMessage(input)
	require.NoError(t, err)

	msg, ok := msgAny.(router.ClientEVM2AnyMessage)
	require.True(t, ok)
	require.Equal(t, common.LeftPadBytes(input.Receiver, 32), msg.Receiver)
	require.Equal(t, input.Data, msg.Data)
	require.Equal(t, common.HexToAddress(feeToken), msg.FeeToken)
	require.Equal(t, input.ExtraArgs, msg.ExtraArgs)
	require.Len(t, msg.TokenAmounts, 1)
	require.Equal(t, common.HexToAddress(input.TokenAmounts[0].Token), msg.TokenAmounts[0].Token)
	require.Zero(t, input.TokenAmounts[0].Amount.Cmp(msg.TokenAmounts[0].Amount))
}

func TestEVMAdapter_BuildMessage_InvalidTokenAddress(t *testing.T) {
	adapter := &EVMAdapter{}

	_, err := adapter.BuildMessage(testadapters.MessageComponents{
		TokenAmounts: []testadapters.TokenAmount{
			{
				Token:  "not-an-address",
				Amount: big.NewInt(1),
			},
		},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "invalid token address")
}

type laneContracts struct {
	router                    string
	onRamp                    string
	feeQuoter                 string
	offRamp                   string
	committeeVerifier         string
	committeeVerifierResolver string
	executor                  string
	addresses                 []datastore.AddressRef
}

// evmFamilySelector is bytes4(keccak256("CCIP ChainFamilySelector EVM")) = 0x2812d52c.
var evmFamilySelector = [4]byte{0x28, 0x12, 0xd5, 0x2c}

func boolPtr(v bool) *bool { return &v }

func deployLaneContracts(t *testing.T, env *deployment.Environment, chain cldf_evm.Chain, chainSelector uint64) laneContracts {
	t.Helper()

	bundle := testsetup.BundleWithFreshReporter(env.OperationsBundle)
	create2FactoryRef, err := contract.MaybeDeployContract(bundle, create2_factory.Deploy, chain, contract.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  chainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{chain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err)

	deploymentReport, err := operations.ExecuteSequence(
		bundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:    chainSelector,
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ContractParams:   testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err)

	out := laneContracts{addresses: deploymentReport.Output.Addresses}
	for _, addr := range deploymentReport.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(op_router.ContractType):
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

	require.NotEmpty(t, out.router)
	require.NotEmpty(t, out.onRamp)
	require.NotEmpty(t, out.feeQuoter)
	require.NotEmpty(t, out.offRamp)
	require.NotEmpty(t, out.committeeVerifier)
	require.NotEmpty(t, out.committeeVerifierResolver)
	require.NotEmpty(t, out.executor)

	return out
}

func buildConfigureChainForLanesInput(
	local laneContracts,
	localSelector uint64,
	remote laneContracts,
	remoteSelector uint64,
) changesetadapters.ConfigureChainForLanesInput {
	return changesetadapters.ConfigureChainForLanesInput{
		ChainSelector: localSelector,
		Router:        common.HexToAddress(local.router).Bytes(),
		OnRamp:        common.HexToAddress(local.onRamp).Bytes(),
		FeeQuoter:     common.HexToAddress(local.feeQuoter).Bytes(),
		OffRamp:       common.HexToAddress(local.offRamp).Bytes(),
		CommitteeVerifiers: []changesetadapters.CommitteeVerifierConfig[datastore.AddressRef]{
			{
				CommitteeVerifier: []datastore.AddressRef{
					{
						Address: local.committeeVerifier,
						Type:    datastore.ContractType(committee_verifier.ContractType),
						Version: committee_verifier.Version,
					},
					{
						Address: local.committeeVerifierResolver,
						Type:    datastore.ContractType(sequences.CommitteeVerifierResolverType),
						Version: semver.MustParse("2.0.0"),
					},
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
					LinkFeeMultiplierPercent:    90,
					USDPerUnitGas:               big.NewInt(20_000_000_000_000),
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

func TestEVMForkCCIPSendAdapter_SendMessageAfterLaneConfiguration(t *testing.T) {
	sourceSelector := uint64(5009297550715157269)
	destSelector := uint64(4356164186791070119)

	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{sourceSelector, destSelector}),
	)
	require.NoError(t, err)

	sourceChain := env.BlockChains.EVMChains()[sourceSelector]
	destChain := env.BlockChains.EVMChains()[destSelector]

	sourceContracts := deployLaneContracts(t, env, sourceChain, sourceSelector)
	destContracts := deployLaneContracts(t, env, destChain, destSelector)

	sourceInput := buildConfigureChainForLanesInput(sourceContracts, sourceSelector, destContracts, destSelector)
	destInput := buildConfigureChainForLanesInput(destContracts, destSelector, sourceContracts, sourceSelector)

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(env.OperationsBundle),
		sequences.ConfigureChainForLanes,
		env.BlockChains,
		sourceInput,
	)
	require.NoError(t, err)

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(env.OperationsBundle),
		sequences.ConfigureChainForLanes,
		env.BlockChains,
		destInput,
	)
	require.NoError(t, err)

	updatedDS := datastore.NewMemoryDataStore()
	require.NoError(t, updatedDS.Merge(env.DataStore))
	for _, ref := range append(sourceContracts.addresses, destContracts.addresses...) {
		addErr := updatedDS.Addresses().Add(ref)
		require.True(t, addErr == nil || errors.Is(addErr, datastore.ErrAddressRefExists))
	}
	env.DataStore = updatedDS.Seal()

	sourceAdapter := NewEVMForkCCIPSendTestAdapter(env, sourceSelector)

	_, tx, msgHasher, err := message_hasher.DeployMessageHasher(sourceChain.DeployerKey, sourceChain.Client)
	require.NoError(t, err)
	_, err = sourceChain.Confirm(tx)
	require.NoError(t, err)

	extraArgs, err := msgHasher.EncodeGenericExtraArgsV3(
		&bind.CallOpts{Context: t.Context()},
		message_hasher.ExtraArgsCodecGenericExtraArgsV3{
			GasLimit:                80_000,
			RequestedFinalityConfig: finality.RawWaitForFinality,
			Ccvs:                    []common.Address{common.HexToAddress(sourceContracts.committeeVerifierResolver)},
			CcvArgs:                 [][]byte{{}},
			Executor:                common.HexToAddress(sourceContracts.executor),
			ExecutorArgs:            []byte{},
			TokenReceiver:           []byte{},
			TokenArgs:               []byte{},
		},
	)
	require.NoError(t, err)

	receiver := common.LeftPadBytes(destChain.DeployerKey.From.Bytes(), 32)
	msg, err := sourceAdapter.BuildMessage(testadapters.MessageComponents{
		DestChainSelector: destSelector,
		Receiver:          receiver,
		Data:              []byte("adapter integration send"),
		ExtraArgs:         extraArgs,
	})
	require.NoError(t, err)

	_, msgID, err := sourceAdapter.SendMessage(t.Context(), destSelector, msg)
	require.NoError(t, err)
	require.NotEmpty(t, msgID)
}
