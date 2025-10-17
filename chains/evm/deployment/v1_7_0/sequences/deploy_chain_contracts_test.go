package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	mock_receiver "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	mock_recv_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_receiver_v2"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestDeployChainContracts_Idempotency(t *testing.T) {
	tests := []struct {
		desc              string
		existingAddresses []datastore.AddressRef
	}{
		{
			desc: "full deployment",
		},
		{
			desc: "partial deployment",
			existingAddresses: []datastore.AddressRef{
				{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(link.ContractType),
					Version:       semver.MustParse("1.0.0"),
					Address:       common.HexToAddress("0x01").Hex(),
				},
				{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(weth.ContractType),
					Version:       semver.MustParse("1.0.0"),
					Address:       common.HexToAddress("0x02").Hex(),
				},
			},
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

			report, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				e.BlockChains.EVMChains()[chainSelector],
				sequences.DeployChainContractsInput{
					ChainSelector:     chainSelector,
					ExistingAddresses: test.existingAddresses,
					ContractParams:    testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, report.Output.BatchOps, 2, "Expected 2 batch operations")
			require.Len(t, report.Output.BatchOps[0].Transactions, 0, "Expected no transactions in first batch operation")
			require.Len(t, report.Output.BatchOps[1].Transactions, 0, "Expected no transactions in second batch operation")

			exists := map[deployment.ContractType]bool{
				rmn_remote.ContractType:           false,
				router.ContractType:               false,
				executor.ContractType:             false,
				link.ContractType:                 false,
				weth.ContractType:                 false,
				committee_verifier.ContractType:   false,
				onramp.ContractType:               false,
				offramp.ContractType:              false,
				fee_quoter.ContractType:           false,
				committee_verifier.ProxyType:      false,
				rmn_proxy.ContractType:            false,
				token_admin_registry.ContractType: false,
				mock_receiver.ContractType:        false,
			}
			for _, addr := range report.Output.Addresses {
				exists[deployment.ContractType(addr.Type)] = true
			}
			for ctype, found := range exists {
				require.True(t, found, "Expected contract of type %s to be deployed", ctype)
			}

			for _, existing := range test.existingAddresses {
				found := false
				for _, addr := range report.Output.Addresses {
					if addr.Type == existing.Type {
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

func TestDeployChainContracts_MultipleDeployments(t *testing.T) {
	t.Run("sequential deployments", func(t *testing.T) {
		e, err := environment.New(t.Context(),
			environment.WithEVMSimulated(t, []uint64{5009297550715157269, 4949039107694359620, 6433500567565415381}),
		)
		require.NoError(t, err, "Failed to create environment")
		require.NotNil(t, e, "Environment should be created")
		evmChains := e.BlockChains.EVMChains()

		// Deploy to each chain sequentially using the same bundle
		var allReports []operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]
		for _, evmChain := range evmChains {
			input := sequences.DeployChainContractsInput{
				ChainSelector:     evmChain.Selector,
				ExistingAddresses: nil,
				ContractParams:    testsetup.CreateBasicContractParams(),
			}

			report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployChainContracts, evmChain, input)
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

		// Deploy to all chains concurrently using the same bundle
		type deployResult struct {
			chainSelector uint64
			report        operations.SequenceReport[sequences.DeployChainContractsInput, seq_core.OnChainOutput]
			err           error
		}

		resultChan := make(chan deployResult, len(evmChains))

		// Launch concurrent deployments
		for _, evmChain := range evmChains {
			go func(chainSel uint64) {
				evmChain := evmChains[chainSel]

				input := sequences.DeployChainContractsInput{
					ChainSelector:     chainSel,
					ExistingAddresses: nil,
					ContractParams:    testsetup.CreateBasicContractParams(),
				}

				report, execErr := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployChainContracts, evmChain, input)
				resultChan <- deployResult{chainSel, report, execErr}
			}(evmChain.Selector)
		}

		// Collect all results
		var results []deployResult
		for i := 0; i < len(evmChains); i++ {
			result := <-resultChan
			results = append(results, result)
		}

		// Verify all deployments succeeded
		require.Len(t, results, len(evmChains), "Expected results for all chains")

		for _, result := range results {
			require.NoError(t, result.err, "Failed to execute sequence for chain %d", result.chainSelector)
			require.NotEmpty(t, result.report.Output.Addresses, "Expected addresses for chain %d", result.chainSelector)
		}
	})
}

func TestDeployChainContracts_MultipleCommitteeVerifiersAndMultipleMockReceiverConfigs(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	// Configure two committee verifiers with different qualifiers and request both for MockReceiver
	params := testsetup.CreateBasicContractParams()
	params.CommitteeVerifier = []sequences.CommitteeVerifierParams{
		{
			Version:       semver.MustParse("1.7.0"),
			FeeAggregator: common.HexToAddress("0x01"),
			SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
				Threshold: 1,
				Signers: []common.Address{
					common.HexToAddress("0x02"),
					common.HexToAddress("0x03"),
				},
			},
			StorageLocation: "https://test.chain.link.fake",
			Qualifier:       "alpha",
		},
		{
			Version:       semver.MustParse("1.7.0"),
			FeeAggregator: common.HexToAddress("0x01"),
			SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
				Threshold: 1,
				Signers: []common.Address{
					common.HexToAddress("0x02"),
					common.HexToAddress("0x03"),
				},
			},
			StorageLocation: "https://test.chain.link.fake",
			Qualifier:       "beta",
		},
	}
	params.MockReceivers = []sequences.MockReceiverParams{
		{
			Version: semver.MustParse("1.7.0"),
			RequiredVerifiers: []datastore.AddressRef{
				{
					ChainSelector: chainSelector,
					Type:          datastore.ContractType(committee_verifier.ContractType),
					Version:       semver.MustParse("1.7.0"),
					Qualifier:     "alpha",
				},
				{
					ChainSelector: chainSelector,
					Type:          datastore.ContractType(committee_verifier.ContractType),
					Version:       semver.MustParse("1.7.0"),
					Qualifier:     "beta",
				},
			},
			Qualifier: "q1",
		},
		{
			Version: semver.MustParse("1.7.0"),
			RequiredVerifiers: []datastore.AddressRef{
				{
					Type:      datastore.ContractType(committee_verifier.ContractType),
					Version:   semver.MustParse("1.7.0"),
					Qualifier: "alpha",
				},
			},
			OptionalVerifiers: []datastore.AddressRef{
				{
					Type:      datastore.ContractType(committee_verifier.ContractType),
					Version:   semver.MustParse("1.7.0"),
					Qualifier: "beta",
				},
			},
			OptionalThreshold: 1,
			Qualifier:         "q2",
		},
	}

	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployChainContractsInput{
			ChainSelector:     chainSelector,
			ExistingAddresses: nil,
			ContractParams:    params,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	// Assert mock receiver properties
	ds := datastore.NewMemoryDataStore()
	for _, addr := range report.Output.Addresses {
		require.NoError(t, ds.Addresses().Add(addr))
	}
	sealed := ds.Seal()

	q1ReceiverRef, err := datastore_utils.FindAndFormatRef(sealed, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(mock_receiver.ContractType),
		Version:       semver.MustParse("1.7.0"),
		Qualifier:     "q1",
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	q2ReceiverRef, err := datastore_utils.FindAndFormatRef(sealed, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(mock_receiver.ContractType),
		Version:       semver.MustParse("1.7.0"),
		Qualifier:     "q2",
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	q1Receiver, err := mock_recv_bindings.NewMockReceiverV2(q1ReceiverRef, e.BlockChains.EVMChains()[chainSelector].Client)
	require.NoError(t, err)

	required, optional, threshold, err := q1Receiver.GetCCVs(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, chainSelector)
	require.NoError(t, err)
	require.Len(t, required, 2)
	require.Len(t, optional, 0)
	require.Equal(t, uint8(0), threshold)

	q2Receiver, err := mock_recv_bindings.NewMockReceiverV2(q2ReceiverRef, e.BlockChains.EVMChains()[chainSelector].Client)
	require.NoError(t, err)

	required, optional, threshold, err = q2Receiver.GetCCVs(&bind.CallOpts{Context: e.OperationsBundle.GetContext()}, chainSelector)
	require.NoError(t, err)
	require.Len(t, required, 1)
	require.Len(t, optional, 1)
	require.Equal(t, uint8(1), threshold)
}
