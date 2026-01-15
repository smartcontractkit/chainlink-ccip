package sequences_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/message_hasher"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/link_token"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
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

			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
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
					ChainSelector:  chainSelector,
					CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
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
				case datastore.ContractType(executor.ProxyType):
					executorAddress = addr.Address
				case datastore.ContractType(committee_verifier.ResolverType):
					committeeVerifierResolver = addr.Address
				}
			}
			ccipMessageSource := common.HexToAddress("0x10").Bytes()
			ccipMessageDest := common.HexToAddress("0x11").Bytes()
			remoteChainSelector := uint64(4356164186791070119)

			configureReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.ConfigureChainForLanes,
				e.BlockChains,
				adapters.ConfigureChainForLanesInput{
					ChainSelector: chainSelector,
					Router:        routerAddress,
					OnRamp:        onRamp,
					CommitteeVerifiers: []adapters.CommitteeVerifierConfig[datastore.AddressRef]{
						{
							CommitteeVerifier: []datastore.AddressRef{
								{
									Address: committeeVerifier,
									Type:    datastore.ContractType(committee_verifier.ContractType),
									Version: committee_verifier.Version,
								},
								{
									Address: committeeVerifierResolver,
									Type:    datastore.ContractType(committee_verifier.ResolverType),
									Version: committee_verifier.Version,
								},
							},
							RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfig{
								remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
							},
						},
					},
					FeeQuoter: feeQuoter,
					OffRamp:   offRamp,
					RemoteChains: map[uint64]adapters.RemoteChainConfig[[]byte, string]{
						remoteChainSelector: {
							AllowTrafficFrom:         true,
							OnRamps:                  [][]byte{ccipMessageSource},
							OffRamp:                  ccipMessageDest,
							DefaultInboundCCVs:       []string{committeeVerifier},
							DefaultOutboundCCVs:      []string{committeeVerifier},
							DefaultExecutor:          executorAddress,
							FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
							ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
							AddressBytesLength:       20,
							BaseExecutionGasCost:     80_000,
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Test FeeQuoter metadata
			require.NotNil(t, configureReport.Output.Metadata, "Metadata should be set")
			require.NotEmpty(t, configureReport.Output.Metadata.Contracts, "Should have contract metadata")
			feeQuoterMetadataFound := false
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				if contractMeta.Address == feeQuoter && contractMeta.ChainSelector == chainSelector {
					feeQuoterMetadataFound = true
					require.Equal(t, feeQuoter, contractMeta.Address, "FeeQuoter metadata should have correct address")
					require.Equal(t, chainSelector, contractMeta.ChainSelector, "FeeQuoter metadata should have correct chain selector")
					if metaMap, ok := contractMeta.Metadata.(map[string]interface{}); ok {
						require.Equal(t, true, metaMap["configured"], "FeeQuoter metadata should have configured=true")
						require.Equal(t, "fee_quoter_configured", metaMap["test_metadata"], "FeeQuoter metadata should have test_metadata")
					} else {
						t.Fatalf("FeeQuoter metadata should be a map[string]interface{}")
					}
					break
				}
			}
			require.True(t, feeQuoterMetadataFound, "Should have found FeeQuoter metadata")

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
			require.Equal(t, common.LeftPadBytes(ccipMessageSource, 32), sourceChainConfig.Output.OnRamps[0], "OnRamp in source chain config should match OnRamp address")
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
			committeeVerifierRemoteChainConfig, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.GetRemoteChainConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, routerAddress, committeeVerifierRemoteChainConfig.Output.Router.Hex(), "Router in CommitteeVerifier remote chain config should match Router address")
			require.False(t, committeeVerifierRemoteChainConfig.Output.AllowlistEnabled, "AllowlistEnabled in CommitteeVerifier remote chain config should be false")

			// Check signature quorum on CommitteeVerifier
			signatureQuorumReport, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.GetSignatureConfig, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, uint8(1), signatureQuorumReport.Output.Threshold, "Threshold in CommitteeVerifier signature config should be 1")
			require.Equal(t, []common.Address{common.HexToAddress("0x01")}, signatureQuorumReport.Output.Signers, "Signers in CommitteeVerifier signature config should match")
			require.Equal(t, remoteChainSelector, signatureQuorumReport.Output.SourceChainSelector, "Source chain selector in CommitteeVerifier signature config should match remote chain selector")

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

func TestConfigureChainForLanes_Metadata(t *testing.T) {
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

			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chainSelector,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{evmChain.DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")

			mockERC20Ref0, err := contract_utils.MaybeDeployContract(e.OperationsBundle, erc20.Deploy, evmChain, contract_utils.DeployInput[erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chainSelector,
				Args: erc20.ConstructorArgs{
					Name:      "Mock ERC20 0",
					Symbol:    "MOCK 0",
					Decimals:  18,
					MaxSupply: big.NewInt(1000000000000000000),
					PreMint:   big.NewInt(1000000000000000000),
					NewOwner:  evmChain.DeployerKey.From,
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy Mock ERC20 0")

			mockERC20Ref1, err := contract_utils.MaybeDeployContract(e.OperationsBundle, erc20.Deploy, evmChain, contract_utils.DeployInput[erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chainSelector,
				Args: erc20.ConstructorArgs{
					Name:      "Mock ERC20 1",
					Symbol:    "MOCK 1",
					Decimals:  18,
					MaxSupply: big.NewInt(1000000000000000000),
					PreMint:   big.NewInt(1000000000000000000),
					NewOwner:  evmChain.DeployerKey.From,
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy Mock ERC20 1")

			mockERC20Ref2, err := contract_utils.MaybeDeployContract(e.OperationsBundle, erc20.Deploy, evmChain, contract_utils.DeployInput[erc20.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *semver.MustParse("1.0.0")),
				ChainSelector:  chainSelector,
				Args: erc20.ConstructorArgs{
					Name:      "Mock ERC20 2",
					Symbol:    "MOCK 2",
					Decimals:  6,
					MaxSupply: big.NewInt(1000000000000000000),
					PreMint:   big.NewInt(1000000000000000000),
					NewOwner:  evmChain.DeployerKey.From,
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy Mock ERC20 2")

			deploymentReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				evmChain,
				sequences.DeployChainContractsInput{
					ChainSelector:  chainSelector,
					CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
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
				case datastore.ContractType(executor.ProxyType):
					executorAddress = addr.Address
				case datastore.ContractType(committee_verifier.ResolverType):
					committeeVerifierResolver = addr.Address
				}
			}

			// Extract LINK and WETH addresses from deployment report
			var linkAddress string
			var wethAddress string
			for _, addr := range deploymentReport.Output.Addresses {
				switch addr.Type {
				case datastore.ContractType(link_token.ContractType):
					linkAddress = addr.Address
				case datastore.ContractType(weth.ContractType):
					wethAddress = addr.Address
				}
			}
			require.NotEmpty(t, linkAddress, "LINK address should be found")
			require.NotEmpty(t, wethAddress, "WETH address should be found")

			nameReport, err := operations.ExecuteOperation(e.OperationsBundle, erc20.Name, evmChain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(mockERC20Ref0.Address),
				Args:          nil,
			})
			require.NoError(t, err, "Failed to read token name")
			require.Equal(t, "Mock ERC20 0", nameReport.Output, "Token name should match")

			nameReport, err = operations.ExecuteOperation(e.OperationsBundle, erc20.Name, evmChain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(mockERC20Ref1.Address),
				Args:          nil,
			})
			require.NoError(t, err, "Failed to read token name")
			require.Equal(t, "Mock ERC20 1", nameReport.Output, "Token name should match")

			nameReport, err = operations.ExecuteOperation(e.OperationsBundle, erc20.Name, evmChain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(mockERC20Ref2.Address),
				Args:          nil,
			})
			require.NoError(t, err, "Failed to read token name")
			require.Equal(t, "Mock ERC20 2", nameReport.Output, "Token name should match")

			// Price mock tokens in fee quoter, so they are added to it.
			_, err = operations.ExecuteOperation(e.OperationsBundle, fee_quoter.UpdatePrices, evmChain, contract.FunctionInput[fee_quoter.PriceUpdates]{
				ChainSelector: evmChain.Selector,
				Address:       common.HexToAddress(feeQuoter),
				Args: fee_quoter.PriceUpdates{
					TokenPriceUpdates: []fee_quoter.TokenPriceUpdate{
						{
							SourceToken: common.HexToAddress(mockERC20Ref0.Address),
							UsdPerToken: big.NewInt(1000000000000000000),
						},
						{
							SourceToken: common.HexToAddress(mockERC20Ref1.Address),
							UsdPerToken: big.NewInt(2000000000000000000),
						},
						{
							SourceToken: common.HexToAddress(mockERC20Ref2.Address),
							UsdPerToken: big.NewInt(3000000000000000000),
						},
					},
				},
			})
			require.NoError(t, err, "ExecuteOperation should not error")

			remoteChainSelector := uint64(4356164186791070119)

			// Pre-configure signature configs on CommitteeVerifier so they exist on-chain
			// when metadata is collected
			preConfiguredConfig := committee_verifier.SignatureConfig{
				SourceChainSelector: remoteChainSelector,
				Threshold:           1,
				Signers:             []common.Address{common.HexToAddress("0x01"), common.HexToAddress("0x02")},
			}
			_, err = operations.ExecuteOperation(e.OperationsBundle, committee_verifier.ApplySignatureConfigs, evmChain, contract.FunctionInput[committee_verifier.SignatureConfigArgs]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(committeeVerifier),
				Args: committee_verifier.SignatureConfigArgs{
					SignatureConfigUpdates: []committee_verifier.SignatureConfig{preConfiguredConfig},
				},
			})
			require.NoError(t, err, "Failed to pre-configure signature configs on CommitteeVerifier")

			// Verify they're actually on-chain
			getAllConfigsReport, err := operations.ExecuteOperation(e.OperationsBundle, committee_verifier.GetAllSignatureConfigs, evmChain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(committeeVerifier),
				Args:          nil,
			})
			require.NoError(t, err, "Failed to verify signature configs are on-chain")
			require.Len(t, getAllConfigsReport.Output, 1, "Should have one signature config on-chain")
			require.Equal(t, preConfiguredConfig.SourceChainSelector, getAllConfigsReport.Output[0].SourceChainSelector, "Signature config should have correct sourceChainSelector")
			require.Equal(t, preConfiguredConfig.Threshold, getAllConfigsReport.Output[0].Threshold, "Signature config should have correct threshold")
			require.Equal(t, preConfiguredConfig.Signers, getAllConfigsReport.Output[0].Signers, "Signature config should have correct signers")

			ccipMessageSource := common.HexToAddress("0x10").Bytes()
			ccipMessageDest := common.HexToAddress("0x11").Bytes()

			configureReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.ConfigureChainForLanes,
				e.BlockChains,
				adapters.ConfigureChainForLanesInput{
					ChainSelector: chainSelector,
					Router:        routerAddress,
					OnRamp:        onRamp,
					CommitteeVerifiers: []adapters.CommitteeVerifierConfig[datastore.AddressRef]{
						{
							CommitteeVerifier: []datastore.AddressRef{
								{
									Address: committeeVerifier,
									Type:    datastore.ContractType(committee_verifier.ContractType),
									Version: committee_verifier.Version,
								},
								{
									Address: committeeVerifierResolver,
									Type:    datastore.ContractType(committee_verifier.ResolverType),
									Version: committee_verifier.Version,
								},
							},
							RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfig{
								remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
							},
						},
					},
					FeeQuoter: feeQuoter,
					OffRamp:   offRamp,
					RemoteChains: map[uint64]adapters.RemoteChainConfig[[]byte, string]{
						remoteChainSelector: {
							AllowTrafficFrom:         true,
							OnRamps:                  [][]byte{ccipMessageSource},
							OffRamp:                  ccipMessageDest,
							DefaultInboundCCVs:       []string{committeeVerifier},
							DefaultOutboundCCVs:      []string{committeeVerifier},
							DefaultExecutor:          executorAddress,
							FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
							ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
							AddressBytesLength:       20,
							BaseExecutionGasCost:     80_000,
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Test fee token metadata
			require.NotNil(t, configureReport.Output.Metadata, "Metadata should be set")
			require.NotEmpty(t, configureReport.Output.Metadata.Contracts, "Should have contract metadata")

			// Expected fee token metadata
			expectedFeeTokens := map[string]struct {
				name     string
				symbol   string
				decimals uint8
			}{
				linkAddress: {
					name:     "LINK",
					symbol:   "LINK",
					decimals: 18,
				},
				wethAddress: {
					name:     "Wrapped Ether",
					symbol:   "WETH",
					decimals: 18,
				},
				mockERC20Ref0.Address: {
					name:     "Mock ERC20 0",
					symbol:   "MOCK 0",
					decimals: 18,
				},
				mockERC20Ref1.Address: {
					name:     "Mock ERC20 1",
					symbol:   "MOCK 1",
					decimals: 18,
				},
				mockERC20Ref2.Address: {
					name:     "Mock ERC20 2",
					symbol:   "MOCK 2",
					decimals: 6,
				},
			}

			// Build a map of contract metadata by address for easy lookup
			metadataByAddress := make(map[string]datastore.ContractMetadata)
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				metadataByAddress[contractMeta.Address] = contractMeta
			}

			// Verify each expected fee token has metadata
			for tokenAddr, expected := range expectedFeeTokens {
				contractMeta, found := metadataByAddress[tokenAddr]
				require.True(t, found, "Should have metadata for fee token %s", tokenAddr)
				require.Equal(t, tokenAddr, contractMeta.Address, "Fee token metadata should have correct address")
				require.Equal(t, chainSelector, contractMeta.ChainSelector, "Fee token metadata should have correct chain selector")

				metaMap, ok := contractMeta.Metadata.(map[string]interface{})
				require.True(t, ok, "Fee token metadata should be a map[string]interface{}")

				require.Equal(t, expected.name, metaMap["name"], "Fee token %s should have correct name", tokenAddr)
				require.Equal(t, expected.symbol, metaMap["symbol"], "Fee token %s should have correct symbol", tokenAddr)
				require.Equal(t, expected.decimals, metaMap["decimals"], "Fee token %s should have correct decimals", tokenAddr)
			}

			// Test CommitteeVerifier signature config metadata
			// Convert expected signers to hex strings for comparison
			expectedSignersHex := make([]string, len(preConfiguredConfig.Signers))
			for i, signer := range preConfiguredConfig.Signers {
				expectedSignersHex[i] = signer.Hex()
			}

			// Find CommitteeVerifier metadata entry for the pre-configured config
			committeeVerifierMetadataFound := false
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				if contractMeta.Address == committeeVerifier && contractMeta.ChainSelector == chainSelector {
					metaMap, ok := contractMeta.Metadata.(map[string]interface{})
					require.True(t, ok, "CommitteeVerifier metadata should be a map[string]interface{}")

					// Check if this is the metadata for our pre-configured signature config
					if sourceChainSelector, ok := metaMap["sourceChainSelector"].(uint64); ok {
						if sourceChainSelector == preConfiguredConfig.SourceChainSelector {
							require.Equal(t, committeeVerifier, contractMeta.Address, "CommitteeVerifier metadata should have correct address")
							require.Equal(t, chainSelector, contractMeta.ChainSelector, "CommitteeVerifier metadata should have correct chain selector")
							require.Equal(t, preConfiguredConfig.SourceChainSelector, metaMap["sourceChainSelector"], "Signature config should have correct sourceChainSelector")
							require.Equal(t, preConfiguredConfig.Threshold, metaMap["threshold"], "Signature config should have correct threshold")

							signers, ok := metaMap["signers"].([]string)
							require.True(t, ok, "Signers should be []string")
							require.Equal(t, expectedSignersHex, signers, "Signature config should have correct signers")

							committeeVerifierMetadataFound = true
							break
						}
					}
				}
			}

			require.True(t, committeeVerifierMetadataFound, "Should have found CommitteeVerifier signature config metadata")
		})
	}
}
