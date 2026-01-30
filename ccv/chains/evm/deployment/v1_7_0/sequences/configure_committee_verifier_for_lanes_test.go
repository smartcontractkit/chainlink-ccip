package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
)

func TestConfigureCommitteeVerifierForLanes(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) sequences.ConfigureCommitteeVerifierForLanesInput
		expectedErr string
	}{
		{
			desc: "happy path",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) sequences.ConfigureCommitteeVerifierForLanesInput {
				var routerAddress string
				var committeeVerifier string
				var committeeVerifierResolver string
				for _, addr := range chainReport.Output.Addresses {
					switch addr.Type {
					case datastore.ContractType(router.ContractType):
						routerAddress = addr.Address
					case datastore.ContractType(committee_verifier.ContractType):
						committeeVerifier = addr.Address
					case datastore.ContractType(committee_verifier.ResolverType):
						committeeVerifierResolver = addr.Address
					}
				}
				remoteChainSelector := uint64(4356164186791070119)
				return sequences.ConfigureCommitteeVerifierForLanesInput{
					ChainSelector: chainReport.Input.ChainSelector,
					Router:        routerAddress,
					CommitteeVerifierConfig: adapters.CommitteeVerifierConfig[datastore.AddressRef]{
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
				}
			},
			expectedErr: "",
		},
		{
			desc: "multiple remote chains",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) sequences.ConfigureCommitteeVerifierForLanesInput {
				var routerAddress string
				var committeeVerifier string
				var committeeVerifierResolver string
				for _, addr := range chainReport.Output.Addresses {
					switch addr.Type {
					case datastore.ContractType(router.ContractType):
						routerAddress = addr.Address
					case datastore.ContractType(committee_verifier.ContractType):
						committeeVerifier = addr.Address
					case datastore.ContractType(committee_verifier.ResolverType):
						committeeVerifierResolver = addr.Address
					}
				}
				remoteChainSelector1 := uint64(4356164186791070119)
				remoteChainSelector2 := uint64(4949039107694359620)
				return sequences.ConfigureCommitteeVerifierForLanesInput{
					ChainSelector: chainReport.Input.ChainSelector,
					Router:        routerAddress,
					CommitteeVerifierConfig: adapters.CommitteeVerifierConfig[datastore.AddressRef]{
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
							remoteChainSelector2: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
						},
					},
				}
			},
			expectedErr: "",
		},
		{
			desc: "with allowlist enabled",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) sequences.ConfigureCommitteeVerifierForLanesInput {
				var routerAddress string
				var committeeVerifier string
				var committeeVerifierResolver string
				for _, addr := range chainReport.Output.Addresses {
					switch addr.Type {
					case datastore.ContractType(router.ContractType):
						routerAddress = addr.Address
					case datastore.ContractType(committee_verifier.ContractType):
						committeeVerifier = addr.Address
					case datastore.ContractType(committee_verifier.ResolverType):
						committeeVerifierResolver = addr.Address
					}
				}
				remoteChainSelector := uint64(4356164186791070119)
				config := testsetup.CreateBasicCommitteeVerifierRemoteChainConfig()
				config.AllowlistEnabled = true
				config.AddedAllowlistedSenders = []string{common.HexToAddress("0x10").Hex(), common.HexToAddress("0x11").Hex()}
				return sequences.ConfigureCommitteeVerifierForLanesInput{
					ChainSelector: chainReport.Input.ChainSelector,
					Router:        routerAddress,
					CommitteeVerifierConfig: adapters.CommitteeVerifierConfig[datastore.AddressRef]{
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
							remoteChainSelector: config,
						},
					},
				}
			},
			expectedErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")
			evmChain := e.BlockChains.EVMChains()[chainSelector]

			// Deploy chain contracts
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
					ContractParams: testsetup.CreateBasicContractParams(),
					CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")

			// Configure committee verifier for lanes
			input := test.makeInput(deploymentReport)
			configureReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.ConfigureCommitteeVerifierForLanes,
				e.BlockChains,
				input,
			)
			if test.expectedErr != "" {
				require.Error(t, err, "ExecuteSequence should error")
				require.Contains(t, err.Error(), test.expectedErr)
				return
			}
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, configureReport.Output.BatchOps, 1, "Expected 1 batch operation in output")

			// Verify configurations for each remote chain
			for remoteSelector, remoteConfig := range input.RemoteChains {
				// Check remote chain config on CommitteeVerifier
				remoteChainConfigReport, err := operations.ExecuteOperation(
					testsetup.BundleWithFreshReporter(e.OperationsBundle),
					committee_verifier.GetRemoteChainConfig,
					evmChain,
					contract.FunctionInput[uint64]{
						ChainSelector: evmChain.Selector,
						Address:       common.HexToAddress(input.CommitteeVerifier[0].Address),
						Args:          remoteSelector,
					},
				)
				require.NoError(t, err, "ExecuteOperation should not error")
				require.Equal(t, common.HexToAddress(input.Router), remoteChainConfigReport.Output.Router, "Router in remote chain config should match")
				require.Equal(t, remoteConfig.AllowlistEnabled, remoteChainConfigReport.Output.AllowlistEnabled, "AllowlistEnabled should match")

				// Check signature config on CommitteeVerifier
				signatureConfigReport, err := operations.ExecuteOperation(
					testsetup.BundleWithFreshReporter(e.OperationsBundle),
					committee_verifier.GetSignatureConfig,
					evmChain,
					contract.FunctionInput[uint64]{
						ChainSelector: evmChain.Selector,
						Address:       common.HexToAddress(input.CommitteeVerifier[0].Address),
						Args:          remoteSelector,
					},
				)
				require.NoError(t, err, "ExecuteOperation should not error")
				require.Equal(t, remoteConfig.SignatureConfig.Threshold, signatureConfigReport.Output.Threshold, "Threshold should match")
				require.Equal(t, remoteSelector, signatureConfigReport.Output.SourceChainSelector, "SourceChainSelector should match")
				expectedSigners := make([]common.Address, len(remoteConfig.SignatureConfig.Signers))
				for i, signer := range remoteConfig.SignatureConfig.Signers {
					expectedSigners[i] = common.HexToAddress(signer)
				}
				require.Equal(t, expectedSigners, signatureConfigReport.Output.Signers, "Signers should match")

				// Check outbound implementation on CommitteeVerifierResolver
				boundResolver, err := versioned_verifier_resolver.NewVersionedVerifierResolver(
					common.HexToAddress(input.CommitteeVerifier[1].Address),
					evmChain.Client,
				)
				require.NoError(t, err, "Failed to instantiate VersionedVerifierResolver")
				outboundImpl, err := boundResolver.GetOutboundImplementation(
					&bind.CallOpts{Context: t.Context()},
					remoteSelector,
					[]byte{},
				)
				require.NoError(t, err, "GetOutboundImplementation should not error")
				require.Equal(t, input.CommitteeVerifier[0].Address, outboundImpl.Hex(), "Outbound implementation verifier should match CommitteeVerifier address")

				// Check inbound implementation on CommitteeVerifierResolver
				versionTagReport, err := operations.ExecuteOperation(
					testsetup.BundleWithFreshReporter(e.OperationsBundle),
					committee_verifier.GetVersionTag,
					evmChain,
					contract.FunctionInput[any]{
						ChainSelector: evmChain.Selector,
						Address:       common.HexToAddress(input.CommitteeVerifier[0].Address),
					},
				)
				require.NoError(t, err, "ExecuteOperation should not error")
				inboundImpl, err := boundResolver.GetInboundImplementation(
					&bind.CallOpts{Context: t.Context()},
					versionTagReport.Output[:],
				)
				require.NoError(t, err, "GetInboundImplementation should not error")
				require.Equal(t, input.CommitteeVerifier[0].Address, inboundImpl.Hex(), "Inbound implementation verifier should match CommitteeVerifier address")
			}
		})
	}
}

func TestConfigureCommitteeVerifierForLanes_RevertWhen_InvalidSupportingContracts(t *testing.T) {
	tests := []struct {
		desc        string
		makeInput   func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) sequences.ConfigureCommitteeVerifierForLanesInput
		expectedErr string
	}{
		{
			desc: "no supporting contracts",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) sequences.ConfigureCommitteeVerifierForLanesInput {
				var routerAddress string
				var committeeVerifier string
				for _, addr := range chainReport.Output.Addresses {
					switch addr.Type {
					case datastore.ContractType(router.ContractType):
						routerAddress = addr.Address
					case datastore.ContractType(committee_verifier.ContractType):
						committeeVerifier = addr.Address
					}
				}
				remoteChainSelector := uint64(4356164186791070119)
				return sequences.ConfigureCommitteeVerifierForLanesInput{
					ChainSelector: chainReport.Input.ChainSelector,
					Router:        routerAddress,
					CommitteeVerifierConfig: adapters.CommitteeVerifierConfig[datastore.AddressRef]{
						CommitteeVerifier: []datastore.AddressRef{
							{
								Address: committeeVerifier,
								Type:    datastore.ContractType(committee_verifier.ContractType),
								Version: committee_verifier.Version,
							},
						},
						RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfig{
							remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
						},
					},
				}
			},
			expectedErr: "committee verifier resolver contract not found",
		},
		{
			desc: "wrong contract type",
			makeInput: func(chainReport operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]) sequences.ConfigureCommitteeVerifierForLanesInput {
				var routerAddress string
				var committeeVerifier string
				for _, addr := range chainReport.Output.Addresses {
					switch addr.Type {
					case datastore.ContractType(router.ContractType):
						routerAddress = addr.Address
					case datastore.ContractType(committee_verifier.ContractType):
						committeeVerifier = addr.Address
					}
				}
				remoteChainSelector := uint64(4356164186791070119)
				return sequences.ConfigureCommitteeVerifierForLanesInput{
					ChainSelector: chainReport.Input.ChainSelector,
					Router:        routerAddress,
					CommitteeVerifierConfig: adapters.CommitteeVerifierConfig[datastore.AddressRef]{
						CommitteeVerifier: []datastore.AddressRef{
							{
								Address: committeeVerifier,
								Type:    datastore.ContractType(committee_verifier.ContractType),
								Version: committee_verifier.Version,
							},
							{
								Address: routerAddress,
								Type:    datastore.ContractType(router.ContractType),
								Version: router.Version,
							},
						},
						RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfig{
							remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
						},
					},
				}
			},
			expectedErr: "committee verifier resolver contract not found",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")
			evmChain := e.BlockChains.EVMChains()[chainSelector]

			// Deploy chain contracts
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

			// Configure committee verifier for lanes
			input := test.makeInput(deploymentReport)
			_, err = operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.ConfigureCommitteeVerifierForLanes,
				e.BlockChains,
				input,
			)
			require.Error(t, err, "ExecuteSequence should error")
			require.Contains(t, err.Error(), test.expectedErr)
		})
	}
}
