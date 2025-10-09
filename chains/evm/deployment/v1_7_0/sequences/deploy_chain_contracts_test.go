package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	mock_receiver "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/off_ramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
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
			e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
				chainSelector: {NumAdditionalAccounts: 1},
			})
			require.NoError(t, err, "Failed to create test environment")
			evmChain := e.BlockChains.EVMChains()[chainSelector]

			report, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.DeployChainContracts,
				evmChain,
				sequences.DeployChainContractsInput{
					ChainSelector:     chainSelector,
					ExistingAddresses: test.existingAddresses,
					ContractParams:    testsetup.CreateBasicContractParams(),
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, report.Output.BatchOps, 1, "Expected 1 empty batch operation")
			require.Len(t, report.Output.BatchOps[0].Transactions, 0, "Expected no transactions")

			exists := map[deployment.ContractType]bool{
				rmn_remote.ContractType:           false,
				router.ContractType:               false,
				executor_onramp.ContractType:      false,
				link.ContractType:                 false,
				weth.ContractType:                 false,
				committee_verifier.ContractType:   false,
				ccv_proxy.ContractType:            false,
				off_ramp.ContractType:             false,
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
		e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
			5009297550715157269: {NumAdditionalAccounts: 1},
			4949039107694359620: {NumAdditionalAccounts: 1},
			6433500567565415381: {NumAdditionalAccounts: 1},
		})
		require.NoError(t, err, "Failed to create test environment")
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
		e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
			5009297550715157269: {NumAdditionalAccounts: 1},
			4949039107694359620: {NumAdditionalAccounts: 1},
			6433500567565415381: {NumAdditionalAccounts: 1},
		})
		require.NoError(t, err, "Failed to create test environment")
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
