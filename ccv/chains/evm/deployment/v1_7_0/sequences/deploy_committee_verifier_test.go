package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func basicParams() sequences.CommitteeVerifierParams {
	return sequences.CommitteeVerifierParams{
		Version:         semver.MustParse("1.7.0"),
		FeeAggregator:   common.HexToAddress("0x02"),
		AllowlistAdmin:  common.HexToAddress("0x03"),
		StorageLocation: "https://test.chain.link.fake",
		SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
			Threshold: 1,
			Signers: []common.Address{
				common.HexToAddress("0x02"),
				common.HexToAddress("0x03"),
				common.HexToAddress("0x04"),
				common.HexToAddress("0x05"),
			},
		},
	}
}

func TestDeployCommitteeVerifier_Idempotency(t *testing.T) {
	tests := []struct {
		desc              string
		existingAddresses []datastore.AddressRef
	}{
		{
			desc: "full deployment",
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
			e.DataStore = datastore.NewMemoryDataStore().Seal()

			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSelector], contract.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  chainSelector,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{e.BlockChains.EVMChains()[chainSelector].DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")

			params := basicParams()
			params.Qualifier = "alpha"

			report, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployCommitteeVerifier,
				e.BlockChains.EVMChains()[chainSelector],
				sequences.DeployCommitteeVerifierInput{
					CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
					ChainSelector:     chainSelector,
					ExistingAddresses: test.existingAddresses,
					Params:            params,
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, report.Output.BatchOps, 1, "Expected 1 empty batch operation")
			require.Len(t, report.Output.BatchOps[0].Transactions, 0, "Expected no transactions")

			// Expect both contract types to be present
			exists := map[deployment.ContractType]bool{
				deployment.ContractType(committee_verifier.ContractType):      false,
				deployment.ContractType(committee_verifier.ResolverType):      false,
				deployment.ContractType(committee_verifier.ResolverProxyType): false,
			}
			for _, addr := range report.Output.Addresses {
				exists[deployment.ContractType(addr.Type)] = true
			}
			for ctype, found := range exists {
				require.True(t, found, "Expected contract of type %s to be deployed", ctype)
			}

			// Verify reuse of existing addresses when provided
			for _, existing := range test.existingAddresses {
				found := false
				for _, addr := range report.Output.Addresses {
					if addr.Type == existing.Type && addr.Qualifier == existing.Qualifier {
						require.Equal(t, existing.Address, addr.Address, "Expected existing address to be reused for %s", existing.Type)
						found = true
						break
					}
				}
				require.True(t, found, "Expected to find existing address for %s", existing.Type)
			}
		})
	}
}

func TestDeployCommitteeVerifier_Idempotency_WithPredeployedCommitteeVerifier(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSelector], contract.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
		ChainSelector:  chainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{e.BlockChains.EVMChains()[chainSelector].DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory")

	params := basicParams()
	params.Qualifier = "alpha"

	// First, deploy the full CommitteeVerifier stack
	deployReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployCommitteeVerifierInput{
			CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
			ChainSelector:   chainSelector,
			Params:          params,
		},
	)
	require.NoError(t, err, "Failed to pre-deploy CommitteeVerifier stack")

	// Now run the sequence with the existing deployed address
	redundantReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployCommitteeVerifierInput{
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ChainSelector:     chainSelector,
			ExistingAddresses: deployReport.Output.Addresses,
			Params:            params,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error with pre-deployed address")

	// Expect all contract types to be present
	exists := map[deployment.ContractType]bool{
		deployment.ContractType(committee_verifier.ContractType):      false,
		deployment.ContractType(committee_verifier.ResolverType):      false,
		deployment.ContractType(committee_verifier.ResolverProxyType): false,
	}
	for _, addr := range redundantReport.Output.Addresses {
		exists[deployment.ContractType(addr.Type)] = true
	}
	for ctype, found := range exists {
		require.True(t, found, "Expected contract of type %s to be present", ctype)
	}

	// Ensure that all addresses in the deployReport exist in the redundantReport
	for _, firstAddr := range deployReport.Output.Addresses {
		found := false
		for _, redundantAddr := range redundantReport.Output.Addresses {
			if redundantAddr.Type == firstAddr.Type &&
				redundantAddr.Qualifier == firstAddr.Qualifier &&
				redundantAddr.Address == firstAddr.Address &&
				redundantAddr.Version.String() == firstAddr.Version.String() {
				found = true
				break
			}
		}
		require.True(t, found, "Expected to find address %s with type %s, qualifier %s, version %s in redundantReport", firstAddr.Address, firstAddr.Type, firstAddr.Qualifier, firstAddr.Version.String())
	}

	// Since SetSignatureConfigs executes immediately, expect a single empty batch operation
	require.Len(t, redundantReport.Output.BatchOps, 1, "Expected 1 empty batch operation")
	require.Len(t, redundantReport.Output.BatchOps[0].Transactions, 0, "Expected no transactions in batch op")
}

func TestDeployCommitteeVerifier_MultipleDeployments(t *testing.T) {
	t.Run("sequential deployments", func(t *testing.T) {
		e, err := environment.New(t.Context(),
			environment.WithEVMSimulated(t, []uint64{5009297550715157269, 4949039107694359620, 6433500567565415381}),
		)
		require.NoError(t, err, "Failed to create environment")
		require.NotNil(t, e, "Environment should be created")
		evmChains := e.BlockChains.EVMChains()

		var allReports []operations.SequenceReport[sequences.DeployCommitteeVerifierInput, seq_core.OnChainOutput]
		for _, evmChain := range evmChains {
			create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract.DeployInput[create2_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
				ChainSelector:  evmChain.Selector,
				Args: create2_factory.ConstructorArgs{
					AllowList: []common.Address{evmChain.DeployerKey.From},
				},
			}, nil)
			require.NoError(t, err, "Failed to deploy CREATE2Factory")

			params := basicParams()
			params.Qualifier = "alpha"
			input := sequences.DeployCommitteeVerifierInput{
				CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
				ChainSelector:     evmChain.Selector,
				ExistingAddresses: nil,
				Params:            params,
			}

			report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployCommitteeVerifier, evmChain, input)
			require.NoError(t, err, "Failed to execute sequence for chain %d", evmChain.Selector)
			require.NotEmpty(t, report.Output.Addresses, "Expected operation reports for chain %d", evmChain.Selector)

			allReports = append(allReports, report)
		}

		// Verify all deployments succeeded
		require.Len(t, allReports, len(evmChains), "Expected reports for all chains")

		for _, report := range allReports {
			require.NotEmpty(t, report.Output.Addresses, "Expected addresses")
		}
	})

	t.Run("concurrent deployments", func(t *testing.T) {
		e, err := environment.New(t.Context(),
			environment.WithEVMSimulated(t, []uint64{5009297550715157269, 4949039107694359620, 6433500567565415381}),
		)
		require.NoError(t, err, "Failed to create environment")
		require.NotNil(t, e, "Environment should be created")
		evmChains := e.BlockChains.EVMChains()

		type deployResult struct {
			chainSelector uint64
			report        operations.SequenceReport[sequences.DeployCommitteeVerifierInput, seq_core.OnChainOutput]
			err           error
		}

		resultChan := make(chan deployResult, len(evmChains))

		for _, evmChain := range evmChains {
			go func(chainSel uint64) {
				evmChain := evmChains[chainSel]

				create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, evmChain, contract.DeployInput[create2_factory.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
					ChainSelector:  evmChain.Selector,
					Args: create2_factory.ConstructorArgs{
						AllowList: []common.Address{evmChain.DeployerKey.From},
					},
				}, nil)
				require.NoError(t, err, "Failed to deploy CREATE2Factory")

				params := basicParams()
				params.Qualifier = "alpha"
				input := sequences.DeployCommitteeVerifierInput{
					CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
					ChainSelector:     chainSel,
					ExistingAddresses: nil,
					Params:            params,
				}

				report, execErr := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployCommitteeVerifier, evmChain, input)
				resultChan <- deployResult{chainSel, report, execErr}
			}(evmChain.Selector)
		}

		var results []deployResult
		for i := 0; i < len(evmChains); i++ {
			result := <-resultChan
			results = append(results, result)
		}

		require.Len(t, results, len(evmChains), "Expected results for all chains")

		for _, result := range results {
			require.NoError(t, result.err, "Failed to execute sequence for chain %d", result.chainSelector)
			require.NotEmpty(t, result.report.Output.Addresses, "Expected addresses for chain %d", result.chainSelector)
		}
	})
}

func TestDeployCommitteeVerifier_MultipleQualifiersOnSameChain(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{5009297550715157269}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	chainSel := uint64(5009297550715157269)

	// First run with qualifier "alpha"
	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
		ChainSelector:  chainSel,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory")

	paramsAlpha := basicParams()
	paramsAlpha.Qualifier = "alpha"
	report1, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSel],
		sequences.DeployCommitteeVerifierInput{
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ChainSelector:     chainSel,
			ExistingAddresses: nil,
			Params:            paramsAlpha,
		},
	)
	require.NoError(t, err)
	addrs1 := report1.Output.Addresses

	find := func(refs []datastore.AddressRef, ct datastore.ContractType, q string) (datastore.AddressRef, bool) {
		for _, r := range refs {
			if r.Type == ct && r.Qualifier == q {
				return r, true
			}
		}
		return datastore.AddressRef{}, false
	}

	alphaCV, ok := find(addrs1, datastore.ContractType(committee_verifier.ContractType), "alpha")
	require.True(t, ok)
	alphaResolver, ok := find(addrs1, datastore.ContractType(committee_verifier.ResolverType), "alpha")
	require.True(t, ok)
	alphaResolverProxy, ok := find(addrs1, datastore.ContractType(committee_verifier.ResolverProxyType), "alpha")
	require.True(t, ok)

	// Second run with qualifier "beta", passing previous addresses as existing
	paramsBeta := basicParams()
	paramsBeta.Qualifier = "beta"
	report2, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSel],
		sequences.DeployCommitteeVerifierInput{
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ChainSelector:     chainSel,
			ExistingAddresses: addrs1,
			Params:            paramsBeta,
		},
	)
	require.NoError(t, err)
	addrs2 := report2.Output.Addresses

	betaCV, ok := find(addrs2, datastore.ContractType(committee_verifier.ContractType), "beta")
	require.True(t, ok)
	betaResolver, ok := find(addrs2, datastore.ContractType(committee_verifier.ResolverType), "beta")
	require.True(t, ok)
	betaResolverProxy, ok := find(addrs2, datastore.ContractType(committee_verifier.ResolverProxyType), "beta")
	require.True(t, ok)

	require.NotEqual(t, alphaCV.Address, betaCV.Address, "expected different addresses for different qualifiers")
	require.NotEqual(t, alphaResolver.Address, betaResolver.Address, "expected different addresses for different qualifiers")
	require.NotEqual(t, alphaResolverProxy.Address, betaResolverProxy.Address, "expected different addresses for different qualifiers")

	// Third run reusing qualifier "alpha" should return the same alpha addresses
	report3, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSel],
		sequences.DeployCommitteeVerifierInput{
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ChainSelector:     chainSel,
			ExistingAddresses: append(addrs1, addrs2...),
			Params:            paramsAlpha,
		},
	)
	require.NoError(t, err)
	addrs3 := report3.Output.Addresses

	reAlphaCV, ok := find(addrs3, datastore.ContractType(committee_verifier.ContractType), "alpha")
	require.True(t, ok)
	reAlphaResolver, ok := find(addrs3, datastore.ContractType(committee_verifier.ResolverType), "alpha")
	require.True(t, ok)
	reAlphaResolverProxy, ok := find(addrs3, datastore.ContractType(committee_verifier.ResolverProxyType), "alpha")
	require.True(t, ok)

	require.Equal(t, alphaCV.Address, reAlphaCV.Address, "expected same address when reusing qualifier")
	require.Equal(t, alphaResolver.Address, reAlphaResolver.Address, "expected same address when reusing qualifier")
	require.Equal(t, alphaResolverProxy.Address, reAlphaResolverProxy.Address, "expected same address when reusing qualifier")
}
