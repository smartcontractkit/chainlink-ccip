package sequences_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
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
			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")

			bundle := operations.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				operations.NewMemoryReporter(),
			)

			chain, err := cldf_evm_provider.NewSimChainProvider(t, 5009297550715157269,
				cldf_evm_provider.SimChainProviderConfig{
					NumAdditionalAccounts: 1,
				},
			).Initialize(t.Context())
			require.NoError(t, err, "Failed to create SimChainProvider")

			chains := cldf_chain.NewBlockChainsFromSlice(
				[]cldf_chain.BlockChain{chain},
			)
			evmChain := chains.EVMChains()[5009297550715157269]

			usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10) // $15
			require.True(t, ok, "Failed to parse USDPerLINK")
			usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10) // $2000
			require.True(t, ok, "Failed to parse USDPerWETH")

			report, err := operations.ExecuteSequence(
				bundle,
				sequences.DeployChainContracts,
				evmChain,
				sequences.DeployChainContractsInput{
					ChainSelector:     5009297550715157269,
					ExistingAddresses: test.existingAddresses,
					ContractParams: sequences.ContractParams{
						RMNRemote:     sequences.RMNRemoteParams{},
						CCVAggregator: sequences.CCVAggregatorParams{},
						CommitOnRamp: sequences.CommitOnRampParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						CCVProxy: sequences.CCVProxyParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						FeeQuoter: sequences.FeeQuoterParams{
							MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
							TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
							LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
							WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
							USDPerLINK:                     usdPerLink,
							USDPerWETH:                     usdPerWeth,
						},
						CommitOffRamp: sequences.CommitOffRampParams{
							SignatureConfigArgs: commit_offramp.SignatureConfigArgs{{
								ConfigDigest: [32]byte{0x01},
								F:            1,
								Signers: []common.Address{
									common.HexToAddress("0x02"),
									common.HexToAddress("0x03"),
									common.HexToAddress("0x04"),
									common.HexToAddress("0x05"),
								},
							}},
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, report.Output.Addresses, 12, "Expected 12 addresses in output")
			require.Len(t, report.Output.Writes, 4, "Expected 4 writes in output")
			for _, write := range report.Output.Writes {
				// Contracts are deployed & still owned by deployer, so all writes should be executed
				require.True(t, write.Executed, "Expected all writes to be executed")
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
		lggr, err := logger.New()
		require.NoError(t, err, "Failed to create logger")

		bundle := operations.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			operations.NewMemoryReporter(),
		)

		// Create multiple chains
		chainSelectors := []uint64{
			5009297550715157269, // Chain 1
			4949039107694359620, // Chain 2
			6433500567565415381, // Chain 3
		}

		var allChains []cldf_chain.BlockChain
		for _, selector := range chainSelectors {
			chain, err := cldf_evm_provider.NewSimChainProvider(t, selector,
				cldf_evm_provider.SimChainProviderConfig{
					NumAdditionalAccounts: 1,
				},
			).Initialize(t.Context())
			require.NoError(t, err, "Failed to create SimChainProvider for chain %d", selector)
			allChains = append(allChains, chain)
		}

		chains := cldf_chain.NewBlockChainsFromSlice(allChains)
		evmChains := chains.EVMChains()

		// Deploy to each chain sequentially using the same bundle
		var allReports []operations.SequenceReport[sequences.DeployChainContractsInput, sequences.DeployChainContractsOutput]
		for _, selector := range chainSelectors {
			evmChain := evmChains[selector]

			usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10) // $15
			require.True(t, ok, "Failed to parse USDPerLINK")
			usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10) // $2000
			require.True(t, ok, "Failed to parse USDPerWETH")

			input := sequences.DeployChainContractsInput{
				ChainSelector:     selector,
				ExistingAddresses: nil,
				ContractParams: sequences.ContractParams{
					RMNRemote:     sequences.RMNRemoteParams{},
					CCVAggregator: sequences.CCVAggregatorParams{},
					CommitOnRamp: sequences.CommitOnRampParams{
						FeeAggregator: common.HexToAddress("0x01"),
					},
					CCVProxy: sequences.CCVProxyParams{
						FeeAggregator: common.HexToAddress("0x01"),
					},
					FeeQuoter: sequences.FeeQuoterParams{
						MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
						TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
						LINKPremiumMultiplierWeiPerEth: 9e17,       // 0.9 ETH
						WETHPremiumMultiplierWeiPerEth: 1e18,       // 1.0 ETH
						USDPerLINK:                     usdPerLink, // $15
						USDPerWETH:                     usdPerWeth, // $2000
					},
					CommitOffRamp: sequences.CommitOffRampParams{
						SignatureConfigArgs: commit_offramp.SignatureConfigArgs{{
							ConfigDigest: [32]byte{0x01},
							F:            1,
							Signers: []common.Address{
								common.HexToAddress("0x02"),
								common.HexToAddress("0x03"),
								common.HexToAddress("0x04"),
								common.HexToAddress("0x05"),
							},
						}},
					},
				},
			}

			report, err := operations.ExecuteSequence(bundle, sequences.DeployChainContracts, evmChain, input)
			require.NoError(t, err, "Failed to execute sequence for chain %d", selector)
			require.NotEmpty(t, report.Output.Addresses, "Expected operation reports for chain %d", selector)

			allReports = append(allReports, report)
		}

		// Verify all deployments succeeded
		require.Len(t, allReports, len(chainSelectors), "Expected reports for all chains")

		for i, report := range allReports {
			require.NotEmpty(t, report.Output.Addresses, "Expected addresses for chain %d", chainSelectors[i])
			require.Len(t, report.Output.Addresses, 12, "Expected 12 addresses deployed for chain %d", chainSelectors[i])
		}
	})

	t.Run("concurrent deployments", func(t *testing.T) {
		lggr, err := logger.New()
		require.NoError(t, err, "Failed to create logger")

		bundle := operations.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			operations.NewMemoryReporter(),
		)

		// Create multiple chains
		chainSelectors := []uint64{
			5009297550715157269, // Chain 1
			4949039107694359620, // Chain 2
			6433500567565415381, // Chain 3
		}

		var allChains []cldf_chain.BlockChain
		for _, selector := range chainSelectors {
			chain, err := cldf_evm_provider.NewSimChainProvider(t, selector,
				cldf_evm_provider.SimChainProviderConfig{
					NumAdditionalAccounts: 1,
				},
			).Initialize(t.Context())
			require.NoError(t, err, "Failed to create SimChainProvider for chain %d", selector)
			allChains = append(allChains, chain)
		}

		chains := cldf_chain.NewBlockChainsFromSlice(allChains)
		evmChains := chains.EVMChains()

		// Deploy to all chains concurrently using the same bundle
		type deployResult struct {
			chainSelector uint64
			report        operations.SequenceReport[sequences.DeployChainContractsInput, sequences.DeployChainContractsOutput]
			err           error
		}

		resultChan := make(chan deployResult, len(chainSelectors))

		usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10) // $15
		require.True(t, ok, "Failed to parse USDPerLINK")
		usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10) // $2000
		require.True(t, ok, "Failed to parse USDPerWETH")

		// Launch concurrent deployments
		for _, selector := range chainSelectors {
			go func(chainSel uint64) {
				evmChain := evmChains[chainSel]

				input := sequences.DeployChainContractsInput{
					ChainSelector:     chainSel,
					ExistingAddresses: nil,
					ContractParams: sequences.ContractParams{
						RMNRemote:     sequences.RMNRemoteParams{},
						CCVAggregator: sequences.CCVAggregatorParams{},
						CommitOnRamp: sequences.CommitOnRampParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						CCVProxy: sequences.CCVProxyParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						FeeQuoter: sequences.FeeQuoterParams{
							MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
							TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
							LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
							WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
							USDPerLINK:                     usdPerLink,
							USDPerWETH:                     usdPerWeth,
						},
						CommitOffRamp: sequences.CommitOffRampParams{
							SignatureConfigArgs: commit_offramp.SignatureConfigArgs{{
								ConfigDigest: [32]byte{0x01},
								F:            1,
								Signers: []common.Address{
									common.HexToAddress("0x02"),
									common.HexToAddress("0x03"),
									common.HexToAddress("0x04"),
									common.HexToAddress("0x05"),
								},
							}},
						},
					},
				}

				report, execErr := operations.ExecuteSequence(bundle, sequences.DeployChainContracts, evmChain, input)
				resultChan <- deployResult{chainSel, report, execErr}
			}(selector)
		}

		// Collect all results
		var results []deployResult
		for i := 0; i < len(chainSelectors); i++ {
			result := <-resultChan
			results = append(results, result)
		}

		// Verify all deployments succeeded
		require.Len(t, results, len(chainSelectors), "Expected results for all chains")

		for _, result := range results {
			require.NoError(t, result.err, "Failed to execute sequence for chain %d", result.chainSelector)
			require.NotEmpty(t, result.report.Output.Addresses, "Expected addresses for chain %d", result.chainSelector)
			require.Len(t, result.report.Output.Addresses, 12, "Expected 12 addresses deployed for chain %d", result.chainSelector)
		}
	})
}
