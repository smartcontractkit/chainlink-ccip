package sequences_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/message_hasher"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestConfigureChainForLanes(t *testing.T) {
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
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")
			evmChain := e.BlockChains.EVMChains()[chainSelector]

			deploymentReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				evmChain,
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSelector,
					ContractParams: testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			var routerAddress string
			var onRamp string
			var feeQuoter string
			var offRamp string
			var committeeVerifier string
			var committeeVerifierResolver string
			var executorAddress string
			var linkTokenAddress string
			for _, addr := range deploymentReport.Output.Addresses {
				switch addr.Type {
				case datastore.ContractType(router.ContractType):
					routerAddress = addr.Address
				case datastore.ContractType(onramp.ContractType):
					onRamp = addr.Address
				case datastore.ContractType(fee_quoter.ContractType):
					feeQuoter = addr.Address
				case datastore.ContractType(offramp.ContractType):
					offRamp = addr.Address
				case datastore.ContractType(committee_verifier.ContractType):
					committeeVerifier = addr.Address
				case datastore.ContractType(executor.ContractType):
					executorAddress = addr.Address
				case datastore.ContractType(committee_verifier.ResolverType):
					committeeVerifierResolver = addr.Address
				case datastore.ContractType(link.ContractType):
					linkTokenAddress = addr.Address
				}
			}
			ccipMessageSource := common.HexToAddress("0x10").Bytes()
			ccipMessageDest := common.HexToAddress("0x11").Bytes()
			remoteChainSelector := uint64(4356164186791070119)

			_, err = operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.ConfigureChainForLanes,
				e.BlockChains,
				adapters.ConfigureChainForLanesInput{
					ChainSelector: chainSelector,
					Router:        routerAddress,
					OnRamp:        onRamp,
					CommitteeVerifiers: []adapters.CommitteeVerifier[string]{
						{
							Implementation: committeeVerifier,
							Resolver:       committeeVerifierResolver,
						},
					},
					FeeQuoter: feeQuoter,
					OffRamp:   offRamp,
					RemoteChains: map[uint64]adapters.RemoteChainConfig[[]byte, string]{
						remoteChainSelector: {
							AllowTrafficFrom:                 true,
							OnRamp:                           ccipMessageSource,
							OffRamp:                          ccipMessageDest,
							DefaultInboundCCVs:               []string{committeeVerifier},
							DefaultOutboundCCVs:              []string{committeeVerifier},
							DefaultExecutor:                  executorAddress,
							CommitteeVerifierDestChainConfig: testsetup.CreateBasicCommitteeVerifierDestChainConfig(),
							FeeQuoterDestChainConfig:         testsetup.CreateBasicFeeQuoterDestChainConfig(),
							ExecutorDestChainConfig:          testsetup.CreateBasicExecutorDestChainConfig(),
							AddressBytesLength:               20,
							BaseExecutionGasCost:             80_000,
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Check onRamps on router
			onRampOnRouter, err := operations.ExecuteOperation(e.OperationsBundle, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(routerAddress),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, onRamp, onRampOnRouter.Output.Hex(), "OnRamp address on router should match OnRamp address")

			// Check offRamps on router
			offRampsOnRouter, err := operations.ExecuteOperation(e.OperationsBundle, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(routerAddress),
				Args:          nil,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Len(t, offRampsOnRouter.Output, 1, "There should be one OffRamp on the router for the remote chain")
			require.Equal(t, offRamp, offRampsOnRouter.Output[0].OffRamp.Hex(), "OffRamp address on router should match OffRamp address")

			// Check sourceChainConfig on OffRamp
			sourceChainConfig, err := operations.ExecuteOperation(e.OperationsBundle, offramp.GetSourceChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(offRamp),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, ccipMessageSource, sourceChainConfig.Output.OnRamp, "OnRamp in source chain config should match OnRamp address")
			require.Len(t, sourceChainConfig.Output.DefaultCCVs, 1, "There should be one DefaultCCV in source chain config")
			require.Equal(t, committeeVerifier, sourceChainConfig.Output.DefaultCCVs[0].Hex(), "DefaultCCV in source chain config should match CommitteeVerifier address")
			require.True(t, sourceChainConfig.Output.IsEnabled, "IsEnabled in source chain config should be true")
			require.Equal(t, routerAddress, sourceChainConfig.Output.Router.Hex(), "Router in source chain config should match Router address")

			// Check destChainConfig on OnRamp
			destChainConfig, err := operations.ExecuteOperation(e.OperationsBundle, onramp.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(onRamp),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, routerAddress, destChainConfig.Output.Router.Hex(), "Router in dest chain config should match Router address")
			require.Equal(t, ccipMessageDest, destChainConfig.Output.OffRamp, "OffRamp in dest chain config should match CCIPMessageDest")
			require.Equal(t, executorAddress, destChainConfig.Output.DefaultExecutor.Hex(), "DefaultExecutor in dest chain config should match configured DefaultExecutor")
			require.Len(t, destChainConfig.Output.DefaultCCVs, 1, "There should be one DefaultCCV in dest chain config")
			require.Equal(t, committeeVerifier, destChainConfig.Output.DefaultCCVs[0].Hex(), "DefaultCCV in dest chain config should match CommitteeVerifier address")

			// Check destChainConfig on CommitteeVerifier
			committeeVerifierDestChainConfig, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, routerAddress, committeeVerifierDestChainConfig.Output.Router.Hex(), "Router in CommitteeVerifier dest chain config should match Router address")
			require.False(t, committeeVerifierDestChainConfig.Output.AllowlistEnabled, "AllowlistEnabled in CommitteeVerifier dest chain config should be false")

			// Check outbound implementation on CommitteeVerifierResolver
			boundResolver, err := versioned_verifier_resolver.NewVersionedVerifierResolver(common.HexToAddress(committeeVerifierResolver), evmChain.Client)
			require.NoError(t, err, "Failed to instantiate VersionedVerifierResolver")
			outboundImpl, err := boundResolver.GetOutboundImplementation(&bind.CallOpts{Context: t.Context()}, remoteChainSelector, []byte{})
			require.NoError(t, err, "GetOutboundImplementation should not error")
			require.Equal(t, committeeVerifier, outboundImpl.Hex(), "Outbound implementation verifier on CommitteeVerifierResolver should match CommitteeVerifier address")

			// Check inbound implementation on CommitteeVerifierResolver
			versionTagReport, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.GetVersionTag, evmChain, contract.FunctionInput[any]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			inboundImpl, err := boundResolver.GetInboundImplementation(&bind.CallOpts{Context: t.Context()}, versionTagReport.Output[:])
			require.NoError(t, err, "GetInboundImplementationForVersion should not error")
			require.Equal(t, committeeVerifier, inboundImpl.Hex(), "Inbound implementation verifier on CommitteeVerifierResolver should match CommitteeVerifier address")

			// Check dest chains on Executor
			ExecutorDestChains, err := operations.ExecuteOperation(e.OperationsBundle, executor.GetDestChains, evmChain, contract.FunctionInput[any]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(executorAddress),
				Args:          nil,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Len(t, ExecutorDestChains.Output, 1, "There should be one dest chain on Executor")
			expectedExecConfig := testsetup.CreateBasicExecutorDestChainConfig()
			gotExecConfig := ExecutorDestChains.Output[0].Config
			require.Equal(t, remoteChainSelector, ExecutorDestChains.Output[0].DestChainSelector, "Dest chain selector on Executor should match remote chain selector")
			require.Equal(t, expectedExecConfig.USDCentsFee, gotExecConfig.UsdCentsFee, "UsdCentsFee in Executor dest chain config should match")
			require.True(t, gotExecConfig.Enabled, "Dest chain selector on Executor should be enabled")

			// For this test, there is no offchain component providing prices. We have to initialize a token price for
			// the fee token (assumed to be LINK) so that fee calculation can succeed.
			_, err = operations.ExecuteOperation(e.OperationsBundle, fee_quoter.UpdatePrices, evmChain, contract.FunctionInput[fee_quoter.PriceUpdates]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(feeQuoter),
				Args: fee_quoter.PriceUpdates{
					TokenPriceUpdates: []fee_quoter.TokenPriceUpdate{
						{
							SourceToken: common.HexToAddress(linkTokenAddress),
							UsdPerToken: big.NewInt(1e18), // $1.00 with 18 decimals
						},
					},
					GasPriceUpdates: []fee_quoter.GasPriceUpdate{
						{
							DestChainSelector: remoteChainSelector,
							UsdPerUnitGas:     big.NewInt(1e6),
						},
					},
				},
			})
			require.NoError(t, err, "ExecuteOperation should not error")

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
					Ccvs:               []common.Address{common.HexToAddress(committeeVerifierResolver)},
					CcvArgs:            [][]byte{{}},
					Executor:           common.HexToAddress(executorAddress),
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

			// Send CCIP message with value
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
