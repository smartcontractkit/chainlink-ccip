package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/message_hasher"
	versioned_verifier_resolver_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	onramp2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
)

func TestConfigureLaneLegAsSourceAndDest(t *testing.T) {
	tests := []struct {
		desc string
	}{
		{
			desc: "valid input",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSelector := uint64(5009297550715157269)
			remoteChainSelector := uint64(4356164186791070119)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")
			evmChain := e.BlockChains.EVMChains()[chainSelector]
			evmChain2 := e.BlockChains.EVMChains()[remoteChainSelector]

			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
				ChainSelector:  chainSelector,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{evmChain.DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")

			deploymentReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				evmChain,
				sequences.DeployChainContractsInput{
					ChainSelector:    chainSelector,
					CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
					ContractParams:   testsetup.CreateBasicContractParams(),
					DeployerKeyOwned: true,
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			var routerAddress string
			var onRampAddr string
			var feeQuoterAddr string
			var offRampAddr string
			var committeeVerifierAddr string
			var committeeVerifierResolverAddr string
			var executorAddr string
			for _, addr := range deploymentReport.Output.Addresses {
				switch addr.Type {
				case datastore.ContractType(router.ContractType):
					routerAddress = addr.Address
				case datastore.ContractType(onramp.ContractType):
					onRampAddr = addr.Address
				case datastore.ContractType(fee_quoter.ContractType):
					feeQuoterAddr = addr.Address
				case datastore.ContractType(offramp.ContractType):
					offRampAddr = addr.Address
				case datastore.ContractType(committee_verifier.ContractType):
					committeeVerifierAddr = addr.Address
				case datastore.ContractType(sequences.ExecutorProxyType):
					executorAddr = addr.Address
				case datastore.ContractType(sequences.CommitteeVerifierResolverType):
					committeeVerifierResolverAddr = addr.Address
				}
			}

			create2FactoryRef, err = contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain2, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
				ChainSelector:  remoteChainSelector,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{evmChain2.DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")

			deploymentReport, err = operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				evmChain2,
				sequences.DeployChainContractsInput{
					ChainSelector:    remoteChainSelector,
					CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
					ContractParams:   testsetup.CreateBasicContractParams(),
					DeployerKeyOwned: true,
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			var remoteRouterAddress string
			var remoteOnRampAddr string
			var remoteFeeQuoterAddr string
			var remoteOffRampAddr string
			var remoteCommitteeVerifierAddr string
			var remoteCommitteeVerifierResolverAddr string
			var remoteExecutorAddr string
			for _, addr := range deploymentReport.Output.Addresses {
				switch addr.Type {
				case datastore.ContractType(router.ContractType):
					remoteRouterAddress = addr.Address
				case datastore.ContractType(onramp.ContractType):
					remoteOnRampAddr = addr.Address
				case datastore.ContractType(fee_quoter.ContractType):
					remoteFeeQuoterAddr = addr.Address
				case datastore.ContractType(offramp.ContractType):
					remoteOffRampAddr = addr.Address
				case datastore.ContractType(committee_verifier.ContractType):
					remoteCommitteeVerifierAddr = addr.Address
				case datastore.ContractType(sequences.ExecutorProxyType):
					remoteExecutorAddr = addr.Address
				case datastore.ContractType(sequences.CommitteeVerifierResolverType):
					remoteCommitteeVerifierResolverAddr = addr.Address
				}
			}

			committeeVerifiers := []lanes.CommitteeVerifierConfig[datastore.AddressRef]{
				{
					CommitteeVerifier: []datastore.AddressRef{
						{
							Address: committeeVerifierAddr,
							Type:    datastore.ContractType(committee_verifier.ContractType),
							Version: committee_verifier.Version,
						},
						{
							Address: committeeVerifierResolverAddr,
							Type:    datastore.ContractType(sequences.CommitteeVerifierResolverType),
							Version: committee_verifier.Version,
						},
					},
					RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainConfig{
						remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
					},
				},
			}

			remoteCommitteeVerifiers := []lanes.CommitteeVerifierConfig[datastore.AddressRef]{
				{
					CommitteeVerifier: []datastore.AddressRef{
						{
							Address: remoteCommitteeVerifierAddr,
							Type:    datastore.ContractType(committee_verifier.ContractType),
							Version: committee_verifier.Version,
						},
						{
							Address: remoteCommitteeVerifierResolverAddr,
							Type:    datastore.ContractType(sequences.CommitteeVerifierResolverType),
							Version: committee_verifier.Version,
						},
					},
					RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainConfig{
						chainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
					},
				},
			}

			localChainDef := &lanes.ChainDefinition{
				Selector:           chainSelector,
				Router:             common.HexToAddress(routerAddress).Bytes(),
				OnRamp:             common.HexToAddress(onRampAddr).Bytes(),
				FeeQuoter:          common.HexToAddress(feeQuoterAddr).Bytes(),
				OffRamp:            common.HexToAddress(offRampAddr).Bytes(),
				CommitteeVerifiers: committeeVerifiers,
				DefaultInboundCCVs: []datastore.AddressRef{{
					Address: committeeVerifierAddr,
				}},
				DefaultOutboundCCVs: []datastore.AddressRef{{
					Address: committeeVerifierAddr,
				}},
				DefaultExecutor: datastore.AddressRef{
					Address: executorAddr,
				},
			}

			remoteChainDef := &lanes.ChainDefinition{
				Selector:           remoteChainSelector,
				Router:             common.HexToAddress(remoteRouterAddress).Bytes(),
				OnRamp:             common.HexToAddress(remoteOnRampAddr).Bytes(),
				FeeQuoter:          common.HexToAddress(remoteFeeQuoterAddr).Bytes(),
				OffRamp:            common.HexToAddress(remoteOffRampAddr).Bytes(),
				CommitteeVerifiers: remoteCommitteeVerifiers,
				DefaultOutboundCCVs: []datastore.AddressRef{
					{Address: remoteCommitteeVerifierAddr},
				},
				DefaultInboundCCVs: []datastore.AddressRef{
					{Address: remoteCommitteeVerifierAddr},
				},
				DefaultExecutor:          datastore.AddressRef{Address: remoteExecutorAddr},
				FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
				ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
				AddressBytesLength:       20,
				BaseExecutionGasCost:     80_000,
			}

			freshBundle := testsetup.BundleWithFreshReporter(e.OperationsBundle)

			// Configure lane leg as source: sets up OnRamp, FeeQuoter, Executor, Router onRamp, CommitteeVerifier
			_, err = operations.ExecuteSequence(
				freshBundle,
				sequences.ConfigureLaneLegAsSource,
				e.BlockChains,
				lanes.UpdateLanesInput{
					Source: localChainDef,
					Dest:   remoteChainDef,
				},
			)
			require.NoError(t, err, "ConfigureLaneLegAsSource should not error")

			// Configure lane leg as dest: sets up OffRamp, Router offRamp, CommitteeVerifier
			_, err = operations.ExecuteSequence(
				testsetup.BundleWithFreshReporter(e.OperationsBundle),
				sequences.ConfigureLaneLegAsDest,
				e.BlockChains,
				lanes.UpdateLanesInput{
					Source: remoteChainDef,
					Dest:   localChainDef,
				},
			)
			require.NoError(t, err, "ConfigureLaneLegAsDest should not error")

			// Check onRamps on router
			onRampOnRouter, err := operations.ExecuteOperation(e.OperationsBundle, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(routerAddress),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, onRampAddr, onRampOnRouter.Output.Hex(), "OnRamp address on router should match OnRamp address")

			// Check offRamps on router
			offRampsOnRouter, err := operations.ExecuteOperation(e.OperationsBundle, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(routerAddress),
				Args:          nil,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Len(t, offRampsOnRouter.Output, 1, "There should be one OffRamp on the router for the remote chain")
			require.Equal(t, offRampAddr, offRampsOnRouter.Output[0].OffRamp.Hex(), "OffRamp address on router should match OffRamp address")

			// Check sourceChainConfig on OffRamp
			sourceChainConfig, err := operations.ExecuteOperation(e.OperationsBundle, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(offRampAddr),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, common.LeftPadBytes(remoteChainDef.OnRamp, 32), sourceChainConfig.Output.OnRamps[0], "OnRamp in source chain config should match OnRamp address")
			require.Len(t, sourceChainConfig.Output.DefaultCCVs, 1, "There should be one DefaultCCV in source chain config")
			require.Equal(t, committeeVerifierAddr, sourceChainConfig.Output.DefaultCCVs[0].Hex(), "DefaultCCV in source chain config should match CommitteeVerifier address")
			require.True(t, sourceChainConfig.Output.IsEnabled, "IsEnabled in source chain config should be true")
			require.Equal(t, routerAddress, sourceChainConfig.Output.Router.Hex(), "Router in source chain config should match Router address")

			// Check destChainConfig on OnRamp
			destChainConfig, err := operations.ExecuteOperation(e.OperationsBundle, onramp.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(onRampAddr),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, routerAddress, destChainConfig.Output.Router.Hex(), "Router in dest chain config should match Router address")
			require.Equal(t, remoteChainDef.OffRamp, destChainConfig.Output.OffRamp, "OffRamp in dest chain config should match CCIPMessageDest")
			require.Equal(t, executorAddr, destChainConfig.Output.DefaultExecutor.Hex(), "DefaultExecutor in dest chain config should match configured DefaultExecutor")
			require.Len(t, destChainConfig.Output.DefaultCCVs, 1, "There should be one DefaultCCV in dest chain config")
			require.Equal(t, committeeVerifierAddr, destChainConfig.Output.DefaultCCVs[0].Hex(), "DefaultCCV in dest chain config should match CommitteeVerifier address")

			// Check destChainConfig on CommitteeVerifier
			committeeVerifierRemoteChainConfig, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.GetRemoteChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(committeeVerifierAddr),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, routerAddress, committeeVerifierRemoteChainConfig.Output.Router.Hex(), "Router in CommitteeVerifier remote chain config should match Router address")
			require.False(t, committeeVerifierRemoteChainConfig.Output.AllowlistEnabled, "AllowlistEnabled in CommitteeVerifier remote chain config should be false")

			// Check signature quorum on CommitteeVerifier
			signatureQuorumReport, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.GetSignatureConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(committeeVerifierAddr),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, uint8(1), signatureQuorumReport.Output.Threshold, "Threshold in CommitteeVerifier signature config should be 1")
			require.Equal(t, []common.Address{common.HexToAddress("0x01")}, signatureQuorumReport.Output.Signers, "Signers in CommitteeVerifier signature config should match")

			// Check outbound implementation on CommitteeVerifierResolver
			boundResolver, err := versioned_verifier_resolver_latest.NewVersionedVerifierResolver(common.HexToAddress(committeeVerifierResolverAddr), evmChain.Client)
			require.NoError(t, err, "Failed to instantiate VersionedVerifierResolver")
			outboundImpl, err := boundResolver.GetOutboundImplementation(&bind.CallOpts{Context: t.Context()}, remoteChainSelector, []byte{})
			require.NoError(t, err, "GetOutboundImplementation should not error")
			require.Equal(t, committeeVerifierAddr, outboundImpl.Hex(), "Outbound implementation verifier on CommitteeVerifierResolver should match CommitteeVerifier address")

			// Check inbound implementation on CommitteeVerifierResolver
			versionTagReport, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.VersionTag, evmChain, contract.FunctionInput[struct{}]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(committeeVerifierAddr),
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			inboundImpl, err := boundResolver.GetInboundImplementation(&bind.CallOpts{Context: t.Context()}, versionTagReport.Output[:])
			require.NoError(t, err, "GetInboundImplementationForVersion should not error")
			require.Equal(t, committeeVerifierAddr, inboundImpl.Hex(), "Inbound implementation verifier on CommitteeVerifierResolver should match CommitteeVerifier address")

			// Check dest chains on Executor
			executorDestChains, err := operations.ExecuteOperation(e.OperationsBundle, executor.GetDestChains, evmChain, contract.FunctionInput[struct{}]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(executorAddr),
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Len(t, executorDestChains.Output, 1, "There should be one dest chain on Executor")
			expectedExecConfig := testsetup.CreateBasicExecutorDestChainConfig()
			gotExecConfig := executorDestChains.Output[0].Config
			require.Equal(t, remoteChainSelector, executorDestChains.Output[0].DestChainSelector, "Dest chain selector on Executor should match remote chain selector")
			require.Equal(t, expectedExecConfig.USDCentsFee, gotExecConfig.UsdCentsFee, "UsdCentsFee in Executor dest chain config should match")
			require.True(t, gotExecConfig.Enabled, "Dest chain selector on Executor should be enabled")

			/////////////////////////////////////////
			// Try sending CCIP message /////////////
			/////////////////////////////////////////

			_, tx, msgHasher, err := message_hasher.DeployMessageHasher(evmChain.DeployerKey, evmChain.Client)
			require.NoError(t, err, "Failed to deploy MessageHasher")
			_, err = evmChain.Confirm(tx)
			require.NoError(t, err, "Failed to confirm MessageHasher deployment")

			extraArgs, err := msgHasher.EncodeGenericExtraArgsV3(
				&bind.CallOpts{Context: t.Context()},
				message_hasher.ExtraArgsCodecGenericExtraArgsV3{
					GasLimit:           80_000,
					BlockConfirmations: 0,
					Ccvs:               []common.Address{common.HexToAddress(committeeVerifierResolverAddr)},
					CcvArgs:            [][]byte{{}},
					Executor:           common.HexToAddress(executorAddr),
					ExecutorArgs:       []byte{},
					TokenReceiver:      []byte{},
					TokenArgs:          []byte{},
				},
			)
			require.NoError(t, err, "EncodeGenericExtraArgsV3 should not error")

			ccipSendArgs := router.CCIPSendArgs{
				DestChainSelector: remoteChainSelector,
				EVM2AnyMessage: router.EVM2AnyMessage{
					Receiver:     common.LeftPadBytes(evmChain.DeployerKey.From.Bytes(), 32),
					Data:         []byte{},
					TokenAmounts: []router.EVMTokenAmount{},
					ExtraArgs:    extraArgs,
				},
			}

			fee, err := operations.ExecuteOperation(e.OperationsBundle, router.GetFee, evmChain, contract.FunctionInput[router.CCIPSendArgs]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(routerAddress),
				Args:          ccipSendArgs,
			})
			require.NoError(t, err, "ExecuteOperation should not error")

			ccipSendArgs.Value = fee.Output
			_, err = operations.ExecuteOperation(e.OperationsBundle, router.CCIPSend, evmChain, contract.FunctionInput[router.CCIPSendArgs]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(routerAddress),
				Args:          ccipSendArgs,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
		})
	}
}

// TestMaybeAddRouterOnRampsAddsConfigArg verifies the expected output of MaybeAddRouterOnRampsAddsConfigArg
// for TestRouter true and TestRouter false when the router has no onRamp for dest, and after ConfigureLaneLegAsSource.
func TestMaybeAddRouterOnRampsAddsConfigArg(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	remoteChainSelector := uint64(4356164186791070119)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector, remoteChainSelector}),
	)
	require.NoError(t, err)
	evmChain := e.BlockChains.EVMChains()[chainSelector]
	evmChain2 := e.BlockChains.EVMChains()[remoteChainSelector]

	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  chainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{evmChain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err)

	deploymentReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		evmChain,
		sequences.DeployChainContractsInput{
			ChainSelector:    chainSelector,
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ContractParams:   testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
			DeployTestRouter: true,
		},
	)
	require.NoError(t, err)

	var testRouterAddress, routerAddress, onRampAddr, feeQuoterAddr, offRampAddr, committeeVerifierAddr, committeeVerifierResolverAddr, executorAddr string
	for _, addr := range deploymentReport.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(router.ContractType):
			routerAddress = addr.Address
		case datastore.ContractType(router.TestRouterContractType):
			testRouterAddress = addr.Address
		case datastore.ContractType(onramp.ContractType):
			onRampAddr = addr.Address
		case datastore.ContractType(fee_quoter.ContractType):
			feeQuoterAddr = addr.Address
		case datastore.ContractType(offramp.ContractType):
			offRampAddr = addr.Address
		case datastore.ContractType(committee_verifier.ContractType):
			committeeVerifierAddr = addr.Address
		case datastore.ContractType(sequences.ExecutorProxyType):
			executorAddr = addr.Address
		case datastore.ContractType(sequences.CommitteeVerifierResolverType):
			committeeVerifierResolverAddr = addr.Address
		}
	}

	create2FactoryRef, err = contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain2, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  remoteChainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{evmChain2.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err)

	deploymentReport, err = operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		evmChain2,
		sequences.DeployChainContractsInput{
			ChainSelector:    remoteChainSelector,
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ContractParams:   testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
			DeployTestRouter: true,
		},
	)
	require.NoError(t, err)

	var remoteTestRouterAddr, remoteRouterAddress, remoteOnRampAddr, remoteFeeQuoterAddr, remoteOffRampAddr, remoteCommitteeVerifierAddr, remoteCommitteeVerifierResolverAddr, remoteExecutorAddr string
	for _, addr := range deploymentReport.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(router.ContractType):
			remoteRouterAddress = addr.Address
		case datastore.ContractType(router.TestRouterContractType):
			remoteTestRouterAddr = addr.Address
		case datastore.ContractType(onramp.ContractType):
			remoteOnRampAddr = addr.Address
		case datastore.ContractType(fee_quoter.ContractType):
			remoteFeeQuoterAddr = addr.Address
		case datastore.ContractType(offramp.ContractType):
			remoteOffRampAddr = addr.Address
		case datastore.ContractType(committee_verifier.ContractType):
			remoteCommitteeVerifierAddr = addr.Address
		case datastore.ContractType(sequences.ExecutorProxyType):
			remoteExecutorAddr = addr.Address
		case datastore.ContractType(sequences.CommitteeVerifierResolverType):
			remoteCommitteeVerifierResolverAddr = addr.Address
		}
	}

	committeeVerifiers := []lanes.CommitteeVerifierConfig[datastore.AddressRef]{
		{
			CommitteeVerifier: []datastore.AddressRef{
				{Address: committeeVerifierAddr, Type: datastore.ContractType(committee_verifier.ContractType), Version: committee_verifier.Version},
				{Address: committeeVerifierResolverAddr, Type: datastore.ContractType(sequences.CommitteeVerifierResolverType), Version: committee_verifier.Version},
			},
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainConfig{
				remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
			},
		},
	}
	remoteCommitteeVerifiers := []lanes.CommitteeVerifierConfig[datastore.AddressRef]{
		{
			CommitteeVerifier: []datastore.AddressRef{
				{Address: remoteCommitteeVerifierAddr, Type: datastore.ContractType(committee_verifier.ContractType), Version: committee_verifier.Version},
				{Address: remoteCommitteeVerifierResolverAddr, Type: datastore.ContractType(sequences.CommitteeVerifierResolverType), Version: committee_verifier.Version},
			},
			RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainConfig{
				chainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
			},
		},
	}

	localChainDef := &lanes.ChainDefinition{
		Selector:            chainSelector,
		Router:              common.HexToAddress(routerAddress).Bytes(),
		OnRamp:              common.HexToAddress(onRampAddr).Bytes(),
		FeeQuoter:           common.HexToAddress(feeQuoterAddr).Bytes(),
		OffRamp:             common.HexToAddress(offRampAddr).Bytes(),
		CommitteeVerifiers:  committeeVerifiers,
		DefaultInboundCCVs:  []datastore.AddressRef{{Address: committeeVerifierAddr}},
		DefaultOutboundCCVs: []datastore.AddressRef{{Address: committeeVerifierAddr}},
		DefaultExecutor:     datastore.AddressRef{Address: executorAddr},
	}
	remoteChainDef := &lanes.ChainDefinition{
		Selector:                 remoteChainSelector,
		Router:                   common.HexToAddress(remoteRouterAddress).Bytes(),
		OnRamp:                   common.HexToAddress(remoteOnRampAddr).Bytes(),
		FeeQuoter:                common.HexToAddress(remoteFeeQuoterAddr).Bytes(),
		OffRamp:                  common.HexToAddress(remoteOffRampAddr).Bytes(),
		CommitteeVerifiers:       remoteCommitteeVerifiers,
		DefaultOutboundCCVs:      []datastore.AddressRef{{Address: remoteCommitteeVerifierAddr}},
		DefaultInboundCCVs:       []datastore.AddressRef{{Address: remoteCommitteeVerifierAddr}},
		DefaultExecutor:          datastore.AddressRef{Address: remoteExecutorAddr},
		FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
		ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
		AddressBytesLength:       20,
		BaseExecutionGasCost:     80_000,
	}

	input := lanes.UpdateLanesInput{
		Source: localChainDef,
		Dest:   remoteChainDef,
	}
	t.Run("TestRouter_true_router_has_no_onRamp", func(t *testing.T) {
		input.TestRouter = true
		localChainDef.Router = common.HexToAddress(testRouterAddress).Bytes()
		remoteChainDef.Router = common.HexToAddress(remoteTestRouterAddr).Bytes()
		b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
		onRampAdds, err := sequences.MaybeAddRouterOnRampsAddsConfigArg(b, evmChain, input)
		require.NoError(t, err)
		require.Len(t, onRampAdds, 1, "when router has no onRamp for dest, TestRouter true should return one OnRamp to add")
		require.Equal(t, remoteChainSelector, onRampAdds[0].DestChainSelector)
		require.Equal(t, common.HexToAddress(onRampAddr), onRampAdds[0].OnRamp)
	})

	t.Run("TestRouter_false_router_has_no_onRamp", func(t *testing.T) {
		b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
		input.TestRouter = false
		localChainDef.Router = common.HexToAddress(routerAddress).Bytes()
		remoteChainDef.Router = common.HexToAddress(remoteRouterAddress).Bytes()
		onRampAdds, err := sequences.MaybeAddRouterOnRampsAddsConfigArg(b, evmChain, input)
		require.NoError(t, err)
		require.Len(t, onRampAdds, 1, "when router has no onRamp for dest, TestRouter false should return one OnRamp to add")
		require.Equal(t, remoteChainSelector, onRampAdds[0].DestChainSelector)
		require.Equal(t, common.HexToAddress(onRampAddr), onRampAdds[0].OnRamp)
	})

	t.Run("router_has_1.6_onRamp", func(t *testing.T) {
		b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
		input.TestRouter = false
		localChainDef.Router = common.HexToAddress(routerAddress).Bytes()
		remoteChainDef.Router = common.HexToAddress(remoteRouterAddress).Bytes()
		// add an onramp with version 1.6.0
		// deploy onramp 1.6.0
		out, err := operations.ExecuteOperation(testsetup.BundleWithFreshReporter(e.OperationsBundle), onramp2.Deploy, evmChain, contract_utils.DeployInput[onramp2.ConstructorArgs]{
			ChainSelector:  chainSelector,
			TypeAndVersion: onramp2.TypeAndVersion,
			Args: onramp2.ConstructorArgs{
				StaticConfig: onramp2.StaticConfig{
					ChainSelector:      chainSelector,
					RmnRemote:          utils.RandomAddress(),
					NonceManager:       utils.RandomAddress(),
					TokenAdminRegistry: utils.RandomAddress(),
				},
				DynamicConfig: onramp2.DynamicConfig{
					FeeQuoter:     utils.RandomAddress(),
					FeeAggregator: utils.RandomAddress(),
				},
			},
		})
		require.NoError(t, err)
		deployedOnRampAddr := common.HexToAddress(out.Output.Address)
		// add this onramp to the router
		_, err = operations.ExecuteOperation(testsetup.BundleWithFreshReporter(e.OperationsBundle), router.ApplyRampUpdates, evmChain, contract_utils.FunctionInput[router.ApplyRampsUpdatesArgs]{
			ChainSelector: chainSelector,
			Address:       common.HexToAddress(routerAddress),
			Args: router.ApplyRampsUpdatesArgs{
				OnRampUpdates: []router.OnRamp{
					{
						DestChainSelector: remoteChainSelector,
						OnRamp:            deployedOnRampAddr,
					},
				},
			},
		})
		require.NoError(t, err)
		b = testsetup.BundleWithFreshReporter(e.OperationsBundle)
		onRampAdds, err := sequences.MaybeAddRouterOnRampsAddsConfigArg(b, evmChain, input)
		require.NoError(t, err)
		require.Empty(t, onRampAdds, "when router has prev version onRamp for dest, TestRouter false should return no adds")
		input.TestRouter = true
		localChainDef.Router = common.HexToAddress(testRouterAddress).Bytes()
		remoteChainDef.Router = common.HexToAddress(remoteTestRouterAddr).Bytes()

		_, err = operations.ExecuteOperation(testsetup.BundleWithFreshReporter(e.OperationsBundle), router.ApplyRampUpdates, evmChain, contract_utils.FunctionInput[router.ApplyRampsUpdatesArgs]{
			ChainSelector: chainSelector,
			Address:       common.HexToAddress(testRouterAddress),
			Args: router.ApplyRampsUpdatesArgs{
				OnRampUpdates: []router.OnRamp{
					{
						DestChainSelector: remoteChainSelector,
						OnRamp:            deployedOnRampAddr,
					},
				},
			},
		})
		onRampAdds, err = sequences.MaybeAddRouterOnRampsAddsConfigArg(b, evmChain, input)
		require.NoError(t, err)
		require.Len(t, onRampAdds, 1, "when router has prev version onRamp for dest, TestRouter false should return one OnRamp to add")
		require.Equal(t, remoteChainSelector, onRampAdds[0].DestChainSelector)
		require.Equal(t, common.HexToAddress(onRampAddr), onRampAdds[0].OnRamp)
	})

	t.Run("TestRouter_true_router_already_has_onRamp", func(t *testing.T) {
		b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
		input.TestRouter = true
		localChainDef.Router = common.HexToAddress(testRouterAddress).Bytes()
		remoteChainDef.Router = common.HexToAddress(remoteTestRouterAddr).Bytes()
		_, err = operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			sequences.ConfigureLaneLegAsSource,
			e.BlockChains,
			lanes.UpdateLanesInput{
				Source:     localChainDef,
				Dest:       remoteChainDef,
				TestRouter: true,
			},
		)
		require.NoError(t, err)
		onRampAdds, err := sequences.MaybeAddRouterOnRampsAddsConfigArg(b, evmChain, input)
		require.NoError(t, err)
		require.Empty(t, onRampAdds, "when router already has same onRamp for dest, TestRouter true should return no adds")
	})

	t.Run("TestRouter_false_router_already_has_onRamp", func(t *testing.T) {
		b := testsetup.BundleWithFreshReporter(e.OperationsBundle)
		input.TestRouter = false
		localChainDef.Router = common.HexToAddress(routerAddress).Bytes()
		remoteChainDef.Router = common.HexToAddress(remoteRouterAddress).Bytes()
		_, err = operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			sequences.ConfigureLaneLegAsSource,
			e.BlockChains,
			lanes.UpdateLanesInput{
				Source: localChainDef,
				Dest:   remoteChainDef,
			},
		)
		require.NoError(t, err)
		onRampAdds, err := sequences.MaybeAddRouterOnRampsAddsConfigArg(b, evmChain, input)
		require.NoError(t, err)
		require.Empty(t, onRampAdds, "when router already has same onRamp for dest, TestRouter false should return no adds")
	})
}
