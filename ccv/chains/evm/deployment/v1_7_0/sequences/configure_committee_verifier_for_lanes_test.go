package sequences_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
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
					CommitteeVerifierConfigWithSignatureConfigPerRemoteChain: adapters.CommitteeVerifierConfigWithSignatureConfigPerRemoteChain[string]{
						CommitteeVerifier: []string{committeeVerifier, committeeVerifierResolver},
						RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfigWithSignatureConfig{
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
					CommitteeVerifierConfigWithSignatureConfigPerRemoteChain: adapters.CommitteeVerifierConfigWithSignatureConfigPerRemoteChain[string]{
						CommitteeVerifier: []string{committeeVerifier, committeeVerifierResolver},
						RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfigWithSignatureConfig{
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
					CommitteeVerifierConfigWithSignatureConfigPerRemoteChain: adapters.CommitteeVerifierConfigWithSignatureConfigPerRemoteChain[string]{
						CommitteeVerifier: []string{committeeVerifier, committeeVerifierResolver},
						RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfigWithSignatureConfig{
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
			chainSelector := uint64(5009297550715157269)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")
			evmChain := e.BlockChains.EVMChains()[chainSelector]

			// Deploy chain contracts
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
						Address:       common.HexToAddress(input.CommitteeVerifier[0]),
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
						Address:       common.HexToAddress(input.CommitteeVerifier[0]),
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
					common.HexToAddress(input.CommitteeVerifier[1]),
					evmChain.Client,
				)
				require.NoError(t, err, "Failed to instantiate VersionedVerifierResolver")
				outboundImpl, err := boundResolver.GetOutboundImplementation(
					&bind.CallOpts{Context: t.Context()},
					remoteSelector,
					[]byte{},
				)
				require.NoError(t, err, "GetOutboundImplementation should not error")
				require.Equal(t, input.CommitteeVerifier[0], outboundImpl.Hex(), "Outbound implementation verifier should match CommitteeVerifier address")

				// Check inbound implementation on CommitteeVerifierResolver
				versionTagReport, err := operations.ExecuteOperation(
					testsetup.BundleWithFreshReporter(e.OperationsBundle),
					committee_verifier.GetVersionTag,
					evmChain,
					contract.FunctionInput[any]{
						ChainSelector: evmChain.Selector,
						Address:       common.HexToAddress(input.CommitteeVerifier[0]),
					},
				)
				require.NoError(t, err, "ExecuteOperation should not error")
				inboundImpl, err := boundResolver.GetInboundImplementation(
					&bind.CallOpts{Context: t.Context()},
					versionTagReport.Output[:],
				)
				require.NoError(t, err, "GetInboundImplementation should not error")
				require.Equal(t, input.CommitteeVerifier[0], inboundImpl.Hex(), "Inbound implementation verifier should match CommitteeVerifier address")
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
					CommitteeVerifierConfigWithSignatureConfigPerRemoteChain: adapters.CommitteeVerifierConfigWithSignatureConfigPerRemoteChain[string]{
						CommitteeVerifier: []string{committeeVerifier},
						RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfigWithSignatureConfig{
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
					CommitteeVerifierConfigWithSignatureConfigPerRemoteChain: adapters.CommitteeVerifierConfigWithSignatureConfigPerRemoteChain[string]{
						CommitteeVerifier: []string{committeeVerifier, routerAddress},
						RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfigWithSignatureConfig{
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
			chainSelector := uint64(5009297550715157269)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create environment")
			require.NotNil(t, e, "Environment should be created")
			evmChain := e.BlockChains.EVMChains()[chainSelector]

			// Deploy chain contracts
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

func TestConfigureCommitteeVerifierForLanes_RemoteChainsToDisconnect(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	remoteChainSelector1 := uint64(4356164186791070119)
	remoteChainSelector2 := uint64(4949039107694359620)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	evmChain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy chain contracts
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
	var committeeVerifier string
	var committeeVerifierResolver string
	for _, addr := range deploymentReport.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(router.ContractType):
			routerAddress = addr.Address
		case datastore.ContractType(committee_verifier.ContractType):
			committeeVerifier = addr.Address
		case datastore.ContractType(committee_verifier.ResolverType):
			committeeVerifierResolver = addr.Address
		}
	}

	// First, configure both remote chains
	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.ConfigureCommitteeVerifierForLanes,
		e.BlockChains,
		sequences.ConfigureCommitteeVerifierForLanesInput{
			ChainSelector: chainSelector,
			Router:        routerAddress,
			CommitteeVerifierConfigWithSignatureConfigPerRemoteChain: adapters.CommitteeVerifierConfigWithSignatureConfigPerRemoteChain[string]{
				CommitteeVerifier: []string{committeeVerifier, committeeVerifierResolver},
				RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfigWithSignatureConfig{
					remoteChainSelector1: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
					remoteChainSelector2: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
				},
			},
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	// Verify both chains have signature configs
	signatureConfig1, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		committee_verifier.GetSignatureConfig,
		evmChain,
		contract.FunctionInput[uint64]{
			ChainSelector: evmChain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
			Args:          remoteChainSelector1,
		},
	)
	require.NoError(t, err, "ExecuteOperation should not error")
	require.Equal(t, remoteChainSelector1, signatureConfig1.Output.SourceChainSelector, "SourceChainSelector should match for chain 1")

	signatureConfig2, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		committee_verifier.GetSignatureConfig,
		evmChain,
		contract.FunctionInput[uint64]{
			ChainSelector: evmChain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
			Args:          remoteChainSelector2,
		},
	)
	require.NoError(t, err, "ExecuteOperation should not error")
	require.Equal(t, remoteChainSelector2, signatureConfig2.Output.SourceChainSelector, "SourceChainSelector should match for chain 2")

	// Now disconnect remoteChainSelector1
	_, err = operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.ConfigureCommitteeVerifierForLanes,
		e.BlockChains,
		sequences.ConfigureCommitteeVerifierForLanesInput{
			ChainSelector:            chainSelector,
			Router:                   routerAddress,
			RemoteChainsToDisconnect: []uint64{remoteChainSelector1},
			CommitteeVerifierConfigWithSignatureConfigPerRemoteChain: adapters.CommitteeVerifierConfigWithSignatureConfigPerRemoteChain[string]{
				CommitteeVerifier: []string{committeeVerifier, committeeVerifierResolver},
				RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfigWithSignatureConfig{
					remoteChainSelector2: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
				},
			},
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	// Verify remoteChainSelector1 signature config is removed
	// Attempting to get signature config for disconnected chain should fail or return empty
	signatureConfig1After, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		committee_verifier.GetSignatureConfig,
		evmChain,
		contract.FunctionInput[uint64]{
			ChainSelector: evmChain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
			Args:          remoteChainSelector1,
		},
	)
	// The signature config should be removed, so this should either error or return zero values
	// Based on the contract implementation, it may revert or return zero values
	if err == nil {
		// If no error, verify it's been removed (zero threshold or empty signers)
		require.True(t, signatureConfig1After.Output.Threshold == 0 || len(signatureConfig1After.Output.Signers) == 0,
			"Signature config should be removed for disconnected chain")
	}

	// Verify remoteChainSelector2 signature config still exists
	signatureConfig2After, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		committee_verifier.GetSignatureConfig,
		evmChain,
		contract.FunctionInput[uint64]{
			ChainSelector: evmChain.Selector,
			Address:       common.HexToAddress(committeeVerifier),
			Args:          remoteChainSelector2,
		},
	)
	require.NoError(t, err, "ExecuteOperation should not error")
	require.Equal(t, remoteChainSelector2, signatureConfig2After.Output.SourceChainSelector, "SourceChainSelector should still match for non-disconnected chain")
	require.NotZero(t, signatureConfig2After.Output.Threshold, "Threshold should not be zero for non-disconnected chain")
	require.NotEmpty(t, signatureConfig2After.Output.Signers, "Signers should not be empty for non-disconnected chain")
}
