package sequences_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
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
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/message_hasher"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/link_token"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_5_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
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

			// Deploy a token and pool, then register it with TokenAdminRegistry
			// This is the lightest weight way to add a token to the registry for testing
			var rmnProxyAddr common.Address
			for _, addr := range deploymentReport.Output.Addresses {
				if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
					rmnProxyAddr = common.HexToAddress(addr.Address)
					break
				}
			}

			tokenAndPoolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployTokenAndPool,
				e.BlockChains.EVMChains()[chainSelector],
				tokens.DeployTokenAndPoolInput{
					Accounts: map[common.Address]*big.Int{
						common.HexToAddress("0x01"): big.NewInt(500_000),
					},
					DeployTokenPoolInput: tokens.DeployTokenPoolInput{
						ChainSel:                         chainSelector,
						TokenPoolType:                    datastore.ContractType("BurnMintTokenPool"),
						TokenPoolVersion:                 semver.MustParse("1.7.0"),
						TokenSymbol:                      "TEST",
						RateLimitAdmin:                   common.HexToAddress("0x01"),
						ThresholdAmountForAdditionalCCVs: big.NewInt(1e18),
						FeeAggregator:                    common.HexToAddress("0x03"),
						ConstructorArgs: tokens.ConstructorArgs{
							Decimals: 18,
							RMNProxy: rmnProxyAddr,
							Router:   common.HexToAddress(routerAddress),
						},
					},
				},
			)
			require.NoError(t, err, "Failed to deploy token and pool")
			require.Len(t, tokenAndPoolReport.Output.Addresses, 3, "Expected 3 addresses (token, pool, advanced pool hooks)")

			tokenAddress := tokenAndPoolReport.Output.Addresses[0].Address
			tokenPoolAddress := tokenAndPoolReport.Output.Addresses[1].Address

			// Find TokenAdminRegistry address
			var tokenAdminRegistryAddress string
			for _, addr := range deploymentReport.Output.Addresses {
				if addr.Type == datastore.ContractType(token_admin_registry.ContractType) {
					tokenAdminRegistryAddress = addr.Address
					break
				}
			}
			require.NotEmpty(t, tokenAdminRegistryAddress, "TokenAdminRegistry address should be found")

			// Register the token with TokenAdminRegistry
			_, err = operations.ExecuteSequence(
				e.OperationsBundle,
				v1_5_0_sequences.RegisterToken,
				e.BlockChains.EVMChains()[chainSelector],
				v1_5_0_sequences.RegisterTokenInput{
					ChainSelector:             chainSelector,
					TokenAddress:              common.HexToAddress(tokenAddress),
					TokenPoolAddress:          common.HexToAddress(tokenPoolAddress),
					ExternalAdmin:             common.Address{}, // Use internal admin
					TokenAdminRegistryAddress: common.HexToAddress(tokenAdminRegistryAddress),
				},
			)
			require.NoError(t, err, "Failed to register token with TokenAdminRegistry")

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
			// when metadata is collected. Use a different selector so it doesn't conflict with the input.
			preConfiguredSelector := uint64(9999999999999999999) // Different from remoteChainSelector
			preConfiguredConfig := committee_verifier.SignatureConfig{
				SourceChainSelector: preConfiguredSelector,
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

			// Pre-configure Router with onRamps and offRamps so they exist on-chain
			// when metadata is collected (similar to how fee tokens are priced before the sequence)
			ccipMessageSource := common.HexToAddress("0x10").Bytes()
			ccipMessageDest := common.HexToAddress("0x11").Bytes()
			_, err = operations.ExecuteOperation(e.OperationsBundle, router.ApplyRampUpdates, evmChain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(routerAddress),
				Args: router.ApplyRampsUpdatesArgs{
					OnRampUpdates: []router.OnRamp{
						{
							DestChainSelector: remoteChainSelector,
							OnRamp:            common.HexToAddress(onRamp),
						},
					},
					OffRampRemoves: []router.OffRamp{},
					OffRampAdds: []router.OffRamp{
						{
							SourceChainSelector: remoteChainSelector,
							OffRamp:             common.HexToAddress(offRamp),
						},
					},
				},
			})
			require.NoError(t, err, "Failed to pre-configure Router with onRamps and offRamps")

			// Verify they're actually on-chain
			getOffRampsReport, err := operations.ExecuteOperation(e.OperationsBundle, router.GetOffRamps, evmChain, contract.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(routerAddress),
				Args:          nil,
			})
			require.NoError(t, err, "Failed to verify offRamps are on-chain")
			require.Len(t, getOffRampsReport.Output, 1, "Should have one offRamp on-chain")
			require.Equal(t, remoteChainSelector, getOffRampsReport.Output[0].SourceChainSelector, "OffRamp should have correct sourceChainSelector")
			require.Equal(t, offRamp, getOffRampsReport.Output[0].OffRamp.Hex(), "OffRamp should have correct address")

			getOnRampReport, err := operations.ExecuteOperation(e.OperationsBundle, router.GetOnRamp, evmChain, contract.FunctionInput[uint64]{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(routerAddress),
				Args:          remoteChainSelector,
			})
			require.NoError(t, err, "Failed to verify onRamp is on-chain")
			require.Equal(t, onRamp, getOnRampReport.Output.Hex(), "OnRamp should have correct address")

			// Deploy a token and pool, then register it with TokenAdminRegistry
			// This is the lightest weight way to add a token to the registry for testing
			var rmnProxyAddr common.Address
			for _, addr := range deploymentReport.Output.Addresses {
				if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
					rmnProxyAddr = common.HexToAddress(addr.Address)
					break
				}
			}

			tokenAndPoolReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.DeployTokenAndPool,
				e.BlockChains.EVMChains()[chainSelector],
				tokens.DeployTokenAndPoolInput{
					Accounts: map[common.Address]*big.Int{
						common.HexToAddress("0x01"): big.NewInt(500_000),
					},
					DeployTokenPoolInput: tokens.DeployTokenPoolInput{
						ChainSel:                         chainSelector,
						TokenPoolType:                    datastore.ContractType("BurnMintTokenPool"),
						TokenPoolVersion:                 semver.MustParse("1.7.0"),
						TokenSymbol:                      "TEST",
						RateLimitAdmin:                   common.HexToAddress("0x01"),
						ThresholdAmountForAdditionalCCVs: big.NewInt(1e18),
						FeeAggregator:                    common.HexToAddress("0x03"),
						ConstructorArgs: tokens.ConstructorArgs{
							Decimals: 18,
							RMNProxy: rmnProxyAddr,
							Router:   common.HexToAddress(routerAddress),
						},
					},
				},
			)
			require.NoError(t, err, "Failed to deploy token and pool")
			require.Len(t, tokenAndPoolReport.Output.Addresses, 3, "Expected 3 addresses (token, pool, advanced pool hooks)")

			tokenAddress := tokenAndPoolReport.Output.Addresses[0].Address
			tokenPoolAddress := tokenAndPoolReport.Output.Addresses[1].Address

			// Find TokenAdminRegistry address
			var tokenAdminRegistryAddress string
			for _, addr := range deploymentReport.Output.Addresses {
				if addr.Type == datastore.ContractType(token_admin_registry.ContractType) {
					tokenAdminRegistryAddress = addr.Address
					break
				}
			}
			require.NotEmpty(t, tokenAdminRegistryAddress, "TokenAdminRegistry address should be found")

			// Register the token with TokenAdminRegistry
			_, err = operations.ExecuteSequence(
				e.OperationsBundle,
				v1_5_0_sequences.RegisterToken,
				e.BlockChains.EVMChains()[chainSelector],
				v1_5_0_sequences.RegisterTokenInput{
					ChainSelector:             chainSelector,
					TokenAddress:              common.HexToAddress(tokenAddress),
					TokenPoolAddress:          common.HexToAddress(tokenPoolAddress),
					ExternalAdmin:             common.Address{}, // Use internal admin
					TokenAdminRegistryAddress: common.HexToAddress(tokenAdminRegistryAddress),
				},
			)
			require.NoError(t, err, "Failed to register token with TokenAdminRegistry")

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

			// Find FeeQuoter metadata and verify fee tokens are nested within it
			feeQuoterMetadataFound := false
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				if contractMeta.Address == feeQuoter && contractMeta.ChainSelector == chainSelector {
					feeQuoterMetadataFound = true
					require.Equal(t, feeQuoter, contractMeta.Address, "FeeQuoter metadata should have correct address")
					require.Equal(t, chainSelector, contractMeta.ChainSelector, "FeeQuoter metadata should have correct chain selector")

					metaMap, ok := contractMeta.Metadata.(map[string]interface{})
					require.True(t, ok, "FeeQuoter metadata should be a map[string]interface{}")

					require.Equal(t, true, metaMap["configured"], "FeeQuoter metadata should have configured=true")
					require.Equal(t, "fee_quoter_configured", metaMap["test_metadata"], "FeeQuoter metadata should have test_metadata")

					// Verify feeTokens list exists
					feeTokensValue, ok := metaMap["feeTokens"]
					require.True(t, ok, "feeTokens should exist in FeeQuoter metadata")
					feeTokensList, ok := feeTokensValue.([]interface{})
					if !ok {
						// Try []map[string]interface{} (might not have gone through JSON round-trip)
						feeTokensListMap, okMap := feeTokensValue.([]map[string]interface{})
						require.True(t, okMap, "feeTokens should be a []interface{} or []map[string]interface{}, got %T", feeTokensValue)
						feeTokensList = make([]interface{}, len(feeTokensListMap))
						for i, m := range feeTokensListMap {
							feeTokensList[i] = m
						}
					}
					require.Len(t, feeTokensList, len(expectedFeeTokens), "Should have correct number of fee tokens")

					// Build a map of fee tokens by address for easy lookup
					feeTokensByAddress := make(map[string]map[string]interface{})
					for _, tokenInterface := range feeTokensList {
						tokenMap, ok := tokenInterface.(map[string]interface{})
						require.True(t, ok, "Each fee token should be a map[string]interface{}")
						address, ok := tokenMap["address"].(string)
						require.True(t, ok, "Fee token should have address field")
						feeTokensByAddress[address] = tokenMap
					}

					// Verify each expected fee token is in the list
					for tokenAddr, expected := range expectedFeeTokens {
						tokenMeta, found := feeTokensByAddress[tokenAddr]
						require.True(t, found, "Should have fee token %s in feeTokens list", tokenAddr)
						require.Equal(t, tokenAddr, tokenMeta["address"], "Fee token should have correct address")

						// Handle chainSelector - can be uint64 or float64 after JSON round-trip
						chainSelectorValue := tokenMeta["chainSelector"]
						switch v := chainSelectorValue.(type) {
						case uint64:
							require.Equal(t, chainSelector, v, "Fee token should have correct chain selector")
						case float64:
							require.Equal(t, float64(chainSelector), v, "Fee token should have correct chain selector")
						default:
							require.Fail(t, "chainSelector should be uint64 or float64, got %T", chainSelectorValue)
						}

						require.Equal(t, expected.name, tokenMeta["name"], "Fee token %s should have correct name", tokenAddr)
						require.Equal(t, expected.symbol, tokenMeta["symbol"], "Fee token %s should have correct symbol", tokenAddr)

						// Handle decimals - can be uint8 or float64 after JSON round-trip
						decimalsValue := tokenMeta["decimals"]
						switch v := decimalsValue.(type) {
						case uint8:
							require.Equal(t, expected.decimals, v, "Fee token %s should have correct decimals", tokenAddr)
						case float64:
							require.Equal(t, float64(expected.decimals), v, "Fee token %s should have correct decimals", tokenAddr)
						default:
							require.Fail(t, "decimals should be uint8 or float64, got %T", decimalsValue)
						}
					}

					break
				}
			}
			require.True(t, feeQuoterMetadataFound, "Should have found FeeQuoter metadata")

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

			// Test Router metadata (onRamps and offRamps)
			routerMetadataFound := false
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				if contractMeta.Address == routerAddress && contractMeta.ChainSelector == chainSelector {
					routerMetadataFound = true
					require.Equal(t, routerAddress, contractMeta.Address, "Router metadata should have correct address")
					require.Equal(t, chainSelector, contractMeta.ChainSelector, "Router metadata should have correct chain selector")

					metaMap, ok := contractMeta.Metadata.(map[string]interface{})
					require.True(t, ok, "Router metadata should be a map[string]interface{}")

					// Verify onRamps
					onRampsValue, ok := metaMap["onRamps"]
					require.True(t, ok, "onRamps should exist in metadata")
					onRampsMap, ok := onRampsValue.(map[string]interface{})
					if !ok {
						// Try map[string]string (might not have gone through JSON round-trip)
						onRampsMapStr, okStr := onRampsValue.(map[string]string)
						require.True(t, okStr, "onRamps should be a map[string]interface{} or map[string]string, got %T", onRampsValue)
						onRampsMap = make(map[string]interface{})
						for k, v := range onRampsMapStr {
							onRampsMap[k] = v
						}
					}
					require.Len(t, onRampsMap, 1, "Should have one onRamp")
					onRampAddrValue, ok := onRampsMap[fmt.Sprintf("%d", remoteChainSelector)]
					require.True(t, ok, "onRamp should exist for remote chain selector")
					onRampAddr, ok := onRampAddrValue.(string)
					require.True(t, ok, "onRamp address should be a string, got %T", onRampAddrValue)
					require.Equal(t, onRamp, onRampAddr, "OnRamp address should match")

					// Verify offRamps
					offRampsValue, ok := metaMap["offRamps"]
					require.True(t, ok, "offRamps should exist in metadata")
					offRampsMap, ok := offRampsValue.(map[string]interface{})
					if !ok {
						// Try map[string][]string (might not have gone through JSON round-trip)
						offRampsMapStr, okStr := offRampsValue.(map[string][]string)
						require.True(t, okStr, "offRamps should be a map[string]interface{} or map[string][]string, got %T", offRampsValue)
						offRampsMap = make(map[string]interface{})
						for k, v := range offRampsMapStr {
							// Convert []string to []interface{}
							interfaceSlice := make([]interface{}, len(v))
							for i, s := range v {
								interfaceSlice[i] = s
							}
							offRampsMap[k] = interfaceSlice
						}
					}
					require.Len(t, offRampsMap, 1, "Should have one offRamp entry")
					offRampAddrsValue, ok := offRampsMap[fmt.Sprintf("%d", remoteChainSelector)]
					require.True(t, ok, "offRamp should exist for remote chain selector")
					offRampAddrs, ok := offRampAddrsValue.([]interface{})
					if !ok {
						// Try []string (might not have gone through JSON round-trip)
						offRampAddrsStr, okStr := offRampAddrsValue.([]string)
						require.True(t, okStr, "offRamp addresses should be []interface{} or []string, got %T", offRampAddrsValue)
						offRampAddrs = make([]interface{}, len(offRampAddrsStr))
						for i, s := range offRampAddrsStr {
							offRampAddrs[i] = s
						}
					}
					require.Len(t, offRampAddrs, 1, "Should have one offRamp address")
					offRampAddrValue, ok := offRampAddrs[0].(string)
					require.True(t, ok, "offRamp address should be a string, got %T", offRampAddrs[0])
					require.Equal(t, offRamp, offRampAddrValue, "OffRamp address should match")

					break
				}
			}
			require.True(t, routerMetadataFound, "Should have found Router metadata")

			// Test OnRamp metadata (destChainConfigs)
			onRampMetadataFound := false
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				if contractMeta.Address == onRamp && contractMeta.ChainSelector == chainSelector {
					onRampMetadataFound = true
					require.Equal(t, onRamp, contractMeta.Address, "OnRamp metadata should have correct address")
					require.Equal(t, chainSelector, contractMeta.ChainSelector, "OnRamp metadata should have correct chain selector")

					metaMap, ok := contractMeta.Metadata.(map[string]interface{})
					require.True(t, ok, "OnRamp metadata should be a map[string]interface{}")

					// Verify destChainConfigs list exists
					destChainConfigsValue, ok := metaMap["destChainConfigs"]
					require.True(t, ok, "destChainConfigs should exist in OnRamp metadata")
					destChainConfigsList, ok := destChainConfigsValue.([]interface{})
					if !ok {
						// Try []map[string]interface{} (might not have gone through JSON round-trip)
						destChainConfigsListMap, okMap := destChainConfigsValue.([]map[string]interface{})
						require.True(t, okMap, "destChainConfigs should be a []interface{} or []map[string]interface{}, got %T", destChainConfigsValue)
						destChainConfigsList = make([]interface{}, len(destChainConfigsListMap))
						for i, m := range destChainConfigsListMap {
							destChainConfigsList[i] = m
						}
					}
					require.Len(t, destChainConfigsList, 1, "Should have one dest chain config")

					// Verify the dest chain config structure
					configMap, ok := destChainConfigsList[0].(map[string]interface{})
					require.True(t, ok, "Each dest chain config should be a map[string]interface{}")

					// Verify destChainSelector
					destChainSelectorValue := configMap["destChainSelector"]
					switch v := destChainSelectorValue.(type) {
					case uint64:
						require.Equal(t, remoteChainSelector, v, "Dest chain config should have correct destChainSelector")
					case float64:
						require.Equal(t, float64(remoteChainSelector), v, "Dest chain config should have correct destChainSelector")
					default:
						require.Fail(t, "destChainSelector should be uint64 or float64, got %T", destChainSelectorValue)
					}

					// Verify router
					routerValue, ok := configMap["router"].(string)
					require.True(t, ok, "router should be a string")
					require.Equal(t, routerAddress, routerValue, "Dest chain config should have correct router address")

					// Verify defaultExecutor
					defaultExecutorValue, ok := configMap["defaultExecutor"].(string)
					require.True(t, ok, "defaultExecutor should be a string")
					require.Equal(t, executorAddress, defaultExecutorValue, "Dest chain config should have correct defaultExecutor address")

					// Verify offRamp
					offRampValue, ok := configMap["offRamp"].(string)
					require.True(t, ok, "offRamp should be a string")
					require.True(t, len(offRampValue) > 0, "offRamp should not be empty")

					// Verify other fields exist
					require.Contains(t, configMap, "addressBytesLength", "Dest chain config should have addressBytesLength")
					require.Contains(t, configMap, "baseExecutionGasCost", "Dest chain config should have baseExecutionGasCost")
					require.Contains(t, configMap, "tokenReceiverAllowed", "Dest chain config should have tokenReceiverAllowed")
					require.Contains(t, configMap, "messageNetworkFeeUSDCents", "Dest chain config should have messageNetworkFeeUSDCents")
					require.Contains(t, configMap, "tokenNetworkFeeUSDCents", "Dest chain config should have tokenNetworkFeeUSDCents")
					require.Contains(t, configMap, "defaultCCVs", "Dest chain config should have defaultCCVs")
					require.Contains(t, configMap, "laneMandatedCCVs", "Dest chain config should have laneMandatedCCVs")

					break
				}
			}
			require.True(t, onRampMetadataFound, "Should have found OnRamp metadata")

			// Test OffRamp metadata (sourceChainConfigs)
			offRampMetadataFound := false
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				if contractMeta.Address == offRamp && contractMeta.ChainSelector == chainSelector {
					offRampMetadataFound = true
					require.Equal(t, offRamp, contractMeta.Address, "OffRamp metadata should have correct address")
					require.Equal(t, chainSelector, contractMeta.ChainSelector, "OffRamp metadata should have correct chain selector")

					metaMap, ok := contractMeta.Metadata.(map[string]interface{})
					require.True(t, ok, "OffRamp metadata should be a map[string]interface{}")

					// Verify sourceChainConfigs list exists
					sourceChainConfigsValue, ok := metaMap["sourceChainConfigs"]
					require.True(t, ok, "sourceChainConfigs should exist in OffRamp metadata")
					sourceChainConfigsList, ok := sourceChainConfigsValue.([]interface{})
					if !ok {
						// Try []map[string]interface{} (might not have gone through JSON round-trip)
						sourceChainConfigsListMap, okMap := sourceChainConfigsValue.([]map[string]interface{})
						require.True(t, okMap, "sourceChainConfigs should be a []interface{} or []map[string]interface{}, got %T", sourceChainConfigsValue)
						sourceChainConfigsList = make([]interface{}, len(sourceChainConfigsListMap))
						for i, m := range sourceChainConfigsListMap {
							sourceChainConfigsList[i] = m
						}
					}
					require.Len(t, sourceChainConfigsList, 1, "Should have one source chain config")

					// Verify the source chain config structure
					configMap, ok := sourceChainConfigsList[0].(map[string]interface{})
					require.True(t, ok, "Each source chain config should be a map[string]interface{}")

					// Verify sourceChainSelector
					sourceChainSelectorValue := configMap["sourceChainSelector"]
					switch v := sourceChainSelectorValue.(type) {
					case uint64:
						require.Equal(t, remoteChainSelector, v, "Source chain config should have correct sourceChainSelector")
					case float64:
						require.Equal(t, float64(remoteChainSelector), v, "Source chain config should have correct sourceChainSelector")
					default:
						require.Fail(t, "sourceChainSelector should be uint64 or float64, got %T", sourceChainSelectorValue)
					}

					// Verify router
					routerValue, ok := configMap["router"].(string)
					require.True(t, ok, "router should be a string")
					require.Equal(t, routerAddress, routerValue, "Source chain config should have correct router address")

					// Verify isEnabled
					isEnabledValue, ok := configMap["isEnabled"].(bool)
					require.True(t, ok, "isEnabled should be a bool")
					require.True(t, isEnabledValue, "Source chain config should be enabled")

					// Verify onRamps
					onRampsValue, ok := configMap["onRamps"].([]interface{})
					if !ok {
						// Try []string (might not have gone through JSON round-trip)
						onRampsStr, okStr := configMap["onRamps"].([]string)
						require.True(t, okStr, "onRamps should be []interface{} or []string, got %T", configMap["onRamps"])
						onRampsValue = make([]interface{}, len(onRampsStr))
						for i, s := range onRampsStr {
							onRampsValue[i] = s
						}
					}
					require.True(t, len(onRampsValue) > 0, "onRamps should not be empty")

					// Verify other fields exist
					require.Contains(t, configMap, "defaultCCVs", "Source chain config should have defaultCCVs")
					require.Contains(t, configMap, "laneMandatedCCVs", "Source chain config should have laneMandatedCCVs")

					break
				}
			}
			require.True(t, offRampMetadataFound, "Should have found OffRamp metadata")

			// Test TokenAdminRegistry metadata (tokens)
			tokenAdminRegistryMetadataFound := false
			for _, contractMeta := range configureReport.Output.Metadata.Contracts {
				if contractMeta.Address == tokenAdminRegistryAddress && contractMeta.ChainSelector == chainSelector {
					tokenAdminRegistryMetadataFound = true
					require.Equal(t, tokenAdminRegistryAddress, contractMeta.Address, "TokenAdminRegistry metadata should have correct address")
					require.Equal(t, chainSelector, contractMeta.ChainSelector, "TokenAdminRegistry metadata should have correct chain selector")

					metaMap, ok := contractMeta.Metadata.(map[string]interface{})
					require.True(t, ok, "TokenAdminRegistry metadata should be a map[string]interface{}")

					// Verify tokens list exists
					tokensValue, ok := metaMap["tokens"]
					require.True(t, ok, "tokens should exist in TokenAdminRegistry metadata")
					tokensList, ok := tokensValue.([]interface{})
					if !ok {
						// Try []map[string]interface{} (might not have gone through JSON round-trip)
						tokensListMap, okMap := tokensValue.([]map[string]interface{})
						require.True(t, okMap, "tokens should be a []interface{} or []map[string]interface{}, got %T", tokensValue)
						tokensList = make([]interface{}, len(tokensListMap))
						for i, m := range tokensListMap {
							tokensList[i] = m
						}
					}
					require.Greater(t, len(tokensList), 0, "Should have at least one token registered")

					// Verify the first token has the expected structure
					tokenObj, ok := tokensList[0].(map[string]interface{})
					require.True(t, ok, "Token should be a map[string]interface{}")

					// Verify token address exists and matches
					tokenAddrValue, ok := tokenObj["address"].(string)
					require.True(t, ok, "Token address should be a string")
					require.Equal(t, tokenAddress, tokenAddrValue, "Token address should match registered token")

					// Verify token has name, symbol, decimals
					require.Contains(t, tokenObj, "name", "Token should have name")
					require.Contains(t, tokenObj, "symbol", "Token should have symbol")
					require.Contains(t, tokenObj, "decimals", "Token should have decimals")

					// Verify admin fields exist
					require.Contains(t, tokenObj, "admin", "Token should have admin")
					require.Contains(t, tokenObj, "pendingAdministrator", "Token should have pendingAdministrator")

					// Verify tokenPool exists and has rmnProxy
					tokenPoolValueRaw, tokenPoolExists := tokenObj["tokenPool"]
					require.True(t, tokenPoolExists, "Token should have tokenPool field")

					// tokenPool might be nil if the token doesn't have a pool set
					if tokenPoolValueRaw != nil {
						tokenPoolValue, ok := tokenPoolValueRaw.(map[string]interface{})
						if !ok {
							// Try to see what type it actually is for debugging
							t.Fatalf("TokenPool should be map[string]interface{} or nil, got %T: %v", tokenPoolValueRaw, tokenPoolValueRaw)
						}
						require.Contains(t, tokenPoolValue, "address", "TokenPool should have address")
						require.Contains(t, tokenPoolValue, "rmnProxy", "TokenPool should have rmnProxy")

						// Verify rmnProxy exists and has owner and arm
						rmnProxyValueRaw, rmnProxyExists := tokenPoolValue["rmnProxy"]
						require.True(t, rmnProxyExists, "TokenPool should have rmnProxy field")
						require.NotNil(t, rmnProxyValueRaw, "RMNProxy should not be nil")

						rmnProxyValue, ok := rmnProxyValueRaw.(map[string]interface{})
						if !ok {
							// Try to see what type it actually is for debugging
							t.Fatalf("RMNProxy should be map[string]interface{}, got %T: %v", rmnProxyValueRaw, rmnProxyValueRaw)
						}
						require.Contains(t, rmnProxyValue, "address", "RMNProxy should have address")
						require.Contains(t, rmnProxyValue, "owner", "RMNProxy should have owner")
						require.Contains(t, rmnProxyValue, "arm", "RMNProxy should have arm")
					} else {
						// If tokenPool is nil, that's okay - some tokens might not have pools
						// But in our test, we registered a token with a pool, so this shouldn't happen
						t.Logf("Warning: tokenPool is nil for token %s, but we registered it with a pool", tokenAddress)
					}

					break
				}
			}
			require.True(t, tokenAdminRegistryMetadataFound, "Should have found TokenAdminRegistry metadata")

			// Output metadata to JSON file for inspection
			metadataJSON, err := json.MarshalIndent(configureReport.Output.Metadata, "", "  ")
			require.NoError(t, err, "Failed to marshal metadata to JSON")
			err = os.WriteFile("test_metadata_output.json", metadataJSON, 0644)
			require.NoError(t, err, "Failed to write metadata to file")
			t.Logf("Metadata written to test_metadata_output.json")
		})
	}
}

func TestConfigureChainForLanes_PendingChangesProjection(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	remoteChainSelector1 := uint64(4356164186791070119)
	remoteChainSelector2 := uint64(1234567890123456789) // Different remote chain

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	evmChain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy contracts
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

	var routerAddress, onRamp, feeQuoter, offRamp, committeeVerifier, committeeVerifierResolver, executorAddress string
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

	// Configure first remote chain (this establishes current state)
	ccipMessageSource1 := common.HexToAddress("0x10").Bytes()
	ccipMessageDest1 := common.HexToAddress("0x11").Bytes()
	_, err = operations.ExecuteSequence(
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
						remoteChainSelector1: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
					},
				},
			},
			FeeQuoter: feeQuoter,
			OffRamp:   offRamp,
			RemoteChains: map[uint64]adapters.RemoteChainConfig[[]byte, string]{
				remoteChainSelector1: {
					AllowTrafficFrom:         true,
					OnRamps:                  [][]byte{ccipMessageSource1},
					OffRamp:                  ccipMessageDest1,
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
	require.NoError(t, err, "First ConfigureChainForLanes should not error")

	// Now configure a second remote chain (this is the "pending" change we're testing)
	// The metadata should project both chains: chain1 from current state, chain2 from pending
	ccipMessageSource2 := common.HexToAddress("0x20").Bytes()
	ccipMessageDest2 := common.HexToAddress("0x21").Bytes()
	configureReport2, err := operations.ExecuteSequence(
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
						remoteChainSelector2: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
					},
				},
			},
			FeeQuoter: feeQuoter,
			OffRamp:   offRamp,
			RemoteChains: map[uint64]adapters.RemoteChainConfig[[]byte, string]{
				remoteChainSelector2: {
					AllowTrafficFrom:         true,
					OnRamps:                  [][]byte{ccipMessageSource2},
					OffRamp:                  ccipMessageDest2,
					DefaultInboundCCVs:       []string{committeeVerifier},
					DefaultOutboundCCVs:      []string{committeeVerifier},
					DefaultExecutor:          executorAddress,
					FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
					ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
					AddressBytesLength:       20,
					BaseExecutionGasCost:     90_000, // Different value to verify pending changes
				},
			},
		},
	)
	require.NoError(t, err, "Second ConfigureChainForLanes should not error")

	// Verify Router metadata shows both remote chains (chain1 from current state, chain2 from pending)
	routerMetadataFound := false
	for _, contractMeta := range configureReport2.Output.Metadata.Contracts {
		if contractMeta.Address == routerAddress && contractMeta.ChainSelector == chainSelector {
			routerMetadataFound = true
			metaMap, ok := contractMeta.Metadata.(map[string]interface{})
			require.True(t, ok, "Router metadata should be a map[string]interface{}")

			// Verify onRamps shows both chains
			onRampsValue, ok := metaMap["onRamps"]
			require.True(t, ok, "onRamps should exist in metadata")
			onRampsMap, ok := onRampsValue.(map[string]interface{})
			if !ok {
				onRampsMapStr, okStr := onRampsValue.(map[string]string)
				require.True(t, okStr, "onRamps should be a map")
				onRampsMap = make(map[string]interface{})
				for k, v := range onRampsMapStr {
					onRampsMap[k] = v
				}
			}
			// Should have both remote chains: chain1 (current) and chain2 (pending)
			require.Len(t, onRampsMap, 2, "Should have onRamps for both remote chains")
			require.Equal(t, onRamp, onRampsMap[fmt.Sprintf("%d", remoteChainSelector1)], "Chain1 onRamp should match current state")
			require.Equal(t, onRamp, onRampsMap[fmt.Sprintf("%d", remoteChainSelector2)], "Chain2 onRamp should match pending update")

			// Verify offRamps shows both chains
			offRampsValue, ok := metaMap["offRamps"]
			require.True(t, ok, "offRamps should exist in metadata")
			offRampsMap, ok := offRampsValue.(map[string]interface{})
			if !ok {
				offRampsMapStr, okStr := offRampsValue.(map[string][]string)
				require.True(t, okStr, "offRamps should be a map")
				offRampsMap = make(map[string]interface{})
				for k, v := range offRampsMapStr {
					interfaceSlice := make([]interface{}, len(v))
					for i, s := range v {
						interfaceSlice[i] = s
					}
					offRampsMap[k] = interfaceSlice
				}
			}
			// Should have both remote chains
			require.Len(t, offRampsMap, 2, "Should have offRamps for both remote chains")
			chain1OffRamps, ok := offRampsMap[fmt.Sprintf("%d", remoteChainSelector1)].([]interface{})
			require.True(t, ok, "Chain1 offRamps should be a slice")
			require.Contains(t, chain1OffRamps, offRamp, "Chain1 should include current offRamp")
			chain2OffRamps, ok := offRampsMap[fmt.Sprintf("%d", remoteChainSelector2)].([]interface{})
			require.True(t, ok, "Chain2 offRamps should be a slice")
			require.Contains(t, chain2OffRamps, offRamp, "Chain2 should include pending offRamp")

			break
		}
	}
	require.True(t, routerMetadataFound, "Should have found Router metadata")

	// Verify OnRamp metadata shows both dest chain configs
	// Chain1 should be from current state (queried from chain)
	// Chain2 should be from pending configs (projected state)
	onRampMetadataFound := false
	for _, contractMeta := range configureReport2.Output.Metadata.Contracts {
		if contractMeta.Address == onRamp && contractMeta.ChainSelector == chainSelector {
			onRampMetadataFound = true
			metaMap, ok := contractMeta.Metadata.(map[string]interface{})
			require.True(t, ok, "OnRamp metadata should be a map[string]interface{}")

			destChainConfigsValue, ok := metaMap["destChainConfigs"]
			require.True(t, ok, "destChainConfigs should exist")
			destChainConfigsList, ok := destChainConfigsValue.([]interface{})
			if !ok {
				destChainConfigsListMap, okMap := destChainConfigsValue.([]map[string]interface{})
				require.True(t, okMap, "destChainConfigs should be a slice")
				destChainConfigsList = make([]interface{}, len(destChainConfigsListMap))
				for i, m := range destChainConfigsListMap {
					destChainConfigsList[i] = m
				}
			}

			// Should have configs for both chains
			require.Len(t, destChainConfigsList, 2, "Should have dest chain configs for both remote chains")

			// Find configs for each chain
			chain1ConfigFound := false
			chain2ConfigFound := false
			for _, configInterface := range destChainConfigsList {
				configMap, ok := configInterface.(map[string]interface{})
				require.True(t, ok, "Config should be a map")

				var destChainSelector uint64
				switch v := configMap["destChainSelector"].(type) {
				case uint64:
					destChainSelector = v
				case float64:
					destChainSelector = uint64(v)
				default:
					continue
				}

				if destChainSelector == remoteChainSelector1 {
					chain1ConfigFound = true
					// Chain1 config should be from current state (has messageNumber, etc.)
					require.Contains(t, configMap, "messageNumber", "Chain1 config should have messageNumber from current state")
					require.Equal(t, uint64(80_000), getUint64FromInterface(configMap["baseExecutionGasCost"]), "Chain1 should have baseExecutionGasCost from first config")
				} else if destChainSelector == remoteChainSelector2 {
					chain2ConfigFound = true
					// Chain2 config should be from pending (projected state)
					// messageNumber should be 0 (default for pending configs)
					messageNumber := getUint64FromInterface(configMap["messageNumber"])
					require.Equal(t, uint64(0), messageNumber, "Chain2 config should have messageNumber=0 from pending config")
					require.Equal(t, uint64(90_000), getUint64FromInterface(configMap["baseExecutionGasCost"]), "Chain2 should have baseExecutionGasCost=90000 from pending config")
					require.Equal(t, routerAddress, configMap["router"], "Chain2 config should have router from pending config")
					require.Equal(t, executorAddress, configMap["defaultExecutor"], "Chain2 config should have defaultExecutor from pending config")
				}
			}
			require.True(t, chain1ConfigFound, "Should have found config for chain1")
			require.True(t, chain2ConfigFound, "Should have found config for chain2")

			break
		}
	}
	require.True(t, onRampMetadataFound, "Should have found OnRamp metadata")

	// Verify OffRamp metadata shows both source chain configs
	offRampMetadataFound := false
	for _, contractMeta := range configureReport2.Output.Metadata.Contracts {
		if contractMeta.Address == offRamp && contractMeta.ChainSelector == chainSelector {
			offRampMetadataFound = true
			metaMap, ok := contractMeta.Metadata.(map[string]interface{})
			require.True(t, ok, "OffRamp metadata should be a map[string]interface{}")

			sourceChainConfigsValue, ok := metaMap["sourceChainConfigs"]
			require.True(t, ok, "sourceChainConfigs should exist")
			sourceChainConfigsList, ok := sourceChainConfigsValue.([]interface{})
			if !ok {
				sourceChainConfigsListMap, okMap := sourceChainConfigsValue.([]map[string]interface{})
				require.True(t, okMap, "sourceChainConfigs should be a slice")
				sourceChainConfigsList = make([]interface{}, len(sourceChainConfigsListMap))
				for i, m := range sourceChainConfigsListMap {
					sourceChainConfigsList[i] = m
				}
			}

			// Should have configs for both chains
			require.Len(t, sourceChainConfigsList, 2, "Should have source chain configs for both remote chains")

			// Find configs for each chain
			chain1ConfigFound := false
			chain2ConfigFound := false
			for _, configInterface := range sourceChainConfigsList {
				configMap, ok := configInterface.(map[string]interface{})
				require.True(t, ok, "Config should be a map")

				var sourceChainSelector uint64
				switch v := configMap["sourceChainSelector"].(type) {
				case uint64:
					sourceChainSelector = v
				case float64:
					sourceChainSelector = uint64(v)
				default:
					continue
				}

				if sourceChainSelector == remoteChainSelector1 {
					chain1ConfigFound = true
					// Chain1 config should be from current state
					require.Contains(t, configMap, "isEnabled", "Chain1 config should have isEnabled from current state")
				} else if sourceChainSelector == remoteChainSelector2 {
					chain2ConfigFound = true
					// Chain2 config should be from pending (projected state)
					require.Equal(t, true, configMap["isEnabled"], "Chain2 config should have isEnabled=true from pending config")
					require.Equal(t, routerAddress, configMap["router"], "Chain2 config should have router from pending config")
				}
			}
			require.True(t, chain1ConfigFound, "Should have found config for chain1")
			require.True(t, chain2ConfigFound, "Should have found config for chain2")

			break
		}
	}
	require.True(t, offRampMetadataFound, "Should have found OffRamp metadata")

	// Verify CommitteeVerifier signature config metadata shows both chains
	// Chain1 should be from current state, Chain2 should be from pending
	committeeVerifierMetadataFound := false
	for _, contractMeta := range configureReport2.Output.Metadata.Contracts {
		if contractMeta.Address == committeeVerifier && contractMeta.ChainSelector == chainSelector {
			metaMap, ok := contractMeta.Metadata.(map[string]interface{})
			require.True(t, ok, "CommitteeVerifier metadata should be a map[string]interface{}")

			// Check if this is the metadata for chain1 or chain2
			sourceChainSelector, ok := metaMap["sourceChainSelector"].(uint64)
			if !ok {
				// Try float64 (from JSON)
				if sourceChainSelectorFloat, okFloat := metaMap["sourceChainSelector"].(float64); okFloat {
					sourceChainSelector = uint64(sourceChainSelectorFloat)
				} else {
					continue
				}
			}

			if sourceChainSelector == remoteChainSelector1 {
				committeeVerifierMetadataFound = true
				// Chain1 config should be from current state (first configuration)
				require.Equal(t, remoteChainSelector1, metaMap["sourceChainSelector"], "Chain1 should have correct sourceChainSelector")
				threshold, ok := metaMap["threshold"].(uint8)
				if !ok {
					thresholdFloat, okFloat := metaMap["threshold"].(float64)
					require.True(t, okFloat, "Threshold should be uint8 or float64")
					threshold = uint8(thresholdFloat)
				}
				require.Equal(t, uint8(1), threshold, "Chain1 should have threshold=1 from first config")
				signers, ok := metaMap["signers"].([]string)
				require.True(t, ok, "Signers should be []string")
				require.Len(t, signers, 1, "Chain1 should have 1 signer from first config")
				require.Equal(t, common.HexToAddress("0x01").Hex(), signers[0], "Chain1 should have correct signer")
			} else if sourceChainSelector == remoteChainSelector2 {
				// Chain2 config should be from pending (second configuration)
				require.Equal(t, remoteChainSelector2, metaMap["sourceChainSelector"], "Chain2 should have correct sourceChainSelector")
				threshold, ok := metaMap["threshold"].(uint8)
				if !ok {
					thresholdFloat, okFloat := metaMap["threshold"].(float64)
					require.True(t, okFloat, "Threshold should be uint8 or float64")
					threshold = uint8(thresholdFloat)
				}
				require.Equal(t, uint8(1), threshold, "Chain2 should have threshold=1 from pending config")
				signers, ok := metaMap["signers"].([]string)
				require.True(t, ok, "Signers should be []string")
				require.Len(t, signers, 1, "Chain2 should have 1 signer from pending config")
				require.Equal(t, common.HexToAddress("0x01").Hex(), signers[0], "Chain2 should have correct signer from pending config")
			}
		}
	}
	require.True(t, committeeVerifierMetadataFound, "Should have found CommitteeVerifier metadata for chain1")

	// Verify we have metadata for both chains
	chain1SigConfigFound := false
	chain2SigConfigFound := false
	for _, contractMeta := range configureReport2.Output.Metadata.Contracts {
		if contractMeta.Address == committeeVerifier && contractMeta.ChainSelector == chainSelector {
			metaMap, ok := contractMeta.Metadata.(map[string]interface{})
			require.True(t, ok, "CommitteeVerifier metadata should be a map[string]interface{}")

			sourceChainSelector, ok := metaMap["sourceChainSelector"].(uint64)
			if !ok {
				if sourceChainSelectorFloat, okFloat := metaMap["sourceChainSelector"].(float64); okFloat {
					sourceChainSelector = uint64(sourceChainSelectorFloat)
				} else {
					continue
				}
			}

			if sourceChainSelector == remoteChainSelector1 {
				chain1SigConfigFound = true
			} else if sourceChainSelector == remoteChainSelector2 {
				chain2SigConfigFound = true
			}
		}
	}
	require.True(t, chain1SigConfigFound, "Should have found signature config for chain1")
	require.True(t, chain2SigConfigFound, "Should have found signature config for chain2")
}

// Helper function to extract uint64 from interface{} (handles both uint64 and float64 from JSON)
func getUint64FromInterface(v interface{}) uint64 {
	switch val := v.(type) {
	case uint64:
		return val
	case uint32:
		return uint64(val)
	case float64:
		return uint64(val)
	case int:
		return uint64(val)
	case int64:
		return uint64(val)
	default:
		return 0
	}
}
