package sequences_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/deployment"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestDeployChainContracts(t *testing.T) {
	tests := []struct {
		desc              string
		existingAddresses []deployment.AddressRef
	}{
		{
			desc: "full deployment",
		},
		{
			desc: "partial deployment",
			existingAddresses: []deployment.AddressRef{
				{
					ChainSelector: 5009297550715157269,
					Type:          link.ContractType,
					Version:       semver.MustParse("1.0.0").String(),
					Address:       common.HexToAddress("0x01").Hex(),
				},
				{
					ChainSelector: 5009297550715157269,
					Type:          weth.ContractType,
					Version:       semver.MustParse("1.0.0").String(),
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

			report, err := operations.ExecuteSequence(
				bundle,
				sequences.DeployChainContracts,
				evmChain,
				sequences.DeployChainContractsInput{
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
						},
						CommitOffRamp: sequences.CommitOffRampParams{
							SignatureConfigArgs: commit_offramp.SignatureConfigArgs{
								ConfigDigest: [32]byte{},
								F:            1,
								Signers: []common.Address{
									common.HexToAddress("0x02"),
									common.HexToAddress("0x03"),
									common.HexToAddress("0x04"),
									common.HexToAddress("0x05"),
								},
							},
						},
					},
				},
			)
			require.NoError(t, err, "ExecuteSequence should not error")
			require.Len(t, report.Output.Addresses, 12, "Expected 12 addresses in output")
			require.Len(t, report.Output.Writes, 3, "Expected 3 writes in output")
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
