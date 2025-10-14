package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func basicParams() sequences.DeployCommitteeVerifierParams {
	return sequences.DeployCommitteeVerifierParams{
		CommitteeVerifierVersion:      semver.MustParse("1.7.0"),
		CommitteeVerifierProxyVersion: semver.MustParse("1.7.0"),
		Args: committee_verifier.ConstructorArgs{
			DynamicConfig: committee_verifier.DynamicConfig{
				FeeQuoter:      common.HexToAddress("0x01"),
				FeeAggregator:  common.HexToAddress("0x02"),
				AllowlistAdmin: common.HexToAddress("0x03"),
			},
			StorageLocation: "https://test.chain.link.fake",
		},
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

			params := basicParams()
			params.Qualifier = "alpha"

			report, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployCommitteeVerifier,
				e.BlockChains.EVMChains()[chainSelector],
				sequences.DeployCommitteeVerifierInput{
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
				deployment.ContractType(committee_verifier.ContractType): false,
				deployment.ContractType(committee_verifier.ProxyType):    false,
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

	params := basicParams()
	params.Qualifier = "alpha"

	// Pre-deploy a real CommitteeVerifier with qualifier "alpha"
	q := params.Qualifier
	deployReport, err := operations.ExecuteOperation(
		e.OperationsBundle,
		committee_verifier.Deploy,
		e.BlockChains.EVMChains()[chainSelector],
		evm_contract.DeployInput[committee_verifier.ConstructorArgs]{
			ChainSelector:  chainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ContractType, *params.CommitteeVerifierVersion),
			Qualifier:      &q,
			Args:           params.Args,
		},
	)
	require.NoError(t, err, "Failed to pre-deploy CommitteeVerifier")

	// Now run the sequence with the existing deployed address
	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployCommitteeVerifierInput{
			ChainSelector:     chainSelector,
			ExistingAddresses: []datastore.AddressRef{deployReport.Output},
			Params:            params,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error with pre-deployed address")

	// Expect both contract types to be present
	exists := map[deployment.ContractType]bool{
		deployment.ContractType(committee_verifier.ContractType): false,
		deployment.ContractType(committee_verifier.ProxyType):    false,
	}
	for _, addr := range report.Output.Addresses {
		exists[deployment.ContractType(addr.Type)] = true
	}
	for ctype, found := range exists {
		require.True(t, found, "Expected contract of type %s to be present", ctype)
	}

	// Ensure the sequence reused the pre-deployed CommitteeVerifier address for qualifier "alpha"
	var reused bool
	for _, addr := range report.Output.Addresses {
		if addr.Type == datastore.ContractType(committee_verifier.ContractType) && addr.Qualifier == params.Qualifier {
			require.Equal(t, deployReport.Output.Address, addr.Address, "Expected existing CommitteeVerifier address to be reused")
			reused = true
			break
		}
	}
	require.True(t, reused, "Expected to find reused CommitteeVerifier address")

	// Since SetSignatureConfigs executes immediately, expect a single empty batch operation
	require.Len(t, report.Output.BatchOps, 1, "Expected 1 empty batch operation")
	require.Len(t, report.Output.BatchOps[0].Transactions, 0, "Expected no transactions in batch op")
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
			params := basicParams()
			params.Qualifier = "alpha"
			input := sequences.DeployCommitteeVerifierInput{
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

				params := basicParams()
				params.Qualifier = "alpha"
				input := sequences.DeployCommitteeVerifierInput{
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
	paramsAlpha := basicParams()
	paramsAlpha.Qualifier = "alpha"
	report1, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSel],
		sequences.DeployCommitteeVerifierInput{
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
	alphaProxy, ok := find(addrs1, datastore.ContractType(committee_verifier.ProxyType), "alpha")
	require.True(t, ok)

	// Second run with qualifier "beta", passing previous addresses as existing
	paramsBeta := basicParams()
	paramsBeta.Qualifier = "beta"
	report2, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSel],
		sequences.DeployCommitteeVerifierInput{
			ChainSelector:     chainSel,
			ExistingAddresses: addrs1,
			Params:            paramsBeta,
		},
	)
	require.NoError(t, err)
	addrs2 := report2.Output.Addresses

	betaCV, ok := find(addrs2, datastore.ContractType(committee_verifier.ContractType), "beta")
	require.True(t, ok)
	betaProxy, ok := find(addrs2, datastore.ContractType(committee_verifier.ProxyType), "beta")
	require.True(t, ok)

	require.NotEqual(t, alphaCV.Address, betaCV.Address, "expected different addresses for different qualifiers")
	require.NotEqual(t, alphaProxy.Address, betaProxy.Address, "expected different addresses for different qualifiers")

	// Third run reusing qualifier "alpha" should return the same alpha addresses
	report3, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployCommitteeVerifier,
		e.BlockChains.EVMChains()[chainSel],
		sequences.DeployCommitteeVerifierInput{
			ChainSelector:     chainSel,
			ExistingAddresses: append(addrs1, addrs2...),
			Params:            paramsAlpha,
		},
	)
	require.NoError(t, err)
	addrs3 := report3.Output.Addresses

	reAlphaCV, ok := find(addrs3, datastore.ContractType(committee_verifier.ContractType), "alpha")
	require.True(t, ok)
	reAlphaProxy, ok := find(addrs3, datastore.ContractType(committee_verifier.ProxyType), "alpha")
	require.True(t, ok)

	require.Equal(t, alphaCV.Address, reAlphaCV.Address, "expected same address when reusing qualifier")
	require.Equal(t, alphaProxy.Address, reAlphaProxy.Address, "expected same address when reusing qualifier")
}
