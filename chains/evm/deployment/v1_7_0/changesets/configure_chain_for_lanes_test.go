package changesets_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestConfigureChainForLanes_Apply(t *testing.T) {
	tests := []struct {
		desc string
	}{
		{
			desc: "valid input",
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

			chainSels := []uint64{5009297550715157269, 4356164186791070119}

			chainSlice := make([]cldf_chain.BlockChain, 0, len(chainSels))
			for _, chainSel := range chainSels {
				chain, err := cldf_evm_provider.NewSimChainProvider(t, chainSel,
					cldf_evm_provider.SimChainProviderConfig{
						NumAdditionalAccounts: 1,
					},
				).Initialize(t.Context())
				require.NoError(t, err, "Failed to create SimChainProvider")

				chainSlice = append(chainSlice, chain)
			}
			chains := cldf_chain.NewBlockChainsFromSlice(chainSlice)

			e := deployment.Environment{
				GetContext:       func() context.Context { return context.Background() },
				Logger:           lggr,
				OperationsBundle: bundle,
				BlockChains:      chains,
				DataStore:        datastore.NewMemoryDataStore().Seal(),
			}

			usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10) // $15
			require.True(t, ok, "Failed to parse USDPerLINK")
			usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10) // $2000
			require.True(t, ok, "Failed to parse USDPerWETH")

			// Deploy both chains
			runningDataStore := datastore.NewMemoryDataStore()
			for _, chainSel := range chainSels {
				out, err := changesets.DeployChainContracts.Apply(e, changesets.DeployChainContractsCfg{
					ChainSel: chainSel,
					Params: sequences.ContractParams{
						RMNRemote:     sequences.RMNRemoteParams{},
						CCVAggregator: sequences.CCVAggregatorParams{},
						CommitteeVerifier: sequences.CommitteeVerifierParams{
							FeeAggregator: common.HexToAddress("0x01"),
							SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
								Threshold: 1,
								Signers: []common.Address{
									common.HexToAddress("0x02"),
									common.HexToAddress("0x03"),
									common.HexToAddress("0x04"),
									common.HexToAddress("0x05"),
								},
							},
							StorageLocation: "https://test.chain.link.fake",
						},
						CCVProxy: sequences.CCVProxyParams{
							FeeAggregator: common.HexToAddress("0x01"),
						},
						ExecutorOnRamp: sequences.ExecutorOnRampParams{
							MaxCCVsPerMsg: 10,
						},
						FeeQuoter: sequences.FeeQuoterParams{
							MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
							TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
							LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
							WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
							USDPerLINK:                     usdPerLink,
							USDPerWETH:                     usdPerWeth,
						},
					},
				})
				require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
				err = runningDataStore.Merge(out.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployChainContracts")
			}
			e.DataStore = runningDataStore.Seal() // Override datastore in environment to include deployed contracts

			_, err = changesets.ConfigureChainForLanes.Apply(e, changesets.ConfigureChainForLanesCfg{
				ChainSel: 5009297550715157269,
				RemoteChains: map[uint64]changesets.RemoteChainConfig{
					4356164186791070119: {
						AllowTrafficFrom: true,
						CCIPMessageSource: datastore.AddressRef{
							Type:    datastore.ContractType(ccv_proxy.ContractType),
							Version: semver.MustParse("1.7.0"),
						},
						CCIPMessageDest: datastore.AddressRef{
							Type:    datastore.ContractType(ccv_aggregator.ContractType),
							Version: semver.MustParse("1.7.0"),
						},
						DefaultCCVOffRamps: []datastore.AddressRef{
							{Type: datastore.ContractType(committee_verifier.ContractType), Version: semver.MustParse("1.7.0")},
						},
						DefaultCCVOnRamps: []datastore.AddressRef{
							{Type: datastore.ContractType(committee_verifier.ContractType), Version: semver.MustParse("1.7.0")},
						},
						DefaultExecutor: datastore.AddressRef{
							Type:    datastore.ContractType(executor_onramp.ContractType),
							Version: semver.MustParse("1.7.0"),
						},
						CommitteeVerifierDestChainConfig: sequences.CommitteeVerifierDestChainConfig{
							AllowlistEnabled: false,
						},
						FeeQuoterDestChainConfig: fee_quoter_v2.DestChainConfig{
							IsEnabled:                         true,
							MaxNumberOfTokensPerMsg:           10,
							MaxDataBytes:                      30_000,
							MaxPerMsgGasLimit:                 3_000_000,
							DestGasOverhead:                   300_000,
							DefaultTokenFeeUSDCents:           25,
							DestGasPerPayloadByteBase:         16,
							DestGasPerPayloadByteHigh:         40,
							DestGasPerPayloadByteThreshold:    3000,
							DestDataAvailabilityOverheadGas:   100,
							DestGasPerDataAvailabilityByte:    16,
							DestDataAvailabilityMultiplierBps: 1,
							DefaultTokenDestGasOverhead:       90_000,
							DefaultTxGasLimit:                 200_000,
							GasMultiplierWeiPerEth:            11e17, // Gas multiplier in wei per eth is scaled by 1e18, so 11e17 is 1.1 = 110%
							NetworkFeeUSDCents:                10,
							ChainFamilySelector:               [4]byte{0x28, 0x12, 0xd5, 0x2c}, // EVM
						},
					},
				},
			})
			require.NoError(t, err, "Failed to apply ConfigureChainForLanes changeset")
		})
	}
}
